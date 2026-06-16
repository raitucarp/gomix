package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderTooltip() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/overlays"

var Tooltip = overlays.Tooltip`, nil),
		Playground("Basic Usage", "Standard implementation of the Tooltip component.", `Tooltip()`, Tooltip().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
