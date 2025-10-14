package di

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath = "svg"
	name    = "Devicons"
)

type devIcons struct {
	el *element.HtmlElement
}

func (i *devIcons) IsIcon()                       {}
func (i *devIcons) Name() string                  { return name }
func (i *devIcons) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &devIcons{el}
	return icon
}
