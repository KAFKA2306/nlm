package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

var itemRE = regexp.MustCompile(`(?s)<item>.*?</item>`)
var pubDateRE = regexp.MustCompile(`(?s)<pubDate>\s*(.*?)\s*</pubDate>`)
var guidRE = regexp.MustCompile(`(?s)<guid(?:\s[^>]*)?>\s*(.*?)\s*</guid>`)
var enclosureRE = regexp.MustCompile(`<enclosure\s+[^>]*url="([^"]+)"`)

type feedItem struct {
	Start    int
	End      int
	Raw      string
	Date     time.Time
	GUID     string
	AudioKey string
}

func parseFeedItems(feed []byte) []feedItem {
	locs := itemRE.FindAllIndex(feed, -1)
	items := make([]feedItem, 0, len(locs))
	for _, loc := range locs {
		raw := string(feed[loc[0]:loc[1]])
		item := feedItem{Start: loc[0], End: loc[1], Raw: raw}
		if match := pubDateRE.FindStringSubmatch(raw); len(match) == 2 {
			item.Date, _ = time.Parse(time.RFC1123Z, strings.TrimSpace(html.UnescapeString(match[1])))
		}
		if match := guidRE.FindStringSubmatch(raw); len(match) == 2 {
			item.GUID = strings.TrimSpace(html.UnescapeString(match[1]))
		}
		if match := enclosureRE.FindStringSubmatch(raw); len(match) == 2 {
			if parsed, err := url.Parse(html.UnescapeString(match[1])); err == nil {
				base := filepath.Base(parsed.Path)
				item.AudioKey = strings.TrimSuffix(base, filepath.Ext(base))
			}
		}
		items = append(items, item)
	}
	return items
}

func insertEpisodes(feed []byte, existing []feedItem, episodes []episode) []byte {
	if len(episodes) == 0 {
		return append([]byte(nil), feed...)
	}
	sort.SliceStable(episodes, func(i, j int) bool { return episodes[i].PublishedAt.After(episodes[j].PublishedAt) })
	buckets := make([][]episode, len(existing)+1)
	for _, ep := range episodes {
		index := len(existing)
		for i, item := range existing {
			if item.Date.IsZero() || !ep.PublishedAt.Before(item.Date) {
				index = i
				break
			}
		}
		buckets[index] = append(buckets[index], ep)
	}

	var out bytes.Buffer
	cursor := 0
	for i, item := range existing {
		out.Write(feed[cursor:item.Start])
		writeEpisodeBucket(&out, buckets[i])
		out.Write(feed[item.Start:item.End])
		cursor = item.End
	}
	if len(existing) == 0 {
		closing := bytes.Index(feed, []byte("</channel>"))
		if closing < 0 {
			return append([]byte(nil), feed...)
		}
		out.Reset()
		out.Write(feed[:closing])
		writeEpisodeBucket(&out, buckets[0])
		out.Write(feed[closing:])
		return out.Bytes()
	}
	out.Write(feed[cursor:])
	if len(buckets[len(existing)]) > 0 {
		result := out.Bytes()
		closing := bytes.LastIndex(result, []byte("</channel>"))
		if closing >= 0 {
			var fixed bytes.Buffer
			fixed.Write(result[:closing])
			writeEpisodeBucket(&fixed, buckets[len(existing)])
			fixed.Write(result[closing:])
			return fixed.Bytes()
		}
	}
	return out.Bytes()
}

func writeEpisodeBucket(out *bytes.Buffer, episodes []episode) {
	if len(episodes) == 0 {
		return
	}
	for _, ep := range episodes {
		if out.Len() > 0 {
			last := out.Bytes()[out.Len()-1]
			if last != '\n' {
				out.WriteByte('\n')
			}
		}
		out.WriteString(renderEpisode(ep))
		out.WriteByte('\n')
	}
}

func renderEpisode(ep episode) string {
	description := fmt.Sprintf("NotebookLMで作成した音声概要。元Notebook: %s", ep.NotebookTitle)
	var b strings.Builder
	b.WriteString("    <item>\n")
	fmt.Fprintf(&b, "      <title>%s</title>\n", escapeXML(ep.Title))
	fmt.Fprintf(&b, "      <description>%s</description>\n", escapeXML(description))
	fmt.Fprintf(&b, "      <pubDate>%s</pubDate>\n", ep.PublishedAt.UTC().Format(time.RFC1123Z))
	fmt.Fprintf(&b, "      <guid isPermaLink=\"false\">%s</guid>\n", escapeXML(ep.AudioID))
	fmt.Fprintf(&b, "      <enclosure url=\"%s\" type=\"audio/mp4\" length=\"%d\"/>\n", escapeXML(ep.EnclosureURL), ep.Length)
	if ep.Duration != "" {
		fmt.Fprintf(&b, "      <itunes:duration>%s</itunes:duration>\n", ep.Duration)
	}
	b.WriteString("      <itunes:explicit>false</itunes:explicit>\n")
	b.WriteString("    </item>")
	return b.String()
}

func escapeXML(value string) string {
	var b bytes.Buffer
	_ = xml.EscapeText(&b, []byte(value))
	return b.String()
}

func validateXML(data []byte) error {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	for {
		_, err := decoder.Token()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

func writeBothFeeds(pathA, pathB string, data []byte) error {
	tmpA, err := writeTemp(pathA, data)
	if err != nil {
		return fmt.Errorf("stage %s: %w", pathA, err)
	}
	defer os.Remove(tmpA)
	tmpB, err := writeTemp(pathB, data)
	if err != nil {
		return fmt.Errorf("stage %s: %w", pathB, err)
	}
	defer os.Remove(tmpB)
	if err := os.Rename(tmpA, pathA); err != nil {
		return fmt.Errorf("replace %s: %w", pathA, err)
	}
	if err := os.Rename(tmpB, pathB); err != nil {
		return fmt.Errorf("replace %s: %w", pathB, err)
	}
	return nil
}

func writeTemp(target string, data []byte) (string, error) {
	dir := filepath.Dir(target)
	file, err := os.CreateTemp(dir, ".podcast-backfill-*")
	if err != nil {
		return "", err
	}
	name := file.Name()
	if _, err := file.Write(data); err != nil {
		file.Close()
		os.Remove(name)
		return "", err
	}
	if err := file.Sync(); err != nil {
		file.Close()
		os.Remove(name)
		return "", err
	}
	if err := file.Close(); err != nil {
		os.Remove(name)
		return "", err
	}
	return name, nil
}

func writeReport(path string, rep report) error {
	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal report: %w", err)
	}
	data = append(data, '\n')
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil && filepath.Dir(path) != "." {
		return fmt.Errorf("create report directory: %w", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write report: %w", err)
	}
	return nil
}
