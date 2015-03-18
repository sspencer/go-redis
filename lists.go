package godis

// BLPop removes and gets the first element in a list or blocks
// until one is available.  Both the key and value are returned.
// You may pass in more than 1 key and BLPop returns an element
// from the first list with an available item.  Specify timeout
// in seconds.   Returns (godis.NIL, godis.NIL) on timeout or error.
// NOTE: argument order is: BLPop(key1 [, key2, key3...], timeout)
// Minimum is 1 key and timeout is mandatory.
func (g *Godis) BLPop(args ...interface{}) (string, string) {
	return g.cmdStringString("BLPOP", args...)
}

// BRPop removes and gets the last element in a list or blocks
// until one is available.  Both the key and value are returned.
// You may pass in more than 1 key and BRPop returns an element
// from the first list with an available item.  Specify timeout
// in seconds.  Returns (godis.NIL, godis.NIL) on timeout or error.
// NOTE: argument order is: BRPop(key1 [, key2, key3...], timeout)
// Minimum is 1 key and timeout is mandatory.
func (g *Godis) BRPop(args ...interface{}) (string, string) {
	return g.cmdStringString("BRPOP", args...)
}

// BRPopLPush pops a value from the source list, pushes it onto destination and
// returns that list element.  Blocks until timeout or element is available.
func (g *Godis) BRPopLPush(source, destination string, timeout int) string {
	return g.cmdString("BRPOPLPUSH", source, destination, timeout)
}

// LIndex gets an element from a list by its index.
func (g *Godis) LIndex(key string, index int) string {
	return g.cmdString("LINDEX", key, index)
}

// LInsertAfter inserts an element after another element in a list. Returns list length
// or -1 if no item was inserted.
func (g *Godis) LInsertAfter(key, pivot, value string) int64 {
	return g.cmdInt("LINSERT", key, "AFTER", pivot, value)
}

// LInsertAfter inserts an element before another element in a list.  Returns list length
// or -1 if no item was inserted.
func (g *Godis) LInsertBefore(key, pivot, value string) int64 {
	return g.cmdInt("LINSERT", key, "BEFORE", pivot, value)
}

// LLen gets the length of a list.
func (g *Godis) LLen(key string) int64 {
	return g.cmdInt("LLEN", key)
}

// LPop remove and get the first element in a list.
func (g *Godis) LPop(key string) string {
	return g.cmdString("LPOP", key)
}

// LPush prepends one or more items to a list and returns the length
// of the list after the operation.
func (g *Godis) LPush(key string, values ...interface{}) int64 {
	return g.cmdInt("LPUSH", args1(key, values...)...)
}

// LPushX prepends a value to a list, only if the list exists.  Returns the length
// of the list after the operation.
func (g *Godis) LPushX(key, value string) int64 {
	return g.cmdInt("LPUSHX", key, value)
}

// LRange gets a range of elements from a list.
func (g *Godis) LRange(key string, start, stop int) []string {
	return g.cmdStrings("LRANGE", key, start, stop)
}

// LRem removes elements from a list.  Returns the number of removed elements.
func (g *Godis) LRem(key string, count int, value string) int64 {
	return g.cmdInt("LREM", key, count, value)
}

// LSet sets the value of an element in a list by its index.
func (g *Godis) LSet(key string, index int, value string) bool {
	return g.cmdString("LSET", key, index, value) == OK
}

// LTrim trims a list to the specified range
func (g *Godis) LTrim(key string, start, stop int) bool {
	return g.cmdString("LTRIM", key, start, stop) == OK
}

// RPop remove and get the last element in a list.
func (g *Godis) RPop(key string) string {
	return g.cmdString("RPOP", key)
}

// RPopLPush removes the last element ina list and prepends it to another list and returns it.
// Can use the same list for both source and destination to process items in a circular fashion.
func (g *Godis) RPopLPush(source, destination string) string {
	return g.cmdString("RPOPLPUSH", source, destination)
}

// RPush appends one or more items to a list and returns the length
// of the list after the operation.
func (g *Godis) RPush(key string, values ...interface{}) int64 {
	return g.cmdInt("RPUSH", args1(key, values...)...)
}

// RPushX appends a value to a list, only if the list exists.  Returns the length
// of the list after the operation.
func (g *Godis) RPushX(key, value string) int64 {
	return g.cmdInt("RPUSHX", key, value)
}
