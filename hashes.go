package godis

import (
	"github.com/garyburd/redigo/redis"
	"math"
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

	reply, err := conn.Do("HEXISTS", key, field)
	g.log.Printf("HEXISTS %s %s\n", key, field)

	if retval, err := redis.Bool(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HEXISTS %s\n", err)
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

// HGetAll gets all the fields of values in a hash and returns in as
// a map
func (g *Godis) HGetAll(key string) map[string]string {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("HGETALL", key)
	g.log.Printf("HGETALL %s\n", key)

	if retval, err := redis.StringMap(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HGETALL %s\n", err)
		return make(map[string]string)

	} else {
		g.Error = nil
		return retval
	}
}

// HIncrBy increments the integer value of a hash field by the given number.  Returns math.MinInt64
// on error.
func (g *Godis) HIncrBy(key, field string, increment int) int64 {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("HINCRBY", key, field, increment)
	g.log.Printf("HINCRBY %s %s %d\n", key, field, increment)

	if retval, err := redis.Int64(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HINCRBY %s\n", err)
		return math.MaxInt64
	} else {
		g.Error = nil
		return retval
	}
}

// HIncrByFloat increments the float value of a key by the given amount.  Return math.MaxFloat64 on error.
func (g *Godis) HIncrByFloat(key, field string, value float64) float64 {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("HINCRBYFLOAT", key, field, value)
	g.log.Printf("HINCRBYFLOAT %s %s \"%f\"\n", key, field, value)

	if retval, err := redis.Float64(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HINCRBYFLOAT %s\n", err)
		return math.MaxFloat64
	} else {
		g.Error = nil
		return retval
	}
}

// HKeys gets all the field names in a hash
func (g *Godis) HKeys(key string) []string {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("HKEYS", key)
	g.log.Printf("HKEYS %s\n", key)

	if retval, err := redis.Strings(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HKEYS %s\n", err)
		return []string{}
	} else {
		g.Error = nil
		return retval
	}
}

// HLen gets the number of fields in a hash
func (g *Godis) HLen(key string) int {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("HLen", key)
	g.log.Printf("HLen %s\n", key)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HLen %s\n", err)
		return 0
	} else {
		g.Error = nil
		return retval
	}
}

// HMGet gets the values of all the given hash fields.
func (g *Godis) HMGet(key string, fields ...interface{}) []string {
	conn := g.pool.Get()
	defer conn.Close()

	keyarg := make([]interface{}, 1)
	keyarg[0] = key
	args := append(keyarg, fields...)
	reply, err := conn.Do("HMGET", args...)
	g.log.Printf("HMGET %v\n", args)

	if retval, err := redis.Strings(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HMGET %s\n", err)
		return []string{}
	} else {
		g.Error = nil
		return retval
	}
}

// HMSet sets multiple hash fields to multiple values
func (g *Godis) HMSet(key string, fieldvals ...interface{}) bool {
	conn := g.pool.Get()
	defer conn.Close()

	keyarg := make([]interface{}, 1)
	keyarg[0] = key
	args := append(keyarg, fieldvals...)
	reply, err := conn.Do("HMSET", args...)
	g.log.Printf("HMSET %v\n", args)

	if retval, err := redis.String(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HMSET %s\n", err)
		return false
	} else {
		g.Error = nil
		return retval == OK
	}
}

// HSet sets the string value of a hash field.
// Returns 1 if field is a new field in the hash and value was set.
// Returns 0 if field already exists in the hash and the value was updated.
// Returns -1 on error.
func (g *Godis) HSet(key, field, value string) int {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("HSET", key, field, value)
	g.log.Printf("HSET %s %s %s\n", key, field, value)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HSET %s\n", err)
		return -1
	} else {
		g.Error = nil
		return retval
	}
}

// HSetNX sets the string value of a hash field if the field does not already exist.
// Returns 1 if field is a new field in the hash and value was set.
// Returns 0 if field already exists in the hash and no operation was performed.
// Returns -1 on error.
func (g *Godis) HSetNX(key, field, value string) int {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("HSETNX", key, field, value)
	g.log.Printf("HSETNX %s %s %s\n", key, field, value)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HSETNX %s\n", err)
		return -1
	} else {
		g.Error = nil
		return retval
	}
}

/*
Since 3.2.0
// HStrlen gets the length of the value of a hash field.
func (g *Godis) HStrlen(key, field string) int {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("HSTRLEN", key, field)
	g.log.Printf("HSTRLEN %s %s\n", key, field)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HSTRLEN %s\n", err)
		return -1
	} else {
		g.Error = nil
		return retval
	}
}
*/

// HVals gets all the field values in a hash
func (g *Godis) HVals(key string) []string {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("HVALS", key)
	g.log.Printf("HVALS %s\n", key)

	if retval, err := redis.Strings(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error HVALS %s\n", err)
		return []string{}
	} else {
		g.Error = nil
		return retval
	}
}
