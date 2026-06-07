package disclosure

import (
	"fmt"

	"github.com/raitucarp/gomix/addons/alpinejs"
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
)

type tabsData struct {
	ActiveTab int `json:"activeTab"`
}

type tabs struct {
	component *components.Comp
	data      tabsData
}

func Tabs(children ...components.IsComponent) *tabs {
	t := &tabs{
		component: &components.Comp{
			El: element.Div().Element(),
		},
		data: tabsData{ActiveTab: 0},
	}

	t.component.El.AddAttribute("data-gomix-component", "Tabs")
	t.component.El.AddAttribute("class", "chakra-tabs")

	t.component.El = alpinejs.Alpine(t.component.El).Data(fmt.Sprintf("{ activeTab: %d }", t.data.ActiveTab)).Element()

	t.component.Children(children...)
	return t
}

func (t *tabs) Element() *element.HtmlElement {
	return t.component.El
}

func (t *tabs) Component() *components.Comp {
	return t.component
}

func (t *tabs) DefaultIndex(index int) *tabs {
	t.data.ActiveTab = index
	t.component.El = alpinejs.Alpine(t.component.El).Data(fmt.Sprintf("{ activeTab: %d }", t.data.ActiveTab)).Element()
	return t
}

type tabList struct {
	component *components.Comp
}

func TabList(children ...components.IsComponent) *tabList {
	tl := &tabList{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	tl.component.El.AddAttribute("data-gomix-component", "TabList")
	tl.component.El.AddAttribute("class", "chakra-tab-list")
	tl.component.El.AddAttribute("role", "tablist")

	for index, child := range children {
		tl.component.Children(
			alpinejs.Alpine(child.Element()).
				On("click", fmt.Sprintf("activeTab = %d", index)).
				Bind("aria-selected", fmt.Sprintf("activeTab === %d", index)).
				Bind("data-selected", fmt.Sprintf("activeTab === %d ? '' : null", index)).
				Element(),
		)
	}

	return tl
}

func (tl *tabList) Element() *element.HtmlElement {
	return tl.component.El
}

func (tl *tabList) Component() *components.Comp {
	return tl.component
}

type tab struct {
	component *components.Comp
}

func Tab(children ...components.IsComponent) *tab {
	t := &tab{
		component: &components.Comp{
			El: element.Button().Element(),
		},
	}

	t.component.El.AddAttribute("data-gomix-component", "Tab")
	t.component.El.AddAttribute("class", "chakra-tab")
	t.component.El.AddAttribute("role", "tab")
	t.component.El.AddAttribute("type", "button")

	t.component.Children(children...)
	return t
}

func (t *tab) Element() *element.HtmlElement {
	return t.component.El
}

func (t *tab) Component() *components.Comp {
	return t.component
}

type tabPanels struct {
	component *components.Comp
}

func TabPanels(children ...components.IsComponent) *tabPanels {
	tp := &tabPanels{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	tp.component.El.AddAttribute("data-gomix-component", "TabPanels")
	tp.component.El.AddAttribute("class", "chakra-tab-panels")

	for index, child := range children {
		tp.component.Children(
			alpinejs.Alpine(child.Element()).
				Show(fmt.Sprintf("activeTab === %d", index)).
				Element(),
		)
	}

	return tp
}

func (tp *tabPanels) Element() *element.HtmlElement {
	return tp.component.El
}

func (tp *tabPanels) Component() *components.Comp {
	return tp.component
}

type tabPanel struct {
	component *components.Comp
}

func TabPanel(children ...components.IsComponent) *tabPanel {
	tp := &tabPanel{
		component: &components.Comp{
			El: element.Div().Element(),
		},
	}

	tp.component.El.AddAttribute("data-gomix-component", "TabPanel")
	tp.component.El.AddAttribute("class", "chakra-tab-panel")
	tp.component.El.AddAttribute("role", "tabpanel")

	tp.component.Children(children...)
	return tp
}

func (tp *tabPanel) Element() *element.HtmlElement {
	return tp.component.El
}

func (tp *tabPanel) Component() *components.Comp {
	return tp.component
}
