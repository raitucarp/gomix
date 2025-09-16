package components

import "golang.org/x/net/html"

type bdi struct {
	comp *Component
}

func (b *bdi) Component() *Component {
	return b.comp
}

func Bdi(child *Component) *bdi {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "bdi",
		},
	}

	comp.Children(child)
	b := &bdi{comp: comp}
	return b
}
