package RedigoEFcore

import (
	EF "github.com/ReadyCore/goef/other"
)

type dbContext struct {
	RedisMode
	ClusterMode
}

func (red *dbContext) Pipe(connkey string, Mode string, DBnumber int, in ...EF.Container) chan interface{} {

	switch Mode {
	case single:
		{
			return red.pipe(connGet(connkey, DBnumber), in...)
		}
	case Cluster:
		{
			return red.pipeCluster(clusterconnGet(connkey), in...)
		}
	}
	return nil
}

func (red *dbContext) DO(key string, Mode string, DBnumber int, in EF.Container) chan interface{} {

	switch Mode {
	case single:
		{
			return red.do(connGet(key, DBnumber), in)
		}
	case Cluster:
		{
			return red.doCluster(clusterconnGet(key), in)
		}
	}

	return nil

}

func (red *dbContext) PipeTwice(key string, Mode string, DBnumber int, input []EF.Container) chan interface{} {

	switch Mode {
	case single:
		{
			return red.pipetwice(connGet(key, DBnumber), input)

		}
	case Cluster:
		{
			return red.pipetwiceCluster(clusterconnGet(key), input)
		}
	}
	return nil

}
