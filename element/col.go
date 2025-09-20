package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type col struct {
	el *HtmlElement
}

func (c *col) Span(span int) *col {
	c.el.AddAttribute("span", strconv.Itoa(span))
	return c
}
func (c *col) Component()            {}
func (c *col) isHtml()               {}
func (c *col) element() *HtmlElement { return c.el }
func (c *col) Render() string        { return c.el.Render() }

func Col(child *HtmlElement) *col {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "col",
		},
	}

	el.Children(child)
	c := &col{el: el}
	return c
}
