package element

import (
	"golang.org/x/net/html"
)

type rt struct {
	el *HtmlElement
}

func (r *rt) Element() *HtmlElement { return r.el }
func (r *rt) Render() string        { return r.el.Render() }

func Rt(children IsElement) *rt {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "rt",
		},
	}

	el.Children(children)
	p := &rt{el: el}

	return p
}
