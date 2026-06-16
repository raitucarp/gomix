package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderRadioDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default radio component.", `formsPkg.Radio()`, formsPkg.Radio().Component()),
		PropsTable([]map[string]string{
			{"name": "ColorScheme", "type": "color string", "description": "Chainable method for ColorScheme."},
			{"name": "Size", "type": "s string", "description": "Chainable method for Size."},
		}),
	).Component().GapBy(value.Rem(2))
}
