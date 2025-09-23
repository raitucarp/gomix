package element

import "golang.org/x/net/html"

type header struct {
	el *HtmlElement
}

func (f *header) Element() *HtmlElement { return f.el }
func (f *header) Render() string        { return f.el.Render() }

func Header(children ...IsElement) *header {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "header",
		},
	}

	el.Children(children...)

	h := &header{el: el}

	return h
}
