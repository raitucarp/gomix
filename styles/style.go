package styles

import (
	"fmt"
	"maps"
	"strings"

	"github.com/raitucarp/gomix/theme"
)

type properties map[string]string

func (p *properties) StyleProp() {}
func (p *properties) GetProps() map[string]string {
	return *p
}
func (p *properties) Variant() StyleVariant {
	return defaultVar
}
func (p *properties) Kind() kind {
	return kindProps
}

type Props map[Prop]string

type style struct {
	props   Props
	variant StyleVariant
	theme   *theme.Theme
}

func (s *style) addProp(prop Prop, value string) {
	s.props[prop] = value
}

type kind string

const (
	kindProps          kind = "props"
	kindVariant        kind = "variant"
	kindSelectorParams kind = "selector-params"
)

type ApplyProp func(s *style) styleProp
type styleProp interface {
	StyleProp()
	GetProps() map[string]string
	Kind() kind
	Variant() StyleVariant
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

func ApplyStyle(styleTheme *theme.Theme, props ...ApplyProp) map[StyleVariant]Props {
	defaultStyle := &style{
		props:   make(Props),
		variant: defaultVar,
		theme:   styleTheme,
	}

	allStyles := map[StyleVariant]Props{}

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
		prop := p(defaultStyle)
		if prop.Kind() != kindVariant {
			continue
		}

		variant := prop.Variant()
		if allStyles[variant] == nil {
			allStyles[variant] = make(Props)
		}

		variantStyle := &style{
			props:   make(Props),
			variant: variant,
			theme:   defaultStyle.theme,
		}

		variantProps := p(variantStyle).GetProps()

		for key, value := range variantProps {
			variantStyle.addProp(Prop(key), value)
		}

		allStyles[variant] = variantStyle.props
		if vChild, ok := prop.(*variantOf); ok {
			if len(vChild.children) > 0 {
				cv := combinedVariants(vChild)
				maps.Copy(allStyles, cv)
			}
		}
	}

	return allStyles
}
