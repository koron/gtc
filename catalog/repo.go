package catalog

var repo = map[string]Catalog{
	"errcheck": {
		Path: "github.com/kisielk/errcheck",
		Desc: "errcheck checks that you checked errors.",
	},
	"gocode": {
		Path: "github.com/nsf/gocode",
		Desc: "An autocompletion daemon for the Go programming language",
	},
	"goimports": {
		Path: "golang.org/x/tools/cmd/goimports",
		Desc: "updates your Go import lines, adding missing ones and removing unreferenced ones.",
	},
	"golint": {
		Path: "github.com/golang/lint/golint",
		Desc: "a linter for Go source code",
	},
	"gotags": {
		Path: "github.com/jstemmer/gotags",
		Desc: "ctags-compatible tag generator for Go",
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
}
