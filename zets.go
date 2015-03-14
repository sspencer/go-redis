package godis

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
)

// From string array with ["member1", "score1", "member2", "score2"], return
// array of ScoreMembers.
func memberScores(values []string) []ScoreMember {
	members := make([]ScoreMember, len(values)/2)
	index := 0
	for i := 0; i < len(values); i += 2 {
		member := values[i]
		score, _ := strconv.ParseFloat(values[i+1], 32)
		members[index] = ScoreMember{score: float32(score), member: member}
		index++
	}

	return members
}

// ZRange returns a range of members in a sorted set, by index
func (g *Godis) ZRange(key string, start, stop int) []string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("ZRANGE", key, start, stop)

	if retval, err := redis.Strings(reply, err); err != nil {
		return EmptyStrings
	} else {
		return retval
	}
}

// ZRange returns a range of members and their scores in a sorted set, by index.
func (g *Godis) ZRangeWithScores(key string, start, stop int) []ScoreMember {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do("ZRANGE", key, start, stop, "WITHSCORES")

	if retval, err := redis.Strings(reply, err); err != nil {
		return EmptyScoreMembers
	} else {
		return memberScores(retval)
	}
}
