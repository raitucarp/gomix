package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	feedbackPkg "github.com/raitucarp/gomix/components/feedback"
)

func RenderToast() components.IsComponent {
	// Alias simulation for documentation
	var Toast = feedbackPkg.Toast

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/feedback"

var Toast = feedback.Toast`, nil),
		Playground("Basic Usage", "Standard implementation of the Toast component.", `Toast()`, Toast().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
