package main

import (
	"fmt"
	"time"
	"sync"
)

type AtomicInt struct{
	value int
	mutex sync.Mutex
}

func (a *AtomicInt) Increment(){
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.value++
}

func (a * AtomicInt) get() int{
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.value
}

func main() {
	var a AtomicInt
	a.Increment()
	go func() {
		a.Increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
