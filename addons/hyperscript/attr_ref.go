package hyperscript

import (
	"fmt"
	"strings"
)

func AttributeOf(v hasAttribute) *attributeRef {
	return v.Attribute()
}

type attributeRef struct {
	attached string
	key      string
	value    string
}

func (ref *attributeRef) IsRef() {}

func (ref *attributeRef) String() string {
	s := []string{}

	if ref.attached != "" {
		s = append(s, fmt.Sprintf("%s.%s", ref.attached, ref.key))
		return strings.Join(s, "")
	}

	if ref.key != "" && ref.value != "" {
		s = append(s, fmt.Sprintf("[@%s='%s']", ref.key, ref.value))
		return strings.Join(s, "")
	}

	s = append(s, fmt.Sprintf("@%s", ref.key))
	return strings.Join(s, "")
}

func Attribute() *attributeRef {
	return &attributeRef{}
}

func (attr *attributeRef) Accept(value ...string) *attributeRef {
	attr.key = "accept"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) AcceptCharset(value ...string) *attributeRef {
	attr.key = "acceptcharset"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}

func (attr *attributeRef) AccessKey(value ...string) *attributeRef {
	attr.key = "accesskey"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Action(value ...string) *attributeRef {
	attr.key = "action"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Align(value ...string) *attributeRef {
	attr.key = "align"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Allow(value ...string) *attributeRef {
	attr.key = "allow"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Alpha(value ...string) *attributeRef {
	attr.key = "alpha"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Alt(value ...string) *attributeRef {
	attr.key = "alt"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) As(value ...string) *attributeRef {
	attr.key = "as"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Async(value ...string) *attributeRef {
	attr.key = "async"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) AutoCapitalize(value ...string) *attributeRef {
	attr.key = "autocapitalize"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) AutoComplete(value ...string) *attributeRef {
	attr.key = "autocomplete"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) AutoPlay(value ...string) *attributeRef {
	attr.key = "autoplay"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Background(value ...string) *attributeRef {
	attr.key = "background"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Bgcolor(value ...string) *attributeRef {
	attr.key = "bgcolor"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Border(value ...string) *attributeRef {
	attr.key = "border"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Capture(value ...string) *attributeRef {
	attr.key = "capture"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Charset(value ...string) *attributeRef {
	attr.key = "charset"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Checked(value ...string) *attributeRef {
	attr.key = "checked"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Cite(value ...string) *attributeRef {
	attr.key = "cite"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Class(value ...string) *attributeRef {
	attr.key = "class"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Color(value ...string) *attributeRef {
	attr.key = "color"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Colorspace(value ...string) *attributeRef {
	attr.key = "colorspace"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Cols(value ...string) *attributeRef {
	attr.key = "cols"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Colspan(value ...string) *attributeRef {
	attr.key = "colspan"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Content(value ...string) *attributeRef {
	attr.key = "content"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) ContentEditable(value ...string) *attributeRef {
	attr.key = "contenteditable"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Controls(value ...string) *attributeRef {
	attr.key = "controls"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Coords(value ...string) *attributeRef {
	attr.key = "coords"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) CrossOrigin(value ...string) *attributeRef {
	attr.key = "crossorigin"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Csp(value ...string) *attributeRef {
	attr.key = "csp"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Data(value ...string) *attributeRef {
	attr.key = "data"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Datetime(value ...string) *attributeRef {
	attr.key = "datetime"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Decoding(value ...string) *attributeRef {
	attr.key = "decoding"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Default(value ...string) *attributeRef {
	attr.key = "default"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Defer(value ...string) *attributeRef {
	attr.key = "defer"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Dir(value ...string) *attributeRef {
	attr.key = "dir"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Dirname(value ...string) *attributeRef {
	attr.key = "dirname"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Disabled(value ...string) *attributeRef {
	attr.key = "disabled"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Download(value ...string) *attributeRef {
	attr.key = "download"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Draggable(value ...string) *attributeRef {
	attr.key = "draggable"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Enctype(value ...string) *attributeRef {
	attr.key = "enctype"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) EnterKeyhint(value ...string) *attributeRef {
	attr.key = "enterkeyhint"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) ElementTiming(value ...string) *attributeRef {
	attr.key = "elementtiming"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) For(value ...string) *attributeRef {
	attr.key = "for"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Form(value ...string) *attributeRef {
	attr.key = "form"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}

func (attr *attributeRef) FormAction(value ...string) *attributeRef {
	attr.key = "formaction"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}

func (attr *attributeRef) Formenctype(value ...string) *attributeRef {
	attr.key = "formenctype"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Formmethod(value ...string) *attributeRef {
	attr.key = "formmethod"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Formnovalidate(value ...string) *attributeRef {
	attr.key = "formnovalidate"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}

func (attr *attributeRef) FormTarget(value ...string) *attributeRef {
	attr.key = "formtarget"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Headers(value ...string) *attributeRef {
	attr.key = "headers"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Height(value ...string) *attributeRef {
	attr.key = "height"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Hidden(value ...string) *attributeRef {
	attr.key = "hidden"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) High(value ...string) *attributeRef {
	attr.key = "high"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Href(value ...string) *attributeRef {
	attr.key = "href"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) HrefLang(value ...string) *attributeRef {
	attr.key = "hreflang"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) HttpEquiv(value ...string) *attributeRef {
	attr.key = "httpequiv"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Id(value ...string) *attributeRef {
	attr.key = "id"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Integrity(value ...string) *attributeRef {
	attr.key = "integrity"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) InputMode(value ...string) *attributeRef {
	attr.key = "inputmode"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) IsMap(value ...string) *attributeRef {
	attr.key = "ismap"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) ItemProp(value ...string) *attributeRef {
	attr.key = "itemprop"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Kind(value ...string) *attributeRef {
	attr.key = "kind"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Label(value ...string) *attributeRef {
	attr.key = "label"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Lang(value ...string) *attributeRef {
	attr.key = "lang"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Language(value ...string) *attributeRef {
	attr.key = "language"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Loading(value ...string) *attributeRef {
	attr.key = "loading"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) List(value ...string) *attributeRef {
	attr.key = "list"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Loop(value ...string) *attributeRef {
	attr.key = "loop"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Low(value ...string) *attributeRef {
	attr.key = "low"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Max(value ...string) *attributeRef {
	attr.key = "max"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) MaxLength(value ...string) *attributeRef {
	attr.key = "maxlength"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) MinLength(value ...string) *attributeRef {
	attr.key = "minlength"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Media(value ...string) *attributeRef {
	attr.key = "media"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Method(value ...string) *attributeRef {
	attr.key = "method"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Min(value ...string) *attributeRef {
	attr.key = "min"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Multiple(value ...string) *attributeRef {
	attr.key = "multiple"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Muted(value ...string) *attributeRef {
	attr.key = "muted"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Name(value ...string) *attributeRef {
	attr.key = "name"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) NoValidate(value ...string) *attributeRef {
	attr.key = "novalidate"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Open(value ...string) *attributeRef {
	attr.key = "open"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Optimum(value ...string) *attributeRef {
	attr.key = "optimum"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Pattern(value ...string) *attributeRef {
	attr.key = "pattern"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Ping(value ...string) *attributeRef {
	attr.key = "ping"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Placeholder(value ...string) *attributeRef {
	attr.key = "placeholder"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) PlaysInline(value ...string) *attributeRef {
	attr.key = "playsinline"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Poster(value ...string) *attributeRef {
	attr.key = "poster"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Preload(value ...string) *attributeRef {
	attr.key = "preload"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Readonly(value ...string) *attributeRef {
	attr.key = "readonly"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) ReferrerPolicy(value ...string) *attributeRef {
	attr.key = "referrerpolicy"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Rel(value ...string) *attributeRef {
	attr.key = "rel"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Required(value ...string) *attributeRef {
	attr.key = "required"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Reversed(value ...string) *attributeRef {
	attr.key = "reversed"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Role(value ...string) *attributeRef {
	attr.key = "role"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Rows(value ...string) *attributeRef {
	attr.key = "rows"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Rowspan(value ...string) *attributeRef {
	attr.key = "rowspan"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Sandbox(value ...string) *attributeRef {
	attr.key = "sandbox"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Scope(value ...string) *attributeRef {
	attr.key = "scope"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Selected(value ...string) *attributeRef {
	attr.key = "selected"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Shape(value ...string) *attributeRef {
	attr.key = "shape"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Size(value ...string) *attributeRef {
	attr.key = "size"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Sizes(value ...string) *attributeRef {
	attr.key = "sizes"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Slot(value ...string) *attributeRef {
	attr.key = "slot"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Span(value ...string) *attributeRef {
	attr.key = "span"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Spellcheck(value ...string) *attributeRef {
	attr.key = "spellcheck"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Src(value ...string) *attributeRef {
	attr.key = "src"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) SrcDoc(value ...string) *attributeRef {
	attr.key = "srcdoc"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) SrcLang(value ...string) *attributeRef {
	attr.key = "srclang"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) SrcSet(value ...string) *attributeRef {
	attr.key = "srcset"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Start(value ...string) *attributeRef {
	attr.key = "start"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Step(value ...string) *attributeRef {
	attr.key = "step"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Style(value ...string) *attributeRef {
	attr.key = "style"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Summary(value ...string) *attributeRef {
	attr.key = "summary"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) TabIndex(value ...string) *attributeRef {
	attr.key = "tabindex"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Target(value ...string) *attributeRef {
	attr.key = "target"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Title(value ...string) *attributeRef {
	attr.key = "title"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Translate(value ...string) *attributeRef {
	attr.key = "translate"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Type(value ...string) *attributeRef {
	attr.key = "type"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) UseMap(value ...string) *attributeRef {
	attr.key = "usemap"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Value(value ...string) *attributeRef {
	attr.key = "value"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Width(value ...string) *attributeRef {
	attr.key = "width"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
func (attr *attributeRef) Wrap(value ...string) *attributeRef {
	attr.key = "wrap"
	if len(value) > 0 {
		attr.key = value[0]
	}
	return attr
}
