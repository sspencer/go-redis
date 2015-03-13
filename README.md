# Godis

Provides a easier interface to Redis, with a method per Redis command and simplified return values.

    package main

    import (
        "fmt"
        "github.com/sspencer/go-redis"
    )

    func main() {
        redis := godis.NewGodis()
        fmt.Println("Value:", redis.Get("message1"))
    }
