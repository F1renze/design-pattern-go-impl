package component

import (
	"fmt"
	"strings"
)

type IComponent interface {
	Add(component IComponent)
	Remove(component IComponent)
	Display(depth int)
}

func NewComponent(name string) IComponent {
	return &Component{name: name, children: make(map[IComponent]bool)}
}

type Component struct {
	name     string
	children map[IComponent]bool
}

func (c *Component) Add(component IComponent) {
	c.children[component] = true
}

func (c *Component) Remove(component IComponent) {
	if _, ok := c.children[component]; !ok {
		return
	}
	delete(c.children, component)
}

func (c *Component) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth) + c.name)
	for k, _ := range c.children {
		k.Display(depth + 2)
	}
}

func NewLeaf(name string) IComponent {
	return &Leaf{name: name}
}

type Leaf struct {
	name string
}

func (l *Leaf) Add(component IComponent) {
	fmt.Println("cannot add to a leaf")
}

func (l *Leaf) Remove(component IComponent) {
	fmt.Println("cannot remove from a leaf")
}

func (l *Leaf) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth)+l.name)
}

