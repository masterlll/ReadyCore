package RedisDB

import (
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
)


type Cluster struct {
	
	ProxyAddressCluster []string
	PassWord            string
	MaxIdle             int
	MaxActive           int
	ConnType            string
	ConnectTimeout      int
	ReadTimeout         int
	WriteTimeout        int
	DBnumber            int
}

type RedisConn struct {
	ProxyAddress   string
	PassWord       string
	MaxIdle        int
	MaxActive      int
	ConnType       string
	ConnectTimeout int
	ReadTimeout    int
	WriteTimeout   int
	DBnumber       int
}

type singleton struct {
	poolist     []*redis.Pool
	Pool        *redis.Pool
	clusterpool redisc.Cluster
	Cluster     Cluster
	RedisConn   RedisConn
}

var instance *singleton
var once sync.Once

// var cLocks [n]sync.RWMutex

func Shared() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}

func (red *singleton) InitRedis() error {

	if err := red.connPool(13); err != nil {
		return err
	}
	return nil
}

func (red *singleton) ClusterPool() error {

	cluster := redisc.Cluster{
		StartupNodes: red.Cluster.ProxyAddressCluster,
		DialOptions: []redis.DialOption{redis.DialConnectTimeout(time.Duration(red.Cluster.ConnectTimeout) * time.Second),
			redis.DialReadTimeout(time.Duration(red.Cluster.ReadTimeout) * time.Millisecond),
			redis.DialWriteTimeout(time.Duration(red.Cluster.WriteTimeout) * time.Millisecond)},
		CreatePool: red.createPool,
	}
	if err := cluster.Refresh(); err != nil {
		///err
	}
	red.clusterpool = cluster
	return nil
}

func (red *singleton) connPool(i int) error {

	var b int
	for b <= i {
		// //cf.ProxyAddress = "127.0.0.1:6379"
		// cf.PassWord = "12345678"
		// //c
		// //連線ＴＹＰＥ , 最大閒置連線　，　最大連線
		// cf.ConnType = "tcp"
		// cf.MaxIdle = 120000
		// cf.MaxActive = 2400000
		// /// 連線　　讀取　寫入　的逾時
		// cf.ConnectTimeout = 10000
		// cf.ReadTimeout = 1000
		// cf.WriteTimeout = 1000
		// ////0
		// //// 選擇ＤＢ　０,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15
		// cf.DBnumber = b
		var p = red.RedisConn
		p.DBnumber = b
		// fmt.Println("cf", cf)
		conn := newPool(p) /////設定連線
		err := conn.TestOnBorrow(conn.Get(), time.Now())
		if err != nil {
			return err
		}
		red.poolist = append(red.poolist, &conn)
		b++
	}

	return nil
}

func RDConn(i int) *redis.Pool {
	return Shared().poolist[i]
}

func ClustorConn() redis.Conn {
	a := Shared().clusterpool.Get()
	return a
}

func (red *singleton) createPool(addr string, opts ...redis.DialOption) (*redis.Pool, error) {
	return &redis.Pool{
		MaxIdle:     red.Cluster.MaxIdle,
		MaxActive:   red.Cluster.MaxActive,
		IdleTimeout: time.Minute,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial( red.Cluster.ConnType, addr, opts...)
			if err != nil {
				fmt.Println(addr, red.Cluster.PassWord)
				fmt.Println(err)
				return nil, err
			}
			if _, err := c.Do("AUTH", red.Cluster.PassWord); err != nil { ///密碼　驗證
				c.Close()
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, nil
}

func newPool(cf RedisConn) redis.Pool {

	a := redis.Pool{
		MaxIdle:   cf.MaxIdle,
		MaxActive: cf.MaxActive, // max number of connections

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(cf.ConnType, cf.ProxyAddress,
				redis.DialConnectTimeout(time.Duration(cf.ConnectTimeout)*time.Millisecond),
				redis.DialReadTimeout(time.Duration(cf.ReadTimeout)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(cf.WriteTimeout)*time.Millisecond),
			)
			if err != nil {
				panic(err.Error())
			}
			if _, err := c.Do("AUTH", cf.PassWord); err != nil { ///密碼　驗證
				c.Close()
			}
			if _, err := c.Do("SELECT", cf.DBnumber); err != nil { /// 選資料庫
				c.Close()
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { //檢查連線狀態
			if time.Since(t) > time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},

		IdleTimeout: 3 * time.Second, // 閒置連線逾時
		Wait:        true,
	}

	return a
}
