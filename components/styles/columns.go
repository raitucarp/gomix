package styles

import "strconv"

func (s *Styles) Columns(number int) *Styles {
	s.addProp(Columns, strconv.Itoa(number))
	return s
}

// func (s *Styles) Columns
