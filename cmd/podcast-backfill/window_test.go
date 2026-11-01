package main

import (
	"testing"
	"time"
)

func TestUntilJSTIncludesEntireDateButNotFollowingDay(t *testing.T) {
	since, err := parseSince("2026-09-01")
	if err != nil {
		t.Fatal(err)
	}
	until, err := parseUntil("2026-09-30")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name string
		at   string
		want bool
	}{
		{"before start", "2026-08-31T23:59:59+09:00", false},
		{"start included", "2026-09-01T00:00:00+09:00", true},
		{"end included", "2026-09-30T23:59:59+09:00", true},
		{"following day excluded", "2026-10-01T00:00:00+09:00", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			when, err := time.Parse(time.RFC3339, tt.at)
			if err != nil {
				t.Fatal(err)
			}
			if got := inWindow(when, since, until); got != tt.want {
				t.Fatalf("inWindow(%s) = %v, want %v", tt.at, got, tt.want)
			}
		})
	}
}

func TestUntilTimestampIncludesExactInstant(t *testing.T) {
	until, err := parseUntil("2026-09-30T12:00:00Z")
	if err != nil {
		t.Fatal(err)
	}
	at := time.Date(2026, 9, 30, 12, 0, 0, 0, time.UTC)
	if !inWindow(at, at.Add(-time.Hour), until) {
		t.Fatal("exact end timestamp must be included")
	}
	if inWindow(at.Add(time.Nanosecond), at.Add(-time.Hour), until) {
		t.Fatal("later timestamp must be excluded")
	}
}

func TestUntilRejectsInvalidDate(t *testing.T) {
	for _, value := range []string{"2026-02-30", "2026/09/30", "tomorrow"} {
		if _, err := parseUntil(value); err == nil {
			t.Fatalf("parseUntil(%q) unexpectedly succeeded", value)
		}
	}
}

func TestWindowWithoutUpperBound(t *testing.T) {
	since := time.Date(2026, 9, 1, 0, 0, 0, 0, time.UTC)
	if !inWindow(since.Add(30*24*time.Hour), since, time.Time{}) {
		t.Fatal("open-ended window must preserve existing behavior")
	}
}
