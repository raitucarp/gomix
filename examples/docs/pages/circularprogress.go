package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	feedbackPkg "github.com/raitucarp/gomix/components/feedback"
)

func RenderCircularProgress() components.IsComponent {
	// Alias simulation for documentation
	var CircularProgress = feedbackPkg.CircularProgress

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/feedback"

var CircularProgress = feedback.CircularProgress`, nil),
		Playground("Basic Usage", "Standard implementation of the CircularProgress component.", `CircularProgress()`, CircularProgress().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
