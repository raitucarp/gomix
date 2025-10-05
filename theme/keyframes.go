package theme

func KeyframeSpin(value string) ThemeParam {
	return func(t *Theme) {
		t.AddKeyframe("spin", value)
	}
}

func KeyframePing(value string) ThemeParam {
	return func(t *Theme) {
		t.AddKeyframe("ping", value)
	}
}

func KeyframePulse(value string) ThemeParam {
	return func(t *Theme) {
		t.AddKeyframe("ping", value)
	}
}

func KeyframeBounce(value string) ThemeParam {
	return func(t *Theme) {
		t.AddKeyframe("bounce", value)
	}
}

func KeyframeVar(variant string, value string) ThemeParam {
	return func(t *Theme) {
		t.AddKeyframe(variant, value)
	}
}
