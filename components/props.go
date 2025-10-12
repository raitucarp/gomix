package components

import (
	"github.com/raitucarp/gomix/styles"
	"github.com/raitucarp/gomix/value"
)

func (c *component) StyleHover(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Hover(props...))
	return c
}

func (c *component) StyleFocus(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Focus(props...))
	return c
}

func (c *component) StyleFocusWithin(props ...styles.ApplyProp) *component {
	c.el.Style(styles.FocusWithin(props...))
	return c
}

func (c *component) StyleFocusVisible(props ...styles.ApplyProp) *component {
	c.el.Style(styles.FocusVisible(props...))
	return c
}

func (c *component) StyleActive(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Active(props...))
	return c
}

func (c *component) StyleVisited(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Visited(props...))
	return c
}

func (c *component) StyleTarget(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Target(props...))
	return c
}

func (c *component) StyleDirectChildren(props ...styles.ApplyProp) *component {
	c.el.Style(styles.DirectChildren(props...))
	return c
}

func (c *component) StyleAllDescendants(props ...styles.ApplyProp) *component {
	c.el.Style(styles.AllDescendants(props...))
	return c
}

func (c *component) StyleHas(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Has(props...))
	return c
}

func (c *component) StyleGroup(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Group(props...))
	return c
}

func (c *component) StylePeer(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Peer(props...))
	return c
}

func (c *component) StyleIn(props ...styles.ApplyProp) *component {
	c.el.Style(styles.In(props...))
	return c
}

func (c *component) StyleNot(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Not(props...))
	return c
}

func (c *component) StyleInert(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Inert(props...))
	return c
}

func (c *component) StyleFirst(props ...styles.ApplyProp) *component {
	c.el.Style(styles.First(props...))
	return c
}

func (c *component) StyleLast(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Last(props...))
	return c
}

func (c *component) StyleOnly(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Only(props...))
	return c
}

func (c *component) StyleOdd(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Odd(props...))
	return c
}

func (c *component) StyleEven(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Even(props...))
	return c
}

func (c *component) StyleFirstOfType(props ...styles.ApplyProp) *component {
	c.el.Style(styles.FirstOfType(props...))
	return c
}

func (c *component) StyleLastOfType(props ...styles.ApplyProp) *component {
	c.el.Style(styles.LastOfType(props...))
	return c
}

func (c *component) StyleOnlyOfType(props ...styles.ApplyProp) *component {
	c.el.Style(styles.OnlyOfType(props...))
	return c
}

func (c *component) StyleNth(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Nth(props...))
	return c
}

func (c *component) StyleNthLast(props ...styles.ApplyProp) *component {
	c.el.Style(styles.NthLast(props...))
	return c
}

func (c *component) StyleNthOfType(props ...styles.ApplyProp) *component {
	c.el.Style(styles.NthOfType(props...))
	return c
}

func (c *component) StyleNthLastOfType(props ...styles.ApplyProp) *component {
	c.el.Style(styles.NthLastOfType(props...))
	return c
}

func (c *component) StyleEmpty(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Empty(props...))
	return c
}

func (c *component) StyleDisabled(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Disabled(props...))
	return c
}

func (c *component) StyleEnabled(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Enabled(props...))
	return c
}

func (c *component) StyleChecked(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Checked(props...))
	return c
}

func (c *component) StyleIndeterminate(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Indeterminate(props...))
	return c
}

func (c *component) StyleDef(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Def(props...))
	return c
}

func (c *component) StyleOptional(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Optional(props...))
	return c
}

func (c *component) StyleRequired(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Required(props...))
	return c
}

func (c *component) StyleValid(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Valid(props...))
	return c
}

func (c *component) StyleInvalid(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Invalid(props...))
	return c
}

func (c *component) StyleUserValid(props ...styles.ApplyProp) *component {
	c.el.Style(styles.UserValid(props...))
	return c
}

func (c *component) StyleUserInvalid(props ...styles.ApplyProp) *component {
	c.el.Style(styles.UserInvalid(props...))
	return c
}

func (c *component) StyleInRange(props ...styles.ApplyProp) *component {
	c.el.Style(styles.InRange(props...))
	return c
}

func (c *component) StyleOutOfRange(props ...styles.ApplyProp) *component {
	c.el.Style(styles.OutOfRange(props...))
	return c
}

func (c *component) StylePlaceholderShown(props ...styles.ApplyProp) *component {
	c.el.Style(styles.PlaceholderShown(props...))
	return c
}

func (c *component) StyleDetailsContent(props ...styles.ApplyProp) *component {
	c.el.Style(styles.DetailsContent(props...))
	return c
}

func (c *component) StyleAutofill(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Autofill(props...))
	return c
}

func (c *component) StyleReadOnly(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ReadOnly(props...))
	return c
}

func (c *component) StyleBefore(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Before(props...))
	return c
}

func (c *component) StyleAfter(props ...styles.ApplyProp) *component {
	c.el.Style(styles.After(props...))
	return c
}

func (c *component) StyleFirstLetter(props ...styles.ApplyProp) *component {
	c.el.Style(styles.FirstLetter(props...))
	return c
}

func (c *component) StyleFirstLine(props ...styles.ApplyProp) *component {
	c.el.Style(styles.FirstLine(props...))
	return c
}

func (c *component) StyleMarker(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Marker(props...))
	return c
}

func (c *component) StyleSelection(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Selection(props...))
	return c
}

func (c *component) StyleFile(props ...styles.ApplyProp) *component {
	c.el.Style(styles.File(props...))
	return c
}

func (c *component) StyleBackdrop(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Backdrop(props...))
	return c
}

func (c *component) StylePlaceholder(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Placeholder(props...))
	return c
}

func (c *component) StyleSm(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Sm(props...))
	return c
}

func (c *component) StyleMd(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Md(props...))
	return c
}

func (c *component) StyleLg(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Lg(props...))
	return c
}

func (c *component) StyleXl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Xl(props...))
	return c
}

func (c *component) StyleTwoXl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.TwoXl(props...))
	return c
}

func (c *component) StyleMin(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Min(props...))
	return c
}

func (c *component) StyleMaxSm(props ...styles.ApplyProp) *component {
	c.el.Style(styles.MaxSm(props...))
	return c
}

func (c *component) StyleMaxMd(props ...styles.ApplyProp) *component {
	c.el.Style(styles.MaxMd(props...))
	return c
}

func (c *component) StyleMaxLg(props ...styles.ApplyProp) *component {
	c.el.Style(styles.MaxLg(props...))
	return c
}

func (c *component) StyleMaxXl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.MaxXl(props...))
	return c
}

func (c *component) StyleMax2xl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Max2xl(props...))
	return c
}

func (c *component) StyleMax(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Max(props...))
	return c
}

func (c *component) StyleContainer3xs(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Container3xs(props...))
	return c
}

func (c *component) StyleContainer2xs(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Container2xs(props...))
	return c
}

func (c *component) StyleContainerXs(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerXs(props...))
	return c
}

func (c *component) StyleContainerSm(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerSm(props...))
	return c
}

func (c *component) StyleContainerMd(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMd(props...))
	return c
}

func (c *component) StyleContainerLg(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerLg(props...))
	return c
}

func (c *component) StyleContainerXl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerXl(props...))
	return c
}

func (c *component) StyleContainer2xl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Container2xl(props...))
	return c
}

func (c *component) StyleContainer3xl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Container3xl(props...))
	return c
}

func (c *component) StyleContainer4xl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Container4xl(props...))
	return c
}

func (c *component) StyleContainer5xl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Container5xl(props...))
	return c
}

func (c *component) StyleContainer6xl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Container6xl(props...))
	return c
}

func (c *component) StyleContainer7xl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Container7xl(props...))
	return c
}

func (c *component) StyleContainerMin(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMin(props...))
	return c
}

func (c *component) StyleContainerMax3xs(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMax3xs(props...))
	return c
}

func (c *component) StyleContainerMax2xs(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMax2xs(props...))
	return c
}

func (c *component) StyleContainerMaxXs(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMaxXs(props...))
	return c
}

func (c *component) StyleContainerMaxSm(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMaxSm(props...))
	return c
}

func (c *component) StyleContainerMaxMd(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMaxMd(props...))
	return c
}

func (c *component) StyleContainerMaxLg(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMaxLg(props...))
	return c
}

func (c *component) StyleContainerMaxXl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMaxXl(props...))
	return c
}

func (c *component) StyleContainerMax2xl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMax2xl(props...))
	return c
}

func (c *component) StyleContainerMax3xl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMax3xl(props...))
	return c
}

func (c *component) StyleContainerMax4xl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMax4xl(props...))
	return c
}

func (c *component) StyleContainerMax5xl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMax5xl(props...))
	return c
}

func (c *component) StyleContainerMax6xl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMax6xl(props...))
	return c
}

func (c *component) StyleContainerMax7xl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMax7xl(props...))
	return c
}

func (c *component) StyleContainerMax(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContainerMax(props...))
	return c
}

func (c *component) StyleDark(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Dark(props...))
	return c
}

func (c *component) StyleMotionSafe(props ...styles.ApplyProp) *component {
	c.el.Style(styles.MotionSafe(props...))
	return c
}

func (c *component) StyleMotionReduce(props ...styles.ApplyProp) *component {
	c.el.Style(styles.MotionReduce(props...))
	return c
}

func (c *component) StyleContrastMore(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContrastMore(props...))
	return c
}

func (c *component) StyleContrastLess(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ContrastLess(props...))
	return c
}

func (c *component) StyleForcedColors(props ...styles.ApplyProp) *component {
	c.el.Style(styles.ForcedColors(props...))
	return c
}

func (c *component) StyleInvertedColors(props ...styles.ApplyProp) *component {
	c.el.Style(styles.InvertedColors(props...))
	return c
}

func (c *component) StylePointerFine(props ...styles.ApplyProp) *component {
	c.el.Style(styles.PointerFine(props...))
	return c
}

func (c *component) StylePointerCoarse(props ...styles.ApplyProp) *component {
	c.el.Style(styles.PointerCoarse(props...))
	return c
}

func (c *component) StylePointerNone(props ...styles.ApplyProp) *component {
	c.el.Style(styles.PointerNone(props...))
	return c
}

func (c *component) StyleAnyPointerFine(props ...styles.ApplyProp) *component {
	c.el.Style(styles.AnyPointerFine(props...))
	return c
}

func (c *component) StyleAnyPointerCoarse(props ...styles.ApplyProp) *component {
	c.el.Style(styles.AnyPointerCoarse(props...))
	return c
}

func (c *component) StyleAnyPointerNone(props ...styles.ApplyProp) *component {
	c.el.Style(styles.AnyPointerNone(props...))
	return c
}

func (c *component) StylePortrait(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Portrait(props...))
	return c
}

func (c *component) StyleLandscape(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Landscape(props...))
	return c
}

func (c *component) StyleNoscript(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Noscript(props...))
	return c
}

func (c *component) StylePrint(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Print(props...))
	return c
}

func (c *component) StyleSupports(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Supports(props...))
	return c
}

func (c *component) StyleAriaBusy(props ...styles.ApplyProp) *component {
	c.el.Style(styles.AriaBusy(props...))
	return c
}

func (c *component) StyleAriaChecked(props ...styles.ApplyProp) *component {
	c.el.Style(styles.AriaChecked(props...))
	return c
}

func (c *component) StyleAriaDisabled(props ...styles.ApplyProp) *component {
	c.el.Style(styles.AriaDisabled(props...))
	return c
}

func (c *component) StyleAriaExpanded(props ...styles.ApplyProp) *component {
	c.el.Style(styles.AriaExpanded(props...))
	return c
}

func (c *component) StyleAriaHidden(props ...styles.ApplyProp) *component {
	c.el.Style(styles.AriaHidden(props...))
	return c
}

func (c *component) StyleAriaPressed(props ...styles.ApplyProp) *component {
	c.el.Style(styles.AriaPressed(props...))
	return c
}

func (c *component) StyleAriaReadonly(props ...styles.ApplyProp) *component {
	c.el.Style(styles.AriaReadonly(props...))
	return c
}

func (c *component) StyleAriaRequired(props ...styles.ApplyProp) *component {
	c.el.Style(styles.AriaRequired(props...))
	return c
}

func (c *component) StyleAriaSelected(props ...styles.ApplyProp) *component {
	c.el.Style(styles.AriaSelected(props...))
	return c
}

func (c *component) StyleAria(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Aria(props...))
	return c
}

func (c *component) StyleData(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Data(props...))
	return c
}

func (c *component) StyleRtl(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Rtl(props...))
	return c
}

func (c *component) StyleLtr(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Ltr(props...))
	return c
}

func (c *component) StyleOpen(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Open(props...))
	return c
}

func (c *component) StyleStarting(props ...styles.ApplyProp) *component {
	c.el.Style(styles.Starting(props...))
	return c
}

func (c *component) Absolute() *component {
	c.el.Style(styles.Absolute())
	return c
}

func (c *component) Accent(val value.Value) *component {
	c.el.Style(styles.Accent(val))
	return c
}

func (c *component) AccentAmber(scale int) *component {
	c.el.Style(styles.AccentAmber(scale))
	return c
}

func (c *component) AccentAmberAlpha(scale int) *component {
	c.el.Style(styles.AccentAmberAlpha(scale))
	return c
}

func (c *component) AccentAmberDark(scale int) *component {
	c.el.Style(styles.AccentAmberDark(scale))
	return c
}

func (c *component) AccentAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentAmberDarkAlpha(scale))
	return c
}

func (c *component) AccentBlack() *component {
	c.el.Style(styles.AccentBlack())
	return c
}

func (c *component) AccentBlackAlpha(scale int) *component {
	c.el.Style(styles.AccentBlackAlpha(scale))
	return c
}

func (c *component) AccentBlue(scale int) *component {
	c.el.Style(styles.AccentBlue(scale))
	return c
}

func (c *component) AccentBlueAlpha(scale int) *component {
	c.el.Style(styles.AccentBlueAlpha(scale))
	return c
}

func (c *component) AccentBlueDark(scale int) *component {
	c.el.Style(styles.AccentBlueDark(scale))
	return c
}

func (c *component) AccentBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentBlueDarkAlpha(scale))
	return c
}

func (c *component) AccentBronze(scale int) *component {
	c.el.Style(styles.AccentBronze(scale))
	return c
}

func (c *component) AccentBronzeAlpha(scale int) *component {
	c.el.Style(styles.AccentBronzeAlpha(scale))
	return c
}

func (c *component) AccentBronzeDark(scale int) *component {
	c.el.Style(styles.AccentBronzeDark(scale))
	return c
}

func (c *component) AccentBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentBronzeDarkAlpha(scale))
	return c
}

func (c *component) AccentBrown(scale int) *component {
	c.el.Style(styles.AccentBrown(scale))
	return c
}

func (c *component) AccentBrownAlpha(scale int) *component {
	c.el.Style(styles.AccentBrownAlpha(scale))
	return c
}

func (c *component) AccentBrownDark(scale int) *component {
	c.el.Style(styles.AccentBrownDark(scale))
	return c
}

func (c *component) AccentBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentBrownDarkAlpha(scale))
	return c
}

func (c *component) AccentCrimson(scale int) *component {
	c.el.Style(styles.AccentCrimson(scale))
	return c
}

func (c *component) AccentCrimsonAlpha(scale int) *component {
	c.el.Style(styles.AccentCrimsonAlpha(scale))
	return c
}

func (c *component) AccentCrimsonDark(scale int) *component {
	c.el.Style(styles.AccentCrimsonDark(scale))
	return c
}

func (c *component) AccentCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentCrimsonDarkAlpha(scale))
	return c
}

func (c *component) AccentCurrent() *component {
	c.el.Style(styles.AccentCurrent())
	return c
}

func (c *component) AccentCyan(scale int) *component {
	c.el.Style(styles.AccentCyan(scale))
	return c
}

func (c *component) AccentCyanAlpha(scale int) *component {
	c.el.Style(styles.AccentCyanAlpha(scale))
	return c
}

func (c *component) AccentCyanDark(scale int) *component {
	c.el.Style(styles.AccentCyanDark(scale))
	return c
}

func (c *component) AccentCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentCyanDarkAlpha(scale))
	return c
}

func (c *component) AccentGold(scale int) *component {
	c.el.Style(styles.AccentGold(scale))
	return c
}

func (c *component) AccentGoldAlpha(scale int) *component {
	c.el.Style(styles.AccentGoldAlpha(scale))
	return c
}

func (c *component) AccentGoldDark(scale int) *component {
	c.el.Style(styles.AccentGoldDark(scale))
	return c
}

func (c *component) AccentGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentGoldDarkAlpha(scale))
	return c
}

func (c *component) AccentGrass(scale int) *component {
	c.el.Style(styles.AccentGrass(scale))
	return c
}

func (c *component) AccentGrassAlpha(scale int) *component {
	c.el.Style(styles.AccentGrassAlpha(scale))
	return c
}

func (c *component) AccentGrassDark(scale int) *component {
	c.el.Style(styles.AccentGrassDark(scale))
	return c
}

func (c *component) AccentGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentGrassDarkAlpha(scale))
	return c
}

func (c *component) AccentGray(scale int) *component {
	c.el.Style(styles.AccentGray(scale))
	return c
}

func (c *component) AccentGrayAlpha(scale int) *component {
	c.el.Style(styles.AccentGrayAlpha(scale))
	return c
}

func (c *component) AccentGrayDark(scale int) *component {
	c.el.Style(styles.AccentGrayDark(scale))
	return c
}

func (c *component) AccentGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentGrayDarkAlpha(scale))
	return c
}

func (c *component) AccentGreen(scale int) *component {
	c.el.Style(styles.AccentGreen(scale))
	return c
}

func (c *component) AccentGreenAlpha(scale int) *component {
	c.el.Style(styles.AccentGreenAlpha(scale))
	return c
}

func (c *component) AccentGreenDark(scale int) *component {
	c.el.Style(styles.AccentGreenDark(scale))
	return c
}

func (c *component) AccentGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentGreenDarkAlpha(scale))
	return c
}

func (c *component) AccentIndigo(scale int) *component {
	c.el.Style(styles.AccentIndigo(scale))
	return c
}

func (c *component) AccentIndigoAlpha(scale int) *component {
	c.el.Style(styles.AccentIndigoAlpha(scale))
	return c
}

func (c *component) AccentIndigoDark(scale int) *component {
	c.el.Style(styles.AccentIndigoDark(scale))
	return c
}

func (c *component) AccentIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentIndigoDarkAlpha(scale))
	return c
}

func (c *component) AccentInherit() *component {
	c.el.Style(styles.AccentInherit())
	return c
}

func (c *component) AccentIris(scale int) *component {
	c.el.Style(styles.AccentIris(scale))
	return c
}

func (c *component) AccentIrisAlpha(scale int) *component {
	c.el.Style(styles.AccentIrisAlpha(scale))
	return c
}

func (c *component) AccentIrisDark(scale int) *component {
	c.el.Style(styles.AccentIrisDark(scale))
	return c
}

func (c *component) AccentIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentIrisDarkAlpha(scale))
	return c
}

func (c *component) AccentJade(scale int) *component {
	c.el.Style(styles.AccentJade(scale))
	return c
}

func (c *component) AccentJadeAlpha(scale int) *component {
	c.el.Style(styles.AccentJadeAlpha(scale))
	return c
}

func (c *component) AccentJadeDark(scale int) *component {
	c.el.Style(styles.AccentJadeDark(scale))
	return c
}

func (c *component) AccentJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentJadeDarkAlpha(scale))
	return c
}

func (c *component) AccentLime(scale int) *component {
	c.el.Style(styles.AccentLime(scale))
	return c
}

func (c *component) AccentLimeAlpha(scale int) *component {
	c.el.Style(styles.AccentLimeAlpha(scale))
	return c
}

func (c *component) AccentLimeDark(scale int) *component {
	c.el.Style(styles.AccentLimeDark(scale))
	return c
}

func (c *component) AccentLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentLimeDarkAlpha(scale))
	return c
}

func (c *component) AccentMauve(scale int) *component {
	c.el.Style(styles.AccentMauve(scale))
	return c
}

func (c *component) AccentMauveAlpha(scale int) *component {
	c.el.Style(styles.AccentMauveAlpha(scale))
	return c
}

func (c *component) AccentMauveDark(scale int) *component {
	c.el.Style(styles.AccentMauveDark(scale))
	return c
}

func (c *component) AccentMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentMauveDarkAlpha(scale))
	return c
}

func (c *component) AccentMint(scale int) *component {
	c.el.Style(styles.AccentMint(scale))
	return c
}

func (c *component) AccentMintAlpha(scale int) *component {
	c.el.Style(styles.AccentMintAlpha(scale))
	return c
}

func (c *component) AccentMintDark(scale int) *component {
	c.el.Style(styles.AccentMintDark(scale))
	return c
}

func (c *component) AccentMintDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentMintDarkAlpha(scale))
	return c
}

func (c *component) AccentOlive(scale int) *component {
	c.el.Style(styles.AccentOlive(scale))
	return c
}

func (c *component) AccentOliveAlpha(scale int) *component {
	c.el.Style(styles.AccentOliveAlpha(scale))
	return c
}

func (c *component) AccentOliveDark(scale int) *component {
	c.el.Style(styles.AccentOliveDark(scale))
	return c
}

func (c *component) AccentOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentOliveDarkAlpha(scale))
	return c
}

func (c *component) AccentOrange(scale int) *component {
	c.el.Style(styles.AccentOrange(scale))
	return c
}

func (c *component) AccentOrangeAlpha(scale int) *component {
	c.el.Style(styles.AccentOrangeAlpha(scale))
	return c
}

func (c *component) AccentOrangeDark(scale int) *component {
	c.el.Style(styles.AccentOrangeDark(scale))
	return c
}

func (c *component) AccentOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentOrangeDarkAlpha(scale))
	return c
}

func (c *component) AccentPink(scale int) *component {
	c.el.Style(styles.AccentPink(scale))
	return c
}

func (c *component) AccentPinkAlpha(scale int) *component {
	c.el.Style(styles.AccentPinkAlpha(scale))
	return c
}

func (c *component) AccentPinkDark(scale int) *component {
	c.el.Style(styles.AccentPinkDark(scale))
	return c
}

func (c *component) AccentPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentPinkDarkAlpha(scale))
	return c
}

func (c *component) AccentPlum(scale int) *component {
	c.el.Style(styles.AccentPlum(scale))
	return c
}

func (c *component) AccentPlumAlpha(scale int) *component {
	c.el.Style(styles.AccentPlumAlpha(scale))
	return c
}

func (c *component) AccentPlumDark(scale int) *component {
	c.el.Style(styles.AccentPlumDark(scale))
	return c
}

func (c *component) AccentPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentPlumDarkAlpha(scale))
	return c
}

func (c *component) AccentPurple(scale int) *component {
	c.el.Style(styles.AccentPurple(scale))
	return c
}

func (c *component) AccentPurpleAlpha(scale int) *component {
	c.el.Style(styles.AccentPurpleAlpha(scale))
	return c
}

func (c *component) AccentPurpleDark(scale int) *component {
	c.el.Style(styles.AccentPurpleDark(scale))
	return c
}

func (c *component) AccentPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentPurpleDarkAlpha(scale))
	return c
}

func (c *component) AccentRed(scale int) *component {
	c.el.Style(styles.AccentRed(scale))
	return c
}

func (c *component) AccentRedAlpha(scale int) *component {
	c.el.Style(styles.AccentRedAlpha(scale))
	return c
}

func (c *component) AccentRedDark(scale int) *component {
	c.el.Style(styles.AccentRedDark(scale))
	return c
}

func (c *component) AccentRedDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentRedDarkAlpha(scale))
	return c
}

func (c *component) AccentRuby(scale int) *component {
	c.el.Style(styles.AccentRuby(scale))
	return c
}

func (c *component) AccentRubyAlpha(scale int) *component {
	c.el.Style(styles.AccentRubyAlpha(scale))
	return c
}

func (c *component) AccentRubyDark(scale int) *component {
	c.el.Style(styles.AccentRubyDark(scale))
	return c
}

func (c *component) AccentRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentRubyDarkAlpha(scale))
	return c
}

func (c *component) AccentSage(scale int) *component {
	c.el.Style(styles.AccentSage(scale))
	return c
}

func (c *component) AccentSageAlpha(scale int) *component {
	c.el.Style(styles.AccentSageAlpha(scale))
	return c
}

func (c *component) AccentSageDark(scale int) *component {
	c.el.Style(styles.AccentSageDark(scale))
	return c
}

func (c *component) AccentSageDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentSageDarkAlpha(scale))
	return c
}

func (c *component) AccentSand(scale int) *component {
	c.el.Style(styles.AccentSand(scale))
	return c
}

func (c *component) AccentSandAlpha(scale int) *component {
	c.el.Style(styles.AccentSandAlpha(scale))
	return c
}

func (c *component) AccentSandDark(scale int) *component {
	c.el.Style(styles.AccentSandDark(scale))
	return c
}

func (c *component) AccentSandDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentSandDarkAlpha(scale))
	return c
}

func (c *component) AccentSky(scale int) *component {
	c.el.Style(styles.AccentSky(scale))
	return c
}

func (c *component) AccentSkyAlpha(scale int) *component {
	c.el.Style(styles.AccentSkyAlpha(scale))
	return c
}

func (c *component) AccentSkyDark(scale int) *component {
	c.el.Style(styles.AccentSkyDark(scale))
	return c
}

func (c *component) AccentSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentSkyDarkAlpha(scale))
	return c
}

func (c *component) AccentSlate(scale int) *component {
	c.el.Style(styles.AccentSlate(scale))
	return c
}

func (c *component) AccentSlateAlpha(scale int) *component {
	c.el.Style(styles.AccentSlateAlpha(scale))
	return c
}

func (c *component) AccentSlateDark(scale int) *component {
	c.el.Style(styles.AccentSlateDark(scale))
	return c
}

func (c *component) AccentSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentSlateDarkAlpha(scale))
	return c
}

func (c *component) AccentTeal(scale int) *component {
	c.el.Style(styles.AccentTeal(scale))
	return c
}

func (c *component) AccentTealAlpha(scale int) *component {
	c.el.Style(styles.AccentTealAlpha(scale))
	return c
}

func (c *component) AccentTealDark(scale int) *component {
	c.el.Style(styles.AccentTealDark(scale))
	return c
}

func (c *component) AccentTealDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentTealDarkAlpha(scale))
	return c
}

func (c *component) AccentTomato(scale int) *component {
	c.el.Style(styles.AccentTomato(scale))
	return c
}

func (c *component) AccentTomatoAlpha(scale int) *component {
	c.el.Style(styles.AccentTomatoAlpha(scale))
	return c
}

func (c *component) AccentTomatoDark(scale int) *component {
	c.el.Style(styles.AccentTomatoDark(scale))
	return c
}

func (c *component) AccentTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentTomatoDarkAlpha(scale))
	return c
}

func (c *component) AccentTransparent() *component {
	c.el.Style(styles.AccentTransparent())
	return c
}

func (c *component) AccentViolet(scale int) *component {
	c.el.Style(styles.AccentViolet(scale))
	return c
}

func (c *component) AccentVioletAlpha(scale int) *component {
	c.el.Style(styles.AccentVioletAlpha(scale))
	return c
}

func (c *component) AccentVioletDark(scale int) *component {
	c.el.Style(styles.AccentVioletDark(scale))
	return c
}

func (c *component) AccentVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentVioletDarkAlpha(scale))
	return c
}

func (c *component) AccentWhite() *component {
	c.el.Style(styles.AccentWhite())
	return c
}

func (c *component) AccentWhiteAlpha(scale int) *component {
	c.el.Style(styles.AccentWhiteAlpha(scale))
	return c
}

func (c *component) AccentYellow(scale int) *component {
	c.el.Style(styles.AccentYellow(scale))
	return c
}

func (c *component) AccentYellowAlpha(scale int) *component {
	c.el.Style(styles.AccentYellowAlpha(scale))
	return c
}

func (c *component) AccentYellowDark(scale int) *component {
	c.el.Style(styles.AccentYellowDark(scale))
	return c
}

func (c *component) AccentYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.AccentYellowDarkAlpha(scale))
	return c
}

func (c *component) AlignBaseline() *component {
	c.el.Style(styles.AlignBaseline())
	return c
}

func (c *component) AlignBottom() *component {
	c.el.Style(styles.AlignBottom())
	return c
}

func (c *component) AlignBy(val value.Value) *component {
	c.el.Style(styles.AlignBy(val))
	return c
}

func (c *component) AlignMiddle() *component {
	c.el.Style(styles.AlignMiddle())
	return c
}

func (c *component) AlignSub() *component {
	c.el.Style(styles.AlignSub())
	return c
}

func (c *component) AlignSuper() *component {
	c.el.Style(styles.AlignSuper())
	return c
}

func (c *component) AlignTextBottom() *component {
	c.el.Style(styles.AlignTextBottom())
	return c
}

func (c *component) AlignTextTop() *component {
	c.el.Style(styles.AlignTextTop())
	return c
}

func (c *component) AlignTop() *component {
	c.el.Style(styles.AlignTop())
	return c
}

func (c *component) Animate(val value.Value) *component {
	c.el.Style(styles.Animate(val))
	return c
}

func (c *component) AnimateBounce() *component {
	c.el.Style(styles.AnimateBounce())
	return c
}

func (c *component) AnimateNone() *component {
	c.el.Style(styles.AnimateNone())
	return c
}

func (c *component) AnimatePing() *component {
	c.el.Style(styles.AnimatePing())
	return c
}

func (c *component) AnimatePulse() *component {
	c.el.Style(styles.AnimatePulse())
	return c
}

func (c *component) AnimateSpin() *component {
	c.el.Style(styles.AnimateSpin())
	return c
}

func (c *component) Antialiazed() *component {
	c.el.Style(styles.Antialiazed())
	return c
}

func (c *component) AppearanceAuto() *component {
	c.el.Style(styles.AppearanceAuto())
	return c
}

func (c *component) AppearanceNone() *component {
	c.el.Style(styles.AppearanceNone())
	return c
}

func (c *component) Aspect(ratio value.Value) *component {
	c.el.Style(styles.Aspect(ratio))
	return c
}

func (c *component) AspectAuto() *component {
	c.el.Style(styles.AspectAuto())
	return c
}

func (c *component) AspectSquare() *component {
	c.el.Style(styles.AspectSquare())
	return c
}

func (c *component) AspectVideo() *component {
	c.el.Style(styles.AspectVideo())
	return c
}

func (c *component) AutoColsAuto() *component {
	c.el.Style(styles.AutoColsAuto())
	return c
}

func (c *component) AutoColsBy(val value.Value) *component {
	c.el.Style(styles.AutoColsBy(val))
	return c
}

func (c *component) AutoColsFr() *component {
	c.el.Style(styles.AutoColsFr())
	return c
}

func (c *component) AutoColsMax() *component {
	c.el.Style(styles.AutoColsMax())
	return c
}

func (c *component) AutoColsMin() *component {
	c.el.Style(styles.AutoColsMin())
	return c
}

func (c *component) AutoRowsAuto() *component {
	c.el.Style(styles.AutoRowsAuto())
	return c
}

func (c *component) AutoRowsBy(val value.Value) *component {
	c.el.Style(styles.AutoRowsBy(val))
	return c
}

func (c *component) AutoRowsFr() *component {
	c.el.Style(styles.AutoRowsFr())
	return c
}

func (c *component) AutoRowsMax() *component {
	c.el.Style(styles.AutoRowsMax())
	return c
}

func (c *component) AutoRowsMin() *component {
	c.el.Style(styles.AutoRowsMin())
	return c
}

func (c *component) BackdropBlur(val value.Value) *component {
	c.el.Style(styles.BackdropBlur(val))
	return c
}

func (c *component) BackdropBlur2xl() *component {
	c.el.Style(styles.BackdropBlur2xl())
	return c
}

func (c *component) BackdropBlur3xl() *component {
	c.el.Style(styles.BackdropBlur3xl())
	return c
}

func (c *component) BackdropBlurLg() *component {
	c.el.Style(styles.BackdropBlurLg())
	return c
}

func (c *component) BackdropBlurMd() *component {
	c.el.Style(styles.BackdropBlurMd())
	return c
}

func (c *component) BackdropBlurNone() *component {
	c.el.Style(styles.BackdropBlurNone())
	return c
}

func (c *component) BackdropBlurSm() *component {
	c.el.Style(styles.BackdropBlurSm())
	return c
}

func (c *component) BackdropBlurXl() *component {
	c.el.Style(styles.BackdropBlurXl())
	return c
}

func (c *component) BackdropBlurXs() *component {
	c.el.Style(styles.BackdropBlurXs())
	return c
}

func (c *component) BackdropBrightness(val any) *component {
	c.el.Style(styles.BackdropBrightness(val))
	return c
}

func (c *component) BackdropContrast(val any) *component {
	c.el.Style(styles.BackdropContrast(val))
	return c
}

func (c *component) BackdropFilter(val value.Value) *component {
	c.el.Style(styles.BackdropFilter(val))
	return c
}

func (c *component) BackdropFilterNone() *component {
	c.el.Style(styles.BackdropFilterNone())
	return c
}

func (c *component) BackdropGrayscale(val ...any) *component {
	c.el.Style(styles.BackdropGrayscale(val...))
	return c
}

func (c *component) BackdropHueRotate(val any) *component {
	c.el.Style(styles.BackdropHueRotate(val))
	return c
}

func (c *component) BackdropInvert(val ...any) *component {
	c.el.Style(styles.BackdropInvert(val...))
	return c
}

func (c *component) BackdropSaturate(val any) *component {
	c.el.Style(styles.BackdropSaturate(val))
	return c
}

func (c *component) BackdropSepia(val ...any) *component {
	c.el.Style(styles.BackdropSepia(val...))
	return c
}

func (c *component) BackfaceHidden() *component {
	c.el.Style(styles.BackfaceHidden())
	return c
}

func (c *component) BackfaceVisible() *component {
	c.el.Style(styles.BackfaceVisible())
	return c
}

func (c *component) Basis(val value.Value) *component {
	c.el.Style(styles.Basis(val))
	return c
}

func (c *component) Basis2xl() *component {
	c.el.Style(styles.Basis2xl())
	return c
}

func (c *component) Basis2xs() *component {
	c.el.Style(styles.Basis2xs())
	return c
}

func (c *component) Basis3xl() *component {
	c.el.Style(styles.Basis3xl())
	return c
}

func (c *component) Basis3xs() *component {
	c.el.Style(styles.Basis3xs())
	return c
}

func (c *component) Basis4xl() *component {
	c.el.Style(styles.Basis4xl())
	return c
}

func (c *component) Basis5xl() *component {
	c.el.Style(styles.Basis5xl())
	return c
}

func (c *component) Basis6xl() *component {
	c.el.Style(styles.Basis6xl())
	return c
}

func (c *component) Basis7xl() *component {
	c.el.Style(styles.Basis7xl())
	return c
}

func (c *component) BasisAuto() *component {
	c.el.Style(styles.BasisAuto())
	return c
}

func (c *component) BasisFraction(fraction float64) *component {
	c.el.Style(styles.BasisFraction(fraction))
	return c
}

func (c *component) BasisFull() *component {
	c.el.Style(styles.BasisFull())
	return c
}

func (c *component) BasisLg() *component {
	c.el.Style(styles.BasisLg())
	return c
}

func (c *component) BasisMd() *component {
	c.el.Style(styles.BasisMd())
	return c
}

func (c *component) BasisSm() *component {
	c.el.Style(styles.BasisSm())
	return c
}

func (c *component) BasisXl() *component {
	c.el.Style(styles.BasisXl())
	return c
}

func (c *component) BasisXs() *component {
	c.el.Style(styles.BasisXs())
	return c
}

func (c *component) Bg(val value.Value) *component {
	c.el.Style(styles.Bg(val))
	return c
}

func (c *component) BgAmber(scale int) *component {
	c.el.Style(styles.BgAmber(scale))
	return c
}

func (c *component) BgAmberAlpha(scale int) *component {
	c.el.Style(styles.BgAmberAlpha(scale))
	return c
}

func (c *component) BgAmberDark(scale int) *component {
	c.el.Style(styles.BgAmberDark(scale))
	return c
}

func (c *component) BgAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.BgAmberDarkAlpha(scale))
	return c
}

func (c *component) BgAuto() *component {
	c.el.Style(styles.BgAuto())
	return c
}

func (c *component) BgBlack() *component {
	c.el.Style(styles.BgBlack())
	return c
}

func (c *component) BgBlackAlpha(scale int) *component {
	c.el.Style(styles.BgBlackAlpha(scale))
	return c
}

func (c *component) BgBlendColor() *component {
	c.el.Style(styles.BgBlendColor())
	return c
}

func (c *component) BgBlendColorBurn() *component {
	c.el.Style(styles.BgBlendColorBurn())
	return c
}

func (c *component) BgBlendColorDodge() *component {
	c.el.Style(styles.BgBlendColorDodge())
	return c
}

func (c *component) BgBlendDarken() *component {
	c.el.Style(styles.BgBlendDarken())
	return c
}

func (c *component) BgBlendDifference() *component {
	c.el.Style(styles.BgBlendDifference())
	return c
}

func (c *component) BgBlendExclusion() *component {
	c.el.Style(styles.BgBlendExclusion())
	return c
}

func (c *component) BgBlendHardLight() *component {
	c.el.Style(styles.BgBlendHardLight())
	return c
}

func (c *component) BgBlendHue() *component {
	c.el.Style(styles.BgBlendHue())
	return c
}

func (c *component) BgBlendLighten() *component {
	c.el.Style(styles.BgBlendLighten())
	return c
}

func (c *component) BgBlendLuminosity() *component {
	c.el.Style(styles.BgBlendLuminosity())
	return c
}

func (c *component) BgBlendMultiply() *component {
	c.el.Style(styles.BgBlendMultiply())
	return c
}

func (c *component) BgBlendNormal() *component {
	c.el.Style(styles.BgBlendNormal())
	return c
}

func (c *component) BgBlendOverlay() *component {
	c.el.Style(styles.BgBlendOverlay())
	return c
}

func (c *component) BgBlendSaturation() *component {
	c.el.Style(styles.BgBlendSaturation())
	return c
}

func (c *component) BgBlendScreen() *component {
	c.el.Style(styles.BgBlendScreen())
	return c
}

func (c *component) BgBlendSoftLight() *component {
	c.el.Style(styles.BgBlendSoftLight())
	return c
}

func (c *component) BgBlue(scale int) *component {
	c.el.Style(styles.BgBlue(scale))
	return c
}

func (c *component) BgBlueAlpha(scale int) *component {
	c.el.Style(styles.BgBlueAlpha(scale))
	return c
}

func (c *component) BgBlueDark(scale int) *component {
	c.el.Style(styles.BgBlueDark(scale))
	return c
}

func (c *component) BgBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.BgBlueDarkAlpha(scale))
	return c
}

func (c *component) BgBottom() *component {
	c.el.Style(styles.BgBottom())
	return c
}

func (c *component) BgBottomLeft() *component {
	c.el.Style(styles.BgBottomLeft())
	return c
}

func (c *component) BgBottomRight() *component {
	c.el.Style(styles.BgBottomRight())
	return c
}

func (c *component) BgBronze(scale int) *component {
	c.el.Style(styles.BgBronze(scale))
	return c
}

func (c *component) BgBronzeAlpha(scale int) *component {
	c.el.Style(styles.BgBronzeAlpha(scale))
	return c
}

func (c *component) BgBronzeDark(scale int) *component {
	c.el.Style(styles.BgBronzeDark(scale))
	return c
}

func (c *component) BgBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.BgBronzeDarkAlpha(scale))
	return c
}

func (c *component) BgBrown(scale int) *component {
	c.el.Style(styles.BgBrown(scale))
	return c
}

func (c *component) BgBrownAlpha(scale int) *component {
	c.el.Style(styles.BgBrownAlpha(scale))
	return c
}

func (c *component) BgBrownDark(scale int) *component {
	c.el.Style(styles.BgBrownDark(scale))
	return c
}

func (c *component) BgBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.BgBrownDarkAlpha(scale))
	return c
}

func (c *component) BgCenter() *component {
	c.el.Style(styles.BgCenter())
	return c
}

func (c *component) BgClipBorder() *component {
	c.el.Style(styles.BgClipBorder())
	return c
}

func (c *component) BgClipContent() *component {
	c.el.Style(styles.BgClipContent())
	return c
}

func (c *component) BgClipPadding() *component {
	c.el.Style(styles.BgClipPadding())
	return c
}

func (c *component) BgClipText() *component {
	c.el.Style(styles.BgClipText())
	return c
}

func (c *component) BgColor(color value.Value) *component {
	c.el.Style(styles.BgColor(color))
	return c
}

func (c *component) BgConic(val value.Value) *component {
	c.el.Style(styles.BgConic(val))
	return c
}

func (c *component) BgContain() *component {
	c.el.Style(styles.BgContain())
	return c
}

func (c *component) BgCover() *component {
	c.el.Style(styles.BgCover())
	return c
}

func (c *component) BgCrimson(scale int) *component {
	c.el.Style(styles.BgCrimson(scale))
	return c
}

func (c *component) BgCrimsonAlpha(scale int) *component {
	c.el.Style(styles.BgCrimsonAlpha(scale))
	return c
}

func (c *component) BgCrimsonDark(scale int) *component {
	c.el.Style(styles.BgCrimsonDark(scale))
	return c
}

func (c *component) BgCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.BgCrimsonDarkAlpha(scale))
	return c
}

func (c *component) BgCurrent() *component {
	c.el.Style(styles.BgCurrent())
	return c
}

func (c *component) BgCyan(scale int) *component {
	c.el.Style(styles.BgCyan(scale))
	return c
}

func (c *component) BgCyanAlpha(scale int) *component {
	c.el.Style(styles.BgCyanAlpha(scale))
	return c
}

func (c *component) BgCyanDark(scale int) *component {
	c.el.Style(styles.BgCyanDark(scale))
	return c
}

func (c *component) BgCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.BgCyanDarkAlpha(scale))
	return c
}

func (c *component) BgFixed() *component {
	c.el.Style(styles.BgFixed())
	return c
}

func (c *component) BgGold(scale int) *component {
	c.el.Style(styles.BgGold(scale))
	return c
}

func (c *component) BgGoldAlpha(scale int) *component {
	c.el.Style(styles.BgGoldAlpha(scale))
	return c
}

func (c *component) BgGoldDark(scale int) *component {
	c.el.Style(styles.BgGoldDark(scale))
	return c
}

func (c *component) BgGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.BgGoldDarkAlpha(scale))
	return c
}

func (c *component) BgGrass(scale int) *component {
	c.el.Style(styles.BgGrass(scale))
	return c
}

func (c *component) BgGrassAlpha(scale int) *component {
	c.el.Style(styles.BgGrassAlpha(scale))
	return c
}

func (c *component) BgGrassDark(scale int) *component {
	c.el.Style(styles.BgGrassDark(scale))
	return c
}

func (c *component) BgGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.BgGrassDarkAlpha(scale))
	return c
}

func (c *component) BgGray(scale int) *component {
	c.el.Style(styles.BgGray(scale))
	return c
}

func (c *component) BgGrayAlpha(scale int) *component {
	c.el.Style(styles.BgGrayAlpha(scale))
	return c
}

func (c *component) BgGrayDark(scale int) *component {
	c.el.Style(styles.BgGrayDark(scale))
	return c
}

func (c *component) BgGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.BgGrayDarkAlpha(scale))
	return c
}

func (c *component) BgGreen(scale int) *component {
	c.el.Style(styles.BgGreen(scale))
	return c
}

func (c *component) BgGreenAlpha(scale int) *component {
	c.el.Style(styles.BgGreenAlpha(scale))
	return c
}

func (c *component) BgGreenDark(scale int) *component {
	c.el.Style(styles.BgGreenDark(scale))
	return c
}

func (c *component) BgGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.BgGreenDarkAlpha(scale))
	return c
}

func (c *component) BgIndigo(scale int) *component {
	c.el.Style(styles.BgIndigo(scale))
	return c
}

func (c *component) BgIndigoAlpha(scale int) *component {
	c.el.Style(styles.BgIndigoAlpha(scale))
	return c
}

func (c *component) BgIndigoDark(scale int) *component {
	c.el.Style(styles.BgIndigoDark(scale))
	return c
}

func (c *component) BgIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.BgIndigoDarkAlpha(scale))
	return c
}

func (c *component) BgInherit() *component {
	c.el.Style(styles.BgInherit())
	return c
}

func (c *component) BgIris(scale int) *component {
	c.el.Style(styles.BgIris(scale))
	return c
}

func (c *component) BgIrisAlpha(scale int) *component {
	c.el.Style(styles.BgIrisAlpha(scale))
	return c
}

func (c *component) BgIrisDark(scale int) *component {
	c.el.Style(styles.BgIrisDark(scale))
	return c
}

func (c *component) BgIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.BgIrisDarkAlpha(scale))
	return c
}

func (c *component) BgJade(scale int) *component {
	c.el.Style(styles.BgJade(scale))
	return c
}

func (c *component) BgJadeAlpha(scale int) *component {
	c.el.Style(styles.BgJadeAlpha(scale))
	return c
}

func (c *component) BgJadeDark(scale int) *component {
	c.el.Style(styles.BgJadeDark(scale))
	return c
}

func (c *component) BgJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.BgJadeDarkAlpha(scale))
	return c
}

func (c *component) BgLeft() *component {
	c.el.Style(styles.BgLeft())
	return c
}

func (c *component) BgLime(scale int) *component {
	c.el.Style(styles.BgLime(scale))
	return c
}

func (c *component) BgLimeAlpha(scale int) *component {
	c.el.Style(styles.BgLimeAlpha(scale))
	return c
}

func (c *component) BgLimeDark(scale int) *component {
	c.el.Style(styles.BgLimeDark(scale))
	return c
}

func (c *component) BgLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.BgLimeDarkAlpha(scale))
	return c
}

func (c *component) BgLinear(val value.Value) *component {
	c.el.Style(styles.BgLinear(val))
	return c
}

func (c *component) BgLinearToB() *component {
	c.el.Style(styles.BgLinearToB())
	return c
}

func (c *component) BgLinearToBl() *component {
	c.el.Style(styles.BgLinearToBl())
	return c
}

func (c *component) BgLinearToBr() *component {
	c.el.Style(styles.BgLinearToBr())
	return c
}

func (c *component) BgLinearToL() *component {
	c.el.Style(styles.BgLinearToL())
	return c
}

func (c *component) BgLinearToR() *component {
	c.el.Style(styles.BgLinearToR())
	return c
}

func (c *component) BgLinearToT() *component {
	c.el.Style(styles.BgLinearToT())
	return c
}

func (c *component) BgLinearToTl() *component {
	c.el.Style(styles.BgLinearToTl())
	return c
}

func (c *component) BgLinearToTr() *component {
	c.el.Style(styles.BgLinearToTr())
	return c
}

func (c *component) BgLocal() *component {
	c.el.Style(styles.BgLocal())
	return c
}

func (c *component) BgMauve(scale int) *component {
	c.el.Style(styles.BgMauve(scale))
	return c
}

func (c *component) BgMauveAlpha(scale int) *component {
	c.el.Style(styles.BgMauveAlpha(scale))
	return c
}

func (c *component) BgMauveDark(scale int) *component {
	c.el.Style(styles.BgMauveDark(scale))
	return c
}

func (c *component) BgMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.BgMauveDarkAlpha(scale))
	return c
}

func (c *component) BgMint(scale int) *component {
	c.el.Style(styles.BgMint(scale))
	return c
}

func (c *component) BgMintAlpha(scale int) *component {
	c.el.Style(styles.BgMintAlpha(scale))
	return c
}

func (c *component) BgMintDark(scale int) *component {
	c.el.Style(styles.BgMintDark(scale))
	return c
}

func (c *component) BgMintDarkAlpha(scale int) *component {
	c.el.Style(styles.BgMintDarkAlpha(scale))
	return c
}

func (c *component) BgNoRepeat() *component {
	c.el.Style(styles.BgNoRepeat())
	return c
}

func (c *component) BgNone() *component {
	c.el.Style(styles.BgNone())
	return c
}

func (c *component) BgOlive(scale int) *component {
	c.el.Style(styles.BgOlive(scale))
	return c
}

func (c *component) BgOliveAlpha(scale int) *component {
	c.el.Style(styles.BgOliveAlpha(scale))
	return c
}

func (c *component) BgOliveDark(scale int) *component {
	c.el.Style(styles.BgOliveDark(scale))
	return c
}

func (c *component) BgOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.BgOliveDarkAlpha(scale))
	return c
}

func (c *component) BgOrange(scale int) *component {
	c.el.Style(styles.BgOrange(scale))
	return c
}

func (c *component) BgOrangeAlpha(scale int) *component {
	c.el.Style(styles.BgOrangeAlpha(scale))
	return c
}

func (c *component) BgOrangeDark(scale int) *component {
	c.el.Style(styles.BgOrangeDark(scale))
	return c
}

func (c *component) BgOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.BgOrangeDarkAlpha(scale))
	return c
}

func (c *component) BgOriginBorder() *component {
	c.el.Style(styles.BgOriginBorder())
	return c
}

func (c *component) BgOriginContent() *component {
	c.el.Style(styles.BgOriginContent())
	return c
}

func (c *component) BgOriginPadding() *component {
	c.el.Style(styles.BgOriginPadding())
	return c
}

func (c *component) BgPink(scale int) *component {
	c.el.Style(styles.BgPink(scale))
	return c
}

func (c *component) BgPinkAlpha(scale int) *component {
	c.el.Style(styles.BgPinkAlpha(scale))
	return c
}

func (c *component) BgPinkDark(scale int) *component {
	c.el.Style(styles.BgPinkDark(scale))
	return c
}

func (c *component) BgPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.BgPinkDarkAlpha(scale))
	return c
}

func (c *component) BgPlum(scale int) *component {
	c.el.Style(styles.BgPlum(scale))
	return c
}

func (c *component) BgPlumAlpha(scale int) *component {
	c.el.Style(styles.BgPlumAlpha(scale))
	return c
}

func (c *component) BgPlumDark(scale int) *component {
	c.el.Style(styles.BgPlumDark(scale))
	return c
}

func (c *component) BgPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.BgPlumDarkAlpha(scale))
	return c
}

func (c *component) BgPosition(val value.Value) *component {
	c.el.Style(styles.BgPosition(val))
	return c
}

func (c *component) BgPurple(scale int) *component {
	c.el.Style(styles.BgPurple(scale))
	return c
}

func (c *component) BgPurpleAlpha(scale int) *component {
	c.el.Style(styles.BgPurpleAlpha(scale))
	return c
}

func (c *component) BgPurpleDark(scale int) *component {
	c.el.Style(styles.BgPurpleDark(scale))
	return c
}

func (c *component) BgPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.BgPurpleDarkAlpha(scale))
	return c
}

func (c *component) BgRadial(val value.Value) *component {
	c.el.Style(styles.BgRadial(val))
	return c
}

func (c *component) BgRed(scale int) *component {
	c.el.Style(styles.BgRed(scale))
	return c
}

func (c *component) BgRedAlpha(scale int) *component {
	c.el.Style(styles.BgRedAlpha(scale))
	return c
}

func (c *component) BgRedDark(scale int) *component {
	c.el.Style(styles.BgRedDark(scale))
	return c
}

func (c *component) BgRedDarkAlpha(scale int) *component {
	c.el.Style(styles.BgRedDarkAlpha(scale))
	return c
}

func (c *component) BgRepeat() *component {
	c.el.Style(styles.BgRepeat())
	return c
}

func (c *component) BgRepeatRound() *component {
	c.el.Style(styles.BgRepeatRound())
	return c
}

func (c *component) BgRepeatSpace() *component {
	c.el.Style(styles.BgRepeatSpace())
	return c
}

func (c *component) BgRepeatX() *component {
	c.el.Style(styles.BgRepeatX())
	return c
}

func (c *component) BgRepeatY() *component {
	c.el.Style(styles.BgRepeatY())
	return c
}

func (c *component) BgRight() *component {
	c.el.Style(styles.BgRight())
	return c
}

func (c *component) BgRuby(scale int) *component {
	c.el.Style(styles.BgRuby(scale))
	return c
}

func (c *component) BgRubyAlpha(scale int) *component {
	c.el.Style(styles.BgRubyAlpha(scale))
	return c
}

func (c *component) BgRubyDark(scale int) *component {
	c.el.Style(styles.BgRubyDark(scale))
	return c
}

func (c *component) BgRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.BgRubyDarkAlpha(scale))
	return c
}

func (c *component) BgSage(scale int) *component {
	c.el.Style(styles.BgSage(scale))
	return c
}

func (c *component) BgSageAlpha(scale int) *component {
	c.el.Style(styles.BgSageAlpha(scale))
	return c
}

func (c *component) BgSageDark(scale int) *component {
	c.el.Style(styles.BgSageDark(scale))
	return c
}

func (c *component) BgSageDarkAlpha(scale int) *component {
	c.el.Style(styles.BgSageDarkAlpha(scale))
	return c
}

func (c *component) BgSand(scale int) *component {
	c.el.Style(styles.BgSand(scale))
	return c
}

func (c *component) BgSandAlpha(scale int) *component {
	c.el.Style(styles.BgSandAlpha(scale))
	return c
}

func (c *component) BgSandDark(scale int) *component {
	c.el.Style(styles.BgSandDark(scale))
	return c
}

func (c *component) BgSandDarkAlpha(scale int) *component {
	c.el.Style(styles.BgSandDarkAlpha(scale))
	return c
}

func (c *component) BgScroll() *component {
	c.el.Style(styles.BgScroll())
	return c
}

func (c *component) BgSize(val value.Value) *component {
	c.el.Style(styles.BgSize(val))
	return c
}

func (c *component) BgSky(scale int) *component {
	c.el.Style(styles.BgSky(scale))
	return c
}

func (c *component) BgSkyAlpha(scale int) *component {
	c.el.Style(styles.BgSkyAlpha(scale))
	return c
}

func (c *component) BgSkyDark(scale int) *component {
	c.el.Style(styles.BgSkyDark(scale))
	return c
}

func (c *component) BgSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.BgSkyDarkAlpha(scale))
	return c
}

func (c *component) BgSlate(scale int) *component {
	c.el.Style(styles.BgSlate(scale))
	return c
}

func (c *component) BgSlateAlpha(scale int) *component {
	c.el.Style(styles.BgSlateAlpha(scale))
	return c
}

func (c *component) BgSlateDark(scale int) *component {
	c.el.Style(styles.BgSlateDark(scale))
	return c
}

func (c *component) BgSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.BgSlateDarkAlpha(scale))
	return c
}

func (c *component) BgTeal(scale int) *component {
	c.el.Style(styles.BgTeal(scale))
	return c
}

func (c *component) BgTealAlpha(scale int) *component {
	c.el.Style(styles.BgTealAlpha(scale))
	return c
}

func (c *component) BgTealDark(scale int) *component {
	c.el.Style(styles.BgTealDark(scale))
	return c
}

func (c *component) BgTealDarkAlpha(scale int) *component {
	c.el.Style(styles.BgTealDarkAlpha(scale))
	return c
}

func (c *component) BgTomato(scale int) *component {
	c.el.Style(styles.BgTomato(scale))
	return c
}

func (c *component) BgTomatoAlpha(scale int) *component {
	c.el.Style(styles.BgTomatoAlpha(scale))
	return c
}

func (c *component) BgTomatoDark(scale int) *component {
	c.el.Style(styles.BgTomatoDark(scale))
	return c
}

func (c *component) BgTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.BgTomatoDarkAlpha(scale))
	return c
}

func (c *component) BgTop() *component {
	c.el.Style(styles.BgTop())
	return c
}

func (c *component) BgTopLeft() *component {
	c.el.Style(styles.BgTopLeft())
	return c
}

func (c *component) BgTopRight() *component {
	c.el.Style(styles.BgTopRight())
	return c
}

func (c *component) BgTransparent() *component {
	c.el.Style(styles.BgTransparent())
	return c
}

func (c *component) BgViolet(scale int) *component {
	c.el.Style(styles.BgViolet(scale))
	return c
}

func (c *component) BgVioletAlpha(scale int) *component {
	c.el.Style(styles.BgVioletAlpha(scale))
	return c
}

func (c *component) BgVioletDark(scale int) *component {
	c.el.Style(styles.BgVioletDark(scale))
	return c
}

func (c *component) BgVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.BgVioletDarkAlpha(scale))
	return c
}

func (c *component) BgWhite() *component {
	c.el.Style(styles.BgWhite())
	return c
}

func (c *component) BgWhiteAlpha(scale int) *component {
	c.el.Style(styles.BgWhiteAlpha(scale))
	return c
}

func (c *component) BgYellow(scale int) *component {
	c.el.Style(styles.BgYellow(scale))
	return c
}

func (c *component) BgYellowAlpha(scale int) *component {
	c.el.Style(styles.BgYellowAlpha(scale))
	return c
}

func (c *component) BgYellowDark(scale int) *component {
	c.el.Style(styles.BgYellowDark(scale))
	return c
}

func (c *component) BgYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.BgYellowDarkAlpha(scale))
	return c
}

func (c *component) Block() *component {
	c.el.Style(styles.Block())
	return c
}

func (c *component) Blur(val value.Value) *component {
	c.el.Style(styles.Blur(val))
	return c
}

func (c *component) Blur2xl() *component {
	c.el.Style(styles.Blur2xl())
	return c
}

func (c *component) Blur3xl() *component {
	c.el.Style(styles.Blur3xl())
	return c
}

func (c *component) BlurLg() *component {
	c.el.Style(styles.BlurLg())
	return c
}

func (c *component) BlurMd() *component {
	c.el.Style(styles.BlurMd())
	return c
}

func (c *component) BlurNone() *component {
	c.el.Style(styles.BlurNone())
	return c
}

func (c *component) BlurSm() *component {
	c.el.Style(styles.BlurSm())
	return c
}

func (c *component) BlurXl() *component {
	c.el.Style(styles.BlurXl())
	return c
}

func (c *component) BlurXs() *component {
	c.el.Style(styles.BlurXs())
	return c
}

func (c *component) Border(val ...value.Value) *component {
	c.el.Style(styles.Border(val...))
	return c
}

func (c *component) BorderAmber(scale int) *component {
	c.el.Style(styles.BorderAmber(scale))
	return c
}

func (c *component) BorderAmberAlpha(scale int) *component {
	c.el.Style(styles.BorderAmberAlpha(scale))
	return c
}

func (c *component) BorderAmberDark(scale int) *component {
	c.el.Style(styles.BorderAmberDark(scale))
	return c
}

func (c *component) BorderAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderAmberDarkAlpha(scale))
	return c
}

func (c *component) BorderB(val ...value.Value) *component {
	c.el.Style(styles.BorderB(val...))
	return c
}

func (c *component) BorderBlack() *component {
	c.el.Style(styles.BorderBlack())
	return c
}

func (c *component) BorderBlackAlpha(scale int) *component {
	c.el.Style(styles.BorderBlackAlpha(scale))
	return c
}

func (c *component) BorderBlue(scale int) *component {
	c.el.Style(styles.BorderBlue(scale))
	return c
}

func (c *component) BorderBlueAlpha(scale int) *component {
	c.el.Style(styles.BorderBlueAlpha(scale))
	return c
}

func (c *component) BorderBlueDark(scale int) *component {
	c.el.Style(styles.BorderBlueDark(scale))
	return c
}

func (c *component) BorderBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBlueDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomAmber(scale int) *component {
	c.el.Style(styles.BorderBottomAmber(scale))
	return c
}

func (c *component) BorderBottomAmberAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomAmberAlpha(scale))
	return c
}

func (c *component) BorderBottomAmberDark(scale int) *component {
	c.el.Style(styles.BorderBottomAmberDark(scale))
	return c
}

func (c *component) BorderBottomAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomAmberDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomBlack() *component {
	c.el.Style(styles.BorderBottomBlack())
	return c
}

func (c *component) BorderBottomBlackAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomBlackAlpha(scale))
	return c
}

func (c *component) BorderBottomBlue(scale int) *component {
	c.el.Style(styles.BorderBottomBlue(scale))
	return c
}

func (c *component) BorderBottomBlueAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomBlueAlpha(scale))
	return c
}

func (c *component) BorderBottomBlueDark(scale int) *component {
	c.el.Style(styles.BorderBottomBlueDark(scale))
	return c
}

func (c *component) BorderBottomBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomBlueDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomBronze(scale int) *component {
	c.el.Style(styles.BorderBottomBronze(scale))
	return c
}

func (c *component) BorderBottomBronzeAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomBronzeAlpha(scale))
	return c
}

func (c *component) BorderBottomBronzeDark(scale int) *component {
	c.el.Style(styles.BorderBottomBronzeDark(scale))
	return c
}

func (c *component) BorderBottomBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomBronzeDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomBrown(scale int) *component {
	c.el.Style(styles.BorderBottomBrown(scale))
	return c
}

func (c *component) BorderBottomBrownAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomBrownAlpha(scale))
	return c
}

func (c *component) BorderBottomBrownDark(scale int) *component {
	c.el.Style(styles.BorderBottomBrownDark(scale))
	return c
}

func (c *component) BorderBottomBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomBrownDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomColor(val value.Value) *component {
	c.el.Style(styles.BorderBottomColor(val))
	return c
}

func (c *component) BorderBottomCrimson(scale int) *component {
	c.el.Style(styles.BorderBottomCrimson(scale))
	return c
}

func (c *component) BorderBottomCrimsonAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomCrimsonAlpha(scale))
	return c
}

func (c *component) BorderBottomCrimsonDark(scale int) *component {
	c.el.Style(styles.BorderBottomCrimsonDark(scale))
	return c
}

func (c *component) BorderBottomCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomCrimsonDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomCurrent() *component {
	c.el.Style(styles.BorderBottomCurrent())
	return c
}

func (c *component) BorderBottomCyan(scale int) *component {
	c.el.Style(styles.BorderBottomCyan(scale))
	return c
}

func (c *component) BorderBottomCyanAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomCyanAlpha(scale))
	return c
}

func (c *component) BorderBottomCyanDark(scale int) *component {
	c.el.Style(styles.BorderBottomCyanDark(scale))
	return c
}

func (c *component) BorderBottomCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomCyanDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomGold(scale int) *component {
	c.el.Style(styles.BorderBottomGold(scale))
	return c
}

func (c *component) BorderBottomGoldAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomGoldAlpha(scale))
	return c
}

func (c *component) BorderBottomGoldDark(scale int) *component {
	c.el.Style(styles.BorderBottomGoldDark(scale))
	return c
}

func (c *component) BorderBottomGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomGoldDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomGrass(scale int) *component {
	c.el.Style(styles.BorderBottomGrass(scale))
	return c
}

func (c *component) BorderBottomGrassAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomGrassAlpha(scale))
	return c
}

func (c *component) BorderBottomGrassDark(scale int) *component {
	c.el.Style(styles.BorderBottomGrassDark(scale))
	return c
}

func (c *component) BorderBottomGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomGrassDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomGray(scale int) *component {
	c.el.Style(styles.BorderBottomGray(scale))
	return c
}

func (c *component) BorderBottomGrayAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomGrayAlpha(scale))
	return c
}

func (c *component) BorderBottomGrayDark(scale int) *component {
	c.el.Style(styles.BorderBottomGrayDark(scale))
	return c
}

func (c *component) BorderBottomGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomGrayDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomGreen(scale int) *component {
	c.el.Style(styles.BorderBottomGreen(scale))
	return c
}

func (c *component) BorderBottomGreenAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomGreenAlpha(scale))
	return c
}

func (c *component) BorderBottomGreenDark(scale int) *component {
	c.el.Style(styles.BorderBottomGreenDark(scale))
	return c
}

func (c *component) BorderBottomGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomGreenDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomIndigo(scale int) *component {
	c.el.Style(styles.BorderBottomIndigo(scale))
	return c
}

func (c *component) BorderBottomIndigoAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomIndigoAlpha(scale))
	return c
}

func (c *component) BorderBottomIndigoDark(scale int) *component {
	c.el.Style(styles.BorderBottomIndigoDark(scale))
	return c
}

func (c *component) BorderBottomIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomIndigoDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomInherit() *component {
	c.el.Style(styles.BorderBottomInherit())
	return c
}

func (c *component) BorderBottomIris(scale int) *component {
	c.el.Style(styles.BorderBottomIris(scale))
	return c
}

func (c *component) BorderBottomIrisAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomIrisAlpha(scale))
	return c
}

func (c *component) BorderBottomIrisDark(scale int) *component {
	c.el.Style(styles.BorderBottomIrisDark(scale))
	return c
}

func (c *component) BorderBottomIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomIrisDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomJade(scale int) *component {
	c.el.Style(styles.BorderBottomJade(scale))
	return c
}

func (c *component) BorderBottomJadeAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomJadeAlpha(scale))
	return c
}

func (c *component) BorderBottomJadeDark(scale int) *component {
	c.el.Style(styles.BorderBottomJadeDark(scale))
	return c
}

func (c *component) BorderBottomJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomJadeDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomLime(scale int) *component {
	c.el.Style(styles.BorderBottomLime(scale))
	return c
}

func (c *component) BorderBottomLimeAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomLimeAlpha(scale))
	return c
}

func (c *component) BorderBottomLimeDark(scale int) *component {
	c.el.Style(styles.BorderBottomLimeDark(scale))
	return c
}

func (c *component) BorderBottomLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomLimeDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomMauve(scale int) *component {
	c.el.Style(styles.BorderBottomMauve(scale))
	return c
}

func (c *component) BorderBottomMauveAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomMauveAlpha(scale))
	return c
}

func (c *component) BorderBottomMauveDark(scale int) *component {
	c.el.Style(styles.BorderBottomMauveDark(scale))
	return c
}

func (c *component) BorderBottomMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomMauveDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomMint(scale int) *component {
	c.el.Style(styles.BorderBottomMint(scale))
	return c
}

func (c *component) BorderBottomMintAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomMintAlpha(scale))
	return c
}

func (c *component) BorderBottomMintDark(scale int) *component {
	c.el.Style(styles.BorderBottomMintDark(scale))
	return c
}

func (c *component) BorderBottomMintDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomMintDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomOlive(scale int) *component {
	c.el.Style(styles.BorderBottomOlive(scale))
	return c
}

func (c *component) BorderBottomOliveAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomOliveAlpha(scale))
	return c
}

func (c *component) BorderBottomOliveDark(scale int) *component {
	c.el.Style(styles.BorderBottomOliveDark(scale))
	return c
}

func (c *component) BorderBottomOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomOliveDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomOrange(scale int) *component {
	c.el.Style(styles.BorderBottomOrange(scale))
	return c
}

func (c *component) BorderBottomOrangeAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomOrangeAlpha(scale))
	return c
}

func (c *component) BorderBottomOrangeDark(scale int) *component {
	c.el.Style(styles.BorderBottomOrangeDark(scale))
	return c
}

func (c *component) BorderBottomOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomOrangeDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomPink(scale int) *component {
	c.el.Style(styles.BorderBottomPink(scale))
	return c
}

func (c *component) BorderBottomPinkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomPinkAlpha(scale))
	return c
}

func (c *component) BorderBottomPinkDark(scale int) *component {
	c.el.Style(styles.BorderBottomPinkDark(scale))
	return c
}

func (c *component) BorderBottomPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomPinkDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomPlum(scale int) *component {
	c.el.Style(styles.BorderBottomPlum(scale))
	return c
}

func (c *component) BorderBottomPlumAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomPlumAlpha(scale))
	return c
}

func (c *component) BorderBottomPlumDark(scale int) *component {
	c.el.Style(styles.BorderBottomPlumDark(scale))
	return c
}

func (c *component) BorderBottomPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomPlumDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomPurple(scale int) *component {
	c.el.Style(styles.BorderBottomPurple(scale))
	return c
}

func (c *component) BorderBottomPurpleAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomPurpleAlpha(scale))
	return c
}

func (c *component) BorderBottomPurpleDark(scale int) *component {
	c.el.Style(styles.BorderBottomPurpleDark(scale))
	return c
}

func (c *component) BorderBottomPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomPurpleDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomRed(scale int) *component {
	c.el.Style(styles.BorderBottomRed(scale))
	return c
}

func (c *component) BorderBottomRedAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomRedAlpha(scale))
	return c
}

func (c *component) BorderBottomRedDark(scale int) *component {
	c.el.Style(styles.BorderBottomRedDark(scale))
	return c
}

func (c *component) BorderBottomRedDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomRedDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomRuby(scale int) *component {
	c.el.Style(styles.BorderBottomRuby(scale))
	return c
}

func (c *component) BorderBottomRubyAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomRubyAlpha(scale))
	return c
}

func (c *component) BorderBottomRubyDark(scale int) *component {
	c.el.Style(styles.BorderBottomRubyDark(scale))
	return c
}

func (c *component) BorderBottomRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomRubyDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomSage(scale int) *component {
	c.el.Style(styles.BorderBottomSage(scale))
	return c
}

func (c *component) BorderBottomSageAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomSageAlpha(scale))
	return c
}

func (c *component) BorderBottomSageDark(scale int) *component {
	c.el.Style(styles.BorderBottomSageDark(scale))
	return c
}

func (c *component) BorderBottomSageDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomSageDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomSand(scale int) *component {
	c.el.Style(styles.BorderBottomSand(scale))
	return c
}

func (c *component) BorderBottomSandAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomSandAlpha(scale))
	return c
}

func (c *component) BorderBottomSandDark(scale int) *component {
	c.el.Style(styles.BorderBottomSandDark(scale))
	return c
}

func (c *component) BorderBottomSandDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomSandDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomSky(scale int) *component {
	c.el.Style(styles.BorderBottomSky(scale))
	return c
}

func (c *component) BorderBottomSkyAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomSkyAlpha(scale))
	return c
}

func (c *component) BorderBottomSkyDark(scale int) *component {
	c.el.Style(styles.BorderBottomSkyDark(scale))
	return c
}

func (c *component) BorderBottomSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomSkyDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomSlate(scale int) *component {
	c.el.Style(styles.BorderBottomSlate(scale))
	return c
}

func (c *component) BorderBottomSlateAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomSlateAlpha(scale))
	return c
}

func (c *component) BorderBottomSlateDark(scale int) *component {
	c.el.Style(styles.BorderBottomSlateDark(scale))
	return c
}

func (c *component) BorderBottomSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomSlateDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomTeal(scale int) *component {
	c.el.Style(styles.BorderBottomTeal(scale))
	return c
}

func (c *component) BorderBottomTealAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomTealAlpha(scale))
	return c
}

func (c *component) BorderBottomTealDark(scale int) *component {
	c.el.Style(styles.BorderBottomTealDark(scale))
	return c
}

func (c *component) BorderBottomTealDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomTealDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomTomato(scale int) *component {
	c.el.Style(styles.BorderBottomTomato(scale))
	return c
}

func (c *component) BorderBottomTomatoAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomTomatoAlpha(scale))
	return c
}

func (c *component) BorderBottomTomatoDark(scale int) *component {
	c.el.Style(styles.BorderBottomTomatoDark(scale))
	return c
}

func (c *component) BorderBottomTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomTomatoDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomTransparent() *component {
	c.el.Style(styles.BorderBottomTransparent())
	return c
}

func (c *component) BorderBottomViolet(scale int) *component {
	c.el.Style(styles.BorderBottomViolet(scale))
	return c
}

func (c *component) BorderBottomVioletAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomVioletAlpha(scale))
	return c
}

func (c *component) BorderBottomVioletDark(scale int) *component {
	c.el.Style(styles.BorderBottomVioletDark(scale))
	return c
}

func (c *component) BorderBottomVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomVioletDarkAlpha(scale))
	return c
}

func (c *component) BorderBottomWhite() *component {
	c.el.Style(styles.BorderBottomWhite())
	return c
}

func (c *component) BorderBottomWhiteAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomWhiteAlpha(scale))
	return c
}

func (c *component) BorderBottomYellow(scale int) *component {
	c.el.Style(styles.BorderBottomYellow(scale))
	return c
}

func (c *component) BorderBottomYellowAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomYellowAlpha(scale))
	return c
}

func (c *component) BorderBottomYellowDark(scale int) *component {
	c.el.Style(styles.BorderBottomYellowDark(scale))
	return c
}

func (c *component) BorderBottomYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBottomYellowDarkAlpha(scale))
	return c
}

func (c *component) BorderBronze(scale int) *component {
	c.el.Style(styles.BorderBronze(scale))
	return c
}

func (c *component) BorderBronzeAlpha(scale int) *component {
	c.el.Style(styles.BorderBronzeAlpha(scale))
	return c
}

func (c *component) BorderBronzeDark(scale int) *component {
	c.el.Style(styles.BorderBronzeDark(scale))
	return c
}

func (c *component) BorderBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBronzeDarkAlpha(scale))
	return c
}

func (c *component) BorderBrown(scale int) *component {
	c.el.Style(styles.BorderBrown(scale))
	return c
}

func (c *component) BorderBrownAlpha(scale int) *component {
	c.el.Style(styles.BorderBrownAlpha(scale))
	return c
}

func (c *component) BorderBrownDark(scale int) *component {
	c.el.Style(styles.BorderBrownDark(scale))
	return c
}

func (c *component) BorderBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderBrownDarkAlpha(scale))
	return c
}

func (c *component) BorderCollapse() *component {
	c.el.Style(styles.BorderCollapse())
	return c
}

func (c *component) BorderColor(val value.Value) *component {
	c.el.Style(styles.BorderColor(val))
	return c
}

func (c *component) BorderCrimson(scale int) *component {
	c.el.Style(styles.BorderCrimson(scale))
	return c
}

func (c *component) BorderCrimsonAlpha(scale int) *component {
	c.el.Style(styles.BorderCrimsonAlpha(scale))
	return c
}

func (c *component) BorderCrimsonDark(scale int) *component {
	c.el.Style(styles.BorderCrimsonDark(scale))
	return c
}

func (c *component) BorderCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderCrimsonDarkAlpha(scale))
	return c
}

func (c *component) BorderCurrent() *component {
	c.el.Style(styles.BorderCurrent())
	return c
}

func (c *component) BorderCyan(scale int) *component {
	c.el.Style(styles.BorderCyan(scale))
	return c
}

func (c *component) BorderCyanAlpha(scale int) *component {
	c.el.Style(styles.BorderCyanAlpha(scale))
	return c
}

func (c *component) BorderCyanDark(scale int) *component {
	c.el.Style(styles.BorderCyanDark(scale))
	return c
}

func (c *component) BorderCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderCyanDarkAlpha(scale))
	return c
}

func (c *component) BorderDashed() *component {
	c.el.Style(styles.BorderDashed())
	return c
}

func (c *component) BorderDotted() *component {
	c.el.Style(styles.BorderDotted())
	return c
}

func (c *component) BorderDouble() *component {
	c.el.Style(styles.BorderDouble())
	return c
}

func (c *component) BorderE(val ...value.Value) *component {
	c.el.Style(styles.BorderE(val...))
	return c
}

func (c *component) BorderEndAmber(scale int) *component {
	c.el.Style(styles.BorderEndAmber(scale))
	return c
}

func (c *component) BorderEndAmberAlpha(scale int) *component {
	c.el.Style(styles.BorderEndAmberAlpha(scale))
	return c
}

func (c *component) BorderEndAmberDark(scale int) *component {
	c.el.Style(styles.BorderEndAmberDark(scale))
	return c
}

func (c *component) BorderEndAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndAmberDarkAlpha(scale))
	return c
}

func (c *component) BorderEndBlack() *component {
	c.el.Style(styles.BorderEndBlack())
	return c
}

func (c *component) BorderEndBlackAlpha(scale int) *component {
	c.el.Style(styles.BorderEndBlackAlpha(scale))
	return c
}

func (c *component) BorderEndBlue(scale int) *component {
	c.el.Style(styles.BorderEndBlue(scale))
	return c
}

func (c *component) BorderEndBlueAlpha(scale int) *component {
	c.el.Style(styles.BorderEndBlueAlpha(scale))
	return c
}

func (c *component) BorderEndBlueDark(scale int) *component {
	c.el.Style(styles.BorderEndBlueDark(scale))
	return c
}

func (c *component) BorderEndBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndBlueDarkAlpha(scale))
	return c
}

func (c *component) BorderEndBronze(scale int) *component {
	c.el.Style(styles.BorderEndBronze(scale))
	return c
}

func (c *component) BorderEndBronzeAlpha(scale int) *component {
	c.el.Style(styles.BorderEndBronzeAlpha(scale))
	return c
}

func (c *component) BorderEndBronzeDark(scale int) *component {
	c.el.Style(styles.BorderEndBronzeDark(scale))
	return c
}

func (c *component) BorderEndBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndBronzeDarkAlpha(scale))
	return c
}

func (c *component) BorderEndBrown(scale int) *component {
	c.el.Style(styles.BorderEndBrown(scale))
	return c
}

func (c *component) BorderEndBrownAlpha(scale int) *component {
	c.el.Style(styles.BorderEndBrownAlpha(scale))
	return c
}

func (c *component) BorderEndBrownDark(scale int) *component {
	c.el.Style(styles.BorderEndBrownDark(scale))
	return c
}

func (c *component) BorderEndBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndBrownDarkAlpha(scale))
	return c
}

func (c *component) BorderEndColor(val value.Value) *component {
	c.el.Style(styles.BorderEndColor(val))
	return c
}

func (c *component) BorderEndCrimson(scale int) *component {
	c.el.Style(styles.BorderEndCrimson(scale))
	return c
}

func (c *component) BorderEndCrimsonAlpha(scale int) *component {
	c.el.Style(styles.BorderEndCrimsonAlpha(scale))
	return c
}

func (c *component) BorderEndCrimsonDark(scale int) *component {
	c.el.Style(styles.BorderEndCrimsonDark(scale))
	return c
}

func (c *component) BorderEndCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndCrimsonDarkAlpha(scale))
	return c
}

func (c *component) BorderEndCurrent() *component {
	c.el.Style(styles.BorderEndCurrent())
	return c
}

func (c *component) BorderEndCyan(scale int) *component {
	c.el.Style(styles.BorderEndCyan(scale))
	return c
}

func (c *component) BorderEndCyanAlpha(scale int) *component {
	c.el.Style(styles.BorderEndCyanAlpha(scale))
	return c
}

func (c *component) BorderEndCyanDark(scale int) *component {
	c.el.Style(styles.BorderEndCyanDark(scale))
	return c
}

func (c *component) BorderEndCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndCyanDarkAlpha(scale))
	return c
}

func (c *component) BorderEndGold(scale int) *component {
	c.el.Style(styles.BorderEndGold(scale))
	return c
}

func (c *component) BorderEndGoldAlpha(scale int) *component {
	c.el.Style(styles.BorderEndGoldAlpha(scale))
	return c
}

func (c *component) BorderEndGoldDark(scale int) *component {
	c.el.Style(styles.BorderEndGoldDark(scale))
	return c
}

func (c *component) BorderEndGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndGoldDarkAlpha(scale))
	return c
}

func (c *component) BorderEndGrass(scale int) *component {
	c.el.Style(styles.BorderEndGrass(scale))
	return c
}

func (c *component) BorderEndGrassAlpha(scale int) *component {
	c.el.Style(styles.BorderEndGrassAlpha(scale))
	return c
}

func (c *component) BorderEndGrassDark(scale int) *component {
	c.el.Style(styles.BorderEndGrassDark(scale))
	return c
}

func (c *component) BorderEndGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndGrassDarkAlpha(scale))
	return c
}

func (c *component) BorderEndGray(scale int) *component {
	c.el.Style(styles.BorderEndGray(scale))
	return c
}

func (c *component) BorderEndGrayAlpha(scale int) *component {
	c.el.Style(styles.BorderEndGrayAlpha(scale))
	return c
}

func (c *component) BorderEndGrayDark(scale int) *component {
	c.el.Style(styles.BorderEndGrayDark(scale))
	return c
}

func (c *component) BorderEndGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndGrayDarkAlpha(scale))
	return c
}

func (c *component) BorderEndGreen(scale int) *component {
	c.el.Style(styles.BorderEndGreen(scale))
	return c
}

func (c *component) BorderEndGreenAlpha(scale int) *component {
	c.el.Style(styles.BorderEndGreenAlpha(scale))
	return c
}

func (c *component) BorderEndGreenDark(scale int) *component {
	c.el.Style(styles.BorderEndGreenDark(scale))
	return c
}

func (c *component) BorderEndGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndGreenDarkAlpha(scale))
	return c
}

func (c *component) BorderEndIndigo(scale int) *component {
	c.el.Style(styles.BorderEndIndigo(scale))
	return c
}

func (c *component) BorderEndIndigoAlpha(scale int) *component {
	c.el.Style(styles.BorderEndIndigoAlpha(scale))
	return c
}

func (c *component) BorderEndIndigoDark(scale int) *component {
	c.el.Style(styles.BorderEndIndigoDark(scale))
	return c
}

func (c *component) BorderEndIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndIndigoDarkAlpha(scale))
	return c
}

func (c *component) BorderEndInherit() *component {
	c.el.Style(styles.BorderEndInherit())
	return c
}

func (c *component) BorderEndIris(scale int) *component {
	c.el.Style(styles.BorderEndIris(scale))
	return c
}

func (c *component) BorderEndIrisAlpha(scale int) *component {
	c.el.Style(styles.BorderEndIrisAlpha(scale))
	return c
}

func (c *component) BorderEndIrisDark(scale int) *component {
	c.el.Style(styles.BorderEndIrisDark(scale))
	return c
}

func (c *component) BorderEndIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndIrisDarkAlpha(scale))
	return c
}

func (c *component) BorderEndJade(scale int) *component {
	c.el.Style(styles.BorderEndJade(scale))
	return c
}

func (c *component) BorderEndJadeAlpha(scale int) *component {
	c.el.Style(styles.BorderEndJadeAlpha(scale))
	return c
}

func (c *component) BorderEndJadeDark(scale int) *component {
	c.el.Style(styles.BorderEndJadeDark(scale))
	return c
}

func (c *component) BorderEndJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndJadeDarkAlpha(scale))
	return c
}

func (c *component) BorderEndLime(scale int) *component {
	c.el.Style(styles.BorderEndLime(scale))
	return c
}

func (c *component) BorderEndLimeAlpha(scale int) *component {
	c.el.Style(styles.BorderEndLimeAlpha(scale))
	return c
}

func (c *component) BorderEndLimeDark(scale int) *component {
	c.el.Style(styles.BorderEndLimeDark(scale))
	return c
}

func (c *component) BorderEndLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndLimeDarkAlpha(scale))
	return c
}

func (c *component) BorderEndMauve(scale int) *component {
	c.el.Style(styles.BorderEndMauve(scale))
	return c
}

func (c *component) BorderEndMauveAlpha(scale int) *component {
	c.el.Style(styles.BorderEndMauveAlpha(scale))
	return c
}

func (c *component) BorderEndMauveDark(scale int) *component {
	c.el.Style(styles.BorderEndMauveDark(scale))
	return c
}

func (c *component) BorderEndMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndMauveDarkAlpha(scale))
	return c
}

func (c *component) BorderEndMint(scale int) *component {
	c.el.Style(styles.BorderEndMint(scale))
	return c
}

func (c *component) BorderEndMintAlpha(scale int) *component {
	c.el.Style(styles.BorderEndMintAlpha(scale))
	return c
}

func (c *component) BorderEndMintDark(scale int) *component {
	c.el.Style(styles.BorderEndMintDark(scale))
	return c
}

func (c *component) BorderEndMintDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndMintDarkAlpha(scale))
	return c
}

func (c *component) BorderEndOlive(scale int) *component {
	c.el.Style(styles.BorderEndOlive(scale))
	return c
}

func (c *component) BorderEndOliveAlpha(scale int) *component {
	c.el.Style(styles.BorderEndOliveAlpha(scale))
	return c
}

func (c *component) BorderEndOliveDark(scale int) *component {
	c.el.Style(styles.BorderEndOliveDark(scale))
	return c
}

func (c *component) BorderEndOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndOliveDarkAlpha(scale))
	return c
}

func (c *component) BorderEndOrange(scale int) *component {
	c.el.Style(styles.BorderEndOrange(scale))
	return c
}

func (c *component) BorderEndOrangeAlpha(scale int) *component {
	c.el.Style(styles.BorderEndOrangeAlpha(scale))
	return c
}

func (c *component) BorderEndOrangeDark(scale int) *component {
	c.el.Style(styles.BorderEndOrangeDark(scale))
	return c
}

func (c *component) BorderEndOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndOrangeDarkAlpha(scale))
	return c
}

func (c *component) BorderEndPink(scale int) *component {
	c.el.Style(styles.BorderEndPink(scale))
	return c
}

func (c *component) BorderEndPinkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndPinkAlpha(scale))
	return c
}

func (c *component) BorderEndPinkDark(scale int) *component {
	c.el.Style(styles.BorderEndPinkDark(scale))
	return c
}

func (c *component) BorderEndPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndPinkDarkAlpha(scale))
	return c
}

func (c *component) BorderEndPlum(scale int) *component {
	c.el.Style(styles.BorderEndPlum(scale))
	return c
}

func (c *component) BorderEndPlumAlpha(scale int) *component {
	c.el.Style(styles.BorderEndPlumAlpha(scale))
	return c
}

func (c *component) BorderEndPlumDark(scale int) *component {
	c.el.Style(styles.BorderEndPlumDark(scale))
	return c
}

func (c *component) BorderEndPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndPlumDarkAlpha(scale))
	return c
}

func (c *component) BorderEndPurple(scale int) *component {
	c.el.Style(styles.BorderEndPurple(scale))
	return c
}

func (c *component) BorderEndPurpleAlpha(scale int) *component {
	c.el.Style(styles.BorderEndPurpleAlpha(scale))
	return c
}

func (c *component) BorderEndPurpleDark(scale int) *component {
	c.el.Style(styles.BorderEndPurpleDark(scale))
	return c
}

func (c *component) BorderEndPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndPurpleDarkAlpha(scale))
	return c
}

func (c *component) BorderEndRed(scale int) *component {
	c.el.Style(styles.BorderEndRed(scale))
	return c
}

func (c *component) BorderEndRedAlpha(scale int) *component {
	c.el.Style(styles.BorderEndRedAlpha(scale))
	return c
}

func (c *component) BorderEndRedDark(scale int) *component {
	c.el.Style(styles.BorderEndRedDark(scale))
	return c
}

func (c *component) BorderEndRedDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndRedDarkAlpha(scale))
	return c
}

func (c *component) BorderEndRuby(scale int) *component {
	c.el.Style(styles.BorderEndRuby(scale))
	return c
}

func (c *component) BorderEndRubyAlpha(scale int) *component {
	c.el.Style(styles.BorderEndRubyAlpha(scale))
	return c
}

func (c *component) BorderEndRubyDark(scale int) *component {
	c.el.Style(styles.BorderEndRubyDark(scale))
	return c
}

func (c *component) BorderEndRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndRubyDarkAlpha(scale))
	return c
}

func (c *component) BorderEndSage(scale int) *component {
	c.el.Style(styles.BorderEndSage(scale))
	return c
}

func (c *component) BorderEndSageAlpha(scale int) *component {
	c.el.Style(styles.BorderEndSageAlpha(scale))
	return c
}

func (c *component) BorderEndSageDark(scale int) *component {
	c.el.Style(styles.BorderEndSageDark(scale))
	return c
}

func (c *component) BorderEndSageDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndSageDarkAlpha(scale))
	return c
}

func (c *component) BorderEndSand(scale int) *component {
	c.el.Style(styles.BorderEndSand(scale))
	return c
}

func (c *component) BorderEndSandAlpha(scale int) *component {
	c.el.Style(styles.BorderEndSandAlpha(scale))
	return c
}

func (c *component) BorderEndSandDark(scale int) *component {
	c.el.Style(styles.BorderEndSandDark(scale))
	return c
}

func (c *component) BorderEndSandDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndSandDarkAlpha(scale))
	return c
}

func (c *component) BorderEndSky(scale int) *component {
	c.el.Style(styles.BorderEndSky(scale))
	return c
}

func (c *component) BorderEndSkyAlpha(scale int) *component {
	c.el.Style(styles.BorderEndSkyAlpha(scale))
	return c
}

func (c *component) BorderEndSkyDark(scale int) *component {
	c.el.Style(styles.BorderEndSkyDark(scale))
	return c
}

func (c *component) BorderEndSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndSkyDarkAlpha(scale))
	return c
}

func (c *component) BorderEndSlate(scale int) *component {
	c.el.Style(styles.BorderEndSlate(scale))
	return c
}

func (c *component) BorderEndSlateAlpha(scale int) *component {
	c.el.Style(styles.BorderEndSlateAlpha(scale))
	return c
}

func (c *component) BorderEndSlateDark(scale int) *component {
	c.el.Style(styles.BorderEndSlateDark(scale))
	return c
}

func (c *component) BorderEndSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndSlateDarkAlpha(scale))
	return c
}

func (c *component) BorderEndTeal(scale int) *component {
	c.el.Style(styles.BorderEndTeal(scale))
	return c
}

func (c *component) BorderEndTealAlpha(scale int) *component {
	c.el.Style(styles.BorderEndTealAlpha(scale))
	return c
}

func (c *component) BorderEndTealDark(scale int) *component {
	c.el.Style(styles.BorderEndTealDark(scale))
	return c
}

func (c *component) BorderEndTealDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndTealDarkAlpha(scale))
	return c
}

func (c *component) BorderEndTomato(scale int) *component {
	c.el.Style(styles.BorderEndTomato(scale))
	return c
}

func (c *component) BorderEndTomatoAlpha(scale int) *component {
	c.el.Style(styles.BorderEndTomatoAlpha(scale))
	return c
}

func (c *component) BorderEndTomatoDark(scale int) *component {
	c.el.Style(styles.BorderEndTomatoDark(scale))
	return c
}

func (c *component) BorderEndTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndTomatoDarkAlpha(scale))
	return c
}

func (c *component) BorderEndTransparent() *component {
	c.el.Style(styles.BorderEndTransparent())
	return c
}

func (c *component) BorderEndViolet(scale int) *component {
	c.el.Style(styles.BorderEndViolet(scale))
	return c
}

func (c *component) BorderEndVioletAlpha(scale int) *component {
	c.el.Style(styles.BorderEndVioletAlpha(scale))
	return c
}

func (c *component) BorderEndVioletDark(scale int) *component {
	c.el.Style(styles.BorderEndVioletDark(scale))
	return c
}

func (c *component) BorderEndVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndVioletDarkAlpha(scale))
	return c
}

func (c *component) BorderEndWhite() *component {
	c.el.Style(styles.BorderEndWhite())
	return c
}

func (c *component) BorderEndWhiteAlpha(scale int) *component {
	c.el.Style(styles.BorderEndWhiteAlpha(scale))
	return c
}

func (c *component) BorderEndYellow(scale int) *component {
	c.el.Style(styles.BorderEndYellow(scale))
	return c
}

func (c *component) BorderEndYellowAlpha(scale int) *component {
	c.el.Style(styles.BorderEndYellowAlpha(scale))
	return c
}

func (c *component) BorderEndYellowDark(scale int) *component {
	c.el.Style(styles.BorderEndYellowDark(scale))
	return c
}

func (c *component) BorderEndYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderEndYellowDarkAlpha(scale))
	return c
}

func (c *component) BorderGold(scale int) *component {
	c.el.Style(styles.BorderGold(scale))
	return c
}

func (c *component) BorderGoldAlpha(scale int) *component {
	c.el.Style(styles.BorderGoldAlpha(scale))
	return c
}

func (c *component) BorderGoldDark(scale int) *component {
	c.el.Style(styles.BorderGoldDark(scale))
	return c
}

func (c *component) BorderGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderGoldDarkAlpha(scale))
	return c
}

func (c *component) BorderGrass(scale int) *component {
	c.el.Style(styles.BorderGrass(scale))
	return c
}

func (c *component) BorderGrassAlpha(scale int) *component {
	c.el.Style(styles.BorderGrassAlpha(scale))
	return c
}

func (c *component) BorderGrassDark(scale int) *component {
	c.el.Style(styles.BorderGrassDark(scale))
	return c
}

func (c *component) BorderGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderGrassDarkAlpha(scale))
	return c
}

func (c *component) BorderGray(scale int) *component {
	c.el.Style(styles.BorderGray(scale))
	return c
}

func (c *component) BorderGrayAlpha(scale int) *component {
	c.el.Style(styles.BorderGrayAlpha(scale))
	return c
}

func (c *component) BorderGrayDark(scale int) *component {
	c.el.Style(styles.BorderGrayDark(scale))
	return c
}

func (c *component) BorderGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderGrayDarkAlpha(scale))
	return c
}

func (c *component) BorderGreen(scale int) *component {
	c.el.Style(styles.BorderGreen(scale))
	return c
}

func (c *component) BorderGreenAlpha(scale int) *component {
	c.el.Style(styles.BorderGreenAlpha(scale))
	return c
}

func (c *component) BorderGreenDark(scale int) *component {
	c.el.Style(styles.BorderGreenDark(scale))
	return c
}

func (c *component) BorderGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderGreenDarkAlpha(scale))
	return c
}

func (c *component) BorderHidden() *component {
	c.el.Style(styles.BorderHidden())
	return c
}

func (c *component) BorderIndigo(scale int) *component {
	c.el.Style(styles.BorderIndigo(scale))
	return c
}

func (c *component) BorderIndigoAlpha(scale int) *component {
	c.el.Style(styles.BorderIndigoAlpha(scale))
	return c
}

func (c *component) BorderIndigoDark(scale int) *component {
	c.el.Style(styles.BorderIndigoDark(scale))
	return c
}

func (c *component) BorderIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderIndigoDarkAlpha(scale))
	return c
}

func (c *component) BorderInherit() *component {
	c.el.Style(styles.BorderInherit())
	return c
}

func (c *component) BorderIris(scale int) *component {
	c.el.Style(styles.BorderIris(scale))
	return c
}

func (c *component) BorderIrisAlpha(scale int) *component {
	c.el.Style(styles.BorderIrisAlpha(scale))
	return c
}

func (c *component) BorderIrisDark(scale int) *component {
	c.el.Style(styles.BorderIrisDark(scale))
	return c
}

func (c *component) BorderIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderIrisDarkAlpha(scale))
	return c
}

func (c *component) BorderJade(scale int) *component {
	c.el.Style(styles.BorderJade(scale))
	return c
}

func (c *component) BorderJadeAlpha(scale int) *component {
	c.el.Style(styles.BorderJadeAlpha(scale))
	return c
}

func (c *component) BorderJadeDark(scale int) *component {
	c.el.Style(styles.BorderJadeDark(scale))
	return c
}

func (c *component) BorderJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderJadeDarkAlpha(scale))
	return c
}

func (c *component) BorderL(val ...value.Value) *component {
	c.el.Style(styles.BorderL(val...))
	return c
}

func (c *component) BorderLeftAmber(scale int) *component {
	c.el.Style(styles.BorderLeftAmber(scale))
	return c
}

func (c *component) BorderLeftAmberAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftAmberAlpha(scale))
	return c
}

func (c *component) BorderLeftAmberDark(scale int) *component {
	c.el.Style(styles.BorderLeftAmberDark(scale))
	return c
}

func (c *component) BorderLeftAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftAmberDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftBlack() *component {
	c.el.Style(styles.BorderLeftBlack())
	return c
}

func (c *component) BorderLeftBlackAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftBlackAlpha(scale))
	return c
}

func (c *component) BorderLeftBlue(scale int) *component {
	c.el.Style(styles.BorderLeftBlue(scale))
	return c
}

func (c *component) BorderLeftBlueAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftBlueAlpha(scale))
	return c
}

func (c *component) BorderLeftBlueDark(scale int) *component {
	c.el.Style(styles.BorderLeftBlueDark(scale))
	return c
}

func (c *component) BorderLeftBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftBlueDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftBronze(scale int) *component {
	c.el.Style(styles.BorderLeftBronze(scale))
	return c
}

func (c *component) BorderLeftBronzeAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftBronzeAlpha(scale))
	return c
}

func (c *component) BorderLeftBronzeDark(scale int) *component {
	c.el.Style(styles.BorderLeftBronzeDark(scale))
	return c
}

func (c *component) BorderLeftBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftBronzeDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftBrown(scale int) *component {
	c.el.Style(styles.BorderLeftBrown(scale))
	return c
}

func (c *component) BorderLeftBrownAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftBrownAlpha(scale))
	return c
}

func (c *component) BorderLeftBrownDark(scale int) *component {
	c.el.Style(styles.BorderLeftBrownDark(scale))
	return c
}

func (c *component) BorderLeftBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftBrownDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftColor(val value.Value) *component {
	c.el.Style(styles.BorderLeftColor(val))
	return c
}

func (c *component) BorderLeftCrimson(scale int) *component {
	c.el.Style(styles.BorderLeftCrimson(scale))
	return c
}

func (c *component) BorderLeftCrimsonAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftCrimsonAlpha(scale))
	return c
}

func (c *component) BorderLeftCrimsonDark(scale int) *component {
	c.el.Style(styles.BorderLeftCrimsonDark(scale))
	return c
}

func (c *component) BorderLeftCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftCrimsonDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftCurrent() *component {
	c.el.Style(styles.BorderLeftCurrent())
	return c
}

func (c *component) BorderLeftCyan(scale int) *component {
	c.el.Style(styles.BorderLeftCyan(scale))
	return c
}

func (c *component) BorderLeftCyanAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftCyanAlpha(scale))
	return c
}

func (c *component) BorderLeftCyanDark(scale int) *component {
	c.el.Style(styles.BorderLeftCyanDark(scale))
	return c
}

func (c *component) BorderLeftCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftCyanDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftGold(scale int) *component {
	c.el.Style(styles.BorderLeftGold(scale))
	return c
}

func (c *component) BorderLeftGoldAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftGoldAlpha(scale))
	return c
}

func (c *component) BorderLeftGoldDark(scale int) *component {
	c.el.Style(styles.BorderLeftGoldDark(scale))
	return c
}

func (c *component) BorderLeftGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftGoldDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftGrass(scale int) *component {
	c.el.Style(styles.BorderLeftGrass(scale))
	return c
}

func (c *component) BorderLeftGrassAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftGrassAlpha(scale))
	return c
}

func (c *component) BorderLeftGrassDark(scale int) *component {
	c.el.Style(styles.BorderLeftGrassDark(scale))
	return c
}

func (c *component) BorderLeftGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftGrassDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftGray(scale int) *component {
	c.el.Style(styles.BorderLeftGray(scale))
	return c
}

func (c *component) BorderLeftGrayAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftGrayAlpha(scale))
	return c
}

func (c *component) BorderLeftGrayDark(scale int) *component {
	c.el.Style(styles.BorderLeftGrayDark(scale))
	return c
}

func (c *component) BorderLeftGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftGrayDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftGreen(scale int) *component {
	c.el.Style(styles.BorderLeftGreen(scale))
	return c
}

func (c *component) BorderLeftGreenAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftGreenAlpha(scale))
	return c
}

func (c *component) BorderLeftGreenDark(scale int) *component {
	c.el.Style(styles.BorderLeftGreenDark(scale))
	return c
}

func (c *component) BorderLeftGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftGreenDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftIndigo(scale int) *component {
	c.el.Style(styles.BorderLeftIndigo(scale))
	return c
}

func (c *component) BorderLeftIndigoAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftIndigoAlpha(scale))
	return c
}

func (c *component) BorderLeftIndigoDark(scale int) *component {
	c.el.Style(styles.BorderLeftIndigoDark(scale))
	return c
}

func (c *component) BorderLeftIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftIndigoDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftInherit() *component {
	c.el.Style(styles.BorderLeftInherit())
	return c
}

func (c *component) BorderLeftIris(scale int) *component {
	c.el.Style(styles.BorderLeftIris(scale))
	return c
}

func (c *component) BorderLeftIrisAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftIrisAlpha(scale))
	return c
}

func (c *component) BorderLeftIrisDark(scale int) *component {
	c.el.Style(styles.BorderLeftIrisDark(scale))
	return c
}

func (c *component) BorderLeftIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftIrisDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftJade(scale int) *component {
	c.el.Style(styles.BorderLeftJade(scale))
	return c
}

func (c *component) BorderLeftJadeAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftJadeAlpha(scale))
	return c
}

func (c *component) BorderLeftJadeDark(scale int) *component {
	c.el.Style(styles.BorderLeftJadeDark(scale))
	return c
}

func (c *component) BorderLeftJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftJadeDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftLime(scale int) *component {
	c.el.Style(styles.BorderLeftLime(scale))
	return c
}

func (c *component) BorderLeftLimeAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftLimeAlpha(scale))
	return c
}

func (c *component) BorderLeftLimeDark(scale int) *component {
	c.el.Style(styles.BorderLeftLimeDark(scale))
	return c
}

func (c *component) BorderLeftLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftLimeDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftMauve(scale int) *component {
	c.el.Style(styles.BorderLeftMauve(scale))
	return c
}

func (c *component) BorderLeftMauveAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftMauveAlpha(scale))
	return c
}

func (c *component) BorderLeftMauveDark(scale int) *component {
	c.el.Style(styles.BorderLeftMauveDark(scale))
	return c
}

func (c *component) BorderLeftMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftMauveDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftMint(scale int) *component {
	c.el.Style(styles.BorderLeftMint(scale))
	return c
}

func (c *component) BorderLeftMintAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftMintAlpha(scale))
	return c
}

func (c *component) BorderLeftMintDark(scale int) *component {
	c.el.Style(styles.BorderLeftMintDark(scale))
	return c
}

func (c *component) BorderLeftMintDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftMintDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftOlive(scale int) *component {
	c.el.Style(styles.BorderLeftOlive(scale))
	return c
}

func (c *component) BorderLeftOliveAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftOliveAlpha(scale))
	return c
}

func (c *component) BorderLeftOliveDark(scale int) *component {
	c.el.Style(styles.BorderLeftOliveDark(scale))
	return c
}

func (c *component) BorderLeftOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftOliveDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftOrange(scale int) *component {
	c.el.Style(styles.BorderLeftOrange(scale))
	return c
}

func (c *component) BorderLeftOrangeAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftOrangeAlpha(scale))
	return c
}

func (c *component) BorderLeftOrangeDark(scale int) *component {
	c.el.Style(styles.BorderLeftOrangeDark(scale))
	return c
}

func (c *component) BorderLeftOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftOrangeDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftPink(scale int) *component {
	c.el.Style(styles.BorderLeftPink(scale))
	return c
}

func (c *component) BorderLeftPinkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftPinkAlpha(scale))
	return c
}

func (c *component) BorderLeftPinkDark(scale int) *component {
	c.el.Style(styles.BorderLeftPinkDark(scale))
	return c
}

func (c *component) BorderLeftPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftPinkDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftPlum(scale int) *component {
	c.el.Style(styles.BorderLeftPlum(scale))
	return c
}

func (c *component) BorderLeftPlumAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftPlumAlpha(scale))
	return c
}

func (c *component) BorderLeftPlumDark(scale int) *component {
	c.el.Style(styles.BorderLeftPlumDark(scale))
	return c
}

func (c *component) BorderLeftPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftPlumDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftPurple(scale int) *component {
	c.el.Style(styles.BorderLeftPurple(scale))
	return c
}

func (c *component) BorderLeftPurpleAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftPurpleAlpha(scale))
	return c
}

func (c *component) BorderLeftPurpleDark(scale int) *component {
	c.el.Style(styles.BorderLeftPurpleDark(scale))
	return c
}

func (c *component) BorderLeftPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftPurpleDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftRed(scale int) *component {
	c.el.Style(styles.BorderLeftRed(scale))
	return c
}

func (c *component) BorderLeftRedAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftRedAlpha(scale))
	return c
}

func (c *component) BorderLeftRedDark(scale int) *component {
	c.el.Style(styles.BorderLeftRedDark(scale))
	return c
}

func (c *component) BorderLeftRedDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftRedDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftRuby(scale int) *component {
	c.el.Style(styles.BorderLeftRuby(scale))
	return c
}

func (c *component) BorderLeftRubyAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftRubyAlpha(scale))
	return c
}

func (c *component) BorderLeftRubyDark(scale int) *component {
	c.el.Style(styles.BorderLeftRubyDark(scale))
	return c
}

func (c *component) BorderLeftRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftRubyDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftSage(scale int) *component {
	c.el.Style(styles.BorderLeftSage(scale))
	return c
}

func (c *component) BorderLeftSageAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftSageAlpha(scale))
	return c
}

func (c *component) BorderLeftSageDark(scale int) *component {
	c.el.Style(styles.BorderLeftSageDark(scale))
	return c
}

func (c *component) BorderLeftSageDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftSageDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftSand(scale int) *component {
	c.el.Style(styles.BorderLeftSand(scale))
	return c
}

func (c *component) BorderLeftSandAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftSandAlpha(scale))
	return c
}

func (c *component) BorderLeftSandDark(scale int) *component {
	c.el.Style(styles.BorderLeftSandDark(scale))
	return c
}

func (c *component) BorderLeftSandDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftSandDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftSky(scale int) *component {
	c.el.Style(styles.BorderLeftSky(scale))
	return c
}

func (c *component) BorderLeftSkyAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftSkyAlpha(scale))
	return c
}

func (c *component) BorderLeftSkyDark(scale int) *component {
	c.el.Style(styles.BorderLeftSkyDark(scale))
	return c
}

func (c *component) BorderLeftSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftSkyDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftSlate(scale int) *component {
	c.el.Style(styles.BorderLeftSlate(scale))
	return c
}

func (c *component) BorderLeftSlateAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftSlateAlpha(scale))
	return c
}

func (c *component) BorderLeftSlateDark(scale int) *component {
	c.el.Style(styles.BorderLeftSlateDark(scale))
	return c
}

func (c *component) BorderLeftSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftSlateDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftTeal(scale int) *component {
	c.el.Style(styles.BorderLeftTeal(scale))
	return c
}

func (c *component) BorderLeftTealAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftTealAlpha(scale))
	return c
}

func (c *component) BorderLeftTealDark(scale int) *component {
	c.el.Style(styles.BorderLeftTealDark(scale))
	return c
}

func (c *component) BorderLeftTealDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftTealDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftTomato(scale int) *component {
	c.el.Style(styles.BorderLeftTomato(scale))
	return c
}

func (c *component) BorderLeftTomatoAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftTomatoAlpha(scale))
	return c
}

func (c *component) BorderLeftTomatoDark(scale int) *component {
	c.el.Style(styles.BorderLeftTomatoDark(scale))
	return c
}

func (c *component) BorderLeftTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftTomatoDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftTransparent() *component {
	c.el.Style(styles.BorderLeftTransparent())
	return c
}

func (c *component) BorderLeftViolet(scale int) *component {
	c.el.Style(styles.BorderLeftViolet(scale))
	return c
}

func (c *component) BorderLeftVioletAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftVioletAlpha(scale))
	return c
}

func (c *component) BorderLeftVioletDark(scale int) *component {
	c.el.Style(styles.BorderLeftVioletDark(scale))
	return c
}

func (c *component) BorderLeftVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftVioletDarkAlpha(scale))
	return c
}

func (c *component) BorderLeftWhite() *component {
	c.el.Style(styles.BorderLeftWhite())
	return c
}

func (c *component) BorderLeftWhiteAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftWhiteAlpha(scale))
	return c
}

func (c *component) BorderLeftYellow(scale int) *component {
	c.el.Style(styles.BorderLeftYellow(scale))
	return c
}

func (c *component) BorderLeftYellowAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftYellowAlpha(scale))
	return c
}

func (c *component) BorderLeftYellowDark(scale int) *component {
	c.el.Style(styles.BorderLeftYellowDark(scale))
	return c
}

func (c *component) BorderLeftYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLeftYellowDarkAlpha(scale))
	return c
}

func (c *component) BorderLime(scale int) *component {
	c.el.Style(styles.BorderLime(scale))
	return c
}

func (c *component) BorderLimeAlpha(scale int) *component {
	c.el.Style(styles.BorderLimeAlpha(scale))
	return c
}

func (c *component) BorderLimeDark(scale int) *component {
	c.el.Style(styles.BorderLimeDark(scale))
	return c
}

func (c *component) BorderLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderLimeDarkAlpha(scale))
	return c
}

func (c *component) BorderMauve(scale int) *component {
	c.el.Style(styles.BorderMauve(scale))
	return c
}

func (c *component) BorderMauveAlpha(scale int) *component {
	c.el.Style(styles.BorderMauveAlpha(scale))
	return c
}

func (c *component) BorderMauveDark(scale int) *component {
	c.el.Style(styles.BorderMauveDark(scale))
	return c
}

func (c *component) BorderMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderMauveDarkAlpha(scale))
	return c
}

func (c *component) BorderMint(scale int) *component {
	c.el.Style(styles.BorderMint(scale))
	return c
}

func (c *component) BorderMintAlpha(scale int) *component {
	c.el.Style(styles.BorderMintAlpha(scale))
	return c
}

func (c *component) BorderMintDark(scale int) *component {
	c.el.Style(styles.BorderMintDark(scale))
	return c
}

func (c *component) BorderMintDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderMintDarkAlpha(scale))
	return c
}

func (c *component) BorderNone() *component {
	c.el.Style(styles.BorderNone())
	return c
}

func (c *component) BorderOlive(scale int) *component {
	c.el.Style(styles.BorderOlive(scale))
	return c
}

func (c *component) BorderOliveAlpha(scale int) *component {
	c.el.Style(styles.BorderOliveAlpha(scale))
	return c
}

func (c *component) BorderOliveDark(scale int) *component {
	c.el.Style(styles.BorderOliveDark(scale))
	return c
}

func (c *component) BorderOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderOliveDarkAlpha(scale))
	return c
}

func (c *component) BorderOrange(scale int) *component {
	c.el.Style(styles.BorderOrange(scale))
	return c
}

func (c *component) BorderOrangeAlpha(scale int) *component {
	c.el.Style(styles.BorderOrangeAlpha(scale))
	return c
}

func (c *component) BorderOrangeDark(scale int) *component {
	c.el.Style(styles.BorderOrangeDark(scale))
	return c
}

func (c *component) BorderOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderOrangeDarkAlpha(scale))
	return c
}

func (c *component) BorderPink(scale int) *component {
	c.el.Style(styles.BorderPink(scale))
	return c
}

func (c *component) BorderPinkAlpha(scale int) *component {
	c.el.Style(styles.BorderPinkAlpha(scale))
	return c
}

func (c *component) BorderPinkDark(scale int) *component {
	c.el.Style(styles.BorderPinkDark(scale))
	return c
}

func (c *component) BorderPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderPinkDarkAlpha(scale))
	return c
}

func (c *component) BorderPlum(scale int) *component {
	c.el.Style(styles.BorderPlum(scale))
	return c
}

func (c *component) BorderPlumAlpha(scale int) *component {
	c.el.Style(styles.BorderPlumAlpha(scale))
	return c
}

func (c *component) BorderPlumDark(scale int) *component {
	c.el.Style(styles.BorderPlumDark(scale))
	return c
}

func (c *component) BorderPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderPlumDarkAlpha(scale))
	return c
}

func (c *component) BorderPurple(scale int) *component {
	c.el.Style(styles.BorderPurple(scale))
	return c
}

func (c *component) BorderPurpleAlpha(scale int) *component {
	c.el.Style(styles.BorderPurpleAlpha(scale))
	return c
}

func (c *component) BorderPurpleDark(scale int) *component {
	c.el.Style(styles.BorderPurpleDark(scale))
	return c
}

func (c *component) BorderPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderPurpleDarkAlpha(scale))
	return c
}

func (c *component) BorderR(val ...value.Value) *component {
	c.el.Style(styles.BorderR(val...))
	return c
}

func (c *component) BorderRed(scale int) *component {
	c.el.Style(styles.BorderRed(scale))
	return c
}

func (c *component) BorderRedAlpha(scale int) *component {
	c.el.Style(styles.BorderRedAlpha(scale))
	return c
}

func (c *component) BorderRedDark(scale int) *component {
	c.el.Style(styles.BorderRedDark(scale))
	return c
}

func (c *component) BorderRedDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRedDarkAlpha(scale))
	return c
}

func (c *component) BorderRightAmber(scale int) *component {
	c.el.Style(styles.BorderRightAmber(scale))
	return c
}

func (c *component) BorderRightAmberAlpha(scale int) *component {
	c.el.Style(styles.BorderRightAmberAlpha(scale))
	return c
}

func (c *component) BorderRightAmberDark(scale int) *component {
	c.el.Style(styles.BorderRightAmberDark(scale))
	return c
}

func (c *component) BorderRightAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightAmberDarkAlpha(scale))
	return c
}

func (c *component) BorderRightBlack() *component {
	c.el.Style(styles.BorderRightBlack())
	return c
}

func (c *component) BorderRightBlackAlpha(scale int) *component {
	c.el.Style(styles.BorderRightBlackAlpha(scale))
	return c
}

func (c *component) BorderRightBlue(scale int) *component {
	c.el.Style(styles.BorderRightBlue(scale))
	return c
}

func (c *component) BorderRightBlueAlpha(scale int) *component {
	c.el.Style(styles.BorderRightBlueAlpha(scale))
	return c
}

func (c *component) BorderRightBlueDark(scale int) *component {
	c.el.Style(styles.BorderRightBlueDark(scale))
	return c
}

func (c *component) BorderRightBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightBlueDarkAlpha(scale))
	return c
}

func (c *component) BorderRightBronze(scale int) *component {
	c.el.Style(styles.BorderRightBronze(scale))
	return c
}

func (c *component) BorderRightBronzeAlpha(scale int) *component {
	c.el.Style(styles.BorderRightBronzeAlpha(scale))
	return c
}

func (c *component) BorderRightBronzeDark(scale int) *component {
	c.el.Style(styles.BorderRightBronzeDark(scale))
	return c
}

func (c *component) BorderRightBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightBronzeDarkAlpha(scale))
	return c
}

func (c *component) BorderRightBrown(scale int) *component {
	c.el.Style(styles.BorderRightBrown(scale))
	return c
}

func (c *component) BorderRightBrownAlpha(scale int) *component {
	c.el.Style(styles.BorderRightBrownAlpha(scale))
	return c
}

func (c *component) BorderRightBrownDark(scale int) *component {
	c.el.Style(styles.BorderRightBrownDark(scale))
	return c
}

func (c *component) BorderRightBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightBrownDarkAlpha(scale))
	return c
}

func (c *component) BorderRightColor(val value.Value) *component {
	c.el.Style(styles.BorderRightColor(val))
	return c
}

func (c *component) BorderRightCrimson(scale int) *component {
	c.el.Style(styles.BorderRightCrimson(scale))
	return c
}

func (c *component) BorderRightCrimsonAlpha(scale int) *component {
	c.el.Style(styles.BorderRightCrimsonAlpha(scale))
	return c
}

func (c *component) BorderRightCrimsonDark(scale int) *component {
	c.el.Style(styles.BorderRightCrimsonDark(scale))
	return c
}

func (c *component) BorderRightCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightCrimsonDarkAlpha(scale))
	return c
}

func (c *component) BorderRightCurrent() *component {
	c.el.Style(styles.BorderRightCurrent())
	return c
}

func (c *component) BorderRightCyan(scale int) *component {
	c.el.Style(styles.BorderRightCyan(scale))
	return c
}

func (c *component) BorderRightCyanAlpha(scale int) *component {
	c.el.Style(styles.BorderRightCyanAlpha(scale))
	return c
}

func (c *component) BorderRightCyanDark(scale int) *component {
	c.el.Style(styles.BorderRightCyanDark(scale))
	return c
}

func (c *component) BorderRightCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightCyanDarkAlpha(scale))
	return c
}

func (c *component) BorderRightGold(scale int) *component {
	c.el.Style(styles.BorderRightGold(scale))
	return c
}

func (c *component) BorderRightGoldAlpha(scale int) *component {
	c.el.Style(styles.BorderRightGoldAlpha(scale))
	return c
}

func (c *component) BorderRightGoldDark(scale int) *component {
	c.el.Style(styles.BorderRightGoldDark(scale))
	return c
}

func (c *component) BorderRightGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightGoldDarkAlpha(scale))
	return c
}

func (c *component) BorderRightGrass(scale int) *component {
	c.el.Style(styles.BorderRightGrass(scale))
	return c
}

func (c *component) BorderRightGrassAlpha(scale int) *component {
	c.el.Style(styles.BorderRightGrassAlpha(scale))
	return c
}

func (c *component) BorderRightGrassDark(scale int) *component {
	c.el.Style(styles.BorderRightGrassDark(scale))
	return c
}

func (c *component) BorderRightGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightGrassDarkAlpha(scale))
	return c
}

func (c *component) BorderRightGray(scale int) *component {
	c.el.Style(styles.BorderRightGray(scale))
	return c
}

func (c *component) BorderRightGrayAlpha(scale int) *component {
	c.el.Style(styles.BorderRightGrayAlpha(scale))
	return c
}

func (c *component) BorderRightGrayDark(scale int) *component {
	c.el.Style(styles.BorderRightGrayDark(scale))
	return c
}

func (c *component) BorderRightGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightGrayDarkAlpha(scale))
	return c
}

func (c *component) BorderRightGreen(scale int) *component {
	c.el.Style(styles.BorderRightGreen(scale))
	return c
}

func (c *component) BorderRightGreenAlpha(scale int) *component {
	c.el.Style(styles.BorderRightGreenAlpha(scale))
	return c
}

func (c *component) BorderRightGreenDark(scale int) *component {
	c.el.Style(styles.BorderRightGreenDark(scale))
	return c
}

func (c *component) BorderRightGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightGreenDarkAlpha(scale))
	return c
}

func (c *component) BorderRightIndigo(scale int) *component {
	c.el.Style(styles.BorderRightIndigo(scale))
	return c
}

func (c *component) BorderRightIndigoAlpha(scale int) *component {
	c.el.Style(styles.BorderRightIndigoAlpha(scale))
	return c
}

func (c *component) BorderRightIndigoDark(scale int) *component {
	c.el.Style(styles.BorderRightIndigoDark(scale))
	return c
}

func (c *component) BorderRightIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightIndigoDarkAlpha(scale))
	return c
}

func (c *component) BorderRightInherit() *component {
	c.el.Style(styles.BorderRightInherit())
	return c
}

func (c *component) BorderRightIris(scale int) *component {
	c.el.Style(styles.BorderRightIris(scale))
	return c
}

func (c *component) BorderRightIrisAlpha(scale int) *component {
	c.el.Style(styles.BorderRightIrisAlpha(scale))
	return c
}

func (c *component) BorderRightIrisDark(scale int) *component {
	c.el.Style(styles.BorderRightIrisDark(scale))
	return c
}

func (c *component) BorderRightIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightIrisDarkAlpha(scale))
	return c
}

func (c *component) BorderRightJade(scale int) *component {
	c.el.Style(styles.BorderRightJade(scale))
	return c
}

func (c *component) BorderRightJadeAlpha(scale int) *component {
	c.el.Style(styles.BorderRightJadeAlpha(scale))
	return c
}

func (c *component) BorderRightJadeDark(scale int) *component {
	c.el.Style(styles.BorderRightJadeDark(scale))
	return c
}

func (c *component) BorderRightJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightJadeDarkAlpha(scale))
	return c
}

func (c *component) BorderRightLime(scale int) *component {
	c.el.Style(styles.BorderRightLime(scale))
	return c
}

func (c *component) BorderRightLimeAlpha(scale int) *component {
	c.el.Style(styles.BorderRightLimeAlpha(scale))
	return c
}

func (c *component) BorderRightLimeDark(scale int) *component {
	c.el.Style(styles.BorderRightLimeDark(scale))
	return c
}

func (c *component) BorderRightLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightLimeDarkAlpha(scale))
	return c
}

func (c *component) BorderRightMauve(scale int) *component {
	c.el.Style(styles.BorderRightMauve(scale))
	return c
}

func (c *component) BorderRightMauveAlpha(scale int) *component {
	c.el.Style(styles.BorderRightMauveAlpha(scale))
	return c
}

func (c *component) BorderRightMauveDark(scale int) *component {
	c.el.Style(styles.BorderRightMauveDark(scale))
	return c
}

func (c *component) BorderRightMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightMauveDarkAlpha(scale))
	return c
}

func (c *component) BorderRightMint(scale int) *component {
	c.el.Style(styles.BorderRightMint(scale))
	return c
}

func (c *component) BorderRightMintAlpha(scale int) *component {
	c.el.Style(styles.BorderRightMintAlpha(scale))
	return c
}

func (c *component) BorderRightMintDark(scale int) *component {
	c.el.Style(styles.BorderRightMintDark(scale))
	return c
}

func (c *component) BorderRightMintDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightMintDarkAlpha(scale))
	return c
}

func (c *component) BorderRightOlive(scale int) *component {
	c.el.Style(styles.BorderRightOlive(scale))
	return c
}

func (c *component) BorderRightOliveAlpha(scale int) *component {
	c.el.Style(styles.BorderRightOliveAlpha(scale))
	return c
}

func (c *component) BorderRightOliveDark(scale int) *component {
	c.el.Style(styles.BorderRightOliveDark(scale))
	return c
}

func (c *component) BorderRightOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightOliveDarkAlpha(scale))
	return c
}

func (c *component) BorderRightOrange(scale int) *component {
	c.el.Style(styles.BorderRightOrange(scale))
	return c
}

func (c *component) BorderRightOrangeAlpha(scale int) *component {
	c.el.Style(styles.BorderRightOrangeAlpha(scale))
	return c
}

func (c *component) BorderRightOrangeDark(scale int) *component {
	c.el.Style(styles.BorderRightOrangeDark(scale))
	return c
}

func (c *component) BorderRightOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightOrangeDarkAlpha(scale))
	return c
}

func (c *component) BorderRightPink(scale int) *component {
	c.el.Style(styles.BorderRightPink(scale))
	return c
}

func (c *component) BorderRightPinkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightPinkAlpha(scale))
	return c
}

func (c *component) BorderRightPinkDark(scale int) *component {
	c.el.Style(styles.BorderRightPinkDark(scale))
	return c
}

func (c *component) BorderRightPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightPinkDarkAlpha(scale))
	return c
}

func (c *component) BorderRightPlum(scale int) *component {
	c.el.Style(styles.BorderRightPlum(scale))
	return c
}

func (c *component) BorderRightPlumAlpha(scale int) *component {
	c.el.Style(styles.BorderRightPlumAlpha(scale))
	return c
}

func (c *component) BorderRightPlumDark(scale int) *component {
	c.el.Style(styles.BorderRightPlumDark(scale))
	return c
}

func (c *component) BorderRightPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightPlumDarkAlpha(scale))
	return c
}

func (c *component) BorderRightPurple(scale int) *component {
	c.el.Style(styles.BorderRightPurple(scale))
	return c
}

func (c *component) BorderRightPurpleAlpha(scale int) *component {
	c.el.Style(styles.BorderRightPurpleAlpha(scale))
	return c
}

func (c *component) BorderRightPurpleDark(scale int) *component {
	c.el.Style(styles.BorderRightPurpleDark(scale))
	return c
}

func (c *component) BorderRightPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightPurpleDarkAlpha(scale))
	return c
}

func (c *component) BorderRightRed(scale int) *component {
	c.el.Style(styles.BorderRightRed(scale))
	return c
}

func (c *component) BorderRightRedAlpha(scale int) *component {
	c.el.Style(styles.BorderRightRedAlpha(scale))
	return c
}

func (c *component) BorderRightRedDark(scale int) *component {
	c.el.Style(styles.BorderRightRedDark(scale))
	return c
}

func (c *component) BorderRightRedDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightRedDarkAlpha(scale))
	return c
}

func (c *component) BorderRightRuby(scale int) *component {
	c.el.Style(styles.BorderRightRuby(scale))
	return c
}

func (c *component) BorderRightRubyAlpha(scale int) *component {
	c.el.Style(styles.BorderRightRubyAlpha(scale))
	return c
}

func (c *component) BorderRightRubyDark(scale int) *component {
	c.el.Style(styles.BorderRightRubyDark(scale))
	return c
}

func (c *component) BorderRightRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightRubyDarkAlpha(scale))
	return c
}

func (c *component) BorderRightSage(scale int) *component {
	c.el.Style(styles.BorderRightSage(scale))
	return c
}

func (c *component) BorderRightSageAlpha(scale int) *component {
	c.el.Style(styles.BorderRightSageAlpha(scale))
	return c
}

func (c *component) BorderRightSageDark(scale int) *component {
	c.el.Style(styles.BorderRightSageDark(scale))
	return c
}

func (c *component) BorderRightSageDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightSageDarkAlpha(scale))
	return c
}

func (c *component) BorderRightSand(scale int) *component {
	c.el.Style(styles.BorderRightSand(scale))
	return c
}

func (c *component) BorderRightSandAlpha(scale int) *component {
	c.el.Style(styles.BorderRightSandAlpha(scale))
	return c
}

func (c *component) BorderRightSandDark(scale int) *component {
	c.el.Style(styles.BorderRightSandDark(scale))
	return c
}

func (c *component) BorderRightSandDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightSandDarkAlpha(scale))
	return c
}

func (c *component) BorderRightSky(scale int) *component {
	c.el.Style(styles.BorderRightSky(scale))
	return c
}

func (c *component) BorderRightSkyAlpha(scale int) *component {
	c.el.Style(styles.BorderRightSkyAlpha(scale))
	return c
}

func (c *component) BorderRightSkyDark(scale int) *component {
	c.el.Style(styles.BorderRightSkyDark(scale))
	return c
}

func (c *component) BorderRightSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightSkyDarkAlpha(scale))
	return c
}

func (c *component) BorderRightSlate(scale int) *component {
	c.el.Style(styles.BorderRightSlate(scale))
	return c
}

func (c *component) BorderRightSlateAlpha(scale int) *component {
	c.el.Style(styles.BorderRightSlateAlpha(scale))
	return c
}

func (c *component) BorderRightSlateDark(scale int) *component {
	c.el.Style(styles.BorderRightSlateDark(scale))
	return c
}

func (c *component) BorderRightSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightSlateDarkAlpha(scale))
	return c
}

func (c *component) BorderRightTeal(scale int) *component {
	c.el.Style(styles.BorderRightTeal(scale))
	return c
}

func (c *component) BorderRightTealAlpha(scale int) *component {
	c.el.Style(styles.BorderRightTealAlpha(scale))
	return c
}

func (c *component) BorderRightTealDark(scale int) *component {
	c.el.Style(styles.BorderRightTealDark(scale))
	return c
}

func (c *component) BorderRightTealDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightTealDarkAlpha(scale))
	return c
}

func (c *component) BorderRightTomato(scale int) *component {
	c.el.Style(styles.BorderRightTomato(scale))
	return c
}

func (c *component) BorderRightTomatoAlpha(scale int) *component {
	c.el.Style(styles.BorderRightTomatoAlpha(scale))
	return c
}

func (c *component) BorderRightTomatoDark(scale int) *component {
	c.el.Style(styles.BorderRightTomatoDark(scale))
	return c
}

func (c *component) BorderRightTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightTomatoDarkAlpha(scale))
	return c
}

func (c *component) BorderRightTransparent() *component {
	c.el.Style(styles.BorderRightTransparent())
	return c
}

func (c *component) BorderRightViolet(scale int) *component {
	c.el.Style(styles.BorderRightViolet(scale))
	return c
}

func (c *component) BorderRightVioletAlpha(scale int) *component {
	c.el.Style(styles.BorderRightVioletAlpha(scale))
	return c
}

func (c *component) BorderRightVioletDark(scale int) *component {
	c.el.Style(styles.BorderRightVioletDark(scale))
	return c
}

func (c *component) BorderRightVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightVioletDarkAlpha(scale))
	return c
}

func (c *component) BorderRightWhite() *component {
	c.el.Style(styles.BorderRightWhite())
	return c
}

func (c *component) BorderRightWhiteAlpha(scale int) *component {
	c.el.Style(styles.BorderRightWhiteAlpha(scale))
	return c
}

func (c *component) BorderRightYellow(scale int) *component {
	c.el.Style(styles.BorderRightYellow(scale))
	return c
}

func (c *component) BorderRightYellowAlpha(scale int) *component {
	c.el.Style(styles.BorderRightYellowAlpha(scale))
	return c
}

func (c *component) BorderRightYellowDark(scale int) *component {
	c.el.Style(styles.BorderRightYellowDark(scale))
	return c
}

func (c *component) BorderRightYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRightYellowDarkAlpha(scale))
	return c
}

func (c *component) BorderRuby(scale int) *component {
	c.el.Style(styles.BorderRuby(scale))
	return c
}

func (c *component) BorderRubyAlpha(scale int) *component {
	c.el.Style(styles.BorderRubyAlpha(scale))
	return c
}

func (c *component) BorderRubyDark(scale int) *component {
	c.el.Style(styles.BorderRubyDark(scale))
	return c
}

func (c *component) BorderRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderRubyDarkAlpha(scale))
	return c
}

func (c *component) BorderS(val ...value.Value) *component {
	c.el.Style(styles.BorderS(val...))
	return c
}

func (c *component) BorderSage(scale int) *component {
	c.el.Style(styles.BorderSage(scale))
	return c
}

func (c *component) BorderSageAlpha(scale int) *component {
	c.el.Style(styles.BorderSageAlpha(scale))
	return c
}

func (c *component) BorderSageDark(scale int) *component {
	c.el.Style(styles.BorderSageDark(scale))
	return c
}

func (c *component) BorderSageDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderSageDarkAlpha(scale))
	return c
}

func (c *component) BorderSand(scale int) *component {
	c.el.Style(styles.BorderSand(scale))
	return c
}

func (c *component) BorderSandAlpha(scale int) *component {
	c.el.Style(styles.BorderSandAlpha(scale))
	return c
}

func (c *component) BorderSandDark(scale int) *component {
	c.el.Style(styles.BorderSandDark(scale))
	return c
}

func (c *component) BorderSandDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderSandDarkAlpha(scale))
	return c
}

func (c *component) BorderSeparate() *component {
	c.el.Style(styles.BorderSeparate())
	return c
}

func (c *component) BorderSky(scale int) *component {
	c.el.Style(styles.BorderSky(scale))
	return c
}

func (c *component) BorderSkyAlpha(scale int) *component {
	c.el.Style(styles.BorderSkyAlpha(scale))
	return c
}

func (c *component) BorderSkyDark(scale int) *component {
	c.el.Style(styles.BorderSkyDark(scale))
	return c
}

func (c *component) BorderSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderSkyDarkAlpha(scale))
	return c
}

func (c *component) BorderSlate(scale int) *component {
	c.el.Style(styles.BorderSlate(scale))
	return c
}

func (c *component) BorderSlateAlpha(scale int) *component {
	c.el.Style(styles.BorderSlateAlpha(scale))
	return c
}

func (c *component) BorderSlateDark(scale int) *component {
	c.el.Style(styles.BorderSlateDark(scale))
	return c
}

func (c *component) BorderSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderSlateDarkAlpha(scale))
	return c
}

func (c *component) BorderSolid() *component {
	c.el.Style(styles.BorderSolid())
	return c
}

func (c *component) BorderSpacing(spacing any) *component {
	c.el.Style(styles.BorderSpacing(spacing))
	return c
}

func (c *component) BorderSpacingX(spacing any) *component {
	c.el.Style(styles.BorderSpacingX(spacing))
	return c
}

func (c *component) BorderSpacingY(spacing any) *component {
	c.el.Style(styles.BorderSpacingY(spacing))
	return c
}

func (c *component) BorderStartAmber(scale int) *component {
	c.el.Style(styles.BorderStartAmber(scale))
	return c
}

func (c *component) BorderStartAmberAlpha(scale int) *component {
	c.el.Style(styles.BorderStartAmberAlpha(scale))
	return c
}

func (c *component) BorderStartAmberDark(scale int) *component {
	c.el.Style(styles.BorderStartAmberDark(scale))
	return c
}

func (c *component) BorderStartAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartAmberDarkAlpha(scale))
	return c
}

func (c *component) BorderStartBlack() *component {
	c.el.Style(styles.BorderStartBlack())
	return c
}

func (c *component) BorderStartBlackAlpha(scale int) *component {
	c.el.Style(styles.BorderStartBlackAlpha(scale))
	return c
}

func (c *component) BorderStartBlue(scale int) *component {
	c.el.Style(styles.BorderStartBlue(scale))
	return c
}

func (c *component) BorderStartBlueAlpha(scale int) *component {
	c.el.Style(styles.BorderStartBlueAlpha(scale))
	return c
}

func (c *component) BorderStartBlueDark(scale int) *component {
	c.el.Style(styles.BorderStartBlueDark(scale))
	return c
}

func (c *component) BorderStartBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartBlueDarkAlpha(scale))
	return c
}

func (c *component) BorderStartBronze(scale int) *component {
	c.el.Style(styles.BorderStartBronze(scale))
	return c
}

func (c *component) BorderStartBronzeAlpha(scale int) *component {
	c.el.Style(styles.BorderStartBronzeAlpha(scale))
	return c
}

func (c *component) BorderStartBronzeDark(scale int) *component {
	c.el.Style(styles.BorderStartBronzeDark(scale))
	return c
}

func (c *component) BorderStartBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartBronzeDarkAlpha(scale))
	return c
}

func (c *component) BorderStartBrown(scale int) *component {
	c.el.Style(styles.BorderStartBrown(scale))
	return c
}

func (c *component) BorderStartBrownAlpha(scale int) *component {
	c.el.Style(styles.BorderStartBrownAlpha(scale))
	return c
}

func (c *component) BorderStartBrownDark(scale int) *component {
	c.el.Style(styles.BorderStartBrownDark(scale))
	return c
}

func (c *component) BorderStartBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartBrownDarkAlpha(scale))
	return c
}

func (c *component) BorderStartColor(val value.Value) *component {
	c.el.Style(styles.BorderStartColor(val))
	return c
}

func (c *component) BorderStartCrimson(scale int) *component {
	c.el.Style(styles.BorderStartCrimson(scale))
	return c
}

func (c *component) BorderStartCrimsonAlpha(scale int) *component {
	c.el.Style(styles.BorderStartCrimsonAlpha(scale))
	return c
}

func (c *component) BorderStartCrimsonDark(scale int) *component {
	c.el.Style(styles.BorderStartCrimsonDark(scale))
	return c
}

func (c *component) BorderStartCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartCrimsonDarkAlpha(scale))
	return c
}

func (c *component) BorderStartCurrent() *component {
	c.el.Style(styles.BorderStartCurrent())
	return c
}

func (c *component) BorderStartCyan(scale int) *component {
	c.el.Style(styles.BorderStartCyan(scale))
	return c
}

func (c *component) BorderStartCyanAlpha(scale int) *component {
	c.el.Style(styles.BorderStartCyanAlpha(scale))
	return c
}

func (c *component) BorderStartCyanDark(scale int) *component {
	c.el.Style(styles.BorderStartCyanDark(scale))
	return c
}

func (c *component) BorderStartCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartCyanDarkAlpha(scale))
	return c
}

func (c *component) BorderStartGold(scale int) *component {
	c.el.Style(styles.BorderStartGold(scale))
	return c
}

func (c *component) BorderStartGoldAlpha(scale int) *component {
	c.el.Style(styles.BorderStartGoldAlpha(scale))
	return c
}

func (c *component) BorderStartGoldDark(scale int) *component {
	c.el.Style(styles.BorderStartGoldDark(scale))
	return c
}

func (c *component) BorderStartGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartGoldDarkAlpha(scale))
	return c
}

func (c *component) BorderStartGrass(scale int) *component {
	c.el.Style(styles.BorderStartGrass(scale))
	return c
}

func (c *component) BorderStartGrassAlpha(scale int) *component {
	c.el.Style(styles.BorderStartGrassAlpha(scale))
	return c
}

func (c *component) BorderStartGrassDark(scale int) *component {
	c.el.Style(styles.BorderStartGrassDark(scale))
	return c
}

func (c *component) BorderStartGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartGrassDarkAlpha(scale))
	return c
}

func (c *component) BorderStartGray(scale int) *component {
	c.el.Style(styles.BorderStartGray(scale))
	return c
}

func (c *component) BorderStartGrayAlpha(scale int) *component {
	c.el.Style(styles.BorderStartGrayAlpha(scale))
	return c
}

func (c *component) BorderStartGrayDark(scale int) *component {
	c.el.Style(styles.BorderStartGrayDark(scale))
	return c
}

func (c *component) BorderStartGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartGrayDarkAlpha(scale))
	return c
}

func (c *component) BorderStartGreen(scale int) *component {
	c.el.Style(styles.BorderStartGreen(scale))
	return c
}

func (c *component) BorderStartGreenAlpha(scale int) *component {
	c.el.Style(styles.BorderStartGreenAlpha(scale))
	return c
}

func (c *component) BorderStartGreenDark(scale int) *component {
	c.el.Style(styles.BorderStartGreenDark(scale))
	return c
}

func (c *component) BorderStartGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartGreenDarkAlpha(scale))
	return c
}

func (c *component) BorderStartIndigo(scale int) *component {
	c.el.Style(styles.BorderStartIndigo(scale))
	return c
}

func (c *component) BorderStartIndigoAlpha(scale int) *component {
	c.el.Style(styles.BorderStartIndigoAlpha(scale))
	return c
}

func (c *component) BorderStartIndigoDark(scale int) *component {
	c.el.Style(styles.BorderStartIndigoDark(scale))
	return c
}

func (c *component) BorderStartIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartIndigoDarkAlpha(scale))
	return c
}

func (c *component) BorderStartInherit() *component {
	c.el.Style(styles.BorderStartInherit())
	return c
}

func (c *component) BorderStartIris(scale int) *component {
	c.el.Style(styles.BorderStartIris(scale))
	return c
}

func (c *component) BorderStartIrisAlpha(scale int) *component {
	c.el.Style(styles.BorderStartIrisAlpha(scale))
	return c
}

func (c *component) BorderStartIrisDark(scale int) *component {
	c.el.Style(styles.BorderStartIrisDark(scale))
	return c
}

func (c *component) BorderStartIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartIrisDarkAlpha(scale))
	return c
}

func (c *component) BorderStartJade(scale int) *component {
	c.el.Style(styles.BorderStartJade(scale))
	return c
}

func (c *component) BorderStartJadeAlpha(scale int) *component {
	c.el.Style(styles.BorderStartJadeAlpha(scale))
	return c
}

func (c *component) BorderStartJadeDark(scale int) *component {
	c.el.Style(styles.BorderStartJadeDark(scale))
	return c
}

func (c *component) BorderStartJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartJadeDarkAlpha(scale))
	return c
}

func (c *component) BorderStartLime(scale int) *component {
	c.el.Style(styles.BorderStartLime(scale))
	return c
}

func (c *component) BorderStartLimeAlpha(scale int) *component {
	c.el.Style(styles.BorderStartLimeAlpha(scale))
	return c
}

func (c *component) BorderStartLimeDark(scale int) *component {
	c.el.Style(styles.BorderStartLimeDark(scale))
	return c
}

func (c *component) BorderStartLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartLimeDarkAlpha(scale))
	return c
}

func (c *component) BorderStartMauve(scale int) *component {
	c.el.Style(styles.BorderStartMauve(scale))
	return c
}

func (c *component) BorderStartMauveAlpha(scale int) *component {
	c.el.Style(styles.BorderStartMauveAlpha(scale))
	return c
}

func (c *component) BorderStartMauveDark(scale int) *component {
	c.el.Style(styles.BorderStartMauveDark(scale))
	return c
}

func (c *component) BorderStartMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartMauveDarkAlpha(scale))
	return c
}

func (c *component) BorderStartMint(scale int) *component {
	c.el.Style(styles.BorderStartMint(scale))
	return c
}

func (c *component) BorderStartMintAlpha(scale int) *component {
	c.el.Style(styles.BorderStartMintAlpha(scale))
	return c
}

func (c *component) BorderStartMintDark(scale int) *component {
	c.el.Style(styles.BorderStartMintDark(scale))
	return c
}

func (c *component) BorderStartMintDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartMintDarkAlpha(scale))
	return c
}

func (c *component) BorderStartOlive(scale int) *component {
	c.el.Style(styles.BorderStartOlive(scale))
	return c
}

func (c *component) BorderStartOliveAlpha(scale int) *component {
	c.el.Style(styles.BorderStartOliveAlpha(scale))
	return c
}

func (c *component) BorderStartOliveDark(scale int) *component {
	c.el.Style(styles.BorderStartOliveDark(scale))
	return c
}

func (c *component) BorderStartOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartOliveDarkAlpha(scale))
	return c
}

func (c *component) BorderStartOrange(scale int) *component {
	c.el.Style(styles.BorderStartOrange(scale))
	return c
}

func (c *component) BorderStartOrangeAlpha(scale int) *component {
	c.el.Style(styles.BorderStartOrangeAlpha(scale))
	return c
}

func (c *component) BorderStartOrangeDark(scale int) *component {
	c.el.Style(styles.BorderStartOrangeDark(scale))
	return c
}

func (c *component) BorderStartOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartOrangeDarkAlpha(scale))
	return c
}

func (c *component) BorderStartPink(scale int) *component {
	c.el.Style(styles.BorderStartPink(scale))
	return c
}

func (c *component) BorderStartPinkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartPinkAlpha(scale))
	return c
}

func (c *component) BorderStartPinkDark(scale int) *component {
	c.el.Style(styles.BorderStartPinkDark(scale))
	return c
}

func (c *component) BorderStartPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartPinkDarkAlpha(scale))
	return c
}

func (c *component) BorderStartPlum(scale int) *component {
	c.el.Style(styles.BorderStartPlum(scale))
	return c
}

func (c *component) BorderStartPlumAlpha(scale int) *component {
	c.el.Style(styles.BorderStartPlumAlpha(scale))
	return c
}

func (c *component) BorderStartPlumDark(scale int) *component {
	c.el.Style(styles.BorderStartPlumDark(scale))
	return c
}

func (c *component) BorderStartPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartPlumDarkAlpha(scale))
	return c
}

func (c *component) BorderStartPurple(scale int) *component {
	c.el.Style(styles.BorderStartPurple(scale))
	return c
}

func (c *component) BorderStartPurpleAlpha(scale int) *component {
	c.el.Style(styles.BorderStartPurpleAlpha(scale))
	return c
}

func (c *component) BorderStartPurpleDark(scale int) *component {
	c.el.Style(styles.BorderStartPurpleDark(scale))
	return c
}

func (c *component) BorderStartPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartPurpleDarkAlpha(scale))
	return c
}

func (c *component) BorderStartRed(scale int) *component {
	c.el.Style(styles.BorderStartRed(scale))
	return c
}

func (c *component) BorderStartRedAlpha(scale int) *component {
	c.el.Style(styles.BorderStartRedAlpha(scale))
	return c
}

func (c *component) BorderStartRedDark(scale int) *component {
	c.el.Style(styles.BorderStartRedDark(scale))
	return c
}

func (c *component) BorderStartRedDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartRedDarkAlpha(scale))
	return c
}

func (c *component) BorderStartRuby(scale int) *component {
	c.el.Style(styles.BorderStartRuby(scale))
	return c
}

func (c *component) BorderStartRubyAlpha(scale int) *component {
	c.el.Style(styles.BorderStartRubyAlpha(scale))
	return c
}

func (c *component) BorderStartRubyDark(scale int) *component {
	c.el.Style(styles.BorderStartRubyDark(scale))
	return c
}

func (c *component) BorderStartRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartRubyDarkAlpha(scale))
	return c
}

func (c *component) BorderStartSage(scale int) *component {
	c.el.Style(styles.BorderStartSage(scale))
	return c
}

func (c *component) BorderStartSageAlpha(scale int) *component {
	c.el.Style(styles.BorderStartSageAlpha(scale))
	return c
}

func (c *component) BorderStartSageDark(scale int) *component {
	c.el.Style(styles.BorderStartSageDark(scale))
	return c
}

func (c *component) BorderStartSageDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartSageDarkAlpha(scale))
	return c
}

func (c *component) BorderStartSand(scale int) *component {
	c.el.Style(styles.BorderStartSand(scale))
	return c
}

func (c *component) BorderStartSandAlpha(scale int) *component {
	c.el.Style(styles.BorderStartSandAlpha(scale))
	return c
}

func (c *component) BorderStartSandDark(scale int) *component {
	c.el.Style(styles.BorderStartSandDark(scale))
	return c
}

func (c *component) BorderStartSandDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartSandDarkAlpha(scale))
	return c
}

func (c *component) BorderStartSky(scale int) *component {
	c.el.Style(styles.BorderStartSky(scale))
	return c
}

func (c *component) BorderStartSkyAlpha(scale int) *component {
	c.el.Style(styles.BorderStartSkyAlpha(scale))
	return c
}

func (c *component) BorderStartSkyDark(scale int) *component {
	c.el.Style(styles.BorderStartSkyDark(scale))
	return c
}

func (c *component) BorderStartSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartSkyDarkAlpha(scale))
	return c
}

func (c *component) BorderStartSlate(scale int) *component {
	c.el.Style(styles.BorderStartSlate(scale))
	return c
}

func (c *component) BorderStartSlateAlpha(scale int) *component {
	c.el.Style(styles.BorderStartSlateAlpha(scale))
	return c
}

func (c *component) BorderStartSlateDark(scale int) *component {
	c.el.Style(styles.BorderStartSlateDark(scale))
	return c
}

func (c *component) BorderStartSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartSlateDarkAlpha(scale))
	return c
}

func (c *component) BorderStartTeal(scale int) *component {
	c.el.Style(styles.BorderStartTeal(scale))
	return c
}

func (c *component) BorderStartTealAlpha(scale int) *component {
	c.el.Style(styles.BorderStartTealAlpha(scale))
	return c
}

func (c *component) BorderStartTealDark(scale int) *component {
	c.el.Style(styles.BorderStartTealDark(scale))
	return c
}

func (c *component) BorderStartTealDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartTealDarkAlpha(scale))
	return c
}

func (c *component) BorderStartTomato(scale int) *component {
	c.el.Style(styles.BorderStartTomato(scale))
	return c
}

func (c *component) BorderStartTomatoAlpha(scale int) *component {
	c.el.Style(styles.BorderStartTomatoAlpha(scale))
	return c
}

func (c *component) BorderStartTomatoDark(scale int) *component {
	c.el.Style(styles.BorderStartTomatoDark(scale))
	return c
}

func (c *component) BorderStartTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartTomatoDarkAlpha(scale))
	return c
}

func (c *component) BorderStartTransparent() *component {
	c.el.Style(styles.BorderStartTransparent())
	return c
}

func (c *component) BorderStartViolet(scale int) *component {
	c.el.Style(styles.BorderStartViolet(scale))
	return c
}

func (c *component) BorderStartVioletAlpha(scale int) *component {
	c.el.Style(styles.BorderStartVioletAlpha(scale))
	return c
}

func (c *component) BorderStartVioletDark(scale int) *component {
	c.el.Style(styles.BorderStartVioletDark(scale))
	return c
}

func (c *component) BorderStartVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartVioletDarkAlpha(scale))
	return c
}

func (c *component) BorderStartWhite() *component {
	c.el.Style(styles.BorderStartWhite())
	return c
}

func (c *component) BorderStartWhiteAlpha(scale int) *component {
	c.el.Style(styles.BorderStartWhiteAlpha(scale))
	return c
}

func (c *component) BorderStartYellow(scale int) *component {
	c.el.Style(styles.BorderStartYellow(scale))
	return c
}

func (c *component) BorderStartYellowAlpha(scale int) *component {
	c.el.Style(styles.BorderStartYellowAlpha(scale))
	return c
}

func (c *component) BorderStartYellowDark(scale int) *component {
	c.el.Style(styles.BorderStartYellowDark(scale))
	return c
}

func (c *component) BorderStartYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderStartYellowDarkAlpha(scale))
	return c
}

func (c *component) BorderT(val ...value.Value) *component {
	c.el.Style(styles.BorderT(val...))
	return c
}

func (c *component) BorderTeal(scale int) *component {
	c.el.Style(styles.BorderTeal(scale))
	return c
}

func (c *component) BorderTealAlpha(scale int) *component {
	c.el.Style(styles.BorderTealAlpha(scale))
	return c
}

func (c *component) BorderTealDark(scale int) *component {
	c.el.Style(styles.BorderTealDark(scale))
	return c
}

func (c *component) BorderTealDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTealDarkAlpha(scale))
	return c
}

func (c *component) BorderTomato(scale int) *component {
	c.el.Style(styles.BorderTomato(scale))
	return c
}

func (c *component) BorderTomatoAlpha(scale int) *component {
	c.el.Style(styles.BorderTomatoAlpha(scale))
	return c
}

func (c *component) BorderTomatoDark(scale int) *component {
	c.el.Style(styles.BorderTomatoDark(scale))
	return c
}

func (c *component) BorderTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTomatoDarkAlpha(scale))
	return c
}

func (c *component) BorderTopAmber(scale int) *component {
	c.el.Style(styles.BorderTopAmber(scale))
	return c
}

func (c *component) BorderTopAmberAlpha(scale int) *component {
	c.el.Style(styles.BorderTopAmberAlpha(scale))
	return c
}

func (c *component) BorderTopAmberDark(scale int) *component {
	c.el.Style(styles.BorderTopAmberDark(scale))
	return c
}

func (c *component) BorderTopAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopAmberDarkAlpha(scale))
	return c
}

func (c *component) BorderTopBlack() *component {
	c.el.Style(styles.BorderTopBlack())
	return c
}

func (c *component) BorderTopBlackAlpha(scale int) *component {
	c.el.Style(styles.BorderTopBlackAlpha(scale))
	return c
}

func (c *component) BorderTopBlue(scale int) *component {
	c.el.Style(styles.BorderTopBlue(scale))
	return c
}

func (c *component) BorderTopBlueAlpha(scale int) *component {
	c.el.Style(styles.BorderTopBlueAlpha(scale))
	return c
}

func (c *component) BorderTopBlueDark(scale int) *component {
	c.el.Style(styles.BorderTopBlueDark(scale))
	return c
}

func (c *component) BorderTopBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopBlueDarkAlpha(scale))
	return c
}

func (c *component) BorderTopBronze(scale int) *component {
	c.el.Style(styles.BorderTopBronze(scale))
	return c
}

func (c *component) BorderTopBronzeAlpha(scale int) *component {
	c.el.Style(styles.BorderTopBronzeAlpha(scale))
	return c
}

func (c *component) BorderTopBronzeDark(scale int) *component {
	c.el.Style(styles.BorderTopBronzeDark(scale))
	return c
}

func (c *component) BorderTopBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopBronzeDarkAlpha(scale))
	return c
}

func (c *component) BorderTopBrown(scale int) *component {
	c.el.Style(styles.BorderTopBrown(scale))
	return c
}

func (c *component) BorderTopBrownAlpha(scale int) *component {
	c.el.Style(styles.BorderTopBrownAlpha(scale))
	return c
}

func (c *component) BorderTopBrownDark(scale int) *component {
	c.el.Style(styles.BorderTopBrownDark(scale))
	return c
}

func (c *component) BorderTopBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopBrownDarkAlpha(scale))
	return c
}

func (c *component) BorderTopColor(val value.Value) *component {
	c.el.Style(styles.BorderTopColor(val))
	return c
}

func (c *component) BorderTopCrimson(scale int) *component {
	c.el.Style(styles.BorderTopCrimson(scale))
	return c
}

func (c *component) BorderTopCrimsonAlpha(scale int) *component {
	c.el.Style(styles.BorderTopCrimsonAlpha(scale))
	return c
}

func (c *component) BorderTopCrimsonDark(scale int) *component {
	c.el.Style(styles.BorderTopCrimsonDark(scale))
	return c
}

func (c *component) BorderTopCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopCrimsonDarkAlpha(scale))
	return c
}

func (c *component) BorderTopCurrent() *component {
	c.el.Style(styles.BorderTopCurrent())
	return c
}

func (c *component) BorderTopCyan(scale int) *component {
	c.el.Style(styles.BorderTopCyan(scale))
	return c
}

func (c *component) BorderTopCyanAlpha(scale int) *component {
	c.el.Style(styles.BorderTopCyanAlpha(scale))
	return c
}

func (c *component) BorderTopCyanDark(scale int) *component {
	c.el.Style(styles.BorderTopCyanDark(scale))
	return c
}

func (c *component) BorderTopCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopCyanDarkAlpha(scale))
	return c
}

func (c *component) BorderTopGold(scale int) *component {
	c.el.Style(styles.BorderTopGold(scale))
	return c
}

func (c *component) BorderTopGoldAlpha(scale int) *component {
	c.el.Style(styles.BorderTopGoldAlpha(scale))
	return c
}

func (c *component) BorderTopGoldDark(scale int) *component {
	c.el.Style(styles.BorderTopGoldDark(scale))
	return c
}

func (c *component) BorderTopGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopGoldDarkAlpha(scale))
	return c
}

func (c *component) BorderTopGrass(scale int) *component {
	c.el.Style(styles.BorderTopGrass(scale))
	return c
}

func (c *component) BorderTopGrassAlpha(scale int) *component {
	c.el.Style(styles.BorderTopGrassAlpha(scale))
	return c
}

func (c *component) BorderTopGrassDark(scale int) *component {
	c.el.Style(styles.BorderTopGrassDark(scale))
	return c
}

func (c *component) BorderTopGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopGrassDarkAlpha(scale))
	return c
}

func (c *component) BorderTopGray(scale int) *component {
	c.el.Style(styles.BorderTopGray(scale))
	return c
}

func (c *component) BorderTopGrayAlpha(scale int) *component {
	c.el.Style(styles.BorderTopGrayAlpha(scale))
	return c
}

func (c *component) BorderTopGrayDark(scale int) *component {
	c.el.Style(styles.BorderTopGrayDark(scale))
	return c
}

func (c *component) BorderTopGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopGrayDarkAlpha(scale))
	return c
}

func (c *component) BorderTopGreen(scale int) *component {
	c.el.Style(styles.BorderTopGreen(scale))
	return c
}

func (c *component) BorderTopGreenAlpha(scale int) *component {
	c.el.Style(styles.BorderTopGreenAlpha(scale))
	return c
}

func (c *component) BorderTopGreenDark(scale int) *component {
	c.el.Style(styles.BorderTopGreenDark(scale))
	return c
}

func (c *component) BorderTopGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopGreenDarkAlpha(scale))
	return c
}

func (c *component) BorderTopIndigo(scale int) *component {
	c.el.Style(styles.BorderTopIndigo(scale))
	return c
}

func (c *component) BorderTopIndigoAlpha(scale int) *component {
	c.el.Style(styles.BorderTopIndigoAlpha(scale))
	return c
}

func (c *component) BorderTopIndigoDark(scale int) *component {
	c.el.Style(styles.BorderTopIndigoDark(scale))
	return c
}

func (c *component) BorderTopIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopIndigoDarkAlpha(scale))
	return c
}

func (c *component) BorderTopInherit() *component {
	c.el.Style(styles.BorderTopInherit())
	return c
}

func (c *component) BorderTopIris(scale int) *component {
	c.el.Style(styles.BorderTopIris(scale))
	return c
}

func (c *component) BorderTopIrisAlpha(scale int) *component {
	c.el.Style(styles.BorderTopIrisAlpha(scale))
	return c
}

func (c *component) BorderTopIrisDark(scale int) *component {
	c.el.Style(styles.BorderTopIrisDark(scale))
	return c
}

func (c *component) BorderTopIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopIrisDarkAlpha(scale))
	return c
}

func (c *component) BorderTopJade(scale int) *component {
	c.el.Style(styles.BorderTopJade(scale))
	return c
}

func (c *component) BorderTopJadeAlpha(scale int) *component {
	c.el.Style(styles.BorderTopJadeAlpha(scale))
	return c
}

func (c *component) BorderTopJadeDark(scale int) *component {
	c.el.Style(styles.BorderTopJadeDark(scale))
	return c
}

func (c *component) BorderTopJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopJadeDarkAlpha(scale))
	return c
}

func (c *component) BorderTopLime(scale int) *component {
	c.el.Style(styles.BorderTopLime(scale))
	return c
}

func (c *component) BorderTopLimeAlpha(scale int) *component {
	c.el.Style(styles.BorderTopLimeAlpha(scale))
	return c
}

func (c *component) BorderTopLimeDark(scale int) *component {
	c.el.Style(styles.BorderTopLimeDark(scale))
	return c
}

func (c *component) BorderTopLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopLimeDarkAlpha(scale))
	return c
}

func (c *component) BorderTopMauve(scale int) *component {
	c.el.Style(styles.BorderTopMauve(scale))
	return c
}

func (c *component) BorderTopMauveAlpha(scale int) *component {
	c.el.Style(styles.BorderTopMauveAlpha(scale))
	return c
}

func (c *component) BorderTopMauveDark(scale int) *component {
	c.el.Style(styles.BorderTopMauveDark(scale))
	return c
}

func (c *component) BorderTopMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopMauveDarkAlpha(scale))
	return c
}

func (c *component) BorderTopMint(scale int) *component {
	c.el.Style(styles.BorderTopMint(scale))
	return c
}

func (c *component) BorderTopMintAlpha(scale int) *component {
	c.el.Style(styles.BorderTopMintAlpha(scale))
	return c
}

func (c *component) BorderTopMintDark(scale int) *component {
	c.el.Style(styles.BorderTopMintDark(scale))
	return c
}

func (c *component) BorderTopMintDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopMintDarkAlpha(scale))
	return c
}

func (c *component) BorderTopOlive(scale int) *component {
	c.el.Style(styles.BorderTopOlive(scale))
	return c
}

func (c *component) BorderTopOliveAlpha(scale int) *component {
	c.el.Style(styles.BorderTopOliveAlpha(scale))
	return c
}

func (c *component) BorderTopOliveDark(scale int) *component {
	c.el.Style(styles.BorderTopOliveDark(scale))
	return c
}

func (c *component) BorderTopOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopOliveDarkAlpha(scale))
	return c
}

func (c *component) BorderTopOrange(scale int) *component {
	c.el.Style(styles.BorderTopOrange(scale))
	return c
}

func (c *component) BorderTopOrangeAlpha(scale int) *component {
	c.el.Style(styles.BorderTopOrangeAlpha(scale))
	return c
}

func (c *component) BorderTopOrangeDark(scale int) *component {
	c.el.Style(styles.BorderTopOrangeDark(scale))
	return c
}

func (c *component) BorderTopOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopOrangeDarkAlpha(scale))
	return c
}

func (c *component) BorderTopPink(scale int) *component {
	c.el.Style(styles.BorderTopPink(scale))
	return c
}

func (c *component) BorderTopPinkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopPinkAlpha(scale))
	return c
}

func (c *component) BorderTopPinkDark(scale int) *component {
	c.el.Style(styles.BorderTopPinkDark(scale))
	return c
}

func (c *component) BorderTopPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopPinkDarkAlpha(scale))
	return c
}

func (c *component) BorderTopPlum(scale int) *component {
	c.el.Style(styles.BorderTopPlum(scale))
	return c
}

func (c *component) BorderTopPlumAlpha(scale int) *component {
	c.el.Style(styles.BorderTopPlumAlpha(scale))
	return c
}

func (c *component) BorderTopPlumDark(scale int) *component {
	c.el.Style(styles.BorderTopPlumDark(scale))
	return c
}

func (c *component) BorderTopPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopPlumDarkAlpha(scale))
	return c
}

func (c *component) BorderTopPurple(scale int) *component {
	c.el.Style(styles.BorderTopPurple(scale))
	return c
}

func (c *component) BorderTopPurpleAlpha(scale int) *component {
	c.el.Style(styles.BorderTopPurpleAlpha(scale))
	return c
}

func (c *component) BorderTopPurpleDark(scale int) *component {
	c.el.Style(styles.BorderTopPurpleDark(scale))
	return c
}

func (c *component) BorderTopPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopPurpleDarkAlpha(scale))
	return c
}

func (c *component) BorderTopRed(scale int) *component {
	c.el.Style(styles.BorderTopRed(scale))
	return c
}

func (c *component) BorderTopRedAlpha(scale int) *component {
	c.el.Style(styles.BorderTopRedAlpha(scale))
	return c
}

func (c *component) BorderTopRedDark(scale int) *component {
	c.el.Style(styles.BorderTopRedDark(scale))
	return c
}

func (c *component) BorderTopRedDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopRedDarkAlpha(scale))
	return c
}

func (c *component) BorderTopRuby(scale int) *component {
	c.el.Style(styles.BorderTopRuby(scale))
	return c
}

func (c *component) BorderTopRubyAlpha(scale int) *component {
	c.el.Style(styles.BorderTopRubyAlpha(scale))
	return c
}

func (c *component) BorderTopRubyDark(scale int) *component {
	c.el.Style(styles.BorderTopRubyDark(scale))
	return c
}

func (c *component) BorderTopRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopRubyDarkAlpha(scale))
	return c
}

func (c *component) BorderTopSage(scale int) *component {
	c.el.Style(styles.BorderTopSage(scale))
	return c
}

func (c *component) BorderTopSageAlpha(scale int) *component {
	c.el.Style(styles.BorderTopSageAlpha(scale))
	return c
}

func (c *component) BorderTopSageDark(scale int) *component {
	c.el.Style(styles.BorderTopSageDark(scale))
	return c
}

func (c *component) BorderTopSageDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopSageDarkAlpha(scale))
	return c
}

func (c *component) BorderTopSand(scale int) *component {
	c.el.Style(styles.BorderTopSand(scale))
	return c
}

func (c *component) BorderTopSandAlpha(scale int) *component {
	c.el.Style(styles.BorderTopSandAlpha(scale))
	return c
}

func (c *component) BorderTopSandDark(scale int) *component {
	c.el.Style(styles.BorderTopSandDark(scale))
	return c
}

func (c *component) BorderTopSandDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopSandDarkAlpha(scale))
	return c
}

func (c *component) BorderTopSky(scale int) *component {
	c.el.Style(styles.BorderTopSky(scale))
	return c
}

func (c *component) BorderTopSkyAlpha(scale int) *component {
	c.el.Style(styles.BorderTopSkyAlpha(scale))
	return c
}

func (c *component) BorderTopSkyDark(scale int) *component {
	c.el.Style(styles.BorderTopSkyDark(scale))
	return c
}

func (c *component) BorderTopSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopSkyDarkAlpha(scale))
	return c
}

func (c *component) BorderTopSlate(scale int) *component {
	c.el.Style(styles.BorderTopSlate(scale))
	return c
}

func (c *component) BorderTopSlateAlpha(scale int) *component {
	c.el.Style(styles.BorderTopSlateAlpha(scale))
	return c
}

func (c *component) BorderTopSlateDark(scale int) *component {
	c.el.Style(styles.BorderTopSlateDark(scale))
	return c
}

func (c *component) BorderTopSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopSlateDarkAlpha(scale))
	return c
}

func (c *component) BorderTopTeal(scale int) *component {
	c.el.Style(styles.BorderTopTeal(scale))
	return c
}

func (c *component) BorderTopTealAlpha(scale int) *component {
	c.el.Style(styles.BorderTopTealAlpha(scale))
	return c
}

func (c *component) BorderTopTealDark(scale int) *component {
	c.el.Style(styles.BorderTopTealDark(scale))
	return c
}

func (c *component) BorderTopTealDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopTealDarkAlpha(scale))
	return c
}

func (c *component) BorderTopTomato(scale int) *component {
	c.el.Style(styles.BorderTopTomato(scale))
	return c
}

func (c *component) BorderTopTomatoAlpha(scale int) *component {
	c.el.Style(styles.BorderTopTomatoAlpha(scale))
	return c
}

func (c *component) BorderTopTomatoDark(scale int) *component {
	c.el.Style(styles.BorderTopTomatoDark(scale))
	return c
}

func (c *component) BorderTopTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopTomatoDarkAlpha(scale))
	return c
}

func (c *component) BorderTopTransparent() *component {
	c.el.Style(styles.BorderTopTransparent())
	return c
}

func (c *component) BorderTopViolet(scale int) *component {
	c.el.Style(styles.BorderTopViolet(scale))
	return c
}

func (c *component) BorderTopVioletAlpha(scale int) *component {
	c.el.Style(styles.BorderTopVioletAlpha(scale))
	return c
}

func (c *component) BorderTopVioletDark(scale int) *component {
	c.el.Style(styles.BorderTopVioletDark(scale))
	return c
}

func (c *component) BorderTopVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopVioletDarkAlpha(scale))
	return c
}

func (c *component) BorderTopWhite() *component {
	c.el.Style(styles.BorderTopWhite())
	return c
}

func (c *component) BorderTopWhiteAlpha(scale int) *component {
	c.el.Style(styles.BorderTopWhiteAlpha(scale))
	return c
}

func (c *component) BorderTopYellow(scale int) *component {
	c.el.Style(styles.BorderTopYellow(scale))
	return c
}

func (c *component) BorderTopYellowAlpha(scale int) *component {
	c.el.Style(styles.BorderTopYellowAlpha(scale))
	return c
}

func (c *component) BorderTopYellowDark(scale int) *component {
	c.el.Style(styles.BorderTopYellowDark(scale))
	return c
}

func (c *component) BorderTopYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderTopYellowDarkAlpha(scale))
	return c
}

func (c *component) BorderTransparent() *component {
	c.el.Style(styles.BorderTransparent())
	return c
}

func (c *component) BorderViolet(scale int) *component {
	c.el.Style(styles.BorderViolet(scale))
	return c
}

func (c *component) BorderVioletAlpha(scale int) *component {
	c.el.Style(styles.BorderVioletAlpha(scale))
	return c
}

func (c *component) BorderVioletDark(scale int) *component {
	c.el.Style(styles.BorderVioletDark(scale))
	return c
}

func (c *component) BorderVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderVioletDarkAlpha(scale))
	return c
}

func (c *component) BorderWhite() *component {
	c.el.Style(styles.BorderWhite())
	return c
}

func (c *component) BorderWhiteAlpha(scale int) *component {
	c.el.Style(styles.BorderWhiteAlpha(scale))
	return c
}

func (c *component) BorderX(val ...value.Value) *component {
	c.el.Style(styles.BorderX(val...))
	return c
}

func (c *component) BorderXAmber(scale int) *component {
	c.el.Style(styles.BorderXAmber(scale))
	return c
}

func (c *component) BorderXAmberAlpha(scale int) *component {
	c.el.Style(styles.BorderXAmberAlpha(scale))
	return c
}

func (c *component) BorderXAmberDark(scale int) *component {
	c.el.Style(styles.BorderXAmberDark(scale))
	return c
}

func (c *component) BorderXAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXAmberDarkAlpha(scale))
	return c
}

func (c *component) BorderXBlack() *component {
	c.el.Style(styles.BorderXBlack())
	return c
}

func (c *component) BorderXBlackAlpha(scale int) *component {
	c.el.Style(styles.BorderXBlackAlpha(scale))
	return c
}

func (c *component) BorderXBlue(scale int) *component {
	c.el.Style(styles.BorderXBlue(scale))
	return c
}

func (c *component) BorderXBlueAlpha(scale int) *component {
	c.el.Style(styles.BorderXBlueAlpha(scale))
	return c
}

func (c *component) BorderXBlueDark(scale int) *component {
	c.el.Style(styles.BorderXBlueDark(scale))
	return c
}

func (c *component) BorderXBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXBlueDarkAlpha(scale))
	return c
}

func (c *component) BorderXBronze(scale int) *component {
	c.el.Style(styles.BorderXBronze(scale))
	return c
}

func (c *component) BorderXBronzeAlpha(scale int) *component {
	c.el.Style(styles.BorderXBronzeAlpha(scale))
	return c
}

func (c *component) BorderXBronzeDark(scale int) *component {
	c.el.Style(styles.BorderXBronzeDark(scale))
	return c
}

func (c *component) BorderXBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXBronzeDarkAlpha(scale))
	return c
}

func (c *component) BorderXBrown(scale int) *component {
	c.el.Style(styles.BorderXBrown(scale))
	return c
}

func (c *component) BorderXBrownAlpha(scale int) *component {
	c.el.Style(styles.BorderXBrownAlpha(scale))
	return c
}

func (c *component) BorderXBrownDark(scale int) *component {
	c.el.Style(styles.BorderXBrownDark(scale))
	return c
}

func (c *component) BorderXBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXBrownDarkAlpha(scale))
	return c
}

func (c *component) BorderXColor(val value.Value) *component {
	c.el.Style(styles.BorderXColor(val))
	return c
}

func (c *component) BorderXCrimson(scale int) *component {
	c.el.Style(styles.BorderXCrimson(scale))
	return c
}

func (c *component) BorderXCrimsonAlpha(scale int) *component {
	c.el.Style(styles.BorderXCrimsonAlpha(scale))
	return c
}

func (c *component) BorderXCrimsonDark(scale int) *component {
	c.el.Style(styles.BorderXCrimsonDark(scale))
	return c
}

func (c *component) BorderXCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXCrimsonDarkAlpha(scale))
	return c
}

func (c *component) BorderXCurrent() *component {
	c.el.Style(styles.BorderXCurrent())
	return c
}

func (c *component) BorderXCyan(scale int) *component {
	c.el.Style(styles.BorderXCyan(scale))
	return c
}

func (c *component) BorderXCyanAlpha(scale int) *component {
	c.el.Style(styles.BorderXCyanAlpha(scale))
	return c
}

func (c *component) BorderXCyanDark(scale int) *component {
	c.el.Style(styles.BorderXCyanDark(scale))
	return c
}

func (c *component) BorderXCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXCyanDarkAlpha(scale))
	return c
}

func (c *component) BorderXGold(scale int) *component {
	c.el.Style(styles.BorderXGold(scale))
	return c
}

func (c *component) BorderXGoldAlpha(scale int) *component {
	c.el.Style(styles.BorderXGoldAlpha(scale))
	return c
}

func (c *component) BorderXGoldDark(scale int) *component {
	c.el.Style(styles.BorderXGoldDark(scale))
	return c
}

func (c *component) BorderXGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXGoldDarkAlpha(scale))
	return c
}

func (c *component) BorderXGrass(scale int) *component {
	c.el.Style(styles.BorderXGrass(scale))
	return c
}

func (c *component) BorderXGrassAlpha(scale int) *component {
	c.el.Style(styles.BorderXGrassAlpha(scale))
	return c
}

func (c *component) BorderXGrassDark(scale int) *component {
	c.el.Style(styles.BorderXGrassDark(scale))
	return c
}

func (c *component) BorderXGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXGrassDarkAlpha(scale))
	return c
}

func (c *component) BorderXGray(scale int) *component {
	c.el.Style(styles.BorderXGray(scale))
	return c
}

func (c *component) BorderXGrayAlpha(scale int) *component {
	c.el.Style(styles.BorderXGrayAlpha(scale))
	return c
}

func (c *component) BorderXGrayDark(scale int) *component {
	c.el.Style(styles.BorderXGrayDark(scale))
	return c
}

func (c *component) BorderXGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXGrayDarkAlpha(scale))
	return c
}

func (c *component) BorderXGreen(scale int) *component {
	c.el.Style(styles.BorderXGreen(scale))
	return c
}

func (c *component) BorderXGreenAlpha(scale int) *component {
	c.el.Style(styles.BorderXGreenAlpha(scale))
	return c
}

func (c *component) BorderXGreenDark(scale int) *component {
	c.el.Style(styles.BorderXGreenDark(scale))
	return c
}

func (c *component) BorderXGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXGreenDarkAlpha(scale))
	return c
}

func (c *component) BorderXIndigo(scale int) *component {
	c.el.Style(styles.BorderXIndigo(scale))
	return c
}

func (c *component) BorderXIndigoAlpha(scale int) *component {
	c.el.Style(styles.BorderXIndigoAlpha(scale))
	return c
}

func (c *component) BorderXIndigoDark(scale int) *component {
	c.el.Style(styles.BorderXIndigoDark(scale))
	return c
}

func (c *component) BorderXIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXIndigoDarkAlpha(scale))
	return c
}

func (c *component) BorderXInherit() *component {
	c.el.Style(styles.BorderXInherit())
	return c
}

func (c *component) BorderXIris(scale int) *component {
	c.el.Style(styles.BorderXIris(scale))
	return c
}

func (c *component) BorderXIrisAlpha(scale int) *component {
	c.el.Style(styles.BorderXIrisAlpha(scale))
	return c
}

func (c *component) BorderXIrisDark(scale int) *component {
	c.el.Style(styles.BorderXIrisDark(scale))
	return c
}

func (c *component) BorderXIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXIrisDarkAlpha(scale))
	return c
}

func (c *component) BorderXJade(scale int) *component {
	c.el.Style(styles.BorderXJade(scale))
	return c
}

func (c *component) BorderXJadeAlpha(scale int) *component {
	c.el.Style(styles.BorderXJadeAlpha(scale))
	return c
}

func (c *component) BorderXJadeDark(scale int) *component {
	c.el.Style(styles.BorderXJadeDark(scale))
	return c
}

func (c *component) BorderXJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXJadeDarkAlpha(scale))
	return c
}

func (c *component) BorderXLime(scale int) *component {
	c.el.Style(styles.BorderXLime(scale))
	return c
}

func (c *component) BorderXLimeAlpha(scale int) *component {
	c.el.Style(styles.BorderXLimeAlpha(scale))
	return c
}

func (c *component) BorderXLimeDark(scale int) *component {
	c.el.Style(styles.BorderXLimeDark(scale))
	return c
}

func (c *component) BorderXLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXLimeDarkAlpha(scale))
	return c
}

func (c *component) BorderXMauve(scale int) *component {
	c.el.Style(styles.BorderXMauve(scale))
	return c
}

func (c *component) BorderXMauveAlpha(scale int) *component {
	c.el.Style(styles.BorderXMauveAlpha(scale))
	return c
}

func (c *component) BorderXMauveDark(scale int) *component {
	c.el.Style(styles.BorderXMauveDark(scale))
	return c
}

func (c *component) BorderXMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXMauveDarkAlpha(scale))
	return c
}

func (c *component) BorderXMint(scale int) *component {
	c.el.Style(styles.BorderXMint(scale))
	return c
}

func (c *component) BorderXMintAlpha(scale int) *component {
	c.el.Style(styles.BorderXMintAlpha(scale))
	return c
}

func (c *component) BorderXMintDark(scale int) *component {
	c.el.Style(styles.BorderXMintDark(scale))
	return c
}

func (c *component) BorderXMintDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXMintDarkAlpha(scale))
	return c
}

func (c *component) BorderXOlive(scale int) *component {
	c.el.Style(styles.BorderXOlive(scale))
	return c
}

func (c *component) BorderXOliveAlpha(scale int) *component {
	c.el.Style(styles.BorderXOliveAlpha(scale))
	return c
}

func (c *component) BorderXOliveDark(scale int) *component {
	c.el.Style(styles.BorderXOliveDark(scale))
	return c
}

func (c *component) BorderXOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXOliveDarkAlpha(scale))
	return c
}

func (c *component) BorderXOrange(scale int) *component {
	c.el.Style(styles.BorderXOrange(scale))
	return c
}

func (c *component) BorderXOrangeAlpha(scale int) *component {
	c.el.Style(styles.BorderXOrangeAlpha(scale))
	return c
}

func (c *component) BorderXOrangeDark(scale int) *component {
	c.el.Style(styles.BorderXOrangeDark(scale))
	return c
}

func (c *component) BorderXOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXOrangeDarkAlpha(scale))
	return c
}

func (c *component) BorderXPink(scale int) *component {
	c.el.Style(styles.BorderXPink(scale))
	return c
}

func (c *component) BorderXPinkAlpha(scale int) *component {
	c.el.Style(styles.BorderXPinkAlpha(scale))
	return c
}

func (c *component) BorderXPinkDark(scale int) *component {
	c.el.Style(styles.BorderXPinkDark(scale))
	return c
}

func (c *component) BorderXPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXPinkDarkAlpha(scale))
	return c
}

func (c *component) BorderXPlum(scale int) *component {
	c.el.Style(styles.BorderXPlum(scale))
	return c
}

func (c *component) BorderXPlumAlpha(scale int) *component {
	c.el.Style(styles.BorderXPlumAlpha(scale))
	return c
}

func (c *component) BorderXPlumDark(scale int) *component {
	c.el.Style(styles.BorderXPlumDark(scale))
	return c
}

func (c *component) BorderXPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXPlumDarkAlpha(scale))
	return c
}

func (c *component) BorderXPurple(scale int) *component {
	c.el.Style(styles.BorderXPurple(scale))
	return c
}

func (c *component) BorderXPurpleAlpha(scale int) *component {
	c.el.Style(styles.BorderXPurpleAlpha(scale))
	return c
}

func (c *component) BorderXPurpleDark(scale int) *component {
	c.el.Style(styles.BorderXPurpleDark(scale))
	return c
}

func (c *component) BorderXPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXPurpleDarkAlpha(scale))
	return c
}

func (c *component) BorderXRed(scale int) *component {
	c.el.Style(styles.BorderXRed(scale))
	return c
}

func (c *component) BorderXRedAlpha(scale int) *component {
	c.el.Style(styles.BorderXRedAlpha(scale))
	return c
}

func (c *component) BorderXRedDark(scale int) *component {
	c.el.Style(styles.BorderXRedDark(scale))
	return c
}

func (c *component) BorderXRedDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXRedDarkAlpha(scale))
	return c
}

func (c *component) BorderXRuby(scale int) *component {
	c.el.Style(styles.BorderXRuby(scale))
	return c
}

func (c *component) BorderXRubyAlpha(scale int) *component {
	c.el.Style(styles.BorderXRubyAlpha(scale))
	return c
}

func (c *component) BorderXRubyDark(scale int) *component {
	c.el.Style(styles.BorderXRubyDark(scale))
	return c
}

func (c *component) BorderXRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXRubyDarkAlpha(scale))
	return c
}

func (c *component) BorderXSage(scale int) *component {
	c.el.Style(styles.BorderXSage(scale))
	return c
}

func (c *component) BorderXSageAlpha(scale int) *component {
	c.el.Style(styles.BorderXSageAlpha(scale))
	return c
}

func (c *component) BorderXSageDark(scale int) *component {
	c.el.Style(styles.BorderXSageDark(scale))
	return c
}

func (c *component) BorderXSageDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXSageDarkAlpha(scale))
	return c
}

func (c *component) BorderXSand(scale int) *component {
	c.el.Style(styles.BorderXSand(scale))
	return c
}

func (c *component) BorderXSandAlpha(scale int) *component {
	c.el.Style(styles.BorderXSandAlpha(scale))
	return c
}

func (c *component) BorderXSandDark(scale int) *component {
	c.el.Style(styles.BorderXSandDark(scale))
	return c
}

func (c *component) BorderXSandDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXSandDarkAlpha(scale))
	return c
}

func (c *component) BorderXSky(scale int) *component {
	c.el.Style(styles.BorderXSky(scale))
	return c
}

func (c *component) BorderXSkyAlpha(scale int) *component {
	c.el.Style(styles.BorderXSkyAlpha(scale))
	return c
}

func (c *component) BorderXSkyDark(scale int) *component {
	c.el.Style(styles.BorderXSkyDark(scale))
	return c
}

func (c *component) BorderXSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXSkyDarkAlpha(scale))
	return c
}

func (c *component) BorderXSlate(scale int) *component {
	c.el.Style(styles.BorderXSlate(scale))
	return c
}

func (c *component) BorderXSlateAlpha(scale int) *component {
	c.el.Style(styles.BorderXSlateAlpha(scale))
	return c
}

func (c *component) BorderXSlateDark(scale int) *component {
	c.el.Style(styles.BorderXSlateDark(scale))
	return c
}

func (c *component) BorderXSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXSlateDarkAlpha(scale))
	return c
}

func (c *component) BorderXTeal(scale int) *component {
	c.el.Style(styles.BorderXTeal(scale))
	return c
}

func (c *component) BorderXTealAlpha(scale int) *component {
	c.el.Style(styles.BorderXTealAlpha(scale))
	return c
}

func (c *component) BorderXTealDark(scale int) *component {
	c.el.Style(styles.BorderXTealDark(scale))
	return c
}

func (c *component) BorderXTealDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXTealDarkAlpha(scale))
	return c
}

func (c *component) BorderXTomato(scale int) *component {
	c.el.Style(styles.BorderXTomato(scale))
	return c
}

func (c *component) BorderXTomatoAlpha(scale int) *component {
	c.el.Style(styles.BorderXTomatoAlpha(scale))
	return c
}

func (c *component) BorderXTomatoDark(scale int) *component {
	c.el.Style(styles.BorderXTomatoDark(scale))
	return c
}

func (c *component) BorderXTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXTomatoDarkAlpha(scale))
	return c
}

func (c *component) BorderXTransparent() *component {
	c.el.Style(styles.BorderXTransparent())
	return c
}

func (c *component) BorderXViolet(scale int) *component {
	c.el.Style(styles.BorderXViolet(scale))
	return c
}

func (c *component) BorderXVioletAlpha(scale int) *component {
	c.el.Style(styles.BorderXVioletAlpha(scale))
	return c
}

func (c *component) BorderXVioletDark(scale int) *component {
	c.el.Style(styles.BorderXVioletDark(scale))
	return c
}

func (c *component) BorderXVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXVioletDarkAlpha(scale))
	return c
}

func (c *component) BorderXWhite() *component {
	c.el.Style(styles.BorderXWhite())
	return c
}

func (c *component) BorderXWhiteAlpha(scale int) *component {
	c.el.Style(styles.BorderXWhiteAlpha(scale))
	return c
}

func (c *component) BorderXYellow(scale int) *component {
	c.el.Style(styles.BorderXYellow(scale))
	return c
}

func (c *component) BorderXYellowAlpha(scale int) *component {
	c.el.Style(styles.BorderXYellowAlpha(scale))
	return c
}

func (c *component) BorderXYellowDark(scale int) *component {
	c.el.Style(styles.BorderXYellowDark(scale))
	return c
}

func (c *component) BorderXYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderXYellowDarkAlpha(scale))
	return c
}

func (c *component) BorderY(val ...value.Value) *component {
	c.el.Style(styles.BorderY(val...))
	return c
}

func (c *component) BorderYAmber(scale int) *component {
	c.el.Style(styles.BorderYAmber(scale))
	return c
}

func (c *component) BorderYAmberAlpha(scale int) *component {
	c.el.Style(styles.BorderYAmberAlpha(scale))
	return c
}

func (c *component) BorderYAmberDark(scale int) *component {
	c.el.Style(styles.BorderYAmberDark(scale))
	return c
}

func (c *component) BorderYAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYAmberDarkAlpha(scale))
	return c
}

func (c *component) BorderYBlack() *component {
	c.el.Style(styles.BorderYBlack())
	return c
}

func (c *component) BorderYBlackAlpha(scale int) *component {
	c.el.Style(styles.BorderYBlackAlpha(scale))
	return c
}

func (c *component) BorderYBlue(scale int) *component {
	c.el.Style(styles.BorderYBlue(scale))
	return c
}

func (c *component) BorderYBlueAlpha(scale int) *component {
	c.el.Style(styles.BorderYBlueAlpha(scale))
	return c
}

func (c *component) BorderYBlueDark(scale int) *component {
	c.el.Style(styles.BorderYBlueDark(scale))
	return c
}

func (c *component) BorderYBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYBlueDarkAlpha(scale))
	return c
}

func (c *component) BorderYBronze(scale int) *component {
	c.el.Style(styles.BorderYBronze(scale))
	return c
}

func (c *component) BorderYBronzeAlpha(scale int) *component {
	c.el.Style(styles.BorderYBronzeAlpha(scale))
	return c
}

func (c *component) BorderYBronzeDark(scale int) *component {
	c.el.Style(styles.BorderYBronzeDark(scale))
	return c
}

func (c *component) BorderYBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYBronzeDarkAlpha(scale))
	return c
}

func (c *component) BorderYBrown(scale int) *component {
	c.el.Style(styles.BorderYBrown(scale))
	return c
}

func (c *component) BorderYBrownAlpha(scale int) *component {
	c.el.Style(styles.BorderYBrownAlpha(scale))
	return c
}

func (c *component) BorderYBrownDark(scale int) *component {
	c.el.Style(styles.BorderYBrownDark(scale))
	return c
}

func (c *component) BorderYBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYBrownDarkAlpha(scale))
	return c
}

func (c *component) BorderYColor(val value.Value) *component {
	c.el.Style(styles.BorderYColor(val))
	return c
}

func (c *component) BorderYCrimson(scale int) *component {
	c.el.Style(styles.BorderYCrimson(scale))
	return c
}

func (c *component) BorderYCrimsonAlpha(scale int) *component {
	c.el.Style(styles.BorderYCrimsonAlpha(scale))
	return c
}

func (c *component) BorderYCrimsonDark(scale int) *component {
	c.el.Style(styles.BorderYCrimsonDark(scale))
	return c
}

func (c *component) BorderYCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYCrimsonDarkAlpha(scale))
	return c
}

func (c *component) BorderYCurrent() *component {
	c.el.Style(styles.BorderYCurrent())
	return c
}

func (c *component) BorderYCyan(scale int) *component {
	c.el.Style(styles.BorderYCyan(scale))
	return c
}

func (c *component) BorderYCyanAlpha(scale int) *component {
	c.el.Style(styles.BorderYCyanAlpha(scale))
	return c
}

func (c *component) BorderYCyanDark(scale int) *component {
	c.el.Style(styles.BorderYCyanDark(scale))
	return c
}

func (c *component) BorderYCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYCyanDarkAlpha(scale))
	return c
}

func (c *component) BorderYGold(scale int) *component {
	c.el.Style(styles.BorderYGold(scale))
	return c
}

func (c *component) BorderYGoldAlpha(scale int) *component {
	c.el.Style(styles.BorderYGoldAlpha(scale))
	return c
}

func (c *component) BorderYGoldDark(scale int) *component {
	c.el.Style(styles.BorderYGoldDark(scale))
	return c
}

func (c *component) BorderYGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYGoldDarkAlpha(scale))
	return c
}

func (c *component) BorderYGrass(scale int) *component {
	c.el.Style(styles.BorderYGrass(scale))
	return c
}

func (c *component) BorderYGrassAlpha(scale int) *component {
	c.el.Style(styles.BorderYGrassAlpha(scale))
	return c
}

func (c *component) BorderYGrassDark(scale int) *component {
	c.el.Style(styles.BorderYGrassDark(scale))
	return c
}

func (c *component) BorderYGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYGrassDarkAlpha(scale))
	return c
}

func (c *component) BorderYGray(scale int) *component {
	c.el.Style(styles.BorderYGray(scale))
	return c
}

func (c *component) BorderYGrayAlpha(scale int) *component {
	c.el.Style(styles.BorderYGrayAlpha(scale))
	return c
}

func (c *component) BorderYGrayDark(scale int) *component {
	c.el.Style(styles.BorderYGrayDark(scale))
	return c
}

func (c *component) BorderYGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYGrayDarkAlpha(scale))
	return c
}

func (c *component) BorderYGreen(scale int) *component {
	c.el.Style(styles.BorderYGreen(scale))
	return c
}

func (c *component) BorderYGreenAlpha(scale int) *component {
	c.el.Style(styles.BorderYGreenAlpha(scale))
	return c
}

func (c *component) BorderYGreenDark(scale int) *component {
	c.el.Style(styles.BorderYGreenDark(scale))
	return c
}

func (c *component) BorderYGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYGreenDarkAlpha(scale))
	return c
}

func (c *component) BorderYIndigo(scale int) *component {
	c.el.Style(styles.BorderYIndigo(scale))
	return c
}

func (c *component) BorderYIndigoAlpha(scale int) *component {
	c.el.Style(styles.BorderYIndigoAlpha(scale))
	return c
}

func (c *component) BorderYIndigoDark(scale int) *component {
	c.el.Style(styles.BorderYIndigoDark(scale))
	return c
}

func (c *component) BorderYIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYIndigoDarkAlpha(scale))
	return c
}

func (c *component) BorderYInherit() *component {
	c.el.Style(styles.BorderYInherit())
	return c
}

func (c *component) BorderYIris(scale int) *component {
	c.el.Style(styles.BorderYIris(scale))
	return c
}

func (c *component) BorderYIrisAlpha(scale int) *component {
	c.el.Style(styles.BorderYIrisAlpha(scale))
	return c
}

func (c *component) BorderYIrisDark(scale int) *component {
	c.el.Style(styles.BorderYIrisDark(scale))
	return c
}

func (c *component) BorderYIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYIrisDarkAlpha(scale))
	return c
}

func (c *component) BorderYJade(scale int) *component {
	c.el.Style(styles.BorderYJade(scale))
	return c
}

func (c *component) BorderYJadeAlpha(scale int) *component {
	c.el.Style(styles.BorderYJadeAlpha(scale))
	return c
}

func (c *component) BorderYJadeDark(scale int) *component {
	c.el.Style(styles.BorderYJadeDark(scale))
	return c
}

func (c *component) BorderYJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYJadeDarkAlpha(scale))
	return c
}

func (c *component) BorderYLime(scale int) *component {
	c.el.Style(styles.BorderYLime(scale))
	return c
}

func (c *component) BorderYLimeAlpha(scale int) *component {
	c.el.Style(styles.BorderYLimeAlpha(scale))
	return c
}

func (c *component) BorderYLimeDark(scale int) *component {
	c.el.Style(styles.BorderYLimeDark(scale))
	return c
}

func (c *component) BorderYLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYLimeDarkAlpha(scale))
	return c
}

func (c *component) BorderYMauve(scale int) *component {
	c.el.Style(styles.BorderYMauve(scale))
	return c
}

func (c *component) BorderYMauveAlpha(scale int) *component {
	c.el.Style(styles.BorderYMauveAlpha(scale))
	return c
}

func (c *component) BorderYMauveDark(scale int) *component {
	c.el.Style(styles.BorderYMauveDark(scale))
	return c
}

func (c *component) BorderYMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYMauveDarkAlpha(scale))
	return c
}

func (c *component) BorderYMint(scale int) *component {
	c.el.Style(styles.BorderYMint(scale))
	return c
}

func (c *component) BorderYMintAlpha(scale int) *component {
	c.el.Style(styles.BorderYMintAlpha(scale))
	return c
}

func (c *component) BorderYMintDark(scale int) *component {
	c.el.Style(styles.BorderYMintDark(scale))
	return c
}

func (c *component) BorderYMintDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYMintDarkAlpha(scale))
	return c
}

func (c *component) BorderYOlive(scale int) *component {
	c.el.Style(styles.BorderYOlive(scale))
	return c
}

func (c *component) BorderYOliveAlpha(scale int) *component {
	c.el.Style(styles.BorderYOliveAlpha(scale))
	return c
}

func (c *component) BorderYOliveDark(scale int) *component {
	c.el.Style(styles.BorderYOliveDark(scale))
	return c
}

func (c *component) BorderYOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYOliveDarkAlpha(scale))
	return c
}

func (c *component) BorderYOrange(scale int) *component {
	c.el.Style(styles.BorderYOrange(scale))
	return c
}

func (c *component) BorderYOrangeAlpha(scale int) *component {
	c.el.Style(styles.BorderYOrangeAlpha(scale))
	return c
}

func (c *component) BorderYOrangeDark(scale int) *component {
	c.el.Style(styles.BorderYOrangeDark(scale))
	return c
}

func (c *component) BorderYOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYOrangeDarkAlpha(scale))
	return c
}

func (c *component) BorderYPink(scale int) *component {
	c.el.Style(styles.BorderYPink(scale))
	return c
}

func (c *component) BorderYPinkAlpha(scale int) *component {
	c.el.Style(styles.BorderYPinkAlpha(scale))
	return c
}

func (c *component) BorderYPinkDark(scale int) *component {
	c.el.Style(styles.BorderYPinkDark(scale))
	return c
}

func (c *component) BorderYPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYPinkDarkAlpha(scale))
	return c
}

func (c *component) BorderYPlum(scale int) *component {
	c.el.Style(styles.BorderYPlum(scale))
	return c
}

func (c *component) BorderYPlumAlpha(scale int) *component {
	c.el.Style(styles.BorderYPlumAlpha(scale))
	return c
}

func (c *component) BorderYPlumDark(scale int) *component {
	c.el.Style(styles.BorderYPlumDark(scale))
	return c
}

func (c *component) BorderYPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYPlumDarkAlpha(scale))
	return c
}

func (c *component) BorderYPurple(scale int) *component {
	c.el.Style(styles.BorderYPurple(scale))
	return c
}

func (c *component) BorderYPurpleAlpha(scale int) *component {
	c.el.Style(styles.BorderYPurpleAlpha(scale))
	return c
}

func (c *component) BorderYPurpleDark(scale int) *component {
	c.el.Style(styles.BorderYPurpleDark(scale))
	return c
}

func (c *component) BorderYPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYPurpleDarkAlpha(scale))
	return c
}

func (c *component) BorderYRed(scale int) *component {
	c.el.Style(styles.BorderYRed(scale))
	return c
}

func (c *component) BorderYRedAlpha(scale int) *component {
	c.el.Style(styles.BorderYRedAlpha(scale))
	return c
}

func (c *component) BorderYRedDark(scale int) *component {
	c.el.Style(styles.BorderYRedDark(scale))
	return c
}

func (c *component) BorderYRedDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYRedDarkAlpha(scale))
	return c
}

func (c *component) BorderYRuby(scale int) *component {
	c.el.Style(styles.BorderYRuby(scale))
	return c
}

func (c *component) BorderYRubyAlpha(scale int) *component {
	c.el.Style(styles.BorderYRubyAlpha(scale))
	return c
}

func (c *component) BorderYRubyDark(scale int) *component {
	c.el.Style(styles.BorderYRubyDark(scale))
	return c
}

func (c *component) BorderYRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYRubyDarkAlpha(scale))
	return c
}

func (c *component) BorderYSage(scale int) *component {
	c.el.Style(styles.BorderYSage(scale))
	return c
}

func (c *component) BorderYSageAlpha(scale int) *component {
	c.el.Style(styles.BorderYSageAlpha(scale))
	return c
}

func (c *component) BorderYSageDark(scale int) *component {
	c.el.Style(styles.BorderYSageDark(scale))
	return c
}

func (c *component) BorderYSageDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYSageDarkAlpha(scale))
	return c
}

func (c *component) BorderYSand(scale int) *component {
	c.el.Style(styles.BorderYSand(scale))
	return c
}

func (c *component) BorderYSandAlpha(scale int) *component {
	c.el.Style(styles.BorderYSandAlpha(scale))
	return c
}

func (c *component) BorderYSandDark(scale int) *component {
	c.el.Style(styles.BorderYSandDark(scale))
	return c
}

func (c *component) BorderYSandDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYSandDarkAlpha(scale))
	return c
}

func (c *component) BorderYSky(scale int) *component {
	c.el.Style(styles.BorderYSky(scale))
	return c
}

func (c *component) BorderYSkyAlpha(scale int) *component {
	c.el.Style(styles.BorderYSkyAlpha(scale))
	return c
}

func (c *component) BorderYSkyDark(scale int) *component {
	c.el.Style(styles.BorderYSkyDark(scale))
	return c
}

func (c *component) BorderYSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYSkyDarkAlpha(scale))
	return c
}

func (c *component) BorderYSlate(scale int) *component {
	c.el.Style(styles.BorderYSlate(scale))
	return c
}

func (c *component) BorderYSlateAlpha(scale int) *component {
	c.el.Style(styles.BorderYSlateAlpha(scale))
	return c
}

func (c *component) BorderYSlateDark(scale int) *component {
	c.el.Style(styles.BorderYSlateDark(scale))
	return c
}

func (c *component) BorderYSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYSlateDarkAlpha(scale))
	return c
}

func (c *component) BorderYTeal(scale int) *component {
	c.el.Style(styles.BorderYTeal(scale))
	return c
}

func (c *component) BorderYTealAlpha(scale int) *component {
	c.el.Style(styles.BorderYTealAlpha(scale))
	return c
}

func (c *component) BorderYTealDark(scale int) *component {
	c.el.Style(styles.BorderYTealDark(scale))
	return c
}

func (c *component) BorderYTealDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYTealDarkAlpha(scale))
	return c
}

func (c *component) BorderYTomato(scale int) *component {
	c.el.Style(styles.BorderYTomato(scale))
	return c
}

func (c *component) BorderYTomatoAlpha(scale int) *component {
	c.el.Style(styles.BorderYTomatoAlpha(scale))
	return c
}

func (c *component) BorderYTomatoDark(scale int) *component {
	c.el.Style(styles.BorderYTomatoDark(scale))
	return c
}

func (c *component) BorderYTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYTomatoDarkAlpha(scale))
	return c
}

func (c *component) BorderYTransparent() *component {
	c.el.Style(styles.BorderYTransparent())
	return c
}

func (c *component) BorderYViolet(scale int) *component {
	c.el.Style(styles.BorderYViolet(scale))
	return c
}

func (c *component) BorderYVioletAlpha(scale int) *component {
	c.el.Style(styles.BorderYVioletAlpha(scale))
	return c
}

func (c *component) BorderYVioletDark(scale int) *component {
	c.el.Style(styles.BorderYVioletDark(scale))
	return c
}

func (c *component) BorderYVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYVioletDarkAlpha(scale))
	return c
}

func (c *component) BorderYWhite() *component {
	c.el.Style(styles.BorderYWhite())
	return c
}

func (c *component) BorderYWhiteAlpha(scale int) *component {
	c.el.Style(styles.BorderYWhiteAlpha(scale))
	return c
}

func (c *component) BorderYYellow(scale int) *component {
	c.el.Style(styles.BorderYYellow(scale))
	return c
}

func (c *component) BorderYYellowAlpha(scale int) *component {
	c.el.Style(styles.BorderYYellowAlpha(scale))
	return c
}

func (c *component) BorderYYellowDark(scale int) *component {
	c.el.Style(styles.BorderYYellowDark(scale))
	return c
}

func (c *component) BorderYYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYYellowDarkAlpha(scale))
	return c
}

func (c *component) BorderYellow(scale int) *component {
	c.el.Style(styles.BorderYellow(scale))
	return c
}

func (c *component) BorderYellowAlpha(scale int) *component {
	c.el.Style(styles.BorderYellowAlpha(scale))
	return c
}

func (c *component) BorderYellowDark(scale int) *component {
	c.el.Style(styles.BorderYellowDark(scale))
	return c
}

func (c *component) BorderYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.BorderYellowDarkAlpha(scale))
	return c
}

func (c *component) Bottom(number int) *component {
	c.el.Style(styles.Bottom(number))
	return c
}

func (c *component) BottomAuto() *component {
	c.el.Style(styles.BottomAuto())
	return c
}

func (c *component) BottomBy(val value.Value) *component {
	c.el.Style(styles.BottomBy(val))
	return c
}

func (c *component) BottomFraction(fraction float64) *component {
	c.el.Style(styles.BottomFraction(fraction))
	return c
}

func (c *component) BottomFull() *component {
	c.el.Style(styles.BottomFull())
	return c
}

func (c *component) BottomPx() *component {
	c.el.Style(styles.BottomPx())
	return c
}

func (c *component) BoxBorder() *component {
	c.el.Style(styles.BoxBorder())
	return c
}

func (c *component) BoxContent() *component {
	c.el.Style(styles.BoxContent())
	return c
}

func (c *component) BoxDecorationClone() *component {
	c.el.Style(styles.BoxDecorationClone())
	return c
}

func (c *component) BoxDecorationSlice() *component {
	c.el.Style(styles.BoxDecorationSlice())
	return c
}

func (c *component) BreakAfterAll() *component {
	c.el.Style(styles.BreakAfterAll())
	return c
}

func (c *component) BreakAfterAuto() *component {
	c.el.Style(styles.BreakAfterAuto())
	return c
}

func (c *component) BreakAfterAvoid() *component {
	c.el.Style(styles.BreakAfterAvoid())
	return c
}

func (c *component) BreakAfterAvoidPage() *component {
	c.el.Style(styles.BreakAfterAvoidPage())
	return c
}

func (c *component) BreakAfterColumn() *component {
	c.el.Style(styles.BreakAfterColumn())
	return c
}

func (c *component) BreakAfterLeft() *component {
	c.el.Style(styles.BreakAfterLeft())
	return c
}

func (c *component) BreakAfterPage() *component {
	c.el.Style(styles.BreakAfterPage())
	return c
}

func (c *component) BreakAfterRight() *component {
	c.el.Style(styles.BreakAfterRight())
	return c
}

func (c *component) BreakAll() *component {
	c.el.Style(styles.BreakAll())
	return c
}

func (c *component) BreakBeforeAll() *component {
	c.el.Style(styles.BreakBeforeAll())
	return c
}

func (c *component) BreakBeforeAuto() *component {
	c.el.Style(styles.BreakBeforeAuto())
	return c
}

func (c *component) BreakBeforeAvoid() *component {
	c.el.Style(styles.BreakBeforeAvoid())
	return c
}

func (c *component) BreakBeforeAvoidPage() *component {
	c.el.Style(styles.BreakBeforeAvoidPage())
	return c
}

func (c *component) BreakBeforeColumn() *component {
	c.el.Style(styles.BreakBeforeColumn())
	return c
}

func (c *component) BreakBeforeLeft() *component {
	c.el.Style(styles.BreakBeforeLeft())
	return c
}

func (c *component) BreakBeforePage() *component {
	c.el.Style(styles.BreakBeforePage())
	return c
}

func (c *component) BreakBeforeRight() *component {
	c.el.Style(styles.BreakBeforeRight())
	return c
}

func (c *component) BreakInsideAuto() *component {
	c.el.Style(styles.BreakInsideAuto())
	return c
}

func (c *component) BreakInsideAvoid() *component {
	c.el.Style(styles.BreakInsideAvoid())
	return c
}

func (c *component) BreakInsideAvoidColumn() *component {
	c.el.Style(styles.BreakInsideAvoidColumn())
	return c
}

func (c *component) BreakInsideAvoidPage() *component {
	c.el.Style(styles.BreakInsideAvoidPage())
	return c
}

func (c *component) BreakKeep() *component {
	c.el.Style(styles.BreakKeep())
	return c
}

func (c *component) BreakNormal() *component {
	c.el.Style(styles.BreakNormal())
	return c
}

func (c *component) Brightness(val any) *component {
	c.el.Style(styles.Brightness(val))
	return c
}

func (c *component) Capitalize() *component {
	c.el.Style(styles.Capitalize())
	return c
}

func (c *component) CaptionBottom() *component {
	c.el.Style(styles.CaptionBottom())
	return c
}

func (c *component) CaptionTop() *component {
	c.el.Style(styles.CaptionTop())
	return c
}

func (c *component) Caret(val value.Value) *component {
	c.el.Style(styles.Caret(val))
	return c
}

func (c *component) CaretAmber(scale int) *component {
	c.el.Style(styles.CaretAmber(scale))
	return c
}

func (c *component) CaretAmberAlpha(scale int) *component {
	c.el.Style(styles.CaretAmberAlpha(scale))
	return c
}

func (c *component) CaretAmberDark(scale int) *component {
	c.el.Style(styles.CaretAmberDark(scale))
	return c
}

func (c *component) CaretAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretAmberDarkAlpha(scale))
	return c
}

func (c *component) CaretBlack() *component {
	c.el.Style(styles.CaretBlack())
	return c
}

func (c *component) CaretBlackAlpha(scale int) *component {
	c.el.Style(styles.CaretBlackAlpha(scale))
	return c
}

func (c *component) CaretBlue(scale int) *component {
	c.el.Style(styles.CaretBlue(scale))
	return c
}

func (c *component) CaretBlueAlpha(scale int) *component {
	c.el.Style(styles.CaretBlueAlpha(scale))
	return c
}

func (c *component) CaretBlueDark(scale int) *component {
	c.el.Style(styles.CaretBlueDark(scale))
	return c
}

func (c *component) CaretBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretBlueDarkAlpha(scale))
	return c
}

func (c *component) CaretBronze(scale int) *component {
	c.el.Style(styles.CaretBronze(scale))
	return c
}

func (c *component) CaretBronzeAlpha(scale int) *component {
	c.el.Style(styles.CaretBronzeAlpha(scale))
	return c
}

func (c *component) CaretBronzeDark(scale int) *component {
	c.el.Style(styles.CaretBronzeDark(scale))
	return c
}

func (c *component) CaretBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretBronzeDarkAlpha(scale))
	return c
}

func (c *component) CaretBrown(scale int) *component {
	c.el.Style(styles.CaretBrown(scale))
	return c
}

func (c *component) CaretBrownAlpha(scale int) *component {
	c.el.Style(styles.CaretBrownAlpha(scale))
	return c
}

func (c *component) CaretBrownDark(scale int) *component {
	c.el.Style(styles.CaretBrownDark(scale))
	return c
}

func (c *component) CaretBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretBrownDarkAlpha(scale))
	return c
}

func (c *component) CaretCrimson(scale int) *component {
	c.el.Style(styles.CaretCrimson(scale))
	return c
}

func (c *component) CaretCrimsonAlpha(scale int) *component {
	c.el.Style(styles.CaretCrimsonAlpha(scale))
	return c
}

func (c *component) CaretCrimsonDark(scale int) *component {
	c.el.Style(styles.CaretCrimsonDark(scale))
	return c
}

func (c *component) CaretCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretCrimsonDarkAlpha(scale))
	return c
}

func (c *component) CaretCurrent() *component {
	c.el.Style(styles.CaretCurrent())
	return c
}

func (c *component) CaretCyan(scale int) *component {
	c.el.Style(styles.CaretCyan(scale))
	return c
}

func (c *component) CaretCyanAlpha(scale int) *component {
	c.el.Style(styles.CaretCyanAlpha(scale))
	return c
}

func (c *component) CaretCyanDark(scale int) *component {
	c.el.Style(styles.CaretCyanDark(scale))
	return c
}

func (c *component) CaretCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretCyanDarkAlpha(scale))
	return c
}

func (c *component) CaretGold(scale int) *component {
	c.el.Style(styles.CaretGold(scale))
	return c
}

func (c *component) CaretGoldAlpha(scale int) *component {
	c.el.Style(styles.CaretGoldAlpha(scale))
	return c
}

func (c *component) CaretGoldDark(scale int) *component {
	c.el.Style(styles.CaretGoldDark(scale))
	return c
}

func (c *component) CaretGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretGoldDarkAlpha(scale))
	return c
}

func (c *component) CaretGrass(scale int) *component {
	c.el.Style(styles.CaretGrass(scale))
	return c
}

func (c *component) CaretGrassAlpha(scale int) *component {
	c.el.Style(styles.CaretGrassAlpha(scale))
	return c
}

func (c *component) CaretGrassDark(scale int) *component {
	c.el.Style(styles.CaretGrassDark(scale))
	return c
}

func (c *component) CaretGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretGrassDarkAlpha(scale))
	return c
}

func (c *component) CaretGray(scale int) *component {
	c.el.Style(styles.CaretGray(scale))
	return c
}

func (c *component) CaretGrayAlpha(scale int) *component {
	c.el.Style(styles.CaretGrayAlpha(scale))
	return c
}

func (c *component) CaretGrayDark(scale int) *component {
	c.el.Style(styles.CaretGrayDark(scale))
	return c
}

func (c *component) CaretGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretGrayDarkAlpha(scale))
	return c
}

func (c *component) CaretGreen(scale int) *component {
	c.el.Style(styles.CaretGreen(scale))
	return c
}

func (c *component) CaretGreenAlpha(scale int) *component {
	c.el.Style(styles.CaretGreenAlpha(scale))
	return c
}

func (c *component) CaretGreenDark(scale int) *component {
	c.el.Style(styles.CaretGreenDark(scale))
	return c
}

func (c *component) CaretGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretGreenDarkAlpha(scale))
	return c
}

func (c *component) CaretIndigo(scale int) *component {
	c.el.Style(styles.CaretIndigo(scale))
	return c
}

func (c *component) CaretIndigoAlpha(scale int) *component {
	c.el.Style(styles.CaretIndigoAlpha(scale))
	return c
}

func (c *component) CaretIndigoDark(scale int) *component {
	c.el.Style(styles.CaretIndigoDark(scale))
	return c
}

func (c *component) CaretIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretIndigoDarkAlpha(scale))
	return c
}

func (c *component) CaretInherit() *component {
	c.el.Style(styles.CaretInherit())
	return c
}

func (c *component) CaretIris(scale int) *component {
	c.el.Style(styles.CaretIris(scale))
	return c
}

func (c *component) CaretIrisAlpha(scale int) *component {
	c.el.Style(styles.CaretIrisAlpha(scale))
	return c
}

func (c *component) CaretIrisDark(scale int) *component {
	c.el.Style(styles.CaretIrisDark(scale))
	return c
}

func (c *component) CaretIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretIrisDarkAlpha(scale))
	return c
}

func (c *component) CaretJade(scale int) *component {
	c.el.Style(styles.CaretJade(scale))
	return c
}

func (c *component) CaretJadeAlpha(scale int) *component {
	c.el.Style(styles.CaretJadeAlpha(scale))
	return c
}

func (c *component) CaretJadeDark(scale int) *component {
	c.el.Style(styles.CaretJadeDark(scale))
	return c
}

func (c *component) CaretJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretJadeDarkAlpha(scale))
	return c
}

func (c *component) CaretLime(scale int) *component {
	c.el.Style(styles.CaretLime(scale))
	return c
}

func (c *component) CaretLimeAlpha(scale int) *component {
	c.el.Style(styles.CaretLimeAlpha(scale))
	return c
}

func (c *component) CaretLimeDark(scale int) *component {
	c.el.Style(styles.CaretLimeDark(scale))
	return c
}

func (c *component) CaretLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretLimeDarkAlpha(scale))
	return c
}

func (c *component) CaretMauve(scale int) *component {
	c.el.Style(styles.CaretMauve(scale))
	return c
}

func (c *component) CaretMauveAlpha(scale int) *component {
	c.el.Style(styles.CaretMauveAlpha(scale))
	return c
}

func (c *component) CaretMauveDark(scale int) *component {
	c.el.Style(styles.CaretMauveDark(scale))
	return c
}

func (c *component) CaretMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretMauveDarkAlpha(scale))
	return c
}

func (c *component) CaretMint(scale int) *component {
	c.el.Style(styles.CaretMint(scale))
	return c
}

func (c *component) CaretMintAlpha(scale int) *component {
	c.el.Style(styles.CaretMintAlpha(scale))
	return c
}

func (c *component) CaretMintDark(scale int) *component {
	c.el.Style(styles.CaretMintDark(scale))
	return c
}

func (c *component) CaretMintDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretMintDarkAlpha(scale))
	return c
}

func (c *component) CaretOlive(scale int) *component {
	c.el.Style(styles.CaretOlive(scale))
	return c
}

func (c *component) CaretOliveAlpha(scale int) *component {
	c.el.Style(styles.CaretOliveAlpha(scale))
	return c
}

func (c *component) CaretOliveDark(scale int) *component {
	c.el.Style(styles.CaretOliveDark(scale))
	return c
}

func (c *component) CaretOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretOliveDarkAlpha(scale))
	return c
}

func (c *component) CaretOrange(scale int) *component {
	c.el.Style(styles.CaretOrange(scale))
	return c
}

func (c *component) CaretOrangeAlpha(scale int) *component {
	c.el.Style(styles.CaretOrangeAlpha(scale))
	return c
}

func (c *component) CaretOrangeDark(scale int) *component {
	c.el.Style(styles.CaretOrangeDark(scale))
	return c
}

func (c *component) CaretOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretOrangeDarkAlpha(scale))
	return c
}

func (c *component) CaretPink(scale int) *component {
	c.el.Style(styles.CaretPink(scale))
	return c
}

func (c *component) CaretPinkAlpha(scale int) *component {
	c.el.Style(styles.CaretPinkAlpha(scale))
	return c
}

func (c *component) CaretPinkDark(scale int) *component {
	c.el.Style(styles.CaretPinkDark(scale))
	return c
}

func (c *component) CaretPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretPinkDarkAlpha(scale))
	return c
}

func (c *component) CaretPlum(scale int) *component {
	c.el.Style(styles.CaretPlum(scale))
	return c
}

func (c *component) CaretPlumAlpha(scale int) *component {
	c.el.Style(styles.CaretPlumAlpha(scale))
	return c
}

func (c *component) CaretPlumDark(scale int) *component {
	c.el.Style(styles.CaretPlumDark(scale))
	return c
}

func (c *component) CaretPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretPlumDarkAlpha(scale))
	return c
}

func (c *component) CaretPurple(scale int) *component {
	c.el.Style(styles.CaretPurple(scale))
	return c
}

func (c *component) CaretPurpleAlpha(scale int) *component {
	c.el.Style(styles.CaretPurpleAlpha(scale))
	return c
}

func (c *component) CaretPurpleDark(scale int) *component {
	c.el.Style(styles.CaretPurpleDark(scale))
	return c
}

func (c *component) CaretPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretPurpleDarkAlpha(scale))
	return c
}

func (c *component) CaretRed(scale int) *component {
	c.el.Style(styles.CaretRed(scale))
	return c
}

func (c *component) CaretRedAlpha(scale int) *component {
	c.el.Style(styles.CaretRedAlpha(scale))
	return c
}

func (c *component) CaretRedDark(scale int) *component {
	c.el.Style(styles.CaretRedDark(scale))
	return c
}

func (c *component) CaretRedDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretRedDarkAlpha(scale))
	return c
}

func (c *component) CaretRuby(scale int) *component {
	c.el.Style(styles.CaretRuby(scale))
	return c
}

func (c *component) CaretRubyAlpha(scale int) *component {
	c.el.Style(styles.CaretRubyAlpha(scale))
	return c
}

func (c *component) CaretRubyDark(scale int) *component {
	c.el.Style(styles.CaretRubyDark(scale))
	return c
}

func (c *component) CaretRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretRubyDarkAlpha(scale))
	return c
}

func (c *component) CaretSage(scale int) *component {
	c.el.Style(styles.CaretSage(scale))
	return c
}

func (c *component) CaretSageAlpha(scale int) *component {
	c.el.Style(styles.CaretSageAlpha(scale))
	return c
}

func (c *component) CaretSageDark(scale int) *component {
	c.el.Style(styles.CaretSageDark(scale))
	return c
}

func (c *component) CaretSageDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretSageDarkAlpha(scale))
	return c
}

func (c *component) CaretSand(scale int) *component {
	c.el.Style(styles.CaretSand(scale))
	return c
}

func (c *component) CaretSandAlpha(scale int) *component {
	c.el.Style(styles.CaretSandAlpha(scale))
	return c
}

func (c *component) CaretSandDark(scale int) *component {
	c.el.Style(styles.CaretSandDark(scale))
	return c
}

func (c *component) CaretSandDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretSandDarkAlpha(scale))
	return c
}

func (c *component) CaretSky(scale int) *component {
	c.el.Style(styles.CaretSky(scale))
	return c
}

func (c *component) CaretSkyAlpha(scale int) *component {
	c.el.Style(styles.CaretSkyAlpha(scale))
	return c
}

func (c *component) CaretSkyDark(scale int) *component {
	c.el.Style(styles.CaretSkyDark(scale))
	return c
}

func (c *component) CaretSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretSkyDarkAlpha(scale))
	return c
}

func (c *component) CaretSlate(scale int) *component {
	c.el.Style(styles.CaretSlate(scale))
	return c
}

func (c *component) CaretSlateAlpha(scale int) *component {
	c.el.Style(styles.CaretSlateAlpha(scale))
	return c
}

func (c *component) CaretSlateDark(scale int) *component {
	c.el.Style(styles.CaretSlateDark(scale))
	return c
}

func (c *component) CaretSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretSlateDarkAlpha(scale))
	return c
}

func (c *component) CaretTeal(scale int) *component {
	c.el.Style(styles.CaretTeal(scale))
	return c
}

func (c *component) CaretTealAlpha(scale int) *component {
	c.el.Style(styles.CaretTealAlpha(scale))
	return c
}

func (c *component) CaretTealDark(scale int) *component {
	c.el.Style(styles.CaretTealDark(scale))
	return c
}

func (c *component) CaretTealDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretTealDarkAlpha(scale))
	return c
}

func (c *component) CaretTomato(scale int) *component {
	c.el.Style(styles.CaretTomato(scale))
	return c
}

func (c *component) CaretTomatoAlpha(scale int) *component {
	c.el.Style(styles.CaretTomatoAlpha(scale))
	return c
}

func (c *component) CaretTomatoDark(scale int) *component {
	c.el.Style(styles.CaretTomatoDark(scale))
	return c
}

func (c *component) CaretTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretTomatoDarkAlpha(scale))
	return c
}

func (c *component) CaretTransparent() *component {
	c.el.Style(styles.CaretTransparent())
	return c
}

func (c *component) CaretViolet(scale int) *component {
	c.el.Style(styles.CaretViolet(scale))
	return c
}

func (c *component) CaretVioletAlpha(scale int) *component {
	c.el.Style(styles.CaretVioletAlpha(scale))
	return c
}

func (c *component) CaretVioletDark(scale int) *component {
	c.el.Style(styles.CaretVioletDark(scale))
	return c
}

func (c *component) CaretVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretVioletDarkAlpha(scale))
	return c
}

func (c *component) CaretWhite() *component {
	c.el.Style(styles.CaretWhite())
	return c
}

func (c *component) CaretWhiteAlpha(scale int) *component {
	c.el.Style(styles.CaretWhiteAlpha(scale))
	return c
}

func (c *component) CaretYellow(scale int) *component {
	c.el.Style(styles.CaretYellow(scale))
	return c
}

func (c *component) CaretYellowAlpha(scale int) *component {
	c.el.Style(styles.CaretYellowAlpha(scale))
	return c
}

func (c *component) CaretYellowDark(scale int) *component {
	c.el.Style(styles.CaretYellowDark(scale))
	return c
}

func (c *component) CaretYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.CaretYellowDarkAlpha(scale))
	return c
}

func (c *component) ClearBoth() *component {
	c.el.Style(styles.ClearBoth())
	return c
}

func (c *component) ClearEnd() *component {
	c.el.Style(styles.ClearEnd())
	return c
}

func (c *component) ClearLeft() *component {
	c.el.Style(styles.ClearLeft())
	return c
}

func (c *component) ClearNone() *component {
	c.el.Style(styles.ClearNone())
	return c
}

func (c *component) ClearRight() *component {
	c.el.Style(styles.ClearRight())
	return c
}

func (c *component) ClearStart() *component {
	c.el.Style(styles.ClearStart())
	return c
}

func (c *component) Col(number int) *component {
	c.el.Style(styles.Col(number))
	return c
}

func (c *component) ColAuto() *component {
	c.el.Style(styles.ColAuto())
	return c
}

func (c *component) ColBy(val value.Value) *component {
	c.el.Style(styles.ColBy(val))
	return c
}

func (c *component) ColEnd(number int) *component {
	c.el.Style(styles.ColEnd(number))
	return c
}

func (c *component) ColEndAuto() *component {
	c.el.Style(styles.ColEndAuto())
	return c
}

func (c *component) ColEndBy(val value.Value) *component {
	c.el.Style(styles.ColEndBy(val))
	return c
}

func (c *component) ColSpan(number int) *component {
	c.el.Style(styles.ColSpan(number))
	return c
}

func (c *component) ColSpanBy(val value.Value) *component {
	c.el.Style(styles.ColSpanBy(val))
	return c
}

func (c *component) ColSpanFull() *component {
	c.el.Style(styles.ColSpanFull())
	return c
}

func (c *component) ColStart(number int) *component {
	c.el.Style(styles.ColStart(number))
	return c
}

func (c *component) ColStartAuto() *component {
	c.el.Style(styles.ColStartAuto())
	return c
}

func (c *component) ColStartBy(val value.Value) *component {
	c.el.Style(styles.ColStartBy(val))
	return c
}

func (c *component) Collapse() *component {
	c.el.Style(styles.Collapse())
	return c
}

func (c *component) Columns(cols value.Value) *component {
	c.el.Style(styles.Columns(cols))
	return c
}

func (c *component) Columns2xl() *component {
	c.el.Style(styles.Columns2xl())
	return c
}

func (c *component) Columns2xs() *component {
	c.el.Style(styles.Columns2xs())
	return c
}

func (c *component) Columns3xl() *component {
	c.el.Style(styles.Columns3xl())
	return c
}

func (c *component) Columns3xs() *component {
	c.el.Style(styles.Columns3xs())
	return c
}

func (c *component) Columns4xl() *component {
	c.el.Style(styles.Columns4xl())
	return c
}

func (c *component) Columns5xl() *component {
	c.el.Style(styles.Columns5xl())
	return c
}

func (c *component) Columns6xl() *component {
	c.el.Style(styles.Columns6xl())
	return c
}

func (c *component) Columns7xl() *component {
	c.el.Style(styles.Columns7xl())
	return c
}

func (c *component) ColumnsAuto() *component {
	c.el.Style(styles.ColumnsAuto())
	return c
}

func (c *component) ColumnsLg() *component {
	c.el.Style(styles.ColumnsLg())
	return c
}

func (c *component) ColumnsMd() *component {
	c.el.Style(styles.ColumnsMd())
	return c
}

func (c *component) ColumnsSm() *component {
	c.el.Style(styles.ColumnsSm())
	return c
}

func (c *component) ColumnsXl() *component {
	c.el.Style(styles.ColumnsXl())
	return c
}

func (c *component) ColumnsXs() *component {
	c.el.Style(styles.ColumnsXs())
	return c
}

func (c *component) Container() *component {
	c.el.Style(styles.Container())
	return c
}

func (c *component) Content(val value.Value) *component {
	c.el.Style(styles.Content(val))
	return c
}

func (c *component) ContentAround() *component {
	c.el.Style(styles.ContentAround())
	return c
}

func (c *component) ContentBaseline() *component {
	c.el.Style(styles.ContentBaseline())
	return c
}

func (c *component) ContentBetween() *component {
	c.el.Style(styles.ContentBetween())
	return c
}

func (c *component) ContentCenter() *component {
	c.el.Style(styles.ContentCenter())
	return c
}

func (c *component) ContentEnd() *component {
	c.el.Style(styles.ContentEnd())
	return c
}

func (c *component) ContentEvenly() *component {
	c.el.Style(styles.ContentEvenly())
	return c
}

func (c *component) ContentNone() *component {
	c.el.Style(styles.ContentNone())
	return c
}

func (c *component) ContentNormal() *component {
	c.el.Style(styles.ContentNormal())
	return c
}

func (c *component) ContentStart() *component {
	c.el.Style(styles.ContentStart())
	return c
}

func (c *component) ContentStretch() *component {
	c.el.Style(styles.ContentStretch())
	return c
}

func (c *component) Contents() *component {
	c.el.Style(styles.Contents())
	return c
}

func (c *component) Contrast(val any) *component {
	c.el.Style(styles.Contrast(val))
	return c
}

func (c *component) Cursor(val value.Value) *component {
	c.el.Style(styles.Cursor(val))
	return c
}

func (c *component) CursorAlias() *component {
	c.el.Style(styles.CursorAlias())
	return c
}

func (c *component) CursorAllScroll() *component {
	c.el.Style(styles.CursorAllScroll())
	return c
}

func (c *component) CursorAuto() *component {
	c.el.Style(styles.CursorAuto())
	return c
}

func (c *component) CursorCell() *component {
	c.el.Style(styles.CursorCell())
	return c
}

func (c *component) CursorColResize() *component {
	c.el.Style(styles.CursorColResize())
	return c
}

func (c *component) CursorContextMenu() *component {
	c.el.Style(styles.CursorContextMenu())
	return c
}

func (c *component) CursorCopy() *component {
	c.el.Style(styles.CursorCopy())
	return c
}

func (c *component) CursorCrosshair() *component {
	c.el.Style(styles.CursorCrosshair())
	return c
}

func (c *component) CursorDefault() *component {
	c.el.Style(styles.CursorDefault())
	return c
}

func (c *component) CursorEResize() *component {
	c.el.Style(styles.CursorEResize())
	return c
}

func (c *component) CursorEwResize() *component {
	c.el.Style(styles.CursorEwResize())
	return c
}

func (c *component) CursorGrab() *component {
	c.el.Style(styles.CursorGrab())
	return c
}

func (c *component) CursorGrabbing() *component {
	c.el.Style(styles.CursorGrabbing())
	return c
}

func (c *component) CursorHelp() *component {
	c.el.Style(styles.CursorHelp())
	return c
}

func (c *component) CursorMove() *component {
	c.el.Style(styles.CursorMove())
	return c
}

func (c *component) CursorNResize() *component {
	c.el.Style(styles.CursorNResize())
	return c
}

func (c *component) CursorNeResize() *component {
	c.el.Style(styles.CursorNeResize())
	return c
}

func (c *component) CursorNeswResize() *component {
	c.el.Style(styles.CursorNeswResize())
	return c
}

func (c *component) CursorNoDrop() *component {
	c.el.Style(styles.CursorNoDrop())
	return c
}

func (c *component) CursorNone() *component {
	c.el.Style(styles.CursorNone())
	return c
}

func (c *component) CursorNotAllowed() *component {
	c.el.Style(styles.CursorNotAllowed())
	return c
}

func (c *component) CursorNsResize() *component {
	c.el.Style(styles.CursorNsResize())
	return c
}

func (c *component) CursorNwResize() *component {
	c.el.Style(styles.CursorNwResize())
	return c
}

func (c *component) CursorNwseResize() *component {
	c.el.Style(styles.CursorNwseResize())
	return c
}

func (c *component) CursorPointer() *component {
	c.el.Style(styles.CursorPointer())
	return c
}

func (c *component) CursorProgress() *component {
	c.el.Style(styles.CursorProgress())
	return c
}

func (c *component) CursorRowResize() *component {
	c.el.Style(styles.CursorRowResize())
	return c
}

func (c *component) CursorSResize() *component {
	c.el.Style(styles.CursorSResize())
	return c
}

func (c *component) CursorSeResize() *component {
	c.el.Style(styles.CursorSeResize())
	return c
}

func (c *component) CursorSwResize() *component {
	c.el.Style(styles.CursorSwResize())
	return c
}

func (c *component) CursorText() *component {
	c.el.Style(styles.CursorText())
	return c
}

func (c *component) CursorVerticalText() *component {
	c.el.Style(styles.CursorVerticalText())
	return c
}

func (c *component) CursorWResize() *component {
	c.el.Style(styles.CursorWResize())
	return c
}

func (c *component) CursorWait() *component {
	c.el.Style(styles.CursorWait())
	return c
}

func (c *component) CursorZoomIn() *component {
	c.el.Style(styles.CursorZoomIn())
	return c
}

func (c *component) CursorZoomOut() *component {
	c.el.Style(styles.CursorZoomOut())
	return c
}

func (c *component) DecorAmber(scale int) *component {
	c.el.Style(styles.DecorAmber(scale))
	return c
}

func (c *component) DecorAmberAlpha(scale int) *component {
	c.el.Style(styles.DecorAmberAlpha(scale))
	return c
}

func (c *component) DecorAmberDark(scale int) *component {
	c.el.Style(styles.DecorAmberDark(scale))
	return c
}

func (c *component) DecorAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorAmberDarkAlpha(scale))
	return c
}

func (c *component) DecorBlack() *component {
	c.el.Style(styles.DecorBlack())
	return c
}

func (c *component) DecorBlackAlpha(scale int) *component {
	c.el.Style(styles.DecorBlackAlpha(scale))
	return c
}

func (c *component) DecorBlue(scale int) *component {
	c.el.Style(styles.DecorBlue(scale))
	return c
}

func (c *component) DecorBlueAlpha(scale int) *component {
	c.el.Style(styles.DecorBlueAlpha(scale))
	return c
}

func (c *component) DecorBlueDark(scale int) *component {
	c.el.Style(styles.DecorBlueDark(scale))
	return c
}

func (c *component) DecorBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorBlueDarkAlpha(scale))
	return c
}

func (c *component) DecorBronze(scale int) *component {
	c.el.Style(styles.DecorBronze(scale))
	return c
}

func (c *component) DecorBronzeAlpha(scale int) *component {
	c.el.Style(styles.DecorBronzeAlpha(scale))
	return c
}

func (c *component) DecorBronzeDark(scale int) *component {
	c.el.Style(styles.DecorBronzeDark(scale))
	return c
}

func (c *component) DecorBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorBronzeDarkAlpha(scale))
	return c
}

func (c *component) DecorBrown(scale int) *component {
	c.el.Style(styles.DecorBrown(scale))
	return c
}

func (c *component) DecorBrownAlpha(scale int) *component {
	c.el.Style(styles.DecorBrownAlpha(scale))
	return c
}

func (c *component) DecorBrownDark(scale int) *component {
	c.el.Style(styles.DecorBrownDark(scale))
	return c
}

func (c *component) DecorBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorBrownDarkAlpha(scale))
	return c
}

func (c *component) DecorColorBy(val value.Value) *component {
	c.el.Style(styles.DecorColorBy(val))
	return c
}

func (c *component) DecorCrimson(scale int) *component {
	c.el.Style(styles.DecorCrimson(scale))
	return c
}

func (c *component) DecorCrimsonAlpha(scale int) *component {
	c.el.Style(styles.DecorCrimsonAlpha(scale))
	return c
}

func (c *component) DecorCrimsonDark(scale int) *component {
	c.el.Style(styles.DecorCrimsonDark(scale))
	return c
}

func (c *component) DecorCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorCrimsonDarkAlpha(scale))
	return c
}

func (c *component) DecorCurrent() *component {
	c.el.Style(styles.DecorCurrent())
	return c
}

func (c *component) DecorCyan(scale int) *component {
	c.el.Style(styles.DecorCyan(scale))
	return c
}

func (c *component) DecorCyanAlpha(scale int) *component {
	c.el.Style(styles.DecorCyanAlpha(scale))
	return c
}

func (c *component) DecorCyanDark(scale int) *component {
	c.el.Style(styles.DecorCyanDark(scale))
	return c
}

func (c *component) DecorCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorCyanDarkAlpha(scale))
	return c
}

func (c *component) DecorGold(scale int) *component {
	c.el.Style(styles.DecorGold(scale))
	return c
}

func (c *component) DecorGoldAlpha(scale int) *component {
	c.el.Style(styles.DecorGoldAlpha(scale))
	return c
}

func (c *component) DecorGoldDark(scale int) *component {
	c.el.Style(styles.DecorGoldDark(scale))
	return c
}

func (c *component) DecorGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorGoldDarkAlpha(scale))
	return c
}

func (c *component) DecorGrass(scale int) *component {
	c.el.Style(styles.DecorGrass(scale))
	return c
}

func (c *component) DecorGrassAlpha(scale int) *component {
	c.el.Style(styles.DecorGrassAlpha(scale))
	return c
}

func (c *component) DecorGrassDark(scale int) *component {
	c.el.Style(styles.DecorGrassDark(scale))
	return c
}

func (c *component) DecorGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorGrassDarkAlpha(scale))
	return c
}

func (c *component) DecorGray(scale int) *component {
	c.el.Style(styles.DecorGray(scale))
	return c
}

func (c *component) DecorGrayAlpha(scale int) *component {
	c.el.Style(styles.DecorGrayAlpha(scale))
	return c
}

func (c *component) DecorGrayDark(scale int) *component {
	c.el.Style(styles.DecorGrayDark(scale))
	return c
}

func (c *component) DecorGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorGrayDarkAlpha(scale))
	return c
}

func (c *component) DecorGreen(scale int) *component {
	c.el.Style(styles.DecorGreen(scale))
	return c
}

func (c *component) DecorGreenAlpha(scale int) *component {
	c.el.Style(styles.DecorGreenAlpha(scale))
	return c
}

func (c *component) DecorGreenDark(scale int) *component {
	c.el.Style(styles.DecorGreenDark(scale))
	return c
}

func (c *component) DecorGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorGreenDarkAlpha(scale))
	return c
}

func (c *component) DecorIndigo(scale int) *component {
	c.el.Style(styles.DecorIndigo(scale))
	return c
}

func (c *component) DecorIndigoAlpha(scale int) *component {
	c.el.Style(styles.DecorIndigoAlpha(scale))
	return c
}

func (c *component) DecorIndigoDark(scale int) *component {
	c.el.Style(styles.DecorIndigoDark(scale))
	return c
}

func (c *component) DecorIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorIndigoDarkAlpha(scale))
	return c
}

func (c *component) DecorInherit() *component {
	c.el.Style(styles.DecorInherit())
	return c
}

func (c *component) DecorIris(scale int) *component {
	c.el.Style(styles.DecorIris(scale))
	return c
}

func (c *component) DecorIrisAlpha(scale int) *component {
	c.el.Style(styles.DecorIrisAlpha(scale))
	return c
}

func (c *component) DecorIrisDark(scale int) *component {
	c.el.Style(styles.DecorIrisDark(scale))
	return c
}

func (c *component) DecorIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorIrisDarkAlpha(scale))
	return c
}

func (c *component) DecorJade(scale int) *component {
	c.el.Style(styles.DecorJade(scale))
	return c
}

func (c *component) DecorJadeAlpha(scale int) *component {
	c.el.Style(styles.DecorJadeAlpha(scale))
	return c
}

func (c *component) DecorJadeDark(scale int) *component {
	c.el.Style(styles.DecorJadeDark(scale))
	return c
}

func (c *component) DecorJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorJadeDarkAlpha(scale))
	return c
}

func (c *component) DecorLime(scale int) *component {
	c.el.Style(styles.DecorLime(scale))
	return c
}

func (c *component) DecorLimeAlpha(scale int) *component {
	c.el.Style(styles.DecorLimeAlpha(scale))
	return c
}

func (c *component) DecorLimeDark(scale int) *component {
	c.el.Style(styles.DecorLimeDark(scale))
	return c
}

func (c *component) DecorLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorLimeDarkAlpha(scale))
	return c
}

func (c *component) DecorMauve(scale int) *component {
	c.el.Style(styles.DecorMauve(scale))
	return c
}

func (c *component) DecorMauveAlpha(scale int) *component {
	c.el.Style(styles.DecorMauveAlpha(scale))
	return c
}

func (c *component) DecorMauveDark(scale int) *component {
	c.el.Style(styles.DecorMauveDark(scale))
	return c
}

func (c *component) DecorMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorMauveDarkAlpha(scale))
	return c
}

func (c *component) DecorMint(scale int) *component {
	c.el.Style(styles.DecorMint(scale))
	return c
}

func (c *component) DecorMintAlpha(scale int) *component {
	c.el.Style(styles.DecorMintAlpha(scale))
	return c
}

func (c *component) DecorMintDark(scale int) *component {
	c.el.Style(styles.DecorMintDark(scale))
	return c
}

func (c *component) DecorMintDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorMintDarkAlpha(scale))
	return c
}

func (c *component) DecorOlive(scale int) *component {
	c.el.Style(styles.DecorOlive(scale))
	return c
}

func (c *component) DecorOliveAlpha(scale int) *component {
	c.el.Style(styles.DecorOliveAlpha(scale))
	return c
}

func (c *component) DecorOliveDark(scale int) *component {
	c.el.Style(styles.DecorOliveDark(scale))
	return c
}

func (c *component) DecorOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorOliveDarkAlpha(scale))
	return c
}

func (c *component) DecorOrange(scale int) *component {
	c.el.Style(styles.DecorOrange(scale))
	return c
}

func (c *component) DecorOrangeAlpha(scale int) *component {
	c.el.Style(styles.DecorOrangeAlpha(scale))
	return c
}

func (c *component) DecorOrangeDark(scale int) *component {
	c.el.Style(styles.DecorOrangeDark(scale))
	return c
}

func (c *component) DecorOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorOrangeDarkAlpha(scale))
	return c
}

func (c *component) DecorPink(scale int) *component {
	c.el.Style(styles.DecorPink(scale))
	return c
}

func (c *component) DecorPinkAlpha(scale int) *component {
	c.el.Style(styles.DecorPinkAlpha(scale))
	return c
}

func (c *component) DecorPinkDark(scale int) *component {
	c.el.Style(styles.DecorPinkDark(scale))
	return c
}

func (c *component) DecorPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorPinkDarkAlpha(scale))
	return c
}

func (c *component) DecorPlum(scale int) *component {
	c.el.Style(styles.DecorPlum(scale))
	return c
}

func (c *component) DecorPlumAlpha(scale int) *component {
	c.el.Style(styles.DecorPlumAlpha(scale))
	return c
}

func (c *component) DecorPlumDark(scale int) *component {
	c.el.Style(styles.DecorPlumDark(scale))
	return c
}

func (c *component) DecorPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorPlumDarkAlpha(scale))
	return c
}

func (c *component) DecorPurple(scale int) *component {
	c.el.Style(styles.DecorPurple(scale))
	return c
}

func (c *component) DecorPurpleAlpha(scale int) *component {
	c.el.Style(styles.DecorPurpleAlpha(scale))
	return c
}

func (c *component) DecorPurpleDark(scale int) *component {
	c.el.Style(styles.DecorPurpleDark(scale))
	return c
}

func (c *component) DecorPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorPurpleDarkAlpha(scale))
	return c
}

func (c *component) DecorRed(scale int) *component {
	c.el.Style(styles.DecorRed(scale))
	return c
}

func (c *component) DecorRedAlpha(scale int) *component {
	c.el.Style(styles.DecorRedAlpha(scale))
	return c
}

func (c *component) DecorRedDark(scale int) *component {
	c.el.Style(styles.DecorRedDark(scale))
	return c
}

func (c *component) DecorRedDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorRedDarkAlpha(scale))
	return c
}

func (c *component) DecorRuby(scale int) *component {
	c.el.Style(styles.DecorRuby(scale))
	return c
}

func (c *component) DecorRubyAlpha(scale int) *component {
	c.el.Style(styles.DecorRubyAlpha(scale))
	return c
}

func (c *component) DecorRubyDark(scale int) *component {
	c.el.Style(styles.DecorRubyDark(scale))
	return c
}

func (c *component) DecorRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorRubyDarkAlpha(scale))
	return c
}

func (c *component) DecorSage(scale int) *component {
	c.el.Style(styles.DecorSage(scale))
	return c
}

func (c *component) DecorSageAlpha(scale int) *component {
	c.el.Style(styles.DecorSageAlpha(scale))
	return c
}

func (c *component) DecorSageDark(scale int) *component {
	c.el.Style(styles.DecorSageDark(scale))
	return c
}

func (c *component) DecorSageDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorSageDarkAlpha(scale))
	return c
}

func (c *component) DecorSand(scale int) *component {
	c.el.Style(styles.DecorSand(scale))
	return c
}

func (c *component) DecorSandAlpha(scale int) *component {
	c.el.Style(styles.DecorSandAlpha(scale))
	return c
}

func (c *component) DecorSandDark(scale int) *component {
	c.el.Style(styles.DecorSandDark(scale))
	return c
}

func (c *component) DecorSandDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorSandDarkAlpha(scale))
	return c
}

func (c *component) DecorSky(scale int) *component {
	c.el.Style(styles.DecorSky(scale))
	return c
}

func (c *component) DecorSkyAlpha(scale int) *component {
	c.el.Style(styles.DecorSkyAlpha(scale))
	return c
}

func (c *component) DecorSkyDark(scale int) *component {
	c.el.Style(styles.DecorSkyDark(scale))
	return c
}

func (c *component) DecorSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorSkyDarkAlpha(scale))
	return c
}

func (c *component) DecorSlate(scale int) *component {
	c.el.Style(styles.DecorSlate(scale))
	return c
}

func (c *component) DecorSlateAlpha(scale int) *component {
	c.el.Style(styles.DecorSlateAlpha(scale))
	return c
}

func (c *component) DecorSlateDark(scale int) *component {
	c.el.Style(styles.DecorSlateDark(scale))
	return c
}

func (c *component) DecorSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorSlateDarkAlpha(scale))
	return c
}

func (c *component) DecorTeal(scale int) *component {
	c.el.Style(styles.DecorTeal(scale))
	return c
}

func (c *component) DecorTealAlpha(scale int) *component {
	c.el.Style(styles.DecorTealAlpha(scale))
	return c
}

func (c *component) DecorTealDark(scale int) *component {
	c.el.Style(styles.DecorTealDark(scale))
	return c
}

func (c *component) DecorTealDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorTealDarkAlpha(scale))
	return c
}

func (c *component) DecorTomato(scale int) *component {
	c.el.Style(styles.DecorTomato(scale))
	return c
}

func (c *component) DecorTomatoAlpha(scale int) *component {
	c.el.Style(styles.DecorTomatoAlpha(scale))
	return c
}

func (c *component) DecorTomatoDark(scale int) *component {
	c.el.Style(styles.DecorTomatoDark(scale))
	return c
}

func (c *component) DecorTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorTomatoDarkAlpha(scale))
	return c
}

func (c *component) DecorTransparent() *component {
	c.el.Style(styles.DecorTransparent())
	return c
}

func (c *component) DecorViolet(scale int) *component {
	c.el.Style(styles.DecorViolet(scale))
	return c
}

func (c *component) DecorVioletAlpha(scale int) *component {
	c.el.Style(styles.DecorVioletAlpha(scale))
	return c
}

func (c *component) DecorVioletDark(scale int) *component {
	c.el.Style(styles.DecorVioletDark(scale))
	return c
}

func (c *component) DecorVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorVioletDarkAlpha(scale))
	return c
}

func (c *component) DecorWhite() *component {
	c.el.Style(styles.DecorWhite())
	return c
}

func (c *component) DecorWhiteAlpha(scale int) *component {
	c.el.Style(styles.DecorWhiteAlpha(scale))
	return c
}

func (c *component) DecorYellow(scale int) *component {
	c.el.Style(styles.DecorYellow(scale))
	return c
}

func (c *component) DecorYellowAlpha(scale int) *component {
	c.el.Style(styles.DecorYellowAlpha(scale))
	return c
}

func (c *component) DecorYellowDark(scale int) *component {
	c.el.Style(styles.DecorYellowDark(scale))
	return c
}

func (c *component) DecorYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.DecorYellowDarkAlpha(scale))
	return c
}

func (c *component) Decoration(number int) *component {
	c.el.Style(styles.Decoration(number))
	return c
}

func (c *component) DecorationAuto() *component {
	c.el.Style(styles.DecorationAuto())
	return c
}

func (c *component) DecorationBy(val value.Value) *component {
	c.el.Style(styles.DecorationBy(val))
	return c
}

func (c *component) DecorationDashed() *component {
	c.el.Style(styles.DecorationDashed())
	return c
}

func (c *component) DecorationDotted() *component {
	c.el.Style(styles.DecorationDotted())
	return c
}

func (c *component) DecorationDouble() *component {
	c.el.Style(styles.DecorationDouble())
	return c
}

func (c *component) DecorationFromFont() *component {
	c.el.Style(styles.DecorationFromFont())
	return c
}

func (c *component) DecorationSolid() *component {
	c.el.Style(styles.DecorationSolid())
	return c
}

func (c *component) DecorationWavy() *component {
	c.el.Style(styles.DecorationWavy())
	return c
}

func (c *component) Delay(val any) *component {
	c.el.Style(styles.Delay(val))
	return c
}

func (c *component) DiagonalFractions() *component {
	c.el.Style(styles.DiagonalFractions())
	return c
}

func (c *component) DivideAmber(scale int) *component {
	c.el.Style(styles.DivideAmber(scale))
	return c
}

func (c *component) DivideAmberAlpha(scale int) *component {
	c.el.Style(styles.DivideAmberAlpha(scale))
	return c
}

func (c *component) DivideAmberDark(scale int) *component {
	c.el.Style(styles.DivideAmberDark(scale))
	return c
}

func (c *component) DivideAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideAmberDarkAlpha(scale))
	return c
}

func (c *component) DivideBlack() *component {
	c.el.Style(styles.DivideBlack())
	return c
}

func (c *component) DivideBlackAlpha(scale int) *component {
	c.el.Style(styles.DivideBlackAlpha(scale))
	return c
}

func (c *component) DivideBlue(scale int) *component {
	c.el.Style(styles.DivideBlue(scale))
	return c
}

func (c *component) DivideBlueAlpha(scale int) *component {
	c.el.Style(styles.DivideBlueAlpha(scale))
	return c
}

func (c *component) DivideBlueDark(scale int) *component {
	c.el.Style(styles.DivideBlueDark(scale))
	return c
}

func (c *component) DivideBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideBlueDarkAlpha(scale))
	return c
}

func (c *component) DivideBronze(scale int) *component {
	c.el.Style(styles.DivideBronze(scale))
	return c
}

func (c *component) DivideBronzeAlpha(scale int) *component {
	c.el.Style(styles.DivideBronzeAlpha(scale))
	return c
}

func (c *component) DivideBronzeDark(scale int) *component {
	c.el.Style(styles.DivideBronzeDark(scale))
	return c
}

func (c *component) DivideBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideBronzeDarkAlpha(scale))
	return c
}

func (c *component) DivideBrown(scale int) *component {
	c.el.Style(styles.DivideBrown(scale))
	return c
}

func (c *component) DivideBrownAlpha(scale int) *component {
	c.el.Style(styles.DivideBrownAlpha(scale))
	return c
}

func (c *component) DivideBrownDark(scale int) *component {
	c.el.Style(styles.DivideBrownDark(scale))
	return c
}

func (c *component) DivideBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideBrownDarkAlpha(scale))
	return c
}

func (c *component) DivideColor(val value.Value) *component {
	c.el.Style(styles.DivideColor(val))
	return c
}

func (c *component) DivideCrimson(scale int) *component {
	c.el.Style(styles.DivideCrimson(scale))
	return c
}

func (c *component) DivideCrimsonAlpha(scale int) *component {
	c.el.Style(styles.DivideCrimsonAlpha(scale))
	return c
}

func (c *component) DivideCrimsonDark(scale int) *component {
	c.el.Style(styles.DivideCrimsonDark(scale))
	return c
}

func (c *component) DivideCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideCrimsonDarkAlpha(scale))
	return c
}

func (c *component) DivideCurrent() *component {
	c.el.Style(styles.DivideCurrent())
	return c
}

func (c *component) DivideCyan(scale int) *component {
	c.el.Style(styles.DivideCyan(scale))
	return c
}

func (c *component) DivideCyanAlpha(scale int) *component {
	c.el.Style(styles.DivideCyanAlpha(scale))
	return c
}

func (c *component) DivideCyanDark(scale int) *component {
	c.el.Style(styles.DivideCyanDark(scale))
	return c
}

func (c *component) DivideCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideCyanDarkAlpha(scale))
	return c
}

func (c *component) DivideDashed() *component {
	c.el.Style(styles.DivideDashed())
	return c
}

func (c *component) DivideDotted() *component {
	c.el.Style(styles.DivideDotted())
	return c
}

func (c *component) DivideDouble() *component {
	c.el.Style(styles.DivideDouble())
	return c
}

func (c *component) DivideGold(scale int) *component {
	c.el.Style(styles.DivideGold(scale))
	return c
}

func (c *component) DivideGoldAlpha(scale int) *component {
	c.el.Style(styles.DivideGoldAlpha(scale))
	return c
}

func (c *component) DivideGoldDark(scale int) *component {
	c.el.Style(styles.DivideGoldDark(scale))
	return c
}

func (c *component) DivideGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideGoldDarkAlpha(scale))
	return c
}

func (c *component) DivideGrass(scale int) *component {
	c.el.Style(styles.DivideGrass(scale))
	return c
}

func (c *component) DivideGrassAlpha(scale int) *component {
	c.el.Style(styles.DivideGrassAlpha(scale))
	return c
}

func (c *component) DivideGrassDark(scale int) *component {
	c.el.Style(styles.DivideGrassDark(scale))
	return c
}

func (c *component) DivideGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideGrassDarkAlpha(scale))
	return c
}

func (c *component) DivideGray(scale int) *component {
	c.el.Style(styles.DivideGray(scale))
	return c
}

func (c *component) DivideGrayAlpha(scale int) *component {
	c.el.Style(styles.DivideGrayAlpha(scale))
	return c
}

func (c *component) DivideGrayDark(scale int) *component {
	c.el.Style(styles.DivideGrayDark(scale))
	return c
}

func (c *component) DivideGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideGrayDarkAlpha(scale))
	return c
}

func (c *component) DivideGreen(scale int) *component {
	c.el.Style(styles.DivideGreen(scale))
	return c
}

func (c *component) DivideGreenAlpha(scale int) *component {
	c.el.Style(styles.DivideGreenAlpha(scale))
	return c
}

func (c *component) DivideGreenDark(scale int) *component {
	c.el.Style(styles.DivideGreenDark(scale))
	return c
}

func (c *component) DivideGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideGreenDarkAlpha(scale))
	return c
}

func (c *component) DivideHidden() *component {
	c.el.Style(styles.DivideHidden())
	return c
}

func (c *component) DivideIndigo(scale int) *component {
	c.el.Style(styles.DivideIndigo(scale))
	return c
}

func (c *component) DivideIndigoAlpha(scale int) *component {
	c.el.Style(styles.DivideIndigoAlpha(scale))
	return c
}

func (c *component) DivideIndigoDark(scale int) *component {
	c.el.Style(styles.DivideIndigoDark(scale))
	return c
}

func (c *component) DivideIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideIndigoDarkAlpha(scale))
	return c
}

func (c *component) DivideInherit() *component {
	c.el.Style(styles.DivideInherit())
	return c
}

func (c *component) DivideIris(scale int) *component {
	c.el.Style(styles.DivideIris(scale))
	return c
}

func (c *component) DivideIrisAlpha(scale int) *component {
	c.el.Style(styles.DivideIrisAlpha(scale))
	return c
}

func (c *component) DivideIrisDark(scale int) *component {
	c.el.Style(styles.DivideIrisDark(scale))
	return c
}

func (c *component) DivideIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideIrisDarkAlpha(scale))
	return c
}

func (c *component) DivideJade(scale int) *component {
	c.el.Style(styles.DivideJade(scale))
	return c
}

func (c *component) DivideJadeAlpha(scale int) *component {
	c.el.Style(styles.DivideJadeAlpha(scale))
	return c
}

func (c *component) DivideJadeDark(scale int) *component {
	c.el.Style(styles.DivideJadeDark(scale))
	return c
}

func (c *component) DivideJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideJadeDarkAlpha(scale))
	return c
}

func (c *component) DivideLime(scale int) *component {
	c.el.Style(styles.DivideLime(scale))
	return c
}

func (c *component) DivideLimeAlpha(scale int) *component {
	c.el.Style(styles.DivideLimeAlpha(scale))
	return c
}

func (c *component) DivideLimeDark(scale int) *component {
	c.el.Style(styles.DivideLimeDark(scale))
	return c
}

func (c *component) DivideLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideLimeDarkAlpha(scale))
	return c
}

func (c *component) DivideMauve(scale int) *component {
	c.el.Style(styles.DivideMauve(scale))
	return c
}

func (c *component) DivideMauveAlpha(scale int) *component {
	c.el.Style(styles.DivideMauveAlpha(scale))
	return c
}

func (c *component) DivideMauveDark(scale int) *component {
	c.el.Style(styles.DivideMauveDark(scale))
	return c
}

func (c *component) DivideMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideMauveDarkAlpha(scale))
	return c
}

func (c *component) DivideMint(scale int) *component {
	c.el.Style(styles.DivideMint(scale))
	return c
}

func (c *component) DivideMintAlpha(scale int) *component {
	c.el.Style(styles.DivideMintAlpha(scale))
	return c
}

func (c *component) DivideMintDark(scale int) *component {
	c.el.Style(styles.DivideMintDark(scale))
	return c
}

func (c *component) DivideMintDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideMintDarkAlpha(scale))
	return c
}

func (c *component) DivideNone() *component {
	c.el.Style(styles.DivideNone())
	return c
}

func (c *component) DivideOlive(scale int) *component {
	c.el.Style(styles.DivideOlive(scale))
	return c
}

func (c *component) DivideOliveAlpha(scale int) *component {
	c.el.Style(styles.DivideOliveAlpha(scale))
	return c
}

func (c *component) DivideOliveDark(scale int) *component {
	c.el.Style(styles.DivideOliveDark(scale))
	return c
}

func (c *component) DivideOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideOliveDarkAlpha(scale))
	return c
}

func (c *component) DivideOrange(scale int) *component {
	c.el.Style(styles.DivideOrange(scale))
	return c
}

func (c *component) DivideOrangeAlpha(scale int) *component {
	c.el.Style(styles.DivideOrangeAlpha(scale))
	return c
}

func (c *component) DivideOrangeDark(scale int) *component {
	c.el.Style(styles.DivideOrangeDark(scale))
	return c
}

func (c *component) DivideOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideOrangeDarkAlpha(scale))
	return c
}

func (c *component) DividePink(scale int) *component {
	c.el.Style(styles.DividePink(scale))
	return c
}

func (c *component) DividePinkAlpha(scale int) *component {
	c.el.Style(styles.DividePinkAlpha(scale))
	return c
}

func (c *component) DividePinkDark(scale int) *component {
	c.el.Style(styles.DividePinkDark(scale))
	return c
}

func (c *component) DividePinkDarkAlpha(scale int) *component {
	c.el.Style(styles.DividePinkDarkAlpha(scale))
	return c
}

func (c *component) DividePlum(scale int) *component {
	c.el.Style(styles.DividePlum(scale))
	return c
}

func (c *component) DividePlumAlpha(scale int) *component {
	c.el.Style(styles.DividePlumAlpha(scale))
	return c
}

func (c *component) DividePlumDark(scale int) *component {
	c.el.Style(styles.DividePlumDark(scale))
	return c
}

func (c *component) DividePlumDarkAlpha(scale int) *component {
	c.el.Style(styles.DividePlumDarkAlpha(scale))
	return c
}

func (c *component) DividePurple(scale int) *component {
	c.el.Style(styles.DividePurple(scale))
	return c
}

func (c *component) DividePurpleAlpha(scale int) *component {
	c.el.Style(styles.DividePurpleAlpha(scale))
	return c
}

func (c *component) DividePurpleDark(scale int) *component {
	c.el.Style(styles.DividePurpleDark(scale))
	return c
}

func (c *component) DividePurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.DividePurpleDarkAlpha(scale))
	return c
}

func (c *component) DivideRed(scale int) *component {
	c.el.Style(styles.DivideRed(scale))
	return c
}

func (c *component) DivideRedAlpha(scale int) *component {
	c.el.Style(styles.DivideRedAlpha(scale))
	return c
}

func (c *component) DivideRedDark(scale int) *component {
	c.el.Style(styles.DivideRedDark(scale))
	return c
}

func (c *component) DivideRedDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideRedDarkAlpha(scale))
	return c
}

func (c *component) DivideRuby(scale int) *component {
	c.el.Style(styles.DivideRuby(scale))
	return c
}

func (c *component) DivideRubyAlpha(scale int) *component {
	c.el.Style(styles.DivideRubyAlpha(scale))
	return c
}

func (c *component) DivideRubyDark(scale int) *component {
	c.el.Style(styles.DivideRubyDark(scale))
	return c
}

func (c *component) DivideRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideRubyDarkAlpha(scale))
	return c
}

func (c *component) DivideSage(scale int) *component {
	c.el.Style(styles.DivideSage(scale))
	return c
}

func (c *component) DivideSageAlpha(scale int) *component {
	c.el.Style(styles.DivideSageAlpha(scale))
	return c
}

func (c *component) DivideSageDark(scale int) *component {
	c.el.Style(styles.DivideSageDark(scale))
	return c
}

func (c *component) DivideSageDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideSageDarkAlpha(scale))
	return c
}

func (c *component) DivideSand(scale int) *component {
	c.el.Style(styles.DivideSand(scale))
	return c
}

func (c *component) DivideSandAlpha(scale int) *component {
	c.el.Style(styles.DivideSandAlpha(scale))
	return c
}

func (c *component) DivideSandDark(scale int) *component {
	c.el.Style(styles.DivideSandDark(scale))
	return c
}

func (c *component) DivideSandDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideSandDarkAlpha(scale))
	return c
}

func (c *component) DivideSky(scale int) *component {
	c.el.Style(styles.DivideSky(scale))
	return c
}

func (c *component) DivideSkyAlpha(scale int) *component {
	c.el.Style(styles.DivideSkyAlpha(scale))
	return c
}

func (c *component) DivideSkyDark(scale int) *component {
	c.el.Style(styles.DivideSkyDark(scale))
	return c
}

func (c *component) DivideSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideSkyDarkAlpha(scale))
	return c
}

func (c *component) DivideSlate(scale int) *component {
	c.el.Style(styles.DivideSlate(scale))
	return c
}

func (c *component) DivideSlateAlpha(scale int) *component {
	c.el.Style(styles.DivideSlateAlpha(scale))
	return c
}

func (c *component) DivideSlateDark(scale int) *component {
	c.el.Style(styles.DivideSlateDark(scale))
	return c
}

func (c *component) DivideSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideSlateDarkAlpha(scale))
	return c
}

func (c *component) DivideSolid() *component {
	c.el.Style(styles.DivideSolid())
	return c
}

func (c *component) DivideTeal(scale int) *component {
	c.el.Style(styles.DivideTeal(scale))
	return c
}

func (c *component) DivideTealAlpha(scale int) *component {
	c.el.Style(styles.DivideTealAlpha(scale))
	return c
}

func (c *component) DivideTealDark(scale int) *component {
	c.el.Style(styles.DivideTealDark(scale))
	return c
}

func (c *component) DivideTealDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideTealDarkAlpha(scale))
	return c
}

func (c *component) DivideTomato(scale int) *component {
	c.el.Style(styles.DivideTomato(scale))
	return c
}

func (c *component) DivideTomatoAlpha(scale int) *component {
	c.el.Style(styles.DivideTomatoAlpha(scale))
	return c
}

func (c *component) DivideTomatoDark(scale int) *component {
	c.el.Style(styles.DivideTomatoDark(scale))
	return c
}

func (c *component) DivideTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideTomatoDarkAlpha(scale))
	return c
}

func (c *component) DivideTransparent() *component {
	c.el.Style(styles.DivideTransparent())
	return c
}

func (c *component) DivideViolet(scale int) *component {
	c.el.Style(styles.DivideViolet(scale))
	return c
}

func (c *component) DivideVioletAlpha(scale int) *component {
	c.el.Style(styles.DivideVioletAlpha(scale))
	return c
}

func (c *component) DivideVioletDark(scale int) *component {
	c.el.Style(styles.DivideVioletDark(scale))
	return c
}

func (c *component) DivideVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideVioletDarkAlpha(scale))
	return c
}

func (c *component) DivideWhite() *component {
	c.el.Style(styles.DivideWhite())
	return c
}

func (c *component) DivideWhiteAlpha(scale int) *component {
	c.el.Style(styles.DivideWhiteAlpha(scale))
	return c
}

func (c *component) DivideX(val ...value.Value) *component {
	c.el.Style(styles.DivideX(val...))
	return c
}

func (c *component) DivideXReverse(val ...value.Value) *component {
	c.el.Style(styles.DivideXReverse(val...))
	return c
}

func (c *component) DivideY(val ...value.Value) *component {
	c.el.Style(styles.DivideY(val...))
	return c
}

func (c *component) DivideYReverse(val ...value.Value) *component {
	c.el.Style(styles.DivideYReverse(val...))
	return c
}

func (c *component) DivideYellow(scale int) *component {
	c.el.Style(styles.DivideYellow(scale))
	return c
}

func (c *component) DivideYellowAlpha(scale int) *component {
	c.el.Style(styles.DivideYellowAlpha(scale))
	return c
}

func (c *component) DivideYellowDark(scale int) *component {
	c.el.Style(styles.DivideYellowDark(scale))
	return c
}

func (c *component) DivideYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.DivideYellowDarkAlpha(scale))
	return c
}

func (c *component) DropShadow(val value.Value) *component {
	c.el.Style(styles.DropShadow(val))
	return c
}

func (c *component) DropShadow2xl() *component {
	c.el.Style(styles.DropShadow2xl())
	return c
}

func (c *component) DropShadow2xs() *component {
	c.el.Style(styles.DropShadow2xs())
	return c
}

func (c *component) DropShadowAmber(scale int) *component {
	c.el.Style(styles.DropShadowAmber(scale))
	return c
}

func (c *component) DropShadowAmberAlpha(scale int) *component {
	c.el.Style(styles.DropShadowAmberAlpha(scale))
	return c
}

func (c *component) DropShadowAmberDark(scale int) *component {
	c.el.Style(styles.DropShadowAmberDark(scale))
	return c
}

func (c *component) DropShadowAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowAmberDarkAlpha(scale))
	return c
}

func (c *component) DropShadowBlack() *component {
	c.el.Style(styles.DropShadowBlack())
	return c
}

func (c *component) DropShadowBlackAlpha(scale int) *component {
	c.el.Style(styles.DropShadowBlackAlpha(scale))
	return c
}

func (c *component) DropShadowBlue(scale int) *component {
	c.el.Style(styles.DropShadowBlue(scale))
	return c
}

func (c *component) DropShadowBlueAlpha(scale int) *component {
	c.el.Style(styles.DropShadowBlueAlpha(scale))
	return c
}

func (c *component) DropShadowBlueDark(scale int) *component {
	c.el.Style(styles.DropShadowBlueDark(scale))
	return c
}

func (c *component) DropShadowBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowBlueDarkAlpha(scale))
	return c
}

func (c *component) DropShadowBronze(scale int) *component {
	c.el.Style(styles.DropShadowBronze(scale))
	return c
}

func (c *component) DropShadowBronzeAlpha(scale int) *component {
	c.el.Style(styles.DropShadowBronzeAlpha(scale))
	return c
}

func (c *component) DropShadowBronzeDark(scale int) *component {
	c.el.Style(styles.DropShadowBronzeDark(scale))
	return c
}

func (c *component) DropShadowBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowBronzeDarkAlpha(scale))
	return c
}

func (c *component) DropShadowBrown(scale int) *component {
	c.el.Style(styles.DropShadowBrown(scale))
	return c
}

func (c *component) DropShadowBrownAlpha(scale int) *component {
	c.el.Style(styles.DropShadowBrownAlpha(scale))
	return c
}

func (c *component) DropShadowBrownDark(scale int) *component {
	c.el.Style(styles.DropShadowBrownDark(scale))
	return c
}

func (c *component) DropShadowBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowBrownDarkAlpha(scale))
	return c
}

func (c *component) DropShadowColor(val value.Value) *component {
	c.el.Style(styles.DropShadowColor(val))
	return c
}

func (c *component) DropShadowCrimson(scale int) *component {
	c.el.Style(styles.DropShadowCrimson(scale))
	return c
}

func (c *component) DropShadowCrimsonAlpha(scale int) *component {
	c.el.Style(styles.DropShadowCrimsonAlpha(scale))
	return c
}

func (c *component) DropShadowCrimsonDark(scale int) *component {
	c.el.Style(styles.DropShadowCrimsonDark(scale))
	return c
}

func (c *component) DropShadowCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowCrimsonDarkAlpha(scale))
	return c
}

func (c *component) DropShadowCurrent() *component {
	c.el.Style(styles.DropShadowCurrent())
	return c
}

func (c *component) DropShadowCyan(scale int) *component {
	c.el.Style(styles.DropShadowCyan(scale))
	return c
}

func (c *component) DropShadowCyanAlpha(scale int) *component {
	c.el.Style(styles.DropShadowCyanAlpha(scale))
	return c
}

func (c *component) DropShadowCyanDark(scale int) *component {
	c.el.Style(styles.DropShadowCyanDark(scale))
	return c
}

func (c *component) DropShadowCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowCyanDarkAlpha(scale))
	return c
}

func (c *component) DropShadowGold(scale int) *component {
	c.el.Style(styles.DropShadowGold(scale))
	return c
}

func (c *component) DropShadowGoldAlpha(scale int) *component {
	c.el.Style(styles.DropShadowGoldAlpha(scale))
	return c
}

func (c *component) DropShadowGoldDark(scale int) *component {
	c.el.Style(styles.DropShadowGoldDark(scale))
	return c
}

func (c *component) DropShadowGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowGoldDarkAlpha(scale))
	return c
}

func (c *component) DropShadowGrass(scale int) *component {
	c.el.Style(styles.DropShadowGrass(scale))
	return c
}

func (c *component) DropShadowGrassAlpha(scale int) *component {
	c.el.Style(styles.DropShadowGrassAlpha(scale))
	return c
}

func (c *component) DropShadowGrassDark(scale int) *component {
	c.el.Style(styles.DropShadowGrassDark(scale))
	return c
}

func (c *component) DropShadowGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowGrassDarkAlpha(scale))
	return c
}

func (c *component) DropShadowGray(scale int) *component {
	c.el.Style(styles.DropShadowGray(scale))
	return c
}

func (c *component) DropShadowGrayAlpha(scale int) *component {
	c.el.Style(styles.DropShadowGrayAlpha(scale))
	return c
}

func (c *component) DropShadowGrayDark(scale int) *component {
	c.el.Style(styles.DropShadowGrayDark(scale))
	return c
}

func (c *component) DropShadowGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowGrayDarkAlpha(scale))
	return c
}

func (c *component) DropShadowGreen(scale int) *component {
	c.el.Style(styles.DropShadowGreen(scale))
	return c
}

func (c *component) DropShadowGreenAlpha(scale int) *component {
	c.el.Style(styles.DropShadowGreenAlpha(scale))
	return c
}

func (c *component) DropShadowGreenDark(scale int) *component {
	c.el.Style(styles.DropShadowGreenDark(scale))
	return c
}

func (c *component) DropShadowGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowGreenDarkAlpha(scale))
	return c
}

func (c *component) DropShadowIndigo(scale int) *component {
	c.el.Style(styles.DropShadowIndigo(scale))
	return c
}

func (c *component) DropShadowIndigoAlpha(scale int) *component {
	c.el.Style(styles.DropShadowIndigoAlpha(scale))
	return c
}

func (c *component) DropShadowIndigoDark(scale int) *component {
	c.el.Style(styles.DropShadowIndigoDark(scale))
	return c
}

func (c *component) DropShadowIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowIndigoDarkAlpha(scale))
	return c
}

func (c *component) DropShadowInherit() *component {
	c.el.Style(styles.DropShadowInherit())
	return c
}

func (c *component) DropShadowIris(scale int) *component {
	c.el.Style(styles.DropShadowIris(scale))
	return c
}

func (c *component) DropShadowIrisAlpha(scale int) *component {
	c.el.Style(styles.DropShadowIrisAlpha(scale))
	return c
}

func (c *component) DropShadowIrisDark(scale int) *component {
	c.el.Style(styles.DropShadowIrisDark(scale))
	return c
}

func (c *component) DropShadowIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowIrisDarkAlpha(scale))
	return c
}

func (c *component) DropShadowJade(scale int) *component {
	c.el.Style(styles.DropShadowJade(scale))
	return c
}

func (c *component) DropShadowJadeAlpha(scale int) *component {
	c.el.Style(styles.DropShadowJadeAlpha(scale))
	return c
}

func (c *component) DropShadowJadeDark(scale int) *component {
	c.el.Style(styles.DropShadowJadeDark(scale))
	return c
}

func (c *component) DropShadowJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowJadeDarkAlpha(scale))
	return c
}

func (c *component) DropShadowLg() *component {
	c.el.Style(styles.DropShadowLg())
	return c
}

func (c *component) DropShadowLime(scale int) *component {
	c.el.Style(styles.DropShadowLime(scale))
	return c
}

func (c *component) DropShadowLimeAlpha(scale int) *component {
	c.el.Style(styles.DropShadowLimeAlpha(scale))
	return c
}

func (c *component) DropShadowLimeDark(scale int) *component {
	c.el.Style(styles.DropShadowLimeDark(scale))
	return c
}

func (c *component) DropShadowLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowLimeDarkAlpha(scale))
	return c
}

func (c *component) DropShadowMauve(scale int) *component {
	c.el.Style(styles.DropShadowMauve(scale))
	return c
}

func (c *component) DropShadowMauveAlpha(scale int) *component {
	c.el.Style(styles.DropShadowMauveAlpha(scale))
	return c
}

func (c *component) DropShadowMauveDark(scale int) *component {
	c.el.Style(styles.DropShadowMauveDark(scale))
	return c
}

func (c *component) DropShadowMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowMauveDarkAlpha(scale))
	return c
}

func (c *component) DropShadowMd() *component {
	c.el.Style(styles.DropShadowMd())
	return c
}

func (c *component) DropShadowMint(scale int) *component {
	c.el.Style(styles.DropShadowMint(scale))
	return c
}

func (c *component) DropShadowMintAlpha(scale int) *component {
	c.el.Style(styles.DropShadowMintAlpha(scale))
	return c
}

func (c *component) DropShadowMintDark(scale int) *component {
	c.el.Style(styles.DropShadowMintDark(scale))
	return c
}

func (c *component) DropShadowMintDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowMintDarkAlpha(scale))
	return c
}

func (c *component) DropShadowNone() *component {
	c.el.Style(styles.DropShadowNone())
	return c
}

func (c *component) DropShadowOlive(scale int) *component {
	c.el.Style(styles.DropShadowOlive(scale))
	return c
}

func (c *component) DropShadowOliveAlpha(scale int) *component {
	c.el.Style(styles.DropShadowOliveAlpha(scale))
	return c
}

func (c *component) DropShadowOliveDark(scale int) *component {
	c.el.Style(styles.DropShadowOliveDark(scale))
	return c
}

func (c *component) DropShadowOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowOliveDarkAlpha(scale))
	return c
}

func (c *component) DropShadowOrange(scale int) *component {
	c.el.Style(styles.DropShadowOrange(scale))
	return c
}

func (c *component) DropShadowOrangeAlpha(scale int) *component {
	c.el.Style(styles.DropShadowOrangeAlpha(scale))
	return c
}

func (c *component) DropShadowOrangeDark(scale int) *component {
	c.el.Style(styles.DropShadowOrangeDark(scale))
	return c
}

func (c *component) DropShadowOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowOrangeDarkAlpha(scale))
	return c
}

func (c *component) DropShadowPink(scale int) *component {
	c.el.Style(styles.DropShadowPink(scale))
	return c
}

func (c *component) DropShadowPinkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowPinkAlpha(scale))
	return c
}

func (c *component) DropShadowPinkDark(scale int) *component {
	c.el.Style(styles.DropShadowPinkDark(scale))
	return c
}

func (c *component) DropShadowPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowPinkDarkAlpha(scale))
	return c
}

func (c *component) DropShadowPlum(scale int) *component {
	c.el.Style(styles.DropShadowPlum(scale))
	return c
}

func (c *component) DropShadowPlumAlpha(scale int) *component {
	c.el.Style(styles.DropShadowPlumAlpha(scale))
	return c
}

func (c *component) DropShadowPlumDark(scale int) *component {
	c.el.Style(styles.DropShadowPlumDark(scale))
	return c
}

func (c *component) DropShadowPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowPlumDarkAlpha(scale))
	return c
}

func (c *component) DropShadowPurple(scale int) *component {
	c.el.Style(styles.DropShadowPurple(scale))
	return c
}

func (c *component) DropShadowPurpleAlpha(scale int) *component {
	c.el.Style(styles.DropShadowPurpleAlpha(scale))
	return c
}

func (c *component) DropShadowPurpleDark(scale int) *component {
	c.el.Style(styles.DropShadowPurpleDark(scale))
	return c
}

func (c *component) DropShadowPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowPurpleDarkAlpha(scale))
	return c
}

func (c *component) DropShadowRed(scale int) *component {
	c.el.Style(styles.DropShadowRed(scale))
	return c
}

func (c *component) DropShadowRedAlpha(scale int) *component {
	c.el.Style(styles.DropShadowRedAlpha(scale))
	return c
}

func (c *component) DropShadowRedDark(scale int) *component {
	c.el.Style(styles.DropShadowRedDark(scale))
	return c
}

func (c *component) DropShadowRedDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowRedDarkAlpha(scale))
	return c
}

func (c *component) DropShadowRuby(scale int) *component {
	c.el.Style(styles.DropShadowRuby(scale))
	return c
}

func (c *component) DropShadowRubyAlpha(scale int) *component {
	c.el.Style(styles.DropShadowRubyAlpha(scale))
	return c
}

func (c *component) DropShadowRubyDark(scale int) *component {
	c.el.Style(styles.DropShadowRubyDark(scale))
	return c
}

func (c *component) DropShadowRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowRubyDarkAlpha(scale))
	return c
}

func (c *component) DropShadowSage(scale int) *component {
	c.el.Style(styles.DropShadowSage(scale))
	return c
}

func (c *component) DropShadowSageAlpha(scale int) *component {
	c.el.Style(styles.DropShadowSageAlpha(scale))
	return c
}

func (c *component) DropShadowSageDark(scale int) *component {
	c.el.Style(styles.DropShadowSageDark(scale))
	return c
}

func (c *component) DropShadowSageDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowSageDarkAlpha(scale))
	return c
}

func (c *component) DropShadowSand(scale int) *component {
	c.el.Style(styles.DropShadowSand(scale))
	return c
}

func (c *component) DropShadowSandAlpha(scale int) *component {
	c.el.Style(styles.DropShadowSandAlpha(scale))
	return c
}

func (c *component) DropShadowSandDark(scale int) *component {
	c.el.Style(styles.DropShadowSandDark(scale))
	return c
}

func (c *component) DropShadowSandDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowSandDarkAlpha(scale))
	return c
}

func (c *component) DropShadowSky(scale int) *component {
	c.el.Style(styles.DropShadowSky(scale))
	return c
}

func (c *component) DropShadowSkyAlpha(scale int) *component {
	c.el.Style(styles.DropShadowSkyAlpha(scale))
	return c
}

func (c *component) DropShadowSkyDark(scale int) *component {
	c.el.Style(styles.DropShadowSkyDark(scale))
	return c
}

func (c *component) DropShadowSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowSkyDarkAlpha(scale))
	return c
}

func (c *component) DropShadowSlate(scale int) *component {
	c.el.Style(styles.DropShadowSlate(scale))
	return c
}

func (c *component) DropShadowSlateAlpha(scale int) *component {
	c.el.Style(styles.DropShadowSlateAlpha(scale))
	return c
}

func (c *component) DropShadowSlateDark(scale int) *component {
	c.el.Style(styles.DropShadowSlateDark(scale))
	return c
}

func (c *component) DropShadowSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowSlateDarkAlpha(scale))
	return c
}

func (c *component) DropShadowSm() *component {
	c.el.Style(styles.DropShadowSm())
	return c
}

func (c *component) DropShadowTeal(scale int) *component {
	c.el.Style(styles.DropShadowTeal(scale))
	return c
}

func (c *component) DropShadowTealAlpha(scale int) *component {
	c.el.Style(styles.DropShadowTealAlpha(scale))
	return c
}

func (c *component) DropShadowTealDark(scale int) *component {
	c.el.Style(styles.DropShadowTealDark(scale))
	return c
}

func (c *component) DropShadowTealDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowTealDarkAlpha(scale))
	return c
}

func (c *component) DropShadowTomato(scale int) *component {
	c.el.Style(styles.DropShadowTomato(scale))
	return c
}

func (c *component) DropShadowTomatoAlpha(scale int) *component {
	c.el.Style(styles.DropShadowTomatoAlpha(scale))
	return c
}

func (c *component) DropShadowTomatoDark(scale int) *component {
	c.el.Style(styles.DropShadowTomatoDark(scale))
	return c
}

func (c *component) DropShadowTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowTomatoDarkAlpha(scale))
	return c
}

func (c *component) DropShadowTransparent() *component {
	c.el.Style(styles.DropShadowTransparent())
	return c
}

func (c *component) DropShadowViolet(scale int) *component {
	c.el.Style(styles.DropShadowViolet(scale))
	return c
}

func (c *component) DropShadowVioletAlpha(scale int) *component {
	c.el.Style(styles.DropShadowVioletAlpha(scale))
	return c
}

func (c *component) DropShadowVioletDark(scale int) *component {
	c.el.Style(styles.DropShadowVioletDark(scale))
	return c
}

func (c *component) DropShadowVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowVioletDarkAlpha(scale))
	return c
}

func (c *component) DropShadowWhite() *component {
	c.el.Style(styles.DropShadowWhite())
	return c
}

func (c *component) DropShadowWhiteAlpha(scale int) *component {
	c.el.Style(styles.DropShadowWhiteAlpha(scale))
	return c
}

func (c *component) DropShadowXl() *component {
	c.el.Style(styles.DropShadowXl())
	return c
}

func (c *component) DropShadowXs() *component {
	c.el.Style(styles.DropShadowXs())
	return c
}

func (c *component) DropShadowYellow(scale int) *component {
	c.el.Style(styles.DropShadowYellow(scale))
	return c
}

func (c *component) DropShadowYellowAlpha(scale int) *component {
	c.el.Style(styles.DropShadowYellowAlpha(scale))
	return c
}

func (c *component) DropShadowYellowDark(scale int) *component {
	c.el.Style(styles.DropShadowYellowDark(scale))
	return c
}

func (c *component) DropShadowYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.DropShadowYellowDarkAlpha(scale))
	return c
}

func (c *component) Duration(val any) *component {
	c.el.Style(styles.Duration(val))
	return c
}

func (c *component) DurationInitial() *component {
	c.el.Style(styles.DurationInitial())
	return c
}

func (c *component) Ease(val value.Value) *component {
	c.el.Style(styles.Ease(val))
	return c
}

func (c *component) EaseIn() *component {
	c.el.Style(styles.EaseIn())
	return c
}

func (c *component) EaseInOut() *component {
	c.el.Style(styles.EaseInOut())
	return c
}

func (c *component) EaseInitial() *component {
	c.el.Style(styles.EaseInitial())
	return c
}

func (c *component) EaseLinear() *component {
	c.el.Style(styles.EaseLinear())
	return c
}

func (c *component) EaseOut() *component {
	c.el.Style(styles.EaseOut())
	return c
}

func (c *component) End(number int) *component {
	c.el.Style(styles.End(number))
	return c
}

func (c *component) EndAuto() *component {
	c.el.Style(styles.EndAuto())
	return c
}

func (c *component) EndBy(val value.Value) *component {
	c.el.Style(styles.EndBy(val))
	return c
}

func (c *component) EndFraction(fraction float64) *component {
	c.el.Style(styles.EndFraction(fraction))
	return c
}

func (c *component) EndFull() *component {
	c.el.Style(styles.EndFull())
	return c
}

func (c *component) EndPx() *component {
	c.el.Style(styles.EndPx())
	return c
}

func (c *component) FieldSizingContent() *component {
	c.el.Style(styles.FieldSizingContent())
	return c
}

func (c *component) FieldSizingFixed() *component {
	c.el.Style(styles.FieldSizingFixed())
	return c
}

func (c *component) Fill(val value.Value) *component {
	c.el.Style(styles.Fill(val))
	return c
}

func (c *component) FillAmber(scale int) *component {
	c.el.Style(styles.FillAmber(scale))
	return c
}

func (c *component) FillAmberAlpha(scale int) *component {
	c.el.Style(styles.FillAmberAlpha(scale))
	return c
}

func (c *component) FillAmberDark(scale int) *component {
	c.el.Style(styles.FillAmberDark(scale))
	return c
}

func (c *component) FillAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.FillAmberDarkAlpha(scale))
	return c
}

func (c *component) FillBlack() *component {
	c.el.Style(styles.FillBlack())
	return c
}

func (c *component) FillBlackAlpha(scale int) *component {
	c.el.Style(styles.FillBlackAlpha(scale))
	return c
}

func (c *component) FillBlue(scale int) *component {
	c.el.Style(styles.FillBlue(scale))
	return c
}

func (c *component) FillBlueAlpha(scale int) *component {
	c.el.Style(styles.FillBlueAlpha(scale))
	return c
}

func (c *component) FillBlueDark(scale int) *component {
	c.el.Style(styles.FillBlueDark(scale))
	return c
}

func (c *component) FillBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.FillBlueDarkAlpha(scale))
	return c
}

func (c *component) FillBronze(scale int) *component {
	c.el.Style(styles.FillBronze(scale))
	return c
}

func (c *component) FillBronzeAlpha(scale int) *component {
	c.el.Style(styles.FillBronzeAlpha(scale))
	return c
}

func (c *component) FillBronzeDark(scale int) *component {
	c.el.Style(styles.FillBronzeDark(scale))
	return c
}

func (c *component) FillBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.FillBronzeDarkAlpha(scale))
	return c
}

func (c *component) FillBrown(scale int) *component {
	c.el.Style(styles.FillBrown(scale))
	return c
}

func (c *component) FillBrownAlpha(scale int) *component {
	c.el.Style(styles.FillBrownAlpha(scale))
	return c
}

func (c *component) FillBrownDark(scale int) *component {
	c.el.Style(styles.FillBrownDark(scale))
	return c
}

func (c *component) FillBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.FillBrownDarkAlpha(scale))
	return c
}

func (c *component) FillCrimson(scale int) *component {
	c.el.Style(styles.FillCrimson(scale))
	return c
}

func (c *component) FillCrimsonAlpha(scale int) *component {
	c.el.Style(styles.FillCrimsonAlpha(scale))
	return c
}

func (c *component) FillCrimsonDark(scale int) *component {
	c.el.Style(styles.FillCrimsonDark(scale))
	return c
}

func (c *component) FillCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.FillCrimsonDarkAlpha(scale))
	return c
}

func (c *component) FillCurrent() *component {
	c.el.Style(styles.FillCurrent())
	return c
}

func (c *component) FillCyan(scale int) *component {
	c.el.Style(styles.FillCyan(scale))
	return c
}

func (c *component) FillCyanAlpha(scale int) *component {
	c.el.Style(styles.FillCyanAlpha(scale))
	return c
}

func (c *component) FillCyanDark(scale int) *component {
	c.el.Style(styles.FillCyanDark(scale))
	return c
}

func (c *component) FillCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.FillCyanDarkAlpha(scale))
	return c
}

func (c *component) FillGold(scale int) *component {
	c.el.Style(styles.FillGold(scale))
	return c
}

func (c *component) FillGoldAlpha(scale int) *component {
	c.el.Style(styles.FillGoldAlpha(scale))
	return c
}

func (c *component) FillGoldDark(scale int) *component {
	c.el.Style(styles.FillGoldDark(scale))
	return c
}

func (c *component) FillGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.FillGoldDarkAlpha(scale))
	return c
}

func (c *component) FillGrass(scale int) *component {
	c.el.Style(styles.FillGrass(scale))
	return c
}

func (c *component) FillGrassAlpha(scale int) *component {
	c.el.Style(styles.FillGrassAlpha(scale))
	return c
}

func (c *component) FillGrassDark(scale int) *component {
	c.el.Style(styles.FillGrassDark(scale))
	return c
}

func (c *component) FillGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.FillGrassDarkAlpha(scale))
	return c
}

func (c *component) FillGray(scale int) *component {
	c.el.Style(styles.FillGray(scale))
	return c
}

func (c *component) FillGrayAlpha(scale int) *component {
	c.el.Style(styles.FillGrayAlpha(scale))
	return c
}

func (c *component) FillGrayDark(scale int) *component {
	c.el.Style(styles.FillGrayDark(scale))
	return c
}

func (c *component) FillGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.FillGrayDarkAlpha(scale))
	return c
}

func (c *component) FillGreen(scale int) *component {
	c.el.Style(styles.FillGreen(scale))
	return c
}

func (c *component) FillGreenAlpha(scale int) *component {
	c.el.Style(styles.FillGreenAlpha(scale))
	return c
}

func (c *component) FillGreenDark(scale int) *component {
	c.el.Style(styles.FillGreenDark(scale))
	return c
}

func (c *component) FillGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.FillGreenDarkAlpha(scale))
	return c
}

func (c *component) FillIndigo(scale int) *component {
	c.el.Style(styles.FillIndigo(scale))
	return c
}

func (c *component) FillIndigoAlpha(scale int) *component {
	c.el.Style(styles.FillIndigoAlpha(scale))
	return c
}

func (c *component) FillIndigoDark(scale int) *component {
	c.el.Style(styles.FillIndigoDark(scale))
	return c
}

func (c *component) FillIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.FillIndigoDarkAlpha(scale))
	return c
}

func (c *component) FillInherit() *component {
	c.el.Style(styles.FillInherit())
	return c
}

func (c *component) FillIris(scale int) *component {
	c.el.Style(styles.FillIris(scale))
	return c
}

func (c *component) FillIrisAlpha(scale int) *component {
	c.el.Style(styles.FillIrisAlpha(scale))
	return c
}

func (c *component) FillIrisDark(scale int) *component {
	c.el.Style(styles.FillIrisDark(scale))
	return c
}

func (c *component) FillIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.FillIrisDarkAlpha(scale))
	return c
}

func (c *component) FillJade(scale int) *component {
	c.el.Style(styles.FillJade(scale))
	return c
}

func (c *component) FillJadeAlpha(scale int) *component {
	c.el.Style(styles.FillJadeAlpha(scale))
	return c
}

func (c *component) FillJadeDark(scale int) *component {
	c.el.Style(styles.FillJadeDark(scale))
	return c
}

func (c *component) FillJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.FillJadeDarkAlpha(scale))
	return c
}

func (c *component) FillLime(scale int) *component {
	c.el.Style(styles.FillLime(scale))
	return c
}

func (c *component) FillLimeAlpha(scale int) *component {
	c.el.Style(styles.FillLimeAlpha(scale))
	return c
}

func (c *component) FillLimeDark(scale int) *component {
	c.el.Style(styles.FillLimeDark(scale))
	return c
}

func (c *component) FillLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.FillLimeDarkAlpha(scale))
	return c
}

func (c *component) FillMauve(scale int) *component {
	c.el.Style(styles.FillMauve(scale))
	return c
}

func (c *component) FillMauveAlpha(scale int) *component {
	c.el.Style(styles.FillMauveAlpha(scale))
	return c
}

func (c *component) FillMauveDark(scale int) *component {
	c.el.Style(styles.FillMauveDark(scale))
	return c
}

func (c *component) FillMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.FillMauveDarkAlpha(scale))
	return c
}

func (c *component) FillMint(scale int) *component {
	c.el.Style(styles.FillMint(scale))
	return c
}

func (c *component) FillMintAlpha(scale int) *component {
	c.el.Style(styles.FillMintAlpha(scale))
	return c
}

func (c *component) FillMintDark(scale int) *component {
	c.el.Style(styles.FillMintDark(scale))
	return c
}

func (c *component) FillMintDarkAlpha(scale int) *component {
	c.el.Style(styles.FillMintDarkAlpha(scale))
	return c
}

func (c *component) FillOlive(scale int) *component {
	c.el.Style(styles.FillOlive(scale))
	return c
}

func (c *component) FillOliveAlpha(scale int) *component {
	c.el.Style(styles.FillOliveAlpha(scale))
	return c
}

func (c *component) FillOliveDark(scale int) *component {
	c.el.Style(styles.FillOliveDark(scale))
	return c
}

func (c *component) FillOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.FillOliveDarkAlpha(scale))
	return c
}

func (c *component) FillOrange(scale int) *component {
	c.el.Style(styles.FillOrange(scale))
	return c
}

func (c *component) FillOrangeAlpha(scale int) *component {
	c.el.Style(styles.FillOrangeAlpha(scale))
	return c
}

func (c *component) FillOrangeDark(scale int) *component {
	c.el.Style(styles.FillOrangeDark(scale))
	return c
}

func (c *component) FillOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.FillOrangeDarkAlpha(scale))
	return c
}

func (c *component) FillPink(scale int) *component {
	c.el.Style(styles.FillPink(scale))
	return c
}

func (c *component) FillPinkAlpha(scale int) *component {
	c.el.Style(styles.FillPinkAlpha(scale))
	return c
}

func (c *component) FillPinkDark(scale int) *component {
	c.el.Style(styles.FillPinkDark(scale))
	return c
}

func (c *component) FillPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.FillPinkDarkAlpha(scale))
	return c
}

func (c *component) FillPlum(scale int) *component {
	c.el.Style(styles.FillPlum(scale))
	return c
}

func (c *component) FillPlumAlpha(scale int) *component {
	c.el.Style(styles.FillPlumAlpha(scale))
	return c
}

func (c *component) FillPlumDark(scale int) *component {
	c.el.Style(styles.FillPlumDark(scale))
	return c
}

func (c *component) FillPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.FillPlumDarkAlpha(scale))
	return c
}

func (c *component) FillPurple(scale int) *component {
	c.el.Style(styles.FillPurple(scale))
	return c
}

func (c *component) FillPurpleAlpha(scale int) *component {
	c.el.Style(styles.FillPurpleAlpha(scale))
	return c
}

func (c *component) FillPurpleDark(scale int) *component {
	c.el.Style(styles.FillPurpleDark(scale))
	return c
}

func (c *component) FillPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.FillPurpleDarkAlpha(scale))
	return c
}

func (c *component) FillRed(scale int) *component {
	c.el.Style(styles.FillRed(scale))
	return c
}

func (c *component) FillRedAlpha(scale int) *component {
	c.el.Style(styles.FillRedAlpha(scale))
	return c
}

func (c *component) FillRedDark(scale int) *component {
	c.el.Style(styles.FillRedDark(scale))
	return c
}

func (c *component) FillRedDarkAlpha(scale int) *component {
	c.el.Style(styles.FillRedDarkAlpha(scale))
	return c
}

func (c *component) FillRuby(scale int) *component {
	c.el.Style(styles.FillRuby(scale))
	return c
}

func (c *component) FillRubyAlpha(scale int) *component {
	c.el.Style(styles.FillRubyAlpha(scale))
	return c
}

func (c *component) FillRubyDark(scale int) *component {
	c.el.Style(styles.FillRubyDark(scale))
	return c
}

func (c *component) FillRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.FillRubyDarkAlpha(scale))
	return c
}

func (c *component) FillSage(scale int) *component {
	c.el.Style(styles.FillSage(scale))
	return c
}

func (c *component) FillSageAlpha(scale int) *component {
	c.el.Style(styles.FillSageAlpha(scale))
	return c
}

func (c *component) FillSageDark(scale int) *component {
	c.el.Style(styles.FillSageDark(scale))
	return c
}

func (c *component) FillSageDarkAlpha(scale int) *component {
	c.el.Style(styles.FillSageDarkAlpha(scale))
	return c
}

func (c *component) FillSand(scale int) *component {
	c.el.Style(styles.FillSand(scale))
	return c
}

func (c *component) FillSandAlpha(scale int) *component {
	c.el.Style(styles.FillSandAlpha(scale))
	return c
}

func (c *component) FillSandDark(scale int) *component {
	c.el.Style(styles.FillSandDark(scale))
	return c
}

func (c *component) FillSandDarkAlpha(scale int) *component {
	c.el.Style(styles.FillSandDarkAlpha(scale))
	return c
}

func (c *component) FillSky(scale int) *component {
	c.el.Style(styles.FillSky(scale))
	return c
}

func (c *component) FillSkyAlpha(scale int) *component {
	c.el.Style(styles.FillSkyAlpha(scale))
	return c
}

func (c *component) FillSkyDark(scale int) *component {
	c.el.Style(styles.FillSkyDark(scale))
	return c
}

func (c *component) FillSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.FillSkyDarkAlpha(scale))
	return c
}

func (c *component) FillSlate(scale int) *component {
	c.el.Style(styles.FillSlate(scale))
	return c
}

func (c *component) FillSlateAlpha(scale int) *component {
	c.el.Style(styles.FillSlateAlpha(scale))
	return c
}

func (c *component) FillSlateDark(scale int) *component {
	c.el.Style(styles.FillSlateDark(scale))
	return c
}

func (c *component) FillSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.FillSlateDarkAlpha(scale))
	return c
}

func (c *component) FillTeal(scale int) *component {
	c.el.Style(styles.FillTeal(scale))
	return c
}

func (c *component) FillTealAlpha(scale int) *component {
	c.el.Style(styles.FillTealAlpha(scale))
	return c
}

func (c *component) FillTealDark(scale int) *component {
	c.el.Style(styles.FillTealDark(scale))
	return c
}

func (c *component) FillTealDarkAlpha(scale int) *component {
	c.el.Style(styles.FillTealDarkAlpha(scale))
	return c
}

func (c *component) FillTomato(scale int) *component {
	c.el.Style(styles.FillTomato(scale))
	return c
}

func (c *component) FillTomatoAlpha(scale int) *component {
	c.el.Style(styles.FillTomatoAlpha(scale))
	return c
}

func (c *component) FillTomatoDark(scale int) *component {
	c.el.Style(styles.FillTomatoDark(scale))
	return c
}

func (c *component) FillTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.FillTomatoDarkAlpha(scale))
	return c
}

func (c *component) FillTransparent() *component {
	c.el.Style(styles.FillTransparent())
	return c
}

func (c *component) FillViolet(scale int) *component {
	c.el.Style(styles.FillViolet(scale))
	return c
}

func (c *component) FillVioletAlpha(scale int) *component {
	c.el.Style(styles.FillVioletAlpha(scale))
	return c
}

func (c *component) FillVioletDark(scale int) *component {
	c.el.Style(styles.FillVioletDark(scale))
	return c
}

func (c *component) FillVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.FillVioletDarkAlpha(scale))
	return c
}

func (c *component) FillWhite() *component {
	c.el.Style(styles.FillWhite())
	return c
}

func (c *component) FillWhiteAlpha(scale int) *component {
	c.el.Style(styles.FillWhiteAlpha(scale))
	return c
}

func (c *component) FillYellow(scale int) *component {
	c.el.Style(styles.FillYellow(scale))
	return c
}

func (c *component) FillYellowAlpha(scale int) *component {
	c.el.Style(styles.FillYellowAlpha(scale))
	return c
}

func (c *component) FillYellowDark(scale int) *component {
	c.el.Style(styles.FillYellowDark(scale))
	return c
}

func (c *component) FillYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.FillYellowDarkAlpha(scale))
	return c
}

func (c *component) Filter(val value.Value) *component {
	c.el.Style(styles.Filter(val))
	return c
}

func (c *component) FilterNone() *component {
	c.el.Style(styles.FilterNone())
	return c
}

func (c *component) Fixed() *component {
	c.el.Style(styles.Fixed())
	return c
}

func (c *component) Flex() *component {
	c.el.Style(styles.Flex())
	return c
}

func (c *component) FlexAuto() *component {
	c.el.Style(styles.FlexAuto())
	return c
}

func (c *component) FlexBy(val value.Value) *component {
	c.el.Style(styles.FlexBy(val))
	return c
}

func (c *component) FlexCol() *component {
	c.el.Style(styles.FlexCol())
	return c
}

func (c *component) FlexColReverse() *component {
	c.el.Style(styles.FlexColReverse())
	return c
}

func (c *component) FlexFraction(fraction float64) *component {
	c.el.Style(styles.FlexFraction(fraction))
	return c
}

func (c *component) FlexInitial() *component {
	c.el.Style(styles.FlexInitial())
	return c
}

func (c *component) FlexNoWrap() *component {
	c.el.Style(styles.FlexNoWrap())
	return c
}

func (c *component) FlexNone() *component {
	c.el.Style(styles.FlexNone())
	return c
}

func (c *component) FlexRow() *component {
	c.el.Style(styles.FlexRow())
	return c
}

func (c *component) FlexRowReverse() *component {
	c.el.Style(styles.FlexRowReverse())
	return c
}

func (c *component) FlexWrap() *component {
	c.el.Style(styles.FlexWrap())
	return c
}

func (c *component) FlexWrapReverse() *component {
	c.el.Style(styles.FlexWrapReverse())
	return c
}

func (c *component) FloatEnd() *component {
	c.el.Style(styles.FloatEnd())
	return c
}

func (c *component) FloatLeft() *component {
	c.el.Style(styles.FloatLeft())
	return c
}

func (c *component) FloatNone() *component {
	c.el.Style(styles.FloatNone())
	return c
}

func (c *component) FloatRight() *component {
	c.el.Style(styles.FloatRight())
	return c
}

func (c *component) FloatStart() *component {
	c.el.Style(styles.FloatStart())
	return c
}

func (c *component) FlowRoot() *component {
	c.el.Style(styles.FlowRoot())
	return c
}

func (c *component) FontBlack() *component {
	c.el.Style(styles.FontBlack())
	return c
}

func (c *component) FontBold() *component {
	c.el.Style(styles.FontBold())
	return c
}

func (c *component) FontExtraBold() *component {
	c.el.Style(styles.FontExtraBold())
	return c
}

func (c *component) FontExtraLight() *component {
	c.el.Style(styles.FontExtraLight())
	return c
}

func (c *component) FontFamilyBy(val value.Value) *component {
	c.el.Style(styles.FontFamilyBy(val))
	return c
}

func (c *component) FontLight() *component {
	c.el.Style(styles.FontLight())
	return c
}

func (c *component) FontMedium() *component {
	c.el.Style(styles.FontMedium())
	return c
}

func (c *component) FontMono() *component {
	c.el.Style(styles.FontMono())
	return c
}

func (c *component) FontNormal() *component {
	c.el.Style(styles.FontNormal())
	return c
}

func (c *component) FontSans() *component {
	c.el.Style(styles.FontSans())
	return c
}

func (c *component) FontSemibold() *component {
	c.el.Style(styles.FontSemibold())
	return c
}

func (c *component) FontSerif() *component {
	c.el.Style(styles.FontSerif())
	return c
}

func (c *component) FontStretchBy(val value.Value) *component {
	c.el.Style(styles.FontStretchBy(val))
	return c
}

func (c *component) FontStretchCondensed() *component {
	c.el.Style(styles.FontStretchCondensed())
	return c
}

func (c *component) FontStretchExpanded() *component {
	c.el.Style(styles.FontStretchExpanded())
	return c
}

func (c *component) FontStretchExtraCondensed() *component {
	c.el.Style(styles.FontStretchExtraCondensed())
	return c
}

func (c *component) FontStretchExtraExpanded() *component {
	c.el.Style(styles.FontStretchExtraExpanded())
	return c
}

func (c *component) FontStretchNormal() *component {
	c.el.Style(styles.FontStretchNormal())
	return c
}

func (c *component) FontStretchPercentage(percentage int) *component {
	c.el.Style(styles.FontStretchPercentage(percentage))
	return c
}

func (c *component) FontStretchSemiCondensed() *component {
	c.el.Style(styles.FontStretchSemiCondensed())
	return c
}

func (c *component) FontStretchSemiExpanded() *component {
	c.el.Style(styles.FontStretchSemiExpanded())
	return c
}

func (c *component) FontStretchUltraCondensed() *component {
	c.el.Style(styles.FontStretchUltraCondensed())
	return c
}

func (c *component) FontStretchUltraExpanded() *component {
	c.el.Style(styles.FontStretchUltraExpanded())
	return c
}

func (c *component) FontThin() *component {
	c.el.Style(styles.FontThin())
	return c
}

func (c *component) FontWeightBy(val value.Value) *component {
	c.el.Style(styles.FontWeightBy(val))
	return c
}

func (c *component) From(val value.Value) *component {
	c.el.Style(styles.From(val))
	return c
}

func (c *component) Gap(number int) *component {
	c.el.Style(styles.Gap(number))
	return c
}

func (c *component) GapBy(val value.Value) *component {
	c.el.Style(styles.GapBy(val))
	return c
}

func (c *component) GapX(number int) *component {
	c.el.Style(styles.GapX(number))
	return c
}

func (c *component) GapXBy(val value.Value) *component {
	c.el.Style(styles.GapXBy(val))
	return c
}

func (c *component) GapY(number int) *component {
	c.el.Style(styles.GapY(number))
	return c
}

func (c *component) GapYBy(val value.Value) *component {
	c.el.Style(styles.GapYBy(val))
	return c
}

func (c *component) Grayscale(val ...any) *component {
	c.el.Style(styles.Grayscale(val...))
	return c
}

func (c *component) Grid() *component {
	c.el.Style(styles.Grid())
	return c
}

func (c *component) GridCols(number int) *component {
	c.el.Style(styles.GridCols(number))
	return c
}

func (c *component) GridColsBy(val value.Value) *component {
	c.el.Style(styles.GridColsBy(val))
	return c
}

func (c *component) GridColsNone() *component {
	c.el.Style(styles.GridColsNone())
	return c
}

func (c *component) GridColsSubgrid() *component {
	c.el.Style(styles.GridColsSubgrid())
	return c
}

func (c *component) GridFlowCol() *component {
	c.el.Style(styles.GridFlowCol())
	return c
}

func (c *component) GridFlowColDense() *component {
	c.el.Style(styles.GridFlowColDense())
	return c
}

func (c *component) GridFlowDense() *component {
	c.el.Style(styles.GridFlowDense())
	return c
}

func (c *component) GridFlowRow() *component {
	c.el.Style(styles.GridFlowRow())
	return c
}

func (c *component) GridFlowRowDense() *component {
	c.el.Style(styles.GridFlowRowDense())
	return c
}

func (c *component) GridRows(number int) *component {
	c.el.Style(styles.GridRows(number))
	return c
}

func (c *component) GridRowsBy(val value.Value) *component {
	c.el.Style(styles.GridRowsBy(val))
	return c
}

func (c *component) GridRowsNone() *component {
	c.el.Style(styles.GridRowsNone())
	return c
}

func (c *component) GridRowsSubgrid() *component {
	c.el.Style(styles.GridRowsSubgrid())
	return c
}

func (c *component) Grow() *component {
	c.el.Style(styles.Grow())
	return c
}

func (c *component) GrowNumber(number int) *component {
	c.el.Style(styles.GrowNumber(number))
	return c
}

func (c *component) GrowValue(val value.Value) *component {
	c.el.Style(styles.GrowValue(val))
	return c
}

func (c *component) H(number int) *component {
	c.el.Style(styles.H(number))
	return c
}

func (c *component) HAuto() *component {
	c.el.Style(styles.HAuto())
	return c
}

func (c *component) HDvh() *component {
	c.el.Style(styles.HDvh())
	return c
}

func (c *component) HDvw() *component {
	c.el.Style(styles.HDvw())
	return c
}

func (c *component) HFit() *component {
	c.el.Style(styles.HFit())
	return c
}

func (c *component) HFraction(fraction float32) *component {
	c.el.Style(styles.HFraction(fraction))
	return c
}

func (c *component) HFull() *component {
	c.el.Style(styles.HFull())
	return c
}

func (c *component) HLh() *component {
	c.el.Style(styles.HLh())
	return c
}

func (c *component) HLvh() *component {
	c.el.Style(styles.HLvh())
	return c
}

func (c *component) HLvw() *component {
	c.el.Style(styles.HLvw())
	return c
}

func (c *component) HMax() *component {
	c.el.Style(styles.HMax())
	return c
}

func (c *component) HMin() *component {
	c.el.Style(styles.HMin())
	return c
}

func (c *component) HPx() *component {
	c.el.Style(styles.HPx())
	return c
}

func (c *component) HScreen() *component {
	c.el.Style(styles.HScreen())
	return c
}

func (c *component) HSvh() *component {
	c.el.Style(styles.HSvh())
	return c
}

func (c *component) HSvw() *component {
	c.el.Style(styles.HSvw())
	return c
}

func (c *component) HValueBy(val value.Value) *component {
	c.el.Style(styles.HValueBy(val))
	return c
}

func (c *component) Hidden() *component {
	c.el.Style(styles.Hidden())
	return c
}

func (c *component) HueRotate(val any) *component {
	c.el.Style(styles.HueRotate(val))
	return c
}

func (c *component) HyphensAuto() *component {
	c.el.Style(styles.HyphensAuto())
	return c
}

func (c *component) HyphensManual() *component {
	c.el.Style(styles.HyphensManual())
	return c
}

func (c *component) HyphensNone() *component {
	c.el.Style(styles.HyphensNone())
	return c
}

func (c *component) Indent(number int) *component {
	c.el.Style(styles.Indent(number))
	return c
}

func (c *component) IndentBy(val value.Value) *component {
	c.el.Style(styles.IndentBy(val))
	return c
}

func (c *component) IndentPx() *component {
	c.el.Style(styles.IndentPx())
	return c
}

func (c *component) Inline() *component {
	c.el.Style(styles.Inline())
	return c
}

func (c *component) InlineBlock() *component {
	c.el.Style(styles.InlineBlock())
	return c
}

func (c *component) InlineFlex() *component {
	c.el.Style(styles.InlineFlex())
	return c
}

func (c *component) InlineGrid() *component {
	c.el.Style(styles.InlineGrid())
	return c
}

func (c *component) InlineTable() *component {
	c.el.Style(styles.InlineTable())
	return c
}

func (c *component) Inset(number int) *component {
	c.el.Style(styles.Inset(number))
	return c
}

func (c *component) InsetAuto() *component {
	c.el.Style(styles.InsetAuto())
	return c
}

func (c *component) InsetBy(val value.Value) *component {
	c.el.Style(styles.InsetBy(val))
	return c
}

func (c *component) InsetFraction(fraction float64) *component {
	c.el.Style(styles.InsetFraction(fraction))
	return c
}

func (c *component) InsetFull() *component {
	c.el.Style(styles.InsetFull())
	return c
}

func (c *component) InsetPx() *component {
	c.el.Style(styles.InsetPx())
	return c
}

func (c *component) InsetRing(val ...value.Value) *component {
	c.el.Style(styles.InsetRing(val...))
	return c
}

func (c *component) InsetRingAmber(scale int) *component {
	c.el.Style(styles.InsetRingAmber(scale))
	return c
}

func (c *component) InsetRingAmberAlpha(scale int) *component {
	c.el.Style(styles.InsetRingAmberAlpha(scale))
	return c
}

func (c *component) InsetRingAmberDark(scale int) *component {
	c.el.Style(styles.InsetRingAmberDark(scale))
	return c
}

func (c *component) InsetRingAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingAmberDarkAlpha(scale))
	return c
}

func (c *component) InsetRingBlack() *component {
	c.el.Style(styles.InsetRingBlack())
	return c
}

func (c *component) InsetRingBlackAlpha(scale int) *component {
	c.el.Style(styles.InsetRingBlackAlpha(scale))
	return c
}

func (c *component) InsetRingBlue(scale int) *component {
	c.el.Style(styles.InsetRingBlue(scale))
	return c
}

func (c *component) InsetRingBlueAlpha(scale int) *component {
	c.el.Style(styles.InsetRingBlueAlpha(scale))
	return c
}

func (c *component) InsetRingBlueDark(scale int) *component {
	c.el.Style(styles.InsetRingBlueDark(scale))
	return c
}

func (c *component) InsetRingBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingBlueDarkAlpha(scale))
	return c
}

func (c *component) InsetRingBronze(scale int) *component {
	c.el.Style(styles.InsetRingBronze(scale))
	return c
}

func (c *component) InsetRingBronzeAlpha(scale int) *component {
	c.el.Style(styles.InsetRingBronzeAlpha(scale))
	return c
}

func (c *component) InsetRingBronzeDark(scale int) *component {
	c.el.Style(styles.InsetRingBronzeDark(scale))
	return c
}

func (c *component) InsetRingBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingBronzeDarkAlpha(scale))
	return c
}

func (c *component) InsetRingBrown(scale int) *component {
	c.el.Style(styles.InsetRingBrown(scale))
	return c
}

func (c *component) InsetRingBrownAlpha(scale int) *component {
	c.el.Style(styles.InsetRingBrownAlpha(scale))
	return c
}

func (c *component) InsetRingBrownDark(scale int) *component {
	c.el.Style(styles.InsetRingBrownDark(scale))
	return c
}

func (c *component) InsetRingBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingBrownDarkAlpha(scale))
	return c
}

func (c *component) InsetRingColor(val value.Value) *component {
	c.el.Style(styles.InsetRingColor(val))
	return c
}

func (c *component) InsetRingCrimson(scale int) *component {
	c.el.Style(styles.InsetRingCrimson(scale))
	return c
}

func (c *component) InsetRingCrimsonAlpha(scale int) *component {
	c.el.Style(styles.InsetRingCrimsonAlpha(scale))
	return c
}

func (c *component) InsetRingCrimsonDark(scale int) *component {
	c.el.Style(styles.InsetRingCrimsonDark(scale))
	return c
}

func (c *component) InsetRingCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingCrimsonDarkAlpha(scale))
	return c
}

func (c *component) InsetRingCurrent() *component {
	c.el.Style(styles.InsetRingCurrent())
	return c
}

func (c *component) InsetRingCyan(scale int) *component {
	c.el.Style(styles.InsetRingCyan(scale))
	return c
}

func (c *component) InsetRingCyanAlpha(scale int) *component {
	c.el.Style(styles.InsetRingCyanAlpha(scale))
	return c
}

func (c *component) InsetRingCyanDark(scale int) *component {
	c.el.Style(styles.InsetRingCyanDark(scale))
	return c
}

func (c *component) InsetRingCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingCyanDarkAlpha(scale))
	return c
}

func (c *component) InsetRingGold(scale int) *component {
	c.el.Style(styles.InsetRingGold(scale))
	return c
}

func (c *component) InsetRingGoldAlpha(scale int) *component {
	c.el.Style(styles.InsetRingGoldAlpha(scale))
	return c
}

func (c *component) InsetRingGoldDark(scale int) *component {
	c.el.Style(styles.InsetRingGoldDark(scale))
	return c
}

func (c *component) InsetRingGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingGoldDarkAlpha(scale))
	return c
}

func (c *component) InsetRingGrass(scale int) *component {
	c.el.Style(styles.InsetRingGrass(scale))
	return c
}

func (c *component) InsetRingGrassAlpha(scale int) *component {
	c.el.Style(styles.InsetRingGrassAlpha(scale))
	return c
}

func (c *component) InsetRingGrassDark(scale int) *component {
	c.el.Style(styles.InsetRingGrassDark(scale))
	return c
}

func (c *component) InsetRingGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingGrassDarkAlpha(scale))
	return c
}

func (c *component) InsetRingGray(scale int) *component {
	c.el.Style(styles.InsetRingGray(scale))
	return c
}

func (c *component) InsetRingGrayAlpha(scale int) *component {
	c.el.Style(styles.InsetRingGrayAlpha(scale))
	return c
}

func (c *component) InsetRingGrayDark(scale int) *component {
	c.el.Style(styles.InsetRingGrayDark(scale))
	return c
}

func (c *component) InsetRingGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingGrayDarkAlpha(scale))
	return c
}

func (c *component) InsetRingGreen(scale int) *component {
	c.el.Style(styles.InsetRingGreen(scale))
	return c
}

func (c *component) InsetRingGreenAlpha(scale int) *component {
	c.el.Style(styles.InsetRingGreenAlpha(scale))
	return c
}

func (c *component) InsetRingGreenDark(scale int) *component {
	c.el.Style(styles.InsetRingGreenDark(scale))
	return c
}

func (c *component) InsetRingGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingGreenDarkAlpha(scale))
	return c
}

func (c *component) InsetRingIndigo(scale int) *component {
	c.el.Style(styles.InsetRingIndigo(scale))
	return c
}

func (c *component) InsetRingIndigoAlpha(scale int) *component {
	c.el.Style(styles.InsetRingIndigoAlpha(scale))
	return c
}

func (c *component) InsetRingIndigoDark(scale int) *component {
	c.el.Style(styles.InsetRingIndigoDark(scale))
	return c
}

func (c *component) InsetRingIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingIndigoDarkAlpha(scale))
	return c
}

func (c *component) InsetRingInherit() *component {
	c.el.Style(styles.InsetRingInherit())
	return c
}

func (c *component) InsetRingIris(scale int) *component {
	c.el.Style(styles.InsetRingIris(scale))
	return c
}

func (c *component) InsetRingIrisAlpha(scale int) *component {
	c.el.Style(styles.InsetRingIrisAlpha(scale))
	return c
}

func (c *component) InsetRingIrisDark(scale int) *component {
	c.el.Style(styles.InsetRingIrisDark(scale))
	return c
}

func (c *component) InsetRingIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingIrisDarkAlpha(scale))
	return c
}

func (c *component) InsetRingJade(scale int) *component {
	c.el.Style(styles.InsetRingJade(scale))
	return c
}

func (c *component) InsetRingJadeAlpha(scale int) *component {
	c.el.Style(styles.InsetRingJadeAlpha(scale))
	return c
}

func (c *component) InsetRingJadeDark(scale int) *component {
	c.el.Style(styles.InsetRingJadeDark(scale))
	return c
}

func (c *component) InsetRingJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingJadeDarkAlpha(scale))
	return c
}

func (c *component) InsetRingLime(scale int) *component {
	c.el.Style(styles.InsetRingLime(scale))
	return c
}

func (c *component) InsetRingLimeAlpha(scale int) *component {
	c.el.Style(styles.InsetRingLimeAlpha(scale))
	return c
}

func (c *component) InsetRingLimeDark(scale int) *component {
	c.el.Style(styles.InsetRingLimeDark(scale))
	return c
}

func (c *component) InsetRingLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingLimeDarkAlpha(scale))
	return c
}

func (c *component) InsetRingMauve(scale int) *component {
	c.el.Style(styles.InsetRingMauve(scale))
	return c
}

func (c *component) InsetRingMauveAlpha(scale int) *component {
	c.el.Style(styles.InsetRingMauveAlpha(scale))
	return c
}

func (c *component) InsetRingMauveDark(scale int) *component {
	c.el.Style(styles.InsetRingMauveDark(scale))
	return c
}

func (c *component) InsetRingMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingMauveDarkAlpha(scale))
	return c
}

func (c *component) InsetRingMint(scale int) *component {
	c.el.Style(styles.InsetRingMint(scale))
	return c
}

func (c *component) InsetRingMintAlpha(scale int) *component {
	c.el.Style(styles.InsetRingMintAlpha(scale))
	return c
}

func (c *component) InsetRingMintDark(scale int) *component {
	c.el.Style(styles.InsetRingMintDark(scale))
	return c
}

func (c *component) InsetRingMintDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingMintDarkAlpha(scale))
	return c
}

func (c *component) InsetRingOlive(scale int) *component {
	c.el.Style(styles.InsetRingOlive(scale))
	return c
}

func (c *component) InsetRingOliveAlpha(scale int) *component {
	c.el.Style(styles.InsetRingOliveAlpha(scale))
	return c
}

func (c *component) InsetRingOliveDark(scale int) *component {
	c.el.Style(styles.InsetRingOliveDark(scale))
	return c
}

func (c *component) InsetRingOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingOliveDarkAlpha(scale))
	return c
}

func (c *component) InsetRingOrange(scale int) *component {
	c.el.Style(styles.InsetRingOrange(scale))
	return c
}

func (c *component) InsetRingOrangeAlpha(scale int) *component {
	c.el.Style(styles.InsetRingOrangeAlpha(scale))
	return c
}

func (c *component) InsetRingOrangeDark(scale int) *component {
	c.el.Style(styles.InsetRingOrangeDark(scale))
	return c
}

func (c *component) InsetRingOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingOrangeDarkAlpha(scale))
	return c
}

func (c *component) InsetRingPink(scale int) *component {
	c.el.Style(styles.InsetRingPink(scale))
	return c
}

func (c *component) InsetRingPinkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingPinkAlpha(scale))
	return c
}

func (c *component) InsetRingPinkDark(scale int) *component {
	c.el.Style(styles.InsetRingPinkDark(scale))
	return c
}

func (c *component) InsetRingPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingPinkDarkAlpha(scale))
	return c
}

func (c *component) InsetRingPlum(scale int) *component {
	c.el.Style(styles.InsetRingPlum(scale))
	return c
}

func (c *component) InsetRingPlumAlpha(scale int) *component {
	c.el.Style(styles.InsetRingPlumAlpha(scale))
	return c
}

func (c *component) InsetRingPlumDark(scale int) *component {
	c.el.Style(styles.InsetRingPlumDark(scale))
	return c
}

func (c *component) InsetRingPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingPlumDarkAlpha(scale))
	return c
}

func (c *component) InsetRingPurple(scale int) *component {
	c.el.Style(styles.InsetRingPurple(scale))
	return c
}

func (c *component) InsetRingPurpleAlpha(scale int) *component {
	c.el.Style(styles.InsetRingPurpleAlpha(scale))
	return c
}

func (c *component) InsetRingPurpleDark(scale int) *component {
	c.el.Style(styles.InsetRingPurpleDark(scale))
	return c
}

func (c *component) InsetRingPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingPurpleDarkAlpha(scale))
	return c
}

func (c *component) InsetRingRed(scale int) *component {
	c.el.Style(styles.InsetRingRed(scale))
	return c
}

func (c *component) InsetRingRedAlpha(scale int) *component {
	c.el.Style(styles.InsetRingRedAlpha(scale))
	return c
}

func (c *component) InsetRingRedDark(scale int) *component {
	c.el.Style(styles.InsetRingRedDark(scale))
	return c
}

func (c *component) InsetRingRedDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingRedDarkAlpha(scale))
	return c
}

func (c *component) InsetRingRuby(scale int) *component {
	c.el.Style(styles.InsetRingRuby(scale))
	return c
}

func (c *component) InsetRingRubyAlpha(scale int) *component {
	c.el.Style(styles.InsetRingRubyAlpha(scale))
	return c
}

func (c *component) InsetRingRubyDark(scale int) *component {
	c.el.Style(styles.InsetRingRubyDark(scale))
	return c
}

func (c *component) InsetRingRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingRubyDarkAlpha(scale))
	return c
}

func (c *component) InsetRingSage(scale int) *component {
	c.el.Style(styles.InsetRingSage(scale))
	return c
}

func (c *component) InsetRingSageAlpha(scale int) *component {
	c.el.Style(styles.InsetRingSageAlpha(scale))
	return c
}

func (c *component) InsetRingSageDark(scale int) *component {
	c.el.Style(styles.InsetRingSageDark(scale))
	return c
}

func (c *component) InsetRingSageDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingSageDarkAlpha(scale))
	return c
}

func (c *component) InsetRingSand(scale int) *component {
	c.el.Style(styles.InsetRingSand(scale))
	return c
}

func (c *component) InsetRingSandAlpha(scale int) *component {
	c.el.Style(styles.InsetRingSandAlpha(scale))
	return c
}

func (c *component) InsetRingSandDark(scale int) *component {
	c.el.Style(styles.InsetRingSandDark(scale))
	return c
}

func (c *component) InsetRingSandDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingSandDarkAlpha(scale))
	return c
}

func (c *component) InsetRingSky(scale int) *component {
	c.el.Style(styles.InsetRingSky(scale))
	return c
}

func (c *component) InsetRingSkyAlpha(scale int) *component {
	c.el.Style(styles.InsetRingSkyAlpha(scale))
	return c
}

func (c *component) InsetRingSkyDark(scale int) *component {
	c.el.Style(styles.InsetRingSkyDark(scale))
	return c
}

func (c *component) InsetRingSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingSkyDarkAlpha(scale))
	return c
}

func (c *component) InsetRingSlate(scale int) *component {
	c.el.Style(styles.InsetRingSlate(scale))
	return c
}

func (c *component) InsetRingSlateAlpha(scale int) *component {
	c.el.Style(styles.InsetRingSlateAlpha(scale))
	return c
}

func (c *component) InsetRingSlateDark(scale int) *component {
	c.el.Style(styles.InsetRingSlateDark(scale))
	return c
}

func (c *component) InsetRingSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingSlateDarkAlpha(scale))
	return c
}

func (c *component) InsetRingTeal(scale int) *component {
	c.el.Style(styles.InsetRingTeal(scale))
	return c
}

func (c *component) InsetRingTealAlpha(scale int) *component {
	c.el.Style(styles.InsetRingTealAlpha(scale))
	return c
}

func (c *component) InsetRingTealDark(scale int) *component {
	c.el.Style(styles.InsetRingTealDark(scale))
	return c
}

func (c *component) InsetRingTealDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingTealDarkAlpha(scale))
	return c
}

func (c *component) InsetRingTomato(scale int) *component {
	c.el.Style(styles.InsetRingTomato(scale))
	return c
}

func (c *component) InsetRingTomatoAlpha(scale int) *component {
	c.el.Style(styles.InsetRingTomatoAlpha(scale))
	return c
}

func (c *component) InsetRingTomatoDark(scale int) *component {
	c.el.Style(styles.InsetRingTomatoDark(scale))
	return c
}

func (c *component) InsetRingTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingTomatoDarkAlpha(scale))
	return c
}

func (c *component) InsetRingTransparent() *component {
	c.el.Style(styles.InsetRingTransparent())
	return c
}

func (c *component) InsetRingViolet(scale int) *component {
	c.el.Style(styles.InsetRingViolet(scale))
	return c
}

func (c *component) InsetRingVioletAlpha(scale int) *component {
	c.el.Style(styles.InsetRingVioletAlpha(scale))
	return c
}

func (c *component) InsetRingVioletDark(scale int) *component {
	c.el.Style(styles.InsetRingVioletDark(scale))
	return c
}

func (c *component) InsetRingVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingVioletDarkAlpha(scale))
	return c
}

func (c *component) InsetRingWhite() *component {
	c.el.Style(styles.InsetRingWhite())
	return c
}

func (c *component) InsetRingWhiteAlpha(scale int) *component {
	c.el.Style(styles.InsetRingWhiteAlpha(scale))
	return c
}

func (c *component) InsetRingYellow(scale int) *component {
	c.el.Style(styles.InsetRingYellow(scale))
	return c
}

func (c *component) InsetRingYellowAlpha(scale int) *component {
	c.el.Style(styles.InsetRingYellowAlpha(scale))
	return c
}

func (c *component) InsetRingYellowDark(scale int) *component {
	c.el.Style(styles.InsetRingYellowDark(scale))
	return c
}

func (c *component) InsetRingYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetRingYellowDarkAlpha(scale))
	return c
}

func (c *component) InsetShadow(val value.Value) *component {
	c.el.Style(styles.InsetShadow(val))
	return c
}

func (c *component) InsetShadow2xl() *component {
	c.el.Style(styles.InsetShadow2xl())
	return c
}

func (c *component) InsetShadow2xs() *component {
	c.el.Style(styles.InsetShadow2xs())
	return c
}

func (c *component) InsetShadowAmber(scale int) *component {
	c.el.Style(styles.InsetShadowAmber(scale))
	return c
}

func (c *component) InsetShadowAmberAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowAmberAlpha(scale))
	return c
}

func (c *component) InsetShadowAmberDark(scale int) *component {
	c.el.Style(styles.InsetShadowAmberDark(scale))
	return c
}

func (c *component) InsetShadowAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowAmberDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowBlack() *component {
	c.el.Style(styles.InsetShadowBlack())
	return c
}

func (c *component) InsetShadowBlackAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowBlackAlpha(scale))
	return c
}

func (c *component) InsetShadowBlue(scale int) *component {
	c.el.Style(styles.InsetShadowBlue(scale))
	return c
}

func (c *component) InsetShadowBlueAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowBlueAlpha(scale))
	return c
}

func (c *component) InsetShadowBlueDark(scale int) *component {
	c.el.Style(styles.InsetShadowBlueDark(scale))
	return c
}

func (c *component) InsetShadowBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowBlueDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowBronze(scale int) *component {
	c.el.Style(styles.InsetShadowBronze(scale))
	return c
}

func (c *component) InsetShadowBronzeAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowBronzeAlpha(scale))
	return c
}

func (c *component) InsetShadowBronzeDark(scale int) *component {
	c.el.Style(styles.InsetShadowBronzeDark(scale))
	return c
}

func (c *component) InsetShadowBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowBronzeDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowBrown(scale int) *component {
	c.el.Style(styles.InsetShadowBrown(scale))
	return c
}

func (c *component) InsetShadowBrownAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowBrownAlpha(scale))
	return c
}

func (c *component) InsetShadowBrownDark(scale int) *component {
	c.el.Style(styles.InsetShadowBrownDark(scale))
	return c
}

func (c *component) InsetShadowBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowBrownDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowColor(val value.Value) *component {
	c.el.Style(styles.InsetShadowColor(val))
	return c
}

func (c *component) InsetShadowCrimson(scale int) *component {
	c.el.Style(styles.InsetShadowCrimson(scale))
	return c
}

func (c *component) InsetShadowCrimsonAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowCrimsonAlpha(scale))
	return c
}

func (c *component) InsetShadowCrimsonDark(scale int) *component {
	c.el.Style(styles.InsetShadowCrimsonDark(scale))
	return c
}

func (c *component) InsetShadowCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowCrimsonDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowCurrent() *component {
	c.el.Style(styles.InsetShadowCurrent())
	return c
}

func (c *component) InsetShadowCyan(scale int) *component {
	c.el.Style(styles.InsetShadowCyan(scale))
	return c
}

func (c *component) InsetShadowCyanAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowCyanAlpha(scale))
	return c
}

func (c *component) InsetShadowCyanDark(scale int) *component {
	c.el.Style(styles.InsetShadowCyanDark(scale))
	return c
}

func (c *component) InsetShadowCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowCyanDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowGold(scale int) *component {
	c.el.Style(styles.InsetShadowGold(scale))
	return c
}

func (c *component) InsetShadowGoldAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowGoldAlpha(scale))
	return c
}

func (c *component) InsetShadowGoldDark(scale int) *component {
	c.el.Style(styles.InsetShadowGoldDark(scale))
	return c
}

func (c *component) InsetShadowGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowGoldDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowGrass(scale int) *component {
	c.el.Style(styles.InsetShadowGrass(scale))
	return c
}

func (c *component) InsetShadowGrassAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowGrassAlpha(scale))
	return c
}

func (c *component) InsetShadowGrassDark(scale int) *component {
	c.el.Style(styles.InsetShadowGrassDark(scale))
	return c
}

func (c *component) InsetShadowGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowGrassDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowGray(scale int) *component {
	c.el.Style(styles.InsetShadowGray(scale))
	return c
}

func (c *component) InsetShadowGrayAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowGrayAlpha(scale))
	return c
}

func (c *component) InsetShadowGrayDark(scale int) *component {
	c.el.Style(styles.InsetShadowGrayDark(scale))
	return c
}

func (c *component) InsetShadowGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowGrayDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowGreen(scale int) *component {
	c.el.Style(styles.InsetShadowGreen(scale))
	return c
}

func (c *component) InsetShadowGreenAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowGreenAlpha(scale))
	return c
}

func (c *component) InsetShadowGreenDark(scale int) *component {
	c.el.Style(styles.InsetShadowGreenDark(scale))
	return c
}

func (c *component) InsetShadowGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowGreenDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowIndigo(scale int) *component {
	c.el.Style(styles.InsetShadowIndigo(scale))
	return c
}

func (c *component) InsetShadowIndigoAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowIndigoAlpha(scale))
	return c
}

func (c *component) InsetShadowIndigoDark(scale int) *component {
	c.el.Style(styles.InsetShadowIndigoDark(scale))
	return c
}

func (c *component) InsetShadowIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowIndigoDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowInherit() *component {
	c.el.Style(styles.InsetShadowInherit())
	return c
}

func (c *component) InsetShadowIris(scale int) *component {
	c.el.Style(styles.InsetShadowIris(scale))
	return c
}

func (c *component) InsetShadowIrisAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowIrisAlpha(scale))
	return c
}

func (c *component) InsetShadowIrisDark(scale int) *component {
	c.el.Style(styles.InsetShadowIrisDark(scale))
	return c
}

func (c *component) InsetShadowIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowIrisDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowJade(scale int) *component {
	c.el.Style(styles.InsetShadowJade(scale))
	return c
}

func (c *component) InsetShadowJadeAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowJadeAlpha(scale))
	return c
}

func (c *component) InsetShadowJadeDark(scale int) *component {
	c.el.Style(styles.InsetShadowJadeDark(scale))
	return c
}

func (c *component) InsetShadowJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowJadeDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowLg() *component {
	c.el.Style(styles.InsetShadowLg())
	return c
}

func (c *component) InsetShadowLime(scale int) *component {
	c.el.Style(styles.InsetShadowLime(scale))
	return c
}

func (c *component) InsetShadowLimeAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowLimeAlpha(scale))
	return c
}

func (c *component) InsetShadowLimeDark(scale int) *component {
	c.el.Style(styles.InsetShadowLimeDark(scale))
	return c
}

func (c *component) InsetShadowLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowLimeDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowMauve(scale int) *component {
	c.el.Style(styles.InsetShadowMauve(scale))
	return c
}

func (c *component) InsetShadowMauveAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowMauveAlpha(scale))
	return c
}

func (c *component) InsetShadowMauveDark(scale int) *component {
	c.el.Style(styles.InsetShadowMauveDark(scale))
	return c
}

func (c *component) InsetShadowMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowMauveDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowMd() *component {
	c.el.Style(styles.InsetShadowMd())
	return c
}

func (c *component) InsetShadowMint(scale int) *component {
	c.el.Style(styles.InsetShadowMint(scale))
	return c
}

func (c *component) InsetShadowMintAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowMintAlpha(scale))
	return c
}

func (c *component) InsetShadowMintDark(scale int) *component {
	c.el.Style(styles.InsetShadowMintDark(scale))
	return c
}

func (c *component) InsetShadowMintDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowMintDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowNone() *component {
	c.el.Style(styles.InsetShadowNone())
	return c
}

func (c *component) InsetShadowOlive(scale int) *component {
	c.el.Style(styles.InsetShadowOlive(scale))
	return c
}

func (c *component) InsetShadowOliveAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowOliveAlpha(scale))
	return c
}

func (c *component) InsetShadowOliveDark(scale int) *component {
	c.el.Style(styles.InsetShadowOliveDark(scale))
	return c
}

func (c *component) InsetShadowOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowOliveDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowOrange(scale int) *component {
	c.el.Style(styles.InsetShadowOrange(scale))
	return c
}

func (c *component) InsetShadowOrangeAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowOrangeAlpha(scale))
	return c
}

func (c *component) InsetShadowOrangeDark(scale int) *component {
	c.el.Style(styles.InsetShadowOrangeDark(scale))
	return c
}

func (c *component) InsetShadowOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowOrangeDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowPink(scale int) *component {
	c.el.Style(styles.InsetShadowPink(scale))
	return c
}

func (c *component) InsetShadowPinkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowPinkAlpha(scale))
	return c
}

func (c *component) InsetShadowPinkDark(scale int) *component {
	c.el.Style(styles.InsetShadowPinkDark(scale))
	return c
}

func (c *component) InsetShadowPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowPinkDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowPlum(scale int) *component {
	c.el.Style(styles.InsetShadowPlum(scale))
	return c
}

func (c *component) InsetShadowPlumAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowPlumAlpha(scale))
	return c
}

func (c *component) InsetShadowPlumDark(scale int) *component {
	c.el.Style(styles.InsetShadowPlumDark(scale))
	return c
}

func (c *component) InsetShadowPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowPlumDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowPurple(scale int) *component {
	c.el.Style(styles.InsetShadowPurple(scale))
	return c
}

func (c *component) InsetShadowPurpleAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowPurpleAlpha(scale))
	return c
}

func (c *component) InsetShadowPurpleDark(scale int) *component {
	c.el.Style(styles.InsetShadowPurpleDark(scale))
	return c
}

func (c *component) InsetShadowPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowPurpleDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowRed(scale int) *component {
	c.el.Style(styles.InsetShadowRed(scale))
	return c
}

func (c *component) InsetShadowRedAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowRedAlpha(scale))
	return c
}

func (c *component) InsetShadowRedDark(scale int) *component {
	c.el.Style(styles.InsetShadowRedDark(scale))
	return c
}

func (c *component) InsetShadowRedDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowRedDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowRuby(scale int) *component {
	c.el.Style(styles.InsetShadowRuby(scale))
	return c
}

func (c *component) InsetShadowRubyAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowRubyAlpha(scale))
	return c
}

func (c *component) InsetShadowRubyDark(scale int) *component {
	c.el.Style(styles.InsetShadowRubyDark(scale))
	return c
}

func (c *component) InsetShadowRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowRubyDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowSage(scale int) *component {
	c.el.Style(styles.InsetShadowSage(scale))
	return c
}

func (c *component) InsetShadowSageAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowSageAlpha(scale))
	return c
}

func (c *component) InsetShadowSageDark(scale int) *component {
	c.el.Style(styles.InsetShadowSageDark(scale))
	return c
}

func (c *component) InsetShadowSageDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowSageDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowSand(scale int) *component {
	c.el.Style(styles.InsetShadowSand(scale))
	return c
}

func (c *component) InsetShadowSandAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowSandAlpha(scale))
	return c
}

func (c *component) InsetShadowSandDark(scale int) *component {
	c.el.Style(styles.InsetShadowSandDark(scale))
	return c
}

func (c *component) InsetShadowSandDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowSandDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowSky(scale int) *component {
	c.el.Style(styles.InsetShadowSky(scale))
	return c
}

func (c *component) InsetShadowSkyAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowSkyAlpha(scale))
	return c
}

func (c *component) InsetShadowSkyDark(scale int) *component {
	c.el.Style(styles.InsetShadowSkyDark(scale))
	return c
}

func (c *component) InsetShadowSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowSkyDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowSlate(scale int) *component {
	c.el.Style(styles.InsetShadowSlate(scale))
	return c
}

func (c *component) InsetShadowSlateAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowSlateAlpha(scale))
	return c
}

func (c *component) InsetShadowSlateDark(scale int) *component {
	c.el.Style(styles.InsetShadowSlateDark(scale))
	return c
}

func (c *component) InsetShadowSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowSlateDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowSm() *component {
	c.el.Style(styles.InsetShadowSm())
	return c
}

func (c *component) InsetShadowTeal(scale int) *component {
	c.el.Style(styles.InsetShadowTeal(scale))
	return c
}

func (c *component) InsetShadowTealAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowTealAlpha(scale))
	return c
}

func (c *component) InsetShadowTealDark(scale int) *component {
	c.el.Style(styles.InsetShadowTealDark(scale))
	return c
}

func (c *component) InsetShadowTealDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowTealDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowTomato(scale int) *component {
	c.el.Style(styles.InsetShadowTomato(scale))
	return c
}

func (c *component) InsetShadowTomatoAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowTomatoAlpha(scale))
	return c
}

func (c *component) InsetShadowTomatoDark(scale int) *component {
	c.el.Style(styles.InsetShadowTomatoDark(scale))
	return c
}

func (c *component) InsetShadowTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowTomatoDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowTransparent() *component {
	c.el.Style(styles.InsetShadowTransparent())
	return c
}

func (c *component) InsetShadowViolet(scale int) *component {
	c.el.Style(styles.InsetShadowViolet(scale))
	return c
}

func (c *component) InsetShadowVioletAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowVioletAlpha(scale))
	return c
}

func (c *component) InsetShadowVioletDark(scale int) *component {
	c.el.Style(styles.InsetShadowVioletDark(scale))
	return c
}

func (c *component) InsetShadowVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowVioletDarkAlpha(scale))
	return c
}

func (c *component) InsetShadowWhite() *component {
	c.el.Style(styles.InsetShadowWhite())
	return c
}

func (c *component) InsetShadowWhiteAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowWhiteAlpha(scale))
	return c
}

func (c *component) InsetShadowXl() *component {
	c.el.Style(styles.InsetShadowXl())
	return c
}

func (c *component) InsetShadowXs() *component {
	c.el.Style(styles.InsetShadowXs())
	return c
}

func (c *component) InsetShadowYellow(scale int) *component {
	c.el.Style(styles.InsetShadowYellow(scale))
	return c
}

func (c *component) InsetShadowYellowAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowYellowAlpha(scale))
	return c
}

func (c *component) InsetShadowYellowDark(scale int) *component {
	c.el.Style(styles.InsetShadowYellowDark(scale))
	return c
}

func (c *component) InsetShadowYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.InsetShadowYellowDarkAlpha(scale))
	return c
}

func (c *component) InsetX(number int) *component {
	c.el.Style(styles.InsetX(number))
	return c
}

func (c *component) InsetXAuto() *component {
	c.el.Style(styles.InsetXAuto())
	return c
}

func (c *component) InsetXBy(val value.Value) *component {
	c.el.Style(styles.InsetXBy(val))
	return c
}

func (c *component) InsetXFraction(fraction float64) *component {
	c.el.Style(styles.InsetXFraction(fraction))
	return c
}

func (c *component) InsetXFull() *component {
	c.el.Style(styles.InsetXFull())
	return c
}

func (c *component) InsetXPx() *component {
	c.el.Style(styles.InsetXPx())
	return c
}

func (c *component) InsetY(number int) *component {
	c.el.Style(styles.InsetY(number))
	return c
}

func (c *component) InsetYAuto() *component {
	c.el.Style(styles.InsetYAuto())
	return c
}

func (c *component) InsetYBy(val value.Value) *component {
	c.el.Style(styles.InsetYBy(val))
	return c
}

func (c *component) InsetYFraction(fraction float64) *component {
	c.el.Style(styles.InsetYFraction(fraction))
	return c
}

func (c *component) InsetYFull() *component {
	c.el.Style(styles.InsetYFull())
	return c
}

func (c *component) InsetYPx() *component {
	c.el.Style(styles.InsetYPx())
	return c
}

func (c *component) Invert(val ...any) *component {
	c.el.Style(styles.Invert(val...))
	return c
}

func (c *component) Invisible() *component {
	c.el.Style(styles.Invisible())
	return c
}

func (c *component) Isolate() *component {
	c.el.Style(styles.Isolate())
	return c
}

func (c *component) IsolationAuto() *component {
	c.el.Style(styles.IsolationAuto())
	return c
}

func (c *component) Italic() *component {
	c.el.Style(styles.Italic())
	return c
}

func (c *component) ItemsBaseline() *component {
	c.el.Style(styles.ItemsBaseline())
	return c
}

func (c *component) ItemsBaselineLast() *component {
	c.el.Style(styles.ItemsBaselineLast())
	return c
}

func (c *component) ItemsCenter() *component {
	c.el.Style(styles.ItemsCenter())
	return c
}

func (c *component) ItemsCenterSafe() *component {
	c.el.Style(styles.ItemsCenterSafe())
	return c
}

func (c *component) ItemsEnd() *component {
	c.el.Style(styles.ItemsEnd())
	return c
}

func (c *component) ItemsEndSafe() *component {
	c.el.Style(styles.ItemsEndSafe())
	return c
}

func (c *component) ItemsStart() *component {
	c.el.Style(styles.ItemsStart())
	return c
}

func (c *component) ItemsStretch() *component {
	c.el.Style(styles.ItemsStretch())
	return c
}

func (c *component) JustifyAround() *component {
	c.el.Style(styles.JustifyAround())
	return c
}

func (c *component) JustifyBaseline() *component {
	c.el.Style(styles.JustifyBaseline())
	return c
}

func (c *component) JustifyBetween() *component {
	c.el.Style(styles.JustifyBetween())
	return c
}

func (c *component) JustifyCenter() *component {
	c.el.Style(styles.JustifyCenter())
	return c
}

func (c *component) JustifyCenterSafe() *component {
	c.el.Style(styles.JustifyCenterSafe())
	return c
}

func (c *component) JustifyEnd() *component {
	c.el.Style(styles.JustifyEnd())
	return c
}

func (c *component) JustifyEndSafe() *component {
	c.el.Style(styles.JustifyEndSafe())
	return c
}

func (c *component) JustifyEvenly() *component {
	c.el.Style(styles.JustifyEvenly())
	return c
}

func (c *component) JustifyItemsCenter() *component {
	c.el.Style(styles.JustifyItemsCenter())
	return c
}

func (c *component) JustifyItemsCenterSafe() *component {
	c.el.Style(styles.JustifyItemsCenterSafe())
	return c
}

func (c *component) JustifyItemsEnd() *component {
	c.el.Style(styles.JustifyItemsEnd())
	return c
}

func (c *component) JustifyItemsEndSafe() *component {
	c.el.Style(styles.JustifyItemsEndSafe())
	return c
}

func (c *component) JustifyItemsNormal() *component {
	c.el.Style(styles.JustifyItemsNormal())
	return c
}

func (c *component) JustifyItemsStart() *component {
	c.el.Style(styles.JustifyItemsStart())
	return c
}

func (c *component) JustifyItemsStretch() *component {
	c.el.Style(styles.JustifyItemsStretch())
	return c
}

func (c *component) JustifyNormal() *component {
	c.el.Style(styles.JustifyNormal())
	return c
}

func (c *component) JustifySelfAuto() *component {
	c.el.Style(styles.JustifySelfAuto())
	return c
}

func (c *component) JustifySelfCenter() *component {
	c.el.Style(styles.JustifySelfCenter())
	return c
}

func (c *component) JustifySelfCenterSafe() *component {
	c.el.Style(styles.JustifySelfCenterSafe())
	return c
}

func (c *component) JustifySelfEnd() *component {
	c.el.Style(styles.JustifySelfEnd())
	return c
}

func (c *component) JustifySelfEndSafe() *component {
	c.el.Style(styles.JustifySelfEndSafe())
	return c
}

func (c *component) JustifySelfStart() *component {
	c.el.Style(styles.JustifySelfStart())
	return c
}

func (c *component) JustifySelfStretch() *component {
	c.el.Style(styles.JustifySelfStretch())
	return c
}

func (c *component) JustifyStart() *component {
	c.el.Style(styles.JustifyStart())
	return c
}

func (c *component) JustifyStretch() *component {
	c.el.Style(styles.JustifyStretch())
	return c
}

func (c *component) Leading(number int) *component {
	c.el.Style(styles.Leading(number))
	return c
}

func (c *component) LeadingBy(val value.Value) *component {
	c.el.Style(styles.LeadingBy(val))
	return c
}

func (c *component) LeadingNone() *component {
	c.el.Style(styles.LeadingNone())
	return c
}

func (c *component) Left(number int) *component {
	c.el.Style(styles.Left(number))
	return c
}

func (c *component) LeftAuto() *component {
	c.el.Style(styles.LeftAuto())
	return c
}

func (c *component) LeftBy(val value.Value) *component {
	c.el.Style(styles.LeftBy(val))
	return c
}

func (c *component) LeftFraction(fraction float64) *component {
	c.el.Style(styles.LeftFraction(fraction))
	return c
}

func (c *component) LeftFull() *component {
	c.el.Style(styles.LeftFull())
	return c
}

func (c *component) LeftPx() *component {
	c.el.Style(styles.LeftPx())
	return c
}

func (c *component) LineClamp(number int) *component {
	c.el.Style(styles.LineClamp(number))
	return c
}

func (c *component) LineClampBy(val value.Value) *component {
	c.el.Style(styles.LineClampBy(val))
	return c
}

func (c *component) LineClampNone() *component {
	c.el.Style(styles.LineClampNone())
	return c
}

func (c *component) LineThrough() *component {
	c.el.Style(styles.LineThrough())
	return c
}

func (c *component) LiningNums() *component {
	c.el.Style(styles.LiningNums())
	return c
}

func (c *component) List(val value.Value) *component {
	c.el.Style(styles.List(val))
	return c
}

func (c *component) ListDecimal() *component {
	c.el.Style(styles.ListDecimal())
	return c
}

func (c *component) ListDisc() *component {
	c.el.Style(styles.ListDisc())
	return c
}

func (c *component) ListImage(val value.Value) *component {
	c.el.Style(styles.ListImage(val))
	return c
}

func (c *component) ListImageNone() *component {
	c.el.Style(styles.ListImageNone())
	return c
}

func (c *component) ListInside() *component {
	c.el.Style(styles.ListInside())
	return c
}

func (c *component) ListItem() *component {
	c.el.Style(styles.ListItem())
	return c
}

func (c *component) ListNone() *component {
	c.el.Style(styles.ListNone())
	return c
}

func (c *component) ListOutside() *component {
	c.el.Style(styles.ListOutside())
	return c
}

func (c *component) Lowercase() *component {
	c.el.Style(styles.Lowercase())
	return c
}

func (c *component) M(number int) *component {
	c.el.Style(styles.M(number))
	return c
}

func (c *component) MAuto() *component {
	c.el.Style(styles.MAuto())
	return c
}

func (c *component) MBy(val value.Value) *component {
	c.el.Style(styles.MBy(val))
	return c
}

func (c *component) MPx() *component {
	c.el.Style(styles.MPx())
	return c
}

func (c *component) Mask(val value.Value) *component {
	c.el.Style(styles.Mask(val))
	return c
}

func (c *component) MaskAdd() *component {
	c.el.Style(styles.MaskAdd())
	return c
}

func (c *component) MaskAlpha() *component {
	c.el.Style(styles.MaskAlpha())
	return c
}

func (c *component) MaskAuto() *component {
	c.el.Style(styles.MaskAuto())
	return c
}

func (c *component) MaskBottom() *component {
	c.el.Style(styles.MaskBottom())
	return c
}

func (c *component) MaskBottomFrom(val any) *component {
	c.el.Style(styles.MaskBottomFrom(val))
	return c
}

func (c *component) MaskBottomLeft() *component {
	c.el.Style(styles.MaskBottomLeft())
	return c
}

func (c *component) MaskBottomRight() *component {
	c.el.Style(styles.MaskBottomRight())
	return c
}

func (c *component) MaskBottomTo(val any) *component {
	c.el.Style(styles.MaskBottomTo(val))
	return c
}

func (c *component) MaskCenter() *component {
	c.el.Style(styles.MaskCenter())
	return c
}

func (c *component) MaskCircle() *component {
	c.el.Style(styles.MaskCircle())
	return c
}

func (c *component) MaskClipBorder() *component {
	c.el.Style(styles.MaskClipBorder())
	return c
}

func (c *component) MaskClipContent() *component {
	c.el.Style(styles.MaskClipContent())
	return c
}

func (c *component) MaskClipFill() *component {
	c.el.Style(styles.MaskClipFill())
	return c
}

func (c *component) MaskClipPadding() *component {
	c.el.Style(styles.MaskClipPadding())
	return c
}

func (c *component) MaskClipStroke() *component {
	c.el.Style(styles.MaskClipStroke())
	return c
}

func (c *component) MaskClipView() *component {
	c.el.Style(styles.MaskClipView())
	return c
}

func (c *component) MaskConic(number float64) *component {
	c.el.Style(styles.MaskConic(number))
	return c
}

func (c *component) MaskConicFrom(val any) *component {
	c.el.Style(styles.MaskConicFrom(val))
	return c
}

func (c *component) MaskConicTo(val any) *component {
	c.el.Style(styles.MaskConicTo(val))
	return c
}

func (c *component) MaskContain() *component {
	c.el.Style(styles.MaskContain())
	return c
}

func (c *component) MaskCover() *component {
	c.el.Style(styles.MaskCover())
	return c
}

func (c *component) MaskEllipse() *component {
	c.el.Style(styles.MaskEllipse())
	return c
}

func (c *component) MaskExclude() *component {
	c.el.Style(styles.MaskExclude())
	return c
}

func (c *component) MaskIntersect() *component {
	c.el.Style(styles.MaskIntersect())
	return c
}

func (c *component) MaskLeft() *component {
	c.el.Style(styles.MaskLeft())
	return c
}

func (c *component) MaskLeftFrom(val any) *component {
	c.el.Style(styles.MaskLeftFrom(val))
	return c
}

func (c *component) MaskLeftTo(val any) *component {
	c.el.Style(styles.MaskLeftTo(val))
	return c
}

func (c *component) MaskLinear(degree int) *component {
	c.el.Style(styles.MaskLinear(degree))
	return c
}

func (c *component) MaskLinearFrom(val any) *component {
	c.el.Style(styles.MaskLinearFrom(val))
	return c
}

func (c *component) MaskLinearTo(val any) *component {
	c.el.Style(styles.MaskLinearTo(val))
	return c
}

func (c *component) MaskLuminance() *component {
	c.el.Style(styles.MaskLuminance())
	return c
}

func (c *component) MaskMatch() *component {
	c.el.Style(styles.MaskMatch())
	return c
}

func (c *component) MaskNoClip() *component {
	c.el.Style(styles.MaskNoClip())
	return c
}

func (c *component) MaskNoRepeat() *component {
	c.el.Style(styles.MaskNoRepeat())
	return c
}

func (c *component) MaskNone() *component {
	c.el.Style(styles.MaskNone())
	return c
}

func (c *component) MaskOriginBorder() *component {
	c.el.Style(styles.MaskOriginBorder())
	return c
}

func (c *component) MaskOriginContent() *component {
	c.el.Style(styles.MaskOriginContent())
	return c
}

func (c *component) MaskOriginFill() *component {
	c.el.Style(styles.MaskOriginFill())
	return c
}

func (c *component) MaskOriginPadding() *component {
	c.el.Style(styles.MaskOriginPadding())
	return c
}

func (c *component) MaskOriginStroke() *component {
	c.el.Style(styles.MaskOriginStroke())
	return c
}

func (c *component) MaskOriginView() *component {
	c.el.Style(styles.MaskOriginView())
	return c
}

func (c *component) MaskPosition(val value.Value) *component {
	c.el.Style(styles.MaskPosition(val))
	return c
}

func (c *component) MaskRadial(val ...any) *component {
	c.el.Style(styles.MaskRadial(val...))
	return c
}

func (c *component) MaskRadialAtBottom() *component {
	c.el.Style(styles.MaskRadialAtBottom())
	return c
}

func (c *component) MaskRadialAtBottomLeft() *component {
	c.el.Style(styles.MaskRadialAtBottomLeft())
	return c
}

func (c *component) MaskRadialAtBottomRight() *component {
	c.el.Style(styles.MaskRadialAtBottomRight())
	return c
}

func (c *component) MaskRadialAtCenter() *component {
	c.el.Style(styles.MaskRadialAtCenter())
	return c
}

func (c *component) MaskRadialAtLeft() *component {
	c.el.Style(styles.MaskRadialAtLeft())
	return c
}

func (c *component) MaskRadialAtRight() *component {
	c.el.Style(styles.MaskRadialAtRight())
	return c
}

func (c *component) MaskRadialAtTop() *component {
	c.el.Style(styles.MaskRadialAtTop())
	return c
}

func (c *component) MaskRadialAtTopLeft() *component {
	c.el.Style(styles.MaskRadialAtTopLeft())
	return c
}

func (c *component) MaskRadialAtTopRight() *component {
	c.el.Style(styles.MaskRadialAtTopRight())
	return c
}

func (c *component) MaskRadialClosestCorner() *component {
	c.el.Style(styles.MaskRadialClosestCorner())
	return c
}

func (c *component) MaskRadialClosestSide() *component {
	c.el.Style(styles.MaskRadialClosestSide())
	return c
}

func (c *component) MaskRadialFarthestCorner() *component {
	c.el.Style(styles.MaskRadialFarthestCorner())
	return c
}

func (c *component) MaskRadialFarthestSide() *component {
	c.el.Style(styles.MaskRadialFarthestSide())
	return c
}

func (c *component) MaskRadialFrom(val any) *component {
	c.el.Style(styles.MaskRadialFrom(val))
	return c
}

func (c *component) MaskRadialTo(val any) *component {
	c.el.Style(styles.MaskRadialTo(val))
	return c
}

func (c *component) MaskRepeat() *component {
	c.el.Style(styles.MaskRepeat())
	return c
}

func (c *component) MaskRepeatRound() *component {
	c.el.Style(styles.MaskRepeatRound())
	return c
}

func (c *component) MaskRepeatSpace() *component {
	c.el.Style(styles.MaskRepeatSpace())
	return c
}

func (c *component) MaskRepeatX() *component {
	c.el.Style(styles.MaskRepeatX())
	return c
}

func (c *component) MaskRepeatY() *component {
	c.el.Style(styles.MaskRepeatY())
	return c
}

func (c *component) MaskRight() *component {
	c.el.Style(styles.MaskRight())
	return c
}

func (c *component) MaskRightFrom(val any) *component {
	c.el.Style(styles.MaskRightFrom(val))
	return c
}

func (c *component) MaskRightTo(val any) *component {
	c.el.Style(styles.MaskRightTo(val))
	return c
}

func (c *component) MaskSize(val value.Value) *component {
	c.el.Style(styles.MaskSize(val))
	return c
}

func (c *component) MaskSubtract() *component {
	c.el.Style(styles.MaskSubtract())
	return c
}

func (c *component) MaskTop() *component {
	c.el.Style(styles.MaskTop())
	return c
}

func (c *component) MaskTopFrom(val any) *component {
	c.el.Style(styles.MaskTopFrom(val))
	return c
}

func (c *component) MaskTopLeft() *component {
	c.el.Style(styles.MaskTopLeft())
	return c
}

func (c *component) MaskTopRight() *component {
	c.el.Style(styles.MaskTopRight())
	return c
}

func (c *component) MaskTopTo(val any) *component {
	c.el.Style(styles.MaskTopTo(val))
	return c
}

func (c *component) MaskTypeAlpha() *component {
	c.el.Style(styles.MaskTypeAlpha())
	return c
}

func (c *component) MaskTypeLuminance() *component {
	c.el.Style(styles.MaskTypeLuminance())
	return c
}

func (c *component) MaskXFrom(val any) *component {
	c.el.Style(styles.MaskXFrom(val))
	return c
}

func (c *component) MaskXTo(val any) *component {
	c.el.Style(styles.MaskXTo(val))
	return c
}

func (c *component) MaskYFrom(val any) *component {
	c.el.Style(styles.MaskYFrom(val))
	return c
}

func (c *component) MaskYTo(val any) *component {
	c.el.Style(styles.MaskYTo(val))
	return c
}

func (c *component) MaxH(number int) *component {
	c.el.Style(styles.MaxH(number))
	return c
}

func (c *component) MaxHBy(val value.Value) *component {
	c.el.Style(styles.MaxHBy(val))
	return c
}

func (c *component) MaxHDvh() *component {
	c.el.Style(styles.MaxHDvh())
	return c
}

func (c *component) MaxHDvw() *component {
	c.el.Style(styles.MaxHDvw())
	return c
}

func (c *component) MaxHFit() *component {
	c.el.Style(styles.MaxHFit())
	return c
}

func (c *component) MaxHFraction(fraction float32) *component {
	c.el.Style(styles.MaxHFraction(fraction))
	return c
}

func (c *component) MaxHFull() *component {
	c.el.Style(styles.MaxHFull())
	return c
}

func (c *component) MaxHLh() *component {
	c.el.Style(styles.MaxHLh())
	return c
}

func (c *component) MaxHLvh() *component {
	c.el.Style(styles.MaxHLvh())
	return c
}

func (c *component) MaxHLvw() *component {
	c.el.Style(styles.MaxHLvw())
	return c
}

func (c *component) MaxHMax() *component {
	c.el.Style(styles.MaxHMax())
	return c
}

func (c *component) MaxHMin() *component {
	c.el.Style(styles.MaxHMin())
	return c
}

func (c *component) MaxHNone() *component {
	c.el.Style(styles.MaxHNone())
	return c
}

func (c *component) MaxHPx() *component {
	c.el.Style(styles.MaxHPx())
	return c
}

func (c *component) MaxHScreen() *component {
	c.el.Style(styles.MaxHScreen())
	return c
}

func (c *component) MaxHSvh() *component {
	c.el.Style(styles.MaxHSvh())
	return c
}

func (c *component) MaxHSvw() *component {
	c.el.Style(styles.MaxHSvw())
	return c
}

func (c *component) MaxW(number int) *component {
	c.el.Style(styles.MaxW(number))
	return c
}

func (c *component) MaxW2xl() *component {
	c.el.Style(styles.MaxW2xl())
	return c
}

func (c *component) MaxW2xs() *component {
	c.el.Style(styles.MaxW2xs())
	return c
}

func (c *component) MaxW3xl() *component {
	c.el.Style(styles.MaxW3xl())
	return c
}

func (c *component) MaxW3xs() *component {
	c.el.Style(styles.MaxW3xs())
	return c
}

func (c *component) MaxW4xl() *component {
	c.el.Style(styles.MaxW4xl())
	return c
}

func (c *component) MaxW5xl() *component {
	c.el.Style(styles.MaxW5xl())
	return c
}

func (c *component) MaxW6xl() *component {
	c.el.Style(styles.MaxW6xl())
	return c
}

func (c *component) MaxW7xl() *component {
	c.el.Style(styles.MaxW7xl())
	return c
}

func (c *component) MaxWBy(val value.Value) *component {
	c.el.Style(styles.MaxWBy(val))
	return c
}

func (c *component) MaxWDvh() *component {
	c.el.Style(styles.MaxWDvh())
	return c
}

func (c *component) MaxWDvw() *component {
	c.el.Style(styles.MaxWDvw())
	return c
}

func (c *component) MaxWFit() *component {
	c.el.Style(styles.MaxWFit())
	return c
}

func (c *component) MaxWFraction(fraction float32) *component {
	c.el.Style(styles.MaxWFraction(fraction))
	return c
}

func (c *component) MaxWFull() *component {
	c.el.Style(styles.MaxWFull())
	return c
}

func (c *component) MaxWLg() *component {
	c.el.Style(styles.MaxWLg())
	return c
}

func (c *component) MaxWLvh() *component {
	c.el.Style(styles.MaxWLvh())
	return c
}

func (c *component) MaxWLvw() *component {
	c.el.Style(styles.MaxWLvw())
	return c
}

func (c *component) MaxWMax() *component {
	c.el.Style(styles.MaxWMax())
	return c
}

func (c *component) MaxWMd() *component {
	c.el.Style(styles.MaxWMd())
	return c
}

func (c *component) MaxWMin() *component {
	c.el.Style(styles.MaxWMin())
	return c
}

func (c *component) MaxWNone() *component {
	c.el.Style(styles.MaxWNone())
	return c
}

func (c *component) MaxWPx() *component {
	c.el.Style(styles.MaxWPx())
	return c
}

func (c *component) MaxWScreen() *component {
	c.el.Style(styles.MaxWScreen())
	return c
}

func (c *component) MaxWSm() *component {
	c.el.Style(styles.MaxWSm())
	return c
}

func (c *component) MaxWSvh() *component {
	c.el.Style(styles.MaxWSvh())
	return c
}

func (c *component) MaxWSvw() *component {
	c.el.Style(styles.MaxWSvw())
	return c
}

func (c *component) MaxWXl() *component {
	c.el.Style(styles.MaxWXl())
	return c
}

func (c *component) MaxWXs() *component {
	c.el.Style(styles.MaxWXs())
	return c
}

func (c *component) Mb(number int) *component {
	c.el.Style(styles.Mb(number))
	return c
}

func (c *component) MbAuto() *component {
	c.el.Style(styles.MbAuto())
	return c
}

func (c *component) MbBy(val value.Value) *component {
	c.el.Style(styles.MbBy(val))
	return c
}

func (c *component) MbPx() *component {
	c.el.Style(styles.MbPx())
	return c
}

func (c *component) Me(number int) *component {
	c.el.Style(styles.Me(number))
	return c
}

func (c *component) MeAuto() *component {
	c.el.Style(styles.MeAuto())
	return c
}

func (c *component) MeBy(val value.Value) *component {
	c.el.Style(styles.MeBy(val))
	return c
}

func (c *component) MePx() *component {
	c.el.Style(styles.MePx())
	return c
}

func (c *component) MinH(number int) *component {
	c.el.Style(styles.MinH(number))
	return c
}

func (c *component) MinHAuto() *component {
	c.el.Style(styles.MinHAuto())
	return c
}

func (c *component) MinHBy(val value.Value) *component {
	c.el.Style(styles.MinHBy(val))
	return c
}

func (c *component) MinHDvh() *component {
	c.el.Style(styles.MinHDvh())
	return c
}

func (c *component) MinHDvw() *component {
	c.el.Style(styles.MinHDvw())
	return c
}

func (c *component) MinHFit() *component {
	c.el.Style(styles.MinHFit())
	return c
}

func (c *component) MinHFraction(fraction float32) *component {
	c.el.Style(styles.MinHFraction(fraction))
	return c
}

func (c *component) MinHFull() *component {
	c.el.Style(styles.MinHFull())
	return c
}

func (c *component) MinHLh() *component {
	c.el.Style(styles.MinHLh())
	return c
}

func (c *component) MinHLvh() *component {
	c.el.Style(styles.MinHLvh())
	return c
}

func (c *component) MinHLvw() *component {
	c.el.Style(styles.MinHLvw())
	return c
}

func (c *component) MinHMax() *component {
	c.el.Style(styles.MinHMax())
	return c
}

func (c *component) MinHMin() *component {
	c.el.Style(styles.MinHMin())
	return c
}

func (c *component) MinHPx() *component {
	c.el.Style(styles.MinHPx())
	return c
}

func (c *component) MinHScreen() *component {
	c.el.Style(styles.MinHScreen())
	return c
}

func (c *component) MinHSvh() *component {
	c.el.Style(styles.MinHSvh())
	return c
}

func (c *component) MinHSvw() *component {
	c.el.Style(styles.MinHSvw())
	return c
}

func (c *component) MinW(number int) *component {
	c.el.Style(styles.MinW(number))
	return c
}

func (c *component) MinW2xl() *component {
	c.el.Style(styles.MinW2xl())
	return c
}

func (c *component) MinW2xs() *component {
	c.el.Style(styles.MinW2xs())
	return c
}

func (c *component) MinW3xl() *component {
	c.el.Style(styles.MinW3xl())
	return c
}

func (c *component) MinW3xs() *component {
	c.el.Style(styles.MinW3xs())
	return c
}

func (c *component) MinW4xl() *component {
	c.el.Style(styles.MinW4xl())
	return c
}

func (c *component) MinW5xl() *component {
	c.el.Style(styles.MinW5xl())
	return c
}

func (c *component) MinW6xl() *component {
	c.el.Style(styles.MinW6xl())
	return c
}

func (c *component) MinW7xl() *component {
	c.el.Style(styles.MinW7xl())
	return c
}

func (c *component) MinWAuto() *component {
	c.el.Style(styles.MinWAuto())
	return c
}

func (c *component) MinWBy(val value.Value) *component {
	c.el.Style(styles.MinWBy(val))
	return c
}

func (c *component) MinWDvh() *component {
	c.el.Style(styles.MinWDvh())
	return c
}

func (c *component) MinWDvw() *component {
	c.el.Style(styles.MinWDvw())
	return c
}

func (c *component) MinWFit() *component {
	c.el.Style(styles.MinWFit())
	return c
}

func (c *component) MinWFraction(fraction float32) *component {
	c.el.Style(styles.MinWFraction(fraction))
	return c
}

func (c *component) MinWFull() *component {
	c.el.Style(styles.MinWFull())
	return c
}

func (c *component) MinWLg() *component {
	c.el.Style(styles.MinWLg())
	return c
}

func (c *component) MinWLvh() *component {
	c.el.Style(styles.MinWLvh())
	return c
}

func (c *component) MinWLvw() *component {
	c.el.Style(styles.MinWLvw())
	return c
}

func (c *component) MinWMax() *component {
	c.el.Style(styles.MinWMax())
	return c
}

func (c *component) MinWMd() *component {
	c.el.Style(styles.MinWMd())
	return c
}

func (c *component) MinWMin() *component {
	c.el.Style(styles.MinWMin())
	return c
}

func (c *component) MinWPx() *component {
	c.el.Style(styles.MinWPx())
	return c
}

func (c *component) MinWScreen() *component {
	c.el.Style(styles.MinWScreen())
	return c
}

func (c *component) MinWSm() *component {
	c.el.Style(styles.MinWSm())
	return c
}

func (c *component) MinWSvh() *component {
	c.el.Style(styles.MinWSvh())
	return c
}

func (c *component) MinWSvw() *component {
	c.el.Style(styles.MinWSvw())
	return c
}

func (c *component) MinWXl() *component {
	c.el.Style(styles.MinWXl())
	return c
}

func (c *component) MinWXs() *component {
	c.el.Style(styles.MinWXs())
	return c
}

func (c *component) MixBlendColor() *component {
	c.el.Style(styles.MixBlendColor())
	return c
}

func (c *component) MixBlendColorBurn() *component {
	c.el.Style(styles.MixBlendColorBurn())
	return c
}

func (c *component) MixBlendColorDodge() *component {
	c.el.Style(styles.MixBlendColorDodge())
	return c
}

func (c *component) MixBlendDarken() *component {
	c.el.Style(styles.MixBlendDarken())
	return c
}

func (c *component) MixBlendDifference() *component {
	c.el.Style(styles.MixBlendDifference())
	return c
}

func (c *component) MixBlendExclusion() *component {
	c.el.Style(styles.MixBlendExclusion())
	return c
}

func (c *component) MixBlendHardLight() *component {
	c.el.Style(styles.MixBlendHardLight())
	return c
}

func (c *component) MixBlendHue() *component {
	c.el.Style(styles.MixBlendHue())
	return c
}

func (c *component) MixBlendLighten() *component {
	c.el.Style(styles.MixBlendLighten())
	return c
}

func (c *component) MixBlendLuminosity() *component {
	c.el.Style(styles.MixBlendLuminosity())
	return c
}

func (c *component) MixBlendMultiply() *component {
	c.el.Style(styles.MixBlendMultiply())
	return c
}

func (c *component) MixBlendNormal() *component {
	c.el.Style(styles.MixBlendNormal())
	return c
}

func (c *component) MixBlendOverlay() *component {
	c.el.Style(styles.MixBlendOverlay())
	return c
}

func (c *component) MixBlendPlusDarker() *component {
	c.el.Style(styles.MixBlendPlusDarker())
	return c
}

func (c *component) MixBlendPlusLighter() *component {
	c.el.Style(styles.MixBlendPlusLighter())
	return c
}

func (c *component) MixBlendSaturation() *component {
	c.el.Style(styles.MixBlendSaturation())
	return c
}

func (c *component) MixBlendScreen() *component {
	c.el.Style(styles.MixBlendScreen())
	return c
}

func (c *component) MixBlendSoftLight() *component {
	c.el.Style(styles.MixBlendSoftLight())
	return c
}

func (c *component) Ml(number int) *component {
	c.el.Style(styles.Ml(number))
	return c
}

func (c *component) MlAuto() *component {
	c.el.Style(styles.MlAuto())
	return c
}

func (c *component) MlBy(val value.Value) *component {
	c.el.Style(styles.MlBy(val))
	return c
}

func (c *component) MlPx() *component {
	c.el.Style(styles.MlPx())
	return c
}

func (c *component) Mr(number int) *component {
	c.el.Style(styles.Mr(number))
	return c
}

func (c *component) MrAuto() *component {
	c.el.Style(styles.MrAuto())
	return c
}

func (c *component) MrBy(val value.Value) *component {
	c.el.Style(styles.MrBy(val))
	return c
}

func (c *component) MrPx() *component {
	c.el.Style(styles.MrPx())
	return c
}

func (c *component) Ms(number int) *component {
	c.el.Style(styles.Ms(number))
	return c
}

func (c *component) MsAuto() *component {
	c.el.Style(styles.MsAuto())
	return c
}

func (c *component) MsBy(val value.Value) *component {
	c.el.Style(styles.MsBy(val))
	return c
}

func (c *component) MsPx() *component {
	c.el.Style(styles.MsPx())
	return c
}

func (c *component) Mt(number int) *component {
	c.el.Style(styles.Mt(number))
	return c
}

func (c *component) MtAuto() *component {
	c.el.Style(styles.MtAuto())
	return c
}

func (c *component) MtBy(val value.Value) *component {
	c.el.Style(styles.MtBy(val))
	return c
}

func (c *component) MtPx() *component {
	c.el.Style(styles.MtPx())
	return c
}

func (c *component) Mx(number int) *component {
	c.el.Style(styles.Mx(number))
	return c
}

func (c *component) MxAuto() *component {
	c.el.Style(styles.MxAuto())
	return c
}

func (c *component) MxBy(val value.Value) *component {
	c.el.Style(styles.MxBy(val))
	return c
}

func (c *component) MxPx() *component {
	c.el.Style(styles.MxPx())
	return c
}

func (c *component) My(number int) *component {
	c.el.Style(styles.My(number))
	return c
}

func (c *component) MyAuto() *component {
	c.el.Style(styles.MyAuto())
	return c
}

func (c *component) MyBy(val value.Value) *component {
	c.el.Style(styles.MyBy(val))
	return c
}

func (c *component) MyPx() *component {
	c.el.Style(styles.MyPx())
	return c
}

func (c *component) NegBottom(number int) *component {
	c.el.Style(styles.NegBottom(number))
	return c
}

func (c *component) NegBottomFraction(fraction float64) *component {
	c.el.Style(styles.NegBottomFraction(fraction))
	return c
}

func (c *component) NegBottomFull() *component {
	c.el.Style(styles.NegBottomFull())
	return c
}

func (c *component) NegBottomPx() *component {
	c.el.Style(styles.NegBottomPx())
	return c
}

func (c *component) NegCol(number int) *component {
	c.el.Style(styles.NegCol(number))
	return c
}

func (c *component) NegColEnd(number int) *component {
	c.el.Style(styles.NegColEnd(number))
	return c
}

func (c *component) NegColStart(number int) *component {
	c.el.Style(styles.NegColStart(number))
	return c
}

func (c *component) NegEnd(number int) *component {
	c.el.Style(styles.NegEnd(number))
	return c
}

func (c *component) NegEndFraction(fraction float64) *component {
	c.el.Style(styles.NegEndFraction(fraction))
	return c
}

func (c *component) NegEndFull() *component {
	c.el.Style(styles.NegEndFull())
	return c
}

func (c *component) NegEndPx() *component {
	c.el.Style(styles.NegEndPx())
	return c
}

func (c *component) NegIndent(number int) *component {
	c.el.Style(styles.NegIndent(number))
	return c
}

func (c *component) NegIndentPx() *component {
	c.el.Style(styles.NegIndentPx())
	return c
}

func (c *component) NegInset(number int) *component {
	c.el.Style(styles.NegInset(number))
	return c
}

func (c *component) NegInsetFraction(fraction float64) *component {
	c.el.Style(styles.NegInsetFraction(fraction))
	return c
}

func (c *component) NegInsetFull() *component {
	c.el.Style(styles.NegInsetFull())
	return c
}

func (c *component) NegInsetPx() *component {
	c.el.Style(styles.NegInsetPx())
	return c
}

func (c *component) NegInsetX(number int) *component {
	c.el.Style(styles.NegInsetX(number))
	return c
}

func (c *component) NegInsetXFraction(fraction float64) *component {
	c.el.Style(styles.NegInsetXFraction(fraction))
	return c
}

func (c *component) NegInsetXFull() *component {
	c.el.Style(styles.NegInsetXFull())
	return c
}

func (c *component) NegInsetXPx() *component {
	c.el.Style(styles.NegInsetXPx())
	return c
}

func (c *component) NegInsetY(number int) *component {
	c.el.Style(styles.NegInsetY(number))
	return c
}

func (c *component) NegInsetYFraction(fraction float64) *component {
	c.el.Style(styles.NegInsetYFraction(fraction))
	return c
}

func (c *component) NegInsetYFull() *component {
	c.el.Style(styles.NegInsetYFull())
	return c
}

func (c *component) NegInsetYPx() *component {
	c.el.Style(styles.NegInsetYPx())
	return c
}

func (c *component) NegLeft(number int) *component {
	c.el.Style(styles.NegLeft(number))
	return c
}

func (c *component) NegLeftFraction(fraction float64) *component {
	c.el.Style(styles.NegLeftFraction(fraction))
	return c
}

func (c *component) NegLeftFull() *component {
	c.el.Style(styles.NegLeftFull())
	return c
}

func (c *component) NegLeftPx() *component {
	c.el.Style(styles.NegLeftPx())
	return c
}

func (c *component) NegM(number int) *component {
	c.el.Style(styles.NegM(number))
	return c
}

func (c *component) NegMPx() *component {
	c.el.Style(styles.NegMPx())
	return c
}

func (c *component) NegMb(number int) *component {
	c.el.Style(styles.NegMb(number))
	return c
}

func (c *component) NegMbPx() *component {
	c.el.Style(styles.NegMbPx())
	return c
}

func (c *component) NegMe(number int) *component {
	c.el.Style(styles.NegMe(number))
	return c
}

func (c *component) NegMePx() *component {
	c.el.Style(styles.NegMePx())
	return c
}

func (c *component) NegMl(number int) *component {
	c.el.Style(styles.NegMl(number))
	return c
}

func (c *component) NegMlPx() *component {
	c.el.Style(styles.NegMlPx())
	return c
}

func (c *component) NegMr(number int) *component {
	c.el.Style(styles.NegMr(number))
	return c
}

func (c *component) NegMrPx() *component {
	c.el.Style(styles.NegMrPx())
	return c
}

func (c *component) NegMs(number int) *component {
	c.el.Style(styles.NegMs(number))
	return c
}

func (c *component) NegMsPx() *component {
	c.el.Style(styles.NegMsPx())
	return c
}

func (c *component) NegMt(number int) *component {
	c.el.Style(styles.NegMt(number))
	return c
}

func (c *component) NegMtPx() *component {
	c.el.Style(styles.NegMtPx())
	return c
}

func (c *component) NegMx(number int) *component {
	c.el.Style(styles.NegMx(number))
	return c
}

func (c *component) NegMxPx() *component {
	c.el.Style(styles.NegMxPx())
	return c
}

func (c *component) NegMy(number int) *component {
	c.el.Style(styles.NegMy(number))
	return c
}

func (c *component) NegMyPx() *component {
	c.el.Style(styles.NegMyPx())
	return c
}

func (c *component) NegOrder(number int) *component {
	c.el.Style(styles.NegOrder(number))
	return c
}

func (c *component) NegRight(number int) *component {
	c.el.Style(styles.NegRight(number))
	return c
}

func (c *component) NegRightFraction(fraction float64) *component {
	c.el.Style(styles.NegRightFraction(fraction))
	return c
}

func (c *component) NegRightFull() *component {
	c.el.Style(styles.NegRightFull())
	return c
}

func (c *component) NegRightPx() *component {
	c.el.Style(styles.NegRightPx())
	return c
}

func (c *component) NegRow(number int) *component {
	c.el.Style(styles.NegRow(number))
	return c
}

func (c *component) NegRowEnd(number int) *component {
	c.el.Style(styles.NegRowEnd(number))
	return c
}

func (c *component) NegRowStart(number int) *component {
	c.el.Style(styles.NegRowStart(number))
	return c
}

func (c *component) NegStart(number int) *component {
	c.el.Style(styles.NegStart(number))
	return c
}

func (c *component) NegStartFraction(fraction float64) *component {
	c.el.Style(styles.NegStartFraction(fraction))
	return c
}

func (c *component) NegStartFull() *component {
	c.el.Style(styles.NegStartFull())
	return c
}

func (c *component) NegStartPx() *component {
	c.el.Style(styles.NegStartPx())
	return c
}

func (c *component) NegTop(number int) *component {
	c.el.Style(styles.NegTop(number))
	return c
}

func (c *component) NegTopFraction(fraction float64) *component {
	c.el.Style(styles.NegTopFraction(fraction))
	return c
}

func (c *component) NegTopFull() *component {
	c.el.Style(styles.NegTopFull())
	return c
}

func (c *component) NegTopPx() *component {
	c.el.Style(styles.NegTopPx())
	return c
}

func (c *component) NegTranslate(val any) *component {
	c.el.Style(styles.NegTranslate(val))
	return c
}

func (c *component) NegTranslateFull() *component {
	c.el.Style(styles.NegTranslateFull())
	return c
}

func (c *component) NegTranslatePx() *component {
	c.el.Style(styles.NegTranslatePx())
	return c
}

func (c *component) NegTranslateX(val any) *component {
	c.el.Style(styles.NegTranslateX(val))
	return c
}

func (c *component) NegTranslateXFull() *component {
	c.el.Style(styles.NegTranslateXFull())
	return c
}

func (c *component) NegTranslateXPx() *component {
	c.el.Style(styles.NegTranslateXPx())
	return c
}

func (c *component) NegTranslateY(val any) *component {
	c.el.Style(styles.NegTranslateY(val))
	return c
}

func (c *component) NegTranslateYFull() *component {
	c.el.Style(styles.NegTranslateYFull())
	return c
}

func (c *component) NegTranslateYPx() *component {
	c.el.Style(styles.NegTranslateYPx())
	return c
}

func (c *component) NegTranslateZ(val any) *component {
	c.el.Style(styles.NegTranslateZ(val))
	return c
}

func (c *component) NegTranslateZPx() *component {
	c.el.Style(styles.NegTranslateZPx())
	return c
}

func (c *component) NegUnderlineOffset(number int) *component {
	c.el.Style(styles.NegUnderlineOffset(number))
	return c
}

func (c *component) NoUnderline() *component {
	c.el.Style(styles.NoUnderline())
	return c
}

func (c *component) NormalCase() *component {
	c.el.Style(styles.NormalCase())
	return c
}

func (c *component) NormalNums() *component {
	c.el.Style(styles.NormalNums())
	return c
}

func (c *component) NotItalic() *component {
	c.el.Style(styles.NotItalic())
	return c
}

func (c *component) NotSrOnly() *component {
	c.el.Style(styles.NotSrOnly())
	return c
}

func (c *component) ObjectBottom() *component {
	c.el.Style(styles.ObjectBottom())
	return c
}

func (c *component) ObjectBottomLeft() *component {
	c.el.Style(styles.ObjectBottomLeft())
	return c
}

func (c *component) ObjectBottomRight() *component {
	c.el.Style(styles.ObjectBottomRight())
	return c
}

func (c *component) ObjectBy(val value.Value) *component {
	c.el.Style(styles.ObjectBy(val))
	return c
}

func (c *component) ObjectCenter() *component {
	c.el.Style(styles.ObjectCenter())
	return c
}

func (c *component) ObjectContain() *component {
	c.el.Style(styles.ObjectContain())
	return c
}

func (c *component) ObjectCover() *component {
	c.el.Style(styles.ObjectCover())
	return c
}

func (c *component) ObjectFill() *component {
	c.el.Style(styles.ObjectFill())
	return c
}

func (c *component) ObjectLeft() *component {
	c.el.Style(styles.ObjectLeft())
	return c
}

func (c *component) ObjectNone() *component {
	c.el.Style(styles.ObjectNone())
	return c
}

func (c *component) ObjectRight() *component {
	c.el.Style(styles.ObjectRight())
	return c
}

func (c *component) ObjectScaleDown() *component {
	c.el.Style(styles.ObjectScaleDown())
	return c
}

func (c *component) ObjectTop() *component {
	c.el.Style(styles.ObjectTop())
	return c
}

func (c *component) ObjectTopLeft() *component {
	c.el.Style(styles.ObjectTopLeft())
	return c
}

func (c *component) ObjectTopRight() *component {
	c.el.Style(styles.ObjectTopRight())
	return c
}

func (c *component) OldStyleNums() *component {
	c.el.Style(styles.OldStyleNums())
	return c
}

func (c *component) Opacity(val value.Value) *component {
	c.el.Style(styles.Opacity(val))
	return c
}

func (c *component) Order(number int) *component {
	c.el.Style(styles.Order(number))
	return c
}

func (c *component) OrderBy(val value.Value) *component {
	c.el.Style(styles.OrderBy(val))
	return c
}

func (c *component) OrderFirst() *component {
	c.el.Style(styles.OrderFirst())
	return c
}

func (c *component) OrderLast() *component {
	c.el.Style(styles.OrderLast())
	return c
}

func (c *component) Ordinal() *component {
	c.el.Style(styles.Ordinal())
	return c
}

func (c *component) Origin(val value.Value) *component {
	c.el.Style(styles.Origin(val))
	return c
}

func (c *component) OriginBottom() *component {
	c.el.Style(styles.OriginBottom())
	return c
}

func (c *component) OriginBottomLeft() *component {
	c.el.Style(styles.OriginBottomLeft())
	return c
}

func (c *component) OriginBottomRight() *component {
	c.el.Style(styles.OriginBottomRight())
	return c
}

func (c *component) OriginCenter() *component {
	c.el.Style(styles.OriginCenter())
	return c
}

func (c *component) OriginLeft() *component {
	c.el.Style(styles.OriginLeft())
	return c
}

func (c *component) OriginRight() *component {
	c.el.Style(styles.OriginRight())
	return c
}

func (c *component) OriginTop() *component {
	c.el.Style(styles.OriginTop())
	return c
}

func (c *component) OriginTopLeft() *component {
	c.el.Style(styles.OriginTopLeft())
	return c
}

func (c *component) OriginTopRight() *component {
	c.el.Style(styles.OriginTopRight())
	return c
}

func (c *component) Outline(val ...value.Value) *component {
	c.el.Style(styles.Outline(val...))
	return c
}

func (c *component) OutlineAmber(scale int) *component {
	c.el.Style(styles.OutlineAmber(scale))
	return c
}

func (c *component) OutlineAmberAlpha(scale int) *component {
	c.el.Style(styles.OutlineAmberAlpha(scale))
	return c
}

func (c *component) OutlineAmberDark(scale int) *component {
	c.el.Style(styles.OutlineAmberDark(scale))
	return c
}

func (c *component) OutlineAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineAmberDarkAlpha(scale))
	return c
}

func (c *component) OutlineBlack() *component {
	c.el.Style(styles.OutlineBlack())
	return c
}

func (c *component) OutlineBlackAlpha(scale int) *component {
	c.el.Style(styles.OutlineBlackAlpha(scale))
	return c
}

func (c *component) OutlineBlue(scale int) *component {
	c.el.Style(styles.OutlineBlue(scale))
	return c
}

func (c *component) OutlineBlueAlpha(scale int) *component {
	c.el.Style(styles.OutlineBlueAlpha(scale))
	return c
}

func (c *component) OutlineBlueDark(scale int) *component {
	c.el.Style(styles.OutlineBlueDark(scale))
	return c
}

func (c *component) OutlineBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineBlueDarkAlpha(scale))
	return c
}

func (c *component) OutlineBronze(scale int) *component {
	c.el.Style(styles.OutlineBronze(scale))
	return c
}

func (c *component) OutlineBronzeAlpha(scale int) *component {
	c.el.Style(styles.OutlineBronzeAlpha(scale))
	return c
}

func (c *component) OutlineBronzeDark(scale int) *component {
	c.el.Style(styles.OutlineBronzeDark(scale))
	return c
}

func (c *component) OutlineBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineBronzeDarkAlpha(scale))
	return c
}

func (c *component) OutlineBrown(scale int) *component {
	c.el.Style(styles.OutlineBrown(scale))
	return c
}

func (c *component) OutlineBrownAlpha(scale int) *component {
	c.el.Style(styles.OutlineBrownAlpha(scale))
	return c
}

func (c *component) OutlineBrownDark(scale int) *component {
	c.el.Style(styles.OutlineBrownDark(scale))
	return c
}

func (c *component) OutlineBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineBrownDarkAlpha(scale))
	return c
}

func (c *component) OutlineColor(val value.Value) *component {
	c.el.Style(styles.OutlineColor(val))
	return c
}

func (c *component) OutlineCrimson(scale int) *component {
	c.el.Style(styles.OutlineCrimson(scale))
	return c
}

func (c *component) OutlineCrimsonAlpha(scale int) *component {
	c.el.Style(styles.OutlineCrimsonAlpha(scale))
	return c
}

func (c *component) OutlineCrimsonDark(scale int) *component {
	c.el.Style(styles.OutlineCrimsonDark(scale))
	return c
}

func (c *component) OutlineCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineCrimsonDarkAlpha(scale))
	return c
}

func (c *component) OutlineCurrent() *component {
	c.el.Style(styles.OutlineCurrent())
	return c
}

func (c *component) OutlineCyan(scale int) *component {
	c.el.Style(styles.OutlineCyan(scale))
	return c
}

func (c *component) OutlineCyanAlpha(scale int) *component {
	c.el.Style(styles.OutlineCyanAlpha(scale))
	return c
}

func (c *component) OutlineCyanDark(scale int) *component {
	c.el.Style(styles.OutlineCyanDark(scale))
	return c
}

func (c *component) OutlineCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineCyanDarkAlpha(scale))
	return c
}

func (c *component) OutlineDashed() *component {
	c.el.Style(styles.OutlineDashed())
	return c
}

func (c *component) OutlineDotted() *component {
	c.el.Style(styles.OutlineDotted())
	return c
}

func (c *component) OutlineDouble() *component {
	c.el.Style(styles.OutlineDouble())
	return c
}

func (c *component) OutlineGold(scale int) *component {
	c.el.Style(styles.OutlineGold(scale))
	return c
}

func (c *component) OutlineGoldAlpha(scale int) *component {
	c.el.Style(styles.OutlineGoldAlpha(scale))
	return c
}

func (c *component) OutlineGoldDark(scale int) *component {
	c.el.Style(styles.OutlineGoldDark(scale))
	return c
}

func (c *component) OutlineGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineGoldDarkAlpha(scale))
	return c
}

func (c *component) OutlineGrass(scale int) *component {
	c.el.Style(styles.OutlineGrass(scale))
	return c
}

func (c *component) OutlineGrassAlpha(scale int) *component {
	c.el.Style(styles.OutlineGrassAlpha(scale))
	return c
}

func (c *component) OutlineGrassDark(scale int) *component {
	c.el.Style(styles.OutlineGrassDark(scale))
	return c
}

func (c *component) OutlineGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineGrassDarkAlpha(scale))
	return c
}

func (c *component) OutlineGray(scale int) *component {
	c.el.Style(styles.OutlineGray(scale))
	return c
}

func (c *component) OutlineGrayAlpha(scale int) *component {
	c.el.Style(styles.OutlineGrayAlpha(scale))
	return c
}

func (c *component) OutlineGrayDark(scale int) *component {
	c.el.Style(styles.OutlineGrayDark(scale))
	return c
}

func (c *component) OutlineGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineGrayDarkAlpha(scale))
	return c
}

func (c *component) OutlineGreen(scale int) *component {
	c.el.Style(styles.OutlineGreen(scale))
	return c
}

func (c *component) OutlineGreenAlpha(scale int) *component {
	c.el.Style(styles.OutlineGreenAlpha(scale))
	return c
}

func (c *component) OutlineGreenDark(scale int) *component {
	c.el.Style(styles.OutlineGreenDark(scale))
	return c
}

func (c *component) OutlineGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineGreenDarkAlpha(scale))
	return c
}

func (c *component) OutlineHidden() *component {
	c.el.Style(styles.OutlineHidden())
	return c
}

func (c *component) OutlineIndigo(scale int) *component {
	c.el.Style(styles.OutlineIndigo(scale))
	return c
}

func (c *component) OutlineIndigoAlpha(scale int) *component {
	c.el.Style(styles.OutlineIndigoAlpha(scale))
	return c
}

func (c *component) OutlineIndigoDark(scale int) *component {
	c.el.Style(styles.OutlineIndigoDark(scale))
	return c
}

func (c *component) OutlineIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineIndigoDarkAlpha(scale))
	return c
}

func (c *component) OutlineInherit() *component {
	c.el.Style(styles.OutlineInherit())
	return c
}

func (c *component) OutlineIris(scale int) *component {
	c.el.Style(styles.OutlineIris(scale))
	return c
}

func (c *component) OutlineIrisAlpha(scale int) *component {
	c.el.Style(styles.OutlineIrisAlpha(scale))
	return c
}

func (c *component) OutlineIrisDark(scale int) *component {
	c.el.Style(styles.OutlineIrisDark(scale))
	return c
}

func (c *component) OutlineIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineIrisDarkAlpha(scale))
	return c
}

func (c *component) OutlineJade(scale int) *component {
	c.el.Style(styles.OutlineJade(scale))
	return c
}

func (c *component) OutlineJadeAlpha(scale int) *component {
	c.el.Style(styles.OutlineJadeAlpha(scale))
	return c
}

func (c *component) OutlineJadeDark(scale int) *component {
	c.el.Style(styles.OutlineJadeDark(scale))
	return c
}

func (c *component) OutlineJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineJadeDarkAlpha(scale))
	return c
}

func (c *component) OutlineLime(scale int) *component {
	c.el.Style(styles.OutlineLime(scale))
	return c
}

func (c *component) OutlineLimeAlpha(scale int) *component {
	c.el.Style(styles.OutlineLimeAlpha(scale))
	return c
}

func (c *component) OutlineLimeDark(scale int) *component {
	c.el.Style(styles.OutlineLimeDark(scale))
	return c
}

func (c *component) OutlineLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineLimeDarkAlpha(scale))
	return c
}

func (c *component) OutlineMauve(scale int) *component {
	c.el.Style(styles.OutlineMauve(scale))
	return c
}

func (c *component) OutlineMauveAlpha(scale int) *component {
	c.el.Style(styles.OutlineMauveAlpha(scale))
	return c
}

func (c *component) OutlineMauveDark(scale int) *component {
	c.el.Style(styles.OutlineMauveDark(scale))
	return c
}

func (c *component) OutlineMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineMauveDarkAlpha(scale))
	return c
}

func (c *component) OutlineMint(scale int) *component {
	c.el.Style(styles.OutlineMint(scale))
	return c
}

func (c *component) OutlineMintAlpha(scale int) *component {
	c.el.Style(styles.OutlineMintAlpha(scale))
	return c
}

func (c *component) OutlineMintDark(scale int) *component {
	c.el.Style(styles.OutlineMintDark(scale))
	return c
}

func (c *component) OutlineMintDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineMintDarkAlpha(scale))
	return c
}

func (c *component) OutlineNone() *component {
	c.el.Style(styles.OutlineNone())
	return c
}

func (c *component) OutlineOffset(val ...value.Value) *component {
	c.el.Style(styles.OutlineOffset(val...))
	return c
}

func (c *component) OutlineOlive(scale int) *component {
	c.el.Style(styles.OutlineOlive(scale))
	return c
}

func (c *component) OutlineOliveAlpha(scale int) *component {
	c.el.Style(styles.OutlineOliveAlpha(scale))
	return c
}

func (c *component) OutlineOliveDark(scale int) *component {
	c.el.Style(styles.OutlineOliveDark(scale))
	return c
}

func (c *component) OutlineOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineOliveDarkAlpha(scale))
	return c
}

func (c *component) OutlineOrange(scale int) *component {
	c.el.Style(styles.OutlineOrange(scale))
	return c
}

func (c *component) OutlineOrangeAlpha(scale int) *component {
	c.el.Style(styles.OutlineOrangeAlpha(scale))
	return c
}

func (c *component) OutlineOrangeDark(scale int) *component {
	c.el.Style(styles.OutlineOrangeDark(scale))
	return c
}

func (c *component) OutlineOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineOrangeDarkAlpha(scale))
	return c
}

func (c *component) OutlinePink(scale int) *component {
	c.el.Style(styles.OutlinePink(scale))
	return c
}

func (c *component) OutlinePinkAlpha(scale int) *component {
	c.el.Style(styles.OutlinePinkAlpha(scale))
	return c
}

func (c *component) OutlinePinkDark(scale int) *component {
	c.el.Style(styles.OutlinePinkDark(scale))
	return c
}

func (c *component) OutlinePinkDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlinePinkDarkAlpha(scale))
	return c
}

func (c *component) OutlinePlum(scale int) *component {
	c.el.Style(styles.OutlinePlum(scale))
	return c
}

func (c *component) OutlinePlumAlpha(scale int) *component {
	c.el.Style(styles.OutlinePlumAlpha(scale))
	return c
}

func (c *component) OutlinePlumDark(scale int) *component {
	c.el.Style(styles.OutlinePlumDark(scale))
	return c
}

func (c *component) OutlinePlumDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlinePlumDarkAlpha(scale))
	return c
}

func (c *component) OutlinePurple(scale int) *component {
	c.el.Style(styles.OutlinePurple(scale))
	return c
}

func (c *component) OutlinePurpleAlpha(scale int) *component {
	c.el.Style(styles.OutlinePurpleAlpha(scale))
	return c
}

func (c *component) OutlinePurpleDark(scale int) *component {
	c.el.Style(styles.OutlinePurpleDark(scale))
	return c
}

func (c *component) OutlinePurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlinePurpleDarkAlpha(scale))
	return c
}

func (c *component) OutlineRed(scale int) *component {
	c.el.Style(styles.OutlineRed(scale))
	return c
}

func (c *component) OutlineRedAlpha(scale int) *component {
	c.el.Style(styles.OutlineRedAlpha(scale))
	return c
}

func (c *component) OutlineRedDark(scale int) *component {
	c.el.Style(styles.OutlineRedDark(scale))
	return c
}

func (c *component) OutlineRedDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineRedDarkAlpha(scale))
	return c
}

func (c *component) OutlineRuby(scale int) *component {
	c.el.Style(styles.OutlineRuby(scale))
	return c
}

func (c *component) OutlineRubyAlpha(scale int) *component {
	c.el.Style(styles.OutlineRubyAlpha(scale))
	return c
}

func (c *component) OutlineRubyDark(scale int) *component {
	c.el.Style(styles.OutlineRubyDark(scale))
	return c
}

func (c *component) OutlineRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineRubyDarkAlpha(scale))
	return c
}

func (c *component) OutlineSage(scale int) *component {
	c.el.Style(styles.OutlineSage(scale))
	return c
}

func (c *component) OutlineSageAlpha(scale int) *component {
	c.el.Style(styles.OutlineSageAlpha(scale))
	return c
}

func (c *component) OutlineSageDark(scale int) *component {
	c.el.Style(styles.OutlineSageDark(scale))
	return c
}

func (c *component) OutlineSageDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineSageDarkAlpha(scale))
	return c
}

func (c *component) OutlineSand(scale int) *component {
	c.el.Style(styles.OutlineSand(scale))
	return c
}

func (c *component) OutlineSandAlpha(scale int) *component {
	c.el.Style(styles.OutlineSandAlpha(scale))
	return c
}

func (c *component) OutlineSandDark(scale int) *component {
	c.el.Style(styles.OutlineSandDark(scale))
	return c
}

func (c *component) OutlineSandDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineSandDarkAlpha(scale))
	return c
}

func (c *component) OutlineSky(scale int) *component {
	c.el.Style(styles.OutlineSky(scale))
	return c
}

func (c *component) OutlineSkyAlpha(scale int) *component {
	c.el.Style(styles.OutlineSkyAlpha(scale))
	return c
}

func (c *component) OutlineSkyDark(scale int) *component {
	c.el.Style(styles.OutlineSkyDark(scale))
	return c
}

func (c *component) OutlineSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineSkyDarkAlpha(scale))
	return c
}

func (c *component) OutlineSlate(scale int) *component {
	c.el.Style(styles.OutlineSlate(scale))
	return c
}

func (c *component) OutlineSlateAlpha(scale int) *component {
	c.el.Style(styles.OutlineSlateAlpha(scale))
	return c
}

func (c *component) OutlineSlateDark(scale int) *component {
	c.el.Style(styles.OutlineSlateDark(scale))
	return c
}

func (c *component) OutlineSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineSlateDarkAlpha(scale))
	return c
}

func (c *component) OutlineSolid() *component {
	c.el.Style(styles.OutlineSolid())
	return c
}

func (c *component) OutlineTeal(scale int) *component {
	c.el.Style(styles.OutlineTeal(scale))
	return c
}

func (c *component) OutlineTealAlpha(scale int) *component {
	c.el.Style(styles.OutlineTealAlpha(scale))
	return c
}

func (c *component) OutlineTealDark(scale int) *component {
	c.el.Style(styles.OutlineTealDark(scale))
	return c
}

func (c *component) OutlineTealDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineTealDarkAlpha(scale))
	return c
}

func (c *component) OutlineTomato(scale int) *component {
	c.el.Style(styles.OutlineTomato(scale))
	return c
}

func (c *component) OutlineTomatoAlpha(scale int) *component {
	c.el.Style(styles.OutlineTomatoAlpha(scale))
	return c
}

func (c *component) OutlineTomatoDark(scale int) *component {
	c.el.Style(styles.OutlineTomatoDark(scale))
	return c
}

func (c *component) OutlineTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineTomatoDarkAlpha(scale))
	return c
}

func (c *component) OutlineTransparent() *component {
	c.el.Style(styles.OutlineTransparent())
	return c
}

func (c *component) OutlineViolet(scale int) *component {
	c.el.Style(styles.OutlineViolet(scale))
	return c
}

func (c *component) OutlineVioletAlpha(scale int) *component {
	c.el.Style(styles.OutlineVioletAlpha(scale))
	return c
}

func (c *component) OutlineVioletDark(scale int) *component {
	c.el.Style(styles.OutlineVioletDark(scale))
	return c
}

func (c *component) OutlineVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineVioletDarkAlpha(scale))
	return c
}

func (c *component) OutlineWhite() *component {
	c.el.Style(styles.OutlineWhite())
	return c
}

func (c *component) OutlineWhiteAlpha(scale int) *component {
	c.el.Style(styles.OutlineWhiteAlpha(scale))
	return c
}

func (c *component) OutlineYellow(scale int) *component {
	c.el.Style(styles.OutlineYellow(scale))
	return c
}

func (c *component) OutlineYellowAlpha(scale int) *component {
	c.el.Style(styles.OutlineYellowAlpha(scale))
	return c
}

func (c *component) OutlineYellowDark(scale int) *component {
	c.el.Style(styles.OutlineYellowDark(scale))
	return c
}

func (c *component) OutlineYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.OutlineYellowDarkAlpha(scale))
	return c
}

func (c *component) OverflowAuto() *component {
	c.el.Style(styles.OverflowAuto())
	return c
}

func (c *component) OverflowClip() *component {
	c.el.Style(styles.OverflowClip())
	return c
}

func (c *component) OverflowHidden() *component {
	c.el.Style(styles.OverflowHidden())
	return c
}

func (c *component) OverflowScroll() *component {
	c.el.Style(styles.OverflowScroll())
	return c
}

func (c *component) OverflowVisible() *component {
	c.el.Style(styles.OverflowVisible())
	return c
}

func (c *component) OverflowXAuto() *component {
	c.el.Style(styles.OverflowXAuto())
	return c
}

func (c *component) OverflowXClip() *component {
	c.el.Style(styles.OverflowXClip())
	return c
}

func (c *component) OverflowXHidden() *component {
	c.el.Style(styles.OverflowXHidden())
	return c
}

func (c *component) OverflowXScroll() *component {
	c.el.Style(styles.OverflowXScroll())
	return c
}

func (c *component) OverflowXVisible() *component {
	c.el.Style(styles.OverflowXVisible())
	return c
}

func (c *component) OverflowYAuto() *component {
	c.el.Style(styles.OverflowYAuto())
	return c
}

func (c *component) OverflowYClip() *component {
	c.el.Style(styles.OverflowYClip())
	return c
}

func (c *component) OverflowYHidden() *component {
	c.el.Style(styles.OverflowYHidden())
	return c
}

func (c *component) OverflowYScroll() *component {
	c.el.Style(styles.OverflowYScroll())
	return c
}

func (c *component) OverflowYVisible() *component {
	c.el.Style(styles.OverflowYVisible())
	return c
}

func (c *component) Overline() *component {
	c.el.Style(styles.Overline())
	return c
}

func (c *component) OverscrollAuto() *component {
	c.el.Style(styles.OverscrollAuto())
	return c
}

func (c *component) OverscrollContain() *component {
	c.el.Style(styles.OverscrollContain())
	return c
}

func (c *component) OverscrollNone() *component {
	c.el.Style(styles.OverscrollNone())
	return c
}

func (c *component) OverscrollXAuto() *component {
	c.el.Style(styles.OverscrollXAuto())
	return c
}

func (c *component) OverscrollXContain() *component {
	c.el.Style(styles.OverscrollXContain())
	return c
}

func (c *component) OverscrollXNone() *component {
	c.el.Style(styles.OverscrollXNone())
	return c
}

func (c *component) OverscrollYAuto() *component {
	c.el.Style(styles.OverscrollYAuto())
	return c
}

func (c *component) OverscrollYContain() *component {
	c.el.Style(styles.OverscrollYContain())
	return c
}

func (c *component) OverscrollYNone() *component {
	c.el.Style(styles.OverscrollYNone())
	return c
}

func (c *component) P(number int) *component {
	c.el.Style(styles.P(number))
	return c
}

func (c *component) PBy(val value.Value) *component {
	c.el.Style(styles.PBy(val))
	return c
}

func (c *component) PPx(number int) *component {
	c.el.Style(styles.PPx(number))
	return c
}

func (c *component) Pb(number int) *component {
	c.el.Style(styles.Pb(number))
	return c
}

func (c *component) PbBy(val value.Value) *component {
	c.el.Style(styles.PbBy(val))
	return c
}

func (c *component) PbPx(number int) *component {
	c.el.Style(styles.PbPx(number))
	return c
}

func (c *component) Pe(number int) *component {
	c.el.Style(styles.Pe(number))
	return c
}

func (c *component) PeBy(val value.Value) *component {
	c.el.Style(styles.PeBy(val))
	return c
}

func (c *component) PePx(number int) *component {
	c.el.Style(styles.PePx(number))
	return c
}

func (c *component) Perspective(val value.Value) *component {
	c.el.Style(styles.Perspective(val))
	return c
}

func (c *component) PerspectiveDistant() *component {
	c.el.Style(styles.PerspectiveDistant())
	return c
}

func (c *component) PerspectiveDramatic() *component {
	c.el.Style(styles.PerspectiveDramatic())
	return c
}

func (c *component) PerspectiveMidrange() *component {
	c.el.Style(styles.PerspectiveMidrange())
	return c
}

func (c *component) PerspectiveNear() *component {
	c.el.Style(styles.PerspectiveNear())
	return c
}

func (c *component) PerspectiveNone() *component {
	c.el.Style(styles.PerspectiveNone())
	return c
}

func (c *component) PerspectiveNormal() *component {
	c.el.Style(styles.PerspectiveNormal())
	return c
}

func (c *component) PerspectiveOrigin(val value.Value) *component {
	c.el.Style(styles.PerspectiveOrigin(val))
	return c
}

func (c *component) PerspectiveOriginBottom() *component {
	c.el.Style(styles.PerspectiveOriginBottom())
	return c
}

func (c *component) PerspectiveOriginBottomLeft() *component {
	c.el.Style(styles.PerspectiveOriginBottomLeft())
	return c
}

func (c *component) PerspectiveOriginBottomRight() *component {
	c.el.Style(styles.PerspectiveOriginBottomRight())
	return c
}

func (c *component) PerspectiveOriginCenter() *component {
	c.el.Style(styles.PerspectiveOriginCenter())
	return c
}

func (c *component) PerspectiveOriginLeft() *component {
	c.el.Style(styles.PerspectiveOriginLeft())
	return c
}

func (c *component) PerspectiveOriginRight() *component {
	c.el.Style(styles.PerspectiveOriginRight())
	return c
}

func (c *component) PerspectiveOriginTop() *component {
	c.el.Style(styles.PerspectiveOriginTop())
	return c
}

func (c *component) PerspectiveOriginTopLeft() *component {
	c.el.Style(styles.PerspectiveOriginTopLeft())
	return c
}

func (c *component) PerspectiveOriginTopRight() *component {
	c.el.Style(styles.PerspectiveOriginTopRight())
	return c
}

func (c *component) Pl(number int) *component {
	c.el.Style(styles.Pl(number))
	return c
}

func (c *component) PlBy(val value.Value) *component {
	c.el.Style(styles.PlBy(val))
	return c
}

func (c *component) PlPx(number int) *component {
	c.el.Style(styles.PlPx(number))
	return c
}

func (c *component) PlaceContentAround() *component {
	c.el.Style(styles.PlaceContentAround())
	return c
}

func (c *component) PlaceContentBaseline() *component {
	c.el.Style(styles.PlaceContentBaseline())
	return c
}

func (c *component) PlaceContentBetween() *component {
	c.el.Style(styles.PlaceContentBetween())
	return c
}

func (c *component) PlaceContentCenter() *component {
	c.el.Style(styles.PlaceContentCenter())
	return c
}

func (c *component) PlaceContentCenterSafe() *component {
	c.el.Style(styles.PlaceContentCenterSafe())
	return c
}

func (c *component) PlaceContentEnd() *component {
	c.el.Style(styles.PlaceContentEnd())
	return c
}

func (c *component) PlaceContentEndSafe() *component {
	c.el.Style(styles.PlaceContentEndSafe())
	return c
}

func (c *component) PlaceContentEvenly() *component {
	c.el.Style(styles.PlaceContentEvenly())
	return c
}

func (c *component) PlaceContentStart() *component {
	c.el.Style(styles.PlaceContentStart())
	return c
}

func (c *component) PlaceContentStretch() *component {
	c.el.Style(styles.PlaceContentStretch())
	return c
}

func (c *component) PlaceItemsBaseline() *component {
	c.el.Style(styles.PlaceItemsBaseline())
	return c
}

func (c *component) PlaceItemsCenter() *component {
	c.el.Style(styles.PlaceItemsCenter())
	return c
}

func (c *component) PlaceItemsCenterSafe() *component {
	c.el.Style(styles.PlaceItemsCenterSafe())
	return c
}

func (c *component) PlaceItemsEnd() *component {
	c.el.Style(styles.PlaceItemsEnd())
	return c
}

func (c *component) PlaceItemsEndSafe() *component {
	c.el.Style(styles.PlaceItemsEndSafe())
	return c
}

func (c *component) PlaceItemsStart() *component {
	c.el.Style(styles.PlaceItemsStart())
	return c
}

func (c *component) PlaceItemsStretch() *component {
	c.el.Style(styles.PlaceItemsStretch())
	return c
}

func (c *component) PlaceSelfAuto() *component {
	c.el.Style(styles.PlaceSelfAuto())
	return c
}

func (c *component) PlaceSelfCenter() *component {
	c.el.Style(styles.PlaceSelfCenter())
	return c
}

func (c *component) PlaceSelfCenterSafe() *component {
	c.el.Style(styles.PlaceSelfCenterSafe())
	return c
}

func (c *component) PlaceSelfEnd() *component {
	c.el.Style(styles.PlaceSelfEnd())
	return c
}

func (c *component) PlaceSelfEndSafe() *component {
	c.el.Style(styles.PlaceSelfEndSafe())
	return c
}

func (c *component) PlaceSelfStart() *component {
	c.el.Style(styles.PlaceSelfStart())
	return c
}

func (c *component) PlaceSelfStretch() *component {
	c.el.Style(styles.PlaceSelfStretch())
	return c
}

func (c *component) PointerEventsAuto() *component {
	c.el.Style(styles.PointerEventsAuto())
	return c
}

func (c *component) PointerEventsNone() *component {
	c.el.Style(styles.PointerEventsNone())
	return c
}

func (c *component) Pr(number int) *component {
	c.el.Style(styles.Pr(number))
	return c
}

func (c *component) PrBy(val value.Value) *component {
	c.el.Style(styles.PrBy(val))
	return c
}

func (c *component) PrPx(number int) *component {
	c.el.Style(styles.PrPx(number))
	return c
}

func (c *component) ProportionalNums() *component {
	c.el.Style(styles.ProportionalNums())
	return c
}

func (c *component) Ps(number int) *component {
	c.el.Style(styles.Ps(number))
	return c
}

func (c *component) PsBy(val value.Value) *component {
	c.el.Style(styles.PsBy(val))
	return c
}

func (c *component) PsPx(number int) *component {
	c.el.Style(styles.PsPx(number))
	return c
}

func (c *component) Pt(number int) *component {
	c.el.Style(styles.Pt(number))
	return c
}

func (c *component) PtBy(val value.Value) *component {
	c.el.Style(styles.PtBy(val))
	return c
}

func (c *component) PtPx(number int) *component {
	c.el.Style(styles.PtPx(number))
	return c
}

func (c *component) Px(number int) *component {
	c.el.Style(styles.Px(number))
	return c
}

func (c *component) PxBy(val value.Value) *component {
	c.el.Style(styles.PxBy(val))
	return c
}

func (c *component) PxPx(number int) *component {
	c.el.Style(styles.PxPx(number))
	return c
}

func (c *component) Py(number int) *component {
	c.el.Style(styles.Py(number))
	return c
}

func (c *component) PyBy(val value.Value) *component {
	c.el.Style(styles.PyBy(val))
	return c
}

func (c *component) PyPx(number int) *component {
	c.el.Style(styles.PyPx(number))
	return c
}

func (c *component) Relative() *component {
	c.el.Style(styles.Relative())
	return c
}

func (c *component) Resize() *component {
	c.el.Style(styles.Resize())
	return c
}

func (c *component) ResizeNone() *component {
	c.el.Style(styles.ResizeNone())
	return c
}

func (c *component) ResizeX() *component {
	c.el.Style(styles.ResizeX())
	return c
}

func (c *component) ResizeY() *component {
	c.el.Style(styles.ResizeY())
	return c
}

func (c *component) Right(number int) *component {
	c.el.Style(styles.Right(number))
	return c
}

func (c *component) RightAuto() *component {
	c.el.Style(styles.RightAuto())
	return c
}

func (c *component) RightBy(val value.Value) *component {
	c.el.Style(styles.RightBy(val))
	return c
}

func (c *component) RightFraction(fraction float64) *component {
	c.el.Style(styles.RightFraction(fraction))
	return c
}

func (c *component) RightFull() *component {
	c.el.Style(styles.RightFull())
	return c
}

func (c *component) RightPx() *component {
	c.el.Style(styles.RightPx())
	return c
}

func (c *component) Ring(val ...value.Value) *component {
	c.el.Style(styles.Ring(val...))
	return c
}

func (c *component) RingAmber(scale int) *component {
	c.el.Style(styles.RingAmber(scale))
	return c
}

func (c *component) RingAmberAlpha(scale int) *component {
	c.el.Style(styles.RingAmberAlpha(scale))
	return c
}

func (c *component) RingAmberDark(scale int) *component {
	c.el.Style(styles.RingAmberDark(scale))
	return c
}

func (c *component) RingAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.RingAmberDarkAlpha(scale))
	return c
}

func (c *component) RingBlack() *component {
	c.el.Style(styles.RingBlack())
	return c
}

func (c *component) RingBlackAlpha(scale int) *component {
	c.el.Style(styles.RingBlackAlpha(scale))
	return c
}

func (c *component) RingBlue(scale int) *component {
	c.el.Style(styles.RingBlue(scale))
	return c
}

func (c *component) RingBlueAlpha(scale int) *component {
	c.el.Style(styles.RingBlueAlpha(scale))
	return c
}

func (c *component) RingBlueDark(scale int) *component {
	c.el.Style(styles.RingBlueDark(scale))
	return c
}

func (c *component) RingBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.RingBlueDarkAlpha(scale))
	return c
}

func (c *component) RingBronze(scale int) *component {
	c.el.Style(styles.RingBronze(scale))
	return c
}

func (c *component) RingBronzeAlpha(scale int) *component {
	c.el.Style(styles.RingBronzeAlpha(scale))
	return c
}

func (c *component) RingBronzeDark(scale int) *component {
	c.el.Style(styles.RingBronzeDark(scale))
	return c
}

func (c *component) RingBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.RingBronzeDarkAlpha(scale))
	return c
}

func (c *component) RingBrown(scale int) *component {
	c.el.Style(styles.RingBrown(scale))
	return c
}

func (c *component) RingBrownAlpha(scale int) *component {
	c.el.Style(styles.RingBrownAlpha(scale))
	return c
}

func (c *component) RingBrownDark(scale int) *component {
	c.el.Style(styles.RingBrownDark(scale))
	return c
}

func (c *component) RingBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.RingBrownDarkAlpha(scale))
	return c
}

func (c *component) RingColor(val value.Value) *component {
	c.el.Style(styles.RingColor(val))
	return c
}

func (c *component) RingCrimson(scale int) *component {
	c.el.Style(styles.RingCrimson(scale))
	return c
}

func (c *component) RingCrimsonAlpha(scale int) *component {
	c.el.Style(styles.RingCrimsonAlpha(scale))
	return c
}

func (c *component) RingCrimsonDark(scale int) *component {
	c.el.Style(styles.RingCrimsonDark(scale))
	return c
}

func (c *component) RingCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.RingCrimsonDarkAlpha(scale))
	return c
}

func (c *component) RingCurrent() *component {
	c.el.Style(styles.RingCurrent())
	return c
}

func (c *component) RingCyan(scale int) *component {
	c.el.Style(styles.RingCyan(scale))
	return c
}

func (c *component) RingCyanAlpha(scale int) *component {
	c.el.Style(styles.RingCyanAlpha(scale))
	return c
}

func (c *component) RingCyanDark(scale int) *component {
	c.el.Style(styles.RingCyanDark(scale))
	return c
}

func (c *component) RingCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.RingCyanDarkAlpha(scale))
	return c
}

func (c *component) RingGold(scale int) *component {
	c.el.Style(styles.RingGold(scale))
	return c
}

func (c *component) RingGoldAlpha(scale int) *component {
	c.el.Style(styles.RingGoldAlpha(scale))
	return c
}

func (c *component) RingGoldDark(scale int) *component {
	c.el.Style(styles.RingGoldDark(scale))
	return c
}

func (c *component) RingGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.RingGoldDarkAlpha(scale))
	return c
}

func (c *component) RingGrass(scale int) *component {
	c.el.Style(styles.RingGrass(scale))
	return c
}

func (c *component) RingGrassAlpha(scale int) *component {
	c.el.Style(styles.RingGrassAlpha(scale))
	return c
}

func (c *component) RingGrassDark(scale int) *component {
	c.el.Style(styles.RingGrassDark(scale))
	return c
}

func (c *component) RingGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.RingGrassDarkAlpha(scale))
	return c
}

func (c *component) RingGray(scale int) *component {
	c.el.Style(styles.RingGray(scale))
	return c
}

func (c *component) RingGrayAlpha(scale int) *component {
	c.el.Style(styles.RingGrayAlpha(scale))
	return c
}

func (c *component) RingGrayDark(scale int) *component {
	c.el.Style(styles.RingGrayDark(scale))
	return c
}

func (c *component) RingGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.RingGrayDarkAlpha(scale))
	return c
}

func (c *component) RingGreen(scale int) *component {
	c.el.Style(styles.RingGreen(scale))
	return c
}

func (c *component) RingGreenAlpha(scale int) *component {
	c.el.Style(styles.RingGreenAlpha(scale))
	return c
}

func (c *component) RingGreenDark(scale int) *component {
	c.el.Style(styles.RingGreenDark(scale))
	return c
}

func (c *component) RingGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.RingGreenDarkAlpha(scale))
	return c
}

func (c *component) RingIndigo(scale int) *component {
	c.el.Style(styles.RingIndigo(scale))
	return c
}

func (c *component) RingIndigoAlpha(scale int) *component {
	c.el.Style(styles.RingIndigoAlpha(scale))
	return c
}

func (c *component) RingIndigoDark(scale int) *component {
	c.el.Style(styles.RingIndigoDark(scale))
	return c
}

func (c *component) RingIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.RingIndigoDarkAlpha(scale))
	return c
}

func (c *component) RingInherit() *component {
	c.el.Style(styles.RingInherit())
	return c
}

func (c *component) RingIris(scale int) *component {
	c.el.Style(styles.RingIris(scale))
	return c
}

func (c *component) RingIrisAlpha(scale int) *component {
	c.el.Style(styles.RingIrisAlpha(scale))
	return c
}

func (c *component) RingIrisDark(scale int) *component {
	c.el.Style(styles.RingIrisDark(scale))
	return c
}

func (c *component) RingIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.RingIrisDarkAlpha(scale))
	return c
}

func (c *component) RingJade(scale int) *component {
	c.el.Style(styles.RingJade(scale))
	return c
}

func (c *component) RingJadeAlpha(scale int) *component {
	c.el.Style(styles.RingJadeAlpha(scale))
	return c
}

func (c *component) RingJadeDark(scale int) *component {
	c.el.Style(styles.RingJadeDark(scale))
	return c
}

func (c *component) RingJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.RingJadeDarkAlpha(scale))
	return c
}

func (c *component) RingLime(scale int) *component {
	c.el.Style(styles.RingLime(scale))
	return c
}

func (c *component) RingLimeAlpha(scale int) *component {
	c.el.Style(styles.RingLimeAlpha(scale))
	return c
}

func (c *component) RingLimeDark(scale int) *component {
	c.el.Style(styles.RingLimeDark(scale))
	return c
}

func (c *component) RingLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.RingLimeDarkAlpha(scale))
	return c
}

func (c *component) RingMauve(scale int) *component {
	c.el.Style(styles.RingMauve(scale))
	return c
}

func (c *component) RingMauveAlpha(scale int) *component {
	c.el.Style(styles.RingMauveAlpha(scale))
	return c
}

func (c *component) RingMauveDark(scale int) *component {
	c.el.Style(styles.RingMauveDark(scale))
	return c
}

func (c *component) RingMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.RingMauveDarkAlpha(scale))
	return c
}

func (c *component) RingMint(scale int) *component {
	c.el.Style(styles.RingMint(scale))
	return c
}

func (c *component) RingMintAlpha(scale int) *component {
	c.el.Style(styles.RingMintAlpha(scale))
	return c
}

func (c *component) RingMintDark(scale int) *component {
	c.el.Style(styles.RingMintDark(scale))
	return c
}

func (c *component) RingMintDarkAlpha(scale int) *component {
	c.el.Style(styles.RingMintDarkAlpha(scale))
	return c
}

func (c *component) RingOlive(scale int) *component {
	c.el.Style(styles.RingOlive(scale))
	return c
}

func (c *component) RingOliveAlpha(scale int) *component {
	c.el.Style(styles.RingOliveAlpha(scale))
	return c
}

func (c *component) RingOliveDark(scale int) *component {
	c.el.Style(styles.RingOliveDark(scale))
	return c
}

func (c *component) RingOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.RingOliveDarkAlpha(scale))
	return c
}

func (c *component) RingOrange(scale int) *component {
	c.el.Style(styles.RingOrange(scale))
	return c
}

func (c *component) RingOrangeAlpha(scale int) *component {
	c.el.Style(styles.RingOrangeAlpha(scale))
	return c
}

func (c *component) RingOrangeDark(scale int) *component {
	c.el.Style(styles.RingOrangeDark(scale))
	return c
}

func (c *component) RingOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.RingOrangeDarkAlpha(scale))
	return c
}

func (c *component) RingPink(scale int) *component {
	c.el.Style(styles.RingPink(scale))
	return c
}

func (c *component) RingPinkAlpha(scale int) *component {
	c.el.Style(styles.RingPinkAlpha(scale))
	return c
}

func (c *component) RingPinkDark(scale int) *component {
	c.el.Style(styles.RingPinkDark(scale))
	return c
}

func (c *component) RingPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.RingPinkDarkAlpha(scale))
	return c
}

func (c *component) RingPlum(scale int) *component {
	c.el.Style(styles.RingPlum(scale))
	return c
}

func (c *component) RingPlumAlpha(scale int) *component {
	c.el.Style(styles.RingPlumAlpha(scale))
	return c
}

func (c *component) RingPlumDark(scale int) *component {
	c.el.Style(styles.RingPlumDark(scale))
	return c
}

func (c *component) RingPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.RingPlumDarkAlpha(scale))
	return c
}

func (c *component) RingPurple(scale int) *component {
	c.el.Style(styles.RingPurple(scale))
	return c
}

func (c *component) RingPurpleAlpha(scale int) *component {
	c.el.Style(styles.RingPurpleAlpha(scale))
	return c
}

func (c *component) RingPurpleDark(scale int) *component {
	c.el.Style(styles.RingPurpleDark(scale))
	return c
}

func (c *component) RingPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.RingPurpleDarkAlpha(scale))
	return c
}

func (c *component) RingRed(scale int) *component {
	c.el.Style(styles.RingRed(scale))
	return c
}

func (c *component) RingRedAlpha(scale int) *component {
	c.el.Style(styles.RingRedAlpha(scale))
	return c
}

func (c *component) RingRedDark(scale int) *component {
	c.el.Style(styles.RingRedDark(scale))
	return c
}

func (c *component) RingRedDarkAlpha(scale int) *component {
	c.el.Style(styles.RingRedDarkAlpha(scale))
	return c
}

func (c *component) RingRuby(scale int) *component {
	c.el.Style(styles.RingRuby(scale))
	return c
}

func (c *component) RingRubyAlpha(scale int) *component {
	c.el.Style(styles.RingRubyAlpha(scale))
	return c
}

func (c *component) RingRubyDark(scale int) *component {
	c.el.Style(styles.RingRubyDark(scale))
	return c
}

func (c *component) RingRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.RingRubyDarkAlpha(scale))
	return c
}

func (c *component) RingSage(scale int) *component {
	c.el.Style(styles.RingSage(scale))
	return c
}

func (c *component) RingSageAlpha(scale int) *component {
	c.el.Style(styles.RingSageAlpha(scale))
	return c
}

func (c *component) RingSageDark(scale int) *component {
	c.el.Style(styles.RingSageDark(scale))
	return c
}

func (c *component) RingSageDarkAlpha(scale int) *component {
	c.el.Style(styles.RingSageDarkAlpha(scale))
	return c
}

func (c *component) RingSand(scale int) *component {
	c.el.Style(styles.RingSand(scale))
	return c
}

func (c *component) RingSandAlpha(scale int) *component {
	c.el.Style(styles.RingSandAlpha(scale))
	return c
}

func (c *component) RingSandDark(scale int) *component {
	c.el.Style(styles.RingSandDark(scale))
	return c
}

func (c *component) RingSandDarkAlpha(scale int) *component {
	c.el.Style(styles.RingSandDarkAlpha(scale))
	return c
}

func (c *component) RingSky(scale int) *component {
	c.el.Style(styles.RingSky(scale))
	return c
}

func (c *component) RingSkyAlpha(scale int) *component {
	c.el.Style(styles.RingSkyAlpha(scale))
	return c
}

func (c *component) RingSkyDark(scale int) *component {
	c.el.Style(styles.RingSkyDark(scale))
	return c
}

func (c *component) RingSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.RingSkyDarkAlpha(scale))
	return c
}

func (c *component) RingSlate(scale int) *component {
	c.el.Style(styles.RingSlate(scale))
	return c
}

func (c *component) RingSlateAlpha(scale int) *component {
	c.el.Style(styles.RingSlateAlpha(scale))
	return c
}

func (c *component) RingSlateDark(scale int) *component {
	c.el.Style(styles.RingSlateDark(scale))
	return c
}

func (c *component) RingSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.RingSlateDarkAlpha(scale))
	return c
}

func (c *component) RingTeal(scale int) *component {
	c.el.Style(styles.RingTeal(scale))
	return c
}

func (c *component) RingTealAlpha(scale int) *component {
	c.el.Style(styles.RingTealAlpha(scale))
	return c
}

func (c *component) RingTealDark(scale int) *component {
	c.el.Style(styles.RingTealDark(scale))
	return c
}

func (c *component) RingTealDarkAlpha(scale int) *component {
	c.el.Style(styles.RingTealDarkAlpha(scale))
	return c
}

func (c *component) RingTomato(scale int) *component {
	c.el.Style(styles.RingTomato(scale))
	return c
}

func (c *component) RingTomatoAlpha(scale int) *component {
	c.el.Style(styles.RingTomatoAlpha(scale))
	return c
}

func (c *component) RingTomatoDark(scale int) *component {
	c.el.Style(styles.RingTomatoDark(scale))
	return c
}

func (c *component) RingTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.RingTomatoDarkAlpha(scale))
	return c
}

func (c *component) RingTransparent() *component {
	c.el.Style(styles.RingTransparent())
	return c
}

func (c *component) RingViolet(scale int) *component {
	c.el.Style(styles.RingViolet(scale))
	return c
}

func (c *component) RingVioletAlpha(scale int) *component {
	c.el.Style(styles.RingVioletAlpha(scale))
	return c
}

func (c *component) RingVioletDark(scale int) *component {
	c.el.Style(styles.RingVioletDark(scale))
	return c
}

func (c *component) RingVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.RingVioletDarkAlpha(scale))
	return c
}

func (c *component) RingWhite() *component {
	c.el.Style(styles.RingWhite())
	return c
}

func (c *component) RingWhiteAlpha(scale int) *component {
	c.el.Style(styles.RingWhiteAlpha(scale))
	return c
}

func (c *component) RingYellow(scale int) *component {
	c.el.Style(styles.RingYellow(scale))
	return c
}

func (c *component) RingYellowAlpha(scale int) *component {
	c.el.Style(styles.RingYellowAlpha(scale))
	return c
}

func (c *component) RingYellowDark(scale int) *component {
	c.el.Style(styles.RingYellowDark(scale))
	return c
}

func (c *component) RingYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.RingYellowDarkAlpha(scale))
	return c
}

func (c *component) Rotate(val any) *component {
	c.el.Style(styles.Rotate(val))
	return c
}

func (c *component) RotateNone() *component {
	c.el.Style(styles.RotateNone())
	return c
}

func (c *component) RotateX(val any) *component {
	c.el.Style(styles.RotateX(val))
	return c
}

func (c *component) RotateY(val any) *component {
	c.el.Style(styles.RotateY(val))
	return c
}

func (c *component) RotateZ(val any) *component {
	c.el.Style(styles.RotateZ(val))
	return c
}

func (c *component) Rounded(val value.Value) *component {
	c.el.Style(styles.Rounded(val))
	return c
}

func (c *component) Rounded2xl() *component {
	c.el.Style(styles.Rounded2xl())
	return c
}

func (c *component) Rounded3xl() *component {
	c.el.Style(styles.Rounded3xl())
	return c
}

func (c *component) Rounded4Xl() *component {
	c.el.Style(styles.Rounded4Xl())
	return c
}

func (c *component) RoundedB(val value.Value) *component {
	c.el.Style(styles.RoundedB(val))
	return c
}

func (c *component) RoundedB2xl() *component {
	c.el.Style(styles.RoundedB2xl())
	return c
}

func (c *component) RoundedB3xl() *component {
	c.el.Style(styles.RoundedB3xl())
	return c
}

func (c *component) RoundedB4Xl() *component {
	c.el.Style(styles.RoundedB4Xl())
	return c
}

func (c *component) RoundedBFull() *component {
	c.el.Style(styles.RoundedBFull())
	return c
}

func (c *component) RoundedBLg() *component {
	c.el.Style(styles.RoundedBLg())
	return c
}

func (c *component) RoundedBMd() *component {
	c.el.Style(styles.RoundedBMd())
	return c
}

func (c *component) RoundedBNone() *component {
	c.el.Style(styles.RoundedBNone())
	return c
}

func (c *component) RoundedBSm() *component {
	c.el.Style(styles.RoundedBSm())
	return c
}

func (c *component) RoundedBXl() *component {
	c.el.Style(styles.RoundedBXl())
	return c
}

func (c *component) RoundedBXs() *component {
	c.el.Style(styles.RoundedBXs())
	return c
}

func (c *component) RoundedBl(val value.Value) *component {
	c.el.Style(styles.RoundedBl(val))
	return c
}

func (c *component) RoundedBl2xl() *component {
	c.el.Style(styles.RoundedBl2xl())
	return c
}

func (c *component) RoundedBl3xl() *component {
	c.el.Style(styles.RoundedBl3xl())
	return c
}

func (c *component) RoundedBl4Xl() *component {
	c.el.Style(styles.RoundedBl4Xl())
	return c
}

func (c *component) RoundedBlFull() *component {
	c.el.Style(styles.RoundedBlFull())
	return c
}

func (c *component) RoundedBlLg() *component {
	c.el.Style(styles.RoundedBlLg())
	return c
}

func (c *component) RoundedBlMd() *component {
	c.el.Style(styles.RoundedBlMd())
	return c
}

func (c *component) RoundedBlNone() *component {
	c.el.Style(styles.RoundedBlNone())
	return c
}

func (c *component) RoundedBlSm() *component {
	c.el.Style(styles.RoundedBlSm())
	return c
}

func (c *component) RoundedBlXl() *component {
	c.el.Style(styles.RoundedBlXl())
	return c
}

func (c *component) RoundedBlXs() *component {
	c.el.Style(styles.RoundedBlXs())
	return c
}

func (c *component) RoundedBr(val value.Value) *component {
	c.el.Style(styles.RoundedBr(val))
	return c
}

func (c *component) RoundedBr2xl() *component {
	c.el.Style(styles.RoundedBr2xl())
	return c
}

func (c *component) RoundedBr3xl() *component {
	c.el.Style(styles.RoundedBr3xl())
	return c
}

func (c *component) RoundedBr4Xl() *component {
	c.el.Style(styles.RoundedBr4Xl())
	return c
}

func (c *component) RoundedBrFull() *component {
	c.el.Style(styles.RoundedBrFull())
	return c
}

func (c *component) RoundedBrLg() *component {
	c.el.Style(styles.RoundedBrLg())
	return c
}

func (c *component) RoundedBrMd() *component {
	c.el.Style(styles.RoundedBrMd())
	return c
}

func (c *component) RoundedBrNone() *component {
	c.el.Style(styles.RoundedBrNone())
	return c
}

func (c *component) RoundedBrSm() *component {
	c.el.Style(styles.RoundedBrSm())
	return c
}

func (c *component) RoundedBrXl() *component {
	c.el.Style(styles.RoundedBrXl())
	return c
}

func (c *component) RoundedBrXs() *component {
	c.el.Style(styles.RoundedBrXs())
	return c
}

func (c *component) RoundedE(val value.Value) *component {
	c.el.Style(styles.RoundedE(val))
	return c
}

func (c *component) RoundedE2xl() *component {
	c.el.Style(styles.RoundedE2xl())
	return c
}

func (c *component) RoundedE3xl() *component {
	c.el.Style(styles.RoundedE3xl())
	return c
}

func (c *component) RoundedE4Xl() *component {
	c.el.Style(styles.RoundedE4Xl())
	return c
}

func (c *component) RoundedEFull() *component {
	c.el.Style(styles.RoundedEFull())
	return c
}

func (c *component) RoundedELg() *component {
	c.el.Style(styles.RoundedELg())
	return c
}

func (c *component) RoundedEMd() *component {
	c.el.Style(styles.RoundedEMd())
	return c
}

func (c *component) RoundedENone() *component {
	c.el.Style(styles.RoundedENone())
	return c
}

func (c *component) RoundedESm() *component {
	c.el.Style(styles.RoundedESm())
	return c
}

func (c *component) RoundedEXl() *component {
	c.el.Style(styles.RoundedEXl())
	return c
}

func (c *component) RoundedEXs() *component {
	c.el.Style(styles.RoundedEXs())
	return c
}

func (c *component) RoundedEe(val value.Value) *component {
	c.el.Style(styles.RoundedEe(val))
	return c
}

func (c *component) RoundedEe2xl() *component {
	c.el.Style(styles.RoundedEe2xl())
	return c
}

func (c *component) RoundedEe3xl() *component {
	c.el.Style(styles.RoundedEe3xl())
	return c
}

func (c *component) RoundedEe4Xl() *component {
	c.el.Style(styles.RoundedEe4Xl())
	return c
}

func (c *component) RoundedEeFull() *component {
	c.el.Style(styles.RoundedEeFull())
	return c
}

func (c *component) RoundedEeLg() *component {
	c.el.Style(styles.RoundedEeLg())
	return c
}

func (c *component) RoundedEeMd() *component {
	c.el.Style(styles.RoundedEeMd())
	return c
}

func (c *component) RoundedEeNone() *component {
	c.el.Style(styles.RoundedEeNone())
	return c
}

func (c *component) RoundedEeSm() *component {
	c.el.Style(styles.RoundedEeSm())
	return c
}

func (c *component) RoundedEeXl() *component {
	c.el.Style(styles.RoundedEeXl())
	return c
}

func (c *component) RoundedEeXs() *component {
	c.el.Style(styles.RoundedEeXs())
	return c
}

func (c *component) RoundedEs(val value.Value) *component {
	c.el.Style(styles.RoundedEs(val))
	return c
}

func (c *component) RoundedEs2xl() *component {
	c.el.Style(styles.RoundedEs2xl())
	return c
}

func (c *component) RoundedEs3xl() *component {
	c.el.Style(styles.RoundedEs3xl())
	return c
}

func (c *component) RoundedEs4Xl() *component {
	c.el.Style(styles.RoundedEs4Xl())
	return c
}

func (c *component) RoundedEsFull() *component {
	c.el.Style(styles.RoundedEsFull())
	return c
}

func (c *component) RoundedEsLg() *component {
	c.el.Style(styles.RoundedEsLg())
	return c
}

func (c *component) RoundedEsMd() *component {
	c.el.Style(styles.RoundedEsMd())
	return c
}

func (c *component) RoundedEsNone() *component {
	c.el.Style(styles.RoundedEsNone())
	return c
}

func (c *component) RoundedEsSm() *component {
	c.el.Style(styles.RoundedEsSm())
	return c
}

func (c *component) RoundedEsXl() *component {
	c.el.Style(styles.RoundedEsXl())
	return c
}

func (c *component) RoundedEsXs() *component {
	c.el.Style(styles.RoundedEsXs())
	return c
}

func (c *component) RoundedFull() *component {
	c.el.Style(styles.RoundedFull())
	return c
}

func (c *component) RoundedL(val value.Value) *component {
	c.el.Style(styles.RoundedL(val))
	return c
}

func (c *component) RoundedL2xl() *component {
	c.el.Style(styles.RoundedL2xl())
	return c
}

func (c *component) RoundedL3xl() *component {
	c.el.Style(styles.RoundedL3xl())
	return c
}

func (c *component) RoundedL4Xl() *component {
	c.el.Style(styles.RoundedL4Xl())
	return c
}

func (c *component) RoundedLFull() *component {
	c.el.Style(styles.RoundedLFull())
	return c
}

func (c *component) RoundedLLg() *component {
	c.el.Style(styles.RoundedLLg())
	return c
}

func (c *component) RoundedLMd() *component {
	c.el.Style(styles.RoundedLMd())
	return c
}

func (c *component) RoundedLNone() *component {
	c.el.Style(styles.RoundedLNone())
	return c
}

func (c *component) RoundedLSm() *component {
	c.el.Style(styles.RoundedLSm())
	return c
}

func (c *component) RoundedLXl() *component {
	c.el.Style(styles.RoundedLXl())
	return c
}

func (c *component) RoundedLXs() *component {
	c.el.Style(styles.RoundedLXs())
	return c
}

func (c *component) RoundedLg() *component {
	c.el.Style(styles.RoundedLg())
	return c
}

func (c *component) RoundedMd() *component {
	c.el.Style(styles.RoundedMd())
	return c
}

func (c *component) RoundedNone() *component {
	c.el.Style(styles.RoundedNone())
	return c
}

func (c *component) RoundedR(val value.Value) *component {
	c.el.Style(styles.RoundedR(val))
	return c
}

func (c *component) RoundedR2xl() *component {
	c.el.Style(styles.RoundedR2xl())
	return c
}

func (c *component) RoundedR3xl() *component {
	c.el.Style(styles.RoundedR3xl())
	return c
}

func (c *component) RoundedR4Xl() *component {
	c.el.Style(styles.RoundedR4Xl())
	return c
}

func (c *component) RoundedRFull() *component {
	c.el.Style(styles.RoundedRFull())
	return c
}

func (c *component) RoundedRLg() *component {
	c.el.Style(styles.RoundedRLg())
	return c
}

func (c *component) RoundedRMd() *component {
	c.el.Style(styles.RoundedRMd())
	return c
}

func (c *component) RoundedRNone() *component {
	c.el.Style(styles.RoundedRNone())
	return c
}

func (c *component) RoundedRSm() *component {
	c.el.Style(styles.RoundedRSm())
	return c
}

func (c *component) RoundedRXl() *component {
	c.el.Style(styles.RoundedRXl())
	return c
}

func (c *component) RoundedRXs() *component {
	c.el.Style(styles.RoundedRXs())
	return c
}

func (c *component) RoundedS(val value.Value) *component {
	c.el.Style(styles.RoundedS(val))
	return c
}

func (c *component) RoundedS2xl() *component {
	c.el.Style(styles.RoundedS2xl())
	return c
}

func (c *component) RoundedS3xl() *component {
	c.el.Style(styles.RoundedS3xl())
	return c
}

func (c *component) RoundedS4Xl() *component {
	c.el.Style(styles.RoundedS4Xl())
	return c
}

func (c *component) RoundedSFull() *component {
	c.el.Style(styles.RoundedSFull())
	return c
}

func (c *component) RoundedSLg() *component {
	c.el.Style(styles.RoundedSLg())
	return c
}

func (c *component) RoundedSMd() *component {
	c.el.Style(styles.RoundedSMd())
	return c
}

func (c *component) RoundedSNone() *component {
	c.el.Style(styles.RoundedSNone())
	return c
}

func (c *component) RoundedSSm() *component {
	c.el.Style(styles.RoundedSSm())
	return c
}

func (c *component) RoundedSXl() *component {
	c.el.Style(styles.RoundedSXl())
	return c
}

func (c *component) RoundedSXs() *component {
	c.el.Style(styles.RoundedSXs())
	return c
}

func (c *component) RoundedSe(val value.Value) *component {
	c.el.Style(styles.RoundedSe(val))
	return c
}

func (c *component) RoundedSe2xl() *component {
	c.el.Style(styles.RoundedSe2xl())
	return c
}

func (c *component) RoundedSe3xl() *component {
	c.el.Style(styles.RoundedSe3xl())
	return c
}

func (c *component) RoundedSe4Xl() *component {
	c.el.Style(styles.RoundedSe4Xl())
	return c
}

func (c *component) RoundedSeFull() *component {
	c.el.Style(styles.RoundedSeFull())
	return c
}

func (c *component) RoundedSeLg() *component {
	c.el.Style(styles.RoundedSeLg())
	return c
}

func (c *component) RoundedSeMd() *component {
	c.el.Style(styles.RoundedSeMd())
	return c
}

func (c *component) RoundedSeNone() *component {
	c.el.Style(styles.RoundedSeNone())
	return c
}

func (c *component) RoundedSeSm() *component {
	c.el.Style(styles.RoundedSeSm())
	return c
}

func (c *component) RoundedSeXl() *component {
	c.el.Style(styles.RoundedSeXl())
	return c
}

func (c *component) RoundedSeXs() *component {
	c.el.Style(styles.RoundedSeXs())
	return c
}

func (c *component) RoundedSm() *component {
	c.el.Style(styles.RoundedSm())
	return c
}

func (c *component) RoundedSs(val value.Value) *component {
	c.el.Style(styles.RoundedSs(val))
	return c
}

func (c *component) RoundedSs2xl() *component {
	c.el.Style(styles.RoundedSs2xl())
	return c
}

func (c *component) RoundedSs3xl() *component {
	c.el.Style(styles.RoundedSs3xl())
	return c
}

func (c *component) RoundedSs4Xl() *component {
	c.el.Style(styles.RoundedSs4Xl())
	return c
}

func (c *component) RoundedSsFull() *component {
	c.el.Style(styles.RoundedSsFull())
	return c
}

func (c *component) RoundedSsLg() *component {
	c.el.Style(styles.RoundedSsLg())
	return c
}

func (c *component) RoundedSsMd() *component {
	c.el.Style(styles.RoundedSsMd())
	return c
}

func (c *component) RoundedSsNone() *component {
	c.el.Style(styles.RoundedSsNone())
	return c
}

func (c *component) RoundedSsSm() *component {
	c.el.Style(styles.RoundedSsSm())
	return c
}

func (c *component) RoundedSsXl() *component {
	c.el.Style(styles.RoundedSsXl())
	return c
}

func (c *component) RoundedSsXs() *component {
	c.el.Style(styles.RoundedSsXs())
	return c
}

func (c *component) RoundedT(val value.Value) *component {
	c.el.Style(styles.RoundedT(val))
	return c
}

func (c *component) RoundedT2xl() *component {
	c.el.Style(styles.RoundedT2xl())
	return c
}

func (c *component) RoundedT3xl() *component {
	c.el.Style(styles.RoundedT3xl())
	return c
}

func (c *component) RoundedT4Xl() *component {
	c.el.Style(styles.RoundedT4Xl())
	return c
}

func (c *component) RoundedTFull() *component {
	c.el.Style(styles.RoundedTFull())
	return c
}

func (c *component) RoundedTLg() *component {
	c.el.Style(styles.RoundedTLg())
	return c
}

func (c *component) RoundedTMd() *component {
	c.el.Style(styles.RoundedTMd())
	return c
}

func (c *component) RoundedTNone() *component {
	c.el.Style(styles.RoundedTNone())
	return c
}

func (c *component) RoundedTSm() *component {
	c.el.Style(styles.RoundedTSm())
	return c
}

func (c *component) RoundedTXl() *component {
	c.el.Style(styles.RoundedTXl())
	return c
}

func (c *component) RoundedTXs() *component {
	c.el.Style(styles.RoundedTXs())
	return c
}

func (c *component) RoundedTl(val value.Value) *component {
	c.el.Style(styles.RoundedTl(val))
	return c
}

func (c *component) RoundedTl2xl() *component {
	c.el.Style(styles.RoundedTl2xl())
	return c
}

func (c *component) RoundedTl3xl() *component {
	c.el.Style(styles.RoundedTl3xl())
	return c
}

func (c *component) RoundedTl4Xl() *component {
	c.el.Style(styles.RoundedTl4Xl())
	return c
}

func (c *component) RoundedTlFull() *component {
	c.el.Style(styles.RoundedTlFull())
	return c
}

func (c *component) RoundedTlLg() *component {
	c.el.Style(styles.RoundedTlLg())
	return c
}

func (c *component) RoundedTlMd() *component {
	c.el.Style(styles.RoundedTlMd())
	return c
}

func (c *component) RoundedTlNone() *component {
	c.el.Style(styles.RoundedTlNone())
	return c
}

func (c *component) RoundedTlSm() *component {
	c.el.Style(styles.RoundedTlSm())
	return c
}

func (c *component) RoundedTlXl() *component {
	c.el.Style(styles.RoundedTlXl())
	return c
}

func (c *component) RoundedTlXs() *component {
	c.el.Style(styles.RoundedTlXs())
	return c
}

func (c *component) RoundedTr(val value.Value) *component {
	c.el.Style(styles.RoundedTr(val))
	return c
}

func (c *component) RoundedTr2xl() *component {
	c.el.Style(styles.RoundedTr2xl())
	return c
}

func (c *component) RoundedTr3xl() *component {
	c.el.Style(styles.RoundedTr3xl())
	return c
}

func (c *component) RoundedTr4Xl() *component {
	c.el.Style(styles.RoundedTr4Xl())
	return c
}

func (c *component) RoundedTrFull() *component {
	c.el.Style(styles.RoundedTrFull())
	return c
}

func (c *component) RoundedTrLg() *component {
	c.el.Style(styles.RoundedTrLg())
	return c
}

func (c *component) RoundedTrMd() *component {
	c.el.Style(styles.RoundedTrMd())
	return c
}

func (c *component) RoundedTrNone() *component {
	c.el.Style(styles.RoundedTrNone())
	return c
}

func (c *component) RoundedTrSm() *component {
	c.el.Style(styles.RoundedTrSm())
	return c
}

func (c *component) RoundedTrXl() *component {
	c.el.Style(styles.RoundedTrXl())
	return c
}

func (c *component) RoundedTrXs() *component {
	c.el.Style(styles.RoundedTrXs())
	return c
}

func (c *component) RoundedXl() *component {
	c.el.Style(styles.RoundedXl())
	return c
}

func (c *component) RoundedXs() *component {
	c.el.Style(styles.RoundedXs())
	return c
}

func (c *component) Row(number int) *component {
	c.el.Style(styles.Row(number))
	return c
}

func (c *component) RowAuto() *component {
	c.el.Style(styles.RowAuto())
	return c
}

func (c *component) RowBy(val value.Value) *component {
	c.el.Style(styles.RowBy(val))
	return c
}

func (c *component) RowEnd(number int) *component {
	c.el.Style(styles.RowEnd(number))
	return c
}

func (c *component) RowEndAuto() *component {
	c.el.Style(styles.RowEndAuto())
	return c
}

func (c *component) RowEndBy(val value.Value) *component {
	c.el.Style(styles.RowEndBy(val))
	return c
}

func (c *component) RowSpan(number int) *component {
	c.el.Style(styles.RowSpan(number))
	return c
}

func (c *component) RowSpanBy(val value.Value) *component {
	c.el.Style(styles.RowSpanBy(val))
	return c
}

func (c *component) RowSpanFull() *component {
	c.el.Style(styles.RowSpanFull())
	return c
}

func (c *component) RowStart(number int) *component {
	c.el.Style(styles.RowStart(number))
	return c
}

func (c *component) RowStartAuto() *component {
	c.el.Style(styles.RowStartAuto())
	return c
}

func (c *component) RowStartBy(val value.Value) *component {
	c.el.Style(styles.RowStartBy(val))
	return c
}

func (c *component) Saturate(val any) *component {
	c.el.Style(styles.Saturate(val))
	return c
}

func (c *component) Scale(val any) *component {
	c.el.Style(styles.Scale(val))
	return c
}

func (c *component) Scale3d() *component {
	c.el.Style(styles.Scale3d())
	return c
}

func (c *component) ScaleNone() *component {
	c.el.Style(styles.ScaleNone())
	return c
}

func (c *component) ScaleX(val any) *component {
	c.el.Style(styles.ScaleX(val))
	return c
}

func (c *component) ScaleY(val any) *component {
	c.el.Style(styles.ScaleY(val))
	return c
}

func (c *component) ScaleZ(val any) *component {
	c.el.Style(styles.ScaleZ(val))
	return c
}

func (c *component) SchemeDark() *component {
	c.el.Style(styles.SchemeDark())
	return c
}

func (c *component) SchemeLight() *component {
	c.el.Style(styles.SchemeLight())
	return c
}

func (c *component) SchemeLightDark() *component {
	c.el.Style(styles.SchemeLightDark())
	return c
}

func (c *component) SchemeNormal() *component {
	c.el.Style(styles.SchemeNormal())
	return c
}

func (c *component) SchemeOnlyDark() *component {
	c.el.Style(styles.SchemeOnlyDark())
	return c
}

func (c *component) SchemeOnlyLight() *component {
	c.el.Style(styles.SchemeOnlyLight())
	return c
}

func (c *component) ScrollAuto() *component {
	c.el.Style(styles.ScrollAuto())
	return c
}

func (c *component) ScrollM(val any) *component {
	c.el.Style(styles.ScrollM(val))
	return c
}

func (c *component) ScrollMb(val any) *component {
	c.el.Style(styles.ScrollMb(val))
	return c
}

func (c *component) ScrollMe(val any) *component {
	c.el.Style(styles.ScrollMe(val))
	return c
}

func (c *component) ScrollMl(val any) *component {
	c.el.Style(styles.ScrollMl(val))
	return c
}

func (c *component) ScrollMr(val any) *component {
	c.el.Style(styles.ScrollMr(val))
	return c
}

func (c *component) ScrollMs(val any) *component {
	c.el.Style(styles.ScrollMs(val))
	return c
}

func (c *component) ScrollMt(val any) *component {
	c.el.Style(styles.ScrollMt(val))
	return c
}

func (c *component) ScrollMx(val any) *component {
	c.el.Style(styles.ScrollMx(val))
	return c
}

func (c *component) ScrollMy(val any) *component {
	c.el.Style(styles.ScrollMy(val))
	return c
}

func (c *component) ScrollP(val any) *component {
	c.el.Style(styles.ScrollP(val))
	return c
}

func (c *component) ScrollPb(val any) *component {
	c.el.Style(styles.ScrollPb(val))
	return c
}

func (c *component) ScrollPe(val any) *component {
	c.el.Style(styles.ScrollPe(val))
	return c
}

func (c *component) ScrollPl(val any) *component {
	c.el.Style(styles.ScrollPl(val))
	return c
}

func (c *component) ScrollPr(val any) *component {
	c.el.Style(styles.ScrollPr(val))
	return c
}

func (c *component) ScrollPs(val any) *component {
	c.el.Style(styles.ScrollPs(val))
	return c
}

func (c *component) ScrollPt(val any) *component {
	c.el.Style(styles.ScrollPt(val))
	return c
}

func (c *component) ScrollPx(val any) *component {
	c.el.Style(styles.ScrollPx(val))
	return c
}

func (c *component) ScrollPy(val any) *component {
	c.el.Style(styles.ScrollPy(val))
	return c
}

func (c *component) ScrollSmooth() *component {
	c.el.Style(styles.ScrollSmooth())
	return c
}

func (c *component) SelectAll() *component {
	c.el.Style(styles.SelectAll())
	return c
}

func (c *component) SelectAuto() *component {
	c.el.Style(styles.SelectAuto())
	return c
}

func (c *component) SelectNone() *component {
	c.el.Style(styles.SelectNone())
	return c
}

func (c *component) SelectText() *component {
	c.el.Style(styles.SelectText())
	return c
}

func (c *component) SelectorParam(param string) *component {
	c.el.Style(styles.SelectorParam(param))
	return c
}

func (c *component) SelfAuto() *component {
	c.el.Style(styles.SelfAuto())
	return c
}

func (c *component) SelfBaseline() *component {
	c.el.Style(styles.SelfBaseline())
	return c
}

func (c *component) SelfBaselineLast() *component {
	c.el.Style(styles.SelfBaselineLast())
	return c
}

func (c *component) SelfCenter() *component {
	c.el.Style(styles.SelfCenter())
	return c
}

func (c *component) SelfCenterSafe() *component {
	c.el.Style(styles.SelfCenterSafe())
	return c
}

func (c *component) SelfEnd() *component {
	c.el.Style(styles.SelfEnd())
	return c
}

func (c *component) SelfEndSafe() *component {
	c.el.Style(styles.SelfEndSafe())
	return c
}

func (c *component) SelfStart() *component {
	c.el.Style(styles.SelfStart())
	return c
}

func (c *component) SelfStretch() *component {
	c.el.Style(styles.SelfStretch())
	return c
}

func (c *component) Sepia(val ...any) *component {
	c.el.Style(styles.Sepia(val...))
	return c
}

func (c *component) Shadow(val value.Value) *component {
	c.el.Style(styles.Shadow(val))
	return c
}

func (c *component) Shadow2xl() *component {
	c.el.Style(styles.Shadow2xl())
	return c
}

func (c *component) Shadow2xs() *component {
	c.el.Style(styles.Shadow2xs())
	return c
}

func (c *component) ShadowAmber(scale int) *component {
	c.el.Style(styles.ShadowAmber(scale))
	return c
}

func (c *component) ShadowAmberAlpha(scale int) *component {
	c.el.Style(styles.ShadowAmberAlpha(scale))
	return c
}

func (c *component) ShadowAmberDark(scale int) *component {
	c.el.Style(styles.ShadowAmberDark(scale))
	return c
}

func (c *component) ShadowAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowAmberDarkAlpha(scale))
	return c
}

func (c *component) ShadowBlack() *component {
	c.el.Style(styles.ShadowBlack())
	return c
}

func (c *component) ShadowBlackAlpha(scale int) *component {
	c.el.Style(styles.ShadowBlackAlpha(scale))
	return c
}

func (c *component) ShadowBlue(scale int) *component {
	c.el.Style(styles.ShadowBlue(scale))
	return c
}

func (c *component) ShadowBlueAlpha(scale int) *component {
	c.el.Style(styles.ShadowBlueAlpha(scale))
	return c
}

func (c *component) ShadowBlueDark(scale int) *component {
	c.el.Style(styles.ShadowBlueDark(scale))
	return c
}

func (c *component) ShadowBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowBlueDarkAlpha(scale))
	return c
}

func (c *component) ShadowBronze(scale int) *component {
	c.el.Style(styles.ShadowBronze(scale))
	return c
}

func (c *component) ShadowBronzeAlpha(scale int) *component {
	c.el.Style(styles.ShadowBronzeAlpha(scale))
	return c
}

func (c *component) ShadowBronzeDark(scale int) *component {
	c.el.Style(styles.ShadowBronzeDark(scale))
	return c
}

func (c *component) ShadowBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowBronzeDarkAlpha(scale))
	return c
}

func (c *component) ShadowBrown(scale int) *component {
	c.el.Style(styles.ShadowBrown(scale))
	return c
}

func (c *component) ShadowBrownAlpha(scale int) *component {
	c.el.Style(styles.ShadowBrownAlpha(scale))
	return c
}

func (c *component) ShadowBrownDark(scale int) *component {
	c.el.Style(styles.ShadowBrownDark(scale))
	return c
}

func (c *component) ShadowBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowBrownDarkAlpha(scale))
	return c
}

func (c *component) ShadowColor(val value.Value) *component {
	c.el.Style(styles.ShadowColor(val))
	return c
}

func (c *component) ShadowCrimson(scale int) *component {
	c.el.Style(styles.ShadowCrimson(scale))
	return c
}

func (c *component) ShadowCrimsonAlpha(scale int) *component {
	c.el.Style(styles.ShadowCrimsonAlpha(scale))
	return c
}

func (c *component) ShadowCrimsonDark(scale int) *component {
	c.el.Style(styles.ShadowCrimsonDark(scale))
	return c
}

func (c *component) ShadowCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowCrimsonDarkAlpha(scale))
	return c
}

func (c *component) ShadowCurrent() *component {
	c.el.Style(styles.ShadowCurrent())
	return c
}

func (c *component) ShadowCyan(scale int) *component {
	c.el.Style(styles.ShadowCyan(scale))
	return c
}

func (c *component) ShadowCyanAlpha(scale int) *component {
	c.el.Style(styles.ShadowCyanAlpha(scale))
	return c
}

func (c *component) ShadowCyanDark(scale int) *component {
	c.el.Style(styles.ShadowCyanDark(scale))
	return c
}

func (c *component) ShadowCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowCyanDarkAlpha(scale))
	return c
}

func (c *component) ShadowGold(scale int) *component {
	c.el.Style(styles.ShadowGold(scale))
	return c
}

func (c *component) ShadowGoldAlpha(scale int) *component {
	c.el.Style(styles.ShadowGoldAlpha(scale))
	return c
}

func (c *component) ShadowGoldDark(scale int) *component {
	c.el.Style(styles.ShadowGoldDark(scale))
	return c
}

func (c *component) ShadowGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowGoldDarkAlpha(scale))
	return c
}

func (c *component) ShadowGrass(scale int) *component {
	c.el.Style(styles.ShadowGrass(scale))
	return c
}

func (c *component) ShadowGrassAlpha(scale int) *component {
	c.el.Style(styles.ShadowGrassAlpha(scale))
	return c
}

func (c *component) ShadowGrassDark(scale int) *component {
	c.el.Style(styles.ShadowGrassDark(scale))
	return c
}

func (c *component) ShadowGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowGrassDarkAlpha(scale))
	return c
}

func (c *component) ShadowGray(scale int) *component {
	c.el.Style(styles.ShadowGray(scale))
	return c
}

func (c *component) ShadowGrayAlpha(scale int) *component {
	c.el.Style(styles.ShadowGrayAlpha(scale))
	return c
}

func (c *component) ShadowGrayDark(scale int) *component {
	c.el.Style(styles.ShadowGrayDark(scale))
	return c
}

func (c *component) ShadowGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowGrayDarkAlpha(scale))
	return c
}

func (c *component) ShadowGreen(scale int) *component {
	c.el.Style(styles.ShadowGreen(scale))
	return c
}

func (c *component) ShadowGreenAlpha(scale int) *component {
	c.el.Style(styles.ShadowGreenAlpha(scale))
	return c
}

func (c *component) ShadowGreenDark(scale int) *component {
	c.el.Style(styles.ShadowGreenDark(scale))
	return c
}

func (c *component) ShadowGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowGreenDarkAlpha(scale))
	return c
}

func (c *component) ShadowIndigo(scale int) *component {
	c.el.Style(styles.ShadowIndigo(scale))
	return c
}

func (c *component) ShadowIndigoAlpha(scale int) *component {
	c.el.Style(styles.ShadowIndigoAlpha(scale))
	return c
}

func (c *component) ShadowIndigoDark(scale int) *component {
	c.el.Style(styles.ShadowIndigoDark(scale))
	return c
}

func (c *component) ShadowIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowIndigoDarkAlpha(scale))
	return c
}

func (c *component) ShadowInherit() *component {
	c.el.Style(styles.ShadowInherit())
	return c
}

func (c *component) ShadowIris(scale int) *component {
	c.el.Style(styles.ShadowIris(scale))
	return c
}

func (c *component) ShadowIrisAlpha(scale int) *component {
	c.el.Style(styles.ShadowIrisAlpha(scale))
	return c
}

func (c *component) ShadowIrisDark(scale int) *component {
	c.el.Style(styles.ShadowIrisDark(scale))
	return c
}

func (c *component) ShadowIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowIrisDarkAlpha(scale))
	return c
}

func (c *component) ShadowJade(scale int) *component {
	c.el.Style(styles.ShadowJade(scale))
	return c
}

func (c *component) ShadowJadeAlpha(scale int) *component {
	c.el.Style(styles.ShadowJadeAlpha(scale))
	return c
}

func (c *component) ShadowJadeDark(scale int) *component {
	c.el.Style(styles.ShadowJadeDark(scale))
	return c
}

func (c *component) ShadowJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowJadeDarkAlpha(scale))
	return c
}

func (c *component) ShadowLg() *component {
	c.el.Style(styles.ShadowLg())
	return c
}

func (c *component) ShadowLime(scale int) *component {
	c.el.Style(styles.ShadowLime(scale))
	return c
}

func (c *component) ShadowLimeAlpha(scale int) *component {
	c.el.Style(styles.ShadowLimeAlpha(scale))
	return c
}

func (c *component) ShadowLimeDark(scale int) *component {
	c.el.Style(styles.ShadowLimeDark(scale))
	return c
}

func (c *component) ShadowLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowLimeDarkAlpha(scale))
	return c
}

func (c *component) ShadowMauve(scale int) *component {
	c.el.Style(styles.ShadowMauve(scale))
	return c
}

func (c *component) ShadowMauveAlpha(scale int) *component {
	c.el.Style(styles.ShadowMauveAlpha(scale))
	return c
}

func (c *component) ShadowMauveDark(scale int) *component {
	c.el.Style(styles.ShadowMauveDark(scale))
	return c
}

func (c *component) ShadowMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowMauveDarkAlpha(scale))
	return c
}

func (c *component) ShadowMd() *component {
	c.el.Style(styles.ShadowMd())
	return c
}

func (c *component) ShadowMint(scale int) *component {
	c.el.Style(styles.ShadowMint(scale))
	return c
}

func (c *component) ShadowMintAlpha(scale int) *component {
	c.el.Style(styles.ShadowMintAlpha(scale))
	return c
}

func (c *component) ShadowMintDark(scale int) *component {
	c.el.Style(styles.ShadowMintDark(scale))
	return c
}

func (c *component) ShadowMintDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowMintDarkAlpha(scale))
	return c
}

func (c *component) ShadowNone() *component {
	c.el.Style(styles.ShadowNone())
	return c
}

func (c *component) ShadowOlive(scale int) *component {
	c.el.Style(styles.ShadowOlive(scale))
	return c
}

func (c *component) ShadowOliveAlpha(scale int) *component {
	c.el.Style(styles.ShadowOliveAlpha(scale))
	return c
}

func (c *component) ShadowOliveDark(scale int) *component {
	c.el.Style(styles.ShadowOliveDark(scale))
	return c
}

func (c *component) ShadowOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowOliveDarkAlpha(scale))
	return c
}

func (c *component) ShadowOrange(scale int) *component {
	c.el.Style(styles.ShadowOrange(scale))
	return c
}

func (c *component) ShadowOrangeAlpha(scale int) *component {
	c.el.Style(styles.ShadowOrangeAlpha(scale))
	return c
}

func (c *component) ShadowOrangeDark(scale int) *component {
	c.el.Style(styles.ShadowOrangeDark(scale))
	return c
}

func (c *component) ShadowOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowOrangeDarkAlpha(scale))
	return c
}

func (c *component) ShadowPink(scale int) *component {
	c.el.Style(styles.ShadowPink(scale))
	return c
}

func (c *component) ShadowPinkAlpha(scale int) *component {
	c.el.Style(styles.ShadowPinkAlpha(scale))
	return c
}

func (c *component) ShadowPinkDark(scale int) *component {
	c.el.Style(styles.ShadowPinkDark(scale))
	return c
}

func (c *component) ShadowPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowPinkDarkAlpha(scale))
	return c
}

func (c *component) ShadowPlum(scale int) *component {
	c.el.Style(styles.ShadowPlum(scale))
	return c
}

func (c *component) ShadowPlumAlpha(scale int) *component {
	c.el.Style(styles.ShadowPlumAlpha(scale))
	return c
}

func (c *component) ShadowPlumDark(scale int) *component {
	c.el.Style(styles.ShadowPlumDark(scale))
	return c
}

func (c *component) ShadowPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowPlumDarkAlpha(scale))
	return c
}

func (c *component) ShadowPurple(scale int) *component {
	c.el.Style(styles.ShadowPurple(scale))
	return c
}

func (c *component) ShadowPurpleAlpha(scale int) *component {
	c.el.Style(styles.ShadowPurpleAlpha(scale))
	return c
}

func (c *component) ShadowPurpleDark(scale int) *component {
	c.el.Style(styles.ShadowPurpleDark(scale))
	return c
}

func (c *component) ShadowPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowPurpleDarkAlpha(scale))
	return c
}

func (c *component) ShadowRed(scale int) *component {
	c.el.Style(styles.ShadowRed(scale))
	return c
}

func (c *component) ShadowRedAlpha(scale int) *component {
	c.el.Style(styles.ShadowRedAlpha(scale))
	return c
}

func (c *component) ShadowRedDark(scale int) *component {
	c.el.Style(styles.ShadowRedDark(scale))
	return c
}

func (c *component) ShadowRedDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowRedDarkAlpha(scale))
	return c
}

func (c *component) ShadowRuby(scale int) *component {
	c.el.Style(styles.ShadowRuby(scale))
	return c
}

func (c *component) ShadowRubyAlpha(scale int) *component {
	c.el.Style(styles.ShadowRubyAlpha(scale))
	return c
}

func (c *component) ShadowRubyDark(scale int) *component {
	c.el.Style(styles.ShadowRubyDark(scale))
	return c
}

func (c *component) ShadowRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowRubyDarkAlpha(scale))
	return c
}

func (c *component) ShadowSage(scale int) *component {
	c.el.Style(styles.ShadowSage(scale))
	return c
}

func (c *component) ShadowSageAlpha(scale int) *component {
	c.el.Style(styles.ShadowSageAlpha(scale))
	return c
}

func (c *component) ShadowSageDark(scale int) *component {
	c.el.Style(styles.ShadowSageDark(scale))
	return c
}

func (c *component) ShadowSageDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowSageDarkAlpha(scale))
	return c
}

func (c *component) ShadowSand(scale int) *component {
	c.el.Style(styles.ShadowSand(scale))
	return c
}

func (c *component) ShadowSandAlpha(scale int) *component {
	c.el.Style(styles.ShadowSandAlpha(scale))
	return c
}

func (c *component) ShadowSandDark(scale int) *component {
	c.el.Style(styles.ShadowSandDark(scale))
	return c
}

func (c *component) ShadowSandDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowSandDarkAlpha(scale))
	return c
}

func (c *component) ShadowSky(scale int) *component {
	c.el.Style(styles.ShadowSky(scale))
	return c
}

func (c *component) ShadowSkyAlpha(scale int) *component {
	c.el.Style(styles.ShadowSkyAlpha(scale))
	return c
}

func (c *component) ShadowSkyDark(scale int) *component {
	c.el.Style(styles.ShadowSkyDark(scale))
	return c
}

func (c *component) ShadowSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowSkyDarkAlpha(scale))
	return c
}

func (c *component) ShadowSlate(scale int) *component {
	c.el.Style(styles.ShadowSlate(scale))
	return c
}

func (c *component) ShadowSlateAlpha(scale int) *component {
	c.el.Style(styles.ShadowSlateAlpha(scale))
	return c
}

func (c *component) ShadowSlateDark(scale int) *component {
	c.el.Style(styles.ShadowSlateDark(scale))
	return c
}

func (c *component) ShadowSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowSlateDarkAlpha(scale))
	return c
}

func (c *component) ShadowSm() *component {
	c.el.Style(styles.ShadowSm())
	return c
}

func (c *component) ShadowTeal(scale int) *component {
	c.el.Style(styles.ShadowTeal(scale))
	return c
}

func (c *component) ShadowTealAlpha(scale int) *component {
	c.el.Style(styles.ShadowTealAlpha(scale))
	return c
}

func (c *component) ShadowTealDark(scale int) *component {
	c.el.Style(styles.ShadowTealDark(scale))
	return c
}

func (c *component) ShadowTealDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowTealDarkAlpha(scale))
	return c
}

func (c *component) ShadowTomato(scale int) *component {
	c.el.Style(styles.ShadowTomato(scale))
	return c
}

func (c *component) ShadowTomatoAlpha(scale int) *component {
	c.el.Style(styles.ShadowTomatoAlpha(scale))
	return c
}

func (c *component) ShadowTomatoDark(scale int) *component {
	c.el.Style(styles.ShadowTomatoDark(scale))
	return c
}

func (c *component) ShadowTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowTomatoDarkAlpha(scale))
	return c
}

func (c *component) ShadowTransparent() *component {
	c.el.Style(styles.ShadowTransparent())
	return c
}

func (c *component) ShadowViolet(scale int) *component {
	c.el.Style(styles.ShadowViolet(scale))
	return c
}

func (c *component) ShadowVioletAlpha(scale int) *component {
	c.el.Style(styles.ShadowVioletAlpha(scale))
	return c
}

func (c *component) ShadowVioletDark(scale int) *component {
	c.el.Style(styles.ShadowVioletDark(scale))
	return c
}

func (c *component) ShadowVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowVioletDarkAlpha(scale))
	return c
}

func (c *component) ShadowWhite() *component {
	c.el.Style(styles.ShadowWhite())
	return c
}

func (c *component) ShadowWhiteAlpha(scale int) *component {
	c.el.Style(styles.ShadowWhiteAlpha(scale))
	return c
}

func (c *component) ShadowXl() *component {
	c.el.Style(styles.ShadowXl())
	return c
}

func (c *component) ShadowXs() *component {
	c.el.Style(styles.ShadowXs())
	return c
}

func (c *component) ShadowYellow(scale int) *component {
	c.el.Style(styles.ShadowYellow(scale))
	return c
}

func (c *component) ShadowYellowAlpha(scale int) *component {
	c.el.Style(styles.ShadowYellowAlpha(scale))
	return c
}

func (c *component) ShadowYellowDark(scale int) *component {
	c.el.Style(styles.ShadowYellowDark(scale))
	return c
}

func (c *component) ShadowYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.ShadowYellowDarkAlpha(scale))
	return c
}

func (c *component) Shrink() *component {
	c.el.Style(styles.Shrink())
	return c
}

func (c *component) ShrinkBy(val value.Value) *component {
	c.el.Style(styles.ShrinkBy(val))
	return c
}

func (c *component) Size(number int) *component {
	c.el.Style(styles.Size(number))
	return c
}

func (c *component) SizeAuto() *component {
	c.el.Style(styles.SizeAuto())
	return c
}

func (c *component) SizeBy(val value.Value) *component {
	c.el.Style(styles.SizeBy(val))
	return c
}

func (c *component) SizeDvh() *component {
	c.el.Style(styles.SizeDvh())
	return c
}

func (c *component) SizeDvw() *component {
	c.el.Style(styles.SizeDvw())
	return c
}

func (c *component) SizeFit() *component {
	c.el.Style(styles.SizeFit())
	return c
}

func (c *component) SizeFraction(fraction float32) *component {
	c.el.Style(styles.SizeFraction(fraction))
	return c
}

func (c *component) SizeLvh() *component {
	c.el.Style(styles.SizeLvh())
	return c
}

func (c *component) SizeLvw() *component {
	c.el.Style(styles.SizeLvw())
	return c
}

func (c *component) SizeMax() *component {
	c.el.Style(styles.SizeMax())
	return c
}

func (c *component) SizeMin() *component {
	c.el.Style(styles.SizeMin())
	return c
}

func (c *component) SizePx() *component {
	c.el.Style(styles.SizePx())
	return c
}

func (c *component) SizeSvh() *component {
	c.el.Style(styles.SizeSvh())
	return c
}

func (c *component) SizeSvw() *component {
	c.el.Style(styles.SizeSvw())
	return c
}

func (c *component) Skew(val any) *component {
	c.el.Style(styles.Skew(val))
	return c
}

func (c *component) SkewX(val any) *component {
	c.el.Style(styles.SkewX(val))
	return c
}

func (c *component) SkewY(val any) *component {
	c.el.Style(styles.SkewY(val))
	return c
}

func (c *component) SlashedZero() *component {
	c.el.Style(styles.SlashedZero())
	return c
}

func (c *component) SnapAlignNone() *component {
	c.el.Style(styles.SnapAlignNone())
	return c
}

func (c *component) SnapAlways() *component {
	c.el.Style(styles.SnapAlways())
	return c
}

func (c *component) SnapBoth() *component {
	c.el.Style(styles.SnapBoth())
	return c
}

func (c *component) SnapCenter() *component {
	c.el.Style(styles.SnapCenter())
	return c
}

func (c *component) SnapEnd() *component {
	c.el.Style(styles.SnapEnd())
	return c
}

func (c *component) SnapMandatory() *component {
	c.el.Style(styles.SnapMandatory())
	return c
}

func (c *component) SnapNone() *component {
	c.el.Style(styles.SnapNone())
	return c
}

func (c *component) SnapNormal() *component {
	c.el.Style(styles.SnapNormal())
	return c
}

func (c *component) SnapProximity() *component {
	c.el.Style(styles.SnapProximity())
	return c
}

func (c *component) SnapStart() *component {
	c.el.Style(styles.SnapStart())
	return c
}

func (c *component) SnapX() *component {
	c.el.Style(styles.SnapX())
	return c
}

func (c *component) SnapY() *component {
	c.el.Style(styles.SnapY())
	return c
}

func (c *component) SrOnly() *component {
	c.el.Style(styles.SrOnly())
	return c
}

func (c *component) StackedFractions() *component {
	c.el.Style(styles.StackedFractions())
	return c
}

func (c *component) Start(number int) *component {
	c.el.Style(styles.Start(number))
	return c
}

func (c *component) StartAuto() *component {
	c.el.Style(styles.StartAuto())
	return c
}

func (c *component) StartBy(val value.Value) *component {
	c.el.Style(styles.StartBy(val))
	return c
}

func (c *component) StartFraction(fraction float64) *component {
	c.el.Style(styles.StartFraction(fraction))
	return c
}

func (c *component) StartFull() *component {
	c.el.Style(styles.StartFull())
	return c
}

func (c *component) StartPx() *component {
	c.el.Style(styles.StartPx())
	return c
}

func (c *component) Static() *component {
	c.el.Style(styles.Static())
	return c
}

func (c *component) Sticky() *component {
	c.el.Style(styles.Sticky())
	return c
}

func (c *component) Stroke(val any) *component {
	c.el.Style(styles.Stroke(val))
	return c
}

func (c *component) StrokeAmber(scale int) *component {
	c.el.Style(styles.StrokeAmber(scale))
	return c
}

func (c *component) StrokeAmberAlpha(scale int) *component {
	c.el.Style(styles.StrokeAmberAlpha(scale))
	return c
}

func (c *component) StrokeAmberDark(scale int) *component {
	c.el.Style(styles.StrokeAmberDark(scale))
	return c
}

func (c *component) StrokeAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeAmberDarkAlpha(scale))
	return c
}

func (c *component) StrokeBlack() *component {
	c.el.Style(styles.StrokeBlack())
	return c
}

func (c *component) StrokeBlackAlpha(scale int) *component {
	c.el.Style(styles.StrokeBlackAlpha(scale))
	return c
}

func (c *component) StrokeBlue(scale int) *component {
	c.el.Style(styles.StrokeBlue(scale))
	return c
}

func (c *component) StrokeBlueAlpha(scale int) *component {
	c.el.Style(styles.StrokeBlueAlpha(scale))
	return c
}

func (c *component) StrokeBlueDark(scale int) *component {
	c.el.Style(styles.StrokeBlueDark(scale))
	return c
}

func (c *component) StrokeBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeBlueDarkAlpha(scale))
	return c
}

func (c *component) StrokeBronze(scale int) *component {
	c.el.Style(styles.StrokeBronze(scale))
	return c
}

func (c *component) StrokeBronzeAlpha(scale int) *component {
	c.el.Style(styles.StrokeBronzeAlpha(scale))
	return c
}

func (c *component) StrokeBronzeDark(scale int) *component {
	c.el.Style(styles.StrokeBronzeDark(scale))
	return c
}

func (c *component) StrokeBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeBronzeDarkAlpha(scale))
	return c
}

func (c *component) StrokeBrown(scale int) *component {
	c.el.Style(styles.StrokeBrown(scale))
	return c
}

func (c *component) StrokeBrownAlpha(scale int) *component {
	c.el.Style(styles.StrokeBrownAlpha(scale))
	return c
}

func (c *component) StrokeBrownDark(scale int) *component {
	c.el.Style(styles.StrokeBrownDark(scale))
	return c
}

func (c *component) StrokeBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeBrownDarkAlpha(scale))
	return c
}

func (c *component) StrokeColor(val value.Value) *component {
	c.el.Style(styles.StrokeColor(val))
	return c
}

func (c *component) StrokeCrimson(scale int) *component {
	c.el.Style(styles.StrokeCrimson(scale))
	return c
}

func (c *component) StrokeCrimsonAlpha(scale int) *component {
	c.el.Style(styles.StrokeCrimsonAlpha(scale))
	return c
}

func (c *component) StrokeCrimsonDark(scale int) *component {
	c.el.Style(styles.StrokeCrimsonDark(scale))
	return c
}

func (c *component) StrokeCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeCrimsonDarkAlpha(scale))
	return c
}

func (c *component) StrokeCurrent() *component {
	c.el.Style(styles.StrokeCurrent())
	return c
}

func (c *component) StrokeCyan(scale int) *component {
	c.el.Style(styles.StrokeCyan(scale))
	return c
}

func (c *component) StrokeCyanAlpha(scale int) *component {
	c.el.Style(styles.StrokeCyanAlpha(scale))
	return c
}

func (c *component) StrokeCyanDark(scale int) *component {
	c.el.Style(styles.StrokeCyanDark(scale))
	return c
}

func (c *component) StrokeCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeCyanDarkAlpha(scale))
	return c
}

func (c *component) StrokeGold(scale int) *component {
	c.el.Style(styles.StrokeGold(scale))
	return c
}

func (c *component) StrokeGoldAlpha(scale int) *component {
	c.el.Style(styles.StrokeGoldAlpha(scale))
	return c
}

func (c *component) StrokeGoldDark(scale int) *component {
	c.el.Style(styles.StrokeGoldDark(scale))
	return c
}

func (c *component) StrokeGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeGoldDarkAlpha(scale))
	return c
}

func (c *component) StrokeGrass(scale int) *component {
	c.el.Style(styles.StrokeGrass(scale))
	return c
}

func (c *component) StrokeGrassAlpha(scale int) *component {
	c.el.Style(styles.StrokeGrassAlpha(scale))
	return c
}

func (c *component) StrokeGrassDark(scale int) *component {
	c.el.Style(styles.StrokeGrassDark(scale))
	return c
}

func (c *component) StrokeGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeGrassDarkAlpha(scale))
	return c
}

func (c *component) StrokeGray(scale int) *component {
	c.el.Style(styles.StrokeGray(scale))
	return c
}

func (c *component) StrokeGrayAlpha(scale int) *component {
	c.el.Style(styles.StrokeGrayAlpha(scale))
	return c
}

func (c *component) StrokeGrayDark(scale int) *component {
	c.el.Style(styles.StrokeGrayDark(scale))
	return c
}

func (c *component) StrokeGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeGrayDarkAlpha(scale))
	return c
}

func (c *component) StrokeGreen(scale int) *component {
	c.el.Style(styles.StrokeGreen(scale))
	return c
}

func (c *component) StrokeGreenAlpha(scale int) *component {
	c.el.Style(styles.StrokeGreenAlpha(scale))
	return c
}

func (c *component) StrokeGreenDark(scale int) *component {
	c.el.Style(styles.StrokeGreenDark(scale))
	return c
}

func (c *component) StrokeGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeGreenDarkAlpha(scale))
	return c
}

func (c *component) StrokeIndigo(scale int) *component {
	c.el.Style(styles.StrokeIndigo(scale))
	return c
}

func (c *component) StrokeIndigoAlpha(scale int) *component {
	c.el.Style(styles.StrokeIndigoAlpha(scale))
	return c
}

func (c *component) StrokeIndigoDark(scale int) *component {
	c.el.Style(styles.StrokeIndigoDark(scale))
	return c
}

func (c *component) StrokeIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeIndigoDarkAlpha(scale))
	return c
}

func (c *component) StrokeInherit() *component {
	c.el.Style(styles.StrokeInherit())
	return c
}

func (c *component) StrokeIris(scale int) *component {
	c.el.Style(styles.StrokeIris(scale))
	return c
}

func (c *component) StrokeIrisAlpha(scale int) *component {
	c.el.Style(styles.StrokeIrisAlpha(scale))
	return c
}

func (c *component) StrokeIrisDark(scale int) *component {
	c.el.Style(styles.StrokeIrisDark(scale))
	return c
}

func (c *component) StrokeIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeIrisDarkAlpha(scale))
	return c
}

func (c *component) StrokeJade(scale int) *component {
	c.el.Style(styles.StrokeJade(scale))
	return c
}

func (c *component) StrokeJadeAlpha(scale int) *component {
	c.el.Style(styles.StrokeJadeAlpha(scale))
	return c
}

func (c *component) StrokeJadeDark(scale int) *component {
	c.el.Style(styles.StrokeJadeDark(scale))
	return c
}

func (c *component) StrokeJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeJadeDarkAlpha(scale))
	return c
}

func (c *component) StrokeLime(scale int) *component {
	c.el.Style(styles.StrokeLime(scale))
	return c
}

func (c *component) StrokeLimeAlpha(scale int) *component {
	c.el.Style(styles.StrokeLimeAlpha(scale))
	return c
}

func (c *component) StrokeLimeDark(scale int) *component {
	c.el.Style(styles.StrokeLimeDark(scale))
	return c
}

func (c *component) StrokeLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeLimeDarkAlpha(scale))
	return c
}

func (c *component) StrokeMauve(scale int) *component {
	c.el.Style(styles.StrokeMauve(scale))
	return c
}

func (c *component) StrokeMauveAlpha(scale int) *component {
	c.el.Style(styles.StrokeMauveAlpha(scale))
	return c
}

func (c *component) StrokeMauveDark(scale int) *component {
	c.el.Style(styles.StrokeMauveDark(scale))
	return c
}

func (c *component) StrokeMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeMauveDarkAlpha(scale))
	return c
}

func (c *component) StrokeMint(scale int) *component {
	c.el.Style(styles.StrokeMint(scale))
	return c
}

func (c *component) StrokeMintAlpha(scale int) *component {
	c.el.Style(styles.StrokeMintAlpha(scale))
	return c
}

func (c *component) StrokeMintDark(scale int) *component {
	c.el.Style(styles.StrokeMintDark(scale))
	return c
}

func (c *component) StrokeMintDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeMintDarkAlpha(scale))
	return c
}

func (c *component) StrokeOlive(scale int) *component {
	c.el.Style(styles.StrokeOlive(scale))
	return c
}

func (c *component) StrokeOliveAlpha(scale int) *component {
	c.el.Style(styles.StrokeOliveAlpha(scale))
	return c
}

func (c *component) StrokeOliveDark(scale int) *component {
	c.el.Style(styles.StrokeOliveDark(scale))
	return c
}

func (c *component) StrokeOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeOliveDarkAlpha(scale))
	return c
}

func (c *component) StrokeOrange(scale int) *component {
	c.el.Style(styles.StrokeOrange(scale))
	return c
}

func (c *component) StrokeOrangeAlpha(scale int) *component {
	c.el.Style(styles.StrokeOrangeAlpha(scale))
	return c
}

func (c *component) StrokeOrangeDark(scale int) *component {
	c.el.Style(styles.StrokeOrangeDark(scale))
	return c
}

func (c *component) StrokeOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeOrangeDarkAlpha(scale))
	return c
}

func (c *component) StrokePink(scale int) *component {
	c.el.Style(styles.StrokePink(scale))
	return c
}

func (c *component) StrokePinkAlpha(scale int) *component {
	c.el.Style(styles.StrokePinkAlpha(scale))
	return c
}

func (c *component) StrokePinkDark(scale int) *component {
	c.el.Style(styles.StrokePinkDark(scale))
	return c
}

func (c *component) StrokePinkDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokePinkDarkAlpha(scale))
	return c
}

func (c *component) StrokePlum(scale int) *component {
	c.el.Style(styles.StrokePlum(scale))
	return c
}

func (c *component) StrokePlumAlpha(scale int) *component {
	c.el.Style(styles.StrokePlumAlpha(scale))
	return c
}

func (c *component) StrokePlumDark(scale int) *component {
	c.el.Style(styles.StrokePlumDark(scale))
	return c
}

func (c *component) StrokePlumDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokePlumDarkAlpha(scale))
	return c
}

func (c *component) StrokePurple(scale int) *component {
	c.el.Style(styles.StrokePurple(scale))
	return c
}

func (c *component) StrokePurpleAlpha(scale int) *component {
	c.el.Style(styles.StrokePurpleAlpha(scale))
	return c
}

func (c *component) StrokePurpleDark(scale int) *component {
	c.el.Style(styles.StrokePurpleDark(scale))
	return c
}

func (c *component) StrokePurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokePurpleDarkAlpha(scale))
	return c
}

func (c *component) StrokeRed(scale int) *component {
	c.el.Style(styles.StrokeRed(scale))
	return c
}

func (c *component) StrokeRedAlpha(scale int) *component {
	c.el.Style(styles.StrokeRedAlpha(scale))
	return c
}

func (c *component) StrokeRedDark(scale int) *component {
	c.el.Style(styles.StrokeRedDark(scale))
	return c
}

func (c *component) StrokeRedDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeRedDarkAlpha(scale))
	return c
}

func (c *component) StrokeRuby(scale int) *component {
	c.el.Style(styles.StrokeRuby(scale))
	return c
}

func (c *component) StrokeRubyAlpha(scale int) *component {
	c.el.Style(styles.StrokeRubyAlpha(scale))
	return c
}

func (c *component) StrokeRubyDark(scale int) *component {
	c.el.Style(styles.StrokeRubyDark(scale))
	return c
}

func (c *component) StrokeRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeRubyDarkAlpha(scale))
	return c
}

func (c *component) StrokeSage(scale int) *component {
	c.el.Style(styles.StrokeSage(scale))
	return c
}

func (c *component) StrokeSageAlpha(scale int) *component {
	c.el.Style(styles.StrokeSageAlpha(scale))
	return c
}

func (c *component) StrokeSageDark(scale int) *component {
	c.el.Style(styles.StrokeSageDark(scale))
	return c
}

func (c *component) StrokeSageDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeSageDarkAlpha(scale))
	return c
}

func (c *component) StrokeSand(scale int) *component {
	c.el.Style(styles.StrokeSand(scale))
	return c
}

func (c *component) StrokeSandAlpha(scale int) *component {
	c.el.Style(styles.StrokeSandAlpha(scale))
	return c
}

func (c *component) StrokeSandDark(scale int) *component {
	c.el.Style(styles.StrokeSandDark(scale))
	return c
}

func (c *component) StrokeSandDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeSandDarkAlpha(scale))
	return c
}

func (c *component) StrokeSky(scale int) *component {
	c.el.Style(styles.StrokeSky(scale))
	return c
}

func (c *component) StrokeSkyAlpha(scale int) *component {
	c.el.Style(styles.StrokeSkyAlpha(scale))
	return c
}

func (c *component) StrokeSkyDark(scale int) *component {
	c.el.Style(styles.StrokeSkyDark(scale))
	return c
}

func (c *component) StrokeSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeSkyDarkAlpha(scale))
	return c
}

func (c *component) StrokeSlate(scale int) *component {
	c.el.Style(styles.StrokeSlate(scale))
	return c
}

func (c *component) StrokeSlateAlpha(scale int) *component {
	c.el.Style(styles.StrokeSlateAlpha(scale))
	return c
}

func (c *component) StrokeSlateDark(scale int) *component {
	c.el.Style(styles.StrokeSlateDark(scale))
	return c
}

func (c *component) StrokeSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeSlateDarkAlpha(scale))
	return c
}

func (c *component) StrokeTeal(scale int) *component {
	c.el.Style(styles.StrokeTeal(scale))
	return c
}

func (c *component) StrokeTealAlpha(scale int) *component {
	c.el.Style(styles.StrokeTealAlpha(scale))
	return c
}

func (c *component) StrokeTealDark(scale int) *component {
	c.el.Style(styles.StrokeTealDark(scale))
	return c
}

func (c *component) StrokeTealDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeTealDarkAlpha(scale))
	return c
}

func (c *component) StrokeTomato(scale int) *component {
	c.el.Style(styles.StrokeTomato(scale))
	return c
}

func (c *component) StrokeTomatoAlpha(scale int) *component {
	c.el.Style(styles.StrokeTomatoAlpha(scale))
	return c
}

func (c *component) StrokeTomatoDark(scale int) *component {
	c.el.Style(styles.StrokeTomatoDark(scale))
	return c
}

func (c *component) StrokeTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeTomatoDarkAlpha(scale))
	return c
}

func (c *component) StrokeTransparent() *component {
	c.el.Style(styles.StrokeTransparent())
	return c
}

func (c *component) StrokeViolet(scale int) *component {
	c.el.Style(styles.StrokeViolet(scale))
	return c
}

func (c *component) StrokeVioletAlpha(scale int) *component {
	c.el.Style(styles.StrokeVioletAlpha(scale))
	return c
}

func (c *component) StrokeVioletDark(scale int) *component {
	c.el.Style(styles.StrokeVioletDark(scale))
	return c
}

func (c *component) StrokeVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeVioletDarkAlpha(scale))
	return c
}

func (c *component) StrokeWhite() *component {
	c.el.Style(styles.StrokeWhite())
	return c
}

func (c *component) StrokeWhiteAlpha(scale int) *component {
	c.el.Style(styles.StrokeWhiteAlpha(scale))
	return c
}

func (c *component) StrokeYellow(scale int) *component {
	c.el.Style(styles.StrokeYellow(scale))
	return c
}

func (c *component) StrokeYellowAlpha(scale int) *component {
	c.el.Style(styles.StrokeYellowAlpha(scale))
	return c
}

func (c *component) StrokeYellowDark(scale int) *component {
	c.el.Style(styles.StrokeYellowDark(scale))
	return c
}

func (c *component) StrokeYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.StrokeYellowDarkAlpha(scale))
	return c
}

func (c *component) SubPixelAntialiazed() *component {
	c.el.Style(styles.SubPixelAntialiazed())
	return c
}

func (c *component) Table() *component {
	c.el.Style(styles.Table())
	return c
}

func (c *component) TableAuto() *component {
	c.el.Style(styles.TableAuto())
	return c
}

func (c *component) TableCaption() *component {
	c.el.Style(styles.TableCaption())
	return c
}

func (c *component) TableCell() *component {
	c.el.Style(styles.TableCell())
	return c
}

func (c *component) TableColumn() *component {
	c.el.Style(styles.TableColumn())
	return c
}

func (c *component) TableColumnGroup() *component {
	c.el.Style(styles.TableColumnGroup())
	return c
}

func (c *component) TableFixed() *component {
	c.el.Style(styles.TableFixed())
	return c
}

func (c *component) TableFooterGroup() *component {
	c.el.Style(styles.TableFooterGroup())
	return c
}

func (c *component) TableHeaderGroup() *component {
	c.el.Style(styles.TableHeaderGroup())
	return c
}

func (c *component) TableRow() *component {
	c.el.Style(styles.TableRow())
	return c
}

func (c *component) TableRowGroup() *component {
	c.el.Style(styles.TableRowGroup())
	return c
}

func (c *component) TabularNums() *component {
	c.el.Style(styles.TabularNums())
	return c
}

func (c *component) Text2xl(lineHeights ...any) *component {
	c.el.Style(styles.Text2xl(lineHeights...))
	return c
}

func (c *component) Text3xl(lineHeights ...any) *component {
	c.el.Style(styles.Text3xl(lineHeights...))
	return c
}

func (c *component) Text4xl(lineHeights ...any) *component {
	c.el.Style(styles.Text4xl(lineHeights...))
	return c
}

func (c *component) Text5xl(lineHeights ...any) *component {
	c.el.Style(styles.Text5xl(lineHeights...))
	return c
}

func (c *component) Text6xl(lineHeights ...any) *component {
	c.el.Style(styles.Text6xl(lineHeights...))
	return c
}

func (c *component) Text7xl(lineHeights ...any) *component {
	c.el.Style(styles.Text7xl(lineHeights...))
	return c
}

func (c *component) Text8xl(lineHeights ...any) *component {
	c.el.Style(styles.Text8xl(lineHeights...))
	return c
}

func (c *component) Text9xl(lineHeights ...any) *component {
	c.el.Style(styles.Text9xl(lineHeights...))
	return c
}

func (c *component) TextAmber(scale int) *component {
	c.el.Style(styles.TextAmber(scale))
	return c
}

func (c *component) TextAmberAlpha(scale int) *component {
	c.el.Style(styles.TextAmberAlpha(scale))
	return c
}

func (c *component) TextAmberDark(scale int) *component {
	c.el.Style(styles.TextAmberDark(scale))
	return c
}

func (c *component) TextAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.TextAmberDarkAlpha(scale))
	return c
}

func (c *component) TextBalance() *component {
	c.el.Style(styles.TextBalance())
	return c
}

func (c *component) TextBase(lineHeights ...any) *component {
	c.el.Style(styles.TextBase(lineHeights...))
	return c
}

func (c *component) TextBlack() *component {
	c.el.Style(styles.TextBlack())
	return c
}

func (c *component) TextBlackAlpha(scale int) *component {
	c.el.Style(styles.TextBlackAlpha(scale))
	return c
}

func (c *component) TextBlue(scale int) *component {
	c.el.Style(styles.TextBlue(scale))
	return c
}

func (c *component) TextBlueAlpha(scale int) *component {
	c.el.Style(styles.TextBlueAlpha(scale))
	return c
}

func (c *component) TextBlueDark(scale int) *component {
	c.el.Style(styles.TextBlueDark(scale))
	return c
}

func (c *component) TextBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.TextBlueDarkAlpha(scale))
	return c
}

func (c *component) TextBronze(scale int) *component {
	c.el.Style(styles.TextBronze(scale))
	return c
}

func (c *component) TextBronzeAlpha(scale int) *component {
	c.el.Style(styles.TextBronzeAlpha(scale))
	return c
}

func (c *component) TextBronzeDark(scale int) *component {
	c.el.Style(styles.TextBronzeDark(scale))
	return c
}

func (c *component) TextBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.TextBronzeDarkAlpha(scale))
	return c
}

func (c *component) TextBrown(scale int) *component {
	c.el.Style(styles.TextBrown(scale))
	return c
}

func (c *component) TextBrownAlpha(scale int) *component {
	c.el.Style(styles.TextBrownAlpha(scale))
	return c
}

func (c *component) TextBrownDark(scale int) *component {
	c.el.Style(styles.TextBrownDark(scale))
	return c
}

func (c *component) TextBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.TextBrownDarkAlpha(scale))
	return c
}

func (c *component) TextCenter() *component {
	c.el.Style(styles.TextCenter())
	return c
}

func (c *component) TextClip() *component {
	c.el.Style(styles.TextClip())
	return c
}

func (c *component) TextColor(val value.Value) *component {
	c.el.Style(styles.TextColor(val))
	return c
}

func (c *component) TextCrimson(scale int) *component {
	c.el.Style(styles.TextCrimson(scale))
	return c
}

func (c *component) TextCrimsonAlpha(scale int) *component {
	c.el.Style(styles.TextCrimsonAlpha(scale))
	return c
}

func (c *component) TextCrimsonDark(scale int) *component {
	c.el.Style(styles.TextCrimsonDark(scale))
	return c
}

func (c *component) TextCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.TextCrimsonDarkAlpha(scale))
	return c
}

func (c *component) TextCurrent() *component {
	c.el.Style(styles.TextCurrent())
	return c
}

func (c *component) TextCyan(scale int) *component {
	c.el.Style(styles.TextCyan(scale))
	return c
}

func (c *component) TextCyanAlpha(scale int) *component {
	c.el.Style(styles.TextCyanAlpha(scale))
	return c
}

func (c *component) TextCyanDark(scale int) *component {
	c.el.Style(styles.TextCyanDark(scale))
	return c
}

func (c *component) TextCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.TextCyanDarkAlpha(scale))
	return c
}

func (c *component) TextEllipsis() *component {
	c.el.Style(styles.TextEllipsis())
	return c
}

func (c *component) TextEnd() *component {
	c.el.Style(styles.TextEnd())
	return c
}

func (c *component) TextGold(scale int) *component {
	c.el.Style(styles.TextGold(scale))
	return c
}

func (c *component) TextGoldAlpha(scale int) *component {
	c.el.Style(styles.TextGoldAlpha(scale))
	return c
}

func (c *component) TextGoldDark(scale int) *component {
	c.el.Style(styles.TextGoldDark(scale))
	return c
}

func (c *component) TextGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.TextGoldDarkAlpha(scale))
	return c
}

func (c *component) TextGrass(scale int) *component {
	c.el.Style(styles.TextGrass(scale))
	return c
}

func (c *component) TextGrassAlpha(scale int) *component {
	c.el.Style(styles.TextGrassAlpha(scale))
	return c
}

func (c *component) TextGrassDark(scale int) *component {
	c.el.Style(styles.TextGrassDark(scale))
	return c
}

func (c *component) TextGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.TextGrassDarkAlpha(scale))
	return c
}

func (c *component) TextGray(scale int) *component {
	c.el.Style(styles.TextGray(scale))
	return c
}

func (c *component) TextGrayAlpha(scale int) *component {
	c.el.Style(styles.TextGrayAlpha(scale))
	return c
}

func (c *component) TextGrayDark(scale int) *component {
	c.el.Style(styles.TextGrayDark(scale))
	return c
}

func (c *component) TextGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.TextGrayDarkAlpha(scale))
	return c
}

func (c *component) TextGreen(scale int) *component {
	c.el.Style(styles.TextGreen(scale))
	return c
}

func (c *component) TextGreenAlpha(scale int) *component {
	c.el.Style(styles.TextGreenAlpha(scale))
	return c
}

func (c *component) TextGreenDark(scale int) *component {
	c.el.Style(styles.TextGreenDark(scale))
	return c
}

func (c *component) TextGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.TextGreenDarkAlpha(scale))
	return c
}

func (c *component) TextIndigo(scale int) *component {
	c.el.Style(styles.TextIndigo(scale))
	return c
}

func (c *component) TextIndigoAlpha(scale int) *component {
	c.el.Style(styles.TextIndigoAlpha(scale))
	return c
}

func (c *component) TextIndigoDark(scale int) *component {
	c.el.Style(styles.TextIndigoDark(scale))
	return c
}

func (c *component) TextIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.TextIndigoDarkAlpha(scale))
	return c
}

func (c *component) TextInherit() *component {
	c.el.Style(styles.TextInherit())
	return c
}

func (c *component) TextIris(scale int) *component {
	c.el.Style(styles.TextIris(scale))
	return c
}

func (c *component) TextIrisAlpha(scale int) *component {
	c.el.Style(styles.TextIrisAlpha(scale))
	return c
}

func (c *component) TextIrisDark(scale int) *component {
	c.el.Style(styles.TextIrisDark(scale))
	return c
}

func (c *component) TextIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.TextIrisDarkAlpha(scale))
	return c
}

func (c *component) TextJade(scale int) *component {
	c.el.Style(styles.TextJade(scale))
	return c
}

func (c *component) TextJadeAlpha(scale int) *component {
	c.el.Style(styles.TextJadeAlpha(scale))
	return c
}

func (c *component) TextJadeDark(scale int) *component {
	c.el.Style(styles.TextJadeDark(scale))
	return c
}

func (c *component) TextJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.TextJadeDarkAlpha(scale))
	return c
}

func (c *component) TextJustify() *component {
	c.el.Style(styles.TextJustify())
	return c
}

func (c *component) TextLeft() *component {
	c.el.Style(styles.TextLeft())
	return c
}

func (c *component) TextLg(lineHeights ...any) *component {
	c.el.Style(styles.TextLg(lineHeights...))
	return c
}

func (c *component) TextLime(scale int) *component {
	c.el.Style(styles.TextLime(scale))
	return c
}

func (c *component) TextLimeAlpha(scale int) *component {
	c.el.Style(styles.TextLimeAlpha(scale))
	return c
}

func (c *component) TextLimeDark(scale int) *component {
	c.el.Style(styles.TextLimeDark(scale))
	return c
}

func (c *component) TextLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.TextLimeDarkAlpha(scale))
	return c
}

func (c *component) TextMauve(scale int) *component {
	c.el.Style(styles.TextMauve(scale))
	return c
}

func (c *component) TextMauveAlpha(scale int) *component {
	c.el.Style(styles.TextMauveAlpha(scale))
	return c
}

func (c *component) TextMauveDark(scale int) *component {
	c.el.Style(styles.TextMauveDark(scale))
	return c
}

func (c *component) TextMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.TextMauveDarkAlpha(scale))
	return c
}

func (c *component) TextMint(scale int) *component {
	c.el.Style(styles.TextMint(scale))
	return c
}

func (c *component) TextMintAlpha(scale int) *component {
	c.el.Style(styles.TextMintAlpha(scale))
	return c
}

func (c *component) TextMintDark(scale int) *component {
	c.el.Style(styles.TextMintDark(scale))
	return c
}

func (c *component) TextMintDarkAlpha(scale int) *component {
	c.el.Style(styles.TextMintDarkAlpha(scale))
	return c
}

func (c *component) TextNoWrap() *component {
	c.el.Style(styles.TextNoWrap())
	return c
}

func (c *component) TextOlive(scale int) *component {
	c.el.Style(styles.TextOlive(scale))
	return c
}

func (c *component) TextOliveAlpha(scale int) *component {
	c.el.Style(styles.TextOliveAlpha(scale))
	return c
}

func (c *component) TextOliveDark(scale int) *component {
	c.el.Style(styles.TextOliveDark(scale))
	return c
}

func (c *component) TextOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.TextOliveDarkAlpha(scale))
	return c
}

func (c *component) TextOrange(scale int) *component {
	c.el.Style(styles.TextOrange(scale))
	return c
}

func (c *component) TextOrangeAlpha(scale int) *component {
	c.el.Style(styles.TextOrangeAlpha(scale))
	return c
}

func (c *component) TextOrangeDark(scale int) *component {
	c.el.Style(styles.TextOrangeDark(scale))
	return c
}

func (c *component) TextOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.TextOrangeDarkAlpha(scale))
	return c
}

func (c *component) TextPink(scale int) *component {
	c.el.Style(styles.TextPink(scale))
	return c
}

func (c *component) TextPinkAlpha(scale int) *component {
	c.el.Style(styles.TextPinkAlpha(scale))
	return c
}

func (c *component) TextPinkDark(scale int) *component {
	c.el.Style(styles.TextPinkDark(scale))
	return c
}

func (c *component) TextPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.TextPinkDarkAlpha(scale))
	return c
}

func (c *component) TextPlum(scale int) *component {
	c.el.Style(styles.TextPlum(scale))
	return c
}

func (c *component) TextPlumAlpha(scale int) *component {
	c.el.Style(styles.TextPlumAlpha(scale))
	return c
}

func (c *component) TextPlumDark(scale int) *component {
	c.el.Style(styles.TextPlumDark(scale))
	return c
}

func (c *component) TextPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.TextPlumDarkAlpha(scale))
	return c
}

func (c *component) TextPretty() *component {
	c.el.Style(styles.TextPretty())
	return c
}

func (c *component) TextPurple(scale int) *component {
	c.el.Style(styles.TextPurple(scale))
	return c
}

func (c *component) TextPurpleAlpha(scale int) *component {
	c.el.Style(styles.TextPurpleAlpha(scale))
	return c
}

func (c *component) TextPurpleDark(scale int) *component {
	c.el.Style(styles.TextPurpleDark(scale))
	return c
}

func (c *component) TextPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.TextPurpleDarkAlpha(scale))
	return c
}

func (c *component) TextRed(scale int) *component {
	c.el.Style(styles.TextRed(scale))
	return c
}

func (c *component) TextRedAlpha(scale int) *component {
	c.el.Style(styles.TextRedAlpha(scale))
	return c
}

func (c *component) TextRedDark(scale int) *component {
	c.el.Style(styles.TextRedDark(scale))
	return c
}

func (c *component) TextRedDarkAlpha(scale int) *component {
	c.el.Style(styles.TextRedDarkAlpha(scale))
	return c
}

func (c *component) TextRight() *component {
	c.el.Style(styles.TextRight())
	return c
}

func (c *component) TextRuby(scale int) *component {
	c.el.Style(styles.TextRuby(scale))
	return c
}

func (c *component) TextRubyAlpha(scale int) *component {
	c.el.Style(styles.TextRubyAlpha(scale))
	return c
}

func (c *component) TextRubyDark(scale int) *component {
	c.el.Style(styles.TextRubyDark(scale))
	return c
}

func (c *component) TextRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.TextRubyDarkAlpha(scale))
	return c
}

func (c *component) TextSage(scale int) *component {
	c.el.Style(styles.TextSage(scale))
	return c
}

func (c *component) TextSageAlpha(scale int) *component {
	c.el.Style(styles.TextSageAlpha(scale))
	return c
}

func (c *component) TextSageDark(scale int) *component {
	c.el.Style(styles.TextSageDark(scale))
	return c
}

func (c *component) TextSageDarkAlpha(scale int) *component {
	c.el.Style(styles.TextSageDarkAlpha(scale))
	return c
}

func (c *component) TextSand(scale int) *component {
	c.el.Style(styles.TextSand(scale))
	return c
}

func (c *component) TextSandAlpha(scale int) *component {
	c.el.Style(styles.TextSandAlpha(scale))
	return c
}

func (c *component) TextSandDark(scale int) *component {
	c.el.Style(styles.TextSandDark(scale))
	return c
}

func (c *component) TextSandDarkAlpha(scale int) *component {
	c.el.Style(styles.TextSandDarkAlpha(scale))
	return c
}

func (c *component) TextShadow(val value.Value) *component {
	c.el.Style(styles.TextShadow(val))
	return c
}

func (c *component) TextShadow2xl() *component {
	c.el.Style(styles.TextShadow2xl())
	return c
}

func (c *component) TextShadow2xs() *component {
	c.el.Style(styles.TextShadow2xs())
	return c
}

func (c *component) TextShadowAmber(scale int) *component {
	c.el.Style(styles.TextShadowAmber(scale))
	return c
}

func (c *component) TextShadowAmberAlpha(scale int) *component {
	c.el.Style(styles.TextShadowAmberAlpha(scale))
	return c
}

func (c *component) TextShadowAmberDark(scale int) *component {
	c.el.Style(styles.TextShadowAmberDark(scale))
	return c
}

func (c *component) TextShadowAmberDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowAmberDarkAlpha(scale))
	return c
}

func (c *component) TextShadowBlack() *component {
	c.el.Style(styles.TextShadowBlack())
	return c
}

func (c *component) TextShadowBlackAlpha(scale int) *component {
	c.el.Style(styles.TextShadowBlackAlpha(scale))
	return c
}

func (c *component) TextShadowBlue(scale int) *component {
	c.el.Style(styles.TextShadowBlue(scale))
	return c
}

func (c *component) TextShadowBlueAlpha(scale int) *component {
	c.el.Style(styles.TextShadowBlueAlpha(scale))
	return c
}

func (c *component) TextShadowBlueDark(scale int) *component {
	c.el.Style(styles.TextShadowBlueDark(scale))
	return c
}

func (c *component) TextShadowBlueDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowBlueDarkAlpha(scale))
	return c
}

func (c *component) TextShadowBronze(scale int) *component {
	c.el.Style(styles.TextShadowBronze(scale))
	return c
}

func (c *component) TextShadowBronzeAlpha(scale int) *component {
	c.el.Style(styles.TextShadowBronzeAlpha(scale))
	return c
}

func (c *component) TextShadowBronzeDark(scale int) *component {
	c.el.Style(styles.TextShadowBronzeDark(scale))
	return c
}

func (c *component) TextShadowBronzeDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowBronzeDarkAlpha(scale))
	return c
}

func (c *component) TextShadowBrown(scale int) *component {
	c.el.Style(styles.TextShadowBrown(scale))
	return c
}

func (c *component) TextShadowBrownAlpha(scale int) *component {
	c.el.Style(styles.TextShadowBrownAlpha(scale))
	return c
}

func (c *component) TextShadowBrownDark(scale int) *component {
	c.el.Style(styles.TextShadowBrownDark(scale))
	return c
}

func (c *component) TextShadowBrownDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowBrownDarkAlpha(scale))
	return c
}

func (c *component) TextShadowColor(val value.Value) *component {
	c.el.Style(styles.TextShadowColor(val))
	return c
}

func (c *component) TextShadowCrimson(scale int) *component {
	c.el.Style(styles.TextShadowCrimson(scale))
	return c
}

func (c *component) TextShadowCrimsonAlpha(scale int) *component {
	c.el.Style(styles.TextShadowCrimsonAlpha(scale))
	return c
}

func (c *component) TextShadowCrimsonDark(scale int) *component {
	c.el.Style(styles.TextShadowCrimsonDark(scale))
	return c
}

func (c *component) TextShadowCrimsonDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowCrimsonDarkAlpha(scale))
	return c
}

func (c *component) TextShadowCurrent() *component {
	c.el.Style(styles.TextShadowCurrent())
	return c
}

func (c *component) TextShadowCyan(scale int) *component {
	c.el.Style(styles.TextShadowCyan(scale))
	return c
}

func (c *component) TextShadowCyanAlpha(scale int) *component {
	c.el.Style(styles.TextShadowCyanAlpha(scale))
	return c
}

func (c *component) TextShadowCyanDark(scale int) *component {
	c.el.Style(styles.TextShadowCyanDark(scale))
	return c
}

func (c *component) TextShadowCyanDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowCyanDarkAlpha(scale))
	return c
}

func (c *component) TextShadowGold(scale int) *component {
	c.el.Style(styles.TextShadowGold(scale))
	return c
}

func (c *component) TextShadowGoldAlpha(scale int) *component {
	c.el.Style(styles.TextShadowGoldAlpha(scale))
	return c
}

func (c *component) TextShadowGoldDark(scale int) *component {
	c.el.Style(styles.TextShadowGoldDark(scale))
	return c
}

func (c *component) TextShadowGoldDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowGoldDarkAlpha(scale))
	return c
}

func (c *component) TextShadowGrass(scale int) *component {
	c.el.Style(styles.TextShadowGrass(scale))
	return c
}

func (c *component) TextShadowGrassAlpha(scale int) *component {
	c.el.Style(styles.TextShadowGrassAlpha(scale))
	return c
}

func (c *component) TextShadowGrassDark(scale int) *component {
	c.el.Style(styles.TextShadowGrassDark(scale))
	return c
}

func (c *component) TextShadowGrassDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowGrassDarkAlpha(scale))
	return c
}

func (c *component) TextShadowGray(scale int) *component {
	c.el.Style(styles.TextShadowGray(scale))
	return c
}

func (c *component) TextShadowGrayAlpha(scale int) *component {
	c.el.Style(styles.TextShadowGrayAlpha(scale))
	return c
}

func (c *component) TextShadowGrayDark(scale int) *component {
	c.el.Style(styles.TextShadowGrayDark(scale))
	return c
}

func (c *component) TextShadowGrayDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowGrayDarkAlpha(scale))
	return c
}

func (c *component) TextShadowGreen(scale int) *component {
	c.el.Style(styles.TextShadowGreen(scale))
	return c
}

func (c *component) TextShadowGreenAlpha(scale int) *component {
	c.el.Style(styles.TextShadowGreenAlpha(scale))
	return c
}

func (c *component) TextShadowGreenDark(scale int) *component {
	c.el.Style(styles.TextShadowGreenDark(scale))
	return c
}

func (c *component) TextShadowGreenDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowGreenDarkAlpha(scale))
	return c
}

func (c *component) TextShadowIndigo(scale int) *component {
	c.el.Style(styles.TextShadowIndigo(scale))
	return c
}

func (c *component) TextShadowIndigoAlpha(scale int) *component {
	c.el.Style(styles.TextShadowIndigoAlpha(scale))
	return c
}

func (c *component) TextShadowIndigoDark(scale int) *component {
	c.el.Style(styles.TextShadowIndigoDark(scale))
	return c
}

func (c *component) TextShadowIndigoDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowIndigoDarkAlpha(scale))
	return c
}

func (c *component) TextShadowInherit() *component {
	c.el.Style(styles.TextShadowInherit())
	return c
}

func (c *component) TextShadowIris(scale int) *component {
	c.el.Style(styles.TextShadowIris(scale))
	return c
}

func (c *component) TextShadowIrisAlpha(scale int) *component {
	c.el.Style(styles.TextShadowIrisAlpha(scale))
	return c
}

func (c *component) TextShadowIrisDark(scale int) *component {
	c.el.Style(styles.TextShadowIrisDark(scale))
	return c
}

func (c *component) TextShadowIrisDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowIrisDarkAlpha(scale))
	return c
}

func (c *component) TextShadowJade(scale int) *component {
	c.el.Style(styles.TextShadowJade(scale))
	return c
}

func (c *component) TextShadowJadeAlpha(scale int) *component {
	c.el.Style(styles.TextShadowJadeAlpha(scale))
	return c
}

func (c *component) TextShadowJadeDark(scale int) *component {
	c.el.Style(styles.TextShadowJadeDark(scale))
	return c
}

func (c *component) TextShadowJadeDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowJadeDarkAlpha(scale))
	return c
}

func (c *component) TextShadowLg() *component {
	c.el.Style(styles.TextShadowLg())
	return c
}

func (c *component) TextShadowLime(scale int) *component {
	c.el.Style(styles.TextShadowLime(scale))
	return c
}

func (c *component) TextShadowLimeAlpha(scale int) *component {
	c.el.Style(styles.TextShadowLimeAlpha(scale))
	return c
}

func (c *component) TextShadowLimeDark(scale int) *component {
	c.el.Style(styles.TextShadowLimeDark(scale))
	return c
}

func (c *component) TextShadowLimeDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowLimeDarkAlpha(scale))
	return c
}

func (c *component) TextShadowMauve(scale int) *component {
	c.el.Style(styles.TextShadowMauve(scale))
	return c
}

func (c *component) TextShadowMauveAlpha(scale int) *component {
	c.el.Style(styles.TextShadowMauveAlpha(scale))
	return c
}

func (c *component) TextShadowMauveDark(scale int) *component {
	c.el.Style(styles.TextShadowMauveDark(scale))
	return c
}

func (c *component) TextShadowMauveDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowMauveDarkAlpha(scale))
	return c
}

func (c *component) TextShadowMd() *component {
	c.el.Style(styles.TextShadowMd())
	return c
}

func (c *component) TextShadowMint(scale int) *component {
	c.el.Style(styles.TextShadowMint(scale))
	return c
}

func (c *component) TextShadowMintAlpha(scale int) *component {
	c.el.Style(styles.TextShadowMintAlpha(scale))
	return c
}

func (c *component) TextShadowMintDark(scale int) *component {
	c.el.Style(styles.TextShadowMintDark(scale))
	return c
}

func (c *component) TextShadowMintDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowMintDarkAlpha(scale))
	return c
}

func (c *component) TextShadowNone() *component {
	c.el.Style(styles.TextShadowNone())
	return c
}

func (c *component) TextShadowOlive(scale int) *component {
	c.el.Style(styles.TextShadowOlive(scale))
	return c
}

func (c *component) TextShadowOliveAlpha(scale int) *component {
	c.el.Style(styles.TextShadowOliveAlpha(scale))
	return c
}

func (c *component) TextShadowOliveDark(scale int) *component {
	c.el.Style(styles.TextShadowOliveDark(scale))
	return c
}

func (c *component) TextShadowOliveDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowOliveDarkAlpha(scale))
	return c
}

func (c *component) TextShadowOrange(scale int) *component {
	c.el.Style(styles.TextShadowOrange(scale))
	return c
}

func (c *component) TextShadowOrangeAlpha(scale int) *component {
	c.el.Style(styles.TextShadowOrangeAlpha(scale))
	return c
}

func (c *component) TextShadowOrangeDark(scale int) *component {
	c.el.Style(styles.TextShadowOrangeDark(scale))
	return c
}

func (c *component) TextShadowOrangeDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowOrangeDarkAlpha(scale))
	return c
}

func (c *component) TextShadowPink(scale int) *component {
	c.el.Style(styles.TextShadowPink(scale))
	return c
}

func (c *component) TextShadowPinkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowPinkAlpha(scale))
	return c
}

func (c *component) TextShadowPinkDark(scale int) *component {
	c.el.Style(styles.TextShadowPinkDark(scale))
	return c
}

func (c *component) TextShadowPinkDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowPinkDarkAlpha(scale))
	return c
}

func (c *component) TextShadowPlum(scale int) *component {
	c.el.Style(styles.TextShadowPlum(scale))
	return c
}

func (c *component) TextShadowPlumAlpha(scale int) *component {
	c.el.Style(styles.TextShadowPlumAlpha(scale))
	return c
}

func (c *component) TextShadowPlumDark(scale int) *component {
	c.el.Style(styles.TextShadowPlumDark(scale))
	return c
}

func (c *component) TextShadowPlumDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowPlumDarkAlpha(scale))
	return c
}

func (c *component) TextShadowPurple(scale int) *component {
	c.el.Style(styles.TextShadowPurple(scale))
	return c
}

func (c *component) TextShadowPurpleAlpha(scale int) *component {
	c.el.Style(styles.TextShadowPurpleAlpha(scale))
	return c
}

func (c *component) TextShadowPurpleDark(scale int) *component {
	c.el.Style(styles.TextShadowPurpleDark(scale))
	return c
}

func (c *component) TextShadowPurpleDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowPurpleDarkAlpha(scale))
	return c
}

func (c *component) TextShadowRed(scale int) *component {
	c.el.Style(styles.TextShadowRed(scale))
	return c
}

func (c *component) TextShadowRedAlpha(scale int) *component {
	c.el.Style(styles.TextShadowRedAlpha(scale))
	return c
}

func (c *component) TextShadowRedDark(scale int) *component {
	c.el.Style(styles.TextShadowRedDark(scale))
	return c
}

func (c *component) TextShadowRedDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowRedDarkAlpha(scale))
	return c
}

func (c *component) TextShadowRuby(scale int) *component {
	c.el.Style(styles.TextShadowRuby(scale))
	return c
}

func (c *component) TextShadowRubyAlpha(scale int) *component {
	c.el.Style(styles.TextShadowRubyAlpha(scale))
	return c
}

func (c *component) TextShadowRubyDark(scale int) *component {
	c.el.Style(styles.TextShadowRubyDark(scale))
	return c
}

func (c *component) TextShadowRubyDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowRubyDarkAlpha(scale))
	return c
}

func (c *component) TextShadowSage(scale int) *component {
	c.el.Style(styles.TextShadowSage(scale))
	return c
}

func (c *component) TextShadowSageAlpha(scale int) *component {
	c.el.Style(styles.TextShadowSageAlpha(scale))
	return c
}

func (c *component) TextShadowSageDark(scale int) *component {
	c.el.Style(styles.TextShadowSageDark(scale))
	return c
}

func (c *component) TextShadowSageDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowSageDarkAlpha(scale))
	return c
}

func (c *component) TextShadowSand(scale int) *component {
	c.el.Style(styles.TextShadowSand(scale))
	return c
}

func (c *component) TextShadowSandAlpha(scale int) *component {
	c.el.Style(styles.TextShadowSandAlpha(scale))
	return c
}

func (c *component) TextShadowSandDark(scale int) *component {
	c.el.Style(styles.TextShadowSandDark(scale))
	return c
}

func (c *component) TextShadowSandDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowSandDarkAlpha(scale))
	return c
}

func (c *component) TextShadowSky(scale int) *component {
	c.el.Style(styles.TextShadowSky(scale))
	return c
}

func (c *component) TextShadowSkyAlpha(scale int) *component {
	c.el.Style(styles.TextShadowSkyAlpha(scale))
	return c
}

func (c *component) TextShadowSkyDark(scale int) *component {
	c.el.Style(styles.TextShadowSkyDark(scale))
	return c
}

func (c *component) TextShadowSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowSkyDarkAlpha(scale))
	return c
}

func (c *component) TextShadowSlate(scale int) *component {
	c.el.Style(styles.TextShadowSlate(scale))
	return c
}

func (c *component) TextShadowSlateAlpha(scale int) *component {
	c.el.Style(styles.TextShadowSlateAlpha(scale))
	return c
}

func (c *component) TextShadowSlateDark(scale int) *component {
	c.el.Style(styles.TextShadowSlateDark(scale))
	return c
}

func (c *component) TextShadowSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowSlateDarkAlpha(scale))
	return c
}

func (c *component) TextShadowSm() *component {
	c.el.Style(styles.TextShadowSm())
	return c
}

func (c *component) TextShadowTeal(scale int) *component {
	c.el.Style(styles.TextShadowTeal(scale))
	return c
}

func (c *component) TextShadowTealAlpha(scale int) *component {
	c.el.Style(styles.TextShadowTealAlpha(scale))
	return c
}

func (c *component) TextShadowTealDark(scale int) *component {
	c.el.Style(styles.TextShadowTealDark(scale))
	return c
}

func (c *component) TextShadowTealDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowTealDarkAlpha(scale))
	return c
}

func (c *component) TextShadowTomato(scale int) *component {
	c.el.Style(styles.TextShadowTomato(scale))
	return c
}

func (c *component) TextShadowTomatoAlpha(scale int) *component {
	c.el.Style(styles.TextShadowTomatoAlpha(scale))
	return c
}

func (c *component) TextShadowTomatoDark(scale int) *component {
	c.el.Style(styles.TextShadowTomatoDark(scale))
	return c
}

func (c *component) TextShadowTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowTomatoDarkAlpha(scale))
	return c
}

func (c *component) TextShadowTransparent() *component {
	c.el.Style(styles.TextShadowTransparent())
	return c
}

func (c *component) TextShadowViolet(scale int) *component {
	c.el.Style(styles.TextShadowViolet(scale))
	return c
}

func (c *component) TextShadowVioletAlpha(scale int) *component {
	c.el.Style(styles.TextShadowVioletAlpha(scale))
	return c
}

func (c *component) TextShadowVioletDark(scale int) *component {
	c.el.Style(styles.TextShadowVioletDark(scale))
	return c
}

func (c *component) TextShadowVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowVioletDarkAlpha(scale))
	return c
}

func (c *component) TextShadowWhite() *component {
	c.el.Style(styles.TextShadowWhite())
	return c
}

func (c *component) TextShadowWhiteAlpha(scale int) *component {
	c.el.Style(styles.TextShadowWhiteAlpha(scale))
	return c
}

func (c *component) TextShadowXl() *component {
	c.el.Style(styles.TextShadowXl())
	return c
}

func (c *component) TextShadowXs() *component {
	c.el.Style(styles.TextShadowXs())
	return c
}

func (c *component) TextShadowYellow(scale int) *component {
	c.el.Style(styles.TextShadowYellow(scale))
	return c
}

func (c *component) TextShadowYellowAlpha(scale int) *component {
	c.el.Style(styles.TextShadowYellowAlpha(scale))
	return c
}

func (c *component) TextShadowYellowDark(scale int) *component {
	c.el.Style(styles.TextShadowYellowDark(scale))
	return c
}

func (c *component) TextShadowYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.TextShadowYellowDarkAlpha(scale))
	return c
}

func (c *component) TextSizeBy(val value.Value) *component {
	c.el.Style(styles.TextSizeBy(val))
	return c
}

func (c *component) TextSky(scale int) *component {
	c.el.Style(styles.TextSky(scale))
	return c
}

func (c *component) TextSkyAlpha(scale int) *component {
	c.el.Style(styles.TextSkyAlpha(scale))
	return c
}

func (c *component) TextSkyDark(scale int) *component {
	c.el.Style(styles.TextSkyDark(scale))
	return c
}

func (c *component) TextSkyDarkAlpha(scale int) *component {
	c.el.Style(styles.TextSkyDarkAlpha(scale))
	return c
}

func (c *component) TextSlate(scale int) *component {
	c.el.Style(styles.TextSlate(scale))
	return c
}

func (c *component) TextSlateAlpha(scale int) *component {
	c.el.Style(styles.TextSlateAlpha(scale))
	return c
}

func (c *component) TextSlateDark(scale int) *component {
	c.el.Style(styles.TextSlateDark(scale))
	return c
}

func (c *component) TextSlateDarkAlpha(scale int) *component {
	c.el.Style(styles.TextSlateDarkAlpha(scale))
	return c
}

func (c *component) TextSm(lineHeights ...any) *component {
	c.el.Style(styles.TextSm(lineHeights...))
	return c
}

func (c *component) TextStart() *component {
	c.el.Style(styles.TextStart())
	return c
}

func (c *component) TextTeal(scale int) *component {
	c.el.Style(styles.TextTeal(scale))
	return c
}

func (c *component) TextTealAlpha(scale int) *component {
	c.el.Style(styles.TextTealAlpha(scale))
	return c
}

func (c *component) TextTealDark(scale int) *component {
	c.el.Style(styles.TextTealDark(scale))
	return c
}

func (c *component) TextTealDarkAlpha(scale int) *component {
	c.el.Style(styles.TextTealDarkAlpha(scale))
	return c
}

func (c *component) TextTomato(scale int) *component {
	c.el.Style(styles.TextTomato(scale))
	return c
}

func (c *component) TextTomatoAlpha(scale int) *component {
	c.el.Style(styles.TextTomatoAlpha(scale))
	return c
}

func (c *component) TextTomatoDark(scale int) *component {
	c.el.Style(styles.TextTomatoDark(scale))
	return c
}

func (c *component) TextTomatoDarkAlpha(scale int) *component {
	c.el.Style(styles.TextTomatoDarkAlpha(scale))
	return c
}

func (c *component) TextTransparent() *component {
	c.el.Style(styles.TextTransparent())
	return c
}

func (c *component) TextViolet(scale int) *component {
	c.el.Style(styles.TextViolet(scale))
	return c
}

func (c *component) TextVioletAlpha(scale int) *component {
	c.el.Style(styles.TextVioletAlpha(scale))
	return c
}

func (c *component) TextVioletDark(scale int) *component {
	c.el.Style(styles.TextVioletDark(scale))
	return c
}

func (c *component) TextVioletDarkAlpha(scale int) *component {
	c.el.Style(styles.TextVioletDarkAlpha(scale))
	return c
}

func (c *component) TextWhite() *component {
	c.el.Style(styles.TextWhite())
	return c
}

func (c *component) TextWhiteAlpha(scale int) *component {
	c.el.Style(styles.TextWhiteAlpha(scale))
	return c
}

func (c *component) TextWrap() *component {
	c.el.Style(styles.TextWrap())
	return c
}

func (c *component) TextXl(lineHeights ...any) *component {
	c.el.Style(styles.TextXl(lineHeights...))
	return c
}

func (c *component) TextXs(lineHeights ...any) *component {
	c.el.Style(styles.TextXs(lineHeights...))
	return c
}

func (c *component) TextYellow(scale int) *component {
	c.el.Style(styles.TextYellow(scale))
	return c
}

func (c *component) TextYellowAlpha(scale int) *component {
	c.el.Style(styles.TextYellowAlpha(scale))
	return c
}

func (c *component) TextYellowDark(scale int) *component {
	c.el.Style(styles.TextYellowDark(scale))
	return c
}

func (c *component) TextYellowDarkAlpha(scale int) *component {
	c.el.Style(styles.TextYellowDarkAlpha(scale))
	return c
}

func (c *component) To(val value.Value) *component {
	c.el.Style(styles.To(val))
	return c
}

func (c *component) Top(number int) *component {
	c.el.Style(styles.Top(number))
	return c
}

func (c *component) TopAuto() *component {
	c.el.Style(styles.TopAuto())
	return c
}

func (c *component) TopBy(val value.Value) *component {
	c.el.Style(styles.TopBy(val))
	return c
}

func (c *component) TopFraction(fraction float64) *component {
	c.el.Style(styles.TopFraction(fraction))
	return c
}

func (c *component) TopFull() *component {
	c.el.Style(styles.TopFull())
	return c
}

func (c *component) TopPx() *component {
	c.el.Style(styles.TopPx())
	return c
}

func (c *component) TouchAuto() *component {
	c.el.Style(styles.TouchAuto())
	return c
}

func (c *component) TouchManipulation() *component {
	c.el.Style(styles.TouchManipulation())
	return c
}

func (c *component) TouchNone() *component {
	c.el.Style(styles.TouchNone())
	return c
}

func (c *component) TouchPanDown() *component {
	c.el.Style(styles.TouchPanDown())
	return c
}

func (c *component) TouchPanLeft() *component {
	c.el.Style(styles.TouchPanLeft())
	return c
}

func (c *component) TouchPanRight() *component {
	c.el.Style(styles.TouchPanRight())
	return c
}

func (c *component) TouchPanUp() *component {
	c.el.Style(styles.TouchPanUp())
	return c
}

func (c *component) TouchPanX() *component {
	c.el.Style(styles.TouchPanX())
	return c
}

func (c *component) TouchPanY() *component {
	c.el.Style(styles.TouchPanY())
	return c
}

func (c *component) TouchPinchZoom() *component {
	c.el.Style(styles.TouchPinchZoom())
	return c
}

func (c *component) TrackingBy(val value.Value) *component {
	c.el.Style(styles.TrackingBy(val))
	return c
}

func (c *component) TrackingNormal() *component {
	c.el.Style(styles.TrackingNormal())
	return c
}

func (c *component) TrackingTight() *component {
	c.el.Style(styles.TrackingTight())
	return c
}

func (c *component) TrackingTighter() *component {
	c.el.Style(styles.TrackingTighter())
	return c
}

func (c *component) TrackingWide() *component {
	c.el.Style(styles.TrackingWide())
	return c
}

func (c *component) TrackingWider() *component {
	c.el.Style(styles.TrackingWider())
	return c
}

func (c *component) TrackingWidest() *component {
	c.el.Style(styles.TrackingWidest())
	return c
}

func (c *component) Transform(val value.Value) *component {
	c.el.Style(styles.Transform(val))
	return c
}

func (c *component) Transform3d() *component {
	c.el.Style(styles.Transform3d())
	return c
}

func (c *component) TransformCpu() *component {
	c.el.Style(styles.TransformCpu())
	return c
}

func (c *component) TransformFkat() *component {
	c.el.Style(styles.TransformFkat())
	return c
}

func (c *component) TransformGpu() *component {
	c.el.Style(styles.TransformGpu())
	return c
}

func (c *component) TransformNone() *component {
	c.el.Style(styles.TransformNone())
	return c
}

func (c *component) Transition(val ...value.Value) *component {
	c.el.Style(styles.Transition(val...))
	return c
}

func (c *component) TransitionAll() *component {
	c.el.Style(styles.TransitionAll())
	return c
}

func (c *component) TransitionColors() *component {
	c.el.Style(styles.TransitionColors())
	return c
}

func (c *component) TransitionDiscrete() *component {
	c.el.Style(styles.TransitionDiscrete())
	return c
}

func (c *component) TransitionNone() *component {
	c.el.Style(styles.TransitionNone())
	return c
}

func (c *component) TransitionNormal() *component {
	c.el.Style(styles.TransitionNormal())
	return c
}

func (c *component) TransitionOpacity() *component {
	c.el.Style(styles.TransitionOpacity())
	return c
}

func (c *component) TransitionShadow() *component {
	c.el.Style(styles.TransitionShadow())
	return c
}

func (c *component) TransitionTransform() *component {
	c.el.Style(styles.TransitionTransform())
	return c
}

func (c *component) Translate(val any) *component {
	c.el.Style(styles.Translate(val))
	return c
}

func (c *component) TranslateFull() *component {
	c.el.Style(styles.TranslateFull())
	return c
}

func (c *component) TranslateNone() *component {
	c.el.Style(styles.TranslateNone())
	return c
}

func (c *component) TranslatePx() *component {
	c.el.Style(styles.TranslatePx())
	return c
}

func (c *component) TranslateX(val any) *component {
	c.el.Style(styles.TranslateX(val))
	return c
}

func (c *component) TranslateXFull() *component {
	c.el.Style(styles.TranslateXFull())
	return c
}

func (c *component) TranslateXPx() *component {
	c.el.Style(styles.TranslateXPx())
	return c
}

func (c *component) TranslateY(val any) *component {
	c.el.Style(styles.TranslateY(val))
	return c
}

func (c *component) TranslateYFull() *component {
	c.el.Style(styles.TranslateYFull())
	return c
}

func (c *component) TranslateYPx() *component {
	c.el.Style(styles.TranslateYPx())
	return c
}

func (c *component) TranslateZ(val any) *component {
	c.el.Style(styles.TranslateZ(val))
	return c
}

func (c *component) TranslateZPx() *component {
	c.el.Style(styles.TranslateZPx())
	return c
}

func (c *component) Truncate() *component {
	c.el.Style(styles.Truncate())
	return c
}

func (c *component) Underline() *component {
	c.el.Style(styles.Underline())
	return c
}

func (c *component) UnderlineOffset(number int) *component {
	c.el.Style(styles.UnderlineOffset(number))
	return c
}

func (c *component) UnderlineOffsetAuto() *component {
	c.el.Style(styles.UnderlineOffsetAuto())
	return c
}

func (c *component) UnderlineOffsetBy(val value.Value) *component {
	c.el.Style(styles.UnderlineOffsetBy(val))
	return c
}

func (c *component) Uppercase() *component {
	c.el.Style(styles.Uppercase())
	return c
}

func (c *component) Via(val value.Value) *component {
	c.el.Style(styles.Via(val))
	return c
}

func (c *component) Visible() *component {
	c.el.Style(styles.Visible())
	return c
}

func (c *component) W(number int) *component {
	c.el.Style(styles.W(number))
	return c
}

func (c *component) W2xl() *component {
	c.el.Style(styles.W2xl())
	return c
}

func (c *component) W2xs() *component {
	c.el.Style(styles.W2xs())
	return c
}

func (c *component) W3xl() *component {
	c.el.Style(styles.W3xl())
	return c
}

func (c *component) W3xs() *component {
	c.el.Style(styles.W3xs())
	return c
}

func (c *component) W4xl() *component {
	c.el.Style(styles.W4xl())
	return c
}

func (c *component) W5xl() *component {
	c.el.Style(styles.W5xl())
	return c
}

func (c *component) W6xl() *component {
	c.el.Style(styles.W6xl())
	return c
}

func (c *component) W7xl() *component {
	c.el.Style(styles.W7xl())
	return c
}

func (c *component) WAuto() *component {
	c.el.Style(styles.WAuto())
	return c
}

func (c *component) WBy(val value.Value) *component {
	c.el.Style(styles.WBy(val))
	return c
}

func (c *component) WDvh() *component {
	c.el.Style(styles.WDvh())
	return c
}

func (c *component) WDvw() *component {
	c.el.Style(styles.WDvw())
	return c
}

func (c *component) WFit() *component {
	c.el.Style(styles.WFit())
	return c
}

func (c *component) WFraction(fraction float32) *component {
	c.el.Style(styles.WFraction(fraction))
	return c
}

func (c *component) WFull() *component {
	c.el.Style(styles.WFull())
	return c
}

func (c *component) WLg() *component {
	c.el.Style(styles.WLg())
	return c
}

func (c *component) WLvh() *component {
	c.el.Style(styles.WLvh())
	return c
}

func (c *component) WLvw() *component {
	c.el.Style(styles.WLvw())
	return c
}

func (c *component) WMax() *component {
	c.el.Style(styles.WMax())
	return c
}

func (c *component) WMd() *component {
	c.el.Style(styles.WMd())
	return c
}

func (c *component) WMin() *component {
	c.el.Style(styles.WMin())
	return c
}

func (c *component) WPx() *component {
	c.el.Style(styles.WPx())
	return c
}

func (c *component) WScreen() *component {
	c.el.Style(styles.WScreen())
	return c
}

func (c *component) WSm() *component {
	c.el.Style(styles.WSm())
	return c
}

func (c *component) WSvh() *component {
	c.el.Style(styles.WSvh())
	return c
}

func (c *component) WSvw() *component {
	c.el.Style(styles.WSvw())
	return c
}

func (c *component) WXl() *component {
	c.el.Style(styles.WXl())
	return c
}

func (c *component) WXs() *component {
	c.el.Style(styles.WXs())
	return c
}

func (c *component) WhitespaceBreakSpaces() *component {
	c.el.Style(styles.WhitespaceBreakSpaces())
	return c
}

func (c *component) WhitespaceNormal() *component {
	c.el.Style(styles.WhitespaceNormal())
	return c
}

func (c *component) WhitespaceNowrap() *component {
	c.el.Style(styles.WhitespaceNowrap())
	return c
}

func (c *component) WhitespacePre() *component {
	c.el.Style(styles.WhitespacePre())
	return c
}

func (c *component) WhitespacePreLine() *component {
	c.el.Style(styles.WhitespacePreLine())
	return c
}

func (c *component) WhitespacePreWrap() *component {
	c.el.Style(styles.WhitespacePreWrap())
	return c
}

func (c *component) WillChange(val value.Value) *component {
	c.el.Style(styles.WillChange(val))
	return c
}

func (c *component) WillChangeAuto() *component {
	c.el.Style(styles.WillChangeAuto())
	return c
}

func (c *component) WillChangeContents() *component {
	c.el.Style(styles.WillChangeContents())
	return c
}

func (c *component) WillChangeScroll() *component {
	c.el.Style(styles.WillChangeScroll())
	return c
}

func (c *component) WillChangeTransform() *component {
	c.el.Style(styles.WillChangeTransform())
	return c
}

func (c *component) WrapAnywhere() *component {
	c.el.Style(styles.WrapAnywhere())
	return c
}

func (c *component) WrapBreakWord() *component {
	c.el.Style(styles.WrapBreakWord())
	return c
}

func (c *component) WrapNormal() *component {
	c.el.Style(styles.WrapNormal())
	return c
}

func (c *component) Z(index int) *component {
	c.el.Style(styles.Z(index))
	return c
}

func (c *component) ZAuto() *component {
	c.el.Style(styles.ZAuto())
	return c
}

func (c *component) ZBy(val value.Value) *component {
	c.el.Style(styles.ZBy(val))
	return c
}
