package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	feedbackPkg "github.com/raitucarp/gomix/components/feedback"
)

func RenderProgress() components.IsComponent {
	// Alias simulation for documentation
	var Progress = feedbackPkg.Progress

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/feedback"

var Progress = feedback.Progress`, nil),
		Playground("Basic Usage", "Standard implementation of the Progress component.", `Progress()`, Progress().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
