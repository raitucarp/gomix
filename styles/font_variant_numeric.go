package styles

func NormalNums() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontVariantNumericProp): "normal",
		}
	}
}

func Ordinal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontVariantNumericProp): "ordinal",
		}
	}
}

func SlashedZero() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontVariantNumericProp): "slashed-zero",
		}
	}
}

func LiningNums() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontVariantNumericProp): "lining-nums",
		}
	}
}

func OldStyleNums() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontVariantNumericProp): "oldstyle-nums",
		}
	}
}

func ProportionalNums() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontVariantNumericProp): "proportional-nums",
		}
	}
}

func TabularNums() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontVariantNumericProp): "tabular-nums",
		}
	}
}

func DiagonalFractions() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontVariantNumericProp): "diagonal-fractions",
		}
	}
}

func StackedFractions() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(fontVariantNumericProp): "stacked-fractions",
		}
	}
}
