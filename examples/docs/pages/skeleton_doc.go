package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	feedbackPkg "github.com/raitucarp/gomix/components/feedback"
)

func RenderSkeletonDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default skeleton component.", `feedbackPkg.Skeleton().Component()`, feedbackPkg.Skeleton().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
