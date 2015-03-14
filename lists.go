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
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	keys = append(keys, timeout)
	reply, err := conn.Do("BLPOP", keys...)

	if retval, err := redis.Strings(reply, err); err != nil {
		return NIL, NIL
	} else {
		return retval[0], retval[1]
	}
}

// BRPop removes and gets the last element in a list or blocks
// until one is available.  Both the key and value are returned.
// You may pass in more than 1 key and BRPop returns an element
// from the first list with an available item.  Specify timeout
// in seconds.  Returns (godis.NIL, godis.NIL) on timeout or error.
func (g *Godis) BRPop(timeout int, keys ...interface{}) (string, string) {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	keys = append(keys, timeout)
	reply, err := conn.Do("BRPOP", keys...)

	if retval, err := redis.Strings(reply, err); err != nil {
		return NIL, NIL
	} else {
		return retval[0], retval[1]
	}
}

// BRPopLPush pops a value from the source list, pushes it onto destination and
// returns that list element.  Blocks until timeout or element is available.
func (g *Godis) BRPopLPush(source, destination string, timeout int) string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("BRPOPLPUSH", source, destination, timeout)

	if retval, err := redis.String(reply, err); err != nil {
		return NIL
	} else {
		return retval
	}
}

// LIndex gets an element from a list by its index.
func (g *Godis) LIndex(key string, index int) string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("LINDEX", key, index)

	if retval, err := redis.String(reply, err); err != nil {
		return NIL
	} else {
		return retval
	}
}

func (g *Godis) linsert(key, location, pivot, value string) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("LINSERT", key, location, pivot, value)

	if retval, err := redis.Int(reply, err); err != nil {
		return -1
	} else {
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
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("LLEN", key)

	if retval, err := redis.Int(reply, err); err != nil {
		return 0
	} else {
		return retval
	}
}

// LPop remove and get the first element in a list.
func (g *Godis) LPop(key string) string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("LPOP", key)

	if retval, err := redis.String(reply, err); err != nil {
		return NIL
	} else {
		return retval
	}
}

// LPush prepends one or more items to a list and returns the length
// of the list after the operation.
func (g *Godis) LPush(key string, values ...interface{}) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	keyarg := make([]interface{}, 1)
	keyarg[0] = key
	args := append(keyarg, values...)
	reply, err := conn.Do("LPUSH", args...)

	if retval, err := redis.Int(reply, err); err != nil {
		return -1
	} else {
		return retval
	}
}

// LPushX prepends a value to a list, only if the list exists.  Returns the length
// of the list after the operation.
func (g *Godis) LPushX(key, value string) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("LPUSHX", key, value)

	if retval, err := redis.Int(reply, err); err != nil {
		return 0
	} else {
		return retval
	}
}

// LRange gets a range of elements from a list.
func (g *Godis) LRange(key string, start, stop int) []string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("LRANGE", key, start, stop)

	if retval, err := redis.Strings(reply, err); err != nil {
		return []string{}
	} else {
		return retval
	}
}

// LRem removes elements from a list.  Returns the number of removed elements.
func (g *Godis) LRem(key string, count int, value string) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("LREM", key, count, value)

	if retval, err := redis.Int(reply, err); err != nil {
		return 0
	} else {
		return retval
	}
}

// LSet sets the value of an element in a list by its index.
func (g *Godis) LSet(key string, index int, value string) bool {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("LSET", key, index, value)

	if retval, err := redis.Bool(reply, err); err != nil {
		return false
	} else {
		return retval
	}
}

// LTrim trims a list to the specified range
func (g *Godis) LTrim(key string, start, stop int) bool {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("LTRIM", key, start, stop)

	if retval, err := redis.Bool(reply, err); err != nil {
		return false
	} else {
		return retval
	}
}

// RPop remove and get the last element in a list.
func (g *Godis) RPop(key string) string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("RPOP", key)

	if retval, err := redis.String(reply, err); err != nil {
		return NIL
	} else {
		return retval
	}
}

// RPopLPush removes the last element ina list and prepends it to another list and returns it.
// Can use the same list for both source and destination to process items in a circular fashion.
func (g *Godis) RPopLPush(source, destination string) string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("RPOPLPUSH", source, destination)

	if retval, err := redis.String(reply, err); err != nil {
		return NIL
	} else {
		return retval
	}
}

// RPush appends one or more items to a list and returns the length
// of the list after the operation.
func (g *Godis) RPush(key string, values ...interface{}) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	keyarg := make([]interface{}, 1)
	keyarg[0] = key
	args := append(keyarg, values...)
	reply, err := conn.Do("RPUSH", args...)

	if retval, err := redis.Int(reply, err); err != nil {
		return -1
	} else {
		return retval
	}
}

// RPushX appends a value to a list, only if the list exists.  Returns the length
// of the list after the operation.
func (g *Godis) RPushX(key, value string) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("RPUSHX", key, value)

	if retval, err := redis.Int(reply, err); err != nil {
		return 0
	} else {
		return retval
	}
}
