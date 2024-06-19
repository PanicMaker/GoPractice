package lru

import "container/list"

type Cache struct {
	maxBytes int64
	nBytes   int64
	list     *list.List
	cache    map[string]*list.Element

	// 当记录被移除时的回调函数
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		list:      list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get 查找节点
func (c *Cache) Get(key string) (Value, bool) {
	if element, ok := c.cache[key]; ok {
		// 将节点移到队尾
		c.list.MoveToBack(element)
		kv := element.Value.(*entry)
		return kv.value, true
	}
	return nil, false
}

// RemoveOldest 删除最近最少访问的节点
func (c *Cache) RemoveOldest() {
	// 取到队首节点
	element := c.list.Front()
	if element != nil {
		c.list.Remove(element)
		kv := element.Value.(*entry)
		// 删除该节点的映射关系
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {

	if element, ok := c.cache[key]; ok {
		// 存在更新对应的值并移动到队尾
		c.list.MoveToBack(element)
		kv := element.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		// 不存在则新增节点并放在队尾
		element := c.list.PushBack(&entry{
			key:   key,
			value: value,
		})
		c.cache[key] = element
		c.nBytes += int64(len(key)) + int64(value.Len())
	}

	// 超过最大容量，移除最近最少访问节点
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.list.Len()
}
