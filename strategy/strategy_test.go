package strategy

import (
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	c := NewContest()
	c.CashContext(Rebate)
	fmt.Println(c.GetResult(99))
}