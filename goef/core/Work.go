package RedigoEFcore

import (
	"errors"
	"strconv"
	"sync"
	EF "github.com/ReadyCore/goef/other"
	"github.com/gomodule/redigo/redis"
)

type work struct {
	convent    convent
	lock       sync.Mutex
	Mode       string
	connKey    string
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

func (p *work) Value() EF.Container {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.value
}

func (p *work) DO(DBnumber int) *convent {
	return p.doHelper(DBnumber)
}

func (p *work) Pipe(DBnumber int) *convent {
	return p.pipeHelper(DBnumber)
}

func (p *work) PipeTWice(DBnumber int, twice EF.Container) chan *convent {

	if p.hashInput.Input != nil {
		return p.pipeTWiceHelper(DBnumber, p.hashInput, twice)
	}
	if p.setInput.Input != nil {
		return p.pipeTWiceHelper(DBnumber, p.setInput, twice)
	}
	if p.listInput.Input != nil {
		return p.pipeTWiceHelper(DBnumber, p.listInput, twice)
	}
	if p.keyInput.Input != nil {
		return p.pipeTWiceHelper(DBnumber, p.keyInput, twice)
	}
	a := p.convent.constructor()
	a.value = nil
	ch1 := make(chan *convent)
	ch1 <- &a
	return ch1
}

////////  convent 結構設定

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
		return 0, errors.New("type == unknown   ,Value conving err")
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
		return "", errors.New("type == unknown  ,Value conving err")
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
		return "", errors.New("type == unknown   ,Value conving err")
	}

}

func (p *convent) ByteArray() ([]byte, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	data, ok := p.value.([]byte)
	if !ok {
		return nil, errors.New("type == unknown  ,Value conving err")
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
		return "", errors.New("type == unknown   ,Value conving err")
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
		return "", errors.New("type == unknown   ,Value conving err")
	}

}
