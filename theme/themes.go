package theme

import (
	"fmt"
	"sort"
	"strings"

	"github.com/raitucarp/gomix/constants"
)

type Namespace = constants.ThemeNamespace
const Color = constants.ThemeColor
const Font = constants.ThemeFont
const Text = constants.ThemeText
const FontWeight = constants.ThemeFontWeight
const Tracking = constants.ThemeTracking
const Leading = constants.ThemeLeading
const Breakpoint = constants.ThemeBreakpoint
const Container = constants.ThemeContainer
const Spacing = constants.ThemeSpacing
const Radius = constants.ThemeRadius
const Shadow = constants.ThemeShadow
const InsetShadow = constants.ThemeInsetShadow
const DropShadow = constants.ThemeDropShadow
const TextShadow = constants.ThemeTextShadow
const Blur = constants.ThemeBlur
const Perspective = constants.ThemePerspective
const Aspect = constants.ThemeAspect
const Ease = constants.ThemeEase
const Animate = constants.ThemeAnimate
const Custom = constants.ThemeCustom


type UtilityClass map[string]string

func (u *UtilityClass) Value(class string) string {
	return (*u)[class]
}

type Variables map[Namespace]UtilityClass

func (v *Variables) ToCSSVariables() (vars map[string]string) {
	vars = make(map[string]string)
	for namespace, class := range *v {
		for key, value := range class {
			if namespace == Custom {
				varName := fmt.Sprintf("--%s",
					key,
				)
				vars[varName] = value
			} else {
				varNameParts := []string{string(namespace)}
				if key != "" {
					varNameParts = append(varNameParts, key)
				}
				varName := strings.Join(varNameParts, "-")
				varName = fmt.Sprintf("--%s", varName)
				vars[varName] = value
			}
		}
	}
	return
}

type Theme struct {
	name       string
	variables  Variables
	keyframes  map[string]string
	properties []customProperty
}

func (t *Theme) RawCSS() string {
	rawFormat := `:root {
		%s
  }
	`
	allVariables := []string{}
	for key, value := range t.variables.ToCSSVariables() {
		allVariables = append(allVariables, strings.Join([]string{key, value}, ":"))
	}

	sort.Strings(allVariables)

	return fmt.Sprintf(rawFormat, strings.Join(allVariables, ";"))
}

type customProperty struct {
	name         string
	inherits     bool
	initialValue any
}

func (t *Theme) UseVarKey(namespace Namespace, class string) string {
	return fmt.Sprintf("var(%s)", t.VarKey(namespace, class))

}

func (t *Theme) VarKey(namespace Namespace, class string) string {
	if namespace == Custom {
		return fmt.Sprintf("--%s",
			class,
		)
	}

	varName := strings.Join([]string{string(namespace), class}, "-")

	return fmt.Sprintf("--%s", varName)
}

func (t *Theme) AddVariable(namespace Namespace, class string, value string) {

	if t.variables[namespace] == nil {
		t.variables[namespace] = map[string]string{}
	}

	t.variables[namespace][class] = value
}

func (t *Theme) AddKeyframe(name string, value string) {
	t.keyframes[name] = value
}

type ThemeParam func(t *Theme)

func CreateTheme(name string, params ...ThemeParam) *Theme {
	theme := &Theme{name: name, variables: Variables{}, keyframes: make(map[string]string)}
	for _, p := range params {
		p(theme)
	}
	return theme
}
