package element

import "golang.org/x/net/html"

type TargetAction string

const (
	TargetActionHide   TargetAction = "hide"
	TargetActionShow   TargetAction = "show"
	TargetActionToggle TargetAction = "toggle"
)

type ButtonType string

const (
	ButtonTypeButton ButtonType = "button"
	ButtonTypeReset  ButtonType = "reset"
	ButtonTypeSubmit ButtonType = "submit"
)

type button struct {
	el *HtmlElement
}

func (b *button) Autofocus() *button {
	b.el.AddAttributeBool("autofocus")
	return b
}

func (b *button) Disabled() *button {
	b.el.AddAttributeBool("disabled")
	return b
}

func (b *button) Form(formId string) *button {
	b.el.AddAttribute("form", formId)
	return b
}

func (b *button) FormEncodingType(formenc string) *button {
	b.el.AddAttribute("formenctype", formenc)
	return b
}

func (b *button) FormMethod(method string) *button {
	b.el.AddAttribute("formmethod", method)
	return b
}

func (b *button) FormNoValidate() *button {
	b.el.AddAttributeBool("formnovalidate")
	return b
}

func (b *button) FormTarget(target AnchorTarget) *button {
	b.el.AddAttribute("formtarget", string(target))
	return b
}

func (b *button) PopoverTarget(target string) *button {
	b.el.AddAttribute("popovertarget", target)
	return b
}

func (b *button) PopoverTargetAction(target TargetAction) *button {
	b.el.AddAttribute("popovertargetaction", string(target))
	return b
}

func (b *button) Name(name string) *button {
	b.el.AddAttribute("name", name)
	return b
}

func (b *button) Type(buttonType ButtonType) *button {
	b.el.AddAttribute("type", string(buttonType))
	return b
}

func (b *button) Value(value string) *button {
	b.el.AddAttribute("value", value)
	return b
}

func (b *button) Component() {}

func (b *button) Element() *HtmlElement { return b.el }
func (b *button) Render() string        { return b.el.Render() }

func Button(children ...IsElement) *button {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "button",
		},
	}

	el.Children(children...)

	b := &button{el: el}
	return b
}
