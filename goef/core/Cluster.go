package RedigoEFcore

type ClusterConn struct {
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

type ClusterConnModel struct {

	/// Cluster
	//clusterpool redisc.Cluster
	//Cluster     Cluster
	RedisHelper
}

// func (red *singleton) ClusterPool() error {

// 	cluster := redisc.Cluster{
// 		StartupNodes: red.Cluster.ProxyAddressCluster,
// 		DialOptions: []redis.DialOption{redis.DialConnectTimeout(time.Duration(red.Cluster.ConnectTimeout) * time.Second),
// 			redis.DialReadTimeout(time.Duration(red.Cluster.ReadTimeout) * time.Millisecond),
// 			redis.DialWriteTimeout(time.Duration(red.Cluster.WriteTimeout) * time.Millisecond)},
// 		CreatePool: red.createPool,
// 	}
// 	if err := cluster.Refresh(); err != nil {
// 		///err
// 	}
// 	red.clusterpool = cluster
// 	return nil
// }

// func ClustorConn() redis.Conn {
// 	a := Shared().clusterpool.Get()
// 	return a
// }

// func (red *singleton) createPool(addr string, opts ...redis.DialOption) (*redis.Pool, error) {
// 	return &redis.Pool{
// 		MaxIdle:     red.Cluster.MaxIdle,
// 		MaxActive:   red.Cluster.MaxActive,
// 		IdleTimeout: time.Minute,
// 		Dial: func() (redis.Conn, error) {
// 			c, err := redis.Dial(red.Cluster.ConnType, addr, opts...)
// 			if err != nil {
// 				fmt.Println(addr, red.Cluster.PassWord)
// 				fmt.Println(err)
// 				return nil, err
// 			}
// 			if _, err := c.Do("AUTH", red.Cluster.PassWord); err != nil { ///密碼　驗證
// 				c.Close()
// 			}
// 			return c, err
// 		},
// 		TestOnBorrow: func(c redis.Conn, t time.Time) error {
// 			_, err := c.Do("PING")
// 			return err
// 		},
// 	}, nil
// }
