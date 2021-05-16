package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestMediator(t *testing.T) {
	sn := "Service A"
	srv := NewServiceA(sn)

	m := NewMediator()
	m.Register(sn, srv)

	reply, err := m.Call(sn, "Hello")
	assert.Nil(t, err, nil)
	assert.NotEmpty(t, reply)
	fmt.Println(reply)

	reply, err = m.Call("not found", "hello")
	assert.ErrorIs(t, ErrServiceNotFound, err)
}
