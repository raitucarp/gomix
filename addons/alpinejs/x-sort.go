package alpinejs

import "encoding/json"

type xSort struct {
	alpineRef *alpine
	isItem    bool
	position  int
	handle    string
	group     string
	isGhost   bool
	asHandle  bool
	config    string
}

func (a *alpine) Sort() *xSort {
	return &xSort{alpineRef: a, position: -1}
}

func (x *xSort) Handle(handle string) *xSort {
	x.handle = handle
	return x
}

func (x *xSort) Group(groupName string) *xSort {
	x.group = groupName
	return x
}

func (x *xSort) Item() *xSort {
	x.isItem = true
	return x
}

func (x *xSort) Position(position int) *xSort {
	x.position = position
	return x
}

func (x *xSort) Ghost() *xSort {
	x.isGhost = true
	return x
}

func (x *xSort) Config(config map[string]any) *xSort {
	c, _ := json.Marshal(config)
	x.config = string(c)

	return x
}

func (x *xSort) Attribute() *alpine {
	attributeName := "x-sort"

	if x.isItem {
		if x.position > -1 {
			x.alpineRef.el.AddAttribute(attributeName+":item", x.handle)
			return x.alpineRef
		}

		x.alpineRef.el.AddAttributeBool(attributeName + ":item")
		return x.alpineRef
	} else if x.asHandle {
		x.alpineRef.el.AddAttributeBool(attributeName + ":handle")
	} else {

		if x.isGhost {
			attributeName += ".ghost"
		}

		if x.config != "" {
			x.alpineRef.el.AddAttribute(attributeName+":config", x.config)
		}

		if x.group != "" {
			x.alpineRef.el.AddAttribute(attributeName+":group", x.group)
		}

		if x.handle != "" {
			x.alpineRef.el.AddAttribute(attributeName, x.handle)
			return x.alpineRef
		}

		x.alpineRef.el.AddAttributeBool(attributeName)
	}

	return x.alpineRef
}
