package RedigoEFcore

import (
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
)

type ClusterConnModel struct {
	ClusterConn ClusterConn
	//RedisHelper
	// Hash  Hash
	// List  List
	// Set   Set
	// Key   Key
	// Other Other
	// Queue Queue
}

// 設定類
func (red *ClusterConnModel) TcpMode() {
	red.ClusterConn.connType = tcp
}
func (red *ClusterConnModel) HostSet(Hosts []string, password string) {

	red.ClusterConn.proxyAddress = Hosts // host
	red.ClusterConn.passWord = password  // 密碼

}

func (red *ClusterConnModel) MaxConnSet(MaxIdle, MaxActive int) {
	red.ClusterConn.maxActive = MaxActive
	red.ClusterConn.maxIdle = MaxIdle
}
func (red *ClusterConnModel) TimeoutSet(Connect, Read, Write, Idleint int) {
	red.ClusterConn.writeTimeout = Write
	red.ClusterConn.connectTimeout = Connect
	red.ClusterConn.readTimeout = Read
	red.ClusterConn.idleTimeout = Idleint

}
func (red *ClusterConnModel) Wait(wait bool) {
	red.ClusterConn.wait = wait
}
func (red *ClusterConnModel) Auth(start bool) {
	red.ClusterConn.auth = start
}
func (red *ClusterConnModel) Default(Hosts []string, password string) {
	red.TcpMode()
	red.MaxConnSet(maxIdle, maxActive)
	red.Auth(Auth)
	red.TimeoutSet(connectTimeout, ReadTimeout, writeTimeout, Idleint)
	red.Wait(Wait)
	red.HostSet(Hosts, password)
}

func (red *ClusterConnModel) Ping() bool {
	c := clusterConnMap[red.ClusterConn.connKey].Get()
	_, err := c.Do("PING")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 連線類

 //// 配新實體　
func (red *ClusterConnModel) NewEnity() (*CEnity, error) {
	if red.ClusterConn.connKey == "" {
		return nil, errors.New("redis not Conning or miss ConnKey ! ")
	}
	if !red.Ping() {
		return nil, errors.New("redis  ping  err ! ")
	}

	Enity := CEnity{}
	Enity.Hash.mode = Cluster // 設定 help mode ..
	Enity.List.mode = Cluster
	Enity.Set.mode = Cluster
	Enity.Key.mode = Cluster
	Enity.Other.mode = Cluster
	Enity.Queue.mode = Cluster

	
	//// 配　連線　ｉｄ

	Enity.Hash.connkey = red.ClusterConn.connKey
	Enity.List.connkey = red.ClusterConn.connKey
	Enity.Set.connkey = red.ClusterConn.connKey
	Enity.Key.connkey = red.ClusterConn.connKey
	Enity.Other.connkey = red.ClusterConn.connKey
	Enity.Queue.connkey = red.ClusterConn.connKey

	return &Enity, nil

}

func (red *ClusterConnModel) ClusterConning() (*CEnity, error) {
	conn, err := red.clusterPool()
	if err != nil {
		return nil, err
	}
	

	uuid := connkey() // 配連線Key   connkey
	clusterConnMap[uuid] = conn

	Enity := CEnity{}
	Enity.Hash.mode = Cluster // 設定 help mode ..
	Enity.List.mode = Cluster
	Enity.Set.mode = Cluster
	Enity.Key.mode = Cluster
	Enity.Other.mode = Cluster
	Enity.Queue.mode = Cluster
	//// 配　連線　ｉｄ
	
	red.ClusterConn.connKey = uuid
	Enity.Hash.connkey = uuid
	Enity.List.connkey = uuid
	Enity.Set.connkey = uuid
	Enity.Key.connkey = uuid
	Enity.Other.connkey = uuid
	Enity.Queue.connkey = uuid
	
	return &Enity, nil

}

// 建立 集群連線
func (red *ClusterConnModel) clusterPool() (*redisc.Cluster, error) {
	cluster := redisc.Cluster{
		StartupNodes: red.ClusterConn.proxyAddress,
		DialOptions: []redis.DialOption{redis.DialConnectTimeout(time.Duration(red.ClusterConn.connectTimeout) * time.Second),
			redis.DialReadTimeout(time.Duration(red.ClusterConn.readTimeout) * time.Second),
			redis.DialWriteTimeout(time.Duration(red.ClusterConn.writeTimeout) * time.Second)},
		CreatePool: red.clustorPoolBuild, //建立redis 連線
	}
	if err := cluster.Refresh(); err != nil {
		///err
	}
	return &cluster, nil
}

//建立redis 連線
func (red *ClusterConnModel) clustorPoolBuild(addr string, opts ...redis.DialOption) (*redis.Pool, error) {
	return &redis.Pool{
		MaxIdle:     red.ClusterConn.maxIdle,
		MaxActive:   red.ClusterConn.maxActive,
		IdleTimeout: time.Duration(red.ClusterConn.idleTimeout) * time.Second, // 閒置連線逾時
		Wait:        red.ClusterConn.wait,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(red.ClusterConn.connType, addr, opts...)
			if err != nil {
				return nil, err
			}
			if red.ClusterConn.auth {
				if !authCheck(c, red.ClusterConn.passWord) {
					return c, errors.New("Auth err ! password ok ?  or Auth still trun off ! ")
				} // 驗證
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, nil
}
