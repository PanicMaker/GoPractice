package main

import (
	"sync"
	"time"
)

type LeakyBucketLimiter struct {
	capacity int       // 最高水位
	current  int       // 当前水位
	rate     int       // 水流速度/秒
	lastTime time.Time // 上次放水时间
	mu       sync.Mutex
}

func NewLeakyBucketLimiter(capacity int, rate int) *LeakyBucketLimiter {
	return &LeakyBucketLimiter{
		capacity: capacity,
		rate:     rate,
		lastTime: time.Now(),
	}
}

func (l *LeakyBucketLimiter) TryAcquire() (bool, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// 尝试放水
	now := time.Now()

	// 距离上次放水的时间
	interval := now.Sub(l.lastTime)
	if interval >= time.Second {
		// 当前水位-距离上次放水的时间(秒)*水流速度
		l.current = max(0, l.current-int(interval/time.Second)*l.rate)
		l.lastTime = now
	}

	// 若到达最高水位，请求失败
	if l.current >= l.capacity {
		return false, nil
	}

	// 若没有到达最高水位，当前水位+1，请求成功
	l.current++
	return true, nil
}
