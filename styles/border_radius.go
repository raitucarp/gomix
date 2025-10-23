package styles

import (
	"github.com/raitucarp/gomix/value"
)

func RoundedXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func Rounded2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func Rounded3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func Rounded4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderRadiusProp): "0",
		}
	}
}

func RoundedFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderRadiusProp): "calc(infinity+1px)",
		}
	}
}

func Rounded(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderRadiusProp): val.Value(),
		}
	}
}

func RoundedSXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedSSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedSMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedSLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedSXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedS2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedS3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedS4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedSNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): "0",
			string(borderEndStartRadiusProp):   "0",
		}
	}
}

func RoundedSFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): "calc(infinity+1px)",
			string(borderEndStartRadiusProp):   "calc(infinity+1px)",
		}
	}
}

func RoundedS(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		val.Value()
		return &Properties{
			string(borderStartStartRadiusProp): val.Value(),
			string(borderEndStartRadiusProp):   val.Value(),
		}
	}
}

//--

func RoundedEXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedESm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedEMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedELg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedEXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedE2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedE3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedE4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedENone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): "0",
			string(borderEndEndRadiusProp):   "0",
		}
	}
}

func RoundedEFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): "calc(infinity+1px)",
			string(borderEndEndRadiusProp):   "calc(infinity+1px)",
		}
	}
}

func RoundedE(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): val.Value(),
			string(borderEndEndRadiusProp):   val.Value(),
		}
	}
}

// --

func RoundedTXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "xs"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedTSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "sm"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedTMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "md"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedTLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "lg"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedTXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "xl"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedT2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "2xl"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedT3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "3xl"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedT4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "4xl"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedTNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):  "0",
			string(borderTopRightRadiusProp): "0",
		}
	}
}

func RoundedTFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):  "calc(infinity+1px)",
			string(borderTopRightRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedT(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):  val.Value(),
			string(borderTopRightRadiusProp): val.Value(),
		}
	}
}

// --

func RoundedRXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themeRadius, "xs"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedRSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themeRadius, "sm"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedRMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themeRadius, "md"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedRLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themeRadius, "lg"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedRXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themeRadius, "xl"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedR2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themeRadius, "2xl"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedR3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themeRadius, "3xl"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedR4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themeRadius, "4xl"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedRNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp):    "0",
			string(borderBottomRightRadiusProp): "0",
		}
	}
}

func RoundedRFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp):    "calc(infinity+1px)",
			string(borderBottomRightRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedR(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp):    val.Value(),
			string(borderBottomRightRadiusProp): val.Value(),
		}
	}
}

// --

func RoundedBXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedBSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedBMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedBLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedBXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedB2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedB3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedB4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedBNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): "0",
			string(borderBottomLeftRadiusProp):  "0",
		}
	}
}

func RoundedBFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): "calc(infinity+1px)",
			string(borderBottomLeftRadiusProp):  "calc(infinity+1px)",
		}
	}
}

func RoundedB(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): val.Value(),
			string(borderBottomLeftRadiusProp):  val.Value(),
		}
	}
}

// --

func RoundedLXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themeRadius, "xs"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedLSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themeRadius, "sm"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedLMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themeRadius, "md"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedLLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themeRadius, "lg"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedLXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themeRadius, "xl"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedL2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themeRadius, "2xl"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedL3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themeRadius, "3xl"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedL4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themeRadius, "4xl"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedLNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):    "0",
			string(borderBottomLeftRadiusProp): "0",
		}
	}
}

func RoundedLFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):    "calc(infinity+1px)",
			string(borderBottomLeftRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedL(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp):    val.Value(),
			string(borderBottomLeftRadiusProp): val.Value(),
		}
	}
}

// --
func RoundedSsXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedSsSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedSsMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedSsLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedSsXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedSs2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedSs3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedSs4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedSsNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): "0",
		}
	}
}

func RoundedSsFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedSs(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartStartRadiusProp): val.Value(),
		}
	}
}

// --

func RoundedSeXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedSeSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedSeMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedSeLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedSeXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedSe2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedSe3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedSe4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedSeNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): "0",
		}
	}
}

func RoundedSeFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedSe(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderStartEndRadiusProp): val.Value(),
		}
	}
}

// --

func RoundedEeXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedEeSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedEeMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedEeLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedEeXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedEe2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedEe3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedEe4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedEeNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndEndRadiusProp): "0",
		}
	}
}

func RoundedEeFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndEndRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedEe(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndEndRadiusProp): val.Value(),
		}
	}
}

// --

func RoundedEsXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedEsSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedEsMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedEsLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedEsXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedEs2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedEs3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedEs4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedEsNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndStartRadiusProp): "0",
		}
	}
}

func RoundedEsFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndStartRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedEs(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderEndStartRadiusProp): val.Value(),
		}
	}
}

// --

func RoundedTlXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedTlSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedTlMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedTlLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedTlXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedTl2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedTl3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedTl4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedTlNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp): "0",
		}
	}
}

func RoundedTlFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedTl(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopLeftRadiusProp): val.Value(),
		}
	}
}

// --

func RoundedTrXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedTrSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedTrMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedTrLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedTrXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedTr2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedTr3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedTr4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedTrNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp): "0",
		}
	}
}

func RoundedTrFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedTr(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderTopRightRadiusProp): val.Value(),
		}
	}
}

// --

func RoundedBrXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedBrSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedBrMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedBrLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedBrXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedBr2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedBr3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedBr4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedBrNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): "0",
		}
	}
}

func RoundedBrFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedBr(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomRightRadiusProp): val.Value(),
		}
	}
}

// --

func RoundedBlXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "xs"),
		}
	}
}

func RoundedBlSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "sm"),
		}
	}
}

func RoundedBlMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "md"),
		}
	}
}

func RoundedBlLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "lg"),
		}
	}
}

func RoundedBlXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "xl"),
		}
	}
}

func RoundedBl2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "2xl"),
		}
	}
}

func RoundedBl3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "3xl"),
		}
	}
}

func RoundedBl4Xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themeRadius, "4xl"),
		}
	}
}

func RoundedBlNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomLeftRadiusProp): "0",
		}
	}
}

func RoundedBlFull() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomLeftRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedBl(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(borderBottomLeftRadiusProp): val.Value(),
		}
	}
}
