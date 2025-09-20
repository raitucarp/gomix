package element

import "golang.org/x/net/html"

type cite struct {
	el *HtmlElement
}

func (c *cite) Component()            {}
func (c *cite) isHtml()               {}
func (c *cite) element() *HtmlElement { return c.el }
func (c *cite) Render() string        { return c.el.Render() }

func Cite(child *HtmlElement) *cite {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "cite",
		},
	}

	el.Children(child)
	c := &cite{el: el}
	return c
}
