package geecache

import (
	"fmt"
	"log"
	"sync"
)

// Getter loads data for a key
type Getter interface {
	Get(key string) ([]byte, error)
}

// GetterFunc implements Getter with a function
type GetterFunc func(key string) ([]byte, error)

// Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// Group 负责缓存值的存储与获取
type Group struct {
	name string

	// 缓存未命中时的获取源数据的回调
	getter    Getter
	mainCache cache
	peers     PeerPicker
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

// NewGroup create a new instance of Group
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("Getter nil")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:   name,
		getter: getter,
		mainCache: cache{
			cacheBytes: cacheBytes,
		},
	}
	groups[name] = g
	return g
}

// GetGroup returns the named group previously created with NewGroup, or
// nil if there is no such group
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

// RegisterPeers 将实现 PeerPicker 的 HTTPPOOL 注入 Group
func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.peers = peers
}

// Get 从 cache 中获取数据
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	if v, ok := g.mainCache.get(key); ok {
		fmt.Println("[GeeCache] hit")
		return v, nil
	}

	// 缓存不存在调用，加载数据
	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error) {
	// 远程节点不为空时调用
	if g.peers != nil {
		if peer, ok := g.peers.PickPeer(key); ok {
			if value, err = g.getFromPeer(peer, key); err != nil {
				log.Println("[GeeCache] Failed to get from peer", err)
			}
			return value, err
		}
	}

	return g.getLocally(key)
}

// 从本地获取数据并放入缓存
func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key) // 通过回调函数获取数据
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

// 从远程节点获取数据
func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	bytes, err := peer.Get(g.name, key)
	if err != nil {
		return ByteView{}, err
	}

	return ByteView{b: bytes}, err
}

// 将数据加入缓存中
func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
