package alpinejs

import "fmt"

type xFor struct {
	alpineRef *alpine
	varName   string
	indexName string
	in        string
}

func (x *xFor) VariableName(name string) *xFor {
	x.varName = name
	return x
}

func (x *xFor) IndexName(name string) *xFor {
	x.indexName = name
	return x
}

func (x *xFor) In(in string) *xFor {
	x.in = in
	return x
}

func (x *xFor) Attribute() *alpine {
	x.alpineRef.el.AddAttribute("x-for", fmt.Sprintf("(%s, %s) in %s", x.varName, x.indexName, x.in))
	return x.alpineRef
}

func (a *alpine) For() *xFor {
	return &xFor{alpineRef: a}
}

func (a *alpine) Key(key string) *alpine {
	a.el.AddAttribute(":key", key)
	return a
}
