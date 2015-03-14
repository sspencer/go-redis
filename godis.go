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
	pooled bool
	conn   redis.Conn
	pool   *redis.Pool
	log    *log.Logger
}

// When a Redis call has not results, return a shared empty variable.
var (
	EmptyValues  = make([]interface{}, 0)
	EmptyStrings = []string{}
)

type dialer func() (redis.Conn, error)

func newDialer(server, password string, logger *log.Logger) dialer {
	return func() (redis.Conn, error) {
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

	}
}

func newPool(server, password string, logger *log.Logger) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        newDialer(server, password, logger),
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func NewGodisPool(server, password string, w io.Writer) *Godis {

	logger := log.New(w, "", log.LstdFlags)
	pool := newPool(server, password, logger)

	return &Godis{true, nil, pool, logger}
}

func NewGodisConn(server, password string, w io.Writer) *Godis {

	logger := log.New(w, "", log.LstdFlags)
	dial := newDialer(server, password, logger)
	conn, err := dial()
	if err != nil {
		panic(err)
	}
	return &Godis{false, conn, nil, logger}
}
