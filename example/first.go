package main

import (
	"fmt"
	"github.com/sspencer/go-redis"
	"time"
)

func main() {
	redis := godis.NewGodis()
	redis.Set("message2", fmt.Sprintf("%v", time.Now()))
	redis.Append("message1", "!")
	fmt.Printf("Message1: %s\nMessage2: %s\n", redis.Get("message1"), redis.Get("message2"))
	fmt.Printf("myflt: %f\n", redis.IncrByFloat("myflt", 0.1))
	fmt.Printf("MSETNX success:", redis.MSetNX("go1", "1", "go2", "2"))
	fmt.Printf("SETBIT:", redis.SetBit("mybit", 20, 1))
	redis.SetRange("message2", 5, " ** 2015 **")
	fmt.Println("BITPOS:", redis.BitPos("mybit", 1))
	fmt.Println("LLEN:", redis.LLen("mylist"))
	fmt.Println("BITCOUNT:", redis.BitCount("mybit"))
	fmt.Println("BITOP:", redis.BitOpOr("newbits", "mybit", "myotherbit"))
	fmt.Println("DECR:", redis.Decr("mycounter"))
	fmt.Println("DECRBY:", redis.DecrBy("mycounter", 3))
	fmt.Println("GETBIT:", redis.GetBit("mybit", 16))
	fmt.Println("GETRANGE:", redis.GetRange("message1", 4, 10))
	fmt.Println("INCR:", redis.Incr("mycounter"))
	fmt.Println("INCRBY:", redis.IncrBy("mycounter", 3))
	fmt.Printf("MGET: %#v\n", redis.MGet("message1", "mycount", "doesnotexist", "message2"))
	fmt.Println("MSET:", redis.MSet("my1", "one", "my2", "two", "my3", 3))
	fmt.Println("PSETEX:", redis.PSetEX("tempstore", 10000, "this will go away in 10,000 millis"))
	fmt.Println("SETEX:", redis.SetEX("tempstore2", 10, "this will go away in 10 seconds"))
	fmt.Println("SETNX:", redis.SetNX("tempstore2", "already set"))
	fmt.Println("STRELEN", redis.Strlen("message1"))
	/*
		key, val := redis.BLPop(60, "list1", "list2")
		fmt.Println("BLPOP key:", key, " val:", val)
	*/
}
