package element

import "golang.org/x/net/html"

type blockquote struct {
	el *HtmlElement
}

func (b *blockquote) Cite(url string) *blockquote {
	b.el.AddAttribute("cite", url)
	return b
}

func (b *blockquote) Children(children ...IsElement) *blockquote {
	b.el.Children(children...)
	return b
}
func (b *blockquote) Component()            {}
func (b *blockquote) isHtml()               {}
func (b *blockquote) Element() *HtmlElement { return b.el }
func (b *blockquote) Render() string        { return b.el.Render() }

func Blockquote(child ...IsElement) *blockquote {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "blockquote",
		},
	}

	el.Children(child...)
	b := &blockquote{el: el}
	return b
}
