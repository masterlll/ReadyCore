package RedigoEFcore

import (
	"fmt"

	EF "github.com/ReadyCore/goef/other"
	//"github.com/gomodule/redigo/redis"
)

type RedisMode struct {
}

func (S *Redis) pipe(DB int, input ...EF.Container) chan interface{} {

	data := make(chan interface{})
	ok := make(chan bool)
	c := RDConn(DB).Get()
	go func(input []EF.Container) {
		for _, p := range input {
			err2 := c.Send(p.Action, p.Input...)
			if err2 != nil {
				fmt.Println("redis  failed:", err2)
			}
		}
		if err := c.Flush(); err != nil { // 清空記憶體　發送
			fmt.Println("err :", err)
		}
		for i := 1; i <= len(input); i++ {
			re, err := c.Receive()
			if err != nil {
				fmt.Println("redis  failed:", err)
			}
			data <- re
		}

		ok <- true
	}(input)
	go func() {
		<-ok
		c.Close()
		close(data)

	}()
	return data
}

func (S *Redis) do(DB int, In EF.Container) chan interface{} {
	fmt.Println("  redis   do ")
	DO := make(chan interface{})
	c := RDConn(DB).Get()

	res, err := c.Do(In.Action, In.Input...)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		DO <- res
		c.Close()
		close(DO)
	}()
	return DO
}

func (S *Redis) pipetwice(DB int, input ...[]EF.Container) chan interface{} {

	data := make(chan interface{})
	ok := make(chan bool)
	c := RDConn(DB).Get()

	for _, in := range input {
		fmt.Println(len(input))
		fmt.Println(input)
		go func(input []EF.Container) {

			err := c.Send(input[0].Action, input[0].Input...)
			if err != nil {
				fmt.Println("redis  failed:", err)
			}

			err2 := c.Send(input[1].Action, input[1].Input...)
			if err2 != nil {
				fmt.Println("redis  failed:", err2)
			}

			if err := c.Flush(); err != nil { // 清空記憶體　發送
				fmt.Println("err :", err)
			}
			for i := 1; i <= len(input); i++ {
				re, err := c.Receive()
				if err != nil {
					fmt.Println("redis  failed:", err)
				}
				data <- re
			}
			ok <- true
		}(in)

	}
	go func() {
		<-ok
		c.Close()
		close(data)

	}()
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
