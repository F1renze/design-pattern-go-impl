package builder

import "fmt"

func NewProduct() *Product {
	p := &Product{}
	p.parts = make([]string, 0, 2)
	return p
}

type Product struct {
	parts []string
}

func (p *Product) Add(part string) {
	p.parts = append(p.parts, part)
}

func (p *Product) Show() {
	fmt.Println("show all parts of product")
	for _, v := range p.parts {
		fmt.Printf("%s\t", v)
	}
	fmt.Println()
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	GetResult() *Product
}

func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{product: NewProduct()}
}

type ConcreteBuilder struct {
	product *Product
}

func (b ConcreteBuilder) BuildPartA() {
	b.product.Add("Part A")
}

func (b ConcreteBuilder) BuildPartB() {
	b.product.Add("Part B")
}

func (b ConcreteBuilder) GetResult() *Product {
	return b.product
}

type Director struct {
}

func (d Director) Construct(builder Builder) {
	builder.BuildPartA()
	builder.BuildPartB()
}
