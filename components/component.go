package components

import (
	"bytes"

	"github.com/raitucarp/gomix/components/styles"
	"github.com/raitucarp/gomix/element"
	"golang.org/x/net/html"
)

type IsComponent interface {
	Element() *element.HtmlElement
	// Render() string
}

type component struct {
	el     *element.HtmlElement
	styles *styles.Styles

	// stylesheets []string
	// scripts     []string
}

func Component(c IsComponent) IsComponent {
	comp := &component{
		el: c.Element(),
	}

	return *comp
}

func cloneNode(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}
	clone := &html.Node{
		Type:      n.Type,
		DataAtom:  n.DataAtom,
		Data:      n.Data,
		Namespace: n.Namespace,
		Attr:      append([]html.Attribute{}, n.Attr...),
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childClone := cloneNode(c)
		clone.AppendChild(childClone)
	}
	return clone
}

func ApplyLayout(layout IsComponent, e IsComponent) IsComponent {

	buffer := bytes.Buffer{}
	html.Render(&buffer, cloneNode(layout.Element().GetNode()))

	node, _ := html.Parse(&buffer)
	r := cloneNode(e.Element().GetNode())

	for desc := range node.Descendants() {
		if desc.Type == html.ElementNode && desc.Data == "slot" {
			desc.Parent.InsertBefore(r, desc)
			desc.Parent.RemoveChild(desc)

		}
	}

	n := element.CreateElementByNode(node)

	return n
}

func (e *component) Children(components ...*component) *component {
	for _, comp := range components {
		e.el.Children(comp.el)
		// e.node.AppendChild(comp.node)
	}
	return e
}

func (e component) Element() *element.HtmlElement {
	return e.el
}

func Render(c IsComponent) string {
	buffer := bytes.Buffer{}
	html.Render(&buffer, c.Element().GetNode())

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
