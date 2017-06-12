package goenv

import (
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Env struct {
	RootDir   string
	ExeSuffix string
}

func New(root string) *Env {
	return &Env{
		RootDir: root,
	}
}

func (env *Env) HasTool(tool string) bool {
	name := filepath.Join(env.RootDir, "bin", tool+env.ExeSuffix)
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

func (env *Env) Update(path string) error {
	c := exec.Command("go", "get", "-v", "-u", path)
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}

func (env *Env) Tools() ([]string, error) {
	files, err := ioutil.ReadDir(filepath.Join(env.RootDir, "bin"))
	if err != nil {
		return nil, err
	}
	var tools []string
	for _, fi := range files {
		if !fi.Mode().IsRegular() {
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

func defaultEnv(bc build.Context) Env {
	gopath := filepath.SplitList(bc.GOPATH)
	if len(gopath) == 0 {
		panic("GOPATH isn't set")
	}
	var exeSuffix string
	if bc.GOOS == "windows" {
		exeSuffix = ".exe"
	}
	return Env{
		RootDir:   gopath[0],
		ExeSuffix: exeSuffix,
	}
}

func init() {
	Default = defaultEnv(build.Default)
}

var Default Env
