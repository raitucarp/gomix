package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	media_and_iconsPkg "github.com/raitucarp/gomix/components/media_and_icons"
)

func RenderImageDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default image component.", `media_and_iconsPkg.Image()`, media_and_iconsPkg.Image().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
