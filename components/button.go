package components

import "golang.org/x/net/html"

type button struct {
	comp *Component
}

func (b *button) Autofocus() *button {
	b.comp.addAttributeBool("autofocus")
	return b
}

func (b *button) Disabled() *button {
	b.comp.addAttributeBool("disabled")
	return b
}

func (b *button) Form(formId string) *button {
	b.comp.addAttribute("form", formId)
	return b
}

func (b *button) FormEncodingType(formenc string) *button {
	b.comp.addAttribute("formenctype", formenc)
	return b
}

func (b *button) FormMethod(method string) *button {
	b.comp.addAttribute("formmethod", method)
	return b
}

func (b *button) FormNoValidate() *button {
	b.comp.addAttributeBool("formnovalidate")
	return b
}

func (b *button) FormTarget(target AnchorTarget) *button {
	b.comp.addAttribute("formtarget", string(target))
	return b
}

func (b *button) PopoverTarget(target string) *button {
	b.comp.addAttribute("popovertarget", target)
	return b
}

func (b *button) PopoverTargetAction(target TargetAction) *button {
	b.comp.addAttribute("popovertargetaction", string(target))
	return b
}

func (b *button) Name(name string) *button {
	b.comp.addAttribute("name", name)
	return b
}

func (b *button) Type(buttonType ButtonType) *button {
	b.comp.addAttribute("type", string(buttonType))
	return b
}

func (b *button) Value(value string) *button {
	b.comp.addAttribute("value", value)
	return b
}

func (b *button) Component() *Component {
	return b.comp
}

func Button(children ...*Component) *button {
	comp := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "button",
		},
	}

	comp.Children(children...)

	b := &button{comp: comp}
	return b
}
