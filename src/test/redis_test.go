package test

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	redis2 "github.com/go-redis/redis/v8"
	"strconv"
	"sync"
	"testing"
	"time"
)

func BenchmarkRedis(b *testing.B) {
	// redigo
	// var pool *redis.Pool
	// pool = &redis.Pool{
	// 	MaxIdle:     10,   // 最初的连接数量
	// 	MaxActive:   1000, // 连接池最大连接数量,（0表示自动定义），按需分配
	// 	IdleTimeout: 300,  // 连接关闭时间 300秒 （300秒不使用自动关闭）
	// 	Dial: func() (redis.Conn, error) { // 要连接的redis数据库
	// 		return redis.Dial("tcp", "localhost:6379")
	// 	},
	// }

	// go-redis
	var goridis = redis2.NewClient(
		&redis2.Options{
			Addr:         "localhost:6379",
			Password:     "",
			MinIdleConns: 10,
			PoolSize:     1000,
		})

	// miniredis
	// mock一个redis server
	miniredi, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer miniredi.Close()

	minigoredisClient := redis2.NewClient(&redis2.Options{
		Addr: miniredi.Addr(),
	})

	// memory
	memory := sync.Map{}

	// b.Run("redigo client Benchmark", func(b *testing.B) {
	// 	for j := 0; j < b.N; j++ {
	// 		conn := pool.Get() // 从连接池，取一个链接
	// 		conn.Do("set", time.Now().String(), 10000, time.Second)
	// 		conn.Do("get", time.Now().String())
	// 		conn.Close()
	// 	}
	// })

	ctx := context.Background()
	b.Run("go-redis client Benchmark", func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			key := strconv.Itoa(j)
			goridis.Set(ctx, key, 1000, time.Second)
			goridis.Get(ctx, key)
		}
	})

	b.Run("miniredis Benchmark", func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			key := strconv.Itoa(j)
			miniredi.Set(key, "1000")
			miniredi.Get(key)
		}
	})

	b.Run("sync.Map Benchmark", func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			key := strconv.Itoa(j)
			memory.Store(key, "1000")
			memory.Load(key)
		}
	})

	b.Run("minigoredisClient Benchmark", func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			key := strconv.Itoa(j)
			minigoredisClient.Set(ctx, key, 1000, time.Second)
			minigoredisClient.Get(ctx, key)
		}
	})
}
