package main

import (
	"fmt"
	"github.com/sspencer/go-redis"
	"os"
)

func main() {
	redis := godis.NewGodisConn(":6379", "", os.Stderr)
	fmt.Println("SADD:", redis.SAdd("myset", "one", "two", "three"))
	fmt.Println("SCARD:", redis.SCard("myset"))
}
