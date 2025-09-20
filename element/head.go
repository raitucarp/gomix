package element

import "golang.org/x/net/html"

type head struct {
	el *HtmlElement
}

func (h *head) Element() *HtmlElement { return h.el }
func (h *head) Render() string        { return h.el.Render() }

type IsHeadElement interface {
	InsideHead()
	Element() *HtmlElement
}

func Head(children ...IsHeadElement) *head {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "head",
		},
	}

	for _, c := range children {
		el.Children(Element(c))
	}

	h := &head{el: el}

	return h
}
