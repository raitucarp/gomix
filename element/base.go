package element

import "golang.org/x/net/html"

type base struct {
	el *HtmlElement
}

func (b *base) Href(url string) *HtmlElement {
	b.el.AddAttribute("href", url)
	return b.el
}

func (b *base) Target(target AnchorTarget) *HtmlElement {
	b.el.AddAttribute("target", string(target))
	return b.el
}

func (b *base) Component()            {}
func (b *base) isHtml()               {}
func (b *base) element() *HtmlElement { return b.el }
func (b *base) Render() string        { return b.el.Render() }

func Base() *base {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "base",
		},
	}

	b := &base{el: el}
	return b
}
