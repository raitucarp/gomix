package styles

import (
	"fmt"

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

func TrackingCustomProperty(customProperty string) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): fmt.Sprintf("var(--%s)", customProperty),
		}
	}
}

func TrackingValue(value CustomValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(letterSpacingProp): value.String(),
		}
	}
}
