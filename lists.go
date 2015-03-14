package godis

import (
	"github.com/garyburd/redigo/redis"
)

// BLPop removes and gets the first element in a list or blocks
// until one is available.  Both the key and value are returned.
// You may pass in more than 1 key and BLPop returns an element
// from the first list with an available item.  Specify timeout
// in seconds.   Returns (godis.NIL, godis.NIL) on timeout or error.
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
		return NIL, NIL
	} else {
		g.Error = nil
		return retval[0], retval[1]
	}
}

// BLPop removes and gets the last element in a list or blocks
// until one is available.  Both the key and value are returned.
// You may pass in more than 1 key and BLPop returns an element
// from the first list with an available item.  Specify timeout
// in seconds.  Returns (godis.NIL, godis.NIL) on timeout or error.
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
		return NIL, NIL
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
		return NIL
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
		return NIL
	} else {
		g.Error = nil
		return retval
	}
}

func (g *Godis) linsert(key, location, pivot, value string) int {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("LINSERT", key, location, pivot, value)
	g.log.Printf("LINSERT %s %s %s %s\n", key, location, pivot, value)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error LINSERT %s\n", err)
		return -1
	} else {
		g.Error = nil
		return retval
	}
}

// LInsertAfter inserts an element after another element in a list. Returns list length
// or -1 if no item was inserted.
func (g *Godis) LInsertAfter(key, pivot, value string) int {
	return g.linsert(key, "AFTER", pivot, value)
}

// LInsertAfter inserts an element before another element in a list.  Returns list length
// or -1 if no item was inserted.
func (g *Godis) LInsertBefore(key, pivot, value string) int {
	return g.linsert(key, "BEFORE", pivot, value)
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

// LPop remove and get the first element in a list.
func (g *Godis) LPop(key string) string {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("LPOP", key)
	g.log.Printf("LPOP %s\n", key)

	if retval, err := redis.String(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error LPOP %s\n", err)
		return NIL
	} else {
		g.Error = nil
		return retval
	}
}

// LPush prepends one or more items to a list and returns the length
// of the list after the operation.
func (g *Godis) LPush(key string, values ...interface{}) int {
	conn := g.pool.Get()
	defer conn.Close()

	keyarg := make([]interface{}, 1)
	keyarg[0] = key
	args := append(keyarg, values...)
	reply, err := conn.Do("LPUSH", args...)
	g.log.Printf("LPUSH %v\n", args)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error LPUSH %s\n", err)
		return -1
	} else {
		g.Error = nil
		return retval
	}
}

// LPushX prepends a value to a list, only if the list exists.  Returns the length
// of the list after the operation.
func (g *Godis) LPushX(key, value string) int {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("LPUSHX", key, value)
	g.log.Printf("LPUSHX %s %s\n", key, value)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error LPUSHX %s\n", err)
		return 0
	} else {
		g.Error = nil
		return retval
	}
}

// LRange gets a range of elements from a list.
func (g *Godis) LRange(key string, start, stop int) []string {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("LRANGE", key, start, stop)
	g.log.Printf("LRANGE %s %d %d\n", key, start, stop)

	if retval, err := redis.Strings(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error LRANGE %s\n", err)
		return []string{}
	} else {
		g.Error = nil
		return retval
	}
}

// LRem removes elements from a list.  Returns the number of removed elements.
func (g *Godis) LRem(key string, count int, value string) int {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("LREM", key, count, value)
	g.log.Printf("LREM %s %d %s\n", key, count, value)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error LREM %s\n", err)
		return 0
	} else {
		g.Error = nil
		return retval
	}
}

// LSet sets the value of an element in a list by its index.
func (g *Godis) LSet(key string, index int, value string) bool {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("LSET", key, index, value)
	g.log.Printf("LSET %s %d %s\n", key, index, value)

	if retval, err := redis.Bool(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error LSET %s\n", err)
		return false
	} else {
		g.Error = nil
		return retval
	}
}

// LTrim trims a list to the specified range
func (g *Godis) LTrim(key string, start, stop int) bool {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("LTRIM", key, start, stop)
	g.log.Printf("LTRIM %s %d %d\n", key, start, stop)

	if retval, err := redis.Bool(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error LTRIM %s\n", err)
		return false
	} else {
		g.Error = nil
		return retval
	}
}

// RPush appends one or more items to a list and returns the length
// of the list after the operation.
func (g *Godis) RPush(key string, values ...interface{}) int {
	conn := g.pool.Get()
	defer conn.Close()

	keyarg := make([]interface{}, 1)
	keyarg[0] = key
	args := append(keyarg, values...)
	reply, err := conn.Do("RPUSH", args...)
	g.log.Printf("RPUSH %v\n", args)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error RPUSH %s\n", err)
		return -1
	} else {
		g.Error = nil
		return retval
	}
}

// RPushX appends a value to a list, only if the list exists.  Returns the length
// of the list after the operation.
func (g *Godis) RPushX(key, value string) int {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("RPUSHX", key, value)
	g.log.Printf("RPUSHX %s %s\n", key, value)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error RPUSHX %s\n", err)
		return 0
	} else {
		g.Error = nil
		return retval
	}
}
