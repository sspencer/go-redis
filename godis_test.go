package godis

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"os"
	"testing"
	"time"
)

// look in redigo internal redistest
type testconn struct {
	*Godis
}

func (t *testconn) Close() error {
	fmt.Println("TESTDIS CLOSE")
	_, err := t.conn.Do("SELECT", "9")
	if err != nil {
		return nil
	}
	_, err = t.conn.Do("FLUSHDB")
	if err != nil {
		return err
	}
	return t.conn.Close()
}

// newTestconn dials the local Redis server and selects database 9. To prevent
// stomping on real data, DialTestDB fails if database 9 contains data. The
// returned connection flushes database 9 on close.
func newTestconn() *testconn {

	logger := log.New(os.Stderr, "", log.LstdFlags)
	conn, err := redis.DialTimeout("tcp", ":6379", 0, 1*time.Second, 1*time.Second)
	if err != nil {
		panic(err)
	}

	conn = redis.NewLoggingConn(conn, logger, "[TESTCONN] ")

	_, err = conn.Do("SELECT", "9")
	if err != nil {
		panic(err)
	}

	n, err := redis.Int(conn.Do("DBSIZE"))
	if err != nil {
		panic(err)
	}

	if n != 0 {
		panic(errors.New("database #9 is not empty, test can not continue"))
	}

	return &testconn{&Godis{false, conn, nil, logger, nil}}
}

// to just run this file: go test -run TestHash
func TestHashCommands(t *testing.T) {
	r := newTestconn()
	defer r.Close()
	key := "track123"
	nope := "nope:not:a:field"
	artist := "Hem"
	title := "Not California"
	album := "Funnel Cloud"
	plays := 21
	share := 0.5

	fields := []string{"artist", "title", "album", "plays", "share"}
	values := []interface{}{artist, title, album, plays, share}
	_ = r.HMSet(key,
		fields[0], values[0], // artist
		fields[1], values[1], // title
		fields[2], values[2], // album
		fields[3], values[3], // plays
		fields[4], values[4], // share
		"x1", 1,
		"x2", 2)

	f := r.HMGet(key, "artist", "title")
	if len(f) != 2 || f[0] != artist || f[1] != title {
		t.Fatal("Something wrong with HMSet or HMGet")
	}

	if cnt := r.HDel(key, "x1", "x2"); cnt != 2 {
		t.Fatal("Count not correct after HDel:", cnt)
	}

	if !r.HExists(key, "artist") || r.HExists(key, nope) {
		t.Fatal("HExists not working")
	}

	if r.HGet(key, "artist") != artist || r.HGet(key, nope) != EmptyString {
		t.Fatal("HGet not working")
	}

	fv := r.HGetAll(key)
	if len(fv) != len(fields) {
		t.Fatalf("HGetAll did not return correct number of fieldvals %d / %d\n", len(fv), len(fields))
	}

	fmt.Printf("%#v\n", fv)

	for index, field := range fields {

		if r1, ok := fv[field]; ok {
			if r1 != fmt.Sprintf("%v", values[index]) {
				t.Fatalf("HGetAll %s value not returned: '%s' == '%s'??\n", field, r1, values[index])
			}
		} else {
			t.Fatalf("HGetAll %s not returned\n", field)
		}
	}

	if r.HIncrBy(key, "plays", 10) != 31 {
		t.Fatalf("HIncrBy not working")
	}

	if r.HIncrBy(key, "plays", -10) != 21 {
		t.Fatalf("HIncrBy not working")
	}
}

func xTestListCommands(t *testing.T) {
	r := newTestconn()
	defer r.Close()

	cnt := r.LPush("mylist", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten")
	if cnt != 10 {
		t.Fatal("LPush returned incorrect item count")
	}

	list, item := r.BLPop("mylist", 1)
	if list != "mylist" || item != "ten" {
		t.Fatalf("BLPop returned %s:%s", list, item)
	}

	list, item = r.BRPop("mylist", 1)
	if list != "mylist" || item != "one" {
		t.Fatalf("BRPop returned %s:%s", list, item)
	}

	item = r.BRPopLPush("mylist", "otherlist", 1)
	if item != "two" {
		t.Fatal("BRPopLPush expected 'two', returned", item)
	}

	if r.LLen("mylist") != 7 || r.LLen("otherlist") != 1 {
		t.Fatal("Unexpected count after BRPopLPush/LLen")
	}

	if item = r.LIndex("mylist", 3); item != "six" {
		t.Fatal("LIndex returned:", item)
	}

	if cnt = r.LInsertAfter("mylist", "eight", "sevenfive"); cnt != 8 {
		t.Fatal("LInsertAfter returned:", cnt)
	}

	if cnt = r.LInsertBefore("mylist", "four", "fourfive"); cnt != 9 {
		t.Fatal("LInsertBefore return:", cnt)
	}

	if item = r.LPop("mylist"); item != "nine" {
		t.Fatal("LPop returned:", item)
	}

	if cnt = r.LPush("mylist", "foo", "bar"); cnt != 10 {
		t.Fatal("LPush returned:", cnt)
	}

	if cnt = r.LPushX("nolist", "nope"); cnt != 0 {
		t.Fatal("LPushX(1) returned:", cnt)
	}

	if cnt = r.LPushX("mylist", "alpha"); cnt != 11 {
		t.Fatal("LPushX(2) returned:", cnt)
	}

	items := r.LRange("mylist", 0, 1)
	if len(items) != 2 || items[0] != "alpha" || items[1] != "bar" {
		t.Fatal("LRange(1) unexpected value")
	}

	items = r.LRange("nononolist", 0, 1)
	if len(items) != 0 {
		t.Fatal("LRange(2) unexpected value")
	}

	if cnt = r.LRem("mylist", -2, "foo"); cnt != 1 {
		t.Fatal("LRem unexpected count:", cnt)
	}

	_ = r.LSet("mylist", 2, "berry")
	if item = r.LIndex("mylist", 2); item != "berry" {
		t.Fatal("LSet did not take:", item)
	}

	_ = r.LTrim("mylist", 3, 10)
	items = r.LRange("mylist", 0, -1)
	if len(items) != 7 || items[0] != "sevenfive" || items[6] != "three" {
		t.Fatal("LTrim unexpected value:", items)
	}

	if item = r.RPop("mylist"); item != "three" {
		t.Fatal("RPop unexpected value:", item)
	}

	if item = r.RPopLPush("mylist", "otherlist"); item != "four" {
		t.Fatal("RPopLPush unexpected value:", item)
	}

	items = r.LRange("otherlist", 0, -1)
	if len(items) != 2 || items[0] != "four" || items[1] != "two" {
		t.Fatal("Unexpected values in otherlist after RPopLPush")
	}

	if cnt = r.RPush("mylist", "tude"); cnt != 6 {
		t.Fatal("RPush unexpected count:", cnt)
	}

	if item = r.LIndex("mylist", 5); item != "tude" {
		t.Fatal("RPush new element not found:", item)
	}

	if cnt = r.RPushX("nolist", "hello"); cnt != 0 {
		t.Fatal("RPushX(1) unexpected value:", cnt)
	}

	if cnt = r.RPushX("otherlist", "hello"); cnt != 3 {
		t.Fatal("RPushX(2) unexpected value:", cnt)
	}
}
