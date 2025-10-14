package gi

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath = "svg"
	name    = "Game Icons"
)

type gameIcon struct {
	el *element.HtmlElement
}

func (i *gameIcon) IsIcon()                       {}
func (i *gameIcon) Name() string                  { return name }
func (i *gameIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &gameIcon{el}
	return icon
}
