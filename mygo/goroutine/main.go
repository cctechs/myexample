package main

import (
	"time"
	"fmt"
)


// coroutine
// 轻量级线程
// 非抢占式多任务处理，由协程主动交出控制权
// 编译器/解释器/虚拟机层面的多任务，不是操作系统层面的多任务
// 多个协程可能在一个或多个线程中运行

// 子程序是协程的一个特例

// goroutine 可能的切换点
// I/O select   channel 等待锁  函数调用 runtime.Gosched()

func main() {
	//var a [10]int
	for i := 0; i < 10; i++ {
		go func(v int) {
			for {
				fmt.Printf("goroutine %d\n", v)
				//a[v]++
				//runtime.Gosched() // 手动交出控制权，让别的协程有机会运行
			}
		}(i)
	}
	time.Sleep(time.Minute)
	//fmt.Println(a)
}
