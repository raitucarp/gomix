package styles

import (
	"github.com/raitucarp/gomix/value"
)

func TrackingTighter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(letterSpacingProp): s.theme.UseVarKey(themeTracking, "tighter"),
		}
	}
}

func TrackingTight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(letterSpacingProp): s.theme.UseVarKey(themeTracking, "tight"),
		}
	}
}

func TrackingNormal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(letterSpacingProp): s.theme.UseVarKey(themeTracking, "normal"),
		}
	}
}

func TrackingWide() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(letterSpacingProp): s.theme.UseVarKey(themeTracking, "wide"),
		}
	}
}

func TrackingWider() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(letterSpacingProp): s.theme.UseVarKey(themeTracking, "wider"),
		}
	}
}

func TrackingWidest() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(letterSpacingProp): s.theme.UseVarKey(themeTracking, "widest"),
		}
	}
}

func TrackingBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(letterSpacingProp): val.Value(),
		}
	}
}
