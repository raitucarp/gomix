package components

import "golang.org/x/net/html"

type body struct {
	comp *Component
}

func (b *body) Component() *Component {
	return b.comp
}

func Body(children ...*Component) *body {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "body",
		},
	}

	comp.Children(children...)
	b := &body{comp: comp}
	return b
}
