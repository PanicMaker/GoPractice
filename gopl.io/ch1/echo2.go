// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var s, sep string
	start := time.Now()
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Since(start))

}
