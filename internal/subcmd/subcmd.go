package subcmd

import (
	"fmt"
	"strings"
)

type Subcmd struct {
	Name string
	Main Main
}

func (sc *Subcmd) run(args []string) error {
	if sc.Main == nil {
		return fmt.Errorf("no entry point")
	}
	return sc.Main(args)
}

type Main func([]string) error

func names(cmds []Subcmd) []string {
	if len(cmds) == 0 {
		return []string{"(none)"}
	}
	n := make([]string, 0, len(cmds))
	m := map[string]struct{}{}
	for _, c := range cmds {
		if c.Name == "" {
			continue
		}
		_, ok := m[c.Name]
		if ok {
			continue
		}
		m[c.Name] = struct{}{}
		n = append(n, c.Name)
	}
	return n
}

func Run(cmds []Subcmd, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("require one of sub-commands: %s",
			strings.Join(names(cmds), ", "))
	}
	n, args := args[0], args[1:]
	for _, c := range cmds {
		if c.Name == n {
			err := c.run(args)
			if err != nil {
				return fmt.Errorf("failed %q sub-command: %s", n, err)
			}
			return nil
		}
	}
	return fmt.Errorf("unknown %q sub-command, should be one of: %s",
		n, strings.Join(names(cmds), ", "))
}
