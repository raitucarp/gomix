package element

import "golang.org/x/net/html"

type address struct {
	el *HtmlElement
}

func (a *address) Children(children ...IsElement) *address {
	a.el.Children(children...)
	return a
}

func (a *address) Component()            {}
func (a *address) isHtml()               {}
func (a *address) Element() *HtmlElement { return a.el }
func (a *address) Render() string        { return a.el.Render() }

func Address(child ...IsElement) *address {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "address",
		},
	}

	el.Children(child...)
	addr := &address{el: el}
	return addr
}
