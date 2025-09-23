package element

import (
	"golang.org/x/net/html"
)

type kbd struct {
	el *HtmlElement
}

func (k *kbd) Element() *HtmlElement { return k.el }
func (k *kbd) Render() string        { return k.el.Render() }

func Kbd(children IsElement) *kbd {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "kbd",
		},
	}

	el.Children(children)

	k := &kbd{el: el}

	return k
}
