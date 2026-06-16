package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderSpacer() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/layout"

var Spacer = layout.Spacer`, nil),
		Playground("Basic Usage", "Standard implementation of the Spacer component.", `HStack(div(text("Left")), Spacer().Component(), div(text("Right"))).Component().WFull()`, HStack(div(text("Left")), Spacer().Component(), div(text("Right"))).Component().WFull()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
