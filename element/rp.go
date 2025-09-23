package element

import (
	"golang.org/x/net/html"
)

type rp struct {
	el *HtmlElement
}

func (r *rp) Element() *HtmlElement { return r.el }
func (r *rp) Render() string        { return r.el.Render() }

func Rp(children IsElement) *rp {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "rp",
		},
	}

	el.Children(children)
	r := &rp{el: el}

	return r
}
