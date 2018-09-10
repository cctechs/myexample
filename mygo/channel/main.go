package main

import (
	"fmt"
	"time"
)
// CSP
// Communication Sequential Process
// 不要通过共享内存来通信，通过通信来共享内存

// channel 是goroutine与goroutine之间的交互
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	/*
		for {
			if n, ok := <-c; !ok {
				break
			} else {
				fmt.Printf("worker %d received %v\n", id, n)
			}
		}
	*/
	//for{
	for r := range c {
		fmt.Printf("work %d received %c\n", id, r)
	}
	//}

}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'A'
	c <- 'B'
	c <- 'C'
	c <- 'D'
	time.Sleep(time.Millisecond)
}

func chanDemo() {
	//var c chan int // c == nil
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'A'
	c <- 'B'
	time.Sleep(time.Second)
	c <- 'C'
	//close(c)
}

func main() {

	chanDemo()
	//bufferedChannel()
	//channelClose()
	//time.Sleep(time.Second)
}
