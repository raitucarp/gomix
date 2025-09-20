package element

import "golang.org/x/net/html"

type bdo struct {
	el *HtmlElement
}

func (b *bdo) Ltr() *bdo {
	b.el.Dir(DirectionLtr)
	return b
}

func (b *bdo) Rtl() *bdo {
	b.el.Dir(DirectionRtl)
	return b
}

func (b *bdo) Component()            {}
func (b *bdo) isHtml()               {}
func (b *bdo) element() *HtmlElement { return b.el }
func (b *bdo) Render() string        { return b.el.Render() }

func Bdo() *bdo {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "bdo",
		},
	}

	b := &bdo{el: el}
	return b
}
