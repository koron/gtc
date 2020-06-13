package goenv

import (
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Env is information of gtc running environment.
type Env struct {
	// RootDir is first element of GOPATH
	RootDir string

	// ExeSuffix is extension of executable file name for the OS.
	// ".exe" for Windows, "" for other OS.
	ExeSuffix string

	// IsWindows is true for Windows.
	IsWindows bool

	EnableModule bool

	Verbose bool
}

// New creates `*Env`
func New(root string) *Env {
	return &Env{
		RootDir: root,
	}
}

func (env *Env) toolName(tool string) string {
	return filepath.Join(env.RootDir, "bin", tool+env.ExeSuffix)
}

// Uninstall uninstalls a tool from "$GOPATH/bin".
func (env *Env) Uninstall(tool string) error {
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
		// on Windows, running command can't be removed. So try to rename with
		// "~" suffix instead.
		if env.IsWindows && os.IsPermission(err) {
			err := os.Rename(n, n+"~")
			if err != nil {
				return err
			}
			if env.Verbose {
				log.Printf("uninstalled with rename: %s", tool)
			}
			return nil
		}
		return err
	}
	log.Printf("uninstalled: %s", tool)
	return nil
}

func (env *Env) touchTool(tool string) error {
	n := env.toolName(tool)
	now := time.Now()
	return os.Chtimes(n, now, now)
}

// HasTool checks a specified tool is installed or not.
func (env *Env) HasTool(tool string) bool {
	name := env.toolName(tool)
	fi, err := os.Stat(name)
	if err != nil {
		return false
	}
	return fi.Mode().IsRegular()
}

func (env *Env) moduleEnvstr() string {
	if env.EnableModule {
		return "GO111MODULE=on"
	}
	return "GO111MODULE=off"
}

// Install installs a tool.
func (env *Env) Install(path string) error {
	args := make([]string, 0, 3)
	args = append(args, "get")
	if env.Verbose {
		args = append(args, "-v")
	}
	args = append(args, path)
	c := exec.Command("go", args...)
	c.Env = os.Environ()
	c.Env = append(c.Env, env.moduleEnvstr())

	b, err := c.CombinedOutput()
	if err != nil {
		os.Stderr.Write(b)
		return err
	}
	return nil
}

// Update update a tool.
func (env *Env) Update(path, tool string) error {
	args := make([]string, 0, 4)
	args = append(args, "get", "-u")
	if env.Verbose {
		args = append(args, "-v")
	}
	args = append(args, path)
	c := exec.Command("go", args...)
	c.Env = os.Environ()
	c.Env = append(c.Env, env.moduleEnvstr())

	b, err := c.CombinedOutput()
	if err != nil {
		os.Stderr.Write(b)
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

// Tools returns all installed tools.
func (env *Env) Tools() ([]string, error) {
	return env.tools(func(os.FileInfo) bool { return true })
}

// OldTools returns list of tools, which are not updated recently.
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

// Default is default `Env` for current running environment.
var Default Env
