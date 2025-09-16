package components

import "golang.org/x/net/html"

type bold struct {
	comp *Component
}

func (b *bold) Component() *Component {
	return b.comp
}

func B(child *Component) *bold {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "b",
		},
	}

	comp.Children(child)
	b := &bold{comp: comp}
	return b
}
