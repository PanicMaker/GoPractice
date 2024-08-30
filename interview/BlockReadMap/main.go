package main

import (
	"fmt"
	"sync"
	"time"
)

type sp interface {
	Out(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type SafeMap struct {
	mu    sync.Mutex
	data  map[string]any
	conds map[string]*sync.Cond
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data:  make(map[string]any),
		conds: make(map[string]*sync.Cond),
	}
}

// Out 方法：存入key/val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
func (m *SafeMap) Out(key string, val any) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[key] = val

	// 如果该key有等待的goroutine，则唤醒它们
	if cond, ok := m.conds[key]; ok {
		cond.Broadcast() // 唤醒所有等待该key的goroutine
	}
}

// Rd 方法：读取一个key，如果key不存在阻塞，等待key存在或者超时
func (m *SafeMap) Rd(key string, timeout time.Duration) interface{} {
	m.mu.Lock()

	// 如果key存在，立即返回
	if val, ok := m.data[key]; ok {
		m.mu.Unlock()
		return val
	}

	// 如果key不存在，创建或获取该key的条件变量
	cond, ok := m.conds[key]
	if !ok {
		cond = sync.NewCond(&m.mu)
		m.conds[key] = cond
	}

	// 开始等待条件满足或超时
	done := make(chan struct{})
	go func() {
		m.mu.Lock()
		cond.Wait() // 等待条件
		close(done)
		m.mu.Unlock()
	}()

	m.mu.Unlock()

	select {
	case <-done:
		// 当唤醒时再次检查数据
		m.mu.Lock()
		defer m.mu.Unlock()
		return m.data[key]
	case <-time.After(timeout):
		// 超时
		fmt.Println("Time Out!")
		return nil
	}

}

func main() {
	// 创建SafeMap实例
	safeMap := NewSafeMap()

	// 启动一个goroutine读取值
	go func() {
		fmt.Println("Waiting for key1...")
		value := safeMap.Rd("key1", 3*time.Second)
		if value != nil {
			fmt.Printf("Got value for key1: %v\n", value)
		} else {
			fmt.Println("Timeout while waiting for key1.")
		}
	}()

	// 模拟一些延迟后设置值
	time.Sleep(1 * time.Second)
	safeMap.Out("key1", "Hello, Go!")

	// 等待goroutine完成
	time.Sleep(4 * time.Second)
}
