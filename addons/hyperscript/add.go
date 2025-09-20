package hyperscript

import "strings"

type AddCommand struct {
	source refParam
	to     refParam
	where  refParam
}

func (c *AddCommand) To(to refParam) *AddCommand {
	c.to = to
	return c
}

func (c *AddCommand) Where(where refParam) *AddCommand {
	c.where = where
	return c
}

func (c *AddCommand) Command() {}
func (c *AddCommand) String() string {
	s := []string{"add"}
	s = append(s, c.source.String())

	if c.to != nil {
		s = append(s, "to", c.to.String())
	}

	if c.where != nil {
		s = append(s, "where", c.where.String())
	}
	return strings.Join(s, " ")
}

func Add(from refParam) *AddCommand {
	return &AddCommand{source: from}
}
