//Package serialqueue provides a serial queue for functions. 
//Queue items are processed in First In First Out (FIFO) order. 
package serialqueue

//New returns a new serial queue.
//Enqueue items like queueObj <- func() {doWork(data)}
func New() chan func() {
	//create channel of type function
	var queue = make(chan func())

	//spawn go routine to read and run functions in the channel
	go func() {
		for true {
			nextFunction := <-queue
			nextFunction()
		}
	}()

	return queue
}