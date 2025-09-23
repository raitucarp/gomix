package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type li struct {
	el *HtmlElement
}

func (l *li) Value(value int) *li {
	l.el.AddAttribute("value", strconv.Itoa(value))
	return l
}
func (l *li) Element() *HtmlElement { return l.el }
func (l *li) Render() string        { return l.el.Render() }

func Li(children ...IsElement) *li {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "li",
		},
	}

	el.Children(children...)

	l := &li{el: el}

	return l
}
