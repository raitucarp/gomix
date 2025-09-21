package element

import "golang.org/x/net/html"

type canvas struct {
	el *HtmlElement
}

func (c *canvas) Component()            {}
func (c *canvas) isHtml()               {}
func (c *canvas) Element() *HtmlElement { return c.el }
func (c *canvas) Render() string        { return c.el.Render() }

func Canvas(message string) *canvas {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "canvas",
		},
	}

	el.Children(Text(message))
	c := &canvas{el: el}
	return c
}
