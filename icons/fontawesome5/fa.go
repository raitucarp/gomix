package fa5

import (
	"embed"

	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/icons"
)

//go:embed svgs/*
var svgs embed.FS

func readSvg(path string) (string, error) {
	f, err := svgs.ReadFile("svgs/" + path + ".svg")
	if err != nil {
		return "", err
	}

	return string(f), err
}

type fontAwesomeIcon struct {
	el *element.HtmlElement
}

func (f *fontAwesomeIcon) IsIcon() {}
func (f *fontAwesomeIcon) Element() *element.HtmlElement {
	return f.el
}

func createIconSvg(path string) icons.IsIcon {
	f, _ := readSvg(path)
	svgEl, _ := element.SvgString(f)
	return &fontAwesomeIcon{
		el: svgEl.Element(),
	}
}
