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
var ruleBracketsPattern = regexp.MustCompile(`\{(.*)\}`)
var hasAtPattern = regexp.MustCompile(`@(.*)\((.*)\)`)
var ruleCssPattern = regexp.MustCompile(`(?P<Rule>&(.*?)})`)

func ExtractCSS(c IsComponent) string {
	el := c.Element()
	css := []string{}
	props := []string{}
	// s, _ := sqids.New(sqids.Options{
	// 	Alphabet: "abcdefghijklmnopqrstuvwxyz",
	// })

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

	slices.Sort(props)
	props = slices.Compact(props)

	variantClassMap := map[string][]string{}
	variantPropsMap := map[string]map[string]string{}
	for desc := range el.GetNode().Descendants() {
		if desc.Type == html.ElementNode {
			classnameAttrIndex := slices.IndexFunc(desc.Attr, func(attr html.Attribute) bool { return attr.Key == "data-classname" })
			if classnameAttrIndex == -1 {
				continue
			}
			className := desc.Attr[classnameAttrIndex].Val

			for _, attr := range desc.Attr {
				if !styleAttrPattern.MatchString(attr.Key) {
					continue
				}

				variantStyleName := strings.Join(strings.Split(attr.Key, "-")[2:], "-")

				if variantPropsMap[variantStyleName] == nil {
					variantPropsMap[variantStyleName] = make(map[string]string)
				}
				variantPropsMap[variantStyleName][className] = attr.Val
				// desc.Attr[attrIndex].Val = strings.Trim(string(s), "[]")
				variantClassMap[variantStyleName] = append(
					variantClassMap[variantStyleName], className,
				)
			}
		}
	}

	for variantName, classNames := range variantClassMap {
		splittedVariant := strings.Split(variantName, ":")

		var computedClassName string
		var placeholder string = "& { %s }"
		if len(splittedVariant) > 1 {
			scope := []string{}
			for _, v := range splittedVariant {
				val := string(styles.VariantNameOfAttr(v))
				submatch := ruleCssPattern.FindAllStringSubmatch(val, -1)
				if len(submatch) > 0 {
					placeholder = submatch[0][0]
				}
				scope = append(scope, ruleBracketsPattern.ReplaceAllString(val, ""))
			}
			computedClassName = strings.Join(scope, " and ")
			computedClassName = strings.ReplaceAll(computedClassName, " and @media", "and ")
		} else {
			val := string(styles.VariantNameOfAttr(variantName))
			submatch := ruleCssPattern.FindAllStringSubmatch(val, -1)
			if len(submatch) > 0 {
				placeholder = submatch[0][0]
			}

			computedClassName = ruleBracketsPattern.ReplaceAllString(string(styles.VariantNameOfAttr(variantName)), "")
		}

		var classPlaceHolder []string
		for _, className := range classNames {
			p := strings.ReplaceAll(placeholder, "&", "."+className)
			p = fmt.Sprintf(p, variantPropsMap[variantName][className])
			classPlaceHolder = append(classPlaceHolder, p)
		}

		var cssText string
		if hasAtPattern.MatchString(computedClassName) {
			cssText = fmt.Sprintf(`
				%s {
					%s
				}
			`, computedClassName, strings.Join(classPlaceHolder, "\n"))
		} else {
			cssText = strings.Join(classPlaceHolder, "\n")
		}

		css = append(css, cssText)
	}

	// classes := map[string]string{}

	for desc := range el.GetNode().Descendants() {
		if desc.Type == html.ElementNode {
			newAttr := []html.Attribute{}
			className := ""
			for _, attr := range desc.Attr {
				if attr.Key == "class" {
					className = attr.Val
				}

				if attr.Key == "data-classname" {
					className = className + " " + attr.Val
				}

				if !styleAttrPattern.MatchString(attr.Key) && attr.Key != "data-classname" {
					newAttr = append(newAttr, attr)
					continue
				}
			}

			classIndex := slices.IndexFunc(desc.Attr, func(a html.Attribute) bool { return a.Key == "class" })
			if classIndex > -1 {
				desc.Attr[classIndex].Val = className
			} else {
				newAttr = append(newAttr, html.Attribute{Key: "class", Val: className})
			}
			desc.Attr = newAttr
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(css)))

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
