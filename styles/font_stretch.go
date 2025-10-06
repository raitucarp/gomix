package styles

import (
	"fmt"

	"github.com/raitucarp/gomix/value"
)

func FontStretchUltraCondensed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "ultra-condensed",
		}
	}
}

func FontStretchExtraCondensed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "extra-condensed",
		}
	}
}

func FontStretchCondensed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "condensed",
		}
	}
}

func FontStretchSemiCondensed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "semi-condensed",
		}
	}
}

func FontStretchNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "normal",
		}
	}
}

func FontStretchSemiExpanded() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "semi-expanded",
		}
	}
}

func FontStretchExpanded() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "expanded",
		}
	}
}

func FontStretchExtraExpanded() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "extra-expanded",
		}
	}
}

func FontStretchUltraExpanded() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): "ultra-expanded",
		}
	}
}

func FontStretchPercentage(percentage int) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): fmt.Sprintf("%d%%", percentage),
		}
	}
}

func FontStretchBy(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontStretchProp): val.Value(),
		}
	}
}
