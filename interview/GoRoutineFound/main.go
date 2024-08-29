package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	nums := []int{23, 32, 78, 43, 76, 65, 345, 762, 915, 86}
	target := 5

	goroutineNum := 4
	size := len(nums) / goroutineNum

	var wg sync.WaitGroup
	result := make(chan bool)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < goroutineNum; i++ {
		start := i * size
		end := start + size
		if i == len(nums)-1 {
			end = len(nums)
		}

		wg.Add(1)
		go find(ctx, nums[start:end], target, result, &wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	found := false
	timeout := false

	// 处理结果
	for {
		select {
		case result, ok := <-result:
			if !ok {
				// channel 关闭
				if found {
					fmt.Println("Found it!")
				} else if timeout {
					fmt.Println("Timeout! Not Found")
				}
				return
			}
			if result {
				found = true
				cancel() // 取消所有 goroutine
			}
		case <-ctx.Done():
			timeout = true
			cancel() // 确保所有 goroutine 都被取消
		}
	}

}

func find(ctx context.Context, arr []int, target int, result chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range arr {
		select {
		case <-ctx.Done():
			return
		default:
			if v == target {
				result <- true
				return
			}
		}
	}

	time.Sleep(5 * time.Second)
	result <- false
}
