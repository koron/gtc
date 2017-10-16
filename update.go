package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/koron/gtc/catalog"
	"github.com/koron/gtc/goenv"
)

var (
	updateDryrun bool
)

func update(args []string) error {
	var (
		all bool
		dur time.Duration
		fs  = flag.NewFlagSet(`"gtc update"`, flag.ExitOnError)
	)
	fs.BoolVar(&updateDryrun, "dryrun", false, "dry run to update")
	fs.BoolVar(&all, "all", false, "update all installed tools")
	fs.DurationVar(&dur, "duration", time.Hour*24*5,
		"threshold to update with \"-all\"")
	err := fs.Parse(args)
	if err != nil {
		return err
	}
	env := goenv.Default
	if all {
		return updateAll(&env, dur)
	}
	return updateTools(&env, fs.Args())
}

func updateAll(env *goenv.Env, dur time.Duration) error {
	tools, err := env.OldTools(time.Now().Add(-dur))
	if err != nil {
		return err
	}
	var all []string
	for _, t := range tools {
		if _, ok := catalog.Find(t); ok {
			all = append(all, t)
		}
	}
	if len(all) == 0 {
		log.Printf("no tools to update")
		return nil
	}
	return updateTools(env, all)
}

func updateTools(env *goenv.Env, tools []string) error {
	if len(tools) == 0 {
		return errors.New("no tools to update")
	}
	for _, prog := range tools {
		err := updateOne(env, prog)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateOne(env *goenv.Env, prog string) error {
	c, ok := catalog.Find(prog)
	if !ok {
		return fmt.Errorf("unknown catalog %q", prog)
	}
	if !env.HasTool(prog) {
		log.Printf("not installed: %s", prog)
		return nil
	}
	if updateDryrun {
		log.Printf("updating (dryrun): %s", prog)
		return nil
	}
	log.Printf("updating: %s", prog)
	err := env.Update(c.Path, prog)
	if err != nil {
		return err
	}
	return nil
}
