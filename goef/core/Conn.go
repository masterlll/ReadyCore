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

func (red *RedisConnModel) TcpMode() {
	red.RedisConn.ConnType = tcp
}
func (red *RedisConnModel) HostSet(Host, password string) {

	red.RedisConn.ProxyAddress = Host // host
	red.RedisConn.PassWord = password // 密碼

	red.RedisHelper.Hash.mode = single // 設定 help mode ..
	red.RedisHelper.List.mode = single
	red.RedisHelper.Set.mode = single
	red.RedisHelper.Key.mode = single
	red.RedisHelper.Other.mode = single
	red.RedisHelper.Queue.mode = single

}

func (red *RedisConnModel) MaxConnSet(MaxIdle, MaxActive int) {
	red.RedisConn.MaxActive = MaxActive
	red.RedisConn.MaxIdle = MaxIdle
}
func (red *RedisConnModel) TimeoutSet(Connect, Read, Write, Idleint int) {
	red.RedisConn.WriteTimeout = Write
	red.RedisConn.ConnectTimeout = Connect
	red.RedisConn.ReadTimeout = Read
	red.RedisConn.IdleTimeout = Idleint

}
func (red *RedisConnModel) Wait(wait bool) {
	red.RedisConn.Wait = wait
}
func (red *RedisConnModel) DBnumberSet(Total int) {
	red.RedisConn.DBnumber = Total
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
	c := connmap[red.RedisConn.connKey][red.RedisConn.DBnumber].Get()
	_, err := c.Do("PING")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 開始連線
func (red *RedisConnModel) RedisConning() error {
	conn, err := red.connPool(red.RedisConn.DBnumber)
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

func (red *RedisConnModel) connPool(i int) (Connpool []*redis.Pool, err error) {
	b := i
	for b <= i {
		conn := newPool(red.RedisConn, b) /////設定連線
		err = conn.TestOnBorrow(conn.Get(), time.Now())
		if err != nil {
			fmt.Println(err)
			return
		}
		Connpool = append(Connpool, &conn)
		b++
	}
	return
}

func connGet(key string, dbnumber int) redis.Conn {
	return connmap[key][dbnumber].Get()
}

func newPool(cf RedisConn, number int) redis.Pool {

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
			if cf.auth {
				if !authCheck(c, cf.PassWord) {
					return c, errors.New("Auth err password ?")
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
		IdleTimeout: 3 * time.Second, // 閒置連線逾時
		Wait:        true,
	}

	return a
}
