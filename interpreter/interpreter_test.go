package interpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlertRuleInterpreter_Interpret(t *testing.T) {
	//str := "(count > 1000 && timeout > 100) || (error > 50 && delay > 10)"
	expr := "count > 1000 && timeout > 100 || error > 50"
	interpreter := NewAlertRuleInterpreter()
	collector := NewMetricCollector()

	collector.AddKey("count", 1001)
	collector.AddKey("timeout", 101)
	collector.AddKey("error", 51)

	interpreter.Add(expr)
	got := interpreter.Interpret(collector.GetStorage())
	assert.Equal(t, got, true)
}