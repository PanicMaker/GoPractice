package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("golang连接redis")

	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	clientCluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{},
	})

	ctx := context.Background()

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	_, err = clientCluster.Ping(ctx).Result()

	//添加键值对
	err = client.Set(ctx, "golang", "yes", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("键golang设置成功")

}
