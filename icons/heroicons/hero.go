package hi

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svg/*
var svgs embed.FS

const (
	svgPath     = "svg"
	name        = "Heroicons"
	solidIcon   = "solid"
	outlineIcon = "outline"
)

type heroIcon struct {
	el *element.HtmlElement
}

func (i *heroIcon) IsIcon()                       {}
func (i *heroIcon) Name() string                  { return name }
func (i *heroIcon) Element() *element.HtmlElement { return i.el }
func newIcon(paths ...string) icons.IsIcon {
	finalPath := []string{svgPath}
	finalPath = append(finalPath, paths...)
	el := icons.CreateSvgElementByFs(svgs, finalPath...)
	icon := &heroIcon{el}
	return icon
}
