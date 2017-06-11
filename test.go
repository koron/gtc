package main

import (
	"fmt"

	"github.com/koron/gtc/goenv"
)

func test(args []string) error {
	fmt.Printf("goenv.Default=%+v\n", goenv.Default)
	return nil
}
