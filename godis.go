package godis

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"os"
	"time"
)

const (
	OK = "OK"
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

/*--------------------------------------------------------------
 * STRING Commands
 *--------------------------------------------------------------*/

// Append a value to a key.  Returns the length of the new string or -1 on error.
func (g *Godis) Append(key string, value string) {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("APPEND", key, value)
	g.log.Printf("APPEND %s \"%s\"\n", key, value)

	// ignore return value (new string length) for now, may not be useful
	if _, err := redis.Int(reply, err); err != nil {
		// handle error
		g.log.Printf("Error APPEND %s\n", err)
		g.Error = err
	} else {
		g.Error = nil
	}
}

// Get the value of a key.  Return "" is key does not exist or upon error.  To
// tell the difference, look at godis.Error.
func (g *Godis) Get(key string) string {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("GET", key)
	g.log.Printf("GET %s\n", key)

	if retval, err := redis.String(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error GET %s\n", err)
		return ""
	} else {
		g.Error = nil
		return retval
	}
}

// Set the string value of a key and return its old value.
//
func (g *Godis) GetSet(key, value string) string {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("GETSET", key, value)
	g.log.Printf("GETSET %s \"%s\"\n", key, value)

	if retval, err := redis.String(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error GETSET %s\n", err)
		return ""
	} else {
		g.Error = nil
		return retval
	}
}

// Set the string value of a key.
func (g *Godis) Set(key string, value string) {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("SET", key, value)
	g.log.Printf("SET %s \"%s\"\n", key, value)

	if _, err := redis.String(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error SET %s\n", err)
	} else {
		g.Error = nil
	}
}
