package components

import "strings"

func (c *Component) addHtmxAttribute(name string, value string) {
	c.addAttribute(strings.Join([]string{"hx", name}, "-"), value)
}

func (c *Component) HxGet(url string) *Component {
	c.addHtmxAttribute("get", url)
	return c
}

func (c *Component) HxTarget(id string) *Component {
	c.addHtmxAttribute("target", id)
	return c
}
