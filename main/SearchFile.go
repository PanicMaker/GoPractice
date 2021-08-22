package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var query = "test"
var matches int

var workerCount = 0
var maxWorkerCount = 32
var searchRequest = make(chan string)
var workerDone = make(chan bool)
var foubfMatch = make(chan bool)

func main() {
	start := time.Now()
	workerCount = 1
	go search("C:\\Users\\HER", true)
	waitForWorkerers()
	fmt.Println(matches, "matches")
	fmt.Println(time.Since(start))
}

func waitForWorkerers() {
	for {
		select {
		case path := <-searchRequest:
			workerCount++
			go search(path, true)
		case <-workerDone:
			workerCount--
			if workerCount == 0 {
				return
			}
		case <-foubfMatch:
			matches++
		}
	}
}

func search(path string, master bool) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query {
				matches++
			}
			if file.IsDir() {
				if workerCount < maxWorkerCount {
					searchRequest <- path + name + "\\"
				} else {
					search(path+name+"\\", false)
				}
			}
		}
		if master {
			workerDone <- true
		}
	}
}
