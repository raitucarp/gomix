package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	data_displayPkg "github.com/raitucarp/gomix/components/data_display"
)

func RenderStat() components.IsComponent {
	// Alias simulation for documentation
	var Stat = data_displayPkg.Stat

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/data_display"

var Stat = data_display.Stat`, nil),
		Playground("Basic Usage", "Standard implementation of the Stat component.", `Stat()`, Stat().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
