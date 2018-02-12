package taskr

import (
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

//Task represents a task object
type Task struct {
	Name   string
	Test   string
	Reward int
}

type TaskSet []Task

func (t *TaskSet) Save(path string) error {
	b, err := yaml.Marshal(&t)
	if err != nil {
		return err
	}

	if err := afero.WriteFile(fs, path, b, 0644); err != nil {
		return err
	}

	return nil
}

//ParseTasks parse tasks from a file and returns them
func ParseTasks(path string) (tasks TaskSet, err error) {
	file, err := afero.ReadFile(fs, path)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(file, &tasks); err != nil {
		return
	}

	return
}

func SampleTask(filename string) error {
	sampleTask := &TaskSet{{
		Name: "compile",
		Test: "go build main.go",
	},
		{
			Name: "lint",
			Test: "golint -set_exit_status",
		}}

	return sampleTask.Save(mkPath(filename + ".yaml"))
}
