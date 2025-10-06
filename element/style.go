package element

import "golang.org/x/net/html"

type styl struct {
	el *HtmlElement
}

func (s *styl) Media(media_query string) *styl {
	s.el.AddAttribute("media", media_query)
	return s
}

func (s *styl) Type(styleType string) *styl {
	s.el.AddAttribute("type", styleType)
	return s
}

func (s *styl) InsideHead()           {}
func (s *styl) Element() *HtmlElement { return s.el }
func (s *styl) Render() string        { return s.el.Render() }

func Style(text string) *styl {
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

	t := &styl{el: el}

	return t
}
