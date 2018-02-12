package taskr

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/afero"
)

func addHook(name string) error {
	hookFile := fmt.Sprintf(".git/hooks/%s", name)
	hookCmd := fmt.Sprintf("if ! taskr %s; then exit 1; fi;", name)
	file, err := fs.OpenFile(mkPath(hookFile), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := afero.ReadFile(fs, hookFile)
	if err != nil || !strings.Contains(string(content), hookCmd) {
		if _, err = file.WriteString(hookCmd); err != nil {
			return err
		}
	}

	return nil
}

func AddPreCommitHook() error {
	return addHook("pre-commit")
}

func AddPostCommitHook() error {
	return addHook("post-commit")
}
