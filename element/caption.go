package element

import "golang.org/x/net/html"

type caption struct {
	el *HtmlElement
}

func (c *caption) Component()            {}
func (c *caption) isHtml()               {}
func (c *caption) Element() *HtmlElement { return c.el }
func (c *caption) Render() string        { return c.el.Render() }

func Caption(child *HtmlElement) *caption {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "caption",
		},
	}

	el.Children(child)
	c := &caption{el: el}
	return c
}
