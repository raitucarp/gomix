package main

import (
	"github.com/raitucarp/gomix"
	"github.com/raitucarp/gomix/addons/alpinejs"
	"github.com/raitucarp/gomix/addons/htmx"
	modern_normalize "github.com/raitucarp/gomix/addons/modern-normalize"
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
	fa5 "github.com/raitucarp/gomix/icons/fontawesome-5"

	googlefonts "github.com/raitucarp/gomix/addons/google-fonts"
	"github.com/raitucarp/gomix/styles"
	"github.com/raitucarp/gomix/value"
)

type isComponent = components.IsComponent
type page_path = gomix.LocationPath
type fragment_path = gomix.FragmentPath
type css = gomix.CSS

// app scope alias
var app = gomix.App
var name = gomix.Name
var addons = gomix.Addons
var web_addons = gomix.WebAddons
var port = gomix.Port
var features = gomix.Features
var logger = gomix.Logger
var app_layout = gomix.Layout
var web = gomix.Web
var scripts = gomix.Scripts
var page_at = gomix.PageAt
var fragment_at = gomix.FragmentAt
var style_global = gomix.StyleGlobal
var style = gomix.Style
var not_found_page = gomix.NotFoundPage
var layout = gomix.PageLayout
var title_template = gomix.TitleTemplate

// fragment scope
var fragment = gomix.FragmentComponent

// page scope
var title_ = gomix.PageTitle
var component = gomix.PageComponent
var fragment_component = gomix.FragmentComponent

// addons scope alias
var _htmx = htmx.Addon
var _modern_normalize = modern_normalize.Addon
var usePlaywriteDeGrundFont = googlefonts.UseGoogleFontPlaywriteDeGrund
var _alpine_addon = alpinejs.Addon
var alpine = alpinejs.Alpine

// element scope
var div = element.Div
var text = element.Text
var slot = components.Slot
var el = element.Element
var button = element.Button
var a = element.A
var span = element.Span
var Htmx = htmx.Htmx

// style scope

var hover = styles.Hover
var dark = styles.Dark
var md = styles.Md
var text_2xl = styles.Text2xl
var text_white = styles.TextWhite
var text_3xl = styles.Text3xl
var text_7xl = styles.Text7xl
var text_red = styles.TextRed
var text_orange = styles.TextOrange
var text_blue = styles.TextBlue
var text_green = styles.TextGreen
var cursor_pointer = styles.CursorPointer
var bg_slate = styles.BgSlate
var inline = styles.Inline
var width_by = styles.WBy
var px = value.Px[float64]
var flex_1 = styles.FlexBy(value.Number(1))
var fill_red = styles.FillRed

// components
var Hstack = components.HStack
var Vstack = components.VStack
var Icon = components.Icon
var List = components.List
var ListItem = components.ListItem

// font awesome
var fa_brand_android = fa5.FaBrandAndroid
