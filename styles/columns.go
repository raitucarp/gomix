package styles

import (
	"github.com/raitucarp/gomix/theme"
	"github.com/raitucarp/gomix/value"
)

func Columns(cols value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): cols.Value(),
		}
	}
}

func Columns3xs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): s.theme.UseVarKey(theme.Container, "3xs"),
		}
	}
}

func Columns2xs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): s.theme.UseVarKey(theme.Container, "2xs"),
		}
	}
}

func ColumnsXs() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): s.theme.UseVarKey(theme.Container, "xs"),
		}
	}
}

func ColumnsSm() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): s.theme.UseVarKey(theme.Container, "sm"),
		}
	}
}

func ColumnsMd() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): s.theme.UseVarKey(theme.Container, "md"),
		}
	}
}

func ColumnsLg() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): s.theme.UseVarKey(theme.Container, "lg"),
		}
	}
}

func ColumnsXl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): s.theme.UseVarKey(theme.Container, "xl"),
		}
	}
}

func Columns2xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): s.theme.UseVarKey(theme.Container, "2xl"),
		}
	}
}

func Columns3xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): s.theme.UseVarKey(theme.Container, "3xl"),
		}
	}
}

func Columns4xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): s.theme.UseVarKey(theme.Container, "4xl"),
		}
	}
}

func Columns5xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): s.theme.UseVarKey(theme.Container, "5xl"),
		}
	}
}

func Columns6xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): s.theme.UseVarKey(theme.Container, "6xl"),
		}
	}
}

func Columns7xl() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): s.theme.UseVarKey(theme.Container, "7xl"),
		}
	}
}

func ColumnsAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(columnsProp): "auto",
		}
	}
}
