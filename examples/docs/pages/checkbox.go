package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderCheckbox() components.IsComponent {
	// Alias simulation for documentation
	var Checkbox = formsPkg.Checkbox

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/forms"

var Checkbox = forms.Checkbox`, nil),
		Playground("Basic Usage", "Standard implementation of the Checkbox component.", `Checkbox()`, Checkbox().Component()),
		Playground("Prop: ColorScheme", "Demonstrating the ColorScheme method.", `HStack(
	Checkbox().ColorScheme("blue"),
	Checkbox().ColorScheme("red"),
	Checkbox().ColorScheme("green")
)`, HStack(Checkbox().ColorScheme("blue").Component(), Checkbox().ColorScheme("red").Component(), Checkbox().ColorScheme("green").Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		Playground("Prop: Size", "Demonstrating the Size method.", `HStack(
	Checkbox().Size("sm"),
	Checkbox().Size("md"),
	Checkbox().Size("lg")
)`, HStack(Checkbox().Size("sm").Component(), Checkbox().Size("md").Component(), Checkbox().Size("lg").Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		Playground("Prop: IsInvalid", "Demonstrating the IsInvalid method.", `HStack(
	Checkbox().IsInvalid(true),
	Checkbox().IsInvalid(false)
)`, HStack(Checkbox().IsInvalid(true).Component(), Checkbox().IsInvalid(false).Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		PropsTable([]map[string]string{
			{"name": "ColorScheme", "type": "color string", "description": "Chainable method for ColorScheme setting Gomix attributes/styles."},
			{"name": "Size", "type": "s string", "description": "Chainable method for Size setting Gomix attributes/styles."},
			{"name": "IsInvalid", "type": "invalid bool", "description": "Chainable method for IsInvalid setting Gomix attributes/styles."},
		}),
	).Component().GapBy(value.Rem(2))
}
