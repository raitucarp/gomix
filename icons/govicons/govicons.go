package gvi

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath     = "svg"
	name        = "Govicons"
)

type govicons struct {
	el *element.HtmlElement
}

func (i *govicons) IsIcon()                       {}
func (i *govicons) Name() string                  { return name }
func (i *govicons) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &govicons{el}
	return icon
}
