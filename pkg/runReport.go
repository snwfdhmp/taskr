package taskr

import (
	"fmt"
)

//RunReport represents a run output
type RunReport struct {
	Tests map[string]testReport
}

type testReport struct {
	State  bool
	Reward int
}

//NewRunReport returns a new RunReport
func NewRunReport() *RunReport {
	return &RunReport{make(map[string]testReport)}
}

//Add adds a record to a run report, containing new test's state, and associated reward
func (r *RunReport) Add(testName string, state bool, reward int) {
	if r.Tests == nil {
		r.Tests = make(map[string]testReport)
	}

	r.Tests[testName] = testReport{state, reward}
}

//Print prints a RunReport with fmt
func (r *RunReport) Print() {
	if len(r.Tests) < 1 {
		return
	}
	total := 0
	for n, t := range r.Tests {
		fmt.Printf("taskr: task '%s' %s.\n", n, wordFor(t.State))
		if t.Reward > 0 {
			fmt.Printf("reward: %d", t.Reward)
		}
		total += t.Reward
	}
	if total > 0 {
		fmt.Println("total rewards:", total)
	}

}
