package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderRadio() components.IsComponent {
	// Alias simulation for documentation
	var Radio = formsPkg.Radio

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/forms"

var Radio = forms.Radio`, nil),
		Playground("Basic Usage", "Standard implementation of the Radio component.", `Radio()`, Radio().Component()),
		Playground("Prop: ColorScheme", "Demonstrating the ColorScheme method.", `HStack(
	Radio().ColorScheme("blue"),
	Radio().ColorScheme("red"),
	Radio().ColorScheme("green")
)`, HStack(Radio().ColorScheme("blue").Component(), Radio().ColorScheme("red").Component(), Radio().ColorScheme("green").Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		Playground("Prop: Size", "Demonstrating the Size method.", `HStack(
	Radio().Size("sm"),
	Radio().Size("md"),
	Radio().Size("lg")
)`, HStack(Radio().Size("sm").Component(), Radio().Size("md").Component(), Radio().Size("lg").Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		PropsTable([]map[string]string{
			{"name": "ColorScheme", "type": "color string", "description": "Chainable method for ColorScheme setting Gomix attributes/styles."},
			{"name": "Size", "type": "s string", "description": "Chainable method for Size setting Gomix attributes/styles."},
		}),
	).Component().GapBy(value.Rem(2))
}
