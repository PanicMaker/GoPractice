package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(&x)
	fmt.Println(p)
	fmt.Println(*p)
	*p = 3
	fmt.Println(x)
}
