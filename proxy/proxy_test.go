package proxy

import "testing"

func TestProxy(t *testing.T)  {
	p := NewProxy()
	go func() {
		p.Request(1)
	}()
	go func() {
		p.Request(2)
	}()

	p.Request(3)
}