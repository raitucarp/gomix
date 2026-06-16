package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	media_and_iconsPkg "github.com/raitucarp/gomix/components/media_and_icons"
)

func RenderAvatarDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default avatar component.", `media_and_iconsPkg.Avatar()`, media_and_iconsPkg.Avatar().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
