package main

import "math/rand"

// https://leetcode.cn/problems/insert-delete-getrandom-o1/description/

type RandomizedSet struct {
	maps  map[int]int
	array []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		maps:  make(map[int]int),
		array: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.maps[val]; !ok {
		this.array = append(this.array, val)
		this.maps[val] = len(this.array) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.maps[val]; ok {
		this.array[i] = this.array[len(this.array)-1] // 将最后一个元素放到要删除的地方
		this.maps[this.array[i]] = i                  // 更新原来最后一个元素在maps中存储的下标
		delete(this.maps, val)
		this.array = this.array[:len(this.array)-1]
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.array[rand.Intn(len(this.array))]
}
