package components

import (
	"github.com/raitucarp/gomix/element"
)

type listComponent struct {
	component *component
	variant   string
}

func (l *listComponent) Ordered(items ...IsComponent) *listComponent {
	l.component.el.ChangeTagName("ol")

	for _, item := range items {
		l.component.el.Children(item)
	}
	return l
}

func (l *listComponent) Unordered(items ...IsComponent) *listComponent {
	l.component.el.ChangeTagName("ul")

	for _, item := range items {
		l.component.el.Children(item)
	}
	return l
}

func (l *listComponent) TypeArabicIndic() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "arabic-indic, -moz-arabic-indic")
	}
	return l
}

func (l *listComponent) TypeArmenian() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "armenian")
	}
	return l
}

func (l *listComponent) TypeBengali() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "bengali, -moz-bengali")
	}
	return l
}

func (l *listComponent) TypeCambodianKhmer() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "cambodian/khmer")
	}
	return l
}

func (l *listComponent) TypeCircle() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "circle")
	}
	return l
}

func (l *listComponent) TypeCjkDecimal() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "cjk-decimal")
	}
	return l
}

func (l *listComponent) TypeCjkEarthlyBranch() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "cjk-earthly-branch, -moz-cjk-earthly-branch")
	}
	return l
}

func (l *listComponent) TypeCjkHeavenlyStem() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "cjk-heavenly-stem, -moz-cjk-heavenly-stem")
	}
	return l
}

func (l *listComponent) TypeCjkIdeographic() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "cjk-ideographic")
	}
	return l
}

func (l *listComponent) TypeDecimal() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "decimal")
	}
	return l
}

func (l *listComponent) TypeDecimalLeadingZero() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "decimal-leading-zero")
	}
	return l
}

func (l *listComponent) TypeDevanagari() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "devanagari, -moz-devanagari")
	}
	return l
}

func (l *listComponent) TypeDisc() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "disc")
	}
	return l
}

func (l *listComponent) TypeDisclosureClosed() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "disclosure-closed")
	}
	return l
}

func (l *listComponent) TypeDisclosureOpen() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "disclosure-open")
	}
	return l
}

func (l *listComponent) TypeEthiopicNumeric() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "ethiopic-numeric")
	}
	return l
}

func (l *listComponent) TypeGeorgian() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "georgian")
	}
	return l
}

func (l *listComponent) TypeGujarati() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "gujarati, -moz-gujarati")
	}
	return l
}

func (l *listComponent) TypeGurmukhi() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "gurmukhi, -moz-gurmukhi")
	}
	return l
}

func (l *listComponent) TypeHebrew() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "hebrew")
	}
	return l
}

func (l *listComponent) TypeHiragana() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "hiragana")
	}
	return l
}

func (l *listComponent) TypeHiraganaIroha() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "hiragana-iroha")
	}
	return l
}

func (l *listComponent) TypeJapaneseFormal() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "japanese-formal")
	}
	return l
}

func (l *listComponent) TypeJapaneseInformal() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "japanese-informal")
	}
	return l
}

func (l *listComponent) TypeKannada() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "kannada, -moz-kannada")
	}
	return l
}

func (l *listComponent) TypeKatakana() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "katakana")
	}
	return l
}

func (l *listComponent) TypeKatakanaIroha() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "katakana-iroha")
	}
	return l
}

func (l *listComponent) TypeKoreanHangulFormal() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "korean-hangul-formal")
	}
	return l
}

func (l *listComponent) TypeKoreanHanjaFormal() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "korean-hanja-formal")
	}
	return l
}

func (l *listComponent) TypeKoreanHanjaInformal() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "korean-hanja-informal")
	}
	return l
}

func (l *listComponent) TypeLao() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "lao, -moz-lao")
	}
	return l
}

func (l *listComponent) TypeLowerAlpha() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "lower-alpha")
	}
	return l
}

func (l *listComponent) TypeLowerArmenian() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "lower-armenian")
	}
	return l
}

func (l *listComponent) TypeLowerGreek() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "lower-greek")
	}
	return l
}

func (l *listComponent) TypeLowerLatin() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "lower-latin")
	}
	return l
}

func (l *listComponent) TypeLowerRoman() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "lower-roman")
	}
	return l
}

func (l *listComponent) TypeMalayalam() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "malayalam, -moz-malayalam")
	}
	return l
}

func (l *listComponent) TypeMongolian() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "mongolian")
	}
	return l
}

func (l *listComponent) TypeMyanmar() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "myanmar, -moz-myanmar")
	}
	return l
}

func (l *listComponent) TypeNone() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "none")
	}
	return l
}

func (l *listComponent) TypeOriya() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "oriya, -moz-oriya")
	}
	return l
}

func (l *listComponent) TypePersian() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "persian, -moz-persian")
	}
	return l
}

func (l *listComponent) TypeSimpChineseFormal() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "simp-chinese-formal")
	}
	return l
}

func (l *listComponent) TypeSimpChineseInformal() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "simp-chinese-informal")
	}
	return l
}

func (l *listComponent) TypeSquare() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "square")
	}
	return l
}

func (l *listComponent) TypeTamil() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "tamil, -moz-tamil")
	}
	return l
}

func (l *listComponent) TypeTelugu() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "telugu, -moz-telugu")
	}
	return l
}

func (l *listComponent) TypeThai() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "thai, -moz-thai")
	}
	return l
}

func (l *listComponent) TypeTibetan() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "tibetan")
	}
	return l
}

func (l *listComponent) TypeTradChineseFormal() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "trad-chinese-formal")
	}
	return l
}

func (l *listComponent) TypeTradChineseInformal() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "trad-chinese-informal")
	}
	return l
}

func (l *listComponent) TypeUpperAlpha() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "upper-alpha")
	}
	return l
}

func (l *listComponent) TypeUpperArmenian() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "upper-armenian")
	}
	return l
}

func (l *listComponent) TypeUpperLatin() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "upper-latin")
	}
	return l
}

func (l *listComponent) TypeUpperRoman() *listComponent {
	if l.variant == "marker" {
		l.component.el.CSSProperty("list-style-type", "upper-roman")
	}
	return l
}

func (l *listComponent) Outside() *listComponent {
	l.component.el.CSSProperty("list-style-position", "outside")
	return l
}

func (l *listComponent) Inside() *listComponent {
	l.component.el.CSSProperty("list-style-position", "inside")
	return l
}

func (l *listComponent) Marker() *listComponent {
	l.variant = "marker"
	return l
}

func (l *listComponent) Plain() *listComponent {
	l.variant = "plain"
	l.TypeNone()
	return l
}

func (l *listComponent) Unstyled() *listComponent {
	l.TypeNone()
	l.component.M(0)
	l.component.P(0)
	return l
}

func (l *listComponent) Component() *component {
	return l.component
}

func (l *listComponent) Element() *element.HtmlElement {
	return l.component.el
}

func List() *listComponent {
	c := &listComponent{
		component: &component{
			el: element.Ol().Element(),
		},
		variant: "marker",
	}

	c.Element().Role("list")

	return c
}

type listItemComponent struct {
	component *component
}

func (l *listItemComponent) Component() *component {
	return l.component
}
func (l *listItemComponent) IsListItem() {}

func (l *listItemComponent) Element() *element.HtmlElement {
	return l.component.el
}

func ListItem(children ...IsComponent) *listItemComponent {
	el := &listItemComponent{
		component: &component{
			el: element.Li().Element(),
		},
	}

	for _, c := range children {

		el.Element().Children(c)
	}

	return el
}
