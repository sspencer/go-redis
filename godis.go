package godis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	OK = "OK"
)

type Godis struct {
	pool  *redis.Pool
	Error error
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
	return &Godis{pool, nil}
}

/*--------------------------------------------------------------
 * STRING Commands
 *--------------------------------------------------------------*/

// Append a value to a key.  Returns the length of the new string or -1 on error.
func (g *Godis) Append(key string, value string) int {
	conn := g.pool.Get()
	defer conn.Close()

	r, err := conn.Do("APPEND", key, value)

	if value, err := redis.Int(r, err); err != nil {
		// handle error
		g.Error = err
		return -1
	} else {
		g.Error = nil
		return value
	}
}

// Get the value of a key.  Return "" is key does not exist or upon error.  To
// tell the difference, look at godis.Error.
func (g *Godis) Get(key string) string {
	conn := g.pool.Get()
	defer conn.Close()

	r, err := conn.Do("GET", key)

	if value, err := redis.String(r, err); err != nil {
		// handle error
		g.Error = err
		return ""
	} else {
		g.Error = nil
		return value
	}
}

// Set the string value of a key.
func (g *Godis) Set(key string, value string) bool {
	conn := g.pool.Get()
	defer conn.Close()

	r, err := conn.Do("SET", key, value)

	if value, err := redis.String(r, err); err != nil {
		// handle error
		g.Error = err
		return false
	} else {
		g.Error = nil
		return value == OK
	}
}
