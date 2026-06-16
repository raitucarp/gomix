package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	media_and_iconsPkg "github.com/raitucarp/gomix/components/media_and_icons"
)

func RenderAvatar() components.IsComponent {
	// Alias simulation for documentation
	var Avatar = media_and_iconsPkg.Avatar

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/media_and_icons"

var Avatar = media_and_icons.Avatar`, nil),
		Playground("Basic Usage", "Standard implementation of the Avatar component.", `Avatar()`, Avatar().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
