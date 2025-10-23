package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func FillInherit() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): "inherit",
		}
	}
}

func FillCurrent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): "currentColor",
		}
	}
}

func FillTransparent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): "transparent",
		}
	}
}

func FillBlack() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, "black"),
		}
	}
}

func FillWhite() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, "white"),
		}
	}
}

func FillAmber(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func FillAmberAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func FillAmberDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func FillAmberDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func FillBlackAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func FillBlue(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func FillBlueAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func FillBlueDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func FillBlueDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func FillBronze(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func FillBronzeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func FillBronzeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func FillBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func FillBrown(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func FillBrownAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func FillBrownDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func FillBrownDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func FillCrimson(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func FillCrimsonAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func FillCrimsonDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func FillCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func FillCyan(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func FillCyanAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func FillCyanDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func FillCyanDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func FillGold(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func FillGoldAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func FillGoldDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func FillGoldDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func FillGrass(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func FillGrassAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func FillGrassDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func FillGrassDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func FillGray(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func FillGrayAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func FillGrayDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func FillGrayDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func FillGreen(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func FillGreenAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func FillGreenDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func FillGreenDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func FillIndigo(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func FillIndigoAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func FillIndigoDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func FillIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func FillIris(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func FillIrisAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func FillIrisDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func FillIrisDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func FillJade(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func FillJadeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func FillJadeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func FillJadeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func FillLime(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func FillLimeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func FillLimeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func FillLimeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func FillMauve(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func FillMauveAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func FillMauveDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func FillMauveDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func FillMint(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func FillMintAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func FillMintDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func FillMintDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func FillOlive(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func FillOliveAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func FillOliveDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func FillOliveDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func FillOrange(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func FillOrangeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func FillOrangeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func FillOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func FillPink(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func FillPinkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func FillPinkDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func FillPinkDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func FillPlum(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func FillPlumAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func FillPlumDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func FillPlumDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func FillPurple(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func FillPurpleAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func FillPurpleDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func FillPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func FillRed(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func FillRedAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func FillRedDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func FillRedDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func FillRuby(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func FillRubyAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func FillRubyDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func FillRubyDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func FillSage(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func FillSageAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func FillSageDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func FillSageDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func FillSand(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func FillSandAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func FillSandDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func FillSandDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func FillSky(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func FillSkyAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func FillSkyDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func FillSkyDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func FillSlate(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func FillSlateAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func FillSlateDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func FillSlateDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func FillTeal(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func FillTealAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func FillTealDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func FillTealDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func FillTomato(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func FillTomatoAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func FillTomatoDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func FillTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func FillViolet(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func FillVioletAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func FillVioletDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func FillVioletDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func FillWhiteAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func FillYellow(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func FillYellowAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func FillYellowDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func FillYellowDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func Fill(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fillProp): val.Value(),
		}
	}
}
