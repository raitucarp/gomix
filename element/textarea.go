package element

import (
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type textarea struct {
	el *HtmlElement
}

func (t *textarea) AutoFocus() *textarea {
	t.el.AddAttributeBool("autofocus")
	return t
}

func (t *textarea) Cols(cols int) *textarea {
	t.el.AddAttribute("cols", strconv.Itoa(cols))
	return t
}

func (t *textarea) Dirname(name string) *textarea {
	t.el.AddAttribute("dirname", strings.Join([]string{name, "dir"}, "."))
	return t
}

func (t *textarea) Disabled() *textarea {
	t.el.AddAttributeBool("disabled")
	return t
}

func (t *textarea) Form(form string) *textarea {
	t.el.AddAttribute("form", form)
	return t
}

func (t *textarea) MaxLength(maxlength int) *textarea {
	t.el.AddAttribute("maxlength", strconv.Itoa(maxlength))
	return t
}

func (t *textarea) Name(name string) *textarea {
	t.el.AddAttribute("name", name)
	return t
}

func (t *textarea) Placeholder(placeholder string) *textarea {
	t.el.AddAttribute("placeholder", placeholder)
	return t
}

func (t *textarea) ReadOnly() *textarea {
	t.el.AddAttributeBool("readonly")
	return t
}

func (t *textarea) Required() *textarea {
	t.el.AddAttributeBool("required")
	return t
}

func (t *textarea) Rows(rows int) *textarea {
	t.el.AddAttribute("rows", strconv.Itoa(rows))
	return t
}

type TextAreaWrap string

const (
	TextAreaWrapSoft TextAreaWrap = "soft"
	TextAreaWrapHard TextAreaWrap = "hard"
)

func (t *textarea) Wrap(wrap TextAreaWrap) *textarea {
	t.el.AddAttribute("wrap", string(wrap))
	return t
}

func (t *textarea) Element() *HtmlElement { return t.el }
func (t *textarea) Render() string        { return t.el.Render() }

func TextArea(children ...IsElement) *textarea {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "textarea",
		},
	}

	el.Children(children...)
	t := &textarea{el: el}

	return t
}
