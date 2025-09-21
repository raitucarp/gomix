package element

import "golang.org/x/net/html"

type code struct {
	el *HtmlElement
}

func (c *code) Component()            {}
func (c *code) isHtml()               {}
func (c *code) Element() *HtmlElement { return c.el }
func (c *code) Render() string        { return c.el.Render() }

func Code(child *HtmlElement) *code {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "code",
		},
	}

	el.Children(child)
	c := &code{el: el}
	return c
}
