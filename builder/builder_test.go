package builder

import "testing"

func TestBuilder(t *testing.T) {
	d := Director{}

	b := NewConcreteBuilder()
	d.Construct(b)
	p := b.GetResult()
	p.Show()
}
