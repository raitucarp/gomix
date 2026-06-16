package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderHeading() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/typography"

var Heading = typography.Heading`, nil),
		Playground("Basic Usage", "Standard implementation of the Heading component.", `Heading(text("Heading"))`, Heading(text("Heading")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
