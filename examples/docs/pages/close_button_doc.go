package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	buttonsPkg "github.com/raitucarp/gomix/components/buttons"
)

func RenderCloseButtonDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default close_button component.", `buttonsPkg.CloseButton()`, buttonsPkg.CloseButton().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
