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
}
