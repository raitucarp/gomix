package styles

import "github.com/raitucarp/gomix/value"

func AutoColsAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridAutoColumnsProp): "auto",
		}
	}
}

func AutoColsMin() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridAutoColumnsProp): "min-content",
		}
	}
}

func AutoColsMax() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridAutoColumnsProp): "max-content",
		}
	}
}

func AutoColsFr() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridAutoColumnsProp): "min-max(0,1fr)",
		}
	}
}

func AutoColsBy(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(gridAutoColumnsProp): val.Value(),
		}
	}
}
