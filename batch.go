package webbatch

import "net/http"

type BatchFn func(w http.ResponseWriter, r *http.Request) bool
type WebBatchRunner struct {
	jobs []BatchFn
}

func Exec(w http.ResponseWriter, r *http.Request, steps ...BatchFn) bool {
	for _, step := range steps {
		if !step(w, r) {
			return false
		}
	}

	return true
}

func (runner *WebBatchRunner) Add(fn BatchFn) *WebBatchRunner {
	runner.jobs = append(runner.jobs, fn)
	return runner
}

func (runner *WebBatchRunner) Run(w http.ResponseWriter, r *http.Request, steps []BatchFn) bool {
	return Exec(w, r, runner.jobs...)
}
