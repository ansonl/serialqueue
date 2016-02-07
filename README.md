# Serial Queue

A Serial Queue implementation for Go.
Tasks enqueued are executed in sequential order. Tasks are dequeued and run in a goroutine. 

...Like a blocking background thread.


Similar to [Apple GCD Serial Dispatch Queue](https://developer.apple.com/library/ios/documentation/General/Conceptual/ConcurrencyProgrammingGuide/OperationQueues/OperationQueues.html#//apple_ref/doc/uid/TP40008091-CH102-SW6). 

### Documentation
[![GoDoc](https://godoc.org/github.com/ansonl/serialqueue?status.svg)](https://godoc.org/github.com/ansonl/serialqueue)
