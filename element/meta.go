package element

import "golang.org/x/net/html"

type meta struct {
	el *HtmlElement
}

func (m *meta) InsideHead()           {}
func (m *meta) Element() *HtmlElement { return m.el }
func (m *meta) Render() string        { return m.el.Render() }

func Meta(attribs map[string]string) *meta {
	attrs := []html.Attribute{}
	for k, v := range attribs {
		attrs = append(attrs, html.Attribute{Key: k, Val: v})
	}

	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "meta",
			Attr: attrs,
		},
	}

	m := &meta{el: el}

	return m
}
