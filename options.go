package batchease

import (
	"time"
)

// BatcherOption
type BatcherOption[T any] func(b *Batcher[T]) error

func WithDefaultConfig[T any]() BatcherOption[T] {
	return func(b *Batcher[T]) error {
		b.size = 5
		b.wait = time.Second
		b.loadBalancer = &RoundRobinLB[T]{}

		return nil
	}
}

func WithSize[T any](size int) BatcherOption[T] {
	return func(b *Batcher[T]) error {
		if size <= 0 {
			return ErrInvalidBatchSize
		}

		b.size = size

		return nil
	}
}

func WithWait[T any](wait time.Duration) BatcherOption[T] {
	return func(b *Batcher[T]) error {
		b.wait = wait

		return nil
	}
}

// WorkerOption
type WorkerOption[T any] func(w *Worker[T]) error
