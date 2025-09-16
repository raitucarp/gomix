package components

import "golang.org/x/net/html"

type caption struct {
	comp *Component
}

func (c *caption) Component() *Component {
	return c.comp
}

func Caption(child *Component) *caption {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "caption",
		},
	}

	comp.Children(child)
	c := &caption{comp: comp}
	return c
}
