package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/koron/gtc/catalog"
	"github.com/koron/gtc/goenv"
)

func install(args []string) error {
	if len(args) == 0 {
		return errors.New("no tools to install")
	}
	env := goenv.Default
	for _, prog := range args {
		err := installOne(env, prog)
		if err != nil {
			return err
		}
	}
	return nil
}

func installOne(env goenv.Env, prog string) error {
	c, ok := catalog.Find(prog)
	if !ok {
		return fmt.Errorf("unknown catalog %q", prog)
	}
	if env.HasTool(prog) {
		log.Printf("already installed: %s", prog)
		return nil
	}
	err := env.Install(c.Path)
	if err != nil {
		return err
	}
	return nil
}
