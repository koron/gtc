package goenv

import (
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Env struct {
	RootDir   string
	ExeSuffix string
	IsWindows bool
}

func New(root string) *Env {
	return &Env{
		RootDir: root,
	}
}

func (env *Env) toolName(tool string) string {
	return filepath.Join(env.RootDir, "bin", tool+env.ExeSuffix)
}

func (env *Env) removeTool(tool string) error {
	n := env.toolName(tool)
	_, err := os.Stat(n)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	err = os.Remove(n)
	if err != nil {
		if env.IsWindows && os.IsPermission(err) {
			return os.Rename(n, n+"~")
		}
		return err
	}
	return nil
}

func (env *Env) touchTool(tool string) error {
	n := env.toolName(tool)
	now := time.Now()
	return os.Chtimes(n, now, now)
}

func (env *Env) HasTool(tool string) bool {
	name := env.toolName(tool)
	fi, err := os.Stat(name)
	if err != nil {
		return false
	}
	return fi.Mode().IsRegular()
}

func (env *Env) Install(path string) error {
	c := exec.Command("go", "get", path)
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

// Update update a tool.
func (env *Env) Update(path, tool string) error {
	c := exec.Command("go", "get", "-v", "-u", path)
	err := c.Run()
	if err != nil {
		return err
	}
	err = env.touchTool(tool)
	if err != nil {
		return err
	}
	return nil
}

func (env *Env) tools(filter func(os.FileInfo) bool) ([]string, error) {
	files, err := ioutil.ReadDir(filepath.Join(env.RootDir, "bin"))
	if err != nil {
		return nil, err
	}
	var tools []string
	for _, fi := range files {
		if !fi.Mode().IsRegular() || !filter(fi) {
			continue
		}
		n := fi.Name()
		if env.ExeSuffix != "" && strings.HasSuffix(n, env.ExeSuffix) {
			n = n[:len(n)-len(env.ExeSuffix)]
		}
		tools = append(tools, n)
	}
	return tools, nil
}

func (env *Env) Tools() ([]string, error) {
	return env.tools(func(os.FileInfo) bool { return true })
}

func (env *Env) OldTools(pivot time.Time) ([]string, error) {
	return env.tools(func(fi os.FileInfo) bool {
		return fi.ModTime().Before(pivot)
	})
}

func defaultEnv(bc build.Context) Env {
	gopath := filepath.SplitList(bc.GOPATH)
	if len(gopath) == 0 {
		panic("GOPATH isn't set")
	}
	var (
		exeSuffix string
		isWindows bool
	)
	if bc.GOOS == "windows" {
		exeSuffix = ".exe"
		isWindows = true
	}
	return Env{
		RootDir:   gopath[0],
		ExeSuffix: exeSuffix,
		IsWindows: isWindows,
	}
}

func init() {
	Default = defaultEnv(build.Default)
}

var Default Env
