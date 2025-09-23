package element

import (
	"strings"

	"golang.org/x/net/html"
)

type source struct {
	el *HtmlElement
}

func (s *source) Media(media string) *source {
	s.el.AddAttribute("media", media)
	return s
}

func (s *source) Sizes(size ...string) *source {
	s.el.AddAttribute("size", strings.Join(size, ", "))
	return s
}

func (s *source) Src(url string) *source {
	s.el.AddAttribute("src", url)
	return s
}

func (s *source) SrcSet(sources ...string) *source {
	s.el.AddAttribute("srcset", strings.Join(sources, ","))
	return s
}

func (s *source) Type(mimeType string) *source {
	s.el.AddAttribute("type", mimeType)
	return s
}

func (s *source) Element() *HtmlElement { return s.el }
func (s *source) Render() string        { return s.el.Render() }

func Source() *source {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "source",
		},
	}

	s := &source{el: el}

	return s
}
