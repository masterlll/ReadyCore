package RedigoEFcore

import (
	"fmt"
	"sync"

	EF "github.com/ReadyCore/goef/other"
)

// type RedisHelper struct {
// 	Hash  Hash
// 	List  List
// 	Set   Set
// 	Key   Key
// 	Other Other
// 	Queue Queue
// }
type Hash struct {
	DBnumber int
	work     work
	mode     string
	connkey  string
}

func (r *Hash) input(Action string, in []interface{}) *work {
	wo := r.work.constructor()
	wo.Mode = r.mode
	wo.hashInput.Input = in
	wo.hashInput.Action = Action
	wo.value = wo.hashInput
	wo.connKey = r.connkey
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
func (r *Hash) HVALS(table string) *work {
	input, _ := EF.MergeValue("HVALS", table, nil, nil)
	return r.input("HVALS", input)
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

func (r *Hash) HLEN(table string) *work {

	input, _ := EF.MergeValue("HLEN", table, nil, nil)

	return r.input("HLEN", input)

}

type Set struct {
	work     work
	mode     string
	DBnumber int
	connkey  string
	lock     sync.Mutex
}

func (r *Set) input(Action string, in []interface{}) *work {
	wo := r.work.constructor()
	wo.Mode = r.mode
	wo.setInput.Input = in
	wo.setInput.Action = Action
	wo.value = wo.setInput
	wo.connKey = r.connkey

	return wo

}

func (r *Set) SADD(table string, value ...interface{}) *work {
	input, _ := EF.MergeValue("SADD", table, nil, value...)
	return r.input("SADD", input)

}
func (r *Set) SREM(table string, value ...interface{}) *work {
	input, _ := EF.MergeValue("SREM", table, nil, value...)
	return r.input("SREM", input)

}

func (r *Set) SCARD(table string, key string) *work {

	input, _ := EF.MergeValue("SCARD", table, key, nil)
	return r.input("SCARD", input)

}

func (r *Set) SPOP(Key string) *work {
	input, _ := EF.MergeValue("SPOP", Key, nil, nil)
	return r.input("SPOP", input)
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
	connkey  string
	mode     string
}

func (r *List) input(Action string, in []interface{}) *work {
	wo := r.work.constructor()
	wo.Mode = r.mode
	wo.listInput.Input = in
	wo.listInput.Action = Action
	wo.value = wo.listInput
	wo.connKey = r.connkey

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
func (r *List) RPUSH(table string, value ...interface{}) *work {

	input, _ := EF.MergeValue("RPUSH", table, nil, value...)
	return r.input("RPUSH", input)
}

func (r *List) LPOP(table string) *work {
	input, _ := EF.MergeValue("LPOP", table, nil, nil)
	return r.input("LPOP", input)
}
func (r *List) RPOP(table string) *work {
	input, _ := EF.MergeValue("RPOP", table, nil, nil)
	return r.input("RPOP", input)
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
	connkey  string
	mode     string
}

func (r *Key) input(Action string, in []interface{}) *work {
	wo := r.work.constructor()
	wo.Mode = r.mode
	wo.keyInput.Input = in
	wo.keyInput.Action = Action
	wo.value = wo.keyInput
	wo.connKey = r.connkey

	return wo
}

func (r *Key) SET(table string, Value string) *work {
	input, _ := EF.MergeValue("SET", table, Value, nil)
	return r.input("SET", input)
}
func (r *Key) DEL(table string) *work {
	input, _ := EF.MergeValue("DEL", table, nil, nil)
	return r.input("DEL", input)
}

func (r *Key) INCRBY(table string, number int) *work {
	input, _ := EF.MergeValue("INCRBY", table, number, nil)
	return r.input("INCRBY", input)
}
func (r *Key) EXPIRE(table string, time int) *work {
	input, _ := EF.MergeValue("EXPIRE", table, nil, time)
	return r.input("EXPIRE", input)
}
func (r *Key) EXISTS(key string) *work {
	input, _ := EF.MergeValue("EXISTS", key, nil, nil)
	return r.input("EXISTS", input)
}

func (r *Key) GET(table string) *work {
	input, _ := EF.MergeValue("GET", table, nil, nil)
	return r.input("GET", input)
}

func (r *Key) GETSET(table string, number int) *work {
	input, _ := EF.MergeValue("GETSET", number, nil, nil)

	return r.input("GETSET", input)
}

func (r *Key) DECRBY(table string, number int) *work {
	input, _ := EF.MergeValue("DECRBY", table, number, nil)
	return r.input("DECRBY", input)
}

type Other struct {
	work     work
	DBnumber int
	connkey  string
	mode     string
}

func (r *Other) input(Action string, in []interface{}) *work {
	wo := r.work.constructor()
	wo.Mode = r.mode
	wo.otherInput.Input = in
	wo.otherInput.Action = Action
	wo.value = wo.otherInput
	wo.connKey = r.connkey

	return wo
}
func (r *Other) SCAN(table string) *work {
	input, _ := EF.MergeValue("SCAN", table, nil, nil)
	return r.input("SCAN", input)
}

type Queue struct {
	mode    string
	connkey string
	convent convent
	lock    sync.Mutex
}

func (p *Queue) QueuePipe(DBnumber int, twice []Container) chan *convent {
	rd := dbContext{}
	ch := make(chan *convent)
	ok := make(chan bool)

	go func() {
		for o := range ok {
			fmt.Println(o)
			close(ch)
		}
	}()

	go func() {
		for i := range rd.Pipe(p.connkey, p.mode, DBnumber, twice...) {
			a := p.convent.constructor()
			a.value = i
			ch <- &a
		}
		ok <- true
		close(ok)

	}()

	return ch
}

// if p.hashInput.Input != nil {
// 	for i := range rd.Pipe(p.connKey, p.Mode, dbnumber, p.hashInput) {
// 		a.value = i
// 	}
// 	return &a
// }
