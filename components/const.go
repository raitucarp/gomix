package components

type Dir string

const (
	DirLtr Dir = "ltr"
	DirRtl Dir = "rtl"
)

type TargetAction string

const (
	TargetActionHide   TargetAction = "hide"
	TargetActionShow   TargetAction = "show"
	TargetActionToggle TargetAction = "toggle"
)

type ButtonType string

const (
	ButtonTypeButton ButtonType = "button"
	ButtonTypeReset  ButtonType = "reset"
	ButtonTypeSubmit ButtonType = "submit"
)
