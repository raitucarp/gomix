package styles

import (
	"strings"
)

type StyleVariant string

const (
	defaultVar       StyleVariant = "& { %s }"
	hover            StyleVariant = "@media (hover: hover) { &:hover { %s } }"
	focus            StyleVariant = "&:focus { %s }"
	focusWithin      StyleVariant = "&:focus-within { %s }"
	focusVisible     StyleVariant = "&:focus-visible { %s }"
	active           StyleVariant = "&:active { %s }"
	visited          StyleVariant = "&:visited { %s }"
	target           StyleVariant = "&:target { %s }"
	directChildren   StyleVariant = "&:is(& > *) { %s }"
	allDescendants   StyleVariant = "&:is(& *) { %s }"
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
	_2xl             StyleVariant = "@media (width >= 96rem) { %s }"
	min              StyleVariant = "@media (width >= ...) { %s }"
	maxSm            StyleVariant = "@media (width < 40rem) { %s }"
	maxMd            StyleVariant = "@media (width < 48rem) { %s }"
	maxLg            StyleVariant = "@media (width < 64rem) { %s }"
	maxXl            StyleVariant = "@media (width < 80rem) { %s }"
	max2xl           StyleVariant = "@media (width < 96rem) { %s }"
	max              StyleVariant = "@media (width < ...) { %s }"
	container3xs     StyleVariant = "@container (width >= 16rem) { %s }"
	container2xs     StyleVariant = "@container (width >= 18rem) { %s }"
	containerXs      StyleVariant = "@container (width >= 20rem) { %s }"
	containerSm      StyleVariant = "@container (width >= 24rem) { %s }"
	containerMd      StyleVariant = "@container (width >= 28rem) { %s }"
	containerLg      StyleVariant = "@container (width >= 32rem) { %s }"
	containerXl      StyleVariant = "@container (width >= 36rem) { %s }"
	container2xl     StyleVariant = "@container (width >= 42rem) { %s }"
	container3xl     StyleVariant = "@container (width >= 48rem) { %s }"
	container4xl     StyleVariant = "@container (width >= 56rem) { %s }"
	container5xl     StyleVariant = "@container (width >= 64rem) { %s }"
	container6xl     StyleVariant = "@container (width >= 72rem) { %s }"
	container7xl     StyleVariant = "@container (width >= 80rem) { %s }"
	containerMin     StyleVariant = "@container (width >= ...) { %s }"
	containerMax3xs  StyleVariant = "@container (width < 16rem) { %s }"
	containerMax2xs  StyleVariant = "@container (width < 18rem) { %s }"
	containerMaxXs   StyleVariant = "@container (width < 20rem) { %s }"
	containerMaxSm   StyleVariant = "@container (width < 24rem) { %s }"
	containerMaxMd   StyleVariant = "@container (width < 28rem) { %s }"
	containerMaxLg   StyleVariant = "@container (width < 32rem) { %s }"
	containerMaxXl   StyleVariant = "@container (width < 36rem) { %s }"
	containerMax2xl  StyleVariant = "@container (width < 42rem) { %s }"
	containerMax3xl  StyleVariant = "@container (width < 48rem) { %s }"
	containerMax4xl  StyleVariant = "@container (width < 56rem) { %s }"
	containerMax5xl  StyleVariant = "@container (width < 64rem) { %s }"
	containerMax6xl  StyleVariant = "@container (width < 72rem) { %s }"
	containerMax7xl  StyleVariant = "@container (width < 80rem) { %s }"
	containerMax     StyleVariant = "@container (width < ...) { %s }"
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
	directChildren:   "star",
	allDescendants:   "double-star",
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
	_2xl:             "2xl",
	min:              "min",
	maxSm:            "max-sm",
	maxMd:            "max-md",
	maxLg:            "max-lg",
	maxXl:            "max-xl",
	max2xl:           "max-2xl",
	max:              "max",
	container3xs:     "container-3xs",
	container2xs:     "container-2xs",
	containerXs:      "container-xs",
	containerSm:      "container-sm",
	containerMd:      "container-md",
	containerLg:      "container-lg",
	containerXl:      "container-xl",
	container2xl:     "container-2xl",
	container3xl:     "container-3xl",
	container4xl:     "container-4xl",
	container5xl:     "container-5xl",
	container6xl:     "container-6xl",
	container7xl:     "container-7xl",
	containerMin:     "container-min",
	containerMax3xs:  "container-max-3xs",
	containerMax2xs:  "container-max-2xs",
	containerMaxXs:   "container-max-xs",
	containerMaxSm:   "container-max-sm",
	containerMaxMd:   "container-max-md",
	containerMaxLg:   "container-max-lg",
	containerMaxXl:   "container-max-xl",
	containerMax2xl:  "container-max-2xl",
	containerMax3xl:  "container-max-3xl",
	containerMax4xl:  "container-max-4xl",
	containerMax5xl:  "container-max-5xl",
	containerMax6xl:  "container-max-6xl",
	containerMax7xl:  "container-max-7xl",
	containerMax:     "container-max",
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
	if attr == "" {
		return defaultVar
	}
	for variant, attrString := range styleVariantMap {
		if attr == attrString {
			return variant
		}
	}
	return StyleVariant(attr)
}

func (v StyleVariant) Name() string {
	if v, ok := styleVariantMap[v]; ok {
		return v
	}
	return string(v)
}

type variantOf struct {
	variant  StyleVariant
	props    Props
	params   []selectorParam
	children []*variantOf
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

func flattenVariants(variant *variantOf) []*variantOf {
	flatVariants := []*variantOf{}
	flatVariants = append(flatVariants, variant)

	for _, v := range variant.children {
		flatVariants = append(flatVariants, flattenVariants(v)...)
	}

	return flatVariants
}

func combinedVariants(variant *variantOf) map[StyleVariant]Props {
	combinedProps := make(map[StyleVariant]Props)
	variants := flattenVariants(variant)

	v := []string{}
	for _, vrnt := range variants {
		v = append(v, vrnt.variant.Name())
		combinedProps[StyleVariant(strings.Join(v, ":"))] = vrnt.props

	}
	return combinedProps
}

type selectorParam struct {
	value string
}

func (s *selectorParam) StyleProp() {}
func (s *selectorParam) GetProps() map[string]string {
	return map[string]string{"value": s.value}
}
func (s *selectorParam) Variant() StyleVariant {
	return defaultVar
}
func (s *selectorParam) Kind() kind {
	return kindSelectorParams
}

func SelectorParam(param string) ApplyProp {
	return func(s *style) styleProp {
		return &selectorParam{
			value: param,
		}
	}
}

func (v StyleVariant) makeStyleProp(props ...ApplyProp) ApplyProp {
	return func(s *style) styleProp {
		variantStyle := &variantOf{
			variant:  v,
			props:    make(Props),
			params:   []selectorParam{},
			children: []*variantOf{},
		}

		for _, p := range props {
			prop := p(s)
			if prop.Kind() == kindVariant {
				if vChild, ok := prop.(*variantOf); ok {
					variantStyle.children = append(variantStyle.children, vChild)
				}
				continue
			}

			if prop.Kind() == kindSelectorParams {
				param := selectorParam{
					value: prop.GetProps()["value"],
				}
				variantStyle.params = append(variantStyle.params, param)
				continue
			}

			varProps := p(s).GetProps()
			for key, value := range varProps {
				variantStyle.addProp(Prop(key), value)
			}

		}
		return variantStyle
	}
}

func Hover(props ...ApplyProp) ApplyProp            { return hover.makeStyleProp(props...) }
func Focus(props ...ApplyProp) ApplyProp            { return focus.makeStyleProp(props...) }
func FocusWithin(props ...ApplyProp) ApplyProp      { return focusWithin.makeStyleProp(props...) }
func FocusVisible(props ...ApplyProp) ApplyProp     { return focusVisible.makeStyleProp(props...) }
func Active(props ...ApplyProp) ApplyProp           { return active.makeStyleProp(props...) }
func Visited(props ...ApplyProp) ApplyProp          { return visited.makeStyleProp(props...) }
func Target(props ...ApplyProp) ApplyProp           { return target.makeStyleProp(props...) }
func DirectChildren(props ...ApplyProp) ApplyProp   { return directChildren.makeStyleProp(props...) }
func AllDescendants(props ...ApplyProp) ApplyProp   { return allDescendants.makeStyleProp(props...) }
func Has(props ...ApplyProp) ApplyProp              { return has.makeStyleProp(props...) }
func Group(props ...ApplyProp) ApplyProp            { return group.makeStyleProp(props...) }
func Peer(props ...ApplyProp) ApplyProp             { return peer.makeStyleProp(props...) }
func In(props ...ApplyProp) ApplyProp               { return in.makeStyleProp(props...) }
func Not(props ...ApplyProp) ApplyProp              { return not.makeStyleProp(props...) }
func Inert(props ...ApplyProp) ApplyProp            { return inert.makeStyleProp(props...) }
func First(props ...ApplyProp) ApplyProp            { return first.makeStyleProp(props...) }
func Last(props ...ApplyProp) ApplyProp             { return last.makeStyleProp(props...) }
func Only(props ...ApplyProp) ApplyProp             { return only.makeStyleProp(props...) }
func Odd(props ...ApplyProp) ApplyProp              { return odd.makeStyleProp(props...) }
func Even(props ...ApplyProp) ApplyProp             { return even.makeStyleProp(props...) }
func FirstOfType(props ...ApplyProp) ApplyProp      { return firstOfType.makeStyleProp(props...) }
func LastOfType(props ...ApplyProp) ApplyProp       { return lastOfType.makeStyleProp(props...) }
func OnlyOfType(props ...ApplyProp) ApplyProp       { return onlyOfType.makeStyleProp(props...) }
func Nth(props ...ApplyProp) ApplyProp              { return nth.makeStyleProp(props...) }
func NthLast(props ...ApplyProp) ApplyProp          { return nthLast.makeStyleProp(props...) }
func NthOfType(props ...ApplyProp) ApplyProp        { return nthOfType.makeStyleProp(props...) }
func NthLastOfType(props ...ApplyProp) ApplyProp    { return nthLastOfType.makeStyleProp(props...) }
func Empty(props ...ApplyProp) ApplyProp            { return empty.makeStyleProp(props...) }
func Disabled(props ...ApplyProp) ApplyProp         { return disabled.makeStyleProp(props...) }
func Enabled(props ...ApplyProp) ApplyProp          { return enabled.makeStyleProp(props...) }
func Checked(props ...ApplyProp) ApplyProp          { return checked.makeStyleProp(props...) }
func Indeterminate(props ...ApplyProp) ApplyProp    { return indeterminate.makeStyleProp(props...) }
func Def(props ...ApplyProp) ApplyProp              { return def.makeStyleProp(props...) }
func Optional(props ...ApplyProp) ApplyProp         { return optional.makeStyleProp(props...) }
func Required(props ...ApplyProp) ApplyProp         { return required.makeStyleProp(props...) }
func Valid(props ...ApplyProp) ApplyProp            { return valid.makeStyleProp(props...) }
func Invalid(props ...ApplyProp) ApplyProp          { return invalid.makeStyleProp(props...) }
func UserValid(props ...ApplyProp) ApplyProp        { return userValid.makeStyleProp(props...) }
func UserInvalid(props ...ApplyProp) ApplyProp      { return userInvalid.makeStyleProp(props...) }
func InRange(props ...ApplyProp) ApplyProp          { return inRange.makeStyleProp(props...) }
func OutOfRange(props ...ApplyProp) ApplyProp       { return outOfRange.makeStyleProp(props...) }
func PlaceholderShown(props ...ApplyProp) ApplyProp { return placeholderShown.makeStyleProp(props...) }
func DetailsContent(props ...ApplyProp) ApplyProp   { return detailsContent.makeStyleProp(props...) }
func Autofill(props ...ApplyProp) ApplyProp         { return autofill.makeStyleProp(props...) }
func ReadOnly(props ...ApplyProp) ApplyProp         { return readOnly.makeStyleProp(props...) }
func Before(props ...ApplyProp) ApplyProp           { return before.makeStyleProp(props...) }
func After(props ...ApplyProp) ApplyProp            { return after.makeStyleProp(props...) }
func FirstLetter(props ...ApplyProp) ApplyProp      { return firstLetter.makeStyleProp(props...) }
func FirstLine(props ...ApplyProp) ApplyProp        { return firstLine.makeStyleProp(props...) }
func Marker(props ...ApplyProp) ApplyProp           { return marker.makeStyleProp(props...) }
func Selection(props ...ApplyProp) ApplyProp        { return selection.makeStyleProp(props...) }
func File(props ...ApplyProp) ApplyProp             { return file.makeStyleProp(props...) }
func Backdrop(props ...ApplyProp) ApplyProp         { return backdrop.makeStyleProp(props...) }
func Placeholder(props ...ApplyProp) ApplyProp      { return placeholder.makeStyleProp(props...) }
func Sm(props ...ApplyProp) ApplyProp               { return sm.makeStyleProp(props...) }
func Md(props ...ApplyProp) ApplyProp               { return md.makeStyleProp(props...) }
func Lg(props ...ApplyProp) ApplyProp               { return lg.makeStyleProp(props...) }
func Xl(props ...ApplyProp) ApplyProp               { return xl.makeStyleProp(props...) }
func TwoXl(props ...ApplyProp) ApplyProp            { return _2xl.makeStyleProp(props...) }
func Min(props ...ApplyProp) ApplyProp              { return min.makeStyleProp(props...) }
func MaxSm(props ...ApplyProp) ApplyProp            { return maxSm.makeStyleProp(props...) }
func MaxMd(props ...ApplyProp) ApplyProp            { return maxMd.makeStyleProp(props...) }
func MaxLg(props ...ApplyProp) ApplyProp            { return maxLg.makeStyleProp(props...) }
func MaxXl(props ...ApplyProp) ApplyProp            { return maxXl.makeStyleProp(props...) }
func Max2xl(props ...ApplyProp) ApplyProp           { return max2xl.makeStyleProp(props...) }
func Max(props ...ApplyProp) ApplyProp              { return max.makeStyleProp(props...) }
func Container3xs(props ...ApplyProp) ApplyProp     { return container3xs.makeStyleProp(props...) }
func Container2xs(props ...ApplyProp) ApplyProp     { return container2xs.makeStyleProp(props...) }
func ContainerXs(props ...ApplyProp) ApplyProp      { return containerXs.makeStyleProp(props...) }
func ContainerSm(props ...ApplyProp) ApplyProp      { return containerSm.makeStyleProp(props...) }
func ContainerMd(props ...ApplyProp) ApplyProp      { return containerMd.makeStyleProp(props...) }
func ContainerLg(props ...ApplyProp) ApplyProp      { return containerLg.makeStyleProp(props...) }
func ContainerXl(props ...ApplyProp) ApplyProp      { return containerXl.makeStyleProp(props...) }
func Container2xl(props ...ApplyProp) ApplyProp     { return container2xl.makeStyleProp(props...) }
func Container3xl(props ...ApplyProp) ApplyProp     { return container3xl.makeStyleProp(props...) }
func Container4xl(props ...ApplyProp) ApplyProp     { return container4xl.makeStyleProp(props...) }
func Container5xl(props ...ApplyProp) ApplyProp     { return container5xl.makeStyleProp(props...) }
func Container6xl(props ...ApplyProp) ApplyProp     { return container6xl.makeStyleProp(props...) }
func Container7xl(props ...ApplyProp) ApplyProp     { return container7xl.makeStyleProp(props...) }
func ContainerMin(props ...ApplyProp) ApplyProp     { return containerMin.makeStyleProp(props...) }
func ContainerMax3xs(props ...ApplyProp) ApplyProp  { return containerMax3xs.makeStyleProp(props...) }
func ContainerMax2xs(props ...ApplyProp) ApplyProp  { return containerMax2xs.makeStyleProp(props...) }
func ContainerMaxXs(props ...ApplyProp) ApplyProp   { return containerMaxXs.makeStyleProp(props...) }
func ContainerMaxSm(props ...ApplyProp) ApplyProp   { return containerMaxSm.makeStyleProp(props...) }
func ContainerMaxMd(props ...ApplyProp) ApplyProp   { return containerMaxMd.makeStyleProp(props...) }
func ContainerMaxLg(props ...ApplyProp) ApplyProp   { return containerMaxLg.makeStyleProp(props...) }
func ContainerMaxXl(props ...ApplyProp) ApplyProp   { return containerMaxXl.makeStyleProp(props...) }
func ContainerMax2xl(props ...ApplyProp) ApplyProp  { return containerMax2xl.makeStyleProp(props...) }
func ContainerMax3xl(props ...ApplyProp) ApplyProp  { return containerMax3xl.makeStyleProp(props...) }
func ContainerMax4xl(props ...ApplyProp) ApplyProp  { return containerMax4xl.makeStyleProp(props...) }
func ContainerMax5xl(props ...ApplyProp) ApplyProp  { return containerMax5xl.makeStyleProp(props...) }
func ContainerMax6xl(props ...ApplyProp) ApplyProp  { return containerMax6xl.makeStyleProp(props...) }
func ContainerMax7xl(props ...ApplyProp) ApplyProp  { return containerMax7xl.makeStyleProp(props...) }
func ContainerMax(props ...ApplyProp) ApplyProp     { return containerMax.makeStyleProp(props...) }
func Dark(props ...ApplyProp) ApplyProp             { return dark.makeStyleProp(props...) }
func MotionSafe(props ...ApplyProp) ApplyProp       { return motionSafe.makeStyleProp(props...) }
func MotionReduce(props ...ApplyProp) ApplyProp     { return motionReduce.makeStyleProp(props...) }
func ContrastMore(props ...ApplyProp) ApplyProp     { return contrastMore.makeStyleProp(props...) }
func ContrastLess(props ...ApplyProp) ApplyProp     { return contrastLess.makeStyleProp(props...) }
func ForcedColors(props ...ApplyProp) ApplyProp     { return forcedColors.makeStyleProp(props...) }
func InvertedColors(props ...ApplyProp) ApplyProp   { return invertedColors.makeStyleProp(props...) }
func PointerFine(props ...ApplyProp) ApplyProp      { return pointerFine.makeStyleProp(props...) }
func PointerCoarse(props ...ApplyProp) ApplyProp    { return pointerCoarse.makeStyleProp(props...) }
func PointerNone(props ...ApplyProp) ApplyProp      { return pointerNone.makeStyleProp(props...) }
func AnyPointerFine(props ...ApplyProp) ApplyProp   { return anyPointerFine.makeStyleProp(props...) }
func AnyPointerCoarse(props ...ApplyProp) ApplyProp { return anyPointerCoarse.makeStyleProp(props...) }
func AnyPointerNone(props ...ApplyProp) ApplyProp   { return anyPointerNone.makeStyleProp(props...) }
func Portrait(props ...ApplyProp) ApplyProp         { return portrait.makeStyleProp(props...) }
func Landscape(props ...ApplyProp) ApplyProp        { return landscape.makeStyleProp(props...) }
func Noscript(props ...ApplyProp) ApplyProp         { return noscript.makeStyleProp(props...) }
func Print(props ...ApplyProp) ApplyProp            { return print.makeStyleProp(props...) }
func Supports(props ...ApplyProp) ApplyProp         { return supports.makeStyleProp(props...) }
func AriaBusy(props ...ApplyProp) ApplyProp         { return ariaBusy.makeStyleProp(props...) }
func AriaChecked(props ...ApplyProp) ApplyProp      { return ariaChecked.makeStyleProp(props...) }
func AriaDisabled(props ...ApplyProp) ApplyProp     { return ariaDisabled.makeStyleProp(props...) }
func AriaExpanded(props ...ApplyProp) ApplyProp     { return ariaExpanded.makeStyleProp(props...) }
func AriaHidden(props ...ApplyProp) ApplyProp       { return ariaHidden.makeStyleProp(props...) }
func AriaPressed(props ...ApplyProp) ApplyProp      { return ariaPressed.makeStyleProp(props...) }
func AriaReadonly(props ...ApplyProp) ApplyProp     { return ariaReadonly.makeStyleProp(props...) }
func AriaRequired(props ...ApplyProp) ApplyProp     { return ariaRequired.makeStyleProp(props...) }
func AriaSelected(props ...ApplyProp) ApplyProp     { return ariaSelected.makeStyleProp(props...) }
func Aria(props ...ApplyProp) ApplyProp             { return aria.makeStyleProp(props...) }
func Data(props ...ApplyProp) ApplyProp             { return data.makeStyleProp(props...) }
func Rtl(props ...ApplyProp) ApplyProp              { return rtl.makeStyleProp(props...) }
func Ltr(props ...ApplyProp) ApplyProp              { return ltr.makeStyleProp(props...) }
func Open(props ...ApplyProp) ApplyProp             { return open.makeStyleProp(props...) }
func Starting(props ...ApplyProp) ApplyProp         { return starting.makeStyleProp(props...) }
