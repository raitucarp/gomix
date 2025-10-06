package styles

type StyleVariant string

const (
	defaultVar       StyleVariant = "& {%s}"
	hover            StyleVariant = "@media (hover: hover) { &:hover { %s } }"
	focus            StyleVariant = "&:focus { %s }"
	focusWithin      StyleVariant = "&:focus-within { %s }"
	focusVisible     StyleVariant = "&:focus-visible { %s }"
	active           StyleVariant = "&:active { %s }"
	visited          StyleVariant = "&:visited { %s }"
	target           StyleVariant = "&:target { %s }"
	star             StyleVariant = "&:is(& > *) { %s }"
	doubleStar       StyleVariant = "&:is(& *) { %s }"
	has              StyleVariant = "&:has(...) { %s }"
	group            StyleVariant = "&:is(:where(.group)... *) { %s }"
	peer             StyleVariant = "&:is(:where(.peer)... ~ *) { %s }"
	in               StyleVariant = ":where(...) & { %s }"
	not              StyleVariant = "&:not(...) { %s }"
	inert            StyleVariant = "&:is([inert], [inert] *) { %s }"
	first            StyleVariant = "&:first-child { %s }"
	last             StyleVariant = "&:last-child { %s }"
	only             StyleVariant = "&:only-child { %s }"
	odd              StyleVariant = "&:nth-child(odd) { %s }"
	even             StyleVariant = "&:nth-child(even) { %s }"
	firstOfType      StyleVariant = "&:first-of-type { %s }"
	lastOfType       StyleVariant = "&:last-of-type { %s }"
	onlyOfType       StyleVariant = "&:only-of-type { %s }"
	nth              StyleVariant = "&:nth-child(...) { %s }"
	nthLast          StyleVariant = "&:nth-last-child(...) { %s }"
	nthOfType        StyleVariant = "&:nth-of-type(...) { %s }"
	nthLastOfType    StyleVariant = "&:nth-last-of-type(...) { %s }"
	empty            StyleVariant = "&:empty { %s }"
	disabled         StyleVariant = "&:disabled { %s }"
	enabled          StyleVariant = "&:enabled { %s }"
	checked          StyleVariant = "&:checked { %s }"
	indeterminate    StyleVariant = "&:indeterminate { %s }"
	def              StyleVariant = "&:default { %s }"
	optional         StyleVariant = "&:optional { %s }"
	required         StyleVariant = "&:required { %s }"
	valid            StyleVariant = "&:valid { %s }"
	invalid          StyleVariant = "&:invalid { %s }"
	userValid        StyleVariant = "&:user-valid { %s }"
	userInvalid      StyleVariant = "&:user-invalid { %s }"
	inRange          StyleVariant = "&:in-range { %s }"
	outOfRange       StyleVariant = "&:out-of-range { %s }"
	placeholderShown StyleVariant = "&:placeholder-shown { %s }"
	detailsContent   StyleVariant = "&:details-content { %s }"
	autofill         StyleVariant = "&:autofill { %s }"
	readOnly         StyleVariant = "&:read-only { %s }"
	before           StyleVariant = "&::before { %s }"
	after            StyleVariant = "&::after { %s }"
	firstLetter      StyleVariant = "&::first-letter { %s }"
	firstLine        StyleVariant = "&::first-line { %s }"
	marker           StyleVariant = "&::marker, & *::marker { %s }"
	selection        StyleVariant = "&::selection { %s }"
	file             StyleVariant = "&::file-selector-button { %s }"
	backdrop         StyleVariant = "&::backdrop { %s }"
	placeholder      StyleVariant = "&::placeholder { %s }"
	sm               StyleVariant = "@media (width >= 40rem) { %s }"
	md               StyleVariant = "@media (width >= 48rem) { %s }"
	lg               StyleVariant = "@media (width >= 64rem) { %s }"
	xl               StyleVariant = "@media (width >= 80rem) { %s }"
	twoxl            StyleVariant = "@media (width >= 96rem) "
	min              StyleVariant = "@media (width >= ...) { %s }"
	maxSm            StyleVariant = "@media (width < 40rem) { %s }"
	maxMd            StyleVariant = "@media (width < 48rem) { %s }"
	maxLg            StyleVariant = "@media (width < 64rem) { %s }"
	maxXl            StyleVariant = "@media (width < 80rem) { %s }"
	maxTwoxl         StyleVariant = "@media (width < 96rem) { %s }"
	max              StyleVariant = "@media (width < ...) { %s }"
	at3xs            StyleVariant = "@container (width >= 16rem) { %s }"
	at2xs            StyleVariant = "@container (width >= 18rem) { %s }"
	atxs             StyleVariant = "@container (width >= 20rem) { %s }"
	atsm             StyleVariant = "@container (width >= 24rem) { %s }"
	atmd             StyleVariant = "@container (width >= 28rem) { %s }"
	atlg             StyleVariant = "@container (width >= 32rem) { %s }"
	atxl             StyleVariant = "@container (width >= 36rem) { %s }"
	at2xl            StyleVariant = "@container (width >= 42rem) { %s }"
	at3xl            StyleVariant = "@container (width >= 48rem) { %s }"
	at4xl            StyleVariant = "@container (width >= 56rem) { %s }"
	at5xl            StyleVariant = "@container (width >= 64rem) { %s }"
	at6xl            StyleVariant = "@container (width >= 72rem) { %s }"
	at7xl            StyleVariant = "@container (width >= 80rem) { %s }"
	atmin            StyleVariant = "@container (width >= ...) { %s }"
	atmax3xs         StyleVariant = "@container (width < 16rem) { %s }"
	atmax2xs         StyleVariant = "@container (width < 18rem) { %s }"
	atmaxxs          StyleVariant = "@container (width < 20rem) { %s }"
	atmaxsm          StyleVariant = "@container (width < 24rem) { %s }"
	atmaxmd          StyleVariant = "@container (width < 28rem) { %s }"
	atmaxlg          StyleVariant = "@container (width < 32rem) { %s }"
	atmaxxl          StyleVariant = "@container (width < 36rem) { %s }"
	atmax2xl         StyleVariant = "@container (width < 42rem) { %s }"
	atmax3xl         StyleVariant = "@container (width < 48rem) { %s }"
	atmax4xl         StyleVariant = "@container (width < 56rem) { %s }"
	atmax5xl         StyleVariant = "@container (width < 64rem) { %s }"
	atmax6xl         StyleVariant = "@container (width < 72rem) { %s }"
	atmax7xl         StyleVariant = "@container (width < 80rem) { %s }"
	atmax            StyleVariant = "@container (width < ...) { %s }"
	dark             StyleVariant = "@media (prefers-color-scheme: dark) { %s }"
	motionSafe       StyleVariant = "@media (prefers-reduced-motion: no-preference) { %s }"
	motionReduce     StyleVariant = "@media (prefers-reduced-motion: reduce) { %s }"
	contrastMore     StyleVariant = "@media (prefers-contrast: more) { %s }"
	contrastLess     StyleVariant = "@media (prefers-contrast: less) { %s }"
	forcedColors     StyleVariant = "@media (forced-colors: active) { %s }"
	invertedColors   StyleVariant = "@media (inverted-colors: inverted) { %s }"
	pointerFine      StyleVariant = "@media (pointer: fine) { %s }"
	pointerCoarse    StyleVariant = "@media (pointer: coarse) { %s }"
	pointerNone      StyleVariant = "@media (pointer: none) { %s }"
	anyPointerFine   StyleVariant = "@media (any-pointer: fine) { %s }"
	anyPointerCoarse StyleVariant = "@media (any-pointer: coarse) { %s }"
	anyPointerNone   StyleVariant = "@media (any-pointer: none) { %s }"
	portrait         StyleVariant = "@media (orientation: portrait) { %s }"
	landscape        StyleVariant = "@media (orientation: landscape) { %s }"
	noscript         StyleVariant = "@media (scripting: none) { %s }"
	print            StyleVariant = "@media print { %s }"
	supports         StyleVariant = "@supports (...) { %s }"
	ariaBusy         StyleVariant = "&[aria-busy=\"true\"] { %s }"
	ariaChecked      StyleVariant = "&[aria-checked=\"true\"] { %s }"
	ariaDisabled     StyleVariant = "&[aria-disabled=\"true\"] { %s }"
	ariaExpanded     StyleVariant = "&[aria-expanded=\"true\"] { %s }"
	ariaHidden       StyleVariant = "&[aria-hidden=\"true\"] { %s }"
	ariaPressed      StyleVariant = "&[aria-pressed=\"true\"] { %s }"
	ariaReadonly     StyleVariant = "&[aria-readonly=\"true\"] { %s }"
	ariaRequired     StyleVariant = "&[aria-required=\"true\"] { %s }"
	ariaSelected     StyleVariant = "&[aria-selected=\"true\"] { %s }"
	aria             StyleVariant = "&[aria-…] { %s }"
	data             StyleVariant = "&[data-…] { %s }"
	rtl              StyleVariant = "&:where(:dir(rtl), [dir=\"rtl\"], [dir=\"rtl\"] *) { %s }"
	ltr              StyleVariant = "&:where(:dir(ltr), [dir=\"ltr\"], [dir=\"ltr\"] *) { %s }"
	open             StyleVariant = "&:is([open], :popover-open, :open) { %s }"
	starting         StyleVariant = "@starting-style { %s }"
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

func (v *variantOf) addProp(prop Prop, value string) {
	v.props[prop] = value
}

func (v StyleVariant) makeStyleProp(props ...ApplyProp) ApplyProp {
	return func(s *style) styleProp {
		variantStyle := &variantOf{
			variant: v,
			props:   s.props,
		}

		for _, p := range props {
			props := p(s).GetProps()
			for key, value := range props {
				variantStyle.addProp(Prop(key), value)
			}

		}
		return variantStyle
	}
}

func Hover(props ...ApplyProp) ApplyProp { return hover.makeStyleProp(props...) }
