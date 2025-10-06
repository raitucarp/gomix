package components

import (
	"bytes"
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/styles"
	"github.com/sqids/sqids-go"
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

var styleAttrPattern = regexp.MustCompile(`data-style-`)

func ExtractCSS(c IsComponent) string {
	// classMap := make(map[string]map[string]string)
	el := c.Element()
	css := []string{}

	props := []string{}

	for desc := range el.GetNode().Descendants() {
		if desc.Type == html.ElementNode {
			for _, attr := range desc.Attr {
				if !styleAttrPattern.MatchString(attr.Key) {
					continue
				}

				props = append(props, strings.Split(attr.Val, ";")...)
			}
		}
	}

	sort.Strings(props)
	props = slices.Compact(props)

	s, _ := sqids.New(sqids.Options{
		Alphabet: "abcdefghijklmnopqrstuvwxyz",
	})

	for desc := range el.GetNode().Descendants() {
		if desc.Type == html.ElementNode {
			// numbers := []uint64{}
			for attrIndex, attr := range desc.Attr {
				if !styleAttrPattern.MatchString(attr.Key) {
					continue
				}

				elProps := strings.Split(attr.Val, ";")
				propNumbers := []uint64{}
				for _, p := range elProps {
					index := slices.IndexFunc(props, func(prop string) bool { return prop == p })
					propNumbers = append(propNumbers, uint64(index))
				}
				slices.Sort(propNumbers)
				propNumbers = slices.Compact(propNumbers)

				className, _ := s.Encode(propNumbers)
				desc.Attr[attrIndex].Val = className
			}
		}
	}

	classes := map[string]string{}
	for desc := range el.GetNode().Descendants() {
		if desc.Type == html.ElementNode {
			for _, attr := range desc.Attr {

				if !styleAttrPattern.MatchString(attr.Key) {
					continue
				}

				key := strings.Join(strings.Split(attr.Key, "-")[2:], "-")
				className := attr.Val
				propIndexes := s.Decode(className)

				properties := []string{}
				for _, propIndex := range propIndexes {
					properties = append(properties, props[propIndex])
				}

				computedClassName := string(styles.VariantNameOfAttr(key))
				computedClassName = strings.ReplaceAll(computedClassName, "&", "."+className)
				classes[computedClassName] = strings.Join(properties, ";")

				existingClass := slices.IndexFunc(desc.Attr, func(a html.Attribute) bool { return a.Key == "class" })

				if existingClass != -1 {
					classes := strings.Split(desc.Attr[existingClass].Val, " ")
					classes = append(classes, className)
					desc.Attr[existingClass].Val = strings.Join(classes, " ")
				} else {
					desc.Attr = append(desc.Attr, html.Attribute{Key: "class", Val: className})
				}
			}

			newAttr := []html.Attribute{}
			for _, attr := range desc.Attr {
				if !styleAttrPattern.MatchString(attr.Key) {
					newAttr = append(newAttr, attr)
					continue
				}
			}
			desc.Attr = newAttr
		}
	}

	for className, properties := range classes {
		css = append(css, fmt.Sprintf(className, properties))
	}

	sort.Strings(css)

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
