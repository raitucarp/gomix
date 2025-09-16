package components

import (
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type area struct {
	comp *Component
}

func (a *area) Alt(text string) *area {
	a.comp.addAttribute("alt", text)
	return a
}

func (a *area) Coords(coordinates ...float32) *area {
	coords := []string{}
	for _, coord := range coordinates {
		coords = append(coords, strconv.Itoa(int(coord)))
	}
	a.comp.addAttribute("coords", strings.Join(coords, ","))
	return a
}

func (a *area) Download(filename string) *area {
	a.comp.addAttribute("download", filename)
	return a
}

func (a *area) Href(url string) *area {
	a.comp.addAttribute("href", url)
	return a
}

func (a *area) HrefLang(language_code string) *area {
	a.comp.addAttribute("hreflang", language_code)
	return a
}

func (a *area) Media(media_query string) *area {
	a.comp.addAttribute("media", media_query)
	return a
}

func (a *area) ReferrerPolicy(policy RefererPolicy) *area {
	a.comp.addAttribute("referrerpolicy", string(policy))
	return a
}

func (a *area) Rel(rel AnchorRel) *area {
	a.comp.addAttribute("rel", string(rel))
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
	a.comp.addAttribute("shape", string(shape))
	return a
}

func (a *area) Target(target AnchorTarget) *area {
	a.comp.addAttribute("target", string(target))
	return a
}

func (a *area) Type(media_type string) *area {
	a.comp.addAttribute("type", media_type)
	return a
}

func (a *area) Component() *Component {
	return a.comp
}

func Area(child *Component) *area {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "area",
		},
	}

	comp.Children(child)
	ar := &area{comp: comp}
	return ar
}
