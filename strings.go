package godis

// Append a value to a key.  Returns the length of the new string or -1 on error.
func (g *Godis) Append(key string, value string) int64 {
	return g.cmdInt("APPEND", key, value)
}

// BitCount counts the number of set bits in a string.
// Function can be invoked with a variable number of parameters:
// BITCOUNT key [start end]
func (g *Godis) BitCount(args ...interface{}) int64 {
	return g.cmdInt("BITCOUNT", args...)
}

// BitOp performs bitwise operations between strings.
// More than 1 source key can be specified.  Returns
// the size of the new key.
// BITOP AND destkey srckey1 [srckey2 srckey3 ... srckeyN]
func (g *Godis) BitOpAnd(args ...interface{}) int64 {
	return g.cmdInt("BITOP AND", args...)
}

// BitOp performs bitwise operations between strings.
// More than 1 source key can be specified.  Returns
// the size of the new key.
// BITOP OR destkey srckey1 [srckey2 srckey3 ... srckeyN]
func (g *Godis) BitOpNot(destkey, srckey string) int64 {
	return g.cmdInt("BITOP NOT", destkey, srckey)
}

// BitOp performs bitwise operations between strings.
// More than 1 source key can be specified.  Returns
// the size of the new key.
// BITOP OR destkey srckey1 [srckey2 srckey3 ... srckeyN]
func (g *Godis) BitOpOr(args ...interface{}) int64 {
	return g.cmdInt("BITOP OR", args...)
}

// BitOp performs bitwise operations between strings.
// More than 1 source key can be specified.  Returns
// the size of the new key.
// BITOP OR destkey srckey1 [srckey2 srckey3 ... srckeyN]
func (g *Godis) BitOpXor(args ...interface{}) int64 {
	return g.cmdInt("BITOP XOR", args...)
}

// BitPos returns the position of the first bit set to 1 or 0 in a string.
// Function can be invoked with a variable number of parameters:
// BITPOS key bit [start] [end]
func (g *Godis) BitPos(args ...interface{}) int64 {
	return g.cmdInt("BITPOS", args...)
}

// Decr decrements the integer value of a key by one.  Returns math.MaxInt64
// on error.
func (g *Godis) Decr(key string) int64 {
	return g.cmdInt("DECR", key)
}

// DecrBy decrements the integer value of a key by the given number.  Returns math.MaxInt64
// on error.
func (g *Godis) DecrBy(key string, decrement int) int64 {
	return g.cmdInt("DECRBY", key, decrement)
}

// Get the value of a key.  Return godis.NIL if key does not exist or upon error.  To
// tell the difference, look at godis.Error.
func (g *Godis) Get(key string) string {
	return g.cmdString("GET", key)
}

// GetBit returns the bit value at offset in the value stored at key.
// Returns -1 on error.
func (g *Godis) GetBit(key string, offset int) int64 {
	return g.cmdInt("GETBIT", key, offset)
}

// Get a substring of the string stored at key.
func (g *Godis) GetRange(key string, start, end int) string {
	return g.cmdString("GetRange", key, start, end)
}

// Set the string value of a key and return its old value.
func (g *Godis) GetSet(key, value string) string {
	return g.cmdString("GETSET", key, value)
}

// Incr increments the integer value of a key by one.  Returns math.MinInt64
// on error.
func (g *Godis) Incr(key string) int64 {
	return g.cmdInt("INCR", key)
}

// IncrBy increments the integer value of a key by the given number.  Returns math.MinInt64
// on error.
func (g *Godis) IncrBy(key string, increment int) int64 {
	return g.cmdInt("INCRBY", key, increment)
}

// IncrByFloat increments the float value of a key by the given amount.  Return math.MaxFloat64 on error.
func (g *Godis) IncrByFloat(key string, increment float64) float64 {
	return g.cmdFloat("INCRBYFLOAT", key, increment)
}

// MGet gets the values of all the given keys.
func (g *Godis) MGet(keys ...interface{}) []string {
	return g.cmdStrings("MGET", keys...)
}

// MSet sets multiple keys to multiple values.
func (g *Godis) MSet(keyvals ...interface{}) bool {
	return g.cmdString("MSET", keyvals...) == OK
}

// MSetNX sets multiple keys to multiple values, only if none of the keys already exist.
func (g *Godis) MSetNX(keyvals ...interface{}) bool {
	return g.cmdString("MSETNX", keyvals...) == OK
}

// Set the string value of a key with an expiration time in milliseconds.
func (g *Godis) PSetEX(key string, millis int, value string) bool {
	return g.cmdString("PSETEX", key, millis, value) == OK
}

// Set the string value of a key.
func (g *Godis) Set(key string, value string) bool {
	return g.cmdString("SET", key, value) == OK
}

// SetBit sets or clears the bit at offset in the string value stored at key.
// Returns -1 on error.
func (g *Godis) SetBit(key string, offset, value int) int64 {
	return g.cmdInt("SETBIT", key, offset, value)
}

// Set the string value of a key with an expiration time in seconds.
func (g *Godis) SetEX(key string, seconds int, value string) bool {
	return g.cmdString("SETEX", key, seconds, value) == OK
}

// Set the string value of a key if the key does not already exist.
func (g *Godis) SetNX(key string, value string) bool {
	return g.cmdInt("SETNX", key, value) == 1
}

// SetRange overwrites part of a string at key starting at the specified offset.
func (g *Godis) SetRange(key string, offset int, value string) int64 {
	return g.cmdInt("SETRANGE", key, offset, value)
}

// Strlen returns the length of the value stored at key.
func (g *Godis) Strlen(key string) int64 {
	return g.cmdInt("STRLEN", key)
}
