package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	buttonsPkg "github.com/raitucarp/gomix/components/buttons"
)

func RenderIconButton() components.IsComponent {
	// Alias simulation for documentation
	var IconButton = buttonsPkg.IconButton

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/buttons"

var IconButton = buttons.IconButton`, nil),
		Playground("Basic Usage", "Standard implementation of the IconButton component.", `IconButton(text("X"))`, IconButton(text("X")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
