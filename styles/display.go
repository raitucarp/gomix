package styles

func Inline() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "inline",
		}
	}
}

func Block() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "block",
		}
	}
}

func InlineBlock() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "inline-block",
		}
	}
}

func FlowRoot() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "flow-root",
		}
	}
}

func Flex() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "flex",
		}
	}
}

func InlineFlex() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "inline-flex",
		}
	}
}

func Grid() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "grid",
		}
	}
}

func InlineGrid() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "inline-grid",
		}
	}
}

func Contents() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "contents",
		}
	}
}

func Table() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "table",
		}
	}
}

func InlineTable() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "inline-table",
		}
	}
}

func TableCaption() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "table-caption",
		}
	}
}

func TableCell() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "table-cell",
		}
	}
}

func TableColumn() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "table-column",
		}
	}
}

func TableColumnGroup() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "table-column-group",
		}
	}
}

func TableFooterGroup() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "table-footer-group",
		}
	}
}

func TableHeaderGroup() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "table-header-group",
		}
	}
}

func TableRowGroup() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "table-row-group",
		}
	}
}

func TableRow() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "table-row",
		}
	}
}

func ListItem() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "list-item",
		}
	}
}

func Hidden() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(displayProp): "none",
		}
	}
}

func SrOnly() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(positionProp):    "absolute",
			string(widthProp):       "1px",
			string(heightProp):      "1px",
			string(paddingProp):     "0",
			string(marginProp):      "0",
			string(overflowProp):    "overflow",
			string(clipProp):        "rect(0, 0, 0, 0)",
			string(whiteSpaceProp):  "nowrap",
			string(borderWidthProp): "0",
		}
	}
}

func NotSrOnly() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(positionProp):   "static",
			string(widthProp):      "auto",
			string(heightProp):     "auto",
			string(paddingProp):    "0",
			string(marginProp):     "0",
			string(overflowProp):   "visible",
			string(clipProp):       "auto",
			string(whiteSpaceProp): "normal",
		}
	}
}
