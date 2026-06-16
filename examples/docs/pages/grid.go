package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	layoutPkg "github.com/raitucarp/gomix/components/layout"
)

func RenderGrid() components.IsComponent {
	// Alias simulation for documentation
	var Grid = layoutPkg.Grid

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/layout"

var Grid = layout.Grid`, nil),
		Playground("Basic Usage", "Standard implementation of the Grid component.", `Grid()`, Grid().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
