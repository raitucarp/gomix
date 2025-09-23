package element

import (
	"golang.org/x/net/html"
)

type template struct {
	el *HtmlElement
}

func (t *template) Element() *HtmlElement { return t.el }
func (t *template) Render() string        { return t.el.Render() }

func Template(children ...IsElement) *template {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "template",
		},
	}

	el.Children(children...)
	t := &template{el: el}

	return t
}
