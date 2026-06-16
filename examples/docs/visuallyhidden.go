package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderVisuallyHidden() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/disclosure"

var VisuallyHidden = disclosure.VisuallyHidden`, nil),
		Playground("Basic Usage", "Standard implementation of the VisuallyHidden component.", `VisuallyHidden(text("Hidden"))`, VisuallyHidden(text("Hidden")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
