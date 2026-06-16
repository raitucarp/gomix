package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	feedbackPkg "github.com/raitucarp/gomix/components/feedback"
)

func RenderCircularProgressDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default circular_progress component.", `feedbackPkg.CircularProgress()`, feedbackPkg.CircularProgress().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
