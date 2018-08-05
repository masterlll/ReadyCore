package RedigoEFcore

import (
	
	db"ReadyCore/ReadyCore/goef/RedisDB"	
	EF "ReadyCore/ReadyCore/goef/other"
	"sync"
)

type work struct {
	convent    convent
	lock       sync.Mutex
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

		a.value = <-rd.DO(DBnumber, p.setInput)
		defer p.lock.Unlock()
		return &a
	}

	if p.hashInput.Input != nil {
		a.value = <-rd.DO(DBnumber, p.hashInput)
		defer p.lock.Unlock()
		return &a
	}

	if p.listInput.Input != nil {

		a.value = <-rd.DO(DBnumber, p.listInput)
		defer p.lock.Unlock()
		return &a

	}

	if p.keyInput.Input != nil {
		a.value = <-rd.DO(DBnumber, p.keyInput)
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

// func (p *work) Queue() string {

// 	if p.hashInput.Input != nil {

// 		for _, i := range p.hashInput.Input {
// 			fmt.Println("DO", i)
// 			return "DO" + i.(string)
// 		}
// 	}
// 	if p.setInput.Input != nil {
// 		for _, i := range p.setInput.Input {
// 			return "DO" + i.(string)
// 		}
// 	}

// 	if p.listInput.Input != nil {

// 		for _, i := range p.listInput.Input {
// 			fmt.Println(i)
// 			return "DO" + i.(string)

// 		}

// 	}
// 	if p.keyInput.Input != nil {

// 		for _, i := range p.keyInput.Input {
// 			fmt.Println(i)
// 			return "DO" + i.(string)
// 		}
// 	}
// 	return "no found "
// }

/////////

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
