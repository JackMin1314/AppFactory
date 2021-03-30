package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	const numtimes = 10

	startT := time.Now()
	var wg sync.WaitGroup
	wg.Add(numtimes)

	// 进行处理
	// 模拟100个超时并发请求,需要等待全部有返回
	for i := 0; i < numtimes; i++ {
		go func() {

			// 仅有这个可以查看request的panic，但是request内部起的goroutine里产生的panic其它goroutine无法捕获。
			defer func() {
				if p := recover(); p != nil {
					fmt.Println("get panic !!")
				}
			}()
			defer wg.Done()
			requestwork(context.Background(), "test")
		}()

	}

	wg.Wait()

	endT := time.Since(startT)
	fmt.Printf("total time:%s\n", endT)
	// 等待所有执行完
	time.Sleep(15 * time.Second)
	fmt.Printf("goroutine num:%d\n", runtime.NumGoroutine())
}

func requestwork(ctx context.Context, req interface{}) error {

	// 比对传入的ctx时间和超时时间
	ctx1, cancel := ShrinkDeadline(ctx, 2*time.Second)
	// ctx1, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	chanErr := make(chan error, 1) // 有缓冲， done <- dowork 不管是否超时都能写入而不卡住goroutine
	/**
	* 一旦requestwork超时立马返回error
	select{
	case e := <- chanErr:
		return e
	case <- ctx1.Done():
		return ctx1.Err()
	}
	**/

	/**
	* 当requestwork 超时时不立刻返回，等待dowork执行完
	select {
	case e := <-chanErr:
		return e
	case <-ctx1.Done():
		// 一直等待 dowork 执行完有结果
		<-chanErr
		cancel()
		return ctx1.Err()
	}
	**/
	panicChan := make(chan interface{}, 1)
	// 新起了一个协程,如果panic则无法被捕捉
	// 启一个协程调用超时处理函数
	go func() {
		// 捕捉panic
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		chanErr <- dowork(req)
	}()

	select {
	case myerr := <-chanErr:
		return myerr

	case <-ctx1.Done():
		// <-chanErr // 阻塞在这里，等待直到dowork返回结果.
		// cancel()
		return ctx1.Err()
	case pc := <-panicChan:
		panic(pc) // request panic 模拟
	}

}

func dowork(value interface{}) error {
	time.Sleep(10 * time.Second)
	// panic("evil intrigue")
	return nil
}

// 超时小于给定的时间，则用超时的时间创建
func ShrinkDeadline(ctx context.Context, timeout time.Duration) (context.Context, func()) {
	if deadline, ok := ctx.Deadline(); ok {
		leftTime := time.Until(deadline)
		if leftTime < timeout {
			timeout = leftTime
		}
	}

	return context.WithDeadline(ctx, time.Now().Add(timeout))
}
