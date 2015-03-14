package godis

import (
	"github.com/garyburd/redigo/redis"
)

// LLen gets the length of a list.
func (g *Godis) LLen(key string) int {
	conn := g.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("LLEN", key)
	g.log.Printf("LLEN %s\n", key)

	if retval, err := redis.Int(reply, err); err != nil {
		// handle error
		g.Error = err
		g.log.Printf("Error LLEN %s\n", err)
		return 0
	} else {
		g.Error = nil
		return retval
	}
}
