package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash Map bytes to uint32
type Hash func(data []byte) uint32

// Map contains all hashed keys
type Map struct {
	hash     Hash
	replicas int            // 虚拟节点倍数
	keys     []int          // 哈希环
	hashMap  map[int]string // 虚拟节点与真实节点的映射表
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}

	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}

	return m
}

// Add adds some keys to the hash
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			// 计算虚拟节点的hash值
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			// 添加到环上
			m.keys = append(m.keys, hash)
			// 增加虚拟节点与真实节点的映射关系
			m.hashMap[hash] = key
		}
	}
	// 哈希值排序
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	// 计算key的哈希值
	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	return m.hashMap[m.keys[idx%len(m.keys)]]
}
