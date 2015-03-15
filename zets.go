package godis

import (
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

// ZCard returns the number of members in a sorted set.
func (g *Godis) ZCard(key string) int64 {
	return g.cmdInt("ZCARD", key)
}

// ZRange returns a range of members in a sorted set, by index
func (g *Godis) ZRange(key string, start, stop int) []string {
	return g.cmdStrings("ZRANGE", start, stop)
}

// ZRange returns a range of members and their scores in a sorted set, by index.
func (g *Godis) ZRangeWithScores(key string, start, stop int) []ScoreMember {
	retval := g.cmdStrings("ZRANGE", key, start, stop, "WITHSCORES")
	return memberScores(retval)
}
