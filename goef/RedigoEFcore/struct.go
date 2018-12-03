package RedigoEFcore

import (
	"sync"

	EF "github.com/ReadyCore/goef/other"
	"github.com/gomodule/redigo/redis"
)

type RedisConn struct {
	ProxyAddress   string // host
	PassWord       string // 密碼
	MaxIdle        int    // 最大閒置連線
	MaxActive      int    // 最大連線
	ConnType       string //連線 模式 TCP ..
	ConnectTimeout int    //連線  Timeout
	ReadTimeout    int    //讀取 Timeout
	WriteTimeout   int    //寫入 Timeout
	DBnumber       int    // DB　總數
	Mode           string // 集群 or 一般
}
type RedisHelper struct {
	Hash  Hash
	List  List
	Set   Set
	Key   Key
	Other Other
	//Queue Queue
}
type Hash struct {
	DBnumber int
	work     work
	mode     string
	ConnPool []*redis.Pool
}

// func (r *Hash) Shared() *RedisConnModel {
// 	return &RedisConnModel{}
// }
func (r *Hash) connGet(i int) *redis.Pool {
	return r.ConnPool[i]
}

func (r *Hash) input(Action string, in []interface{}) *work {
	wo := r.work.constructor()
	wo.Mode = r.mode
	wo.hashInput.Input = in
	wo.hashInput.Action = Action
	wo.value = wo.hashInput
	switch r.mode {
	case single:
		{
			i := r.DBnumber
			for i == 0 {
				wo.ConnPool = append(wo.ConnPool, r.connGet(i))
				i--
			}

		}
	case Cluster:
		{
		}
	}
	return wo
}

func (r *Hash) HGET(table string, key string) *work {
	input, _ := EF.MergeValue("HGET", table, key, nil)
	return r.input("HGET", input)
}
func (r *Hash) HMGET(table string, key ...interface{}) *work {
	input, _ := EF.MergeValue("HMGET", table, nil, key...)
	return r.input("HMGET", input)
}

func (r *Hash) HGETALL(table string) *work {
	input, _ := EF.MergeValue("HGETALL", table, nil, nil)
	return r.input("HGETALL", input)
}

func (r *Hash) HDEL(table string, key string) *work {

	input, _ := EF.MergeValue("HDEL", table, key, nil)
	return r.input("HDEL", input)

}
func (r *Hash) HEXISTS(table string, key string) *work {

	input, _ := EF.MergeValue("HEXISTS", table, key, nil)
	return r.input("HEXISTS", input)

}
func (r *Hash) HSET(table string, key string, value ...interface{}) *work {

	input, _ := EF.MergeValue("HSET", table, key, value...)
	return r.input("HSET", input)
}

type Set struct {
	work     work
	mode     string
	DBnumber int
	ConnPool []*redis.Pool
	lock     sync.Mutex
}

func (r *Set) connGet(i int) *redis.Pool {
	return r.ConnPool[i]
}
func (r *Set) input(Action string, in []interface{}) *work {
	wo := r.work.constructor()
	wo.Mode = r.mode
	wo.setInput.Input = in
	wo.setInput.Action = Action
	wo.value = wo.setInput
	switch r.mode {
	case single:
		{
			i := r.DBnumber
			for i == 0 {
				wo.ConnPool = append(wo.ConnPool, r.connGet(i))
				i--
			}
		}
	case Cluster:
		{
		}
	}
	return wo

}

func (r *Set) SADD(table string, value ...interface{}) *work {
	input, _ := EF.MergeValue("SADD", table, nil, value...)
	return r.input("SADD", input)

}

func (r *Set) SCARD(table string, key string) *work {

	input, _ := EF.MergeValue("SCARD", table, key, nil)
	return r.input("SCARD", input)

}

func (r *Set) SDIFF(table string, key string) *work {
	input, _ := EF.MergeValue("SDIFF", table, key, nil)
	return r.input("SDIFF", input)

}

func (r *Set) SINTER(table string, key string) *work {
	input, _ := EF.MergeValue("SINTER", table, key, nil)
	return r.input("SINTER", input)
}

func (r *Set) SISMEMBER(table string, key string) *work {
	input, _ := EF.MergeValue("SISMEMBER", table, key, nil)
	return r.input("SISMEMBER", input)
}

func (r *Set) SMEMBERS(table string) *work {
	input, _ := EF.MergeValue("SMEMBERS", table, nil, nil)
	return r.input("SMEMBERS", input)

}

func (r *Set) SSCAN(table string, key string) *work {
	input, _ := EF.MergeValue("SSCAN", table, key, nil)
	return r.input("SSCAN", input)
}

type List struct {
	work     work
	DBnumber int
	ConnPool []*redis.Pool
	mode     string
}

func (r *List) connGet(i int) *redis.Pool {
	return r.ConnPool[i]
}
func (r *List) input(Action string, in []interface{}) *work {
	wo := r.work.constructor()
	wo.Mode = r.mode
	wo.listInput.Input = in
	wo.listInput.Action = Action
	wo.value = wo.listInput
	switch r.mode {
	case single:
		{
			i := r.DBnumber
			for i == 0 {
				wo.ConnPool = append(wo.ConnPool, r.connGet(i))
				i--
			}
		}
	case Cluster:
		{
		}
	}
	return wo
}
func (r *List) LSET(table string, index string, value ...interface{}) *work {
	input, _ := EF.MergeValue("LSET", table, index, value)
	return r.input("LSET", input)
}
func (r *List) LPUSH(table string, value interface{}) *work {

	input, _ := EF.MergeValue("LPUSH", table, value, nil)
	return r.input("LPUSH", input)
}
func (r *List) LINDEX(table string, index string) *work {
	input, _ := EF.MergeValue("LINDEX", table, index, nil)
	return r.input("LINDEX", input)
}
func (r *List) LLEN(table string) *work {
	input, _ := EF.MergeValue("LLEN", table, nil, nil)
	return r.input("LLEN", input)
}

type Key struct {
	work     work
	DBnumber int
	ConnPool []*redis.Pool
	mode     string
}

func (r *Key) connGet(i int) *redis.Pool {
	return r.ConnPool[i]
}

func (r *Key) input(Action string, in []interface{}) *work {
	wo := r.work.constructor()
	wo.Mode = r.mode
	wo.keyInput.Input = in
	wo.keyInput.Action = Action
	wo.value = wo.keyInput
	switch r.mode {
	case single:
		{
			i := r.DBnumber
			for i == 0 {
				wo.ConnPool = append(wo.ConnPool, r.connGet(i))
				i--
			}
		}
	case Cluster:
		{

		}
	}
	return wo
}
func (r *Key) DEL(table string) *work {
	input, _ := EF.MergeValue("DEL", table, nil, nil)
	return r.input("DEL", input)
}
func (r *Key) EXPIRE(table string, time int) *work {
	input, _ := EF.MergeValue("EXPIRE", table, nil, time)
	return r.input("EXPIRE", input)
}

type Other struct {
	work     work
	DBnumber int
	ConnPool []*redis.Pool
	mode     string
}

func (r *Other) connGet(i int) *redis.Pool {
	return r.ConnPool[i]
}

func (r *Other) input(Action string, in []interface{}) *work {
	wo := r.work.constructor()
	wo.Mode = r.mode
	wo.otherInput.Input = in
	wo.otherInput.Action = Action
	wo.value = wo.otherInput
	switch r.mode {
	case single:
		{
			i := r.DBnumber
			for i == 0 {
				wo.ConnPool = append(wo.ConnPool, r.connGet(i))
				i--
			}
		}
	case Cluster:
		{
		}
	}
	return wo
}
func (r *Other) SCAN(table string) *work {
	input, _ := EF.MergeValue("SCAN", table, nil, nil)
	return r.input("SCAN", input)
}
