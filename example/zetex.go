package main

import (
	"fmt"
	"github.com/sspencer/go-redis"
	"os"
)

func main() {
	redis := godis.NewGodisConn(":6379", "", os.Stderr)
	fmt.Println("ZRANGE:", redis.ZRange("myzet", 0, -1))
	fmt.Println("ZRANGE WITHSCORES:", redis.ZRangeWithScores("myzet", 0, -1))
}
