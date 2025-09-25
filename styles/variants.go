package styles

type variant string

const (
	defaultVar       variant = ""
	hover            variant = "@media (hover: hover) { &:hover }"
	focus            variant = "&:focus"
	focusWithin      variant = "&:focus-within"
	focusVisible     variant = "&:focus-visible"
	active           variant = "&:active"
	visited          variant = "&:visited"
	target           variant = "&:target"
	star             variant = "&:is(& > *)"
	doubleStar       variant = "&:is(& *)"
	has              variant = "&:has(...)"
	group            variant = "&:is(:where(.group)... *)"
	peer             variant = "&:is(:where(.peer)... ~ *)"
	in               variant = ":where(...) &"
	not              variant = "&:not(...)"
	inert            variant = "&:is([inert], [inert] *)"
	first            variant = "&:first-child"
	last             variant = "&:last-child"
	only             variant = "&:only-child"
	odd              variant = "&:nth-child(odd)"
	even             variant = "&:nth-child(even)"
	firstOfType      variant = "&:first-of-type"
	lastOfType       variant = "&:last-of-type"
	onlyOfType       variant = "&:only-of-type"
	nth              variant = "&:nth-child(...)"
	nthLast          variant = "&:nth-last-child(...)"
	nthOfType        variant = "&:nth-of-type(...)"
	nthLastOfType    variant = "&:nth-last-of-type(...)"
	empty            variant = "&:empty"
	disabled         variant = "&:disabled"
	enabled          variant = "&:enabled"
	checked          variant = "&:checked"
	indeterminate    variant = "&:indeterminate"
	def              variant = "&:default"
	optional         variant = "&:optional"
	required         variant = "&:required"
	valid            variant = "&:valid"
	invalid          variant = "&:invalid"
	userValid        variant = "&:user-valid"
	userInvalid      variant = "&:user-invalid"
	inRange          variant = "&:in-range"
	outOfRange       variant = "&:out-of-range"
	placeholderShown variant = "&:placeholder-shown"
	detailsContent   variant = "&:details-content"
	autofill         variant = "&:autofill"
	readOnly         variant = "&:read-only"
	before           variant = "&::before"
	after            variant = "&::after"
	firstLetter      variant = "&::first-letter"
	firstLine        variant = "&::first-line"
	marker           variant = "&::marker, & *::marker"
	selection        variant = "&::selection"
	file             variant = "&::file-selector-button"
	backdrop         variant = "&::backdrop"
	placeholder      variant = "&::placeholder"
	sm               variant = "@media (width >= 40rem)"
	md               variant = "@media (width >= 48rem)"
	lg               variant = "@media (width >= 64rem)"
	xl               variant = "@media (width >= 80rem)"
	twoxl            variant = "@media (width >= 96rem)"
	min              variant = "@media (width >= ...)"
	maxSm            variant = "@media (width < 40rem)"
	maxMd            variant = "@media (width < 48rem)"
	maxLg            variant = "@media (width < 64rem)"
	maxXl            variant = "@media (width < 80rem)"
	maxTwoxl         variant = "@media (width < 96rem)"
	max              variant = "@media (width < ...)"
	at3xs            variant = "@container (width >= 16rem)"
	at2xs            variant = "@container (width >= 18rem)"
	atxs             variant = "@container (width >= 20rem)"
	atsm             variant = "@container (width >= 24rem)"
	atmd             variant = "@container (width >= 28rem)"
	atlg             variant = "@container (width >= 32rem)"
	atxl             variant = "@container (width >= 36rem)"
	at2xl            variant = "@container (width >= 42rem)"
	at3xl            variant = "@container (width >= 48rem)"
	at4xl            variant = "@container (width >= 56rem)"
	at5xl            variant = "@container (width >= 64rem)"
	at6xl            variant = "@container (width >= 72rem)"
	at7xl            variant = "@container (width >= 80rem)"
	atmin            variant = "@container (width >= ...)"
	atmax3xs         variant = "@container (width < 16rem)"
	atmax2xs         variant = "@container (width < 18rem)"
	atmaxxs          variant = "@container (width < 20rem)"
	atmaxsm          variant = "@container (width < 24rem)"
	atmaxmd          variant = "@container (width < 28rem)"
	atmaxlg          variant = "@container (width < 32rem)"
	atmaxxl          variant = "@container (width < 36rem)"
	atmax2xl         variant = "@container (width < 42rem)"
	atmax3xl         variant = "@container (width < 48rem)"
	atmax4xl         variant = "@container (width < 56rem)"
	atmax5xl         variant = "@container (width < 64rem)"
	atmax6xl         variant = "@container (width < 72rem)"
	atmax7xl         variant = "@container (width < 80rem)"
	atmax            variant = "@container (width < ...)"
	dark             variant = "@media (prefers-color-scheme: dark)"
	motionSafe       variant = "@media (prefers-reduced-motion: no-preference)"
	motionReduce     variant = "@media (prefers-reduced-motion: reduce)"
	contrastMore     variant = "@media (prefers-contrast: more)"
	contrastLess     variant = "@media (prefers-contrast: less)"
	forcedColors     variant = "@media (forced-colors: active)"
	invertedColors   variant = "@media (inverted-colors: inverted)"
	pointerFine      variant = "@media (pointer: fine)"
	pointerCoarse    variant = "@media (pointer: coarse)"
	pointerNone      variant = "@media (pointer: none)"
	anyPointerFine   variant = "@media (any-pointer: fine)"
	anyPointerCoarse variant = "@media (any-pointer: coarse)"
	anyPointerNone   variant = "@media (any-pointer: none)"
	portrait         variant = "@media (orientation: portrait)"
	landscape        variant = "@media (orientation: landscape)"
	noscript         variant = "@media (scripting: none)"
	print            variant = "@media print"
	supports         variant = "@supports (...)"
	ariaBusy         variant = "&[aria-busy=\"true\"]"
	ariaChecked      variant = "&[aria-checked=\"true\"]"
	ariaDisabled     variant = "&[aria-disabled=\"true\"]"
	ariaExpanded     variant = "&[aria-expanded=\"true\"]"
	ariaHidden       variant = "&[aria-hidden=\"true\"]"
	ariaPressed      variant = "&[aria-pressed=\"true\"]"
	ariaReadonly     variant = "&[aria-readonly=\"true\"]"
	ariaRequired     variant = "&[aria-required=\"true\"]"
	ariaSelected     variant = "&[aria-selected=\"true\"]"
	aria             variant = "&[aria-…]"
	data             variant = "&[data-…]"
	rtl              variant = "&:where(:dir(rtl), [dir=\"rtl\"], [dir=\"rtl\"] *)"
	ltr              variant = "&:where(:dir(ltr), [dir=\"ltr\"], [dir=\"ltr\"] *)"
	open             variant = "&:is([open], :popover-open, :open)"
	starting         variant = "@starting-style"
)

type variantOf struct {
	variant variant
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
func (v *variantOf) Variant() variant {
	return v.variant
}

func (v variant) makeStyleProp(props ...ApplyProp) ApplyProp {
	return func(s *style) styleProp {

		return &variantOf{
			variant: v,
			props:   s.props,
		}
	}
}

func Hover(props ...ApplyProp) ApplyProp {
	return hover.makeStyleProp(props...)
}
