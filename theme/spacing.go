package theme

import (
	"github.com/raitucarp/gomix/value"
)

func SpacingVar(spacing value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Spacing, "", spacing.Value())
	}
}

func BreakpointSm(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Breakpoint, "sm", breakpoint.Value())
	}
}

func BreakpointMd(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Breakpoint, "md", breakpoint.Value())
	}
}

func BreakpointLg(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Breakpoint, "lg", breakpoint.Value())
	}
}

func BreakpointXl(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Breakpoint, "xl", breakpoint.Value())
	}
}

func Breakpoint2xl(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Breakpoint, "2xl", breakpoint.Value())
	}
}

func Container3xs(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Container, "3xs", breakpoint.Value())
	}
}

func Container2xs(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Container, "2xs", breakpoint.Value())
	}
}

func ContainerXs(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Container, "xs", breakpoint.Value())
	}
}

func ContainerSm(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Container, "sm", breakpoint.Value())
	}
}

func ContainerMd(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Container, "md", breakpoint.Value())
	}
}

func ContainerLg(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Container, "lg", breakpoint.Value())
	}
}

func ContainerXl(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Container, "xl", breakpoint.Value())
	}
}

func Container2xl(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Container, "2xl", breakpoint.Value())
	}
}

func Container3xl(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Container, "3xl", breakpoint.Value())
	}
}

func Container4xl(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Container, "4xl", breakpoint.Value())
	}
}

func Container5xl(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Container, "5xl", breakpoint.Value())
	}
}

func Container6xl(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Container, "6xl", breakpoint.Value())
	}
}

func Container7xl(breakpoint value.Value) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Container, "7xl", breakpoint.Value())
	}
}
