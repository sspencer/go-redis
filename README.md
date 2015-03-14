# Godis

Provides a easier interface to Redis, with a method per Redis command and simplified return values.
Code is not aiming to be exhaustive in covering all Redis commands, more of an exploration of
what a simplified interface would look like.

Uses: [github.com/garyburd/redigo/redis](https://github.com/garyburd/redigo/redis/)

Simple GET:

    package main

    import (
        "fmt"
        "github.com/sspencer/go-redis"
    )

    func main() {
        redis := godis.NewGodis()
        fmt.Println("Value:", redis.Get("message1"))
    }
