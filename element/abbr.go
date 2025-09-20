package element

import "golang.org/x/net/html"

type abbreviation struct {
	el *HtmlElement
}

func (a *abbreviation) Children(children ...IsElement) *abbreviation {
	a.el.Children(children...)
	return a
}

func (a *abbreviation) Component()            {}
func (a *abbreviation) isHtml()               {}
func (a *abbreviation) Element() *HtmlElement { return a.el }
func (a *abbreviation) Render() string        { return a.el.Render() }

func Abbr(child *HtmlElement) *abbreviation {
	comp := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "abbr",
		},
	}

	comp.Children(child)
	abbr := &abbreviation{el: comp}
	return abbr
}
