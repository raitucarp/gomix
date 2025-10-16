package styles

func NormalNums() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontVariantNumericProp): "normal",
		}
	}
}

func Ordinal() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontVariantNumericProp): "ordinal",
		}
	}
}

func SlashedZero() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontVariantNumericProp): "slashed-zero",
		}
	}
}

func LiningNums() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontVariantNumericProp): "lining-nums",
		}
	}
}

func OldStyleNums() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontVariantNumericProp): "oldstyle-nums",
		}
	}
}

func ProportionalNums() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontVariantNumericProp): "proportional-nums",
		}
	}
}

func TabularNums() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontVariantNumericProp): "tabular-nums",
		}
	}
}

func DiagonalFractions() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontVariantNumericProp): "diagonal-fractions",
		}
	}
}

func StackedFractions() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(fontVariantNumericProp): "stacked-fractions",
		}
	}
}
