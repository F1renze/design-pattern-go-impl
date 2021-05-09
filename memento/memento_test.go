package memento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemento(t *testing.T) {
	r := &Runner{
		in: &InputText{},
		sh: NewSnapshotHolder(),
	}

	tt := []struct {
		In  string
		Out string
	}{
		{In: "hello", Out: ""},
		{In: ":list", Out: "hello"},
		{In: "world", Out: ""},
		{In: ":list", Out: "helloworld"},
		{In: ":undo", Out: ""},
		{In: ":list", Out: "hello"},
	}

	for _, v := range tt {
		got := r.Run(v.In)
		if len(v.Out) > 0 {
			assert.Equal(t, got, v.Out)
		}
	}
}
