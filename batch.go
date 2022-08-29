package webbach

import "net/http"

type Batch struct {
	steps []Step
}

func (batch *Batch) AddStep(step Step) *Batch {
	batch.steps = append(batch.steps, step)
	return batch
}

func (batch *Batch) Execute(writer http.ResponseWriter, request *http.Request, steps ...Step) bool {
	batch.steps = append(batch.steps, steps...)

	for _, step := range batch.steps {
		if !step(writer, request) {
			return false
		}
	}

	return true
}
