package main

import (
	"fmt"
	"strconv"
	"time"

	db "github.com/ReadyCore/goef/RedigoEFcore"
)

func main() {

	p := db.RedisConnModel{}
	p.Default("127.0.0.1:6379", "")
	p.RedisConning()

	//fmt.Println(p.Hash.HSET("ffd", "bb", "ss").DO(0).Int64ToString())

	//var aa []int

	t1 := time.Now() // get current time

	for i := 1; i <= 10000; i++ {
		p.Hash.HSET("ffd", strconv.Itoa(i)+"ccc", "ssss").Pipe(0).Value()
	}

	fmt.Println("App elapsed: ", time.Since(t1))
	// if err := rd.Shared().InitRedis(); err != nil {
	// 	fmt.Println("err", err)
	// }
	// db := Core.Redis{}
	// fmt.Println(db.Hash.HSET("AAAA", "AAAA", "123456789").Pipe(13).Value())
	// fmt.Println(db.Hash.HGET("AAAA", "AAAA").Pipe(13).Value())
	// fmt.Println(db.Hash.HEXISTS("AAAA", "AAAA").Pipe(13).Value())
	// fmt.Println(db.Hash.HSET("s", "s", "s").DO(13).Value())
	// fmt.Println(db.Hash.HGET("s", "s").DO(13).Value())
	// fmt.Println(db.Hash.HEXISTS("s", "s").DO(13).Value())
}

// func test() {

// 	rds := rd.DbContext{}
// 	time.Sleep(1 * time.Millisecond)
// 	go func() {
// 		C := []EF.Container{}
// 		for i := 2000001; i <= 3000000; i++ {

// 			a := EF.Container{}
// 			a.Action = "HSET"
// 			a.DB = 6
// 			a.Input = append(a.Input, i)
// 			a.Input = append(a.Input, "HSET")
// 			a.Input = append(a.Input, "HSET")
// 			C = append(C, a)

// 		}
// 		t1 := time.Now()
// 		for i := range rds.Pipe(10, C...) {
// 			i = i

// 		}
// 		fmt.Println("time1", time.Since(t1)) //印出時間

// 	}()
// 	time.Sleep(1 * time.Millisecond)
// 	time.Sleep(1 * time.Millisecond)
// 	go func() {
// 		C := []EF.Container{}
// 		for i := 1000001; i <= 2000000; i++ {

// 			a := EF.Container{}
// 			a.Action = "HSET"
// 			a.DB = 5
// 			a.Input = append(a.Input, i)
// 			a.Input = append(a.Input, "HSET")
// 			a.Input = append(a.Input, "HSET")
// 			C = append(C, a)

// 		}
// 		t1 := time.Now()
// 		for i := range rds.Pipe(10, C...) {
// 			i = i

// 		}
// 		fmt.Println("time1", time.Since(t1)) //印出時間

// 	}()
// 	time.Sleep(1 * time.Millisecond)

// 	C := []EF.Container{}
// 	for i := 1; i <= 10; i++ {

// 		a := EF.Container{}
// 		a.Action = "DEL"
// 		a.DB = 4
// 		a.Input = append(a.Input, i)
// 		a.Input = append(a.Input, "HSET")
// 		a.Input = append(a.Input, "HSET")
// 		C = append(C, a)

// 		bbb := rds.DO(1, a, 1)
// 		fmt.Println("do", <-bbb)

// 	}
// 	t1 := time.Now()
// 	for i := range rds.Pipe(1, C...) {
// 		i = i
// 		fmt.Println(i)

// 	}
// 	fmt.Println("time1", time.Since(t1)) //印出時間

// }
