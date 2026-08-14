package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	defaultGlossaryDataPath = "data/glossary/terms.yaml"
	defaultGlossaryDocPath  = "docs/glossary.md"
)

var glossaryIDPattern = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

type glossaryDocument struct {
	Version     int            `yaml:"version"`
	VerifiedAt  string         `yaml:"verified_at"`
	SourceScope string         `yaml:"source_scope"`
	Terms       []glossaryTerm `yaml:"terms"`
}

type glossaryTerm struct {
	ID           string           `yaml:"id"`
	Term         string           `yaml:"term"`
	Aliases      []string         `yaml:"aliases"`
	Domain       string           `yaml:"domain"`
	DefinitionJA string           `yaml:"definition_ja"`
	WhyItMatters string           `yaml:"why_it_matters"`
	Example      string           `yaml:"example"`
	RelatedTerms []string         `yaml:"related_terms"`
	Sources      []glossarySource `yaml:"sources"`
	VerifiedAt   string           `yaml:"verified_at"`
	Status       string           `yaml:"status"`
}

type glossarySource struct {
	Title      string `yaml:"title"`
	URL        string `yaml:"url"`
	SourceType string `yaml:"source_type"`
}

// glossary is a local, authentication-free command family. It is handled before
// the NotebookLM command dispatcher so learning data can be inspected and
// validated without browser credentials.
func init() {
	if len(os.Args) < 2 || os.Args[1] != "glossary" {
		return
	}
	if err := runGlossaryCLI(os.Args[2:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "nlm glossary: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func runGlossaryCLI(args []string, out io.Writer) error {
	if len(args) == 0 {
		return errors.New("usage: nlm glossary <list|search|show|check|generate> [arguments]")
	}

	dataPath, err := locateGlossaryDataPath()
	if err != nil {
		return err
	}
	doc, err := loadGlossary(dataPath)
	if err != nil {
		return err
	}

	switch args[0] {
	case "list":
		if len(args) != 1 {
			return errors.New("usage: nlm glossary list")
		}
		if err := validateGlossary(doc); err != nil {
			return err
		}
		return listGlossary(out, doc)
	case "search":
		if len(args) < 2 {
			return errors.New("usage: nlm glossary search <query>")
		}
		if err := validateGlossary(doc); err != nil {
			return err
		}
		return searchGlossary(out, doc, strings.Join(args[1:], " "))
	case "show":
		if len(args) < 2 {
			return errors.New("usage: nlm glossary show <term|id|alias>")
		}
		if err := validateGlossary(doc); err != nil {
			return err
		}
		return showGlossary(out, doc, strings.Join(args[1:], " "))
	case "check":
		if len(args) != 1 {
			return errors.New("usage: nlm glossary check")
		}
		if err := validateGlossary(doc); err != nil {
			return err
		}
		fmt.Fprintf(out, "OK: glossary v%d, %d terms, verified_at=%s\n", doc.Version, len(doc.Terms), doc.VerifiedAt)
		return nil
	case "generate":
		if len(args) > 2 {
			return errors.New("usage: nlm glossary generate [output-path]")
		}
		if err := validateGlossary(doc); err != nil {
			return err
		}
		outputPath := ""
		if len(args) == 2 {
			outputPath = args[1]
		} else {
			outputPath, err = defaultGlossaryOutputPath(dataPath)
			if err != nil {
				return err
			}
		}
		content := renderGlossaryMarkdown(doc)
		if err := os.WriteFile(outputPath, content, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", outputPath, err)
		}
		fmt.Fprintf(out, "generated %s (%d terms)\n", outputPath, len(doc.Terms))
		return nil
	default:
		return fmt.Errorf("unknown glossary command %q; expected list, search, show, check, or generate", args[0])
	}
}

func locateGlossaryDataPath() (string, error) {
	if path := strings.TrimSpace(os.Getenv("NLM_GLOSSARY_PATH")); path != "" {
		if _, err := os.Stat(path); err != nil {
			return "", fmt.Errorf("NLM_GLOSSARY_PATH %s: %w", path, err)
		}
		return path, nil
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get working directory: %w", err)
	}
	for {
		candidate := filepath.Join(dir, defaultGlossaryDataPath)
		if info, statErr := os.Stat(candidate); statErr == nil && !info.IsDir() {
			return candidate, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("could not locate %s; run inside the nlm repository or set NLM_GLOSSARY_PATH", defaultGlossaryDataPath)
}

func defaultGlossaryOutputPath(dataPath string) (string, error) {
	abs, err := filepath.Abs(dataPath)
	if err != nil {
		return "", fmt.Errorf("resolve glossary path: %w", err)
	}
	glossaryDir := filepath.Dir(abs)
	dataDir := filepath.Dir(glossaryDir)
	root := filepath.Dir(dataDir)
	if filepath.Base(dataDir) != "data" || filepath.Base(glossaryDir) != "glossary" {
		return "", errors.New("cannot infer repository root from glossary path; pass an explicit output path")
	}
	return filepath.Join(root, defaultGlossaryDocPath), nil
}

func loadGlossary(path string) (*glossaryDocument, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open glossary: %w", err)
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	dec.KnownFields(true)
	var doc glossaryDocument
	if err := dec.Decode(&doc); err != nil {
		return nil, fmt.Errorf("decode glossary YAML: %w", err)
	}
	var trailing any
	if err := dec.Decode(&trailing); err != io.EOF {
		if err == nil {
			return nil, errors.New("decode glossary YAML: multiple YAML documents are not allowed")
		}
		return nil, fmt.Errorf("decode glossary YAML: %w", err)
	}
	return &doc, nil
}

func validateGlossary(doc *glossaryDocument) error {
	var problems []string
	if doc == nil {
		return errors.New("glossary is nil")
	}
	if doc.Version != 1 {
		problems = append(problems, fmt.Sprintf("version must be 1, got %d", doc.Version))
	}
	if strings.TrimSpace(doc.SourceScope) == "" {
		problems = append(problems, "source_scope is required")
	}
	if err := validateDate(doc.VerifiedAt); err != nil {
		problems = append(problems, "verified_at: "+err.Error())
	}
	if len(doc.Terms) == 0 {
		problems = append(problems, "terms must contain at least one entry")
	}

	ids := make(map[string]int, len(doc.Terms))
	terms := make(map[string]int, len(doc.Terms))
	for i, term := range doc.Terms {
		prefix := fmt.Sprintf("terms[%d]", i)
		if strings.TrimSpace(term.ID) == "" {
			problems = append(problems, prefix+".id is required")
		} else {
			if !glossaryIDPattern.MatchString(term.ID) {
				problems = append(problems, fmt.Sprintf("%s.id %q must match %s", prefix, term.ID, glossaryIDPattern.String()))
			}
			if first, ok := ids[term.ID]; ok {
				problems = append(problems, fmt.Sprintf("duplicate id %q at terms[%d] and terms[%d]", term.ID, first, i))
			} else {
				ids[term.ID] = i
			}
		}

		canonical := normalizeLookup(term.Term)
		if canonical == "" {
			problems = append(problems, prefix+".term is required")
		} else if first, ok := terms[canonical]; ok {
			problems = append(problems, fmt.Sprintf("duplicate term %q at terms[%d] and terms[%d]", term.Term, first, i))
		} else {
			terms[canonical] = i
		}

		if term.Aliases == nil {
			problems = append(problems, prefix+".aliases is required (use [] when empty)")
		}
		requiredString(&problems, prefix+".domain", term.Domain)
		requiredString(&problems, prefix+".definition_ja", term.DefinitionJA)
		requiredString(&problems, prefix+".why_it_matters", term.WhyItMatters)
		requiredString(&problems, prefix+".example", term.Example)
		if term.RelatedTerms == nil {
			problems = append(problems, prefix+".related_terms is required (use [] when empty)")
		}
		if term.Sources == nil {
			problems = append(problems, prefix+".sources is required (use [] when empty)")
		}
		if err := validateDate(term.VerifiedAt); err != nil {
			problems = append(problems, prefix+".verified_at: "+err.Error())
		}
		if term.Status != "verified" && term.Status != "needs_review" {
			problems = append(problems, fmt.Sprintf("%s.status must be verified or needs_review, got %q", prefix, term.Status))
		}
		if term.Status == "verified" && len(term.Sources) == 0 {
			problems = append(problems, prefix+".sources must contain at least one source for verified entries")
		}
		for j, source := range term.Sources {
			sourcePrefix := fmt.Sprintf("%s.sources[%d]", prefix, j)
			requiredString(&problems, sourcePrefix+".title", source.Title)
			requiredString(&problems, sourcePrefix+".source_type", source.SourceType)
			if err := validateHTTPURL(source.URL); err != nil {
				problems = append(problems, sourcePrefix+".url: "+err.Error())
			}
		}
	}

	for i, term := range doc.Terms {
		seen := map[string]struct{}{}
		for _, related := range term.RelatedTerms {
			if _, duplicate := seen[related]; duplicate {
				problems = append(problems, fmt.Sprintf("terms[%d].related_terms contains duplicate %q", i, related))
				continue
			}
			seen[related] = struct{}{}
			if _, ok := ids[related]; !ok {
				problems = append(problems, fmt.Sprintf("terms[%d].related_terms references unknown id %q", i, related))
			}
		}
	}

	if len(problems) > 0 {
		sort.Strings(problems)
		return fmt.Errorf("glossary validation failed:\n- %s", strings.Join(problems, "\n- "))
	}
	return nil
}

func requiredString(problems *[]string, field, value string) {
	if strings.TrimSpace(value) == "" {
		*problems = append(*problems, field+" is required")
	}
}

func validateDate(value string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New("date is required")
	}
	parsed, err := time.Parse("2006-01-02", value)
	if err != nil || parsed.Format("2006-01-02") != value {
		return fmt.Errorf("must be YYYY-MM-DD, got %q", value)
	}
	return nil
}

func validateHTTPURL(value string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New("URL is required")
	}
	parsed, err := url.Parse(value)
	if err != nil || parsed.Host == "" || (parsed.Scheme != "https" && parsed.Scheme != "http") {
		return fmt.Errorf("must be an absolute http(s) URL, got %q", value)
	}
	return nil
}

func normalizeLookup(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func listGlossary(out io.Writer, doc *glossaryDocument) error {
	terms := append([]glossaryTerm(nil), doc.Terms...)
	sort.Slice(terms, func(i, j int) bool {
		return strings.ToLower(terms[i].Term) < strings.ToLower(terms[j].Term)
	})
	w := tabwriter.NewWriter(out, 0, 4, 2, ' ', 0)
	fmt.Fprintln(w, "TERM\tDOMAIN\tSTATUS\tID")
	for _, term := range terms {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", term.Term, term.Domain, term.Status, term.ID)
	}
	return w.Flush()
}

func searchGlossary(out io.Writer, doc *glossaryDocument, query string) error {
	needle := normalizeLookup(query)
	if needle == "" {
		return errors.New("search query must not be empty")
	}
	var matches []glossaryTerm
	for _, term := range doc.Terms {
		haystack := []string{term.ID, term.Term, term.Domain, term.DefinitionJA, term.WhyItMatters, term.Example}
		haystack = append(haystack, term.Aliases...)
		matched := false
		for _, value := range haystack {
			if strings.Contains(normalizeLookup(value), needle) {
				matched = true
				break
			}
		}
		if matched {
			matches = append(matches, term)
		}
	}
	sort.Slice(matches, func(i, j int) bool {
		return strings.ToLower(matches[i].Term) < strings.ToLower(matches[j].Term)
	})
	if len(matches) == 0 {
		fmt.Fprintf(out, "No glossary terms matched %q.\n", query)
		return nil
	}
	w := tabwriter.NewWriter(out, 0, 4, 2, ' ', 0)
	fmt.Fprintln(w, "TERM\tDOMAIN\tSTATUS\tID")
	for _, term := range matches {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", term.Term, term.Domain, term.Status, term.ID)
	}
	return w.Flush()
}

func showGlossary(out io.Writer, doc *glossaryDocument, query string) error {
	needle := normalizeLookup(query)
	if needle == "" {
		return errors.New("term must not be empty")
	}
	for _, term := range doc.Terms {
		if normalizeLookup(term.ID) == needle || normalizeLookup(term.Term) == needle || containsNormalized(term.Aliases, needle) {
			writeGlossaryTerm(out, term)
			return nil
		}
	}
	return fmt.Errorf("term %q not found", query)
}

func containsNormalized(values []string, needle string) bool {
	for _, value := range values {
		if normalizeLookup(value) == needle {
			return true
		}
	}
	return false
}

func writeGlossaryTerm(out io.Writer, term glossaryTerm) {
	fmt.Fprintln(out, term.Term)
	fmt.Fprintf(out, "ID: %s\n", term.ID)
	fmt.Fprintf(out, "Domain: %s\n", term.Domain)
	fmt.Fprintf(out, "Status: %s\n", term.Status)
	if len(term.Aliases) > 0 {
		fmt.Fprintf(out, "Aliases: %s\n", strings.Join(term.Aliases, ", "))
	} else {
		fmt.Fprintln(out, "Aliases: -")
	}
	fmt.Fprintf(out, "Definition: %s\n", term.DefinitionJA)
	fmt.Fprintf(out, "Why it matters: %s\n", term.WhyItMatters)
	fmt.Fprintf(out, "Example: %s\n", term.Example)
	if len(term.RelatedTerms) > 0 {
		fmt.Fprintf(out, "Related: %s\n", strings.Join(term.RelatedTerms, ", "))
	} else {
		fmt.Fprintln(out, "Related: -")
	}
	fmt.Fprintf(out, "Verified at: %s\n", term.VerifiedAt)
	fmt.Fprintln(out, "Sources:")
	for _, source := range term.Sources {
		fmt.Fprintf(out, "- %s — %s (%s)\n", source.Title, source.URL, source.SourceType)
	}
}

func renderGlossaryMarkdown(doc *glossaryDocument) []byte {
	var b bytes.Buffer
	fmt.Fprintln(&b, "# 学習用語集")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "> `data/glossary/terms.yaml` から自動生成。詳しく見たい語は `nlm glossary show <term>`。直接編集しないでください。")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "| 用語 | 一言説明 | 一次情報 |")
	fmt.Fprintln(&b, "|---|---|---|")
	for _, term := range doc.Terms {
		source := "-"
		if len(term.Sources) > 0 {
			source = fmt.Sprintf("[%s](%s)", markdownTableCell(term.Sources[0].Title), term.Sources[0].URL)
		}
		fmt.Fprintf(
			&b,
			"| %s | %s | %s |\n",
			markdownTableCell(term.Term),
			markdownTableCell(shortGlossaryDefinition(term.DefinitionJA)),
			source,
		)
	}
	return b.Bytes()
}

func shortGlossaryDefinition(value string) string {
	value = strings.Join(strings.Fields(value), " ")
	if i := strings.Index(value, "。"); i >= 0 {
		return value[:i+len("。")]
	}
	return value
}

func markdownTableCell(value string) string {
	value = strings.Join(strings.Fields(value), " ")
	return strings.ReplaceAll(value, "|", `\|`)
}
