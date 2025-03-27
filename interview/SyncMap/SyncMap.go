package main

import "sync"

type SyncMap interface {
	Store(key any, value any)
	Load(key any) (any, bool)
}

type MutexMap struct {
	data map[any]any
	mu   sync.Mutex
}

func (m *MutexMap) Store(key any, value any) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *MutexMap) Load(key any) (v any, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok = m.data[key]
	return
}

type RWMutexMap struct {
	data map[any]any
	mu   sync.RWMutex
}

func (m *RWMutexMap) Store(key any, value any) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *RWMutexMap) Load(key any) (v any, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok = m.data[key]
	return
}
