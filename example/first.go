package main

import (
	"fmt"
	"github.com/sspencer/go-redis"
)

func main() {
	redis := godis.NewGodis()
	fmt.Println("Success:", redis.Set("message2", "godis starts"))
	fmt.Println("Append Value:", redis.Append("message1", "!"))
	fmt.Println("Value:", redis.Get("message1"))
}
