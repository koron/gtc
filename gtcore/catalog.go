package gtcore

import (
	"encoding/json"
	"os"
	"sort"
)

// Catalog is a collection of tools.
type Catalog map[string]Tool

// NewCatalog creates a Catalog from Tool array.
func NewCatalog(tools ...Tool) Catalog {
	c := Catalog{}
	c.Merge(tools...)
	return c
}

// Merge merges tools to a Catalog.
func (c Catalog) Merge(tools ...Tool) Catalog {
	for _, tool := range tools {
		c[tool.CmdName()] = tool
	}
	return c
}

// Names returns all name of tools.
func (c Catalog) Names() []string {
	names := make([]string, 0, len(c))
	for k := range c {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}

// Run runs an operation on a Catalog.
func (c Catalog) Run(args []string) error {
	return run(c, args)
}

// DefaultCatalog provides a catalog of default tools.
var DefaultCatalog Catalog

func init() {
	SetupDefaultCatalog()
}

// SetupDefaultCatalog loads/setups DefaultCatalog with tools from JSON files.
func SetupDefaultCatalog(names ...string) error {
	cc := Catalog{}
	for _, name := range names {
		f, err := os.Open(name)
		if err != nil {
			return err
		}
		var tools []Tool
		err = json.NewDecoder(f).Decode(&tools)
		if err != nil {
			return err
		}
		cc.Merge(tools...)
	}
	DefaultCatalog = cc.
		Merge(defaultTools...).
		Merge(platformTools...)
	return nil
}
