package components

import "golang.org/x/net/html"

type dfn struct {
	comp *Component
}

func (d *dfn) Component() *Component {
	return d.comp
}

func Dfn(child *Component) *dfn {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "dfn",
		},
	}

	comp.Children(child)
	d := &dfn{comp: comp}
	return d
}
