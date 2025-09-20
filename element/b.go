package element

import "golang.org/x/net/html"

type bold struct {
	el *HtmlElement
}

func (b *bold) Children(children ...IsElement) *bold {
	b.el.Children(children...)
	return b
}
func (b *bold) Component()            {}
func (b *bold) isHtml()               {}
func (b *bold) element() *HtmlElement { return b.el }
func (b *bold) Render() string        { return b.el.Render() }

func B(child *HtmlElement) *bold {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "b",
		},
	}

	el.Children(child)
	b := &bold{el: el}
	return b
}
