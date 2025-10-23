package components

import (
	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/value"
)

type aspectRatio struct {
	component *Comp
}

func (a *aspectRatio) Ratio(antecedent, consequent int) *aspectRatio {
	a.Component().Aspect(value.Fraction(value.Number(antecedent), value.Number(consequent)))
	return a
}

func (a *aspectRatio) Component() *Comp {
	return a.component
}

func AspectRatio() *aspectRatio {
	c := &aspectRatio{
		component: &Comp{
			El: element.Div().Element(),
		},
	}

	return c
}
