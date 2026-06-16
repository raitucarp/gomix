package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderSlider() components.IsComponent {
	// Alias simulation for documentation
	var Slider = formsPkg.Slider

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/forms"

var Slider = forms.Slider`, nil),
		Playground("Basic Usage", "Standard implementation of the Slider component.", `Slider()`, Slider().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
