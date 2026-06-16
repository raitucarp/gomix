package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderPinInput() components.IsComponent {
	// Alias simulation for documentation
	var PinInput = formsPkg.PinInput

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/forms"

var PinInput = forms.PinInput`, nil),
		Playground("Basic Usage", "Standard implementation of the PinInput component.", `PinInput()`, PinInput().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
