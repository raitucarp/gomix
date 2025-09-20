package element

import "golang.org/x/net/html"

type aside struct {
	el *HtmlElement
}

func (a *aside) Children(children ...IsElement) *aside {
	a.el.Children(children...)
	return a
}
func (a *aside) Component()            {}
func (a *aside) isHtml()               {}
func (a *aside) element() *HtmlElement { return a.el }
func (a *aside) Render() string        { return a.el.Render() }

func Aside(child ...IsElement) *aside {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "aside",
		},
	}

	el.Children(child...)
	as := &aside{el: el}
	return as
}
