package theme

import "github.com/raitucarp/gomix/value"

func RadiusXs(radius value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Radius, "xs", radius.Value())
	}
}

func RadiusSm(radius value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Radius, "sm", radius.Value())
	}
}

func RadiusMd(radius value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Radius, "md", radius.Value())
	}
}

func RadiusLg(radius value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Radius, "lg", radius.Value())
	}
}

func RadiusXl(radius value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Radius, "xl", radius.Value())
	}
}

func Radius2xl(radius value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Radius, "2xl", radius.Value())
	}
}

func Radius3xl(radius value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Radius, "3xl", radius.Value())
	}
}

func Radius4xl(radius value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Radius, "4xl", radius.Value())
	}
}

func RadiusVar(variant string, radius value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Radius, variant, radius.Value())
	}
}
