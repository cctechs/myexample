package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 传统的同步机制
// WaitGroup channel 实现
// Mutex   锁
// Cond    条件变量

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(
				rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func myworker() chan <- int{
	c := make(chan int)

	go func() {
		for r := range c{
			fmt.Println(r)
		}
	}()
	return c
}

func myFunc(){
	sum := 0
	sum++
	fmt.Println("sum:", sum)
}

func main() {
	{
		f := myFunc
		f()
		f()
		return
	}
	{
		w := myworker()
		go func() {
			for{
				time.Sleep(time.Millisecond*100)
				w <- rand.Int()
			}
		}()
		time.Sleep(time.Second)
		return
	}
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int

	chanTime := time.After(10 * time.Second)
	chanTick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout, queue:", len(values))
		case <- chanTick:
			fmt.Println("queue:", len(values))
		case <-chanTime:
			fmt.Println("bye")
			return
		}
	}

}
