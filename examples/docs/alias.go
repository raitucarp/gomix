package main



import (
	"github.com/raitucarp/gomix"
	"github.com/raitucarp/gomix/components"
	"github.com/raitucarp/gomix/element"
	"github.com/raitucarp/gomix/addons/alpinejs"
	"github.com/raitucarp/gomix/components/buttons"
	"github.com/raitucarp/gomix/components/data_display"
	"github.com/raitucarp/gomix/components/disclosure"
	"github.com/raitucarp/gomix/icons/fontawesome-5"
	"github.com/raitucarp/gomix/components/feedback"
	"github.com/raitucarp/gomix/components/forms"
	"github.com/raitucarp/gomix/addons/google-fonts"
	"github.com/raitucarp/gomix/addons/htmx"
	layoutComponent "github.com/raitucarp/gomix/components/layout"
	"github.com/raitucarp/gomix/components/media_and_icons"
	"github.com/raitucarp/gomix/addons/modern-normalize"
	"github.com/raitucarp/gomix/components/overlays"
	"github.com/raitucarp/gomix/styles"
	"github.com/raitucarp/gomix/components/typography"
	"github.com/raitucarp/gomix/value"
)

type isComponent = components.IsComponent
type page_path = gomix.LocationPath
type fragment_path = gomix.FragmentPath
type css = gomix.CSS
// app scope alias
// fragment scope
// page scope
// addons scope alias
// element scope
// style scope
// components
// font awesome
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
var title_template = gomix.TitleTemplate
var fragment = gomix.FragmentComponent
var title_ = gomix.PageTitle
var component = gomix.PageComponent
var fragment_component = gomix.FragmentComponent
var _htmx = htmx.Addon
var _modern_normalize = modern_normalize.Addon
var usePlaywriteDeGrundFont = googlefonts.UseGoogleFontPlaywriteDeGrund
var _alpine_addon = alpinejs.Alpine
var alpine = alpinejs.Alpine
var div = element.Div
var aside = element.Aside
var Main = element.Main
var text = element.Text
var slot = components.Slot
var el = element.Element
var button = element.Button
var a = element.A
var span = element.Span
var Htmx = htmx.Htmx
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
var bg_slate_alpha = styles.BgSlateAlpha
var inline = styles.Inline
var width_by = styles.WBy
var px = value.Px[float64]
var flex_1 = styles.FlexBy(value.Number(1))
var fill_red = styles.FillRed
var overflowY_scroll = styles.OverflowYScroll
var HStack = layoutComponent.HStack
var VStack = layoutComponent.VStack
var Text = components.Text
var Icon = media_and_icons.Icon
var List = components.List
var Link = components.Link
var ListItem = components.ListItem
var ForEachComponentLink = components.ForEach[ComponentLink]
var ForComponentGroup = components.ForEach[ComponentGroup]
var Accordion = disclosure.Accordion
var AccordionItem = disclosure.AccordionItem
var fa_brand_android = fa5.FaBrandAndroid
var rem = value.Rem[float64]
var Alert = feedback.Alert
var AlertDialog = overlays.AlertDialog
var Avatar = media_and_icons.Avatar
var Badge = data_display.Badge
var Button = buttons.Button
var Center = layoutComponent.Center
var Checkbox = forms.Checkbox
var CircularProgress = feedback.CircularProgress
var CloseButton = buttons.CloseButton
var Code = typography.Code
var Divider = data_display.Divider
var Drawer = overlays.Drawer
var Grid = layoutComponent.Grid
var Heading = typography.Heading
var IconButton = buttons.IconButton
var Image = media_and_icons.Image
var Input = forms.Input
var Kbd = typography.Kbd
var Menu = overlays.Menu
var Modal = overlays.Modal
var NumberInput = forms.NumberInput
var PinInput = forms.PinInput
var Popover = overlays.Popover
var Progress = feedback.Progress
var Radio = forms.Radio
var SelectNative = forms.SelectNative
var SimpleGrid = layoutComponent.SimpleGrid
var Skeleton = feedback.Skeleton
var Slider = forms.Slider
var Spacer = layoutComponent.Spacer
var Spinner = feedback.Spinner
var Stack = layoutComponent.Stack
var Stat = data_display.Stat
var SwitchCmp = forms.SwitchCmp
var Tab = disclosure.Tab
var TabList = disclosure.TabList
var TabPanel = disclosure.TabPanel
var TabPanels = disclosure.TabPanels
var Table = data_display.Table
var Tabs = disclosure.Tabs
var Tag = data_display.Tag
var TextCmp = typography.TextCmp
var Textarea = forms.Textarea
var Toast = feedback.Toast
var Tooltip = overlays.Tooltip
var UseDisclosureComponents = disclosure.UseDisclosureComponents
var VisuallyHidden = disclosure.VisuallyHidden
var Wrap = layoutComponent.Wrap
var static = gomix.Static
var static_paths = gomix.StaticGenerationPath
var layout = gomix.PageLayout
var h1 = element.H1
var h2 = element.H2
