package main

import (
	"fmt"
	"sync"
)
// CSP
// Communication Sequential Process
// 不要通过共享内存来通信，通过通信来共享内存

type worker struct {
	in chan int
	done func()
}

// channel 是goroutine与goroutine之间的交互
func createWorker(id int, wg* sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func doWorker(id int, w worker) {
	for r := range w.in {
		fmt.Printf("work %d received %c\n", id, r)
		go func() {
			w.done()
		}()
	}
}

func chanDemo() {
	//var c chan int // c == nil
	var wg sync.WaitGroup
	var works [10]worker
	for i := 0; i < 10; i++ {
		works[i] = createWorker(i, &wg)
	}

	for i, w := range works{
		wg.Add(1)
		w.in <- 'a' + i

	}

	for i, w := range works{
		wg.Add(1)
		w.in <- 'A' + i

	}

	wg.Wait()
}

func main() {

	chanDemo()
	//bufferedChannel()
	//channelClose()
	//time.Sleep(time.Second)
}
