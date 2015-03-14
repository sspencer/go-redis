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
		return 0
	} else {
		return retval
	}
}

func (g *Godis) setOp(cmd, key string, keys ...interface{}) []string {
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
	reply, err := conn.Do(cmd, args...)

	if retval, err := redis.Strings(reply, err); err != nil {
		return EmptyStrings
	} else {
		return retval
	}
}

func (g *Godis) setOpStore(cmd, destination, key string, keys ...interface{}) int {
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
	reply, err := conn.Do(cmd, args...)

	if retval, err := redis.Int(reply, err); err != nil {
		return 0
	} else {
		return retval
	}
}

// SDiff returns the members of the set resulting from the difference between
// the first set and all the successive sets.
func (g *Godis) SDiff(key string, keys ...interface{}) []string {
	return g.setOp("SDIFF", key, keys...)
}

// SDiffStore stores the members of the set resulting from the difference between
// the first set and all the successive sets into destination and returns the count
// of members in destination.
func (g *Godis) SDiffStore(destination, key string, keys ...interface{}) int {
	return g.setOpStore("SDIFFSTORE", destination, key, keys...)
}

// SInter returns the members of the set resulting from the intersection of all the
// given sets.
func (g *Godis) SInter(key string, keys ...interface{}) []string {
	return g.setOp("SINTER", key, keys...)
}

// SInterStore stores the members of the set resulting from the intersection of all the
// given sets into destination.
func (g *Godis) SInterStore(destination, key string, keys ...interface{}) int {
	return g.setOpStore("SINTERSTORE", destination, key, keys...)
}

// SIsMember returns if a member is a member of the set stored at key
func (g *Godis) SIsMember(key, member string) bool {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("SISMEMBER", key, member)

	if retval, err := redis.Int(reply, err); err != nil {
		return false
	} else {
		return retval == 1
	}
}

// SMembers returns all members in a set
func (g *Godis) SMembers(key string) []string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("SMEMBERS", key)

	if retval, err := redis.Strings(reply, err); err != nil {
		return EmptyStrings
	} else {
		return retval
	}
}

// SMove moves an element from one list to another.  Returns true if element
// was moved, false if element is not a member of source and no operation was performed.
func (g *Godis) SMove(source, destination, member string) bool {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("SMOVE", source, destination, member)

	if retval, err := redis.Int(reply, err); err != nil {
		return false
	} else {
		return retval == 1
	}
}

// SPop removes and returns one or more random members from a set.
func (g *Godis) SPop(key string) string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("SPOP", key)

	if retval, err := redis.String(reply, err); err != nil {
		return ""
	} else {
		return retval
	}
}

// SRandMember returns a random member for a set
func (g *Godis) SRandMember(key string) string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("SRANDMEMBER", key)

	if retval, err := redis.String(reply, err); err != nil {
		return ""
	} else {
		return retval
	}
}

// SRem removes one or more members from a set and returns
// the number of elements removed.
func (g *Godis) SRem(key string, members ...interface{}) int {
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
	reply, err := conn.Do("SREM", args...)

	if retval, err := redis.Int(reply, err); err != nil {
		return 0
	} else {
		return retval
	}
}

// SUnion returns the members of the set resulting from the union of all the given sets.
func (g *Godis) SUnion(key string, keys ...interface{}) []string {
	return g.setOp("SUNION", key, keys...)
}

// SUnionStore stores the members of the set resulting from the union between
// the first set and all the successive sets into destination and returns the count
// of members in destination.
func (g *Godis) SUnionStore(destination, key string, keys ...interface{}) int {
	return g.setOpStore("SUNIONSTORE", destination, key, keys...)
}
