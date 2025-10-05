package styles

import "github.com/raitucarp/gomix/theme"

func TrackingTighter() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): s.theme.UseVarKey(theme.Tracking, "tighter"),
		}
	}
}

func TrackingTight() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): s.theme.UseVarKey(theme.Tracking, "tight"),
		}
	}
}

func TrackingNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): s.theme.UseVarKey(theme.Tracking, "normal"),
		}
	}
}

func TrackingWide() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): s.theme.UseVarKey(theme.Tracking, "wide"),
		}
	}
}

func TrackingWider() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): s.theme.UseVarKey(theme.Tracking, "wider"),
		}
	}
}

func TrackingWidest() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): s.theme.UseVarKey(theme.Tracking, "widest"),
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
