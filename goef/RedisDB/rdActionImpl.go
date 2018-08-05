package RedisDB

import (
	"fmt"
	
	EF "ReadyCore/ReadyCore/goef/other"
	//"github.com/gomodule/redigo/redis"
)

type Redis struct {
}

// pipe
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

func (S *Redis) pipetwice(DB int, input ...EF.Container) chan interface{} {

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
