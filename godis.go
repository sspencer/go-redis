package godis

import (
	"github.com/garyburd/redigo/redis"
	"io"
	"log"
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

func newPool(server, password string, logger *log.Logger) *redis.Pool {
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

			if logger != nil {
				loggingConn := redis.NewLoggingConn(c, logger, "[GODIS] ")
				return loggingConn, err
			} else {
				return c, err
			}

		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func NewGodisPool(w io.Writer) *Godis {

	logger := log.New(w, "[GODIS]", log.LstdFlags)
	pool := newPool(":6379", "", logger)

	return &Godis{pool, nil, logger}
}
