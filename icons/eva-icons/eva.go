package ev

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed icons/*
var svgFiles embed.FS

const (
	svgPath     = "icons"
	name        = "Eva Icons"
	outlineIcon = "outline"
	fillIcon    = "fill"
)

type evaIcon struct {
	el *element.HtmlElement
}

func (i *evaIcon) IsIcon()                       {}
func (i *evaIcon) Name() string                  { return name }
func (i *evaIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgFiles, finalPath...)
	icon := &evaIcon{el}
	return icon
}
