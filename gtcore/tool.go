package gtcore

import (
	"path"
)

type Tool struct {
	Path string `json:"path"`
	Desc string `json:"desc"`

	// Name is command name (OPTION). If empty, extract from Path.
	Name string `json:"name,omitempty`
}

func (tool Tool) CmdName() string {
	if tool.Name != "" {
		return tool.Name
	}
	return path.Base(tool.Path)
}
