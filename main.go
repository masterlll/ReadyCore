package main

import (
	"fmt"
	"strconv"
	"time"

	db "github.com/ReadyCore/goef/core"
)

func main() {
	//	Defeault()
	Cluster()

}

func Defeault() {

	p := db.RedisConnModel{}
	p.Default("127.0.0.1:6379", "")
	p.Auth(false)
	p.RedisConning()
	fmt.Println(p.Ping(), "ping ")
	/////////////////////////////
	// // Do
	data := p.Hash.HSET("aaa", "aa", "aaa").DO(0)
	fmt.Println(data)

	// Pipe
	data2 := p.Hash.HSET("bb", "bb", "bb").Pipe(0)
	fmt.Println(data2)

	//PipeTWice
	in := p.Hash.HSET("cc", "cc", "cc").Value()
	for v := range p.Hash.HSET("vv", "vv", "v").PipeTWice(0, in) {
		fmt.Println(v, "asadsa")
		time.Sleep(time.Microsecond * 100)
	}
	//QueuePipe
	go func() {
		var aa []db.Container
		for i := 1; i <= 1000; i++ {
			aa = append(aa, p.Hash.HSET("f123ss56d", strconv.Itoa(i)+"cdcc", "ssss").Value())
		}
		t1 := time.Now() // get current time
		fmt.Println("go")
		for m := range p.Queue.QueuePipe(0, aa) {
			m = m
		}
		fmt.Println("App elapsed: ", time.Since(t1))
	}()

	var aa []db.Container
	for i := 1; i <= 1000; i++ {
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

	data := p.Hash.HSET("aaa", "aa", "aaa").DO(0)
	fmt.Println(data)
	data2 := p.Hash.HSET("bb", "bb", "bb").Pipe(0)
	fmt.Println(data2)

	in := p.Hash.HSET("cc", "cc", "cc").Value()
	for v := range p.Hash.HSET("vv", "vv", "v").PipeTWice(0, in) {
		fmt.Println(v, "asadsa")
	}

	go func() {
		var aa []db.Container
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
	var aa []db.Container
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
