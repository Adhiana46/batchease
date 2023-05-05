package batchease

import (
	"sync"
	"time"
)

type HandleFunc[T any] func(worker *Worker[T], data []T)

type Batcher[T any] struct {
	// Size per batch
	size int
	// wait time before items processed
	wait time.Duration
	// items to be processed
	items []T

	// slice of available workers
	workers []*Worker[T]
	// maximum workers size the batcher has
	workerSize int
	// current active worker, selected by load balancer
	currentWorker *Worker[T]
	// load balancer, to determine who process the items
	loadBalancer LoadBalancer[T]

	mutex *sync.Mutex
}

func New[T any](opts ...BatcherOption[T]) (*Batcher[T], error) {
	instance := Batcher[T]{
		items: []T{},
		mutex: &sync.Mutex{},
	}

	if len(opts) == 0 {
		opts = append(opts, WithDefaultConfig[T]())
	}

	for _, opt := range opts {
		if err := opt(&instance); err != nil {
			return &Batcher[T]{}, err
		}
	}

	return &instance, nil
}

func (b *Batcher[T]) AddWorker(w *Worker[T]) error {
	b.workers = append(b.workers, w)
	b.workerSize++

	return nil
}

func (b *Batcher[T]) AddItem(item T) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// TODO: middlewares

	b.items = append(b.items, item)

	if len(b.items) >= b.size {
		b.doWork()
	}
}

func (b *Batcher[T]) SetHandleFn(handleFn HandleFunc[T]) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for _, worker := range b.workers {
		worker.handleFn = handleFn
	}
}

func (b *Batcher[T]) Run() {
	for range time.Tick(b.wait) {
		b.mutex.Lock()
		b.doWork()
		b.mutex.Unlock()
	}
}

func (b *Batcher[T]) Shutdown() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.doWork()
}

func (b *Batcher[T]) doWork() {
	if len(b.items) == 0 {
		return
	}

	go func(w Worker[T], data []T) {
		w.do(data)
	}(*b.currentWorker, b.items)

	b.currentWorker = b.loadBalancer.Resolve(b.workers)

	b.items = []T{}
}
