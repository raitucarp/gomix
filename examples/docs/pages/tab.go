package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	disclosurePkg "github.com/raitucarp/gomix/components/disclosure"
)

func RenderTab() components.IsComponent {
	// Alias simulation for documentation
	var Tab = disclosurePkg.Tab

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/disclosure"

var Tab = disclosure.Tab`, nil),
		Playground("Basic Usage", "Standard implementation of the Tab component.", `Tab()`, Tab().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
