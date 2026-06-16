package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	data_displayPkg "github.com/raitucarp/gomix/components/data_display"
)

func RenderDivider() components.IsComponent {
	// Alias simulation for documentation
	var Divider = data_displayPkg.Divider

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/data_display"

var Divider = data_display.Divider`, nil),
		Playground("Basic Usage", "Standard implementation of the Divider component.", `Divider()`, Divider().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
