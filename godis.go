package godis

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"os"
	"time"
)

const (
	OK  = "OK"
	NIL = "(*GODIS NIL TOKEN*)"
)

type Godis struct {
	pool  *redis.Pool
	Error error
	log   *log.Logger
}

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func NewGodis() *Godis {
	pool := newPool(":6379", "")

	/*
		out, err := os.Create("/tmp/godis.log")
		if nil != err {
			panic(err.Error())
		}
	*/
	out := os.Stderr

	logger := log.New(out, "[GODIS] ", log.LstdFlags)
	return &Godis{pool, nil, logger}
}
