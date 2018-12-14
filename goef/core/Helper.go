package RedigoEFcore

import (
	EF "github.com/ReadyCore/goef/other"
)

func (p *work) doHelper(dbnumber int) *convent {
	rd := dbContext{}
	a := p.convent.constructor()
	if p.setInput.Input != nil {
		a.value = <-rd.DO(p.connKey, p.Mode, dbnumber, p.setInput)
		return &a
	}
	if p.hashInput.Input != nil {
		a.value = <-rd.DO(p.connKey, p.Mode, dbnumber, p.hashInput)
		return &a
	}
	if p.listInput.Input != nil {
		a.value = <-rd.DO(p.connKey, p.Mode, dbnumber, p.listInput)
		return &a
	}
	if p.keyInput.Input != nil {
		a.value = <-rd.DO(p.connKey, p.Mode, dbnumber, p.keyInput)
		return &a
	}
	a.value = nil
	return &a
}
func (p *work) pipeHelper(dbnumber int) *convent {
	rd := dbContext{}
	a := p.convent.constructor()
	if p.hashInput.Input != nil {
		for i := range rd.Pipe(p.connKey, p.Mode, dbnumber, p.hashInput) {
			a.value = i
		}
		return &a
	}
	if p.setInput.Input != nil {
		for i := range rd.Pipe(p.connKey, p.Mode, dbnumber, p.setInput) {
			a.value = i
		}
		return &a
	}

	if p.listInput.Input != nil {

		for i := range rd.Pipe(p.connKey, p.Mode, dbnumber, p.listInput) {
			a.value = i
		}
		return &a
	}
	if p.keyInput.Input != nil {
		for i := range rd.Pipe(p.connKey, p.Mode, dbnumber, p.keyInput) {
			a.value = i
		}
		return &a
	}
	a.value = nil

	return &a
}

func (p *work) pipeTWiceHelper(DBnumber int, in, twice EF.Container) chan *convent {
	rd := dbContext{}
	ch := make(chan *convent)
	input := []EF.Container{in, twice}
	ok := make(chan bool)

	go func() {
		for i := range rd.PipeTwice(p.connKey, p.Mode, DBnumber, input) {
			a := p.convent.constructor()
			a.value = i
			ch <- &a
		}
		ok <- true
	}()
	go func() {
		<-ok
		close(ch)
		close(ok)
	}()
	return ch
}
