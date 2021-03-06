package gtcore

var defaultTools = []Tool{
	{
		Path: "github.com/achiku/planter",
		Desc: "Generate PlantUML ER diagram textual description from PostgreSQL tables",
	},
	{
		Path: "github.com/client9/misspell/cmd/misspell",
		Desc: "Correct commonly misspelled English words in source files",
	},
	{
		Path: "github.com/derekparker/delve/cmd/dlv",
		Desc: "Delve is a debugger for the Go programming language.",
	},
	{
		Path: "github.com/fzipp/gocyclo",
		Desc: "Calculate cyclomatic complexities of functions in Go source code.",
	},
	{
		Path: "github.com/go-swagger/go-swagger/cmd/swagger",
		Desc: "Swagger 2.0 implementation for go https://goswagger.io",
	},
	{
		Path: "github.com/golang/protobuf/protoc-gen-go",
		Desc: "Go support for Google's protocol buffers",
	},
	{
		Path: "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway",
		Desc: "gRPC to JSON proxy generator following the gRPC HTTP spec - gateway",
	},
	{
		Path: "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger",
		Desc: "gRPC to JSON proxy generator following the gRPC HTTP spec - swagger",
	},
	{
		Path: "github.com/jstemmer/gotags",
		Desc: "ctags-compatible tag generator for Go",
	},
	{
		Path: "github.com/kisielk/errcheck",
		Desc: "errcheck checks that you checked errors.",
	},
	{
		Path: "github.com/koron/gtc",
		Desc: "Go tools catalog",
	},
	{
		Path: "github.com/koron/tmpl",
		Desc: "simple template expander",
	},
	{
		Path: "github.com/mattn/jvgrep",
		Desc: "grep for japanese vimmer",
	},
	{
		Path: "github.com/nsf/gocode",
		Desc: "An autocompletion daemon for the Go programming language",
	},
	{
		Path: "golang.org/x/lint/golint",
		Desc: "a linter for Go source code",
	},
	{
		Path: "golang.org/x/tools/cmd/goimports",
		Desc: "updates your Go import lines, adding missing ones and removing unreferenced ones.",
	},
	{
		Path: "golang.org/x/tools/cmd/gorename",
		Desc: "The gorename command performs precise type-safe renaming of identifiers in Go source code.",
	},
	{
		Path: "golang.org/x/tools/cmd/goyacc",
		Desc: "Goyacc is a version of yacc for Go.",
	},
	{
		Path: "golang.org/x/tools/gopls",
		Desc: "The gols command is an LSP server for Go.",
	},
	{
		Path: "honnef.co/go/tools/cmd/staticcheck",
		Desc: "staticcheck offers extensive analysis of Go code, covering a myriad of categories.",
	},
	{
		Path: "honnef.co/go/tools/cmd/unused",
		Desc: "unused checks Go code for unused constants, variables, functions and types (deprecated)",
	},
}
