package gtcore

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/koron/gtc/goenv"
)

var (
	listFilter     string
	listShowPath   bool
	listShowDesc   bool
	listShowModule bool
	listShowLatest bool
)

func list(fs *flag.FlagSet, args []string) error {
	fs.StringVar(&listFilter, "filter", "installed",
		`filter by status: "installed", "notinstalled" or "unknown"`)
	fs.BoolVar(&listShowPath, "path", false, `show path of catalogs`)
	fs.BoolVar(&listShowDesc, "desc", false, `show desc of catalogs`)
	fs.BoolVar(&listShowModule, "mod", false, `show module path of catalogs`)
	fs.BoolVar(&listShowLatest, "latest", false, `show latest version`)
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
	fmt.Fprintf(w, "%s\n", tool.CmdName())
	if listShowPath {
		fmt.Fprintf(w, "    Path: %s\n", tool.Path)
	}
	if listShowDesc {
		fmt.Fprintf(w, "    Desc: %s\n", tool.Desc)
	}
	if listShowModule {
		fmt.Fprintf(w, "  Module: %s\n", tool.ModulePath())
	}
	if listShowLatest {
		info, err := tool.Latest(context.Background())
		var latest string
		if err != nil {
			latest = fmt.Sprintf("(error: %s)", err)
		} else {
			latest = info.Version
		}
		fmt.Fprintf(w, "  Latest: %s\n", latest)
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
