package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderIcon() components.IsComponent {
	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/media_and_icons"

var Icon = media_and_icons.Icon`, nil),
		Playground("Basic Usage", "Standard implementation of the Icon component.", `Icon(fa_brand_android())`, Icon(fa_brand_android()).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
