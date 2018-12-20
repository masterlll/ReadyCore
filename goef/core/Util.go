package RedigoEFcore

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
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

var connmap map[string][]*redis.Pool          /// 一般連線  連線池
var clusterConnMap map[string]*redisc.Cluster /// 集群連線   連線池

type RedisConn struct { // 一般設定

	proxyAddress   string // host
	passWord       string // 密碼
	connKey        string // 連結KEY
	maxIdle        int    // 最大閒置連線
	maxActive      int    // 最大連線
	connType       string //連線 模式 TCP ..
	idleTimeout    int    // 閒置連線逾時
	connectTimeout int    //連線  Timeout
	readTimeout    int    //讀取 Timeout
	writeTimeout   int    //寫入 Timeout
	dBnumber       int    // DB　總數
	mode           string // 集群 or 一般
	wait           bool   //等待
	auth           bool   //密碼驗證
}

type ClusterConn struct { ///  集群 設定
	proxyAddress   []string // host
	connKey        string   // 密碼
	passWord       string   // 連結KEY
	maxIdle        int      // 最大閒置連線
	maxActive      int      // 最大連線
	connType       string   //連線 模式 TCP ..
	idleTimeout    int      // 閒置連線逾時
	connectTimeout int      // 連線  Timeout
	readTimeout    int      //讀取 Timeout
	writeTimeout   int      //寫入 Timeout
	dBnumber       int      // DB　總數
	mode           string   // 集群 or 一般
	wait           bool     //等待
	auth           bool     //密碼驗證
}



//// 使用　新實體　執行操作　 v0.2.0
type Enity struct { //預設實體
	Hash  Hash
	List  List
	Set   Set
	Key   Key
	Other Other
	Queue Queue
}
//// 使用　新集群實體　執行操作　 v0.2.0  
type CEnity struct { //預設 集群　實體
	Hash  Hash
	List  List
	Set   Set
	Key   Key
	Other Other
	Queue Queue
}

type Container struct { // 容器
	Action string
	DB     int
	Input  []interface{}
	Twice  []interface{}
}

// 初始化
func init() {
	connmap = make(map[string][]*redis.Pool)
	clusterConnMap = make(map[string]*redisc.Cluster)
}

func connGet(key string, dbnumber int) redis.Conn { //一般連線 拿取連線
	return connmap[key][dbnumber].Get()
}
func clusterconnGet(key string) redis.Conn { //集群連線 拿取連線
	return clusterConnMap[key].Get()
}

func connkey() (connkey string) { // 生成 connkey
	uuid, _ := uuid.NewV4()
	connkey = uuid.String()
	return
}

// 密碼 驗證
func authCheck(c redis.Conn, pwd string) bool {
	if _, err := c.Do("AUTH", pwd); err != nil {
		fmt.Println(err)
		c.Close()
		return false
	}
	return true
}
