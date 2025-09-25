package themes

var DefaultVariables = map[string]*Variables{
	"FontSans": NewVariable(Font, "sans", `ui-sans-serif, system-ui, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"`),


}

/*
DefaultTheme = Theme(
	fonts(
		sans(),
		serif(),
		mono(),
	), 

	colors(
		redColor()
	)
)
*/

// --font-serif: ui-serif, Georgia, Cambria, "Times New Roman", Times, serif;
// --font-mono: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;

func Default() *Theme {
	theme := NewTheme("default")

	// for _, v := range DefaultVariables {
	// 	for namespace, utility := range *v {

	// 		theme.AddVariable(namespace)
	// 	}
		
	// }
	// variables := Variables{
	// 	Animate: make(UtilityClass),

	// }
	// a := variables[Animate]
	// theme.AddVariable(Animate)

	return theme
}
