package ico

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed icons/*
var svgFiles embed.FS

const (
	svgPath     = "icons"
	name        = "Iconoir"
	regularIcon = "regular"
	solidIcon   = "solid"
)

type iconoirIcon struct {
	el *element.HtmlElement
}

func (i *iconoirIcon) IsIcon()                       {}
func (i *iconoirIcon) Name() string                  { return name }
func (i *iconoirIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &iconoirIcon{el}
	return icon
}
