package styles

import "github.com/raitucarp/gomix/themes"

func RoundedXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func Rounded2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func Rounded3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func Rounded4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRadiusProp): "0",
		}
	}
}

func RoundedFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRadiusProp): "calc(infinity+1px)",
		}
	}
}

func Rounded(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderRadiusProp): value.Value(),
		}
	}
}

func RoundedSXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedSSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedSMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedSLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedSXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedS2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedS3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedS4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
			string(borderEndStartRadiusProp):   s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedSNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): "0",
			string(borderEndStartRadiusProp):   "0",
		}
	}
}

func RoundedSFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): "calc(infinity+1px)",
			string(borderEndStartRadiusProp):   "calc(infinity+1px)",
		}
	}
}

func RoundedS(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): value.Value(),
			string(borderEndStartRadiusProp):   value.Value(),
		}
	}
}

//--

func RoundedEXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedESm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedEMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedELg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedEXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedE2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedE3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedE4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
			string(borderEndEndRadiusProp):   s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedENone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): "0",
			string(borderEndEndRadiusProp):   "0",
		}
	}
}

func RoundedEFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): "calc(infinity+1px)",
			string(borderEndEndRadiusProp):   "calc(infinity+1px)",
		}
	}
}

func RoundedE(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): value.Value(),
			string(borderEndEndRadiusProp):   value.Value(),
		}
	}
}

// --

func RoundedTXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "xs"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedTSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "sm"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedTMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "md"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedTLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "lg"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedTXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "xl"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedT2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "2xl"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedT3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "3xl"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedT4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "4xl"),
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedTNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):  "0",
			string(borderTopRightRadiusProp): "0",
		}
	}
}

func RoundedTFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):  "calc(infinity+1px)",
			string(borderTopRightRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedT(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):  value.Value(),
			string(borderTopRightRadiusProp): value.Value(),
		}
	}
}

// --

func RoundedRXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themes.Radius, "xs"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedRSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themes.Radius, "sm"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedRMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themes.Radius, "md"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedRLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themes.Radius, "lg"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedRXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themes.Radius, "xl"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedR2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themes.Radius, "2xl"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedR3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themes.Radius, "3xl"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedR4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp):    s.theme.UseVarKey(themes.Radius, "4xl"),
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedRNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp):    "0",
			string(borderBottomRightRadiusProp): "0",
		}
	}
}

func RoundedRFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp):    "calc(infinity+1px)",
			string(borderBottomRightRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedR(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp):    value.Value(),
			string(borderBottomRightRadiusProp): value.Value(),
		}
	}
}

// --

func RoundedBXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedBSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedBMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedBLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedBXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedB2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedB3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedB4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
			string(borderBottomLeftRadiusProp):  s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedBNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): "0",
			string(borderBottomLeftRadiusProp):  "0",
		}
	}
}

func RoundedBFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): "calc(infinity+1px)",
			string(borderBottomLeftRadiusProp):  "calc(infinity+1px)",
		}
	}
}

func RoundedB(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): value.Value(),
			string(borderBottomLeftRadiusProp):  value.Value(),
		}
	}
}

// --

func RoundedLXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themes.Radius, "xs"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedLSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themes.Radius, "sm"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedLMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themes.Radius, "md"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedLLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themes.Radius, "lg"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedLXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themes.Radius, "xl"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedL2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themes.Radius, "2xl"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedL3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themes.Radius, "3xl"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedL4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):    s.theme.UseVarKey(themes.Radius, "4xl"),
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedLNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):    "0",
			string(borderBottomLeftRadiusProp): "0",
		}
	}
}

func RoundedLFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):    "calc(infinity+1px)",
			string(borderBottomLeftRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedL(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp):    value.Value(),
			string(borderBottomLeftRadiusProp): value.Value(),
		}
	}
}

// --
func RoundedSsXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedSsSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedSsMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedSsLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedSsXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedSs2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedSs3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedSs4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedSsNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): "0",
		}
	}
}

func RoundedSsFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedSs(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartStartRadiusProp): value.Value(),
		}
	}
}

// --

func RoundedSeXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedSeSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedSeMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedSeLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedSeXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedSe2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedSe3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedSe4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedSeNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): "0",
		}
	}
}

func RoundedSeFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedSe(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderStartEndRadiusProp): value.Value(),
		}
	}
}

// --

func RoundedEeXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedEeSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedEeMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedEeLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedEeXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedEe2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedEe3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedEe4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndEndRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedEeNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndEndRadiusProp): "0",
		}
	}
}

func RoundedEeFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndEndRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedEe(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndEndRadiusProp): value.Value(),
		}
	}
}

// --

func RoundedEsXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedEsSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedEsMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedEsLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedEsXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedEs2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedEs3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedEs4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndStartRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedEsNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndStartRadiusProp): "0",
		}
	}
}

func RoundedEsFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndStartRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedEs(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderEndStartRadiusProp): value.Value(),
		}
	}
}

// --

func RoundedTlXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedTlSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedTlMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedTlLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedTlXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedTl2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedTl3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedTl4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedTlNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp): "0",
		}
	}
}

func RoundedTlFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedTl(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopLeftRadiusProp): value.Value(),
		}
	}
}

// --

func RoundedTrXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedTrSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedTrMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedTrLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedTrXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedTr2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedTr3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedTr4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedTrNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp): "0",
		}
	}
}

func RoundedTrFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedTr(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderTopRightRadiusProp): value.Value(),
		}
	}
}

// --

func RoundedBrXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedBrSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedBrMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedBrLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedBrXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedBr2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedBr3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedBr4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedBrNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): "0",
		}
	}
}

func RoundedBrFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedBr(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomRightRadiusProp): value.Value(),
		}
	}
}

// --

func RoundedBlXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "xs"),
		}
	}
}

func RoundedBlSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "sm"),
		}
	}
}

func RoundedBlMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "md"),
		}
	}
}

func RoundedBlLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "lg"),
		}
	}
}

func RoundedBlXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "xl"),
		}
	}
}

func RoundedBl2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "2xl"),
		}
	}
}

func RoundedBl3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "3xl"),
		}
	}
}

func RoundedBl4Xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomLeftRadiusProp): s.theme.UseVarKey(themes.Radius, "4xl"),
		}
	}
}

func RoundedBlNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomLeftRadiusProp): "0",
		}
	}
}

func RoundedBlFull() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomLeftRadiusProp): "calc(infinity+1px)",
		}
	}
}

func RoundedBl(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(borderBottomLeftRadiusProp): value.Value(),
		}
	}
}
