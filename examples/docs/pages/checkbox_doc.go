package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderCheckboxDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default checkbox component.", `formsPkg.Checkbox()`, formsPkg.Checkbox().Component()),
		PropsTable([]map[string]string{
			{"name": "ColorScheme", "type": "color string", "description": "Chainable method for ColorScheme."},
			{"name": "Size", "type": "s string", "description": "Chainable method for Size."},
			{"name": "IsInvalid", "type": "invalid bool", "description": "Chainable method for IsInvalid."},
		}),
	).Component().GapBy(value.Rem(2))
}
