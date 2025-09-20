package hyperscript

type hs struct {
	script string
}

func HyperScript(script string) *hs {
	return &hs{script: script}
}

func (h *hs) String() string {
	return h.script
}


func Addon() {
	
}