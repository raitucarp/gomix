package bi

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgFiles embed.FS

const (
	svgPath     = "svg"
	name        = "Box Icons"
	logosIcon   = "logos"
	solidIcon   = "solid"
	regularIcon = "regular"
)

type boxIcon struct {
	el *element.HtmlElement
}

func (i *boxIcon) IsIcon()                       {}
func (i *boxIcon) Name() string                  { return name }
func (i *boxIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &boxIcon{el}
	return icon
}
