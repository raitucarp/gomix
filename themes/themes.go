package themes

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
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
	InsetShadow Namespace = "insetshadow"
	DropShadow  Namespace = "dropshadow"
	Blur        Namespace = "blur"
	Perspective Namespace = "perspective"
	Aspect      Namespace = "aspect"
	Ease        Namespace = "ease"
	Animate     Namespace = "animate"
)

type UtilityClass map[string]string

func (u *UtilityClass) Value(class string) string {
	return (*u)[class]
}

type Variables map[Namespace]UtilityClass

func NewVariable(namespace Namespace, class string, value ...string) *Variables {
	v := make(Variables)
	v[namespace] = make(UtilityClass)
	v[namespace][class] = strings.Join(value, ",")
	return &v
}

func (v *Variables) ToCSSVariables() (vars []string) {
	for namespace, class := range *v {
		for c := range class {
			vars = append(vars, fmt.Sprintf("--%s-%s",
				namespace, strcase.ToDelimited(c, '-'),
			))
		}
	}
	return
}

type Theme struct {
	name      string
	variables Variables
}

func (t *Theme) UseVarKey(namespace Namespace, class string) string {
	return fmt.Sprintf("var(%s)", t.VarKey(namespace, class))

}

func (t *Theme) VarKey(namespace Namespace, class string) string {
	return fmt.Sprintf("--%s-%s",
		namespace, strcase.ToDelimited(class, '-'),
	)
}

func (t *Theme) AddVariable(namespace Namespace, class string, value string) {
	if t.variables[namespace] == nil {
		t.variables[namespace] = map[string]string{}
	}

	t.variables[namespace][class] = value
}

func (t *Theme) ListVariables() (vars []string) {
	return t.variables.ToCSSVariables()
}

func NewTheme(name string) *Theme {
	return &Theme{name: name, variables: Variables{}}
}
