package webbach

import (
	"context"
	"net/http"
)

type Batch struct {
	steps []Step
	ctx   context.Context
}

func (batch *Batch) AddStep(step Step) *Batch {
	batch.steps = append(batch.steps, step)
	return batch
}

// Executes every step in the batch queue.
// Does not execute subsequent steps to one that returns false.
func (batch *Batch) Execute(writer http.ResponseWriter, request *http.Request, steps ...Step) bool {
	batch.steps = append(batch.steps, steps...)

	for _, step := range batch.steps {
		if !step(writer, request, &batch.ctx) {
			return false
		}
	}

	return true
}

// Executes every step in the batch queue.
// Does not execute subsequent steps to one that returns false.
func Execute(writer http.ResponseWriter, request *http.Request, steps ...Step) bool {
	batch := Batch{
		steps: steps,
	}

	return batch.Execute(writer, request)
}
