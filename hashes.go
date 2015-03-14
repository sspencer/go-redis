package godis

import (
	"github.com/garyburd/redigo/redis"
	"math"
)

// HDel deletes one of more hash fields.
func (g *Godis) HDel(key string, fields ...interface{}) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	keyarg := make([]interface{}, 1)
	keyarg[0] = key
	args := append(keyarg, fields...)
	reply, err := conn.Do("HDEL", args...)

	if retval, err := redis.Int(reply, err); err != nil {
		return 0
	} else {
		return retval
	}
}

// HExists determines if a hash field exists
func (g *Godis) HExists(key, field string) bool {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("HEXISTS", key, field)

	if retval, err := redis.Bool(reply, err); err != nil {
		return false
	} else {
		return retval
	}
}

// HGet gets the value of a hash field
func (g *Godis) HGet(key, field string) string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("HGET", key, field)

	if retval, err := redis.String(reply, err); err != nil {
		return NIL
	} else {
		return retval
	}
}

// HGetAll gets all the fields of values in a hash and returns in as
// a map
func (g *Godis) HGetAll(key string) map[string]string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("HGETALL", key)

	if retval, err := redis.StringMap(reply, err); err != nil {
		return make(map[string]string)

	} else {
		return retval
	}
}

// HIncrBy increments the integer value of a hash field by the given number.  Returns math.MinInt64
// on error.
func (g *Godis) HIncrBy(key, field string, increment int) int64 {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("HINCRBY", key, field, increment)

	if retval, err := redis.Int64(reply, err); err != nil {
		return math.MaxInt64
	} else {
		return retval
	}
}

// HIncrByFloat increments the float value of a key by the given amount.  Return math.MaxFloat64 on error.
func (g *Godis) HIncrByFloat(key, field string, value float64) float64 {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("HINCRBYFLOAT", key, field, value)

	if retval, err := redis.Float64(reply, err); err != nil {
		return math.MaxFloat64
	} else {
		return retval
	}
}

// HKeys gets all the field names in a hash
func (g *Godis) HKeys(key string) []string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("HKEYS", key)

	if retval, err := redis.Strings(reply, err); err != nil {
		return []string{}
	} else {
		return retval
	}
}

// HLen gets the number of fields in a hash
func (g *Godis) HLen(key string) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("HLen", key)

	if retval, err := redis.Int(reply, err); err != nil {
		return 0
	} else {
		return retval
	}
}

// HMGet gets the values of all the given hash fields.
func (g *Godis) HMGet(key string, fields ...interface{}) []string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	keyarg := make([]interface{}, 1)
	keyarg[0] = key
	args := append(keyarg, fields...)
	reply, err := conn.Do("HMGET", args...)

	if retval, err := redis.Strings(reply, err); err != nil {
		return []string{}
	} else {
		return retval
	}
}

// HMSet sets multiple hash fields to multiple values
func (g *Godis) HMSet(key string, fieldvals ...interface{}) bool {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	keyarg := make([]interface{}, 1)
	keyarg[0] = key
	args := append(keyarg, fieldvals...)
	reply, err := conn.Do("HMSET", args...)

	if retval, err := redis.String(reply, err); err != nil {
		return false
	} else {
		return retval == OK
	}
}

// HSet sets the string value of a hash field.
// Returns 1 if field is a new field in the hash and value was set.
// Returns 0 if field already exists in the hash and the value was updated.
// Returns -1 on error.
func (g *Godis) HSet(key, field, value string) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("HSET", key, field, value)

	if retval, err := redis.Int(reply, err); err != nil {
		return -1
	} else {
		return retval
	}
}

// HSetNX sets the string value of a hash field if the field does not already exist.
// Returns 1 if field is a new field in the hash and value was set.
// Returns 0 if field already exists in the hash and no operation was performed.
// Returns -1 on error.
func (g *Godis) HSetNX(key, field, value string) int {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("HSETNX", key, field, value)

	if retval, err := redis.Int(reply, err); err != nil {
		return -1
	} else {
		return retval
	}
}

// HVals gets all the field values in a hash
func (g *Godis) HVals(key string) []string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("HVALS", key)

	if retval, err := redis.Strings(reply, err); err != nil {
		return []string{}
	} else {
		return retval
	}
}
