package RedigoEFcore

import (
	EF "github.com/ReadyCore/goef/other"
	"github.com/gomodule/redigo/redis"
)

//	"fmt"

//"time"

type DbContext struct {
	RedisMode
	//	ClusterMode
}

func (red *DbContext) Pipe(Conn redis.Conn, Mode string, in ...EF.Container) chan interface{} {

	return red.pipe(Conn, in...)

}

func (red *DbContext) DO(Conn redis.Conn, Mode string, in EF.Container) chan interface{} {

	//	if Mode != EF.ModeCluster {
	return red.do(Conn, in)
	//	}
	//	return red.ClusterMode.do(DB, in)

}

func (red *DbContext) PipeTwice(DB int, input ...[]EF.Container) chan interface{} {

	return red.pipetwice(DB, input...)

}
