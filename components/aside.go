package components

import "golang.org/x/net/html"

type aside struct {
	comp *Component
}

func (a *aside) Component() *Component {
	return a.comp
}

func Aside(child ...*Component) *aside {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "aside",
		},
	}

	comp.Children(child...)
	as := &aside{comp: comp}
	return as
}
