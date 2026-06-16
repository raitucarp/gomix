package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderCircularProgress() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/feedback"

var CircularProgress = feedback.CircularProgress`, nil),
		Playground("Basic Usage", "Standard implementation of the CircularProgress component.", `CircularProgress()`, CircularProgress().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
