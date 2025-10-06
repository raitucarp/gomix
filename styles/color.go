package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func TextInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): "inherit",
		}
	}
}

func TextCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): "currentColor",
		}
	}
}

func TextTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): "transparent",
		}
	}
}

func TextBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func TextWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func TextAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func TextAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func TextAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func TextAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func TextBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func TextBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func TextBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func TextBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func TextBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func TextBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func TextBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func TextBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func TextBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func TextBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func TextBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func TextBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func TextBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func TextCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func TextCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func TextCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func TextCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func TextCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func TextCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func TextCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func TextCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func TextGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func TextGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func TextGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func TextGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func TextGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func TextGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func TextGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func TextGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func TextGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func TextGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func TextGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func TextGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func TextGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func TextGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func TextGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func TextGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func TextIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func TextIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func TextIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func TextIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func TextIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func TextIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func TextIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func TextIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func TextJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func TextJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func TextJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func TextJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func TextLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func TextLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func TextLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func TextLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func TextMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func TextMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func TextMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func TextMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func TextMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func TextMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func TextMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func TextMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func TextOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func TextOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func TextOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func TextOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func TextOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func TextOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func TextOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func TextOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func TextPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func TextPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func TextPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func TextPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func TextPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func TextPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func TextPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func TextPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func TextPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func TextPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func TextPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func TextPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func TextRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func TextRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func TextRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func TextRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func TextRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func TextRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func TextRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func TextRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func TextSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func TextSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func TextSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func TextSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func TextSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func TextSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func TextSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func TextSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func TextSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func TextSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func TextSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func TextSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func TextSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func TextSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func TextSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func TextSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func TextTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func TextTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func TextTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func TextTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func TextTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func TextTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func TextTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func TextTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func TextViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func TextVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func TextVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func TextVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func TextWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func TextYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func TextYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func TextYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func TextYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func TextColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(colorProp): value.Value(),
		}
	}
}
