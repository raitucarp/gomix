package theme

import (
	"github.com/raitucarp/gomix/value"
)

func TextXs(size value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "xs", size.Value())
	}
}

func LineHeightXs(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "xs--line-height", val.Value())
	}
}

func TextSm(size value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "sm", size.Value())
	}
}

func LineHeightSm(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "sm--line-height", val.Value())
	}
}

func TextBase(size value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "sm", size.Value())
	}
}

func LineHeightBase(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "sm--line-height", val.Value())
	}
}

func TextLg(size value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "lg", size.Value())
	}
}

func LineHeightLg(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "lg--line-height", val.Value())
	}
}

func TextXl(size value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "xl", size.Value())
	}
}

func LineHeightXl(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "xl--line-height", val.Value())
	}
}

func Text2xl(size value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "2xl", size.Value())
	}
}

func LineHeight2xl(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "2xl--line-height", val.Value())
	}
}

func Text3xl(size value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "3xl", size.Value())
	}
}

func LineHeight3xl(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "3xl--line-height", val.Value())
	}
}

func Text4xl(size value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "4xl", size.Value())
	}
}

func LineHeight4xl(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "4xl--line-height", val.Value())
	}
}

func Text5xl(size value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "5xl", size.Value())
	}
}

func LineHeight5xl(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "5xl--line-height", val.Value())
	}
}

func Text6xl(size value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "6xl", size.Value())
	}
}

func LineHeight6xl(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "6xl--line-height", val.Value())
	}
}

func Text7xl(size value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "7xl", size.Value())
	}
}

func LineHeight7xl(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "7xl--line-height", val.Value())
	}
}

func Text8xl(size value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "8xl", size.Value())
	}
}

func LineHeight8xl(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "8xl--line-height", val.Value())
	}
}

func Text9xl(size value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "9xl", size.Value())
	}
}

func LineHeight9xl(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Text, "9xl--line-height", val.Value())
	}
}

func TrackingTighter(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Tracking, "tighter", val.Value())
	}
}

func TrackingTight(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Tracking, "tight", val.Value())
	}
}

func TrackingNormal(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Tracking, "normal", val.Value())
	}
}

func TrackingWide(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Tracking, "wide", val.Value())
	}
}

func TrackingWider(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Tracking, "wider", val.Value())
	}
}

func TrackingWidest(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Tracking, "widest", val.Value())
	}
}

func LeadingTight(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Leading, "tight", val.Value())
	}
}

func LeadingSnug(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Leading, "snug", val.Value())
	}
}

func LeadingNormal(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Leading, "normal", val.Value())
	}
}

func LeadingRelaxed(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Leading, "relaxed", val.Value())
	}
}

func LeadingLoose(val value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Leading, "loose", val.Value())
	}
}
