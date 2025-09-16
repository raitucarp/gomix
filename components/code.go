package components

import "golang.org/x/net/html"

type code struct {
	comp *Component
}

func (c *code) Component() *Component {
	return c.comp
}

func Code(child *Component) *code {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "code",
		},
	}

	comp.Children(child)
	c := &code{comp: comp}
	return c
}
