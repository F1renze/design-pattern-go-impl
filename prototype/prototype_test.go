package prototype

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrototype(t *testing.T) {
	now := time.Now()
	nums := []int{2, 4, 5}
	p := NewConcretePrototype("01", nums, now)
	clone := p.Clone().(*ConcretePrototype)
	assert.True(t, clone.UpdatedAt.Equal(p.UpdatedAt))
	assert.Equal(t, p.Code, clone.Code)
	assert.Equal(t, p.Nums, clone.Nums)
}
