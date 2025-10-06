package styles

import "github.com/raitucarp/gomix/value"

func CursorAuto() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "auto",
		}

	}
}
func CursorDefault() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "default",
		}

	}
}
func CursorPointer() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "pointer",
		}

	}
}
func CursorWait() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "wait",
		}

	}
}
func CursorText() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "text",
		}

	}
}
func CursorMove() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "move",
		}

	}
}
func CursorHelp() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "help",
		}

	}
}
func CursorNotAllowed() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "not-allowed",
		}

	}
}
func CursorNone() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "none",
		}

	}
}
func CursorContextMenu() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "context-menu",
		}

	}
}
func CursorProgress() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "progress",
		}

	}
}
func CursorCell() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "cell",
		}

	}
}
func CursorCrosshair() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "crosshair",
		}

	}
}
func CursorVerticalText() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "vertical-text",
		}

	}
}
func CursorAlias() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "alias",
		}

	}
}
func CursorCopy() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "copy",
		}

	}
}
func CursorNoDrop() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "no-drop",
		}

	}
}
func CursorGrab() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "grab",
		}

	}
}
func CursorGrabbing() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "grabbing",
		}

	}
}
func CursorAllScroll() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "all-scroll",
		}

	}
}
func CursorColResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "col-resize",
		}

	}
}
func CursorRowResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "row-resize",
		}

	}
}
func CursorNResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "n-resize",
		}

	}
}
func CursorEResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "e-resize",
		}

	}
}
func CursorSResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "s-resize",
		}

	}
}
func CursorWResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "w-resize",
		}

	}
}
func CursorNeResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "ne-resize",
		}

	}
}
func CursorNwResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "nw-resize",
		}

	}
}
func CursorSeResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "se-resize",
		}

	}
}
func CursorSwResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "sw-resize",
		}

	}
}
func CursorEwResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "ew-resize",
		}

	}
}
func CursorNsResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "ns-resize",
		}

	}
}
func CursorNeswResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "nesw-resize",
		}

	}
}
func CursorNwseResize() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "nwse-resize",
		}

	}
}
func CursorZoomIn() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "zoom-in",
		}

	}
}
func CursorZoomOut() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): "zoom-out",
		}

	}
}
func Cursor(val value.Value) ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(cursorProp): val.Value(),
		}

	}
}
