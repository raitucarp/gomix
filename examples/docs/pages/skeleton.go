package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	feedbackPkg "github.com/raitucarp/gomix/components/feedback"
)

func RenderSkeleton() components.IsComponent {
	// Alias simulation for documentation
	var Skeleton = feedbackPkg.Skeleton

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/feedback"

var Skeleton = feedback.Skeleton`, nil),
		Playground("Basic Usage", "Standard implementation of the Skeleton component.", `Skeleton().Component()`, Skeleton().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
