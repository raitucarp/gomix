package theme

import (
	"fmt"
	"sort"
	"strings"
)

type Namespace string

const (
	Color       Namespace = "color"
	Font        Namespace = "font"
	Text        Namespace = "text"
	FontWeight  Namespace = "fontweight"
	Tracking    Namespace = "tracking"
	Leading     Namespace = "leading"
	Breakpoint  Namespace = "breakpoint"
	Container   Namespace = "container"
	Spacing     Namespace = "spacing"
	Radius      Namespace = "radius"
	Shadow      Namespace = "shadow"
	InsetShadow Namespace = "inset-shadow"
	DropShadow  Namespace = "drop-shadow"
	TextShadow  Namespace = "text-shadow"
	Blur        Namespace = "blur"
	Perspective Namespace = "perspective"
	Aspect      Namespace = "aspect"
	Ease        Namespace = "ease"
	Animate     Namespace = "animate"
	Custom      Namespace = ""
)

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
				varName := fmt.Sprintf("--%s-%s",
					namespace, key,
				)
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

	return fmt.Sprintf("--%s-%s",
		namespace, class,
	)
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
