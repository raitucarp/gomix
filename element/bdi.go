package element

import "golang.org/x/net/html"

type bdi struct {
	el *HtmlElement
}

func (b *bdi) Children(children ...IsElement) *bdi {
	b.el.Children(children...)
	return b
}
func (b *bdi) Component()            {}
func (b *bdi) isHtml()               {}
func (b *bdi) element() *HtmlElement { return b.el }
func (b *bdi) Render() string        { return b.el.Render() }

func Bdi(child *HtmlElement) *bdi {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "bdi",
		},
	}

	el.Children(child)
	b := &bdi{el: el}
	return b
}
