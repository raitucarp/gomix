package styles

import "github.com/raitucarp/gomix/value"

func AlignBaseline() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(verticalAlignProp): "baseline",
		}
	}
}

func AlignTop() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(verticalAlignProp): "top",
		}
	}
}

func AlignMiddle() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(verticalAlignProp): "middle",
		}
	}
}

func AlignBottom() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(verticalAlignProp): "bottom",
		}
	}
}

func AlignTextTop() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(verticalAlignProp): "text-top",
		}
	}
}

func AlignTextBottom() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(verticalAlignProp): "text-bottom",
		}
	}
}

func AlignSub() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(verticalAlignProp): "sub",
		}
	}

}

func AlignSuper() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(verticalAlignProp): "super",
		}
	}
}

func AlignBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(verticalAlignProp): val.Value(),
		}
	}
}
