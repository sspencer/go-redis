package godis

// SAdd adds one or more members to a set.  Returns the number
// of elements added to the set (not including members already present).
func (g *Godis) SAdd(key string, members ...interface{}) int64 {
	return g.cmdInt("SADD", args1(key, members...)...)
}

// SCard returns the number of members in a set.
func (g *Godis) SCard(key string) int64 {
	return g.cmdInt("SCARD", key)
}

// SDiff returns the members of the set resulting from the difference between
// the first set and all the successive sets.
func (g *Godis) SDiff(keys ...interface{}) []string {
	return g.cmdStrings("SDIFF", keys...)
}

// SDiffStore stores the members of the set resulting from the difference between
// the first set and all the successive sets into destination and returns the count
// of members in destination.
func (g *Godis) SDiffStore(destination string, keys ...interface{}) int64 {
	return g.cmdInt("SDIFFSTORE", args1(destination, keys...)...)
}

// SInter returns the members of the set resulting from the intersection of all the
// given sets.
func (g *Godis) SInter(keys ...interface{}) []string {
	return g.cmdStrings("SINTER", keys...)
}

// SInterStore stores the members of the set resulting from the intersection of all the
// given sets into destination.
func (g *Godis) SInterStore(destination string, keys ...interface{}) int64 {
	return g.cmdInt("SINTERSTORE", args1(destination, keys...)...)
}

// SIsMember returns if a member is a member of the set stored at key
func (g *Godis) SIsMember(key, member string) bool {
	return g.cmdInt("SISMEMBER", key, member) == 1
}

// SMembers returns all members in a set
func (g *Godis) SMembers(key string) []string {
	return g.cmdStrings("SMEMBERS", key)
}

// SMove moves an element from one list to another.  Returns true if element
// was moved, false if element is not a member of source and no operation was performed.
func (g *Godis) SMove(source, destination, member string) bool {
	return g.cmdInt("SMOVE", source, destination, member) == 1
}

// SPop removes and returns one or more random members from a set.
func (g *Godis) SPop(key string) string {
	return g.cmdString("SPOP", key)
}

// SRandMember returns a random member for a set
func (g *Godis) SRandMember(key string) string {
	return g.cmdString("SRANDMEMBER", key)
}

// SRem removes one or more members from a set and returns
// the number of elements removed.
func (g *Godis) SRem(key string, members ...interface{}) int64 {
	return g.cmdInt("SREM", args1(key, members...)...)
}

// SUnion returns the members of the set resulting from the union of all the given sets.
func (g *Godis) SUnion(keys ...interface{}) []string {
	return g.cmdStrings("SUNION", keys...)
}

// SUnionStore stores the members of the set resulting from the union between
// the first set and all the successive sets into destination and returns the count
// of members in destination.
func (g *Godis) SUnionStore(destination string, keys ...interface{}) int64 {
	return g.cmdInt("SUNIONSTORE", args1(destination, keys...)...)
}
