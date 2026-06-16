package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderIconButton() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/buttons"

var IconButton = buttons.IconButton`, nil),
		Playground("Basic Usage", "Standard implementation of the IconButton component.", `IconButton(text("X"))`, IconButton(text("X")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
