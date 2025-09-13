package gtcore

import (
	"context"
	"log"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/koron-go/jsonhttpc"
)

// Tool represents a well known tool.
type Tool struct {
	Path string `json:"path"`
	Desc string `json:"desc"`

	// Name is command name (OPTION). If empty, extract from Path.
	Name string `json:"name,omitempty"`

	Module string `json:"module,omitempty"`
}

// CmdName returns name of a tool.
func (tool Tool) CmdName() string {
	if tool.Name != "" {
		return tool.Name
	}
	return path.Base(tool.Path)
}

// ModulePath returns module path
func (tool Tool) ModulePath() string {
	if tool.Module != "" {
		return tool.Module
	}
	x := strings.SplitN(tool.Path, "/", 4)
	if len(x) < 3 {
		log.Printf("[WARN] short module name, may be wrong: %s", tool.Path)
		return tool.Path
	}
	return strings.Join(x[:3], "/")
}

type Info struct {
	Version string
	Time    time.Time
}

var proxyClient *jsonhttpc.Client

func init() {
	u, err := url.Parse("https://proxy.golang.org/")
	if err != nil {
		panic(err)
	}
	proxyClient = jsonhttpc.New(u)
}

func (tool Tool) Latest(ctx context.Context) (*Info, error) {
	info := new(Info)
	err := proxyClient.Do(ctx, "GET", tool.ModulePath()+"/@latest", nil, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
