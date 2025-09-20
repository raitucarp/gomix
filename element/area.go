package element

import (
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type area struct {
	el *HtmlElement
}

func (a *area) Alt(text string) *area {
	a.el.AddAttribute("alt", text)
	return a
}

func (a *area) Coords(coordinates ...float32) *area {
	coords := []string{}
	for _, coord := range coordinates {
		coords = append(coords, strconv.Itoa(int(coord)))
	}
	a.el.AddAttribute("coords", strings.Join(coords, ","))
	return a
}

func (a *area) Download(filename string) *area {
	a.el.AddAttribute("download", filename)
	return a
}

func (a *area) Href(url string) *area {
	a.el.AddAttribute("href", url)
	return a
}

func (a *area) HrefLang(language_code string) *area {
	a.el.AddAttribute("hreflang", language_code)
	return a
}

func (a *area) Media(media_query string) *area {
	a.el.AddAttribute("media", media_query)
	return a
}

func (a *area) ReferrerPolicy(policy RefererPolicy) *area {
	a.el.AddAttribute("referrerpolicy", string(policy))
	return a
}

func (a *area) Rel(rel AnchorRel) *area {
	a.el.AddAttribute("rel", string(rel))
	return a
}

type AreaShape string

const (
	AreaShapeDefault AreaShape = "default"
	AreaShapeRect    AreaShape = "rect"
	AreaShapeCircle  AreaShape = "circle"
	AreaShapePoly    AreaShape = "poly"
)

func (a *area) Shape(shape AreaShape) *area {
	a.el.AddAttribute("shape", string(shape))
	return a
}

func (a *area) Target(target AnchorTarget) *area {
	a.el.AddAttribute("target", string(target))
	return a
}

func (a *area) Type(media_type string) *area {
	a.el.AddAttribute("type", media_type)
	return a
}

func (a *area) Component()            {}
func (a *area) isHtml()               {}
func (a *area) Element() *HtmlElement { return a.el }
func (a *area) Render() string        { return a.el.Render() }

func Area() *area {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "area",
		},
	}

	ar := &area{el: el}
	return ar
}
