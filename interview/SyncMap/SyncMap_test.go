package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkSyncMap(b *testing.B) {
	maps := [...]SyncMap{
		&MutexMap{
			data: make(map[any]any),
		},
		&RWMutexMap{
			data: make(map[any]any),
		},
		&sync.Map{},
	}

	// 测试对于同一个 key 的 n-1 并发读和 1 并发写的性能
	for _, m := range maps {
		b.Run(fmt.Sprintf("%T", m), func(b *testing.B) {
			var i int64
			b.RunParallel(func(pb *testing.PB) {
				// 记录并发执行的 goroutine id
				gid := int(atomic.AddInt64(&i, 1) - 1)

				if gid == 0 {
					// gid 为 0 的 goroutine 负责并发写
					for i := 0; pb.Next(); i++ {
						m.Store(0, i)
					}
				} else {
					// gid 不为 0 的 goroutine 负责并发读
					for pb.Next() {
						m.Load(0)
					}
				}
			})
		})
	}
}
