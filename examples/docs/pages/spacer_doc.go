package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	layoutPkg "github.com/raitucarp/gomix/components/layout"
)

func RenderSpacerDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default spacer component.", `HStack(div(text("Left")), layoutPkg.Spacer().Component(), div(text("Right"))).Component().WFull()`, HStack(div(text("Left")), layoutPkg.Spacer().Component(), div(text("Right"))).Component().WFull()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
