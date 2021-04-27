package strategy

import "testing"

func TestContext(t *testing.T) {
	c := NewContest()
	c.CashContext(Normal)
	c.GetResult(99)
}