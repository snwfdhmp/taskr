/*
Package taskr is built to keep you going forward.
See README.md or read it at https://github.com/snwfdhmp/taskr
Documentation can be found at https://godoc.org/github.com/snwfdhmp/taskr/pkg
*/
package taskr

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"gopkg.in/src-d/go-git.v4"
)

var (
	fs      = afero.NewOsFs()
	gitRepo *git.Repository

	WD string = "./"
)

func appDir() string {
	return filepath.Join(WD, ".taskr")
}

func wordFor(b bool) string {
	if b == true {
		return "completed"
	}
	return "failed"
}

func mkPath(p string) string {
	return filepath.Join(WD, p)
}

// commit:
// 	test: true

func gitHead() (string, error) {
	head, err := gitRepo.Head()
	if err != nil {
		return "genesis", nil
	}

	return head.Hash().String(), nil
}

func MkInit() error {
	if err := fs.MkdirAll(appDir(), 0700); err != nil {
		return err
	}

	if _, err := fs.Create(historyPath()); err != nil {
		return err
	}

	return nil
}

func searchGitRepo() (repo *git.Repository) {
	var err error
	for abs := "./"; abs != "/" && err == nil; abs, err = filepath.Abs(filepath.Join(abs, "..")) {
		if repo, err = git.PlainOpen(abs); err == nil {
			WD = abs
			return
		}
	}
	return
}

func init() {
	gitRepo = searchGitRepo()
	if gitRepo == nil {
		wd, err := os.Getwd()
		if err == nil {
			fmt.Println("fatal: working dir:", wd)
		}
		fmt.Println("fatal: not a git repository (or any of the parent directories)")
		os.Exit(1)
	}
}
