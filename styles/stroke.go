package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func StrokeInherit() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): "inherit",
		}
	}
}

func StrokeCurrent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): "currentColor",
		}
	}
}

func StrokeTransparent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): "transparent",
		}
	}
}

func StrokeBlack() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, "black"),
		}
	}
}

func StrokeWhite() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, "white"),
		}
	}
}

func StrokeAmber(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func StrokeAmberAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func StrokeAmberDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func StrokeAmberDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func StrokeBlackAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func StrokeBlue(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func StrokeBlueAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func StrokeBlueDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func StrokeBlueDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func StrokeBronze(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func StrokeBronzeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func StrokeBronzeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func StrokeBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func StrokeBrown(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func StrokeBrownAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func StrokeBrownDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func StrokeBrownDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func StrokeCrimson(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func StrokeCrimsonAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func StrokeCrimsonDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func StrokeCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func StrokeCyan(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func StrokeCyanAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func StrokeCyanDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func StrokeCyanDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func StrokeGold(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func StrokeGoldAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func StrokeGoldDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func StrokeGoldDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func StrokeGrass(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func StrokeGrassAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func StrokeGrassDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func StrokeGrassDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func StrokeGray(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func StrokeGrayAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func StrokeGrayDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func StrokeGrayDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func StrokeGreen(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func StrokeGreenAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func StrokeGreenDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func StrokeGreenDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func StrokeIndigo(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func StrokeIndigoAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func StrokeIndigoDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func StrokeIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func StrokeIris(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func StrokeIrisAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func StrokeIrisDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func StrokeIrisDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func StrokeJade(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func StrokeJadeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func StrokeJadeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func StrokeJadeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func StrokeLime(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func StrokeLimeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func StrokeLimeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func StrokeLimeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func StrokeMauve(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func StrokeMauveAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func StrokeMauveDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func StrokeMauveDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func StrokeMint(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func StrokeMintAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func StrokeMintDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func StrokeMintDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func StrokeOlive(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func StrokeOliveAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func StrokeOliveDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func StrokeOliveDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func StrokeOrange(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func StrokeOrangeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func StrokeOrangeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func StrokeOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func StrokePink(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func StrokePinkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func StrokePinkDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func StrokePinkDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func StrokePlum(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func StrokePlumAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func StrokePlumDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func StrokePlumDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func StrokePurple(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func StrokePurpleAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func StrokePurpleDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func StrokePurpleDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func StrokeRed(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func StrokeRedAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func StrokeRedDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func StrokeRedDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func StrokeRuby(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func StrokeRubyAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func StrokeRubyDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func StrokeRubyDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func StrokeSage(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func StrokeSageAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func StrokeSageDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func StrokeSageDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func StrokeSand(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func StrokeSandAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func StrokeSandDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func StrokeSandDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func StrokeSky(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func StrokeSkyAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func StrokeSkyDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func StrokeSkyDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func StrokeSlate(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func StrokeSlateAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func StrokeSlateDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func StrokeSlateDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func StrokeTeal(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func StrokeTealAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func StrokeTealDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func StrokeTealDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func StrokeTomato(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func StrokeTomatoAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func StrokeTomatoDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func StrokeTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func StrokeViolet(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func StrokeVioletAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func StrokeVioletDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func StrokeVioletDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func StrokeWhiteAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func StrokeYellow(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func StrokeYellowAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func StrokeYellowDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func StrokeYellowDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func StrokeColor(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(strokeProp): val.Value(),
		}
	}
}
