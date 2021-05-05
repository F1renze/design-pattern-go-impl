package observer

import "testing"

func TestObserver(t *testing.T) {
	s := NewConcreteSubject()

	s.Attach(NewConcreteObserver("A", s))
	s.Attach(NewConcreteObserver("B", s))
	s.Attach(NewConcreteObserver("C", s))

	s.SetSubjectState("Updated")

	s.Notify()
}
