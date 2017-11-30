package gtcore

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/koron/gtc/goenv"
)

var (
	listFilter   string
	listShowPath bool
	listShowDesc bool
)

func list(fs *flag.FlagSet, args []string) error {
	fs.StringVar(&listFilter, "filter", "installed",
		`filter by status: "installed", "notinstalled" or "unknown"`)
	fs.BoolVar(&listShowPath, "path", false, `show path of catalogs`)
	fs.BoolVar(&listShowDesc, "desc", false, `show desc of catalogs`)
	if err := fs.Parse(args); err != nil {
		return err
	}

	env := goenv.Default
	switch listFilter {
	case "installed":
		return listInstalled(env)
	case "notinstalled":
		return listNotInstalled(env)
	case "unknown":
		return listUnknown(env)
	default:
		return fmt.Errorf("unsupported filter: %s", listFilter)
	}
}

func listPrint(w io.Writer, tool Tool) {
	w.Write([]byte(tool.CmdName()))
	w.Write([]byte("\n"))
	if listShowPath {
		w.Write([]byte("  "))
		w.Write([]byte(tool.Path))
		w.Write([]byte("\n"))
	}
	if listShowDesc {
		w.Write([]byte("  "))
		w.Write([]byte(tool.Desc))
		w.Write([]byte("\n"))
	}
}

func listNotInstalled(env goenv.Env) error {
	for _, prog := range catalogNames() {
		if env.HasTool(prog) {
			continue
		}
		tool, _ := catalogFind(prog)
		listPrint(os.Stdout, tool)
	}
	return nil
}

func listInstalled(env goenv.Env) error {
	for _, prog := range catalogNames() {
		if !env.HasTool(prog) {
			continue
		}
		tool, _ := catalogFind(prog)
		listPrint(os.Stdout, tool)
	}
	return nil
}

func listUnknown(env goenv.Env) error {
	tools, err := env.Tools()
	if err != nil {
		return err
	}
	for _, t := range tools {
		_, ok := catalogFind(t)
		if ok {
			continue
		}
		fmt.Println(t)
	}
	return nil
}
