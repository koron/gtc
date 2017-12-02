package main

import (
	"fmt"
	"os"

	"github.com/koron/gtc/gtcore"
)

func run() error {
	if custom := os.Getenv("GTC_CATALOG_FILE"); custom != "" {
		err := gtcore.SetupDefaultCatalog(custom)
		if err != nil {
			return err
		}
	}
	return gtcore.Run(os.Args)
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
