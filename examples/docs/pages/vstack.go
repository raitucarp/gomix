package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	layoutPkg "github.com/raitucarp/gomix/components/layout"
)

func RenderVStack() components.IsComponent {
	// Alias simulation for documentation
	var VStack = layoutPkg.VStack

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/layout"

var VStack = layout.VStack`, nil),
		Playground("Basic Usage", "Standard implementation of the VStack component.", `VStack()`, VStack().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
