package goenv

import (
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
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
