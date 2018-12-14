package RedigoEFcore

import (
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisConnModel struct {
	RedisConn RedisConn
	RedisHelper
}

// 設定類
func (red *RedisConnModel) TcpMode() {
	red.RedisConn.connType = tcp
}
func (red *RedisConnModel) HostSet(Host, password string) {

	red.RedisConn.proxyAddress = Host // host
	red.RedisConn.passWord = password // 密碼

	red.RedisHelper.Hash.mode = single // 設定 help mode ..
	red.RedisHelper.List.mode = single
	red.RedisHelper.Set.mode = single
	red.RedisHelper.Key.mode = single
	red.RedisHelper.Other.mode = single
	red.RedisHelper.Queue.mode = single
}

func (red *RedisConnModel) MaxConnSet(MaxIdle, MaxActive int) {
	red.RedisConn.maxActive = MaxActive
	red.RedisConn.maxIdle = MaxIdle
}
func (red *RedisConnModel) TimeoutSet(Connect, Read, Write, Idleint int) {
	red.RedisConn.writeTimeout = Write
	red.RedisConn.connectTimeout = Connect
	red.RedisConn.readTimeout = Read
	red.RedisConn.idleTimeout = Idleint

}
func (red *RedisConnModel) Wait(wait bool) {
	red.RedisConn.wait = wait
}
func (red *RedisConnModel) DBnumberSet(Total int) {
	red.RedisConn.dBnumber = Total
}
func (red *RedisConnModel) Auth(start bool) {
	red.RedisConn.auth = start
}

func (red *RedisConnModel) Default(Host, password string) {
	red.TcpMode()
	red.MaxConnSet(maxIdle, maxActive)
	red.DBnumberSet(dbnumber)
	red.Auth(Auth)
	red.TimeoutSet(connectTimeout, ReadTimeout, writeTimeout, Idleint)
	red.Wait(Wait)
	red.HostSet(Host, password)
}

func (red *RedisConnModel) Ping() bool {
	c := connmap[red.RedisConn.connKey][red.RedisConn.dBnumber].Get()
	_, err := c.Do("PING")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

/// 連線類 方法///

// 開始連線
func (red *RedisConnModel) RedisConning() error {
	conn, err := red.connPool(red.RedisConn.dBnumber)
	if err != nil {
		return err
	}

	uuid := connkey() // 配連線Key   connkey
	connmap[uuid] = conn
	red.RedisConn.connKey = uuid
	red.RedisHelper.Hash.connkey = uuid
	red.RedisHelper.List.connkey = uuid
	red.RedisHelper.Set.connkey = uuid
	red.RedisHelper.Key.connkey = uuid
	red.RedisHelper.Other.connkey = uuid
	red.RedisHelper.Queue.connkey = uuid
	return nil
}

/// 建立連線
func (red *RedisConnModel) connPool(i int) (Connpool []*redis.Pool, err error) {
	b := i
	for b <= i {
		conn := newPool(red.RedisConn, b) /////設定連線

		err = conn.TestOnBorrow(conn.Get(), time.Now())
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		Connpool = append(Connpool, &conn)
		b++
	}
	return
}

func newPool(cf RedisConn, number int) redis.Pool {

	a := redis.Pool{
		MaxIdle:     cf.maxIdle,
		MaxActive:   cf.maxActive,                                // max number of connections
		IdleTimeout: time.Duration(cf.idleTimeout) * time.Second, // 閒置連線逾時
		Wait:        cf.wait,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(cf.connType, cf.proxyAddress,
				redis.DialConnectTimeout(time.Duration(cf.connectTimeout)*time.Millisecond),
				redis.DialReadTimeout(time.Duration(cf.readTimeout)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(cf.writeTimeout)*time.Millisecond),
			)
			if err != nil {
				panic(err.Error())
			}
			if cf.auth {
				if !authCheck(c, cf.passWord) {
					return nil, errors.New("Auth err ! password ok ?  or Auth still trun off ! ")
				} // 驗證
			}
			if _, err := c.Do("SELECT", number); err != nil { /// 選資料庫
				fmt.Println(err)
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
	}

	return a
}
