package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
)

func OutlineInherit() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): "inherit",
		}
	}
}

func OutlineCurrent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): "currentColor",
		}
	}
}

func OutlineTransparent() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): "transparent",
		}
	}
}

func OutlineBlack() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func OutlineWhite() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}

func OutlineAmber(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func OutlineAmberAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func OutlineAmberDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func OutlineAmberDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func OutlineBlackAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func OutlineBlue(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func OutlineBlueAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func OutlineBlueDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func OutlineBlueDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func OutlineBronze(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func OutlineBronzeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func OutlineBronzeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func OutlineBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func OutlineBrown(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func OutlineBrownAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func OutlineBrownDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func OutlineBrownDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func OutlineCrimson(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func OutlineCrimsonAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func OutlineCrimsonDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func OutlineCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func OutlineCyan(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func OutlineCyanAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func OutlineCyanDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func OutlineCyanDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func OutlineGold(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func OutlineGoldAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func OutlineGoldDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func OutlineGoldDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func OutlineGrass(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func OutlineGrassAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func OutlineGrassDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func OutlineGrassDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func OutlineGray(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func OutlineGrayAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func OutlineGrayDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func OutlineGrayDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func OutlineGreen(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func OutlineGreenAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func OutlineGreenDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func OutlineGreenDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func OutlineIndigo(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func OutlineIndigoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func OutlineIndigoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func OutlineIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func OutlineIris(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func OutlineIrisAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func OutlineIrisDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func OutlineIrisDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func OutlineJade(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func OutlineJadeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func OutlineJadeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func OutlineJadeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func OutlineLime(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func OutlineLimeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func OutlineLimeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func OutlineLimeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func OutlineMauve(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func OutlineMauveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func OutlineMauveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func OutlineMauveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func OutlineMint(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func OutlineMintAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func OutlineMintDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func OutlineMintDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func OutlineOlive(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func OutlineOliveAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func OutlineOliveDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func OutlineOliveDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func OutlineOrange(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func OutlineOrangeAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func OutlineOrangeDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func OutlineOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func OutlinePink(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func OutlinePinkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func OutlinePinkDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func OutlinePinkDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func OutlinePlum(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func OutlinePlumAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func OutlinePlumDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func OutlinePlumDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func OutlinePurple(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func OutlinePurpleAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func OutlinePurpleDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func OutlinePurpleDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func OutlineRed(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func OutlineRedAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func OutlineRedDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func OutlineRedDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func OutlineRuby(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func OutlineRubyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func OutlineRubyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func OutlineRubyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func OutlineSage(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func OutlineSageAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func OutlineSageDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func OutlineSageDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func OutlineSand(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func OutlineSandAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func OutlineSandDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func OutlineSandDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func OutlineSky(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func OutlineSkyAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func OutlineSkyDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func OutlineSkyDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func OutlineSlate(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func OutlineSlateAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func OutlineSlateDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func OutlineSlateDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func OutlineTeal(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func OutlineTealAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func OutlineTealDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func OutlineTealDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func OutlineTomato(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func OutlineTomatoAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func OutlineTomatoDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func OutlineTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func OutlineViolet(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func OutlineVioletAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func OutlineVioletDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func OutlineVioletDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func OutlineWhiteAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func OutlineYellow(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func OutlineYellowAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func OutlineYellowDark(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func OutlineYellowDarkAlpha(scale int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func OutlineColor(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(outlineColorProp): value.Value(),
		}
	}
}
