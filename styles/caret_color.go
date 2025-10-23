package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func CaretInherit() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): "inherit",
		}
	}
}

func CaretCurrent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): "currentColor",
		}
	}
}

func CaretTransparent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): "transparent",
		}
	}
}

func CaretBlack() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, "black"),
		}
	}
}

func CaretWhite() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, "white"),
		}
	}
}

func CaretAmber(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func CaretAmberAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func CaretAmberDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func CaretAmberDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func CaretBlackAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func CaretBlue(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func CaretBlueAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func CaretBlueDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func CaretBlueDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func CaretBronze(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func CaretBronzeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func CaretBronzeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func CaretBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func CaretBrown(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func CaretBrownAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func CaretBrownDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func CaretBrownDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func CaretCrimson(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func CaretCrimsonAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func CaretCrimsonDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func CaretCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func CaretCyan(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func CaretCyanAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func CaretCyanDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func CaretCyanDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func CaretGold(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func CaretGoldAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func CaretGoldDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func CaretGoldDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func CaretGrass(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func CaretGrassAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func CaretGrassDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func CaretGrassDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func CaretGray(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func CaretGrayAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func CaretGrayDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func CaretGrayDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func CaretGreen(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func CaretGreenAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func CaretGreenDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func CaretGreenDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func CaretIndigo(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func CaretIndigoAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func CaretIndigoDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func CaretIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func CaretIris(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func CaretIrisAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func CaretIrisDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func CaretIrisDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func CaretJade(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func CaretJadeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func CaretJadeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func CaretJadeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func CaretLime(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func CaretLimeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func CaretLimeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func CaretLimeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func CaretMauve(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func CaretMauveAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func CaretMauveDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func CaretMauveDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func CaretMint(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func CaretMintAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func CaretMintDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func CaretMintDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func CaretOlive(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func CaretOliveAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func CaretOliveDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func CaretOliveDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func CaretOrange(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func CaretOrangeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func CaretOrangeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func CaretOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func CaretPink(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func CaretPinkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func CaretPinkDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func CaretPinkDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func CaretPlum(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func CaretPlumAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func CaretPlumDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func CaretPlumDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func CaretPurple(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func CaretPurpleAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func CaretPurpleDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func CaretPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func CaretRed(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func CaretRedAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func CaretRedDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func CaretRedDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func CaretRuby(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func CaretRubyAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func CaretRubyDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func CaretRubyDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func CaretSage(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func CaretSageAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func CaretSageDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func CaretSageDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func CaretSand(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func CaretSandAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func CaretSandDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func CaretSandDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func CaretSky(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func CaretSkyAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func CaretSkyDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func CaretSkyDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func CaretSlate(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func CaretSlateAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func CaretSlateDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func CaretSlateDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func CaretTeal(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func CaretTealAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func CaretTealDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func CaretTealDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func CaretTomato(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func CaretTomatoAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func CaretTomatoDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func CaretTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func CaretViolet(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func CaretVioletAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func CaretVioletDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func CaretVioletDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func CaretWhiteAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func CaretYellow(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func CaretYellowAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func CaretYellowDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func CaretYellowDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): s.theme.UseVarKey(themeColor, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func Caret(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(caretColorProp): val.Value(),
		}
	}
}
