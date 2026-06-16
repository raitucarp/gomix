package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderSliderDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default slider component.", `formsPkg.Slider()`, formsPkg.Slider().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
