package element

import "golang.org/x/net/html"

type br struct {
	el *HtmlElement
}

func (b *br) Component()            {}
func (b *br) isHtml()               {}
func (b *br) element() *HtmlElement { return b.el }
func (b *br) Render() string        { return b.el.Render() }

func Br() *br {
	comp := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "br",
		},
	}

	b := &br{el: comp}
	return b
}
