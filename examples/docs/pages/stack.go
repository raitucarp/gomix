package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	layoutPkg "github.com/raitucarp/gomix/components/layout"
)

func RenderStack() components.IsComponent {
	// Alias simulation for documentation
	var Stack = layoutPkg.Stack

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/layout"

var Stack = layout.Stack`, nil),
		Playground("Basic Usage", "Standard implementation of the Stack component.", `Stack()`, Stack().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
