package element

import "golang.org/x/net/html"

type title struct {
	el *HtmlElement
}

func (t *title) InsideHead()           {}
func (t *title) Element() *HtmlElement { return t.el }
func (t *title) Render() string        { return t.el.Render() }

func Title(text string) *title {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "title",
			FirstChild: &html.Node{
				Type: html.TextNode,
				Data: text,
			},
		},
	}

	t := &title{el: el}

	return t
}
