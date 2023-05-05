package batchease

import "errors"

var (
	ErrInvalidBatchSize      = errors.New("invalid batch size")
	ErrInvalidHandleFn       = errors.New("invalid handle function")
	ErrInvalidNumberOfWorker = errors.New("invalid number of workers")
)
