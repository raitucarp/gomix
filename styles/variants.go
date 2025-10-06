package styles

type StyleVariant string

const (
	defaultVar       StyleVariant = ""
	hover            StyleVariant = "@media (hover: hover) { &:hover }"
	focus            StyleVariant = "&:focus"
	focusWithin      StyleVariant = "&:focus-within"
	focusVisible     StyleVariant = "&:focus-visible"
	active           StyleVariant = "&:active"
	visited          StyleVariant = "&:visited"
	target           StyleVariant = "&:target"
	star             StyleVariant = "&:is(& > *)"
	doubleStar       StyleVariant = "&:is(& *)"
	has              StyleVariant = "&:has(...)"
	group            StyleVariant = "&:is(:where(.group)... *)"
	peer             StyleVariant = "&:is(:where(.peer)... ~ *)"
	in               StyleVariant = ":where(...) &"
	not              StyleVariant = "&:not(...)"
	inert            StyleVariant = "&:is([inert], [inert] *)"
	first            StyleVariant = "&:first-child"
	last             StyleVariant = "&:last-child"
	only             StyleVariant = "&:only-child"
	odd              StyleVariant = "&:nth-child(odd)"
	even             StyleVariant = "&:nth-child(even)"
	firstOfType      StyleVariant = "&:first-of-type"
	lastOfType       StyleVariant = "&:last-of-type"
	onlyOfType       StyleVariant = "&:only-of-type"
	nth              StyleVariant = "&:nth-child(...)"
	nthLast          StyleVariant = "&:nth-last-child(...)"
	nthOfType        StyleVariant = "&:nth-of-type(...)"
	nthLastOfType    StyleVariant = "&:nth-last-of-type(...)"
	empty            StyleVariant = "&:empty"
	disabled         StyleVariant = "&:disabled"
	enabled          StyleVariant = "&:enabled"
	checked          StyleVariant = "&:checked"
	indeterminate    StyleVariant = "&:indeterminate"
	def              StyleVariant = "&:default"
	optional         StyleVariant = "&:optional"
	required         StyleVariant = "&:required"
	valid            StyleVariant = "&:valid"
	invalid          StyleVariant = "&:invalid"
	userValid        StyleVariant = "&:user-valid"
	userInvalid      StyleVariant = "&:user-invalid"
	inRange          StyleVariant = "&:in-range"
	outOfRange       StyleVariant = "&:out-of-range"
	placeholderShown StyleVariant = "&:placeholder-shown"
	detailsContent   StyleVariant = "&:details-content"
	autofill         StyleVariant = "&:autofill"
	readOnly         StyleVariant = "&:read-only"
	before           StyleVariant = "&::before"
	after            StyleVariant = "&::after"
	firstLetter      StyleVariant = "&::first-letter"
	firstLine        StyleVariant = "&::first-line"
	marker           StyleVariant = "&::marker, & *::marker"
	selection        StyleVariant = "&::selection"
	file             StyleVariant = "&::file-selector-button"
	backdrop         StyleVariant = "&::backdrop"
	placeholder      StyleVariant = "&::placeholder"
	sm               StyleVariant = "@media (width >= 40rem)"
	md               StyleVariant = "@media (width >= 48rem)"
	lg               StyleVariant = "@media (width >= 64rem)"
	xl               StyleVariant = "@media (width >= 80rem)"
	twoxl            StyleVariant = "@media (width >= 96rem)"
	min              StyleVariant = "@media (width >= ...)"
	maxSm            StyleVariant = "@media (width < 40rem)"
	maxMd            StyleVariant = "@media (width < 48rem)"
	maxLg            StyleVariant = "@media (width < 64rem)"
	maxXl            StyleVariant = "@media (width < 80rem)"
	maxTwoxl         StyleVariant = "@media (width < 96rem)"
	max              StyleVariant = "@media (width < ...)"
	at3xs            StyleVariant = "@container (width >= 16rem)"
	at2xs            StyleVariant = "@container (width >= 18rem)"
	atxs             StyleVariant = "@container (width >= 20rem)"
	atsm             StyleVariant = "@container (width >= 24rem)"
	atmd             StyleVariant = "@container (width >= 28rem)"
	atlg             StyleVariant = "@container (width >= 32rem)"
	atxl             StyleVariant = "@container (width >= 36rem)"
	at2xl            StyleVariant = "@container (width >= 42rem)"
	at3xl            StyleVariant = "@container (width >= 48rem)"
	at4xl            StyleVariant = "@container (width >= 56rem)"
	at5xl            StyleVariant = "@container (width >= 64rem)"
	at6xl            StyleVariant = "@container (width >= 72rem)"
	at7xl            StyleVariant = "@container (width >= 80rem)"
	atmin            StyleVariant = "@container (width >= ...)"
	atmax3xs         StyleVariant = "@container (width < 16rem)"
	atmax2xs         StyleVariant = "@container (width < 18rem)"
	atmaxxs          StyleVariant = "@container (width < 20rem)"
	atmaxsm          StyleVariant = "@container (width < 24rem)"
	atmaxmd          StyleVariant = "@container (width < 28rem)"
	atmaxlg          StyleVariant = "@container (width < 32rem)"
	atmaxxl          StyleVariant = "@container (width < 36rem)"
	atmax2xl         StyleVariant = "@container (width < 42rem)"
	atmax3xl         StyleVariant = "@container (width < 48rem)"
	atmax4xl         StyleVariant = "@container (width < 56rem)"
	atmax5xl         StyleVariant = "@container (width < 64rem)"
	atmax6xl         StyleVariant = "@container (width < 72rem)"
	atmax7xl         StyleVariant = "@container (width < 80rem)"
	atmax            StyleVariant = "@container (width < ...)"
	dark             StyleVariant = "@media (prefers-color-scheme: dark)"
	motionSafe       StyleVariant = "@media (prefers-reduced-motion: no-preference)"
	motionReduce     StyleVariant = "@media (prefers-reduced-motion: reduce)"
	contrastMore     StyleVariant = "@media (prefers-contrast: more)"
	contrastLess     StyleVariant = "@media (prefers-contrast: less)"
	forcedColors     StyleVariant = "@media (forced-colors: active)"
	invertedColors   StyleVariant = "@media (inverted-colors: inverted)"
	pointerFine      StyleVariant = "@media (pointer: fine)"
	pointerCoarse    StyleVariant = "@media (pointer: coarse)"
	pointerNone      StyleVariant = "@media (pointer: none)"
	anyPointerFine   StyleVariant = "@media (any-pointer: fine)"
	anyPointerCoarse StyleVariant = "@media (any-pointer: coarse)"
	anyPointerNone   StyleVariant = "@media (any-pointer: none)"
	portrait         StyleVariant = "@media (orientation: portrait)"
	landscape        StyleVariant = "@media (orientation: landscape)"
	noscript         StyleVariant = "@media (scripting: none)"
	print            StyleVariant = "@media print"
	supports         StyleVariant = "@supports (...)"
	ariaBusy         StyleVariant = "&[aria-busy=\"true\"]"
	ariaChecked      StyleVariant = "&[aria-checked=\"true\"]"
	ariaDisabled     StyleVariant = "&[aria-disabled=\"true\"]"
	ariaExpanded     StyleVariant = "&[aria-expanded=\"true\"]"
	ariaHidden       StyleVariant = "&[aria-hidden=\"true\"]"
	ariaPressed      StyleVariant = "&[aria-pressed=\"true\"]"
	ariaReadonly     StyleVariant = "&[aria-readonly=\"true\"]"
	ariaRequired     StyleVariant = "&[aria-required=\"true\"]"
	ariaSelected     StyleVariant = "&[aria-selected=\"true\"]"
	aria             StyleVariant = "&[aria-…]"
	data             StyleVariant = "&[data-…]"
	rtl              StyleVariant = "&:where(:dir(rtl), [dir=\"rtl\"], [dir=\"rtl\"] *)"
	ltr              StyleVariant = "&:where(:dir(ltr), [dir=\"ltr\"], [dir=\"ltr\"] *)"
	open             StyleVariant = "&:is([open], :popover-open, :open)"
	starting         StyleVariant = "@starting-style"
)

var styleVariantMap = map[StyleVariant]string{
	defaultVar:       "default",
	hover:            "hover",
	focus:            "focus",
	focusWithin:      "focus-within",
	focusVisible:     "focus-visible",
	active:           "active",
	visited:          "visited",
	target:           "target",
	star:             "star",
	doubleStar:       "double-star",
	has:              "has",
	group:            "group",
	peer:             "peer",
	in:               "in",
	not:              "not",
	inert:            "inert",
	first:            "first",
	last:             "last",
	only:             "only",
	odd:              "odd",
	even:             "even",
	firstOfType:      "first-of-type",
	lastOfType:       "last-of-type",
	onlyOfType:       "only-of-type",
	nth:              "nth",
	nthLast:          "nth-last",
	nthOfType:        "nth-of-type",
	nthLastOfType:    "nth-last-of-type",
	empty:            "empty",
	disabled:         "disabled",
	enabled:          "enabled",
	checked:          "checked",
	indeterminate:    "indeterminate",
	def:              "def",
	optional:         "optional",
	required:         "required",
	valid:            "valid",
	invalid:          "invalid",
	userValid:        "user-valid",
	userInvalid:      "user-invalid",
	inRange:          "in-range",
	outOfRange:       "out-of-range",
	placeholderShown: "placeholder-shown",
	detailsContent:   "details-content",
	autofill:         "autofill",
	readOnly:         "read-only",
	before:           "before",
	after:            "after",
	firstLetter:      "first-letter",
	firstLine:        "first-line",
	marker:           "marker",
	selection:        "selection",
	file:             "file",
	backdrop:         "backdrop",
	placeholder:      "placeholder",
	sm:               "sm",
	md:               "md",
	lg:               "lg",
	xl:               "xl",
	twoxl:            "twoxl",
	min:              "min",
	maxSm:            "max-sm",
	maxMd:            "max-md",
	maxLg:            "max-lg",
	maxXl:            "max-xl",
	maxTwoxl:         "max-twoxl",
	max:              "max",
	at3xs:            "at3xs",
	at2xs:            "at2xs",
	atxs:             "atxs",
	atsm:             "atsm",
	atmd:             "atmd",
	atlg:             "atlg",
	atxl:             "atxl",
	at2xl:            "at2xl",
	at3xl:            "at3xl",
	at4xl:            "at4xl",
	at5xl:            "at5xl",
	at6xl:            "at6xl",
	at7xl:            "at7xl",
	atmin:            "atmin",
	atmax3xs:         "atmax3xs",
	atmax2xs:         "atmax2xs",
	atmaxxs:          "atmaxxs",
	atmaxsm:          "atmaxsm",
	atmaxmd:          "atmaxmd",
	atmaxlg:          "atmaxlg",
	atmaxxl:          "atmaxxl",
	atmax2xl:         "atmax2xl",
	atmax3xl:         "atmax3xl",
	atmax4xl:         "atmax4xl",
	atmax5xl:         "atmax5xl",
	atmax6xl:         "atmax6xl",
	atmax7xl:         "atmax7xl",
	atmax:            "atmax",
	dark:             "dark",
	motionSafe:       "motion-safe",
	motionReduce:     "motion-reduce",
	contrastMore:     "contrast-more",
	contrastLess:     "contrast-less",
	forcedColors:     "forced-colors",
	invertedColors:   "inverted-colors",
	pointerFine:      "pointer-fine",
	pointerCoarse:    "pointer-coarse",
	pointerNone:      "pointer-none",
	anyPointerFine:   "any-pointer-fine",
	anyPointerCoarse: "any-pointer-coarse",
	anyPointerNone:   "any-pointer-none",
	portrait:         "portrait",
	landscape:        "landscape",
	noscript:         "noscript",
	print:            "print",
	supports:         "supports",
	ariaBusy:         "aria-busy",
	ariaChecked:      "aria-checked",
	ariaDisabled:     "aria-disabled",
	ariaExpanded:     "aria-expanded",
	ariaHidden:       "aria-hidden",
	ariaPressed:      "aria-pressed",
	ariaReadonly:     "aria-readonly",
	ariaRequired:     "aria-required",
	ariaSelected:     "aria-selected",
	aria:             "aria",
	data:             "data",
	rtl:              "rtl",
	ltr:              "ltr",
	open:             "open",
	starting:         "starting",
}

func VariantNameOfAttr(attr string) StyleVariant {
	for variant, attrString := range styleVariantMap {
		if attr == attrString {
			return variant
		}
	}
	return defaultVar
}

func (v StyleVariant) Name() string {
	return styleVariantMap[v]
}

type variantOf struct {
	variant StyleVariant
	props   Props
}

func (v *variantOf) StyleProp() {}
func (v *variantOf) GetProps() map[string]string {
	props := map[string]string{}

	for key, val := range v.props {
		props[string(key)] = val
	}

	return props
}
func (v *variantOf) Kind() kind {
	return kindVariant
}

func (v *variantOf) Variant() StyleVariant {
	return v.variant
}

func (v StyleVariant) makeStyleProp(props ...ApplyProp) ApplyProp {
	return func(s *style) styleProp {

		return &variantOf{
			variant: v,
			props:   s.props,
		}
	}
}

func Hover(props ...ApplyProp) ApplyProp { return hover.makeStyleProp(props...) }
