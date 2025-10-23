package styles

import (
	"fmt"
	"maps"
	"regexp"
	"sort"
	"strings"

	"github.com/raitucarp/gomix/constants"
	"github.com/raitucarp/gomix/theme"
)

const themeColor = constants.ThemeColor
const themeFont = constants.ThemeFont
const themeText = constants.ThemeText
const themeFontWeight = constants.ThemeFontWeight
const themeTracking = constants.ThemeTracking
const themeLeading = constants.ThemeLeading
const themeBreakpoint = constants.ThemeBreakpoint
const themeContainer = constants.ThemeContainer
const themeSpacing = constants.ThemeSpacing
const themeRadius = constants.ThemeRadius
const themeShadow = constants.ThemeShadow
const themeInsetShadow = constants.ThemeInsetShadow
const themeDropShadow = constants.ThemeDropShadow
const themeTextShadow = constants.ThemeTextShadow
const themeBlur = constants.ThemeBlur
const themePerspective = constants.ThemePerspective
const themeAspect = constants.ThemeAspect
const themeEase = constants.ThemeEase
const themeAnimate = constants.ThemeAnimate
const themeCustom = constants.ThemeCustom

type Properties map[string]string

func (p *Properties) StyleProp() {}
func (p *Properties) GetProps() map[string]string {
	return *p
}
func (p *Properties) Variant() StyleVariant {
	return defaultVar
}
func (p *Properties) Kind() kind {
	return kindProps
}

type Props map[Prop]string

type Style struct {
	props   Props
	variant StyleVariant
	theme   *theme.Theme
}

func (s *Style) addProp(prop Prop, value string) {
	s.props[prop] = value
}

type kind string

const (
	kindProps          kind = "props"
	kindVariant        kind = "variant"
	kindSelectorParams kind = "selector-params"
)

type ApplyProp func(s *Style) StyleProp
type StyleProp interface {
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

func CSSProperty(property string, value string) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			property: value,
		}
	}
}

func ApplyStyle(styleTheme *theme.Theme, props ...ApplyProp) map[StyleVariant]Props {
	defaultStyle := &Style{
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

		variantStyle := &Style{
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

var ruleBracketsPattern = regexp.MustCompile(`\{(.*)\}`)
var hasAtPattern = regexp.MustCompile(`@(.*)\((.*)\)`)
var ruleCssPattern = regexp.MustCompile(`(?P<Rule>&(.*?)})`)

func ExtractCSSFromStyle(variantClassMap map[string][]string, variantPropsMap map[string]map[string]string) []string {
	css := []string{}

	for variantName, classNames := range variantClassMap {
		splittedVariant := strings.Split(variantName, ":")

		var computedClassName string
		var placeholder string = "& { %s }"
		if len(splittedVariant) > 1 {
			scope := []string{}
			for _, v := range splittedVariant {
				val := string(VariantNameOfAttr(v))
				submatch := ruleCssPattern.FindAllStringSubmatch(val, -1)
				if len(submatch) > 0 {
					placeholder = submatch[0][0]
				}
				scope = append(scope, ruleBracketsPattern.ReplaceAllString(val, ""))
			}
			computedClassName = strings.Join(scope, " and ")
			computedClassName = strings.ReplaceAll(computedClassName, " and @media", "and ")
		} else {
			val := string(VariantNameOfAttr(variantName))
			submatch := ruleCssPattern.FindAllStringSubmatch(val, -1)
			if len(submatch) > 0 {
				placeholder = submatch[0][0]
			}

			computedClassName = ruleBracketsPattern.ReplaceAllString(string(VariantNameOfAttr(variantName)), "")
		}

		var classPlaceHolder []string
		for _, className := range classNames {
			p := strings.ReplaceAll(placeholder, "&", className)
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

	sort.Sort(sort.Reverse(sort.StringSlice(css)))

	return css
}
