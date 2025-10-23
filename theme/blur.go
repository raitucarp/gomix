package theme

import (
	"github.com/raitucarp/gomix/value"
)

func BlurXs(blurSize value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Blur, "xs", blurSize.Value())
	}
}

func BlurSm(blurSize value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Blur, "sm", blurSize.Value())
	}
}

func BlurMd(blurSize value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Blur, "md", blurSize.Value())
	}
}

func BlurLg(blurSize value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Blur, "lg", blurSize.Value())
	}
}

func BlurXl(blurSize value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Blur, "xl", blurSize.Value())
	}
}

func Blur2xl(blurSize value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Blur, "2xl", blurSize.Value())
	}
}

func Blur3xl(blurSize value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Blur, "3xl", blurSize.Value())
	}
}

func BlurVar(variant string, blurSize value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Blur, variant, blurSize.Value())
	}
}
