package element

import (
	"golang.org/x/net/html"
)

type wbr struct {
	el *HtmlElement
}

func (u *wbr) Element() *HtmlElement { return u.el }
func (u *wbr) Render() string        { return u.el.Render() }

func Wbr(children ...IsElement) *wbr {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "wbr",
		},
	}

	el.Children(children...)
	u := &wbr{el: el}

	return u
}
