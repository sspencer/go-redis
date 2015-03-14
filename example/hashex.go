package main

import (
	"fmt"
	"github.com/sspencer/go-redis"
)

func main() {
	redis := godis.NewGodis()
	fmt.Println("HGET:", redis.HGet("track123", "title"))
	fmt.Println("HDEL:", redis.HDel("track123", "these", "do", "not", "exist"))
	fmt.Println("HEXISTS:", redis.HExists("track123", "artist"))
	fmt.Println("HINCRBY:", redis.HIncrBy("track123", "plays", 1))
	fmt.Println("HINCRBYFLOAT:", redis.HIncrByFloat("track123", "share", 0.001))
	fmt.Println("HGETALL:", redis.HGetAll("track123"))
	fmt.Println("HKEYS:", redis.HKeys("track123"))
	fmt.Println("HVALS:", redis.HVals("track123"))
	fmt.Println("HLEN:", redis.HLen("track123"))
}
