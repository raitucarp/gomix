package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	disclosurePkg "github.com/raitucarp/gomix/components/disclosure"
)

func RenderVisuallyHidden() components.IsComponent {
	// Alias simulation for documentation
	var VisuallyHidden = disclosurePkg.VisuallyHidden

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/disclosure"

var VisuallyHidden = disclosure.VisuallyHidden`, nil),
		Playground("Basic Usage", "Standard implementation of the VisuallyHidden component.", `VisuallyHidden(text("Hidden"))`, VisuallyHidden(text("Hidden")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
