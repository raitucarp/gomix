package element

import "golang.org/x/net/html"

type style struct {
	el *HtmlElement
}

func (s *style) Media(media_query string) *style {
	s.el.AddAttribute("media", media_query)
	return s
}

func (s *style) Type(styleType string) *style {
	s.el.AddAttribute("type", styleType)
	return s
}

func (s *style) InsideHead()           {}
func (s *style) Element() *HtmlElement { return s.el }
func (s *style) Render() string        { return s.el.Render() }

func Style(text string) *style {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "style",
			FirstChild: &html.Node{
				Type: html.TextNode,
				Data: text,
			},
		},
	}

	t := &style{el: el}

	return t
}
