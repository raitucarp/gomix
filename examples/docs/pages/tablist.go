package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	disclosurePkg "github.com/raitucarp/gomix/components/disclosure"
)

func RenderTabList() components.IsComponent {
	// Alias simulation for documentation
	var TabList = disclosurePkg.TabList

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/disclosure"

var TabList = disclosure.TabList`, nil),
		Playground("Basic Usage", "Standard implementation of the TabList component.", `TabList()`, TabList().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
