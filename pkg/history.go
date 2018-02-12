package taskr

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

//History represents a repo runs history
type History struct {
	States map[string]map[string]bool
	Tip    string
}

func historyPath() string {
	return filepath.Join(appDir(), "history.yaml")
}

// func (h *History) GetStates(commit string) (map[string]bool, bool) {
// 	s, ok := h.States[commit]
// 	return s, ok
// }

// func (h *History) GetLastState(testName string) (bool, bool) {
// 	s, ok := h.States[h.Tip][testName]
// 	return s, ok
// }
//

//OpenHistory opens history from the filesystem and returns it
func OpenHistory() (*History, error) {
	var h History

	file, err := afero.ReadFile(fs, historyPath()) //@todo config
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(file, &h); err != nil {
		return nil, err
	}

	return &h, nil
}

//Run runs a bunch of tasks
func (h *History) Run(tasks ...Task) (report RunReport, err error) {
	if gitRepo == nil {
		fmt.Println("fatal: git repo is nil")
		os.Exit(-1)
	}

	commit, err := gitHead()
	if err != nil {
		return
	}

	// if _, ok := h.States[commit]; ok {
	// 	err = fmt.Errorf("commit '%.6s' has already been tested", commit)
	// 	return
	// }

	for _, t := range tasks {
		var state bool
		if err := exec.Command(os.Getenv("SHELL"), "-c", fmt.Sprintf(t.Test)).Run(); err != nil {
			state = false
		} else {
			state = true
		}
		lastState, ok := h.States[h.Tip][t.Name]
		h.Add(commit, t.Name, state)
		if !ok {
			lastState = false
		}
		if lastState == state {
			continue
		}
		reward := t.Reward
		if lastState == true && state == false {
			err = fmt.Errorf("regression: task '%s' has returned error, did not on previous commit", t.Name)
			return
		}
		report.Add(t.Name, state, reward)
	}

	h.Tip = commit

	return
}

//Add adds a key/value pair record of a test (name:state) for a specific commit
func (h *History) Add(commit string, testName string, testValue bool) {
	if h.States == nil {
		h.States = make(map[string]map[string]bool)
	}
	if _, ok := h.States[commit]; !ok {
		h.States[commit] = make(map[string]bool)
	}

	h.States[commit][testName] = testValue
}

//Save saves history in filesystem
func (h *History) Save() error {
	if err := fs.MkdirAll(appDir(), 0700); err != nil {
		return err
	}

	out, err := fs.OpenFile(filepath.Join(appDir(), "history.yaml"), os.O_CREATE|os.O_WRONLY, 0700)
	if err != nil {
		return err
	}

	b, err := yaml.Marshal(h)
	if err != nil {
		return err
	}

	if _, err := out.Write(b); err != nil {
		return err
	}

	return nil
}
