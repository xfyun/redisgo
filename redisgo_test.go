package redisgo

import (
	"testing"
	"time"
)

//测试对key值的操作
func TestKeyOp(t *testing.T) {
	var (
		addr        = "192.168.86.60:6379"
		pwd         = ""
		maxIdle     = 3
		idleTimeOut = time.Second * 200
		db          = 0
	)
	redisInst, err := NewRedisInst(addr, pwd, maxIdle, db, idleTimeOut)
	if err != nil {
		t.Errorf("create redis inst error:%s", err.Error())
		return
	}

	var key = "test_key"
	//SET
	if _, err = redisInst.Set(key, "hello"); err != nil {
		t.Errorf("SET error:%s", err.Error())
		return
	}

	//EXISTS
	if exists, err := redisInst.Exists(key); err != nil || !exists {
		t.Errorf("Exists error")
		return
	}

	//EXPIRE
	if err = redisInst.Expire(key, 30); err != nil {
		t.Errorf("EXPIRE error:%s", err.Error())
		return
	}

	//Ttl
	if expired, err := redisInst.Ttl(key); err != nil || expired < 0 {
		t.Errorf("Ttl error")
		return
	}

	//PExpire
	if err = redisInst.PExpire(key, 3000); err != nil {
		t.Errorf("PExpire error:%s", err.Error())
		return
	}

	//PTtl
	if expired, err := redisInst.PTtl(key); err != nil || expired < 0 {
		t.Errorf("PTtl error")
		return
	}

	//Persist
	if err = redisInst.Persist(key); err != nil {
		t.Errorf("Persist error:%s", err.Error())
		return
	}

	//Rename
	if err = redisInst.Rename(key, "test2"); err != nil {
		t.Errorf("Rename  error:%s", err.Error())
		return
	}

	//Del
	if err = redisInst.Del(key); err != nil {
		t.Errorf("Del  error:%s", err.Error())
		return
	}
}

//测试String类型操作
func TestStringOp(t *testing.T) {
	var (
		addr        = "192.168.86.60:6379"
		pwd         = ""
		maxIdle     = 3
		idleTimeOut = time.Second * 200
		db          = 0
	)

	redisInst, err := NewRedisInst(addr, pwd, maxIdle, db, idleTimeOut)
	if err != nil {
		t.Errorf("create redis inst error:%s", err.Error())
		return
	}

	var key = "test_string"

	//clean key
	exists, err := redisInst.Exists(key)
	if err != nil {
		t.Errorf("exists error:%s", err.Error())
		return
	}

	if exists {
		if err = redisInst.Del(key); err != nil {
			t.Errorf("Del error:%s", err.Error())
			return
		}
	}

	//SET(value is string)
	if _, err = redisInst.Set(key, "hello string"); err != nil {
		t.Errorf("Set error:%s", err.Error())
		return
	}

	//GETString
	if val, err := redisInst.GetString(key); err != nil || val != "hello string" {
		t.Errorf("GetString error")
		return
	}

	//SET(value is object)
	var a = map[string]int{"hello": 1}
	var b = make(map[string]int)
	if _, err = redisInst.Set(key, a); err != nil {
		t.Errorf("Set error:%s", err.Error())
		return
	}

	//SETEX
	if _, err = redisInst.SetEx(key, a, 30); err != nil {
		t.Errorf("SetEx error:%s", err.Error())
		return
	}

	//PSETEXS
	if _, err = redisInst.PSetEx(key, a, 3000); err != nil {
		t.Errorf("PSetEx error:%s", err.Error())
		return
	}

	//GetObject
	if err = redisInst.GetObject(key, &b); err != nil {
		t.Errorf("GetObject error:%s", err.Error())
		return
	}

	//SET(value is int)
	if _, err = redisInst.Set(key, 0); err != nil {
		t.Errorf("Set error:%s", err.Error())
		return
	}

	//Incr
	if valIncr, err := redisInst.Incr(key); err != nil && valIncr != 1 {
		t.Errorf("Incr error")
		return
	}

	//Decr
	if valDecr, err := redisInst.Decr(key); err != nil && valDecr != 0 {
		t.Errorf("Decr error")
		return
	}

	//IncrBy
	if valIncrBy, err := redisInst.IncrBy(key, 100); err != nil && valIncrBy != 100 {
		t.Errorf("valIncrBy error")
		return
	}

	//DecrBy
	if valDecrBy, err := redisInst.DecrBy(key, 100); err != nil && valDecrBy != 0 {
		t.Errorf("DecrBy error")
		return
	}
}

//hash操作测试
func TestHashOp(t *testing.T) {
	var (
		addr        = "192.168.86.60:6379"
		pwd         = ""
		maxIdle     = 3
		idleTimeOut = time.Second * 200
		db          = 0
	)

	redisInst, err := NewRedisInst(addr, pwd, maxIdle, db, idleTimeOut)
	if err != nil {
		t.Errorf("create redis inst error:%s", err.Error())
		return
	}

	var key = "test_hash"
	var field = "1"

	var exists bool
	//clean key
	if exists, err = redisInst.Exists(key); err != nil {
		t.Errorf("exists error:%s", err.Error())
		return
	}

	if exists {
		if err = redisInst.Del(key); err != nil {
			t.Errorf("Del error:%s", err.Error())
			return
		}
	}

	//HSET(value is string type)
	if _, err = redisInst.Hset(key, field, "hello"); err != nil {
		t.Errorf("HSET error:%s", err.Error())
		return
	}

	//HGET
	if _, err = redisInst.Hget(key, field); err != nil {
		t.Errorf("HGET error")
		return
	}

	//HgetString
	if valString, err := redisInst.HgetString(key, field); err != nil || valString != "hello" {
		t.Errorf("HgetString error")
		return
	}

	var a = map[int]string{1: "hero"}
	var b = make(map[int]string)

	//HSET(value is object type)
	if _, err = redisInst.Hset(key, field, a); err != nil {
		t.Errorf("HSET error:%s", err.Error())
		return
	}

	//HgetObject
	if err = redisInst.HgetObject(key, field, &b); err != nil {
		t.Errorf("HgetObject error:%s", err.Error())
		return
	}

	//Hmset
	var p1, p2 struct {
		Title  string `redis:"title"`
		Author string `redis:"author"`
		Body   string `redis:"body"`
	}

	p1.Title = "Example"
	p1.Author = "Gary"
	p1.Body = "Hello"

	if err = redisInst.Hmset(key, &p1); err != nil {
		t.Errorf("Hmset error:%s", err.Error())
		return
	}

	//Hgetall
	if err = redisInst.HgetAll(key, &p2); err != nil {
		t.Errorf("HgetAll error:%s", err.Error())
		return
	}

	//Hdel
	if err = redisInst.Hdel(key, field); err != nil {
		t.Errorf("Hdel error:%s", err.Error())
		return
	}

	//Hexists
	if exists, err = redisInst.Hexists(key, field); err != nil || exists {
		t.Errorf("Hexists error")
		return
	}
}

//list操作测试
func TestListOp(t *testing.T) {
	var (
		addr        = "192.168.86.60:6379"
		pwd         = ""
		maxIdle     = 3
		idleTimeOut = time.Second * 200
		db          = 0
	)

	redisInst, err := NewRedisInst(addr, pwd, maxIdle, db, idleTimeOut)
	if err != nil {
		t.Errorf("create redis inst error:%s", err.Error())
		return
	}

	var key = "test_list"

	//clean key
	exists, err := redisInst.Exists(key)
	if err != nil {
		t.Errorf("exists error:%s", err.Error())
		return
	}

	if exists {
		if err = redisInst.Del(key); err != nil {
			t.Errorf("Del error:%s", err.Error())
			return
		}
	}

	//Lpush
	if err = redisInst.Lpush(key, "11", "22", "33"); err != nil {
		t.Errorf("Lpush error:%s", err.Error())
		return
	}

	//Llen
	if len, err := redisInst.Llen(key); err != nil || len != 3 {
		t.Errorf("Llen error:%s", err.Error())
		return
	}

	//Lrange
	if _, err = redisInst.LrangeString(key, 0, 2); err != nil {
		t.Errorf("LrangeString error:%s", err.Error())
		return
	}

	//Lpop
	if lpopElem, err := redisInst.LpopString(key); err != nil || lpopElem != "33" {
		t.Errorf("Lpop error:%s", err.Error())
		return
	}

	if err = redisInst.Del(key); err != nil {
		t.Errorf("Del error:%s", err.Error())
		return
	}

	//Rpush
	if err = redisInst.Rpush(key, "11", "22", "33"); err != nil {
		t.Errorf("Rpush error:%s", err.Error())
		return
	}

	//Rpop
	if elem, err := redisInst.RpopString(key); err != nil || elem != "33" {
		t.Errorf("Rpop error:%s", err.Error())
		return
	}
}

//set操作测试
func TestSetOp(t *testing.T) {
	var (
		addr        = "192.168.86.60:6379"
		pwd         = ""
		maxIdle     = 3
		idleTimeOut = time.Second * 200
		db          = 0
	)
	redisInst, err := NewRedisInst(addr, pwd, maxIdle, db, idleTimeOut)
	if err != nil {
		t.Errorf("create redis inst error:%s", err.Error())
		return
	}

	var key = "test_set"

	//clean key
	exists, err := redisInst.Exists(key)
	if err != nil {
		t.Errorf("exists error:%s", err.Error())
		return
	}

	if exists {
		if err = redisInst.Del(key); err != nil {
			t.Errorf("Del error:%s", err.Error())
			return
		}
	}

	//Sadd
	if err = redisInst.Sadd(key, "11", "22", "33"); err != nil {
		t.Errorf("Sadd error:%s", err.Error())
		return
	}

	//Scard
	if memberNum, err := redisInst.Scard(key); err != nil || memberNum != 3 {
		t.Errorf("Scard error")
		return
	}

	//Sismember
	if isMember, err := redisInst.Sismember(key, "11"); err != nil || !isMember {
		t.Errorf("Sismember error：%s", err.Error())
		return
	}

	//Smembers
	if members, err := redisInst.Smembers(key); err != nil || len(members) != 3 {
		t.Errorf("Smembers error:%s", err.Error())
		return
	}

	//Srem
	if err = redisInst.Srem(key, "11", "22"); err != nil {
		t.Errorf("Srem error:%s", err.Error())
		return
	}

	var key1 = "test_set1"
	//Sadd
	if err = redisInst.Sadd(key1, "33", "66"); err != nil {
		t.Errorf("Sadd error:%s", err.Error())
		return
	}

	//Sunion
	if unionMembers, err := redisInst.Sunion(key, key1); err != nil || len(unionMembers) != 2 {
		t.Errorf("Sunion error:%s", err.Error())
		return
	}

	var key2 = "test_set2"
	//Suionstrore
	if err = redisInst.Suionstrore(key2, key, key1); err != nil {
		t.Errorf("Suionstrore error:%s", err.Error())
		return
	}

	//Sdiff
	if diffMembers, err := redisInst.Sdiff(key1, key); err != nil || len(diffMembers) != 1 {
		t.Errorf("Sdiff error:%s", err.Error())
		return
	}

	//Sdiffstore
	var key3 = "test_set3"
	if err = redisInst.Sdiffstore(key3, key, key1); err != nil {
		t.Errorf("Sdiffstore error:%s", err.Error())
		return
	}

	//Smove
	if err = redisInst.Smove(key1, key, "66"); err != nil {
		t.Errorf("Smove error:%s", err.Error())
		return
	}
}

//sorted set操作测试
func TestSortedSetOp(t *testing.T) {
	var (
		addr        = "192.168.86.60:6379"
		pwd         = ""
		maxIdle     = 3
		idleTimeOut = time.Second * 200
		db          = 0
	)
	redisInst, err := NewRedisInst(addr, pwd, maxIdle, db, idleTimeOut)
	if err != nil {
		t.Errorf("create redis inst error:%s", err.Error())
		return
	}

	var key = "test_sortedset"

	//clean key
	exists, err := redisInst.Exists(key)
	if err != nil {
		t.Errorf("exists error:%s", err.Error())
		return
	}

	if exists {
		if err = redisInst.Del(key); err != nil {
			t.Errorf("Del error:%s", err.Error())
			return
		}
	}

	//Zadd
	if err = redisInst.Zadd(key, 1, "11", 3, "22", 2, "33"); err != nil {
		t.Errorf("Sadd error:%s", err.Error())
		return
	}

	//Zcard
	if memberNum, err := redisInst.Zcard(key); err != nil || memberNum != 3 {
		t.Errorf("Zcard error")
		return
	}

	//Zrange
	if membersScore, err := redisInst.Zrange(key, 0, -1); err != nil || len(membersScore) != 3 {
		t.Errorf("Zrange error")
		return
	}

	//Zrevrange
	if membersScore, err := redisInst.Zrevrange(key, 0, -1); err != nil || len(membersScore) != 3 {
		t.Errorf("Zrevrange error")
		return
	}

	//Zcount
	if count, err := redisInst.Zcount(key, 1, 2); err != nil || count != 2 {
		t.Errorf("Zcount error")
		return
	}

	//Zrank
	if rank, err := redisInst.Zrank(key, "22"); err != nil || rank != 2 {
		t.Errorf("Zrank error")
		return
	}

	//Zrevrank
	if rank, err := redisInst.Zrevrank(key, "22"); err != nil || rank != 0 {
		t.Errorf("Zrevrank error")
		return
	}

	//Zscore
	if score, err := redisInst.Zscore(key, "22"); err != nil || score != 3 {
		t.Errorf("Zrevrank error")
		return
	}

	//ZremRangeByRank
	if err = redisInst.ZremRangeByRank(key, 0, 0); err != nil {
		t.Errorf("ZremRangeByRank error:%s", err.Error())
		return
	}

	//ZremRangeByRank
	if err = redisInst.ZremRangeByScore(key, 2, 2); err != nil {
		t.Errorf("ZremRangeByScore error:%s", err.Error())
		return
	}

	//Zrem
	if err = redisInst.Zrem(key, "22", "11"); err != nil {
		t.Errorf("Zrem error:%s", err.Error())
		return
	}
}

//发布和订阅测试
func TestMsgOp(t *testing.T) {
	var (
		addr        = "192.168.86.60:6379"
		pwd         = ""
		maxIdle     = 3
		idleTimeOut = time.Second * 200
		db          = 0
	)
	redisInst, err := NewRedisInst(addr, pwd, maxIdle, db, idleTimeOut)
	if err != nil {
		t.Errorf("create redis inst error:%s", err.Error())
		return
	}

	var chan1, chan2, chan3 = "test_chan1", "test_chan2", "test_chan3"

	//publish
	redisInst.Publish(chan1, "hello1")
	redisInst.Publish(chan1, "hello55")
	redisInst.Publish(chan2, "hello2")
	redisInst.Publish(chan3, "hello3")

	//subscribe
	go redisInst.Subscribe(chan1, chan2, chan3)
	time.Sleep(time.Second * 5)
}
