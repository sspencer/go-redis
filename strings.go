package godis

import (
	"github.com/garyburd/redigo/redis"
	"math"
)

// Append a value to a key.  Returns the length of the new string or -1 on error.
func (g *Godis) Append(key string, value string) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("APPEND", key, value)

	// ignore return value (new string length) for now, may not be useful
	if retval, err := redis.Int(reply, err); err != nil {
		return -1
	} else {
		return retval
	}
}

// BitCount counts the number of set bits in a string.
// Function can be invoked with a variable number of parameters:
// BITCOUNT key [start end]
func (g *Godis) BitCount(args ...interface{}) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("BITCOUNT", args...)

	if retval, err := redis.Int(reply, err); err != nil {
		return -1
	} else {
		return retval
	}
}

func (g *Godis) bitop(operation string, args ...interface{}) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	op := make([]interface{}, 1)
	op[0] = operation
	args = append(op, args...)

	reply, err := conn.Do("BITOP", args...)

	if retval, err := redis.Int(reply, err); err != nil {
		return 0
	} else {
		return retval
	}
}

// BitOp performs bitwise operations between strings.
// More than 1 source key can be specified.  Returns
// the size of the new key.
// BITOP AND destkey srckey1 [srckey2 srckey3 ... srckeyN]
func (g *Godis) BitOpAnd(args ...interface{}) int {
	return g.bitop("AND", args...)
}

// BitOp performs bitwise operations between strings.
// More than 1 source key can be specified.  Returns
// the size of the new key.
// BITOP OR destkey srckey1 [srckey2 srckey3 ... srckeyN]
func (g *Godis) BitOpNot(destkey, srckey string) int {
	args := make([]interface{}, 2)

	args[0] = destkey
	args[1] = srckey
	return g.bitop("NOT", args...)
}

// BitOp performs bitwise operations between strings.
// More than 1 source key can be specified.  Returns
// the size of the new key.
// BITOP OR destkey srckey1 [srckey2 srckey3 ... srckeyN]
func (g *Godis) BitOpOr(args ...interface{}) int {
	return g.bitop("OR", args...)
}

// BitOp performs bitwise operations between strings.
// More than 1 source key can be specified.  Returns
// the size of the new key.
// BITOP OR destkey srckey1 [srckey2 srckey3 ... srckeyN]
func (g *Godis) BitOpXor(args ...interface{}) int {
	return g.bitop("XOR", args...)
}

// BitPos returns the position of the first bit set to 1 or 0 in a string.
// Function can be invoked with a variable number of parameters:
// BITPOS key bit [start] [end]
func (g *Godis) BitPos(args ...interface{}) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("BITPOS", args...)

	if retval, err := redis.Int(reply, err); err != nil {
		return -1
	} else {
		return retval
	}
}

// Decr decrements the integer value of a key by one.  Returns math.MaxInt64
// on error.
func (g *Godis) Decr(key string) int64 {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("DECR", key)

	if retval, err := redis.Int64(reply, err); err != nil {
		return math.MaxInt64
	} else {
		return retval
	}
}

// DecrBy decrements the integer value of a key by the given number.  Returns math.MaxInt64
// on error.
func (g *Godis) DecrBy(key string, decrement int) int64 {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("DECRBY", key, decrement)

	if retval, err := redis.Int64(reply, err); err != nil {
		return math.MaxInt64
	} else {
		return retval
	}
}

// Get the value of a key.  Return godis.NIL if key does not exist or upon error.  To
// tell the difference, look at godis.Error.
func (g *Godis) Get(key string) string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("GET", key)

	if retval, err := redis.String(reply, err); err != nil {
		return NIL
	} else {
		return retval
	}
}

// GetBit returns the bit value at offset in the value stored at key.
// Returns -1 on error.
func (g *Godis) GetBit(key string, offset int) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("GETBIT", key, offset)

	if retval, err := redis.Int(reply, err); err != nil {
		return -1
	} else {
		return retval
	}
}

// Get a substring of the string stored at key.
func (g *Godis) GetRange(key string, start, end int) string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("GETRANGE", key, start, end)

	if retval, err := redis.String(reply, err); err != nil {
		return NIL
	} else {
		return retval
	}
}

// Set the string value of a key and return its old value.
func (g *Godis) GetSet(key, value string) string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("GETSET", key, value)

	if retval, err := redis.String(reply, err); err != nil {
		return NIL
	} else {
		return retval
	}
}

// Incr increments the integer value of a key by one.  Returns math.MinInt64
// on error.
func (g *Godis) Incr(key string) int64 {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("INCR", key)

	if retval, err := redis.Int64(reply, err); err != nil {
		return math.MinInt64
	} else {
		return retval
	}
}

// IncrBy increments the integer value of a key by the given number.  Returns math.MinInt64
// on error.
func (g *Godis) IncrBy(key string, increment int) int64 {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("INCRBY", key, increment)

	if retval, err := redis.Int64(reply, err); err != nil {
		return math.MaxInt64
	} else {
		return retval
	}
}

// IncrByFloat increments the float value of a key by the given amount.  Return math.MaxFloat64 on error.
func (g *Godis) IncrByFloat(key string, value float64) float64 {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("INCRBYFLOAT", key, value)

	if retval, err := redis.Float64(reply, err); err != nil {
		return math.MaxFloat64
	} else {
		return retval
	}
}

// MGet gets the values of all the given keys.
func (g *Godis) MGet(keys ...interface{}) []string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("MGET", keys...)

	if retval, err := redis.Strings(reply, err); err != nil {
		return EmptyStrings
	} else {
		return retval
	}
}

// MSet sets multiple keys to multiple values.
func (g *Godis) MSet(keyvals ...interface{}) bool {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("MSET", keyvals...)

	if retval, err := redis.String(reply, err); err != nil {
		return false
	} else {
		return retval == OK
	}
}

// MSetNX sets multiple keys to multiple values, only if none of the keys already exist.
func (g *Godis) MSetNX(keyvals ...interface{}) bool {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("MSETNX", keyvals...)

	if retval, err := redis.Int(reply, err); err != nil {
		return false
	} else {
		return !(retval == 0)
	}
}

// Set the string value of a key with an expiration time in milliseconds.
func (g *Godis) PSetEX(key string, millis int, value string) bool {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("PSETEX", key, millis, value)

	if retval, err := redis.String(reply, err); err != nil {
		return false
	} else {
		return retval == OK
	}
}

// Set the string value of a key.
func (g *Godis) Set(key string, value string) bool {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("SET", key, value)

	if retval, err := redis.String(reply, err); err != nil {
		return false
	} else {
		return retval == OK
	}
}

// SetBit sets or clears the bit at offset in the string value stored at key.
// Returns -1 on error.
func (g *Godis) SetBit(key string, offset, value int) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("SETBIT", key, offset, value)

	if retval, err := redis.Int(reply, err); err != nil {
		return -1
	} else {
		return retval
	}
}

// Set the string value of a key with an expiration time in seconds.
func (g *Godis) SetEX(key string, seconds int, value string) bool {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("SETEX", key, seconds, value)

	if retval, err := redis.String(reply, err); err != nil {
		return false
	} else {
		return retval == OK
	}
}

// Set the string value of a key if the key does not already exist.
func (g *Godis) SetNX(key string, value string) bool {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("SETNX", key, value)

	if retval, err := redis.Int(reply, err); err != nil {
		return false
	} else {
		return retval == 1
	}
}

// SetRange overwrites part of a string at key starting at the specified offset.
func (g *Godis) SetRange(key string, offset int, value string) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("SETRANGE", key, offset, value)

	if retval, err := redis.Int(reply, err); err != nil {
		return -1
	} else {
		return retval
	}
}

// Strlen returns the length of the value stored at key.
func (g *Godis) Strlen(key string) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("STRLEN", key)

	if retval, err := redis.Int(reply, err); err != nil {
		return -1
	} else {
		return retval
	}
}
