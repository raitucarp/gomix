package pi

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed assets/*
var svgFiles embed.FS

const (
	svgPath     = "assets"
	name        = "Phosphor Icons"
	boldIcon    = "bold"
	duotoneIcon = "duotone"
	fillIcon    = "fill"
	lightIcon   = "light"
	regularIcon = "regular"
	thinIcon    = "thin"
)

type phosphorIcon struct {
	el *element.HtmlElement
}

func (i *phosphorIcon) IsIcon()                       {}
func (i *phosphorIcon) Name() string                  { return name }
func (i *phosphorIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &phosphorIcon{el}
	return icon
}
