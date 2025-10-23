package theme

import (
	"github.com/raitucarp/gomix/value"
)

func PerspectiveDramatic(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Perspective, "dramatic", val.Value())
	}
}

func PerspectiveNear(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Perspective, "near", val.Value())
	}
}

func PerspectiveNormal(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Perspective, "normal", val.Value())
	}
}

func PerspectiveMidrange(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Perspective, "midrange", val.Value())
	}
}

func PerspectiveDistant(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Perspective, "distant", val.Value())
	}
}

func PerspectiveVar(variant string, val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Perspective, variant, val.Value())
	}
}
