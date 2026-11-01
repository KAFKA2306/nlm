package main

import (
	"fmt"
	"os"

	"github.com/tmc/nlm/internal/glossary"
)

func main() {
	if err := glossary.RunCLI(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "glossary: %v\n", err)
		os.Exit(1)
	}
}
