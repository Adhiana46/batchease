# Basic Example

This is a Go program that demonstrates how to use the batchease library to create a batch processing system that can handle incoming data in batches.

The batchease library is imported from the github.com/Adhiana46/batchease package. It provides a simple way to create a batch processor that can process incoming data in batches, and allows you to configure the batch size, the number of workers, and the wait time between batches.

In this program, the batchease library is used to create a batch processor that can handle incoming integers. The program generates integers using the atomic package to increment the value of a variable i in a loop, and then adds the value to the batch processor using the AddItem method. Two separate loops generate the integers, and each one adds the integers to the batch processor.

The batch processor is configured to have a batch size of 500, 5 workers, and a wait time of 1000 milliseconds between batches. When a batch is processed, the program prints the number of messages in the batch, and the total of the last item in the batch. This is done by passing a function to the WithWorkers configuration option that prints the output.

The program runs the batch processor for 30 seconds, and then shuts it down using the Shutdown method.

## Files

- [main.go](main.go) - example source code, the **most interesting file for you**
- [go.mod](go.mod) - Go modules dependencies, you can find more information at [Go wiki](https://github.com/golang/go/wiki/Modules)
- [go.sum](go.sum) - Go modules checksums

## Requirements

To run this example you will need Golang installed. See the [installation guide](https://go.dev/doc/install).

```bash
> go run main.go
worker #1 	 | 500 messages sent, total sent 500 
worker #3 	 | 500 messages sent, total sent 1000 
worker #1 	 | 500 messages sent, total sent 1500 
worker #4 	 | 500 messages sent, total sent 2000 
worker #0 	 | 500 messages sent, total sent 2500 
worker #3 	 | 500 messages sent, total sent 3000
...
...
...
...
worker #4 	 | 500 messages sent, total sent 456632271 
worker #0 	 | 500 messages sent, total sent 456632771 
worker #3 	 | 500 messages sent, total sent 456633271 
worker #2 	 | 500 messages sent, total sent 456633771 
worker #4 	 | 500 messages sent, total sent 456634271
```