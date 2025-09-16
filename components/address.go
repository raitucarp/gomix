package components

import "golang.org/x/net/html"

type address struct {
	comp *Component
}

func (address *address) Component() *Component {
	return address.comp
}

func Address(child ...*Component) *address {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "address",
		},
	}

	comp.Children(child...)
	addr := &address{comp: comp}
	return addr
}
