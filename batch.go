// The package webbatch provides a way to executed multiple treatment in a sequential order.
package webbatch

import "net/http"

// A function that is capable of treating a request and sending a response if needed
type BatchFn func(w http.ResponseWriter, r *http.Request) bool

// A container for steps to be run in queue
type WebBatchRunner struct {
	jobs []BatchFn
}

// Exec executes an infinite number of steps, functions that take a http.ResponseWriter and an http.Request
// each of these steps return a boolean that indicate if we should continue the treatment down to the next
// step in the queue.
func Exec(w http.ResponseWriter, r *http.Request, steps ...BatchFn) bool {
	for _, step := range steps {
		if !step(w, r) {
			return false
		}
	}

	return true
}

// Adds a new step to the queue
func (runner *WebBatchRunner) Add(fn BatchFn) *WebBatchRunner {
	runner.jobs = append(runner.jobs, fn)
	return runner
}

// Runs all the steps in the queue. Will stop when at least one of the steps returns true or when
// there are no more steps to execute.
func (runner *WebBatchRunner) Run(w http.ResponseWriter, r *http.Request, steps []BatchFn) bool {
	return Exec(w, r, runner.jobs...)
}
