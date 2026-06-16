package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	media_and_iconsPkg "github.com/raitucarp/gomix/components/media_and_icons"
	"github.com/raitucarp/gomix/icons/fontawesome-5"
)

func RenderIcon() components.IsComponent {
	// Alias simulation for documentation
	var Icon = media_and_iconsPkg.Icon

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/media_and_icons"

var Icon = media_and_icons.Icon`, nil),
		Playground("Basic Usage", "Standard implementation of the Icon component.", `Icon(fa5.FaBrand500px())`, Icon(fa5.FaBrand500px()).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
