package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func BorderInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): "inherit",
		}
	}
}

func BorderCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): "currentColor",
		}
	}
}

func BorderTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): "transparent",
		}
	}
}

func BorderBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func BorderWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}
func BorderAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func BorderAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func BorderAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func BorderAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func BorderBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func BorderBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func BorderBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func BorderBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func BorderBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func BorderBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func BorderBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func BorderBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func BorderBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func BorderBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func BorderBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func BorderBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func BorderBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func BorderCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func BorderCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func BorderCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func BorderCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func BorderCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func BorderCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func BorderCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func BorderCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func BorderGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func BorderGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func BorderGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func BorderGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func BorderGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func BorderGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func BorderGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func BorderGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func BorderGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func BorderGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func BorderGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func BorderGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func BorderGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func BorderGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func BorderGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func BorderGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func BorderIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func BorderIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func BorderIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func BorderIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func BorderIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func BorderIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func BorderIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func BorderIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func BorderJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func BorderJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func BorderJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func BorderJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func BorderLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func BorderLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func BorderLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func BorderLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func BorderMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func BorderMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func BorderMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func BorderMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func BorderMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func BorderMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func BorderMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func BorderMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func BorderOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func BorderOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func BorderOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func BorderOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func BorderOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func BorderOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func BorderOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func BorderOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func BorderPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func BorderPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func BorderPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func BorderPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func BorderPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func BorderPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func BorderPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func BorderPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func BorderPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func BorderPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func BorderPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func BorderPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func BorderRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func BorderRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func BorderRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func BorderRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func BorderRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func BorderRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func BorderRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func BorderRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func BorderSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func BorderSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func BorderSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func BorderSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func BorderSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func BorderSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func BorderSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func BorderSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func BorderSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func BorderSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func BorderSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func BorderSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func BorderSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func BorderSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func BorderSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func BorderSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func BorderTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func BorderTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func BorderTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func BorderTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func BorderTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func BorderTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func BorderTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func BorderTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func BorderViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func BorderVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func BorderVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func BorderVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func BorderWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func BorderYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func BorderYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func BorderYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func BorderYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func BorderColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderColorProp): value.Value(),
		}
	}
}

//--

func BorderXInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): "inherit",
		}
	}
}

func BorderXCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): "currentColor",
		}
	}
}

func BorderXTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): "transparent",
		}
	}
}

func BorderXBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func BorderXWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}
func BorderXAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func BorderXAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func BorderXAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func BorderXAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func BorderXBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func BorderXBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func BorderXBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func BorderXBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func BorderXBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func BorderXBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func BorderXBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func BorderXBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func BorderXBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func BorderXBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func BorderXBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func BorderXBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func BorderXBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func BorderXCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func BorderXCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func BorderXCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func BorderXCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func BorderXCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func BorderXCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func BorderXCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func BorderXCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func BorderXGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func BorderXGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func BorderXGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func BorderXGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func BorderXGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func BorderXGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func BorderXGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func BorderXGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func BorderXGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func BorderXGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func BorderXGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func BorderXGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func BorderXGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func BorderXGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func BorderXGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func BorderXGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func BorderXIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func BorderXIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func BorderXIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func BorderXIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func BorderXIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func BorderXIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func BorderXIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func BorderXIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func BorderXJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func BorderXJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func BorderXJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func BorderXJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func BorderXLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func BorderXLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func BorderXLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func BorderXLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func BorderXMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func BorderXMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func BorderXMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func BorderXMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func BorderXMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func BorderXMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func BorderXMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func BorderXMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func BorderXOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func BorderXOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func BorderXOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func BorderXOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func BorderXOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func BorderXOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func BorderXOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func BorderXOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func BorderXPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func BorderXPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func BorderXPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func BorderXPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func BorderXPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func BorderXPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func BorderXPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func BorderXPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func BorderXPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func BorderXPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func BorderXPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func BorderXPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func BorderXRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func BorderXRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func BorderXRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func BorderXRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func BorderXRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func BorderXRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func BorderXRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func BorderXRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func BorderXSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func BorderXSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func BorderXSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func BorderXSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func BorderXSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func BorderXSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func BorderXSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func BorderXSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func BorderXSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func BorderXSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func BorderXSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func BorderXSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func BorderXSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func BorderXSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func BorderXSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func BorderXSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func BorderXTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func BorderXTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func BorderXTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func BorderXTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func BorderXTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func BorderXTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func BorderXTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func BorderXTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func BorderXViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func BorderXVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func BorderXVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func BorderXVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func BorderXWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func BorderXYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func BorderXYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func BorderXYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func BorderXYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func BorderXColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineColorProp): value.Value(),
		}
	}
}

// --

func BorderYInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): "inherit",
		}
	}
}

func BorderYCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): "currentColor",
		}
	}
}

func BorderYTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): "transparent",
		}
	}
}

func BorderYBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func BorderYWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}
func BorderYAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func BorderYAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func BorderYAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func BorderYAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func BorderYBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func BorderYBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func BorderYBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func BorderYBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func BorderYBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func BorderYBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func BorderYBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func BorderYBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func BorderYBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func BorderYBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func BorderYBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func BorderYBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func BorderYBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func BorderYCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func BorderYCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func BorderYCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func BorderYCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func BorderYCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func BorderYCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func BorderYCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func BorderYCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func BorderYGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func BorderYGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func BorderYGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func BorderYGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func BorderYGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func BorderYGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func BorderYGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func BorderYGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func BorderYGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func BorderYGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func BorderYGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func BorderYGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func BorderYGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func BorderYGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func BorderYGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func BorderYGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func BorderYIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func BorderYIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func BorderYIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func BorderYIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func BorderYIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func BorderYIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func BorderYIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func BorderYIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func BorderYJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func BorderYJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func BorderYJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func BorderYJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func BorderYLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func BorderYLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func BorderYLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func BorderYLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func BorderYMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func BorderYMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func BorderYMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func BorderYMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func BorderYMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func BorderYMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func BorderYMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func BorderYMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func BorderYOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func BorderYOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func BorderYOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func BorderYOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func BorderYOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func BorderYOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func BorderYOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func BorderYOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func BorderYPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func BorderYPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func BorderYPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func BorderYPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func BorderYPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func BorderYPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func BorderYPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func BorderYPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func BorderYPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func BorderYPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func BorderYPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func BorderYPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func BorderYRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func BorderYRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func BorderYRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func BorderYRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func BorderYRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func BorderYRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func BorderYRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func BorderYRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func BorderYSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func BorderYSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func BorderYSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func BorderYSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func BorderYSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func BorderYSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func BorderYSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func BorderYSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func BorderYSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func BorderYSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func BorderYSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func BorderYSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func BorderYSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func BorderYSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func BorderYSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func BorderYSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func BorderYTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func BorderYTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func BorderYTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func BorderYTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func BorderYTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func BorderYTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func BorderYTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func BorderYTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func BorderYViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func BorderYVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func BorderYVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func BorderYVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func BorderYWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func BorderYYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func BorderYYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func BorderYYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func BorderYYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func BorderYColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBlockColorProp): value.Value(),
		}
	}
}

// --

func BorderStartInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): "inherit",
		}
	}
}

func BorderStartCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): "currentColor",
		}
	}
}

func BorderStartTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): "transparent",
		}
	}
}

func BorderStartBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func BorderStartWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}
func BorderStartAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func BorderStartAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func BorderStartAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func BorderStartAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func BorderStartBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func BorderStartBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func BorderStartBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func BorderStartBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func BorderStartBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func BorderStartBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func BorderStartBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func BorderStartBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func BorderStartBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func BorderStartBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func BorderStartCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func BorderStartCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func BorderStartCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func BorderStartCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func BorderStartCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func BorderStartCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func BorderStartGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func BorderStartGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func BorderStartGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func BorderStartGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func BorderStartGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func BorderStartGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func BorderStartGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func BorderStartGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func BorderStartGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func BorderStartGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func BorderStartGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func BorderStartGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func BorderStartIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func BorderStartIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func BorderStartIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func BorderStartIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func BorderStartIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func BorderStartIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func BorderStartJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func BorderStartJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func BorderStartJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func BorderStartLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func BorderStartLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func BorderStartLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func BorderStartMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func BorderStartMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func BorderStartMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func BorderStartMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func BorderStartMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func BorderStartMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func BorderStartOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func BorderStartOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func BorderStartOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func BorderStartOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func BorderStartOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func BorderStartOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func BorderStartPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func BorderStartPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func BorderStartPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func BorderStartPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func BorderStartPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func BorderStartPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func BorderStartPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func BorderStartPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func BorderStartPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func BorderStartRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func BorderStartRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func BorderStartRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func BorderStartRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func BorderStartRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func BorderStartRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func BorderStartSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func BorderStartSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func BorderStartSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func BorderStartSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func BorderStartSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func BorderStartSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func BorderStartSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func BorderStartSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func BorderStartSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func BorderStartSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func BorderStartSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func BorderStartSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func BorderStartTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func BorderStartTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func BorderStartTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func BorderStartTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func BorderStartTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func BorderStartTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func BorderStartVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func BorderStartVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func BorderStartVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func BorderStartWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func BorderStartYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func BorderStartYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func BorderStartYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func BorderStartYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func BorderStartColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineStartColorProp): value.Value(),
		}
	}
}

// --

func BorderEndInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): "inherit",
		}
	}
}

func BorderEndCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): "currentColor",
		}
	}
}

func BorderEndTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): "transparent",
		}
	}
}

func BorderEndBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func BorderEndWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}
func BorderEndAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func BorderEndAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func BorderEndAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func BorderEndAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func BorderEndBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func BorderEndBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func BorderEndBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func BorderEndBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func BorderEndBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func BorderEndBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func BorderEndBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func BorderEndBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func BorderEndBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func BorderEndBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func BorderEndCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func BorderEndCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func BorderEndCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func BorderEndCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func BorderEndCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func BorderEndCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func BorderEndGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func BorderEndGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func BorderEndGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func BorderEndGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func BorderEndGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func BorderEndGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func BorderEndGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func BorderEndGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func BorderEndGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func BorderEndGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func BorderEndGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func BorderEndGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func BorderEndIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func BorderEndIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func BorderEndIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func BorderEndIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func BorderEndIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func BorderEndIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func BorderEndJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func BorderEndJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func BorderEndJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func BorderEndLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func BorderEndLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func BorderEndLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func BorderEndMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func BorderEndMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func BorderEndMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func BorderEndMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func BorderEndMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func BorderEndMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func BorderEndOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func BorderEndOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func BorderEndOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func BorderEndOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func BorderEndOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func BorderEndOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func BorderEndPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func BorderEndPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func BorderEndPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func BorderEndPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func BorderEndPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func BorderEndPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func BorderEndPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func BorderEndPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func BorderEndPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func BorderEndRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func BorderEndRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func BorderEndRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func BorderEndRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func BorderEndRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func BorderEndRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func BorderEndSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func BorderEndSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func BorderEndSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func BorderEndSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func BorderEndSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func BorderEndSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func BorderEndSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func BorderEndSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func BorderEndSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func BorderEndSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func BorderEndSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func BorderEndSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func BorderEndTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func BorderEndTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func BorderEndTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func BorderEndTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func BorderEndTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func BorderEndTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func BorderEndVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func BorderEndVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func BorderEndVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func BorderEndWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func BorderEndYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func BorderEndYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func BorderEndYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func BorderEndYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func BorderEndColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderInlineEndColorProp): value.Value(),
		}
	}
}

// --

func BorderTopInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): "inherit",
		}
	}
}

func BorderTopCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): "currentColor",
		}
	}
}

func BorderTopTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): "transparent",
		}
	}
}

func BorderTopBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func BorderTopWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}
func BorderTopAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func BorderTopAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func BorderTopAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func BorderTopAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func BorderTopBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func BorderTopBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func BorderTopBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func BorderTopBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func BorderTopBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func BorderTopBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func BorderTopBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func BorderTopBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func BorderTopBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func BorderTopBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func BorderTopCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func BorderTopCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func BorderTopCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func BorderTopCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func BorderTopCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func BorderTopCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func BorderTopGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func BorderTopGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func BorderTopGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func BorderTopGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func BorderTopGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func BorderTopGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func BorderTopGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func BorderTopGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func BorderTopGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func BorderTopGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func BorderTopGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func BorderTopGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func BorderTopIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func BorderTopIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func BorderTopIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func BorderTopIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func BorderTopIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func BorderTopIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func BorderTopJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func BorderTopJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func BorderTopJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func BorderTopLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func BorderTopLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func BorderTopLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func BorderTopMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func BorderTopMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func BorderTopMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func BorderTopMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func BorderTopMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func BorderTopMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func BorderTopOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func BorderTopOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func BorderTopOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func BorderTopOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func BorderTopOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func BorderTopOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func BorderTopPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func BorderTopPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func BorderTopPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func BorderTopPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func BorderTopPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func BorderTopPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func BorderTopPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func BorderTopPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func BorderTopPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func BorderTopRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func BorderTopRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func BorderTopRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func BorderTopRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func BorderTopRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func BorderTopRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func BorderTopSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func BorderTopSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func BorderTopSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func BorderTopSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func BorderTopSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func BorderTopSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func BorderTopSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func BorderTopSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func BorderTopSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func BorderTopSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func BorderTopSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func BorderTopSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func BorderTopTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func BorderTopTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func BorderTopTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func BorderTopTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func BorderTopTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func BorderTopTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func BorderTopVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func BorderTopVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func BorderTopVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func BorderTopWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func BorderTopYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func BorderTopYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func BorderTopYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func BorderTopYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func BorderTopColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopColorProp): value.Value(),
		}
	}
}

// --

func BorderRightInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): "inherit",
		}
	}
}

func BorderRightCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): "currentColor",
		}
	}
}

func BorderRightTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): "transparent",
		}
	}
}

func BorderRightBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func BorderRightWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func BorderRightAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func BorderRightAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func BorderRightAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func BorderRightAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func BorderRightBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func BorderRightBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func BorderRightBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func BorderRightBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func BorderRightBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func BorderRightBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func BorderRightBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func BorderRightBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func BorderRightBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func BorderRightBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func BorderRightCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func BorderRightCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func BorderRightCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func BorderRightCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func BorderRightCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func BorderRightCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func BorderRightGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func BorderRightGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func BorderRightGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func BorderRightGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func BorderRightGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func BorderRightGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func BorderRightGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func BorderRightGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func BorderRightGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func BorderRightGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func BorderRightGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func BorderRightGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func BorderRightIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func BorderRightIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func BorderRightIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func BorderRightIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func BorderRightIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func BorderRightIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func BorderRightJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func BorderRightJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func BorderRightJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func BorderRightLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func BorderRightLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func BorderRightLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func BorderRightMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func BorderRightMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func BorderRightMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func BorderRightMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func BorderRightMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func BorderRightMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func BorderRightOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func BorderRightOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func BorderRightOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func BorderRightOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func BorderRightOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func BorderRightOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func BorderRightPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func BorderRightPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func BorderRightPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func BorderRightPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func BorderRightPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func BorderRightPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func BorderRightPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func BorderRightPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func BorderRightPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func BorderRightRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func BorderRightRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func BorderRightRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func BorderRightRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func BorderRightRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func BorderRightRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func BorderRightSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func BorderRightSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func BorderRightSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func BorderRightSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func BorderRightSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func BorderRightSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func BorderRightSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func BorderRightSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func BorderRightSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func BorderRightSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func BorderRightSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func BorderRightSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func BorderRightTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func BorderRightTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func BorderRightTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func BorderRightTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func BorderRightTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func BorderRightTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func BorderRightVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func BorderRightVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func BorderRightVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func BorderRightWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func BorderRightYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func BorderRightYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func BorderRightYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func BorderRightYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func BorderRightColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRightColorProp): value.Value(),
		}
	}
}

// --

func BorderBottomInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): "inherit",
		}
	}
}

func BorderBottomCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): "currentColor",
		}
	}
}

func BorderBottomTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): "transparent",
		}
	}
}

func BorderBottomBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func BorderBottomWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func BorderBottomAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func BorderBottomAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func BorderBottomAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func BorderBottomAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func BorderBottomBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func BorderBottomBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func BorderBottomBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func BorderBottomBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func BorderBottomBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func BorderBottomBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func BorderBottomBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func BorderBottomBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func BorderBottomBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func BorderBottomBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func BorderBottomCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func BorderBottomCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func BorderBottomCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func BorderBottomCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func BorderBottomCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func BorderBottomCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func BorderBottomGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func BorderBottomGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func BorderBottomGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func BorderBottomGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func BorderBottomGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func BorderBottomGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func BorderBottomGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func BorderBottomGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func BorderBottomGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func BorderBottomGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func BorderBottomGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func BorderBottomGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func BorderBottomIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func BorderBottomIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func BorderBottomIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func BorderBottomIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func BorderBottomIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func BorderBottomIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func BorderBottomJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func BorderBottomJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func BorderBottomJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func BorderBottomLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func BorderBottomLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func BorderBottomLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func BorderBottomMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func BorderBottomMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func BorderBottomMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func BorderBottomMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func BorderBottomMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func BorderBottomMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func BorderBottomOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func BorderBottomOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func BorderBottomOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func BorderBottomOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func BorderBottomOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func BorderBottomOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func BorderBottomPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func BorderBottomPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func BorderBottomPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func BorderBottomPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func BorderBottomPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func BorderBottomPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func BorderBottomPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func BorderBottomPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func BorderBottomPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func BorderBottomRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func BorderBottomRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func BorderBottomRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func BorderBottomRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func BorderBottomRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func BorderBottomRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func BorderBottomSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func BorderBottomSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func BorderBottomSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func BorderBottomSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func BorderBottomSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func BorderBottomSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func BorderBottomSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func BorderBottomSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func BorderBottomSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func BorderBottomSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func BorderBottomSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func BorderBottomSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func BorderBottomTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func BorderBottomTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func BorderBottomTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func BorderBottomTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func BorderBottomTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func BorderBottomTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func BorderBottomVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func BorderBottomVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func BorderBottomVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func BorderBottomWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func BorderBottomYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func BorderBottomYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func BorderBottomYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func BorderBottomYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func BorderBottomColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomColorProp): value.Value(),
		}
	}
}

// --

func BorderLeftInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): "inherit",
		}
	}
}

func BorderLeftCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): "currentColor",
		}
	}
}

func BorderLeftTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): "transparent",
		}
	}
}

func BorderLeftBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func BorderLeftWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func BorderLeftAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func BorderLeftAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func BorderLeftAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func BorderLeftAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func BorderLeftBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func BorderLeftBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func BorderLeftBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func BorderLeftBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func BorderLeftBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func BorderLeftBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func BorderLeftBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func BorderLeftBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func BorderLeftBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func BorderLeftBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func BorderLeftCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func BorderLeftCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func BorderLeftCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func BorderLeftCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func BorderLeftCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func BorderLeftCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func BorderLeftGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func BorderLeftGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func BorderLeftGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func BorderLeftGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func BorderLeftGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func BorderLeftGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func BorderLeftGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func BorderLeftGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func BorderLeftGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func BorderLeftGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func BorderLeftGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func BorderLeftGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func BorderLeftIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func BorderLeftIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func BorderLeftIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func BorderLeftIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func BorderLeftIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func BorderLeftIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func BorderLeftJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func BorderLeftJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func BorderLeftJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func BorderLeftLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func BorderLeftLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func BorderLeftLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func BorderLeftMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func BorderLeftMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func BorderLeftMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func BorderLeftMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func BorderLeftMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func BorderLeftMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func BorderLeftOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func BorderLeftOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func BorderLeftOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func BorderLeftOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func BorderLeftOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func BorderLeftOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftPink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func BorderLeftPinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func BorderLeftPinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func BorderLeftPinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftPlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func BorderLeftPlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func BorderLeftPlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func BorderLeftPlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftPurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func BorderLeftPurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func BorderLeftPurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func BorderLeftPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func BorderLeftRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func BorderLeftRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func BorderLeftRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func BorderLeftRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func BorderLeftRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func BorderLeftRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func BorderLeftSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func BorderLeftSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func BorderLeftSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func BorderLeftSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func BorderLeftSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func BorderLeftSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func BorderLeftSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func BorderLeftSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func BorderLeftSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func BorderLeftSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func BorderLeftSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func BorderLeftSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func BorderLeftTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func BorderLeftTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func BorderLeftTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func BorderLeftTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func BorderLeftTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func BorderLeftTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func BorderLeftVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func BorderLeftVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func BorderLeftVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func BorderLeftWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func BorderLeftYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func BorderLeftYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func BorderLeftYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func BorderLeftYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func BorderLeftColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderLeftColorProp): value.Value(),
		}
	}
}

//--

func DivideInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): "inherit",
		}
	}
}

func DivideCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): "currentColor",
		}
	}
}

func DivideTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): "transparent",
		}
	}
}

func DivideBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func DivideWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func DivideAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func DivideAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func DivideAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func DivideAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func DivideBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func DivideBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func DivideBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func DivideBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func DivideBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func DivideBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func DivideBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func DivideBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func DivideBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func DivideBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func DivideBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func DivideBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func DivideBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func DivideCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func DivideCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func DivideCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func DivideCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func DivideCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func DivideCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func DivideCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func DivideCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func DivideGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func DivideGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func DivideGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func DivideGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func DivideGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func DivideGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func DivideGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func DivideGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func DivideGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func DivideGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func DivideGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func DivideGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func DivideGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func DivideGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func DivideGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func DivideGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func DivideIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func DivideIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func DivideIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func DivideIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func DivideIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func DivideIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func DivideIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func DivideIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func DivideJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func DivideJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func DivideJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func DivideJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func DivideLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func DivideLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func DivideLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func DivideLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func DivideMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func DivideMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func DivideMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func DivideMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func DivideMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func DivideMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func DivideMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func DivideMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func DivideOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func DivideOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func DivideOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func DivideOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func DivideOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func DivideOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func DivideOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func DivideOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func DividePink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func DividePinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func DividePinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func DividePinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func DividePlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func DividePlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func DividePlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func DividePlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func DividePurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func DividePurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func DividePurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func DividePurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func DivideRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func DivideRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func DivideRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func DivideRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func DivideRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func DivideRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func DivideRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func DivideRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func DivideSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func DivideSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func DivideSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func DivideSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func DivideSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func DivideSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func DivideSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func DivideSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func DivideSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func DivideSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func DivideSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func DivideSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func DivideSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func DivideSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func DivideSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func DivideSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func DivideTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func DivideTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func DivideTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func DivideTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func DivideTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func DivideTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func DivideTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func DivideTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func DivideViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func DivideVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func DivideVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func DivideVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func DivideWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func DivideYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func DivideYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func DivideYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func DivideYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}
func DivideColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			notLastChildProp(borderColorProp): value.Value(),
		}
	}
}
