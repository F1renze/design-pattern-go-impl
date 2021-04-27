package strategy

import "math"

type IStrategy interface {
	AcceptCash(money float64) float64
}

// ----- strategy impl

type CashNormal struct {
}

func (c CashNormal) AcceptCash(money float64) float64 {
	return money
}

type CashRebate struct {
	moneyRebate float64
}

func (c CashRebate) AcceptCash(money float64) float64 {
	return money * c.moneyRebate
}

type CashReturn struct {
	moneyCondition float64
	moneyReturn    float64
}

func (c CashReturn) AcceptCash(money float64) float64 {
	if money >= c.moneyCondition {
		money -= math.Floor(money/c.moneyCondition) * c.moneyReturn
	}
	return money
}

// ----- strategy end

type Kind int

const (
	Normal Kind = iota
	Rebate
	Return
)

func NewContest() Context{
	return Context{}
}

type Context struct {
	strategy IStrategy
}

func (c Context) CashContext(kind Kind) {
	switch kind {
	case Normal:
		c.strategy = CashNormal{}
	case Rebate:
		// TODO param
		c.strategy = CashRebate{}
	case Return:
		c.strategy = CashReturn{}
	}
}

func (c Context) GetResult(money float64) float64 {
	return c.strategy.AcceptCash(money)
}
