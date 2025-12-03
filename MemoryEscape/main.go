package main

import "fmt"

// 场景一：返回局部变量的指针
func pointerEscape() *int {
	x := 10   // x 原本在栈上
	return &x // 逃逸！因为外部还在引用它
}

// 场景二：interface{} 动态类型（fmt.Println）
func interfaceEscape() {
	y := 20
	fmt.Println(y) // 逃逸！Println 接收的是 interface{}，编译器无法确定具体类型
}

// 场景三：栈空间不足（大对象）
func bigStackEscape() {
	// 这是一个巨大的切片，栈（默认2KB-8KB）放不下
	// 虽然 Go 栈可以动态扩容，但太大时编译器会直接扔到堆上
	s := make([]int, 10000, 10000)
	// 注意：如果只是 make([]int, 10, 10)，通常不会逃逸
	_ = s
}

// 场景四：闭包引用
func closureEscape() func() int {
	n := 100 // n 被闭包捕获
	return func() int {
		n++ // 逃逸！n 的生命周期超过了 closureEscape 函数
		return n
	}
}

// 对照组：正常传值（不逃逸）
func noEscape(a int) int {
	b := a + 1
	return b
}

func main() {
	_ = pointerEscape()
	interfaceEscape()
	bigStackEscape()
	_ = closureEscape()

	val := 10
	_ = noEscape(val) // 纯栈操作，不逃逸
}
