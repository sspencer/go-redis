package main

import (
	"fmt"
	"github.com/sspencer/go-redis"
	"os"
)

func main() {
	redis := godis.NewGodisConn(":6379", "", os.Stderr)
	fmt.Println("SADD:", redis.SAdd("myset", "one", "two", "three"))
	fmt.Println("SADD:", redis.SAdd("myset2", "three", "four", "five"))
	fmt.Println("SCARD:", redis.SCard("myset"))
	fmt.Println("SDIFF:", redis.SDiff("myset", "myset2"))
	fmt.Println("SDIFFSTORE:", redis.SDiffStore("mynewset", "myset", "myset2"))
}
