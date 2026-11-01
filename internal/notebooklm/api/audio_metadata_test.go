package api

import (
	"strings"
	"testing"
)

func TestParseAudioArtifactMetadata(t *testing.T) {
	// Artifact IDs must be available without requiring audio URL or bytes.
	got, err := parseAudioArtifactMetadata([]byte(`[[[" audio-123 ","Episode",2,null,null,null,null]]]`), "notebook-1")
	if err != nil {
		t.Fatal(err)
	}
	if got.ProjectID != "notebook-1" || got.AudioID != "audio-123" || got.Title != "Episode" {
		t.Fatalf("unexpected metadata: %+v", got)
	}
	if got.AudioData != "" {
		t.Fatal("metadata lookup must not return downloaded audio")
	}
}

func TestParseAudioArtifactMetadataNoAudio(t *testing.T) {
	for _, input := range []string{`[]`, `[[]]`} {
		_, err := parseAudioArtifactMetadata([]byte(input), "notebook-1")
		if err == nil || !strings.Contains(err.Error(), "no audio") {
			t.Fatalf("input %s: expected skippable no-audio error, got %v", input, err)
		}
	}
}

func TestParseAudioArtifactMetadataRejectsMalformed(t *testing.T) {
	for _, input := range []string{`invalid`, `[["not an artifact"]]`, `[[[]]]`} {
		if _, err := parseAudioArtifactMetadata([]byte(input), "notebook-1"); err == nil {
			t.Fatalf("input %s: expected parsing error", input)
		}
	}
}
