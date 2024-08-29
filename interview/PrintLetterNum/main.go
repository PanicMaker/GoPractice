package main

import (
	"fmt"
	"sync"
)

func main() {
	letter, num := make(chan bool), make(chan bool)
	var wg sync.WaitGroup

	go func() {
		i := 1

		for {
			select {
			case <-num:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	wg.Add(1)
	go func() {
		i := 'A'
		for {
			select {
			case <-letter:
				if i > 'Z' {
					wg.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				num <- true
			}
		}

	}()

	num <- true
	wg.Wait()
}
