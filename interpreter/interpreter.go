package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

type IExpression interface {
	Add(expression string) IExpression
	Interpret(stat map[string]int) bool
}

type IMetricCollector interface {
	AddKey(key string, num int)
	GetStorage() map[string]int
}

func NewAlertRuleInterpreter() *AlertRuleInterpreter {
	return &AlertRuleInterpreter{}
}

type AlertRuleInterpreter struct {
	expression IExpression
}

// Add 解析表达式
func (i *AlertRuleInterpreter) Add(expr string) IExpression {
	e := &OrExpression{}
	i.expression = e.Add(expr)
	return i
}

func (i *AlertRuleInterpreter) Interpret(stat map[string]int) bool {
	return i.expression.Interpret(stat)
}

func NewMetricCollector() IMetricCollector {
	return &MetricCollector{storage: map[string]int{}}
}

type MetricCollector struct {
	storage map[string]int
}

func (c *MetricCollector) AddKey(key string, num int) {
	c.storage[key] = num
}

func (c *MetricCollector) GetStorage() map[string]int {
	storageCopy := make(map[string]int, len(c.storage))

	for k, v := range c.storage {
		storageCopy[k] = v
	}
	return storageCopy
}

/**
假设自定义的告警规则只包含“||、&&、>、<、==”这

五个运算符，其中，“>、<、==”运算符的优先级高于“||、&&”运算符，“&&”运算

符优先级高于“||”。在表达式中，任意元素之间需要通过空格来分隔。
*/

type LogicExpression struct {
}

func (e LogicExpression) InterpretExpressions(stat map[string]int, expressions []IExpression) bool {
	for _, expr := range expressions {
		if !expr.Interpret(stat) {
			return false
		}
	}
	return true
}

type CompareExpression struct {
}

func (e CompareExpression) ParseKeyValue(expr, sep string) (key string, val int) {
	vals := strings.Split(expr, sep)
	if len(vals) != 2 {
		panic(fmt.Sprintf("expression '%s' is invalid", expr))
	}
	key = strings.TrimSpace(vals[0])
	val, _ = strconv.Atoi(strings.TrimSpace(vals[0]))
	return
}

func (e CompareExpression) CompareInterpret(stat map[string]int, key string, compare func(val int) bool) bool {
	val, ok := stat[key]
	if !ok {
		return false
	}
	return compare(val)
}

// OrExpression || 表达式
type OrExpression struct {
	LogicExpression

	expressions []IExpression
}

func (e *OrExpression) Add(expression string) IExpression {
	elements := strings.Split(expression, "||")
	var expr IExpression
	for _, strExpr := range elements {
		expr = &AndExpression{}
		e.expressions = append(e.expressions, expr.Add(strExpr))
	}
	return e
}

func (e *OrExpression) Interpret(stat map[string]int) bool {
	return e.InterpretExpressions(stat, e.expressions)
}

// AndExpression && 表达式
type AndExpression struct {
	LogicExpression

	expressions []IExpression
}

func (e *AndExpression) Add(expression string) IExpression {
	elements := strings.Split(expression, "&&")
	var expr IExpression
	for _, strExpr := range elements {
		switch {
		case strings.Contains(strExpr, ">"):
			expr = &GreaterExpression{}
		case strings.Contains(strExpr, "<"):
			expr = &LessExpression{}
		case strings.Contains(strExpr, "=="):
			expr = &EqualExpression{}
		default:
			panic(fmt.Sprintf("expression '%s' is invalid", strExpr))
		}
		e.expressions = append(e.expressions, expr.Add(strExpr))
	}
	return e
}

func (e *AndExpression) Interpret(stat map[string]int) bool {
	return e.InterpretExpressions(stat, e.expressions)
}

// GreaterExpression > 表达式
type GreaterExpression struct {
	CompareExpression

	key string
	val int
}

func (e *GreaterExpression) Add(expr string) IExpression {
	e.key, e.val = e.ParseKeyValue(expr, ">")
	return e
}

func (e *GreaterExpression) Interpret(stat map[string]int) bool {
	return e.CompareInterpret(stat, e.key, func(val int) bool {
		return val > e.val
	})
}

type LessExpression struct {
	CompareExpression

	key string
	val int
}

func (e LessExpression) Add(expr string) IExpression {
	e.key, e.val = e.ParseKeyValue(expr, "<")
	return e
}

func (e LessExpression) Interpret(stat map[string]int) bool {
	return e.CompareInterpret(stat, e.key, func(val int) bool {
		return val < e.val
	})
}

type EqualExpression struct {
	CompareExpression

	key string
	val int
}

func (e EqualExpression) Add(expr string) IExpression {
	e.key, e.val = e.ParseKeyValue(expr, "==")
	return e
}

func (e EqualExpression) Interpret(stat map[string]int) bool {
	return e.CompareInterpret(stat, e.key, func(val int) bool {
		return val == e.val
	})
}
