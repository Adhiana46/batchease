package batchease

// BatcherOption
type BatcherOption[T any] func(b *Batcher[T]) error

// WorkerOption
type WorkerOption[T any] func(w *Worker[T]) error
