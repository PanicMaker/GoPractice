package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	var wg = sync.WaitGroup{}
	//nums := make(chan string, 5)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//nums <- getNum()
			fmt.Fprintln(os.Stdout, time.Now().String()+" "+getNum())
		}()
	}
	wg.Wait()
}

func getNum() string {
	var dataUrl = "http://127.0.0.1:9999/"
	resp, err := http.Get(dataUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	//var res tokenRes
	//_ = json.Unmarshal(body, &res)
	// fmt.Printf("%+v", res)
	return string(body)
}
