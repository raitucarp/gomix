package element

import (
	"golang.org/x/net/html"
)

type ul struct {
	el *HtmlElement
}

func (u *ul) Element() *HtmlElement { return u.el }
func (u *ul) Render() string        { return u.el.Render() }

func Ul(children ...IsElement) *ul {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "ul",
		},
	}

	el.Children(children...)
	u := &ul{el: el}

	return u
}
