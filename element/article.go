package element

import "golang.org/x/net/html"

type article struct {
	el *HtmlElement
}

func (a *article) Children(children ...IsElement) *article {
	a.el.Children(children...)
	return a
}
func (a *article) Component()            {}
func (a *article) isHtml()               {}
func (a *article) element() *HtmlElement { return a.el }
func (a *article) Render() string        { return a.el.Render() }

func Article(children ...IsElement) *article {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "article",
		},
	}

	el.Children(children...)
	art := &article{el: el}
	return art
}
