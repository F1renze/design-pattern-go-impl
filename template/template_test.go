package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayTemplate(t *testing.T) {
	pt := NewPayTemplate()
	assert.NoError(t, pt.Pay("wechat"))
	assert.NoError(t, pt.Pay("alipay"))
	assert.NoError(t, pt.Pay("union"))
	assert.ErrorIs(t, pt.Pay("btc"), ErrPaymentNotExist)
}
