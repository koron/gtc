package main

import (
	"flag"
	"fmt"

	"github.com/koron/gtc/catalog"
	"github.com/koron/gtc/goenv"
)

func list(args []string) error {
	var (
		filter string
		fs  = flag.NewFlagSet(`"gtc list"`, flag.ExitOnError)
	)
	fs.StringVar(&filter, "filter", "installed",
		`filter by status: "installed", "notinstalled" or "unknown"`)
	err := fs.Parse(args)
	if err != nil {
		return err
	}

	env := goenv.Default
	switch filter {
	case "installed":
		return listInstalled(env)
	case "notinstalled":
		return listNotInstalled(env)
	case "unknown":
		return listUnknown(env)
	default:
		return fmt.Errorf("unsupported filter: %s", filter)
	}
}

func listNotInstalled(env goenv.Env) error {
	for _, prog := range catalog.Names() {
		if env.HasTool(prog) {
			continue
		}
		c, _ := catalog.Find(prog)
		fmt.Printf("%s\t%s\n", prog, c.Desc)
	}
	return nil
}

func listInstalled(env goenv.Env) error {
	for _, prog := range catalog.Names() {
		if !env.HasTool(prog) {
			continue
		}
		fmt.Println(prog)
	}
	return nil
}

func listUnknown(env goenv.Env) error {
	tools, err := env.Tools()
	if err != nil {
		return err
	}
	for _, t := range tools {
		_, ok := catalog.Find(t)
		if ok {
			continue
		}
		fmt.Println(t)
	}
	return nil
}
