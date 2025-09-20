package hyperscript

type stringer interface {
	String() string
}

func String(s stringer) string {
	return s.String()
}
