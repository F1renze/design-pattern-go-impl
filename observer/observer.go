package observer

import "fmt"

type ISubject interface {
	Attach(observer IObserver)
	Detach(observer IObserver)
	Notify()

	GetSubjectState() string
	SetSubjectState(s string)
}

func NewSubject() *Subject {
	return &Subject{[]IObserver{}}
}

// 父类，代码复用
type Subject struct {
	observers []IObserver
}

func (s *Subject) Attach(o IObserver) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Detach(o IObserver) {
	//TODO
}

func (s *Subject) Notify() {
	for _, o := range s.observers {
		o.Update()
	}
}

func (s *Subject) GetSubjectState() string {
	return ""
}

func (s *Subject) SetSubjectState(state string) {
	return
}

func NewConcreteSubject() ISubject {
	return &ConcreteSubject{
		Subject:      NewSubject(),
		SubjectState: "",
	}
}

// 子类重写 GetSubjectState 方法
type ConcreteSubject struct {
	*Subject

	SubjectState string
}

func (s *ConcreteSubject) GetSubjectState() string {
	return s.SubjectState
}

func (s *ConcreteSubject) SetSubjectState(state string) {
	s.SubjectState = state
}

// 观察者
type IObserver interface {
	Update()
}

func NewConcreteObserver(name string, subject ISubject) IObserver {
	return &ConcreteObserver{
		name:          name,
		observerState: "",
		subject:       subject,
	}
}

type ConcreteObserver struct {
	name          string
	observerState string
	subject       ISubject
}

func (o *ConcreteObserver) Update() {
	o.observerState = o.subject.GetSubjectState()
	fmt.Printf("New state of observer '%s' is '%s'\n", o.name, o.observerState)
}
