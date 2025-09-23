package element

import (
	"golang.org/x/net/html"
)

type tfoot struct {
	el *HtmlElement
}

func (t *tfoot) Element() *HtmlElement { return t.el }
func (t *tfoot) Render() string        { return t.el.Render() }

func TFoot(children ...IsElement) *tfoot {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "tfoot",
		},
	}

	el.Children(children...)
	t := &tfoot{el: el}

	return t
}
