package godis

// HDel deletes one of more hash fields.
func (g *Godis) HDel(key string, fields ...interface{}) int64 {
	return g.cmdInt("HDEL", args1(key, fields...)...)
}

// HExists determines if a hash field exists
func (g *Godis) HExists(key, field string) bool {
	return g.cmdInt("HEXISTS", key, field) == 1
}

// HGet gets the value of a hash field
func (g *Godis) HGet(key, field string) string {
	return g.cmdString("HGET", key, field)
}

// HGetAll gets all the fields of values in a hash and returns in as
// a map
func (g *Godis) HGetAll(key string) map[string]string {
	return g.cmdMap("HGETALL", key)
}

// HIncrBy increments the integer value of a hash field by the given number.
func (g *Godis) HIncrBy(key, field string, increment int) int64 {
	return g.cmdInt("HINCRBY", key, field, increment)
}

// HIncrByFloat increments the float value of a key by the given amount.
func (g *Godis) HIncrByFloat(key, field string, increment float64) float64 {
	return g.cmdFloat("HINCRBYFLOAT", key, field, increment)
}

// HKeys gets all the field names in a hash
func (g *Godis) HKeys(key string) []string {
	return g.cmdStrings("HKEYS", key)
}

// HLen gets the number of fields in a hash
func (g *Godis) HLen(key string) int64 {
	return g.cmdInt("HLEN", key)
}

// HMGet gets the values of all the given hash fields.
func (g *Godis) HMGet(key string, fields ...interface{}) []string {
	return g.cmdStrings("HMGET", args1(key, fields...)...)
}

// HMSet sets multiple hash fields to multiple values
func (g *Godis) HMSet(key string, fieldvals ...interface{}) bool {
	return g.cmdString("HMSET", args1(key, fieldvals...)...) == OK
}

// HSet sets the string value of a hash field.
// Returns 1 if field is a new field in the hash and value was set.
// Returns 0 if field already exists in the hash and the value was updated.
func (g *Godis) HSet(key, field, value string) int64 {
	return g.cmdInt("HSET", key, field, value)
}

// HSetNX sets the string value of a hash field if the field does not already exist.
// Returns 1 if field is a new field in the hash and value was set.
// Returns 0 if field already exists in the hash and no operation was performed.
func (g *Godis) HSetNX(key, field, value string) int64 {
	return g.cmdInt("HSETNX", key, field, value)
}

// HVals gets all the field values in a hash
func (g *Godis) HVals(key string) []string {
	return g.cmdStrings("HVALS", key)
}
