package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func validGlossaryFixture() *glossaryDocument {
	return &glossaryDocument{
		Version:     1,
		VerifiedAt:  "2026-08-14",
		SourceScope: "test fixture",
		Terms: []glossaryTerm{
			{
				ID:           "alpha",
				Term:         "Alpha",
				Aliases:      []string{"A"},
				Domain:       "test",
				DefinitionJA: "Alphaの定義。",
				WhyItMatters: "Alphaが重要な理由。",
				Example:      "Alphaの例。",
				RelatedTerms: []string{"beta"},
				Sources: []glossarySource{{
					Title:      "Alpha spec",
					URL:        "https://example.com/alpha",
					SourceType: "official_docs",
				}},
				VerifiedAt: "2026-08-14",
				Status:     "verified",
			},
			{
				ID:           "beta",
				Term:         "Beta",
				Aliases:      []string{},
				Domain:       "test",
				DefinitionJA: "Betaの定義。",
				WhyItMatters: "Betaが重要な理由。",
				Example:      "Betaの例。",
				RelatedTerms: []string{"alpha"},
				Sources: []glossarySource{{
					Title:      "Beta spec",
					URL:        "https://example.com/beta",
					SourceType: "standard",
				}},
				VerifiedAt: "2026-08-14",
				Status:     "verified",
			},
		},
	}
}

func TestValidateGlossary(t *testing.T) {
	if err := validateGlossary(validGlossaryFixture()); err != nil {
		t.Fatalf("validateGlossary() error = %v", err)
	}
}

func TestValidateGlossaryRejectsDuplicateAndBrokenReference(t *testing.T) {
	doc := validGlossaryFixture()
	doc.Terms[1].ID = "alpha"
	doc.Terms[1].Term = " alpha "
	doc.Terms[0].RelatedTerms = []string{"missing"}

	err := validateGlossary(doc)
	if err == nil {
		t.Fatal("validateGlossary() unexpectedly succeeded")
	}
	message := err.Error()
	for _, want := range []string{"duplicate id", "duplicate term", "references unknown id"} {
		if !strings.Contains(message, want) {
			t.Fatalf("validation error %q does not contain %q", message, want)
		}
	}
}

func TestValidateGlossaryRejectsBadSourceAndDate(t *testing.T) {
	doc := validGlossaryFixture()
	doc.Terms[0].VerifiedAt = "2026/08/14"
	doc.Terms[0].Sources[0].URL = "javascript:alert(1)"

	err := validateGlossary(doc)
	if err == nil {
		t.Fatal("validateGlossary() unexpectedly succeeded")
	}
	message := err.Error()
	for _, want := range []string{"must be YYYY-MM-DD", "absolute http(s) URL"} {
		if !strings.Contains(message, want) {
			t.Fatalf("validation error %q does not contain %q", message, want)
		}
	}
}

func TestLoadGlossaryRejectsUnknownField(t *testing.T) {
	path := filepath.Join(t.TempDir(), "terms.yaml")
	content := `version: 1
verified_at: "2026-08-14"
source_scope: test
unexpected: true
terms: []
`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := loadGlossary(path); err == nil || !strings.Contains(err.Error(), "field unexpected not found") {
		t.Fatalf("loadGlossary() error = %v, want strict unknown-field failure", err)
	}
}

func TestSearchAndShowGlossary(t *testing.T) {
	doc := validGlossaryFixture()

	var search bytes.Buffer
	if err := searchGlossary(&search, doc, "A"); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(search.String(), "Alpha") {
		t.Fatalf("search output = %q, want Alpha", search.String())
	}

	var show bytes.Buffer
	if err := showGlossary(&show, doc, "a"); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Alpha", "Definition: Alphaの定義。", "https://example.com/alpha"} {
		if !strings.Contains(show.String(), want) {
			t.Fatalf("show output = %q, want %q", show.String(), want)
		}
	}
}

func TestRenderGlossaryMarkdownDeterministic(t *testing.T) {
	doc := validGlossaryFixture()
	first := renderGlossaryMarkdown(doc)
	second := renderGlossaryMarkdown(doc)
	if !bytes.Equal(first, second) {
		t.Fatal("renderGlossaryMarkdown() is not deterministic")
	}
	for _, want := range []string{
		"直接編集しないでください",
		"## Alpha",
		"**なぜ重要か:** Alphaが重要な理由。",
		"[Alpha spec](https://example.com/alpha)",
	} {
		if !bytes.Contains(first, []byte(want)) {
			t.Fatalf("generated markdown missing %q", want)
		}
	}
}

func TestCanonicalGlossaryPassesValidation(t *testing.T) {
	t.Setenv("NLM_GLOSSARY_PATH", "")
	path, err := locateGlossaryDataPath()
	if err != nil {
		t.Fatal(err)
	}
	doc, err := loadGlossary(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := validateGlossary(doc); err != nil {
		t.Fatal(err)
	}
	if len(doc.Terms) < 10 {
		t.Fatalf("canonical glossary has %d terms, want at least 10", len(doc.Terms))
	}
}

func TestRunGlossaryCLICheckAndGenerate(t *testing.T) {
	dir := t.TempDir()
	dataDir := filepath.Join(dir, "data", "glossary")
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		t.Fatal(err)
	}
	dataPath := filepath.Join(dataDir, "terms.yaml")
	fixture := `version: 1
verified_at: "2026-08-14"
source_scope: fixture
terms:
  - id: alpha
    term: Alpha
    aliases: [A]
    domain: test
    definition_ja: Alphaの定義。
    why_it_matters: Alphaが重要な理由。
    example: Alphaの例。
    related_terms: []
    sources:
      - title: Alpha spec
        url: https://example.com/alpha
        source_type: official_docs
    verified_at: "2026-08-14"
    status: verified
`
	if err := os.WriteFile(dataPath, []byte(fixture), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("NLM_GLOSSARY_PATH", dataPath)

	var check bytes.Buffer
	if err := runGlossaryCLI([]string{"check"}, &check); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(check.String(), "OK: glossary v1, 1 terms") {
		t.Fatalf("check output = %q", check.String())
	}

	output := filepath.Join(dir, "glossary.md")
	var generate bytes.Buffer
	if err := runGlossaryCLI([]string{"generate", output}, &generate); err != nil {
		t.Fatal(err)
	}
	generated, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(generated, []byte("## Alpha")) {
		t.Fatalf("generated markdown = %q", generated)
	}
}
