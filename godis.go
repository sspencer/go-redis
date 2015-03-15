package godis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"log"
	"math"
	"time"
)

const (
	OK  = "OK"
	NIL = "(*GODIS NIL TOKEN*)"
)

type Godis struct {
	pooled bool
	conn   redis.Conn
	pool   *redis.Pool
	log    *log.Logger
}

type ScoreMember struct {
	score  float32
	member string
}

func (m ScoreMember) String() string {
	return fmt.Sprintf("(%s:%0.2f)", m.member, m.score)
}

// When a Redis call has not results, return a shared empty variable.
var (
	EmptyValues       = make([]interface{}, 0)
	EmptyStrings      = []string{}
	EmptyScoreMembers = []ScoreMember{}
	EmptyStringMap    = make(map[string]string)
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

	return &Godis{true, nil, pool, logger}
}

func NewGodisConn(server, password string, w io.Writer) *Godis {

	logger := log.New(w, "", log.LstdFlags)
	dial := newDialer(server, password, logger)
	conn, err := dial()
	if err != nil {
		panic(err)
	}
	return &Godis{false, conn, nil, logger}
}

func args1(v1 interface{}, values ...interface{}) []interface{} {
	args := make([]interface{}, 1)
	args[0] = v1
	args = append(args, values...)
	return args
}

func args2(v1, v2 interface{}, values ...interface{}) []interface{} {
	args := make([]interface{}, 2)
	args[0] = v1
	args[1] = v2
	args = append(args, values...)
	return args
}

func args3(v1, v2, v3 interface{}, values ...interface{}) []interface{} {
	args := make([]interface{}, 1)
	args[0] = v1
	args[1] = v2
	args[2] = v3
	args = append(args, values...)
	return args
}

// Reusuable redis function that matches pattern of sending a
// command and receiving an int.  Returns math.maxInt64 on error.
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
		return math.MaxInt64
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
		return math.MaxFloat64
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
		return NIL
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
		return NIL, NIL
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
		return EmptyStrings
	} else {
		return retval
	}
}
