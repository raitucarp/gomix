package element

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"slices"
	"strconv"
	"strings"

	"github.com/raitucarp/gomix/styles"
	"github.com/raitucarp/gomix/theme"
	"golang.org/x/net/html"
)

type Direction string

const (
	DirectionLtr  Direction = "ltr"
	DirectionRtl  Direction = "rtl"
	DirectionAuto Direction = "auto"
)

type KeyHint string

const (
	KeyHintDone     KeyHint = "Done"
	KeyHintEnter    KeyHint = "Enter"
	KeyHintGo       KeyHint = "Go"
	KeyHintNext     KeyHint = "Next"
	KeyHintPrevious KeyHint = "Previous"
	KeyHintSearch   KeyHint = "Search"
	KeyHintSend     KeyHint = "Send"
)

type InputMode string

const (
	InputModeDecimal InputMode = "decimal"
	InputModeEmail   InputMode = "text"
	InputModeNone    InputMode = "none"
	InputModeNumeric InputMode = "numeric"
	InputModeSearch  InputMode = "search"
	InputModeTel     InputMode = "tel"
	InputModeText    InputMode = "text"
	InputModeUrl     InputMode = "url"
)

type TranslateOptions string

const (
	TranslateYes TranslateOptions = "yes"
	TranslateNo  TranslateOptions = "no"
)

type globalAttributes interface {
	AccessKey(character string) *HtmlElement
	Class(classnames ...string) *HtmlElement
	ContentEditable(editable bool) *HtmlElement
	Data(name string, value string) *HtmlElement
	Dir(dir Direction) *HtmlElement
	Draggable(draggable bool) *HtmlElement
	EnterKeyHint(hint KeyHint) *HtmlElement
	Hidden() *HtmlElement
	Id(id string) *HtmlElement
	Inert() *HtmlElement
	InputMode(mode InputMode) *HtmlElement
	Lang(langcode LanguageCode) *HtmlElement
	Popover() *HtmlElement
	Spellcheck(check bool) *HtmlElement
	Style(styles map[string]any) *HtmlElement
	TabIndex(number int) *HtmlElement
	Title(text string) *HtmlElement
	Translate(translate TranslateOptions) *HtmlElement
}

type IsElement interface {
	Element() *HtmlElement
}

func Element(el IsElement) *HtmlElement {
	a := el
	return a.Element()
}

type HtmlElement struct {
	node          *html.Node
	theme         *theme.Theme
	style         map[styles.StyleVariant]styles.Props
	prefClassName string
}

func (e *HtmlElement) AddAttribute(key string, value string) {
	currentAttrIndex := slices.IndexFunc(e.node.Attr, func(attr html.Attribute) bool {
		return key == attr.Key
	})

	if currentAttrIndex > -1 {
		e.node.Attr[currentAttrIndex].Val = value
	} else {
		e.node.Attr = append(e.node.Attr, html.Attribute{Namespace: "", Key: key, Val: value})
	}
}

func (e *HtmlElement) AddAttributeBool(key string) {
	currentAttrIndex := slices.IndexFunc(e.node.Attr, func(attr html.Attribute) bool {
		return key == attr.Key
	})

	if currentAttrIndex < 0 {
		e.node.Attr = append(e.node.Attr, html.Attribute{Namespace: "", Key: key})
	}
}

func (e *HtmlElement) GetAttribute(key string) string {
	for _, attr := range e.node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}

	return ""
}

func (e *HtmlElement) GetData(key string) string {
	return e.GetAttribute("data-" + key)
}

func (e *HtmlElement) Children(children ...IsElement) *HtmlElement {
	for _, el := range children {
		e.node.AppendChild(el.Element().GetNode())
	}
	return e
}

func (e *HtmlElement) Render() string {
	var buffer bytes.Buffer
	html.Render(&buffer, e.node)

	return buffer.String()
}

func (e *HtmlElement) AccessKey(character string) *HtmlElement {
	e.AddAttribute("accesskey", character)
	return e
}

func (e *HtmlElement) Class(classnames ...string) *HtmlElement {
	existingClassName := []string{}
	for _, attr := range e.node.Attr {
		if attr.Key == "class" {
			existingClassName = append(existingClassName, strings.Split(attr.Val, " ")...)
			break
		}
	}

	existingClassName = append(existingClassName, classnames...)
	e.AddAttribute("class", strings.Join(existingClassName, " "))
	return e
}

func (e *HtmlElement) ContentEditable(editable bool) *HtmlElement {
	if editable {
		e.AddAttributeBool("contenteditable")
	} else {
		for i, attr := range e.node.Attr {
			if attr.Key == "contenteditable" {
				e.node.Attr = slices.Delete(e.node.Attr, i, i+1)
			}
		}
	}
	return e
}

func (e *HtmlElement) Data(name string, value string) *HtmlElement {
	e.AddAttribute(strings.Join([]string{"data", name}, "-"), value)

	return e
}

func (e *HtmlElement) Dir(dir Direction) *HtmlElement {
	e.AddAttribute("dir", string(dir))
	return e
}
func (e *HtmlElement) Draggable(draggable bool) *HtmlElement {
	e.AddAttributeBool("draggable")
	return e
}
func (e *HtmlElement) EnterKeyHint(hint KeyHint) *HtmlElement {
	e.AddAttribute("enterkeyhint", string(hint))
	return e
}
func (e *HtmlElement) Hidden() *HtmlElement {
	e.AddAttributeBool("hidden")
	return e
}
func (e *HtmlElement) Id(id string) *HtmlElement {
	e.AddAttribute("id", id)
	return e
}
func (e *HtmlElement) Inert() *HtmlElement {
	e.AddAttributeBool("inert")
	return e
}
func (e *HtmlElement) InputMode(mode InputMode) *HtmlElement {
	e.AddAttribute("inputmode", string(mode))
	return e
}
func (e *HtmlElement) Lang(langcode LanguageCode) *HtmlElement {
	e.AddAttribute("lang", string(langcode))
	return e
}
func (e *HtmlElement) Popover() *HtmlElement {
	e.AddAttributeBool("popover")
	return e
}
func (e *HtmlElement) Spellcheck(check bool) *HtmlElement {
	e.AddAttribute("spellcheck", strconv.FormatBool(check))
	return e
}

func (e *HtmlElement) Role(role string) *HtmlElement {
	e.AddAttribute("role", role)
	return e
}

func (e *HtmlElement) getClassSha() string {
	hasher := sha256.New()
	hasher.Write([]byte(e.Render()))

	return "c" + hex.EncodeToString(hasher.Sum(nil))[0:5]
}

func (e *HtmlElement) PreferredClassName(className string) *HtmlElement {
	e.prefClassName = className
	return e
}

func (e *HtmlElement) ClassName() string {
	className := e.prefClassName

	if className == "" {
		className = e.getClassSha()
	}
	return className
}

func (e *HtmlElement) Style(props ...styles.ApplyProp) *HtmlElement {
	e.style = styles.ApplyStyle(e.theme, props...)

	for variant, props := range e.style {
		styleProps := []string{}
		for prop, value := range props {
			styleProps = append(styleProps, strings.Join([]string{string(prop), value}, ":"))
		}

		styleKey := "style-" + variant.Name()
		styleValue := strings.Join(styleProps, ";")
		currentValue := e.GetData(styleKey)
		if currentValue != "" {
			styleValue = currentValue + ";" + styleValue
		}

		e.Data(styleKey, styleValue)
	}

	e.Data("classname", e.ClassName())

	return e
}
func (e *HtmlElement) TabIndex(number int) *HtmlElement {
	e.AddAttribute("tabindex", strconv.Itoa(number))
	return e
}

func (e *HtmlElement) CSSProperty(key string, value string) *HtmlElement {
	e.Style(styles.CSSProperty(key, value))
	return e
}

func (e *HtmlElement) Title(text string) *HtmlElement {
	e.AddAttribute("title", text)
	return e
}

func (e *HtmlElement) Translate(translate TranslateOptions) *HtmlElement {
	e.AddAttribute("translate", string(translate))
	return e
}

func (e *HtmlElement) Aria(aria string, value string) *HtmlElement {
	e.AddAttribute("aria-"+aria, value)
	return e
}

func (e *HtmlElement) GetNode() *html.Node {
	return e.node
}

func (e *HtmlElement) ChangeTagName(tagName string) *html.Node {
	e.node.Data = tagName
	return e.node
}

func (e *HtmlElement) Element() *HtmlElement {
	return e
}

func CreateElementByNode(node *html.Node) *HtmlElement {
	return &HtmlElement{node: node}
}
