package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	layoutPkg "github.com/raitucarp/gomix/components/layout"
)

func RenderCenterDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default center component.", `layoutPkg.Center(text("Center"))`, layoutPkg.Center(text("Center")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
