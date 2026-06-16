package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	feedbackPkg "github.com/raitucarp/gomix/components/feedback"
)

func RenderToastDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default toast component.", `feedbackPkg.Toast()`, feedbackPkg.Toast().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
