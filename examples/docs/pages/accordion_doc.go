package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	disclosurePkg "github.com/raitucarp/gomix/components/disclosure"
)

func RenderAccordionDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default accordion component.", `disclosurePkg.Accordion()`, disclosurePkg.Accordion().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
