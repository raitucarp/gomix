package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	feedbackPkg "github.com/raitucarp/gomix/components/feedback"
)

func RenderSpinner() components.IsComponent {
	// Alias simulation for documentation
	var Spinner = feedbackPkg.Spinner

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/feedback"

var Spinner = feedback.Spinner`, nil),
		Playground("Basic Usage", "Standard implementation of the Spinner component.", `Spinner()`, Spinner().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
