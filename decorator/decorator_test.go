package decorator

import "testing"

func TestDecorator(t *testing.T) {
	c := NewComponent()
	d1 := NewDecorator(1)
	d2 := NewDecorator(2)

	d1.SetComponent(c)
	d2.SetComponent(d1)
	d2.Operation()
}
