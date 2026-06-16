package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	disclosurePkg "github.com/raitucarp/gomix/components/disclosure"
)

func RenderVisuallyHiddenDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default visually_hidden component.", `disclosurePkg.VisuallyHidden(text("Hidden"))`, disclosurePkg.VisuallyHidden(text("Hidden")).Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
