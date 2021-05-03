package template

import (
	"errors"
	"fmt"
)

type IPay interface {
	Pay() error
}

type WechatPay struct {
}

func (p WechatPay)Pay() error {
	fmt.Println("微信支付中...")
	return nil
}

type AliPay struct {
}

func (p AliPay)Pay() error {
	fmt.Println("支付宝支付中...")
	return nil
}

type UnionPay struct {
}

func (p UnionPay)Pay() error {
	fmt.Println("银联支付中...")
	return nil
}

func NewPayTemplate() *PayTemplate {
	return &PayTemplate{subMap: map[string]IPay{
		"wechat": WechatPay{},
		"alipay": AliPay{},
		"union": UnionPay{},
	}}
}

var ErrPaymentNotExist = errors.New("payment not exist")

type PayTemplate struct {
	subMap map[string]IPay
}

func (t *PayTemplate)Pay(payment string) error {
	p, ok := t.subMap[payment]
	if !ok {
		return ErrPaymentNotExist
	}
	fmt.Println("开始支付")
	fmt.Println("更新订单状态...")
	return p.Pay()
}