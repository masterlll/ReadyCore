package RedisDB

import (
	"fmt"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

type singleton struct {
	poolist []*redis.Pool
	Pool    *redis.Pool
}

type RDConnModel struct {
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

	if err := red.poolarray(13); err != nil {
		return err
	}
	return nil
}

func (red *singleton) poolarray(i int) error {

	for b := 0; b <= i; b++ {
		var cf RDConnModel
		//cf.ProxyAddress = "127.0.0.1:6379"
		cf.PassWord = "12345678"
		//c
		//連線ＴＹＰＥ , 最大閒置連線　，　最大連線
		cf.ConnType = "tcp"
		cf.MaxIdle = 120000
		cf.MaxActive = 2400000
		/// 連線　　讀取　寫入　的逾時
		cf.ConnectTimeout = 10000
		cf.ReadTimeout = 1000
		cf.WriteTimeout = 1000
		////0
		//// 選擇ＤＢ　０,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15
		cf.DBnumber = b
		// fmt.Println("cf", cf)
		conn := newPool(cf) /////設定連線

		err := conn.TestOnBorrow(conn.Get(), time.Now())
		if err != nil {


			return err
		}
		red.poolist = append(red.poolist, &conn)
	}

	return nil
}

func RDConn(i int) *redis.Pool {
	return Shared().poolist[i]
}

func newPool(cf RDConnModel) redis.Pool {

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
