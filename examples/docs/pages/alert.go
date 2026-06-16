package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	feedbackPkg "github.com/raitucarp/gomix/components/feedback"
)

func RenderAlert() components.IsComponent {
	// Alias simulation for documentation
	var Alert = feedbackPkg.Alert

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/feedback"

var Alert = feedback.Alert`, nil),
		Playground("Basic Usage", "Standard implementation of the Alert component.", `Alert()`, Alert().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
