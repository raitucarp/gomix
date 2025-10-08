package main

import (
	"github.com/raitucarp/gomix"
	"github.com/raitucarp/gomix/addons/htmx"
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/styles"
)

type isComponent = components.IsComponent
type page_path = gomix.LocationPath
type fragment_path = gomix.FragmentPath

// app scope alias
var app = gomix.App
var name = gomix.Name
var addons = gomix.Addons
var web_addons = gomix.WebAddons
var port = gomix.Port
var features = gomix.Features
var logger = gomix.Logger
var layout = gomix.Layout
var webpage = gomix.Webpage
var scripts = gomix.Scripts
var at = gomix.PageAt
var fragment = gomix.FragmentAt

// addons scope alias
var _htmx = htmx.Addon

// element scope
var div = element.Div
var text = element.Text
var slot = components.Slot
var el = element.Element
var Htmx = htmx.Htmx
var button = element.Button
var a = element.A

// style scope

var hover = styles.Hover
var dark = styles.Dark
var md = styles.Md
var text_2xl = styles.Text2xl
var text_3xl = styles.Text3xl
var text_7xl = styles.Text7xl
var text_red = styles.TextRed
var text_orange = styles.TextOrange
var text_blue = styles.TextBlue
var text_green = styles.TextGreen
var cursor_pointer = styles.CursorPointer
