package gtcore

import (
	"flag"
	"fmt"

	"github.com/koron/gtc/goenv"
)

func test(fs *flag.FlagSet, args []string) error {
	if err := fs.Parse(args); err != nil {
		return err
	}
	fmt.Printf("goenv.Default=%+v\n", goenv.Default)
	return nil
}
