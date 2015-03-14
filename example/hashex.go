package main

import (
	"fmt"
	"github.com/sspencer/go-redis"
	"os"
)

func main() {
	redis := godis.NewGodisConn(":6379", "", os.Stderr)
	fmt.Println("HGET:", redis.HGet("track123", "title"))
	fmt.Println("HDEL:", redis.HDel("track123", "these", "do", "not", "exist"))
	fmt.Println("HEXISTS:", redis.HExists("track123", "artist"))
	fmt.Println("HINCRBY:", redis.HIncrBy("track123", "plays", 1))
	fmt.Println("HINCRBYFLOAT:", redis.HIncrByFloat("track123", "share", 0.001))
	fmt.Println("HGETALL:", redis.HGetAll("track123"))
	fmt.Println("HKEYS:", redis.HKeys("track123"))
	fmt.Println("HVALS:", redis.HVals("track123"))
	fmt.Println("HLEN:", redis.HLen("track123"))
	fmt.Println("HMGET:", redis.HMGet("track123", "artist", "title", "plays"))
	fmt.Println("HMSET:", redis.HMSet("track124", "artist", "Natalia Clavier", "title", "Simple"))
	fmt.Println("HSET:", redis.HSet("track124", "plays", "1"))
	fmt.Println("HSETNX:", redis.HSetNX("track124", "plays", "1"))
}
