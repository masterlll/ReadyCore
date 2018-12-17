package RedigoEFcore

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type RedisMode struct {
}

func (S *RedisMode) pipe(Conn redis.Conn, input ...Container) chan interface{} {

	data := make(chan interface{})
	ok := make(chan bool)
	go func(input []Container, C redis.Conn) {
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
		fmt.Println("pipe close ")

	}()
	return data
}

func (S *RedisMode) do(Conn redis.Conn, In Container) chan interface{} {

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
		fmt.Println("DO close ")
	}()
	return DO
}
func (S *RedisMode) pipetwice(Conn redis.Conn, input []Container) chan interface{} {

	data := make(chan interface{})
	ok := make(chan bool)

	go func(input []Container) {
		for _, i := range input {
			err := Conn.Send(i.Action, i.Input...)
			if err != nil {
				fmt.Println("redis  failed:", err)
				data <- err
			}
		}
		if err := Conn.Flush(); err != nil { // 清空記憶體　發送
			fmt.Println("err :", err)
		}
		re, err := Conn.Receive()
		if err != nil {
			fmt.Println("redis  failed:", err)
			data <- err
		}
		data <- re
		fmt.Println("aa")
		ok <- true
	}(input)

	go func() {
		<-ok
		Conn.Close()
		close(data)
		close(ok)
		fmt.Println("closer")
	}()
	return data
}

type ClusterMode struct {
}

func (S *ClusterMode) pipeCluster(Conn redis.Conn, input ...Container) chan interface{} {
	data := make(chan interface{})
	//ok := make(chan redis.Conn)
	ok := make(chan bool)
	go func() {
		<-ok
		Conn.Close()

		//fmt.Print("conn and chan close")
		close(ok)
		//	fmt.Print("conn and chan close 1")
		close(data)
		//	close(data)
		fmt.Print("conn and chan close")
	}()
	go func(c redis.Conn, input []Container) {

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

		ok <- true
		fmt.Print("ok close")
		//time.Sleep(time.Millisecond * 1000)
	}(Conn, input)

	return data
}

func (S *ClusterMode) doCluster(Conn redis.Conn, In Container) chan interface{} {
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
		fmt.Print("conn and chan close")
	}()
	return DO
}

// 	go func() {
// 		for c := range ok {
// 			c.Close()
// 			fmt.Print("conn close")
// 		}
// 		close(data)
// 		fmt.Print("chan close")
// 	}()
// 	return data
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
