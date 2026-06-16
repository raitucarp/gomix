package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	feedbackPkg "github.com/raitucarp/gomix/components/feedback"
)

func RenderProgressDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default progress component.", `feedbackPkg.Progress()`, feedbackPkg.Progress().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
