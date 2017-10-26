package gtcore

import (
	"path"
)

type Tool struct {
	Path string
	Desc string
}

func (c Tool) Name() string {
	return path.Base(c.Path)
}
