package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderTag() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/data_display"

var Tag = data_display.Tag`, nil),
		Playground("Basic Usage", "Standard implementation of the Tag component.", `Tag(text("Tag"))`, Tag(text("Tag")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
