package RedigoEFcore

import (
	"sync"

	db "github.com/ReadyCore/goef/RedisDB"
	EF "github.com/ReadyCore/goef/other"
)

type work struct {
	convent    convent
	lock       sync.Mutex
	Mode       int
	value      EF.Container
	hashInput  EF.Container
	setInput   EF.Container
	listInput  EF.Container
	keyInput   EF.Container
	otherInput EF.Container
}

func (p *work) constructor() *work {
	return p
}
func (p *work) Value() EF.Container {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.value
}

func (p *work) DO(DBnumber int) *convent {
	rd := db.DbContext{}

	p.lock.Lock()
	a := p.convent.constructor()
	if p.setInput.Input != nil {

		a.value = <-rd.DO(DBnumber, p.setInput, p.Mode)
		defer p.lock.Unlock()
		return &a
	}

	if p.hashInput.Input != nil {
		a.value = <-rd.DO(DBnumber, p.hashInput, p.Mode)
		defer p.lock.Unlock()
		return &a
	}

	if p.listInput.Input != nil {

		a.value = <-rd.DO(DBnumber, p.listInput, p.Mode)
		defer p.lock.Unlock()
		return &a

	}

	if p.keyInput.Input != nil {
		a.value = <-rd.DO(DBnumber, p.keyInput, p.Mode)
		defer p.lock.Unlock()
		return &a
	}

	a.value = nil

	defer p.lock.Unlock()
	return &a
}

func (p *work) Pipe(DBnumber int) *convent {
	rd := db.DbContext{}
	p.lock.Lock()
	a := p.convent.constructor()

	if p.hashInput.Input != nil {
		for i := range rd.Pipe(DBnumber, p.hashInput) {
			a.value = i
		}
		defer p.lock.Unlock()
		return &a
	}
	if p.setInput.Input != nil {

		for i := range rd.Pipe(DBnumber, p.setInput) {
			a.value = i
		}
		defer p.lock.Unlock()
		return &a
	}

	if p.listInput.Input != nil {

		for i := range rd.Pipe(DBnumber, p.listInput) {
			a.value = i

		}
		defer p.lock.Unlock()
		return &a

	}

	if p.keyInput.Input != nil {
		for i := range rd.Pipe(DBnumber, p.keyInput) {
			a.value = i
		}

		defer p.lock.Unlock()
		return &a

	}

	a.value = nil

	defer p.lock.Unlock()
	return &a
}

func (p *work) PipeTWice(DBnumber int, twice EF.Container) chan *convent {
	rd := db.DbContext{}

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

func (p *Queue) QueuePipe(DBnumber int, twice []EF.Container) chan *convent {
	rd := db.DbContext{}
	p.lock.Lock()
	ch := make(chan *convent)
	ok := make(chan bool)
	go func() {
		<-ok
		close(ch)
	}()
	go func() {
		for i := range rd.Pipe(DBnumber, twice...) {
			a := p.convent.constructor()
			a.value = i
			//	fmt.Println(a)
			ch <- &a
			//fmt.Println("hi")
		}
		ok <- true
	}()

	defer p.lock.Unlock()
	return ch
}

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

func (p *convent) Int64() int64 {

	p.lock.Lock()
	a := p.value.(int64)
	defer p.lock.Unlock()
	return a
}
func (p *convent) Value() interface{} {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.value
}
func (p *convent) ToString() string {
	p.lock.Lock()
	a := p.value.(string)
	defer p.lock.Unlock()
	return a
}
