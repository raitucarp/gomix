package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	feedbackPkg "github.com/raitucarp/gomix/components/feedback"
)

func RenderAlertDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default alert component.", `feedbackPkg.Alert()`, feedbackPkg.Alert().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
