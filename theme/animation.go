package theme

import "github.com/raitucarp/gomix/value"

func EaseIn(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Ease, "in", val.Value())
	}
}

func EaseOut(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Ease, "in", val.Value())
	}
}

func EaseInOut(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Ease, "in-out", val.Value())
	}
}

func EaseVar(variant string, val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Ease, variant, val.Value())
	}
}

func AnimateSpin(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Animate, "spin", val.Value())
	}
}

func AnimatePing(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Animate, "ping", val.Value())
	}
}

func AnimatePulse(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Animate, "pulse", val.Value())
	}
}

func AnimateBounce(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Animate, "bounce", val.Value())
	}
}

func AnimateVar(variant string, val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Animate, variant, val.Value())
	}
}
