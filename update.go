package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/koron/gtc/catalog"
	"github.com/koron/gtc/goenv"
)

func update(args []string) error {
	if len(args) == 0 {
		return errors.New("no tools to update")
	}
	env := goenv.Default
	for _, prog := range args {
		err := updateOne(env, prog)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateOne(env goenv.Env, prog string) error {
	c, ok := catalog.Find(prog)
	if !ok {
		return fmt.Errorf("unknown catalog %q", prog)
	}
	if !env.HasTool(prog) {
		log.Printf("not installed: %s", prog)
		return nil
	}
	err := env.Update(c.Path)
	if err != nil {
		return err
	}
	return nil
}
