package proxy

import (
	"fmt"
	"sync"
)

type Subject interface {
	Request(i int)
}

func NewRealSubject() Subject {
	return RealSubject{}
}

type RealSubject struct {
}

func (s RealSubject) Request(i int)  {
	fmt.Println(i, "do sth")
}

func NewProxy() Subject {
	return &Proxy{subject: NewRealSubject()}
}

type Proxy struct {
	lock sync.Mutex
	subject Subject
}

func (p *Proxy) Request(i int)  {
	p.lock.Lock()
	p.subject.Request(i)
	p.lock.Unlock()
}
