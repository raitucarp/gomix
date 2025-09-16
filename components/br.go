package components

import "golang.org/x/net/html"

type br struct {
	comp *Component
}

func (b *br) Component() *Component {
	return b.comp
}

func Br() *br {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "br",
		},
	}

	b := &br{comp: comp}
	return b
}
