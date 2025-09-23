package element

import (
	"golang.org/x/net/html"
)

type ruby struct {
	el *HtmlElement
}

func (r *ruby) Element() *HtmlElement { return r.el }
func (r *ruby) Render() string        { return r.el.Render() }

func Ruby(children ...IsElement) *ruby {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "ruby",
		},
	}

	el.Children(children...)
	r := &ruby{el: el}

	return r
}
