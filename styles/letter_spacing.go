package styles

import (
	"github.com/raitucarp/gomix/theme"
	"github.com/raitucarp/gomix/value"
)

func TrackingTighter() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(letterSpacingProp): s.theme.UseVarKey(theme.Tracking, "tighter"),
		}
	}
}

func TrackingTight() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(letterSpacingProp): s.theme.UseVarKey(theme.Tracking, "tight"),
		}
	}
}

func TrackingNormal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(letterSpacingProp): s.theme.UseVarKey(theme.Tracking, "normal"),
		}
	}
}

func TrackingWide() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(letterSpacingProp): s.theme.UseVarKey(theme.Tracking, "wide"),
		}
	}
}

func TrackingWider() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(letterSpacingProp): s.theme.UseVarKey(theme.Tracking, "wider"),
		}
	}
}

func TrackingWidest() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(letterSpacingProp): s.theme.UseVarKey(theme.Tracking, "widest"),
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
