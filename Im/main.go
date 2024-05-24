package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	go func() {
		// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
	}()

	server := NewServer("0.0.0.0", 123)
	server.Start()

}
