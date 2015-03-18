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

// ZCount counts the members in a sorted set with scores within the given values.
// Note that the values do not have to be integers, but could be strings like
// "-inf", "+inf", "(1".  Read the Redis doc.
func (g *Godis) ZCount(key, min, max string) int64 {
	return g.cmdInt("ZCOUNT", key, min, max)
}

// ZIncrBy increments the score of a member in a sorted set.
func (g *Godis) ZIncrBy(key string, increment float64, member string) float64 {
	return g.cmdFloat("ZINCRBY", key, increment, member)
}

// Skipped for now - ZInterStore

// ZLexCount counts the number of members in a sorted set between a
// given lexicographical range.
func (g *Godis) ZLexCount(key, min, max string) int64 {
	return g.cmdInt("ZLEXCOUNT", key, min, max)
}

// ZRange returns a range of members in a sorted set, by index
func (g *Godis) ZRange(key string, start, stop int) []string {
	return g.cmdStrings("ZRANGE", key, start, stop)
}

// ZRange returns a range of members and their scores in a sorted set, by index.
func (g *Godis) ZRangeWithScores(key string, start, stop int) []ScoreMember {
	retval := g.cmdStrings("ZRANGE", key, start, stop, "WITHSCORES")
	return memberScores(retval)
}

// ZRangeByLex return a range of members ina sorted set, by lexicographical range.
func (g *Godis) ZRangeByLex(key, min, max string) []string {
	return g.cmdStrings("ZRANGEBYLEX", key, min, max)
}

// ZRangeByLex return a range of members ina sorted set, by lexicographical range.
func (g *Godis) ZRangeByLexLimit(key, min, max, string, offset, count int) []string {
	return g.cmdStrings("ZRANGEBYLEX", key, min, max, "LIMIT", offset, count)
}

// ZRevRangeByLex return a range of members ina sorted set, by lexicographical range,
// ordered from higher to lower strings.
func (g *Godis) ZRevRangeByLex(key, max, min string) []string {
	return g.cmdStrings("ZREVRANGEBYLEX", key, max, min)
}

// ZRevRangeByLex return a range of members ina sorted set, by lexicographical range,
// ordered from higher to lower strings.
func (g *Godis) ZRevRangeByLexLimit(key, max, min, string, offset, count int) []string {
	return g.cmdStrings("ZRANGEBYLEX", key, max, min, "LIMIT", offset, count)
}

// ZRangeByScore returns a range of members in a sorted set, by score.
func (g *Godis) ZRangeByScore(key, min, max string) []string {
	return g.cmdStrings("ZRANGEBYSCORE", key, min, max)
}

// ZRangeByScoreLimit returns a range of members in a sorted set, by score,
// with LIMIT.
func (g *Godis) ZRangeByScoreLimit(key, min, max string, offset, count int) []string {
	return g.cmdStrings("ZRANGEBYSCORE", key, min, max, "LIMIT", offset, count)
}

// ZRangeByScoreWithScores return a range of members and their score in a sorted set,
// by score.
func (g *Godis) ZRangeByScoreWithScores(key, min, max string) []ScoreMember {
	retval := g.cmdStrings("ZRANGEBYSCORE", key, min, max, "WITHSCORES")
	return memberScores(retval)
}

// ZRangeByScoreWithScoresLimit returns a range of members and their score in a sorted set,
// by score, with LIMIT.
func (g *Godis) ZRangeByScoreWithScoresLimit(key, min, max string, offset, count int) []ScoreMember {
	retval := g.cmdStrings("ZRANGEBYSCORE", key, min, max, "WITHSCORES", "LIMIT", offset, count)
	return memberScores(retval)
}

// ZRank determines the index of a member in a sorted set.
func (g *Godis) ZRank(key, member string) int64 {
	return g.cmdInt("ZRANK", key, member)
}

// ZRem removes one or more members from a sorted set.
func (g *Godis) ZRem(key string, members ...interface{}) int64 {
	return g.cmdInt("ZREM", args1(key, members...)...)
}
