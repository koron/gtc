package main

import (
	"fmt"
	"os"

	"github.com/koron/gtc/gtcore"
)

func main() {
	err := gtcore.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}
