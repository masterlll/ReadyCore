package RedisDB

import (
	//	"fmt"
	EF "github.com/ReadyCore/goef/other"
	//"time"
)

type DbContext struct {
	Redis
	ClusterMode
}

func (red *DbContext) Pipe(DB int, in ...EF.Container) chan interface{} {

	return red.pipe(DB, in...)

}

func (red *DbContext) DO(DB int, in EF.Container, Mode int) chan interface{} {

	if Mode != EF.ModeCluster {
		return red.Redis.do(DB, in)
	}
	return red.ClusterMode.do(DB, in)

}

func (red *DbContext) PipeTwice(DB int, input ...[]EF.Container) chan interface{} {

	return red.pipetwice(DB, input...)

}
