package element

import "golang.org/x/net/html"

type fragmentElement struct {
	el *HtmlElement
}

func (a *fragmentElement) Children(children ...IsElement) *fragmentElement {
	a.el.Children(children...)
	return a
}
func (a *fragmentElement) Component()            {}
func (a *fragmentElement) isHtml()               {}
func (a *fragmentElement) Element() *HtmlElement { return a.el }
func (a *fragmentElement) Render() string        { return a.el.Render() }

func Fragment(children ...IsElement) *fragmentElement {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "fragment",
		},
	}

	el.Children(children...)
	art := &fragmentElement{el: el}
	return art
}
