package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type colgroup struct {
	el *HtmlElement
}

func (c *colgroup) Span(span int) *colgroup {
	c.el.AddAttribute("span", strconv.Itoa(span))
	return c
}

func (c *colgroup) Component()            {}
func (c *colgroup) isHtml()               {}
func (c *colgroup) element() *HtmlElement { return c.el }
func (c *colgroup) Render() string        { return c.el.Render() }

func ColGroup(child *HtmlElement) *colgroup {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "colgroup",
		},
	}

	el.Children(child)
	c := &colgroup{el: el}
	return c
}
