package RedigoEFcore

import (
	"fmt"
	"time"

	EF "github.com/ReadyCore/goef/other"
	"github.com/gomodule/redigo/redis"
	//"github.com/gomodule/redigo/redis"
)

type RedisMode struct {
}

func (S *RedisMode) pipe(Conn redis.Conn, input ...EF.Container) chan interface{} {

	data := make(chan interface{})
	//ok := make(chan bool)
	ok := make(chan redis.Conn)
	t1 := time.Now()
	go func(input []EF.Container, C redis.Conn) {

		for _, p := range input {
			err2 := C.Send(p.Action, p.Input...)
			if err2 != nil {
				//	logger.Err(err2, p.Action+" pipe().send")
				data <- err2
				continue
			}
			if err := C.Flush(); err != nil { // 清空記憶體　發送
				//logger.Err(err, " pipe().Flush")
				data <- err
				continue
			}
			re, err := C.Receive()
			if err != nil {
				//	logger.Err(err, " pipe().Receive")
				data <- err
				continue
			}
			data <- re
		}
		ok <- C
		close(ok)
	}(input, Conn)
	go func() {
		c := <-ok
		c.Close()
		close(data)
	}()
	return data
	//	c := RDConn(DB).Get()
	// go func(input []EF.Container, C redis.Conn) {
	// 	for _, p := range input {
	// 		err2 := C.Send(p.Action, p.Input...)
	// 		if err2 != nil {
	// 			fmt.Println("redis  failed:", err2)
	// 		}
	// 	}
	// 	if err := C.Flush(); err != nil { // 清空記憶體　發送
	// 		fmt.Println("err :", err)
	// 	}
	// 	for i := 1; i <= len(input); i++ {
	// 		re, err := C.Receive()
	// 		if err != nil {
	// 			fmt.Println("redis  failed:", err)
	// 		}
	// 		data <- re
	// 	}
	// 	ok <- true
	// }(input, Conn)
	// go func() {
	// 	<-ok
	// 	Conn.Close()
	// 	close(data)

	// }()
	fmt.Println("App elapsed: ", time.Since(t1))
	return data
}

func (S *RedisMode) do(Conn redis.Conn, In EF.Container) chan interface{} {
	fmt.Println("  redis   do ")
	DO := make(chan interface{})

	res, err := Conn.Do(In.Action, In.Input...)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		DO <- res
		Conn.Close()
		close(DO)
	}()
	return DO
}
func (S *RedisMode) pipetwice(DB int, input ...[]EF.Container) chan interface{} {

	data := make(chan interface{})
	//ok := make(chan bool)
	// c := RDConn(DB).Get()

	// for _, in := range input {
	// 	fmt.Println(len(input))
	// 	fmt.Println(input)
	// 	go func(input []EF.Container) {

	// 		err := c.Send(input[0].Action, input[0].Input...)
	// 		if err != nil {
	// 			fmt.Println("redis  failed:", err)
	// 		}

	// 		err2 := c.Send(input[1].Action, input[1].Input...)
	// 		if err2 != nil {
	// 			fmt.Println("redis  failed:", err2)
	// 		}

	// 		if err := c.Flush(); err != nil { // 清空記憶體　發送
	// 			fmt.Println("err :", err)
	// 		}
	// 		for i := 1; i <= len(input); i++ {
	// 			re, err := c.Receive()
	// 			if err != nil {
	// 				fmt.Println("redis  failed:", err)
	// 			}
	// 			data <- re
	// 		}
	// 		ok <- true
	// 	}(in)

	// }
	// go func() {
	// 	<-ok
	// 	c.Close()
	// 	close(data)

	// }()
	return data
}

// type ClusterMode struct {
// }

// func (S *ClusterMode) do(DB int, In EF.Container) chan interface{} {
// 	fmt.Println("  cluster  do ")

// 	DO := make(chan interface{})
// 	ok := make(chan redis.Conn)
// 	time.Sleep(1 * time.Millisecond)
// 	go func() {
// 		c := ClustorConn()
// 		res, err := c.Do(In.Action, In.Input...)
// 		if err != nil {
// 			///
// 		}
// 		DO <- res
// 		ok <- c
// 	}()
// 	go func() {
// 		for c := range ok {
// 			c.Close()
// 		}
// 		close(DO)
// 	}()
// 	return DO
// }
