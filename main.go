package main

import (
	"fmt"
	"strconv"
	"time"

	db "github.com/ReadyCore/goef/core"
	ef "github.com/ReadyCore/goef/other"
)

func main() {

	Cluster()
}

func Defeault() {

	p := db.RedisConnModel{}
	p.Default("127.0.0.1:6379", "")
	p.Auth(false)
	p.RedisConning()
	fmt.Println(p.Ping(), "ping ")
	//fmt.Println(p.Hash.HSET("ffd", "bb", "ss").DO(0).Int64ToString())
	var aa []ef.Container
	for i := 1; i <= 1000000; i++ {
		aa = append(aa, p.Hash.HSET("f123156d", strconv.Itoa(i)+"cdcc", "ssss").Value())
	}
	t1 := time.Now() // get current time
	fmt.Println("go")
	for m := range p.Queue.QueuePipe(0, aa) {
		m = m
	}
	fmt.Println("App elapsed: ", time.Since(t1))
}

func Cluster() {
	p := db.ClusterConnModel{}
	p.Default([]string{"192.168.6.85:7000", "192.168.6.85:7001", "192.168.6.85:7002", "192.168.6.85:7003", "192.168.6.85:7004", "192.168.6.85:7005"}, "12345678")
	//	p.Auth(false)
	p.ClusterConning()
	fmt.Println(p.Ping(), "ping ?")

	// data := p.Hash.HSET("aaa", "aa", "aaa").DO(0)
	// fmt.Println(data)
	// data2 := p.Hash.HSET("bb", "bb", "bb").Pipe(0)
	// fmt.Println(data2)
	// in := p.Hash.HSET("cc", "cc", "cc").Value()
	// for v := range p.Hash.HSET("vv", "vv", "v").PipeTWice(0, in) {
	// 	fmt.Println(v, "asadsa")
	// }
	go func() {
		var aa []ef.Container
		for i := 1; i <= 1000; i++ {
			aa = append(aa, p.Hash.HSET("f123106d", strconv.Itoa(i)+"cdcc", "ssss").Value())
		}
		t1 := time.Now()
		fmt.Println("go")
		for m := range p.Queue.QueuePipe(0, aa) {
			m = m
		}
		fmt.Println("App elapsed: ", time.Since(t1))
		time.Sleep(time.Millisecond * 100) //資源要關乾淨
	}()

	var aa []ef.Container
	for i := 1; i <= 1000; i++ {
		aa = append(aa, p.Hash.HSET("f123156d", strconv.Itoa(i)+"cdcc", "ssss").Value())
	}
	t1 := time.Now()
	fmt.Println("go")
	for m := range p.Queue.QueuePipe(0, aa) {
		m = m
	}
	fmt.Println("App elapsed: ", time.Since(t1))
	time.Sleep(time.Millisecond * 100) //資源要關乾淨

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
