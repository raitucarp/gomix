package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	buttonsPkg "github.com/raitucarp/gomix/components/buttons"
)

func RenderCloseButton() components.IsComponent {
	// Alias simulation for documentation
	var CloseButton = buttonsPkg.CloseButton

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/buttons"

var CloseButton = buttons.CloseButton`, nil),
		Playground("Basic Usage", "Standard implementation of the CloseButton component.", `CloseButton()`, CloseButton().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
