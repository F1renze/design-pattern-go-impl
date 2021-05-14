package command

import "testing"

func TestCommand(t *testing.T)  {
	r := NewReceiver("ikun")

	singCommand := &SingCommand{receiver: r}
	danceCommand := &DanceCommand{receiver: r}
	playBasketBallCommand := &PlayBasketBallCommand{receiver: r}

	i := NewInvoker()
	i.SetCommand(singCommand)
	i.SetCommand(danceCommand)
	i.SetCommand(playBasketBallCommand)

	i.CancelCommand(playBasketBallCommand)

	i.Run()
}
