package hyperscript

type Command interface {
	Command()
	String() string
}

