package main

import (
	"encoding/binary"
	"strings"
	"testing"
	"time"
)

func TestParseSinceDateUsesJST(t *testing.T) {
	got, err := parseSince("2026-06-15")
	if err != nil {
		t.Fatal(err)
	}
	if got.Format(time.RFC3339) != "2026-06-15T00:00:00+09:00" {
		t.Fatalf("unexpected date: %s", got.Format(time.RFC3339))
	}
}

func TestNotebookTimePrefersModified(t *testing.T) {
	created := time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)
	modified := time.Date(2026, 8, 20, 3, 4, 5, 0, time.UTC)
	got, ok := notebookTimeFromValues(modified, created)
	if !ok || !got.Equal(modified) {
		t.Fatalf("got %v, %v; want modified time", got, ok)
	}
}

func TestParseFeedItemsUsesEnclosureAudioID(t *testing.T) {
	feed := []byte(`<rss><channel><item>
<title>old</title><pubDate>Sat, 20 Jun 2026 00:00:00 +0000</pubDate>
<guid isPermaLink="false">legacy-guid</guid>
<enclosure url="https://archive.org/download/kafka-podcast-vault-v2/abc-123.m4a" type="audio/mp4" length="1"/>
</item></channel></rss>`)
	items := parseFeedItems(feed)
	if len(items) != 1 {
		t.Fatalf("got %d items", len(items))
	}
	if items[0].GUID != "legacy-guid" || items[0].AudioKey != "abc-123" {
		t.Fatalf("unexpected parsed item: %#v", items[0])
	}
}

func TestInsertEpisodesPlacesNewerEpisodeBeforeExistingWithoutRewritingOldItem(t *testing.T) {
	feed := []byte(`<?xml version="1.0"?><rss xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd"><channel>
  <title>KAFKA探究室</title>
    <item>
      <title>old &amp; preserved</title>
      <pubDate>Sat, 20 Jun 2026 00:00:00 +0000</pubDate>
      <guid isPermaLink="false">old-guid</guid>
      <enclosure url="https://archive.org/download/kafka-podcast-vault-v2/old-audio.m4a" type="audio/mp4" length="1"/>
    </item>
</channel></rss>`)
	oldRaw := parseFeedItems(feed)[0].Raw
	ep := episode{
		AudioID:       "new-audio",
		NotebookTitle: "Notebook A",
		Title:         "new & title",
		PublishedAt:   time.Date(2026, 8, 20, 0, 0, 0, 0, time.UTC),
		Length:        123,
		Duration:      "00:01:02",
		EnclosureURL:  "https://archive.org/download/kafka-podcast-vault-v2/new-audio.m4a",
	}
	updated := insertEpisodes(feed, parseFeedItems(feed), []episode{ep})
	if err := validateXML(updated); err != nil {
		t.Fatalf("invalid XML: %v\n%s", err, updated)
	}
	text := string(updated)
	if strings.Index(text, "new-audio") > strings.Index(text, "old-guid") {
		t.Fatalf("new episode was not inserted before old episode:\n%s", text)
	}
	if !strings.Contains(text, oldRaw) {
		t.Fatal("existing item was rewritten")
	}
	if !strings.Contains(text, "new &amp; title") {
		t.Fatal("new title was not XML-escaped")
	}
}

func TestMP4DurationV0(t *testing.T) {
	mvhdPayload := make([]byte, 20)
	mvhdPayload[0] = 0
	binary.BigEndian.PutUint32(mvhdPayload[12:16], 1000)
	binary.BigEndian.PutUint32(mvhdPayload[16:20], 62_000)
	data := mp4Box("moov", mp4Box("mvhd", mvhdPayload))
	got, err := mp4Duration(data)
	if err != nil {
		t.Fatal(err)
	}
	if got != "00:01:02" {
		t.Fatalf("got %q", got)
	}
}

func TestMP4DurationV1(t *testing.T) {
	mvhdPayload := make([]byte, 32)
	mvhdPayload[0] = 1
	binary.BigEndian.PutUint32(mvhdPayload[20:24], 48_000)
	binary.BigEndian.PutUint64(mvhdPayload[24:32], 48_000*3601)
	data := mp4Box("moov", mp4Box("mvhd", mvhdPayload))
	got, err := mp4Duration(data)
	if err != nil {
		t.Fatal(err)
	}
	if got != "01:00:01" {
		t.Fatalf("got %q", got)
	}
}

func mp4Box(kind string, payload []byte) []byte {
	box := make([]byte, 8+len(payload))
	binary.BigEndian.PutUint32(box[:4], uint32(len(box)))
	copy(box[4:8], kind)
	copy(box[8:], payload)
	return box
}
