package styles

import (
	"fmt"
	"strings"

	"github.com/raitucarp/gomix/theme"
)

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
	theme   *theme.Theme
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

func notLastChildProp(prop Prop) string {
	return fmt.Sprintf("& > :not(:last-child) | %s", prop)
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
		// theme:   theme.DefaultTheme(),
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
