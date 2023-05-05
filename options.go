package batchease

import (
	"strconv"
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

func WithWorkers[T any](n int, handleFn HandleFunc[T]) BatcherOption[T] {
	return func(b *Batcher[T]) error {
		if handleFn == nil {
			return ErrInvalidHandleFn
		}

		if n <= 0 {
			return ErrInvalidNumberOfWorker
		}

		b.workerSize = n
		b.workers = make([]*Worker[T], n)

		for i := 0; i < n; i++ {
			w, err := NewWorker[T](
				strconv.Itoa(i),
				WorkerWithHandleFn[T](handleFn),
			)
			if err != nil {
				return err
			}

			b.workers[i] = w
			if i == 0 {
				b.currentWorker = b.workers[i]
			}
		}

		return nil
	}
}

// WorkerOption
type WorkerOption[T any] func(w *Worker[T]) error

func WorkerWithHandleFn[T any](handleFn HandleFunc[T]) WorkerOption[T] {
	return func(w *Worker[T]) error {
		if handleFn == nil {
			return ErrInvalidHandleFn
		}

		w.handleFn = handleFn

		return nil
	}
}
