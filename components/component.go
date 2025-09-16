package components

import (
	"bytes"
	"slices"
	"strings"

	"github.com/raitucarp/gomix/components/styles"
	"golang.org/x/net/html"
)

type AsComponentInter interface {
	Component() *Component
}

func AsComponent(c AsComponentInter) *Component {
	return c.Component()
}

type Component struct {
	node   *html.Node
	styles *styles.Styles

	stylesheets []string
	scripts     []string
}

type Components []*Component

func (e *Component) addAttribute(key string, value string) {
	e.node.Attr = append(e.node.Attr, html.Attribute{Namespace: "", Key: key, Val: value})
	e.node.Attr = slices.CompactFunc(e.node.Attr, func(attr1 html.Attribute, attr2 html.Attribute) bool {
		return attr1.Key == attr2.Key
	})
}

func (e *Component) addAttributeBool(key string) {
	e.node.Attr = append(e.node.Attr, html.Attribute{Namespace: "", Key: key})
	e.node.Attr = slices.CompactFunc(e.node.Attr, func(attr1 html.Attribute, attr2 html.Attribute) bool {
		return attr1.Key == attr2.Key
	})
}

func (e *Component) Id(id string) *Component {
	e.addAttribute("id", id)
	return e
}

func (e *Component) Class(classnames ...string) *Component {
	e.addAttribute("class", strings.Join(classnames, " "))
	return e
}

func (c *Component) Title(title string) *Component {
	c.addAttribute("title", title)
	return c
}

func (c *Component) Lang(lang string) *Component {
	c.addAttribute("lang", lang)
	return c
}

func (c *Component) Aria(name string, value string) *Component {
	c.addAttribute("aria-"+name, value)
	return c
}

func (c *Component) Role(role string) *Component {
	c.addAttribute("role", role)
	return c
}

func (e *Component) Styles(s *styles.Styles) *Component {
	e.styles = s
	return e
}

func (e *Component) GetNode() *html.Node {
	return e.node
}

func ApplyLayout(layout *Component, e *Component) *Component {
	buffer := bytes.Buffer{}
	html.Render(&buffer, layout.node)

	node, _ := html.Parse(&buffer)
	r := *e.node

	for desc := range node.Descendants() {
		if desc.Type == html.ElementNode && desc.Data == "slot" {
			desc.Parent.InsertBefore(&r, desc)
			desc.Parent.RemoveChild(desc)

		}
	}

	n := &Component{node: node, styles: e.styles, stylesheets: e.stylesheets, scripts: e.scripts}

	return n
}

func (e *Component) Children(components ...*Component) *Component {
	for _, comp := range components {
		e.node.AppendChild(comp.node)
	}
	return e
}

func Slot() *Component {
	slot := &Component{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "slot",
		},
	}

	return slot
}

func Render(component *Component) string {
	buffer := bytes.Buffer{}
	html.Render(&buffer, component.node)

	node, _ := html.Parse(&buffer)

	for desc := range node.Descendants() {
		if desc.Type == html.ElementNode && desc.Data == "slot" {

			desc.Parent.RemoveChild(desc)

		}
	}

	buffer2 := bytes.Buffer{}
	html.Render(&buffer2, node)

	return buffer2.String()
}
