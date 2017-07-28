package catalog

import (
	"path"
	"sort"
)

type Catalog struct {
	Path string
	Desc string
}

func (c Catalog) Name() string {
	return path.Base(c.Path)
}

func Find(name string) (Catalog, bool) {
	c, ok := repo[name]
	return c, ok
}

func Names() []string {
	names := make([]string, 0, len(repo))
	for k, _ := range repo {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}
