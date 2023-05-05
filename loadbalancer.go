package batchease

type LoadBalancer[T any] interface {
	Resolve(workers []*Worker[T]) *Worker[T]
}

type RoundRobinLB[T any] struct {
	// current index
	i int32
}

func (lb *RoundRobinLB[T]) Resolve(workers []*Worker[T]) *Worker[T] {
	var n int32 = int32(len(workers))

	if n <= 1 {
		return workers[0]
	}

	lb.i = (lb.i + 1) % n

	return workers[lb.i]
}
