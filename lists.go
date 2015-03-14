package godis

import (
	"github.com/garyburd/redigo/redis"
)

// BLPop removes and gets the first element in a list or blocks
// until one is available.  Both the key and value are returned.
// You may pass in more than 1 key and BLPop returns an element
// from the first list with an available item.  Specify timeout
// in seconds.   Returns ("","") on timeout or error.
func (g *Godis) BLPop(timeout int, keys ...interface{}) (string, string) {
	conn := g.pool.Get()
	defer conn.Close()

	keys = append(keys, timeout)
	reply, err := conn.Do("BLPOP", keys...)
	g.log.Printf("BLPOP %v\n", keys)

	if retval, err := redis.Strings(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error BLPOP %s\n", err)
		return "", ""
	} else {
		g.Error = nil
		return retval[0], retval[1]
	}
}

// BLPop removes and gets the last element in a list or blocks
// until one is available.  Both the key and value are returned.
// You may pass in more than 1 key and BLPop returns an element
// from the first list with an available item.  Specify timeout
// in seconds.  Returns ("","") on timeout or error.
func (g *Godis) BRPop(timeout int, keys ...interface{}) (string, string) {
	conn := g.pool.Get()
	defer conn.Close()

	keys = append(keys, timeout)
	reply, err := conn.Do("BRPOP", keys...)
	g.log.Printf("BRPOP %v\n", keys)

	if retval, err := redis.Strings(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error BRPOP %s\n", err)
		return "", ""
	} else {
		g.Error = nil
		return retval[0], retval[1]
	}
}

// BRPopLPush pops a value from the source list, pushes it onto destination and
// returns that list element.  Blocks until timeout or element is available.
func (g *Godis) BRPopLPush(source, destination string, timeout int) string {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("BRPOPLPUSH", source, destination, timeout)
	g.log.Printf("BRPOPLPUSH %s %s %s\n", source, destination, timeout)

	if retval, err := redis.String(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error BRPOPLPUSH %s\n", err)
		return ""
	} else {
		g.Error = nil
		return retval
	}
}

// LIndex gets an element from a list by its index.
func (g *Godis) LIndex(key string, index int) string {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("LINDEX", key, index)
	g.log.Printf("LINDEX %s %d\n", key, index)

	if retval, err := redis.String(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error LINDEX %s\n", err)
		return ""
	} else {
		g.Error = nil
		return retval
	}
}

// LLen gets the length of a list.
func (g *Godis) LLen(key string) int {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("LLEN", key)
	g.log.Printf("LLEN %s\n", key)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error LLEN %s\n", err)
		return 0
	} else {
		g.Error = nil
		return retval
	}
}
