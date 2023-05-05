package batchease

import "sync"

type Worker[T any] struct {
	ID       string
	handleFn HandleFunc[T]
	mutex    *sync.Mutex
}

func NewWorker[T any](id string, opts ...WorkerOption[T]) (*Worker[T], error) {
	instance := Worker[T]{
		ID:    id,
		mutex: &sync.Mutex{},
	}

	for _, opt := range opts {
		if err := opt(&instance); err != nil {
			return &Worker[T]{}, err
		}
	}

	if instance.handleFn == nil {
		return &Worker[T]{}, ErrInvalidHandleFn
	}

	return &instance, nil
}

func (w *Worker[T]) do(data []T) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.handleFn(w, data)
}
