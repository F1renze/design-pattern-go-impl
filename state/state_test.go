package state

import "testing"

func TestContext(t *testing.T) {
	c := NewContext(ConcreteStateA{})

	c.Request()
	c.Request()
	c.Request()
	c.Request()
}
