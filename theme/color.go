package theme

import "fmt"

func ColorGray(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("gray-%d", scale), value)
	}
}
func ColorGrayAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("gray-alpha-%d", scale), value)
	}
}
func ColorGrayDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("gray-dark-%d", scale), value)
	}
}
func ColorGrayDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("gray-dark-alpha-%d", scale), value)
	}
}
func ColorMauve(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("mauve-%d", scale), value)
	}
}
func ColorMauveAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("mauve-alpha-%d", scale), value)
	}
}
func ColorMauveDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("mauve-dark-%d", scale), value)
	}
}
func ColorMauveDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("mauve-dark-alpha-%d", scale), value)
	}
}
func ColorSlate(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("slate-%d", scale), value)
	}
}
func ColorSlateAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("slate-alpha-%d", scale), value)
	}
}
func ColorSlateDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("slate-dark-%d", scale), value)
	}
}
func ColorSlateDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("slate-dark-alpha-%d", scale), value)
	}
}
func ColorSage(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("sage-%d", scale), value)
	}
}
func ColorSageAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("sage-alpha-%d", scale), value)
	}
}
func ColorSageDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("sage-dark-%d", scale), value)
	}
}
func ColorSageDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("sage-dark-alpha-%d", scale), value)
	}
}
func ColorOlive(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("olive-%d", scale), value)
	}
}
func ColorOliveAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("olive-alpha-%d", scale), value)
	}
}
func ColorOliveDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("olive-dark-%d", scale), value)
	}
}
func ColorOliveDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("olive-dark-alpha-%d", scale), value)
	}
}
func ColorSand(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("sand-%d", scale), value)
	}
}
func ColorSandAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("sand-alpha-%d", scale), value)
	}
}
func ColorSandDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("sand-dark-%d", scale), value)
	}
}
func ColorSandDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("sand-dark-alpha-%d", scale), value)
	}
}
func ColorGold(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("gold-%d", scale), value)
	}
}
func ColorGoldAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("gold-alpha-%d", scale), value)
	}
}
func ColorGoldDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("gold-dark-%d", scale), value)
	}
}
func ColorGoldDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("gold-dark-alpha-%d", scale), value)
	}
}
func ColorBronze(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("bronze-%d", scale), value)
	}
}
func ColorBronzeAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("bronze-alpha-%d", scale), value)
	}
}
func ColorBronzeDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("bronze-dark-%d", scale), value)
	}
}
func ColorBronzeDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("bronze-dark-alpha-%d", scale), value)
	}
}
func ColorBrown(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("brown-%d", scale), value)
	}
}
func ColorBrownAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("brown-alpha-%d", scale), value)
	}
}
func ColorBrownDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("brown-dark-%d", scale), value)
	}
}
func ColorBrownDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("brown-dark-alpha-%d", scale), value)
	}
}
func ColorYellow(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("yellow-%d", scale), value)
	}
}
func ColorYellowAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("yellow-alpha-%d", scale), value)
	}
}
func ColorYellowDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("yellow-dark-%d", scale), value)
	}
}
func ColorYellowDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("yellow-dark-alpha-%d", scale), value)
	}
}
func ColorAmber(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("amber-%d", scale), value)
	}
}
func ColorAmberAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("amber-alpha-%d", scale), value)
	}
}
func ColorAmberDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("amber-dark-%d", scale), value)
	}
}
func ColorAmberDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("amber-dark-alpha-%d", scale), value)
	}
}
func ColorOrange(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("orange-%d", scale), value)
	}
}
func ColorOrangeAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("orange-alpha-%d", scale), value)
	}
}
func ColorOrangeDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("orange-dark-%d", scale), value)
	}
}
func ColorOrangeDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("orange-dark-alpha-%d", scale), value)
	}
}
func ColorTomato(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("tomato-%d", scale), value)
	}
}
func ColorTomatoAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("tomato-alpha-%d", scale), value)
	}
}
func ColorTomatoDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("tomato-dark-%d", scale), value)
	}
}
func ColorTomatoDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("tomato-dark-alpha-%d", scale), value)
	}
}
func ColorRed(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("red-%d", scale), value)
	}
}
func ColorRedAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("red-alpha-%d", scale), value)
	}
}
func ColorRedDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("red-dark-%d", scale), value)
	}
}
func ColorRedDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("red-dark-alpha-%d", scale), value)
	}
}
func ColorRuby(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("ruby-%d", scale), value)
	}
}
func ColorRubyAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("ruby-alpha-%d", scale), value)
	}
}
func ColorRubyDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("ruby-dark-%d", scale), value)
	}
}
func ColorRubyDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("ruby-dark-alpha-%d", scale), value)
	}
}
func ColorCrimson(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("crimson-%d", scale), value)
	}
}
func ColorCrimsonAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("crimson-alpha-%d", scale), value)
	}
}
func ColorCrimsonDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("crimson-dark-%d", scale), value)
	}
}
func ColorCrimsonDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("crimson-dark-alpha-%d", scale), value)
	}
}
func ColorPink(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("pink-%d", scale), value)
	}
}
func ColorPinkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("pink-alpha-%d", scale), value)
	}
}
func ColorPinkDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("pink-dark-%d", scale), value)
	}
}
func ColorPinkDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("pink-dark-alpha-%d", scale), value)
	}
}
func ColorPlum(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("plum-%d", scale), value)
	}
}
func ColorPlumAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("plum-alpha-%d", scale), value)
	}
}
func ColorPlumDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("plum-dark-%d", scale), value)
	}
}
func ColorPlumDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("plum-dark-alpha-%d", scale), value)
	}
}
func ColorPurple(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("purple-%d", scale), value)
	}
}
func ColorPurpleAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("purple-alpha-%d", scale), value)
	}
}
func ColorPurpleDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("purple-dark-%d", scale), value)
	}
}
func ColorPurpleDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("purple-dark-alpha-%d", scale), value)
	}
}
func ColorViolet(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("violet-%d", scale), value)
	}
}
func ColorVioletAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("violet-alpha-%d", scale), value)
	}
}
func ColorVioletDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("violet-dark-%d", scale), value)
	}
}
func ColorVioletDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("violet-dark-alpha-%d", scale), value)
	}
}
func ColorIris(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("iris-%d", scale), value)
	}
}
func ColorIrisAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("iris-alpha-%d", scale), value)
	}
}
func ColorIrisDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("iris-dark-%d", scale), value)
	}
}
func ColorIrisDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("iris-dark-alpha-%d", scale), value)
	}
}
func ColorIndigo(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("indigo-%d", scale), value)
	}
}
func ColorIndigoAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("indigo-alpha-%d", scale), value)
	}
}
func ColorIndigoDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("indigo-dark-%d", scale), value)
	}
}
func ColorIndigoDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("indigo-dark-alpha-%d", scale), value)
	}
}
func ColorBlue(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("blue-%d", scale), value)
	}
}
func ColorBlueAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("blue-alpha-%d", scale), value)
	}
}
func ColorBlueDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("blue-dark-%d", scale), value)
	}
}
func ColorBlueDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("blue-dark-alpha-%d", scale), value)
	}
}
func ColorCyan(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("cyan-%d", scale), value)
	}
}
func ColorCyanAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("cyan-alpha-%d", scale), value)
	}
}
func ColorCyanDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("cyan-dark-%d", scale), value)
	}
}
func ColorCyanDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("cyan-dark-alpha-%d", scale), value)
	}
}
func ColorTeal(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("teal-%d", scale), value)
	}
}
func ColorTealAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("teal-alpha-%d", scale), value)
	}
}
func ColorTealDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("teal-dark-%d", scale), value)
	}
}
func ColorTealDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("teal-dark-alpha-%d", scale), value)
	}
}
func ColorJade(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("jade-%d", scale), value)
	}
}
func ColorJadeAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("jade-alpha-%d", scale), value)
	}
}
func ColorJadeDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("jade-dark-%d", scale), value)
	}
}
func ColorJadeDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("jade-dark-alpha-%d", scale), value)
	}
}
func ColorGreen(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("green-%d", scale), value)
	}
}
func ColorGreenAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("green-alpha-%d", scale), value)
	}
}
func ColorGreenDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("green-dark-%d", scale), value)
	}
}
func ColorGreenDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("green-dark-alpha-%d", scale), value)
	}
}
func ColorGrass(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("grass-%d", scale), value)
	}
}
func ColorGrassAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("grass-alpha-%d", scale), value)
	}
}
func ColorGrassDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("grass-dark-%d", scale), value)
	}
}
func ColorGrassDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("grass-dark-alpha-%d", scale), value)
	}
}
func ColorLime(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("lime-%d", scale), value)
	}
}
func ColorLimeAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("lime-alpha-%d", scale), value)
	}
}
func ColorLimeDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("lime-dark-%d", scale), value)
	}
}
func ColorLimeDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("lime-dark-alpha-%d", scale), value)
	}
}
func ColorMint(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("mint-%d", scale), value)
	}
}
func ColorMintAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("mint-alpha-%d", scale), value)
	}
}
func ColorMintDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("mint-dark-%d", scale), value)
	}
}
func ColorMintDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("mint-dark-alpha-%d", scale), value)
	}
}
func ColorSky(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("sky-%d", scale), value)
	}
}
func ColorSkyAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("sky-alpha-%d", scale), value)
	}
}
func ColorSkyDark(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("sky-dark-%d", scale), value)
	}
}
func ColorSkyDarkAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("sky-dark-alpha-%d", scale), value)
	}
}
func ColorBlackAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("black-alpha-%d", scale), value)
	}
}
func ColorWhiteAlpha(scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("white-alpha-%d", scale), value)
	}
}

func ColorBlack() ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, "black", "#000")
	}
}
func ColorWhite() ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, "white", "#fff")
	}
}

func ColorVar(name string, scale int, value string) ThemeParam {
	return func(t *Theme) {
		t.AddVariable(Color, fmt.Sprintf("%s-%d", name, scale), value)
	}
}
