package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	buttonsPkg "github.com/raitucarp/gomix/components/buttons"
)

func RenderButtonDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default button component.", `buttonsPkg.Button(text("Click me"))`, buttonsPkg.Button(text("Click me")).Component()),
		Playground("Variants", "Different visual variants.", `HStack(
	buttonsPkg.Button(text("Solid")).Variant("solid").Component(),
	buttonsPkg.Button(text("Outline")).Variant("outline").Component(),
	buttonsPkg.Button(text("Ghost")).Variant("ghost").Component(),
	buttonsPkg.Button(text("Link")).Variant("link").Component()
)`, HStack(buttonsPkg.Button(text("Solid")).Variant("solid").Component(), buttonsPkg.Button(text("Outline")).Variant("outline").Component(), buttonsPkg.Button(text("Ghost")).Variant("ghost").Component(), buttonsPkg.Button(text("Link")).Variant("link").Component()).Component().GapBy(value.Rem(1))),
		PropsTable([]map[string]string{
			{"name": "Variant", "type": "v string", "description": "Chainable method for Variant."},
			{"name": "Size", "type": "s string", "description": "Chainable method for Size."},
			{"name": "ColorScheme", "type": "color string", "description": "Chainable method for ColorScheme."},
			{"name": "IsLoading", "type": "loading bool", "description": "Chainable method for IsLoading."},
		}),
	).Component().GapBy(value.Rem(2))
}
