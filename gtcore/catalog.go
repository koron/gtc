package gtcore

import "sort"

type Catalog map[string]Tool

func NewCatalog(tools ...Tool) Catalog {
	c := Catalog{}
	c.Merge(tools...)
	return c
}

func (c Catalog) Merge(tools ...Tool) Catalog {
	for _, tool := range tools {
		c[tool.CmdName()] = tool
	}
	return c
}

func (c Catalog) Names() []string {
	names := make([]string, 0, len(c))
	for k, _ := range c {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}

func (c Catalog) Run(args []string) error {
	return run(c, args)
}
