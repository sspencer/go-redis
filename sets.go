package godis

import (
	"github.com/garyburd/redigo/redis"
)

// SAdd adds one or more members to a set.  Returns the number
// of elements added to the set (not including members already present).
func (g *Godis) SAdd(key string, members ...interface{}) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	keyarg := make([]interface{}, 1)
	keyarg[0] = key
	args := append(keyarg, members...)
	reply, err := conn.Do("SADD", args...)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.log.Printf("Error SADD %s\n", err)
		return 0
	} else {
		return retval
	}
}

// SCard returns the number of members in a set.
func (g *Godis) SCard(key string) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("SCARD", key)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.log.Printf("Error SCARD %s\n", err)
		return 0
	} else {
		return retval
	}
}
