package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	layoutPkg "github.com/raitucarp/gomix/components/layout"
)

func RenderHStack() components.IsComponent {
	// Alias simulation for documentation
	var HStack = layoutPkg.HStack

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/layout"

var HStack = layout.HStack`, nil),
		Playground("Basic Usage", "Standard implementation of the HStack component.", `HStack()`, HStack().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
