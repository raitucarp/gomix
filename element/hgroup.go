package element

import "golang.org/x/net/html"

type hgroup struct {
	el *HtmlElement
}

func (f *hgroup) Element() *HtmlElement { return f.el }
func (f *hgroup) Render() string        { return f.el.Render() }

func HGroup(children ...IsElement) *hgroup {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "hgroup",
		},
	}

	el.Children(children...)

	h := &hgroup{el: el}

	return h
}
