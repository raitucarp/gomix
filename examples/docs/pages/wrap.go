package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	layoutPkg "github.com/raitucarp/gomix/components/layout"
)

func RenderWrap() components.IsComponent {
	// Alias simulation for documentation
	var Wrap = layoutPkg.Wrap

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/layout"

var Wrap = layout.Wrap`, nil),
		Playground("Basic Usage", "Standard implementation of the Wrap component.", `Wrap()`, Wrap().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
