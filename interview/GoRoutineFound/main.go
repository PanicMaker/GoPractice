package main

import "sync"

func main() {
	var wg sync.WaitGroup
	n := 10

	target := 5

	for i := 0; i < n; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()

		}()

	}

	wg.Wait()

}
