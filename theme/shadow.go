package theme

func Shadow2xs(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Shadow, "2xs", value)
	}
}

func ShadowXs(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Shadow, "xs", value)
	}
}

func ShadowSm(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Shadow, "sm", value)
	}
}

func ShadowMd(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Shadow, "md", value)
	}
}

func ShadowLg(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Shadow, "lg", value)
	}
}

func ShadowXl(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Shadow, "xl", value)
	}
}

func Shadow2xl(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Shadow, "2xl", value)
	}
}

func ShadowVar(variant string, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Shadow, variant, value)
	}
}

func InsetShadow2xs(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(InsetShadow, "2xs", value)
	}
}

func InsetShadowXs(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(InsetShadow, "xs", value)
	}
}

func InsetShadowSm(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(InsetShadow, "sm", value)
	}
}

func InsetShadowVar(variant string, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(InsetShadow, variant, value)
	}
}

func DropShadowXs(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(DropShadow, "xs", value)
	}
}

func DropShadowSm(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(DropShadow, "sm", value)
	}
}

func DropShadowMd(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(DropShadow, "md", value)
	}
}

func DropShadowLg(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(DropShadow, "lg", value)
	}
}

func DropShadowXl(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(DropShadow, "xl", value)
	}
}

func DropShadow2xl(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(DropShadow, "2xl", value)
	}
}

func DropShadowVar(variant string, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(DropShadow, variant, value)
	}
}

func TextShadow2xs(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(TextShadow, "2xs", value)
	}
}

func TextShadowXs(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(TextShadow, "xs", value)
	}
}

func TextShadowSm(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(TextShadow, "sm", value)
	}
}

func TextShadowMd(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(TextShadow, "md", value)
	}
}

func TextShadowLg(value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(TextShadow, "lg", value)
	}
}

func TextShadowVar(variant string, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(TextShadow, variant, value)
	}
}
