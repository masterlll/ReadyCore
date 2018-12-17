package RedigoEFcore

//"fmt"

type dbContext struct {
	RedisMode
	ClusterMode
}

func (red *dbContext) Pipe(connkey string, Mode string, DBnumber int, in ...Container) chan interface{} {

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

func (red *dbContext) DO(key string, Mode string, DBnumber int, in Container) chan interface{} {

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

func (red *dbContext) PipeTwice(key string, Mode string, DBnumber int, input []Container) chan interface{} {

	switch Mode {
	case single:
		{
			return red.pipetwice(connGet(key, DBnumber), input)

		}
	case Cluster:
		{
			ch := make(chan interface{})
			go func() {
				for _, in := range input {
					ch <- red.doCluster(clusterconnGet(key), in)
				}
				close(ch)
			}()
			return ch
		}
	}
	return nil

}
