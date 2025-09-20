package element

import "golang.org/x/net/html"

type body struct {
	el *HtmlElement
}

func (b *body) Children(children ...IsElement) *body {
	b.el.Children(children...)
	return b
}
func (b *body) Component()            {}
func (b *body) isHtml()               {}
func (b *body) Element() *HtmlElement { return b.el }
func (b *body) Render() string        { return b.el.Render() }

func Body(children ...IsElement) *body {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "body",
		},
	}

	el.Children(children...)
	b := &body{el: el}
	return b
}
