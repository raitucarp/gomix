package styles

import (
	"github.com/raitucarp/gomix/themes"
)

type Prop string

const (
	aspectRatioProp         Prop = "aspect-ratio"
	columnsProp             Prop = "columns"
	breakAfterProp          Prop = "break-after"
	breakBeforeProp         Prop = "break-before"
	breakInsideProp         Prop = "break-inside"
	boxDecorationProp       Prop = "box-decoration"
	boxSizingProp           Prop = "box-sizing"
	displayProp             Prop = "display"
	positionProp            Prop = "position"
	widthProp               Prop = "width"
	heightProp              Prop = "height"
	paddingProp             Prop = "padding"
	marginProp              Prop = "margin"
	overflowProp            Prop = "overflow"
	clipProp                Prop = "hidden"
	whiteSpaceProp          Prop = "white-space"
	borderWidthProp         Prop = "border-width"
	floatProp               Prop = "float"
	clearProp               Prop = "clear"
	isolationProp           Prop = "isolation"
	objectFitProp           Prop = "object-fit"
	objectPositionProp      Prop = "object-position"
	overflowXProp           Prop = "overflow-x"
	overflowYProp           Prop = "overflow-y"
	overscrollBehaviorProp  Prop = "overscroll-behavior"
	overscrollBehaviorXProp Prop = "overscroll-behavior-x"
	overscrollBehaviorYProp Prop = "overscroll-behavior-y"
	visibilityProp          Prop = "visibility"
	zIndexProp              Prop = "z-index"
	insetProp               Prop = "inset"
	insetInlineProp         Prop = "inset-inline"
	insetBlockProp          Prop = "inset-block"
	insetStartProp          Prop = "inset-inline-start"
	insetEndProp            Prop = "inset-inline-end"
	topProp                 Prop = "top"
	bottomProp              Prop = "bottom"
	leftProp                Prop = "left"
	rightProp               Prop = "right"
)

type customValue interface {
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
