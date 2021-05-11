package component

import "testing"

func TestComponent(t *testing.T)  {
	root := NewComponent("root")

	com1 := NewComponent("com1")
	com2 := NewComponent("com2")

	root.Add(com1)
	root.Add(com2)

	com1.Add(NewLeaf("leaf1"))

	com2.Add(NewLeaf("leaf2"))

	root.Display(1)
}
