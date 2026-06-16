package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderInputDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default input component.", `formsPkg.Input()`, formsPkg.Input().Component()),
		Playground("Variants", "Different visual variants.", `HStack(
	formsPkg.Input().Variant("outline").Component(),
	formsPkg.Input().Variant("filled").Component(),
	formsPkg.Input().Variant("flushed").Component()
)`, HStack(formsPkg.Input().Variant("outline").Component(), formsPkg.Input().Variant("filled").Component(), formsPkg.Input().Variant("flushed").Component()).Component().GapBy(value.Rem(1))),
		PropsTable([]map[string]string{
			{"name": "Variant", "type": "v string", "description": "Chainable method for Variant."},
			{"name": "Size", "type": "s string", "description": "Chainable method for Size."},
			{"name": "IsInvalid", "type": "invalid bool", "description": "Chainable method for IsInvalid."},
			{"name": "IsRequired", "type": "req bool", "description": "Chainable method for IsRequired."},
		}),
	).Component().GapBy(value.Rem(2))
}
