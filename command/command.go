package command

import "fmt"

// TODO use func as command
type ICommand interface {
	Execute()
}

// IReceiver contains all commands
type IReceiver interface {
	Sing()
	Dance()
	PlayBasketBall()
}

type IInvoker interface {
	SetCommand(c ICommand)
	CancelCommand(c ICommand)
	Run()
}

func NewInvoker() IInvoker {
	return &Invoker{commands: make(map[ICommand]bool)}
}

type Invoker struct {
	commands map[ICommand]bool
}

func (i *Invoker) SetCommand(c ICommand) {
	i.commands[c] = true
}

func (i *Invoker) CancelCommand(c ICommand) {
	if _, ok := i.commands[c]; !ok {
		return
	}
	delete(i.commands, c)
}

func (i *Invoker) Run() {
	for k, _ := range i.commands {
		k.Execute()
	}
}

func NewReceiver(name string) IReceiver {
	return &Receiver{name: name}
}

type Receiver struct {
	name string
}

func (r *Receiver) Sing() {
	fmt.Println("姬霓太美")
}

func (r *Receiver) Dance() {
	fmt.Println("唱、跳 rap，⛹")
}

func (r *Receiver) PlayBasketBall() {
	fmt.Println("不会打篮球⛹")
}


type SingCommand struct {
	receiver IReceiver
}

func (c *SingCommand) Execute() {
	c.receiver.Sing()
}

type DanceCommand struct {
	receiver IReceiver
}

func (c *DanceCommand)Execute()  {
	c.receiver.Dance()
}

type PlayBasketBallCommand struct {
	receiver IReceiver
}

func (c *PlayBasketBallCommand) Execute() {
	c.receiver.PlayBasketBall()
}


