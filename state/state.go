package state

import (
	"fmt"
	"reflect"
)

type IState interface {
	Handle(ctx IContext)
}

type ConcreteStateA struct {
}

func (s ConcreteStateA) Handle(ctx IContext) {
	ctx.SetState(ConcreteStateB{})
}

type ConcreteStateB struct {
}

func (s ConcreteStateB)Handle(ctx IContext)  {
	ctx.SetState(ConcreteStateA{})
}

type IContext interface {
	GetState() IState
	SetState(state IState)
	Request()
}

func NewContext(s IState) IContext {
	return &Context{state: s}
}

type Context struct {
	state IState
}

func (c *Context)GetState() IState {
	return c.state
}

func (c *Context)SetState(s IState)  {
	c.state = s
	fmt.Println("current state: ", reflect.TypeOf(s))
}

func (c *Context)Request()  {
	c.state.Handle(c)
}
