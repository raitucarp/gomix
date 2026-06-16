package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderSimpleGrid() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/layout"

var SimpleGrid = layout.SimpleGrid`, nil),
		Playground("Basic Usage", "Standard implementation of the SimpleGrid component.", `SimpleGrid()`, SimpleGrid().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
