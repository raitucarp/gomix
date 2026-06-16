package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	layoutPkg "github.com/raitucarp/gomix/components/layout"
)

func RenderSpacer() components.IsComponent {
	// Alias simulation for documentation
	var Spacer = layoutPkg.Spacer

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/layout"

var Spacer = layout.Spacer`, nil),
		Playground("Basic Usage", "Standard implementation of the Spacer component.", `HStack(div(text("Left")), Spacer().Component(), div(text("Right"))).Component().WFull()`, HStack(div(text("Left")), Spacer().Component(), div(text("Right"))).Component().WFull()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
