package prototype

import (
	"encoding/json"
	"time"
)

type Prototype interface {
	Clone() Prototype
}

func NewConcretePrototype(code string, nums []int, updatedAt time.Time) *ConcretePrototype {
	return &ConcretePrototype{
		Code:      code,
		Nums:      nums,
		UpdatedAt: updatedAt,
	}
}

type ConcretePrototype struct {
	Code      string
	Nums      []int
	UpdatedAt time.Time
}

// deep copy
func (p *ConcretePrototype) Clone() Prototype {
	var clone ConcretePrototype
	b, _ := json.Marshal(p)
	json.Unmarshal(b, &clone)
	return &clone
}

type ConcretePrototypeList []*ConcretePrototype

// deep copy
func (p ConcretePrototypeList) Clone() Prototype {
	clone := make(ConcretePrototypeList, len(p))

	for i, v := range p {
		clone[i] = v.Clone().(*ConcretePrototype)
	}
	return clone
}
