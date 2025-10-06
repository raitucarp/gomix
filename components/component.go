package components

import (
	"bytes"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/styles"
	"golang.org/x/net/html"
)

type IsComponent interface {
	Element() *element.HtmlElement
	// Render() string
}

type component struct {
	el *element.HtmlElement
	// styles *styles.Styles

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

var styleAttrPattern = regexp.MustCompile(`style\-`)

func ExtractCSS(c IsComponent) string {
	classMap := make(map[string]map[string]string)
	el := c.Element()

	// className := ""
	// for key, attr := range el.GetNode().Attr {

	// }

	for desc := range el.GetNode().Descendants() {
		if desc.Type == html.ElementNode {
			className := ""
			for index, attr := range desc.Attr {
				if !styleAttrPattern.MatchString(attr.Key) {
					continue
				}

				styleData := strings.Split(attr.Key, "-")

				if styleData[2] == "classname" {
					className = attr.Val
					classMap[className] = make(map[string]string)
					desc.Attr = slices.Delete(desc.Attr, index, index+1)

					break
				}
			}

			existingClass := []string{}
			for key, attr := range desc.Attr {
				if attr.Key == "class" {
					existingClass = strings.Split(attr.Val, " ")
				}
				if !styleAttrPattern.MatchString(attr.Key) {
					continue
				}

				styleData := strings.Split(attr.Key, "-")
				if styleData[2] == "classname" {
					continue
				}

				attrVariantName := strings.Join(styleData[1:], "-")
				classMap[className][string(styles.VariantNameOfAttr(attrVariantName))] = attr.Val
				desc.Attr = slices.Delete(desc.Attr, key, key+1)
			}

			if className != "" {
				existingClass = append(existingClass, className)
				desc.Attr = append(desc.Attr, html.Attribute{Key: "class", Val: strings.Join(existingClass, " ")})
			}
		}
	}

	css := []string{}
	for className, variantMap := range classMap {
		for variant, props := range variantMap {
			var s string
			if variant == "" {
				s = fmt.Sprintf(`.%s {
					%s
				}`, className, props)

			} else {
				s = fmt.Sprintf(`.%s%s {
					%s
				}`, className, variant, props)

			}
			css = append(css, s)
		}
	}

	return strings.Join(css, "\n")
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
