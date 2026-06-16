package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	data_displayPkg "github.com/raitucarp/gomix/components/data_display"
)

func RenderTable() components.IsComponent {
	// Alias simulation for documentation
	var Table = data_displayPkg.Table

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/data_display"

var Table = data_display.Table`, nil),
		Playground("Basic Usage", "Standard implementation of the Table component.", `Table()`, Table().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
