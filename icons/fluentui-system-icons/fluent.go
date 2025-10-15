package fl

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed icons/*
var svgFiles embed.FS

const (
	svgPath = "icons"
	name    = "FluentUI System Icons"
	filledIcons = "filled"
	regularIcons = "regular"
)

type fluentUiIcons struct {
	el *element.HtmlElement
}

func (i *fluentUiIcons) IsIcon()                       {}
func (i *fluentUiIcons) Name() string                  { return name }
func (i *fluentUiIcons) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &fluentUiIcons{el}
	return icon
}
