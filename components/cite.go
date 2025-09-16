package components

import "golang.org/x/net/html"

type cite struct {
	comp *Component
}

func (c *cite) Component() *Component {
	return c.comp
}

func Cite(child *Component) *cite {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "cite",
		},
	}

	comp.Children(child)
	c := &cite{comp: comp}
	return c
}
