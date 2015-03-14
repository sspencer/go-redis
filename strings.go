package godis

import (
	"github.com/garyburd/redigo/redis"
	"math"
)

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

// BitCount counts the number of set bits in a string.
// Function can be invoked with a variable number of parameters:
// BITCOUNT key [start end]
func (g *Godis) BitCount(args ...interface{}) int {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("BITCOUNT", args...)
	g.log.Printf("BITCOUNT %v", args)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error BITCOUNT %s\n", err)
		return -1
	} else {
		g.Error = nil
		return retval
	}
}

func (g *Godis) bitop(operation string, args ...interface{}) int {
	conn := g.pool.Get()
	defer conn.Close()

	op := make([]interface{}, 1)
	op[0] = operation
	args = append(op, args...)

	reply, err := conn.Do("BITOP", args...)
	g.log.Printf("BITOP %v", args)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error BITOP %s %s\n", operation, err)
		return 0
	} else {
		g.Error = nil
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
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("BITPOS", args...)
	g.log.Printf("BITPOS %v", args)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error BITPOS %s\n", err)
		return -1
	} else {
		g.Error = nil
		return retval
	}
}

// Decr decrements the integer value of a key by one.  Returns math.MaxInt64
// on error.
func (g *Godis) Decr(key string) int64 {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("DECR", key)
	g.log.Printf("DECR %s\n", key)

	if retval, err := redis.Int64(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error DECR %s\n", err)
		return math.MaxInt64
	} else {
		g.Error = nil
		return retval
	}
}

// Decr decrements the integer value of a key by the given number.  Returns math.MaxInt64
// on error.
func (g *Godis) DecrBy(key string, decrement int) int64 {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("DECRBY", key, decrement)
	g.log.Printf("DECRBY %s %d\n", key, decrement)

	if retval, err := redis.Int64(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error DECRBY %s\n", err)
		return math.MaxInt64
	} else {
		g.Error = nil
		return retval
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

// GetBit returns the bit value at offset in the value stored at key.
// Returns -1 on error.
func (g *Godis) GetBit(key string, offset int) int {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("GETBIT", key, offset)
	g.log.Printf("GETBIT %s %d\n", key, offset)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error GETBIT %s\n", err)
		return -1
	} else {
		g.Error = nil
		return retval
	}
}

// Get a substring of the string stored at key.
func (g *Godis) GetRange(key string, start, end int) string {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("GETRANGE", key, start, end)
	g.log.Printf("GETRANGE %s %d %d\n", key, start, end)

	if retval, err := redis.String(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error GETRANGE %s\n", err)
		return ""
	} else {
		g.Error = nil
		return retval
	}
}

// Set the string value of a key and return its old value.
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

// Increment the float value of a key by the given amount.  Return math.MaxFloat64 on error.
func (g *Godis) IncrByFloat(key string, value float64) float64 {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("INCRBYFLOAT", key, value)
	g.log.Printf("INCRBYFLOAT %s \"%f\"\n", key, value)

	if retval, err := redis.Float64(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error INCRBYFLOAT %s\n", err)
		return math.MaxFloat64
	} else {
		g.Error = nil
		return retval
	}
}

// MSetNX sets multiple keys to multiple values, only if none of the keys already exist.
func (g *Godis) MSetNX(keyvals ...interface{}) bool {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("MSETNX", keyvals...)
	g.log.Printf("MSETNX %v", keyvals)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error MULTIPLE %s\n", err)
		return false
	} else {
		g.Error = nil
		return !(retval == 0)
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

// SetBit sets or clears the bit at offset in the string value stored at key.
// Returns -1 on error.
func (g *Godis) SetBit(key string, offset, value int) int {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("SETBIT", key, offset, value)
	g.log.Printf("SETBIT %s %d %d\n", key, offset, value)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error SETBIT %s\n", err)
		return -1
	} else {
		g.Error = nil
		return retval
	}
}

// SetRange overwrites part of a string at key starting at the specified offset.
func (g *Godis) SetRange(key string, offset int, value string) {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("SETRANGE", key, offset, value)
	g.log.Printf("SETRANGE %s %d \"%s\"\n", key, offset, value)

	if _, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error SETRANGE %s\n", err)
	} else {
		g.Error = nil
	}
}
