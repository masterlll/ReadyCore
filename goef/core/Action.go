package RedigoEFcore

import (
	EF "github.com/ReadyCore/goef/other"
)

//	"fmt"

//"time"

type DbContext struct {
	RedisMode
	//	ClusterMode
}

func (red *DbContext) Pipe(connkey string, Mode string, DBnumber int, in ...EF.Container) chan interface{} {

	switch Mode {
	case single:
		{
			//for DBnumber == 0 {

			return red.pipe(connGet(connkey, DBnumber), in...)

			//wo.ConnPool = append(wo.ConnPool, connGet(r.connkey, i))
			//fmt.Println("csc", wo.ConnPool[0].Get())
			//DBnumber--
			//}

		}
	case Cluster:
		{
		}
	}
	return nil
}

func (red *DbContext) DO(key string, Mode string, DBnumber int, in EF.Container) chan interface{} {

	switch Mode {
	case single:
		{
			//for DBnumber == 0 {
			return red.do(connGet("test123", DBnumber), in)

			//wo.ConnPool = append(wo.ConnPool, connGet(r.connkey, i))
			//fmt.Println("csc", wo.ConnPool[0].Get())
			//DBnumber--
			//}

		}
	case Cluster:
		{
		}
	}

	//	if Mode != EF.ModeCluster {
	return nil
	//	}
	//	return red.ClusterMode.do(DB, in)

}

func (red *DbContext) PipeTwice(DB int, input ...[]EF.Container) chan interface{} {

	return red.pipetwice(DB, input...)

}
