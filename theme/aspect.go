package theme

import "github.com/raitucarp/gomix/value"

func AspectVideo(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Aspect, "video", val.Value())
	}
}

func AspectVar(variant string, val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Aspect, variant, val.Value())
	}
}
