package main

import (
	"fmt"
	"github.com/sspencer/go-redis"
	"os"
	"time"
)

func main() {
	redis := godis.NewGodisConn(":6379", "", os.Stderr)
	// key, val := redis.BLPop(60, "list1", "list2")
	// fmt.Println("BLPOP key:", key, " val:", val)
	// fmt.Println("BRPOPLPUSH:", redis.BRPopLPush("list1", "list2", 2))
	// fmt.Println("LINSERT:", redis.LInsertBefore("list2", "four", "five"))
	//fmt.Println("LPOP", redis.LPop("list2"))
	fmt.Println("LPUSH:", redis.LPush("list3", "one", "two", "three"))
	fmt.Println("LPUSHX:", redis.LPushX("list30000", "zero"))
	fmt.Println("RPUSH:", redis.RPush("list3", "one hundred", "two hundred", "three hundred"))
	fmt.Println("RPUSHX:", redis.RPushX("list30000", "zero"))
	fmt.Println("LREM:", redis.LRem("list3", 0, "two"))
	fmt.Println("LSET:", redis.LSet("list3", 0, "front of list"))
	fmt.Println("LTRIM:", redis.LTrim("list3", 0, 6))
	fmt.Println("LRANGE:", redis.LRange("list3", 0, -1))
	fmt.Println("RPOP:", redis.RPop("list3"))
	fmt.Println("LRANGE:", redis.LRange("list3", 0, -1))
	fmt.Println("RPUSHLPOP:", redis.RPopLPush("list3", "list3"))
	fmt.Println("LRANGE:", redis.LRange("list3", 0, -1))
}
