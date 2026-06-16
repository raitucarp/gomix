package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	disclosurePkg "github.com/raitucarp/gomix/components/disclosure"
)

func RenderTabPanel() components.IsComponent {
	// Alias simulation for documentation
	var TabPanel = disclosurePkg.TabPanel

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/disclosure"

var TabPanel = disclosure.TabPanel`, nil),
		Playground("Basic Usage", "Standard implementation of the TabPanel component.", `TabPanel()`, TabPanel().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
