package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func Shadow2xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "2xs"),
		}
	}
}

func ShadowXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "xs"),
		}
	}
}

func ShadowSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "sm"),
		}
	}
}

func ShadowMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "md"),
		}
	}
}

func ShadowLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "lg"),
		}
	}
}

func ShadowXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "xl"),
		}
	}
}

func Shadow2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "2xl"),
		}
	}
}

func ShadowNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): "0 0 #0000",
		}
	}
}

func Shadow(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): value.Value(),
		}
	}
}

func ShadowInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): "inherit",
		}
	}
}

func ShadowCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): "currentColor",
		}
	}
}

func ShadowTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): "transparent",
		}
	}
}

func ShadowBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func ShadowWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func ShadowAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func ShadowAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func ShadowAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func ShadowAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func ShadowBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func ShadowBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func ShadowBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func ShadowBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func ShadowBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func ShadowBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func ShadowBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func ShadowBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func ShadowBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func ShadowBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func ShadowBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func ShadowBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func ShadowBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func ShadowCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func ShadowCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func ShadowCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func ShadowCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func ShadowCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func ShadowCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func ShadowCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func ShadowCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func ShadowGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func ShadowGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func ShadowGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func ShadowGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func ShadowGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func ShadowGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func ShadowGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func ShadowGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func ShadowGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func ShadowGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func ShadowGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func ShadowGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func ShadowGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func ShadowGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func ShadowGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func ShadowGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func ShadowIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func ShadowIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func ShadowIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func ShadowIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func ShadowIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func ShadowIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func ShadowIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func ShadowIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func ShadowJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func ShadowJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func ShadowJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func ShadowJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func ShadowLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func ShadowLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func ShadowLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func ShadowLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func ShadowMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func ShadowMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func ShadowMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func ShadowMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func ShadowMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func ShadowMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func ShadowMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func ShadowMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func ShadowOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func ShadowOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func ShadowOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func ShadowOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func ShadowOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func ShadowOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func ShadowOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func ShadowOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func ShadowPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func ShadowPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func ShadowPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func ShadowPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func ShadowPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func ShadowPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func ShadowPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func ShadowPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func ShadowPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func ShadowPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func ShadowPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func ShadowPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func ShadowRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func ShadowRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func ShadowRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func ShadowRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func ShadowRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func ShadowRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func ShadowRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func ShadowRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func ShadowSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func ShadowSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func ShadowSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func ShadowSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func ShadowSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func ShadowSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func ShadowSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func ShadowSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func ShadowSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func ShadowSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func ShadowSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func ShadowSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func ShadowSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func ShadowSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func ShadowSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func ShadowSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func ShadowTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func ShadowTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func ShadowTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func ShadowTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func ShadowTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func ShadowTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func ShadowTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func ShadowTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func ShadowViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func ShadowVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func ShadowVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func ShadowVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func ShadowWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func ShadowYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func ShadowYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func ShadowYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func ShadowYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func ShadowColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(shadowColorProp): value.Value(),
		}
	}
}

//--

func InsetShadow2xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "2xs"),
		}
	}
}

func InsetShadowXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "xs"),
		}
	}
}

func InsetShadowSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "sm"),
		}
	}
}

func InsetShadowMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "md"),
		}
	}
}

func InsetShadowLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "lg"),
		}
	}
}

func InsetShadowXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "xl"),
		}
	}
}

func InsetShadow2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): s.theme.UseVarKey(theme.Shadow, "2xl"),
		}
	}
}

func InsetShadowNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): "inset 0 0 #0000",
		}
	}
}

func InsetShadow(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(boxShadowProp): value.Value(),
		}
	}
}

func InsetShadowInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): "inherit",
		}
	}
}

func InsetShadowCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): "currentColor",
		}
	}
}

func InsetShadowTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): "transparent",
		}
	}
}

func InsetShadowBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func InsetShadowWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func InsetShadowAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func InsetShadowAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func InsetShadowAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func InsetShadowAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func InsetShadowBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func InsetShadowBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func InsetShadowBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func InsetShadowBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func InsetShadowBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func InsetShadowBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func InsetShadowBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func InsetShadowBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func InsetShadowBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func InsetShadowBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func InsetShadowCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func InsetShadowCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func InsetShadowCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func InsetShadowCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func InsetShadowCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func InsetShadowCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func InsetShadowGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func InsetShadowGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func InsetShadowGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func InsetShadowGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func InsetShadowGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func InsetShadowGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func InsetShadowGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func InsetShadowGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func InsetShadowGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func InsetShadowGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func InsetShadowGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func InsetShadowGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func InsetShadowIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func InsetShadowIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func InsetShadowIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func InsetShadowIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func InsetShadowIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func InsetShadowIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func InsetShadowJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func InsetShadowJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func InsetShadowJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func InsetShadowLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func InsetShadowLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func InsetShadowLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func InsetShadowMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func InsetShadowMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func InsetShadowMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func InsetShadowMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func InsetShadowMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func InsetShadowMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func InsetShadowOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func InsetShadowOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func InsetShadowOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func InsetShadowOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func InsetShadowOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func InsetShadowOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func InsetShadowPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func InsetShadowPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func InsetShadowPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func InsetShadowPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func InsetShadowPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func InsetShadowPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func InsetShadowPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func InsetShadowPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func InsetShadowPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func InsetShadowRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func InsetShadowRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func InsetShadowRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func InsetShadowRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func InsetShadowRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func InsetShadowRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func InsetShadowSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func InsetShadowSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func InsetShadowSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func InsetShadowSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func InsetShadowSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func InsetShadowSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func InsetShadowSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func InsetShadowSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func InsetShadowSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func InsetShadowSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func InsetShadowSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func InsetShadowSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func InsetShadowTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func InsetShadowTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func InsetShadowTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func InsetShadowTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func InsetShadowTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func InsetShadowTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func InsetShadowVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func InsetShadowVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func InsetShadowVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func InsetShadowWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func InsetShadowYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func InsetShadowYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func InsetShadowYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func InsetShadowYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func InsetShadowColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetShadowColorProp): value.Value(),
		}
	}
}

// --

func Ring(value ...customValue) ApplyProp {
	propValue := ""
	if len(value) == 0 {
		propValue = "0 0 0 1px"
	} else {

		switch v := value[0].(type) {
		case *val:
			switch n := v.value.(type) {
			case int:
				propValue = fmt.Sprintf("0 0 0 %dpx", n)
			default:
				propValue = fmt.Sprintf("0 0 0 %s", v.Value())
			}
		case *customVariableProp:
			propValue = fmt.Sprintf("0 0 0 %s", v.Value())
		}
	}

	return func(s *style) styleProp {
		return &properties{
			string(ringShadowProp): propValue,
		}
	}
}

func RingInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): "inherit",
		}
	}
}

func RingCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): "currentColor",
		}
	}
}

func RingTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): "transparent",
		}
	}
}

func RingBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func RingWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func RingAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func RingAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func RingAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func RingAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func RingBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func RingBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func RingBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func RingBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func RingBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func RingBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func RingBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func RingBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func RingBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func RingBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func RingBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func RingBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func RingBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func RingCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func RingCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func RingCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func RingCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func RingCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func RingCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func RingCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func RingCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func RingGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func RingGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func RingGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func RingGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func RingGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func RingGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func RingGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func RingGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func RingGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func RingGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func RingGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func RingGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func RingGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func RingGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func RingGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func RingGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func RingIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func RingIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func RingIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func RingIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func RingIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func RingIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func RingIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func RingIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func RingJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func RingJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func RingJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func RingJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func RingLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func RingLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func RingLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func RingLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func RingMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func RingMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func RingMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func RingMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func RingMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func RingMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func RingMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func RingMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func RingOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func RingOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func RingOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func RingOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func RingOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func RingOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func RingOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func RingOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func RingPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func RingPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func RingPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func RingPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func RingPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func RingPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func RingPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func RingPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func RingPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func RingPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func RingPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func RingPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func RingRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func RingRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func RingRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func RingRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func RingRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func RingRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func RingRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func RingRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func RingSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func RingSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func RingSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func RingSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func RingSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func RingSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func RingSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func RingSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func RingSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func RingSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func RingSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func RingSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func RingSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func RingSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func RingSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func RingSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func RingTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func RingTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func RingTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func RingTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func RingTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func RingTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func RingTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func RingTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func RingViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func RingVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func RingVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func RingVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func RingWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func RingYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func RingYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func RingYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}

func RingYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func RingColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(ringColorProp): value.Value(),
		}
	}
}

//--

func InsetRing(value ...customValue) ApplyProp {
	propValue := ""
	if len(value) == 0 {
		propValue = "0 0 0 1px"
	} else {

		switch v := value[0].(type) {
		case *val:
			switch n := v.value.(type) {
			case int:
				propValue = fmt.Sprintf("0 0 0 %dpx", n)
			default:
				propValue = fmt.Sprintf("0 0 0 %s", v.Value())
			}
		case *customVariableProp:
			propValue = fmt.Sprintf("0 0 0 %s", v.Value())
		}
	}

	return func(s *style) styleProp {
		return &properties{
			string(insetRingShadowProp): propValue,
		}
	}
}

func InsetRingInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): "inherit",
		}
	}
}

func InsetRingCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): "currentColor",
		}
	}
}

func InsetRingTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): "transparent",
		}
	}
}

func InsetRingBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func InsetRingWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}
func InsetRingAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func InsetRingAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func InsetRingAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func InsetRingAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func InsetRingBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func InsetRingBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func InsetRingBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func InsetRingBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func InsetRingBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func InsetRingBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func InsetRingBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func InsetRingBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func InsetRingBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func InsetRingBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func InsetRingCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func InsetRingCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func InsetRingCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func InsetRingCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func InsetRingCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func InsetRingCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func InsetRingGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func InsetRingGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func InsetRingGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func InsetRingGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func InsetRingGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func InsetRingGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func InsetRingGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func InsetRingGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func InsetRingGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func InsetRingGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func InsetRingGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func InsetRingGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func InsetRingIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func InsetRingIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func InsetRingIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func InsetRingIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func InsetRingIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func InsetRingIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func InsetRingJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func InsetRingJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func InsetRingJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func InsetRingLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func InsetRingLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func InsetRingLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func InsetRingMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func InsetRingMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func InsetRingMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func InsetRingMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func InsetRingMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func InsetRingMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func InsetRingOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func InsetRingOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func InsetRingOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func InsetRingOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func InsetRingOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func InsetRingOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func InsetRingPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func InsetRingPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func InsetRingPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func InsetRingPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func InsetRingPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func InsetRingPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func InsetRingPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func InsetRingPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func InsetRingPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func InsetRingRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func InsetRingRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func InsetRingRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func InsetRingRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func InsetRingRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func InsetRingRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func InsetRingSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func InsetRingSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func InsetRingSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func InsetRingSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func InsetRingSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func InsetRingSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func InsetRingSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func InsetRingSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func InsetRingSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func InsetRingSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func InsetRingSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func InsetRingSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func InsetRingTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func InsetRingTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func InsetRingTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func InsetRingTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func InsetRingTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func InsetRingTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func InsetRingVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func InsetRingVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func InsetRingVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func InsetRingWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func InsetRingYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func InsetRingYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func InsetRingYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}

func InsetRingYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func InsetRingColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(insetRingColorProp): value.Value(),
		}
	}
}
