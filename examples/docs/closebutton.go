package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderCloseButton() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/buttons"

var CloseButton = buttons.CloseButton`, nil),
		Playground("Basic Usage", "Standard implementation of the CloseButton component.", `CloseButton()`, CloseButton().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
