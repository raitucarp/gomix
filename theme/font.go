package theme

import (
	"fmt"
	"strings"

)

func FontSans(fonts ...string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Font, "sans", strings.Join(fonts, ", "))
	}
}

func FontSerif(fonts ...string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Font, "serif", strings.Join(fonts, ", "))
	}
}

func FontMono(fonts ...string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Font, "mono", strings.Join(fonts, ", "))
	}
}

func FontVar(name string, fonts ...string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Font, name, strings.Join(fonts, ", "))
	}
}

func FontThin(weight int) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(FontWeight, "thin", fmt.Sprintf("%d", weight))
	}
}

func FontExtraLight(weight int) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(FontWeight, "extralight", fmt.Sprintf("%d", weight))
	}
}

func FontLight(weight int) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(FontWeight, "light", fmt.Sprintf("%d", weight))
	}
}

func FontNormal(weight int) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(FontWeight, "normal", fmt.Sprintf("%d", weight))
	}
}

func FontMedium(weight int) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(FontWeight, "medium", fmt.Sprintf("%d", weight))
	}
}

func FontSemibold(weight int) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(FontWeight, "semibold", fmt.Sprintf("%d", weight))
	}
}

func FontBold(weight int) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(FontWeight, "bold", fmt.Sprintf("%d", weight))
	}
}

func FontExtraBold(weight int) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(FontWeight, "extrabold", fmt.Sprintf("%d", weight))
	}
}

func FontBlack(weight int) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(FontWeight, "black", fmt.Sprintf("%d", weight))
	}
}
