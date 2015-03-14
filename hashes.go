package godis

import (
	"github.com/garyburd/redigo/redis"
)

// HDel deletes one of more hash fields.
func (g *Godis) HDel(key string, fields ...interface{}) int {
	conn := g.pool.Get()
	defer conn.Close()

	keyarg := make([]interface{}, 1)
	keyarg[0] = key
	args := append(keyarg, fields...)
	reply, err := conn.Do("HDEL", args...)
	g.log.Printf("HDEL %v\n", args)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HDel %s\n", err)
		return 0
	} else {
		g.Error = nil
		return retval
	}
}

// HExists determines if a hash field exists
func (g *Godis) HExists(key, field string) bool {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("HExists", key, field)
	g.log.Printf("HExists %s %s\n", key, field)

	if retval, err := redis.Bool(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HExists %s\n", err)
		return false
	} else {
		g.Error = nil
		return retval
	}
}

// HGet gets the value of a hash field
func (g *Godis) HGet(key, field string) string {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("HGET", key, field)
	g.log.Printf("HGET %s %s\n", key, field)

	if retval, err := redis.String(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HGET %s\n", err)
		return NIL
	} else {
		g.Error = nil
		return retval
	}
}
