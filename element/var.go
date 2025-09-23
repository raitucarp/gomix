package element

import (
	"golang.org/x/net/html"
)

type varEl struct {
	el *HtmlElement
}

func (v *varEl) Element() *HtmlElement { return v.el }
func (v *varEl) Render() string        { return v.el.Render() }

func Var(children ...IsElement) *varEl {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "var",
		},
	}

	el.Children(children...)
	v := &varEl{el: el}

	return v
}
