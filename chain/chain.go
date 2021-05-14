package chain

import "fmt"

type IHandler interface {
	SetSuccessor(handler IHandler)
	Handle()
	doHandle() bool
}

// HandlerTemplate
type HandlerTemplate struct {
	sub       IHandler
	successor IHandler
}

func (t *HandlerTemplate) SetSuccessor(handler IHandler) {
	t.successor = handler
}

func (t *HandlerTemplate) Handle() {
	ok := t.sub.doHandle()
	if !ok && t.successor != nil {
		t.successor.Handle()
	}
}

func (t *HandlerTemplate)doHandle() bool {
	fmt.Println("handler template do not handle")
	return false
}

type HandlerChain struct {
	head IHandler
	tail IHandler
}

func (c *HandlerChain)AddHandler(h IHandler)  {
	h.SetSuccessor(nil)
	if c.head == nil {
		c.head, c.tail =  h, h
	}
	c.tail.SetSuccessor(h)
	c.tail = h
}

func (c *HandlerChain)Handle()  {
	if c.head != nil {
		c.head.Handle()
	}
}

func NewHandlerA() IHandler {
	h := &HandlerA{&HandlerTemplate{}}
	h.sub = h
	return h
}

type HandlerA struct {
	*HandlerTemplate
}

func (h *HandlerA) doHandle() bool {
	fmt.Println("HandlerA do sth interesting 🕺")
	return false
}

func NewHandlerB() IHandler {
	h := &HandlerB{&HandlerTemplate{}}
	h.sub = h
	return h
}

type HandlerB struct {
	*HandlerTemplate
}

func (h *HandlerB) doHandle() bool {
	fmt.Println("HandlerB do sth 👊")
	return false
}
