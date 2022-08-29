package webbach

import (
	"net/http"
	"testing"
)

type FakeResponseWriter struct {
	header http.Header
}

func (w FakeResponseWriter) Header() http.Header {
	return w.header
}

func (w FakeResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (w FakeResponseWriter) WriteHeader(int) {

}

var steps []Step
var stepsResults []bool

func beforeEach() {
	steps = []Step{
		func(http.ResponseWriter, *http.Request) bool {
			stepsResults[0] = true
			return true
		},
		func(http.ResponseWriter, *http.Request) bool {
			stepsResults[1] = false
			return false
		},
		func(http.ResponseWriter, *http.Request) bool {
			stepsResults[2] = true
			return true
		},
		func(http.ResponseWriter, *http.Request) bool {
			stepsResults[3] = true
			return true
		},
		func(http.ResponseWriter, *http.Request) bool {
			stepsResults[4] = true
			return true
		},
	}

	stepsResults = []bool{
		false,
		false,
		false,
		false,
		false,
	}
}

var w FakeResponseWriter = FakeResponseWriter{}
var r http.Request = http.Request{
	Method: "GET",
}

func TestExecute(t *testing.T) {
	t.Run("Should not execute following steps if step 2 returns false", func(t *testing.T) {
		beforeEach()

		executionResult := Execute(w, &r, steps...)

		if executionResult {
			t.Fail()
		}

		// if step 3 was executed
		if stepsResults[2] {
			t.Fail()
		}
	})

	t.Run("Should execute following steps if step 2 returns true", func(t *testing.T) {
		beforeEach()
		steps[1] = func(http.ResponseWriter, *http.Request) bool {
			stepsResults[1] = true
			return true
		}

		executionResult := Execute(w, &r, steps...)

		if !executionResult {
			t.Fail()
		}

		// if step 3 was executed
		for _, result := range stepsResults {
			if !result {
				t.Fail()
			}
		}
	})
}

func TestBatchWithAddStep(t *testing.T) {
	t.Run("Should not execute following steps if step 2 returns false", func(t *testing.T) {
		beforeEach()

		batch := Batch{}
		for _, step := range steps {
			batch.AddStep(step)
		}

		executionResult := batch.Execute(w, &r)
		if executionResult {
			t.Fail()
		}

		// if step 3 was executed
		if stepsResults[2] {
			t.Fail()
		}
	})

	t.Run("Should execute following steps if step 2 returns true", func(t *testing.T) {
		beforeEach()
		steps[1] = func(http.ResponseWriter, *http.Request) bool {
			stepsResults[1] = true
			return true
		}

		batch := Batch{}
		for _, step := range steps {
			batch.AddStep(step)
		}

		executionResult := batch.Execute(w, &r)
		if !executionResult {
			t.Fail()
		}

		// if step 3 was executed
		for _, result := range stepsResults {
			if !result {
				t.Fail()
			}
		}
	})
}
