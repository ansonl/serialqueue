//Package serialqueue provides provides tests for github.com/ansonl/serialqueue. 
package serialqueue_test

import (
	"testing"
	"fmt"
	"sync"
	"github.com/ansonl/serialqueue"
	)

func TestQueue(t *testing.T) {
	//Create new serial queue
	queue := serialqueue.New()

	//Number of times to loop
	var loops = 100

	//Queue output will be added here
	var queueOutput string

	//WaitGroup for determining when queue output is finished
	var wg sync.WaitGroup

	//Create function to place in queue
	var printTest = func(i int) {
		queueOutput = fmt.Sprintf("%v%v",queueOutput, i)
		wg.Done()
	}

	//Add functions to queue
	var i int;
	for i=0;i<loops;i++ {
		wg.Add(1)
		t:=i
		queue <- func() {printTest(t)}
	}

	//Generate correct output
	var correctOutput string
	for i=0;i<loops;i++ {
		correctOutput = fmt.Sprintf("%v%v", correctOutput, i)		
	}

	//Wait until all functions in queue are done
	wg.Wait()

	//Compare queue output with correct output
	if queueOutput != correctOutput {
		t.Errorf("Serial Queue produced %v, want %v", queueOutput, correctOutput);
	}
}