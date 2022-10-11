package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "A.领导领导 B.督察指导 C.领导监督 D.监督监督 "
	reg := regexp.MustCompile(`[A-E].([^A-E\d\s]*)`)

	matchStr := reg.FindAllString(str, -1)
	fmt.Println(matchStr)

	for _, subStr := range matchStr {
		fmt.Println(subStr)
	}
}
