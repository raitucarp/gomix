package styles

import "github.com/raitucarp/gomix/value"

func CursorAuto() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "auto",
		}

	}
}
func CursorDefault() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "default",
		}

	}
}
func CursorPointer() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "pointer",
		}

	}
}
func CursorWait() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "wait",
		}

	}
}
func CursorText() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "text",
		}

	}
}
func CursorMove() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "move",
		}

	}
}
func CursorHelp() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "help",
		}

	}
}
func CursorNotAllowed() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "not-allowed",
		}

	}
}
func CursorNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "none",
		}

	}
}
func CursorContextMenu() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "context-menu",
		}

	}
}
func CursorProgress() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "progress",
		}

	}
}
func CursorCell() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "cell",
		}

	}
}
func CursorCrosshair() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "crosshair",
		}

	}
}
func CursorVerticalText() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "vertical-text",
		}

	}
}
func CursorAlias() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "alias",
		}

	}
}
func CursorCopy() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "copy",
		}

	}
}
func CursorNoDrop() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "no-drop",
		}

	}
}
func CursorGrab() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "grab",
		}

	}
}
func CursorGrabbing() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "grabbing",
		}

	}
}
func CursorAllScroll() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "all-scroll",
		}

	}
}
func CursorColResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "col-resize",
		}

	}
}
func CursorRowResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "row-resize",
		}

	}
}
func CursorNResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "n-resize",
		}

	}
}
func CursorEResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "e-resize",
		}

	}
}
func CursorSResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "s-resize",
		}

	}
}
func CursorWResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "w-resize",
		}

	}
}
func CursorNeResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "ne-resize",
		}

	}
}
func CursorNwResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "nw-resize",
		}

	}
}
func CursorSeResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "se-resize",
		}

	}
}
func CursorSwResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "sw-resize",
		}

	}
}
func CursorEwResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "ew-resize",
		}

	}
}
func CursorNsResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "ns-resize",
		}

	}
}
func CursorNeswResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "nesw-resize",
		}

	}
}
func CursorNwseResize() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "nwse-resize",
		}

	}
}
func CursorZoomIn() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "zoom-in",
		}

	}
}
func CursorZoomOut() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): "zoom-out",
		}

	}
}
func Cursor(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(cursorProp): val.Value(),
		}

	}
}
