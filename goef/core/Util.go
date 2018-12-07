package RedigoEFcore

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"
)

const (
	maxIdle        = 120000
	maxActive      = 240000
	connectTimeout = 1000
	ReadTimeout    = 1000
	writeTimeout   = 1000
	Idleint        = 3
	Wait           = true
	Auth           = true
	dbnumber       = 0
	tcp            = "tcp"
	single         = "single"
	Cluster        = "Cluster"
)

var connmap map[string][]*redis.Pool

func init() {
	connmap = make(map[string][]*redis.Pool)
}
func connkey() (connkey string) { // 生成 connkey
	uuid, _ := uuid.NewV4()
	connkey = uuid.String()
	return
}
func authCheck(c redis.Conn, pwd string) bool { // 密碼 驗證
	if _, err := c.Do("AUTH", pwd); err != nil { ///密碼　驗證
		fmt.Println(err)
		c.Close()
		return false
	}
	return true
}
