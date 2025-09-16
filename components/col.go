package components

import (
	"strconv"

	"golang.org/x/net/html"
)

type col struct {
	comp *Component
}

func (c *col) Span(span int) *col {
	c.comp.addAttribute("span", strconv.Itoa(span))
	return c
}

func (c *col) Component() *Component {
	return c.comp
}

func Col(child *Component) *col {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "col",
		},
	}

	comp.Children(child)
	c := &col{comp: comp}
	return c
}
