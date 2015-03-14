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
	fmt.Println("SINTER:", redis.SInter("myset", "myset2"))
	fmt.Println("SINTERSTORE:", redis.SInterStore("mynewset", "myset", "myset2"))
	fmt.Println("SISMEMBER true:", redis.SIsMember("myset", "one"))
	fmt.Println("SISMEMBER false:", redis.SIsMember("myset", "whatever"))
	fmt.Println("SMEMBERS:", redis.SMembers("myset"))
	fmt.Println("SMOVE:", redis.SMove("myset2", "myset", "hello"))
	fmt.Println("SPOP:", redis.SPop("myset2"))
	fmt.Println("SRANDMEMBER:", redis.SRandMember("myset2"))
}
