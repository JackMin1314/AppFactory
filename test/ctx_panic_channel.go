package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	const numtimes = 1000

	startT := time.Now()
	var wg sync.WaitGroup
	wg.Add(numtimes)

	for i := 0; i < numtimes; i++ {
		go func() {

			// 仅有这个可以查看request的panic，但是request内部起的goroutine里产生的panic其它goroutine无法捕获。
			defer func() {
				if p := recover(); p != nil {
					fmt.Println("get panic !!")
				}
			}()
			defer wg.Done()
			request(context.Background(), "test")
		}()

	}

	wg.Wait()

	endT := time.Since(startT)
	fmt.Printf("total time:%s\n", endT)
	time.Sleep(15 * time.Second)
	fmt.Printf("goroutine num:%d\n", runtime.NumGoroutine())
}

func request(ctx context.Context, req interface{}) error {

	ctx1, cancel := ShrinkDeadline(ctx, 2*time.Second)
	// ctx1, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	/*
		// 解决每个超时请求占用一个goroutine的bug，添加buffer size 1
		errchan := make(chan error,1) // 不管是否超时都能写入而不卡住goroutine
		go func() {
			errchan <- workjob(req)
		}()

		select {
		case myerr := <-errchan:
			return myerr

		case <-ctx.Done():
			return ctx.Err()
		}
	*/

	//
	errchan := make(chan error, 0)
	panicChan := make(chan interface{}, 1)
	// 新起了一个协程,如果panic则无法被捕捉
	go func() {
		// 捕捉panic
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		errchan <- workjob(req)
	}()

	select {
	case myerr := <-errchan:
		return myerr

	case <-ctx1.Done():
		<-errchan // 阻塞在这里，等待直到workjob执行完或返回结果.
		cancel()
		return ctx1.Err()
	case pc := <-panicChan:
		panic(pc) // request panic
	}

}

func workjob(value interface{}) error {
	time.Sleep(10 * time.Second)
	panic("evil intrigue")
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