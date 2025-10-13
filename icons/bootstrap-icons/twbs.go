package bs

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed icons/*
var svgFiles embed.FS

const (
	svgPath = "icons"
	name    = "Bootstrap Icons"
)

type bootstrapIcon struct {
	el *element.HtmlElement
}

func (i *bootstrapIcon) IsIcon()                       {}
func (i *bootstrapIcon) Name() string                  { return name }
func (i *bootstrapIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &bootstrapIcon{el}
	return icon
}
