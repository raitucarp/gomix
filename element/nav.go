package element

import (
	"golang.org/x/net/html"
)

type nav struct {
	el *HtmlElement
}

func (n *nav) Element() *HtmlElement { return n.el }
func (n *nav) Render() string        { return n.el.Render() }

func Nav(children ...IsElement) *nav {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "nav",
		},
	}

	el.Children(children...)

	n := &nav{el: el}

	return n
}
