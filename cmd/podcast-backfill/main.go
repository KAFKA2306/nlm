package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	pb "github.com/tmc/nlm/gen/notebooklm/v1alpha1"
	"github.com/tmc/nlm/internal/notebooklm/api"
)

const (
	defaultArchiveID = "kafka-podcast-vault-v2"
	defaultFeedPath  = "podcast/feed.xml"
	defaultPagesFeed = "docs/feed.xml"
)

type options struct {
	Since       time.Time
	Months      int
	Apply       bool
	FeedPath    string
	PagesFeed   string
	ArchiveID   string
	ReportPath  string
	Limit       int
	HTTPTimeout time.Duration
}

type episode struct {
	AudioID       string    `json:"audio_id"`
	NotebookTitle string    `json:"notebook_title"`
	Title         string    `json:"title"`
	PublishedAt   time.Time `json:"published_at"`
	Filename      string    `json:"filename"`
	Length        int       `json:"length"`
	Duration      string    `json:"duration,omitempty"`
	EnclosureURL  string    `json:"enclosure_url"`
}

type skipped struct {
	NotebookTitle string `json:"notebook_title"`
	Reason        string `json:"reason"`
}

type report struct {
	StartedAt   time.Time `json:"started_at"`
	Since       time.Time `json:"since"`
	Apply       bool      `json:"apply"`
	Candidates  int       `json:"candidates"`
	Published   []episode `json:"published"`
	Skipped     []skipped `json:"skipped"`
	FeedUpdated bool      `json:"feed_updated"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "podcast-backfill: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	opts, err := parseOptions()
	if err != nil {
		return err
	}

	nlmToken := strings.TrimSpace(os.Getenv("NLM_AUTH_TOKEN"))
	nlmCookies := strings.TrimSpace(os.Getenv("NLM_COOKIES"))
	if nlmToken == "" || nlmCookies == "" {
		return errors.New("NLM_AUTH_TOKEN and NLM_COOKIES are required")
	}

	feedBytes, err := os.ReadFile(opts.FeedPath)
	if err != nil {
		return fmt.Errorf("read canonical feed: %w", err)
	}
	pagesFeedBytes, err := os.ReadFile(opts.PagesFeed)
	if err != nil {
		return fmt.Errorf("read Pages feed: %w", err)
	}
	if !bytes.Equal(feedBytes, pagesFeedBytes) {
		return fmt.Errorf("feed invariant failed: %s and %s differ", opts.FeedPath, opts.PagesFeed)
	}
	if err := validateXML(feedBytes); err != nil {
		return fmt.Errorf("existing feed is invalid XML: %w", err)
	}

	existingItems := parseFeedItems(feedBytes)
	existingKeys := map[string]struct{}{}
	for _, item := range existingItems {
		if item.GUID != "" {
			existingKeys[item.GUID] = struct{}{}
		}
		if item.AudioKey != "" {
			existingKeys[item.AudioKey] = struct{}{}
		}
	}

	started := time.Now().UTC()
	rep := report{StartedAt: started, Since: opts.Since.UTC(), Apply: opts.Apply}

	client := api.New(nlmToken, nlmCookies)
	notebooks, err := client.ListRecentlyViewedProjects()
	if err != nil {
		return fmt.Errorf("list NotebookLM notebooks: %w", err)
	}

	type candidate struct {
		ID    string
		Title string
		When  time.Time
	}
	var candidates []candidate
	for _, notebook := range notebooks {
		when, ok := notebookActivityTime(notebook.GetMetadata())
		if !ok || when.Before(opts.Since) {
			continue
		}
		candidates = append(candidates, candidate{
			ID: notebook.GetProjectId(), Title: strings.TrimSpace(notebook.GetTitle()), When: when,
		})
	}
	sort.SliceStable(candidates, func(i, j int) bool { return candidates[i].When.After(candidates[j].When) })
	if opts.Limit > 0 && len(candidates) > opts.Limit {
		candidates = candidates[:opts.Limit]
	}
	rep.Candidates = len(candidates)
	fmt.Printf("NotebookLM backfill window: %s .. now\n", opts.Since.Format(time.RFC3339))
	fmt.Printf("Candidates: %d\n", len(candidates))

	var archive *archiveClient
	if opts.Apply {
		accessKey := firstNonEmpty(os.Getenv("IA_ACCESS_KEY"), os.Getenv("ARCHIVE_ACCESS_KEY"), os.Getenv("ARCHIVE_ORG_ACCESS_KEY"))
		secretKey := firstNonEmpty(os.Getenv("IA_SECRET_KEY"), os.Getenv("ARCHIVE_SECRET_KEY"), os.Getenv("ARCHIVE_ORG_SECRET_KEY"))
		if accessKey == "" || secretKey == "" {
			return errors.New("Archive.org credentials are required: IA_ACCESS_KEY/IA_SECRET_KEY (or ARCHIVE_* aliases)")
		}
		archive = &archiveClient{
			identifier: opts.ArchiveID,
			accessKey:  accessKey,
			secretKey:  secretKey,
			http:       &http.Client{Timeout: opts.HTTPTimeout},
		}
	}

	var newEpisodes []episode
	for _, notebook := range candidates {
		fmt.Printf("Inspecting %s (%s)\n", notebook.Title, notebook.When.Format(time.RFC3339))
		audio, err := client.DownloadAudioOverview(notebook.ID)
		if err != nil {
			if isSkippableAudioError(err) {
				rep.Skipped = append(rep.Skipped, skipped{NotebookTitle: notebook.Title, Reason: compactReason(err)})
				fmt.Printf("  skip: %s\n", compactReason(err))
				continue
			}
			return fmt.Errorf("download audio for notebook %q: %w", notebook.Title, err)
		}
		if audio == nil || strings.TrimSpace(audio.AudioID) == "" {
			rep.Skipped = append(rep.Skipped, skipped{NotebookTitle: notebook.Title, Reason: "audio overview has no stable audio ID"})
			continue
		}
		audioID := strings.TrimSpace(audio.AudioID)
		if _, exists := existingKeys[audioID]; exists {
			rep.Skipped = append(rep.Skipped, skipped{NotebookTitle: notebook.Title, Reason: "already present in RSS"})
			fmt.Printf("  skip: already present in RSS (%s)\n", audioID)
			continue
		}
		data, err := decodeAudioData(audio.AudioData)
		if err != nil {
			return fmt.Errorf("decode audio %s: %w", audioID, err)
		}
		filename := audioID + ".m4a"
		title := strings.TrimSpace(audio.Title)
		if title == "" {
			title = notebook.Title
		}
		if title == "" {
			title = "NotebookLM Audio Overview"
		}
		duration, _ := mp4Duration(data)
		ep := episode{
			AudioID:       audioID,
			NotebookTitle: notebook.Title,
			Title:         title,
			PublishedAt:   notebook.When.UTC(),
			Filename:      filename,
			Length:        len(data),
			Duration:      duration,
			EnclosureURL:  fmt.Sprintf("https://archive.org/download/%s/%s", opts.ArchiveID, filename),
		}

		if opts.Apply {
			ctx, cancel := context.WithTimeout(context.Background(), opts.HTTPTimeout)
			exists, err := archive.exists(ctx, filename)
			cancel()
			if err != nil {
				return fmt.Errorf("check Archive.org for %s: %w", filename, err)
			}
			if !exists {
				ctx, cancel = context.WithTimeout(context.Background(), opts.HTTPTimeout)
				err = archive.upload(ctx, filename, data)
				cancel()
				if err != nil {
					return fmt.Errorf("upload %s to Archive.org: %w", filename, err)
				}
				if err := archive.waitAvailable(filename, 20, 3*time.Second); err != nil {
					return err
				}
				fmt.Printf("  uploaded: %s\n", filename)
			} else {
				fmt.Printf("  archive exists: %s\n", filename)
			}
		}
		newEpisodes = append(newEpisodes, ep)
		existingKeys[audioID] = struct{}{}
	}

	if len(newEpisodes) > 0 {
		updated := insertEpisodes(feedBytes, existingItems, newEpisodes)
		if err := validateXML(updated); err != nil {
			return fmt.Errorf("generated feed is invalid XML: %w", err)
		}
		if opts.Apply {
			if err := writeBothFeeds(opts.FeedPath, opts.PagesFeed, updated); err != nil {
				return err
			}
			rep.FeedUpdated = true
		}
		rep.Published = append(rep.Published, newEpisodes...)
	}

	if opts.ReportPath != "" {
		if err := writeReport(opts.ReportPath, rep); err != nil {
			return err
		}
	}
	fmt.Printf("Ready episodes: %d, skipped: %d, feed updated: %v\n", len(rep.Published), len(rep.Skipped), rep.FeedUpdated)
	if !opts.Apply {
		fmt.Println("Dry run only. Re-run with --apply to upload and write RSS.")
	}
	return nil
}

func parseOptions() (options, error) {
	var sinceText string
	var months int
	var apply bool
	var feedPath, pagesFeed, archiveID, reportPath string
	var limit int
	flag.StringVar(&sinceText, "since", "", "include NotebookLM work on/after YYYY-MM-DD or RFC3339")
	flag.IntVar(&months, "months", 3, "lookback in calendar months when --since is omitted")
	flag.BoolVar(&apply, "apply", false, "upload missing audio and update RSS feeds")
	flag.StringVar(&feedPath, "feed", defaultFeedPath, "canonical podcast RSS path")
	flag.StringVar(&pagesFeed, "pages-feed", defaultPagesFeed, "GitHub Pages RSS path")
	flag.StringVar(&archiveID, "archive-id", defaultArchiveID, "Archive.org item identifier")
	flag.StringVar(&reportPath, "report", "", "optional JSON report path")
	flag.IntVar(&limit, "limit", 0, "optional maximum number of notebooks to inspect")
	flag.Parse()
	if months < 1 {
		return options{}, errors.New("--months must be >= 1")
	}
	if archiveID == "" {
		return options{}, errors.New("--archive-id cannot be empty")
	}
	var since time.Time
	var err error
	if strings.TrimSpace(sinceText) != "" {
		since, err = parseSince(sinceText)
		if err != nil {
			return options{}, err
		}
	} else {
		jst := time.FixedZone("JST", 9*60*60)
		now := time.Now().In(jst)
		since = now.AddDate(0, -months, 0)
	}
	return options{
		Since: since, Months: months, Apply: apply, FeedPath: feedPath, PagesFeed: pagesFeed,
		ArchiveID: archiveID, ReportPath: reportPath, Limit: limit, HTTPTimeout: 5 * time.Minute,
	}, nil
}

func parseSince(value string) (time.Time, error) {
	value = strings.TrimSpace(value)
	if t, err := time.Parse(time.RFC3339, value); err == nil {
		return t, nil
	}
	jst := time.FixedZone("JST", 9*60*60)
	if t, err := time.ParseInLocation("2006-01-02", value, jst); err == nil {
		return t, nil
	}
	return time.Time{}, fmt.Errorf("invalid --since %q: use YYYY-MM-DD or RFC3339", value)
}

func notebookActivityTime(meta *pb.ProjectMetadata) (time.Time, bool) {
	if meta == nil {
		return time.Time{}, false
	}
	var modified, created time.Time
	if ts := meta.GetModifiedTime(); ts != nil {
		modified = ts.AsTime()
	}
	if ts := meta.GetCreateTime(); ts != nil {
		created = ts.AsTime()
	}
	return notebookTimeFromValues(modified, created)
}

func notebookTimeFromValues(modified, created time.Time) (time.Time, bool) {
	if !modified.IsZero() && modified.Year() > 2000 {
		return modified, true
	}
	if !created.IsZero() && created.Year() > 2000 {
		return created, true
	}
	return time.Time{}, false
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

func decodeAudioData(encoded string) ([]byte, error) {
	if strings.TrimSpace(encoded) == "" {
		return nil, errors.New("audio payload is empty")
	}
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("decoded audio payload is empty")
	}
	return data, nil
}

func isSkippableAudioError(err error) bool {
	if err == nil {
		return false
	}
	text := strings.ToLower(err.Error())
	for _, needle := range []string{"no audio", "not found", "does not exist", "audio may not be ready", "audio url list not found"} {
		if strings.Contains(text, needle) {
			return true
		}
	}
	return false
}

func compactReason(err error) string {
	if err == nil {
		return ""
	}
	text := strings.ReplaceAll(err.Error(), "\n", " ")
	if len(text) > 180 {
		text = text[:177] + "..."
	}
	return text
}
