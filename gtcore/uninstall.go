package gtcore

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/koron/gtc/goenv"
)

func uninstall(fs *flag.FlagSet, args []string) error {
	env := goenv.Default
	fs.BoolVar(&env.Verbose, "verbose", false, "verbose message")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() == 0 {
		return errors.New("no tools to uninstall")
	}

	failed := false
	for _, prog := range fs.Args() {
		err := uninstallOne(env, prog)
		if err != nil {
			failed = true
			log.Printf("failed to uninstall %s: %s", prog, err)
			continue
		}
	}
	if failed {
		return errors.New("failed to uninstall one or more tools")
	}
	return nil
}

func uninstallOne(env goenv.Env, prog string) error {
	tool, ok := catalogFind(prog)
	if !ok {
		return fmt.Errorf("unknown catalog %q", prog)
	}
	if !env.HasTool(prog) {
		log.Printf("not installed: %s", prog)
		return nil
	}
	err := env.Uninstall(tool.CmdName())
	if err != nil {
		return err
	}
	return nil
}
