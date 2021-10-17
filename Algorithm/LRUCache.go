package main

import "fmt"

type LRUCache struct {
	capacity int
	cache    map[int]*DListNode
	list     *DoubleLinkedList
}

type DListNode struct {
	Key, Val   int
	Prev, Next *DListNode
}

type DoubleLinkedList struct {
	head, tail *DListNode
	size       int
}

func initDoubleLinkList() *DoubleLinkedList {
	l := &DoubleLinkedList{
		head: &DListNode{},
		tail: &DListNode{},
		size: 0,
	}
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

func (l *DoubleLinkedList) addList(x *DListNode) {
	x.Prev = l.tail.Prev
	x.Next = l.tail
	l.tail.Prev.Next = x
	l.tail.Prev = x
	l.size++
}

func (l *DoubleLinkedList) remove(x *DListNode) {
	x.Prev.Next = x.Next
	x.Next.Prev = x.Prev
	l.size--
}

func (l *DoubleLinkedList) removeFirst() *DListNode {
	if l.head.Next == l.tail {
		return nil
	}
	tmp := l.head.Next
	l.remove(tmp)
	return tmp
}

func (l *DoubleLinkedList) getSize() int {
	return l.size
}

func (l *DoubleLinkedList) show() {
	cur := l.head.Next
	var nums []int
	for cur != l.tail {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}
	fmt.Println(nums)
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    map[int]*DListNode{},
		list:     initDoubleLinkList(),
	}
}

func (c *LRUCache) makeRecently(key int) {
	node := c.cache[key]
	c.list.remove(node)
	c.list.addList(node)

}

func (c *LRUCache) addRecently(key, val int) {
	node := &DListNode{
		Key: key,
		Val: val,
	}
	c.list.addList(node)
	c.cache[key] = node
}

func (c *LRUCache) deleteKey(key int) {
	node := c.cache[key]
	c.list.remove(node)
	delete(c.cache, key)
}

func (c *LRUCache) removeLeastRecently() {
	node := c.list.removeFirst()
	key := node.Key
	delete(c.cache, key)
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.cache[key]; !ok {
		return -1
	}
	c.makeRecently(key)
	c.list.show()
	return c.cache[key].Val
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.cache[key]; ok {
		delete(c.cache, key)
		c.addRecently(key, value)
	}

	if c.capacity == c.list.getSize() {
		c.removeLeastRecently()
	}
	c.addRecently(key, value)
	c.list.show()
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Get(1)
	obj.Put(3, 3)
	obj.Get(2)
	obj.Put(4, 4)
	obj.Get(1)
	obj.Get(3)
	obj.Get(4)
}
