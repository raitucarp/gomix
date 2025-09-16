package components

import (
	"strconv"

	"golang.org/x/net/html"
)

type colgroup struct {
	comp *Component
}

func (c *colgroup) Span(span int) *colgroup {
	c.comp.addAttribute("span", strconv.Itoa(span))
	return c
}

func (c *colgroup) Component() *Component {
	return c.comp
}

func ColGroup(child *Component) *colgroup {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "colgroup",
		},
	}

	comp.Children(child)
	c := &colgroup{comp: comp}
	return c
}
