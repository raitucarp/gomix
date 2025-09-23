package element

import "golang.org/x/net/html"

type i struct {
	el *HtmlElement
}

func (i *i) Element() *HtmlElement { return i.el }
func (i *i) Render() string        { return i.el.Render() }

func I(children IsElement) *i {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "i",
		},
	}

	el.Children(children)

	it := &i{el: el}

	return it
}
