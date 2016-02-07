# Serial Queue
[![GoDoc](https://godoc.org/github.com/ansonl/serialqueue?status.svg)](https://godoc.org/github.com/ansonl/serialqueue)

> A blocking background thread for deterministic sequential operation.

A Serial Queue implementation for Go.
Tasks are enqueued in desired order of execution. Tasks are dequeued sequentially and run in a goroutine. 

Similar to [Apple GCD Serial Dispatch Queue](https://developer.apple.com/library/ios/documentation/General/Conceptual/ConcurrencyProgrammingGuide/OperationQueues/OperationQueues.html#//apple_ref/doc/uid/TP40008091-CH102-SW6). 

