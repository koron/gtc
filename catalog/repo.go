package catalog

var repo = map[string]Catalog{
	"dep": {
		Path: "github.com/golang/dep/cmd/dep",
		Desc: "Go dependency tool",
	},
	"errcheck": {
		Path: "github.com/kisielk/errcheck",
		Desc: "errcheck checks that you checked errors.",
	},
	"gocode": {
		Path: "github.com/nsf/gocode",
		Desc: "An autocompletion daemon for the Go programming language",
	},
	"gocyclo" : {
		Path: "github.com/fzipp/gocyclo",
		Desc: "Calculate cyclomatic complexities of functions in Go source code.",
	},
	"goimports": {
		Path: "golang.org/x/tools/cmd/goimports",
		Desc: "updates your Go import lines, adding missing ones and removing unreferenced ones.",
	},
	"golint": {
		Path: "github.com/golang/lint/golint",
		Desc: "a linter for Go source code",
	},
	"gorename": {
		Path: "golang.org/x/tools/cmd/gorename",
		Desc: "The gorename command performs precise type-safe renaming of identifiers in Go source code.",
	},
	"gotags": {
		Path: "github.com/jstemmer/gotags",
		Desc: "ctags-compatible tag generator for Go",
	},
	"goyacc": {
		Path: "golang.org/x/tools/cmd/goyacc",
		Desc: "Goyacc is a version of yacc for Go.",
	},
	"gtc": {
		Path: "github.com/koron/gtc",
		Desc: "Go tools catalog",
	},
	"jvgrep": {
		Path: "github.com/mattn/jvgrep",
		Desc: "grep for japanese vimmer",
	},
	"misspell": {
		Path: "github.com/client9/misspell/cmd/misspell",
		Desc: "Correct commonly misspelled English words in source files",
	},
	"planter": {
		Path: "github.com/achiku/planter",
		Desc: "Generate PlantUML ER diagram textual description from PostgreSQL tables",
	},
	"swagger": {
		Path: "github.com/go-swagger/go-swagger/cmd/swagger",
		Desc: "Swagger 2.0 implementation for go https://goswagger.io",
	},
	"unused": {
		Path: "honnef.co/go/tools/cmd/unused",
		Desc: "unused checks Go code for unused constants, variables, functions and types.",
	},
}
