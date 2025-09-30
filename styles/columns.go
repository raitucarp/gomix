package styles

import (
	"github.com/raitucarp/gomix/themes"
)

func Columns(cols customValue) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): cols.Value(),
		}
	}
}

func Columns3xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): s.theme.UseVarKey(themes.Container, "3xs"),
		}
	}
}

func Columns2xs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): s.theme.UseVarKey(themes.Container, "2xs"),
		}
	}
}

func ColumnsXs() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): s.theme.UseVarKey(themes.Container, "xs"),
		}
	}
}

func ColumnsSm() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): s.theme.UseVarKey(themes.Container, "sm"),
		}
	}
}

func ColumnsMd() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): s.theme.UseVarKey(themes.Container, "md"),
		}
	}
}

func ColumnsLg() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): s.theme.UseVarKey(themes.Container, "lg"),
		}
	}
}

func ColumnsXl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): s.theme.UseVarKey(themes.Container, "xl"),
		}
	}
}

func Columns2xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): s.theme.UseVarKey(themes.Container, "2xl"),
		}
	}
}

func Columns3xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): s.theme.UseVarKey(themes.Container, "3xl"),
		}
	}
}

func Columns4xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): s.theme.UseVarKey(themes.Container, "4xl"),
		}
	}
}

func Columns5xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): s.theme.UseVarKey(themes.Container, "5xl"),
		}
	}
}

func Columns6xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): s.theme.UseVarKey(themes.Container, "6xl"),
		}
	}
}

func Columns7xl() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): s.theme.UseVarKey(themes.Container, "7xl"),
		}
	}
}

func ColumnsAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(columnsProp): "auto",
		}
	}
}
