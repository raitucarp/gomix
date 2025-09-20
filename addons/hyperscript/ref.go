package hyperscript

type refParam interface {
	String() string
	IsRef()
}
