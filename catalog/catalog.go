package catalog

import "sort"

type Catalog struct {
	Path string
	Desc string
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
