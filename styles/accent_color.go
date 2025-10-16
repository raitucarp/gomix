package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/theme"
	"github.com/raitucarp/gomix/value"
)

func AccentInherit() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): "inherit",
		}
	}
}

func AccentCurrent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): "currentColor",
		}
	}
}

func AccentTransparent() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): "transparent",
		}
	}
}

func AccentBlack() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, "black"),
		}
	}
}

func AccentWhite() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, "white"),
		}
	}
}
func AccentAmber(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-%d", scale)),
		}
	}
}
func AccentAmberAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-alpha-%d", scale)),
		}
	}
}
func AccentAmberDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-%d", scale)),
		}
	}
}
func AccentAmberDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("amber-dark-alpha-%d", scale)),
		}
	}
}
func AccentBlackAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("black-alpha-%d", scale)),
		}
	}
}
func AccentBlue(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-%d", scale)),
		}
	}
}
func AccentBlueAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-alpha-%d", scale)),
		}
	}
}
func AccentBlueDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-%d", scale)),
		}
	}
}
func AccentBlueDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("blue-dark-alpha-%d", scale)),
		}
	}
}
func AccentBronze(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-%d", scale)),
		}
	}
}
func AccentBronzeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-alpha-%d", scale)),
		}
	}
}
func AccentBronzeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-%d", scale)),
		}
	}
}
func AccentBronzeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("bronze-dark-alpha-%d", scale)),
		}
	}
}
func AccentBrown(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-%d", scale)),
		}
	}
}
func AccentBrownAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-alpha-%d", scale)),
		}
	}
}
func AccentBrownDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-%d", scale)),
		}
	}
}
func AccentBrownDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("brown-dark-alpha-%d", scale)),
		}
	}
}
func AccentCrimson(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-%d", scale)),
		}
	}
}
func AccentCrimsonAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-alpha-%d", scale)),
		}
	}
}
func AccentCrimsonDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-%d", scale)),
		}
	}
}
func AccentCrimsonDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("crimson-dark-alpha-%d", scale)),
		}
	}
}
func AccentCyan(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-%d", scale)),
		}
	}
}
func AccentCyanAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-alpha-%d", scale)),
		}
	}
}
func AccentCyanDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-%d", scale)),
		}
	}
}
func AccentCyanDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("cyan-dark-alpha-%d", scale)),
		}
	}
}
func AccentGold(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-%d", scale)),
		}
	}
}
func AccentGoldAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-alpha-%d", scale)),
		}
	}
}
func AccentGoldDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-%d", scale)),
		}
	}
}
func AccentGoldDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gold-dark-alpha-%d", scale)),
		}
	}
}
func AccentGrass(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-%d", scale)),
		}
	}
}
func AccentGrassAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-alpha-%d", scale)),
		}
	}
}
func AccentGrassDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-%d", scale)),
		}
	}
}
func AccentGrassDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("grass-dark-alpha-%d", scale)),
		}
	}
}
func AccentGray(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-%d", scale)),
		}
	}
}
func AccentGrayAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-alpha-%d", scale)),
		}
	}
}
func AccentGrayDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-%d", scale)),
		}
	}
}
func AccentGrayDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("gray-dark-alpha-%d", scale)),
		}
	}
}
func AccentGreen(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-%d", scale)),
		}
	}
}
func AccentGreenAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-alpha-%d", scale)),
		}
	}
}
func AccentGreenDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-%d", scale)),
		}
	}
}
func AccentGreenDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("green-dark-alpha-%d", scale)),
		}
	}
}
func AccentIndigo(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-%d", scale)),
		}
	}
}
func AccentIndigoAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-alpha-%d", scale)),
		}
	}
}
func AccentIndigoDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-%d", scale)),
		}
	}
}
func AccentIndigoDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("indigo-dark-alpha-%d", scale)),
		}
	}
}
func AccentIris(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-%d", scale)),
		}
	}
}
func AccentIrisAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-alpha-%d", scale)),
		}
	}
}
func AccentIrisDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-%d", scale)),
		}
	}
}
func AccentIrisDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("iris-dark-alpha-%d", scale)),
		}
	}
}
func AccentJade(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-%d", scale)),
		}
	}
}
func AccentJadeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-alpha-%d", scale)),
		}
	}
}
func AccentJadeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-%d", scale)),
		}
	}
}
func AccentJadeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("jade-dark-alpha-%d", scale)),
		}
	}
}
func AccentLime(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-%d", scale)),
		}
	}
}
func AccentLimeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-alpha-%d", scale)),
		}
	}
}
func AccentLimeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-%d", scale)),
		}
	}
}
func AccentLimeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("lime-dark-alpha-%d", scale)),
		}
	}
}
func AccentMauve(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-%d", scale)),
		}
	}
}
func AccentMauveAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-alpha-%d", scale)),
		}
	}
}
func AccentMauveDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-%d", scale)),
		}
	}
}
func AccentMauveDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mauve-dark-alpha-%d", scale)),
		}
	}
}
func AccentMint(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-%d", scale)),
		}
	}
}
func AccentMintAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-alpha-%d", scale)),
		}
	}
}
func AccentMintDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-%d", scale)),
		}
	}
}
func AccentMintDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("mint-dark-alpha-%d", scale)),
		}
	}
}
func AccentOlive(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-%d", scale)),
		}
	}
}
func AccentOliveAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-alpha-%d", scale)),
		}
	}
}
func AccentOliveDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-%d", scale)),
		}
	}
}
func AccentOliveDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("olive-dark-alpha-%d", scale)),
		}
	}
}
func AccentOrange(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-%d", scale)),
		}
	}
}
func AccentOrangeAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-alpha-%d", scale)),
		}
	}
}
func AccentOrangeDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-%d", scale)),
		}
	}
}
func AccentOrangeDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("orange-dark-alpha-%d", scale)),
		}
	}
}
func AccentPink(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-%d", scale)),
		}
	}
}
func AccentPinkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-alpha-%d", scale)),
		}
	}
}
func AccentPinkDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-%d", scale)),
		}
	}
}
func AccentPinkDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("pink-dark-alpha-%d", scale)),
		}
	}
}
func AccentPlum(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-%d", scale)),
		}
	}
}
func AccentPlumAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-alpha-%d", scale)),
		}
	}
}
func AccentPlumDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-%d", scale)),
		}
	}
}
func AccentPlumDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("plum-dark-alpha-%d", scale)),
		}
	}
}
func AccentPurple(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-%d", scale)),
		}
	}
}
func AccentPurpleAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-alpha-%d", scale)),
		}
	}
}
func AccentPurpleDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-%d", scale)),
		}
	}
}
func AccentPurpleDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("purple-dark-alpha-%d", scale)),
		}
	}
}
func AccentRed(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-%d", scale)),
		}
	}
}
func AccentRedAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-alpha-%d", scale)),
		}
	}
}
func AccentRedDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-%d", scale)),
		}
	}
}
func AccentRedDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("red-dark-alpha-%d", scale)),
		}
	}
}
func AccentRuby(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-%d", scale)),
		}
	}
}
func AccentRubyAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-alpha-%d", scale)),
		}
	}
}
func AccentRubyDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-%d", scale)),
		}
	}
}
func AccentRubyDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("ruby-dark-alpha-%d", scale)),
		}
	}
}
func AccentSage(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-%d", scale)),
		}
	}
}
func AccentSageAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-alpha-%d", scale)),
		}
	}
}
func AccentSageDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-%d", scale)),
		}
	}
}
func AccentSageDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sage-dark-alpha-%d", scale)),
		}
	}
}
func AccentSand(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-%d", scale)),
		}
	}
}
func AccentSandAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-alpha-%d", scale)),
		}
	}
}
func AccentSandDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-%d", scale)),
		}
	}
}
func AccentSandDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sand-dark-alpha-%d", scale)),
		}
	}
}
func AccentSky(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-%d", scale)),
		}
	}
}
func AccentSkyAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-alpha-%d", scale)),
		}
	}
}
func AccentSkyDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-%d", scale)),
		}
	}
}
func AccentSkyDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("sky-dark-alpha-%d", scale)),
		}
	}
}
func AccentSlate(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-%d", scale)),
		}
	}
}
func AccentSlateAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-alpha-%d", scale)),
		}
	}
}
func AccentSlateDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-%d", scale)),
		}
	}
}
func AccentSlateDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("slate-dark-alpha-%d", scale)),
		}
	}
}
func AccentTeal(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-%d", scale)),
		}
	}
}
func AccentTealAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-alpha-%d", scale)),
		}
	}
}
func AccentTealDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-%d", scale)),
		}
	}
}
func AccentTealDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("teal-dark-alpha-%d", scale)),
		}
	}
}
func AccentTomato(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-%d", scale)),
		}
	}
}
func AccentTomatoAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-alpha-%d", scale)),
		}
	}
}
func AccentTomatoDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-%d", scale)),
		}
	}
}
func AccentTomatoDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("tomato-dark-alpha-%d", scale)),
		}
	}
}
func AccentViolet(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-%d", scale)),
		}
	}
}
func AccentVioletAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-alpha-%d", scale)),
		}
	}
}
func AccentVioletDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-%d", scale)),
		}
	}
}
func AccentVioletDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("violet-dark-alpha-%d", scale)),
		}
	}
}
func AccentWhiteAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("white-alpha-%d", scale)),
		}
	}
}
func AccentYellow(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-%d", scale)),
		}
	}
}
func AccentYellowAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-alpha-%d", scale)),
		}
	}
}
func AccentYellowDark(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-%d", scale)),
		}
	}
}
func AccentYellowDarkAlpha(scale int) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): s.theme.UseVarKey(theme.Color, fmt.Sprintf("yellow-dark-alpha-%d", scale)),
		}
	}
}

func Accent(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(accentColorProp): val.Value(),
		}
	}
}
