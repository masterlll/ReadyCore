package RedigoEFcore

import (
	"errors"
	"strconv"
	"sync"

	"github.com/gomodule/redigo/redis"

	EF "github.com/ReadyCore/goef/other"
)

type work struct {
	ConnPool   []*redis.Pool
	convent    convent
	lock       sync.Mutex
	Mode       string
	value      EF.Container
	hashInput  EF.Container
	setInput   EF.Container
	listInput  EF.Container
	keyInput   EF.Container
	otherInput EF.Container
}

func (p *work) constructor() *work {
	return &work{}
}

func (p *work) connHelper(DBnumber int) (Conn redis.Conn) {

	switch p.Mode {

	case single:
		{
			Conn = p.ConnPool[DBnumber].Get()

		}

	case Cluster:
		{

		}

	}
	return
}
func (p *work) doHelper(conn redis.Conn) *convent {
	rd := DbContext{}
	a := p.convent.constructor()
	if p.setInput.Input != nil {
		a.value = <-rd.DO(conn, p.Mode, p.setInput)
		return &a
	}

	if p.hashInput.Input != nil {
		a.value = <-rd.DO(conn, p.Mode, p.hashInput)
		return &a
	}
	if p.listInput.Input != nil {
		a.value = <-rd.DO(conn, p.Mode, p.listInput)
		return &a
	}
	if p.keyInput.Input != nil {
		a.value = <-rd.DO(conn, p.Mode, p.keyInput)
		return &a
	}
	a.value = nil
	return &a
}
func (p *work) pipeHelper(conn redis.Conn) *convent {

	rd := DbContext{}

	a := p.convent.constructor()
	if p.hashInput.Input != nil {
		for i := range rd.Pipe(conn, p.Mode, p.hashInput) {
			a.value = i
		}
		return &a
	}
	if p.setInput.Input != nil {

		for i := range rd.Pipe(conn, p.Mode, p.setInput) {
			a.value = i
		}
		return &a
	}

	if p.listInput.Input != nil {

		for i := range rd.Pipe(conn, p.Mode, p.listInput) {
			a.value = i
		}
		return &a
	}
	if p.keyInput.Input != nil {
		for i := range rd.Pipe(conn, p.Mode, p.keyInput) {
			a.value = i
		}
		return &a
	}
	a.value = nil

	return &a
}

func (p *work) Value() EF.Container {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.value
}

func (p *work) DO(DBnumber int) *convent {
	return p.doHelper(p.connHelper(DBnumber))
}

func (p *work) Pipe(DBnumber int) *convent {

	return p.pipeHelper(p.connHelper(DBnumber))

}

func (p *work) PipeTWice(DBnumber int, twice EF.Container) chan *convent {
	rd := DbContext{}

	p.lock.Lock()

	if p.hashInput.Input != nil {
		ch := make(chan *convent)
		var input []EF.Container
		input = append(input, p.hashInput)
		input = append(input, twice)
		ok := make(chan bool)
		go func() {
			<-ok
			close(ch)
		}()
		go func() {
			for i := range rd.PipeTwice(DBnumber, input) {
				a := p.convent.constructor()
				a.value = i

				ch <- &a

			}
			ok <- true
		}()

		defer p.lock.Unlock()
		return ch
	}
	if p.setInput.Input != nil {

		ch := make(chan *convent)
		var input []EF.Container
		input = append(input, p.setInput)
		input = append(input, twice)
		ok := make(chan bool)
		go func() {
			for i := range rd.PipeTwice(DBnumber, input) {
				a := p.convent.constructor()
				a.value = i
				ch <- &a

			}
			ok <- true
		}()
		go func() {
			<-ok
			close(ch)
		}()
		defer p.lock.Unlock()
		return ch
	}

	if p.listInput.Input != nil {

		ch := make(chan *convent)
		ok := make(chan bool)
		var input []EF.Container
		input = append(input, p.listInput)
		input = append(input, twice)

		go func() {
			for i := range rd.PipeTwice(DBnumber, input) {
				a := p.convent.constructor()
				a.value = i
				ch <- &a

			}
			ok <- true
		}()
		go func() {
			<-ok
			close(ch)
		}()
		defer p.lock.Unlock()
		return ch

	}

	if p.keyInput.Input != nil {
		ch := make(chan *convent)
		var input []EF.Container
		input = append(input, p.keyInput)
		input = append(input, twice)
		ok := make(chan bool)
		go func() {
			for i := range rd.PipeTwice(DBnumber, input) {
				a := p.convent.constructor()
				a.value = i
				ch <- &a

			}
			ok <- true
		}()
		go func() {
			<-ok
			close(ch)
		}()
		defer p.lock.Unlock()
		return ch

	}
	a := p.convent.constructor()
	a.value = nil
	ch1 := make(chan *convent)
	ch1 <- &a

	defer p.lock.Unlock()
	return ch1
}

type Queue struct {
	convent convent
	lock    sync.Mutex
}

// func (p *Queue) QueuePipe(DBnumber int, twice []EF.Container) chan *convent {
// 	rd := DbContext{}
// 	p.lock.Lock()
// 	ch := make(chan *convent)
// 	ok := make(chan bool)
// 	go func() {
// 		<-ok
// 		close(ch)
// 	}()
// 	go func() {
// 		for i := range rd.Pipe(DBnumber, twice...) {
// 			a := p.convent.constructor()
// 			a.value = i
// 			//	fmt.Println(a)
// 			ch <- &a
// 			//fmt.Println("hi")
// 		}
// 		ok <- true
// 	}()

// 	defer p.lock.Unlock()
// 	return ch
// }

// 封印
// func (p *Queue) QueuePipeTWice(DBnumber int, twice ...[]EF.Container) chan *convent {
// 	rd := db.DbContext{}
// 	p.lock.Lock()
// 	ch := make(chan *convent)
// 	ok := make(chan bool)
// 	go func() {
// 		for i := range rd.PipeTwice(DBnumber, twice...) {
// 			a := p.convent.constructor()
// 			a.value = i
// 			ch <- &a
// 		}
// 		ok <- true

// 	}()
// 	go func() {
// 		<-ok
// 		close(ch)
// 	}()
// 	defer p.lock.Unlock()
// 	return ch
// }

type convent struct {
	lock  sync.Mutex
	value interface{}
}

func (p *convent) constructor() convent {
	a := convent{}
	return a
}

func (p *convent) Int64() (int64, error) {

	p.lock.Lock()
	defer p.lock.Unlock()

	switch p.value.(type) {
	case int64:
		return p.value.(int64), nil
	case string:
		return 0, errors.New("type == string ,Value conving err")
	case nil:
		return 0, errors.New("type == nil ,Value conving err")
	case redis.Error:
		return 0, p.value.(redis.Error)
	case error:
		return 0, p.value.(error)
	default:
		return 0, errors.New("type == nuknown   ,Value conving err")
	}

}
func (p *convent) Int64ToString() (string, error) {

	p.lock.Lock()
	defer p.lock.Unlock()

	switch p.value.(type) {
	case int64:
		a := strconv.FormatInt(p.value.(int64), 10)
		return a, nil
	case string:
		return "", errors.New("type == string ,Value conving err")
	case nil:
		return "", errors.New("type == nil ,Value conving err")
	case redis.Error:
		return "", p.value.(redis.Error)
	case error:
		return "", p.value.(error)
	default:
		return "", errors.New("type == nuknown   ,Value conving err")
	}

}
func (p *convent) Value() (interface{}, error) {

	p.lock.Lock()
	defer p.lock.Unlock()
	switch p.value.(type) {
	case interface{}:
		return p.value, nil
	case string:
		return nil, errors.New("type == string ,Value conving err")
	case nil:
		return p.value, nil
	case redis.Error:
		return nil, p.value.(redis.Error)
	case error:
		return nil, p.value.(error)

	default:
		return "", errors.New("type == nuknown   ,Value conving err")
	}

}

func (p *convent) ByteArray() ([]byte, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	data, ok := p.value.([]byte)
	if !ok {
		return nil, errors.New("type == nuknown  ,Value conving err")
	}
	return data, nil
}

func (p *convent) ToString() (string, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	switch p.value.(type) {
	case string:
		a := p.value.(string)
		return a, nil
	case nil:
		return "", errors.New("type == nil ,Value conving err")
	case redis.Error:
		return "", p.value.(redis.Error)
	case []byte:
		return "", errors.New("type == []byte ,Value conving err")
	case error:
		return "", p.value.(error)
	default:
		return "", errors.New("type == nuknown   ,Value conving err")
	}

}

func (p *convent) ByteTostring() (string, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	switch p.value.(type) {
	case []byte:
		return string(p.value.([]byte)), nil
	case string:
		return p.value.(string), nil
	case nil:
		return "", errors.New("type == nil ,Value conving err")
	case redis.Error:
		return "", p.value.(redis.Error)
	case error:
		return "", p.value.(error)
	default:
		return "", errors.New("type == nuknown   ,Value conving err")
	}

}
