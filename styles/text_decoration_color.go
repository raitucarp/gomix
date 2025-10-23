package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func DecorInherit() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): "inherit",
		}
	}
}

func DecorCurrent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): "currentColor",
		}
	}
}

func DecorTransparent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): "transparent",
		}
	}
}

func DecorBlack() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, "black"),
		}
	}
}

func DecorWhite() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, "white"),
		}
	}
}

func DecorAmber(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func DecorAmberAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func DecorAmberDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func DecorAmberDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func DecorBlackAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func DecorBlue(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func DecorBlueAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func DecorBlueDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func DecorBlueDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func DecorBronze(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func DecorBronzeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func DecorBronzeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func DecorBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func DecorBrown(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func DecorBrownAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func DecorBrownDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func DecorBrownDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func DecorCrimson(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func DecorCrimsonAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func DecorCrimsonDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func DecorCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func DecorCyan(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func DecorCyanAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func DecorCyanDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func DecorCyanDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func DecorGold(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func DecorGoldAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func DecorGoldDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func DecorGoldDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func DecorGrass(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func DecorGrassAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func DecorGrassDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func DecorGrassDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func DecorGray(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func DecorGrayAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func DecorGrayDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func DecorGrayDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func DecorGreen(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func DecorGreenAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func DecorGreenDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func DecorGreenDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func DecorIndigo(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func DecorIndigoAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func DecorIndigoDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func DecorIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func DecorIris(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func DecorIrisAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func DecorIrisDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func DecorIrisDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func DecorJade(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func DecorJadeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func DecorJadeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func DecorJadeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func DecorLime(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func DecorLimeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func DecorLimeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func DecorLimeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func DecorMauve(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func DecorMauveAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func DecorMauveDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func DecorMauveDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func DecorMint(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func DecorMintAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func DecorMintDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func DecorMintDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func DecorOlive(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func DecorOliveAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func DecorOliveDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func DecorOliveDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func DecorOrange(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func DecorOrangeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func DecorOrangeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func DecorOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func DecorPink(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func DecorPinkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func DecorPinkDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func DecorPinkDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func DecorPlum(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func DecorPlumAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func DecorPlumDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func DecorPlumDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func DecorPurple(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func DecorPurpleAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func DecorPurpleDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func DecorPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func DecorRed(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func DecorRedAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func DecorRedDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func DecorRedDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func DecorRuby(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func DecorRubyAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func DecorRubyDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func DecorRubyDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func DecorSage(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func DecorSageAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func DecorSageDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func DecorSageDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func DecorSand(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func DecorSandAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func DecorSandDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func DecorSandDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func DecorSky(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func DecorSkyAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func DecorSkyDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func DecorSkyDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func DecorSlate(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func DecorSlateAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func DecorSlateDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func DecorSlateDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func DecorTeal(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func DecorTealAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func DecorTealDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func DecorTealDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func DecorTomato(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func DecorTomatoAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func DecorTomatoDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func DecorTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func DecorViolet(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func DecorVioletAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func DecorVioletDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func DecorVioletDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func DecorWhiteAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func DecorYellow(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func DecorYellowAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func DecorYellowDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func DecorYellowDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func DecorColorBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(textDecorColorProp): val.Value(),
		}
	}
}
