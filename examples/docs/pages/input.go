package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	formsPkg "github.com/raitucarp/gomix/components/forms"
)

func RenderInput() components.IsComponent {
	// Alias simulation for documentation
	var Input = formsPkg.Input

	return VStack(
		Playground("Import & Setup", "Define an alias to use the component cleanly throughout your project.", `import "github.com/raitucarp/gomix/components/forms"

var Input = forms.Input`, nil),
		Playground("Basic Usage", "Standard implementation of the Input component.", `Input()`, Input().Component()),
		Playground("Prop: Variant", "Demonstrating the Variant method.", `HStack(
	Input().Variant("solid"),
	Input().Variant("outline"),
	Input().Variant("ghost")
)`, HStack(Input().Variant("solid").Component(), Input().Variant("outline").Component(), Input().Variant("ghost").Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		Playground("Prop: Size", "Demonstrating the Size method.", `HStack(
	Input().Size("sm"),
	Input().Size("md"),
	Input().Size("lg")
)`, HStack(Input().Size("sm").Component(), Input().Size("md").Component(), Input().Size("lg").Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		Playground("Prop: IsInvalid", "Demonstrating the IsInvalid method.", `HStack(
	Input().IsInvalid(true),
	Input().IsInvalid(false)
)`, HStack(Input().IsInvalid(true).Component(), Input().IsInvalid(false).Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		Playground("Prop: IsRequired", "Demonstrating the IsRequired method.", `HStack(
	Input().IsRequired(true),
	Input().IsRequired(false)
)`, HStack(Input().IsRequired(true).Component(), Input().IsRequired(false).Component()).Component().GapBy(value.Rem(1)).FlexWrap()),
		PropsTable([]map[string]string{
			{"name": "Variant", "type": "v string", "description": "Chainable method for Variant setting Gomix attributes/styles."},
			{"name": "Size", "type": "s string", "description": "Chainable method for Size setting Gomix attributes/styles."},
			{"name": "IsInvalid", "type": "invalid bool", "description": "Chainable method for IsInvalid setting Gomix attributes/styles."},
			{"name": "IsRequired", "type": "req bool", "description": "Chainable method for IsRequired setting Gomix attributes/styles."},
		}),
	).Component().GapBy(value.Rem(2))
}
