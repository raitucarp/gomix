package styles

import (
	"github.com/raitucarp/gomix/themes"
)

func TrackingTighter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): s.theme.UseVarKey(themes.Tracking, "tighter"),
		}
	}
}

func TrackingTight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): s.theme.UseVarKey(themes.Tracking, "tight"),
		}
	}
}

func TrackingNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): s.theme.UseVarKey(themes.Tracking, "normal"),
		}
	}
}

func TrackingWide() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): s.theme.UseVarKey(themes.Tracking, "wide"),
		}
	}
}

func TrackingWider() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): s.theme.UseVarKey(themes.Tracking, "wider"),
		}
	}
}

func TrackingWidest() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): s.theme.UseVarKey(themes.Tracking, "widest"),
		}
	}
}

func TrackingBy(value customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): value.Value(),
		}
	}
}
