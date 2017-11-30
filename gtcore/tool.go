package gtcore

import (
	"path"
)

type Tool struct {
	Path string `json:"path"`
	Desc string `json:"desc"`

	Name string `json:"name,omitempty` // command name (OPTION). extract from Path if empty.
}

func (tool Tool) CmdName() string {
	if tool.Name != "" {
		return tool.Name
	}
	return path.Base(tool.Path)
}
