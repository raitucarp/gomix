package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	media_and_iconsPkg "github.com/raitucarp/gomix/components/media_and_icons"
	"github.com/raitucarp/gomix/icons/fontawesome-5"
)

func RenderIconDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default icon component.", `media_and_iconsPkg.Icon(fa5.FaBrand500px())`, media_and_iconsPkg.Icon(fa5.FaBrand500px()).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
