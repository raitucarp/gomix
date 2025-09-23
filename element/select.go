package element

import (
	"strconv"

	"golang.org/x/net/html"
)

type selectEl struct {
	el *HtmlElement
}

func (s *selectEl) AutoFocus() *selectEl {
	s.el.AddAttributeBool("autofocus")
	return s
}

func (s *selectEl) Disabled() *selectEl {
	s.el.AddAttributeBool("disabled")
	return s
}

func (s *selectEl) Form(formId string) *selectEl {
	s.el.AddAttribute("form", formId)
	return s
}

func (s *selectEl) Multiple() *selectEl {
	s.el.AddAttributeBool("multiple")
	return s
}

func (s *selectEl) Name(name string) *selectEl {
	s.el.AddAttribute("name", name)
	return s
}

func (s *selectEl) Required() *selectEl {
	s.el.AddAttributeBool("required")
	return s
}

func (s *selectEl) Size(size int) *selectEl {
	s.el.AddAttribute("size", strconv.Itoa(size))
	return s
}

func (s *selectEl) Element() *HtmlElement { return s.el }
func (s *selectEl) Render() string        { return s.el.Render() }

func Select(children ...IsElement) *selectEl {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "select",
		},
	}

	el.Children(children...)
	s := &selectEl{el: el}

	return s
}
