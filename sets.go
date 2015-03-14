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

	args := make([]interface{}, 1)
	args[0] = key
	args = append(args, members...)
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

// SDiff returns the members of the set resulting from the difference between
// the first set and all the successive sets.
func (g *Godis) SDiff(key string, keys ...interface{}) []string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	args := make([]interface{}, 1)
	args[0] = key
	args = append(args, keys...)
	reply, err := conn.Do("SDIFF", args...)

	if retval, err := redis.Strings(reply, err); err != nil {
		// handle error
		g.log.Printf("Error SDIFF %s\n", err)
		return []string{}
	} else {
		return retval
	}
}

// SDiffStore stores the members of the set resulting from the difference between
// the first set and all the successive sets into destination and returns the count
// of members in destination.
func (g *Godis) SDiffStore(destination, key string, keys ...interface{}) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	args := make([]interface{}, 2)
	args[0] = destination
	args[1] = key
	args = append(args, keys...)
	reply, err := conn.Do("SDIFFSTORE", args...)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.log.Printf("Error SDIFFSTORE %s\n", err)
		return 0
	} else {
		return retval
	}
}
