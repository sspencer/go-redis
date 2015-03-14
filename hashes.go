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
