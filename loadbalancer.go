package batchease

type LoadBalancer[T any] interface {
	Resolve(workers []*Worker[T]) *Worker[T]
}
