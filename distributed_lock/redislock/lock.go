package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

//var client = redis.NewClusterClient(&redis.ClusterOptions{
//	Addrs:    []string{"localhost:6380", "localhost:6381", "localhost:6382", "localhost:6383", "localhost:6384", "localhost:6385"},
//	Password: "", // no password set
//})

func NewUUID() string {
	uuidStr := strings.Replace(uuid.New().String(), "-", "", -1)
	return uuidStr
}

func lock(i int, myfunc func()) {
	defer wg.Done()
	var lockKey = "mylockr"
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	//lock
	uuid := NewUUID()
	lockSuccess, err := client.SetNX(ctx, lockKey, uuid, time.Second*5).Result()
	if err != nil || !lockSuccess {
		fmt.Printf("%d get lock fail\n", i)
		return
	} else {
		fmt.Printf("%d get lock\n", i)
	}
	//run func
	myfunc()
	//unlock
	var luaScript = redis.NewScript(`
		if redis.call("get", KEYS[1]) == ARGV[1]
			then
				return redis.call("del", KEYS[1])
			else
				return 0
		end
	`)
	rs, _ := luaScript.Run(ctx, client, []string{lockKey}, uuid).Result()
	if rs == 0 {
		fmt.Println("unlock fail")
	} else {
		fmt.Println("unlock")
	}
}

// do action
var counter int64

func incr() {
	counter++
	fmt.Printf("after incr is %d\n", counter)
}

// 5 goroutine compete lock
var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		num := i
		go func() {
			lock(num, incr)
		}()
	}
	wg.Wait()
	fmt.Printf("final counter is %d \n", counter)
}
