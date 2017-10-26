package gtcore

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
	updateDryrun   bool
	updateFlagAll  bool
	updateDuration time.Duration
)

func update(fs *flag.FlagSet, args []string) error {
	fs.BoolVar(&updateDryrun, "dryrun", false, "dry run to update")
	fs.BoolVar(&updateFlagAll, "all", false, "update all installed tools")
	fs.DurationVar(&updateDuration, "duration", time.Hour*24*5,
		"threshold to update with \"-all\"")
	if err := fs.Parse(args); err != nil {
		return err
	}

	env := goenv.Default
	if updateFlagAll {
		return updateAll(&env, updateDuration)
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
