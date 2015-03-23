package godis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"log"
	"time"
)

const (
	OK = "OK"
)

type Godis struct {
	pooled bool
	conn   redis.Conn
	pool   *redis.Pool
	log    *log.Logger
	err    error
}

type ScoreMember struct {
	score  float32
	member string
}

func (m ScoreMember) String() string {
	return fmt.Sprintf("(%s:%0.2f)", m.member, m.score)
}

// When a Redis call has no results or there's an error,
// return a shared empty variable.
var (
	EmptyInt          int64 = 0
	EmptyFloat              = 0.0
	EmptyString             = ""
	EmptyValues             = make([]interface{}, 0)
	EmptyStrings            = []string{}
	EmptyScoreMembers       = []ScoreMember{}
	EmptyStringMap          = make(map[string]string)
)

type dialer func() (redis.Conn, error)

func newDialer(server, password string, logger *log.Logger) dialer {
	return func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", server)
		if err != nil {
			return nil, err
		}

		if len(password) > 0 {
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
		}

		if logger != nil {
			loggingConn := redis.NewLoggingConn(c, logger, "[GODIS] ")
			return loggingConn, err
		} else {
			return c, err
		}

	}
}

func newPool(server, password string, logger *log.Logger) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        newDialer(server, password, logger),
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func NewGodisPool(server, password string, w io.Writer) *Godis {

	logger := log.New(w, "", log.LstdFlags)
	pool := newPool(server, password, logger)

	return &Godis{true, nil, pool, logger, nil}
}

func NewGodisConn(server, password string, dbIndex int, w io.Writer) *Godis {

	logger := log.New(w, "", log.LstdFlags)
	dial := newDialer(server, password, logger)
	conn, err := dial()
	if err != nil {
		panic(err)
	}

	// dbIndex is 0 by default, so only try to change if it's different
	if dbIndex > 0 {
		_, err = conn.Do("SELECT", dbIndex)
		if err != nil {
			panic(err)
		}
	}

	return &Godis{false, conn, nil, logger, nil}
}

func args1(v1 interface{}, values ...interface{}) []interface{} {
	args := make([]interface{}, 1)
	args[0] = v1
	args = append(args, values...)
	return args
}

func (g *Godis) Close() {
	if !g.pooled {
		g.conn.Close()
	}
}

// Reusuable redis function that matches pattern of sending a
// command and receiving an int.
func (g *Godis) cmdInt(cmd string, values ...interface{}) int64 {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do(cmd, values...)

	if retval, err := redis.Int64(reply, err); err != nil {
		g.err = err
		return EmptyInt
	} else {
		return retval
	}
}

// Reusuable redis function that matches pattern of sending a
// command and receiving an int.  Returns math.maxInt64 on error.
func (g *Godis) cmdFloat(cmd string, values ...interface{}) float64 {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do(cmd, values...)

	if retval, err := redis.Float64(reply, err); err != nil {
		g.err = err
		return EmptyFloat
	} else {
		return retval
	}
}

// Reusuable redis function that matches pattern of sending a
// command and receiving a string.
func (g *Godis) cmdString(cmd string, values ...interface{}) string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do(cmd, values...)

	if retval, err := redis.String(reply, err); err != nil {
		g.err = err
		return EmptyString
	} else {
		return retval
	}
}

// Reusuable redis function that matches pattern of sending a
// command and receiving a string.
func (g *Godis) cmdMap(cmd string, values ...interface{}) map[string]string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do(cmd, values...)

	if retval, err := redis.StringMap(reply, err); err != nil {
		g.err = err
		return EmptyStringMap
	} else {
		return retval
	}
}

// Reusuable redis function that matches pattern of sending a
// command and receiving a (string, string).
func (g *Godis) cmdStringString(cmd string, values ...interface{}) (string, string) {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do(cmd, values...)

	if retval, err := redis.Strings(reply, err); err != nil {
		g.err = err
		return EmptyString, EmptyString
	} else {
		return retval[0], retval[1]
	}
}

// Reusuable redis function that matches pattern of sending a
// command and receiving a string.
func (g *Godis) cmdStrings(cmd string, values ...interface{}) []string {
	var conn redis.Conn
	if g.pooled {
		conn = g.pool.Get()
		defer conn.Close()
	} else {
		conn = g.conn
	}

	reply, err := conn.Do(cmd, values...)

	if retval, err := redis.Strings(reply, err); err != nil {
		g.err = err
		return EmptyStrings
	} else {
		return retval
	}
}
