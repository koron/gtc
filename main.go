package main

import (
	"fmt"
	"os"

	"github.com/koron/gtc/internal/subcmd"
)

var cmds = []subcmd.Subcmd{
	{
		Name: "install",
		Main: install,
	},
	{
		Name: "list",
		Main: list,
	},
	{
		Name: "update",
		Main: update,
	},
	{
		Name: "test",
		Main: test,
	},
}

func main() {
	err := subcmd.Run(cmds, os.Args[1:])
	if err != nil {
		fmt.Println(err.Error())
	}
}
