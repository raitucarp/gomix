package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	buttonsPkg "github.com/raitucarp/gomix/components/buttons"
)

func RenderIconButtonDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default icon_button component.", `buttonsPkg.IconButton(text("X"))`, buttonsPkg.IconButton(text("X")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
