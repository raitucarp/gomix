package element

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type svg struct {
	el *HtmlElement
}

func (s *svg) ViewBox(minX, minY, width, height int) *svg {
	s.el.AddAttribute("viewBox", fmt.Sprintf("%d %d %d %d", minX, minY, width, height))
	return s
}

func (s *svg) Component()            {}
func (s *svg) Element() *HtmlElement { return s.el }
func (s *svg) Render() string        { return s.el.Render() }

func SvgString(svgString string) (s *svg, err error) {
	el, err := html.ParseFragment(strings.NewReader(svgString), &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Body,
		Data:     "body",
	})

	if err != nil {
		return
	}

	s = &svg{el: &HtmlElement{node: el[0]}}
	return s, nil
}

func Svg() *svg {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "svg",
		},
	}

	s := &svg{el: el}
	return s
}
