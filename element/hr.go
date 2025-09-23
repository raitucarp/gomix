package element

import "golang.org/x/net/html"

type hr struct {
	el *HtmlElement
}

func (f *hr) Element() *HtmlElement { return f.el }
func (f *hr) Render() string        { return f.el.Render() }

func Hr() *hr {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "hr",
		},
	}

	h := &hr{el: el}

	return h
}
