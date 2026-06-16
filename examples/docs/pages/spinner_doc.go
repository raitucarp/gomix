package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	feedbackPkg "github.com/raitucarp/gomix/components/feedback"
)

func RenderSpinnerDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default spinner component.", `feedbackPkg.Spinner()`, feedbackPkg.Spinner().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
