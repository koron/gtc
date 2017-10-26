package gtcore

import (
	"path"
)

type Tool struct {
	Path string `json:"path"`
	Desc string `json:"desc"`
}

func (c Tool) Name() string {
	return path.Base(c.Path)
}
