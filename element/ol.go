package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type ol struct {
	el *HtmlElement
}

func (o *ol) Reversed() *ol {
	o.el.AddAttributeBool("reversed")
	return o
}

func (o *ol) Start(number int) *ol {
	o.el.AddAttribute("start", strconv.Itoa(number))
	return o
}

func (o *ol) Element() *HtmlElement { return o.el }
func (o *ol) Render() string        { return o.el.Render() }

func Ol(children ...IsElement) *ol {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "ol",
		},
	}

	el.Children(children...)

	o := &ol{el: el}

	return o
}
