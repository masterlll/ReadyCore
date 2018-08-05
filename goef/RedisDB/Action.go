package RedisDB

import (
	//	"fmt"
	EF "ReadyCore/ReadyCore/goef/other"
	//"time"
)

type DbContext struct {
	Redis
}

func (red *DbContext) Pipe(DB int, in ...EF.Container) chan interface{} {

	return red.pipe(DB, in...)

}

func (red *DbContext) DO(DB int, in EF.Container) chan interface{} {

	return red.do(DB, in)

}
