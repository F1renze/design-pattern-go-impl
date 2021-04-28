package decorator

import "fmt"

type IComponent interface {
	Operation()
}

func NewComponent() IComponent {
	return &Component{}
}

type Component struct {
}

func (c *Component) Operation() {
	fmt.Println("component do sth")
}

// ------- decorators
type IDecorator interface {
	SetComponent(c IComponent)
	Operation()
}

func NewDecorator(code int) IDecorator {
	return &DecoratorA{code: code}
}

type DecoratorA struct {
	code int
	c    IComponent
}

func (d *DecoratorA) SetComponent(c IComponent) {
	d.c = c
}

func (d *DecoratorA) Operation() {
	fmt.Println("decorator ", d.code, " setup")
	d.c.Operation()
	fmt.Println("decorator ", d.code, " teardown")
}
