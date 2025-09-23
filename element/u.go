package element

import (
	"golang.org/x/net/html"
)

type uEl struct {
	el *HtmlElement
}

func (u *uEl) Element() *HtmlElement { return u.el }
func (u *uEl) Render() string        { return u.el.Render() }

func U(children ...IsElement) *uEl {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "u",
		},
	}

	el.Children(children...)
	u := &uEl{el: el}

	return u
}
