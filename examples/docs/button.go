package main

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
)

func RenderButton() components.IsComponent {
	return VStack(
		// Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/buttons"

var Button = buttons.Button`, nil),
		Playground("Basic Usage", "Standard implementation of the Button component.", `Button(text("Button")).ColorScheme("blue")`, Button(text("Button")).ColorScheme("blue").Component()),
		Playground("Prop: Variant", "Demonstrating the Variant method.", `HStack(
	Button(text("Button")).Variant("solid"),
	Button(text("Button")).Variant("outline"),
	Button(text("Button")).Variant("ghost")
)`, HStack(Button(text("Button")).Variant("solid").Component(), Button(text("Button")).Variant("outline").Component(), Button(text("Button")).Variant("ghost").Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		Playground("Prop: Size", "Demonstrating the Size method.", `HStack(
	Button(text("Button")).Size("sm"),
	Button(text("Button")).Size("md"),
	Button(text("Button")).Size("lg")
)`, HStack(Button(text("Button")).Size("sm").Component(), Button(text("Button")).Size("md").Component(), Button(text("Button")).Size("lg").Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		Playground("Prop: ColorScheme", "Demonstrating the ColorScheme method.", `HStack(
	Button(text("Button")).ColorScheme("blue"),
	Button(text("Button")).ColorScheme("red"),
	Button(text("Button")).ColorScheme("green")
)`, HStack(Button(text("Button")).ColorScheme("blue").Component(), Button(text("Button")).ColorScheme("red").Component(), Button(text("Button")).ColorScheme("green").Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		Playground("Prop: IsLoading", "Demonstrating the IsLoading method.", `HStack(
	Button(text("Button")).IsLoading(true),
	Button(text("Button")).IsLoading(false)
)`, HStack(Button(text("Button")).IsLoading(true).Component(), Button(text("Button")).IsLoading(false).Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		PropsTable([]map[string]string{
			{"name": "Variant", "type": "v string", "description": "Chainable method for Variant setting Gomix attributes/styles."},
			{"name": "Size", "type": "s string", "description": "Chainable method for Size setting Gomix attributes/styles."},
			{"name": "ColorScheme", "type": "color string", "description": "Chainable method for ColorScheme setting Gomix attributes/styles."},
			{"name": "IsLoading", "type": "loading bool", "description": "Chainable method for IsLoading setting Gomix attributes/styles."},
		}),
	).Component().GapBy(value.Rem(2))
}
