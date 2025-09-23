package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type object struct {
	el *HtmlElement
}

func (o *object) Data(url string) *object {
	o.el.AddAttribute("data", url)
	return o
}

func (o *object) Form(formId string) *object {
	o.el.AddAttribute("form", formId)
	return o
}

func (o *object) Height(pixels int) *object {
	o.el.AddAttribute("height", strconv.Itoa(pixels))
	return o
}

func (o *object) Name(name string) *object {
	o.el.AddAttribute("name", name)
	return o
}

func (o *object) Type(mediaType string) *object {
	o.el.AddAttribute("type", mediaType)
	return o
}

func (o *object) UseMap(url string) *object {
	o.el.AddAttribute("usemap", url)
	return o
}

func (o *object) Width(pixels int) *object {
	o.el.AddAttribute("width", strconv.Itoa(pixels))
	return o
}

func (o *object) Element() *HtmlElement { return o.el }
func (o *object) Render() string        { return o.el.Render() }

func Object(children IsElement) *object {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "object",
		},
	}

	el.Children(children)

	o := &object{el: el}

	return o
}
