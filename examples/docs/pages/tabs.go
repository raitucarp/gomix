package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	disclosurePkg "github.com/raitucarp/gomix/components/disclosure"
)

func RenderTabs() components.IsComponent {
	// Alias simulation for documentation
	var Tabs = disclosurePkg.Tabs

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/disclosure"

var Tabs = disclosure.Tabs`, nil),
		Playground("Basic Usage", "Standard implementation of the Tabs component.", `Tabs()`, Tabs().Component()),
		PropsTable([]map[string]string{
			{"name": "DefaultIndex", "type": "index int", "description": "Chainable method for DefaultIndex setting Gomix attributes/styles."},
		}),
	).Component().GapBy(value.Rem(2))
}
