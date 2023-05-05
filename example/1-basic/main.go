package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/Adhiana46/batchease"
)

func main() {
	totalMessages := atomic.Int64{}

	batcher, err := batchease.New(
		batchease.WithDefaultConfig[int](),
		batchease.WithSize[int](500),
		batchease.WithWorkers(5, func(worker *batchease.Worker[int], data []int) {
			newTotalMessages := totalMessages.Add(int64(len(data)))

			fmt.Printf("worker #%s \t | %d messages sent, total sent %d \n", worker.ID, len(data), newTotalMessages)
		}),
		batchease.WithWait[int](1000*time.Millisecond),
	)
	if err != nil {
		panic(err)
	}

	var i int64 = 0
	go func() {
		for {
			batcher.AddItem(int(atomic.AddInt64(&i, 1)))
			// time.Sleep(time.Millisecond)
		}
	}()
	go func() {
		for {
			batcher.AddItem(int(atomic.AddInt64(&i, 1)))
			// time.Sleep(time.Millisecond)
		}
	}()

	go batcher.Run()

	time.Sleep(30 * time.Second)
	batcher.Shutdown()
}
