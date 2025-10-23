package disclosure

import (
	"encoding/json"
	"fmt"

	"github.com/raitucarp/gomix/addons/alpinejs"
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/components/layout"
	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"

	// "github.com/raitucarp/gomix/addons/alpinejs"
	fa5 "github.com/raitucarp/gomix/icons/fontawesome-5"
)

type accordionData struct {
	OpenItems   []int `json:"openItems"`
	Collapsible bool  `json:"collapsible"`
	Multiple    bool  `json:"multiple"`
}

type accordion struct {
	component *components.Comp
	data      accordionData
}

func Accordion(items ...components.IsComponent) *accordion {
	a := &accordion{
		component: &components.Comp{
			El: element.Div().Element(),
		},
		data: accordionData{OpenItems: []int{}, Collapsible: true, Multiple: false},
	}

	params, _ := json.Marshal(a.data)
	a.component.El = alpinejs.Alpine(a.component.El).Data(fmt.Sprintf("accordion(%s)", params)).Element()

	for index, item := range items {
		a.Component().Children(
			alpinejs.Alpine(item.Element()).Data(map[string]int{"accordion_index": index}),
		)
	}

	return a
}

func (a *accordion) Element() *element.HtmlElement {
	return a.component.El
}

func (a *accordion) Component() *components.Comp {
	return a.component
}

func (a *accordion) Collapsible() *components.Comp {
	a.data.Collapsible = true
	return a.component
}

func (a *accordion) Multiple() *components.Comp {
	a.data.Multiple = true
	return a.component
}

type accordionItem struct {
	component     *components.Comp
	collapsedIcon icons.IsIcon
	expandedIcon  icons.IsIcon

	header  components.IsComponent
	content components.IsComponent
}

func AccordionItem(header components.IsComponent, content components.IsComponent) *accordionItem {
	a := &accordionItem{
		header:        header,
		content:       content,
		collapsedIcon: fa5.FaSolidChevronDown(),
		expandedIcon:  fa5.FaSolidChevronUp(),
	}
	a.initContent()

	return a
}

func (a *accordionItem) initContent() *accordionItem {
	a.component = &components.Comp{
		El: element.Div().Element(),
	}

	collapsedActionComponent := alpinejs.Alpine(
		components.Icon(a.collapsedIcon).
			Xl().
			Element().
			Aria("controls", "accordion_index"),
	).Cloak().
		Show("!isOpen(accordion_index)")

	expandedActionComponent := alpinejs.Alpine(
		components.Icon(a.expandedIcon).
			Xl().
			Element().
			Aria("controls", "accordion_index"),
	).Cloak().
		Show("isOpen(accordion_index)")

	headerComponent := alpinejs.Alpine(
		layout.
			HStack(
				a.header,
				collapsedActionComponent,
				expandedActionComponent,
			).
			Component().
			JustifyBetween().
			ItemsCenter().
			Element(),
	).
		On("click", "toggle(accordion_index)")

	a.Component().Children(
		headerComponent,
		alpinejs.Alpine(a.content).
			Show("isOpen(accordion_index)").
			Collapse().
			Attribute(),
	)
	return a
}

func (a *accordionItem) Element() *element.HtmlElement {

	return a.component.El
}

func (a *accordionItem) Component() *components.Comp {

	return a.component
}

func (a *accordionItem) CollapseIcon(icon icons.IsIcon) *components.Comp {
	a.collapsedIcon = icon
	a.initContent()
	return a.component
}

func (a *accordionItem) ExpandIcon(icon icons.IsIcon) *components.Comp {
	a.expandedIcon = icon
	a.initContent()
	return a.component
}
