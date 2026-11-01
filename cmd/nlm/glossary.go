package main

import (
	"fmt"
	"os"

	"github.com/tmc/nlm/internal/glossary"
)

// Preserve the legacy 'nlm glossary' command during CLI migration.
func init() {
	if len(os.Args) < 2 || os.Args[1] != "glossary" {
		return
	}
	if err := glossary.RunCLI(os.Args[2:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "nlm glossary: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
