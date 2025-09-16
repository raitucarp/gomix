package components

import "golang.org/x/net/html"

type dd struct {
	comp *Component
}

func (d *dd) Component() *Component {
	return d.comp
}

func Dd(child *Component) *dd {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "dd",
		},
	}

	comp.Children(child)
	d := &dd{comp: comp}
	return d
}
