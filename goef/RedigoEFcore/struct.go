package RedigoEFcore

//"fmt"

import (
	"fmt"
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
	DBnumber       int    // DB號碼
	Mode           string // 集群 or 一般
}
type RedisHelper struct {
	Pool    *redis.Pool
	poolist []*redis.Pool
	Hash    Hash
	List    List
	Set     Set
	Key     Key
	Other   Other
	/// Q
	//Queue Queue
}

// type RDState struct {
// 	Err   error
// 	Stats string
// }

type Hash struct {
	work work
	mode string
}

func (r *Hash) HGET(table string, key string) *work {

	input, _ := EF.MergeValue("HGET", table, key, nil)

	//r.lock.Lock()

	switch r.mode {
	case single:
		{

		}
	case Cluster:
		{

		}

	}
	b := r.work.constructor()

	b.lock.Lock()

	b.hashInput.Input = input
	b.hashInput.Action = "HGET"
	b.value = b.hashInput
	b.lock.Unlock()
	//	defer r.lock.Unlock()
	return b

}
func (r *Hash) HMGET(table string, key ...interface{}) *work {

	input, _ := EF.MergeValue("HMGET", table, nil, key...)

	//r.lock.Lock()
	b := r.work.constructor()
	b.lock.Lock()

	b.hashInput.Input = input
	b.hashInput.Action = "HMGET"
	b.value = b.hashInput
	b.lock.Unlock()
	//	defer r.lock.Unlock()
	return b

}

func (r *Hash) HGETALL(table string) *work {

	input, _ := EF.MergeValue("HGETALL", table, nil, nil)
	fmt.Println(input)

	//r.lock.Lock()
	b := r.work.constructor()
	b.lock.Lock()

	b.hashInput.Input = input
	b.hashInput.Action = "HGETALL"
	b.value = b.hashInput
	b.lock.Unlock()
	//	defer r.lock.Unlock()
	return b

}

func (r *Hash) HDEL(table string, key string) *work {

	input, _ := EF.MergeValue("HDEL", table, key, nil)

	b := r.work.constructor()
	b.lock.Lock()

	b.hashInput.Input = input
	b.hashInput.Action = "HDEL"
	b.value = b.hashInput
	b.lock.Unlock()

	return b

}
func (r *Hash) HEXISTS(table string, key string) *work {

	input, _ := EF.MergeValue("HEXISTS", table, key, nil)

	b := r.work.constructor()
	b.lock.Lock()

	b.hashInput.Input = input
	b.hashInput.Action = "HEXISTS"
	b.value = b.hashInput
	b.lock.Unlock()
	return b

}
func (r *Hash) HSET(table string, key string, value ...interface{}) *work {

	input, _ := EF.MergeValue("HSET", table, key, value...)

	b := r.work.constructor()
	b.lock.Lock()

	b.hashInput.Input = input
	b.hashInput.Action = "HSET"
	b.value = b.hashInput
	b.lock.Unlock()
	return b

}

type Set struct {
	work work
	mode string
	lock sync.Mutex
}

func (r *Set) SADD(table string, value ...interface{}) *work {

	input, _ := EF.MergeValue("SADD", table, nil, value...)

	fmt.Println(input)

	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.setInput.Input = input
	b.value = b.setInput
	b.setInput.Action = "SADD"
	b.value = b.setInput
	b.lock.Unlock()
	return b

}

func (r *Set) SCARD(table string, key string) *work {

	input, _ := EF.MergeValue("SCARD", table, key, nil)

	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.setInput.Input = input
	b.setInput.Action = "SCARD"
	b.value = b.setInput

	b.lock.Unlock()
	return b

}

func (r *Set) SDIFF(table string, key string) *work {

	input, _ := EF.MergeValue("SDIFF", table, key, nil)

	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.setInput.Input = input
	b.setInput.Action = "SDIFF"
	b.value = b.setInput
	b.lock.Unlock()
	return b

}

func (r *Set) SINTER(table string, key string) *work {

	input, _ := EF.MergeValue("SINTER", table, key, nil)

	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.setInput.Input = input
	b.setInput.Action = "SINTER"
	b.value = b.setInput
	b.lock.Unlock()

	return b

}

func (r *Set) SISMEMBER(table string, key string) *work {

	input, _ := EF.MergeValue("SISMEMBER", table, key, nil)

	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.setInput.Input = input
	b.value = b.setInput
	b.setInput.Action = "SISMEMBER"
	b.lock.Unlock()
	return b

}

func (r *Set) SMEMBERS(table string) *work {

	input, _ := EF.MergeValue("SMEMBERS", table, nil, nil)

	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.setInput.Input = input
	b.value = b.setInput
	b.setInput.Action = "SMEMBERS"
	b.lock.Unlock()

	return b

}

func (r *Set) SSCAN(table string, key string) *work {

	input, _ := EF.MergeValue("SSCAN", table, key, nil)

	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.setInput.Input = input
	b.value = b.setInput
	b.setInput.Action = "SSCAN"
	b.lock.Unlock()
	return b

}

type List struct {
	work work
	mode string
}

func (r *List) LSET(table string, index string, value ...interface{}) *work {

	input, _ := EF.MergeValue("LSET", table, index, value)

	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.listInput.Input = input
	b.listInput.Action = "LSET"
	b.value = b.listInput
	b.lock.Unlock()
	return b

}

func (r *List) LPUSH(table string, value interface{}) *work {

	input, _ := EF.MergeValue("LPUSH", table, value, nil)

	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.listInput.Input = input
	b.listInput.Action = "LPUSH"
	b.value = b.listInput
	b.lock.Unlock()
	return b

}

func (r *List) LINDEX(table string, index string) *work {

	input, _ := EF.MergeValue("LINDEX", table, index, nil)

	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.listInput.Input = input
	b.listInput.Action = "LINDEX"
	b.value = b.listInput
	b.lock.Unlock()
	return b

}

func (r *List) LLEN(table string) *work {

	input, _ := EF.MergeValue("LLEN", table, nil, nil)

	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.listInput.Input = input
	b.listInput.Action = "LLEN"
	b.value = b.listInput
	b.lock.Unlock()
	return b

}

type Key struct {
	work work
	mode string
}

func (r *Key) DEL(table string) *work {

	input, _ := EF.MergeValue("DEL", table, nil, nil)

	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.keyInput.Input = input
	b.keyInput.Action = "DEL"
	b.value = b.keyInput

	b.lock.Unlock()
	return b

}

func (r *Key) EXPIRE(table string, time int) *work {

	input, _ := EF.MergeValue("EXPIRE", table, nil, time)
	fmt.Println(input)
	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.keyInput.Input = input
	b.keyInput.Action = "EXPIRE"
	b.value = b.keyInput

	b.lock.Unlock()
	return b

}

type Other struct {
	work work
	mode string
}

func (r *Other) SCAN(table string) *work {

	input, _ := EF.MergeValue("SCAN", table, nil, nil)
	b := r.work.constructor()
	b.lock.Lock()
	b.Mode = r.mode
	b.otherInput.Input = input
	b.otherInput.Action = "SCAN"
	b.value = b.otherInput
	b.lock.Unlock()

	return b

}
