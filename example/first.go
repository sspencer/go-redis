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
}
