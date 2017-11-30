package gtcore

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/koron/gtc/goenv"
)

func install(fs *flag.FlagSet, args []string) error {
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() == 0 {
		return errors.New("no tools to install")
	}
	env := goenv.Default
	for _, prog := range fs.Args() {
		err := installOne(env, prog)
		if err != nil {
			return err
		}
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
