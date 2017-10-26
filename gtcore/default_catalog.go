package gtcore

import (
	"encoding/json"
	"log"
	"os"
)

var defaultTools = []Tool{
	{
		Path: "github.com/golang/dep/cmd/dep",
		Desc: "Go dependency tool",
	},
	{
		Path: "github.com/kisielk/errcheck",
		Desc: "errcheck checks that you checked errors.",
	},
	{
		Path: "github.com/nsf/gocode",
		Desc: "An autocompletion daemon for the Go programming language",
	},
	{
		Path: "github.com/fzipp/gocyclo",
		Desc: "Calculate cyclomatic complexities of functions in Go source code.",
	},
	{
		Path: "golang.org/x/tools/cmd/goimports",
		Desc: "updates your Go import lines, adding missing ones and removing unreferenced ones.",
	},
	{
		Path: "github.com/golang/lint/golint",
		Desc: "a linter for Go source code",
	},
	{
		Path: "golang.org/x/tools/cmd/gorename",
		Desc: "The gorename command performs precise type-safe renaming of identifiers in Go source code.",
	},
	{
		Path: "github.com/jstemmer/gotags",
		Desc: "ctags-compatible tag generator for Go",
	},
	{
		Path: "golang.org/x/tools/cmd/goyacc",
		Desc: "Goyacc is a version of yacc for Go.",
	},
	{
		Path: "github.com/koron/gtc",
		Desc: "Go tools catalog",
	},
	{
		Path: "github.com/mattn/jvgrep",
		Desc: "grep for japanese vimmer",
	},
	{
		Path: "github.com/client9/misspell/cmd/misspell",
		Desc: "Correct commonly misspelled English words in source files",
	},
	{
		Path: "github.com/achiku/planter",
		Desc: "Generate PlantUML ER diagram textual description from PostgreSQL tables",
	},
	{
		Path: "github.com/go-swagger/go-swagger/cmd/swagger",
		Desc: "Swagger 2.0 implementation for go https://goswagger.io",
	},
	{
		Path: "honnef.co/go/tools/cmd/unused",
		Desc: "unused checks Go code for unused constants, variables, functions and types.",
	},
}

// DefaultCatalog provides a catalog of default tools.
var DefaultCatalog Catalog

func init() {
	catalogFile := os.Getenv("GTC_CATALOG_FILE")
	if catalogFile == "" {
		DefaultCatalog = NewCatalog(defaultTools...)
	} else {
		f, err := os.Open(catalogFile)
		if err != nil {
			log.Fatalf("cannot open catalog file: %v", err)
		}
		defer f.Close()

		var tools []Tool
		err = json.NewDecoder(f).Decode(&tools)
		if err != nil {
			log.Fatalf("cannot parse catalog file: %v", err)
		}
		DefaultCatalog = NewCatalog(tools...)
	}
}
