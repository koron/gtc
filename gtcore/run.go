package gtcore

import (
	"errors"
	"path/filepath"

	subcmd "github.com/koron/go-subcmd"
)

var cmds = subcmd.Subcmds{
	"install": subcmd.Main2(install),
	"list":    subcmd.Main2(list),
	"update":  subcmd.Main2(update),
	"test":    subcmd.Main2(test),
}

var currentCatalog Catalog

func run(c Catalog, args[]string) error {
	if len(args) < 1 {
		return errors.New("one ore more argurements required")
	}
	name := filepath.Base(args[0])
	name = name[:len(name)-len(filepath.Ext(name))]
	currentCatalog = c
	return cmds.RunWithName(name, args[1:])
}

// Run runs the core of "gtc" command.
func Run(args []string) error {
	return run(DefaultCatalog, args)
}
