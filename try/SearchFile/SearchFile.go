package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var query = "test"
var matches int

func main() {
	start := time.Now()
	search("/Users/vincent/")
	fmt.Println(matches, "matches")
	fmt.Println(time.Since(start))
}

func search(path string) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query {
				matches++
			}
			if file.IsDir() {
				search(path + name + "/")
			}
		}
	}
}
