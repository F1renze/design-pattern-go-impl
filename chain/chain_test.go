package chain

import "testing"

func TestChainOfResponsibility(t *testing.T)  {
	hc := &HandlerChain{}
	hc.AddHandler(NewHandlerA())
	hc.AddHandler(NewHandlerB())
	hc.Handle()
}
