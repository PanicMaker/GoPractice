package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 场景：在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。
// 每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100

type Ban struct {
	visitIPs map[string]time.Time
	lock     sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	b := &Ban{visitIPs: make(map[string]time.Time)}

	timer := time.NewTicker(1 * time.Minute)

	go func() {
		for {
			select {
			case <-timer.C:
				b.lock.Lock()
				for k, v := range b.visitIPs {
					if time.Since(v) >= 3*time.Minute {
						delete(b.visitIPs, k)
					}
				}
				b.lock.Unlock()
			case <-ctx.Done():
				return
			}
		}
	}()

	return b
}

func (o *Ban) visit(ip string) bool {
	o.lock.Lock()
	defer o.lock.Unlock()

	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}

func main() {
	var wg sync.WaitGroup

	success := int64(0)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ban := NewBan(ctx)

	wg.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func() {
				defer wg.Done()

				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				} else {
					fmt.Printf("%v, %s \n", time.Now().Format(time.RFC3339Nano), ip)
				}
			}()
		}

	}
	wg.Wait()

	fmt.Println("success:", success)
}
