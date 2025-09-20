package hyperscript

type possesiveExpression interface {
	Props() string
	Attribute() *attributeRef
}

type hasAttribute interface {
	Attribute() *attributeRef
}

type myExpression struct {
	property  string
	attribute *attributeRef
}

func My() *myExpression {
	return &myExpression{attribute: &attributeRef{attached: "my"}}
}

func (p *myExpression) Props() string {
	return ""
}

func (p *myExpression) Attribute() *attributeRef {
	return p.attribute
}
