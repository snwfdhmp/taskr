# taskr ⏩ go forward

**Taskr keeps you going forward by applying strict no-regression rules on your repo.**

```bash
$ git commit -m "update api /home and /welcome"
taskr: regression: task 'welcomeSayHello' has returned error, did not on previous commit.
taskr: abort commit.
```

**It helps your team know exactly what has to be done, and how it will be rewarded.**

```bash
$ git commit -m "fixed linting issues"
taskr: this commit don't introduce any regression.
taskr: task lint done, reward: 5.
[master f263b5e] fixed linting issues
 3 files changed, 12 insertions(+), 1 deletion(-)
```

## Getting started

### Requirements

- [go](https://golang.org/doc/install) (easy install)

After you have go installed, **make sure the go binary directory is in your $PATH**. Go binary directory is located at *$(go env GOPATH)/bin*

If for any reason, you can't add GOPATH/bin to your $PATH, replace every call to  `taskr` by `$(go env GOPATH)/bin/taskr`

### Download

```bash
go get -u github.com/snwfdhmp/taskr/...
```

### Install

```bash
go install github.com/snwfdhmp/taskr
```

Now, in any git repo, run:

```bash
taskr init

```

### Let's start !

Let's create a new git repository.

```bash
$ mkdir myRepo
$ cd myRepo
$ git init
Initialized empty Git repository in /Users/snwfdhmp/myRepo/.git/
```

Now **init taskr** in this repository. This works the same way in an existing repository.

```bash
$ taskr init
taskr inited successfully.
```

Now let's see the *taskr.yaml* file the previous just created.

```yaml
- name: compile
  test: go build main.go
  reward: 0
- name: lint
  test: golint -set_exit_status
  reward: 0
```

By default, taskr will create an example *taskr.yaml* file for golang.
Each test is defined by its name, a command to run to determine if the test is pass or fail (based on exit status: 0 = pass, any other = fail), and a reward for the developer if the task is completed (will be used in upcoming versions).

Let's add a bunch of code in a *main.go*.

```go
package main

import (
	"fmt"
)

func SayHello() {
	fmt.Println("Hello world !")
}

func main() {
	SayHello()
	return
}
```

Now **commit** our changes.

```bash
$ git add .taskr taskr.yaml
$ git add main.go
$ git commit -m "taskr init"
taskr: task 'compile' completed.
[master (root-commit) 1eb4392] taskr init
 3 files changed, 20 insertions(+)
 create mode 100644 .taskr/history.yaml
 create mode 100644 main.go
 create mode 100644 taskr.yaml
```

We can see that the **task 'compile' has been completed** with this commit.

## Go forward

Taskr is built to keep going forward. **Regressions are automatically blocked** and taskr will abort any commit introducing one.

For example, let's add a mistake in our previous code :

```go
package main

import (
	"fmt"
)

func SayHello() {
	fmt.Println("Hello world !")
}

func main() {
	SayHe110[] //<- this won't compile
	return
}
```

Remember we previously completed the task 'compile'.

Let's try to commit this thing.

```bash
$ git add main.go
$ git commit -m "update main.go"
taskr: regression: task 'compile' has returned error, did not on previous commit.
taskr: abort commit.
```

Taskr isn't accepting our commit because the task 'compile' now fails.
We have to get the test of compile (it was `go build main.go`) to pass again.

## Complete tasks

**To complete tasks, you have to pass their tests successfully** (exit status 0).

For example, remember the default task 'lint':

```yaml
- name: lint
  test: golint -set_exit_status
  reward: 0
```

We have to get `golint -set_exit_status` to exit with status code 0.

Let's try

```bash
$ golint -set_exit_status
main.go:7:1: exported function SayHello should have comment or be unexported
Found 1 lint suggestions; failing.
```

What we have to do is fix this issue.

Let's change our *main.go*

```go
...

//SayHello prints "Hello world !" to the user
func SayHello() {
	fmt.Println("Hello world !")
}

...
```

Now retry the go linter.

```bash
$ golint -set_exit_status
```

Great ! Let's commit those changes.

```bash
$ git add main.go
$ git commit -m "commented func for linting"
taskr: task 'lint' completed.
[master f265f60] commented func for linting
 2 files changed, 1 insertion(+)
 create mode 100755 main
```

Now you know how to **complete tasks with taskr**.

## Golang package

Taskr is built upon the golang package **taskr**. It can be used by importing *github.com/snwfdhmp/taskr/pkg*

```go
import (
	"github.com/snwfdhmp/taskr/pkg"
)
```

The package name is **taskr**.

Then, for example

```go
history, err := taskr.OpenHistory() //error handling is omitted for readability purposes
tasks, err := taskr.ParseTasks()
report, err := history.Run(tasks...)

for _, t := range report.Tests {
	fmt.Println("Test", t.Name, "completed. Congratulations")
}
```

## Documentation

The package documentation can be found [here on godoc](https://godoc.org/github.com/snwfdhmp/taskr/pkg).

## Author

[snwfdhmp](https://github.com/snwfdhmp)
[Talk to me on LinkedIn](https://www.linkedin.com/in/martin-joly-951b8913b/)