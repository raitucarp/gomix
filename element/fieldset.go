package element

import "golang.org/x/net/html"

type fieldset struct {
	el *HtmlElement
}

func (f *fieldset) Disabled() *fieldset {
	f.el.AddAttributeBool("disabled")
	return f
}

func (f *fieldset) Form(id string) *fieldset {
	f.el.AddAttribute("form", id)
	return f
}

func (f *fieldset) Name(text string) *fieldset {
	f.el.AddAttribute("name", text)
	return f
}

func (f *fieldset) Element() *HtmlElement { return f.el }
func (f *fieldset) Render() string        { return f.el.Render() }

func FieldSet(children ...IsElement) *fieldset {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "fieldset",
		},
	}

	el.Children(children...)

	f := &fieldset{el: el}

	return f
}
