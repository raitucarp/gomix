package styles

import (
	"strings"

	"github.com/raitucarp/gomix/themes"
)

type Prop string

const (
	AspectRatio Prop = "aspect-ratio"
	Columns     Prop = "columns"
)

type Props map[Prop]string

type Styles struct {
	props Props
	theme *themes.Theme
}

func (s *Styles) addProp(prop Prop, value string) {
	s.props[prop] = value
}

func (s *Styles) bundle() (styles []string) {
	for key, val := range s.props {
		styles = append(styles, strings.Join([]string{string(key), val}, ":")+";")
	}
	return
}

func New() *Styles {
	return &Styles{
		props: make(map[Prop]string),
		theme: themes.Default(),
	}
}
