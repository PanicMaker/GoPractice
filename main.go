package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	ch := make(chan int, 1)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writer := io.Writer(f)
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
			_, err := writer.Write(IntToBytes(v))
			if err != nil {
				panic(err)
			}
		default:
			return
		}
	}
	wg.Wait()
}

func IntToBytes(intNum int) []byte {
	uint16Num := uint16(intNum)
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, uint16Num)
	return buf.Bytes()
}
