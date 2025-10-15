package tb

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed icons/*
var svgFiles embed.FS

const (
	svgPath      = "icons"
	name         = "Tabler Icons"
	filledIcons  = "filled"
	outlineIcons = "outline"
)

type tablerIcon struct {
	el *element.HtmlElement
}

func (i *tablerIcon) IsIcon()                       {}
func (i *tablerIcon) Name() string                  { return name }
func (i *tablerIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &tablerIcon{el}
	return icon
}
