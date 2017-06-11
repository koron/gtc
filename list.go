package main

import (
	"flag"
	"fmt"

	"github.com/koron/gtc/catalog"
	"github.com/koron/gtc/goenv"
)

func list(args []string) error {
	var (
		not bool
		fs  = flag.NewFlagSet(`"gtc list"`, flag.ExitOnError)
	)
	fs.BoolVar(&not, "not", false, "list not-installed tools")
	err := fs.Parse(args)
	if err != nil {
		return err
	}

	env := goenv.Default
	if not {
		return listNotInstalled(env)
	}
	return listInstalled(env)
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
