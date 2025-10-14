package im

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath = "svg"
	name    = "IcoMoon Free"
)

type icomoonFreeIcon struct {
	el *element.HtmlElement
}

func (i *icomoonFreeIcon) IsIcon()                       {}
func (i *icomoonFreeIcon) Name() string                  { return name }
func (i *icomoonFreeIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &icomoonFreeIcon{el}
	return icon
}
