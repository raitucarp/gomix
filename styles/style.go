package styles

import (
	"fmt"
	"strings"

	"github.com/raitucarp/gomix/themes"
)

type Prop string

const (
	alignContentProp        Prop = "align-content"
	alignItemsProp          Prop = "align-items"
	alignSelfProp           Prop = "align-self"
	aspectRatioProp         Prop = "aspect-ratio"
	borderWidthProp         Prop = "border-width"
	bottomProp              Prop = "bottom"
	boxDecorationProp       Prop = "box-decoration"
	boxSizingProp           Prop = "box-sizing"
	breakAfterProp          Prop = "break-after"
	breakBeforeProp         Prop = "break-before"
	breakInsideProp         Prop = "break-inside"
	clearProp               Prop = "clear"
	clipProp                Prop = "hidden"
	colorProp               Prop = "color"
	columnsProp             Prop = "columns"
	contentProp             Prop = "content"
	displayProp             Prop = "display"
	flexBasisProp           Prop = "flex-basis"
	flexDirectionProp       Prop = "flex-direction"
	flexGrowProp            Prop = "flex-grow"
	flexProp                Prop = "flex"
	flexShrinkProp          Prop = "flex-shrink"
	flexWrapProp            Prop = "flex-wrap"
	floatProp               Prop = "float"
	fontFamilyProp          Prop = "font-family"
	fontSizeProp            Prop = "font-size"
	fontStretchProp         Prop = "font-stretch"
	fontStyleProp           Prop = "font-style"
	fontVariantNumericProp  Prop = "font-variant-numeric"
	fontWeightProp          Prop = "font-weight"
	gapProp                 Prop = "gap"
	gapXProp                Prop = "gap-x"
	gapYProp                Prop = "gap-y"
	gridAutoColumnsProp     Prop = "grid-auto-columns"
	gridAutoFlowProp        Prop = "grid-auto-flow"
	gridAutoRowsProp        Prop = "grid-auto-rows"
	gridColumnEndProp       Prop = "grid-column-end"
	gridColumnProp          Prop = "grid-column"
	gridColumnStartProp     Prop = "grid-column-start"
	gridRowEndProp          Prop = "grid-row-end"
	gridRowProp             Prop = "grid-row"
	gridRowStartProp        Prop = "grid-row-start"
	gridTemplateColumnsProp Prop = "grid-template-columns"
	gridTemplateRowsProp    Prop = "grid-template-rows"
	heightProp              Prop = "height"
	insetBlockProp          Prop = "inset-block"
	insetEndProp            Prop = "inset-inline-end"
	insetInlineProp         Prop = "inset-inline"
	insetProp               Prop = "inset"
	insetStartProp          Prop = "inset-inline-start"
	isolationProp           Prop = "isolation"
	justifyContentProp      Prop = "justify-content"
	justifyItemsProp        Prop = "justify-items"
	justifySelfProp         Prop = "justify-self"
	leftProp                Prop = "left"
	letterSpacingProp       Prop = "letter-spacing"
	lineHeightProp          Prop = "line-height"
	listStyleImageProp      Prop = "list-style-image"
	listStylePositionProp   Prop = "list-style-position"
	listStyleTypeProp       Prop = "list-style-type"
	marginBlockProp         Prop = "margin-block"
	marginBottomProp        Prop = "margin-bottom"
	marginInlineEndProp     Prop = "margin-inline-end"
	marginInlineProp        Prop = "margin-inline"
	marginInlineStartProp   Prop = "margin-inline-start"
	marginLeftProp          Prop = "margin-left"
	marginProp              Prop = "margin"
	marginRightProp         Prop = "margin-right"
	marginTopProp           Prop = "margin-top"
	maxHeightProp           Prop = "max-height"
	maxWidthProp            Prop = "max-width"
	minHeightProp           Prop = "min-height"
	minWidthProp            Prop = "min-width"
	mozOsxFontSmoothingProp Prop = "-moz-osx-font-smoothing"
	objectFitProp           Prop = "object-fit"
	objectPositionProp      Prop = "object-position"
	orderProp               Prop = "order"
	overflowProp            Prop = "overflow"
	overflowWrapProp        Prop = "overflow-wrap"
	overflowXProp           Prop = "overflow-x"
	overflowYProp           Prop = "overflow-y"
	overscrollBehaviorProp  Prop = "overscroll-behavior"
	overscrollBehaviorXProp Prop = "overscroll-behavior-x"
	overscrollBehaviorYProp Prop = "overscroll-behavior-y"
	paddingBlockProp        Prop = "padding-block"
	paddingBottomProp       Prop = "padding-bottom"
	paddingInlineEndProp    Prop = "padding-inline-end"
	paddingInlineProp       Prop = "padding-inline"
	paddingInlineStartProp  Prop = "padding-inline-start"
	paddingLeftProp         Prop = "padding-left"
	paddingProp             Prop = "padding"
	paddingRightProp        Prop = "padding-right"
	paddingTopProp          Prop = "padding-top"
	placeContentProp        Prop = "place-content"
	placeItemsProp          Prop = "place-items"
	placeSelfProp           Prop = "place-self"
	positionProp            Prop = "position"
	rightProp               Prop = "right"
	textAlignProp           Prop = "text-align"
	textDecorColorProp      Prop = "text-decoration-color"
	textDecorLineProp       Prop = "text-decoration-line"
	textDecorStyleProp      Prop = "text-decoration-style"
	textDecorThicknessProp  Prop = "text-decoration-thickness"
	textIndentProp          Prop = "text-indent"
	textOverflowProp        Prop = "text-overflow"
	textTransformProp       Prop = "text-transform"
	textUnderlineOffsetProp Prop = "text-underline-offset"
	textWrapProp            Prop = "text-wrap"
	topProp                 Prop = "top"
	verticalAlignProp       Prop = "vertical-align"
	visibilityProp          Prop = "visibility"
	webkitBoxOrientProp     Prop = "-webkit-box-orient"
	webkitFontSmoothingProp Prop = "-webkit-font-smoothing"
	webkitLineClampProp     Prop = "-webkit-line-clamp"
	whiteSpaceProp          Prop = "white-space"
	widthProp               Prop = "width"
	wordBreakProp           Prop = "word-break"
	zIndexProp              Prop = "z-index"
)

type CustomValue interface {
	String() string
}

type properties map[string]string

func (p *properties) StyleProp() {}
func (p *properties) GetProps() map[string]string {
	return *p
}
func (p *properties) Variant() variant {
	return defaultVar
}
func (p *properties) Kind() kind {
	return kindProps
}

type Props map[Prop]string

type style struct {
	props   Props
	variant variant
	theme   *themes.Theme
}

func (s *style) addProp(prop Prop, value string) {
	s.props[prop] = value
}

type kind string

const (
	kindProps   kind = "props"
	kindVariant kind = "variant"
)

type ApplyProp func(s *style) styleProp
type styleProp interface {
	StyleProp()
	GetProps() map[string]string
	Kind() kind
	Variant() variant
}

type atRules struct {
	name   string
	params []string
	prop   Prop
}

func (r atRules) String() string {
	return fmt.Sprintf("@%s (%s) | %s", r.name, strings.Join(r.params, ", "), r.prop)
}

func mediaQueryProp(prop Prop, params ...string) atRules {
	return atRules{name: "media", params: params, prop: prop}
}

/*
styles(
	aspectRatioVideo(),
)
*/

func Style(props ...ApplyProp) map[variant]Props {
	defaultStyle := &style{
		props:   make(Props),
		variant: defaultVar,
		// theme:   themes.DefaultTheme(),
	}

	allStyles := map[variant]Props{}

	for _, p := range props {
		if p(defaultStyle).Kind() == kindProps {
			props := p(defaultStyle).GetProps()
			for key, value := range props {
				defaultStyle.addProp(Prop(key), value)
			}
		}
	}

	allStyles[defaultVar] = defaultStyle.props

	for _, p := range props {
		if p(defaultStyle).Kind() == kindVariant {
			variant := p(defaultStyle).Variant()
			if allStyles[variant] == nil {
				allStyles[variant] = make(Props)
			}

			variantStyle := &style{
				props:   make(Props),
				variant: variant,
				theme:   defaultStyle.theme,
			}

			props := p(variantStyle).GetProps()
			for key, value := range props {
				variantStyle.addProp(Prop(key), value)
			}

			allStyles[variant] = variantStyle.props
		}
	}

	return allStyles
}
