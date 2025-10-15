package ri

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed icons/*
var svgFiles embed.FS

const (
	svgPath     = "icons"
	name        = "Remix Icon"
)

type remixIcon struct {
	el *element.HtmlElement
}

func (i *remixIcon) IsIcon()                       {}
func (i *remixIcon) Name() string                  { return name }
func (i *remixIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &remixIcon{el}
	return icon
}
