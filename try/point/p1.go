package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var p = f()
	fmt.Println(*p)
	fmt.Println(f() == f()) // 同一个函数生成变量的内存地址也不同
	fmt.Println(f(), f())
	v := 1
	fmt.Println(incr(&v))
	fmt.Println(time.Since(start))
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
