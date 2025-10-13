package fa5

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svgs/*
var svgs embed.FS

const (
	svgPath     = "svgs"
	name        = "font-awesome-5"
	fontBrands  = "brands"
	fontSolid   = "solid"
	fontRegular = "regular"
)

type fontAwesomeIcon struct {
	el *element.HtmlElement
}

func (i *fontAwesomeIcon) IsIcon()                       {}
func (i *fontAwesomeIcon) Name() string                  { return name }
func (i *fontAwesomeIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgs, finalPath...)
	icon := &fontAwesomeIcon{el}
	return icon
}
