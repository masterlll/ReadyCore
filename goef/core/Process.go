package RedigoEFcore

import (
	"fmt"

	EF "github.com/ReadyCore/goef/other"
	"github.com/gomodule/redigo/redis"
	//"github.com/gomodule/redigo/redis"
)

type RedisMode struct {
}

func (S *RedisMode) pipe(Conn redis.Conn, input ...EF.Container) chan interface{} {

	data := make(chan interface{})
	ok := make(chan bool)
	go func(input []EF.Container, C redis.Conn) {
		for _, p := range input {
			err2 := C.Send(p.Action, p.Input...)
			if err2 != nil {
				fmt.Println("redis  failed:", err2)
				data <- err2
			}
		}
		if err := C.Flush(); err != nil { // 清空記憶體　發送
			fmt.Println("err :", err)
			data <- err
		}
		for i := 1; i <= len(input); i++ {
			re, err := C.Receive()
			if err != nil {
				fmt.Println("redis  failed:", err)
				data <- err
			}
			data <- re
		}
		ok <- true
	}(input, Conn)
	go func() {
		<-ok
		Conn.Close()
		close(data)
		close(ok)
	}()
	return data
}

func (S *RedisMode) do(Conn redis.Conn, In EF.Container) chan interface{} {

	DO := make(chan interface{})
	res, err := Conn.Do(In.Action, In.Input...)
	if err != nil {
		fmt.Println(err)
		DO <- err
	}
	go func() {
		DO <- res
		Conn.Close()
		close(DO)
	}()
	return DO
}
func (S *RedisMode) pipetwice(Conn redis.Conn, input []EF.Container) chan interface{} {

	data := make(chan interface{})
	ok := make(chan bool)

	go func(input []EF.Container) {
		err := Conn.Send(input[0].Action, input[0].Input...)
		if err != nil {
			fmt.Println("redis  failed:", err)
		}
		err2 := Conn.Send(input[1].Action, input[1].Input...)
		if err2 != nil {
			fmt.Println("redis  failed:", err2)
		}
		if err := Conn.Flush(); err != nil { // 清空記憶體　發送
			fmt.Println("err :", err)
		}
		for i := 1; i <= len(input); i++ {
			re, err := Conn.Receive()
			if err != nil {
				fmt.Println("redis  failed:", err)
			}
			data <- re
		}
		ok <- true
	}(input)

	go func() {
		<-ok
		Conn.Close()
		close(data)

	}()
	return data
}

type ClusterMode struct {
}

func (S *ClusterMode) pipeCluster(Conn redis.Conn, input ...EF.Container) chan interface{} {
	data := make(chan interface{})
	ok := make(chan redis.Conn)
	go func(c redis.Conn, input []EF.Container) {

		for _, p := range input {
			err2 := c.Send(p.Action, p.Input...)
			if err2 != nil {
				fmt.Println(err2, p.Action+" pipe().send")
				data <- err2
				continue
			}
			if err := c.Flush(); err != nil { // 清空記憶體　發送
				fmt.Println(err, " pipe().Flush")
				data <- err
				continue
			}
			re, err := c.Receive()
			if err != nil {
				fmt.Println(err, " pipe().Receive")
				data <- err
				continue
			}
			data <- re
		}
		ok <- c
		close(ok)
	}(Conn, input)
	go func() {
		c := <-ok
		c.Close()
		close(data)
	}()
	return data
}

func (S *ClusterMode) doCluster(Conn redis.Conn, In EF.Container) chan interface{} {
	DO := make(chan interface{})
	var value interface{}
	res, err := Conn.Do(In.Action, In.Input...)
	if err != nil {
		fmt.Println(err, In.Action+" do().do")
		value = err
	} else {
		value = res
	}
	go func() {
		DO <- value
		Conn.Close()
		close(DO)
	}()
	return DO
}
func (S *ClusterMode) pipetwiceCluster(Conn redis.Conn, input []EF.Container) chan interface{} {

	data := make(chan interface{})
	ok := make(chan redis.Conn)

	go func(c redis.Conn, input []EF.Container) {
		for _, v := range input {
			err := c.Send(v.Action, v.Input...)
			if err != nil {
				fmt.Println(err, v.Action+" pipetwice().send")
				data <- err
			}
		}
		if err := c.Flush(); err != nil { // 清空記憶體　發送
			fmt.Println(err, " pipetwice().Flush")
			data <- err
		}
		for i := 1; i <= len(input); i++ {
			re, err := c.Receive()
			if err != nil {
				fmt.Println(err, "pipetwice().Receive")
				data <- err
			}
			data <- re
		}
		ok <- Conn
		close(ok)
	}(Conn, input)

	go func() {
		for c := range ok {
			c.Close()
		}
		close(data)

	}()
	return data
}

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
