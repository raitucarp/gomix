package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	media_and_iconsPkg "github.com/raitucarp/gomix/components/media_and_icons"
)

func RenderImage() components.IsComponent {
	// Alias simulation for documentation
	var Image = media_and_iconsPkg.Image

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/media_and_icons"

var Image = media_and_icons.Image`, nil),
		Playground("Basic Usage", "Standard implementation of the Image component.", `Image()`, Image().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
