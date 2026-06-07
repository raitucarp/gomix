package overlays

import (
	"fmt"

	"github.com/raitucarp/gomix/addons/alpinejs"
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type modal struct {
	component *components.Comp
}

func Modal(children ...components.IsComponent) *modal {
	m := &modal{
		component: &components.Comp{
			El: element.Dialog(element.Text("")).Element(),
		},
	}

	m.component.El.AddAttribute("data-gomix-component", "Modal")
	m.component.El.AddAttribute("class", "chakra-modal")

	m.component.El = alpinejs.Alpine(m.component.El).
		Data("{ isOpen: false }").
		Bind("open", "isOpen").
		On("@keydown.escape.window", "isOpen = false").
		Element()

	m.component.Children(children...)
	return m
}

func (m *modal) Element() *element.HtmlElement {
	return m.component.El
}

func (m *modal) Component() *components.Comp {
	return m.component
}

func (m *modal) IsOpen(isOpen bool) *modal {
	m.component.El = alpinejs.Alpine(m.component.El).
		Data(fmt.Sprintf("{ isOpen: %t }", isOpen)).
		Element()
	return m
}
