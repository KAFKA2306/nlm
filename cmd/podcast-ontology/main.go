package main

import (
	"encoding/csv"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Feed       string            `yaml:"feed"`
	Stages     []string          `yaml:"stages"`
	Done       []string          `yaml:"done_stages"`
	OutDir     string            `yaml:"out_dir"`
	Predicates map[string]string `yaml:"predicates"`
}

type Enclosure struct {
	URL    string `xml:"url,attr"`
	Type   string `xml:"type,attr"`
	Length int64  `xml:"length,attr"`
}

type Item struct {
	Title    string    `xml:"title"`
	PubDate  string    `xml:"pubDate"`
	GUID     string    `xml:"guid"`
	Duration string    `xml:"duration"`
	Enc      Enclosure `xml:"enclosure"`
}

type Feed struct {
	Title string `xml:"channel>title"`
	Items []Item `xml:"channel>item"`
}

func write(dir, name string, rows [][]string) {
	f, _ := os.Create(filepath.Join(dir, name))
	defer f.Close()
	csv.NewWriter(f).WriteAll(rows)
}

func seconds(d string) int {
	total := 0
	for _, p := range strings.Split(d, ":") {
		n, _ := strconv.Atoi(p)
		total = total*60 + n
	}
	return total
}

func order(stages []string, name string) int {
	for i, s := range stages {
		if s == name {
			return i + 1
		}
	}
	return 0
}

func dimension(header string, keys map[string]bool) [][]string {
	rows := [][]string{{header}}
	names := make([]string, 0, len(keys))
	for k := range keys {
		names = append(names, k)
	}
	sort.Strings(names)
	for _, k := range names {
		rows = append(rows, []string{k})
	}
	return rows
}

func main() {
	var cfg Config
	raw, _ := os.ReadFile(os.Args[1])
	yaml.Unmarshal(raw, &cfg)
	var feed Feed
	resp, _ := http.Get(cfg.Feed)
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	xml.Unmarshal(data, &feed)
	os.MkdirAll(cfg.OutDir, 0o755)

	dates, hosts, formats := map[string]bool{}, map[string]bool{}, map[string]bool{}
	fact := [][]string{{"guid", "channel", "date", "host", "format", "duration_sec", "bytes"}}
	ledger := [][]string{{"guid", "date", "title", "url", "bytes", "duration"}}
	triples := [][]string{{"subject", "predicate", "object"}}
	progress := [][]string{{"guid", "stage", "stage_order", "status", "at"}}
	p := cfg.Predicates
	latest := ""
	for _, it := range feed.Items {
		t, _ := time.Parse(time.RFC1123, it.PubDate)
		date := t.Format("2006-01-02")
		u, _ := url.Parse(it.Enc.URL)
		if date > latest {
			latest = date
		}
		dates[date], hosts[u.Host], formats[it.Enc.Type] = true, true, true
		sec, size := strconv.Itoa(seconds(it.Duration)), strconv.FormatInt(it.Enc.Length, 10)
		fact = append(fact, []string{it.GUID, feed.Title, date, u.Host, it.Enc.Type, sec, size})
		ledger = append(ledger, []string{it.GUID, date, it.Title, it.Enc.URL, size, it.Duration})
		for _, st := range cfg.Done {
			progress = append(progress, []string{it.GUID, st, strconv.Itoa(order(cfg.Stages, st)), "done", date})
			triples = append(triples, []string{it.GUID, p["progress"], st})
		}
		triples = append(triples,
			[]string{it.GUID, p["title"], it.Title},
			[]string{it.GUID, p["channel"], feed.Title},
			[]string{it.GUID, p["date"], date},
			[]string{it.GUID, p["host"], u.Host},
			[]string{it.GUID, p["format"], it.Enc.Type},
			[]string{it.GUID, p["duration"], sec},
			[]string{it.GUID, p["bytes"], size},
		)
	}
	write(cfg.OutDir, "dim_channel.csv", [][]string{{"channel"}, {feed.Title}})
	write(cfg.OutDir, "dim_date.csv", dimension("date", dates))
	write(cfg.OutDir, "dim_host.csv", dimension("host", hosts))
	write(cfg.OutDir, "dim_format.csv", dimension("format", formats))
	stageDim := [][]string{{"stage", "stage_order"}}
	for _, st := range cfg.Stages {
		stageDim = append(stageDim, []string{st, strconv.Itoa(order(cfg.Stages, st))})
	}
	write(cfg.OutDir, "dim_stage.csv", stageDim)
	write(cfg.OutDir, "fact_progress.csv", progress)
	write(cfg.OutDir, "fact_episode.csv", fact)
	write(cfg.OutDir, "triples.csv", triples)
	write(cfg.OutDir, "ledger.csv", ledger)
	write(cfg.OutDir, "status.csv", [][]string{
		{"feed", "http_status", "last_modified", "episodes", "latest_pub_date", "fetched_at"},
		{cfg.Feed, strconv.Itoa(resp.StatusCode), resp.Header.Get("Last-Modified"), strconv.Itoa(len(feed.Items)), latest, time.Now().UTC().Format(time.RFC3339)},
	})
}
