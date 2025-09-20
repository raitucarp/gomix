package element

import "golang.org/x/net/html"

type htmlTag struct {
	el *HtmlElement
}

func (h *htmlTag) Element() *HtmlElement { return h.el }
func (h *htmlTag) Render() string        { return h.el.Render() }
func (h *htmlTag) Lang(lang LanguageCode) *htmlTag {
	return h
}

func Html(children ...IsElement) *htmlTag {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "html",
		},
	}

	el.Children(children...)
	h := &htmlTag{el: el}

	return h
}
