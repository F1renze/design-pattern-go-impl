package strategy

import (
	"math"
)

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

type IContext interface {
	CashContext(kind Kind)
	GetResult(money float64) float64
}

func NewContest() IContext{
	return &Context{}
}

type Context struct {
	strategy IStrategy
}

// use ptr receiver instead value receiver, so that changing can work
func (c *Context) CashContext(kind Kind) {
	switch kind {
	case Normal:
		c.strategy = CashNormal{}
	case Rebate:
		// TODO param
		c.strategy = CashRebate{moneyRebate: 0.8}
	case Return:
		c.strategy = CashReturn{moneyCondition: 300, moneyReturn: 50}
	}
}

func (c *Context) GetResult(money float64) float64 {
	return c.strategy.AcceptCash(money)
}
