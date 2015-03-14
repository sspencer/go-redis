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
}
