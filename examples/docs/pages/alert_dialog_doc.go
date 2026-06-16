package pages

import (
    "github.com/raitucarp/gomix/components"
    "github.com/raitucarp/gomix/value"
	overlaysPkg "github.com/raitucarp/gomix/components/overlays"
)

func RenderAlertDialogDoc() components.IsComponent {
	return VStack(
		Playground("Basic Usage", "The default alert_dialog component.", `overlaysPkg.AlertDialog()`, overlaysPkg.AlertDialog().Component()),
		PropsTable([]map[string]string{
		}),
	).Component().GapBy(value.Rem(2))
}
