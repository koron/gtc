package gtcore

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/koron/gtc/goenv"
)

func install(fs *flag.FlagSet, args []string) error {
	env := goenv.Default
	fs.BoolVar(&env.EnableModule, "module", false, "use module to install")
	fs.BoolVar(&env.Verbose, "verbose", false, "verbose message")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() == 0 {
		return errors.New("no tools to install")
	}
	failed := false
	for _, prog := range fs.Args() {
		err := installOne(env, prog)
		if err != nil {
			failed = true
			log.Printf("failed to install %s: %s", prog, err)
			continue
		}
	}
	if failed {
		return errors.New("failed to install one or more tools")
	}
	return nil
}

func installOne(env goenv.Env, prog string) error {
	tool, ok := catalogFind(prog)
	if !ok {
		return fmt.Errorf("unknown catalog %q", prog)
	}
	if env.HasTool(prog) {
		log.Printf("already installed: %s", prog)
		return nil
	}
	err := env.Install(tool.Path)
	if err != nil {
		return err
	}
	return nil
}
