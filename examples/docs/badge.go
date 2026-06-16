package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderBadge() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/data_display"

var Badge = data_display.Badge`, nil),
		Playground("Basic Usage", "Standard implementation of the Badge component.", `Badge(text("New"))`, Badge(text("New")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
