package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	layoutPkg "github.com/raitucarp/gomix/components/layout"
)

func RenderCenter() components.IsComponent {
	// Alias simulation for documentation
	var Center = layoutPkg.Center

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/layout"

var Center = layout.Center`, nil),
		Playground("Basic Usage", "Standard implementation of the Center component.", `Center(text("Center"))`, Center(text("Center")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
