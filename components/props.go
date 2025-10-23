package components

import (
	"github.com/raitucarp/gomix/styles"
	"github.com/raitucarp/gomix/value"
)

func (c *Comp) StyleHover(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Hover(props...))
	return c
}

func (c *Comp) StyleFocus(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Focus(props...))
	return c
}

func (c *Comp) StyleFocusWithin(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.FocusWithin(props...))
	return c
}

func (c *Comp) StyleFocusVisible(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.FocusVisible(props...))
	return c
}

func (c *Comp) StyleActive(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Active(props...))
	return c
}

func (c *Comp) StyleVisited(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Visited(props...))
	return c
}

func (c *Comp) StyleTarget(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Target(props...))
	return c
}

func (c *Comp) StyleDirectChildren(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.DirectChildren(props...))
	return c
}

func (c *Comp) StyleAllDescendants(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.AllDescendants(props...))
	return c
}

func (c *Comp) StyleHas(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Has(props...))
	return c
}

func (c *Comp) StyleGroup(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Group(props...))
	return c
}

func (c *Comp) StylePeer(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Peer(props...))
	return c
}

func (c *Comp) StyleIn(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.In(props...))
	return c
}

func (c *Comp) StyleNot(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Not(props...))
	return c
}

func (c *Comp) StyleInert(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Inert(props...))
	return c
}

func (c *Comp) StyleFirst(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.First(props...))
	return c
}

func (c *Comp) StyleLast(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Last(props...))
	return c
}

func (c *Comp) StyleOnly(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Only(props...))
	return c
}

func (c *Comp) StyleOdd(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Odd(props...))
	return c
}

func (c *Comp) StyleEven(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Even(props...))
	return c
}

func (c *Comp) StyleFirstOfType(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.FirstOfType(props...))
	return c
}

func (c *Comp) StyleLastOfType(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.LastOfType(props...))
	return c
}

func (c *Comp) StyleOnlyOfType(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.OnlyOfType(props...))
	return c
}

func (c *Comp) StyleNth(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Nth(props...))
	return c
}

func (c *Comp) StyleNthLast(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.NthLast(props...))
	return c
}

func (c *Comp) StyleNthOfType(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.NthOfType(props...))
	return c
}

func (c *Comp) StyleNthLastOfType(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.NthLastOfType(props...))
	return c
}

func (c *Comp) StyleEmpty(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Empty(props...))
	return c
}

func (c *Comp) StyleDisabled(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Disabled(props...))
	return c
}

func (c *Comp) StyleEnabled(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Enabled(props...))
	return c
}

func (c *Comp) StyleChecked(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Checked(props...))
	return c
}

func (c *Comp) StyleIndeterminate(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Indeterminate(props...))
	return c
}

func (c *Comp) StyleDef(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Def(props...))
	return c
}

func (c *Comp) StyleOptional(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Optional(props...))
	return c
}

func (c *Comp) StyleRequired(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Required(props...))
	return c
}

func (c *Comp) StyleValid(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Valid(props...))
	return c
}

func (c *Comp) StyleInvalid(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Invalid(props...))
	return c
}

func (c *Comp) StyleUserValid(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.UserValid(props...))
	return c
}

func (c *Comp) StyleUserInvalid(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.UserInvalid(props...))
	return c
}

func (c *Comp) StyleInRange(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.InRange(props...))
	return c
}

func (c *Comp) StyleOutOfRange(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.OutOfRange(props...))
	return c
}

func (c *Comp) StylePlaceholderShown(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.PlaceholderShown(props...))
	return c
}

func (c *Comp) StyleDetailsContent(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.DetailsContent(props...))
	return c
}

func (c *Comp) StyleAutofill(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Autofill(props...))
	return c
}

func (c *Comp) StyleReadOnly(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ReadOnly(props...))
	return c
}

func (c *Comp) StyleBefore(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Before(props...))
	return c
}

func (c *Comp) StyleAfter(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.After(props...))
	return c
}

func (c *Comp) StyleFirstLetter(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.FirstLetter(props...))
	return c
}

func (c *Comp) StyleFirstLine(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.FirstLine(props...))
	return c
}

func (c *Comp) StyleMarker(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Marker(props...))
	return c
}

func (c *Comp) StyleSelection(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Selection(props...))
	return c
}

func (c *Comp) StyleFile(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.File(props...))
	return c
}

func (c *Comp) StyleBackdrop(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Backdrop(props...))
	return c
}

func (c *Comp) StylePlaceholder(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Placeholder(props...))
	return c
}

func (c *Comp) StyleSm(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Sm(props...))
	return c
}

func (c *Comp) StyleMd(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Md(props...))
	return c
}

func (c *Comp) StyleLg(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Lg(props...))
	return c
}

func (c *Comp) StyleXl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Xl(props...))
	return c
}

func (c *Comp) StyleTwoXl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.TwoXl(props...))
	return c
}

func (c *Comp) StyleMin(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Min(props...))
	return c
}

func (c *Comp) StyleMaxSm(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.MaxSm(props...))
	return c
}

func (c *Comp) StyleMaxMd(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.MaxMd(props...))
	return c
}

func (c *Comp) StyleMaxLg(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.MaxLg(props...))
	return c
}

func (c *Comp) StyleMaxXl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.MaxXl(props...))
	return c
}

func (c *Comp) StyleMax2xl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Max2xl(props...))
	return c
}

func (c *Comp) StyleMax(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Max(props...))
	return c
}

func (c *Comp) StyleContainer3xs(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Container3xs(props...))
	return c
}

func (c *Comp) StyleContainer2xs(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Container2xs(props...))
	return c
}

func (c *Comp) StyleContainerXs(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerXs(props...))
	return c
}

func (c *Comp) StyleContainerSm(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerSm(props...))
	return c
}

func (c *Comp) StyleContainerMd(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMd(props...))
	return c
}

func (c *Comp) StyleContainerLg(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerLg(props...))
	return c
}

func (c *Comp) StyleContainerXl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerXl(props...))
	return c
}

func (c *Comp) StyleContainer2xl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Container2xl(props...))
	return c
}

func (c *Comp) StyleContainer3xl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Container3xl(props...))
	return c
}

func (c *Comp) StyleContainer4xl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Container4xl(props...))
	return c
}

func (c *Comp) StyleContainer5xl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Container5xl(props...))
	return c
}

func (c *Comp) StyleContainer6xl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Container6xl(props...))
	return c
}

func (c *Comp) StyleContainer7xl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Container7xl(props...))
	return c
}

func (c *Comp) StyleContainerMin(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMin(props...))
	return c
}

func (c *Comp) StyleContainerMax3xs(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMax3xs(props...))
	return c
}

func (c *Comp) StyleContainerMax2xs(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMax2xs(props...))
	return c
}

func (c *Comp) StyleContainerMaxXs(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMaxXs(props...))
	return c
}

func (c *Comp) StyleContainerMaxSm(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMaxSm(props...))
	return c
}

func (c *Comp) StyleContainerMaxMd(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMaxMd(props...))
	return c
}

func (c *Comp) StyleContainerMaxLg(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMaxLg(props...))
	return c
}

func (c *Comp) StyleContainerMaxXl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMaxXl(props...))
	return c
}

func (c *Comp) StyleContainerMax2xl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMax2xl(props...))
	return c
}

func (c *Comp) StyleContainerMax3xl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMax3xl(props...))
	return c
}

func (c *Comp) StyleContainerMax4xl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMax4xl(props...))
	return c
}

func (c *Comp) StyleContainerMax5xl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMax5xl(props...))
	return c
}

func (c *Comp) StyleContainerMax6xl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMax6xl(props...))
	return c
}

func (c *Comp) StyleContainerMax7xl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMax7xl(props...))
	return c
}

func (c *Comp) StyleContainerMax(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContainerMax(props...))
	return c
}

func (c *Comp) StyleDark(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Dark(props...))
	return c
}

func (c *Comp) StyleMotionSafe(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.MotionSafe(props...))
	return c
}

func (c *Comp) StyleMotionReduce(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.MotionReduce(props...))
	return c
}

func (c *Comp) StyleContrastMore(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContrastMore(props...))
	return c
}

func (c *Comp) StyleContrastLess(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ContrastLess(props...))
	return c
}

func (c *Comp) StyleForcedColors(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.ForcedColors(props...))
	return c
}

func (c *Comp) StyleInvertedColors(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.InvertedColors(props...))
	return c
}

func (c *Comp) StylePointerFine(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.PointerFine(props...))
	return c
}

func (c *Comp) StylePointerCoarse(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.PointerCoarse(props...))
	return c
}

func (c *Comp) StylePointerNone(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.PointerNone(props...))
	return c
}

func (c *Comp) StyleAnyPointerFine(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.AnyPointerFine(props...))
	return c
}

func (c *Comp) StyleAnyPointerCoarse(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.AnyPointerCoarse(props...))
	return c
}

func (c *Comp) StyleAnyPointerNone(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.AnyPointerNone(props...))
	return c
}

func (c *Comp) StylePortrait(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Portrait(props...))
	return c
}

func (c *Comp) StyleLandscape(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Landscape(props...))
	return c
}

func (c *Comp) StyleNoscript(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Noscript(props...))
	return c
}

func (c *Comp) StylePrint(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Print(props...))
	return c
}

func (c *Comp) StyleSupports(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Supports(props...))
	return c
}

func (c *Comp) StyleAriaBusy(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.AriaBusy(props...))
	return c
}

func (c *Comp) StyleAriaChecked(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.AriaChecked(props...))
	return c
}

func (c *Comp) StyleAriaDisabled(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.AriaDisabled(props...))
	return c
}

func (c *Comp) StyleAriaExpanded(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.AriaExpanded(props...))
	return c
}

func (c *Comp) StyleAriaHidden(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.AriaHidden(props...))
	return c
}

func (c *Comp) StyleAriaPressed(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.AriaPressed(props...))
	return c
}

func (c *Comp) StyleAriaReadonly(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.AriaReadonly(props...))
	return c
}

func (c *Comp) StyleAriaRequired(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.AriaRequired(props...))
	return c
}

func (c *Comp) StyleAriaSelected(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.AriaSelected(props...))
	return c
}

func (c *Comp) StyleAria(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Aria(props...))
	return c
}

func (c *Comp) StyleData(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Data(props...))
	return c
}

func (c *Comp) StyleRtl(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Rtl(props...))
	return c
}

func (c *Comp) StyleLtr(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Ltr(props...))
	return c
}

func (c *Comp) StyleOpen(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Open(props...))
	return c
}

func (c *Comp) StyleStarting(props ...styles.ApplyProp) *Comp {
	c.El.Style(styles.Starting(props...))
	return c
}

func (c *Comp) Absolute() *Comp {
	c.El.Style(styles.Absolute())
	return c
}

func (c *Comp) Accent(val value.Value) *Comp {
	c.El.Style(styles.Accent(val))
	return c
}

func (c *Comp) AccentAmber(scale int) *Comp {
	c.El.Style(styles.AccentAmber(scale))
	return c
}

func (c *Comp) AccentAmberAlpha(scale int) *Comp {
	c.El.Style(styles.AccentAmberAlpha(scale))
	return c
}

func (c *Comp) AccentAmberDark(scale int) *Comp {
	c.El.Style(styles.AccentAmberDark(scale))
	return c
}

func (c *Comp) AccentAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentAmberDarkAlpha(scale))
	return c
}

func (c *Comp) AccentBlack() *Comp {
	c.El.Style(styles.AccentBlack())
	return c
}

func (c *Comp) AccentBlackAlpha(scale int) *Comp {
	c.El.Style(styles.AccentBlackAlpha(scale))
	return c
}

func (c *Comp) AccentBlue(scale int) *Comp {
	c.El.Style(styles.AccentBlue(scale))
	return c
}

func (c *Comp) AccentBlueAlpha(scale int) *Comp {
	c.El.Style(styles.AccentBlueAlpha(scale))
	return c
}

func (c *Comp) AccentBlueDark(scale int) *Comp {
	c.El.Style(styles.AccentBlueDark(scale))
	return c
}

func (c *Comp) AccentBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentBlueDarkAlpha(scale))
	return c
}

func (c *Comp) AccentBronze(scale int) *Comp {
	c.El.Style(styles.AccentBronze(scale))
	return c
}

func (c *Comp) AccentBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.AccentBronzeAlpha(scale))
	return c
}

func (c *Comp) AccentBronzeDark(scale int) *Comp {
	c.El.Style(styles.AccentBronzeDark(scale))
	return c
}

func (c *Comp) AccentBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) AccentBrown(scale int) *Comp {
	c.El.Style(styles.AccentBrown(scale))
	return c
}

func (c *Comp) AccentBrownAlpha(scale int) *Comp {
	c.El.Style(styles.AccentBrownAlpha(scale))
	return c
}

func (c *Comp) AccentBrownDark(scale int) *Comp {
	c.El.Style(styles.AccentBrownDark(scale))
	return c
}

func (c *Comp) AccentBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentBrownDarkAlpha(scale))
	return c
}

func (c *Comp) AccentCrimson(scale int) *Comp {
	c.El.Style(styles.AccentCrimson(scale))
	return c
}

func (c *Comp) AccentCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.AccentCrimsonAlpha(scale))
	return c
}

func (c *Comp) AccentCrimsonDark(scale int) *Comp {
	c.El.Style(styles.AccentCrimsonDark(scale))
	return c
}

func (c *Comp) AccentCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) AccentCurrent() *Comp {
	c.El.Style(styles.AccentCurrent())
	return c
}

func (c *Comp) AccentCyan(scale int) *Comp {
	c.El.Style(styles.AccentCyan(scale))
	return c
}

func (c *Comp) AccentCyanAlpha(scale int) *Comp {
	c.El.Style(styles.AccentCyanAlpha(scale))
	return c
}

func (c *Comp) AccentCyanDark(scale int) *Comp {
	c.El.Style(styles.AccentCyanDark(scale))
	return c
}

func (c *Comp) AccentCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentCyanDarkAlpha(scale))
	return c
}

func (c *Comp) AccentGold(scale int) *Comp {
	c.El.Style(styles.AccentGold(scale))
	return c
}

func (c *Comp) AccentGoldAlpha(scale int) *Comp {
	c.El.Style(styles.AccentGoldAlpha(scale))
	return c
}

func (c *Comp) AccentGoldDark(scale int) *Comp {
	c.El.Style(styles.AccentGoldDark(scale))
	return c
}

func (c *Comp) AccentGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentGoldDarkAlpha(scale))
	return c
}

func (c *Comp) AccentGrass(scale int) *Comp {
	c.El.Style(styles.AccentGrass(scale))
	return c
}

func (c *Comp) AccentGrassAlpha(scale int) *Comp {
	c.El.Style(styles.AccentGrassAlpha(scale))
	return c
}

func (c *Comp) AccentGrassDark(scale int) *Comp {
	c.El.Style(styles.AccentGrassDark(scale))
	return c
}

func (c *Comp) AccentGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentGrassDarkAlpha(scale))
	return c
}

func (c *Comp) AccentGray(scale int) *Comp {
	c.El.Style(styles.AccentGray(scale))
	return c
}

func (c *Comp) AccentGrayAlpha(scale int) *Comp {
	c.El.Style(styles.AccentGrayAlpha(scale))
	return c
}

func (c *Comp) AccentGrayDark(scale int) *Comp {
	c.El.Style(styles.AccentGrayDark(scale))
	return c
}

func (c *Comp) AccentGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentGrayDarkAlpha(scale))
	return c
}

func (c *Comp) AccentGreen(scale int) *Comp {
	c.El.Style(styles.AccentGreen(scale))
	return c
}

func (c *Comp) AccentGreenAlpha(scale int) *Comp {
	c.El.Style(styles.AccentGreenAlpha(scale))
	return c
}

func (c *Comp) AccentGreenDark(scale int) *Comp {
	c.El.Style(styles.AccentGreenDark(scale))
	return c
}

func (c *Comp) AccentGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentGreenDarkAlpha(scale))
	return c
}

func (c *Comp) AccentIndigo(scale int) *Comp {
	c.El.Style(styles.AccentIndigo(scale))
	return c
}

func (c *Comp) AccentIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.AccentIndigoAlpha(scale))
	return c
}

func (c *Comp) AccentIndigoDark(scale int) *Comp {
	c.El.Style(styles.AccentIndigoDark(scale))
	return c
}

func (c *Comp) AccentIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) AccentInherit() *Comp {
	c.El.Style(styles.AccentInherit())
	return c
}

func (c *Comp) AccentIris(scale int) *Comp {
	c.El.Style(styles.AccentIris(scale))
	return c
}

func (c *Comp) AccentIrisAlpha(scale int) *Comp {
	c.El.Style(styles.AccentIrisAlpha(scale))
	return c
}

func (c *Comp) AccentIrisDark(scale int) *Comp {
	c.El.Style(styles.AccentIrisDark(scale))
	return c
}

func (c *Comp) AccentIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentIrisDarkAlpha(scale))
	return c
}

func (c *Comp) AccentJade(scale int) *Comp {
	c.El.Style(styles.AccentJade(scale))
	return c
}

func (c *Comp) AccentJadeAlpha(scale int) *Comp {
	c.El.Style(styles.AccentJadeAlpha(scale))
	return c
}

func (c *Comp) AccentJadeDark(scale int) *Comp {
	c.El.Style(styles.AccentJadeDark(scale))
	return c
}

func (c *Comp) AccentJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentJadeDarkAlpha(scale))
	return c
}

func (c *Comp) AccentLime(scale int) *Comp {
	c.El.Style(styles.AccentLime(scale))
	return c
}

func (c *Comp) AccentLimeAlpha(scale int) *Comp {
	c.El.Style(styles.AccentLimeAlpha(scale))
	return c
}

func (c *Comp) AccentLimeDark(scale int) *Comp {
	c.El.Style(styles.AccentLimeDark(scale))
	return c
}

func (c *Comp) AccentLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentLimeDarkAlpha(scale))
	return c
}

func (c *Comp) AccentMauve(scale int) *Comp {
	c.El.Style(styles.AccentMauve(scale))
	return c
}

func (c *Comp) AccentMauveAlpha(scale int) *Comp {
	c.El.Style(styles.AccentMauveAlpha(scale))
	return c
}

func (c *Comp) AccentMauveDark(scale int) *Comp {
	c.El.Style(styles.AccentMauveDark(scale))
	return c
}

func (c *Comp) AccentMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentMauveDarkAlpha(scale))
	return c
}

func (c *Comp) AccentMint(scale int) *Comp {
	c.El.Style(styles.AccentMint(scale))
	return c
}

func (c *Comp) AccentMintAlpha(scale int) *Comp {
	c.El.Style(styles.AccentMintAlpha(scale))
	return c
}

func (c *Comp) AccentMintDark(scale int) *Comp {
	c.El.Style(styles.AccentMintDark(scale))
	return c
}

func (c *Comp) AccentMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentMintDarkAlpha(scale))
	return c
}

func (c *Comp) AccentOlive(scale int) *Comp {
	c.El.Style(styles.AccentOlive(scale))
	return c
}

func (c *Comp) AccentOliveAlpha(scale int) *Comp {
	c.El.Style(styles.AccentOliveAlpha(scale))
	return c
}

func (c *Comp) AccentOliveDark(scale int) *Comp {
	c.El.Style(styles.AccentOliveDark(scale))
	return c
}

func (c *Comp) AccentOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentOliveDarkAlpha(scale))
	return c
}

func (c *Comp) AccentOrange(scale int) *Comp {
	c.El.Style(styles.AccentOrange(scale))
	return c
}

func (c *Comp) AccentOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.AccentOrangeAlpha(scale))
	return c
}

func (c *Comp) AccentOrangeDark(scale int) *Comp {
	c.El.Style(styles.AccentOrangeDark(scale))
	return c
}

func (c *Comp) AccentOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) AccentPink(scale int) *Comp {
	c.El.Style(styles.AccentPink(scale))
	return c
}

func (c *Comp) AccentPinkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentPinkAlpha(scale))
	return c
}

func (c *Comp) AccentPinkDark(scale int) *Comp {
	c.El.Style(styles.AccentPinkDark(scale))
	return c
}

func (c *Comp) AccentPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentPinkDarkAlpha(scale))
	return c
}

func (c *Comp) AccentPlum(scale int) *Comp {
	c.El.Style(styles.AccentPlum(scale))
	return c
}

func (c *Comp) AccentPlumAlpha(scale int) *Comp {
	c.El.Style(styles.AccentPlumAlpha(scale))
	return c
}

func (c *Comp) AccentPlumDark(scale int) *Comp {
	c.El.Style(styles.AccentPlumDark(scale))
	return c
}

func (c *Comp) AccentPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentPlumDarkAlpha(scale))
	return c
}

func (c *Comp) AccentPurple(scale int) *Comp {
	c.El.Style(styles.AccentPurple(scale))
	return c
}

func (c *Comp) AccentPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.AccentPurpleAlpha(scale))
	return c
}

func (c *Comp) AccentPurpleDark(scale int) *Comp {
	c.El.Style(styles.AccentPurpleDark(scale))
	return c
}

func (c *Comp) AccentPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) AccentRed(scale int) *Comp {
	c.El.Style(styles.AccentRed(scale))
	return c
}

func (c *Comp) AccentRedAlpha(scale int) *Comp {
	c.El.Style(styles.AccentRedAlpha(scale))
	return c
}

func (c *Comp) AccentRedDark(scale int) *Comp {
	c.El.Style(styles.AccentRedDark(scale))
	return c
}

func (c *Comp) AccentRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentRedDarkAlpha(scale))
	return c
}

func (c *Comp) AccentRuby(scale int) *Comp {
	c.El.Style(styles.AccentRuby(scale))
	return c
}

func (c *Comp) AccentRubyAlpha(scale int) *Comp {
	c.El.Style(styles.AccentRubyAlpha(scale))
	return c
}

func (c *Comp) AccentRubyDark(scale int) *Comp {
	c.El.Style(styles.AccentRubyDark(scale))
	return c
}

func (c *Comp) AccentRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentRubyDarkAlpha(scale))
	return c
}

func (c *Comp) AccentSage(scale int) *Comp {
	c.El.Style(styles.AccentSage(scale))
	return c
}

func (c *Comp) AccentSageAlpha(scale int) *Comp {
	c.El.Style(styles.AccentSageAlpha(scale))
	return c
}

func (c *Comp) AccentSageDark(scale int) *Comp {
	c.El.Style(styles.AccentSageDark(scale))
	return c
}

func (c *Comp) AccentSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentSageDarkAlpha(scale))
	return c
}

func (c *Comp) AccentSand(scale int) *Comp {
	c.El.Style(styles.AccentSand(scale))
	return c
}

func (c *Comp) AccentSandAlpha(scale int) *Comp {
	c.El.Style(styles.AccentSandAlpha(scale))
	return c
}

func (c *Comp) AccentSandDark(scale int) *Comp {
	c.El.Style(styles.AccentSandDark(scale))
	return c
}

func (c *Comp) AccentSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentSandDarkAlpha(scale))
	return c
}

func (c *Comp) AccentSky(scale int) *Comp {
	c.El.Style(styles.AccentSky(scale))
	return c
}

func (c *Comp) AccentSkyAlpha(scale int) *Comp {
	c.El.Style(styles.AccentSkyAlpha(scale))
	return c
}

func (c *Comp) AccentSkyDark(scale int) *Comp {
	c.El.Style(styles.AccentSkyDark(scale))
	return c
}

func (c *Comp) AccentSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentSkyDarkAlpha(scale))
	return c
}

func (c *Comp) AccentSlate(scale int) *Comp {
	c.El.Style(styles.AccentSlate(scale))
	return c
}

func (c *Comp) AccentSlateAlpha(scale int) *Comp {
	c.El.Style(styles.AccentSlateAlpha(scale))
	return c
}

func (c *Comp) AccentSlateDark(scale int) *Comp {
	c.El.Style(styles.AccentSlateDark(scale))
	return c
}

func (c *Comp) AccentSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentSlateDarkAlpha(scale))
	return c
}

func (c *Comp) AccentTeal(scale int) *Comp {
	c.El.Style(styles.AccentTeal(scale))
	return c
}

func (c *Comp) AccentTealAlpha(scale int) *Comp {
	c.El.Style(styles.AccentTealAlpha(scale))
	return c
}

func (c *Comp) AccentTealDark(scale int) *Comp {
	c.El.Style(styles.AccentTealDark(scale))
	return c
}

func (c *Comp) AccentTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentTealDarkAlpha(scale))
	return c
}

func (c *Comp) AccentTomato(scale int) *Comp {
	c.El.Style(styles.AccentTomato(scale))
	return c
}

func (c *Comp) AccentTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.AccentTomatoAlpha(scale))
	return c
}

func (c *Comp) AccentTomatoDark(scale int) *Comp {
	c.El.Style(styles.AccentTomatoDark(scale))
	return c
}

func (c *Comp) AccentTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) AccentTransparent() *Comp {
	c.El.Style(styles.AccentTransparent())
	return c
}

func (c *Comp) AccentViolet(scale int) *Comp {
	c.El.Style(styles.AccentViolet(scale))
	return c
}

func (c *Comp) AccentVioletAlpha(scale int) *Comp {
	c.El.Style(styles.AccentVioletAlpha(scale))
	return c
}

func (c *Comp) AccentVioletDark(scale int) *Comp {
	c.El.Style(styles.AccentVioletDark(scale))
	return c
}

func (c *Comp) AccentVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentVioletDarkAlpha(scale))
	return c
}

func (c *Comp) AccentWhite() *Comp {
	c.El.Style(styles.AccentWhite())
	return c
}

func (c *Comp) AccentWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.AccentWhiteAlpha(scale))
	return c
}

func (c *Comp) AccentYellow(scale int) *Comp {
	c.El.Style(styles.AccentYellow(scale))
	return c
}

func (c *Comp) AccentYellowAlpha(scale int) *Comp {
	c.El.Style(styles.AccentYellowAlpha(scale))
	return c
}

func (c *Comp) AccentYellowDark(scale int) *Comp {
	c.El.Style(styles.AccentYellowDark(scale))
	return c
}

func (c *Comp) AccentYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.AccentYellowDarkAlpha(scale))
	return c
}

func (c *Comp) AlignBaseline() *Comp {
	c.El.Style(styles.AlignBaseline())
	return c
}

func (c *Comp) AlignBottom() *Comp {
	c.El.Style(styles.AlignBottom())
	return c
}

func (c *Comp) AlignBy(val value.Value) *Comp {
	c.El.Style(styles.AlignBy(val))
	return c
}

func (c *Comp) AlignMiddle() *Comp {
	c.El.Style(styles.AlignMiddle())
	return c
}

func (c *Comp) AlignSub() *Comp {
	c.El.Style(styles.AlignSub())
	return c
}

func (c *Comp) AlignSuper() *Comp {
	c.El.Style(styles.AlignSuper())
	return c
}

func (c *Comp) AlignTextBottom() *Comp {
	c.El.Style(styles.AlignTextBottom())
	return c
}

func (c *Comp) AlignTextTop() *Comp {
	c.El.Style(styles.AlignTextTop())
	return c
}

func (c *Comp) AlignTop() *Comp {
	c.El.Style(styles.AlignTop())
	return c
}

func (c *Comp) Animate(val value.Value) *Comp {
	c.El.Style(styles.Animate(val))
	return c
}

func (c *Comp) AnimateBounce() *Comp {
	c.El.Style(styles.AnimateBounce())
	return c
}

func (c *Comp) AnimateNone() *Comp {
	c.El.Style(styles.AnimateNone())
	return c
}

func (c *Comp) AnimatePing() *Comp {
	c.El.Style(styles.AnimatePing())
	return c
}

func (c *Comp) AnimatePulse() *Comp {
	c.El.Style(styles.AnimatePulse())
	return c
}

func (c *Comp) AnimateSpin() *Comp {
	c.El.Style(styles.AnimateSpin())
	return c
}

func (c *Comp) Antialiazed() *Comp {
	c.El.Style(styles.Antialiazed())
	return c
}

func (c *Comp) AppearanceAuto() *Comp {
	c.El.Style(styles.AppearanceAuto())
	return c
}

func (c *Comp) AppearanceNone() *Comp {
	c.El.Style(styles.AppearanceNone())
	return c
}

func (c *Comp) Aspect(ratio value.Value) *Comp {
	c.El.Style(styles.Aspect(ratio))
	return c
}

func (c *Comp) AspectAuto() *Comp {
	c.El.Style(styles.AspectAuto())
	return c
}

func (c *Comp) AspectSquare() *Comp {
	c.El.Style(styles.AspectSquare())
	return c
}

func (c *Comp) AspectVideo() *Comp {
	c.El.Style(styles.AspectVideo())
	return c
}

func (c *Comp) AutoColsAuto() *Comp {
	c.El.Style(styles.AutoColsAuto())
	return c
}

func (c *Comp) AutoColsBy(val value.Value) *Comp {
	c.El.Style(styles.AutoColsBy(val))
	return c
}

func (c *Comp) AutoColsFr() *Comp {
	c.El.Style(styles.AutoColsFr())
	return c
}

func (c *Comp) AutoColsMax() *Comp {
	c.El.Style(styles.AutoColsMax())
	return c
}

func (c *Comp) AutoColsMin() *Comp {
	c.El.Style(styles.AutoColsMin())
	return c
}

func (c *Comp) AutoRowsAuto() *Comp {
	c.El.Style(styles.AutoRowsAuto())
	return c
}

func (c *Comp) AutoRowsBy(val value.Value) *Comp {
	c.El.Style(styles.AutoRowsBy(val))
	return c
}

func (c *Comp) AutoRowsFr() *Comp {
	c.El.Style(styles.AutoRowsFr())
	return c
}

func (c *Comp) AutoRowsMax() *Comp {
	c.El.Style(styles.AutoRowsMax())
	return c
}

func (c *Comp) AutoRowsMin() *Comp {
	c.El.Style(styles.AutoRowsMin())
	return c
}

func (c *Comp) BackdropBlur(val value.Value) *Comp {
	c.El.Style(styles.BackdropBlur(val))
	return c
}

func (c *Comp) BackdropBlur2xl() *Comp {
	c.El.Style(styles.BackdropBlur2xl())
	return c
}

func (c *Comp) BackdropBlur3xl() *Comp {
	c.El.Style(styles.BackdropBlur3xl())
	return c
}

func (c *Comp) BackdropBlurLg() *Comp {
	c.El.Style(styles.BackdropBlurLg())
	return c
}

func (c *Comp) BackdropBlurMd() *Comp {
	c.El.Style(styles.BackdropBlurMd())
	return c
}

func (c *Comp) BackdropBlurNone() *Comp {
	c.El.Style(styles.BackdropBlurNone())
	return c
}

func (c *Comp) BackdropBlurSm() *Comp {
	c.El.Style(styles.BackdropBlurSm())
	return c
}

func (c *Comp) BackdropBlurXl() *Comp {
	c.El.Style(styles.BackdropBlurXl())
	return c
}

func (c *Comp) BackdropBlurXs() *Comp {
	c.El.Style(styles.BackdropBlurXs())
	return c
}

func (c *Comp) BackdropBrightness(val any) *Comp {
	c.El.Style(styles.BackdropBrightness(val))
	return c
}

func (c *Comp) BackdropContrast(val any) *Comp {
	c.El.Style(styles.BackdropContrast(val))
	return c
}

func (c *Comp) BackdropFilter(val value.Value) *Comp {
	c.El.Style(styles.BackdropFilter(val))
	return c
}

func (c *Comp) BackdropFilterNone() *Comp {
	c.El.Style(styles.BackdropFilterNone())
	return c
}

func (c *Comp) BackdropGrayscale(val ...any) *Comp {
	c.El.Style(styles.BackdropGrayscale(val...))
	return c
}

func (c *Comp) BackdropHueRotate(val any) *Comp {
	c.El.Style(styles.BackdropHueRotate(val))
	return c
}

func (c *Comp) BackdropInvert(val ...any) *Comp {
	c.El.Style(styles.BackdropInvert(val...))
	return c
}

func (c *Comp) BackdropSaturate(val any) *Comp {
	c.El.Style(styles.BackdropSaturate(val))
	return c
}

func (c *Comp) BackdropSepia(val ...any) *Comp {
	c.El.Style(styles.BackdropSepia(val...))
	return c
}

func (c *Comp) BackfaceHidden() *Comp {
	c.El.Style(styles.BackfaceHidden())
	return c
}

func (c *Comp) BackfaceVisible() *Comp {
	c.El.Style(styles.BackfaceVisible())
	return c
}

func (c *Comp) Basis(val value.Value) *Comp {
	c.El.Style(styles.Basis(val))
	return c
}

func (c *Comp) Basis2xl() *Comp {
	c.El.Style(styles.Basis2xl())
	return c
}

func (c *Comp) Basis2xs() *Comp {
	c.El.Style(styles.Basis2xs())
	return c
}

func (c *Comp) Basis3xl() *Comp {
	c.El.Style(styles.Basis3xl())
	return c
}

func (c *Comp) Basis3xs() *Comp {
	c.El.Style(styles.Basis3xs())
	return c
}

func (c *Comp) Basis4xl() *Comp {
	c.El.Style(styles.Basis4xl())
	return c
}

func (c *Comp) Basis5xl() *Comp {
	c.El.Style(styles.Basis5xl())
	return c
}

func (c *Comp) Basis6xl() *Comp {
	c.El.Style(styles.Basis6xl())
	return c
}

func (c *Comp) Basis7xl() *Comp {
	c.El.Style(styles.Basis7xl())
	return c
}

func (c *Comp) BasisAuto() *Comp {
	c.El.Style(styles.BasisAuto())
	return c
}

func (c *Comp) BasisFraction(fraction float64) *Comp {
	c.El.Style(styles.BasisFraction(fraction))
	return c
}

func (c *Comp) BasisFull() *Comp {
	c.El.Style(styles.BasisFull())
	return c
}

func (c *Comp) BasisLg() *Comp {
	c.El.Style(styles.BasisLg())
	return c
}

func (c *Comp) BasisMd() *Comp {
	c.El.Style(styles.BasisMd())
	return c
}

func (c *Comp) BasisSm() *Comp {
	c.El.Style(styles.BasisSm())
	return c
}

func (c *Comp) BasisXl() *Comp {
	c.El.Style(styles.BasisXl())
	return c
}

func (c *Comp) BasisXs() *Comp {
	c.El.Style(styles.BasisXs())
	return c
}

func (c *Comp) Bg(val value.Value) *Comp {
	c.El.Style(styles.Bg(val))
	return c
}

func (c *Comp) BgAmber(scale int) *Comp {
	c.El.Style(styles.BgAmber(scale))
	return c
}

func (c *Comp) BgAmberAlpha(scale int) *Comp {
	c.El.Style(styles.BgAmberAlpha(scale))
	return c
}

func (c *Comp) BgAmberDark(scale int) *Comp {
	c.El.Style(styles.BgAmberDark(scale))
	return c
}

func (c *Comp) BgAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgAmberDarkAlpha(scale))
	return c
}

func (c *Comp) BgAuto() *Comp {
	c.El.Style(styles.BgAuto())
	return c
}

func (c *Comp) BgBlack() *Comp {
	c.El.Style(styles.BgBlack())
	return c
}

func (c *Comp) BgBlackAlpha(scale int) *Comp {
	c.El.Style(styles.BgBlackAlpha(scale))
	return c
}

func (c *Comp) BgBlendColor() *Comp {
	c.El.Style(styles.BgBlendColor())
	return c
}

func (c *Comp) BgBlendColorBurn() *Comp {
	c.El.Style(styles.BgBlendColorBurn())
	return c
}

func (c *Comp) BgBlendColorDodge() *Comp {
	c.El.Style(styles.BgBlendColorDodge())
	return c
}

func (c *Comp) BgBlendDarken() *Comp {
	c.El.Style(styles.BgBlendDarken())
	return c
}

func (c *Comp) BgBlendDifference() *Comp {
	c.El.Style(styles.BgBlendDifference())
	return c
}

func (c *Comp) BgBlendExclusion() *Comp {
	c.El.Style(styles.BgBlendExclusion())
	return c
}

func (c *Comp) BgBlendHardLight() *Comp {
	c.El.Style(styles.BgBlendHardLight())
	return c
}

func (c *Comp) BgBlendHue() *Comp {
	c.El.Style(styles.BgBlendHue())
	return c
}

func (c *Comp) BgBlendLighten() *Comp {
	c.El.Style(styles.BgBlendLighten())
	return c
}

func (c *Comp) BgBlendLuminosity() *Comp {
	c.El.Style(styles.BgBlendLuminosity())
	return c
}

func (c *Comp) BgBlendMultiply() *Comp {
	c.El.Style(styles.BgBlendMultiply())
	return c
}

func (c *Comp) BgBlendNormal() *Comp {
	c.El.Style(styles.BgBlendNormal())
	return c
}

func (c *Comp) BgBlendOverlay() *Comp {
	c.El.Style(styles.BgBlendOverlay())
	return c
}

func (c *Comp) BgBlendSaturation() *Comp {
	c.El.Style(styles.BgBlendSaturation())
	return c
}

func (c *Comp) BgBlendScreen() *Comp {
	c.El.Style(styles.BgBlendScreen())
	return c
}

func (c *Comp) BgBlendSoftLight() *Comp {
	c.El.Style(styles.BgBlendSoftLight())
	return c
}

func (c *Comp) BgBlue(scale int) *Comp {
	c.El.Style(styles.BgBlue(scale))
	return c
}

func (c *Comp) BgBlueAlpha(scale int) *Comp {
	c.El.Style(styles.BgBlueAlpha(scale))
	return c
}

func (c *Comp) BgBlueDark(scale int) *Comp {
	c.El.Style(styles.BgBlueDark(scale))
	return c
}

func (c *Comp) BgBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgBlueDarkAlpha(scale))
	return c
}

func (c *Comp) BgBottom() *Comp {
	c.El.Style(styles.BgBottom())
	return c
}

func (c *Comp) BgBottomLeft() *Comp {
	c.El.Style(styles.BgBottomLeft())
	return c
}

func (c *Comp) BgBottomRight() *Comp {
	c.El.Style(styles.BgBottomRight())
	return c
}

func (c *Comp) BgBronze(scale int) *Comp {
	c.El.Style(styles.BgBronze(scale))
	return c
}

func (c *Comp) BgBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.BgBronzeAlpha(scale))
	return c
}

func (c *Comp) BgBronzeDark(scale int) *Comp {
	c.El.Style(styles.BgBronzeDark(scale))
	return c
}

func (c *Comp) BgBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) BgBrown(scale int) *Comp {
	c.El.Style(styles.BgBrown(scale))
	return c
}

func (c *Comp) BgBrownAlpha(scale int) *Comp {
	c.El.Style(styles.BgBrownAlpha(scale))
	return c
}

func (c *Comp) BgBrownDark(scale int) *Comp {
	c.El.Style(styles.BgBrownDark(scale))
	return c
}

func (c *Comp) BgBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgBrownDarkAlpha(scale))
	return c
}

func (c *Comp) BgCenter() *Comp {
	c.El.Style(styles.BgCenter())
	return c
}

func (c *Comp) BgClipBorder() *Comp {
	c.El.Style(styles.BgClipBorder())
	return c
}

func (c *Comp) BgClipContent() *Comp {
	c.El.Style(styles.BgClipContent())
	return c
}

func (c *Comp) BgClipPadding() *Comp {
	c.El.Style(styles.BgClipPadding())
	return c
}

func (c *Comp) BgClipText() *Comp {
	c.El.Style(styles.BgClipText())
	return c
}

func (c *Comp) BgColor(color value.Value) *Comp {
	c.El.Style(styles.BgColor(color))
	return c
}

func (c *Comp) BgConic(val value.Value) *Comp {
	c.El.Style(styles.BgConic(val))
	return c
}

func (c *Comp) BgContain() *Comp {
	c.El.Style(styles.BgContain())
	return c
}

func (c *Comp) BgCover() *Comp {
	c.El.Style(styles.BgCover())
	return c
}

func (c *Comp) BgCrimson(scale int) *Comp {
	c.El.Style(styles.BgCrimson(scale))
	return c
}

func (c *Comp) BgCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.BgCrimsonAlpha(scale))
	return c
}

func (c *Comp) BgCrimsonDark(scale int) *Comp {
	c.El.Style(styles.BgCrimsonDark(scale))
	return c
}

func (c *Comp) BgCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) BgCurrent() *Comp {
	c.El.Style(styles.BgCurrent())
	return c
}

func (c *Comp) BgCyan(scale int) *Comp {
	c.El.Style(styles.BgCyan(scale))
	return c
}

func (c *Comp) BgCyanAlpha(scale int) *Comp {
	c.El.Style(styles.BgCyanAlpha(scale))
	return c
}

func (c *Comp) BgCyanDark(scale int) *Comp {
	c.El.Style(styles.BgCyanDark(scale))
	return c
}

func (c *Comp) BgCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgCyanDarkAlpha(scale))
	return c
}

func (c *Comp) BgFixed() *Comp {
	c.El.Style(styles.BgFixed())
	return c
}

func (c *Comp) BgGold(scale int) *Comp {
	c.El.Style(styles.BgGold(scale))
	return c
}

func (c *Comp) BgGoldAlpha(scale int) *Comp {
	c.El.Style(styles.BgGoldAlpha(scale))
	return c
}

func (c *Comp) BgGoldDark(scale int) *Comp {
	c.El.Style(styles.BgGoldDark(scale))
	return c
}

func (c *Comp) BgGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgGoldDarkAlpha(scale))
	return c
}

func (c *Comp) BgGrass(scale int) *Comp {
	c.El.Style(styles.BgGrass(scale))
	return c
}

func (c *Comp) BgGrassAlpha(scale int) *Comp {
	c.El.Style(styles.BgGrassAlpha(scale))
	return c
}

func (c *Comp) BgGrassDark(scale int) *Comp {
	c.El.Style(styles.BgGrassDark(scale))
	return c
}

func (c *Comp) BgGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgGrassDarkAlpha(scale))
	return c
}

func (c *Comp) BgGray(scale int) *Comp {
	c.El.Style(styles.BgGray(scale))
	return c
}

func (c *Comp) BgGrayAlpha(scale int) *Comp {
	c.El.Style(styles.BgGrayAlpha(scale))
	return c
}

func (c *Comp) BgGrayDark(scale int) *Comp {
	c.El.Style(styles.BgGrayDark(scale))
	return c
}

func (c *Comp) BgGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgGrayDarkAlpha(scale))
	return c
}

func (c *Comp) BgGreen(scale int) *Comp {
	c.El.Style(styles.BgGreen(scale))
	return c
}

func (c *Comp) BgGreenAlpha(scale int) *Comp {
	c.El.Style(styles.BgGreenAlpha(scale))
	return c
}

func (c *Comp) BgGreenDark(scale int) *Comp {
	c.El.Style(styles.BgGreenDark(scale))
	return c
}

func (c *Comp) BgGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgGreenDarkAlpha(scale))
	return c
}

func (c *Comp) BgIndigo(scale int) *Comp {
	c.El.Style(styles.BgIndigo(scale))
	return c
}

func (c *Comp) BgIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.BgIndigoAlpha(scale))
	return c
}

func (c *Comp) BgIndigoDark(scale int) *Comp {
	c.El.Style(styles.BgIndigoDark(scale))
	return c
}

func (c *Comp) BgIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) BgInherit() *Comp {
	c.El.Style(styles.BgInherit())
	return c
}

func (c *Comp) BgIris(scale int) *Comp {
	c.El.Style(styles.BgIris(scale))
	return c
}

func (c *Comp) BgIrisAlpha(scale int) *Comp {
	c.El.Style(styles.BgIrisAlpha(scale))
	return c
}

func (c *Comp) BgIrisDark(scale int) *Comp {
	c.El.Style(styles.BgIrisDark(scale))
	return c
}

func (c *Comp) BgIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgIrisDarkAlpha(scale))
	return c
}

func (c *Comp) BgJade(scale int) *Comp {
	c.El.Style(styles.BgJade(scale))
	return c
}

func (c *Comp) BgJadeAlpha(scale int) *Comp {
	c.El.Style(styles.BgJadeAlpha(scale))
	return c
}

func (c *Comp) BgJadeDark(scale int) *Comp {
	c.El.Style(styles.BgJadeDark(scale))
	return c
}

func (c *Comp) BgJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgJadeDarkAlpha(scale))
	return c
}

func (c *Comp) BgLeft() *Comp {
	c.El.Style(styles.BgLeft())
	return c
}

func (c *Comp) BgLime(scale int) *Comp {
	c.El.Style(styles.BgLime(scale))
	return c
}

func (c *Comp) BgLimeAlpha(scale int) *Comp {
	c.El.Style(styles.BgLimeAlpha(scale))
	return c
}

func (c *Comp) BgLimeDark(scale int) *Comp {
	c.El.Style(styles.BgLimeDark(scale))
	return c
}

func (c *Comp) BgLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgLimeDarkAlpha(scale))
	return c
}

func (c *Comp) BgLinear(val value.Value) *Comp {
	c.El.Style(styles.BgLinear(val))
	return c
}

func (c *Comp) BgLinearToB() *Comp {
	c.El.Style(styles.BgLinearToB())
	return c
}

func (c *Comp) BgLinearToBl() *Comp {
	c.El.Style(styles.BgLinearToBl())
	return c
}

func (c *Comp) BgLinearToBr() *Comp {
	c.El.Style(styles.BgLinearToBr())
	return c
}

func (c *Comp) BgLinearToL() *Comp {
	c.El.Style(styles.BgLinearToL())
	return c
}

func (c *Comp) BgLinearToR() *Comp {
	c.El.Style(styles.BgLinearToR())
	return c
}

func (c *Comp) BgLinearToT() *Comp {
	c.El.Style(styles.BgLinearToT())
	return c
}

func (c *Comp) BgLinearToTl() *Comp {
	c.El.Style(styles.BgLinearToTl())
	return c
}

func (c *Comp) BgLinearToTr() *Comp {
	c.El.Style(styles.BgLinearToTr())
	return c
}

func (c *Comp) BgLocal() *Comp {
	c.El.Style(styles.BgLocal())
	return c
}

func (c *Comp) BgMauve(scale int) *Comp {
	c.El.Style(styles.BgMauve(scale))
	return c
}

func (c *Comp) BgMauveAlpha(scale int) *Comp {
	c.El.Style(styles.BgMauveAlpha(scale))
	return c
}

func (c *Comp) BgMauveDark(scale int) *Comp {
	c.El.Style(styles.BgMauveDark(scale))
	return c
}

func (c *Comp) BgMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgMauveDarkAlpha(scale))
	return c
}

func (c *Comp) BgMint(scale int) *Comp {
	c.El.Style(styles.BgMint(scale))
	return c
}

func (c *Comp) BgMintAlpha(scale int) *Comp {
	c.El.Style(styles.BgMintAlpha(scale))
	return c
}

func (c *Comp) BgMintDark(scale int) *Comp {
	c.El.Style(styles.BgMintDark(scale))
	return c
}

func (c *Comp) BgMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgMintDarkAlpha(scale))
	return c
}

func (c *Comp) BgNoRepeat() *Comp {
	c.El.Style(styles.BgNoRepeat())
	return c
}

func (c *Comp) BgNone() *Comp {
	c.El.Style(styles.BgNone())
	return c
}

func (c *Comp) BgOlive(scale int) *Comp {
	c.El.Style(styles.BgOlive(scale))
	return c
}

func (c *Comp) BgOliveAlpha(scale int) *Comp {
	c.El.Style(styles.BgOliveAlpha(scale))
	return c
}

func (c *Comp) BgOliveDark(scale int) *Comp {
	c.El.Style(styles.BgOliveDark(scale))
	return c
}

func (c *Comp) BgOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgOliveDarkAlpha(scale))
	return c
}

func (c *Comp) BgOrange(scale int) *Comp {
	c.El.Style(styles.BgOrange(scale))
	return c
}

func (c *Comp) BgOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.BgOrangeAlpha(scale))
	return c
}

func (c *Comp) BgOrangeDark(scale int) *Comp {
	c.El.Style(styles.BgOrangeDark(scale))
	return c
}

func (c *Comp) BgOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) BgOriginBorder() *Comp {
	c.El.Style(styles.BgOriginBorder())
	return c
}

func (c *Comp) BgOriginContent() *Comp {
	c.El.Style(styles.BgOriginContent())
	return c
}

func (c *Comp) BgOriginPadding() *Comp {
	c.El.Style(styles.BgOriginPadding())
	return c
}

func (c *Comp) BgPink(scale int) *Comp {
	c.El.Style(styles.BgPink(scale))
	return c
}

func (c *Comp) BgPinkAlpha(scale int) *Comp {
	c.El.Style(styles.BgPinkAlpha(scale))
	return c
}

func (c *Comp) BgPinkDark(scale int) *Comp {
	c.El.Style(styles.BgPinkDark(scale))
	return c
}

func (c *Comp) BgPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgPinkDarkAlpha(scale))
	return c
}

func (c *Comp) BgPlum(scale int) *Comp {
	c.El.Style(styles.BgPlum(scale))
	return c
}

func (c *Comp) BgPlumAlpha(scale int) *Comp {
	c.El.Style(styles.BgPlumAlpha(scale))
	return c
}

func (c *Comp) BgPlumDark(scale int) *Comp {
	c.El.Style(styles.BgPlumDark(scale))
	return c
}

func (c *Comp) BgPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgPlumDarkAlpha(scale))
	return c
}

func (c *Comp) BgPosition(val value.Value) *Comp {
	c.El.Style(styles.BgPosition(val))
	return c
}

func (c *Comp) BgPurple(scale int) *Comp {
	c.El.Style(styles.BgPurple(scale))
	return c
}

func (c *Comp) BgPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.BgPurpleAlpha(scale))
	return c
}

func (c *Comp) BgPurpleDark(scale int) *Comp {
	c.El.Style(styles.BgPurpleDark(scale))
	return c
}

func (c *Comp) BgPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) BgRadial(val value.Value) *Comp {
	c.El.Style(styles.BgRadial(val))
	return c
}

func (c *Comp) BgRed(scale int) *Comp {
	c.El.Style(styles.BgRed(scale))
	return c
}

func (c *Comp) BgRedAlpha(scale int) *Comp {
	c.El.Style(styles.BgRedAlpha(scale))
	return c
}

func (c *Comp) BgRedDark(scale int) *Comp {
	c.El.Style(styles.BgRedDark(scale))
	return c
}

func (c *Comp) BgRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgRedDarkAlpha(scale))
	return c
}

func (c *Comp) BgRepeat() *Comp {
	c.El.Style(styles.BgRepeat())
	return c
}

func (c *Comp) BgRepeatRound() *Comp {
	c.El.Style(styles.BgRepeatRound())
	return c
}

func (c *Comp) BgRepeatSpace() *Comp {
	c.El.Style(styles.BgRepeatSpace())
	return c
}

func (c *Comp) BgRepeatX() *Comp {
	c.El.Style(styles.BgRepeatX())
	return c
}

func (c *Comp) BgRepeatY() *Comp {
	c.El.Style(styles.BgRepeatY())
	return c
}

func (c *Comp) BgRight() *Comp {
	c.El.Style(styles.BgRight())
	return c
}

func (c *Comp) BgRuby(scale int) *Comp {
	c.El.Style(styles.BgRuby(scale))
	return c
}

func (c *Comp) BgRubyAlpha(scale int) *Comp {
	c.El.Style(styles.BgRubyAlpha(scale))
	return c
}

func (c *Comp) BgRubyDark(scale int) *Comp {
	c.El.Style(styles.BgRubyDark(scale))
	return c
}

func (c *Comp) BgRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgRubyDarkAlpha(scale))
	return c
}

func (c *Comp) BgSage(scale int) *Comp {
	c.El.Style(styles.BgSage(scale))
	return c
}

func (c *Comp) BgSageAlpha(scale int) *Comp {
	c.El.Style(styles.BgSageAlpha(scale))
	return c
}

func (c *Comp) BgSageDark(scale int) *Comp {
	c.El.Style(styles.BgSageDark(scale))
	return c
}

func (c *Comp) BgSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgSageDarkAlpha(scale))
	return c
}

func (c *Comp) BgSand(scale int) *Comp {
	c.El.Style(styles.BgSand(scale))
	return c
}

func (c *Comp) BgSandAlpha(scale int) *Comp {
	c.El.Style(styles.BgSandAlpha(scale))
	return c
}

func (c *Comp) BgSandDark(scale int) *Comp {
	c.El.Style(styles.BgSandDark(scale))
	return c
}

func (c *Comp) BgSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgSandDarkAlpha(scale))
	return c
}

func (c *Comp) BgScroll() *Comp {
	c.El.Style(styles.BgScroll())
	return c
}

func (c *Comp) BgSize(val value.Value) *Comp {
	c.El.Style(styles.BgSize(val))
	return c
}

func (c *Comp) BgSky(scale int) *Comp {
	c.El.Style(styles.BgSky(scale))
	return c
}

func (c *Comp) BgSkyAlpha(scale int) *Comp {
	c.El.Style(styles.BgSkyAlpha(scale))
	return c
}

func (c *Comp) BgSkyDark(scale int) *Comp {
	c.El.Style(styles.BgSkyDark(scale))
	return c
}

func (c *Comp) BgSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgSkyDarkAlpha(scale))
	return c
}

func (c *Comp) BgSlate(scale int) *Comp {
	c.El.Style(styles.BgSlate(scale))
	return c
}

func (c *Comp) BgSlateAlpha(scale int) *Comp {
	c.El.Style(styles.BgSlateAlpha(scale))
	return c
}

func (c *Comp) BgSlateDark(scale int) *Comp {
	c.El.Style(styles.BgSlateDark(scale))
	return c
}

func (c *Comp) BgSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgSlateDarkAlpha(scale))
	return c
}

func (c *Comp) BgTeal(scale int) *Comp {
	c.El.Style(styles.BgTeal(scale))
	return c
}

func (c *Comp) BgTealAlpha(scale int) *Comp {
	c.El.Style(styles.BgTealAlpha(scale))
	return c
}

func (c *Comp) BgTealDark(scale int) *Comp {
	c.El.Style(styles.BgTealDark(scale))
	return c
}

func (c *Comp) BgTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgTealDarkAlpha(scale))
	return c
}

func (c *Comp) BgTomato(scale int) *Comp {
	c.El.Style(styles.BgTomato(scale))
	return c
}

func (c *Comp) BgTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.BgTomatoAlpha(scale))
	return c
}

func (c *Comp) BgTomatoDark(scale int) *Comp {
	c.El.Style(styles.BgTomatoDark(scale))
	return c
}

func (c *Comp) BgTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) BgTop() *Comp {
	c.El.Style(styles.BgTop())
	return c
}

func (c *Comp) BgTopLeft() *Comp {
	c.El.Style(styles.BgTopLeft())
	return c
}

func (c *Comp) BgTopRight() *Comp {
	c.El.Style(styles.BgTopRight())
	return c
}

func (c *Comp) BgTransparent() *Comp {
	c.El.Style(styles.BgTransparent())
	return c
}

func (c *Comp) BgViolet(scale int) *Comp {
	c.El.Style(styles.BgViolet(scale))
	return c
}

func (c *Comp) BgVioletAlpha(scale int) *Comp {
	c.El.Style(styles.BgVioletAlpha(scale))
	return c
}

func (c *Comp) BgVioletDark(scale int) *Comp {
	c.El.Style(styles.BgVioletDark(scale))
	return c
}

func (c *Comp) BgVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgVioletDarkAlpha(scale))
	return c
}

func (c *Comp) BgWhite() *Comp {
	c.El.Style(styles.BgWhite())
	return c
}

func (c *Comp) BgWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.BgWhiteAlpha(scale))
	return c
}

func (c *Comp) BgYellow(scale int) *Comp {
	c.El.Style(styles.BgYellow(scale))
	return c
}

func (c *Comp) BgYellowAlpha(scale int) *Comp {
	c.El.Style(styles.BgYellowAlpha(scale))
	return c
}

func (c *Comp) BgYellowDark(scale int) *Comp {
	c.El.Style(styles.BgYellowDark(scale))
	return c
}

func (c *Comp) BgYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BgYellowDarkAlpha(scale))
	return c
}

func (c *Comp) Block() *Comp {
	c.El.Style(styles.Block())
	return c
}

func (c *Comp) Blur(val value.Value) *Comp {
	c.El.Style(styles.Blur(val))
	return c
}

func (c *Comp) Blur2xl() *Comp {
	c.El.Style(styles.Blur2xl())
	return c
}

func (c *Comp) Blur3xl() *Comp {
	c.El.Style(styles.Blur3xl())
	return c
}

func (c *Comp) BlurLg() *Comp {
	c.El.Style(styles.BlurLg())
	return c
}

func (c *Comp) BlurMd() *Comp {
	c.El.Style(styles.BlurMd())
	return c
}

func (c *Comp) BlurNone() *Comp {
	c.El.Style(styles.BlurNone())
	return c
}

func (c *Comp) BlurSm() *Comp {
	c.El.Style(styles.BlurSm())
	return c
}

func (c *Comp) BlurXl() *Comp {
	c.El.Style(styles.BlurXl())
	return c
}

func (c *Comp) BlurXs() *Comp {
	c.El.Style(styles.BlurXs())
	return c
}

func (c *Comp) Border(val ...value.Value) *Comp {
	c.El.Style(styles.Border(val...))
	return c
}

func (c *Comp) BorderAmber(scale int) *Comp {
	c.El.Style(styles.BorderAmber(scale))
	return c
}

func (c *Comp) BorderAmberAlpha(scale int) *Comp {
	c.El.Style(styles.BorderAmberAlpha(scale))
	return c
}

func (c *Comp) BorderAmberDark(scale int) *Comp {
	c.El.Style(styles.BorderAmberDark(scale))
	return c
}

func (c *Comp) BorderAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderAmberDarkAlpha(scale))
	return c
}

func (c *Comp) BorderB(val ...value.Value) *Comp {
	c.El.Style(styles.BorderB(val...))
	return c
}

func (c *Comp) BorderBlack() *Comp {
	c.El.Style(styles.BorderBlack())
	return c
}

func (c *Comp) BorderBlackAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBlackAlpha(scale))
	return c
}

func (c *Comp) BorderBlue(scale int) *Comp {
	c.El.Style(styles.BorderBlue(scale))
	return c
}

func (c *Comp) BorderBlueAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBlueAlpha(scale))
	return c
}

func (c *Comp) BorderBlueDark(scale int) *Comp {
	c.El.Style(styles.BorderBlueDark(scale))
	return c
}

func (c *Comp) BorderBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBlueDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomAmber(scale int) *Comp {
	c.El.Style(styles.BorderBottomAmber(scale))
	return c
}

func (c *Comp) BorderBottomAmberAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomAmberAlpha(scale))
	return c
}

func (c *Comp) BorderBottomAmberDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomAmberDark(scale))
	return c
}

func (c *Comp) BorderBottomAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomAmberDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomBlack() *Comp {
	c.El.Style(styles.BorderBottomBlack())
	return c
}

func (c *Comp) BorderBottomBlackAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomBlackAlpha(scale))
	return c
}

func (c *Comp) BorderBottomBlue(scale int) *Comp {
	c.El.Style(styles.BorderBottomBlue(scale))
	return c
}

func (c *Comp) BorderBottomBlueAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomBlueAlpha(scale))
	return c
}

func (c *Comp) BorderBottomBlueDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomBlueDark(scale))
	return c
}

func (c *Comp) BorderBottomBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomBlueDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomBronze(scale int) *Comp {
	c.El.Style(styles.BorderBottomBronze(scale))
	return c
}

func (c *Comp) BorderBottomBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomBronzeAlpha(scale))
	return c
}

func (c *Comp) BorderBottomBronzeDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomBronzeDark(scale))
	return c
}

func (c *Comp) BorderBottomBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomBrown(scale int) *Comp {
	c.El.Style(styles.BorderBottomBrown(scale))
	return c
}

func (c *Comp) BorderBottomBrownAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomBrownAlpha(scale))
	return c
}

func (c *Comp) BorderBottomBrownDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomBrownDark(scale))
	return c
}

func (c *Comp) BorderBottomBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomBrownDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomColor(val value.Value) *Comp {
	c.El.Style(styles.BorderBottomColor(val))
	return c
}

func (c *Comp) BorderBottomCrimson(scale int) *Comp {
	c.El.Style(styles.BorderBottomCrimson(scale))
	return c
}

func (c *Comp) BorderBottomCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomCrimsonAlpha(scale))
	return c
}

func (c *Comp) BorderBottomCrimsonDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomCrimsonDark(scale))
	return c
}

func (c *Comp) BorderBottomCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomCurrent() *Comp {
	c.El.Style(styles.BorderBottomCurrent())
	return c
}

func (c *Comp) BorderBottomCyan(scale int) *Comp {
	c.El.Style(styles.BorderBottomCyan(scale))
	return c
}

func (c *Comp) BorderBottomCyanAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomCyanAlpha(scale))
	return c
}

func (c *Comp) BorderBottomCyanDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomCyanDark(scale))
	return c
}

func (c *Comp) BorderBottomCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomCyanDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomGold(scale int) *Comp {
	c.El.Style(styles.BorderBottomGold(scale))
	return c
}

func (c *Comp) BorderBottomGoldAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomGoldAlpha(scale))
	return c
}

func (c *Comp) BorderBottomGoldDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomGoldDark(scale))
	return c
}

func (c *Comp) BorderBottomGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomGoldDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomGrass(scale int) *Comp {
	c.El.Style(styles.BorderBottomGrass(scale))
	return c
}

func (c *Comp) BorderBottomGrassAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomGrassAlpha(scale))
	return c
}

func (c *Comp) BorderBottomGrassDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomGrassDark(scale))
	return c
}

func (c *Comp) BorderBottomGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomGrassDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomGray(scale int) *Comp {
	c.El.Style(styles.BorderBottomGray(scale))
	return c
}

func (c *Comp) BorderBottomGrayAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomGrayAlpha(scale))
	return c
}

func (c *Comp) BorderBottomGrayDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomGrayDark(scale))
	return c
}

func (c *Comp) BorderBottomGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomGrayDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomGreen(scale int) *Comp {
	c.El.Style(styles.BorderBottomGreen(scale))
	return c
}

func (c *Comp) BorderBottomGreenAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomGreenAlpha(scale))
	return c
}

func (c *Comp) BorderBottomGreenDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomGreenDark(scale))
	return c
}

func (c *Comp) BorderBottomGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomGreenDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomIndigo(scale int) *Comp {
	c.El.Style(styles.BorderBottomIndigo(scale))
	return c
}

func (c *Comp) BorderBottomIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomIndigoAlpha(scale))
	return c
}

func (c *Comp) BorderBottomIndigoDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomIndigoDark(scale))
	return c
}

func (c *Comp) BorderBottomIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomInherit() *Comp {
	c.El.Style(styles.BorderBottomInherit())
	return c
}

func (c *Comp) BorderBottomIris(scale int) *Comp {
	c.El.Style(styles.BorderBottomIris(scale))
	return c
}

func (c *Comp) BorderBottomIrisAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomIrisAlpha(scale))
	return c
}

func (c *Comp) BorderBottomIrisDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomIrisDark(scale))
	return c
}

func (c *Comp) BorderBottomIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomIrisDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomJade(scale int) *Comp {
	c.El.Style(styles.BorderBottomJade(scale))
	return c
}

func (c *Comp) BorderBottomJadeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomJadeAlpha(scale))
	return c
}

func (c *Comp) BorderBottomJadeDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomJadeDark(scale))
	return c
}

func (c *Comp) BorderBottomJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomJadeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomLime(scale int) *Comp {
	c.El.Style(styles.BorderBottomLime(scale))
	return c
}

func (c *Comp) BorderBottomLimeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomLimeAlpha(scale))
	return c
}

func (c *Comp) BorderBottomLimeDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomLimeDark(scale))
	return c
}

func (c *Comp) BorderBottomLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomLimeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomMauve(scale int) *Comp {
	c.El.Style(styles.BorderBottomMauve(scale))
	return c
}

func (c *Comp) BorderBottomMauveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomMauveAlpha(scale))
	return c
}

func (c *Comp) BorderBottomMauveDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomMauveDark(scale))
	return c
}

func (c *Comp) BorderBottomMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomMauveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomMint(scale int) *Comp {
	c.El.Style(styles.BorderBottomMint(scale))
	return c
}

func (c *Comp) BorderBottomMintAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomMintAlpha(scale))
	return c
}

func (c *Comp) BorderBottomMintDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomMintDark(scale))
	return c
}

func (c *Comp) BorderBottomMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomMintDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomOlive(scale int) *Comp {
	c.El.Style(styles.BorderBottomOlive(scale))
	return c
}

func (c *Comp) BorderBottomOliveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomOliveAlpha(scale))
	return c
}

func (c *Comp) BorderBottomOliveDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomOliveDark(scale))
	return c
}

func (c *Comp) BorderBottomOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomOliveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomOrange(scale int) *Comp {
	c.El.Style(styles.BorderBottomOrange(scale))
	return c
}

func (c *Comp) BorderBottomOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomOrangeAlpha(scale))
	return c
}

func (c *Comp) BorderBottomOrangeDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomOrangeDark(scale))
	return c
}

func (c *Comp) BorderBottomOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomPink(scale int) *Comp {
	c.El.Style(styles.BorderBottomPink(scale))
	return c
}

func (c *Comp) BorderBottomPinkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomPinkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomPinkDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomPinkDark(scale))
	return c
}

func (c *Comp) BorderBottomPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomPinkDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomPlum(scale int) *Comp {
	c.El.Style(styles.BorderBottomPlum(scale))
	return c
}

func (c *Comp) BorderBottomPlumAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomPlumAlpha(scale))
	return c
}

func (c *Comp) BorderBottomPlumDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomPlumDark(scale))
	return c
}

func (c *Comp) BorderBottomPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomPlumDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomPurple(scale int) *Comp {
	c.El.Style(styles.BorderBottomPurple(scale))
	return c
}

func (c *Comp) BorderBottomPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomPurpleAlpha(scale))
	return c
}

func (c *Comp) BorderBottomPurpleDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomPurpleDark(scale))
	return c
}

func (c *Comp) BorderBottomPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomRed(scale int) *Comp {
	c.El.Style(styles.BorderBottomRed(scale))
	return c
}

func (c *Comp) BorderBottomRedAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomRedAlpha(scale))
	return c
}

func (c *Comp) BorderBottomRedDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomRedDark(scale))
	return c
}

func (c *Comp) BorderBottomRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomRedDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomRuby(scale int) *Comp {
	c.El.Style(styles.BorderBottomRuby(scale))
	return c
}

func (c *Comp) BorderBottomRubyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomRubyAlpha(scale))
	return c
}

func (c *Comp) BorderBottomRubyDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomRubyDark(scale))
	return c
}

func (c *Comp) BorderBottomRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomRubyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomSage(scale int) *Comp {
	c.El.Style(styles.BorderBottomSage(scale))
	return c
}

func (c *Comp) BorderBottomSageAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomSageAlpha(scale))
	return c
}

func (c *Comp) BorderBottomSageDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomSageDark(scale))
	return c
}

func (c *Comp) BorderBottomSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomSageDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomSand(scale int) *Comp {
	c.El.Style(styles.BorderBottomSand(scale))
	return c
}

func (c *Comp) BorderBottomSandAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomSandAlpha(scale))
	return c
}

func (c *Comp) BorderBottomSandDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomSandDark(scale))
	return c
}

func (c *Comp) BorderBottomSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomSandDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomSky(scale int) *Comp {
	c.El.Style(styles.BorderBottomSky(scale))
	return c
}

func (c *Comp) BorderBottomSkyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomSkyAlpha(scale))
	return c
}

func (c *Comp) BorderBottomSkyDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomSkyDark(scale))
	return c
}

func (c *Comp) BorderBottomSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomSkyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomSlate(scale int) *Comp {
	c.El.Style(styles.BorderBottomSlate(scale))
	return c
}

func (c *Comp) BorderBottomSlateAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomSlateAlpha(scale))
	return c
}

func (c *Comp) BorderBottomSlateDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomSlateDark(scale))
	return c
}

func (c *Comp) BorderBottomSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomSlateDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomTeal(scale int) *Comp {
	c.El.Style(styles.BorderBottomTeal(scale))
	return c
}

func (c *Comp) BorderBottomTealAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomTealAlpha(scale))
	return c
}

func (c *Comp) BorderBottomTealDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomTealDark(scale))
	return c
}

func (c *Comp) BorderBottomTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomTealDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomTomato(scale int) *Comp {
	c.El.Style(styles.BorderBottomTomato(scale))
	return c
}

func (c *Comp) BorderBottomTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomTomatoAlpha(scale))
	return c
}

func (c *Comp) BorderBottomTomatoDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomTomatoDark(scale))
	return c
}

func (c *Comp) BorderBottomTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomTransparent() *Comp {
	c.El.Style(styles.BorderBottomTransparent())
	return c
}

func (c *Comp) BorderBottomViolet(scale int) *Comp {
	c.El.Style(styles.BorderBottomViolet(scale))
	return c
}

func (c *Comp) BorderBottomVioletAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomVioletAlpha(scale))
	return c
}

func (c *Comp) BorderBottomVioletDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomVioletDark(scale))
	return c
}

func (c *Comp) BorderBottomVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomVioletDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBottomWhite() *Comp {
	c.El.Style(styles.BorderBottomWhite())
	return c
}

func (c *Comp) BorderBottomWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomWhiteAlpha(scale))
	return c
}

func (c *Comp) BorderBottomYellow(scale int) *Comp {
	c.El.Style(styles.BorderBottomYellow(scale))
	return c
}

func (c *Comp) BorderBottomYellowAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomYellowAlpha(scale))
	return c
}

func (c *Comp) BorderBottomYellowDark(scale int) *Comp {
	c.El.Style(styles.BorderBottomYellowDark(scale))
	return c
}

func (c *Comp) BorderBottomYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBottomYellowDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBronze(scale int) *Comp {
	c.El.Style(styles.BorderBronze(scale))
	return c
}

func (c *Comp) BorderBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBronzeAlpha(scale))
	return c
}

func (c *Comp) BorderBronzeDark(scale int) *Comp {
	c.El.Style(styles.BorderBronzeDark(scale))
	return c
}

func (c *Comp) BorderBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderBrown(scale int) *Comp {
	c.El.Style(styles.BorderBrown(scale))
	return c
}

func (c *Comp) BorderBrownAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBrownAlpha(scale))
	return c
}

func (c *Comp) BorderBrownDark(scale int) *Comp {
	c.El.Style(styles.BorderBrownDark(scale))
	return c
}

func (c *Comp) BorderBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderBrownDarkAlpha(scale))
	return c
}

func (c *Comp) BorderCollapse() *Comp {
	c.El.Style(styles.BorderCollapse())
	return c
}

func (c *Comp) BorderColor(val value.Value) *Comp {
	c.El.Style(styles.BorderColor(val))
	return c
}

func (c *Comp) BorderCrimson(scale int) *Comp {
	c.El.Style(styles.BorderCrimson(scale))
	return c
}

func (c *Comp) BorderCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.BorderCrimsonAlpha(scale))
	return c
}

func (c *Comp) BorderCrimsonDark(scale int) *Comp {
	c.El.Style(styles.BorderCrimsonDark(scale))
	return c
}

func (c *Comp) BorderCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) BorderCurrent() *Comp {
	c.El.Style(styles.BorderCurrent())
	return c
}

func (c *Comp) BorderCyan(scale int) *Comp {
	c.El.Style(styles.BorderCyan(scale))
	return c
}

func (c *Comp) BorderCyanAlpha(scale int) *Comp {
	c.El.Style(styles.BorderCyanAlpha(scale))
	return c
}

func (c *Comp) BorderCyanDark(scale int) *Comp {
	c.El.Style(styles.BorderCyanDark(scale))
	return c
}

func (c *Comp) BorderCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderCyanDarkAlpha(scale))
	return c
}

func (c *Comp) BorderDashed() *Comp {
	c.El.Style(styles.BorderDashed())
	return c
}

func (c *Comp) BorderDotted() *Comp {
	c.El.Style(styles.BorderDotted())
	return c
}

func (c *Comp) BorderDouble() *Comp {
	c.El.Style(styles.BorderDouble())
	return c
}

func (c *Comp) BorderE(val ...value.Value) *Comp {
	c.El.Style(styles.BorderE(val...))
	return c
}

func (c *Comp) BorderEndAmber(scale int) *Comp {
	c.El.Style(styles.BorderEndAmber(scale))
	return c
}

func (c *Comp) BorderEndAmberAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndAmberAlpha(scale))
	return c
}

func (c *Comp) BorderEndAmberDark(scale int) *Comp {
	c.El.Style(styles.BorderEndAmberDark(scale))
	return c
}

func (c *Comp) BorderEndAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndAmberDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndBlack() *Comp {
	c.El.Style(styles.BorderEndBlack())
	return c
}

func (c *Comp) BorderEndBlackAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndBlackAlpha(scale))
	return c
}

func (c *Comp) BorderEndBlue(scale int) *Comp {
	c.El.Style(styles.BorderEndBlue(scale))
	return c
}

func (c *Comp) BorderEndBlueAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndBlueAlpha(scale))
	return c
}

func (c *Comp) BorderEndBlueDark(scale int) *Comp {
	c.El.Style(styles.BorderEndBlueDark(scale))
	return c
}

func (c *Comp) BorderEndBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndBlueDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndBronze(scale int) *Comp {
	c.El.Style(styles.BorderEndBronze(scale))
	return c
}

func (c *Comp) BorderEndBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndBronzeAlpha(scale))
	return c
}

func (c *Comp) BorderEndBronzeDark(scale int) *Comp {
	c.El.Style(styles.BorderEndBronzeDark(scale))
	return c
}

func (c *Comp) BorderEndBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndBrown(scale int) *Comp {
	c.El.Style(styles.BorderEndBrown(scale))
	return c
}

func (c *Comp) BorderEndBrownAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndBrownAlpha(scale))
	return c
}

func (c *Comp) BorderEndBrownDark(scale int) *Comp {
	c.El.Style(styles.BorderEndBrownDark(scale))
	return c
}

func (c *Comp) BorderEndBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndBrownDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndColor(val value.Value) *Comp {
	c.El.Style(styles.BorderEndColor(val))
	return c
}

func (c *Comp) BorderEndCrimson(scale int) *Comp {
	c.El.Style(styles.BorderEndCrimson(scale))
	return c
}

func (c *Comp) BorderEndCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndCrimsonAlpha(scale))
	return c
}

func (c *Comp) BorderEndCrimsonDark(scale int) *Comp {
	c.El.Style(styles.BorderEndCrimsonDark(scale))
	return c
}

func (c *Comp) BorderEndCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndCurrent() *Comp {
	c.El.Style(styles.BorderEndCurrent())
	return c
}

func (c *Comp) BorderEndCyan(scale int) *Comp {
	c.El.Style(styles.BorderEndCyan(scale))
	return c
}

func (c *Comp) BorderEndCyanAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndCyanAlpha(scale))
	return c
}

func (c *Comp) BorderEndCyanDark(scale int) *Comp {
	c.El.Style(styles.BorderEndCyanDark(scale))
	return c
}

func (c *Comp) BorderEndCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndCyanDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndGold(scale int) *Comp {
	c.El.Style(styles.BorderEndGold(scale))
	return c
}

func (c *Comp) BorderEndGoldAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndGoldAlpha(scale))
	return c
}

func (c *Comp) BorderEndGoldDark(scale int) *Comp {
	c.El.Style(styles.BorderEndGoldDark(scale))
	return c
}

func (c *Comp) BorderEndGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndGoldDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndGrass(scale int) *Comp {
	c.El.Style(styles.BorderEndGrass(scale))
	return c
}

func (c *Comp) BorderEndGrassAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndGrassAlpha(scale))
	return c
}

func (c *Comp) BorderEndGrassDark(scale int) *Comp {
	c.El.Style(styles.BorderEndGrassDark(scale))
	return c
}

func (c *Comp) BorderEndGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndGrassDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndGray(scale int) *Comp {
	c.El.Style(styles.BorderEndGray(scale))
	return c
}

func (c *Comp) BorderEndGrayAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndGrayAlpha(scale))
	return c
}

func (c *Comp) BorderEndGrayDark(scale int) *Comp {
	c.El.Style(styles.BorderEndGrayDark(scale))
	return c
}

func (c *Comp) BorderEndGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndGrayDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndGreen(scale int) *Comp {
	c.El.Style(styles.BorderEndGreen(scale))
	return c
}

func (c *Comp) BorderEndGreenAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndGreenAlpha(scale))
	return c
}

func (c *Comp) BorderEndGreenDark(scale int) *Comp {
	c.El.Style(styles.BorderEndGreenDark(scale))
	return c
}

func (c *Comp) BorderEndGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndGreenDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndIndigo(scale int) *Comp {
	c.El.Style(styles.BorderEndIndigo(scale))
	return c
}

func (c *Comp) BorderEndIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndIndigoAlpha(scale))
	return c
}

func (c *Comp) BorderEndIndigoDark(scale int) *Comp {
	c.El.Style(styles.BorderEndIndigoDark(scale))
	return c
}

func (c *Comp) BorderEndIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndInherit() *Comp {
	c.El.Style(styles.BorderEndInherit())
	return c
}

func (c *Comp) BorderEndIris(scale int) *Comp {
	c.El.Style(styles.BorderEndIris(scale))
	return c
}

func (c *Comp) BorderEndIrisAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndIrisAlpha(scale))
	return c
}

func (c *Comp) BorderEndIrisDark(scale int) *Comp {
	c.El.Style(styles.BorderEndIrisDark(scale))
	return c
}

func (c *Comp) BorderEndIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndIrisDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndJade(scale int) *Comp {
	c.El.Style(styles.BorderEndJade(scale))
	return c
}

func (c *Comp) BorderEndJadeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndJadeAlpha(scale))
	return c
}

func (c *Comp) BorderEndJadeDark(scale int) *Comp {
	c.El.Style(styles.BorderEndJadeDark(scale))
	return c
}

func (c *Comp) BorderEndJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndJadeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndLime(scale int) *Comp {
	c.El.Style(styles.BorderEndLime(scale))
	return c
}

func (c *Comp) BorderEndLimeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndLimeAlpha(scale))
	return c
}

func (c *Comp) BorderEndLimeDark(scale int) *Comp {
	c.El.Style(styles.BorderEndLimeDark(scale))
	return c
}

func (c *Comp) BorderEndLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndLimeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndMauve(scale int) *Comp {
	c.El.Style(styles.BorderEndMauve(scale))
	return c
}

func (c *Comp) BorderEndMauveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndMauveAlpha(scale))
	return c
}

func (c *Comp) BorderEndMauveDark(scale int) *Comp {
	c.El.Style(styles.BorderEndMauveDark(scale))
	return c
}

func (c *Comp) BorderEndMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndMauveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndMint(scale int) *Comp {
	c.El.Style(styles.BorderEndMint(scale))
	return c
}

func (c *Comp) BorderEndMintAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndMintAlpha(scale))
	return c
}

func (c *Comp) BorderEndMintDark(scale int) *Comp {
	c.El.Style(styles.BorderEndMintDark(scale))
	return c
}

func (c *Comp) BorderEndMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndMintDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndOlive(scale int) *Comp {
	c.El.Style(styles.BorderEndOlive(scale))
	return c
}

func (c *Comp) BorderEndOliveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndOliveAlpha(scale))
	return c
}

func (c *Comp) BorderEndOliveDark(scale int) *Comp {
	c.El.Style(styles.BorderEndOliveDark(scale))
	return c
}

func (c *Comp) BorderEndOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndOliveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndOrange(scale int) *Comp {
	c.El.Style(styles.BorderEndOrange(scale))
	return c
}

func (c *Comp) BorderEndOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndOrangeAlpha(scale))
	return c
}

func (c *Comp) BorderEndOrangeDark(scale int) *Comp {
	c.El.Style(styles.BorderEndOrangeDark(scale))
	return c
}

func (c *Comp) BorderEndOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndPink(scale int) *Comp {
	c.El.Style(styles.BorderEndPink(scale))
	return c
}

func (c *Comp) BorderEndPinkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndPinkAlpha(scale))
	return c
}

func (c *Comp) BorderEndPinkDark(scale int) *Comp {
	c.El.Style(styles.BorderEndPinkDark(scale))
	return c
}

func (c *Comp) BorderEndPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndPinkDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndPlum(scale int) *Comp {
	c.El.Style(styles.BorderEndPlum(scale))
	return c
}

func (c *Comp) BorderEndPlumAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndPlumAlpha(scale))
	return c
}

func (c *Comp) BorderEndPlumDark(scale int) *Comp {
	c.El.Style(styles.BorderEndPlumDark(scale))
	return c
}

func (c *Comp) BorderEndPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndPlumDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndPurple(scale int) *Comp {
	c.El.Style(styles.BorderEndPurple(scale))
	return c
}

func (c *Comp) BorderEndPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndPurpleAlpha(scale))
	return c
}

func (c *Comp) BorderEndPurpleDark(scale int) *Comp {
	c.El.Style(styles.BorderEndPurpleDark(scale))
	return c
}

func (c *Comp) BorderEndPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndRed(scale int) *Comp {
	c.El.Style(styles.BorderEndRed(scale))
	return c
}

func (c *Comp) BorderEndRedAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndRedAlpha(scale))
	return c
}

func (c *Comp) BorderEndRedDark(scale int) *Comp {
	c.El.Style(styles.BorderEndRedDark(scale))
	return c
}

func (c *Comp) BorderEndRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndRedDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndRuby(scale int) *Comp {
	c.El.Style(styles.BorderEndRuby(scale))
	return c
}

func (c *Comp) BorderEndRubyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndRubyAlpha(scale))
	return c
}

func (c *Comp) BorderEndRubyDark(scale int) *Comp {
	c.El.Style(styles.BorderEndRubyDark(scale))
	return c
}

func (c *Comp) BorderEndRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndRubyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndSage(scale int) *Comp {
	c.El.Style(styles.BorderEndSage(scale))
	return c
}

func (c *Comp) BorderEndSageAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndSageAlpha(scale))
	return c
}

func (c *Comp) BorderEndSageDark(scale int) *Comp {
	c.El.Style(styles.BorderEndSageDark(scale))
	return c
}

func (c *Comp) BorderEndSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndSageDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndSand(scale int) *Comp {
	c.El.Style(styles.BorderEndSand(scale))
	return c
}

func (c *Comp) BorderEndSandAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndSandAlpha(scale))
	return c
}

func (c *Comp) BorderEndSandDark(scale int) *Comp {
	c.El.Style(styles.BorderEndSandDark(scale))
	return c
}

func (c *Comp) BorderEndSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndSandDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndSky(scale int) *Comp {
	c.El.Style(styles.BorderEndSky(scale))
	return c
}

func (c *Comp) BorderEndSkyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndSkyAlpha(scale))
	return c
}

func (c *Comp) BorderEndSkyDark(scale int) *Comp {
	c.El.Style(styles.BorderEndSkyDark(scale))
	return c
}

func (c *Comp) BorderEndSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndSkyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndSlate(scale int) *Comp {
	c.El.Style(styles.BorderEndSlate(scale))
	return c
}

func (c *Comp) BorderEndSlateAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndSlateAlpha(scale))
	return c
}

func (c *Comp) BorderEndSlateDark(scale int) *Comp {
	c.El.Style(styles.BorderEndSlateDark(scale))
	return c
}

func (c *Comp) BorderEndSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndSlateDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndTeal(scale int) *Comp {
	c.El.Style(styles.BorderEndTeal(scale))
	return c
}

func (c *Comp) BorderEndTealAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndTealAlpha(scale))
	return c
}

func (c *Comp) BorderEndTealDark(scale int) *Comp {
	c.El.Style(styles.BorderEndTealDark(scale))
	return c
}

func (c *Comp) BorderEndTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndTealDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndTomato(scale int) *Comp {
	c.El.Style(styles.BorderEndTomato(scale))
	return c
}

func (c *Comp) BorderEndTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndTomatoAlpha(scale))
	return c
}

func (c *Comp) BorderEndTomatoDark(scale int) *Comp {
	c.El.Style(styles.BorderEndTomatoDark(scale))
	return c
}

func (c *Comp) BorderEndTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndTransparent() *Comp {
	c.El.Style(styles.BorderEndTransparent())
	return c
}

func (c *Comp) BorderEndViolet(scale int) *Comp {
	c.El.Style(styles.BorderEndViolet(scale))
	return c
}

func (c *Comp) BorderEndVioletAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndVioletAlpha(scale))
	return c
}

func (c *Comp) BorderEndVioletDark(scale int) *Comp {
	c.El.Style(styles.BorderEndVioletDark(scale))
	return c
}

func (c *Comp) BorderEndVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndVioletDarkAlpha(scale))
	return c
}

func (c *Comp) BorderEndWhite() *Comp {
	c.El.Style(styles.BorderEndWhite())
	return c
}

func (c *Comp) BorderEndWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndWhiteAlpha(scale))
	return c
}

func (c *Comp) BorderEndYellow(scale int) *Comp {
	c.El.Style(styles.BorderEndYellow(scale))
	return c
}

func (c *Comp) BorderEndYellowAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndYellowAlpha(scale))
	return c
}

func (c *Comp) BorderEndYellowDark(scale int) *Comp {
	c.El.Style(styles.BorderEndYellowDark(scale))
	return c
}

func (c *Comp) BorderEndYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderEndYellowDarkAlpha(scale))
	return c
}

func (c *Comp) BorderGold(scale int) *Comp {
	c.El.Style(styles.BorderGold(scale))
	return c
}

func (c *Comp) BorderGoldAlpha(scale int) *Comp {
	c.El.Style(styles.BorderGoldAlpha(scale))
	return c
}

func (c *Comp) BorderGoldDark(scale int) *Comp {
	c.El.Style(styles.BorderGoldDark(scale))
	return c
}

func (c *Comp) BorderGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderGoldDarkAlpha(scale))
	return c
}

func (c *Comp) BorderGrass(scale int) *Comp {
	c.El.Style(styles.BorderGrass(scale))
	return c
}

func (c *Comp) BorderGrassAlpha(scale int) *Comp {
	c.El.Style(styles.BorderGrassAlpha(scale))
	return c
}

func (c *Comp) BorderGrassDark(scale int) *Comp {
	c.El.Style(styles.BorderGrassDark(scale))
	return c
}

func (c *Comp) BorderGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderGrassDarkAlpha(scale))
	return c
}

func (c *Comp) BorderGray(scale int) *Comp {
	c.El.Style(styles.BorderGray(scale))
	return c
}

func (c *Comp) BorderGrayAlpha(scale int) *Comp {
	c.El.Style(styles.BorderGrayAlpha(scale))
	return c
}

func (c *Comp) BorderGrayDark(scale int) *Comp {
	c.El.Style(styles.BorderGrayDark(scale))
	return c
}

func (c *Comp) BorderGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderGrayDarkAlpha(scale))
	return c
}

func (c *Comp) BorderGreen(scale int) *Comp {
	c.El.Style(styles.BorderGreen(scale))
	return c
}

func (c *Comp) BorderGreenAlpha(scale int) *Comp {
	c.El.Style(styles.BorderGreenAlpha(scale))
	return c
}

func (c *Comp) BorderGreenDark(scale int) *Comp {
	c.El.Style(styles.BorderGreenDark(scale))
	return c
}

func (c *Comp) BorderGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderGreenDarkAlpha(scale))
	return c
}

func (c *Comp) BorderHidden() *Comp {
	c.El.Style(styles.BorderHidden())
	return c
}

func (c *Comp) BorderIndigo(scale int) *Comp {
	c.El.Style(styles.BorderIndigo(scale))
	return c
}

func (c *Comp) BorderIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderIndigoAlpha(scale))
	return c
}

func (c *Comp) BorderIndigoDark(scale int) *Comp {
	c.El.Style(styles.BorderIndigoDark(scale))
	return c
}

func (c *Comp) BorderIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderInherit() *Comp {
	c.El.Style(styles.BorderInherit())
	return c
}

func (c *Comp) BorderIris(scale int) *Comp {
	c.El.Style(styles.BorderIris(scale))
	return c
}

func (c *Comp) BorderIrisAlpha(scale int) *Comp {
	c.El.Style(styles.BorderIrisAlpha(scale))
	return c
}

func (c *Comp) BorderIrisDark(scale int) *Comp {
	c.El.Style(styles.BorderIrisDark(scale))
	return c
}

func (c *Comp) BorderIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderIrisDarkAlpha(scale))
	return c
}

func (c *Comp) BorderJade(scale int) *Comp {
	c.El.Style(styles.BorderJade(scale))
	return c
}

func (c *Comp) BorderJadeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderJadeAlpha(scale))
	return c
}

func (c *Comp) BorderJadeDark(scale int) *Comp {
	c.El.Style(styles.BorderJadeDark(scale))
	return c
}

func (c *Comp) BorderJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderJadeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderL(val ...value.Value) *Comp {
	c.El.Style(styles.BorderL(val...))
	return c
}

func (c *Comp) BorderLeftAmber(scale int) *Comp {
	c.El.Style(styles.BorderLeftAmber(scale))
	return c
}

func (c *Comp) BorderLeftAmberAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftAmberAlpha(scale))
	return c
}

func (c *Comp) BorderLeftAmberDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftAmberDark(scale))
	return c
}

func (c *Comp) BorderLeftAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftAmberDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftBlack() *Comp {
	c.El.Style(styles.BorderLeftBlack())
	return c
}

func (c *Comp) BorderLeftBlackAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftBlackAlpha(scale))
	return c
}

func (c *Comp) BorderLeftBlue(scale int) *Comp {
	c.El.Style(styles.BorderLeftBlue(scale))
	return c
}

func (c *Comp) BorderLeftBlueAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftBlueAlpha(scale))
	return c
}

func (c *Comp) BorderLeftBlueDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftBlueDark(scale))
	return c
}

func (c *Comp) BorderLeftBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftBlueDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftBronze(scale int) *Comp {
	c.El.Style(styles.BorderLeftBronze(scale))
	return c
}

func (c *Comp) BorderLeftBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftBronzeAlpha(scale))
	return c
}

func (c *Comp) BorderLeftBronzeDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftBronzeDark(scale))
	return c
}

func (c *Comp) BorderLeftBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftBrown(scale int) *Comp {
	c.El.Style(styles.BorderLeftBrown(scale))
	return c
}

func (c *Comp) BorderLeftBrownAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftBrownAlpha(scale))
	return c
}

func (c *Comp) BorderLeftBrownDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftBrownDark(scale))
	return c
}

func (c *Comp) BorderLeftBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftBrownDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftColor(val value.Value) *Comp {
	c.El.Style(styles.BorderLeftColor(val))
	return c
}

func (c *Comp) BorderLeftCrimson(scale int) *Comp {
	c.El.Style(styles.BorderLeftCrimson(scale))
	return c
}

func (c *Comp) BorderLeftCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftCrimsonAlpha(scale))
	return c
}

func (c *Comp) BorderLeftCrimsonDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftCrimsonDark(scale))
	return c
}

func (c *Comp) BorderLeftCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftCurrent() *Comp {
	c.El.Style(styles.BorderLeftCurrent())
	return c
}

func (c *Comp) BorderLeftCyan(scale int) *Comp {
	c.El.Style(styles.BorderLeftCyan(scale))
	return c
}

func (c *Comp) BorderLeftCyanAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftCyanAlpha(scale))
	return c
}

func (c *Comp) BorderLeftCyanDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftCyanDark(scale))
	return c
}

func (c *Comp) BorderLeftCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftCyanDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftGold(scale int) *Comp {
	c.El.Style(styles.BorderLeftGold(scale))
	return c
}

func (c *Comp) BorderLeftGoldAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftGoldAlpha(scale))
	return c
}

func (c *Comp) BorderLeftGoldDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftGoldDark(scale))
	return c
}

func (c *Comp) BorderLeftGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftGoldDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftGrass(scale int) *Comp {
	c.El.Style(styles.BorderLeftGrass(scale))
	return c
}

func (c *Comp) BorderLeftGrassAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftGrassAlpha(scale))
	return c
}

func (c *Comp) BorderLeftGrassDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftGrassDark(scale))
	return c
}

func (c *Comp) BorderLeftGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftGrassDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftGray(scale int) *Comp {
	c.El.Style(styles.BorderLeftGray(scale))
	return c
}

func (c *Comp) BorderLeftGrayAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftGrayAlpha(scale))
	return c
}

func (c *Comp) BorderLeftGrayDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftGrayDark(scale))
	return c
}

func (c *Comp) BorderLeftGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftGrayDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftGreen(scale int) *Comp {
	c.El.Style(styles.BorderLeftGreen(scale))
	return c
}

func (c *Comp) BorderLeftGreenAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftGreenAlpha(scale))
	return c
}

func (c *Comp) BorderLeftGreenDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftGreenDark(scale))
	return c
}

func (c *Comp) BorderLeftGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftGreenDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftIndigo(scale int) *Comp {
	c.El.Style(styles.BorderLeftIndigo(scale))
	return c
}

func (c *Comp) BorderLeftIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftIndigoAlpha(scale))
	return c
}

func (c *Comp) BorderLeftIndigoDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftIndigoDark(scale))
	return c
}

func (c *Comp) BorderLeftIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftInherit() *Comp {
	c.El.Style(styles.BorderLeftInherit())
	return c
}

func (c *Comp) BorderLeftIris(scale int) *Comp {
	c.El.Style(styles.BorderLeftIris(scale))
	return c
}

func (c *Comp) BorderLeftIrisAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftIrisAlpha(scale))
	return c
}

func (c *Comp) BorderLeftIrisDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftIrisDark(scale))
	return c
}

func (c *Comp) BorderLeftIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftIrisDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftJade(scale int) *Comp {
	c.El.Style(styles.BorderLeftJade(scale))
	return c
}

func (c *Comp) BorderLeftJadeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftJadeAlpha(scale))
	return c
}

func (c *Comp) BorderLeftJadeDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftJadeDark(scale))
	return c
}

func (c *Comp) BorderLeftJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftJadeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftLime(scale int) *Comp {
	c.El.Style(styles.BorderLeftLime(scale))
	return c
}

func (c *Comp) BorderLeftLimeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftLimeAlpha(scale))
	return c
}

func (c *Comp) BorderLeftLimeDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftLimeDark(scale))
	return c
}

func (c *Comp) BorderLeftLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftLimeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftMauve(scale int) *Comp {
	c.El.Style(styles.BorderLeftMauve(scale))
	return c
}

func (c *Comp) BorderLeftMauveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftMauveAlpha(scale))
	return c
}

func (c *Comp) BorderLeftMauveDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftMauveDark(scale))
	return c
}

func (c *Comp) BorderLeftMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftMauveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftMint(scale int) *Comp {
	c.El.Style(styles.BorderLeftMint(scale))
	return c
}

func (c *Comp) BorderLeftMintAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftMintAlpha(scale))
	return c
}

func (c *Comp) BorderLeftMintDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftMintDark(scale))
	return c
}

func (c *Comp) BorderLeftMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftMintDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftOlive(scale int) *Comp {
	c.El.Style(styles.BorderLeftOlive(scale))
	return c
}

func (c *Comp) BorderLeftOliveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftOliveAlpha(scale))
	return c
}

func (c *Comp) BorderLeftOliveDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftOliveDark(scale))
	return c
}

func (c *Comp) BorderLeftOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftOliveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftOrange(scale int) *Comp {
	c.El.Style(styles.BorderLeftOrange(scale))
	return c
}

func (c *Comp) BorderLeftOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftOrangeAlpha(scale))
	return c
}

func (c *Comp) BorderLeftOrangeDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftOrangeDark(scale))
	return c
}

func (c *Comp) BorderLeftOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftPink(scale int) *Comp {
	c.El.Style(styles.BorderLeftPink(scale))
	return c
}

func (c *Comp) BorderLeftPinkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftPinkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftPinkDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftPinkDark(scale))
	return c
}

func (c *Comp) BorderLeftPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftPinkDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftPlum(scale int) *Comp {
	c.El.Style(styles.BorderLeftPlum(scale))
	return c
}

func (c *Comp) BorderLeftPlumAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftPlumAlpha(scale))
	return c
}

func (c *Comp) BorderLeftPlumDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftPlumDark(scale))
	return c
}

func (c *Comp) BorderLeftPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftPlumDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftPurple(scale int) *Comp {
	c.El.Style(styles.BorderLeftPurple(scale))
	return c
}

func (c *Comp) BorderLeftPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftPurpleAlpha(scale))
	return c
}

func (c *Comp) BorderLeftPurpleDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftPurpleDark(scale))
	return c
}

func (c *Comp) BorderLeftPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftRed(scale int) *Comp {
	c.El.Style(styles.BorderLeftRed(scale))
	return c
}

func (c *Comp) BorderLeftRedAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftRedAlpha(scale))
	return c
}

func (c *Comp) BorderLeftRedDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftRedDark(scale))
	return c
}

func (c *Comp) BorderLeftRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftRedDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftRuby(scale int) *Comp {
	c.El.Style(styles.BorderLeftRuby(scale))
	return c
}

func (c *Comp) BorderLeftRubyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftRubyAlpha(scale))
	return c
}

func (c *Comp) BorderLeftRubyDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftRubyDark(scale))
	return c
}

func (c *Comp) BorderLeftRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftRubyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftSage(scale int) *Comp {
	c.El.Style(styles.BorderLeftSage(scale))
	return c
}

func (c *Comp) BorderLeftSageAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftSageAlpha(scale))
	return c
}

func (c *Comp) BorderLeftSageDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftSageDark(scale))
	return c
}

func (c *Comp) BorderLeftSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftSageDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftSand(scale int) *Comp {
	c.El.Style(styles.BorderLeftSand(scale))
	return c
}

func (c *Comp) BorderLeftSandAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftSandAlpha(scale))
	return c
}

func (c *Comp) BorderLeftSandDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftSandDark(scale))
	return c
}

func (c *Comp) BorderLeftSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftSandDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftSky(scale int) *Comp {
	c.El.Style(styles.BorderLeftSky(scale))
	return c
}

func (c *Comp) BorderLeftSkyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftSkyAlpha(scale))
	return c
}

func (c *Comp) BorderLeftSkyDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftSkyDark(scale))
	return c
}

func (c *Comp) BorderLeftSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftSkyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftSlate(scale int) *Comp {
	c.El.Style(styles.BorderLeftSlate(scale))
	return c
}

func (c *Comp) BorderLeftSlateAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftSlateAlpha(scale))
	return c
}

func (c *Comp) BorderLeftSlateDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftSlateDark(scale))
	return c
}

func (c *Comp) BorderLeftSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftSlateDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftTeal(scale int) *Comp {
	c.El.Style(styles.BorderLeftTeal(scale))
	return c
}

func (c *Comp) BorderLeftTealAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftTealAlpha(scale))
	return c
}

func (c *Comp) BorderLeftTealDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftTealDark(scale))
	return c
}

func (c *Comp) BorderLeftTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftTealDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftTomato(scale int) *Comp {
	c.El.Style(styles.BorderLeftTomato(scale))
	return c
}

func (c *Comp) BorderLeftTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftTomatoAlpha(scale))
	return c
}

func (c *Comp) BorderLeftTomatoDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftTomatoDark(scale))
	return c
}

func (c *Comp) BorderLeftTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftTransparent() *Comp {
	c.El.Style(styles.BorderLeftTransparent())
	return c
}

func (c *Comp) BorderLeftViolet(scale int) *Comp {
	c.El.Style(styles.BorderLeftViolet(scale))
	return c
}

func (c *Comp) BorderLeftVioletAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftVioletAlpha(scale))
	return c
}

func (c *Comp) BorderLeftVioletDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftVioletDark(scale))
	return c
}

func (c *Comp) BorderLeftVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftVioletDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLeftWhite() *Comp {
	c.El.Style(styles.BorderLeftWhite())
	return c
}

func (c *Comp) BorderLeftWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftWhiteAlpha(scale))
	return c
}

func (c *Comp) BorderLeftYellow(scale int) *Comp {
	c.El.Style(styles.BorderLeftYellow(scale))
	return c
}

func (c *Comp) BorderLeftYellowAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftYellowAlpha(scale))
	return c
}

func (c *Comp) BorderLeftYellowDark(scale int) *Comp {
	c.El.Style(styles.BorderLeftYellowDark(scale))
	return c
}

func (c *Comp) BorderLeftYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLeftYellowDarkAlpha(scale))
	return c
}

func (c *Comp) BorderLime(scale int) *Comp {
	c.El.Style(styles.BorderLime(scale))
	return c
}

func (c *Comp) BorderLimeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLimeAlpha(scale))
	return c
}

func (c *Comp) BorderLimeDark(scale int) *Comp {
	c.El.Style(styles.BorderLimeDark(scale))
	return c
}

func (c *Comp) BorderLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderLimeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderMauve(scale int) *Comp {
	c.El.Style(styles.BorderMauve(scale))
	return c
}

func (c *Comp) BorderMauveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderMauveAlpha(scale))
	return c
}

func (c *Comp) BorderMauveDark(scale int) *Comp {
	c.El.Style(styles.BorderMauveDark(scale))
	return c
}

func (c *Comp) BorderMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderMauveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderMint(scale int) *Comp {
	c.El.Style(styles.BorderMint(scale))
	return c
}

func (c *Comp) BorderMintAlpha(scale int) *Comp {
	c.El.Style(styles.BorderMintAlpha(scale))
	return c
}

func (c *Comp) BorderMintDark(scale int) *Comp {
	c.El.Style(styles.BorderMintDark(scale))
	return c
}

func (c *Comp) BorderMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderMintDarkAlpha(scale))
	return c
}

func (c *Comp) BorderNone() *Comp {
	c.El.Style(styles.BorderNone())
	return c
}

func (c *Comp) BorderOlive(scale int) *Comp {
	c.El.Style(styles.BorderOlive(scale))
	return c
}

func (c *Comp) BorderOliveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderOliveAlpha(scale))
	return c
}

func (c *Comp) BorderOliveDark(scale int) *Comp {
	c.El.Style(styles.BorderOliveDark(scale))
	return c
}

func (c *Comp) BorderOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderOliveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderOrange(scale int) *Comp {
	c.El.Style(styles.BorderOrange(scale))
	return c
}

func (c *Comp) BorderOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderOrangeAlpha(scale))
	return c
}

func (c *Comp) BorderOrangeDark(scale int) *Comp {
	c.El.Style(styles.BorderOrangeDark(scale))
	return c
}

func (c *Comp) BorderOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderPink(scale int) *Comp {
	c.El.Style(styles.BorderPink(scale))
	return c
}

func (c *Comp) BorderPinkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderPinkAlpha(scale))
	return c
}

func (c *Comp) BorderPinkDark(scale int) *Comp {
	c.El.Style(styles.BorderPinkDark(scale))
	return c
}

func (c *Comp) BorderPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderPinkDarkAlpha(scale))
	return c
}

func (c *Comp) BorderPlum(scale int) *Comp {
	c.El.Style(styles.BorderPlum(scale))
	return c
}

func (c *Comp) BorderPlumAlpha(scale int) *Comp {
	c.El.Style(styles.BorderPlumAlpha(scale))
	return c
}

func (c *Comp) BorderPlumDark(scale int) *Comp {
	c.El.Style(styles.BorderPlumDark(scale))
	return c
}

func (c *Comp) BorderPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderPlumDarkAlpha(scale))
	return c
}

func (c *Comp) BorderPurple(scale int) *Comp {
	c.El.Style(styles.BorderPurple(scale))
	return c
}

func (c *Comp) BorderPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.BorderPurpleAlpha(scale))
	return c
}

func (c *Comp) BorderPurpleDark(scale int) *Comp {
	c.El.Style(styles.BorderPurpleDark(scale))
	return c
}

func (c *Comp) BorderPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) BorderR(val ...value.Value) *Comp {
	c.El.Style(styles.BorderR(val...))
	return c
}

func (c *Comp) BorderRed(scale int) *Comp {
	c.El.Style(styles.BorderRed(scale))
	return c
}

func (c *Comp) BorderRedAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRedAlpha(scale))
	return c
}

func (c *Comp) BorderRedDark(scale int) *Comp {
	c.El.Style(styles.BorderRedDark(scale))
	return c
}

func (c *Comp) BorderRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRedDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightAmber(scale int) *Comp {
	c.El.Style(styles.BorderRightAmber(scale))
	return c
}

func (c *Comp) BorderRightAmberAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightAmberAlpha(scale))
	return c
}

func (c *Comp) BorderRightAmberDark(scale int) *Comp {
	c.El.Style(styles.BorderRightAmberDark(scale))
	return c
}

func (c *Comp) BorderRightAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightAmberDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightBlack() *Comp {
	c.El.Style(styles.BorderRightBlack())
	return c
}

func (c *Comp) BorderRightBlackAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightBlackAlpha(scale))
	return c
}

func (c *Comp) BorderRightBlue(scale int) *Comp {
	c.El.Style(styles.BorderRightBlue(scale))
	return c
}

func (c *Comp) BorderRightBlueAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightBlueAlpha(scale))
	return c
}

func (c *Comp) BorderRightBlueDark(scale int) *Comp {
	c.El.Style(styles.BorderRightBlueDark(scale))
	return c
}

func (c *Comp) BorderRightBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightBlueDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightBronze(scale int) *Comp {
	c.El.Style(styles.BorderRightBronze(scale))
	return c
}

func (c *Comp) BorderRightBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightBronzeAlpha(scale))
	return c
}

func (c *Comp) BorderRightBronzeDark(scale int) *Comp {
	c.El.Style(styles.BorderRightBronzeDark(scale))
	return c
}

func (c *Comp) BorderRightBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightBrown(scale int) *Comp {
	c.El.Style(styles.BorderRightBrown(scale))
	return c
}

func (c *Comp) BorderRightBrownAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightBrownAlpha(scale))
	return c
}

func (c *Comp) BorderRightBrownDark(scale int) *Comp {
	c.El.Style(styles.BorderRightBrownDark(scale))
	return c
}

func (c *Comp) BorderRightBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightBrownDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightColor(val value.Value) *Comp {
	c.El.Style(styles.BorderRightColor(val))
	return c
}

func (c *Comp) BorderRightCrimson(scale int) *Comp {
	c.El.Style(styles.BorderRightCrimson(scale))
	return c
}

func (c *Comp) BorderRightCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightCrimsonAlpha(scale))
	return c
}

func (c *Comp) BorderRightCrimsonDark(scale int) *Comp {
	c.El.Style(styles.BorderRightCrimsonDark(scale))
	return c
}

func (c *Comp) BorderRightCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightCurrent() *Comp {
	c.El.Style(styles.BorderRightCurrent())
	return c
}

func (c *Comp) BorderRightCyan(scale int) *Comp {
	c.El.Style(styles.BorderRightCyan(scale))
	return c
}

func (c *Comp) BorderRightCyanAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightCyanAlpha(scale))
	return c
}

func (c *Comp) BorderRightCyanDark(scale int) *Comp {
	c.El.Style(styles.BorderRightCyanDark(scale))
	return c
}

func (c *Comp) BorderRightCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightCyanDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightGold(scale int) *Comp {
	c.El.Style(styles.BorderRightGold(scale))
	return c
}

func (c *Comp) BorderRightGoldAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightGoldAlpha(scale))
	return c
}

func (c *Comp) BorderRightGoldDark(scale int) *Comp {
	c.El.Style(styles.BorderRightGoldDark(scale))
	return c
}

func (c *Comp) BorderRightGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightGoldDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightGrass(scale int) *Comp {
	c.El.Style(styles.BorderRightGrass(scale))
	return c
}

func (c *Comp) BorderRightGrassAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightGrassAlpha(scale))
	return c
}

func (c *Comp) BorderRightGrassDark(scale int) *Comp {
	c.El.Style(styles.BorderRightGrassDark(scale))
	return c
}

func (c *Comp) BorderRightGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightGrassDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightGray(scale int) *Comp {
	c.El.Style(styles.BorderRightGray(scale))
	return c
}

func (c *Comp) BorderRightGrayAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightGrayAlpha(scale))
	return c
}

func (c *Comp) BorderRightGrayDark(scale int) *Comp {
	c.El.Style(styles.BorderRightGrayDark(scale))
	return c
}

func (c *Comp) BorderRightGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightGrayDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightGreen(scale int) *Comp {
	c.El.Style(styles.BorderRightGreen(scale))
	return c
}

func (c *Comp) BorderRightGreenAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightGreenAlpha(scale))
	return c
}

func (c *Comp) BorderRightGreenDark(scale int) *Comp {
	c.El.Style(styles.BorderRightGreenDark(scale))
	return c
}

func (c *Comp) BorderRightGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightGreenDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightIndigo(scale int) *Comp {
	c.El.Style(styles.BorderRightIndigo(scale))
	return c
}

func (c *Comp) BorderRightIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightIndigoAlpha(scale))
	return c
}

func (c *Comp) BorderRightIndigoDark(scale int) *Comp {
	c.El.Style(styles.BorderRightIndigoDark(scale))
	return c
}

func (c *Comp) BorderRightIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightInherit() *Comp {
	c.El.Style(styles.BorderRightInherit())
	return c
}

func (c *Comp) BorderRightIris(scale int) *Comp {
	c.El.Style(styles.BorderRightIris(scale))
	return c
}

func (c *Comp) BorderRightIrisAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightIrisAlpha(scale))
	return c
}

func (c *Comp) BorderRightIrisDark(scale int) *Comp {
	c.El.Style(styles.BorderRightIrisDark(scale))
	return c
}

func (c *Comp) BorderRightIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightIrisDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightJade(scale int) *Comp {
	c.El.Style(styles.BorderRightJade(scale))
	return c
}

func (c *Comp) BorderRightJadeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightJadeAlpha(scale))
	return c
}

func (c *Comp) BorderRightJadeDark(scale int) *Comp {
	c.El.Style(styles.BorderRightJadeDark(scale))
	return c
}

func (c *Comp) BorderRightJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightJadeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightLime(scale int) *Comp {
	c.El.Style(styles.BorderRightLime(scale))
	return c
}

func (c *Comp) BorderRightLimeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightLimeAlpha(scale))
	return c
}

func (c *Comp) BorderRightLimeDark(scale int) *Comp {
	c.El.Style(styles.BorderRightLimeDark(scale))
	return c
}

func (c *Comp) BorderRightLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightLimeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightMauve(scale int) *Comp {
	c.El.Style(styles.BorderRightMauve(scale))
	return c
}

func (c *Comp) BorderRightMauveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightMauveAlpha(scale))
	return c
}

func (c *Comp) BorderRightMauveDark(scale int) *Comp {
	c.El.Style(styles.BorderRightMauveDark(scale))
	return c
}

func (c *Comp) BorderRightMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightMauveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightMint(scale int) *Comp {
	c.El.Style(styles.BorderRightMint(scale))
	return c
}

func (c *Comp) BorderRightMintAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightMintAlpha(scale))
	return c
}

func (c *Comp) BorderRightMintDark(scale int) *Comp {
	c.El.Style(styles.BorderRightMintDark(scale))
	return c
}

func (c *Comp) BorderRightMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightMintDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightOlive(scale int) *Comp {
	c.El.Style(styles.BorderRightOlive(scale))
	return c
}

func (c *Comp) BorderRightOliveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightOliveAlpha(scale))
	return c
}

func (c *Comp) BorderRightOliveDark(scale int) *Comp {
	c.El.Style(styles.BorderRightOliveDark(scale))
	return c
}

func (c *Comp) BorderRightOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightOliveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightOrange(scale int) *Comp {
	c.El.Style(styles.BorderRightOrange(scale))
	return c
}

func (c *Comp) BorderRightOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightOrangeAlpha(scale))
	return c
}

func (c *Comp) BorderRightOrangeDark(scale int) *Comp {
	c.El.Style(styles.BorderRightOrangeDark(scale))
	return c
}

func (c *Comp) BorderRightOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightPink(scale int) *Comp {
	c.El.Style(styles.BorderRightPink(scale))
	return c
}

func (c *Comp) BorderRightPinkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightPinkAlpha(scale))
	return c
}

func (c *Comp) BorderRightPinkDark(scale int) *Comp {
	c.El.Style(styles.BorderRightPinkDark(scale))
	return c
}

func (c *Comp) BorderRightPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightPinkDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightPlum(scale int) *Comp {
	c.El.Style(styles.BorderRightPlum(scale))
	return c
}

func (c *Comp) BorderRightPlumAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightPlumAlpha(scale))
	return c
}

func (c *Comp) BorderRightPlumDark(scale int) *Comp {
	c.El.Style(styles.BorderRightPlumDark(scale))
	return c
}

func (c *Comp) BorderRightPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightPlumDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightPurple(scale int) *Comp {
	c.El.Style(styles.BorderRightPurple(scale))
	return c
}

func (c *Comp) BorderRightPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightPurpleAlpha(scale))
	return c
}

func (c *Comp) BorderRightPurpleDark(scale int) *Comp {
	c.El.Style(styles.BorderRightPurpleDark(scale))
	return c
}

func (c *Comp) BorderRightPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightRed(scale int) *Comp {
	c.El.Style(styles.BorderRightRed(scale))
	return c
}

func (c *Comp) BorderRightRedAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightRedAlpha(scale))
	return c
}

func (c *Comp) BorderRightRedDark(scale int) *Comp {
	c.El.Style(styles.BorderRightRedDark(scale))
	return c
}

func (c *Comp) BorderRightRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightRedDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightRuby(scale int) *Comp {
	c.El.Style(styles.BorderRightRuby(scale))
	return c
}

func (c *Comp) BorderRightRubyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightRubyAlpha(scale))
	return c
}

func (c *Comp) BorderRightRubyDark(scale int) *Comp {
	c.El.Style(styles.BorderRightRubyDark(scale))
	return c
}

func (c *Comp) BorderRightRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightRubyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightSage(scale int) *Comp {
	c.El.Style(styles.BorderRightSage(scale))
	return c
}

func (c *Comp) BorderRightSageAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightSageAlpha(scale))
	return c
}

func (c *Comp) BorderRightSageDark(scale int) *Comp {
	c.El.Style(styles.BorderRightSageDark(scale))
	return c
}

func (c *Comp) BorderRightSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightSageDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightSand(scale int) *Comp {
	c.El.Style(styles.BorderRightSand(scale))
	return c
}

func (c *Comp) BorderRightSandAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightSandAlpha(scale))
	return c
}

func (c *Comp) BorderRightSandDark(scale int) *Comp {
	c.El.Style(styles.BorderRightSandDark(scale))
	return c
}

func (c *Comp) BorderRightSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightSandDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightSky(scale int) *Comp {
	c.El.Style(styles.BorderRightSky(scale))
	return c
}

func (c *Comp) BorderRightSkyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightSkyAlpha(scale))
	return c
}

func (c *Comp) BorderRightSkyDark(scale int) *Comp {
	c.El.Style(styles.BorderRightSkyDark(scale))
	return c
}

func (c *Comp) BorderRightSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightSkyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightSlate(scale int) *Comp {
	c.El.Style(styles.BorderRightSlate(scale))
	return c
}

func (c *Comp) BorderRightSlateAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightSlateAlpha(scale))
	return c
}

func (c *Comp) BorderRightSlateDark(scale int) *Comp {
	c.El.Style(styles.BorderRightSlateDark(scale))
	return c
}

func (c *Comp) BorderRightSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightSlateDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightTeal(scale int) *Comp {
	c.El.Style(styles.BorderRightTeal(scale))
	return c
}

func (c *Comp) BorderRightTealAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightTealAlpha(scale))
	return c
}

func (c *Comp) BorderRightTealDark(scale int) *Comp {
	c.El.Style(styles.BorderRightTealDark(scale))
	return c
}

func (c *Comp) BorderRightTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightTealDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightTomato(scale int) *Comp {
	c.El.Style(styles.BorderRightTomato(scale))
	return c
}

func (c *Comp) BorderRightTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightTomatoAlpha(scale))
	return c
}

func (c *Comp) BorderRightTomatoDark(scale int) *Comp {
	c.El.Style(styles.BorderRightTomatoDark(scale))
	return c
}

func (c *Comp) BorderRightTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightTransparent() *Comp {
	c.El.Style(styles.BorderRightTransparent())
	return c
}

func (c *Comp) BorderRightViolet(scale int) *Comp {
	c.El.Style(styles.BorderRightViolet(scale))
	return c
}

func (c *Comp) BorderRightVioletAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightVioletAlpha(scale))
	return c
}

func (c *Comp) BorderRightVioletDark(scale int) *Comp {
	c.El.Style(styles.BorderRightVioletDark(scale))
	return c
}

func (c *Comp) BorderRightVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightVioletDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRightWhite() *Comp {
	c.El.Style(styles.BorderRightWhite())
	return c
}

func (c *Comp) BorderRightWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightWhiteAlpha(scale))
	return c
}

func (c *Comp) BorderRightYellow(scale int) *Comp {
	c.El.Style(styles.BorderRightYellow(scale))
	return c
}

func (c *Comp) BorderRightYellowAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightYellowAlpha(scale))
	return c
}

func (c *Comp) BorderRightYellowDark(scale int) *Comp {
	c.El.Style(styles.BorderRightYellowDark(scale))
	return c
}

func (c *Comp) BorderRightYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRightYellowDarkAlpha(scale))
	return c
}

func (c *Comp) BorderRuby(scale int) *Comp {
	c.El.Style(styles.BorderRuby(scale))
	return c
}

func (c *Comp) BorderRubyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRubyAlpha(scale))
	return c
}

func (c *Comp) BorderRubyDark(scale int) *Comp {
	c.El.Style(styles.BorderRubyDark(scale))
	return c
}

func (c *Comp) BorderRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderRubyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderS(val ...value.Value) *Comp {
	c.El.Style(styles.BorderS(val...))
	return c
}

func (c *Comp) BorderSage(scale int) *Comp {
	c.El.Style(styles.BorderSage(scale))
	return c
}

func (c *Comp) BorderSageAlpha(scale int) *Comp {
	c.El.Style(styles.BorderSageAlpha(scale))
	return c
}

func (c *Comp) BorderSageDark(scale int) *Comp {
	c.El.Style(styles.BorderSageDark(scale))
	return c
}

func (c *Comp) BorderSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderSageDarkAlpha(scale))
	return c
}

func (c *Comp) BorderSand(scale int) *Comp {
	c.El.Style(styles.BorderSand(scale))
	return c
}

func (c *Comp) BorderSandAlpha(scale int) *Comp {
	c.El.Style(styles.BorderSandAlpha(scale))
	return c
}

func (c *Comp) BorderSandDark(scale int) *Comp {
	c.El.Style(styles.BorderSandDark(scale))
	return c
}

func (c *Comp) BorderSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderSandDarkAlpha(scale))
	return c
}

func (c *Comp) BorderSeparate() *Comp {
	c.El.Style(styles.BorderSeparate())
	return c
}

func (c *Comp) BorderSky(scale int) *Comp {
	c.El.Style(styles.BorderSky(scale))
	return c
}

func (c *Comp) BorderSkyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderSkyAlpha(scale))
	return c
}

func (c *Comp) BorderSkyDark(scale int) *Comp {
	c.El.Style(styles.BorderSkyDark(scale))
	return c
}

func (c *Comp) BorderSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderSkyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderSlate(scale int) *Comp {
	c.El.Style(styles.BorderSlate(scale))
	return c
}

func (c *Comp) BorderSlateAlpha(scale int) *Comp {
	c.El.Style(styles.BorderSlateAlpha(scale))
	return c
}

func (c *Comp) BorderSlateDark(scale int) *Comp {
	c.El.Style(styles.BorderSlateDark(scale))
	return c
}

func (c *Comp) BorderSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderSlateDarkAlpha(scale))
	return c
}

func (c *Comp) BorderSolid() *Comp {
	c.El.Style(styles.BorderSolid())
	return c
}

func (c *Comp) BorderSpacing(spacing any) *Comp {
	c.El.Style(styles.BorderSpacing(spacing))
	return c
}

func (c *Comp) BorderSpacingX(spacing any) *Comp {
	c.El.Style(styles.BorderSpacingX(spacing))
	return c
}

func (c *Comp) BorderSpacingY(spacing any) *Comp {
	c.El.Style(styles.BorderSpacingY(spacing))
	return c
}

func (c *Comp) BorderStartAmber(scale int) *Comp {
	c.El.Style(styles.BorderStartAmber(scale))
	return c
}

func (c *Comp) BorderStartAmberAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartAmberAlpha(scale))
	return c
}

func (c *Comp) BorderStartAmberDark(scale int) *Comp {
	c.El.Style(styles.BorderStartAmberDark(scale))
	return c
}

func (c *Comp) BorderStartAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartAmberDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartBlack() *Comp {
	c.El.Style(styles.BorderStartBlack())
	return c
}

func (c *Comp) BorderStartBlackAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartBlackAlpha(scale))
	return c
}

func (c *Comp) BorderStartBlue(scale int) *Comp {
	c.El.Style(styles.BorderStartBlue(scale))
	return c
}

func (c *Comp) BorderStartBlueAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartBlueAlpha(scale))
	return c
}

func (c *Comp) BorderStartBlueDark(scale int) *Comp {
	c.El.Style(styles.BorderStartBlueDark(scale))
	return c
}

func (c *Comp) BorderStartBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartBlueDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartBronze(scale int) *Comp {
	c.El.Style(styles.BorderStartBronze(scale))
	return c
}

func (c *Comp) BorderStartBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartBronzeAlpha(scale))
	return c
}

func (c *Comp) BorderStartBronzeDark(scale int) *Comp {
	c.El.Style(styles.BorderStartBronzeDark(scale))
	return c
}

func (c *Comp) BorderStartBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartBrown(scale int) *Comp {
	c.El.Style(styles.BorderStartBrown(scale))
	return c
}

func (c *Comp) BorderStartBrownAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartBrownAlpha(scale))
	return c
}

func (c *Comp) BorderStartBrownDark(scale int) *Comp {
	c.El.Style(styles.BorderStartBrownDark(scale))
	return c
}

func (c *Comp) BorderStartBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartBrownDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartColor(val value.Value) *Comp {
	c.El.Style(styles.BorderStartColor(val))
	return c
}

func (c *Comp) BorderStartCrimson(scale int) *Comp {
	c.El.Style(styles.BorderStartCrimson(scale))
	return c
}

func (c *Comp) BorderStartCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartCrimsonAlpha(scale))
	return c
}

func (c *Comp) BorderStartCrimsonDark(scale int) *Comp {
	c.El.Style(styles.BorderStartCrimsonDark(scale))
	return c
}

func (c *Comp) BorderStartCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartCurrent() *Comp {
	c.El.Style(styles.BorderStartCurrent())
	return c
}

func (c *Comp) BorderStartCyan(scale int) *Comp {
	c.El.Style(styles.BorderStartCyan(scale))
	return c
}

func (c *Comp) BorderStartCyanAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartCyanAlpha(scale))
	return c
}

func (c *Comp) BorderStartCyanDark(scale int) *Comp {
	c.El.Style(styles.BorderStartCyanDark(scale))
	return c
}

func (c *Comp) BorderStartCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartCyanDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartGold(scale int) *Comp {
	c.El.Style(styles.BorderStartGold(scale))
	return c
}

func (c *Comp) BorderStartGoldAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartGoldAlpha(scale))
	return c
}

func (c *Comp) BorderStartGoldDark(scale int) *Comp {
	c.El.Style(styles.BorderStartGoldDark(scale))
	return c
}

func (c *Comp) BorderStartGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartGoldDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartGrass(scale int) *Comp {
	c.El.Style(styles.BorderStartGrass(scale))
	return c
}

func (c *Comp) BorderStartGrassAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartGrassAlpha(scale))
	return c
}

func (c *Comp) BorderStartGrassDark(scale int) *Comp {
	c.El.Style(styles.BorderStartGrassDark(scale))
	return c
}

func (c *Comp) BorderStartGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartGrassDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartGray(scale int) *Comp {
	c.El.Style(styles.BorderStartGray(scale))
	return c
}

func (c *Comp) BorderStartGrayAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartGrayAlpha(scale))
	return c
}

func (c *Comp) BorderStartGrayDark(scale int) *Comp {
	c.El.Style(styles.BorderStartGrayDark(scale))
	return c
}

func (c *Comp) BorderStartGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartGrayDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartGreen(scale int) *Comp {
	c.El.Style(styles.BorderStartGreen(scale))
	return c
}

func (c *Comp) BorderStartGreenAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartGreenAlpha(scale))
	return c
}

func (c *Comp) BorderStartGreenDark(scale int) *Comp {
	c.El.Style(styles.BorderStartGreenDark(scale))
	return c
}

func (c *Comp) BorderStartGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartGreenDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartIndigo(scale int) *Comp {
	c.El.Style(styles.BorderStartIndigo(scale))
	return c
}

func (c *Comp) BorderStartIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartIndigoAlpha(scale))
	return c
}

func (c *Comp) BorderStartIndigoDark(scale int) *Comp {
	c.El.Style(styles.BorderStartIndigoDark(scale))
	return c
}

func (c *Comp) BorderStartIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartInherit() *Comp {
	c.El.Style(styles.BorderStartInherit())
	return c
}

func (c *Comp) BorderStartIris(scale int) *Comp {
	c.El.Style(styles.BorderStartIris(scale))
	return c
}

func (c *Comp) BorderStartIrisAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartIrisAlpha(scale))
	return c
}

func (c *Comp) BorderStartIrisDark(scale int) *Comp {
	c.El.Style(styles.BorderStartIrisDark(scale))
	return c
}

func (c *Comp) BorderStartIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartIrisDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartJade(scale int) *Comp {
	c.El.Style(styles.BorderStartJade(scale))
	return c
}

func (c *Comp) BorderStartJadeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartJadeAlpha(scale))
	return c
}

func (c *Comp) BorderStartJadeDark(scale int) *Comp {
	c.El.Style(styles.BorderStartJadeDark(scale))
	return c
}

func (c *Comp) BorderStartJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartJadeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartLime(scale int) *Comp {
	c.El.Style(styles.BorderStartLime(scale))
	return c
}

func (c *Comp) BorderStartLimeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartLimeAlpha(scale))
	return c
}

func (c *Comp) BorderStartLimeDark(scale int) *Comp {
	c.El.Style(styles.BorderStartLimeDark(scale))
	return c
}

func (c *Comp) BorderStartLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartLimeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartMauve(scale int) *Comp {
	c.El.Style(styles.BorderStartMauve(scale))
	return c
}

func (c *Comp) BorderStartMauveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartMauveAlpha(scale))
	return c
}

func (c *Comp) BorderStartMauveDark(scale int) *Comp {
	c.El.Style(styles.BorderStartMauveDark(scale))
	return c
}

func (c *Comp) BorderStartMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartMauveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartMint(scale int) *Comp {
	c.El.Style(styles.BorderStartMint(scale))
	return c
}

func (c *Comp) BorderStartMintAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartMintAlpha(scale))
	return c
}

func (c *Comp) BorderStartMintDark(scale int) *Comp {
	c.El.Style(styles.BorderStartMintDark(scale))
	return c
}

func (c *Comp) BorderStartMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartMintDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartOlive(scale int) *Comp {
	c.El.Style(styles.BorderStartOlive(scale))
	return c
}

func (c *Comp) BorderStartOliveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartOliveAlpha(scale))
	return c
}

func (c *Comp) BorderStartOliveDark(scale int) *Comp {
	c.El.Style(styles.BorderStartOliveDark(scale))
	return c
}

func (c *Comp) BorderStartOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartOliveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartOrange(scale int) *Comp {
	c.El.Style(styles.BorderStartOrange(scale))
	return c
}

func (c *Comp) BorderStartOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartOrangeAlpha(scale))
	return c
}

func (c *Comp) BorderStartOrangeDark(scale int) *Comp {
	c.El.Style(styles.BorderStartOrangeDark(scale))
	return c
}

func (c *Comp) BorderStartOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartPink(scale int) *Comp {
	c.El.Style(styles.BorderStartPink(scale))
	return c
}

func (c *Comp) BorderStartPinkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartPinkAlpha(scale))
	return c
}

func (c *Comp) BorderStartPinkDark(scale int) *Comp {
	c.El.Style(styles.BorderStartPinkDark(scale))
	return c
}

func (c *Comp) BorderStartPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartPinkDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartPlum(scale int) *Comp {
	c.El.Style(styles.BorderStartPlum(scale))
	return c
}

func (c *Comp) BorderStartPlumAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartPlumAlpha(scale))
	return c
}

func (c *Comp) BorderStartPlumDark(scale int) *Comp {
	c.El.Style(styles.BorderStartPlumDark(scale))
	return c
}

func (c *Comp) BorderStartPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartPlumDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartPurple(scale int) *Comp {
	c.El.Style(styles.BorderStartPurple(scale))
	return c
}

func (c *Comp) BorderStartPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartPurpleAlpha(scale))
	return c
}

func (c *Comp) BorderStartPurpleDark(scale int) *Comp {
	c.El.Style(styles.BorderStartPurpleDark(scale))
	return c
}

func (c *Comp) BorderStartPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartRed(scale int) *Comp {
	c.El.Style(styles.BorderStartRed(scale))
	return c
}

func (c *Comp) BorderStartRedAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartRedAlpha(scale))
	return c
}

func (c *Comp) BorderStartRedDark(scale int) *Comp {
	c.El.Style(styles.BorderStartRedDark(scale))
	return c
}

func (c *Comp) BorderStartRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartRedDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartRuby(scale int) *Comp {
	c.El.Style(styles.BorderStartRuby(scale))
	return c
}

func (c *Comp) BorderStartRubyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartRubyAlpha(scale))
	return c
}

func (c *Comp) BorderStartRubyDark(scale int) *Comp {
	c.El.Style(styles.BorderStartRubyDark(scale))
	return c
}

func (c *Comp) BorderStartRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartRubyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartSage(scale int) *Comp {
	c.El.Style(styles.BorderStartSage(scale))
	return c
}

func (c *Comp) BorderStartSageAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartSageAlpha(scale))
	return c
}

func (c *Comp) BorderStartSageDark(scale int) *Comp {
	c.El.Style(styles.BorderStartSageDark(scale))
	return c
}

func (c *Comp) BorderStartSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartSageDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartSand(scale int) *Comp {
	c.El.Style(styles.BorderStartSand(scale))
	return c
}

func (c *Comp) BorderStartSandAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartSandAlpha(scale))
	return c
}

func (c *Comp) BorderStartSandDark(scale int) *Comp {
	c.El.Style(styles.BorderStartSandDark(scale))
	return c
}

func (c *Comp) BorderStartSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartSandDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartSky(scale int) *Comp {
	c.El.Style(styles.BorderStartSky(scale))
	return c
}

func (c *Comp) BorderStartSkyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartSkyAlpha(scale))
	return c
}

func (c *Comp) BorderStartSkyDark(scale int) *Comp {
	c.El.Style(styles.BorderStartSkyDark(scale))
	return c
}

func (c *Comp) BorderStartSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartSkyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartSlate(scale int) *Comp {
	c.El.Style(styles.BorderStartSlate(scale))
	return c
}

func (c *Comp) BorderStartSlateAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartSlateAlpha(scale))
	return c
}

func (c *Comp) BorderStartSlateDark(scale int) *Comp {
	c.El.Style(styles.BorderStartSlateDark(scale))
	return c
}

func (c *Comp) BorderStartSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartSlateDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartTeal(scale int) *Comp {
	c.El.Style(styles.BorderStartTeal(scale))
	return c
}

func (c *Comp) BorderStartTealAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartTealAlpha(scale))
	return c
}

func (c *Comp) BorderStartTealDark(scale int) *Comp {
	c.El.Style(styles.BorderStartTealDark(scale))
	return c
}

func (c *Comp) BorderStartTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartTealDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartTomato(scale int) *Comp {
	c.El.Style(styles.BorderStartTomato(scale))
	return c
}

func (c *Comp) BorderStartTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartTomatoAlpha(scale))
	return c
}

func (c *Comp) BorderStartTomatoDark(scale int) *Comp {
	c.El.Style(styles.BorderStartTomatoDark(scale))
	return c
}

func (c *Comp) BorderStartTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartTransparent() *Comp {
	c.El.Style(styles.BorderStartTransparent())
	return c
}

func (c *Comp) BorderStartViolet(scale int) *Comp {
	c.El.Style(styles.BorderStartViolet(scale))
	return c
}

func (c *Comp) BorderStartVioletAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartVioletAlpha(scale))
	return c
}

func (c *Comp) BorderStartVioletDark(scale int) *Comp {
	c.El.Style(styles.BorderStartVioletDark(scale))
	return c
}

func (c *Comp) BorderStartVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartVioletDarkAlpha(scale))
	return c
}

func (c *Comp) BorderStartWhite() *Comp {
	c.El.Style(styles.BorderStartWhite())
	return c
}

func (c *Comp) BorderStartWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartWhiteAlpha(scale))
	return c
}

func (c *Comp) BorderStartYellow(scale int) *Comp {
	c.El.Style(styles.BorderStartYellow(scale))
	return c
}

func (c *Comp) BorderStartYellowAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartYellowAlpha(scale))
	return c
}

func (c *Comp) BorderStartYellowDark(scale int) *Comp {
	c.El.Style(styles.BorderStartYellowDark(scale))
	return c
}

func (c *Comp) BorderStartYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderStartYellowDarkAlpha(scale))
	return c
}

func (c *Comp) BorderT(val ...value.Value) *Comp {
	c.El.Style(styles.BorderT(val...))
	return c
}

func (c *Comp) BorderTeal(scale int) *Comp {
	c.El.Style(styles.BorderTeal(scale))
	return c
}

func (c *Comp) BorderTealAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTealAlpha(scale))
	return c
}

func (c *Comp) BorderTealDark(scale int) *Comp {
	c.El.Style(styles.BorderTealDark(scale))
	return c
}

func (c *Comp) BorderTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTealDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTomato(scale int) *Comp {
	c.El.Style(styles.BorderTomato(scale))
	return c
}

func (c *Comp) BorderTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTomatoAlpha(scale))
	return c
}

func (c *Comp) BorderTomatoDark(scale int) *Comp {
	c.El.Style(styles.BorderTomatoDark(scale))
	return c
}

func (c *Comp) BorderTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopAmber(scale int) *Comp {
	c.El.Style(styles.BorderTopAmber(scale))
	return c
}

func (c *Comp) BorderTopAmberAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopAmberAlpha(scale))
	return c
}

func (c *Comp) BorderTopAmberDark(scale int) *Comp {
	c.El.Style(styles.BorderTopAmberDark(scale))
	return c
}

func (c *Comp) BorderTopAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopAmberDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopBlack() *Comp {
	c.El.Style(styles.BorderTopBlack())
	return c
}

func (c *Comp) BorderTopBlackAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopBlackAlpha(scale))
	return c
}

func (c *Comp) BorderTopBlue(scale int) *Comp {
	c.El.Style(styles.BorderTopBlue(scale))
	return c
}

func (c *Comp) BorderTopBlueAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopBlueAlpha(scale))
	return c
}

func (c *Comp) BorderTopBlueDark(scale int) *Comp {
	c.El.Style(styles.BorderTopBlueDark(scale))
	return c
}

func (c *Comp) BorderTopBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopBlueDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopBronze(scale int) *Comp {
	c.El.Style(styles.BorderTopBronze(scale))
	return c
}

func (c *Comp) BorderTopBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopBronzeAlpha(scale))
	return c
}

func (c *Comp) BorderTopBronzeDark(scale int) *Comp {
	c.El.Style(styles.BorderTopBronzeDark(scale))
	return c
}

func (c *Comp) BorderTopBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopBrown(scale int) *Comp {
	c.El.Style(styles.BorderTopBrown(scale))
	return c
}

func (c *Comp) BorderTopBrownAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopBrownAlpha(scale))
	return c
}

func (c *Comp) BorderTopBrownDark(scale int) *Comp {
	c.El.Style(styles.BorderTopBrownDark(scale))
	return c
}

func (c *Comp) BorderTopBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopBrownDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopColor(val value.Value) *Comp {
	c.El.Style(styles.BorderTopColor(val))
	return c
}

func (c *Comp) BorderTopCrimson(scale int) *Comp {
	c.El.Style(styles.BorderTopCrimson(scale))
	return c
}

func (c *Comp) BorderTopCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopCrimsonAlpha(scale))
	return c
}

func (c *Comp) BorderTopCrimsonDark(scale int) *Comp {
	c.El.Style(styles.BorderTopCrimsonDark(scale))
	return c
}

func (c *Comp) BorderTopCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopCurrent() *Comp {
	c.El.Style(styles.BorderTopCurrent())
	return c
}

func (c *Comp) BorderTopCyan(scale int) *Comp {
	c.El.Style(styles.BorderTopCyan(scale))
	return c
}

func (c *Comp) BorderTopCyanAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopCyanAlpha(scale))
	return c
}

func (c *Comp) BorderTopCyanDark(scale int) *Comp {
	c.El.Style(styles.BorderTopCyanDark(scale))
	return c
}

func (c *Comp) BorderTopCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopCyanDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopGold(scale int) *Comp {
	c.El.Style(styles.BorderTopGold(scale))
	return c
}

func (c *Comp) BorderTopGoldAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopGoldAlpha(scale))
	return c
}

func (c *Comp) BorderTopGoldDark(scale int) *Comp {
	c.El.Style(styles.BorderTopGoldDark(scale))
	return c
}

func (c *Comp) BorderTopGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopGoldDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopGrass(scale int) *Comp {
	c.El.Style(styles.BorderTopGrass(scale))
	return c
}

func (c *Comp) BorderTopGrassAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopGrassAlpha(scale))
	return c
}

func (c *Comp) BorderTopGrassDark(scale int) *Comp {
	c.El.Style(styles.BorderTopGrassDark(scale))
	return c
}

func (c *Comp) BorderTopGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopGrassDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopGray(scale int) *Comp {
	c.El.Style(styles.BorderTopGray(scale))
	return c
}

func (c *Comp) BorderTopGrayAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopGrayAlpha(scale))
	return c
}

func (c *Comp) BorderTopGrayDark(scale int) *Comp {
	c.El.Style(styles.BorderTopGrayDark(scale))
	return c
}

func (c *Comp) BorderTopGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopGrayDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopGreen(scale int) *Comp {
	c.El.Style(styles.BorderTopGreen(scale))
	return c
}

func (c *Comp) BorderTopGreenAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopGreenAlpha(scale))
	return c
}

func (c *Comp) BorderTopGreenDark(scale int) *Comp {
	c.El.Style(styles.BorderTopGreenDark(scale))
	return c
}

func (c *Comp) BorderTopGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopGreenDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopIndigo(scale int) *Comp {
	c.El.Style(styles.BorderTopIndigo(scale))
	return c
}

func (c *Comp) BorderTopIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopIndigoAlpha(scale))
	return c
}

func (c *Comp) BorderTopIndigoDark(scale int) *Comp {
	c.El.Style(styles.BorderTopIndigoDark(scale))
	return c
}

func (c *Comp) BorderTopIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopInherit() *Comp {
	c.El.Style(styles.BorderTopInherit())
	return c
}

func (c *Comp) BorderTopIris(scale int) *Comp {
	c.El.Style(styles.BorderTopIris(scale))
	return c
}

func (c *Comp) BorderTopIrisAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopIrisAlpha(scale))
	return c
}

func (c *Comp) BorderTopIrisDark(scale int) *Comp {
	c.El.Style(styles.BorderTopIrisDark(scale))
	return c
}

func (c *Comp) BorderTopIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopIrisDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopJade(scale int) *Comp {
	c.El.Style(styles.BorderTopJade(scale))
	return c
}

func (c *Comp) BorderTopJadeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopJadeAlpha(scale))
	return c
}

func (c *Comp) BorderTopJadeDark(scale int) *Comp {
	c.El.Style(styles.BorderTopJadeDark(scale))
	return c
}

func (c *Comp) BorderTopJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopJadeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopLime(scale int) *Comp {
	c.El.Style(styles.BorderTopLime(scale))
	return c
}

func (c *Comp) BorderTopLimeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopLimeAlpha(scale))
	return c
}

func (c *Comp) BorderTopLimeDark(scale int) *Comp {
	c.El.Style(styles.BorderTopLimeDark(scale))
	return c
}

func (c *Comp) BorderTopLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopLimeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopMauve(scale int) *Comp {
	c.El.Style(styles.BorderTopMauve(scale))
	return c
}

func (c *Comp) BorderTopMauveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopMauveAlpha(scale))
	return c
}

func (c *Comp) BorderTopMauveDark(scale int) *Comp {
	c.El.Style(styles.BorderTopMauveDark(scale))
	return c
}

func (c *Comp) BorderTopMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopMauveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopMint(scale int) *Comp {
	c.El.Style(styles.BorderTopMint(scale))
	return c
}

func (c *Comp) BorderTopMintAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopMintAlpha(scale))
	return c
}

func (c *Comp) BorderTopMintDark(scale int) *Comp {
	c.El.Style(styles.BorderTopMintDark(scale))
	return c
}

func (c *Comp) BorderTopMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopMintDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopOlive(scale int) *Comp {
	c.El.Style(styles.BorderTopOlive(scale))
	return c
}

func (c *Comp) BorderTopOliveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopOliveAlpha(scale))
	return c
}

func (c *Comp) BorderTopOliveDark(scale int) *Comp {
	c.El.Style(styles.BorderTopOliveDark(scale))
	return c
}

func (c *Comp) BorderTopOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopOliveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopOrange(scale int) *Comp {
	c.El.Style(styles.BorderTopOrange(scale))
	return c
}

func (c *Comp) BorderTopOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopOrangeAlpha(scale))
	return c
}

func (c *Comp) BorderTopOrangeDark(scale int) *Comp {
	c.El.Style(styles.BorderTopOrangeDark(scale))
	return c
}

func (c *Comp) BorderTopOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopPink(scale int) *Comp {
	c.El.Style(styles.BorderTopPink(scale))
	return c
}

func (c *Comp) BorderTopPinkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopPinkAlpha(scale))
	return c
}

func (c *Comp) BorderTopPinkDark(scale int) *Comp {
	c.El.Style(styles.BorderTopPinkDark(scale))
	return c
}

func (c *Comp) BorderTopPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopPinkDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopPlum(scale int) *Comp {
	c.El.Style(styles.BorderTopPlum(scale))
	return c
}

func (c *Comp) BorderTopPlumAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopPlumAlpha(scale))
	return c
}

func (c *Comp) BorderTopPlumDark(scale int) *Comp {
	c.El.Style(styles.BorderTopPlumDark(scale))
	return c
}

func (c *Comp) BorderTopPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopPlumDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopPurple(scale int) *Comp {
	c.El.Style(styles.BorderTopPurple(scale))
	return c
}

func (c *Comp) BorderTopPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopPurpleAlpha(scale))
	return c
}

func (c *Comp) BorderTopPurpleDark(scale int) *Comp {
	c.El.Style(styles.BorderTopPurpleDark(scale))
	return c
}

func (c *Comp) BorderTopPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopRed(scale int) *Comp {
	c.El.Style(styles.BorderTopRed(scale))
	return c
}

func (c *Comp) BorderTopRedAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopRedAlpha(scale))
	return c
}

func (c *Comp) BorderTopRedDark(scale int) *Comp {
	c.El.Style(styles.BorderTopRedDark(scale))
	return c
}

func (c *Comp) BorderTopRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopRedDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopRuby(scale int) *Comp {
	c.El.Style(styles.BorderTopRuby(scale))
	return c
}

func (c *Comp) BorderTopRubyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopRubyAlpha(scale))
	return c
}

func (c *Comp) BorderTopRubyDark(scale int) *Comp {
	c.El.Style(styles.BorderTopRubyDark(scale))
	return c
}

func (c *Comp) BorderTopRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopRubyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopSage(scale int) *Comp {
	c.El.Style(styles.BorderTopSage(scale))
	return c
}

func (c *Comp) BorderTopSageAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopSageAlpha(scale))
	return c
}

func (c *Comp) BorderTopSageDark(scale int) *Comp {
	c.El.Style(styles.BorderTopSageDark(scale))
	return c
}

func (c *Comp) BorderTopSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopSageDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopSand(scale int) *Comp {
	c.El.Style(styles.BorderTopSand(scale))
	return c
}

func (c *Comp) BorderTopSandAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopSandAlpha(scale))
	return c
}

func (c *Comp) BorderTopSandDark(scale int) *Comp {
	c.El.Style(styles.BorderTopSandDark(scale))
	return c
}

func (c *Comp) BorderTopSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopSandDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopSky(scale int) *Comp {
	c.El.Style(styles.BorderTopSky(scale))
	return c
}

func (c *Comp) BorderTopSkyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopSkyAlpha(scale))
	return c
}

func (c *Comp) BorderTopSkyDark(scale int) *Comp {
	c.El.Style(styles.BorderTopSkyDark(scale))
	return c
}

func (c *Comp) BorderTopSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopSkyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopSlate(scale int) *Comp {
	c.El.Style(styles.BorderTopSlate(scale))
	return c
}

func (c *Comp) BorderTopSlateAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopSlateAlpha(scale))
	return c
}

func (c *Comp) BorderTopSlateDark(scale int) *Comp {
	c.El.Style(styles.BorderTopSlateDark(scale))
	return c
}

func (c *Comp) BorderTopSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopSlateDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopTeal(scale int) *Comp {
	c.El.Style(styles.BorderTopTeal(scale))
	return c
}

func (c *Comp) BorderTopTealAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopTealAlpha(scale))
	return c
}

func (c *Comp) BorderTopTealDark(scale int) *Comp {
	c.El.Style(styles.BorderTopTealDark(scale))
	return c
}

func (c *Comp) BorderTopTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopTealDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopTomato(scale int) *Comp {
	c.El.Style(styles.BorderTopTomato(scale))
	return c
}

func (c *Comp) BorderTopTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopTomatoAlpha(scale))
	return c
}

func (c *Comp) BorderTopTomatoDark(scale int) *Comp {
	c.El.Style(styles.BorderTopTomatoDark(scale))
	return c
}

func (c *Comp) BorderTopTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopTransparent() *Comp {
	c.El.Style(styles.BorderTopTransparent())
	return c
}

func (c *Comp) BorderTopViolet(scale int) *Comp {
	c.El.Style(styles.BorderTopViolet(scale))
	return c
}

func (c *Comp) BorderTopVioletAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopVioletAlpha(scale))
	return c
}

func (c *Comp) BorderTopVioletDark(scale int) *Comp {
	c.El.Style(styles.BorderTopVioletDark(scale))
	return c
}

func (c *Comp) BorderTopVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopVioletDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTopWhite() *Comp {
	c.El.Style(styles.BorderTopWhite())
	return c
}

func (c *Comp) BorderTopWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopWhiteAlpha(scale))
	return c
}

func (c *Comp) BorderTopYellow(scale int) *Comp {
	c.El.Style(styles.BorderTopYellow(scale))
	return c
}

func (c *Comp) BorderTopYellowAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopYellowAlpha(scale))
	return c
}

func (c *Comp) BorderTopYellowDark(scale int) *Comp {
	c.El.Style(styles.BorderTopYellowDark(scale))
	return c
}

func (c *Comp) BorderTopYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderTopYellowDarkAlpha(scale))
	return c
}

func (c *Comp) BorderTransparent() *Comp {
	c.El.Style(styles.BorderTransparent())
	return c
}

func (c *Comp) BorderViolet(scale int) *Comp {
	c.El.Style(styles.BorderViolet(scale))
	return c
}

func (c *Comp) BorderVioletAlpha(scale int) *Comp {
	c.El.Style(styles.BorderVioletAlpha(scale))
	return c
}

func (c *Comp) BorderVioletDark(scale int) *Comp {
	c.El.Style(styles.BorderVioletDark(scale))
	return c
}

func (c *Comp) BorderVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderVioletDarkAlpha(scale))
	return c
}

func (c *Comp) BorderWhite() *Comp {
	c.El.Style(styles.BorderWhite())
	return c
}

func (c *Comp) BorderWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.BorderWhiteAlpha(scale))
	return c
}

func (c *Comp) BorderX(val ...value.Value) *Comp {
	c.El.Style(styles.BorderX(val...))
	return c
}

func (c *Comp) BorderXAmber(scale int) *Comp {
	c.El.Style(styles.BorderXAmber(scale))
	return c
}

func (c *Comp) BorderXAmberAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXAmberAlpha(scale))
	return c
}

func (c *Comp) BorderXAmberDark(scale int) *Comp {
	c.El.Style(styles.BorderXAmberDark(scale))
	return c
}

func (c *Comp) BorderXAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXAmberDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXBlack() *Comp {
	c.El.Style(styles.BorderXBlack())
	return c
}

func (c *Comp) BorderXBlackAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXBlackAlpha(scale))
	return c
}

func (c *Comp) BorderXBlue(scale int) *Comp {
	c.El.Style(styles.BorderXBlue(scale))
	return c
}

func (c *Comp) BorderXBlueAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXBlueAlpha(scale))
	return c
}

func (c *Comp) BorderXBlueDark(scale int) *Comp {
	c.El.Style(styles.BorderXBlueDark(scale))
	return c
}

func (c *Comp) BorderXBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXBlueDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXBronze(scale int) *Comp {
	c.El.Style(styles.BorderXBronze(scale))
	return c
}

func (c *Comp) BorderXBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXBronzeAlpha(scale))
	return c
}

func (c *Comp) BorderXBronzeDark(scale int) *Comp {
	c.El.Style(styles.BorderXBronzeDark(scale))
	return c
}

func (c *Comp) BorderXBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXBrown(scale int) *Comp {
	c.El.Style(styles.BorderXBrown(scale))
	return c
}

func (c *Comp) BorderXBrownAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXBrownAlpha(scale))
	return c
}

func (c *Comp) BorderXBrownDark(scale int) *Comp {
	c.El.Style(styles.BorderXBrownDark(scale))
	return c
}

func (c *Comp) BorderXBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXBrownDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXColor(val value.Value) *Comp {
	c.El.Style(styles.BorderXColor(val))
	return c
}

func (c *Comp) BorderXCrimson(scale int) *Comp {
	c.El.Style(styles.BorderXCrimson(scale))
	return c
}

func (c *Comp) BorderXCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXCrimsonAlpha(scale))
	return c
}

func (c *Comp) BorderXCrimsonDark(scale int) *Comp {
	c.El.Style(styles.BorderXCrimsonDark(scale))
	return c
}

func (c *Comp) BorderXCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXCurrent() *Comp {
	c.El.Style(styles.BorderXCurrent())
	return c
}

func (c *Comp) BorderXCyan(scale int) *Comp {
	c.El.Style(styles.BorderXCyan(scale))
	return c
}

func (c *Comp) BorderXCyanAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXCyanAlpha(scale))
	return c
}

func (c *Comp) BorderXCyanDark(scale int) *Comp {
	c.El.Style(styles.BorderXCyanDark(scale))
	return c
}

func (c *Comp) BorderXCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXCyanDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXGold(scale int) *Comp {
	c.El.Style(styles.BorderXGold(scale))
	return c
}

func (c *Comp) BorderXGoldAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXGoldAlpha(scale))
	return c
}

func (c *Comp) BorderXGoldDark(scale int) *Comp {
	c.El.Style(styles.BorderXGoldDark(scale))
	return c
}

func (c *Comp) BorderXGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXGoldDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXGrass(scale int) *Comp {
	c.El.Style(styles.BorderXGrass(scale))
	return c
}

func (c *Comp) BorderXGrassAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXGrassAlpha(scale))
	return c
}

func (c *Comp) BorderXGrassDark(scale int) *Comp {
	c.El.Style(styles.BorderXGrassDark(scale))
	return c
}

func (c *Comp) BorderXGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXGrassDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXGray(scale int) *Comp {
	c.El.Style(styles.BorderXGray(scale))
	return c
}

func (c *Comp) BorderXGrayAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXGrayAlpha(scale))
	return c
}

func (c *Comp) BorderXGrayDark(scale int) *Comp {
	c.El.Style(styles.BorderXGrayDark(scale))
	return c
}

func (c *Comp) BorderXGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXGrayDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXGreen(scale int) *Comp {
	c.El.Style(styles.BorderXGreen(scale))
	return c
}

func (c *Comp) BorderXGreenAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXGreenAlpha(scale))
	return c
}

func (c *Comp) BorderXGreenDark(scale int) *Comp {
	c.El.Style(styles.BorderXGreenDark(scale))
	return c
}

func (c *Comp) BorderXGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXGreenDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXIndigo(scale int) *Comp {
	c.El.Style(styles.BorderXIndigo(scale))
	return c
}

func (c *Comp) BorderXIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXIndigoAlpha(scale))
	return c
}

func (c *Comp) BorderXIndigoDark(scale int) *Comp {
	c.El.Style(styles.BorderXIndigoDark(scale))
	return c
}

func (c *Comp) BorderXIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXInherit() *Comp {
	c.El.Style(styles.BorderXInherit())
	return c
}

func (c *Comp) BorderXIris(scale int) *Comp {
	c.El.Style(styles.BorderXIris(scale))
	return c
}

func (c *Comp) BorderXIrisAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXIrisAlpha(scale))
	return c
}

func (c *Comp) BorderXIrisDark(scale int) *Comp {
	c.El.Style(styles.BorderXIrisDark(scale))
	return c
}

func (c *Comp) BorderXIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXIrisDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXJade(scale int) *Comp {
	c.El.Style(styles.BorderXJade(scale))
	return c
}

func (c *Comp) BorderXJadeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXJadeAlpha(scale))
	return c
}

func (c *Comp) BorderXJadeDark(scale int) *Comp {
	c.El.Style(styles.BorderXJadeDark(scale))
	return c
}

func (c *Comp) BorderXJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXJadeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXLime(scale int) *Comp {
	c.El.Style(styles.BorderXLime(scale))
	return c
}

func (c *Comp) BorderXLimeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXLimeAlpha(scale))
	return c
}

func (c *Comp) BorderXLimeDark(scale int) *Comp {
	c.El.Style(styles.BorderXLimeDark(scale))
	return c
}

func (c *Comp) BorderXLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXLimeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXMauve(scale int) *Comp {
	c.El.Style(styles.BorderXMauve(scale))
	return c
}

func (c *Comp) BorderXMauveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXMauveAlpha(scale))
	return c
}

func (c *Comp) BorderXMauveDark(scale int) *Comp {
	c.El.Style(styles.BorderXMauveDark(scale))
	return c
}

func (c *Comp) BorderXMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXMauveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXMint(scale int) *Comp {
	c.El.Style(styles.BorderXMint(scale))
	return c
}

func (c *Comp) BorderXMintAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXMintAlpha(scale))
	return c
}

func (c *Comp) BorderXMintDark(scale int) *Comp {
	c.El.Style(styles.BorderXMintDark(scale))
	return c
}

func (c *Comp) BorderXMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXMintDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXOlive(scale int) *Comp {
	c.El.Style(styles.BorderXOlive(scale))
	return c
}

func (c *Comp) BorderXOliveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXOliveAlpha(scale))
	return c
}

func (c *Comp) BorderXOliveDark(scale int) *Comp {
	c.El.Style(styles.BorderXOliveDark(scale))
	return c
}

func (c *Comp) BorderXOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXOliveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXOrange(scale int) *Comp {
	c.El.Style(styles.BorderXOrange(scale))
	return c
}

func (c *Comp) BorderXOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXOrangeAlpha(scale))
	return c
}

func (c *Comp) BorderXOrangeDark(scale int) *Comp {
	c.El.Style(styles.BorderXOrangeDark(scale))
	return c
}

func (c *Comp) BorderXOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXPink(scale int) *Comp {
	c.El.Style(styles.BorderXPink(scale))
	return c
}

func (c *Comp) BorderXPinkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXPinkAlpha(scale))
	return c
}

func (c *Comp) BorderXPinkDark(scale int) *Comp {
	c.El.Style(styles.BorderXPinkDark(scale))
	return c
}

func (c *Comp) BorderXPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXPinkDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXPlum(scale int) *Comp {
	c.El.Style(styles.BorderXPlum(scale))
	return c
}

func (c *Comp) BorderXPlumAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXPlumAlpha(scale))
	return c
}

func (c *Comp) BorderXPlumDark(scale int) *Comp {
	c.El.Style(styles.BorderXPlumDark(scale))
	return c
}

func (c *Comp) BorderXPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXPlumDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXPurple(scale int) *Comp {
	c.El.Style(styles.BorderXPurple(scale))
	return c
}

func (c *Comp) BorderXPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXPurpleAlpha(scale))
	return c
}

func (c *Comp) BorderXPurpleDark(scale int) *Comp {
	c.El.Style(styles.BorderXPurpleDark(scale))
	return c
}

func (c *Comp) BorderXPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXRed(scale int) *Comp {
	c.El.Style(styles.BorderXRed(scale))
	return c
}

func (c *Comp) BorderXRedAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXRedAlpha(scale))
	return c
}

func (c *Comp) BorderXRedDark(scale int) *Comp {
	c.El.Style(styles.BorderXRedDark(scale))
	return c
}

func (c *Comp) BorderXRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXRedDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXRuby(scale int) *Comp {
	c.El.Style(styles.BorderXRuby(scale))
	return c
}

func (c *Comp) BorderXRubyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXRubyAlpha(scale))
	return c
}

func (c *Comp) BorderXRubyDark(scale int) *Comp {
	c.El.Style(styles.BorderXRubyDark(scale))
	return c
}

func (c *Comp) BorderXRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXRubyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXSage(scale int) *Comp {
	c.El.Style(styles.BorderXSage(scale))
	return c
}

func (c *Comp) BorderXSageAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXSageAlpha(scale))
	return c
}

func (c *Comp) BorderXSageDark(scale int) *Comp {
	c.El.Style(styles.BorderXSageDark(scale))
	return c
}

func (c *Comp) BorderXSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXSageDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXSand(scale int) *Comp {
	c.El.Style(styles.BorderXSand(scale))
	return c
}

func (c *Comp) BorderXSandAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXSandAlpha(scale))
	return c
}

func (c *Comp) BorderXSandDark(scale int) *Comp {
	c.El.Style(styles.BorderXSandDark(scale))
	return c
}

func (c *Comp) BorderXSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXSandDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXSky(scale int) *Comp {
	c.El.Style(styles.BorderXSky(scale))
	return c
}

func (c *Comp) BorderXSkyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXSkyAlpha(scale))
	return c
}

func (c *Comp) BorderXSkyDark(scale int) *Comp {
	c.El.Style(styles.BorderXSkyDark(scale))
	return c
}

func (c *Comp) BorderXSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXSkyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXSlate(scale int) *Comp {
	c.El.Style(styles.BorderXSlate(scale))
	return c
}

func (c *Comp) BorderXSlateAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXSlateAlpha(scale))
	return c
}

func (c *Comp) BorderXSlateDark(scale int) *Comp {
	c.El.Style(styles.BorderXSlateDark(scale))
	return c
}

func (c *Comp) BorderXSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXSlateDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXTeal(scale int) *Comp {
	c.El.Style(styles.BorderXTeal(scale))
	return c
}

func (c *Comp) BorderXTealAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXTealAlpha(scale))
	return c
}

func (c *Comp) BorderXTealDark(scale int) *Comp {
	c.El.Style(styles.BorderXTealDark(scale))
	return c
}

func (c *Comp) BorderXTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXTealDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXTomato(scale int) *Comp {
	c.El.Style(styles.BorderXTomato(scale))
	return c
}

func (c *Comp) BorderXTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXTomatoAlpha(scale))
	return c
}

func (c *Comp) BorderXTomatoDark(scale int) *Comp {
	c.El.Style(styles.BorderXTomatoDark(scale))
	return c
}

func (c *Comp) BorderXTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXTransparent() *Comp {
	c.El.Style(styles.BorderXTransparent())
	return c
}

func (c *Comp) BorderXViolet(scale int) *Comp {
	c.El.Style(styles.BorderXViolet(scale))
	return c
}

func (c *Comp) BorderXVioletAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXVioletAlpha(scale))
	return c
}

func (c *Comp) BorderXVioletDark(scale int) *Comp {
	c.El.Style(styles.BorderXVioletDark(scale))
	return c
}

func (c *Comp) BorderXVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXVioletDarkAlpha(scale))
	return c
}

func (c *Comp) BorderXWhite() *Comp {
	c.El.Style(styles.BorderXWhite())
	return c
}

func (c *Comp) BorderXWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXWhiteAlpha(scale))
	return c
}

func (c *Comp) BorderXYellow(scale int) *Comp {
	c.El.Style(styles.BorderXYellow(scale))
	return c
}

func (c *Comp) BorderXYellowAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXYellowAlpha(scale))
	return c
}

func (c *Comp) BorderXYellowDark(scale int) *Comp {
	c.El.Style(styles.BorderXYellowDark(scale))
	return c
}

func (c *Comp) BorderXYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderXYellowDarkAlpha(scale))
	return c
}

func (c *Comp) BorderY(val ...value.Value) *Comp {
	c.El.Style(styles.BorderY(val...))
	return c
}

func (c *Comp) BorderYAmber(scale int) *Comp {
	c.El.Style(styles.BorderYAmber(scale))
	return c
}

func (c *Comp) BorderYAmberAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYAmberAlpha(scale))
	return c
}

func (c *Comp) BorderYAmberDark(scale int) *Comp {
	c.El.Style(styles.BorderYAmberDark(scale))
	return c
}

func (c *Comp) BorderYAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYAmberDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYBlack() *Comp {
	c.El.Style(styles.BorderYBlack())
	return c
}

func (c *Comp) BorderYBlackAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYBlackAlpha(scale))
	return c
}

func (c *Comp) BorderYBlue(scale int) *Comp {
	c.El.Style(styles.BorderYBlue(scale))
	return c
}

func (c *Comp) BorderYBlueAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYBlueAlpha(scale))
	return c
}

func (c *Comp) BorderYBlueDark(scale int) *Comp {
	c.El.Style(styles.BorderYBlueDark(scale))
	return c
}

func (c *Comp) BorderYBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYBlueDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYBronze(scale int) *Comp {
	c.El.Style(styles.BorderYBronze(scale))
	return c
}

func (c *Comp) BorderYBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYBronzeAlpha(scale))
	return c
}

func (c *Comp) BorderYBronzeDark(scale int) *Comp {
	c.El.Style(styles.BorderYBronzeDark(scale))
	return c
}

func (c *Comp) BorderYBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYBrown(scale int) *Comp {
	c.El.Style(styles.BorderYBrown(scale))
	return c
}

func (c *Comp) BorderYBrownAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYBrownAlpha(scale))
	return c
}

func (c *Comp) BorderYBrownDark(scale int) *Comp {
	c.El.Style(styles.BorderYBrownDark(scale))
	return c
}

func (c *Comp) BorderYBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYBrownDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYColor(val value.Value) *Comp {
	c.El.Style(styles.BorderYColor(val))
	return c
}

func (c *Comp) BorderYCrimson(scale int) *Comp {
	c.El.Style(styles.BorderYCrimson(scale))
	return c
}

func (c *Comp) BorderYCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYCrimsonAlpha(scale))
	return c
}

func (c *Comp) BorderYCrimsonDark(scale int) *Comp {
	c.El.Style(styles.BorderYCrimsonDark(scale))
	return c
}

func (c *Comp) BorderYCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYCurrent() *Comp {
	c.El.Style(styles.BorderYCurrent())
	return c
}

func (c *Comp) BorderYCyan(scale int) *Comp {
	c.El.Style(styles.BorderYCyan(scale))
	return c
}

func (c *Comp) BorderYCyanAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYCyanAlpha(scale))
	return c
}

func (c *Comp) BorderYCyanDark(scale int) *Comp {
	c.El.Style(styles.BorderYCyanDark(scale))
	return c
}

func (c *Comp) BorderYCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYCyanDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYGold(scale int) *Comp {
	c.El.Style(styles.BorderYGold(scale))
	return c
}

func (c *Comp) BorderYGoldAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYGoldAlpha(scale))
	return c
}

func (c *Comp) BorderYGoldDark(scale int) *Comp {
	c.El.Style(styles.BorderYGoldDark(scale))
	return c
}

func (c *Comp) BorderYGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYGoldDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYGrass(scale int) *Comp {
	c.El.Style(styles.BorderYGrass(scale))
	return c
}

func (c *Comp) BorderYGrassAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYGrassAlpha(scale))
	return c
}

func (c *Comp) BorderYGrassDark(scale int) *Comp {
	c.El.Style(styles.BorderYGrassDark(scale))
	return c
}

func (c *Comp) BorderYGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYGrassDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYGray(scale int) *Comp {
	c.El.Style(styles.BorderYGray(scale))
	return c
}

func (c *Comp) BorderYGrayAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYGrayAlpha(scale))
	return c
}

func (c *Comp) BorderYGrayDark(scale int) *Comp {
	c.El.Style(styles.BorderYGrayDark(scale))
	return c
}

func (c *Comp) BorderYGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYGrayDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYGreen(scale int) *Comp {
	c.El.Style(styles.BorderYGreen(scale))
	return c
}

func (c *Comp) BorderYGreenAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYGreenAlpha(scale))
	return c
}

func (c *Comp) BorderYGreenDark(scale int) *Comp {
	c.El.Style(styles.BorderYGreenDark(scale))
	return c
}

func (c *Comp) BorderYGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYGreenDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYIndigo(scale int) *Comp {
	c.El.Style(styles.BorderYIndigo(scale))
	return c
}

func (c *Comp) BorderYIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYIndigoAlpha(scale))
	return c
}

func (c *Comp) BorderYIndigoDark(scale int) *Comp {
	c.El.Style(styles.BorderYIndigoDark(scale))
	return c
}

func (c *Comp) BorderYIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYInherit() *Comp {
	c.El.Style(styles.BorderYInherit())
	return c
}

func (c *Comp) BorderYIris(scale int) *Comp {
	c.El.Style(styles.BorderYIris(scale))
	return c
}

func (c *Comp) BorderYIrisAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYIrisAlpha(scale))
	return c
}

func (c *Comp) BorderYIrisDark(scale int) *Comp {
	c.El.Style(styles.BorderYIrisDark(scale))
	return c
}

func (c *Comp) BorderYIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYIrisDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYJade(scale int) *Comp {
	c.El.Style(styles.BorderYJade(scale))
	return c
}

func (c *Comp) BorderYJadeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYJadeAlpha(scale))
	return c
}

func (c *Comp) BorderYJadeDark(scale int) *Comp {
	c.El.Style(styles.BorderYJadeDark(scale))
	return c
}

func (c *Comp) BorderYJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYJadeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYLime(scale int) *Comp {
	c.El.Style(styles.BorderYLime(scale))
	return c
}

func (c *Comp) BorderYLimeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYLimeAlpha(scale))
	return c
}

func (c *Comp) BorderYLimeDark(scale int) *Comp {
	c.El.Style(styles.BorderYLimeDark(scale))
	return c
}

func (c *Comp) BorderYLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYLimeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYMauve(scale int) *Comp {
	c.El.Style(styles.BorderYMauve(scale))
	return c
}

func (c *Comp) BorderYMauveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYMauveAlpha(scale))
	return c
}

func (c *Comp) BorderYMauveDark(scale int) *Comp {
	c.El.Style(styles.BorderYMauveDark(scale))
	return c
}

func (c *Comp) BorderYMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYMauveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYMint(scale int) *Comp {
	c.El.Style(styles.BorderYMint(scale))
	return c
}

func (c *Comp) BorderYMintAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYMintAlpha(scale))
	return c
}

func (c *Comp) BorderYMintDark(scale int) *Comp {
	c.El.Style(styles.BorderYMintDark(scale))
	return c
}

func (c *Comp) BorderYMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYMintDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYOlive(scale int) *Comp {
	c.El.Style(styles.BorderYOlive(scale))
	return c
}

func (c *Comp) BorderYOliveAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYOliveAlpha(scale))
	return c
}

func (c *Comp) BorderYOliveDark(scale int) *Comp {
	c.El.Style(styles.BorderYOliveDark(scale))
	return c
}

func (c *Comp) BorderYOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYOliveDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYOrange(scale int) *Comp {
	c.El.Style(styles.BorderYOrange(scale))
	return c
}

func (c *Comp) BorderYOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYOrangeAlpha(scale))
	return c
}

func (c *Comp) BorderYOrangeDark(scale int) *Comp {
	c.El.Style(styles.BorderYOrangeDark(scale))
	return c
}

func (c *Comp) BorderYOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYPink(scale int) *Comp {
	c.El.Style(styles.BorderYPink(scale))
	return c
}

func (c *Comp) BorderYPinkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYPinkAlpha(scale))
	return c
}

func (c *Comp) BorderYPinkDark(scale int) *Comp {
	c.El.Style(styles.BorderYPinkDark(scale))
	return c
}

func (c *Comp) BorderYPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYPinkDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYPlum(scale int) *Comp {
	c.El.Style(styles.BorderYPlum(scale))
	return c
}

func (c *Comp) BorderYPlumAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYPlumAlpha(scale))
	return c
}

func (c *Comp) BorderYPlumDark(scale int) *Comp {
	c.El.Style(styles.BorderYPlumDark(scale))
	return c
}

func (c *Comp) BorderYPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYPlumDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYPurple(scale int) *Comp {
	c.El.Style(styles.BorderYPurple(scale))
	return c
}

func (c *Comp) BorderYPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYPurpleAlpha(scale))
	return c
}

func (c *Comp) BorderYPurpleDark(scale int) *Comp {
	c.El.Style(styles.BorderYPurpleDark(scale))
	return c
}

func (c *Comp) BorderYPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYRed(scale int) *Comp {
	c.El.Style(styles.BorderYRed(scale))
	return c
}

func (c *Comp) BorderYRedAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYRedAlpha(scale))
	return c
}

func (c *Comp) BorderYRedDark(scale int) *Comp {
	c.El.Style(styles.BorderYRedDark(scale))
	return c
}

func (c *Comp) BorderYRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYRedDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYRuby(scale int) *Comp {
	c.El.Style(styles.BorderYRuby(scale))
	return c
}

func (c *Comp) BorderYRubyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYRubyAlpha(scale))
	return c
}

func (c *Comp) BorderYRubyDark(scale int) *Comp {
	c.El.Style(styles.BorderYRubyDark(scale))
	return c
}

func (c *Comp) BorderYRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYRubyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYSage(scale int) *Comp {
	c.El.Style(styles.BorderYSage(scale))
	return c
}

func (c *Comp) BorderYSageAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYSageAlpha(scale))
	return c
}

func (c *Comp) BorderYSageDark(scale int) *Comp {
	c.El.Style(styles.BorderYSageDark(scale))
	return c
}

func (c *Comp) BorderYSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYSageDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYSand(scale int) *Comp {
	c.El.Style(styles.BorderYSand(scale))
	return c
}

func (c *Comp) BorderYSandAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYSandAlpha(scale))
	return c
}

func (c *Comp) BorderYSandDark(scale int) *Comp {
	c.El.Style(styles.BorderYSandDark(scale))
	return c
}

func (c *Comp) BorderYSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYSandDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYSky(scale int) *Comp {
	c.El.Style(styles.BorderYSky(scale))
	return c
}

func (c *Comp) BorderYSkyAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYSkyAlpha(scale))
	return c
}

func (c *Comp) BorderYSkyDark(scale int) *Comp {
	c.El.Style(styles.BorderYSkyDark(scale))
	return c
}

func (c *Comp) BorderYSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYSkyDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYSlate(scale int) *Comp {
	c.El.Style(styles.BorderYSlate(scale))
	return c
}

func (c *Comp) BorderYSlateAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYSlateAlpha(scale))
	return c
}

func (c *Comp) BorderYSlateDark(scale int) *Comp {
	c.El.Style(styles.BorderYSlateDark(scale))
	return c
}

func (c *Comp) BorderYSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYSlateDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYTeal(scale int) *Comp {
	c.El.Style(styles.BorderYTeal(scale))
	return c
}

func (c *Comp) BorderYTealAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYTealAlpha(scale))
	return c
}

func (c *Comp) BorderYTealDark(scale int) *Comp {
	c.El.Style(styles.BorderYTealDark(scale))
	return c
}

func (c *Comp) BorderYTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYTealDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYTomato(scale int) *Comp {
	c.El.Style(styles.BorderYTomato(scale))
	return c
}

func (c *Comp) BorderYTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYTomatoAlpha(scale))
	return c
}

func (c *Comp) BorderYTomatoDark(scale int) *Comp {
	c.El.Style(styles.BorderYTomatoDark(scale))
	return c
}

func (c *Comp) BorderYTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYTransparent() *Comp {
	c.El.Style(styles.BorderYTransparent())
	return c
}

func (c *Comp) BorderYViolet(scale int) *Comp {
	c.El.Style(styles.BorderYViolet(scale))
	return c
}

func (c *Comp) BorderYVioletAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYVioletAlpha(scale))
	return c
}

func (c *Comp) BorderYVioletDark(scale int) *Comp {
	c.El.Style(styles.BorderYVioletDark(scale))
	return c
}

func (c *Comp) BorderYVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYVioletDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYWhite() *Comp {
	c.El.Style(styles.BorderYWhite())
	return c
}

func (c *Comp) BorderYWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYWhiteAlpha(scale))
	return c
}

func (c *Comp) BorderYYellow(scale int) *Comp {
	c.El.Style(styles.BorderYYellow(scale))
	return c
}

func (c *Comp) BorderYYellowAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYYellowAlpha(scale))
	return c
}

func (c *Comp) BorderYYellowDark(scale int) *Comp {
	c.El.Style(styles.BorderYYellowDark(scale))
	return c
}

func (c *Comp) BorderYYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYYellowDarkAlpha(scale))
	return c
}

func (c *Comp) BorderYellow(scale int) *Comp {
	c.El.Style(styles.BorderYellow(scale))
	return c
}

func (c *Comp) BorderYellowAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYellowAlpha(scale))
	return c
}

func (c *Comp) BorderYellowDark(scale int) *Comp {
	c.El.Style(styles.BorderYellowDark(scale))
	return c
}

func (c *Comp) BorderYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.BorderYellowDarkAlpha(scale))
	return c
}

func (c *Comp) Bottom(number int) *Comp {
	c.El.Style(styles.Bottom(number))
	return c
}

func (c *Comp) BottomAuto() *Comp {
	c.El.Style(styles.BottomAuto())
	return c
}

func (c *Comp) BottomBy(val value.Value) *Comp {
	c.El.Style(styles.BottomBy(val))
	return c
}

func (c *Comp) BottomFraction(fraction float64) *Comp {
	c.El.Style(styles.BottomFraction(fraction))
	return c
}

func (c *Comp) BottomFull() *Comp {
	c.El.Style(styles.BottomFull())
	return c
}

func (c *Comp) BottomPx() *Comp {
	c.El.Style(styles.BottomPx())
	return c
}

func (c *Comp) BoxBorder() *Comp {
	c.El.Style(styles.BoxBorder())
	return c
}

func (c *Comp) BoxContent() *Comp {
	c.El.Style(styles.BoxContent())
	return c
}

func (c *Comp) BoxDecorationClone() *Comp {
	c.El.Style(styles.BoxDecorationClone())
	return c
}

func (c *Comp) BoxDecorationSlice() *Comp {
	c.El.Style(styles.BoxDecorationSlice())
	return c
}

func (c *Comp) BreakAfterAll() *Comp {
	c.El.Style(styles.BreakAfterAll())
	return c
}

func (c *Comp) BreakAfterAuto() *Comp {
	c.El.Style(styles.BreakAfterAuto())
	return c
}

func (c *Comp) BreakAfterAvoid() *Comp {
	c.El.Style(styles.BreakAfterAvoid())
	return c
}

func (c *Comp) BreakAfterAvoidPage() *Comp {
	c.El.Style(styles.BreakAfterAvoidPage())
	return c
}

func (c *Comp) BreakAfterColumn() *Comp {
	c.El.Style(styles.BreakAfterColumn())
	return c
}

func (c *Comp) BreakAfterLeft() *Comp {
	c.El.Style(styles.BreakAfterLeft())
	return c
}

func (c *Comp) BreakAfterPage() *Comp {
	c.El.Style(styles.BreakAfterPage())
	return c
}

func (c *Comp) BreakAfterRight() *Comp {
	c.El.Style(styles.BreakAfterRight())
	return c
}

func (c *Comp) BreakAll() *Comp {
	c.El.Style(styles.BreakAll())
	return c
}

func (c *Comp) BreakBeforeAll() *Comp {
	c.El.Style(styles.BreakBeforeAll())
	return c
}

func (c *Comp) BreakBeforeAuto() *Comp {
	c.El.Style(styles.BreakBeforeAuto())
	return c
}

func (c *Comp) BreakBeforeAvoid() *Comp {
	c.El.Style(styles.BreakBeforeAvoid())
	return c
}

func (c *Comp) BreakBeforeAvoidPage() *Comp {
	c.El.Style(styles.BreakBeforeAvoidPage())
	return c
}

func (c *Comp) BreakBeforeColumn() *Comp {
	c.El.Style(styles.BreakBeforeColumn())
	return c
}

func (c *Comp) BreakBeforeLeft() *Comp {
	c.El.Style(styles.BreakBeforeLeft())
	return c
}

func (c *Comp) BreakBeforePage() *Comp {
	c.El.Style(styles.BreakBeforePage())
	return c
}

func (c *Comp) BreakBeforeRight() *Comp {
	c.El.Style(styles.BreakBeforeRight())
	return c
}

func (c *Comp) BreakInsideAuto() *Comp {
	c.El.Style(styles.BreakInsideAuto())
	return c
}

func (c *Comp) BreakInsideAvoid() *Comp {
	c.El.Style(styles.BreakInsideAvoid())
	return c
}

func (c *Comp) BreakInsideAvoidColumn() *Comp {
	c.El.Style(styles.BreakInsideAvoidColumn())
	return c
}

func (c *Comp) BreakInsideAvoidPage() *Comp {
	c.El.Style(styles.BreakInsideAvoidPage())
	return c
}

func (c *Comp) BreakKeep() *Comp {
	c.El.Style(styles.BreakKeep())
	return c
}

func (c *Comp) BreakNormal() *Comp {
	c.El.Style(styles.BreakNormal())
	return c
}

func (c *Comp) Brightness(val any) *Comp {
	c.El.Style(styles.Brightness(val))
	return c
}

func (c *Comp) Capitalize() *Comp {
	c.El.Style(styles.Capitalize())
	return c
}

func (c *Comp) CaptionBottom() *Comp {
	c.El.Style(styles.CaptionBottom())
	return c
}

func (c *Comp) CaptionTop() *Comp {
	c.El.Style(styles.CaptionTop())
	return c
}

func (c *Comp) Caret(val value.Value) *Comp {
	c.El.Style(styles.Caret(val))
	return c
}

func (c *Comp) CaretAmber(scale int) *Comp {
	c.El.Style(styles.CaretAmber(scale))
	return c
}

func (c *Comp) CaretAmberAlpha(scale int) *Comp {
	c.El.Style(styles.CaretAmberAlpha(scale))
	return c
}

func (c *Comp) CaretAmberDark(scale int) *Comp {
	c.El.Style(styles.CaretAmberDark(scale))
	return c
}

func (c *Comp) CaretAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretAmberDarkAlpha(scale))
	return c
}

func (c *Comp) CaretBlack() *Comp {
	c.El.Style(styles.CaretBlack())
	return c
}

func (c *Comp) CaretBlackAlpha(scale int) *Comp {
	c.El.Style(styles.CaretBlackAlpha(scale))
	return c
}

func (c *Comp) CaretBlue(scale int) *Comp {
	c.El.Style(styles.CaretBlue(scale))
	return c
}

func (c *Comp) CaretBlueAlpha(scale int) *Comp {
	c.El.Style(styles.CaretBlueAlpha(scale))
	return c
}

func (c *Comp) CaretBlueDark(scale int) *Comp {
	c.El.Style(styles.CaretBlueDark(scale))
	return c
}

func (c *Comp) CaretBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretBlueDarkAlpha(scale))
	return c
}

func (c *Comp) CaretBronze(scale int) *Comp {
	c.El.Style(styles.CaretBronze(scale))
	return c
}

func (c *Comp) CaretBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.CaretBronzeAlpha(scale))
	return c
}

func (c *Comp) CaretBronzeDark(scale int) *Comp {
	c.El.Style(styles.CaretBronzeDark(scale))
	return c
}

func (c *Comp) CaretBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) CaretBrown(scale int) *Comp {
	c.El.Style(styles.CaretBrown(scale))
	return c
}

func (c *Comp) CaretBrownAlpha(scale int) *Comp {
	c.El.Style(styles.CaretBrownAlpha(scale))
	return c
}

func (c *Comp) CaretBrownDark(scale int) *Comp {
	c.El.Style(styles.CaretBrownDark(scale))
	return c
}

func (c *Comp) CaretBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretBrownDarkAlpha(scale))
	return c
}

func (c *Comp) CaretCrimson(scale int) *Comp {
	c.El.Style(styles.CaretCrimson(scale))
	return c
}

func (c *Comp) CaretCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.CaretCrimsonAlpha(scale))
	return c
}

func (c *Comp) CaretCrimsonDark(scale int) *Comp {
	c.El.Style(styles.CaretCrimsonDark(scale))
	return c
}

func (c *Comp) CaretCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) CaretCurrent() *Comp {
	c.El.Style(styles.CaretCurrent())
	return c
}

func (c *Comp) CaretCyan(scale int) *Comp {
	c.El.Style(styles.CaretCyan(scale))
	return c
}

func (c *Comp) CaretCyanAlpha(scale int) *Comp {
	c.El.Style(styles.CaretCyanAlpha(scale))
	return c
}

func (c *Comp) CaretCyanDark(scale int) *Comp {
	c.El.Style(styles.CaretCyanDark(scale))
	return c
}

func (c *Comp) CaretCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretCyanDarkAlpha(scale))
	return c
}

func (c *Comp) CaretGold(scale int) *Comp {
	c.El.Style(styles.CaretGold(scale))
	return c
}

func (c *Comp) CaretGoldAlpha(scale int) *Comp {
	c.El.Style(styles.CaretGoldAlpha(scale))
	return c
}

func (c *Comp) CaretGoldDark(scale int) *Comp {
	c.El.Style(styles.CaretGoldDark(scale))
	return c
}

func (c *Comp) CaretGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretGoldDarkAlpha(scale))
	return c
}

func (c *Comp) CaretGrass(scale int) *Comp {
	c.El.Style(styles.CaretGrass(scale))
	return c
}

func (c *Comp) CaretGrassAlpha(scale int) *Comp {
	c.El.Style(styles.CaretGrassAlpha(scale))
	return c
}

func (c *Comp) CaretGrassDark(scale int) *Comp {
	c.El.Style(styles.CaretGrassDark(scale))
	return c
}

func (c *Comp) CaretGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretGrassDarkAlpha(scale))
	return c
}

func (c *Comp) CaretGray(scale int) *Comp {
	c.El.Style(styles.CaretGray(scale))
	return c
}

func (c *Comp) CaretGrayAlpha(scale int) *Comp {
	c.El.Style(styles.CaretGrayAlpha(scale))
	return c
}

func (c *Comp) CaretGrayDark(scale int) *Comp {
	c.El.Style(styles.CaretGrayDark(scale))
	return c
}

func (c *Comp) CaretGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretGrayDarkAlpha(scale))
	return c
}

func (c *Comp) CaretGreen(scale int) *Comp {
	c.El.Style(styles.CaretGreen(scale))
	return c
}

func (c *Comp) CaretGreenAlpha(scale int) *Comp {
	c.El.Style(styles.CaretGreenAlpha(scale))
	return c
}

func (c *Comp) CaretGreenDark(scale int) *Comp {
	c.El.Style(styles.CaretGreenDark(scale))
	return c
}

func (c *Comp) CaretGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretGreenDarkAlpha(scale))
	return c
}

func (c *Comp) CaretIndigo(scale int) *Comp {
	c.El.Style(styles.CaretIndigo(scale))
	return c
}

func (c *Comp) CaretIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.CaretIndigoAlpha(scale))
	return c
}

func (c *Comp) CaretIndigoDark(scale int) *Comp {
	c.El.Style(styles.CaretIndigoDark(scale))
	return c
}

func (c *Comp) CaretIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) CaretInherit() *Comp {
	c.El.Style(styles.CaretInherit())
	return c
}

func (c *Comp) CaretIris(scale int) *Comp {
	c.El.Style(styles.CaretIris(scale))
	return c
}

func (c *Comp) CaretIrisAlpha(scale int) *Comp {
	c.El.Style(styles.CaretIrisAlpha(scale))
	return c
}

func (c *Comp) CaretIrisDark(scale int) *Comp {
	c.El.Style(styles.CaretIrisDark(scale))
	return c
}

func (c *Comp) CaretIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretIrisDarkAlpha(scale))
	return c
}

func (c *Comp) CaretJade(scale int) *Comp {
	c.El.Style(styles.CaretJade(scale))
	return c
}

func (c *Comp) CaretJadeAlpha(scale int) *Comp {
	c.El.Style(styles.CaretJadeAlpha(scale))
	return c
}

func (c *Comp) CaretJadeDark(scale int) *Comp {
	c.El.Style(styles.CaretJadeDark(scale))
	return c
}

func (c *Comp) CaretJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretJadeDarkAlpha(scale))
	return c
}

func (c *Comp) CaretLime(scale int) *Comp {
	c.El.Style(styles.CaretLime(scale))
	return c
}

func (c *Comp) CaretLimeAlpha(scale int) *Comp {
	c.El.Style(styles.CaretLimeAlpha(scale))
	return c
}

func (c *Comp) CaretLimeDark(scale int) *Comp {
	c.El.Style(styles.CaretLimeDark(scale))
	return c
}

func (c *Comp) CaretLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretLimeDarkAlpha(scale))
	return c
}

func (c *Comp) CaretMauve(scale int) *Comp {
	c.El.Style(styles.CaretMauve(scale))
	return c
}

func (c *Comp) CaretMauveAlpha(scale int) *Comp {
	c.El.Style(styles.CaretMauveAlpha(scale))
	return c
}

func (c *Comp) CaretMauveDark(scale int) *Comp {
	c.El.Style(styles.CaretMauveDark(scale))
	return c
}

func (c *Comp) CaretMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretMauveDarkAlpha(scale))
	return c
}

func (c *Comp) CaretMint(scale int) *Comp {
	c.El.Style(styles.CaretMint(scale))
	return c
}

func (c *Comp) CaretMintAlpha(scale int) *Comp {
	c.El.Style(styles.CaretMintAlpha(scale))
	return c
}

func (c *Comp) CaretMintDark(scale int) *Comp {
	c.El.Style(styles.CaretMintDark(scale))
	return c
}

func (c *Comp) CaretMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretMintDarkAlpha(scale))
	return c
}

func (c *Comp) CaretOlive(scale int) *Comp {
	c.El.Style(styles.CaretOlive(scale))
	return c
}

func (c *Comp) CaretOliveAlpha(scale int) *Comp {
	c.El.Style(styles.CaretOliveAlpha(scale))
	return c
}

func (c *Comp) CaretOliveDark(scale int) *Comp {
	c.El.Style(styles.CaretOliveDark(scale))
	return c
}

func (c *Comp) CaretOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretOliveDarkAlpha(scale))
	return c
}

func (c *Comp) CaretOrange(scale int) *Comp {
	c.El.Style(styles.CaretOrange(scale))
	return c
}

func (c *Comp) CaretOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.CaretOrangeAlpha(scale))
	return c
}

func (c *Comp) CaretOrangeDark(scale int) *Comp {
	c.El.Style(styles.CaretOrangeDark(scale))
	return c
}

func (c *Comp) CaretOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) CaretPink(scale int) *Comp {
	c.El.Style(styles.CaretPink(scale))
	return c
}

func (c *Comp) CaretPinkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretPinkAlpha(scale))
	return c
}

func (c *Comp) CaretPinkDark(scale int) *Comp {
	c.El.Style(styles.CaretPinkDark(scale))
	return c
}

func (c *Comp) CaretPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretPinkDarkAlpha(scale))
	return c
}

func (c *Comp) CaretPlum(scale int) *Comp {
	c.El.Style(styles.CaretPlum(scale))
	return c
}

func (c *Comp) CaretPlumAlpha(scale int) *Comp {
	c.El.Style(styles.CaretPlumAlpha(scale))
	return c
}

func (c *Comp) CaretPlumDark(scale int) *Comp {
	c.El.Style(styles.CaretPlumDark(scale))
	return c
}

func (c *Comp) CaretPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretPlumDarkAlpha(scale))
	return c
}

func (c *Comp) CaretPurple(scale int) *Comp {
	c.El.Style(styles.CaretPurple(scale))
	return c
}

func (c *Comp) CaretPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.CaretPurpleAlpha(scale))
	return c
}

func (c *Comp) CaretPurpleDark(scale int) *Comp {
	c.El.Style(styles.CaretPurpleDark(scale))
	return c
}

func (c *Comp) CaretPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) CaretRed(scale int) *Comp {
	c.El.Style(styles.CaretRed(scale))
	return c
}

func (c *Comp) CaretRedAlpha(scale int) *Comp {
	c.El.Style(styles.CaretRedAlpha(scale))
	return c
}

func (c *Comp) CaretRedDark(scale int) *Comp {
	c.El.Style(styles.CaretRedDark(scale))
	return c
}

func (c *Comp) CaretRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretRedDarkAlpha(scale))
	return c
}

func (c *Comp) CaretRuby(scale int) *Comp {
	c.El.Style(styles.CaretRuby(scale))
	return c
}

func (c *Comp) CaretRubyAlpha(scale int) *Comp {
	c.El.Style(styles.CaretRubyAlpha(scale))
	return c
}

func (c *Comp) CaretRubyDark(scale int) *Comp {
	c.El.Style(styles.CaretRubyDark(scale))
	return c
}

func (c *Comp) CaretRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretRubyDarkAlpha(scale))
	return c
}

func (c *Comp) CaretSage(scale int) *Comp {
	c.El.Style(styles.CaretSage(scale))
	return c
}

func (c *Comp) CaretSageAlpha(scale int) *Comp {
	c.El.Style(styles.CaretSageAlpha(scale))
	return c
}

func (c *Comp) CaretSageDark(scale int) *Comp {
	c.El.Style(styles.CaretSageDark(scale))
	return c
}

func (c *Comp) CaretSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretSageDarkAlpha(scale))
	return c
}

func (c *Comp) CaretSand(scale int) *Comp {
	c.El.Style(styles.CaretSand(scale))
	return c
}

func (c *Comp) CaretSandAlpha(scale int) *Comp {
	c.El.Style(styles.CaretSandAlpha(scale))
	return c
}

func (c *Comp) CaretSandDark(scale int) *Comp {
	c.El.Style(styles.CaretSandDark(scale))
	return c
}

func (c *Comp) CaretSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretSandDarkAlpha(scale))
	return c
}

func (c *Comp) CaretSky(scale int) *Comp {
	c.El.Style(styles.CaretSky(scale))
	return c
}

func (c *Comp) CaretSkyAlpha(scale int) *Comp {
	c.El.Style(styles.CaretSkyAlpha(scale))
	return c
}

func (c *Comp) CaretSkyDark(scale int) *Comp {
	c.El.Style(styles.CaretSkyDark(scale))
	return c
}

func (c *Comp) CaretSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretSkyDarkAlpha(scale))
	return c
}

func (c *Comp) CaretSlate(scale int) *Comp {
	c.El.Style(styles.CaretSlate(scale))
	return c
}

func (c *Comp) CaretSlateAlpha(scale int) *Comp {
	c.El.Style(styles.CaretSlateAlpha(scale))
	return c
}

func (c *Comp) CaretSlateDark(scale int) *Comp {
	c.El.Style(styles.CaretSlateDark(scale))
	return c
}

func (c *Comp) CaretSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretSlateDarkAlpha(scale))
	return c
}

func (c *Comp) CaretTeal(scale int) *Comp {
	c.El.Style(styles.CaretTeal(scale))
	return c
}

func (c *Comp) CaretTealAlpha(scale int) *Comp {
	c.El.Style(styles.CaretTealAlpha(scale))
	return c
}

func (c *Comp) CaretTealDark(scale int) *Comp {
	c.El.Style(styles.CaretTealDark(scale))
	return c
}

func (c *Comp) CaretTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretTealDarkAlpha(scale))
	return c
}

func (c *Comp) CaretTomato(scale int) *Comp {
	c.El.Style(styles.CaretTomato(scale))
	return c
}

func (c *Comp) CaretTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.CaretTomatoAlpha(scale))
	return c
}

func (c *Comp) CaretTomatoDark(scale int) *Comp {
	c.El.Style(styles.CaretTomatoDark(scale))
	return c
}

func (c *Comp) CaretTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) CaretTransparent() *Comp {
	c.El.Style(styles.CaretTransparent())
	return c
}

func (c *Comp) CaretViolet(scale int) *Comp {
	c.El.Style(styles.CaretViolet(scale))
	return c
}

func (c *Comp) CaretVioletAlpha(scale int) *Comp {
	c.El.Style(styles.CaretVioletAlpha(scale))
	return c
}

func (c *Comp) CaretVioletDark(scale int) *Comp {
	c.El.Style(styles.CaretVioletDark(scale))
	return c
}

func (c *Comp) CaretVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretVioletDarkAlpha(scale))
	return c
}

func (c *Comp) CaretWhite() *Comp {
	c.El.Style(styles.CaretWhite())
	return c
}

func (c *Comp) CaretWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.CaretWhiteAlpha(scale))
	return c
}

func (c *Comp) CaretYellow(scale int) *Comp {
	c.El.Style(styles.CaretYellow(scale))
	return c
}

func (c *Comp) CaretYellowAlpha(scale int) *Comp {
	c.El.Style(styles.CaretYellowAlpha(scale))
	return c
}

func (c *Comp) CaretYellowDark(scale int) *Comp {
	c.El.Style(styles.CaretYellowDark(scale))
	return c
}

func (c *Comp) CaretYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.CaretYellowDarkAlpha(scale))
	return c
}

func (c *Comp) ClearBoth() *Comp {
	c.El.Style(styles.ClearBoth())
	return c
}

func (c *Comp) ClearEnd() *Comp {
	c.El.Style(styles.ClearEnd())
	return c
}

func (c *Comp) ClearLeft() *Comp {
	c.El.Style(styles.ClearLeft())
	return c
}

func (c *Comp) ClearNone() *Comp {
	c.El.Style(styles.ClearNone())
	return c
}

func (c *Comp) ClearRight() *Comp {
	c.El.Style(styles.ClearRight())
	return c
}

func (c *Comp) ClearStart() *Comp {
	c.El.Style(styles.ClearStart())
	return c
}

func (c *Comp) Col(number int) *Comp {
	c.El.Style(styles.Col(number))
	return c
}

func (c *Comp) ColAuto() *Comp {
	c.El.Style(styles.ColAuto())
	return c
}

func (c *Comp) ColBy(val value.Value) *Comp {
	c.El.Style(styles.ColBy(val))
	return c
}

func (c *Comp) ColEnd(number int) *Comp {
	c.El.Style(styles.ColEnd(number))
	return c
}

func (c *Comp) ColEndAuto() *Comp {
	c.El.Style(styles.ColEndAuto())
	return c
}

func (c *Comp) ColEndBy(val value.Value) *Comp {
	c.El.Style(styles.ColEndBy(val))
	return c
}

func (c *Comp) ColSpan(number int) *Comp {
	c.El.Style(styles.ColSpan(number))
	return c
}

func (c *Comp) ColSpanBy(val value.Value) *Comp {
	c.El.Style(styles.ColSpanBy(val))
	return c
}

func (c *Comp) ColSpanFull() *Comp {
	c.El.Style(styles.ColSpanFull())
	return c
}

func (c *Comp) ColStart(number int) *Comp {
	c.El.Style(styles.ColStart(number))
	return c
}

func (c *Comp) ColStartAuto() *Comp {
	c.El.Style(styles.ColStartAuto())
	return c
}

func (c *Comp) ColStartBy(val value.Value) *Comp {
	c.El.Style(styles.ColStartBy(val))
	return c
}

func (c *Comp) Collapse() *Comp {
	c.El.Style(styles.Collapse())
	return c
}

func (c *Comp) Columns(cols value.Value) *Comp {
	c.El.Style(styles.Columns(cols))
	return c
}

func (c *Comp) Columns2xl() *Comp {
	c.El.Style(styles.Columns2xl())
	return c
}

func (c *Comp) Columns2xs() *Comp {
	c.El.Style(styles.Columns2xs())
	return c
}

func (c *Comp) Columns3xl() *Comp {
	c.El.Style(styles.Columns3xl())
	return c
}

func (c *Comp) Columns3xs() *Comp {
	c.El.Style(styles.Columns3xs())
	return c
}

func (c *Comp) Columns4xl() *Comp {
	c.El.Style(styles.Columns4xl())
	return c
}

func (c *Comp) Columns5xl() *Comp {
	c.El.Style(styles.Columns5xl())
	return c
}

func (c *Comp) Columns6xl() *Comp {
	c.El.Style(styles.Columns6xl())
	return c
}

func (c *Comp) Columns7xl() *Comp {
	c.El.Style(styles.Columns7xl())
	return c
}

func (c *Comp) ColumnsAuto() *Comp {
	c.El.Style(styles.ColumnsAuto())
	return c
}

func (c *Comp) ColumnsLg() *Comp {
	c.El.Style(styles.ColumnsLg())
	return c
}

func (c *Comp) ColumnsMd() *Comp {
	c.El.Style(styles.ColumnsMd())
	return c
}

func (c *Comp) ColumnsSm() *Comp {
	c.El.Style(styles.ColumnsSm())
	return c
}

func (c *Comp) ColumnsXl() *Comp {
	c.El.Style(styles.ColumnsXl())
	return c
}

func (c *Comp) ColumnsXs() *Comp {
	c.El.Style(styles.ColumnsXs())
	return c
}

func (c *Comp) Container() *Comp {
	c.El.Style(styles.Container())
	return c
}

func (c *Comp) Content(val value.Value) *Comp {
	c.El.Style(styles.Content(val))
	return c
}

func (c *Comp) ContentAround() *Comp {
	c.El.Style(styles.ContentAround())
	return c
}

func (c *Comp) ContentBaseline() *Comp {
	c.El.Style(styles.ContentBaseline())
	return c
}

func (c *Comp) ContentBetween() *Comp {
	c.El.Style(styles.ContentBetween())
	return c
}

func (c *Comp) ContentCenter() *Comp {
	c.El.Style(styles.ContentCenter())
	return c
}

func (c *Comp) ContentEnd() *Comp {
	c.El.Style(styles.ContentEnd())
	return c
}

func (c *Comp) ContentEvenly() *Comp {
	c.El.Style(styles.ContentEvenly())
	return c
}

func (c *Comp) ContentNone() *Comp {
	c.El.Style(styles.ContentNone())
	return c
}

func (c *Comp) ContentNormal() *Comp {
	c.El.Style(styles.ContentNormal())
	return c
}

func (c *Comp) ContentStart() *Comp {
	c.El.Style(styles.ContentStart())
	return c
}

func (c *Comp) ContentStretch() *Comp {
	c.El.Style(styles.ContentStretch())
	return c
}

func (c *Comp) Contents() *Comp {
	c.El.Style(styles.Contents())
	return c
}

func (c *Comp) Contrast(val any) *Comp {
	c.El.Style(styles.Contrast(val))
	return c
}

func (c *Comp) Cursor(val value.Value) *Comp {
	c.El.Style(styles.Cursor(val))
	return c
}

func (c *Comp) CursorAlias() *Comp {
	c.El.Style(styles.CursorAlias())
	return c
}

func (c *Comp) CursorAllScroll() *Comp {
	c.El.Style(styles.CursorAllScroll())
	return c
}

func (c *Comp) CursorAuto() *Comp {
	c.El.Style(styles.CursorAuto())
	return c
}

func (c *Comp) CursorCell() *Comp {
	c.El.Style(styles.CursorCell())
	return c
}

func (c *Comp) CursorColResize() *Comp {
	c.El.Style(styles.CursorColResize())
	return c
}

func (c *Comp) CursorContextMenu() *Comp {
	c.El.Style(styles.CursorContextMenu())
	return c
}

func (c *Comp) CursorCopy() *Comp {
	c.El.Style(styles.CursorCopy())
	return c
}

func (c *Comp) CursorCrosshair() *Comp {
	c.El.Style(styles.CursorCrosshair())
	return c
}

func (c *Comp) CursorDefault() *Comp {
	c.El.Style(styles.CursorDefault())
	return c
}

func (c *Comp) CursorEResize() *Comp {
	c.El.Style(styles.CursorEResize())
	return c
}

func (c *Comp) CursorEwResize() *Comp {
	c.El.Style(styles.CursorEwResize())
	return c
}

func (c *Comp) CursorGrab() *Comp {
	c.El.Style(styles.CursorGrab())
	return c
}

func (c *Comp) CursorGrabbing() *Comp {
	c.El.Style(styles.CursorGrabbing())
	return c
}

func (c *Comp) CursorHelp() *Comp {
	c.El.Style(styles.CursorHelp())
	return c
}

func (c *Comp) CursorMove() *Comp {
	c.El.Style(styles.CursorMove())
	return c
}

func (c *Comp) CursorNResize() *Comp {
	c.El.Style(styles.CursorNResize())
	return c
}

func (c *Comp) CursorNeResize() *Comp {
	c.El.Style(styles.CursorNeResize())
	return c
}

func (c *Comp) CursorNeswResize() *Comp {
	c.El.Style(styles.CursorNeswResize())
	return c
}

func (c *Comp) CursorNoDrop() *Comp {
	c.El.Style(styles.CursorNoDrop())
	return c
}

func (c *Comp) CursorNone() *Comp {
	c.El.Style(styles.CursorNone())
	return c
}

func (c *Comp) CursorNotAllowed() *Comp {
	c.El.Style(styles.CursorNotAllowed())
	return c
}

func (c *Comp) CursorNsResize() *Comp {
	c.El.Style(styles.CursorNsResize())
	return c
}

func (c *Comp) CursorNwResize() *Comp {
	c.El.Style(styles.CursorNwResize())
	return c
}

func (c *Comp) CursorNwseResize() *Comp {
	c.El.Style(styles.CursorNwseResize())
	return c
}

func (c *Comp) CursorPointer() *Comp {
	c.El.Style(styles.CursorPointer())
	return c
}

func (c *Comp) CursorProgress() *Comp {
	c.El.Style(styles.CursorProgress())
	return c
}

func (c *Comp) CursorRowResize() *Comp {
	c.El.Style(styles.CursorRowResize())
	return c
}

func (c *Comp) CursorSResize() *Comp {
	c.El.Style(styles.CursorSResize())
	return c
}

func (c *Comp) CursorSeResize() *Comp {
	c.El.Style(styles.CursorSeResize())
	return c
}

func (c *Comp) CursorSwResize() *Comp {
	c.El.Style(styles.CursorSwResize())
	return c
}

func (c *Comp) CursorText() *Comp {
	c.El.Style(styles.CursorText())
	return c
}

func (c *Comp) CursorVerticalText() *Comp {
	c.El.Style(styles.CursorVerticalText())
	return c
}

func (c *Comp) CursorWResize() *Comp {
	c.El.Style(styles.CursorWResize())
	return c
}

func (c *Comp) CursorWait() *Comp {
	c.El.Style(styles.CursorWait())
	return c
}

func (c *Comp) CursorZoomIn() *Comp {
	c.El.Style(styles.CursorZoomIn())
	return c
}

func (c *Comp) CursorZoomOut() *Comp {
	c.El.Style(styles.CursorZoomOut())
	return c
}

func (c *Comp) DecorAmber(scale int) *Comp {
	c.El.Style(styles.DecorAmber(scale))
	return c
}

func (c *Comp) DecorAmberAlpha(scale int) *Comp {
	c.El.Style(styles.DecorAmberAlpha(scale))
	return c
}

func (c *Comp) DecorAmberDark(scale int) *Comp {
	c.El.Style(styles.DecorAmberDark(scale))
	return c
}

func (c *Comp) DecorAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorAmberDarkAlpha(scale))
	return c
}

func (c *Comp) DecorBlack() *Comp {
	c.El.Style(styles.DecorBlack())
	return c
}

func (c *Comp) DecorBlackAlpha(scale int) *Comp {
	c.El.Style(styles.DecorBlackAlpha(scale))
	return c
}

func (c *Comp) DecorBlue(scale int) *Comp {
	c.El.Style(styles.DecorBlue(scale))
	return c
}

func (c *Comp) DecorBlueAlpha(scale int) *Comp {
	c.El.Style(styles.DecorBlueAlpha(scale))
	return c
}

func (c *Comp) DecorBlueDark(scale int) *Comp {
	c.El.Style(styles.DecorBlueDark(scale))
	return c
}

func (c *Comp) DecorBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorBlueDarkAlpha(scale))
	return c
}

func (c *Comp) DecorBronze(scale int) *Comp {
	c.El.Style(styles.DecorBronze(scale))
	return c
}

func (c *Comp) DecorBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.DecorBronzeAlpha(scale))
	return c
}

func (c *Comp) DecorBronzeDark(scale int) *Comp {
	c.El.Style(styles.DecorBronzeDark(scale))
	return c
}

func (c *Comp) DecorBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) DecorBrown(scale int) *Comp {
	c.El.Style(styles.DecorBrown(scale))
	return c
}

func (c *Comp) DecorBrownAlpha(scale int) *Comp {
	c.El.Style(styles.DecorBrownAlpha(scale))
	return c
}

func (c *Comp) DecorBrownDark(scale int) *Comp {
	c.El.Style(styles.DecorBrownDark(scale))
	return c
}

func (c *Comp) DecorBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorBrownDarkAlpha(scale))
	return c
}

func (c *Comp) DecorColorBy(val value.Value) *Comp {
	c.El.Style(styles.DecorColorBy(val))
	return c
}

func (c *Comp) DecorCrimson(scale int) *Comp {
	c.El.Style(styles.DecorCrimson(scale))
	return c
}

func (c *Comp) DecorCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.DecorCrimsonAlpha(scale))
	return c
}

func (c *Comp) DecorCrimsonDark(scale int) *Comp {
	c.El.Style(styles.DecorCrimsonDark(scale))
	return c
}

func (c *Comp) DecorCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) DecorCurrent() *Comp {
	c.El.Style(styles.DecorCurrent())
	return c
}

func (c *Comp) DecorCyan(scale int) *Comp {
	c.El.Style(styles.DecorCyan(scale))
	return c
}

func (c *Comp) DecorCyanAlpha(scale int) *Comp {
	c.El.Style(styles.DecorCyanAlpha(scale))
	return c
}

func (c *Comp) DecorCyanDark(scale int) *Comp {
	c.El.Style(styles.DecorCyanDark(scale))
	return c
}

func (c *Comp) DecorCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorCyanDarkAlpha(scale))
	return c
}

func (c *Comp) DecorGold(scale int) *Comp {
	c.El.Style(styles.DecorGold(scale))
	return c
}

func (c *Comp) DecorGoldAlpha(scale int) *Comp {
	c.El.Style(styles.DecorGoldAlpha(scale))
	return c
}

func (c *Comp) DecorGoldDark(scale int) *Comp {
	c.El.Style(styles.DecorGoldDark(scale))
	return c
}

func (c *Comp) DecorGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorGoldDarkAlpha(scale))
	return c
}

func (c *Comp) DecorGrass(scale int) *Comp {
	c.El.Style(styles.DecorGrass(scale))
	return c
}

func (c *Comp) DecorGrassAlpha(scale int) *Comp {
	c.El.Style(styles.DecorGrassAlpha(scale))
	return c
}

func (c *Comp) DecorGrassDark(scale int) *Comp {
	c.El.Style(styles.DecorGrassDark(scale))
	return c
}

func (c *Comp) DecorGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorGrassDarkAlpha(scale))
	return c
}

func (c *Comp) DecorGray(scale int) *Comp {
	c.El.Style(styles.DecorGray(scale))
	return c
}

func (c *Comp) DecorGrayAlpha(scale int) *Comp {
	c.El.Style(styles.DecorGrayAlpha(scale))
	return c
}

func (c *Comp) DecorGrayDark(scale int) *Comp {
	c.El.Style(styles.DecorGrayDark(scale))
	return c
}

func (c *Comp) DecorGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorGrayDarkAlpha(scale))
	return c
}

func (c *Comp) DecorGreen(scale int) *Comp {
	c.El.Style(styles.DecorGreen(scale))
	return c
}

func (c *Comp) DecorGreenAlpha(scale int) *Comp {
	c.El.Style(styles.DecorGreenAlpha(scale))
	return c
}

func (c *Comp) DecorGreenDark(scale int) *Comp {
	c.El.Style(styles.DecorGreenDark(scale))
	return c
}

func (c *Comp) DecorGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorGreenDarkAlpha(scale))
	return c
}

func (c *Comp) DecorIndigo(scale int) *Comp {
	c.El.Style(styles.DecorIndigo(scale))
	return c
}

func (c *Comp) DecorIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.DecorIndigoAlpha(scale))
	return c
}

func (c *Comp) DecorIndigoDark(scale int) *Comp {
	c.El.Style(styles.DecorIndigoDark(scale))
	return c
}

func (c *Comp) DecorIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) DecorInherit() *Comp {
	c.El.Style(styles.DecorInherit())
	return c
}

func (c *Comp) DecorIris(scale int) *Comp {
	c.El.Style(styles.DecorIris(scale))
	return c
}

func (c *Comp) DecorIrisAlpha(scale int) *Comp {
	c.El.Style(styles.DecorIrisAlpha(scale))
	return c
}

func (c *Comp) DecorIrisDark(scale int) *Comp {
	c.El.Style(styles.DecorIrisDark(scale))
	return c
}

func (c *Comp) DecorIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorIrisDarkAlpha(scale))
	return c
}

func (c *Comp) DecorJade(scale int) *Comp {
	c.El.Style(styles.DecorJade(scale))
	return c
}

func (c *Comp) DecorJadeAlpha(scale int) *Comp {
	c.El.Style(styles.DecorJadeAlpha(scale))
	return c
}

func (c *Comp) DecorJadeDark(scale int) *Comp {
	c.El.Style(styles.DecorJadeDark(scale))
	return c
}

func (c *Comp) DecorJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorJadeDarkAlpha(scale))
	return c
}

func (c *Comp) DecorLime(scale int) *Comp {
	c.El.Style(styles.DecorLime(scale))
	return c
}

func (c *Comp) DecorLimeAlpha(scale int) *Comp {
	c.El.Style(styles.DecorLimeAlpha(scale))
	return c
}

func (c *Comp) DecorLimeDark(scale int) *Comp {
	c.El.Style(styles.DecorLimeDark(scale))
	return c
}

func (c *Comp) DecorLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorLimeDarkAlpha(scale))
	return c
}

func (c *Comp) DecorMauve(scale int) *Comp {
	c.El.Style(styles.DecorMauve(scale))
	return c
}

func (c *Comp) DecorMauveAlpha(scale int) *Comp {
	c.El.Style(styles.DecorMauveAlpha(scale))
	return c
}

func (c *Comp) DecorMauveDark(scale int) *Comp {
	c.El.Style(styles.DecorMauveDark(scale))
	return c
}

func (c *Comp) DecorMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorMauveDarkAlpha(scale))
	return c
}

func (c *Comp) DecorMint(scale int) *Comp {
	c.El.Style(styles.DecorMint(scale))
	return c
}

func (c *Comp) DecorMintAlpha(scale int) *Comp {
	c.El.Style(styles.DecorMintAlpha(scale))
	return c
}

func (c *Comp) DecorMintDark(scale int) *Comp {
	c.El.Style(styles.DecorMintDark(scale))
	return c
}

func (c *Comp) DecorMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorMintDarkAlpha(scale))
	return c
}

func (c *Comp) DecorOlive(scale int) *Comp {
	c.El.Style(styles.DecorOlive(scale))
	return c
}

func (c *Comp) DecorOliveAlpha(scale int) *Comp {
	c.El.Style(styles.DecorOliveAlpha(scale))
	return c
}

func (c *Comp) DecorOliveDark(scale int) *Comp {
	c.El.Style(styles.DecorOliveDark(scale))
	return c
}

func (c *Comp) DecorOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorOliveDarkAlpha(scale))
	return c
}

func (c *Comp) DecorOrange(scale int) *Comp {
	c.El.Style(styles.DecorOrange(scale))
	return c
}

func (c *Comp) DecorOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.DecorOrangeAlpha(scale))
	return c
}

func (c *Comp) DecorOrangeDark(scale int) *Comp {
	c.El.Style(styles.DecorOrangeDark(scale))
	return c
}

func (c *Comp) DecorOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) DecorPink(scale int) *Comp {
	c.El.Style(styles.DecorPink(scale))
	return c
}

func (c *Comp) DecorPinkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorPinkAlpha(scale))
	return c
}

func (c *Comp) DecorPinkDark(scale int) *Comp {
	c.El.Style(styles.DecorPinkDark(scale))
	return c
}

func (c *Comp) DecorPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorPinkDarkAlpha(scale))
	return c
}

func (c *Comp) DecorPlum(scale int) *Comp {
	c.El.Style(styles.DecorPlum(scale))
	return c
}

func (c *Comp) DecorPlumAlpha(scale int) *Comp {
	c.El.Style(styles.DecorPlumAlpha(scale))
	return c
}

func (c *Comp) DecorPlumDark(scale int) *Comp {
	c.El.Style(styles.DecorPlumDark(scale))
	return c
}

func (c *Comp) DecorPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorPlumDarkAlpha(scale))
	return c
}

func (c *Comp) DecorPurple(scale int) *Comp {
	c.El.Style(styles.DecorPurple(scale))
	return c
}

func (c *Comp) DecorPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.DecorPurpleAlpha(scale))
	return c
}

func (c *Comp) DecorPurpleDark(scale int) *Comp {
	c.El.Style(styles.DecorPurpleDark(scale))
	return c
}

func (c *Comp) DecorPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) DecorRed(scale int) *Comp {
	c.El.Style(styles.DecorRed(scale))
	return c
}

func (c *Comp) DecorRedAlpha(scale int) *Comp {
	c.El.Style(styles.DecorRedAlpha(scale))
	return c
}

func (c *Comp) DecorRedDark(scale int) *Comp {
	c.El.Style(styles.DecorRedDark(scale))
	return c
}

func (c *Comp) DecorRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorRedDarkAlpha(scale))
	return c
}

func (c *Comp) DecorRuby(scale int) *Comp {
	c.El.Style(styles.DecorRuby(scale))
	return c
}

func (c *Comp) DecorRubyAlpha(scale int) *Comp {
	c.El.Style(styles.DecorRubyAlpha(scale))
	return c
}

func (c *Comp) DecorRubyDark(scale int) *Comp {
	c.El.Style(styles.DecorRubyDark(scale))
	return c
}

func (c *Comp) DecorRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorRubyDarkAlpha(scale))
	return c
}

func (c *Comp) DecorSage(scale int) *Comp {
	c.El.Style(styles.DecorSage(scale))
	return c
}

func (c *Comp) DecorSageAlpha(scale int) *Comp {
	c.El.Style(styles.DecorSageAlpha(scale))
	return c
}

func (c *Comp) DecorSageDark(scale int) *Comp {
	c.El.Style(styles.DecorSageDark(scale))
	return c
}

func (c *Comp) DecorSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorSageDarkAlpha(scale))
	return c
}

func (c *Comp) DecorSand(scale int) *Comp {
	c.El.Style(styles.DecorSand(scale))
	return c
}

func (c *Comp) DecorSandAlpha(scale int) *Comp {
	c.El.Style(styles.DecorSandAlpha(scale))
	return c
}

func (c *Comp) DecorSandDark(scale int) *Comp {
	c.El.Style(styles.DecorSandDark(scale))
	return c
}

func (c *Comp) DecorSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorSandDarkAlpha(scale))
	return c
}

func (c *Comp) DecorSky(scale int) *Comp {
	c.El.Style(styles.DecorSky(scale))
	return c
}

func (c *Comp) DecorSkyAlpha(scale int) *Comp {
	c.El.Style(styles.DecorSkyAlpha(scale))
	return c
}

func (c *Comp) DecorSkyDark(scale int) *Comp {
	c.El.Style(styles.DecorSkyDark(scale))
	return c
}

func (c *Comp) DecorSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorSkyDarkAlpha(scale))
	return c
}

func (c *Comp) DecorSlate(scale int) *Comp {
	c.El.Style(styles.DecorSlate(scale))
	return c
}

func (c *Comp) DecorSlateAlpha(scale int) *Comp {
	c.El.Style(styles.DecorSlateAlpha(scale))
	return c
}

func (c *Comp) DecorSlateDark(scale int) *Comp {
	c.El.Style(styles.DecorSlateDark(scale))
	return c
}

func (c *Comp) DecorSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorSlateDarkAlpha(scale))
	return c
}

func (c *Comp) DecorTeal(scale int) *Comp {
	c.El.Style(styles.DecorTeal(scale))
	return c
}

func (c *Comp) DecorTealAlpha(scale int) *Comp {
	c.El.Style(styles.DecorTealAlpha(scale))
	return c
}

func (c *Comp) DecorTealDark(scale int) *Comp {
	c.El.Style(styles.DecorTealDark(scale))
	return c
}

func (c *Comp) DecorTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorTealDarkAlpha(scale))
	return c
}

func (c *Comp) DecorTomato(scale int) *Comp {
	c.El.Style(styles.DecorTomato(scale))
	return c
}

func (c *Comp) DecorTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.DecorTomatoAlpha(scale))
	return c
}

func (c *Comp) DecorTomatoDark(scale int) *Comp {
	c.El.Style(styles.DecorTomatoDark(scale))
	return c
}

func (c *Comp) DecorTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) DecorTransparent() *Comp {
	c.El.Style(styles.DecorTransparent())
	return c
}

func (c *Comp) DecorViolet(scale int) *Comp {
	c.El.Style(styles.DecorViolet(scale))
	return c
}

func (c *Comp) DecorVioletAlpha(scale int) *Comp {
	c.El.Style(styles.DecorVioletAlpha(scale))
	return c
}

func (c *Comp) DecorVioletDark(scale int) *Comp {
	c.El.Style(styles.DecorVioletDark(scale))
	return c
}

func (c *Comp) DecorVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorVioletDarkAlpha(scale))
	return c
}

func (c *Comp) DecorWhite() *Comp {
	c.El.Style(styles.DecorWhite())
	return c
}

func (c *Comp) DecorWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.DecorWhiteAlpha(scale))
	return c
}

func (c *Comp) DecorYellow(scale int) *Comp {
	c.El.Style(styles.DecorYellow(scale))
	return c
}

func (c *Comp) DecorYellowAlpha(scale int) *Comp {
	c.El.Style(styles.DecorYellowAlpha(scale))
	return c
}

func (c *Comp) DecorYellowDark(scale int) *Comp {
	c.El.Style(styles.DecorYellowDark(scale))
	return c
}

func (c *Comp) DecorYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DecorYellowDarkAlpha(scale))
	return c
}

func (c *Comp) Decoration(number int) *Comp {
	c.El.Style(styles.Decoration(number))
	return c
}

func (c *Comp) DecorationAuto() *Comp {
	c.El.Style(styles.DecorationAuto())
	return c
}

func (c *Comp) DecorationBy(val value.Value) *Comp {
	c.El.Style(styles.DecorationBy(val))
	return c
}

func (c *Comp) DecorationDashed() *Comp {
	c.El.Style(styles.DecorationDashed())
	return c
}

func (c *Comp) DecorationDotted() *Comp {
	c.El.Style(styles.DecorationDotted())
	return c
}

func (c *Comp) DecorationDouble() *Comp {
	c.El.Style(styles.DecorationDouble())
	return c
}

func (c *Comp) DecorationFromFont() *Comp {
	c.El.Style(styles.DecorationFromFont())
	return c
}

func (c *Comp) DecorationSolid() *Comp {
	c.El.Style(styles.DecorationSolid())
	return c
}

func (c *Comp) DecorationWavy() *Comp {
	c.El.Style(styles.DecorationWavy())
	return c
}

func (c *Comp) Delay(val any) *Comp {
	c.El.Style(styles.Delay(val))
	return c
}

func (c *Comp) DiagonalFractions() *Comp {
	c.El.Style(styles.DiagonalFractions())
	return c
}

func (c *Comp) DivideAmber(scale int) *Comp {
	c.El.Style(styles.DivideAmber(scale))
	return c
}

func (c *Comp) DivideAmberAlpha(scale int) *Comp {
	c.El.Style(styles.DivideAmberAlpha(scale))
	return c
}

func (c *Comp) DivideAmberDark(scale int) *Comp {
	c.El.Style(styles.DivideAmberDark(scale))
	return c
}

func (c *Comp) DivideAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideAmberDarkAlpha(scale))
	return c
}

func (c *Comp) DivideBlack() *Comp {
	c.El.Style(styles.DivideBlack())
	return c
}

func (c *Comp) DivideBlackAlpha(scale int) *Comp {
	c.El.Style(styles.DivideBlackAlpha(scale))
	return c
}

func (c *Comp) DivideBlue(scale int) *Comp {
	c.El.Style(styles.DivideBlue(scale))
	return c
}

func (c *Comp) DivideBlueAlpha(scale int) *Comp {
	c.El.Style(styles.DivideBlueAlpha(scale))
	return c
}

func (c *Comp) DivideBlueDark(scale int) *Comp {
	c.El.Style(styles.DivideBlueDark(scale))
	return c
}

func (c *Comp) DivideBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideBlueDarkAlpha(scale))
	return c
}

func (c *Comp) DivideBronze(scale int) *Comp {
	c.El.Style(styles.DivideBronze(scale))
	return c
}

func (c *Comp) DivideBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.DivideBronzeAlpha(scale))
	return c
}

func (c *Comp) DivideBronzeDark(scale int) *Comp {
	c.El.Style(styles.DivideBronzeDark(scale))
	return c
}

func (c *Comp) DivideBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) DivideBrown(scale int) *Comp {
	c.El.Style(styles.DivideBrown(scale))
	return c
}

func (c *Comp) DivideBrownAlpha(scale int) *Comp {
	c.El.Style(styles.DivideBrownAlpha(scale))
	return c
}

func (c *Comp) DivideBrownDark(scale int) *Comp {
	c.El.Style(styles.DivideBrownDark(scale))
	return c
}

func (c *Comp) DivideBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideBrownDarkAlpha(scale))
	return c
}

func (c *Comp) DivideColor(val value.Value) *Comp {
	c.El.Style(styles.DivideColor(val))
	return c
}

func (c *Comp) DivideCrimson(scale int) *Comp {
	c.El.Style(styles.DivideCrimson(scale))
	return c
}

func (c *Comp) DivideCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.DivideCrimsonAlpha(scale))
	return c
}

func (c *Comp) DivideCrimsonDark(scale int) *Comp {
	c.El.Style(styles.DivideCrimsonDark(scale))
	return c
}

func (c *Comp) DivideCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) DivideCurrent() *Comp {
	c.El.Style(styles.DivideCurrent())
	return c
}

func (c *Comp) DivideCyan(scale int) *Comp {
	c.El.Style(styles.DivideCyan(scale))
	return c
}

func (c *Comp) DivideCyanAlpha(scale int) *Comp {
	c.El.Style(styles.DivideCyanAlpha(scale))
	return c
}

func (c *Comp) DivideCyanDark(scale int) *Comp {
	c.El.Style(styles.DivideCyanDark(scale))
	return c
}

func (c *Comp) DivideCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideCyanDarkAlpha(scale))
	return c
}

func (c *Comp) DivideDashed() *Comp {
	c.El.Style(styles.DivideDashed())
	return c
}

func (c *Comp) DivideDotted() *Comp {
	c.El.Style(styles.DivideDotted())
	return c
}

func (c *Comp) DivideDouble() *Comp {
	c.El.Style(styles.DivideDouble())
	return c
}

func (c *Comp) DivideGold(scale int) *Comp {
	c.El.Style(styles.DivideGold(scale))
	return c
}

func (c *Comp) DivideGoldAlpha(scale int) *Comp {
	c.El.Style(styles.DivideGoldAlpha(scale))
	return c
}

func (c *Comp) DivideGoldDark(scale int) *Comp {
	c.El.Style(styles.DivideGoldDark(scale))
	return c
}

func (c *Comp) DivideGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideGoldDarkAlpha(scale))
	return c
}

func (c *Comp) DivideGrass(scale int) *Comp {
	c.El.Style(styles.DivideGrass(scale))
	return c
}

func (c *Comp) DivideGrassAlpha(scale int) *Comp {
	c.El.Style(styles.DivideGrassAlpha(scale))
	return c
}

func (c *Comp) DivideGrassDark(scale int) *Comp {
	c.El.Style(styles.DivideGrassDark(scale))
	return c
}

func (c *Comp) DivideGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideGrassDarkAlpha(scale))
	return c
}

func (c *Comp) DivideGray(scale int) *Comp {
	c.El.Style(styles.DivideGray(scale))
	return c
}

func (c *Comp) DivideGrayAlpha(scale int) *Comp {
	c.El.Style(styles.DivideGrayAlpha(scale))
	return c
}

func (c *Comp) DivideGrayDark(scale int) *Comp {
	c.El.Style(styles.DivideGrayDark(scale))
	return c
}

func (c *Comp) DivideGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideGrayDarkAlpha(scale))
	return c
}

func (c *Comp) DivideGreen(scale int) *Comp {
	c.El.Style(styles.DivideGreen(scale))
	return c
}

func (c *Comp) DivideGreenAlpha(scale int) *Comp {
	c.El.Style(styles.DivideGreenAlpha(scale))
	return c
}

func (c *Comp) DivideGreenDark(scale int) *Comp {
	c.El.Style(styles.DivideGreenDark(scale))
	return c
}

func (c *Comp) DivideGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideGreenDarkAlpha(scale))
	return c
}

func (c *Comp) DivideHidden() *Comp {
	c.El.Style(styles.DivideHidden())
	return c
}

func (c *Comp) DivideIndigo(scale int) *Comp {
	c.El.Style(styles.DivideIndigo(scale))
	return c
}

func (c *Comp) DivideIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.DivideIndigoAlpha(scale))
	return c
}

func (c *Comp) DivideIndigoDark(scale int) *Comp {
	c.El.Style(styles.DivideIndigoDark(scale))
	return c
}

func (c *Comp) DivideIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) DivideInherit() *Comp {
	c.El.Style(styles.DivideInherit())
	return c
}

func (c *Comp) DivideIris(scale int) *Comp {
	c.El.Style(styles.DivideIris(scale))
	return c
}

func (c *Comp) DivideIrisAlpha(scale int) *Comp {
	c.El.Style(styles.DivideIrisAlpha(scale))
	return c
}

func (c *Comp) DivideIrisDark(scale int) *Comp {
	c.El.Style(styles.DivideIrisDark(scale))
	return c
}

func (c *Comp) DivideIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideIrisDarkAlpha(scale))
	return c
}

func (c *Comp) DivideJade(scale int) *Comp {
	c.El.Style(styles.DivideJade(scale))
	return c
}

func (c *Comp) DivideJadeAlpha(scale int) *Comp {
	c.El.Style(styles.DivideJadeAlpha(scale))
	return c
}

func (c *Comp) DivideJadeDark(scale int) *Comp {
	c.El.Style(styles.DivideJadeDark(scale))
	return c
}

func (c *Comp) DivideJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideJadeDarkAlpha(scale))
	return c
}

func (c *Comp) DivideLime(scale int) *Comp {
	c.El.Style(styles.DivideLime(scale))
	return c
}

func (c *Comp) DivideLimeAlpha(scale int) *Comp {
	c.El.Style(styles.DivideLimeAlpha(scale))
	return c
}

func (c *Comp) DivideLimeDark(scale int) *Comp {
	c.El.Style(styles.DivideLimeDark(scale))
	return c
}

func (c *Comp) DivideLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideLimeDarkAlpha(scale))
	return c
}

func (c *Comp) DivideMauve(scale int) *Comp {
	c.El.Style(styles.DivideMauve(scale))
	return c
}

func (c *Comp) DivideMauveAlpha(scale int) *Comp {
	c.El.Style(styles.DivideMauveAlpha(scale))
	return c
}

func (c *Comp) DivideMauveDark(scale int) *Comp {
	c.El.Style(styles.DivideMauveDark(scale))
	return c
}

func (c *Comp) DivideMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideMauveDarkAlpha(scale))
	return c
}

func (c *Comp) DivideMint(scale int) *Comp {
	c.El.Style(styles.DivideMint(scale))
	return c
}

func (c *Comp) DivideMintAlpha(scale int) *Comp {
	c.El.Style(styles.DivideMintAlpha(scale))
	return c
}

func (c *Comp) DivideMintDark(scale int) *Comp {
	c.El.Style(styles.DivideMintDark(scale))
	return c
}

func (c *Comp) DivideMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideMintDarkAlpha(scale))
	return c
}

func (c *Comp) DivideNone() *Comp {
	c.El.Style(styles.DivideNone())
	return c
}

func (c *Comp) DivideOlive(scale int) *Comp {
	c.El.Style(styles.DivideOlive(scale))
	return c
}

func (c *Comp) DivideOliveAlpha(scale int) *Comp {
	c.El.Style(styles.DivideOliveAlpha(scale))
	return c
}

func (c *Comp) DivideOliveDark(scale int) *Comp {
	c.El.Style(styles.DivideOliveDark(scale))
	return c
}

func (c *Comp) DivideOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideOliveDarkAlpha(scale))
	return c
}

func (c *Comp) DivideOrange(scale int) *Comp {
	c.El.Style(styles.DivideOrange(scale))
	return c
}

func (c *Comp) DivideOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.DivideOrangeAlpha(scale))
	return c
}

func (c *Comp) DivideOrangeDark(scale int) *Comp {
	c.El.Style(styles.DivideOrangeDark(scale))
	return c
}

func (c *Comp) DivideOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) DividePink(scale int) *Comp {
	c.El.Style(styles.DividePink(scale))
	return c
}

func (c *Comp) DividePinkAlpha(scale int) *Comp {
	c.El.Style(styles.DividePinkAlpha(scale))
	return c
}

func (c *Comp) DividePinkDark(scale int) *Comp {
	c.El.Style(styles.DividePinkDark(scale))
	return c
}

func (c *Comp) DividePinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DividePinkDarkAlpha(scale))
	return c
}

func (c *Comp) DividePlum(scale int) *Comp {
	c.El.Style(styles.DividePlum(scale))
	return c
}

func (c *Comp) DividePlumAlpha(scale int) *Comp {
	c.El.Style(styles.DividePlumAlpha(scale))
	return c
}

func (c *Comp) DividePlumDark(scale int) *Comp {
	c.El.Style(styles.DividePlumDark(scale))
	return c
}

func (c *Comp) DividePlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DividePlumDarkAlpha(scale))
	return c
}

func (c *Comp) DividePurple(scale int) *Comp {
	c.El.Style(styles.DividePurple(scale))
	return c
}

func (c *Comp) DividePurpleAlpha(scale int) *Comp {
	c.El.Style(styles.DividePurpleAlpha(scale))
	return c
}

func (c *Comp) DividePurpleDark(scale int) *Comp {
	c.El.Style(styles.DividePurpleDark(scale))
	return c
}

func (c *Comp) DividePurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DividePurpleDarkAlpha(scale))
	return c
}

func (c *Comp) DivideRed(scale int) *Comp {
	c.El.Style(styles.DivideRed(scale))
	return c
}

func (c *Comp) DivideRedAlpha(scale int) *Comp {
	c.El.Style(styles.DivideRedAlpha(scale))
	return c
}

func (c *Comp) DivideRedDark(scale int) *Comp {
	c.El.Style(styles.DivideRedDark(scale))
	return c
}

func (c *Comp) DivideRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideRedDarkAlpha(scale))
	return c
}

func (c *Comp) DivideRuby(scale int) *Comp {
	c.El.Style(styles.DivideRuby(scale))
	return c
}

func (c *Comp) DivideRubyAlpha(scale int) *Comp {
	c.El.Style(styles.DivideRubyAlpha(scale))
	return c
}

func (c *Comp) DivideRubyDark(scale int) *Comp {
	c.El.Style(styles.DivideRubyDark(scale))
	return c
}

func (c *Comp) DivideRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideRubyDarkAlpha(scale))
	return c
}

func (c *Comp) DivideSage(scale int) *Comp {
	c.El.Style(styles.DivideSage(scale))
	return c
}

func (c *Comp) DivideSageAlpha(scale int) *Comp {
	c.El.Style(styles.DivideSageAlpha(scale))
	return c
}

func (c *Comp) DivideSageDark(scale int) *Comp {
	c.El.Style(styles.DivideSageDark(scale))
	return c
}

func (c *Comp) DivideSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideSageDarkAlpha(scale))
	return c
}

func (c *Comp) DivideSand(scale int) *Comp {
	c.El.Style(styles.DivideSand(scale))
	return c
}

func (c *Comp) DivideSandAlpha(scale int) *Comp {
	c.El.Style(styles.DivideSandAlpha(scale))
	return c
}

func (c *Comp) DivideSandDark(scale int) *Comp {
	c.El.Style(styles.DivideSandDark(scale))
	return c
}

func (c *Comp) DivideSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideSandDarkAlpha(scale))
	return c
}

func (c *Comp) DivideSky(scale int) *Comp {
	c.El.Style(styles.DivideSky(scale))
	return c
}

func (c *Comp) DivideSkyAlpha(scale int) *Comp {
	c.El.Style(styles.DivideSkyAlpha(scale))
	return c
}

func (c *Comp) DivideSkyDark(scale int) *Comp {
	c.El.Style(styles.DivideSkyDark(scale))
	return c
}

func (c *Comp) DivideSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideSkyDarkAlpha(scale))
	return c
}

func (c *Comp) DivideSlate(scale int) *Comp {
	c.El.Style(styles.DivideSlate(scale))
	return c
}

func (c *Comp) DivideSlateAlpha(scale int) *Comp {
	c.El.Style(styles.DivideSlateAlpha(scale))
	return c
}

func (c *Comp) DivideSlateDark(scale int) *Comp {
	c.El.Style(styles.DivideSlateDark(scale))
	return c
}

func (c *Comp) DivideSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideSlateDarkAlpha(scale))
	return c
}

func (c *Comp) DivideSolid() *Comp {
	c.El.Style(styles.DivideSolid())
	return c
}

func (c *Comp) DivideTeal(scale int) *Comp {
	c.El.Style(styles.DivideTeal(scale))
	return c
}

func (c *Comp) DivideTealAlpha(scale int) *Comp {
	c.El.Style(styles.DivideTealAlpha(scale))
	return c
}

func (c *Comp) DivideTealDark(scale int) *Comp {
	c.El.Style(styles.DivideTealDark(scale))
	return c
}

func (c *Comp) DivideTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideTealDarkAlpha(scale))
	return c
}

func (c *Comp) DivideTomato(scale int) *Comp {
	c.El.Style(styles.DivideTomato(scale))
	return c
}

func (c *Comp) DivideTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.DivideTomatoAlpha(scale))
	return c
}

func (c *Comp) DivideTomatoDark(scale int) *Comp {
	c.El.Style(styles.DivideTomatoDark(scale))
	return c
}

func (c *Comp) DivideTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) DivideTransparent() *Comp {
	c.El.Style(styles.DivideTransparent())
	return c
}

func (c *Comp) DivideViolet(scale int) *Comp {
	c.El.Style(styles.DivideViolet(scale))
	return c
}

func (c *Comp) DivideVioletAlpha(scale int) *Comp {
	c.El.Style(styles.DivideVioletAlpha(scale))
	return c
}

func (c *Comp) DivideVioletDark(scale int) *Comp {
	c.El.Style(styles.DivideVioletDark(scale))
	return c
}

func (c *Comp) DivideVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideVioletDarkAlpha(scale))
	return c
}

func (c *Comp) DivideWhite() *Comp {
	c.El.Style(styles.DivideWhite())
	return c
}

func (c *Comp) DivideWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.DivideWhiteAlpha(scale))
	return c
}

func (c *Comp) DivideX(val ...value.Value) *Comp {
	c.El.Style(styles.DivideX(val...))
	return c
}

func (c *Comp) DivideXReverse(val ...value.Value) *Comp {
	c.El.Style(styles.DivideXReverse(val...))
	return c
}

func (c *Comp) DivideY(val ...value.Value) *Comp {
	c.El.Style(styles.DivideY(val...))
	return c
}

func (c *Comp) DivideYReverse(val ...value.Value) *Comp {
	c.El.Style(styles.DivideYReverse(val...))
	return c
}

func (c *Comp) DivideYellow(scale int) *Comp {
	c.El.Style(styles.DivideYellow(scale))
	return c
}

func (c *Comp) DivideYellowAlpha(scale int) *Comp {
	c.El.Style(styles.DivideYellowAlpha(scale))
	return c
}

func (c *Comp) DivideYellowDark(scale int) *Comp {
	c.El.Style(styles.DivideYellowDark(scale))
	return c
}

func (c *Comp) DivideYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DivideYellowDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadow(val value.Value) *Comp {
	c.El.Style(styles.DropShadow(val))
	return c
}

func (c *Comp) DropShadow2xl() *Comp {
	c.El.Style(styles.DropShadow2xl())
	return c
}

func (c *Comp) DropShadow2xs() *Comp {
	c.El.Style(styles.DropShadow2xs())
	return c
}

func (c *Comp) DropShadowAmber(scale int) *Comp {
	c.El.Style(styles.DropShadowAmber(scale))
	return c
}

func (c *Comp) DropShadowAmberAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowAmberAlpha(scale))
	return c
}

func (c *Comp) DropShadowAmberDark(scale int) *Comp {
	c.El.Style(styles.DropShadowAmberDark(scale))
	return c
}

func (c *Comp) DropShadowAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowAmberDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowBlack() *Comp {
	c.El.Style(styles.DropShadowBlack())
	return c
}

func (c *Comp) DropShadowBlackAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowBlackAlpha(scale))
	return c
}

func (c *Comp) DropShadowBlue(scale int) *Comp {
	c.El.Style(styles.DropShadowBlue(scale))
	return c
}

func (c *Comp) DropShadowBlueAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowBlueAlpha(scale))
	return c
}

func (c *Comp) DropShadowBlueDark(scale int) *Comp {
	c.El.Style(styles.DropShadowBlueDark(scale))
	return c
}

func (c *Comp) DropShadowBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowBlueDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowBronze(scale int) *Comp {
	c.El.Style(styles.DropShadowBronze(scale))
	return c
}

func (c *Comp) DropShadowBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowBronzeAlpha(scale))
	return c
}

func (c *Comp) DropShadowBronzeDark(scale int) *Comp {
	c.El.Style(styles.DropShadowBronzeDark(scale))
	return c
}

func (c *Comp) DropShadowBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowBrown(scale int) *Comp {
	c.El.Style(styles.DropShadowBrown(scale))
	return c
}

func (c *Comp) DropShadowBrownAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowBrownAlpha(scale))
	return c
}

func (c *Comp) DropShadowBrownDark(scale int) *Comp {
	c.El.Style(styles.DropShadowBrownDark(scale))
	return c
}

func (c *Comp) DropShadowBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowBrownDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowColor(val value.Value) *Comp {
	c.El.Style(styles.DropShadowColor(val))
	return c
}

func (c *Comp) DropShadowCrimson(scale int) *Comp {
	c.El.Style(styles.DropShadowCrimson(scale))
	return c
}

func (c *Comp) DropShadowCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowCrimsonAlpha(scale))
	return c
}

func (c *Comp) DropShadowCrimsonDark(scale int) *Comp {
	c.El.Style(styles.DropShadowCrimsonDark(scale))
	return c
}

func (c *Comp) DropShadowCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowCurrent() *Comp {
	c.El.Style(styles.DropShadowCurrent())
	return c
}

func (c *Comp) DropShadowCyan(scale int) *Comp {
	c.El.Style(styles.DropShadowCyan(scale))
	return c
}

func (c *Comp) DropShadowCyanAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowCyanAlpha(scale))
	return c
}

func (c *Comp) DropShadowCyanDark(scale int) *Comp {
	c.El.Style(styles.DropShadowCyanDark(scale))
	return c
}

func (c *Comp) DropShadowCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowCyanDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowGold(scale int) *Comp {
	c.El.Style(styles.DropShadowGold(scale))
	return c
}

func (c *Comp) DropShadowGoldAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowGoldAlpha(scale))
	return c
}

func (c *Comp) DropShadowGoldDark(scale int) *Comp {
	c.El.Style(styles.DropShadowGoldDark(scale))
	return c
}

func (c *Comp) DropShadowGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowGoldDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowGrass(scale int) *Comp {
	c.El.Style(styles.DropShadowGrass(scale))
	return c
}

func (c *Comp) DropShadowGrassAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowGrassAlpha(scale))
	return c
}

func (c *Comp) DropShadowGrassDark(scale int) *Comp {
	c.El.Style(styles.DropShadowGrassDark(scale))
	return c
}

func (c *Comp) DropShadowGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowGrassDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowGray(scale int) *Comp {
	c.El.Style(styles.DropShadowGray(scale))
	return c
}

func (c *Comp) DropShadowGrayAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowGrayAlpha(scale))
	return c
}

func (c *Comp) DropShadowGrayDark(scale int) *Comp {
	c.El.Style(styles.DropShadowGrayDark(scale))
	return c
}

func (c *Comp) DropShadowGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowGrayDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowGreen(scale int) *Comp {
	c.El.Style(styles.DropShadowGreen(scale))
	return c
}

func (c *Comp) DropShadowGreenAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowGreenAlpha(scale))
	return c
}

func (c *Comp) DropShadowGreenDark(scale int) *Comp {
	c.El.Style(styles.DropShadowGreenDark(scale))
	return c
}

func (c *Comp) DropShadowGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowGreenDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowIndigo(scale int) *Comp {
	c.El.Style(styles.DropShadowIndigo(scale))
	return c
}

func (c *Comp) DropShadowIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowIndigoAlpha(scale))
	return c
}

func (c *Comp) DropShadowIndigoDark(scale int) *Comp {
	c.El.Style(styles.DropShadowIndigoDark(scale))
	return c
}

func (c *Comp) DropShadowIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowInherit() *Comp {
	c.El.Style(styles.DropShadowInherit())
	return c
}

func (c *Comp) DropShadowIris(scale int) *Comp {
	c.El.Style(styles.DropShadowIris(scale))
	return c
}

func (c *Comp) DropShadowIrisAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowIrisAlpha(scale))
	return c
}

func (c *Comp) DropShadowIrisDark(scale int) *Comp {
	c.El.Style(styles.DropShadowIrisDark(scale))
	return c
}

func (c *Comp) DropShadowIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowIrisDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowJade(scale int) *Comp {
	c.El.Style(styles.DropShadowJade(scale))
	return c
}

func (c *Comp) DropShadowJadeAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowJadeAlpha(scale))
	return c
}

func (c *Comp) DropShadowJadeDark(scale int) *Comp {
	c.El.Style(styles.DropShadowJadeDark(scale))
	return c
}

func (c *Comp) DropShadowJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowJadeDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowLg() *Comp {
	c.El.Style(styles.DropShadowLg())
	return c
}

func (c *Comp) DropShadowLime(scale int) *Comp {
	c.El.Style(styles.DropShadowLime(scale))
	return c
}

func (c *Comp) DropShadowLimeAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowLimeAlpha(scale))
	return c
}

func (c *Comp) DropShadowLimeDark(scale int) *Comp {
	c.El.Style(styles.DropShadowLimeDark(scale))
	return c
}

func (c *Comp) DropShadowLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowLimeDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowMauve(scale int) *Comp {
	c.El.Style(styles.DropShadowMauve(scale))
	return c
}

func (c *Comp) DropShadowMauveAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowMauveAlpha(scale))
	return c
}

func (c *Comp) DropShadowMauveDark(scale int) *Comp {
	c.El.Style(styles.DropShadowMauveDark(scale))
	return c
}

func (c *Comp) DropShadowMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowMauveDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowMd() *Comp {
	c.El.Style(styles.DropShadowMd())
	return c
}

func (c *Comp) DropShadowMint(scale int) *Comp {
	c.El.Style(styles.DropShadowMint(scale))
	return c
}

func (c *Comp) DropShadowMintAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowMintAlpha(scale))
	return c
}

func (c *Comp) DropShadowMintDark(scale int) *Comp {
	c.El.Style(styles.DropShadowMintDark(scale))
	return c
}

func (c *Comp) DropShadowMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowMintDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowNone() *Comp {
	c.El.Style(styles.DropShadowNone())
	return c
}

func (c *Comp) DropShadowOlive(scale int) *Comp {
	c.El.Style(styles.DropShadowOlive(scale))
	return c
}

func (c *Comp) DropShadowOliveAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowOliveAlpha(scale))
	return c
}

func (c *Comp) DropShadowOliveDark(scale int) *Comp {
	c.El.Style(styles.DropShadowOliveDark(scale))
	return c
}

func (c *Comp) DropShadowOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowOliveDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowOrange(scale int) *Comp {
	c.El.Style(styles.DropShadowOrange(scale))
	return c
}

func (c *Comp) DropShadowOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowOrangeAlpha(scale))
	return c
}

func (c *Comp) DropShadowOrangeDark(scale int) *Comp {
	c.El.Style(styles.DropShadowOrangeDark(scale))
	return c
}

func (c *Comp) DropShadowOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowPink(scale int) *Comp {
	c.El.Style(styles.DropShadowPink(scale))
	return c
}

func (c *Comp) DropShadowPinkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowPinkAlpha(scale))
	return c
}

func (c *Comp) DropShadowPinkDark(scale int) *Comp {
	c.El.Style(styles.DropShadowPinkDark(scale))
	return c
}

func (c *Comp) DropShadowPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowPinkDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowPlum(scale int) *Comp {
	c.El.Style(styles.DropShadowPlum(scale))
	return c
}

func (c *Comp) DropShadowPlumAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowPlumAlpha(scale))
	return c
}

func (c *Comp) DropShadowPlumDark(scale int) *Comp {
	c.El.Style(styles.DropShadowPlumDark(scale))
	return c
}

func (c *Comp) DropShadowPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowPlumDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowPurple(scale int) *Comp {
	c.El.Style(styles.DropShadowPurple(scale))
	return c
}

func (c *Comp) DropShadowPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowPurpleAlpha(scale))
	return c
}

func (c *Comp) DropShadowPurpleDark(scale int) *Comp {
	c.El.Style(styles.DropShadowPurpleDark(scale))
	return c
}

func (c *Comp) DropShadowPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowRed(scale int) *Comp {
	c.El.Style(styles.DropShadowRed(scale))
	return c
}

func (c *Comp) DropShadowRedAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowRedAlpha(scale))
	return c
}

func (c *Comp) DropShadowRedDark(scale int) *Comp {
	c.El.Style(styles.DropShadowRedDark(scale))
	return c
}

func (c *Comp) DropShadowRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowRedDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowRuby(scale int) *Comp {
	c.El.Style(styles.DropShadowRuby(scale))
	return c
}

func (c *Comp) DropShadowRubyAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowRubyAlpha(scale))
	return c
}

func (c *Comp) DropShadowRubyDark(scale int) *Comp {
	c.El.Style(styles.DropShadowRubyDark(scale))
	return c
}

func (c *Comp) DropShadowRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowRubyDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowSage(scale int) *Comp {
	c.El.Style(styles.DropShadowSage(scale))
	return c
}

func (c *Comp) DropShadowSageAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowSageAlpha(scale))
	return c
}

func (c *Comp) DropShadowSageDark(scale int) *Comp {
	c.El.Style(styles.DropShadowSageDark(scale))
	return c
}

func (c *Comp) DropShadowSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowSageDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowSand(scale int) *Comp {
	c.El.Style(styles.DropShadowSand(scale))
	return c
}

func (c *Comp) DropShadowSandAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowSandAlpha(scale))
	return c
}

func (c *Comp) DropShadowSandDark(scale int) *Comp {
	c.El.Style(styles.DropShadowSandDark(scale))
	return c
}

func (c *Comp) DropShadowSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowSandDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowSky(scale int) *Comp {
	c.El.Style(styles.DropShadowSky(scale))
	return c
}

func (c *Comp) DropShadowSkyAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowSkyAlpha(scale))
	return c
}

func (c *Comp) DropShadowSkyDark(scale int) *Comp {
	c.El.Style(styles.DropShadowSkyDark(scale))
	return c
}

func (c *Comp) DropShadowSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowSkyDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowSlate(scale int) *Comp {
	c.El.Style(styles.DropShadowSlate(scale))
	return c
}

func (c *Comp) DropShadowSlateAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowSlateAlpha(scale))
	return c
}

func (c *Comp) DropShadowSlateDark(scale int) *Comp {
	c.El.Style(styles.DropShadowSlateDark(scale))
	return c
}

func (c *Comp) DropShadowSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowSlateDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowSm() *Comp {
	c.El.Style(styles.DropShadowSm())
	return c
}

func (c *Comp) DropShadowTeal(scale int) *Comp {
	c.El.Style(styles.DropShadowTeal(scale))
	return c
}

func (c *Comp) DropShadowTealAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowTealAlpha(scale))
	return c
}

func (c *Comp) DropShadowTealDark(scale int) *Comp {
	c.El.Style(styles.DropShadowTealDark(scale))
	return c
}

func (c *Comp) DropShadowTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowTealDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowTomato(scale int) *Comp {
	c.El.Style(styles.DropShadowTomato(scale))
	return c
}

func (c *Comp) DropShadowTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowTomatoAlpha(scale))
	return c
}

func (c *Comp) DropShadowTomatoDark(scale int) *Comp {
	c.El.Style(styles.DropShadowTomatoDark(scale))
	return c
}

func (c *Comp) DropShadowTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowTransparent() *Comp {
	c.El.Style(styles.DropShadowTransparent())
	return c
}

func (c *Comp) DropShadowViolet(scale int) *Comp {
	c.El.Style(styles.DropShadowViolet(scale))
	return c
}

func (c *Comp) DropShadowVioletAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowVioletAlpha(scale))
	return c
}

func (c *Comp) DropShadowVioletDark(scale int) *Comp {
	c.El.Style(styles.DropShadowVioletDark(scale))
	return c
}

func (c *Comp) DropShadowVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowVioletDarkAlpha(scale))
	return c
}

func (c *Comp) DropShadowWhite() *Comp {
	c.El.Style(styles.DropShadowWhite())
	return c
}

func (c *Comp) DropShadowWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowWhiteAlpha(scale))
	return c
}

func (c *Comp) DropShadowXl() *Comp {
	c.El.Style(styles.DropShadowXl())
	return c
}

func (c *Comp) DropShadowXs() *Comp {
	c.El.Style(styles.DropShadowXs())
	return c
}

func (c *Comp) DropShadowYellow(scale int) *Comp {
	c.El.Style(styles.DropShadowYellow(scale))
	return c
}

func (c *Comp) DropShadowYellowAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowYellowAlpha(scale))
	return c
}

func (c *Comp) DropShadowYellowDark(scale int) *Comp {
	c.El.Style(styles.DropShadowYellowDark(scale))
	return c
}

func (c *Comp) DropShadowYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.DropShadowYellowDarkAlpha(scale))
	return c
}

func (c *Comp) Duration(val any) *Comp {
	c.El.Style(styles.Duration(val))
	return c
}

func (c *Comp) DurationInitial() *Comp {
	c.El.Style(styles.DurationInitial())
	return c
}

func (c *Comp) Ease(val value.Value) *Comp {
	c.El.Style(styles.Ease(val))
	return c
}

func (c *Comp) EaseIn() *Comp {
	c.El.Style(styles.EaseIn())
	return c
}

func (c *Comp) EaseInOut() *Comp {
	c.El.Style(styles.EaseInOut())
	return c
}

func (c *Comp) EaseInitial() *Comp {
	c.El.Style(styles.EaseInitial())
	return c
}

func (c *Comp) EaseLinear() *Comp {
	c.El.Style(styles.EaseLinear())
	return c
}

func (c *Comp) EaseOut() *Comp {
	c.El.Style(styles.EaseOut())
	return c
}

func (c *Comp) End(number int) *Comp {
	c.El.Style(styles.End(number))
	return c
}

func (c *Comp) EndAuto() *Comp {
	c.El.Style(styles.EndAuto())
	return c
}

func (c *Comp) EndBy(val value.Value) *Comp {
	c.El.Style(styles.EndBy(val))
	return c
}

func (c *Comp) EndFraction(fraction float64) *Comp {
	c.El.Style(styles.EndFraction(fraction))
	return c
}

func (c *Comp) EndFull() *Comp {
	c.El.Style(styles.EndFull())
	return c
}

func (c *Comp) EndPx() *Comp {
	c.El.Style(styles.EndPx())
	return c
}

func (c *Comp) FieldSizingContent() *Comp {
	c.El.Style(styles.FieldSizingContent())
	return c
}

func (c *Comp) FieldSizingFixed() *Comp {
	c.El.Style(styles.FieldSizingFixed())
	return c
}

func (c *Comp) Fill(val value.Value) *Comp {
	c.El.Style(styles.Fill(val))
	return c
}

func (c *Comp) FillAmber(scale int) *Comp {
	c.El.Style(styles.FillAmber(scale))
	return c
}

func (c *Comp) FillAmberAlpha(scale int) *Comp {
	c.El.Style(styles.FillAmberAlpha(scale))
	return c
}

func (c *Comp) FillAmberDark(scale int) *Comp {
	c.El.Style(styles.FillAmberDark(scale))
	return c
}

func (c *Comp) FillAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillAmberDarkAlpha(scale))
	return c
}

func (c *Comp) FillBlack() *Comp {
	c.El.Style(styles.FillBlack())
	return c
}

func (c *Comp) FillBlackAlpha(scale int) *Comp {
	c.El.Style(styles.FillBlackAlpha(scale))
	return c
}

func (c *Comp) FillBlue(scale int) *Comp {
	c.El.Style(styles.FillBlue(scale))
	return c
}

func (c *Comp) FillBlueAlpha(scale int) *Comp {
	c.El.Style(styles.FillBlueAlpha(scale))
	return c
}

func (c *Comp) FillBlueDark(scale int) *Comp {
	c.El.Style(styles.FillBlueDark(scale))
	return c
}

func (c *Comp) FillBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillBlueDarkAlpha(scale))
	return c
}

func (c *Comp) FillBronze(scale int) *Comp {
	c.El.Style(styles.FillBronze(scale))
	return c
}

func (c *Comp) FillBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.FillBronzeAlpha(scale))
	return c
}

func (c *Comp) FillBronzeDark(scale int) *Comp {
	c.El.Style(styles.FillBronzeDark(scale))
	return c
}

func (c *Comp) FillBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) FillBrown(scale int) *Comp {
	c.El.Style(styles.FillBrown(scale))
	return c
}

func (c *Comp) FillBrownAlpha(scale int) *Comp {
	c.El.Style(styles.FillBrownAlpha(scale))
	return c
}

func (c *Comp) FillBrownDark(scale int) *Comp {
	c.El.Style(styles.FillBrownDark(scale))
	return c
}

func (c *Comp) FillBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillBrownDarkAlpha(scale))
	return c
}

func (c *Comp) FillCrimson(scale int) *Comp {
	c.El.Style(styles.FillCrimson(scale))
	return c
}

func (c *Comp) FillCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.FillCrimsonAlpha(scale))
	return c
}

func (c *Comp) FillCrimsonDark(scale int) *Comp {
	c.El.Style(styles.FillCrimsonDark(scale))
	return c
}

func (c *Comp) FillCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) FillCurrent() *Comp {
	c.El.Style(styles.FillCurrent())
	return c
}

func (c *Comp) FillCyan(scale int) *Comp {
	c.El.Style(styles.FillCyan(scale))
	return c
}

func (c *Comp) FillCyanAlpha(scale int) *Comp {
	c.El.Style(styles.FillCyanAlpha(scale))
	return c
}

func (c *Comp) FillCyanDark(scale int) *Comp {
	c.El.Style(styles.FillCyanDark(scale))
	return c
}

func (c *Comp) FillCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillCyanDarkAlpha(scale))
	return c
}

func (c *Comp) FillGold(scale int) *Comp {
	c.El.Style(styles.FillGold(scale))
	return c
}

func (c *Comp) FillGoldAlpha(scale int) *Comp {
	c.El.Style(styles.FillGoldAlpha(scale))
	return c
}

func (c *Comp) FillGoldDark(scale int) *Comp {
	c.El.Style(styles.FillGoldDark(scale))
	return c
}

func (c *Comp) FillGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillGoldDarkAlpha(scale))
	return c
}

func (c *Comp) FillGrass(scale int) *Comp {
	c.El.Style(styles.FillGrass(scale))
	return c
}

func (c *Comp) FillGrassAlpha(scale int) *Comp {
	c.El.Style(styles.FillGrassAlpha(scale))
	return c
}

func (c *Comp) FillGrassDark(scale int) *Comp {
	c.El.Style(styles.FillGrassDark(scale))
	return c
}

func (c *Comp) FillGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillGrassDarkAlpha(scale))
	return c
}

func (c *Comp) FillGray(scale int) *Comp {
	c.El.Style(styles.FillGray(scale))
	return c
}

func (c *Comp) FillGrayAlpha(scale int) *Comp {
	c.El.Style(styles.FillGrayAlpha(scale))
	return c
}

func (c *Comp) FillGrayDark(scale int) *Comp {
	c.El.Style(styles.FillGrayDark(scale))
	return c
}

func (c *Comp) FillGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillGrayDarkAlpha(scale))
	return c
}

func (c *Comp) FillGreen(scale int) *Comp {
	c.El.Style(styles.FillGreen(scale))
	return c
}

func (c *Comp) FillGreenAlpha(scale int) *Comp {
	c.El.Style(styles.FillGreenAlpha(scale))
	return c
}

func (c *Comp) FillGreenDark(scale int) *Comp {
	c.El.Style(styles.FillGreenDark(scale))
	return c
}

func (c *Comp) FillGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillGreenDarkAlpha(scale))
	return c
}

func (c *Comp) FillIndigo(scale int) *Comp {
	c.El.Style(styles.FillIndigo(scale))
	return c
}

func (c *Comp) FillIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.FillIndigoAlpha(scale))
	return c
}

func (c *Comp) FillIndigoDark(scale int) *Comp {
	c.El.Style(styles.FillIndigoDark(scale))
	return c
}

func (c *Comp) FillIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) FillInherit() *Comp {
	c.El.Style(styles.FillInherit())
	return c
}

func (c *Comp) FillIris(scale int) *Comp {
	c.El.Style(styles.FillIris(scale))
	return c
}

func (c *Comp) FillIrisAlpha(scale int) *Comp {
	c.El.Style(styles.FillIrisAlpha(scale))
	return c
}

func (c *Comp) FillIrisDark(scale int) *Comp {
	c.El.Style(styles.FillIrisDark(scale))
	return c
}

func (c *Comp) FillIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillIrisDarkAlpha(scale))
	return c
}

func (c *Comp) FillJade(scale int) *Comp {
	c.El.Style(styles.FillJade(scale))
	return c
}

func (c *Comp) FillJadeAlpha(scale int) *Comp {
	c.El.Style(styles.FillJadeAlpha(scale))
	return c
}

func (c *Comp) FillJadeDark(scale int) *Comp {
	c.El.Style(styles.FillJadeDark(scale))
	return c
}

func (c *Comp) FillJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillJadeDarkAlpha(scale))
	return c
}

func (c *Comp) FillLime(scale int) *Comp {
	c.El.Style(styles.FillLime(scale))
	return c
}

func (c *Comp) FillLimeAlpha(scale int) *Comp {
	c.El.Style(styles.FillLimeAlpha(scale))
	return c
}

func (c *Comp) FillLimeDark(scale int) *Comp {
	c.El.Style(styles.FillLimeDark(scale))
	return c
}

func (c *Comp) FillLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillLimeDarkAlpha(scale))
	return c
}

func (c *Comp) FillMauve(scale int) *Comp {
	c.El.Style(styles.FillMauve(scale))
	return c
}

func (c *Comp) FillMauveAlpha(scale int) *Comp {
	c.El.Style(styles.FillMauveAlpha(scale))
	return c
}

func (c *Comp) FillMauveDark(scale int) *Comp {
	c.El.Style(styles.FillMauveDark(scale))
	return c
}

func (c *Comp) FillMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillMauveDarkAlpha(scale))
	return c
}

func (c *Comp) FillMint(scale int) *Comp {
	c.El.Style(styles.FillMint(scale))
	return c
}

func (c *Comp) FillMintAlpha(scale int) *Comp {
	c.El.Style(styles.FillMintAlpha(scale))
	return c
}

func (c *Comp) FillMintDark(scale int) *Comp {
	c.El.Style(styles.FillMintDark(scale))
	return c
}

func (c *Comp) FillMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillMintDarkAlpha(scale))
	return c
}

func (c *Comp) FillOlive(scale int) *Comp {
	c.El.Style(styles.FillOlive(scale))
	return c
}

func (c *Comp) FillOliveAlpha(scale int) *Comp {
	c.El.Style(styles.FillOliveAlpha(scale))
	return c
}

func (c *Comp) FillOliveDark(scale int) *Comp {
	c.El.Style(styles.FillOliveDark(scale))
	return c
}

func (c *Comp) FillOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillOliveDarkAlpha(scale))
	return c
}

func (c *Comp) FillOrange(scale int) *Comp {
	c.El.Style(styles.FillOrange(scale))
	return c
}

func (c *Comp) FillOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.FillOrangeAlpha(scale))
	return c
}

func (c *Comp) FillOrangeDark(scale int) *Comp {
	c.El.Style(styles.FillOrangeDark(scale))
	return c
}

func (c *Comp) FillOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) FillPink(scale int) *Comp {
	c.El.Style(styles.FillPink(scale))
	return c
}

func (c *Comp) FillPinkAlpha(scale int) *Comp {
	c.El.Style(styles.FillPinkAlpha(scale))
	return c
}

func (c *Comp) FillPinkDark(scale int) *Comp {
	c.El.Style(styles.FillPinkDark(scale))
	return c
}

func (c *Comp) FillPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillPinkDarkAlpha(scale))
	return c
}

func (c *Comp) FillPlum(scale int) *Comp {
	c.El.Style(styles.FillPlum(scale))
	return c
}

func (c *Comp) FillPlumAlpha(scale int) *Comp {
	c.El.Style(styles.FillPlumAlpha(scale))
	return c
}

func (c *Comp) FillPlumDark(scale int) *Comp {
	c.El.Style(styles.FillPlumDark(scale))
	return c
}

func (c *Comp) FillPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillPlumDarkAlpha(scale))
	return c
}

func (c *Comp) FillPurple(scale int) *Comp {
	c.El.Style(styles.FillPurple(scale))
	return c
}

func (c *Comp) FillPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.FillPurpleAlpha(scale))
	return c
}

func (c *Comp) FillPurpleDark(scale int) *Comp {
	c.El.Style(styles.FillPurpleDark(scale))
	return c
}

func (c *Comp) FillPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) FillRed(scale int) *Comp {
	c.El.Style(styles.FillRed(scale))
	return c
}

func (c *Comp) FillRedAlpha(scale int) *Comp {
	c.El.Style(styles.FillRedAlpha(scale))
	return c
}

func (c *Comp) FillRedDark(scale int) *Comp {
	c.El.Style(styles.FillRedDark(scale))
	return c
}

func (c *Comp) FillRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillRedDarkAlpha(scale))
	return c
}

func (c *Comp) FillRuby(scale int) *Comp {
	c.El.Style(styles.FillRuby(scale))
	return c
}

func (c *Comp) FillRubyAlpha(scale int) *Comp {
	c.El.Style(styles.FillRubyAlpha(scale))
	return c
}

func (c *Comp) FillRubyDark(scale int) *Comp {
	c.El.Style(styles.FillRubyDark(scale))
	return c
}

func (c *Comp) FillRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillRubyDarkAlpha(scale))
	return c
}

func (c *Comp) FillSage(scale int) *Comp {
	c.El.Style(styles.FillSage(scale))
	return c
}

func (c *Comp) FillSageAlpha(scale int) *Comp {
	c.El.Style(styles.FillSageAlpha(scale))
	return c
}

func (c *Comp) FillSageDark(scale int) *Comp {
	c.El.Style(styles.FillSageDark(scale))
	return c
}

func (c *Comp) FillSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillSageDarkAlpha(scale))
	return c
}

func (c *Comp) FillSand(scale int) *Comp {
	c.El.Style(styles.FillSand(scale))
	return c
}

func (c *Comp) FillSandAlpha(scale int) *Comp {
	c.El.Style(styles.FillSandAlpha(scale))
	return c
}

func (c *Comp) FillSandDark(scale int) *Comp {
	c.El.Style(styles.FillSandDark(scale))
	return c
}

func (c *Comp) FillSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillSandDarkAlpha(scale))
	return c
}

func (c *Comp) FillSky(scale int) *Comp {
	c.El.Style(styles.FillSky(scale))
	return c
}

func (c *Comp) FillSkyAlpha(scale int) *Comp {
	c.El.Style(styles.FillSkyAlpha(scale))
	return c
}

func (c *Comp) FillSkyDark(scale int) *Comp {
	c.El.Style(styles.FillSkyDark(scale))
	return c
}

func (c *Comp) FillSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillSkyDarkAlpha(scale))
	return c
}

func (c *Comp) FillSlate(scale int) *Comp {
	c.El.Style(styles.FillSlate(scale))
	return c
}

func (c *Comp) FillSlateAlpha(scale int) *Comp {
	c.El.Style(styles.FillSlateAlpha(scale))
	return c
}

func (c *Comp) FillSlateDark(scale int) *Comp {
	c.El.Style(styles.FillSlateDark(scale))
	return c
}

func (c *Comp) FillSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillSlateDarkAlpha(scale))
	return c
}

func (c *Comp) FillTeal(scale int) *Comp {
	c.El.Style(styles.FillTeal(scale))
	return c
}

func (c *Comp) FillTealAlpha(scale int) *Comp {
	c.El.Style(styles.FillTealAlpha(scale))
	return c
}

func (c *Comp) FillTealDark(scale int) *Comp {
	c.El.Style(styles.FillTealDark(scale))
	return c
}

func (c *Comp) FillTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillTealDarkAlpha(scale))
	return c
}

func (c *Comp) FillTomato(scale int) *Comp {
	c.El.Style(styles.FillTomato(scale))
	return c
}

func (c *Comp) FillTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.FillTomatoAlpha(scale))
	return c
}

func (c *Comp) FillTomatoDark(scale int) *Comp {
	c.El.Style(styles.FillTomatoDark(scale))
	return c
}

func (c *Comp) FillTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) FillTransparent() *Comp {
	c.El.Style(styles.FillTransparent())
	return c
}

func (c *Comp) FillViolet(scale int) *Comp {
	c.El.Style(styles.FillViolet(scale))
	return c
}

func (c *Comp) FillVioletAlpha(scale int) *Comp {
	c.El.Style(styles.FillVioletAlpha(scale))
	return c
}

func (c *Comp) FillVioletDark(scale int) *Comp {
	c.El.Style(styles.FillVioletDark(scale))
	return c
}

func (c *Comp) FillVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillVioletDarkAlpha(scale))
	return c
}

func (c *Comp) FillWhite() *Comp {
	c.El.Style(styles.FillWhite())
	return c
}

func (c *Comp) FillWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.FillWhiteAlpha(scale))
	return c
}

func (c *Comp) FillYellow(scale int) *Comp {
	c.El.Style(styles.FillYellow(scale))
	return c
}

func (c *Comp) FillYellowAlpha(scale int) *Comp {
	c.El.Style(styles.FillYellowAlpha(scale))
	return c
}

func (c *Comp) FillYellowDark(scale int) *Comp {
	c.El.Style(styles.FillYellowDark(scale))
	return c
}

func (c *Comp) FillYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.FillYellowDarkAlpha(scale))
	return c
}

func (c *Comp) Filter(val value.Value) *Comp {
	c.El.Style(styles.Filter(val))
	return c
}

func (c *Comp) FilterNone() *Comp {
	c.El.Style(styles.FilterNone())
	return c
}

func (c *Comp) Fixed() *Comp {
	c.El.Style(styles.Fixed())
	return c
}

func (c *Comp) Flex() *Comp {
	c.El.Style(styles.Flex())
	return c
}

func (c *Comp) FlexAuto() *Comp {
	c.El.Style(styles.FlexAuto())
	return c
}

func (c *Comp) FlexBy(val value.Value) *Comp {
	c.El.Style(styles.FlexBy(val))
	return c
}

func (c *Comp) FlexCol() *Comp {
	c.El.Style(styles.FlexCol())
	return c
}

func (c *Comp) FlexColReverse() *Comp {
	c.El.Style(styles.FlexColReverse())
	return c
}

func (c *Comp) FlexFraction(fraction float64) *Comp {
	c.El.Style(styles.FlexFraction(fraction))
	return c
}

func (c *Comp) FlexInitial() *Comp {
	c.El.Style(styles.FlexInitial())
	return c
}

func (c *Comp) FlexNoWrap() *Comp {
	c.El.Style(styles.FlexNoWrap())
	return c
}

func (c *Comp) FlexNone() *Comp {
	c.El.Style(styles.FlexNone())
	return c
}

func (c *Comp) FlexRow() *Comp {
	c.El.Style(styles.FlexRow())
	return c
}

func (c *Comp) FlexRowReverse() *Comp {
	c.El.Style(styles.FlexRowReverse())
	return c
}

func (c *Comp) FlexWrap() *Comp {
	c.El.Style(styles.FlexWrap())
	return c
}

func (c *Comp) FlexWrapReverse() *Comp {
	c.El.Style(styles.FlexWrapReverse())
	return c
}

func (c *Comp) FloatEnd() *Comp {
	c.El.Style(styles.FloatEnd())
	return c
}

func (c *Comp) FloatLeft() *Comp {
	c.El.Style(styles.FloatLeft())
	return c
}

func (c *Comp) FloatNone() *Comp {
	c.El.Style(styles.FloatNone())
	return c
}

func (c *Comp) FloatRight() *Comp {
	c.El.Style(styles.FloatRight())
	return c
}

func (c *Comp) FloatStart() *Comp {
	c.El.Style(styles.FloatStart())
	return c
}

func (c *Comp) FlowRoot() *Comp {
	c.El.Style(styles.FlowRoot())
	return c
}

func (c *Comp) FontBlack() *Comp {
	c.El.Style(styles.FontBlack())
	return c
}

func (c *Comp) FontBold() *Comp {
	c.El.Style(styles.FontBold())
	return c
}

func (c *Comp) FontExtraBold() *Comp {
	c.El.Style(styles.FontExtraBold())
	return c
}

func (c *Comp) FontExtraLight() *Comp {
	c.El.Style(styles.FontExtraLight())
	return c
}

func (c *Comp) FontFamilyBy(val value.Value) *Comp {
	c.El.Style(styles.FontFamilyBy(val))
	return c
}

func (c *Comp) FontLight() *Comp {
	c.El.Style(styles.FontLight())
	return c
}

func (c *Comp) FontMedium() *Comp {
	c.El.Style(styles.FontMedium())
	return c
}

func (c *Comp) FontMono() *Comp {
	c.El.Style(styles.FontMono())
	return c
}

func (c *Comp) FontNormal() *Comp {
	c.El.Style(styles.FontNormal())
	return c
}

func (c *Comp) FontSans() *Comp {
	c.El.Style(styles.FontSans())
	return c
}

func (c *Comp) FontSemibold() *Comp {
	c.El.Style(styles.FontSemibold())
	return c
}

func (c *Comp) FontSerif() *Comp {
	c.El.Style(styles.FontSerif())
	return c
}

func (c *Comp) FontStretchBy(val value.Value) *Comp {
	c.El.Style(styles.FontStretchBy(val))
	return c
}

func (c *Comp) FontStretchCondensed() *Comp {
	c.El.Style(styles.FontStretchCondensed())
	return c
}

func (c *Comp) FontStretchExpanded() *Comp {
	c.El.Style(styles.FontStretchExpanded())
	return c
}

func (c *Comp) FontStretchExtraCondensed() *Comp {
	c.El.Style(styles.FontStretchExtraCondensed())
	return c
}

func (c *Comp) FontStretchExtraExpanded() *Comp {
	c.El.Style(styles.FontStretchExtraExpanded())
	return c
}

func (c *Comp) FontStretchNormal() *Comp {
	c.El.Style(styles.FontStretchNormal())
	return c
}

func (c *Comp) FontStretchPercentage(percentage int) *Comp {
	c.El.Style(styles.FontStretchPercentage(percentage))
	return c
}

func (c *Comp) FontStretchSemiCondensed() *Comp {
	c.El.Style(styles.FontStretchSemiCondensed())
	return c
}

func (c *Comp) FontStretchSemiExpanded() *Comp {
	c.El.Style(styles.FontStretchSemiExpanded())
	return c
}

func (c *Comp) FontStretchUltraCondensed() *Comp {
	c.El.Style(styles.FontStretchUltraCondensed())
	return c
}

func (c *Comp) FontStretchUltraExpanded() *Comp {
	c.El.Style(styles.FontStretchUltraExpanded())
	return c
}

func (c *Comp) FontThin() *Comp {
	c.El.Style(styles.FontThin())
	return c
}

func (c *Comp) FontWeightBy(val value.Value) *Comp {
	c.El.Style(styles.FontWeightBy(val))
	return c
}

func (c *Comp) From(val value.Value) *Comp {
	c.El.Style(styles.From(val))
	return c
}

func (c *Comp) Gap(number int) *Comp {
	c.El.Style(styles.Gap(number))
	return c
}

func (c *Comp) GapBy(val value.Value) *Comp {
	c.El.Style(styles.GapBy(val))
	return c
}

func (c *Comp) GapX(number int) *Comp {
	c.El.Style(styles.GapX(number))
	return c
}

func (c *Comp) GapXBy(val value.Value) *Comp {
	c.El.Style(styles.GapXBy(val))
	return c
}

func (c *Comp) GapY(number int) *Comp {
	c.El.Style(styles.GapY(number))
	return c
}

func (c *Comp) GapYBy(val value.Value) *Comp {
	c.El.Style(styles.GapYBy(val))
	return c
}

func (c *Comp) Grayscale(val ...any) *Comp {
	c.El.Style(styles.Grayscale(val...))
	return c
}

func (c *Comp) Grid() *Comp {
	c.El.Style(styles.Grid())
	return c
}

func (c *Comp) GridCols(number int) *Comp {
	c.El.Style(styles.GridCols(number))
	return c
}

func (c *Comp) GridColsBy(val value.Value) *Comp {
	c.El.Style(styles.GridColsBy(val))
	return c
}

func (c *Comp) GridColsNone() *Comp {
	c.El.Style(styles.GridColsNone())
	return c
}

func (c *Comp) GridColsSubgrid() *Comp {
	c.El.Style(styles.GridColsSubgrid())
	return c
}

func (c *Comp) GridFlowCol() *Comp {
	c.El.Style(styles.GridFlowCol())
	return c
}

func (c *Comp) GridFlowColDense() *Comp {
	c.El.Style(styles.GridFlowColDense())
	return c
}

func (c *Comp) GridFlowDense() *Comp {
	c.El.Style(styles.GridFlowDense())
	return c
}

func (c *Comp) GridFlowRow() *Comp {
	c.El.Style(styles.GridFlowRow())
	return c
}

func (c *Comp) GridFlowRowDense() *Comp {
	c.El.Style(styles.GridFlowRowDense())
	return c
}

func (c *Comp) GridRows(number int) *Comp {
	c.El.Style(styles.GridRows(number))
	return c
}

func (c *Comp) GridRowsBy(val value.Value) *Comp {
	c.El.Style(styles.GridRowsBy(val))
	return c
}

func (c *Comp) GridRowsNone() *Comp {
	c.El.Style(styles.GridRowsNone())
	return c
}

func (c *Comp) GridRowsSubgrid() *Comp {
	c.El.Style(styles.GridRowsSubgrid())
	return c
}

func (c *Comp) Grow() *Comp {
	c.El.Style(styles.Grow())
	return c
}

func (c *Comp) GrowNumber(number int) *Comp {
	c.El.Style(styles.GrowNumber(number))
	return c
}

func (c *Comp) GrowValue(val value.Value) *Comp {
	c.El.Style(styles.GrowValue(val))
	return c
}

func (c *Comp) H(number int) *Comp {
	c.El.Style(styles.H(number))
	return c
}

func (c *Comp) HAuto() *Comp {
	c.El.Style(styles.HAuto())
	return c
}

func (c *Comp) HDvh() *Comp {
	c.El.Style(styles.HDvh())
	return c
}

func (c *Comp) HDvw() *Comp {
	c.El.Style(styles.HDvw())
	return c
}

func (c *Comp) HFit() *Comp {
	c.El.Style(styles.HFit())
	return c
}

func (c *Comp) HFraction(fraction float32) *Comp {
	c.El.Style(styles.HFraction(fraction))
	return c
}

func (c *Comp) HFull() *Comp {
	c.El.Style(styles.HFull())
	return c
}

func (c *Comp) HLh() *Comp {
	c.El.Style(styles.HLh())
	return c
}

func (c *Comp) HLvh() *Comp {
	c.El.Style(styles.HLvh())
	return c
}

func (c *Comp) HLvw() *Comp {
	c.El.Style(styles.HLvw())
	return c
}

func (c *Comp) HMax() *Comp {
	c.El.Style(styles.HMax())
	return c
}

func (c *Comp) HMin() *Comp {
	c.El.Style(styles.HMin())
	return c
}

func (c *Comp) HPx() *Comp {
	c.El.Style(styles.HPx())
	return c
}

func (c *Comp) HScreen() *Comp {
	c.El.Style(styles.HScreen())
	return c
}

func (c *Comp) HSvh() *Comp {
	c.El.Style(styles.HSvh())
	return c
}

func (c *Comp) HSvw() *Comp {
	c.El.Style(styles.HSvw())
	return c
}

func (c *Comp) HValueBy(val value.Value) *Comp {
	c.El.Style(styles.HValueBy(val))
	return c
}

func (c *Comp) Hidden() *Comp {
	c.El.Style(styles.Hidden())
	return c
}

func (c *Comp) HueRotate(val any) *Comp {
	c.El.Style(styles.HueRotate(val))
	return c
}

func (c *Comp) HyphensAuto() *Comp {
	c.El.Style(styles.HyphensAuto())
	return c
}

func (c *Comp) HyphensManual() *Comp {
	c.El.Style(styles.HyphensManual())
	return c
}

func (c *Comp) HyphensNone() *Comp {
	c.El.Style(styles.HyphensNone())
	return c
}

func (c *Comp) Indent(number int) *Comp {
	c.El.Style(styles.Indent(number))
	return c
}

func (c *Comp) IndentBy(val value.Value) *Comp {
	c.El.Style(styles.IndentBy(val))
	return c
}

func (c *Comp) IndentPx() *Comp {
	c.El.Style(styles.IndentPx())
	return c
}

func (c *Comp) Inline() *Comp {
	c.El.Style(styles.Inline())
	return c
}

func (c *Comp) InlineBlock() *Comp {
	c.El.Style(styles.InlineBlock())
	return c
}

func (c *Comp) InlineFlex() *Comp {
	c.El.Style(styles.InlineFlex())
	return c
}

func (c *Comp) InlineGrid() *Comp {
	c.El.Style(styles.InlineGrid())
	return c
}

func (c *Comp) InlineTable() *Comp {
	c.El.Style(styles.InlineTable())
	return c
}

func (c *Comp) Inset(number int) *Comp {
	c.El.Style(styles.Inset(number))
	return c
}

func (c *Comp) InsetAuto() *Comp {
	c.El.Style(styles.InsetAuto())
	return c
}

func (c *Comp) InsetBy(val value.Value) *Comp {
	c.El.Style(styles.InsetBy(val))
	return c
}

func (c *Comp) InsetFraction(fraction float64) *Comp {
	c.El.Style(styles.InsetFraction(fraction))
	return c
}

func (c *Comp) InsetFull() *Comp {
	c.El.Style(styles.InsetFull())
	return c
}

func (c *Comp) InsetPx() *Comp {
	c.El.Style(styles.InsetPx())
	return c
}

func (c *Comp) InsetRing(val ...value.Value) *Comp {
	c.El.Style(styles.InsetRing(val...))
	return c
}

func (c *Comp) InsetRingAmber(scale int) *Comp {
	c.El.Style(styles.InsetRingAmber(scale))
	return c
}

func (c *Comp) InsetRingAmberAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingAmberAlpha(scale))
	return c
}

func (c *Comp) InsetRingAmberDark(scale int) *Comp {
	c.El.Style(styles.InsetRingAmberDark(scale))
	return c
}

func (c *Comp) InsetRingAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingAmberDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingBlack() *Comp {
	c.El.Style(styles.InsetRingBlack())
	return c
}

func (c *Comp) InsetRingBlackAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingBlackAlpha(scale))
	return c
}

func (c *Comp) InsetRingBlue(scale int) *Comp {
	c.El.Style(styles.InsetRingBlue(scale))
	return c
}

func (c *Comp) InsetRingBlueAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingBlueAlpha(scale))
	return c
}

func (c *Comp) InsetRingBlueDark(scale int) *Comp {
	c.El.Style(styles.InsetRingBlueDark(scale))
	return c
}

func (c *Comp) InsetRingBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingBlueDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingBronze(scale int) *Comp {
	c.El.Style(styles.InsetRingBronze(scale))
	return c
}

func (c *Comp) InsetRingBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingBronzeAlpha(scale))
	return c
}

func (c *Comp) InsetRingBronzeDark(scale int) *Comp {
	c.El.Style(styles.InsetRingBronzeDark(scale))
	return c
}

func (c *Comp) InsetRingBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingBrown(scale int) *Comp {
	c.El.Style(styles.InsetRingBrown(scale))
	return c
}

func (c *Comp) InsetRingBrownAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingBrownAlpha(scale))
	return c
}

func (c *Comp) InsetRingBrownDark(scale int) *Comp {
	c.El.Style(styles.InsetRingBrownDark(scale))
	return c
}

func (c *Comp) InsetRingBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingBrownDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingColor(val value.Value) *Comp {
	c.El.Style(styles.InsetRingColor(val))
	return c
}

func (c *Comp) InsetRingCrimson(scale int) *Comp {
	c.El.Style(styles.InsetRingCrimson(scale))
	return c
}

func (c *Comp) InsetRingCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingCrimsonAlpha(scale))
	return c
}

func (c *Comp) InsetRingCrimsonDark(scale int) *Comp {
	c.El.Style(styles.InsetRingCrimsonDark(scale))
	return c
}

func (c *Comp) InsetRingCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingCurrent() *Comp {
	c.El.Style(styles.InsetRingCurrent())
	return c
}

func (c *Comp) InsetRingCyan(scale int) *Comp {
	c.El.Style(styles.InsetRingCyan(scale))
	return c
}

func (c *Comp) InsetRingCyanAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingCyanAlpha(scale))
	return c
}

func (c *Comp) InsetRingCyanDark(scale int) *Comp {
	c.El.Style(styles.InsetRingCyanDark(scale))
	return c
}

func (c *Comp) InsetRingCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingCyanDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingGold(scale int) *Comp {
	c.El.Style(styles.InsetRingGold(scale))
	return c
}

func (c *Comp) InsetRingGoldAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingGoldAlpha(scale))
	return c
}

func (c *Comp) InsetRingGoldDark(scale int) *Comp {
	c.El.Style(styles.InsetRingGoldDark(scale))
	return c
}

func (c *Comp) InsetRingGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingGoldDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingGrass(scale int) *Comp {
	c.El.Style(styles.InsetRingGrass(scale))
	return c
}

func (c *Comp) InsetRingGrassAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingGrassAlpha(scale))
	return c
}

func (c *Comp) InsetRingGrassDark(scale int) *Comp {
	c.El.Style(styles.InsetRingGrassDark(scale))
	return c
}

func (c *Comp) InsetRingGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingGrassDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingGray(scale int) *Comp {
	c.El.Style(styles.InsetRingGray(scale))
	return c
}

func (c *Comp) InsetRingGrayAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingGrayAlpha(scale))
	return c
}

func (c *Comp) InsetRingGrayDark(scale int) *Comp {
	c.El.Style(styles.InsetRingGrayDark(scale))
	return c
}

func (c *Comp) InsetRingGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingGrayDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingGreen(scale int) *Comp {
	c.El.Style(styles.InsetRingGreen(scale))
	return c
}

func (c *Comp) InsetRingGreenAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingGreenAlpha(scale))
	return c
}

func (c *Comp) InsetRingGreenDark(scale int) *Comp {
	c.El.Style(styles.InsetRingGreenDark(scale))
	return c
}

func (c *Comp) InsetRingGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingGreenDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingIndigo(scale int) *Comp {
	c.El.Style(styles.InsetRingIndigo(scale))
	return c
}

func (c *Comp) InsetRingIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingIndigoAlpha(scale))
	return c
}

func (c *Comp) InsetRingIndigoDark(scale int) *Comp {
	c.El.Style(styles.InsetRingIndigoDark(scale))
	return c
}

func (c *Comp) InsetRingIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingInherit() *Comp {
	c.El.Style(styles.InsetRingInherit())
	return c
}

func (c *Comp) InsetRingIris(scale int) *Comp {
	c.El.Style(styles.InsetRingIris(scale))
	return c
}

func (c *Comp) InsetRingIrisAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingIrisAlpha(scale))
	return c
}

func (c *Comp) InsetRingIrisDark(scale int) *Comp {
	c.El.Style(styles.InsetRingIrisDark(scale))
	return c
}

func (c *Comp) InsetRingIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingIrisDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingJade(scale int) *Comp {
	c.El.Style(styles.InsetRingJade(scale))
	return c
}

func (c *Comp) InsetRingJadeAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingJadeAlpha(scale))
	return c
}

func (c *Comp) InsetRingJadeDark(scale int) *Comp {
	c.El.Style(styles.InsetRingJadeDark(scale))
	return c
}

func (c *Comp) InsetRingJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingJadeDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingLime(scale int) *Comp {
	c.El.Style(styles.InsetRingLime(scale))
	return c
}

func (c *Comp) InsetRingLimeAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingLimeAlpha(scale))
	return c
}

func (c *Comp) InsetRingLimeDark(scale int) *Comp {
	c.El.Style(styles.InsetRingLimeDark(scale))
	return c
}

func (c *Comp) InsetRingLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingLimeDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingMauve(scale int) *Comp {
	c.El.Style(styles.InsetRingMauve(scale))
	return c
}

func (c *Comp) InsetRingMauveAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingMauveAlpha(scale))
	return c
}

func (c *Comp) InsetRingMauveDark(scale int) *Comp {
	c.El.Style(styles.InsetRingMauveDark(scale))
	return c
}

func (c *Comp) InsetRingMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingMauveDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingMint(scale int) *Comp {
	c.El.Style(styles.InsetRingMint(scale))
	return c
}

func (c *Comp) InsetRingMintAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingMintAlpha(scale))
	return c
}

func (c *Comp) InsetRingMintDark(scale int) *Comp {
	c.El.Style(styles.InsetRingMintDark(scale))
	return c
}

func (c *Comp) InsetRingMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingMintDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingOlive(scale int) *Comp {
	c.El.Style(styles.InsetRingOlive(scale))
	return c
}

func (c *Comp) InsetRingOliveAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingOliveAlpha(scale))
	return c
}

func (c *Comp) InsetRingOliveDark(scale int) *Comp {
	c.El.Style(styles.InsetRingOliveDark(scale))
	return c
}

func (c *Comp) InsetRingOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingOliveDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingOrange(scale int) *Comp {
	c.El.Style(styles.InsetRingOrange(scale))
	return c
}

func (c *Comp) InsetRingOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingOrangeAlpha(scale))
	return c
}

func (c *Comp) InsetRingOrangeDark(scale int) *Comp {
	c.El.Style(styles.InsetRingOrangeDark(scale))
	return c
}

func (c *Comp) InsetRingOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingPink(scale int) *Comp {
	c.El.Style(styles.InsetRingPink(scale))
	return c
}

func (c *Comp) InsetRingPinkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingPinkAlpha(scale))
	return c
}

func (c *Comp) InsetRingPinkDark(scale int) *Comp {
	c.El.Style(styles.InsetRingPinkDark(scale))
	return c
}

func (c *Comp) InsetRingPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingPinkDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingPlum(scale int) *Comp {
	c.El.Style(styles.InsetRingPlum(scale))
	return c
}

func (c *Comp) InsetRingPlumAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingPlumAlpha(scale))
	return c
}

func (c *Comp) InsetRingPlumDark(scale int) *Comp {
	c.El.Style(styles.InsetRingPlumDark(scale))
	return c
}

func (c *Comp) InsetRingPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingPlumDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingPurple(scale int) *Comp {
	c.El.Style(styles.InsetRingPurple(scale))
	return c
}

func (c *Comp) InsetRingPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingPurpleAlpha(scale))
	return c
}

func (c *Comp) InsetRingPurpleDark(scale int) *Comp {
	c.El.Style(styles.InsetRingPurpleDark(scale))
	return c
}

func (c *Comp) InsetRingPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingRed(scale int) *Comp {
	c.El.Style(styles.InsetRingRed(scale))
	return c
}

func (c *Comp) InsetRingRedAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingRedAlpha(scale))
	return c
}

func (c *Comp) InsetRingRedDark(scale int) *Comp {
	c.El.Style(styles.InsetRingRedDark(scale))
	return c
}

func (c *Comp) InsetRingRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingRedDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingRuby(scale int) *Comp {
	c.El.Style(styles.InsetRingRuby(scale))
	return c
}

func (c *Comp) InsetRingRubyAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingRubyAlpha(scale))
	return c
}

func (c *Comp) InsetRingRubyDark(scale int) *Comp {
	c.El.Style(styles.InsetRingRubyDark(scale))
	return c
}

func (c *Comp) InsetRingRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingRubyDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingSage(scale int) *Comp {
	c.El.Style(styles.InsetRingSage(scale))
	return c
}

func (c *Comp) InsetRingSageAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingSageAlpha(scale))
	return c
}

func (c *Comp) InsetRingSageDark(scale int) *Comp {
	c.El.Style(styles.InsetRingSageDark(scale))
	return c
}

func (c *Comp) InsetRingSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingSageDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingSand(scale int) *Comp {
	c.El.Style(styles.InsetRingSand(scale))
	return c
}

func (c *Comp) InsetRingSandAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingSandAlpha(scale))
	return c
}

func (c *Comp) InsetRingSandDark(scale int) *Comp {
	c.El.Style(styles.InsetRingSandDark(scale))
	return c
}

func (c *Comp) InsetRingSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingSandDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingSky(scale int) *Comp {
	c.El.Style(styles.InsetRingSky(scale))
	return c
}

func (c *Comp) InsetRingSkyAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingSkyAlpha(scale))
	return c
}

func (c *Comp) InsetRingSkyDark(scale int) *Comp {
	c.El.Style(styles.InsetRingSkyDark(scale))
	return c
}

func (c *Comp) InsetRingSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingSkyDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingSlate(scale int) *Comp {
	c.El.Style(styles.InsetRingSlate(scale))
	return c
}

func (c *Comp) InsetRingSlateAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingSlateAlpha(scale))
	return c
}

func (c *Comp) InsetRingSlateDark(scale int) *Comp {
	c.El.Style(styles.InsetRingSlateDark(scale))
	return c
}

func (c *Comp) InsetRingSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingSlateDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingTeal(scale int) *Comp {
	c.El.Style(styles.InsetRingTeal(scale))
	return c
}

func (c *Comp) InsetRingTealAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingTealAlpha(scale))
	return c
}

func (c *Comp) InsetRingTealDark(scale int) *Comp {
	c.El.Style(styles.InsetRingTealDark(scale))
	return c
}

func (c *Comp) InsetRingTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingTealDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingTomato(scale int) *Comp {
	c.El.Style(styles.InsetRingTomato(scale))
	return c
}

func (c *Comp) InsetRingTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingTomatoAlpha(scale))
	return c
}

func (c *Comp) InsetRingTomatoDark(scale int) *Comp {
	c.El.Style(styles.InsetRingTomatoDark(scale))
	return c
}

func (c *Comp) InsetRingTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingTransparent() *Comp {
	c.El.Style(styles.InsetRingTransparent())
	return c
}

func (c *Comp) InsetRingViolet(scale int) *Comp {
	c.El.Style(styles.InsetRingViolet(scale))
	return c
}

func (c *Comp) InsetRingVioletAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingVioletAlpha(scale))
	return c
}

func (c *Comp) InsetRingVioletDark(scale int) *Comp {
	c.El.Style(styles.InsetRingVioletDark(scale))
	return c
}

func (c *Comp) InsetRingVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingVioletDarkAlpha(scale))
	return c
}

func (c *Comp) InsetRingWhite() *Comp {
	c.El.Style(styles.InsetRingWhite())
	return c
}

func (c *Comp) InsetRingWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingWhiteAlpha(scale))
	return c
}

func (c *Comp) InsetRingYellow(scale int) *Comp {
	c.El.Style(styles.InsetRingYellow(scale))
	return c
}

func (c *Comp) InsetRingYellowAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingYellowAlpha(scale))
	return c
}

func (c *Comp) InsetRingYellowDark(scale int) *Comp {
	c.El.Style(styles.InsetRingYellowDark(scale))
	return c
}

func (c *Comp) InsetRingYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetRingYellowDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadow(val value.Value) *Comp {
	c.El.Style(styles.InsetShadow(val))
	return c
}

func (c *Comp) InsetShadow2xl() *Comp {
	c.El.Style(styles.InsetShadow2xl())
	return c
}

func (c *Comp) InsetShadow2xs() *Comp {
	c.El.Style(styles.InsetShadow2xs())
	return c
}

func (c *Comp) InsetShadowAmber(scale int) *Comp {
	c.El.Style(styles.InsetShadowAmber(scale))
	return c
}

func (c *Comp) InsetShadowAmberAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowAmberAlpha(scale))
	return c
}

func (c *Comp) InsetShadowAmberDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowAmberDark(scale))
	return c
}

func (c *Comp) InsetShadowAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowAmberDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowBlack() *Comp {
	c.El.Style(styles.InsetShadowBlack())
	return c
}

func (c *Comp) InsetShadowBlackAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowBlackAlpha(scale))
	return c
}

func (c *Comp) InsetShadowBlue(scale int) *Comp {
	c.El.Style(styles.InsetShadowBlue(scale))
	return c
}

func (c *Comp) InsetShadowBlueAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowBlueAlpha(scale))
	return c
}

func (c *Comp) InsetShadowBlueDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowBlueDark(scale))
	return c
}

func (c *Comp) InsetShadowBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowBlueDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowBronze(scale int) *Comp {
	c.El.Style(styles.InsetShadowBronze(scale))
	return c
}

func (c *Comp) InsetShadowBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowBronzeAlpha(scale))
	return c
}

func (c *Comp) InsetShadowBronzeDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowBronzeDark(scale))
	return c
}

func (c *Comp) InsetShadowBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowBrown(scale int) *Comp {
	c.El.Style(styles.InsetShadowBrown(scale))
	return c
}

func (c *Comp) InsetShadowBrownAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowBrownAlpha(scale))
	return c
}

func (c *Comp) InsetShadowBrownDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowBrownDark(scale))
	return c
}

func (c *Comp) InsetShadowBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowBrownDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowColor(val value.Value) *Comp {
	c.El.Style(styles.InsetShadowColor(val))
	return c
}

func (c *Comp) InsetShadowCrimson(scale int) *Comp {
	c.El.Style(styles.InsetShadowCrimson(scale))
	return c
}

func (c *Comp) InsetShadowCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowCrimsonAlpha(scale))
	return c
}

func (c *Comp) InsetShadowCrimsonDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowCrimsonDark(scale))
	return c
}

func (c *Comp) InsetShadowCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowCurrent() *Comp {
	c.El.Style(styles.InsetShadowCurrent())
	return c
}

func (c *Comp) InsetShadowCyan(scale int) *Comp {
	c.El.Style(styles.InsetShadowCyan(scale))
	return c
}

func (c *Comp) InsetShadowCyanAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowCyanAlpha(scale))
	return c
}

func (c *Comp) InsetShadowCyanDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowCyanDark(scale))
	return c
}

func (c *Comp) InsetShadowCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowCyanDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowGold(scale int) *Comp {
	c.El.Style(styles.InsetShadowGold(scale))
	return c
}

func (c *Comp) InsetShadowGoldAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowGoldAlpha(scale))
	return c
}

func (c *Comp) InsetShadowGoldDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowGoldDark(scale))
	return c
}

func (c *Comp) InsetShadowGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowGoldDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowGrass(scale int) *Comp {
	c.El.Style(styles.InsetShadowGrass(scale))
	return c
}

func (c *Comp) InsetShadowGrassAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowGrassAlpha(scale))
	return c
}

func (c *Comp) InsetShadowGrassDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowGrassDark(scale))
	return c
}

func (c *Comp) InsetShadowGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowGrassDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowGray(scale int) *Comp {
	c.El.Style(styles.InsetShadowGray(scale))
	return c
}

func (c *Comp) InsetShadowGrayAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowGrayAlpha(scale))
	return c
}

func (c *Comp) InsetShadowGrayDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowGrayDark(scale))
	return c
}

func (c *Comp) InsetShadowGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowGrayDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowGreen(scale int) *Comp {
	c.El.Style(styles.InsetShadowGreen(scale))
	return c
}

func (c *Comp) InsetShadowGreenAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowGreenAlpha(scale))
	return c
}

func (c *Comp) InsetShadowGreenDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowGreenDark(scale))
	return c
}

func (c *Comp) InsetShadowGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowGreenDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowIndigo(scale int) *Comp {
	c.El.Style(styles.InsetShadowIndigo(scale))
	return c
}

func (c *Comp) InsetShadowIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowIndigoAlpha(scale))
	return c
}

func (c *Comp) InsetShadowIndigoDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowIndigoDark(scale))
	return c
}

func (c *Comp) InsetShadowIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowInherit() *Comp {
	c.El.Style(styles.InsetShadowInherit())
	return c
}

func (c *Comp) InsetShadowIris(scale int) *Comp {
	c.El.Style(styles.InsetShadowIris(scale))
	return c
}

func (c *Comp) InsetShadowIrisAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowIrisAlpha(scale))
	return c
}

func (c *Comp) InsetShadowIrisDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowIrisDark(scale))
	return c
}

func (c *Comp) InsetShadowIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowIrisDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowJade(scale int) *Comp {
	c.El.Style(styles.InsetShadowJade(scale))
	return c
}

func (c *Comp) InsetShadowJadeAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowJadeAlpha(scale))
	return c
}

func (c *Comp) InsetShadowJadeDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowJadeDark(scale))
	return c
}

func (c *Comp) InsetShadowJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowJadeDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowLg() *Comp {
	c.El.Style(styles.InsetShadowLg())
	return c
}

func (c *Comp) InsetShadowLime(scale int) *Comp {
	c.El.Style(styles.InsetShadowLime(scale))
	return c
}

func (c *Comp) InsetShadowLimeAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowLimeAlpha(scale))
	return c
}

func (c *Comp) InsetShadowLimeDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowLimeDark(scale))
	return c
}

func (c *Comp) InsetShadowLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowLimeDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowMauve(scale int) *Comp {
	c.El.Style(styles.InsetShadowMauve(scale))
	return c
}

func (c *Comp) InsetShadowMauveAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowMauveAlpha(scale))
	return c
}

func (c *Comp) InsetShadowMauveDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowMauveDark(scale))
	return c
}

func (c *Comp) InsetShadowMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowMauveDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowMd() *Comp {
	c.El.Style(styles.InsetShadowMd())
	return c
}

func (c *Comp) InsetShadowMint(scale int) *Comp {
	c.El.Style(styles.InsetShadowMint(scale))
	return c
}

func (c *Comp) InsetShadowMintAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowMintAlpha(scale))
	return c
}

func (c *Comp) InsetShadowMintDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowMintDark(scale))
	return c
}

func (c *Comp) InsetShadowMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowMintDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowNone() *Comp {
	c.El.Style(styles.InsetShadowNone())
	return c
}

func (c *Comp) InsetShadowOlive(scale int) *Comp {
	c.El.Style(styles.InsetShadowOlive(scale))
	return c
}

func (c *Comp) InsetShadowOliveAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowOliveAlpha(scale))
	return c
}

func (c *Comp) InsetShadowOliveDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowOliveDark(scale))
	return c
}

func (c *Comp) InsetShadowOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowOliveDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowOrange(scale int) *Comp {
	c.El.Style(styles.InsetShadowOrange(scale))
	return c
}

func (c *Comp) InsetShadowOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowOrangeAlpha(scale))
	return c
}

func (c *Comp) InsetShadowOrangeDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowOrangeDark(scale))
	return c
}

func (c *Comp) InsetShadowOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowPink(scale int) *Comp {
	c.El.Style(styles.InsetShadowPink(scale))
	return c
}

func (c *Comp) InsetShadowPinkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowPinkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowPinkDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowPinkDark(scale))
	return c
}

func (c *Comp) InsetShadowPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowPinkDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowPlum(scale int) *Comp {
	c.El.Style(styles.InsetShadowPlum(scale))
	return c
}

func (c *Comp) InsetShadowPlumAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowPlumAlpha(scale))
	return c
}

func (c *Comp) InsetShadowPlumDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowPlumDark(scale))
	return c
}

func (c *Comp) InsetShadowPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowPlumDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowPurple(scale int) *Comp {
	c.El.Style(styles.InsetShadowPurple(scale))
	return c
}

func (c *Comp) InsetShadowPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowPurpleAlpha(scale))
	return c
}

func (c *Comp) InsetShadowPurpleDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowPurpleDark(scale))
	return c
}

func (c *Comp) InsetShadowPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowRed(scale int) *Comp {
	c.El.Style(styles.InsetShadowRed(scale))
	return c
}

func (c *Comp) InsetShadowRedAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowRedAlpha(scale))
	return c
}

func (c *Comp) InsetShadowRedDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowRedDark(scale))
	return c
}

func (c *Comp) InsetShadowRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowRedDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowRuby(scale int) *Comp {
	c.El.Style(styles.InsetShadowRuby(scale))
	return c
}

func (c *Comp) InsetShadowRubyAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowRubyAlpha(scale))
	return c
}

func (c *Comp) InsetShadowRubyDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowRubyDark(scale))
	return c
}

func (c *Comp) InsetShadowRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowRubyDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowSage(scale int) *Comp {
	c.El.Style(styles.InsetShadowSage(scale))
	return c
}

func (c *Comp) InsetShadowSageAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowSageAlpha(scale))
	return c
}

func (c *Comp) InsetShadowSageDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowSageDark(scale))
	return c
}

func (c *Comp) InsetShadowSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowSageDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowSand(scale int) *Comp {
	c.El.Style(styles.InsetShadowSand(scale))
	return c
}

func (c *Comp) InsetShadowSandAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowSandAlpha(scale))
	return c
}

func (c *Comp) InsetShadowSandDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowSandDark(scale))
	return c
}

func (c *Comp) InsetShadowSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowSandDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowSky(scale int) *Comp {
	c.El.Style(styles.InsetShadowSky(scale))
	return c
}

func (c *Comp) InsetShadowSkyAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowSkyAlpha(scale))
	return c
}

func (c *Comp) InsetShadowSkyDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowSkyDark(scale))
	return c
}

func (c *Comp) InsetShadowSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowSkyDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowSlate(scale int) *Comp {
	c.El.Style(styles.InsetShadowSlate(scale))
	return c
}

func (c *Comp) InsetShadowSlateAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowSlateAlpha(scale))
	return c
}

func (c *Comp) InsetShadowSlateDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowSlateDark(scale))
	return c
}

func (c *Comp) InsetShadowSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowSlateDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowSm() *Comp {
	c.El.Style(styles.InsetShadowSm())
	return c
}

func (c *Comp) InsetShadowTeal(scale int) *Comp {
	c.El.Style(styles.InsetShadowTeal(scale))
	return c
}

func (c *Comp) InsetShadowTealAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowTealAlpha(scale))
	return c
}

func (c *Comp) InsetShadowTealDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowTealDark(scale))
	return c
}

func (c *Comp) InsetShadowTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowTealDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowTomato(scale int) *Comp {
	c.El.Style(styles.InsetShadowTomato(scale))
	return c
}

func (c *Comp) InsetShadowTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowTomatoAlpha(scale))
	return c
}

func (c *Comp) InsetShadowTomatoDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowTomatoDark(scale))
	return c
}

func (c *Comp) InsetShadowTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowTransparent() *Comp {
	c.El.Style(styles.InsetShadowTransparent())
	return c
}

func (c *Comp) InsetShadowViolet(scale int) *Comp {
	c.El.Style(styles.InsetShadowViolet(scale))
	return c
}

func (c *Comp) InsetShadowVioletAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowVioletAlpha(scale))
	return c
}

func (c *Comp) InsetShadowVioletDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowVioletDark(scale))
	return c
}

func (c *Comp) InsetShadowVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowVioletDarkAlpha(scale))
	return c
}

func (c *Comp) InsetShadowWhite() *Comp {
	c.El.Style(styles.InsetShadowWhite())
	return c
}

func (c *Comp) InsetShadowWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowWhiteAlpha(scale))
	return c
}

func (c *Comp) InsetShadowXl() *Comp {
	c.El.Style(styles.InsetShadowXl())
	return c
}

func (c *Comp) InsetShadowXs() *Comp {
	c.El.Style(styles.InsetShadowXs())
	return c
}

func (c *Comp) InsetShadowYellow(scale int) *Comp {
	c.El.Style(styles.InsetShadowYellow(scale))
	return c
}

func (c *Comp) InsetShadowYellowAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowYellowAlpha(scale))
	return c
}

func (c *Comp) InsetShadowYellowDark(scale int) *Comp {
	c.El.Style(styles.InsetShadowYellowDark(scale))
	return c
}

func (c *Comp) InsetShadowYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.InsetShadowYellowDarkAlpha(scale))
	return c
}

func (c *Comp) InsetX(number int) *Comp {
	c.El.Style(styles.InsetX(number))
	return c
}

func (c *Comp) InsetXAuto() *Comp {
	c.El.Style(styles.InsetXAuto())
	return c
}

func (c *Comp) InsetXBy(val value.Value) *Comp {
	c.El.Style(styles.InsetXBy(val))
	return c
}

func (c *Comp) InsetXFraction(fraction float64) *Comp {
	c.El.Style(styles.InsetXFraction(fraction))
	return c
}

func (c *Comp) InsetXFull() *Comp {
	c.El.Style(styles.InsetXFull())
	return c
}

func (c *Comp) InsetXPx() *Comp {
	c.El.Style(styles.InsetXPx())
	return c
}

func (c *Comp) InsetY(number int) *Comp {
	c.El.Style(styles.InsetY(number))
	return c
}

func (c *Comp) InsetYAuto() *Comp {
	c.El.Style(styles.InsetYAuto())
	return c
}

func (c *Comp) InsetYBy(val value.Value) *Comp {
	c.El.Style(styles.InsetYBy(val))
	return c
}

func (c *Comp) InsetYFraction(fraction float64) *Comp {
	c.El.Style(styles.InsetYFraction(fraction))
	return c
}

func (c *Comp) InsetYFull() *Comp {
	c.El.Style(styles.InsetYFull())
	return c
}

func (c *Comp) InsetYPx() *Comp {
	c.El.Style(styles.InsetYPx())
	return c
}

func (c *Comp) Invert(val ...any) *Comp {
	c.El.Style(styles.Invert(val...))
	return c
}

func (c *Comp) Invisible() *Comp {
	c.El.Style(styles.Invisible())
	return c
}

func (c *Comp) Isolate() *Comp {
	c.El.Style(styles.Isolate())
	return c
}

func (c *Comp) IsolationAuto() *Comp {
	c.El.Style(styles.IsolationAuto())
	return c
}

func (c *Comp) Italic() *Comp {
	c.El.Style(styles.Italic())
	return c
}

func (c *Comp) ItemsBaseline() *Comp {
	c.El.Style(styles.ItemsBaseline())
	return c
}

func (c *Comp) ItemsBaselineLast() *Comp {
	c.El.Style(styles.ItemsBaselineLast())
	return c
}

func (c *Comp) ItemsCenter() *Comp {
	c.El.Style(styles.ItemsCenter())
	return c
}

func (c *Comp) ItemsCenterSafe() *Comp {
	c.El.Style(styles.ItemsCenterSafe())
	return c
}

func (c *Comp) ItemsEnd() *Comp {
	c.El.Style(styles.ItemsEnd())
	return c
}

func (c *Comp) ItemsEndSafe() *Comp {
	c.El.Style(styles.ItemsEndSafe())
	return c
}

func (c *Comp) ItemsStart() *Comp {
	c.El.Style(styles.ItemsStart())
	return c
}

func (c *Comp) ItemsStretch() *Comp {
	c.El.Style(styles.ItemsStretch())
	return c
}

func (c *Comp) JustifyAround() *Comp {
	c.El.Style(styles.JustifyAround())
	return c
}

func (c *Comp) JustifyBaseline() *Comp {
	c.El.Style(styles.JustifyBaseline())
	return c
}

func (c *Comp) JustifyBetween() *Comp {
	c.El.Style(styles.JustifyBetween())
	return c
}

func (c *Comp) JustifyCenter() *Comp {
	c.El.Style(styles.JustifyCenter())
	return c
}

func (c *Comp) JustifyCenterSafe() *Comp {
	c.El.Style(styles.JustifyCenterSafe())
	return c
}

func (c *Comp) JustifyEnd() *Comp {
	c.El.Style(styles.JustifyEnd())
	return c
}

func (c *Comp) JustifyEndSafe() *Comp {
	c.El.Style(styles.JustifyEndSafe())
	return c
}

func (c *Comp) JustifyEvenly() *Comp {
	c.El.Style(styles.JustifyEvenly())
	return c
}

func (c *Comp) JustifyItemsCenter() *Comp {
	c.El.Style(styles.JustifyItemsCenter())
	return c
}

func (c *Comp) JustifyItemsCenterSafe() *Comp {
	c.El.Style(styles.JustifyItemsCenterSafe())
	return c
}

func (c *Comp) JustifyItemsEnd() *Comp {
	c.El.Style(styles.JustifyItemsEnd())
	return c
}

func (c *Comp) JustifyItemsEndSafe() *Comp {
	c.El.Style(styles.JustifyItemsEndSafe())
	return c
}

func (c *Comp) JustifyItemsNormal() *Comp {
	c.El.Style(styles.JustifyItemsNormal())
	return c
}

func (c *Comp) JustifyItemsStart() *Comp {
	c.El.Style(styles.JustifyItemsStart())
	return c
}

func (c *Comp) JustifyItemsStretch() *Comp {
	c.El.Style(styles.JustifyItemsStretch())
	return c
}

func (c *Comp) JustifyNormal() *Comp {
	c.El.Style(styles.JustifyNormal())
	return c
}

func (c *Comp) JustifySelfAuto() *Comp {
	c.El.Style(styles.JustifySelfAuto())
	return c
}

func (c *Comp) JustifySelfCenter() *Comp {
	c.El.Style(styles.JustifySelfCenter())
	return c
}

func (c *Comp) JustifySelfCenterSafe() *Comp {
	c.El.Style(styles.JustifySelfCenterSafe())
	return c
}

func (c *Comp) JustifySelfEnd() *Comp {
	c.El.Style(styles.JustifySelfEnd())
	return c
}

func (c *Comp) JustifySelfEndSafe() *Comp {
	c.El.Style(styles.JustifySelfEndSafe())
	return c
}

func (c *Comp) JustifySelfStart() *Comp {
	c.El.Style(styles.JustifySelfStart())
	return c
}

func (c *Comp) JustifySelfStretch() *Comp {
	c.El.Style(styles.JustifySelfStretch())
	return c
}

func (c *Comp) JustifyStart() *Comp {
	c.El.Style(styles.JustifyStart())
	return c
}

func (c *Comp) JustifyStretch() *Comp {
	c.El.Style(styles.JustifyStretch())
	return c
}

func (c *Comp) Leading(number int) *Comp {
	c.El.Style(styles.Leading(number))
	return c
}

func (c *Comp) LeadingBy(val value.Value) *Comp {
	c.El.Style(styles.LeadingBy(val))
	return c
}

func (c *Comp) LeadingNone() *Comp {
	c.El.Style(styles.LeadingNone())
	return c
}

func (c *Comp) Left(number int) *Comp {
	c.El.Style(styles.Left(number))
	return c
}

func (c *Comp) LeftAuto() *Comp {
	c.El.Style(styles.LeftAuto())
	return c
}

func (c *Comp) LeftBy(val value.Value) *Comp {
	c.El.Style(styles.LeftBy(val))
	return c
}

func (c *Comp) LeftFraction(fraction float64) *Comp {
	c.El.Style(styles.LeftFraction(fraction))
	return c
}

func (c *Comp) LeftFull() *Comp {
	c.El.Style(styles.LeftFull())
	return c
}

func (c *Comp) LeftPx() *Comp {
	c.El.Style(styles.LeftPx())
	return c
}

func (c *Comp) LineClamp(number int) *Comp {
	c.El.Style(styles.LineClamp(number))
	return c
}

func (c *Comp) LineClampBy(val value.Value) *Comp {
	c.El.Style(styles.LineClampBy(val))
	return c
}

func (c *Comp) LineClampNone() *Comp {
	c.El.Style(styles.LineClampNone())
	return c
}

func (c *Comp) LineThrough() *Comp {
	c.El.Style(styles.LineThrough())
	return c
}

func (c *Comp) LiningNums() *Comp {
	c.El.Style(styles.LiningNums())
	return c
}

func (c *Comp) List(val value.Value) *Comp {
	c.El.Style(styles.List(val))
	return c
}

func (c *Comp) ListDecimal() *Comp {
	c.El.Style(styles.ListDecimal())
	return c
}

func (c *Comp) ListDisc() *Comp {
	c.El.Style(styles.ListDisc())
	return c
}

func (c *Comp) ListImage(val value.Value) *Comp {
	c.El.Style(styles.ListImage(val))
	return c
}

func (c *Comp) ListImageNone() *Comp {
	c.El.Style(styles.ListImageNone())
	return c
}

func (c *Comp) ListInside() *Comp {
	c.El.Style(styles.ListInside())
	return c
}

func (c *Comp) ListItem() *Comp {
	c.El.Style(styles.ListItem())
	return c
}

func (c *Comp) ListNone() *Comp {
	c.El.Style(styles.ListNone())
	return c
}

func (c *Comp) ListOutside() *Comp {
	c.El.Style(styles.ListOutside())
	return c
}

func (c *Comp) Lowercase() *Comp {
	c.El.Style(styles.Lowercase())
	return c
}

func (c *Comp) M(number int) *Comp {
	c.El.Style(styles.M(number))
	return c
}

func (c *Comp) MAuto() *Comp {
	c.El.Style(styles.MAuto())
	return c
}

func (c *Comp) MBy(val value.Value) *Comp {
	c.El.Style(styles.MBy(val))
	return c
}

func (c *Comp) MPx() *Comp {
	c.El.Style(styles.MPx())
	return c
}

func (c *Comp) Mask(val value.Value) *Comp {
	c.El.Style(styles.Mask(val))
	return c
}

func (c *Comp) MaskAdd() *Comp {
	c.El.Style(styles.MaskAdd())
	return c
}

func (c *Comp) MaskAlpha() *Comp {
	c.El.Style(styles.MaskAlpha())
	return c
}

func (c *Comp) MaskAuto() *Comp {
	c.El.Style(styles.MaskAuto())
	return c
}

func (c *Comp) MaskBottom() *Comp {
	c.El.Style(styles.MaskBottom())
	return c
}

func (c *Comp) MaskBottomFrom(val any) *Comp {
	c.El.Style(styles.MaskBottomFrom(val))
	return c
}

func (c *Comp) MaskBottomLeft() *Comp {
	c.El.Style(styles.MaskBottomLeft())
	return c
}

func (c *Comp) MaskBottomRight() *Comp {
	c.El.Style(styles.MaskBottomRight())
	return c
}

func (c *Comp) MaskBottomTo(val any) *Comp {
	c.El.Style(styles.MaskBottomTo(val))
	return c
}

func (c *Comp) MaskCenter() *Comp {
	c.El.Style(styles.MaskCenter())
	return c
}

func (c *Comp) MaskCircle() *Comp {
	c.El.Style(styles.MaskCircle())
	return c
}

func (c *Comp) MaskClipBorder() *Comp {
	c.El.Style(styles.MaskClipBorder())
	return c
}

func (c *Comp) MaskClipContent() *Comp {
	c.El.Style(styles.MaskClipContent())
	return c
}

func (c *Comp) MaskClipFill() *Comp {
	c.El.Style(styles.MaskClipFill())
	return c
}

func (c *Comp) MaskClipPadding() *Comp {
	c.El.Style(styles.MaskClipPadding())
	return c
}

func (c *Comp) MaskClipStroke() *Comp {
	c.El.Style(styles.MaskClipStroke())
	return c
}

func (c *Comp) MaskClipView() *Comp {
	c.El.Style(styles.MaskClipView())
	return c
}

func (c *Comp) MaskConic(number float64) *Comp {
	c.El.Style(styles.MaskConic(number))
	return c
}

func (c *Comp) MaskConicFrom(val any) *Comp {
	c.El.Style(styles.MaskConicFrom(val))
	return c
}

func (c *Comp) MaskConicTo(val any) *Comp {
	c.El.Style(styles.MaskConicTo(val))
	return c
}

func (c *Comp) MaskContain() *Comp {
	c.El.Style(styles.MaskContain())
	return c
}

func (c *Comp) MaskCover() *Comp {
	c.El.Style(styles.MaskCover())
	return c
}

func (c *Comp) MaskEllipse() *Comp {
	c.El.Style(styles.MaskEllipse())
	return c
}

func (c *Comp) MaskExclude() *Comp {
	c.El.Style(styles.MaskExclude())
	return c
}

func (c *Comp) MaskIntersect() *Comp {
	c.El.Style(styles.MaskIntersect())
	return c
}

func (c *Comp) MaskLeft() *Comp {
	c.El.Style(styles.MaskLeft())
	return c
}

func (c *Comp) MaskLeftFrom(val any) *Comp {
	c.El.Style(styles.MaskLeftFrom(val))
	return c
}

func (c *Comp) MaskLeftTo(val any) *Comp {
	c.El.Style(styles.MaskLeftTo(val))
	return c
}

func (c *Comp) MaskLinear(degree int) *Comp {
	c.El.Style(styles.MaskLinear(degree))
	return c
}

func (c *Comp) MaskLinearFrom(val any) *Comp {
	c.El.Style(styles.MaskLinearFrom(val))
	return c
}

func (c *Comp) MaskLinearTo(val any) *Comp {
	c.El.Style(styles.MaskLinearTo(val))
	return c
}

func (c *Comp) MaskLuminance() *Comp {
	c.El.Style(styles.MaskLuminance())
	return c
}

func (c *Comp) MaskMatch() *Comp {
	c.El.Style(styles.MaskMatch())
	return c
}

func (c *Comp) MaskNoClip() *Comp {
	c.El.Style(styles.MaskNoClip())
	return c
}

func (c *Comp) MaskNoRepeat() *Comp {
	c.El.Style(styles.MaskNoRepeat())
	return c
}

func (c *Comp) MaskNone() *Comp {
	c.El.Style(styles.MaskNone())
	return c
}

func (c *Comp) MaskOriginBorder() *Comp {
	c.El.Style(styles.MaskOriginBorder())
	return c
}

func (c *Comp) MaskOriginContent() *Comp {
	c.El.Style(styles.MaskOriginContent())
	return c
}

func (c *Comp) MaskOriginFill() *Comp {
	c.El.Style(styles.MaskOriginFill())
	return c
}

func (c *Comp) MaskOriginPadding() *Comp {
	c.El.Style(styles.MaskOriginPadding())
	return c
}

func (c *Comp) MaskOriginStroke() *Comp {
	c.El.Style(styles.MaskOriginStroke())
	return c
}

func (c *Comp) MaskOriginView() *Comp {
	c.El.Style(styles.MaskOriginView())
	return c
}

func (c *Comp) MaskPosition(val value.Value) *Comp {
	c.El.Style(styles.MaskPosition(val))
	return c
}

func (c *Comp) MaskRadial(val ...any) *Comp {
	c.El.Style(styles.MaskRadial(val...))
	return c
}

func (c *Comp) MaskRadialAtBottom() *Comp {
	c.El.Style(styles.MaskRadialAtBottom())
	return c
}

func (c *Comp) MaskRadialAtBottomLeft() *Comp {
	c.El.Style(styles.MaskRadialAtBottomLeft())
	return c
}

func (c *Comp) MaskRadialAtBottomRight() *Comp {
	c.El.Style(styles.MaskRadialAtBottomRight())
	return c
}

func (c *Comp) MaskRadialAtCenter() *Comp {
	c.El.Style(styles.MaskRadialAtCenter())
	return c
}

func (c *Comp) MaskRadialAtLeft() *Comp {
	c.El.Style(styles.MaskRadialAtLeft())
	return c
}

func (c *Comp) MaskRadialAtRight() *Comp {
	c.El.Style(styles.MaskRadialAtRight())
	return c
}

func (c *Comp) MaskRadialAtTop() *Comp {
	c.El.Style(styles.MaskRadialAtTop())
	return c
}

func (c *Comp) MaskRadialAtTopLeft() *Comp {
	c.El.Style(styles.MaskRadialAtTopLeft())
	return c
}

func (c *Comp) MaskRadialAtTopRight() *Comp {
	c.El.Style(styles.MaskRadialAtTopRight())
	return c
}

func (c *Comp) MaskRadialClosestCorner() *Comp {
	c.El.Style(styles.MaskRadialClosestCorner())
	return c
}

func (c *Comp) MaskRadialClosestSide() *Comp {
	c.El.Style(styles.MaskRadialClosestSide())
	return c
}

func (c *Comp) MaskRadialFarthestCorner() *Comp {
	c.El.Style(styles.MaskRadialFarthestCorner())
	return c
}

func (c *Comp) MaskRadialFarthestSide() *Comp {
	c.El.Style(styles.MaskRadialFarthestSide())
	return c
}

func (c *Comp) MaskRadialFrom(val any) *Comp {
	c.El.Style(styles.MaskRadialFrom(val))
	return c
}

func (c *Comp) MaskRadialTo(val any) *Comp {
	c.El.Style(styles.MaskRadialTo(val))
	return c
}

func (c *Comp) MaskRepeat() *Comp {
	c.El.Style(styles.MaskRepeat())
	return c
}

func (c *Comp) MaskRepeatRound() *Comp {
	c.El.Style(styles.MaskRepeatRound())
	return c
}

func (c *Comp) MaskRepeatSpace() *Comp {
	c.El.Style(styles.MaskRepeatSpace())
	return c
}

func (c *Comp) MaskRepeatX() *Comp {
	c.El.Style(styles.MaskRepeatX())
	return c
}

func (c *Comp) MaskRepeatY() *Comp {
	c.El.Style(styles.MaskRepeatY())
	return c
}

func (c *Comp) MaskRight() *Comp {
	c.El.Style(styles.MaskRight())
	return c
}

func (c *Comp) MaskRightFrom(val any) *Comp {
	c.El.Style(styles.MaskRightFrom(val))
	return c
}

func (c *Comp) MaskRightTo(val any) *Comp {
	c.El.Style(styles.MaskRightTo(val))
	return c
}

func (c *Comp) MaskSize(val value.Value) *Comp {
	c.El.Style(styles.MaskSize(val))
	return c
}

func (c *Comp) MaskSubtract() *Comp {
	c.El.Style(styles.MaskSubtract())
	return c
}

func (c *Comp) MaskTop() *Comp {
	c.El.Style(styles.MaskTop())
	return c
}

func (c *Comp) MaskTopFrom(val any) *Comp {
	c.El.Style(styles.MaskTopFrom(val))
	return c
}

func (c *Comp) MaskTopLeft() *Comp {
	c.El.Style(styles.MaskTopLeft())
	return c
}

func (c *Comp) MaskTopRight() *Comp {
	c.El.Style(styles.MaskTopRight())
	return c
}

func (c *Comp) MaskTopTo(val any) *Comp {
	c.El.Style(styles.MaskTopTo(val))
	return c
}

func (c *Comp) MaskTypeAlpha() *Comp {
	c.El.Style(styles.MaskTypeAlpha())
	return c
}

func (c *Comp) MaskTypeLuminance() *Comp {
	c.El.Style(styles.MaskTypeLuminance())
	return c
}

func (c *Comp) MaskXFrom(val any) *Comp {
	c.El.Style(styles.MaskXFrom(val))
	return c
}

func (c *Comp) MaskXTo(val any) *Comp {
	c.El.Style(styles.MaskXTo(val))
	return c
}

func (c *Comp) MaskYFrom(val any) *Comp {
	c.El.Style(styles.MaskYFrom(val))
	return c
}

func (c *Comp) MaskYTo(val any) *Comp {
	c.El.Style(styles.MaskYTo(val))
	return c
}

func (c *Comp) MaxH(number int) *Comp {
	c.El.Style(styles.MaxH(number))
	return c
}

func (c *Comp) MaxHBy(val value.Value) *Comp {
	c.El.Style(styles.MaxHBy(val))
	return c
}

func (c *Comp) MaxHDvh() *Comp {
	c.El.Style(styles.MaxHDvh())
	return c
}

func (c *Comp) MaxHDvw() *Comp {
	c.El.Style(styles.MaxHDvw())
	return c
}

func (c *Comp) MaxHFit() *Comp {
	c.El.Style(styles.MaxHFit())
	return c
}

func (c *Comp) MaxHFraction(fraction float32) *Comp {
	c.El.Style(styles.MaxHFraction(fraction))
	return c
}

func (c *Comp) MaxHFull() *Comp {
	c.El.Style(styles.MaxHFull())
	return c
}

func (c *Comp) MaxHLh() *Comp {
	c.El.Style(styles.MaxHLh())
	return c
}

func (c *Comp) MaxHLvh() *Comp {
	c.El.Style(styles.MaxHLvh())
	return c
}

func (c *Comp) MaxHLvw() *Comp {
	c.El.Style(styles.MaxHLvw())
	return c
}

func (c *Comp) MaxHMax() *Comp {
	c.El.Style(styles.MaxHMax())
	return c
}

func (c *Comp) MaxHMin() *Comp {
	c.El.Style(styles.MaxHMin())
	return c
}

func (c *Comp) MaxHNone() *Comp {
	c.El.Style(styles.MaxHNone())
	return c
}

func (c *Comp) MaxHPx() *Comp {
	c.El.Style(styles.MaxHPx())
	return c
}

func (c *Comp) MaxHScreen() *Comp {
	c.El.Style(styles.MaxHScreen())
	return c
}

func (c *Comp) MaxHSvh() *Comp {
	c.El.Style(styles.MaxHSvh())
	return c
}

func (c *Comp) MaxHSvw() *Comp {
	c.El.Style(styles.MaxHSvw())
	return c
}

func (c *Comp) MaxW(number int) *Comp {
	c.El.Style(styles.MaxW(number))
	return c
}

func (c *Comp) MaxW2xl() *Comp {
	c.El.Style(styles.MaxW2xl())
	return c
}

func (c *Comp) MaxW2xs() *Comp {
	c.El.Style(styles.MaxW2xs())
	return c
}

func (c *Comp) MaxW3xl() *Comp {
	c.El.Style(styles.MaxW3xl())
	return c
}

func (c *Comp) MaxW3xs() *Comp {
	c.El.Style(styles.MaxW3xs())
	return c
}

func (c *Comp) MaxW4xl() *Comp {
	c.El.Style(styles.MaxW4xl())
	return c
}

func (c *Comp) MaxW5xl() *Comp {
	c.El.Style(styles.MaxW5xl())
	return c
}

func (c *Comp) MaxW6xl() *Comp {
	c.El.Style(styles.MaxW6xl())
	return c
}

func (c *Comp) MaxW7xl() *Comp {
	c.El.Style(styles.MaxW7xl())
	return c
}

func (c *Comp) MaxWBy(val value.Value) *Comp {
	c.El.Style(styles.MaxWBy(val))
	return c
}

func (c *Comp) MaxWDvh() *Comp {
	c.El.Style(styles.MaxWDvh())
	return c
}

func (c *Comp) MaxWDvw() *Comp {
	c.El.Style(styles.MaxWDvw())
	return c
}

func (c *Comp) MaxWFit() *Comp {
	c.El.Style(styles.MaxWFit())
	return c
}

func (c *Comp) MaxWFraction(fraction float32) *Comp {
	c.El.Style(styles.MaxWFraction(fraction))
	return c
}

func (c *Comp) MaxWFull() *Comp {
	c.El.Style(styles.MaxWFull())
	return c
}

func (c *Comp) MaxWLg() *Comp {
	c.El.Style(styles.MaxWLg())
	return c
}

func (c *Comp) MaxWLvh() *Comp {
	c.El.Style(styles.MaxWLvh())
	return c
}

func (c *Comp) MaxWLvw() *Comp {
	c.El.Style(styles.MaxWLvw())
	return c
}

func (c *Comp) MaxWMax() *Comp {
	c.El.Style(styles.MaxWMax())
	return c
}

func (c *Comp) MaxWMd() *Comp {
	c.El.Style(styles.MaxWMd())
	return c
}

func (c *Comp) MaxWMin() *Comp {
	c.El.Style(styles.MaxWMin())
	return c
}

func (c *Comp) MaxWNone() *Comp {
	c.El.Style(styles.MaxWNone())
	return c
}

func (c *Comp) MaxWPx() *Comp {
	c.El.Style(styles.MaxWPx())
	return c
}

func (c *Comp) MaxWScreen() *Comp {
	c.El.Style(styles.MaxWScreen())
	return c
}

func (c *Comp) MaxWSm() *Comp {
	c.El.Style(styles.MaxWSm())
	return c
}

func (c *Comp) MaxWSvh() *Comp {
	c.El.Style(styles.MaxWSvh())
	return c
}

func (c *Comp) MaxWSvw() *Comp {
	c.El.Style(styles.MaxWSvw())
	return c
}

func (c *Comp) MaxWXl() *Comp {
	c.El.Style(styles.MaxWXl())
	return c
}

func (c *Comp) MaxWXs() *Comp {
	c.El.Style(styles.MaxWXs())
	return c
}

func (c *Comp) Mb(number int) *Comp {
	c.El.Style(styles.Mb(number))
	return c
}

func (c *Comp) MbAuto() *Comp {
	c.El.Style(styles.MbAuto())
	return c
}

func (c *Comp) MbBy(val value.Value) *Comp {
	c.El.Style(styles.MbBy(val))
	return c
}

func (c *Comp) MbPx() *Comp {
	c.El.Style(styles.MbPx())
	return c
}

func (c *Comp) Me(number int) *Comp {
	c.El.Style(styles.Me(number))
	return c
}

func (c *Comp) MeAuto() *Comp {
	c.El.Style(styles.MeAuto())
	return c
}

func (c *Comp) MeBy(val value.Value) *Comp {
	c.El.Style(styles.MeBy(val))
	return c
}

func (c *Comp) MePx() *Comp {
	c.El.Style(styles.MePx())
	return c
}

func (c *Comp) MinH(number int) *Comp {
	c.El.Style(styles.MinH(number))
	return c
}

func (c *Comp) MinHAuto() *Comp {
	c.El.Style(styles.MinHAuto())
	return c
}

func (c *Comp) MinHBy(val value.Value) *Comp {
	c.El.Style(styles.MinHBy(val))
	return c
}

func (c *Comp) MinHDvh() *Comp {
	c.El.Style(styles.MinHDvh())
	return c
}

func (c *Comp) MinHDvw() *Comp {
	c.El.Style(styles.MinHDvw())
	return c
}

func (c *Comp) MinHFit() *Comp {
	c.El.Style(styles.MinHFit())
	return c
}

func (c *Comp) MinHFraction(fraction float32) *Comp {
	c.El.Style(styles.MinHFraction(fraction))
	return c
}

func (c *Comp) MinHFull() *Comp {
	c.El.Style(styles.MinHFull())
	return c
}

func (c *Comp) MinHLh() *Comp {
	c.El.Style(styles.MinHLh())
	return c
}

func (c *Comp) MinHLvh() *Comp {
	c.El.Style(styles.MinHLvh())
	return c
}

func (c *Comp) MinHLvw() *Comp {
	c.El.Style(styles.MinHLvw())
	return c
}

func (c *Comp) MinHMax() *Comp {
	c.El.Style(styles.MinHMax())
	return c
}

func (c *Comp) MinHMin() *Comp {
	c.El.Style(styles.MinHMin())
	return c
}

func (c *Comp) MinHPx() *Comp {
	c.El.Style(styles.MinHPx())
	return c
}

func (c *Comp) MinHScreen() *Comp {
	c.El.Style(styles.MinHScreen())
	return c
}

func (c *Comp) MinHSvh() *Comp {
	c.El.Style(styles.MinHSvh())
	return c
}

func (c *Comp) MinHSvw() *Comp {
	c.El.Style(styles.MinHSvw())
	return c
}

func (c *Comp) MinW(number int) *Comp {
	c.El.Style(styles.MinW(number))
	return c
}

func (c *Comp) MinW2xl() *Comp {
	c.El.Style(styles.MinW2xl())
	return c
}

func (c *Comp) MinW2xs() *Comp {
	c.El.Style(styles.MinW2xs())
	return c
}

func (c *Comp) MinW3xl() *Comp {
	c.El.Style(styles.MinW3xl())
	return c
}

func (c *Comp) MinW3xs() *Comp {
	c.El.Style(styles.MinW3xs())
	return c
}

func (c *Comp) MinW4xl() *Comp {
	c.El.Style(styles.MinW4xl())
	return c
}

func (c *Comp) MinW5xl() *Comp {
	c.El.Style(styles.MinW5xl())
	return c
}

func (c *Comp) MinW6xl() *Comp {
	c.El.Style(styles.MinW6xl())
	return c
}

func (c *Comp) MinW7xl() *Comp {
	c.El.Style(styles.MinW7xl())
	return c
}

func (c *Comp) MinWAuto() *Comp {
	c.El.Style(styles.MinWAuto())
	return c
}

func (c *Comp) MinWBy(val value.Value) *Comp {
	c.El.Style(styles.MinWBy(val))
	return c
}

func (c *Comp) MinWDvh() *Comp {
	c.El.Style(styles.MinWDvh())
	return c
}

func (c *Comp) MinWDvw() *Comp {
	c.El.Style(styles.MinWDvw())
	return c
}

func (c *Comp) MinWFit() *Comp {
	c.El.Style(styles.MinWFit())
	return c
}

func (c *Comp) MinWFraction(fraction float32) *Comp {
	c.El.Style(styles.MinWFraction(fraction))
	return c
}

func (c *Comp) MinWFull() *Comp {
	c.El.Style(styles.MinWFull())
	return c
}

func (c *Comp) MinWLg() *Comp {
	c.El.Style(styles.MinWLg())
	return c
}

func (c *Comp) MinWLvh() *Comp {
	c.El.Style(styles.MinWLvh())
	return c
}

func (c *Comp) MinWLvw() *Comp {
	c.El.Style(styles.MinWLvw())
	return c
}

func (c *Comp) MinWMax() *Comp {
	c.El.Style(styles.MinWMax())
	return c
}

func (c *Comp) MinWMd() *Comp {
	c.El.Style(styles.MinWMd())
	return c
}

func (c *Comp) MinWMin() *Comp {
	c.El.Style(styles.MinWMin())
	return c
}

func (c *Comp) MinWPx() *Comp {
	c.El.Style(styles.MinWPx())
	return c
}

func (c *Comp) MinWScreen() *Comp {
	c.El.Style(styles.MinWScreen())
	return c
}

func (c *Comp) MinWSm() *Comp {
	c.El.Style(styles.MinWSm())
	return c
}

func (c *Comp) MinWSvh() *Comp {
	c.El.Style(styles.MinWSvh())
	return c
}

func (c *Comp) MinWSvw() *Comp {
	c.El.Style(styles.MinWSvw())
	return c
}

func (c *Comp) MinWXl() *Comp {
	c.El.Style(styles.MinWXl())
	return c
}

func (c *Comp) MinWXs() *Comp {
	c.El.Style(styles.MinWXs())
	return c
}

func (c *Comp) MixBlendColor() *Comp {
	c.El.Style(styles.MixBlendColor())
	return c
}

func (c *Comp) MixBlendColorBurn() *Comp {
	c.El.Style(styles.MixBlendColorBurn())
	return c
}

func (c *Comp) MixBlendColorDodge() *Comp {
	c.El.Style(styles.MixBlendColorDodge())
	return c
}

func (c *Comp) MixBlendDarken() *Comp {
	c.El.Style(styles.MixBlendDarken())
	return c
}

func (c *Comp) MixBlendDifference() *Comp {
	c.El.Style(styles.MixBlendDifference())
	return c
}

func (c *Comp) MixBlendExclusion() *Comp {
	c.El.Style(styles.MixBlendExclusion())
	return c
}

func (c *Comp) MixBlendHardLight() *Comp {
	c.El.Style(styles.MixBlendHardLight())
	return c
}

func (c *Comp) MixBlendHue() *Comp {
	c.El.Style(styles.MixBlendHue())
	return c
}

func (c *Comp) MixBlendLighten() *Comp {
	c.El.Style(styles.MixBlendLighten())
	return c
}

func (c *Comp) MixBlendLuminosity() *Comp {
	c.El.Style(styles.MixBlendLuminosity())
	return c
}

func (c *Comp) MixBlendMultiply() *Comp {
	c.El.Style(styles.MixBlendMultiply())
	return c
}

func (c *Comp) MixBlendNormal() *Comp {
	c.El.Style(styles.MixBlendNormal())
	return c
}

func (c *Comp) MixBlendOverlay() *Comp {
	c.El.Style(styles.MixBlendOverlay())
	return c
}

func (c *Comp) MixBlendPlusDarker() *Comp {
	c.El.Style(styles.MixBlendPlusDarker())
	return c
}

func (c *Comp) MixBlendPlusLighter() *Comp {
	c.El.Style(styles.MixBlendPlusLighter())
	return c
}

func (c *Comp) MixBlendSaturation() *Comp {
	c.El.Style(styles.MixBlendSaturation())
	return c
}

func (c *Comp) MixBlendScreen() *Comp {
	c.El.Style(styles.MixBlendScreen())
	return c
}

func (c *Comp) MixBlendSoftLight() *Comp {
	c.El.Style(styles.MixBlendSoftLight())
	return c
}

func (c *Comp) Ml(number int) *Comp {
	c.El.Style(styles.Ml(number))
	return c
}

func (c *Comp) MlAuto() *Comp {
	c.El.Style(styles.MlAuto())
	return c
}

func (c *Comp) MlBy(val value.Value) *Comp {
	c.El.Style(styles.MlBy(val))
	return c
}

func (c *Comp) MlPx() *Comp {
	c.El.Style(styles.MlPx())
	return c
}

func (c *Comp) Mr(number int) *Comp {
	c.El.Style(styles.Mr(number))
	return c
}

func (c *Comp) MrAuto() *Comp {
	c.El.Style(styles.MrAuto())
	return c
}

func (c *Comp) MrBy(val value.Value) *Comp {
	c.El.Style(styles.MrBy(val))
	return c
}

func (c *Comp) MrPx() *Comp {
	c.El.Style(styles.MrPx())
	return c
}

func (c *Comp) Ms(number int) *Comp {
	c.El.Style(styles.Ms(number))
	return c
}

func (c *Comp) MsAuto() *Comp {
	c.El.Style(styles.MsAuto())
	return c
}

func (c *Comp) MsBy(val value.Value) *Comp {
	c.El.Style(styles.MsBy(val))
	return c
}

func (c *Comp) MsPx() *Comp {
	c.El.Style(styles.MsPx())
	return c
}

func (c *Comp) Mt(number int) *Comp {
	c.El.Style(styles.Mt(number))
	return c
}

func (c *Comp) MtAuto() *Comp {
	c.El.Style(styles.MtAuto())
	return c
}

func (c *Comp) MtBy(val value.Value) *Comp {
	c.El.Style(styles.MtBy(val))
	return c
}

func (c *Comp) MtPx() *Comp {
	c.El.Style(styles.MtPx())
	return c
}

func (c *Comp) Mx(number int) *Comp {
	c.El.Style(styles.Mx(number))
	return c
}

func (c *Comp) MxAuto() *Comp {
	c.El.Style(styles.MxAuto())
	return c
}

func (c *Comp) MxBy(val value.Value) *Comp {
	c.El.Style(styles.MxBy(val))
	return c
}

func (c *Comp) MxPx() *Comp {
	c.El.Style(styles.MxPx())
	return c
}

func (c *Comp) My(number int) *Comp {
	c.El.Style(styles.My(number))
	return c
}

func (c *Comp) MyAuto() *Comp {
	c.El.Style(styles.MyAuto())
	return c
}

func (c *Comp) MyBy(val value.Value) *Comp {
	c.El.Style(styles.MyBy(val))
	return c
}

func (c *Comp) MyPx() *Comp {
	c.El.Style(styles.MyPx())
	return c
}

func (c *Comp) NegBottom(number int) *Comp {
	c.El.Style(styles.NegBottom(number))
	return c
}

func (c *Comp) NegBottomFraction(fraction float64) *Comp {
	c.El.Style(styles.NegBottomFraction(fraction))
	return c
}

func (c *Comp) NegBottomFull() *Comp {
	c.El.Style(styles.NegBottomFull())
	return c
}

func (c *Comp) NegBottomPx() *Comp {
	c.El.Style(styles.NegBottomPx())
	return c
}

func (c *Comp) NegCol(number int) *Comp {
	c.El.Style(styles.NegCol(number))
	return c
}

func (c *Comp) NegColEnd(number int) *Comp {
	c.El.Style(styles.NegColEnd(number))
	return c
}

func (c *Comp) NegColStart(number int) *Comp {
	c.El.Style(styles.NegColStart(number))
	return c
}

func (c *Comp) NegEnd(number int) *Comp {
	c.El.Style(styles.NegEnd(number))
	return c
}

func (c *Comp) NegEndFraction(fraction float64) *Comp {
	c.El.Style(styles.NegEndFraction(fraction))
	return c
}

func (c *Comp) NegEndFull() *Comp {
	c.El.Style(styles.NegEndFull())
	return c
}

func (c *Comp) NegEndPx() *Comp {
	c.El.Style(styles.NegEndPx())
	return c
}

func (c *Comp) NegIndent(number int) *Comp {
	c.El.Style(styles.NegIndent(number))
	return c
}

func (c *Comp) NegIndentPx() *Comp {
	c.El.Style(styles.NegIndentPx())
	return c
}

func (c *Comp) NegInset(number int) *Comp {
	c.El.Style(styles.NegInset(number))
	return c
}

func (c *Comp) NegInsetFraction(fraction float64) *Comp {
	c.El.Style(styles.NegInsetFraction(fraction))
	return c
}

func (c *Comp) NegInsetFull() *Comp {
	c.El.Style(styles.NegInsetFull())
	return c
}

func (c *Comp) NegInsetPx() *Comp {
	c.El.Style(styles.NegInsetPx())
	return c
}

func (c *Comp) NegInsetX(number int) *Comp {
	c.El.Style(styles.NegInsetX(number))
	return c
}

func (c *Comp) NegInsetXFraction(fraction float64) *Comp {
	c.El.Style(styles.NegInsetXFraction(fraction))
	return c
}

func (c *Comp) NegInsetXFull() *Comp {
	c.El.Style(styles.NegInsetXFull())
	return c
}

func (c *Comp) NegInsetXPx() *Comp {
	c.El.Style(styles.NegInsetXPx())
	return c
}

func (c *Comp) NegInsetY(number int) *Comp {
	c.El.Style(styles.NegInsetY(number))
	return c
}

func (c *Comp) NegInsetYFraction(fraction float64) *Comp {
	c.El.Style(styles.NegInsetYFraction(fraction))
	return c
}

func (c *Comp) NegInsetYFull() *Comp {
	c.El.Style(styles.NegInsetYFull())
	return c
}

func (c *Comp) NegInsetYPx() *Comp {
	c.El.Style(styles.NegInsetYPx())
	return c
}

func (c *Comp) NegLeft(number int) *Comp {
	c.El.Style(styles.NegLeft(number))
	return c
}

func (c *Comp) NegLeftFraction(fraction float64) *Comp {
	c.El.Style(styles.NegLeftFraction(fraction))
	return c
}

func (c *Comp) NegLeftFull() *Comp {
	c.El.Style(styles.NegLeftFull())
	return c
}

func (c *Comp) NegLeftPx() *Comp {
	c.El.Style(styles.NegLeftPx())
	return c
}

func (c *Comp) NegM(number int) *Comp {
	c.El.Style(styles.NegM(number))
	return c
}

func (c *Comp) NegMPx() *Comp {
	c.El.Style(styles.NegMPx())
	return c
}

func (c *Comp) NegMb(number int) *Comp {
	c.El.Style(styles.NegMb(number))
	return c
}

func (c *Comp) NegMbPx() *Comp {
	c.El.Style(styles.NegMbPx())
	return c
}

func (c *Comp) NegMe(number int) *Comp {
	c.El.Style(styles.NegMe(number))
	return c
}

func (c *Comp) NegMePx() *Comp {
	c.El.Style(styles.NegMePx())
	return c
}

func (c *Comp) NegMl(number int) *Comp {
	c.El.Style(styles.NegMl(number))
	return c
}

func (c *Comp) NegMlPx() *Comp {
	c.El.Style(styles.NegMlPx())
	return c
}

func (c *Comp) NegMr(number int) *Comp {
	c.El.Style(styles.NegMr(number))
	return c
}

func (c *Comp) NegMrPx() *Comp {
	c.El.Style(styles.NegMrPx())
	return c
}

func (c *Comp) NegMs(number int) *Comp {
	c.El.Style(styles.NegMs(number))
	return c
}

func (c *Comp) NegMsPx() *Comp {
	c.El.Style(styles.NegMsPx())
	return c
}

func (c *Comp) NegMt(number int) *Comp {
	c.El.Style(styles.NegMt(number))
	return c
}

func (c *Comp) NegMtPx() *Comp {
	c.El.Style(styles.NegMtPx())
	return c
}

func (c *Comp) NegMx(number int) *Comp {
	c.El.Style(styles.NegMx(number))
	return c
}

func (c *Comp) NegMxPx() *Comp {
	c.El.Style(styles.NegMxPx())
	return c
}

func (c *Comp) NegMy(number int) *Comp {
	c.El.Style(styles.NegMy(number))
	return c
}

func (c *Comp) NegMyPx() *Comp {
	c.El.Style(styles.NegMyPx())
	return c
}

func (c *Comp) NegOrder(number int) *Comp {
	c.El.Style(styles.NegOrder(number))
	return c
}

func (c *Comp) NegRight(number int) *Comp {
	c.El.Style(styles.NegRight(number))
	return c
}

func (c *Comp) NegRightFraction(fraction float64) *Comp {
	c.El.Style(styles.NegRightFraction(fraction))
	return c
}

func (c *Comp) NegRightFull() *Comp {
	c.El.Style(styles.NegRightFull())
	return c
}

func (c *Comp) NegRightPx() *Comp {
	c.El.Style(styles.NegRightPx())
	return c
}

func (c *Comp) NegRow(number int) *Comp {
	c.El.Style(styles.NegRow(number))
	return c
}

func (c *Comp) NegRowEnd(number int) *Comp {
	c.El.Style(styles.NegRowEnd(number))
	return c
}

func (c *Comp) NegRowStart(number int) *Comp {
	c.El.Style(styles.NegRowStart(number))
	return c
}

func (c *Comp) NegStart(number int) *Comp {
	c.El.Style(styles.NegStart(number))
	return c
}

func (c *Comp) NegStartFraction(fraction float64) *Comp {
	c.El.Style(styles.NegStartFraction(fraction))
	return c
}

func (c *Comp) NegStartFull() *Comp {
	c.El.Style(styles.NegStartFull())
	return c
}

func (c *Comp) NegStartPx() *Comp {
	c.El.Style(styles.NegStartPx())
	return c
}

func (c *Comp) NegTop(number int) *Comp {
	c.El.Style(styles.NegTop(number))
	return c
}

func (c *Comp) NegTopFraction(fraction float64) *Comp {
	c.El.Style(styles.NegTopFraction(fraction))
	return c
}

func (c *Comp) NegTopFull() *Comp {
	c.El.Style(styles.NegTopFull())
	return c
}

func (c *Comp) NegTopPx() *Comp {
	c.El.Style(styles.NegTopPx())
	return c
}

func (c *Comp) NegTranslate(val any) *Comp {
	c.El.Style(styles.NegTranslate(val))
	return c
}

func (c *Comp) NegTranslateFull() *Comp {
	c.El.Style(styles.NegTranslateFull())
	return c
}

func (c *Comp) NegTranslatePx() *Comp {
	c.El.Style(styles.NegTranslatePx())
	return c
}

func (c *Comp) NegTranslateX(val any) *Comp {
	c.El.Style(styles.NegTranslateX(val))
	return c
}

func (c *Comp) NegTranslateXFull() *Comp {
	c.El.Style(styles.NegTranslateXFull())
	return c
}

func (c *Comp) NegTranslateXPx() *Comp {
	c.El.Style(styles.NegTranslateXPx())
	return c
}

func (c *Comp) NegTranslateY(val any) *Comp {
	c.El.Style(styles.NegTranslateY(val))
	return c
}

func (c *Comp) NegTranslateYFull() *Comp {
	c.El.Style(styles.NegTranslateYFull())
	return c
}

func (c *Comp) NegTranslateYPx() *Comp {
	c.El.Style(styles.NegTranslateYPx())
	return c
}

func (c *Comp) NegTranslateZ(val any) *Comp {
	c.El.Style(styles.NegTranslateZ(val))
	return c
}

func (c *Comp) NegTranslateZPx() *Comp {
	c.El.Style(styles.NegTranslateZPx())
	return c
}

func (c *Comp) NegUnderlineOffset(number int) *Comp {
	c.El.Style(styles.NegUnderlineOffset(number))
	return c
}

func (c *Comp) NoUnderline() *Comp {
	c.El.Style(styles.NoUnderline())
	return c
}

func (c *Comp) NormalCase() *Comp {
	c.El.Style(styles.NormalCase())
	return c
}

func (c *Comp) NormalNums() *Comp {
	c.El.Style(styles.NormalNums())
	return c
}

func (c *Comp) NotItalic() *Comp {
	c.El.Style(styles.NotItalic())
	return c
}

func (c *Comp) NotSrOnly() *Comp {
	c.El.Style(styles.NotSrOnly())
	return c
}

func (c *Comp) ObjectBottom() *Comp {
	c.El.Style(styles.ObjectBottom())
	return c
}

func (c *Comp) ObjectBottomLeft() *Comp {
	c.El.Style(styles.ObjectBottomLeft())
	return c
}

func (c *Comp) ObjectBottomRight() *Comp {
	c.El.Style(styles.ObjectBottomRight())
	return c
}

func (c *Comp) ObjectBy(val value.Value) *Comp {
	c.El.Style(styles.ObjectBy(val))
	return c
}

func (c *Comp) ObjectCenter() *Comp {
	c.El.Style(styles.ObjectCenter())
	return c
}

func (c *Comp) ObjectContain() *Comp {
	c.El.Style(styles.ObjectContain())
	return c
}

func (c *Comp) ObjectCover() *Comp {
	c.El.Style(styles.ObjectCover())
	return c
}

func (c *Comp) ObjectFill() *Comp {
	c.El.Style(styles.ObjectFill())
	return c
}

func (c *Comp) ObjectLeft() *Comp {
	c.El.Style(styles.ObjectLeft())
	return c
}

func (c *Comp) ObjectNone() *Comp {
	c.El.Style(styles.ObjectNone())
	return c
}

func (c *Comp) ObjectRight() *Comp {
	c.El.Style(styles.ObjectRight())
	return c
}

func (c *Comp) ObjectScaleDown() *Comp {
	c.El.Style(styles.ObjectScaleDown())
	return c
}

func (c *Comp) ObjectTop() *Comp {
	c.El.Style(styles.ObjectTop())
	return c
}

func (c *Comp) ObjectTopLeft() *Comp {
	c.El.Style(styles.ObjectTopLeft())
	return c
}

func (c *Comp) ObjectTopRight() *Comp {
	c.El.Style(styles.ObjectTopRight())
	return c
}

func (c *Comp) OldStyleNums() *Comp {
	c.El.Style(styles.OldStyleNums())
	return c
}

func (c *Comp) Opacity(val value.Value) *Comp {
	c.El.Style(styles.Opacity(val))
	return c
}

func (c *Comp) Order(number int) *Comp {
	c.El.Style(styles.Order(number))
	return c
}

func (c *Comp) OrderBy(val value.Value) *Comp {
	c.El.Style(styles.OrderBy(val))
	return c
}

func (c *Comp) OrderFirst() *Comp {
	c.El.Style(styles.OrderFirst())
	return c
}

func (c *Comp) OrderLast() *Comp {
	c.El.Style(styles.OrderLast())
	return c
}

func (c *Comp) Ordinal() *Comp {
	c.El.Style(styles.Ordinal())
	return c
}

func (c *Comp) Origin(val value.Value) *Comp {
	c.El.Style(styles.Origin(val))
	return c
}

func (c *Comp) OriginBottom() *Comp {
	c.El.Style(styles.OriginBottom())
	return c
}

func (c *Comp) OriginBottomLeft() *Comp {
	c.El.Style(styles.OriginBottomLeft())
	return c
}

func (c *Comp) OriginBottomRight() *Comp {
	c.El.Style(styles.OriginBottomRight())
	return c
}

func (c *Comp) OriginCenter() *Comp {
	c.El.Style(styles.OriginCenter())
	return c
}

func (c *Comp) OriginLeft() *Comp {
	c.El.Style(styles.OriginLeft())
	return c
}

func (c *Comp) OriginRight() *Comp {
	c.El.Style(styles.OriginRight())
	return c
}

func (c *Comp) OriginTop() *Comp {
	c.El.Style(styles.OriginTop())
	return c
}

func (c *Comp) OriginTopLeft() *Comp {
	c.El.Style(styles.OriginTopLeft())
	return c
}

func (c *Comp) OriginTopRight() *Comp {
	c.El.Style(styles.OriginTopRight())
	return c
}

func (c *Comp) Outline(val ...value.Value) *Comp {
	c.El.Style(styles.Outline(val...))
	return c
}

func (c *Comp) OutlineAmber(scale int) *Comp {
	c.El.Style(styles.OutlineAmber(scale))
	return c
}

func (c *Comp) OutlineAmberAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineAmberAlpha(scale))
	return c
}

func (c *Comp) OutlineAmberDark(scale int) *Comp {
	c.El.Style(styles.OutlineAmberDark(scale))
	return c
}

func (c *Comp) OutlineAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineAmberDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineBlack() *Comp {
	c.El.Style(styles.OutlineBlack())
	return c
}

func (c *Comp) OutlineBlackAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineBlackAlpha(scale))
	return c
}

func (c *Comp) OutlineBlue(scale int) *Comp {
	c.El.Style(styles.OutlineBlue(scale))
	return c
}

func (c *Comp) OutlineBlueAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineBlueAlpha(scale))
	return c
}

func (c *Comp) OutlineBlueDark(scale int) *Comp {
	c.El.Style(styles.OutlineBlueDark(scale))
	return c
}

func (c *Comp) OutlineBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineBlueDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineBronze(scale int) *Comp {
	c.El.Style(styles.OutlineBronze(scale))
	return c
}

func (c *Comp) OutlineBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineBronzeAlpha(scale))
	return c
}

func (c *Comp) OutlineBronzeDark(scale int) *Comp {
	c.El.Style(styles.OutlineBronzeDark(scale))
	return c
}

func (c *Comp) OutlineBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineBrown(scale int) *Comp {
	c.El.Style(styles.OutlineBrown(scale))
	return c
}

func (c *Comp) OutlineBrownAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineBrownAlpha(scale))
	return c
}

func (c *Comp) OutlineBrownDark(scale int) *Comp {
	c.El.Style(styles.OutlineBrownDark(scale))
	return c
}

func (c *Comp) OutlineBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineBrownDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineColor(val value.Value) *Comp {
	c.El.Style(styles.OutlineColor(val))
	return c
}

func (c *Comp) OutlineCrimson(scale int) *Comp {
	c.El.Style(styles.OutlineCrimson(scale))
	return c
}

func (c *Comp) OutlineCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineCrimsonAlpha(scale))
	return c
}

func (c *Comp) OutlineCrimsonDark(scale int) *Comp {
	c.El.Style(styles.OutlineCrimsonDark(scale))
	return c
}

func (c *Comp) OutlineCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineCurrent() *Comp {
	c.El.Style(styles.OutlineCurrent())
	return c
}

func (c *Comp) OutlineCyan(scale int) *Comp {
	c.El.Style(styles.OutlineCyan(scale))
	return c
}

func (c *Comp) OutlineCyanAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineCyanAlpha(scale))
	return c
}

func (c *Comp) OutlineCyanDark(scale int) *Comp {
	c.El.Style(styles.OutlineCyanDark(scale))
	return c
}

func (c *Comp) OutlineCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineCyanDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineDashed() *Comp {
	c.El.Style(styles.OutlineDashed())
	return c
}

func (c *Comp) OutlineDotted() *Comp {
	c.El.Style(styles.OutlineDotted())
	return c
}

func (c *Comp) OutlineDouble() *Comp {
	c.El.Style(styles.OutlineDouble())
	return c
}

func (c *Comp) OutlineGold(scale int) *Comp {
	c.El.Style(styles.OutlineGold(scale))
	return c
}

func (c *Comp) OutlineGoldAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineGoldAlpha(scale))
	return c
}

func (c *Comp) OutlineGoldDark(scale int) *Comp {
	c.El.Style(styles.OutlineGoldDark(scale))
	return c
}

func (c *Comp) OutlineGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineGoldDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineGrass(scale int) *Comp {
	c.El.Style(styles.OutlineGrass(scale))
	return c
}

func (c *Comp) OutlineGrassAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineGrassAlpha(scale))
	return c
}

func (c *Comp) OutlineGrassDark(scale int) *Comp {
	c.El.Style(styles.OutlineGrassDark(scale))
	return c
}

func (c *Comp) OutlineGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineGrassDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineGray(scale int) *Comp {
	c.El.Style(styles.OutlineGray(scale))
	return c
}

func (c *Comp) OutlineGrayAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineGrayAlpha(scale))
	return c
}

func (c *Comp) OutlineGrayDark(scale int) *Comp {
	c.El.Style(styles.OutlineGrayDark(scale))
	return c
}

func (c *Comp) OutlineGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineGrayDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineGreen(scale int) *Comp {
	c.El.Style(styles.OutlineGreen(scale))
	return c
}

func (c *Comp) OutlineGreenAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineGreenAlpha(scale))
	return c
}

func (c *Comp) OutlineGreenDark(scale int) *Comp {
	c.El.Style(styles.OutlineGreenDark(scale))
	return c
}

func (c *Comp) OutlineGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineGreenDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineHidden() *Comp {
	c.El.Style(styles.OutlineHidden())
	return c
}

func (c *Comp) OutlineIndigo(scale int) *Comp {
	c.El.Style(styles.OutlineIndigo(scale))
	return c
}

func (c *Comp) OutlineIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineIndigoAlpha(scale))
	return c
}

func (c *Comp) OutlineIndigoDark(scale int) *Comp {
	c.El.Style(styles.OutlineIndigoDark(scale))
	return c
}

func (c *Comp) OutlineIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineInherit() *Comp {
	c.El.Style(styles.OutlineInherit())
	return c
}

func (c *Comp) OutlineIris(scale int) *Comp {
	c.El.Style(styles.OutlineIris(scale))
	return c
}

func (c *Comp) OutlineIrisAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineIrisAlpha(scale))
	return c
}

func (c *Comp) OutlineIrisDark(scale int) *Comp {
	c.El.Style(styles.OutlineIrisDark(scale))
	return c
}

func (c *Comp) OutlineIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineIrisDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineJade(scale int) *Comp {
	c.El.Style(styles.OutlineJade(scale))
	return c
}

func (c *Comp) OutlineJadeAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineJadeAlpha(scale))
	return c
}

func (c *Comp) OutlineJadeDark(scale int) *Comp {
	c.El.Style(styles.OutlineJadeDark(scale))
	return c
}

func (c *Comp) OutlineJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineJadeDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineLime(scale int) *Comp {
	c.El.Style(styles.OutlineLime(scale))
	return c
}

func (c *Comp) OutlineLimeAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineLimeAlpha(scale))
	return c
}

func (c *Comp) OutlineLimeDark(scale int) *Comp {
	c.El.Style(styles.OutlineLimeDark(scale))
	return c
}

func (c *Comp) OutlineLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineLimeDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineMauve(scale int) *Comp {
	c.El.Style(styles.OutlineMauve(scale))
	return c
}

func (c *Comp) OutlineMauveAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineMauveAlpha(scale))
	return c
}

func (c *Comp) OutlineMauveDark(scale int) *Comp {
	c.El.Style(styles.OutlineMauveDark(scale))
	return c
}

func (c *Comp) OutlineMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineMauveDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineMint(scale int) *Comp {
	c.El.Style(styles.OutlineMint(scale))
	return c
}

func (c *Comp) OutlineMintAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineMintAlpha(scale))
	return c
}

func (c *Comp) OutlineMintDark(scale int) *Comp {
	c.El.Style(styles.OutlineMintDark(scale))
	return c
}

func (c *Comp) OutlineMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineMintDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineNone() *Comp {
	c.El.Style(styles.OutlineNone())
	return c
}

func (c *Comp) OutlineOffset(val ...value.Value) *Comp {
	c.El.Style(styles.OutlineOffset(val...))
	return c
}

func (c *Comp) OutlineOlive(scale int) *Comp {
	c.El.Style(styles.OutlineOlive(scale))
	return c
}

func (c *Comp) OutlineOliveAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineOliveAlpha(scale))
	return c
}

func (c *Comp) OutlineOliveDark(scale int) *Comp {
	c.El.Style(styles.OutlineOliveDark(scale))
	return c
}

func (c *Comp) OutlineOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineOliveDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineOrange(scale int) *Comp {
	c.El.Style(styles.OutlineOrange(scale))
	return c
}

func (c *Comp) OutlineOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineOrangeAlpha(scale))
	return c
}

func (c *Comp) OutlineOrangeDark(scale int) *Comp {
	c.El.Style(styles.OutlineOrangeDark(scale))
	return c
}

func (c *Comp) OutlineOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) OutlinePink(scale int) *Comp {
	c.El.Style(styles.OutlinePink(scale))
	return c
}

func (c *Comp) OutlinePinkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlinePinkAlpha(scale))
	return c
}

func (c *Comp) OutlinePinkDark(scale int) *Comp {
	c.El.Style(styles.OutlinePinkDark(scale))
	return c
}

func (c *Comp) OutlinePinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlinePinkDarkAlpha(scale))
	return c
}

func (c *Comp) OutlinePlum(scale int) *Comp {
	c.El.Style(styles.OutlinePlum(scale))
	return c
}

func (c *Comp) OutlinePlumAlpha(scale int) *Comp {
	c.El.Style(styles.OutlinePlumAlpha(scale))
	return c
}

func (c *Comp) OutlinePlumDark(scale int) *Comp {
	c.El.Style(styles.OutlinePlumDark(scale))
	return c
}

func (c *Comp) OutlinePlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlinePlumDarkAlpha(scale))
	return c
}

func (c *Comp) OutlinePurple(scale int) *Comp {
	c.El.Style(styles.OutlinePurple(scale))
	return c
}

func (c *Comp) OutlinePurpleAlpha(scale int) *Comp {
	c.El.Style(styles.OutlinePurpleAlpha(scale))
	return c
}

func (c *Comp) OutlinePurpleDark(scale int) *Comp {
	c.El.Style(styles.OutlinePurpleDark(scale))
	return c
}

func (c *Comp) OutlinePurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlinePurpleDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineRed(scale int) *Comp {
	c.El.Style(styles.OutlineRed(scale))
	return c
}

func (c *Comp) OutlineRedAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineRedAlpha(scale))
	return c
}

func (c *Comp) OutlineRedDark(scale int) *Comp {
	c.El.Style(styles.OutlineRedDark(scale))
	return c
}

func (c *Comp) OutlineRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineRedDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineRuby(scale int) *Comp {
	c.El.Style(styles.OutlineRuby(scale))
	return c
}

func (c *Comp) OutlineRubyAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineRubyAlpha(scale))
	return c
}

func (c *Comp) OutlineRubyDark(scale int) *Comp {
	c.El.Style(styles.OutlineRubyDark(scale))
	return c
}

func (c *Comp) OutlineRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineRubyDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineSage(scale int) *Comp {
	c.El.Style(styles.OutlineSage(scale))
	return c
}

func (c *Comp) OutlineSageAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineSageAlpha(scale))
	return c
}

func (c *Comp) OutlineSageDark(scale int) *Comp {
	c.El.Style(styles.OutlineSageDark(scale))
	return c
}

func (c *Comp) OutlineSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineSageDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineSand(scale int) *Comp {
	c.El.Style(styles.OutlineSand(scale))
	return c
}

func (c *Comp) OutlineSandAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineSandAlpha(scale))
	return c
}

func (c *Comp) OutlineSandDark(scale int) *Comp {
	c.El.Style(styles.OutlineSandDark(scale))
	return c
}

func (c *Comp) OutlineSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineSandDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineSky(scale int) *Comp {
	c.El.Style(styles.OutlineSky(scale))
	return c
}

func (c *Comp) OutlineSkyAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineSkyAlpha(scale))
	return c
}

func (c *Comp) OutlineSkyDark(scale int) *Comp {
	c.El.Style(styles.OutlineSkyDark(scale))
	return c
}

func (c *Comp) OutlineSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineSkyDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineSlate(scale int) *Comp {
	c.El.Style(styles.OutlineSlate(scale))
	return c
}

func (c *Comp) OutlineSlateAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineSlateAlpha(scale))
	return c
}

func (c *Comp) OutlineSlateDark(scale int) *Comp {
	c.El.Style(styles.OutlineSlateDark(scale))
	return c
}

func (c *Comp) OutlineSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineSlateDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineSolid() *Comp {
	c.El.Style(styles.OutlineSolid())
	return c
}

func (c *Comp) OutlineTeal(scale int) *Comp {
	c.El.Style(styles.OutlineTeal(scale))
	return c
}

func (c *Comp) OutlineTealAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineTealAlpha(scale))
	return c
}

func (c *Comp) OutlineTealDark(scale int) *Comp {
	c.El.Style(styles.OutlineTealDark(scale))
	return c
}

func (c *Comp) OutlineTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineTealDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineTomato(scale int) *Comp {
	c.El.Style(styles.OutlineTomato(scale))
	return c
}

func (c *Comp) OutlineTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineTomatoAlpha(scale))
	return c
}

func (c *Comp) OutlineTomatoDark(scale int) *Comp {
	c.El.Style(styles.OutlineTomatoDark(scale))
	return c
}

func (c *Comp) OutlineTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineTransparent() *Comp {
	c.El.Style(styles.OutlineTransparent())
	return c
}

func (c *Comp) OutlineViolet(scale int) *Comp {
	c.El.Style(styles.OutlineViolet(scale))
	return c
}

func (c *Comp) OutlineVioletAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineVioletAlpha(scale))
	return c
}

func (c *Comp) OutlineVioletDark(scale int) *Comp {
	c.El.Style(styles.OutlineVioletDark(scale))
	return c
}

func (c *Comp) OutlineVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineVioletDarkAlpha(scale))
	return c
}

func (c *Comp) OutlineWhite() *Comp {
	c.El.Style(styles.OutlineWhite())
	return c
}

func (c *Comp) OutlineWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineWhiteAlpha(scale))
	return c
}

func (c *Comp) OutlineYellow(scale int) *Comp {
	c.El.Style(styles.OutlineYellow(scale))
	return c
}

func (c *Comp) OutlineYellowAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineYellowAlpha(scale))
	return c
}

func (c *Comp) OutlineYellowDark(scale int) *Comp {
	c.El.Style(styles.OutlineYellowDark(scale))
	return c
}

func (c *Comp) OutlineYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.OutlineYellowDarkAlpha(scale))
	return c
}

func (c *Comp) OverflowAuto() *Comp {
	c.El.Style(styles.OverflowAuto())
	return c
}

func (c *Comp) OverflowClip() *Comp {
	c.El.Style(styles.OverflowClip())
	return c
}

func (c *Comp) OverflowHidden() *Comp {
	c.El.Style(styles.OverflowHidden())
	return c
}

func (c *Comp) OverflowScroll() *Comp {
	c.El.Style(styles.OverflowScroll())
	return c
}

func (c *Comp) OverflowVisible() *Comp {
	c.El.Style(styles.OverflowVisible())
	return c
}

func (c *Comp) OverflowXAuto() *Comp {
	c.El.Style(styles.OverflowXAuto())
	return c
}

func (c *Comp) OverflowXClip() *Comp {
	c.El.Style(styles.OverflowXClip())
	return c
}

func (c *Comp) OverflowXHidden() *Comp {
	c.El.Style(styles.OverflowXHidden())
	return c
}

func (c *Comp) OverflowXScroll() *Comp {
	c.El.Style(styles.OverflowXScroll())
	return c
}

func (c *Comp) OverflowXVisible() *Comp {
	c.El.Style(styles.OverflowXVisible())
	return c
}

func (c *Comp) OverflowYAuto() *Comp {
	c.El.Style(styles.OverflowYAuto())
	return c
}

func (c *Comp) OverflowYClip() *Comp {
	c.El.Style(styles.OverflowYClip())
	return c
}

func (c *Comp) OverflowYHidden() *Comp {
	c.El.Style(styles.OverflowYHidden())
	return c
}

func (c *Comp) OverflowYScroll() *Comp {
	c.El.Style(styles.OverflowYScroll())
	return c
}

func (c *Comp) OverflowYVisible() *Comp {
	c.El.Style(styles.OverflowYVisible())
	return c
}

func (c *Comp) Overline() *Comp {
	c.El.Style(styles.Overline())
	return c
}

func (c *Comp) OverscrollAuto() *Comp {
	c.El.Style(styles.OverscrollAuto())
	return c
}

func (c *Comp) OverscrollContain() *Comp {
	c.El.Style(styles.OverscrollContain())
	return c
}

func (c *Comp) OverscrollNone() *Comp {
	c.El.Style(styles.OverscrollNone())
	return c
}

func (c *Comp) OverscrollXAuto() *Comp {
	c.El.Style(styles.OverscrollXAuto())
	return c
}

func (c *Comp) OverscrollXContain() *Comp {
	c.El.Style(styles.OverscrollXContain())
	return c
}

func (c *Comp) OverscrollXNone() *Comp {
	c.El.Style(styles.OverscrollXNone())
	return c
}

func (c *Comp) OverscrollYAuto() *Comp {
	c.El.Style(styles.OverscrollYAuto())
	return c
}

func (c *Comp) OverscrollYContain() *Comp {
	c.El.Style(styles.OverscrollYContain())
	return c
}

func (c *Comp) OverscrollYNone() *Comp {
	c.El.Style(styles.OverscrollYNone())
	return c
}

func (c *Comp) P(number int) *Comp {
	c.El.Style(styles.P(number))
	return c
}

func (c *Comp) PBy(val value.Value) *Comp {
	c.El.Style(styles.PBy(val))
	return c
}

func (c *Comp) PPx(number int) *Comp {
	c.El.Style(styles.PPx(number))
	return c
}

func (c *Comp) Pb(number int) *Comp {
	c.El.Style(styles.Pb(number))
	return c
}

func (c *Comp) PbBy(val value.Value) *Comp {
	c.El.Style(styles.PbBy(val))
	return c
}

func (c *Comp) PbPx(number int) *Comp {
	c.El.Style(styles.PbPx(number))
	return c
}

func (c *Comp) Pe(number int) *Comp {
	c.El.Style(styles.Pe(number))
	return c
}

func (c *Comp) PeBy(val value.Value) *Comp {
	c.El.Style(styles.PeBy(val))
	return c
}

func (c *Comp) PePx(number int) *Comp {
	c.El.Style(styles.PePx(number))
	return c
}

func (c *Comp) Perspective(val value.Value) *Comp {
	c.El.Style(styles.Perspective(val))
	return c
}

func (c *Comp) PerspectiveDistant() *Comp {
	c.El.Style(styles.PerspectiveDistant())
	return c
}

func (c *Comp) PerspectiveDramatic() *Comp {
	c.El.Style(styles.PerspectiveDramatic())
	return c
}

func (c *Comp) PerspectiveMidrange() *Comp {
	c.El.Style(styles.PerspectiveMidrange())
	return c
}

func (c *Comp) PerspectiveNear() *Comp {
	c.El.Style(styles.PerspectiveNear())
	return c
}

func (c *Comp) PerspectiveNone() *Comp {
	c.El.Style(styles.PerspectiveNone())
	return c
}

func (c *Comp) PerspectiveNormal() *Comp {
	c.El.Style(styles.PerspectiveNormal())
	return c
}

func (c *Comp) PerspectiveOrigin(val value.Value) *Comp {
	c.El.Style(styles.PerspectiveOrigin(val))
	return c
}

func (c *Comp) PerspectiveOriginBottom() *Comp {
	c.El.Style(styles.PerspectiveOriginBottom())
	return c
}

func (c *Comp) PerspectiveOriginBottomLeft() *Comp {
	c.El.Style(styles.PerspectiveOriginBottomLeft())
	return c
}

func (c *Comp) PerspectiveOriginBottomRight() *Comp {
	c.El.Style(styles.PerspectiveOriginBottomRight())
	return c
}

func (c *Comp) PerspectiveOriginCenter() *Comp {
	c.El.Style(styles.PerspectiveOriginCenter())
	return c
}

func (c *Comp) PerspectiveOriginLeft() *Comp {
	c.El.Style(styles.PerspectiveOriginLeft())
	return c
}

func (c *Comp) PerspectiveOriginRight() *Comp {
	c.El.Style(styles.PerspectiveOriginRight())
	return c
}

func (c *Comp) PerspectiveOriginTop() *Comp {
	c.El.Style(styles.PerspectiveOriginTop())
	return c
}

func (c *Comp) PerspectiveOriginTopLeft() *Comp {
	c.El.Style(styles.PerspectiveOriginTopLeft())
	return c
}

func (c *Comp) PerspectiveOriginTopRight() *Comp {
	c.El.Style(styles.PerspectiveOriginTopRight())
	return c
}

func (c *Comp) Pl(number int) *Comp {
	c.El.Style(styles.Pl(number))
	return c
}

func (c *Comp) PlBy(val value.Value) *Comp {
	c.El.Style(styles.PlBy(val))
	return c
}

func (c *Comp) PlPx(number int) *Comp {
	c.El.Style(styles.PlPx(number))
	return c
}

func (c *Comp) PlaceContentAround() *Comp {
	c.El.Style(styles.PlaceContentAround())
	return c
}

func (c *Comp) PlaceContentBaseline() *Comp {
	c.El.Style(styles.PlaceContentBaseline())
	return c
}

func (c *Comp) PlaceContentBetween() *Comp {
	c.El.Style(styles.PlaceContentBetween())
	return c
}

func (c *Comp) PlaceContentCenter() *Comp {
	c.El.Style(styles.PlaceContentCenter())
	return c
}

func (c *Comp) PlaceContentCenterSafe() *Comp {
	c.El.Style(styles.PlaceContentCenterSafe())
	return c
}

func (c *Comp) PlaceContentEnd() *Comp {
	c.El.Style(styles.PlaceContentEnd())
	return c
}

func (c *Comp) PlaceContentEndSafe() *Comp {
	c.El.Style(styles.PlaceContentEndSafe())
	return c
}

func (c *Comp) PlaceContentEvenly() *Comp {
	c.El.Style(styles.PlaceContentEvenly())
	return c
}

func (c *Comp) PlaceContentStart() *Comp {
	c.El.Style(styles.PlaceContentStart())
	return c
}

func (c *Comp) PlaceContentStretch() *Comp {
	c.El.Style(styles.PlaceContentStretch())
	return c
}

func (c *Comp) PlaceItemsBaseline() *Comp {
	c.El.Style(styles.PlaceItemsBaseline())
	return c
}

func (c *Comp) PlaceItemsCenter() *Comp {
	c.El.Style(styles.PlaceItemsCenter())
	return c
}

func (c *Comp) PlaceItemsCenterSafe() *Comp {
	c.El.Style(styles.PlaceItemsCenterSafe())
	return c
}

func (c *Comp) PlaceItemsEnd() *Comp {
	c.El.Style(styles.PlaceItemsEnd())
	return c
}

func (c *Comp) PlaceItemsEndSafe() *Comp {
	c.El.Style(styles.PlaceItemsEndSafe())
	return c
}

func (c *Comp) PlaceItemsStart() *Comp {
	c.El.Style(styles.PlaceItemsStart())
	return c
}

func (c *Comp) PlaceItemsStretch() *Comp {
	c.El.Style(styles.PlaceItemsStretch())
	return c
}

func (c *Comp) PlaceSelfAuto() *Comp {
	c.El.Style(styles.PlaceSelfAuto())
	return c
}

func (c *Comp) PlaceSelfCenter() *Comp {
	c.El.Style(styles.PlaceSelfCenter())
	return c
}

func (c *Comp) PlaceSelfCenterSafe() *Comp {
	c.El.Style(styles.PlaceSelfCenterSafe())
	return c
}

func (c *Comp) PlaceSelfEnd() *Comp {
	c.El.Style(styles.PlaceSelfEnd())
	return c
}

func (c *Comp) PlaceSelfEndSafe() *Comp {
	c.El.Style(styles.PlaceSelfEndSafe())
	return c
}

func (c *Comp) PlaceSelfStart() *Comp {
	c.El.Style(styles.PlaceSelfStart())
	return c
}

func (c *Comp) PlaceSelfStretch() *Comp {
	c.El.Style(styles.PlaceSelfStretch())
	return c
}

func (c *Comp) PointerEventsAuto() *Comp {
	c.El.Style(styles.PointerEventsAuto())
	return c
}

func (c *Comp) PointerEventsNone() *Comp {
	c.El.Style(styles.PointerEventsNone())
	return c
}

func (c *Comp) Pr(number int) *Comp {
	c.El.Style(styles.Pr(number))
	return c
}

func (c *Comp) PrBy(val value.Value) *Comp {
	c.El.Style(styles.PrBy(val))
	return c
}

func (c *Comp) PrPx(number int) *Comp {
	c.El.Style(styles.PrPx(number))
	return c
}

func (c *Comp) ProportionalNums() *Comp {
	c.El.Style(styles.ProportionalNums())
	return c
}

func (c *Comp) Ps(number int) *Comp {
	c.El.Style(styles.Ps(number))
	return c
}

func (c *Comp) PsBy(val value.Value) *Comp {
	c.El.Style(styles.PsBy(val))
	return c
}

func (c *Comp) PsPx(number int) *Comp {
	c.El.Style(styles.PsPx(number))
	return c
}

func (c *Comp) Pt(number int) *Comp {
	c.El.Style(styles.Pt(number))
	return c
}

func (c *Comp) PtBy(val value.Value) *Comp {
	c.El.Style(styles.PtBy(val))
	return c
}

func (c *Comp) PtPx(number int) *Comp {
	c.El.Style(styles.PtPx(number))
	return c
}

func (c *Comp) Px(number int) *Comp {
	c.El.Style(styles.Px(number))
	return c
}

func (c *Comp) PxBy(val value.Value) *Comp {
	c.El.Style(styles.PxBy(val))
	return c
}

func (c *Comp) PxPx(number int) *Comp {
	c.El.Style(styles.PxPx(number))
	return c
}

func (c *Comp) Py(number int) *Comp {
	c.El.Style(styles.Py(number))
	return c
}

func (c *Comp) PyBy(val value.Value) *Comp {
	c.El.Style(styles.PyBy(val))
	return c
}

func (c *Comp) PyPx(number int) *Comp {
	c.El.Style(styles.PyPx(number))
	return c
}

func (c *Comp) Relative() *Comp {
	c.El.Style(styles.Relative())
	return c
}

func (c *Comp) Resize() *Comp {
	c.El.Style(styles.Resize())
	return c
}

func (c *Comp) ResizeNone() *Comp {
	c.El.Style(styles.ResizeNone())
	return c
}

func (c *Comp) ResizeX() *Comp {
	c.El.Style(styles.ResizeX())
	return c
}

func (c *Comp) ResizeY() *Comp {
	c.El.Style(styles.ResizeY())
	return c
}

func (c *Comp) Right(number int) *Comp {
	c.El.Style(styles.Right(number))
	return c
}

func (c *Comp) RightAuto() *Comp {
	c.El.Style(styles.RightAuto())
	return c
}

func (c *Comp) RightBy(val value.Value) *Comp {
	c.El.Style(styles.RightBy(val))
	return c
}

func (c *Comp) RightFraction(fraction float64) *Comp {
	c.El.Style(styles.RightFraction(fraction))
	return c
}

func (c *Comp) RightFull() *Comp {
	c.El.Style(styles.RightFull())
	return c
}

func (c *Comp) RightPx() *Comp {
	c.El.Style(styles.RightPx())
	return c
}

func (c *Comp) Ring(val ...value.Value) *Comp {
	c.El.Style(styles.Ring(val...))
	return c
}

func (c *Comp) RingAmber(scale int) *Comp {
	c.El.Style(styles.RingAmber(scale))
	return c
}

func (c *Comp) RingAmberAlpha(scale int) *Comp {
	c.El.Style(styles.RingAmberAlpha(scale))
	return c
}

func (c *Comp) RingAmberDark(scale int) *Comp {
	c.El.Style(styles.RingAmberDark(scale))
	return c
}

func (c *Comp) RingAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingAmberDarkAlpha(scale))
	return c
}

func (c *Comp) RingBlack() *Comp {
	c.El.Style(styles.RingBlack())
	return c
}

func (c *Comp) RingBlackAlpha(scale int) *Comp {
	c.El.Style(styles.RingBlackAlpha(scale))
	return c
}

func (c *Comp) RingBlue(scale int) *Comp {
	c.El.Style(styles.RingBlue(scale))
	return c
}

func (c *Comp) RingBlueAlpha(scale int) *Comp {
	c.El.Style(styles.RingBlueAlpha(scale))
	return c
}

func (c *Comp) RingBlueDark(scale int) *Comp {
	c.El.Style(styles.RingBlueDark(scale))
	return c
}

func (c *Comp) RingBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingBlueDarkAlpha(scale))
	return c
}

func (c *Comp) RingBronze(scale int) *Comp {
	c.El.Style(styles.RingBronze(scale))
	return c
}

func (c *Comp) RingBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.RingBronzeAlpha(scale))
	return c
}

func (c *Comp) RingBronzeDark(scale int) *Comp {
	c.El.Style(styles.RingBronzeDark(scale))
	return c
}

func (c *Comp) RingBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) RingBrown(scale int) *Comp {
	c.El.Style(styles.RingBrown(scale))
	return c
}

func (c *Comp) RingBrownAlpha(scale int) *Comp {
	c.El.Style(styles.RingBrownAlpha(scale))
	return c
}

func (c *Comp) RingBrownDark(scale int) *Comp {
	c.El.Style(styles.RingBrownDark(scale))
	return c
}

func (c *Comp) RingBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingBrownDarkAlpha(scale))
	return c
}

func (c *Comp) RingColor(val value.Value) *Comp {
	c.El.Style(styles.RingColor(val))
	return c
}

func (c *Comp) RingCrimson(scale int) *Comp {
	c.El.Style(styles.RingCrimson(scale))
	return c
}

func (c *Comp) RingCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.RingCrimsonAlpha(scale))
	return c
}

func (c *Comp) RingCrimsonDark(scale int) *Comp {
	c.El.Style(styles.RingCrimsonDark(scale))
	return c
}

func (c *Comp) RingCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) RingCurrent() *Comp {
	c.El.Style(styles.RingCurrent())
	return c
}

func (c *Comp) RingCyan(scale int) *Comp {
	c.El.Style(styles.RingCyan(scale))
	return c
}

func (c *Comp) RingCyanAlpha(scale int) *Comp {
	c.El.Style(styles.RingCyanAlpha(scale))
	return c
}

func (c *Comp) RingCyanDark(scale int) *Comp {
	c.El.Style(styles.RingCyanDark(scale))
	return c
}

func (c *Comp) RingCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingCyanDarkAlpha(scale))
	return c
}

func (c *Comp) RingGold(scale int) *Comp {
	c.El.Style(styles.RingGold(scale))
	return c
}

func (c *Comp) RingGoldAlpha(scale int) *Comp {
	c.El.Style(styles.RingGoldAlpha(scale))
	return c
}

func (c *Comp) RingGoldDark(scale int) *Comp {
	c.El.Style(styles.RingGoldDark(scale))
	return c
}

func (c *Comp) RingGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingGoldDarkAlpha(scale))
	return c
}

func (c *Comp) RingGrass(scale int) *Comp {
	c.El.Style(styles.RingGrass(scale))
	return c
}

func (c *Comp) RingGrassAlpha(scale int) *Comp {
	c.El.Style(styles.RingGrassAlpha(scale))
	return c
}

func (c *Comp) RingGrassDark(scale int) *Comp {
	c.El.Style(styles.RingGrassDark(scale))
	return c
}

func (c *Comp) RingGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingGrassDarkAlpha(scale))
	return c
}

func (c *Comp) RingGray(scale int) *Comp {
	c.El.Style(styles.RingGray(scale))
	return c
}

func (c *Comp) RingGrayAlpha(scale int) *Comp {
	c.El.Style(styles.RingGrayAlpha(scale))
	return c
}

func (c *Comp) RingGrayDark(scale int) *Comp {
	c.El.Style(styles.RingGrayDark(scale))
	return c
}

func (c *Comp) RingGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingGrayDarkAlpha(scale))
	return c
}

func (c *Comp) RingGreen(scale int) *Comp {
	c.El.Style(styles.RingGreen(scale))
	return c
}

func (c *Comp) RingGreenAlpha(scale int) *Comp {
	c.El.Style(styles.RingGreenAlpha(scale))
	return c
}

func (c *Comp) RingGreenDark(scale int) *Comp {
	c.El.Style(styles.RingGreenDark(scale))
	return c
}

func (c *Comp) RingGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingGreenDarkAlpha(scale))
	return c
}

func (c *Comp) RingIndigo(scale int) *Comp {
	c.El.Style(styles.RingIndigo(scale))
	return c
}

func (c *Comp) RingIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.RingIndigoAlpha(scale))
	return c
}

func (c *Comp) RingIndigoDark(scale int) *Comp {
	c.El.Style(styles.RingIndigoDark(scale))
	return c
}

func (c *Comp) RingIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) RingInherit() *Comp {
	c.El.Style(styles.RingInherit())
	return c
}

func (c *Comp) RingIris(scale int) *Comp {
	c.El.Style(styles.RingIris(scale))
	return c
}

func (c *Comp) RingIrisAlpha(scale int) *Comp {
	c.El.Style(styles.RingIrisAlpha(scale))
	return c
}

func (c *Comp) RingIrisDark(scale int) *Comp {
	c.El.Style(styles.RingIrisDark(scale))
	return c
}

func (c *Comp) RingIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingIrisDarkAlpha(scale))
	return c
}

func (c *Comp) RingJade(scale int) *Comp {
	c.El.Style(styles.RingJade(scale))
	return c
}

func (c *Comp) RingJadeAlpha(scale int) *Comp {
	c.El.Style(styles.RingJadeAlpha(scale))
	return c
}

func (c *Comp) RingJadeDark(scale int) *Comp {
	c.El.Style(styles.RingJadeDark(scale))
	return c
}

func (c *Comp) RingJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingJadeDarkAlpha(scale))
	return c
}

func (c *Comp) RingLime(scale int) *Comp {
	c.El.Style(styles.RingLime(scale))
	return c
}

func (c *Comp) RingLimeAlpha(scale int) *Comp {
	c.El.Style(styles.RingLimeAlpha(scale))
	return c
}

func (c *Comp) RingLimeDark(scale int) *Comp {
	c.El.Style(styles.RingLimeDark(scale))
	return c
}

func (c *Comp) RingLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingLimeDarkAlpha(scale))
	return c
}

func (c *Comp) RingMauve(scale int) *Comp {
	c.El.Style(styles.RingMauve(scale))
	return c
}

func (c *Comp) RingMauveAlpha(scale int) *Comp {
	c.El.Style(styles.RingMauveAlpha(scale))
	return c
}

func (c *Comp) RingMauveDark(scale int) *Comp {
	c.El.Style(styles.RingMauveDark(scale))
	return c
}

func (c *Comp) RingMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingMauveDarkAlpha(scale))
	return c
}

func (c *Comp) RingMint(scale int) *Comp {
	c.El.Style(styles.RingMint(scale))
	return c
}

func (c *Comp) RingMintAlpha(scale int) *Comp {
	c.El.Style(styles.RingMintAlpha(scale))
	return c
}

func (c *Comp) RingMintDark(scale int) *Comp {
	c.El.Style(styles.RingMintDark(scale))
	return c
}

func (c *Comp) RingMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingMintDarkAlpha(scale))
	return c
}

func (c *Comp) RingOlive(scale int) *Comp {
	c.El.Style(styles.RingOlive(scale))
	return c
}

func (c *Comp) RingOliveAlpha(scale int) *Comp {
	c.El.Style(styles.RingOliveAlpha(scale))
	return c
}

func (c *Comp) RingOliveDark(scale int) *Comp {
	c.El.Style(styles.RingOliveDark(scale))
	return c
}

func (c *Comp) RingOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingOliveDarkAlpha(scale))
	return c
}

func (c *Comp) RingOrange(scale int) *Comp {
	c.El.Style(styles.RingOrange(scale))
	return c
}

func (c *Comp) RingOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.RingOrangeAlpha(scale))
	return c
}

func (c *Comp) RingOrangeDark(scale int) *Comp {
	c.El.Style(styles.RingOrangeDark(scale))
	return c
}

func (c *Comp) RingOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) RingPink(scale int) *Comp {
	c.El.Style(styles.RingPink(scale))
	return c
}

func (c *Comp) RingPinkAlpha(scale int) *Comp {
	c.El.Style(styles.RingPinkAlpha(scale))
	return c
}

func (c *Comp) RingPinkDark(scale int) *Comp {
	c.El.Style(styles.RingPinkDark(scale))
	return c
}

func (c *Comp) RingPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingPinkDarkAlpha(scale))
	return c
}

func (c *Comp) RingPlum(scale int) *Comp {
	c.El.Style(styles.RingPlum(scale))
	return c
}

func (c *Comp) RingPlumAlpha(scale int) *Comp {
	c.El.Style(styles.RingPlumAlpha(scale))
	return c
}

func (c *Comp) RingPlumDark(scale int) *Comp {
	c.El.Style(styles.RingPlumDark(scale))
	return c
}

func (c *Comp) RingPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingPlumDarkAlpha(scale))
	return c
}

func (c *Comp) RingPurple(scale int) *Comp {
	c.El.Style(styles.RingPurple(scale))
	return c
}

func (c *Comp) RingPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.RingPurpleAlpha(scale))
	return c
}

func (c *Comp) RingPurpleDark(scale int) *Comp {
	c.El.Style(styles.RingPurpleDark(scale))
	return c
}

func (c *Comp) RingPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) RingRed(scale int) *Comp {
	c.El.Style(styles.RingRed(scale))
	return c
}

func (c *Comp) RingRedAlpha(scale int) *Comp {
	c.El.Style(styles.RingRedAlpha(scale))
	return c
}

func (c *Comp) RingRedDark(scale int) *Comp {
	c.El.Style(styles.RingRedDark(scale))
	return c
}

func (c *Comp) RingRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingRedDarkAlpha(scale))
	return c
}

func (c *Comp) RingRuby(scale int) *Comp {
	c.El.Style(styles.RingRuby(scale))
	return c
}

func (c *Comp) RingRubyAlpha(scale int) *Comp {
	c.El.Style(styles.RingRubyAlpha(scale))
	return c
}

func (c *Comp) RingRubyDark(scale int) *Comp {
	c.El.Style(styles.RingRubyDark(scale))
	return c
}

func (c *Comp) RingRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingRubyDarkAlpha(scale))
	return c
}

func (c *Comp) RingSage(scale int) *Comp {
	c.El.Style(styles.RingSage(scale))
	return c
}

func (c *Comp) RingSageAlpha(scale int) *Comp {
	c.El.Style(styles.RingSageAlpha(scale))
	return c
}

func (c *Comp) RingSageDark(scale int) *Comp {
	c.El.Style(styles.RingSageDark(scale))
	return c
}

func (c *Comp) RingSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingSageDarkAlpha(scale))
	return c
}

func (c *Comp) RingSand(scale int) *Comp {
	c.El.Style(styles.RingSand(scale))
	return c
}

func (c *Comp) RingSandAlpha(scale int) *Comp {
	c.El.Style(styles.RingSandAlpha(scale))
	return c
}

func (c *Comp) RingSandDark(scale int) *Comp {
	c.El.Style(styles.RingSandDark(scale))
	return c
}

func (c *Comp) RingSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingSandDarkAlpha(scale))
	return c
}

func (c *Comp) RingSky(scale int) *Comp {
	c.El.Style(styles.RingSky(scale))
	return c
}

func (c *Comp) RingSkyAlpha(scale int) *Comp {
	c.El.Style(styles.RingSkyAlpha(scale))
	return c
}

func (c *Comp) RingSkyDark(scale int) *Comp {
	c.El.Style(styles.RingSkyDark(scale))
	return c
}

func (c *Comp) RingSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingSkyDarkAlpha(scale))
	return c
}

func (c *Comp) RingSlate(scale int) *Comp {
	c.El.Style(styles.RingSlate(scale))
	return c
}

func (c *Comp) RingSlateAlpha(scale int) *Comp {
	c.El.Style(styles.RingSlateAlpha(scale))
	return c
}

func (c *Comp) RingSlateDark(scale int) *Comp {
	c.El.Style(styles.RingSlateDark(scale))
	return c
}

func (c *Comp) RingSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingSlateDarkAlpha(scale))
	return c
}

func (c *Comp) RingTeal(scale int) *Comp {
	c.El.Style(styles.RingTeal(scale))
	return c
}

func (c *Comp) RingTealAlpha(scale int) *Comp {
	c.El.Style(styles.RingTealAlpha(scale))
	return c
}

func (c *Comp) RingTealDark(scale int) *Comp {
	c.El.Style(styles.RingTealDark(scale))
	return c
}

func (c *Comp) RingTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingTealDarkAlpha(scale))
	return c
}

func (c *Comp) RingTomato(scale int) *Comp {
	c.El.Style(styles.RingTomato(scale))
	return c
}

func (c *Comp) RingTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.RingTomatoAlpha(scale))
	return c
}

func (c *Comp) RingTomatoDark(scale int) *Comp {
	c.El.Style(styles.RingTomatoDark(scale))
	return c
}

func (c *Comp) RingTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) RingTransparent() *Comp {
	c.El.Style(styles.RingTransparent())
	return c
}

func (c *Comp) RingViolet(scale int) *Comp {
	c.El.Style(styles.RingViolet(scale))
	return c
}

func (c *Comp) RingVioletAlpha(scale int) *Comp {
	c.El.Style(styles.RingVioletAlpha(scale))
	return c
}

func (c *Comp) RingVioletDark(scale int) *Comp {
	c.El.Style(styles.RingVioletDark(scale))
	return c
}

func (c *Comp) RingVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingVioletDarkAlpha(scale))
	return c
}

func (c *Comp) RingWhite() *Comp {
	c.El.Style(styles.RingWhite())
	return c
}

func (c *Comp) RingWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.RingWhiteAlpha(scale))
	return c
}

func (c *Comp) RingYellow(scale int) *Comp {
	c.El.Style(styles.RingYellow(scale))
	return c
}

func (c *Comp) RingYellowAlpha(scale int) *Comp {
	c.El.Style(styles.RingYellowAlpha(scale))
	return c
}

func (c *Comp) RingYellowDark(scale int) *Comp {
	c.El.Style(styles.RingYellowDark(scale))
	return c
}

func (c *Comp) RingYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.RingYellowDarkAlpha(scale))
	return c
}

func (c *Comp) Rotate(val any) *Comp {
	c.El.Style(styles.Rotate(val))
	return c
}

func (c *Comp) RotateNone() *Comp {
	c.El.Style(styles.RotateNone())
	return c
}

func (c *Comp) RotateX(val any) *Comp {
	c.El.Style(styles.RotateX(val))
	return c
}

func (c *Comp) RotateY(val any) *Comp {
	c.El.Style(styles.RotateY(val))
	return c
}

func (c *Comp) RotateZ(val any) *Comp {
	c.El.Style(styles.RotateZ(val))
	return c
}

func (c *Comp) Rounded(val value.Value) *Comp {
	c.El.Style(styles.Rounded(val))
	return c
}

func (c *Comp) Rounded2xl() *Comp {
	c.El.Style(styles.Rounded2xl())
	return c
}

func (c *Comp) Rounded3xl() *Comp {
	c.El.Style(styles.Rounded3xl())
	return c
}

func (c *Comp) Rounded4Xl() *Comp {
	c.El.Style(styles.Rounded4Xl())
	return c
}

func (c *Comp) RoundedB(val value.Value) *Comp {
	c.El.Style(styles.RoundedB(val))
	return c
}

func (c *Comp) RoundedB2xl() *Comp {
	c.El.Style(styles.RoundedB2xl())
	return c
}

func (c *Comp) RoundedB3xl() *Comp {
	c.El.Style(styles.RoundedB3xl())
	return c
}

func (c *Comp) RoundedB4Xl() *Comp {
	c.El.Style(styles.RoundedB4Xl())
	return c
}

func (c *Comp) RoundedBFull() *Comp {
	c.El.Style(styles.RoundedBFull())
	return c
}

func (c *Comp) RoundedBLg() *Comp {
	c.El.Style(styles.RoundedBLg())
	return c
}

func (c *Comp) RoundedBMd() *Comp {
	c.El.Style(styles.RoundedBMd())
	return c
}

func (c *Comp) RoundedBNone() *Comp {
	c.El.Style(styles.RoundedBNone())
	return c
}

func (c *Comp) RoundedBSm() *Comp {
	c.El.Style(styles.RoundedBSm())
	return c
}

func (c *Comp) RoundedBXl() *Comp {
	c.El.Style(styles.RoundedBXl())
	return c
}

func (c *Comp) RoundedBXs() *Comp {
	c.El.Style(styles.RoundedBXs())
	return c
}

func (c *Comp) RoundedBl(val value.Value) *Comp {
	c.El.Style(styles.RoundedBl(val))
	return c
}

func (c *Comp) RoundedBl2xl() *Comp {
	c.El.Style(styles.RoundedBl2xl())
	return c
}

func (c *Comp) RoundedBl3xl() *Comp {
	c.El.Style(styles.RoundedBl3xl())
	return c
}

func (c *Comp) RoundedBl4Xl() *Comp {
	c.El.Style(styles.RoundedBl4Xl())
	return c
}

func (c *Comp) RoundedBlFull() *Comp {
	c.El.Style(styles.RoundedBlFull())
	return c
}

func (c *Comp) RoundedBlLg() *Comp {
	c.El.Style(styles.RoundedBlLg())
	return c
}

func (c *Comp) RoundedBlMd() *Comp {
	c.El.Style(styles.RoundedBlMd())
	return c
}

func (c *Comp) RoundedBlNone() *Comp {
	c.El.Style(styles.RoundedBlNone())
	return c
}

func (c *Comp) RoundedBlSm() *Comp {
	c.El.Style(styles.RoundedBlSm())
	return c
}

func (c *Comp) RoundedBlXl() *Comp {
	c.El.Style(styles.RoundedBlXl())
	return c
}

func (c *Comp) RoundedBlXs() *Comp {
	c.El.Style(styles.RoundedBlXs())
	return c
}

func (c *Comp) RoundedBr(val value.Value) *Comp {
	c.El.Style(styles.RoundedBr(val))
	return c
}

func (c *Comp) RoundedBr2xl() *Comp {
	c.El.Style(styles.RoundedBr2xl())
	return c
}

func (c *Comp) RoundedBr3xl() *Comp {
	c.El.Style(styles.RoundedBr3xl())
	return c
}

func (c *Comp) RoundedBr4Xl() *Comp {
	c.El.Style(styles.RoundedBr4Xl())
	return c
}

func (c *Comp) RoundedBrFull() *Comp {
	c.El.Style(styles.RoundedBrFull())
	return c
}

func (c *Comp) RoundedBrLg() *Comp {
	c.El.Style(styles.RoundedBrLg())
	return c
}

func (c *Comp) RoundedBrMd() *Comp {
	c.El.Style(styles.RoundedBrMd())
	return c
}

func (c *Comp) RoundedBrNone() *Comp {
	c.El.Style(styles.RoundedBrNone())
	return c
}

func (c *Comp) RoundedBrSm() *Comp {
	c.El.Style(styles.RoundedBrSm())
	return c
}

func (c *Comp) RoundedBrXl() *Comp {
	c.El.Style(styles.RoundedBrXl())
	return c
}

func (c *Comp) RoundedBrXs() *Comp {
	c.El.Style(styles.RoundedBrXs())
	return c
}

func (c *Comp) RoundedE(val value.Value) *Comp {
	c.El.Style(styles.RoundedE(val))
	return c
}

func (c *Comp) RoundedE2xl() *Comp {
	c.El.Style(styles.RoundedE2xl())
	return c
}

func (c *Comp) RoundedE3xl() *Comp {
	c.El.Style(styles.RoundedE3xl())
	return c
}

func (c *Comp) RoundedE4Xl() *Comp {
	c.El.Style(styles.RoundedE4Xl())
	return c
}

func (c *Comp) RoundedEFull() *Comp {
	c.El.Style(styles.RoundedEFull())
	return c
}

func (c *Comp) RoundedELg() *Comp {
	c.El.Style(styles.RoundedELg())
	return c
}

func (c *Comp) RoundedEMd() *Comp {
	c.El.Style(styles.RoundedEMd())
	return c
}

func (c *Comp) RoundedENone() *Comp {
	c.El.Style(styles.RoundedENone())
	return c
}

func (c *Comp) RoundedESm() *Comp {
	c.El.Style(styles.RoundedESm())
	return c
}

func (c *Comp) RoundedEXl() *Comp {
	c.El.Style(styles.RoundedEXl())
	return c
}

func (c *Comp) RoundedEXs() *Comp {
	c.El.Style(styles.RoundedEXs())
	return c
}

func (c *Comp) RoundedEe(val value.Value) *Comp {
	c.El.Style(styles.RoundedEe(val))
	return c
}

func (c *Comp) RoundedEe2xl() *Comp {
	c.El.Style(styles.RoundedEe2xl())
	return c
}

func (c *Comp) RoundedEe3xl() *Comp {
	c.El.Style(styles.RoundedEe3xl())
	return c
}

func (c *Comp) RoundedEe4Xl() *Comp {
	c.El.Style(styles.RoundedEe4Xl())
	return c
}

func (c *Comp) RoundedEeFull() *Comp {
	c.El.Style(styles.RoundedEeFull())
	return c
}

func (c *Comp) RoundedEeLg() *Comp {
	c.El.Style(styles.RoundedEeLg())
	return c
}

func (c *Comp) RoundedEeMd() *Comp {
	c.El.Style(styles.RoundedEeMd())
	return c
}

func (c *Comp) RoundedEeNone() *Comp {
	c.El.Style(styles.RoundedEeNone())
	return c
}

func (c *Comp) RoundedEeSm() *Comp {
	c.El.Style(styles.RoundedEeSm())
	return c
}

func (c *Comp) RoundedEeXl() *Comp {
	c.El.Style(styles.RoundedEeXl())
	return c
}

func (c *Comp) RoundedEeXs() *Comp {
	c.El.Style(styles.RoundedEeXs())
	return c
}

func (c *Comp) RoundedEs(val value.Value) *Comp {
	c.El.Style(styles.RoundedEs(val))
	return c
}

func (c *Comp) RoundedEs2xl() *Comp {
	c.El.Style(styles.RoundedEs2xl())
	return c
}

func (c *Comp) RoundedEs3xl() *Comp {
	c.El.Style(styles.RoundedEs3xl())
	return c
}

func (c *Comp) RoundedEs4Xl() *Comp {
	c.El.Style(styles.RoundedEs4Xl())
	return c
}

func (c *Comp) RoundedEsFull() *Comp {
	c.El.Style(styles.RoundedEsFull())
	return c
}

func (c *Comp) RoundedEsLg() *Comp {
	c.El.Style(styles.RoundedEsLg())
	return c
}

func (c *Comp) RoundedEsMd() *Comp {
	c.El.Style(styles.RoundedEsMd())
	return c
}

func (c *Comp) RoundedEsNone() *Comp {
	c.El.Style(styles.RoundedEsNone())
	return c
}

func (c *Comp) RoundedEsSm() *Comp {
	c.El.Style(styles.RoundedEsSm())
	return c
}

func (c *Comp) RoundedEsXl() *Comp {
	c.El.Style(styles.RoundedEsXl())
	return c
}

func (c *Comp) RoundedEsXs() *Comp {
	c.El.Style(styles.RoundedEsXs())
	return c
}

func (c *Comp) RoundedFull() *Comp {
	c.El.Style(styles.RoundedFull())
	return c
}

func (c *Comp) RoundedL(val value.Value) *Comp {
	c.El.Style(styles.RoundedL(val))
	return c
}

func (c *Comp) RoundedL2xl() *Comp {
	c.El.Style(styles.RoundedL2xl())
	return c
}

func (c *Comp) RoundedL3xl() *Comp {
	c.El.Style(styles.RoundedL3xl())
	return c
}

func (c *Comp) RoundedL4Xl() *Comp {
	c.El.Style(styles.RoundedL4Xl())
	return c
}

func (c *Comp) RoundedLFull() *Comp {
	c.El.Style(styles.RoundedLFull())
	return c
}

func (c *Comp) RoundedLLg() *Comp {
	c.El.Style(styles.RoundedLLg())
	return c
}

func (c *Comp) RoundedLMd() *Comp {
	c.El.Style(styles.RoundedLMd())
	return c
}

func (c *Comp) RoundedLNone() *Comp {
	c.El.Style(styles.RoundedLNone())
	return c
}

func (c *Comp) RoundedLSm() *Comp {
	c.El.Style(styles.RoundedLSm())
	return c
}

func (c *Comp) RoundedLXl() *Comp {
	c.El.Style(styles.RoundedLXl())
	return c
}

func (c *Comp) RoundedLXs() *Comp {
	c.El.Style(styles.RoundedLXs())
	return c
}

func (c *Comp) RoundedLg() *Comp {
	c.El.Style(styles.RoundedLg())
	return c
}

func (c *Comp) RoundedMd() *Comp {
	c.El.Style(styles.RoundedMd())
	return c
}

func (c *Comp) RoundedNone() *Comp {
	c.El.Style(styles.RoundedNone())
	return c
}

func (c *Comp) RoundedR(val value.Value) *Comp {
	c.El.Style(styles.RoundedR(val))
	return c
}

func (c *Comp) RoundedR2xl() *Comp {
	c.El.Style(styles.RoundedR2xl())
	return c
}

func (c *Comp) RoundedR3xl() *Comp {
	c.El.Style(styles.RoundedR3xl())
	return c
}

func (c *Comp) RoundedR4Xl() *Comp {
	c.El.Style(styles.RoundedR4Xl())
	return c
}

func (c *Comp) RoundedRFull() *Comp {
	c.El.Style(styles.RoundedRFull())
	return c
}

func (c *Comp) RoundedRLg() *Comp {
	c.El.Style(styles.RoundedRLg())
	return c
}

func (c *Comp) RoundedRMd() *Comp {
	c.El.Style(styles.RoundedRMd())
	return c
}

func (c *Comp) RoundedRNone() *Comp {
	c.El.Style(styles.RoundedRNone())
	return c
}

func (c *Comp) RoundedRSm() *Comp {
	c.El.Style(styles.RoundedRSm())
	return c
}

func (c *Comp) RoundedRXl() *Comp {
	c.El.Style(styles.RoundedRXl())
	return c
}

func (c *Comp) RoundedRXs() *Comp {
	c.El.Style(styles.RoundedRXs())
	return c
}

func (c *Comp) RoundedS(val value.Value) *Comp {
	c.El.Style(styles.RoundedS(val))
	return c
}

func (c *Comp) RoundedS2xl() *Comp {
	c.El.Style(styles.RoundedS2xl())
	return c
}

func (c *Comp) RoundedS3xl() *Comp {
	c.El.Style(styles.RoundedS3xl())
	return c
}

func (c *Comp) RoundedS4Xl() *Comp {
	c.El.Style(styles.RoundedS4Xl())
	return c
}

func (c *Comp) RoundedSFull() *Comp {
	c.El.Style(styles.RoundedSFull())
	return c
}

func (c *Comp) RoundedSLg() *Comp {
	c.El.Style(styles.RoundedSLg())
	return c
}

func (c *Comp) RoundedSMd() *Comp {
	c.El.Style(styles.RoundedSMd())
	return c
}

func (c *Comp) RoundedSNone() *Comp {
	c.El.Style(styles.RoundedSNone())
	return c
}

func (c *Comp) RoundedSSm() *Comp {
	c.El.Style(styles.RoundedSSm())
	return c
}

func (c *Comp) RoundedSXl() *Comp {
	c.El.Style(styles.RoundedSXl())
	return c
}

func (c *Comp) RoundedSXs() *Comp {
	c.El.Style(styles.RoundedSXs())
	return c
}

func (c *Comp) RoundedSe(val value.Value) *Comp {
	c.El.Style(styles.RoundedSe(val))
	return c
}

func (c *Comp) RoundedSe2xl() *Comp {
	c.El.Style(styles.RoundedSe2xl())
	return c
}

func (c *Comp) RoundedSe3xl() *Comp {
	c.El.Style(styles.RoundedSe3xl())
	return c
}

func (c *Comp) RoundedSe4Xl() *Comp {
	c.El.Style(styles.RoundedSe4Xl())
	return c
}

func (c *Comp) RoundedSeFull() *Comp {
	c.El.Style(styles.RoundedSeFull())
	return c
}

func (c *Comp) RoundedSeLg() *Comp {
	c.El.Style(styles.RoundedSeLg())
	return c
}

func (c *Comp) RoundedSeMd() *Comp {
	c.El.Style(styles.RoundedSeMd())
	return c
}

func (c *Comp) RoundedSeNone() *Comp {
	c.El.Style(styles.RoundedSeNone())
	return c
}

func (c *Comp) RoundedSeSm() *Comp {
	c.El.Style(styles.RoundedSeSm())
	return c
}

func (c *Comp) RoundedSeXl() *Comp {
	c.El.Style(styles.RoundedSeXl())
	return c
}

func (c *Comp) RoundedSeXs() *Comp {
	c.El.Style(styles.RoundedSeXs())
	return c
}

func (c *Comp) RoundedSm() *Comp {
	c.El.Style(styles.RoundedSm())
	return c
}

func (c *Comp) RoundedSs(val value.Value) *Comp {
	c.El.Style(styles.RoundedSs(val))
	return c
}

func (c *Comp) RoundedSs2xl() *Comp {
	c.El.Style(styles.RoundedSs2xl())
	return c
}

func (c *Comp) RoundedSs3xl() *Comp {
	c.El.Style(styles.RoundedSs3xl())
	return c
}

func (c *Comp) RoundedSs4Xl() *Comp {
	c.El.Style(styles.RoundedSs4Xl())
	return c
}

func (c *Comp) RoundedSsFull() *Comp {
	c.El.Style(styles.RoundedSsFull())
	return c
}

func (c *Comp) RoundedSsLg() *Comp {
	c.El.Style(styles.RoundedSsLg())
	return c
}

func (c *Comp) RoundedSsMd() *Comp {
	c.El.Style(styles.RoundedSsMd())
	return c
}

func (c *Comp) RoundedSsNone() *Comp {
	c.El.Style(styles.RoundedSsNone())
	return c
}

func (c *Comp) RoundedSsSm() *Comp {
	c.El.Style(styles.RoundedSsSm())
	return c
}

func (c *Comp) RoundedSsXl() *Comp {
	c.El.Style(styles.RoundedSsXl())
	return c
}

func (c *Comp) RoundedSsXs() *Comp {
	c.El.Style(styles.RoundedSsXs())
	return c
}

func (c *Comp) RoundedT(val value.Value) *Comp {
	c.El.Style(styles.RoundedT(val))
	return c
}

func (c *Comp) RoundedT2xl() *Comp {
	c.El.Style(styles.RoundedT2xl())
	return c
}

func (c *Comp) RoundedT3xl() *Comp {
	c.El.Style(styles.RoundedT3xl())
	return c
}

func (c *Comp) RoundedT4Xl() *Comp {
	c.El.Style(styles.RoundedT4Xl())
	return c
}

func (c *Comp) RoundedTFull() *Comp {
	c.El.Style(styles.RoundedTFull())
	return c
}

func (c *Comp) RoundedTLg() *Comp {
	c.El.Style(styles.RoundedTLg())
	return c
}

func (c *Comp) RoundedTMd() *Comp {
	c.El.Style(styles.RoundedTMd())
	return c
}

func (c *Comp) RoundedTNone() *Comp {
	c.El.Style(styles.RoundedTNone())
	return c
}

func (c *Comp) RoundedTSm() *Comp {
	c.El.Style(styles.RoundedTSm())
	return c
}

func (c *Comp) RoundedTXl() *Comp {
	c.El.Style(styles.RoundedTXl())
	return c
}

func (c *Comp) RoundedTXs() *Comp {
	c.El.Style(styles.RoundedTXs())
	return c
}

func (c *Comp) RoundedTl(val value.Value) *Comp {
	c.El.Style(styles.RoundedTl(val))
	return c
}

func (c *Comp) RoundedTl2xl() *Comp {
	c.El.Style(styles.RoundedTl2xl())
	return c
}

func (c *Comp) RoundedTl3xl() *Comp {
	c.El.Style(styles.RoundedTl3xl())
	return c
}

func (c *Comp) RoundedTl4Xl() *Comp {
	c.El.Style(styles.RoundedTl4Xl())
	return c
}

func (c *Comp) RoundedTlFull() *Comp {
	c.El.Style(styles.RoundedTlFull())
	return c
}

func (c *Comp) RoundedTlLg() *Comp {
	c.El.Style(styles.RoundedTlLg())
	return c
}

func (c *Comp) RoundedTlMd() *Comp {
	c.El.Style(styles.RoundedTlMd())
	return c
}

func (c *Comp) RoundedTlNone() *Comp {
	c.El.Style(styles.RoundedTlNone())
	return c
}

func (c *Comp) RoundedTlSm() *Comp {
	c.El.Style(styles.RoundedTlSm())
	return c
}

func (c *Comp) RoundedTlXl() *Comp {
	c.El.Style(styles.RoundedTlXl())
	return c
}

func (c *Comp) RoundedTlXs() *Comp {
	c.El.Style(styles.RoundedTlXs())
	return c
}

func (c *Comp) RoundedTr(val value.Value) *Comp {
	c.El.Style(styles.RoundedTr(val))
	return c
}

func (c *Comp) RoundedTr2xl() *Comp {
	c.El.Style(styles.RoundedTr2xl())
	return c
}

func (c *Comp) RoundedTr3xl() *Comp {
	c.El.Style(styles.RoundedTr3xl())
	return c
}

func (c *Comp) RoundedTr4Xl() *Comp {
	c.El.Style(styles.RoundedTr4Xl())
	return c
}

func (c *Comp) RoundedTrFull() *Comp {
	c.El.Style(styles.RoundedTrFull())
	return c
}

func (c *Comp) RoundedTrLg() *Comp {
	c.El.Style(styles.RoundedTrLg())
	return c
}

func (c *Comp) RoundedTrMd() *Comp {
	c.El.Style(styles.RoundedTrMd())
	return c
}

func (c *Comp) RoundedTrNone() *Comp {
	c.El.Style(styles.RoundedTrNone())
	return c
}

func (c *Comp) RoundedTrSm() *Comp {
	c.El.Style(styles.RoundedTrSm())
	return c
}

func (c *Comp) RoundedTrXl() *Comp {
	c.El.Style(styles.RoundedTrXl())
	return c
}

func (c *Comp) RoundedTrXs() *Comp {
	c.El.Style(styles.RoundedTrXs())
	return c
}

func (c *Comp) RoundedXl() *Comp {
	c.El.Style(styles.RoundedXl())
	return c
}

func (c *Comp) RoundedXs() *Comp {
	c.El.Style(styles.RoundedXs())
	return c
}

func (c *Comp) Row(number int) *Comp {
	c.El.Style(styles.Row(number))
	return c
}

func (c *Comp) RowAuto() *Comp {
	c.El.Style(styles.RowAuto())
	return c
}

func (c *Comp) RowBy(val value.Value) *Comp {
	c.El.Style(styles.RowBy(val))
	return c
}

func (c *Comp) RowEnd(number int) *Comp {
	c.El.Style(styles.RowEnd(number))
	return c
}

func (c *Comp) RowEndAuto() *Comp {
	c.El.Style(styles.RowEndAuto())
	return c
}

func (c *Comp) RowEndBy(val value.Value) *Comp {
	c.El.Style(styles.RowEndBy(val))
	return c
}

func (c *Comp) RowSpan(number int) *Comp {
	c.El.Style(styles.RowSpan(number))
	return c
}

func (c *Comp) RowSpanBy(val value.Value) *Comp {
	c.El.Style(styles.RowSpanBy(val))
	return c
}

func (c *Comp) RowSpanFull() *Comp {
	c.El.Style(styles.RowSpanFull())
	return c
}

func (c *Comp) RowStart(number int) *Comp {
	c.El.Style(styles.RowStart(number))
	return c
}

func (c *Comp) RowStartAuto() *Comp {
	c.El.Style(styles.RowStartAuto())
	return c
}

func (c *Comp) RowStartBy(val value.Value) *Comp {
	c.El.Style(styles.RowStartBy(val))
	return c
}

func (c *Comp) Saturate(val any) *Comp {
	c.El.Style(styles.Saturate(val))
	return c
}

func (c *Comp) Scale(val any) *Comp {
	c.El.Style(styles.Scale(val))
	return c
}

func (c *Comp) Scale3d() *Comp {
	c.El.Style(styles.Scale3d())
	return c
}

func (c *Comp) ScaleNone() *Comp {
	c.El.Style(styles.ScaleNone())
	return c
}

func (c *Comp) ScaleX(val any) *Comp {
	c.El.Style(styles.ScaleX(val))
	return c
}

func (c *Comp) ScaleY(val any) *Comp {
	c.El.Style(styles.ScaleY(val))
	return c
}

func (c *Comp) ScaleZ(val any) *Comp {
	c.El.Style(styles.ScaleZ(val))
	return c
}

func (c *Comp) SchemeDark() *Comp {
	c.El.Style(styles.SchemeDark())
	return c
}

func (c *Comp) SchemeLight() *Comp {
	c.El.Style(styles.SchemeLight())
	return c
}

func (c *Comp) SchemeLightDark() *Comp {
	c.El.Style(styles.SchemeLightDark())
	return c
}

func (c *Comp) SchemeNormal() *Comp {
	c.El.Style(styles.SchemeNormal())
	return c
}

func (c *Comp) SchemeOnlyDark() *Comp {
	c.El.Style(styles.SchemeOnlyDark())
	return c
}

func (c *Comp) SchemeOnlyLight() *Comp {
	c.El.Style(styles.SchemeOnlyLight())
	return c
}

func (c *Comp) ScrollAuto() *Comp {
	c.El.Style(styles.ScrollAuto())
	return c
}

func (c *Comp) ScrollM(val any) *Comp {
	c.El.Style(styles.ScrollM(val))
	return c
}

func (c *Comp) ScrollMb(val any) *Comp {
	c.El.Style(styles.ScrollMb(val))
	return c
}

func (c *Comp) ScrollMe(val any) *Comp {
	c.El.Style(styles.ScrollMe(val))
	return c
}

func (c *Comp) ScrollMl(val any) *Comp {
	c.El.Style(styles.ScrollMl(val))
	return c
}

func (c *Comp) ScrollMr(val any) *Comp {
	c.El.Style(styles.ScrollMr(val))
	return c
}

func (c *Comp) ScrollMs(val any) *Comp {
	c.El.Style(styles.ScrollMs(val))
	return c
}

func (c *Comp) ScrollMt(val any) *Comp {
	c.El.Style(styles.ScrollMt(val))
	return c
}

func (c *Comp) ScrollMx(val any) *Comp {
	c.El.Style(styles.ScrollMx(val))
	return c
}

func (c *Comp) ScrollMy(val any) *Comp {
	c.El.Style(styles.ScrollMy(val))
	return c
}

func (c *Comp) ScrollP(val any) *Comp {
	c.El.Style(styles.ScrollP(val))
	return c
}

func (c *Comp) ScrollPb(val any) *Comp {
	c.El.Style(styles.ScrollPb(val))
	return c
}

func (c *Comp) ScrollPe(val any) *Comp {
	c.El.Style(styles.ScrollPe(val))
	return c
}

func (c *Comp) ScrollPl(val any) *Comp {
	c.El.Style(styles.ScrollPl(val))
	return c
}

func (c *Comp) ScrollPr(val any) *Comp {
	c.El.Style(styles.ScrollPr(val))
	return c
}

func (c *Comp) ScrollPs(val any) *Comp {
	c.El.Style(styles.ScrollPs(val))
	return c
}

func (c *Comp) ScrollPt(val any) *Comp {
	c.El.Style(styles.ScrollPt(val))
	return c
}

func (c *Comp) ScrollPx(val any) *Comp {
	c.El.Style(styles.ScrollPx(val))
	return c
}

func (c *Comp) ScrollPy(val any) *Comp {
	c.El.Style(styles.ScrollPy(val))
	return c
}

func (c *Comp) ScrollSmooth() *Comp {
	c.El.Style(styles.ScrollSmooth())
	return c
}

func (c *Comp) SelectAll() *Comp {
	c.El.Style(styles.SelectAll())
	return c
}

func (c *Comp) SelectAuto() *Comp {
	c.El.Style(styles.SelectAuto())
	return c
}

func (c *Comp) SelectNone() *Comp {
	c.El.Style(styles.SelectNone())
	return c
}

func (c *Comp) SelectText() *Comp {
	c.El.Style(styles.SelectText())
	return c
}

func (c *Comp) SelectorParam(param string) *Comp {
	c.El.Style(styles.SelectorParam(param))
	return c
}

func (c *Comp) SelfAuto() *Comp {
	c.El.Style(styles.SelfAuto())
	return c
}

func (c *Comp) SelfBaseline() *Comp {
	c.El.Style(styles.SelfBaseline())
	return c
}

func (c *Comp) SelfBaselineLast() *Comp {
	c.El.Style(styles.SelfBaselineLast())
	return c
}

func (c *Comp) SelfCenter() *Comp {
	c.El.Style(styles.SelfCenter())
	return c
}

func (c *Comp) SelfCenterSafe() *Comp {
	c.El.Style(styles.SelfCenterSafe())
	return c
}

func (c *Comp) SelfEnd() *Comp {
	c.El.Style(styles.SelfEnd())
	return c
}

func (c *Comp) SelfEndSafe() *Comp {
	c.El.Style(styles.SelfEndSafe())
	return c
}

func (c *Comp) SelfStart() *Comp {
	c.El.Style(styles.SelfStart())
	return c
}

func (c *Comp) SelfStretch() *Comp {
	c.El.Style(styles.SelfStretch())
	return c
}

func (c *Comp) Sepia(val ...any) *Comp {
	c.El.Style(styles.Sepia(val...))
	return c
}

func (c *Comp) Shadow(val value.Value) *Comp {
	c.El.Style(styles.Shadow(val))
	return c
}

func (c *Comp) Shadow2xl() *Comp {
	c.El.Style(styles.Shadow2xl())
	return c
}

func (c *Comp) Shadow2xs() *Comp {
	c.El.Style(styles.Shadow2xs())
	return c
}

func (c *Comp) ShadowAmber(scale int) *Comp {
	c.El.Style(styles.ShadowAmber(scale))
	return c
}

func (c *Comp) ShadowAmberAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowAmberAlpha(scale))
	return c
}

func (c *Comp) ShadowAmberDark(scale int) *Comp {
	c.El.Style(styles.ShadowAmberDark(scale))
	return c
}

func (c *Comp) ShadowAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowAmberDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowBlack() *Comp {
	c.El.Style(styles.ShadowBlack())
	return c
}

func (c *Comp) ShadowBlackAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowBlackAlpha(scale))
	return c
}

func (c *Comp) ShadowBlue(scale int) *Comp {
	c.El.Style(styles.ShadowBlue(scale))
	return c
}

func (c *Comp) ShadowBlueAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowBlueAlpha(scale))
	return c
}

func (c *Comp) ShadowBlueDark(scale int) *Comp {
	c.El.Style(styles.ShadowBlueDark(scale))
	return c
}

func (c *Comp) ShadowBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowBlueDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowBronze(scale int) *Comp {
	c.El.Style(styles.ShadowBronze(scale))
	return c
}

func (c *Comp) ShadowBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowBronzeAlpha(scale))
	return c
}

func (c *Comp) ShadowBronzeDark(scale int) *Comp {
	c.El.Style(styles.ShadowBronzeDark(scale))
	return c
}

func (c *Comp) ShadowBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowBrown(scale int) *Comp {
	c.El.Style(styles.ShadowBrown(scale))
	return c
}

func (c *Comp) ShadowBrownAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowBrownAlpha(scale))
	return c
}

func (c *Comp) ShadowBrownDark(scale int) *Comp {
	c.El.Style(styles.ShadowBrownDark(scale))
	return c
}

func (c *Comp) ShadowBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowBrownDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowColor(val value.Value) *Comp {
	c.El.Style(styles.ShadowColor(val))
	return c
}

func (c *Comp) ShadowCrimson(scale int) *Comp {
	c.El.Style(styles.ShadowCrimson(scale))
	return c
}

func (c *Comp) ShadowCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowCrimsonAlpha(scale))
	return c
}

func (c *Comp) ShadowCrimsonDark(scale int) *Comp {
	c.El.Style(styles.ShadowCrimsonDark(scale))
	return c
}

func (c *Comp) ShadowCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowCurrent() *Comp {
	c.El.Style(styles.ShadowCurrent())
	return c
}

func (c *Comp) ShadowCyan(scale int) *Comp {
	c.El.Style(styles.ShadowCyan(scale))
	return c
}

func (c *Comp) ShadowCyanAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowCyanAlpha(scale))
	return c
}

func (c *Comp) ShadowCyanDark(scale int) *Comp {
	c.El.Style(styles.ShadowCyanDark(scale))
	return c
}

func (c *Comp) ShadowCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowCyanDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowGold(scale int) *Comp {
	c.El.Style(styles.ShadowGold(scale))
	return c
}

func (c *Comp) ShadowGoldAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowGoldAlpha(scale))
	return c
}

func (c *Comp) ShadowGoldDark(scale int) *Comp {
	c.El.Style(styles.ShadowGoldDark(scale))
	return c
}

func (c *Comp) ShadowGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowGoldDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowGrass(scale int) *Comp {
	c.El.Style(styles.ShadowGrass(scale))
	return c
}

func (c *Comp) ShadowGrassAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowGrassAlpha(scale))
	return c
}

func (c *Comp) ShadowGrassDark(scale int) *Comp {
	c.El.Style(styles.ShadowGrassDark(scale))
	return c
}

func (c *Comp) ShadowGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowGrassDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowGray(scale int) *Comp {
	c.El.Style(styles.ShadowGray(scale))
	return c
}

func (c *Comp) ShadowGrayAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowGrayAlpha(scale))
	return c
}

func (c *Comp) ShadowGrayDark(scale int) *Comp {
	c.El.Style(styles.ShadowGrayDark(scale))
	return c
}

func (c *Comp) ShadowGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowGrayDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowGreen(scale int) *Comp {
	c.El.Style(styles.ShadowGreen(scale))
	return c
}

func (c *Comp) ShadowGreenAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowGreenAlpha(scale))
	return c
}

func (c *Comp) ShadowGreenDark(scale int) *Comp {
	c.El.Style(styles.ShadowGreenDark(scale))
	return c
}

func (c *Comp) ShadowGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowGreenDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowIndigo(scale int) *Comp {
	c.El.Style(styles.ShadowIndigo(scale))
	return c
}

func (c *Comp) ShadowIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowIndigoAlpha(scale))
	return c
}

func (c *Comp) ShadowIndigoDark(scale int) *Comp {
	c.El.Style(styles.ShadowIndigoDark(scale))
	return c
}

func (c *Comp) ShadowIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowInherit() *Comp {
	c.El.Style(styles.ShadowInherit())
	return c
}

func (c *Comp) ShadowIris(scale int) *Comp {
	c.El.Style(styles.ShadowIris(scale))
	return c
}

func (c *Comp) ShadowIrisAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowIrisAlpha(scale))
	return c
}

func (c *Comp) ShadowIrisDark(scale int) *Comp {
	c.El.Style(styles.ShadowIrisDark(scale))
	return c
}

func (c *Comp) ShadowIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowIrisDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowJade(scale int) *Comp {
	c.El.Style(styles.ShadowJade(scale))
	return c
}

func (c *Comp) ShadowJadeAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowJadeAlpha(scale))
	return c
}

func (c *Comp) ShadowJadeDark(scale int) *Comp {
	c.El.Style(styles.ShadowJadeDark(scale))
	return c
}

func (c *Comp) ShadowJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowJadeDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowLg() *Comp {
	c.El.Style(styles.ShadowLg())
	return c
}

func (c *Comp) ShadowLime(scale int) *Comp {
	c.El.Style(styles.ShadowLime(scale))
	return c
}

func (c *Comp) ShadowLimeAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowLimeAlpha(scale))
	return c
}

func (c *Comp) ShadowLimeDark(scale int) *Comp {
	c.El.Style(styles.ShadowLimeDark(scale))
	return c
}

func (c *Comp) ShadowLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowLimeDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowMauve(scale int) *Comp {
	c.El.Style(styles.ShadowMauve(scale))
	return c
}

func (c *Comp) ShadowMauveAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowMauveAlpha(scale))
	return c
}

func (c *Comp) ShadowMauveDark(scale int) *Comp {
	c.El.Style(styles.ShadowMauveDark(scale))
	return c
}

func (c *Comp) ShadowMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowMauveDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowMd() *Comp {
	c.El.Style(styles.ShadowMd())
	return c
}

func (c *Comp) ShadowMint(scale int) *Comp {
	c.El.Style(styles.ShadowMint(scale))
	return c
}

func (c *Comp) ShadowMintAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowMintAlpha(scale))
	return c
}

func (c *Comp) ShadowMintDark(scale int) *Comp {
	c.El.Style(styles.ShadowMintDark(scale))
	return c
}

func (c *Comp) ShadowMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowMintDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowNone() *Comp {
	c.El.Style(styles.ShadowNone())
	return c
}

func (c *Comp) ShadowOlive(scale int) *Comp {
	c.El.Style(styles.ShadowOlive(scale))
	return c
}

func (c *Comp) ShadowOliveAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowOliveAlpha(scale))
	return c
}

func (c *Comp) ShadowOliveDark(scale int) *Comp {
	c.El.Style(styles.ShadowOliveDark(scale))
	return c
}

func (c *Comp) ShadowOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowOliveDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowOrange(scale int) *Comp {
	c.El.Style(styles.ShadowOrange(scale))
	return c
}

func (c *Comp) ShadowOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowOrangeAlpha(scale))
	return c
}

func (c *Comp) ShadowOrangeDark(scale int) *Comp {
	c.El.Style(styles.ShadowOrangeDark(scale))
	return c
}

func (c *Comp) ShadowOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowPink(scale int) *Comp {
	c.El.Style(styles.ShadowPink(scale))
	return c
}

func (c *Comp) ShadowPinkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowPinkAlpha(scale))
	return c
}

func (c *Comp) ShadowPinkDark(scale int) *Comp {
	c.El.Style(styles.ShadowPinkDark(scale))
	return c
}

func (c *Comp) ShadowPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowPinkDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowPlum(scale int) *Comp {
	c.El.Style(styles.ShadowPlum(scale))
	return c
}

func (c *Comp) ShadowPlumAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowPlumAlpha(scale))
	return c
}

func (c *Comp) ShadowPlumDark(scale int) *Comp {
	c.El.Style(styles.ShadowPlumDark(scale))
	return c
}

func (c *Comp) ShadowPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowPlumDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowPurple(scale int) *Comp {
	c.El.Style(styles.ShadowPurple(scale))
	return c
}

func (c *Comp) ShadowPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowPurpleAlpha(scale))
	return c
}

func (c *Comp) ShadowPurpleDark(scale int) *Comp {
	c.El.Style(styles.ShadowPurpleDark(scale))
	return c
}

func (c *Comp) ShadowPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowRed(scale int) *Comp {
	c.El.Style(styles.ShadowRed(scale))
	return c
}

func (c *Comp) ShadowRedAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowRedAlpha(scale))
	return c
}

func (c *Comp) ShadowRedDark(scale int) *Comp {
	c.El.Style(styles.ShadowRedDark(scale))
	return c
}

func (c *Comp) ShadowRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowRedDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowRuby(scale int) *Comp {
	c.El.Style(styles.ShadowRuby(scale))
	return c
}

func (c *Comp) ShadowRubyAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowRubyAlpha(scale))
	return c
}

func (c *Comp) ShadowRubyDark(scale int) *Comp {
	c.El.Style(styles.ShadowRubyDark(scale))
	return c
}

func (c *Comp) ShadowRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowRubyDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowSage(scale int) *Comp {
	c.El.Style(styles.ShadowSage(scale))
	return c
}

func (c *Comp) ShadowSageAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowSageAlpha(scale))
	return c
}

func (c *Comp) ShadowSageDark(scale int) *Comp {
	c.El.Style(styles.ShadowSageDark(scale))
	return c
}

func (c *Comp) ShadowSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowSageDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowSand(scale int) *Comp {
	c.El.Style(styles.ShadowSand(scale))
	return c
}

func (c *Comp) ShadowSandAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowSandAlpha(scale))
	return c
}

func (c *Comp) ShadowSandDark(scale int) *Comp {
	c.El.Style(styles.ShadowSandDark(scale))
	return c
}

func (c *Comp) ShadowSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowSandDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowSky(scale int) *Comp {
	c.El.Style(styles.ShadowSky(scale))
	return c
}

func (c *Comp) ShadowSkyAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowSkyAlpha(scale))
	return c
}

func (c *Comp) ShadowSkyDark(scale int) *Comp {
	c.El.Style(styles.ShadowSkyDark(scale))
	return c
}

func (c *Comp) ShadowSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowSkyDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowSlate(scale int) *Comp {
	c.El.Style(styles.ShadowSlate(scale))
	return c
}

func (c *Comp) ShadowSlateAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowSlateAlpha(scale))
	return c
}

func (c *Comp) ShadowSlateDark(scale int) *Comp {
	c.El.Style(styles.ShadowSlateDark(scale))
	return c
}

func (c *Comp) ShadowSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowSlateDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowSm() *Comp {
	c.El.Style(styles.ShadowSm())
	return c
}

func (c *Comp) ShadowTeal(scale int) *Comp {
	c.El.Style(styles.ShadowTeal(scale))
	return c
}

func (c *Comp) ShadowTealAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowTealAlpha(scale))
	return c
}

func (c *Comp) ShadowTealDark(scale int) *Comp {
	c.El.Style(styles.ShadowTealDark(scale))
	return c
}

func (c *Comp) ShadowTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowTealDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowTomato(scale int) *Comp {
	c.El.Style(styles.ShadowTomato(scale))
	return c
}

func (c *Comp) ShadowTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowTomatoAlpha(scale))
	return c
}

func (c *Comp) ShadowTomatoDark(scale int) *Comp {
	c.El.Style(styles.ShadowTomatoDark(scale))
	return c
}

func (c *Comp) ShadowTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowTransparent() *Comp {
	c.El.Style(styles.ShadowTransparent())
	return c
}

func (c *Comp) ShadowViolet(scale int) *Comp {
	c.El.Style(styles.ShadowViolet(scale))
	return c
}

func (c *Comp) ShadowVioletAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowVioletAlpha(scale))
	return c
}

func (c *Comp) ShadowVioletDark(scale int) *Comp {
	c.El.Style(styles.ShadowVioletDark(scale))
	return c
}

func (c *Comp) ShadowVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowVioletDarkAlpha(scale))
	return c
}

func (c *Comp) ShadowWhite() *Comp {
	c.El.Style(styles.ShadowWhite())
	return c
}

func (c *Comp) ShadowWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowWhiteAlpha(scale))
	return c
}

func (c *Comp) ShadowXl() *Comp {
	c.El.Style(styles.ShadowXl())
	return c
}

func (c *Comp) ShadowXs() *Comp {
	c.El.Style(styles.ShadowXs())
	return c
}

func (c *Comp) ShadowYellow(scale int) *Comp {
	c.El.Style(styles.ShadowYellow(scale))
	return c
}

func (c *Comp) ShadowYellowAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowYellowAlpha(scale))
	return c
}

func (c *Comp) ShadowYellowDark(scale int) *Comp {
	c.El.Style(styles.ShadowYellowDark(scale))
	return c
}

func (c *Comp) ShadowYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.ShadowYellowDarkAlpha(scale))
	return c
}

func (c *Comp) Shrink() *Comp {
	c.El.Style(styles.Shrink())
	return c
}

func (c *Comp) ShrinkBy(val value.Value) *Comp {
	c.El.Style(styles.ShrinkBy(val))
	return c
}

func (c *Comp) Size(number int) *Comp {
	c.El.Style(styles.Size(number))
	return c
}

func (c *Comp) SizeAuto() *Comp {
	c.El.Style(styles.SizeAuto())
	return c
}

func (c *Comp) SizeBy(val value.Value) *Comp {
	c.El.Style(styles.SizeBy(val))
	return c
}

func (c *Comp) SizeDvh() *Comp {
	c.El.Style(styles.SizeDvh())
	return c
}

func (c *Comp) SizeDvw() *Comp {
	c.El.Style(styles.SizeDvw())
	return c
}

func (c *Comp) SizeFit() *Comp {
	c.El.Style(styles.SizeFit())
	return c
}

func (c *Comp) SizeFraction(fraction float32) *Comp {
	c.El.Style(styles.SizeFraction(fraction))
	return c
}

func (c *Comp) SizeLvh() *Comp {
	c.El.Style(styles.SizeLvh())
	return c
}

func (c *Comp) SizeLvw() *Comp {
	c.El.Style(styles.SizeLvw())
	return c
}

func (c *Comp) SizeMax() *Comp {
	c.El.Style(styles.SizeMax())
	return c
}

func (c *Comp) SizeMin() *Comp {
	c.El.Style(styles.SizeMin())
	return c
}

func (c *Comp) SizePx() *Comp {
	c.El.Style(styles.SizePx())
	return c
}

func (c *Comp) SizeSvh() *Comp {
	c.El.Style(styles.SizeSvh())
	return c
}

func (c *Comp) SizeSvw() *Comp {
	c.El.Style(styles.SizeSvw())
	return c
}

func (c *Comp) Skew(val any) *Comp {
	c.El.Style(styles.Skew(val))
	return c
}

func (c *Comp) SkewX(val any) *Comp {
	c.El.Style(styles.SkewX(val))
	return c
}

func (c *Comp) SkewY(val any) *Comp {
	c.El.Style(styles.SkewY(val))
	return c
}

func (c *Comp) SlashedZero() *Comp {
	c.El.Style(styles.SlashedZero())
	return c
}

func (c *Comp) SnapAlignNone() *Comp {
	c.El.Style(styles.SnapAlignNone())
	return c
}

func (c *Comp) SnapAlways() *Comp {
	c.El.Style(styles.SnapAlways())
	return c
}

func (c *Comp) SnapBoth() *Comp {
	c.El.Style(styles.SnapBoth())
	return c
}

func (c *Comp) SnapCenter() *Comp {
	c.El.Style(styles.SnapCenter())
	return c
}

func (c *Comp) SnapEnd() *Comp {
	c.El.Style(styles.SnapEnd())
	return c
}

func (c *Comp) SnapMandatory() *Comp {
	c.El.Style(styles.SnapMandatory())
	return c
}

func (c *Comp) SnapNone() *Comp {
	c.El.Style(styles.SnapNone())
	return c
}

func (c *Comp) SnapNormal() *Comp {
	c.El.Style(styles.SnapNormal())
	return c
}

func (c *Comp) SnapProximity() *Comp {
	c.El.Style(styles.SnapProximity())
	return c
}

func (c *Comp) SnapStart() *Comp {
	c.El.Style(styles.SnapStart())
	return c
}

func (c *Comp) SnapX() *Comp {
	c.El.Style(styles.SnapX())
	return c
}

func (c *Comp) SnapY() *Comp {
	c.El.Style(styles.SnapY())
	return c
}

func (c *Comp) SrOnly() *Comp {
	c.El.Style(styles.SrOnly())
	return c
}

func (c *Comp) StackedFractions() *Comp {
	c.El.Style(styles.StackedFractions())
	return c
}

func (c *Comp) Start(number int) *Comp {
	c.El.Style(styles.Start(number))
	return c
}

func (c *Comp) StartAuto() *Comp {
	c.El.Style(styles.StartAuto())
	return c
}

func (c *Comp) StartBy(val value.Value) *Comp {
	c.El.Style(styles.StartBy(val))
	return c
}

func (c *Comp) StartFraction(fraction float64) *Comp {
	c.El.Style(styles.StartFraction(fraction))
	return c
}

func (c *Comp) StartFull() *Comp {
	c.El.Style(styles.StartFull())
	return c
}

func (c *Comp) StartPx() *Comp {
	c.El.Style(styles.StartPx())
	return c
}

func (c *Comp) Static() *Comp {
	c.El.Style(styles.Static())
	return c
}

func (c *Comp) Sticky() *Comp {
	c.El.Style(styles.Sticky())
	return c
}

func (c *Comp) Stroke(val any) *Comp {
	c.El.Style(styles.Stroke(val))
	return c
}

func (c *Comp) StrokeAmber(scale int) *Comp {
	c.El.Style(styles.StrokeAmber(scale))
	return c
}

func (c *Comp) StrokeAmberAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeAmberAlpha(scale))
	return c
}

func (c *Comp) StrokeAmberDark(scale int) *Comp {
	c.El.Style(styles.StrokeAmberDark(scale))
	return c
}

func (c *Comp) StrokeAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeAmberDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeBlack() *Comp {
	c.El.Style(styles.StrokeBlack())
	return c
}

func (c *Comp) StrokeBlackAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeBlackAlpha(scale))
	return c
}

func (c *Comp) StrokeBlue(scale int) *Comp {
	c.El.Style(styles.StrokeBlue(scale))
	return c
}

func (c *Comp) StrokeBlueAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeBlueAlpha(scale))
	return c
}

func (c *Comp) StrokeBlueDark(scale int) *Comp {
	c.El.Style(styles.StrokeBlueDark(scale))
	return c
}

func (c *Comp) StrokeBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeBlueDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeBronze(scale int) *Comp {
	c.El.Style(styles.StrokeBronze(scale))
	return c
}

func (c *Comp) StrokeBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeBronzeAlpha(scale))
	return c
}

func (c *Comp) StrokeBronzeDark(scale int) *Comp {
	c.El.Style(styles.StrokeBronzeDark(scale))
	return c
}

func (c *Comp) StrokeBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeBrown(scale int) *Comp {
	c.El.Style(styles.StrokeBrown(scale))
	return c
}

func (c *Comp) StrokeBrownAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeBrownAlpha(scale))
	return c
}

func (c *Comp) StrokeBrownDark(scale int) *Comp {
	c.El.Style(styles.StrokeBrownDark(scale))
	return c
}

func (c *Comp) StrokeBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeBrownDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeColor(val value.Value) *Comp {
	c.El.Style(styles.StrokeColor(val))
	return c
}

func (c *Comp) StrokeCrimson(scale int) *Comp {
	c.El.Style(styles.StrokeCrimson(scale))
	return c
}

func (c *Comp) StrokeCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeCrimsonAlpha(scale))
	return c
}

func (c *Comp) StrokeCrimsonDark(scale int) *Comp {
	c.El.Style(styles.StrokeCrimsonDark(scale))
	return c
}

func (c *Comp) StrokeCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeCurrent() *Comp {
	c.El.Style(styles.StrokeCurrent())
	return c
}

func (c *Comp) StrokeCyan(scale int) *Comp {
	c.El.Style(styles.StrokeCyan(scale))
	return c
}

func (c *Comp) StrokeCyanAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeCyanAlpha(scale))
	return c
}

func (c *Comp) StrokeCyanDark(scale int) *Comp {
	c.El.Style(styles.StrokeCyanDark(scale))
	return c
}

func (c *Comp) StrokeCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeCyanDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeGold(scale int) *Comp {
	c.El.Style(styles.StrokeGold(scale))
	return c
}

func (c *Comp) StrokeGoldAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeGoldAlpha(scale))
	return c
}

func (c *Comp) StrokeGoldDark(scale int) *Comp {
	c.El.Style(styles.StrokeGoldDark(scale))
	return c
}

func (c *Comp) StrokeGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeGoldDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeGrass(scale int) *Comp {
	c.El.Style(styles.StrokeGrass(scale))
	return c
}

func (c *Comp) StrokeGrassAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeGrassAlpha(scale))
	return c
}

func (c *Comp) StrokeGrassDark(scale int) *Comp {
	c.El.Style(styles.StrokeGrassDark(scale))
	return c
}

func (c *Comp) StrokeGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeGrassDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeGray(scale int) *Comp {
	c.El.Style(styles.StrokeGray(scale))
	return c
}

func (c *Comp) StrokeGrayAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeGrayAlpha(scale))
	return c
}

func (c *Comp) StrokeGrayDark(scale int) *Comp {
	c.El.Style(styles.StrokeGrayDark(scale))
	return c
}

func (c *Comp) StrokeGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeGrayDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeGreen(scale int) *Comp {
	c.El.Style(styles.StrokeGreen(scale))
	return c
}

func (c *Comp) StrokeGreenAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeGreenAlpha(scale))
	return c
}

func (c *Comp) StrokeGreenDark(scale int) *Comp {
	c.El.Style(styles.StrokeGreenDark(scale))
	return c
}

func (c *Comp) StrokeGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeGreenDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeIndigo(scale int) *Comp {
	c.El.Style(styles.StrokeIndigo(scale))
	return c
}

func (c *Comp) StrokeIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeIndigoAlpha(scale))
	return c
}

func (c *Comp) StrokeIndigoDark(scale int) *Comp {
	c.El.Style(styles.StrokeIndigoDark(scale))
	return c
}

func (c *Comp) StrokeIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeInherit() *Comp {
	c.El.Style(styles.StrokeInherit())
	return c
}

func (c *Comp) StrokeIris(scale int) *Comp {
	c.El.Style(styles.StrokeIris(scale))
	return c
}

func (c *Comp) StrokeIrisAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeIrisAlpha(scale))
	return c
}

func (c *Comp) StrokeIrisDark(scale int) *Comp {
	c.El.Style(styles.StrokeIrisDark(scale))
	return c
}

func (c *Comp) StrokeIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeIrisDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeJade(scale int) *Comp {
	c.El.Style(styles.StrokeJade(scale))
	return c
}

func (c *Comp) StrokeJadeAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeJadeAlpha(scale))
	return c
}

func (c *Comp) StrokeJadeDark(scale int) *Comp {
	c.El.Style(styles.StrokeJadeDark(scale))
	return c
}

func (c *Comp) StrokeJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeJadeDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeLime(scale int) *Comp {
	c.El.Style(styles.StrokeLime(scale))
	return c
}

func (c *Comp) StrokeLimeAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeLimeAlpha(scale))
	return c
}

func (c *Comp) StrokeLimeDark(scale int) *Comp {
	c.El.Style(styles.StrokeLimeDark(scale))
	return c
}

func (c *Comp) StrokeLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeLimeDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeMauve(scale int) *Comp {
	c.El.Style(styles.StrokeMauve(scale))
	return c
}

func (c *Comp) StrokeMauveAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeMauveAlpha(scale))
	return c
}

func (c *Comp) StrokeMauveDark(scale int) *Comp {
	c.El.Style(styles.StrokeMauveDark(scale))
	return c
}

func (c *Comp) StrokeMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeMauveDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeMint(scale int) *Comp {
	c.El.Style(styles.StrokeMint(scale))
	return c
}

func (c *Comp) StrokeMintAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeMintAlpha(scale))
	return c
}

func (c *Comp) StrokeMintDark(scale int) *Comp {
	c.El.Style(styles.StrokeMintDark(scale))
	return c
}

func (c *Comp) StrokeMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeMintDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeOlive(scale int) *Comp {
	c.El.Style(styles.StrokeOlive(scale))
	return c
}

func (c *Comp) StrokeOliveAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeOliveAlpha(scale))
	return c
}

func (c *Comp) StrokeOliveDark(scale int) *Comp {
	c.El.Style(styles.StrokeOliveDark(scale))
	return c
}

func (c *Comp) StrokeOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeOliveDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeOrange(scale int) *Comp {
	c.El.Style(styles.StrokeOrange(scale))
	return c
}

func (c *Comp) StrokeOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeOrangeAlpha(scale))
	return c
}

func (c *Comp) StrokeOrangeDark(scale int) *Comp {
	c.El.Style(styles.StrokeOrangeDark(scale))
	return c
}

func (c *Comp) StrokeOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) StrokePink(scale int) *Comp {
	c.El.Style(styles.StrokePink(scale))
	return c
}

func (c *Comp) StrokePinkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokePinkAlpha(scale))
	return c
}

func (c *Comp) StrokePinkDark(scale int) *Comp {
	c.El.Style(styles.StrokePinkDark(scale))
	return c
}

func (c *Comp) StrokePinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokePinkDarkAlpha(scale))
	return c
}

func (c *Comp) StrokePlum(scale int) *Comp {
	c.El.Style(styles.StrokePlum(scale))
	return c
}

func (c *Comp) StrokePlumAlpha(scale int) *Comp {
	c.El.Style(styles.StrokePlumAlpha(scale))
	return c
}

func (c *Comp) StrokePlumDark(scale int) *Comp {
	c.El.Style(styles.StrokePlumDark(scale))
	return c
}

func (c *Comp) StrokePlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokePlumDarkAlpha(scale))
	return c
}

func (c *Comp) StrokePurple(scale int) *Comp {
	c.El.Style(styles.StrokePurple(scale))
	return c
}

func (c *Comp) StrokePurpleAlpha(scale int) *Comp {
	c.El.Style(styles.StrokePurpleAlpha(scale))
	return c
}

func (c *Comp) StrokePurpleDark(scale int) *Comp {
	c.El.Style(styles.StrokePurpleDark(scale))
	return c
}

func (c *Comp) StrokePurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokePurpleDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeRed(scale int) *Comp {
	c.El.Style(styles.StrokeRed(scale))
	return c
}

func (c *Comp) StrokeRedAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeRedAlpha(scale))
	return c
}

func (c *Comp) StrokeRedDark(scale int) *Comp {
	c.El.Style(styles.StrokeRedDark(scale))
	return c
}

func (c *Comp) StrokeRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeRedDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeRuby(scale int) *Comp {
	c.El.Style(styles.StrokeRuby(scale))
	return c
}

func (c *Comp) StrokeRubyAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeRubyAlpha(scale))
	return c
}

func (c *Comp) StrokeRubyDark(scale int) *Comp {
	c.El.Style(styles.StrokeRubyDark(scale))
	return c
}

func (c *Comp) StrokeRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeRubyDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeSage(scale int) *Comp {
	c.El.Style(styles.StrokeSage(scale))
	return c
}

func (c *Comp) StrokeSageAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeSageAlpha(scale))
	return c
}

func (c *Comp) StrokeSageDark(scale int) *Comp {
	c.El.Style(styles.StrokeSageDark(scale))
	return c
}

func (c *Comp) StrokeSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeSageDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeSand(scale int) *Comp {
	c.El.Style(styles.StrokeSand(scale))
	return c
}

func (c *Comp) StrokeSandAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeSandAlpha(scale))
	return c
}

func (c *Comp) StrokeSandDark(scale int) *Comp {
	c.El.Style(styles.StrokeSandDark(scale))
	return c
}

func (c *Comp) StrokeSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeSandDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeSky(scale int) *Comp {
	c.El.Style(styles.StrokeSky(scale))
	return c
}

func (c *Comp) StrokeSkyAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeSkyAlpha(scale))
	return c
}

func (c *Comp) StrokeSkyDark(scale int) *Comp {
	c.El.Style(styles.StrokeSkyDark(scale))
	return c
}

func (c *Comp) StrokeSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeSkyDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeSlate(scale int) *Comp {
	c.El.Style(styles.StrokeSlate(scale))
	return c
}

func (c *Comp) StrokeSlateAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeSlateAlpha(scale))
	return c
}

func (c *Comp) StrokeSlateDark(scale int) *Comp {
	c.El.Style(styles.StrokeSlateDark(scale))
	return c
}

func (c *Comp) StrokeSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeSlateDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeTeal(scale int) *Comp {
	c.El.Style(styles.StrokeTeal(scale))
	return c
}

func (c *Comp) StrokeTealAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeTealAlpha(scale))
	return c
}

func (c *Comp) StrokeTealDark(scale int) *Comp {
	c.El.Style(styles.StrokeTealDark(scale))
	return c
}

func (c *Comp) StrokeTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeTealDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeTomato(scale int) *Comp {
	c.El.Style(styles.StrokeTomato(scale))
	return c
}

func (c *Comp) StrokeTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeTomatoAlpha(scale))
	return c
}

func (c *Comp) StrokeTomatoDark(scale int) *Comp {
	c.El.Style(styles.StrokeTomatoDark(scale))
	return c
}

func (c *Comp) StrokeTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeTransparent() *Comp {
	c.El.Style(styles.StrokeTransparent())
	return c
}

func (c *Comp) StrokeViolet(scale int) *Comp {
	c.El.Style(styles.StrokeViolet(scale))
	return c
}

func (c *Comp) StrokeVioletAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeVioletAlpha(scale))
	return c
}

func (c *Comp) StrokeVioletDark(scale int) *Comp {
	c.El.Style(styles.StrokeVioletDark(scale))
	return c
}

func (c *Comp) StrokeVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeVioletDarkAlpha(scale))
	return c
}

func (c *Comp) StrokeWhite() *Comp {
	c.El.Style(styles.StrokeWhite())
	return c
}

func (c *Comp) StrokeWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeWhiteAlpha(scale))
	return c
}

func (c *Comp) StrokeYellow(scale int) *Comp {
	c.El.Style(styles.StrokeYellow(scale))
	return c
}

func (c *Comp) StrokeYellowAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeYellowAlpha(scale))
	return c
}

func (c *Comp) StrokeYellowDark(scale int) *Comp {
	c.El.Style(styles.StrokeYellowDark(scale))
	return c
}

func (c *Comp) StrokeYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.StrokeYellowDarkAlpha(scale))
	return c
}

func (c *Comp) SubPixelAntialiazed() *Comp {
	c.El.Style(styles.SubPixelAntialiazed())
	return c
}

func (c *Comp) Table() *Comp {
	c.El.Style(styles.Table())
	return c
}

func (c *Comp) TableAuto() *Comp {
	c.El.Style(styles.TableAuto())
	return c
}

func (c *Comp) TableCaption() *Comp {
	c.El.Style(styles.TableCaption())
	return c
}

func (c *Comp) TableCell() *Comp {
	c.El.Style(styles.TableCell())
	return c
}

func (c *Comp) TableColumn() *Comp {
	c.El.Style(styles.TableColumn())
	return c
}

func (c *Comp) TableColumnGroup() *Comp {
	c.El.Style(styles.TableColumnGroup())
	return c
}

func (c *Comp) TableFixed() *Comp {
	c.El.Style(styles.TableFixed())
	return c
}

func (c *Comp) TableFooterGroup() *Comp {
	c.El.Style(styles.TableFooterGroup())
	return c
}

func (c *Comp) TableHeaderGroup() *Comp {
	c.El.Style(styles.TableHeaderGroup())
	return c
}

func (c *Comp) TableRow() *Comp {
	c.El.Style(styles.TableRow())
	return c
}

func (c *Comp) TableRowGroup() *Comp {
	c.El.Style(styles.TableRowGroup())
	return c
}

func (c *Comp) TabularNums() *Comp {
	c.El.Style(styles.TabularNums())
	return c
}

func (c *Comp) Text2xl(lineHeights ...any) *Comp {
	c.El.Style(styles.Text2xl(lineHeights...))
	return c
}

func (c *Comp) Text3xl(lineHeights ...any) *Comp {
	c.El.Style(styles.Text3xl(lineHeights...))
	return c
}

func (c *Comp) Text4xl(lineHeights ...any) *Comp {
	c.El.Style(styles.Text4xl(lineHeights...))
	return c
}

func (c *Comp) Text5xl(lineHeights ...any) *Comp {
	c.El.Style(styles.Text5xl(lineHeights...))
	return c
}

func (c *Comp) Text6xl(lineHeights ...any) *Comp {
	c.El.Style(styles.Text6xl(lineHeights...))
	return c
}

func (c *Comp) Text7xl(lineHeights ...any) *Comp {
	c.El.Style(styles.Text7xl(lineHeights...))
	return c
}

func (c *Comp) Text8xl(lineHeights ...any) *Comp {
	c.El.Style(styles.Text8xl(lineHeights...))
	return c
}

func (c *Comp) Text9xl(lineHeights ...any) *Comp {
	c.El.Style(styles.Text9xl(lineHeights...))
	return c
}

func (c *Comp) TextAmber(scale int) *Comp {
	c.El.Style(styles.TextAmber(scale))
	return c
}

func (c *Comp) TextAmberAlpha(scale int) *Comp {
	c.El.Style(styles.TextAmberAlpha(scale))
	return c
}

func (c *Comp) TextAmberDark(scale int) *Comp {
	c.El.Style(styles.TextAmberDark(scale))
	return c
}

func (c *Comp) TextAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextAmberDarkAlpha(scale))
	return c
}

func (c *Comp) TextBalance() *Comp {
	c.El.Style(styles.TextBalance())
	return c
}

func (c *Comp) TextBase(lineHeights ...any) *Comp {
	c.El.Style(styles.TextBase(lineHeights...))
	return c
}

func (c *Comp) TextBlack() *Comp {
	c.El.Style(styles.TextBlack())
	return c
}

func (c *Comp) TextBlackAlpha(scale int) *Comp {
	c.El.Style(styles.TextBlackAlpha(scale))
	return c
}

func (c *Comp) TextBlue(scale int) *Comp {
	c.El.Style(styles.TextBlue(scale))
	return c
}

func (c *Comp) TextBlueAlpha(scale int) *Comp {
	c.El.Style(styles.TextBlueAlpha(scale))
	return c
}

func (c *Comp) TextBlueDark(scale int) *Comp {
	c.El.Style(styles.TextBlueDark(scale))
	return c
}

func (c *Comp) TextBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextBlueDarkAlpha(scale))
	return c
}

func (c *Comp) TextBronze(scale int) *Comp {
	c.El.Style(styles.TextBronze(scale))
	return c
}

func (c *Comp) TextBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.TextBronzeAlpha(scale))
	return c
}

func (c *Comp) TextBronzeDark(scale int) *Comp {
	c.El.Style(styles.TextBronzeDark(scale))
	return c
}

func (c *Comp) TextBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) TextBrown(scale int) *Comp {
	c.El.Style(styles.TextBrown(scale))
	return c
}

func (c *Comp) TextBrownAlpha(scale int) *Comp {
	c.El.Style(styles.TextBrownAlpha(scale))
	return c
}

func (c *Comp) TextBrownDark(scale int) *Comp {
	c.El.Style(styles.TextBrownDark(scale))
	return c
}

func (c *Comp) TextBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextBrownDarkAlpha(scale))
	return c
}

func (c *Comp) TextCenter() *Comp {
	c.El.Style(styles.TextCenter())
	return c
}

func (c *Comp) TextClip() *Comp {
	c.El.Style(styles.TextClip())
	return c
}

func (c *Comp) TextColor(val value.Value) *Comp {
	c.El.Style(styles.TextColor(val))
	return c
}

func (c *Comp) TextCrimson(scale int) *Comp {
	c.El.Style(styles.TextCrimson(scale))
	return c
}

func (c *Comp) TextCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.TextCrimsonAlpha(scale))
	return c
}

func (c *Comp) TextCrimsonDark(scale int) *Comp {
	c.El.Style(styles.TextCrimsonDark(scale))
	return c
}

func (c *Comp) TextCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) TextCurrent() *Comp {
	c.El.Style(styles.TextCurrent())
	return c
}

func (c *Comp) TextCyan(scale int) *Comp {
	c.El.Style(styles.TextCyan(scale))
	return c
}

func (c *Comp) TextCyanAlpha(scale int) *Comp {
	c.El.Style(styles.TextCyanAlpha(scale))
	return c
}

func (c *Comp) TextCyanDark(scale int) *Comp {
	c.El.Style(styles.TextCyanDark(scale))
	return c
}

func (c *Comp) TextCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextCyanDarkAlpha(scale))
	return c
}

func (c *Comp) TextEllipsis() *Comp {
	c.El.Style(styles.TextEllipsis())
	return c
}

func (c *Comp) TextEnd() *Comp {
	c.El.Style(styles.TextEnd())
	return c
}

func (c *Comp) TextGold(scale int) *Comp {
	c.El.Style(styles.TextGold(scale))
	return c
}

func (c *Comp) TextGoldAlpha(scale int) *Comp {
	c.El.Style(styles.TextGoldAlpha(scale))
	return c
}

func (c *Comp) TextGoldDark(scale int) *Comp {
	c.El.Style(styles.TextGoldDark(scale))
	return c
}

func (c *Comp) TextGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextGoldDarkAlpha(scale))
	return c
}

func (c *Comp) TextGrass(scale int) *Comp {
	c.El.Style(styles.TextGrass(scale))
	return c
}

func (c *Comp) TextGrassAlpha(scale int) *Comp {
	c.El.Style(styles.TextGrassAlpha(scale))
	return c
}

func (c *Comp) TextGrassDark(scale int) *Comp {
	c.El.Style(styles.TextGrassDark(scale))
	return c
}

func (c *Comp) TextGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextGrassDarkAlpha(scale))
	return c
}

func (c *Comp) TextGray(scale int) *Comp {
	c.El.Style(styles.TextGray(scale))
	return c
}

func (c *Comp) TextGrayAlpha(scale int) *Comp {
	c.El.Style(styles.TextGrayAlpha(scale))
	return c
}

func (c *Comp) TextGrayDark(scale int) *Comp {
	c.El.Style(styles.TextGrayDark(scale))
	return c
}

func (c *Comp) TextGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextGrayDarkAlpha(scale))
	return c
}

func (c *Comp) TextGreen(scale int) *Comp {
	c.El.Style(styles.TextGreen(scale))
	return c
}

func (c *Comp) TextGreenAlpha(scale int) *Comp {
	c.El.Style(styles.TextGreenAlpha(scale))
	return c
}

func (c *Comp) TextGreenDark(scale int) *Comp {
	c.El.Style(styles.TextGreenDark(scale))
	return c
}

func (c *Comp) TextGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextGreenDarkAlpha(scale))
	return c
}

func (c *Comp) TextIndigo(scale int) *Comp {
	c.El.Style(styles.TextIndigo(scale))
	return c
}

func (c *Comp) TextIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.TextIndigoAlpha(scale))
	return c
}

func (c *Comp) TextIndigoDark(scale int) *Comp {
	c.El.Style(styles.TextIndigoDark(scale))
	return c
}

func (c *Comp) TextIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) TextInherit() *Comp {
	c.El.Style(styles.TextInherit())
	return c
}

func (c *Comp) TextIris(scale int) *Comp {
	c.El.Style(styles.TextIris(scale))
	return c
}

func (c *Comp) TextIrisAlpha(scale int) *Comp {
	c.El.Style(styles.TextIrisAlpha(scale))
	return c
}

func (c *Comp) TextIrisDark(scale int) *Comp {
	c.El.Style(styles.TextIrisDark(scale))
	return c
}

func (c *Comp) TextIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextIrisDarkAlpha(scale))
	return c
}

func (c *Comp) TextJade(scale int) *Comp {
	c.El.Style(styles.TextJade(scale))
	return c
}

func (c *Comp) TextJadeAlpha(scale int) *Comp {
	c.El.Style(styles.TextJadeAlpha(scale))
	return c
}

func (c *Comp) TextJadeDark(scale int) *Comp {
	c.El.Style(styles.TextJadeDark(scale))
	return c
}

func (c *Comp) TextJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextJadeDarkAlpha(scale))
	return c
}

func (c *Comp) TextJustify() *Comp {
	c.El.Style(styles.TextJustify())
	return c
}

func (c *Comp) TextLeft() *Comp {
	c.El.Style(styles.TextLeft())
	return c
}

func (c *Comp) TextLg(lineHeights ...any) *Comp {
	c.El.Style(styles.TextLg(lineHeights...))
	return c
}

func (c *Comp) TextLime(scale int) *Comp {
	c.El.Style(styles.TextLime(scale))
	return c
}

func (c *Comp) TextLimeAlpha(scale int) *Comp {
	c.El.Style(styles.TextLimeAlpha(scale))
	return c
}

func (c *Comp) TextLimeDark(scale int) *Comp {
	c.El.Style(styles.TextLimeDark(scale))
	return c
}

func (c *Comp) TextLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextLimeDarkAlpha(scale))
	return c
}

func (c *Comp) TextMauve(scale int) *Comp {
	c.El.Style(styles.TextMauve(scale))
	return c
}

func (c *Comp) TextMauveAlpha(scale int) *Comp {
	c.El.Style(styles.TextMauveAlpha(scale))
	return c
}

func (c *Comp) TextMauveDark(scale int) *Comp {
	c.El.Style(styles.TextMauveDark(scale))
	return c
}

func (c *Comp) TextMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextMauveDarkAlpha(scale))
	return c
}

func (c *Comp) TextMint(scale int) *Comp {
	c.El.Style(styles.TextMint(scale))
	return c
}

func (c *Comp) TextMintAlpha(scale int) *Comp {
	c.El.Style(styles.TextMintAlpha(scale))
	return c
}

func (c *Comp) TextMintDark(scale int) *Comp {
	c.El.Style(styles.TextMintDark(scale))
	return c
}

func (c *Comp) TextMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextMintDarkAlpha(scale))
	return c
}

func (c *Comp) TextNoWrap() *Comp {
	c.El.Style(styles.TextNoWrap())
	return c
}

func (c *Comp) TextOlive(scale int) *Comp {
	c.El.Style(styles.TextOlive(scale))
	return c
}

func (c *Comp) TextOliveAlpha(scale int) *Comp {
	c.El.Style(styles.TextOliveAlpha(scale))
	return c
}

func (c *Comp) TextOliveDark(scale int) *Comp {
	c.El.Style(styles.TextOliveDark(scale))
	return c
}

func (c *Comp) TextOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextOliveDarkAlpha(scale))
	return c
}

func (c *Comp) TextOrange(scale int) *Comp {
	c.El.Style(styles.TextOrange(scale))
	return c
}

func (c *Comp) TextOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.TextOrangeAlpha(scale))
	return c
}

func (c *Comp) TextOrangeDark(scale int) *Comp {
	c.El.Style(styles.TextOrangeDark(scale))
	return c
}

func (c *Comp) TextOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) TextPink(scale int) *Comp {
	c.El.Style(styles.TextPink(scale))
	return c
}

func (c *Comp) TextPinkAlpha(scale int) *Comp {
	c.El.Style(styles.TextPinkAlpha(scale))
	return c
}

func (c *Comp) TextPinkDark(scale int) *Comp {
	c.El.Style(styles.TextPinkDark(scale))
	return c
}

func (c *Comp) TextPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextPinkDarkAlpha(scale))
	return c
}

func (c *Comp) TextPlum(scale int) *Comp {
	c.El.Style(styles.TextPlum(scale))
	return c
}

func (c *Comp) TextPlumAlpha(scale int) *Comp {
	c.El.Style(styles.TextPlumAlpha(scale))
	return c
}

func (c *Comp) TextPlumDark(scale int) *Comp {
	c.El.Style(styles.TextPlumDark(scale))
	return c
}

func (c *Comp) TextPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextPlumDarkAlpha(scale))
	return c
}

func (c *Comp) TextPretty() *Comp {
	c.El.Style(styles.TextPretty())
	return c
}

func (c *Comp) TextPurple(scale int) *Comp {
	c.El.Style(styles.TextPurple(scale))
	return c
}

func (c *Comp) TextPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.TextPurpleAlpha(scale))
	return c
}

func (c *Comp) TextPurpleDark(scale int) *Comp {
	c.El.Style(styles.TextPurpleDark(scale))
	return c
}

func (c *Comp) TextPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) TextRed(scale int) *Comp {
	c.El.Style(styles.TextRed(scale))
	return c
}

func (c *Comp) TextRedAlpha(scale int) *Comp {
	c.El.Style(styles.TextRedAlpha(scale))
	return c
}

func (c *Comp) TextRedDark(scale int) *Comp {
	c.El.Style(styles.TextRedDark(scale))
	return c
}

func (c *Comp) TextRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextRedDarkAlpha(scale))
	return c
}

func (c *Comp) TextRight() *Comp {
	c.El.Style(styles.TextRight())
	return c
}

func (c *Comp) TextRuby(scale int) *Comp {
	c.El.Style(styles.TextRuby(scale))
	return c
}

func (c *Comp) TextRubyAlpha(scale int) *Comp {
	c.El.Style(styles.TextRubyAlpha(scale))
	return c
}

func (c *Comp) TextRubyDark(scale int) *Comp {
	c.El.Style(styles.TextRubyDark(scale))
	return c
}

func (c *Comp) TextRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextRubyDarkAlpha(scale))
	return c
}

func (c *Comp) TextSage(scale int) *Comp {
	c.El.Style(styles.TextSage(scale))
	return c
}

func (c *Comp) TextSageAlpha(scale int) *Comp {
	c.El.Style(styles.TextSageAlpha(scale))
	return c
}

func (c *Comp) TextSageDark(scale int) *Comp {
	c.El.Style(styles.TextSageDark(scale))
	return c
}

func (c *Comp) TextSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextSageDarkAlpha(scale))
	return c
}

func (c *Comp) TextSand(scale int) *Comp {
	c.El.Style(styles.TextSand(scale))
	return c
}

func (c *Comp) TextSandAlpha(scale int) *Comp {
	c.El.Style(styles.TextSandAlpha(scale))
	return c
}

func (c *Comp) TextSandDark(scale int) *Comp {
	c.El.Style(styles.TextSandDark(scale))
	return c
}

func (c *Comp) TextSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextSandDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadow(val value.Value) *Comp {
	c.El.Style(styles.TextShadow(val))
	return c
}

func (c *Comp) TextShadow2xl() *Comp {
	c.El.Style(styles.TextShadow2xl())
	return c
}

func (c *Comp) TextShadow2xs() *Comp {
	c.El.Style(styles.TextShadow2xs())
	return c
}

func (c *Comp) TextShadowAmber(scale int) *Comp {
	c.El.Style(styles.TextShadowAmber(scale))
	return c
}

func (c *Comp) TextShadowAmberAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowAmberAlpha(scale))
	return c
}

func (c *Comp) TextShadowAmberDark(scale int) *Comp {
	c.El.Style(styles.TextShadowAmberDark(scale))
	return c
}

func (c *Comp) TextShadowAmberDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowAmberDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowBlack() *Comp {
	c.El.Style(styles.TextShadowBlack())
	return c
}

func (c *Comp) TextShadowBlackAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowBlackAlpha(scale))
	return c
}

func (c *Comp) TextShadowBlue(scale int) *Comp {
	c.El.Style(styles.TextShadowBlue(scale))
	return c
}

func (c *Comp) TextShadowBlueAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowBlueAlpha(scale))
	return c
}

func (c *Comp) TextShadowBlueDark(scale int) *Comp {
	c.El.Style(styles.TextShadowBlueDark(scale))
	return c
}

func (c *Comp) TextShadowBlueDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowBlueDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowBronze(scale int) *Comp {
	c.El.Style(styles.TextShadowBronze(scale))
	return c
}

func (c *Comp) TextShadowBronzeAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowBronzeAlpha(scale))
	return c
}

func (c *Comp) TextShadowBronzeDark(scale int) *Comp {
	c.El.Style(styles.TextShadowBronzeDark(scale))
	return c
}

func (c *Comp) TextShadowBronzeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowBronzeDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowBrown(scale int) *Comp {
	c.El.Style(styles.TextShadowBrown(scale))
	return c
}

func (c *Comp) TextShadowBrownAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowBrownAlpha(scale))
	return c
}

func (c *Comp) TextShadowBrownDark(scale int) *Comp {
	c.El.Style(styles.TextShadowBrownDark(scale))
	return c
}

func (c *Comp) TextShadowBrownDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowBrownDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowColor(val value.Value) *Comp {
	c.El.Style(styles.TextShadowColor(val))
	return c
}

func (c *Comp) TextShadowCrimson(scale int) *Comp {
	c.El.Style(styles.TextShadowCrimson(scale))
	return c
}

func (c *Comp) TextShadowCrimsonAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowCrimsonAlpha(scale))
	return c
}

func (c *Comp) TextShadowCrimsonDark(scale int) *Comp {
	c.El.Style(styles.TextShadowCrimsonDark(scale))
	return c
}

func (c *Comp) TextShadowCrimsonDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowCrimsonDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowCurrent() *Comp {
	c.El.Style(styles.TextShadowCurrent())
	return c
}

func (c *Comp) TextShadowCyan(scale int) *Comp {
	c.El.Style(styles.TextShadowCyan(scale))
	return c
}

func (c *Comp) TextShadowCyanAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowCyanAlpha(scale))
	return c
}

func (c *Comp) TextShadowCyanDark(scale int) *Comp {
	c.El.Style(styles.TextShadowCyanDark(scale))
	return c
}

func (c *Comp) TextShadowCyanDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowCyanDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowGold(scale int) *Comp {
	c.El.Style(styles.TextShadowGold(scale))
	return c
}

func (c *Comp) TextShadowGoldAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowGoldAlpha(scale))
	return c
}

func (c *Comp) TextShadowGoldDark(scale int) *Comp {
	c.El.Style(styles.TextShadowGoldDark(scale))
	return c
}

func (c *Comp) TextShadowGoldDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowGoldDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowGrass(scale int) *Comp {
	c.El.Style(styles.TextShadowGrass(scale))
	return c
}

func (c *Comp) TextShadowGrassAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowGrassAlpha(scale))
	return c
}

func (c *Comp) TextShadowGrassDark(scale int) *Comp {
	c.El.Style(styles.TextShadowGrassDark(scale))
	return c
}

func (c *Comp) TextShadowGrassDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowGrassDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowGray(scale int) *Comp {
	c.El.Style(styles.TextShadowGray(scale))
	return c
}

func (c *Comp) TextShadowGrayAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowGrayAlpha(scale))
	return c
}

func (c *Comp) TextShadowGrayDark(scale int) *Comp {
	c.El.Style(styles.TextShadowGrayDark(scale))
	return c
}

func (c *Comp) TextShadowGrayDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowGrayDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowGreen(scale int) *Comp {
	c.El.Style(styles.TextShadowGreen(scale))
	return c
}

func (c *Comp) TextShadowGreenAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowGreenAlpha(scale))
	return c
}

func (c *Comp) TextShadowGreenDark(scale int) *Comp {
	c.El.Style(styles.TextShadowGreenDark(scale))
	return c
}

func (c *Comp) TextShadowGreenDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowGreenDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowIndigo(scale int) *Comp {
	c.El.Style(styles.TextShadowIndigo(scale))
	return c
}

func (c *Comp) TextShadowIndigoAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowIndigoAlpha(scale))
	return c
}

func (c *Comp) TextShadowIndigoDark(scale int) *Comp {
	c.El.Style(styles.TextShadowIndigoDark(scale))
	return c
}

func (c *Comp) TextShadowIndigoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowIndigoDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowInherit() *Comp {
	c.El.Style(styles.TextShadowInherit())
	return c
}

func (c *Comp) TextShadowIris(scale int) *Comp {
	c.El.Style(styles.TextShadowIris(scale))
	return c
}

func (c *Comp) TextShadowIrisAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowIrisAlpha(scale))
	return c
}

func (c *Comp) TextShadowIrisDark(scale int) *Comp {
	c.El.Style(styles.TextShadowIrisDark(scale))
	return c
}

func (c *Comp) TextShadowIrisDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowIrisDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowJade(scale int) *Comp {
	c.El.Style(styles.TextShadowJade(scale))
	return c
}

func (c *Comp) TextShadowJadeAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowJadeAlpha(scale))
	return c
}

func (c *Comp) TextShadowJadeDark(scale int) *Comp {
	c.El.Style(styles.TextShadowJadeDark(scale))
	return c
}

func (c *Comp) TextShadowJadeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowJadeDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowLg() *Comp {
	c.El.Style(styles.TextShadowLg())
	return c
}

func (c *Comp) TextShadowLime(scale int) *Comp {
	c.El.Style(styles.TextShadowLime(scale))
	return c
}

func (c *Comp) TextShadowLimeAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowLimeAlpha(scale))
	return c
}

func (c *Comp) TextShadowLimeDark(scale int) *Comp {
	c.El.Style(styles.TextShadowLimeDark(scale))
	return c
}

func (c *Comp) TextShadowLimeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowLimeDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowMauve(scale int) *Comp {
	c.El.Style(styles.TextShadowMauve(scale))
	return c
}

func (c *Comp) TextShadowMauveAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowMauveAlpha(scale))
	return c
}

func (c *Comp) TextShadowMauveDark(scale int) *Comp {
	c.El.Style(styles.TextShadowMauveDark(scale))
	return c
}

func (c *Comp) TextShadowMauveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowMauveDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowMd() *Comp {
	c.El.Style(styles.TextShadowMd())
	return c
}

func (c *Comp) TextShadowMint(scale int) *Comp {
	c.El.Style(styles.TextShadowMint(scale))
	return c
}

func (c *Comp) TextShadowMintAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowMintAlpha(scale))
	return c
}

func (c *Comp) TextShadowMintDark(scale int) *Comp {
	c.El.Style(styles.TextShadowMintDark(scale))
	return c
}

func (c *Comp) TextShadowMintDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowMintDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowNone() *Comp {
	c.El.Style(styles.TextShadowNone())
	return c
}

func (c *Comp) TextShadowOlive(scale int) *Comp {
	c.El.Style(styles.TextShadowOlive(scale))
	return c
}

func (c *Comp) TextShadowOliveAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowOliveAlpha(scale))
	return c
}

func (c *Comp) TextShadowOliveDark(scale int) *Comp {
	c.El.Style(styles.TextShadowOliveDark(scale))
	return c
}

func (c *Comp) TextShadowOliveDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowOliveDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowOrange(scale int) *Comp {
	c.El.Style(styles.TextShadowOrange(scale))
	return c
}

func (c *Comp) TextShadowOrangeAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowOrangeAlpha(scale))
	return c
}

func (c *Comp) TextShadowOrangeDark(scale int) *Comp {
	c.El.Style(styles.TextShadowOrangeDark(scale))
	return c
}

func (c *Comp) TextShadowOrangeDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowOrangeDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowPink(scale int) *Comp {
	c.El.Style(styles.TextShadowPink(scale))
	return c
}

func (c *Comp) TextShadowPinkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowPinkAlpha(scale))
	return c
}

func (c *Comp) TextShadowPinkDark(scale int) *Comp {
	c.El.Style(styles.TextShadowPinkDark(scale))
	return c
}

func (c *Comp) TextShadowPinkDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowPinkDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowPlum(scale int) *Comp {
	c.El.Style(styles.TextShadowPlum(scale))
	return c
}

func (c *Comp) TextShadowPlumAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowPlumAlpha(scale))
	return c
}

func (c *Comp) TextShadowPlumDark(scale int) *Comp {
	c.El.Style(styles.TextShadowPlumDark(scale))
	return c
}

func (c *Comp) TextShadowPlumDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowPlumDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowPurple(scale int) *Comp {
	c.El.Style(styles.TextShadowPurple(scale))
	return c
}

func (c *Comp) TextShadowPurpleAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowPurpleAlpha(scale))
	return c
}

func (c *Comp) TextShadowPurpleDark(scale int) *Comp {
	c.El.Style(styles.TextShadowPurpleDark(scale))
	return c
}

func (c *Comp) TextShadowPurpleDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowPurpleDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowRed(scale int) *Comp {
	c.El.Style(styles.TextShadowRed(scale))
	return c
}

func (c *Comp) TextShadowRedAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowRedAlpha(scale))
	return c
}

func (c *Comp) TextShadowRedDark(scale int) *Comp {
	c.El.Style(styles.TextShadowRedDark(scale))
	return c
}

func (c *Comp) TextShadowRedDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowRedDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowRuby(scale int) *Comp {
	c.El.Style(styles.TextShadowRuby(scale))
	return c
}

func (c *Comp) TextShadowRubyAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowRubyAlpha(scale))
	return c
}

func (c *Comp) TextShadowRubyDark(scale int) *Comp {
	c.El.Style(styles.TextShadowRubyDark(scale))
	return c
}

func (c *Comp) TextShadowRubyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowRubyDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowSage(scale int) *Comp {
	c.El.Style(styles.TextShadowSage(scale))
	return c
}

func (c *Comp) TextShadowSageAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowSageAlpha(scale))
	return c
}

func (c *Comp) TextShadowSageDark(scale int) *Comp {
	c.El.Style(styles.TextShadowSageDark(scale))
	return c
}

func (c *Comp) TextShadowSageDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowSageDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowSand(scale int) *Comp {
	c.El.Style(styles.TextShadowSand(scale))
	return c
}

func (c *Comp) TextShadowSandAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowSandAlpha(scale))
	return c
}

func (c *Comp) TextShadowSandDark(scale int) *Comp {
	c.El.Style(styles.TextShadowSandDark(scale))
	return c
}

func (c *Comp) TextShadowSandDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowSandDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowSky(scale int) *Comp {
	c.El.Style(styles.TextShadowSky(scale))
	return c
}

func (c *Comp) TextShadowSkyAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowSkyAlpha(scale))
	return c
}

func (c *Comp) TextShadowSkyDark(scale int) *Comp {
	c.El.Style(styles.TextShadowSkyDark(scale))
	return c
}

func (c *Comp) TextShadowSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowSkyDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowSlate(scale int) *Comp {
	c.El.Style(styles.TextShadowSlate(scale))
	return c
}

func (c *Comp) TextShadowSlateAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowSlateAlpha(scale))
	return c
}

func (c *Comp) TextShadowSlateDark(scale int) *Comp {
	c.El.Style(styles.TextShadowSlateDark(scale))
	return c
}

func (c *Comp) TextShadowSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowSlateDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowSm() *Comp {
	c.El.Style(styles.TextShadowSm())
	return c
}

func (c *Comp) TextShadowTeal(scale int) *Comp {
	c.El.Style(styles.TextShadowTeal(scale))
	return c
}

func (c *Comp) TextShadowTealAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowTealAlpha(scale))
	return c
}

func (c *Comp) TextShadowTealDark(scale int) *Comp {
	c.El.Style(styles.TextShadowTealDark(scale))
	return c
}

func (c *Comp) TextShadowTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowTealDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowTomato(scale int) *Comp {
	c.El.Style(styles.TextShadowTomato(scale))
	return c
}

func (c *Comp) TextShadowTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowTomatoAlpha(scale))
	return c
}

func (c *Comp) TextShadowTomatoDark(scale int) *Comp {
	c.El.Style(styles.TextShadowTomatoDark(scale))
	return c
}

func (c *Comp) TextShadowTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowTransparent() *Comp {
	c.El.Style(styles.TextShadowTransparent())
	return c
}

func (c *Comp) TextShadowViolet(scale int) *Comp {
	c.El.Style(styles.TextShadowViolet(scale))
	return c
}

func (c *Comp) TextShadowVioletAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowVioletAlpha(scale))
	return c
}

func (c *Comp) TextShadowVioletDark(scale int) *Comp {
	c.El.Style(styles.TextShadowVioletDark(scale))
	return c
}

func (c *Comp) TextShadowVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowVioletDarkAlpha(scale))
	return c
}

func (c *Comp) TextShadowWhite() *Comp {
	c.El.Style(styles.TextShadowWhite())
	return c
}

func (c *Comp) TextShadowWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowWhiteAlpha(scale))
	return c
}

func (c *Comp) TextShadowXl() *Comp {
	c.El.Style(styles.TextShadowXl())
	return c
}

func (c *Comp) TextShadowXs() *Comp {
	c.El.Style(styles.TextShadowXs())
	return c
}

func (c *Comp) TextShadowYellow(scale int) *Comp {
	c.El.Style(styles.TextShadowYellow(scale))
	return c
}

func (c *Comp) TextShadowYellowAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowYellowAlpha(scale))
	return c
}

func (c *Comp) TextShadowYellowDark(scale int) *Comp {
	c.El.Style(styles.TextShadowYellowDark(scale))
	return c
}

func (c *Comp) TextShadowYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextShadowYellowDarkAlpha(scale))
	return c
}

func (c *Comp) TextSizeBy(val value.Value) *Comp {
	c.El.Style(styles.TextSizeBy(val))
	return c
}

func (c *Comp) TextSky(scale int) *Comp {
	c.El.Style(styles.TextSky(scale))
	return c
}

func (c *Comp) TextSkyAlpha(scale int) *Comp {
	c.El.Style(styles.TextSkyAlpha(scale))
	return c
}

func (c *Comp) TextSkyDark(scale int) *Comp {
	c.El.Style(styles.TextSkyDark(scale))
	return c
}

func (c *Comp) TextSkyDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextSkyDarkAlpha(scale))
	return c
}

func (c *Comp) TextSlate(scale int) *Comp {
	c.El.Style(styles.TextSlate(scale))
	return c
}

func (c *Comp) TextSlateAlpha(scale int) *Comp {
	c.El.Style(styles.TextSlateAlpha(scale))
	return c
}

func (c *Comp) TextSlateDark(scale int) *Comp {
	c.El.Style(styles.TextSlateDark(scale))
	return c
}

func (c *Comp) TextSlateDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextSlateDarkAlpha(scale))
	return c
}

func (c *Comp) TextSm(lineHeights ...any) *Comp {
	c.El.Style(styles.TextSm(lineHeights...))
	return c
}

func (c *Comp) TextStart() *Comp {
	c.El.Style(styles.TextStart())
	return c
}

func (c *Comp) TextTeal(scale int) *Comp {
	c.El.Style(styles.TextTeal(scale))
	return c
}

func (c *Comp) TextTealAlpha(scale int) *Comp {
	c.El.Style(styles.TextTealAlpha(scale))
	return c
}

func (c *Comp) TextTealDark(scale int) *Comp {
	c.El.Style(styles.TextTealDark(scale))
	return c
}

func (c *Comp) TextTealDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextTealDarkAlpha(scale))
	return c
}

func (c *Comp) TextTomato(scale int) *Comp {
	c.El.Style(styles.TextTomato(scale))
	return c
}

func (c *Comp) TextTomatoAlpha(scale int) *Comp {
	c.El.Style(styles.TextTomatoAlpha(scale))
	return c
}

func (c *Comp) TextTomatoDark(scale int) *Comp {
	c.El.Style(styles.TextTomatoDark(scale))
	return c
}

func (c *Comp) TextTomatoDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextTomatoDarkAlpha(scale))
	return c
}

func (c *Comp) TextTransparent() *Comp {
	c.El.Style(styles.TextTransparent())
	return c
}

func (c *Comp) TextViolet(scale int) *Comp {
	c.El.Style(styles.TextViolet(scale))
	return c
}

func (c *Comp) TextVioletAlpha(scale int) *Comp {
	c.El.Style(styles.TextVioletAlpha(scale))
	return c
}

func (c *Comp) TextVioletDark(scale int) *Comp {
	c.El.Style(styles.TextVioletDark(scale))
	return c
}

func (c *Comp) TextVioletDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextVioletDarkAlpha(scale))
	return c
}

func (c *Comp) TextWhite() *Comp {
	c.El.Style(styles.TextWhite())
	return c
}

func (c *Comp) TextWhiteAlpha(scale int) *Comp {
	c.El.Style(styles.TextWhiteAlpha(scale))
	return c
}

func (c *Comp) TextWrap() *Comp {
	c.El.Style(styles.TextWrap())
	return c
}

func (c *Comp) TextXl(lineHeights ...any) *Comp {
	c.El.Style(styles.TextXl(lineHeights...))
	return c
}

func (c *Comp) TextXs(lineHeights ...any) *Comp {
	c.El.Style(styles.TextXs(lineHeights...))
	return c
}

func (c *Comp) TextYellow(scale int) *Comp {
	c.El.Style(styles.TextYellow(scale))
	return c
}

func (c *Comp) TextYellowAlpha(scale int) *Comp {
	c.El.Style(styles.TextYellowAlpha(scale))
	return c
}

func (c *Comp) TextYellowDark(scale int) *Comp {
	c.El.Style(styles.TextYellowDark(scale))
	return c
}

func (c *Comp) TextYellowDarkAlpha(scale int) *Comp {
	c.El.Style(styles.TextYellowDarkAlpha(scale))
	return c
}

func (c *Comp) To(val value.Value) *Comp {
	c.El.Style(styles.To(val))
	return c
}

func (c *Comp) Top(number int) *Comp {
	c.El.Style(styles.Top(number))
	return c
}

func (c *Comp) TopAuto() *Comp {
	c.El.Style(styles.TopAuto())
	return c
}

func (c *Comp) TopBy(val value.Value) *Comp {
	c.El.Style(styles.TopBy(val))
	return c
}

func (c *Comp) TopFraction(fraction float64) *Comp {
	c.El.Style(styles.TopFraction(fraction))
	return c
}

func (c *Comp) TopFull() *Comp {
	c.El.Style(styles.TopFull())
	return c
}

func (c *Comp) TopPx() *Comp {
	c.El.Style(styles.TopPx())
	return c
}

func (c *Comp) TouchAuto() *Comp {
	c.El.Style(styles.TouchAuto())
	return c
}

func (c *Comp) TouchManipulation() *Comp {
	c.El.Style(styles.TouchManipulation())
	return c
}

func (c *Comp) TouchNone() *Comp {
	c.El.Style(styles.TouchNone())
	return c
}

func (c *Comp) TouchPanDown() *Comp {
	c.El.Style(styles.TouchPanDown())
	return c
}

func (c *Comp) TouchPanLeft() *Comp {
	c.El.Style(styles.TouchPanLeft())
	return c
}

func (c *Comp) TouchPanRight() *Comp {
	c.El.Style(styles.TouchPanRight())
	return c
}

func (c *Comp) TouchPanUp() *Comp {
	c.El.Style(styles.TouchPanUp())
	return c
}

func (c *Comp) TouchPanX() *Comp {
	c.El.Style(styles.TouchPanX())
	return c
}

func (c *Comp) TouchPanY() *Comp {
	c.El.Style(styles.TouchPanY())
	return c
}

func (c *Comp) TouchPinchZoom() *Comp {
	c.El.Style(styles.TouchPinchZoom())
	return c
}

func (c *Comp) TrackingBy(val value.Value) *Comp {
	c.El.Style(styles.TrackingBy(val))
	return c
}

func (c *Comp) TrackingNormal() *Comp {
	c.El.Style(styles.TrackingNormal())
	return c
}

func (c *Comp) TrackingTight() *Comp {
	c.El.Style(styles.TrackingTight())
	return c
}

func (c *Comp) TrackingTighter() *Comp {
	c.El.Style(styles.TrackingTighter())
	return c
}

func (c *Comp) TrackingWide() *Comp {
	c.El.Style(styles.TrackingWide())
	return c
}

func (c *Comp) TrackingWider() *Comp {
	c.El.Style(styles.TrackingWider())
	return c
}

func (c *Comp) TrackingWidest() *Comp {
	c.El.Style(styles.TrackingWidest())
	return c
}

func (c *Comp) Transform(val value.Value) *Comp {
	c.El.Style(styles.Transform(val))
	return c
}

func (c *Comp) Transform3d() *Comp {
	c.El.Style(styles.Transform3d())
	return c
}

func (c *Comp) TransformCpu() *Comp {
	c.El.Style(styles.TransformCpu())
	return c
}

func (c *Comp) TransformFkat() *Comp {
	c.El.Style(styles.TransformFkat())
	return c
}

func (c *Comp) TransformGpu() *Comp {
	c.El.Style(styles.TransformGpu())
	return c
}

func (c *Comp) TransformNone() *Comp {
	c.El.Style(styles.TransformNone())
	return c
}

func (c *Comp) Transition(val ...value.Value) *Comp {
	c.El.Style(styles.Transition(val...))
	return c
}

func (c *Comp) TransitionAll() *Comp {
	c.El.Style(styles.TransitionAll())
	return c
}

func (c *Comp) TransitionColors() *Comp {
	c.El.Style(styles.TransitionColors())
	return c
}

func (c *Comp) TransitionDiscrete() *Comp {
	c.El.Style(styles.TransitionDiscrete())
	return c
}

func (c *Comp) TransitionNone() *Comp {
	c.El.Style(styles.TransitionNone())
	return c
}

func (c *Comp) TransitionNormal() *Comp {
	c.El.Style(styles.TransitionNormal())
	return c
}

func (c *Comp) TransitionOpacity() *Comp {
	c.El.Style(styles.TransitionOpacity())
	return c
}

func (c *Comp) TransitionShadow() *Comp {
	c.El.Style(styles.TransitionShadow())
	return c
}

func (c *Comp) TransitionTransform() *Comp {
	c.El.Style(styles.TransitionTransform())
	return c
}

func (c *Comp) Translate(val any) *Comp {
	c.El.Style(styles.Translate(val))
	return c
}

func (c *Comp) TranslateFull() *Comp {
	c.El.Style(styles.TranslateFull())
	return c
}

func (c *Comp) TranslateNone() *Comp {
	c.El.Style(styles.TranslateNone())
	return c
}

func (c *Comp) TranslatePx() *Comp {
	c.El.Style(styles.TranslatePx())
	return c
}

func (c *Comp) TranslateX(val any) *Comp {
	c.El.Style(styles.TranslateX(val))
	return c
}

func (c *Comp) TranslateXFull() *Comp {
	c.El.Style(styles.TranslateXFull())
	return c
}

func (c *Comp) TranslateXPx() *Comp {
	c.El.Style(styles.TranslateXPx())
	return c
}

func (c *Comp) TranslateY(val any) *Comp {
	c.El.Style(styles.TranslateY(val))
	return c
}

func (c *Comp) TranslateYFull() *Comp {
	c.El.Style(styles.TranslateYFull())
	return c
}

func (c *Comp) TranslateYPx() *Comp {
	c.El.Style(styles.TranslateYPx())
	return c
}

func (c *Comp) TranslateZ(val any) *Comp {
	c.El.Style(styles.TranslateZ(val))
	return c
}

func (c *Comp) TranslateZPx() *Comp {
	c.El.Style(styles.TranslateZPx())
	return c
}

func (c *Comp) Truncate() *Comp {
	c.El.Style(styles.Truncate())
	return c
}

func (c *Comp) Underline() *Comp {
	c.El.Style(styles.Underline())
	return c
}

func (c *Comp) UnderlineOffset(number int) *Comp {
	c.El.Style(styles.UnderlineOffset(number))
	return c
}

func (c *Comp) UnderlineOffsetAuto() *Comp {
	c.El.Style(styles.UnderlineOffsetAuto())
	return c
}

func (c *Comp) UnderlineOffsetBy(val value.Value) *Comp {
	c.El.Style(styles.UnderlineOffsetBy(val))
	return c
}

func (c *Comp) Uppercase() *Comp {
	c.El.Style(styles.Uppercase())
	return c
}

func (c *Comp) Via(val value.Value) *Comp {
	c.El.Style(styles.Via(val))
	return c
}

func (c *Comp) Visible() *Comp {
	c.El.Style(styles.Visible())
	return c
}

func (c *Comp) W(number int) *Comp {
	c.El.Style(styles.W(number))
	return c
}

func (c *Comp) W2xl() *Comp {
	c.El.Style(styles.W2xl())
	return c
}

func (c *Comp) W2xs() *Comp {
	c.El.Style(styles.W2xs())
	return c
}

func (c *Comp) W3xl() *Comp {
	c.El.Style(styles.W3xl())
	return c
}

func (c *Comp) W3xs() *Comp {
	c.El.Style(styles.W3xs())
	return c
}

func (c *Comp) W4xl() *Comp {
	c.El.Style(styles.W4xl())
	return c
}

func (c *Comp) W5xl() *Comp {
	c.El.Style(styles.W5xl())
	return c
}

func (c *Comp) W6xl() *Comp {
	c.El.Style(styles.W6xl())
	return c
}

func (c *Comp) W7xl() *Comp {
	c.El.Style(styles.W7xl())
	return c
}

func (c *Comp) WAuto() *Comp {
	c.El.Style(styles.WAuto())
	return c
}

func (c *Comp) WBy(val value.Value) *Comp {
	c.El.Style(styles.WBy(val))
	return c
}

func (c *Comp) WDvh() *Comp {
	c.El.Style(styles.WDvh())
	return c
}

func (c *Comp) WDvw() *Comp {
	c.El.Style(styles.WDvw())
	return c
}

func (c *Comp) WFit() *Comp {
	c.El.Style(styles.WFit())
	return c
}

func (c *Comp) WFraction(fraction float32) *Comp {
	c.El.Style(styles.WFraction(fraction))
	return c
}

func (c *Comp) WFull() *Comp {
	c.El.Style(styles.WFull())
	return c
}

func (c *Comp) WLg() *Comp {
	c.El.Style(styles.WLg())
	return c
}

func (c *Comp) WLvh() *Comp {
	c.El.Style(styles.WLvh())
	return c
}

func (c *Comp) WLvw() *Comp {
	c.El.Style(styles.WLvw())
	return c
}

func (c *Comp) WMax() *Comp {
	c.El.Style(styles.WMax())
	return c
}

func (c *Comp) WMd() *Comp {
	c.El.Style(styles.WMd())
	return c
}

func (c *Comp) WMin() *Comp {
	c.El.Style(styles.WMin())
	return c
}

func (c *Comp) WPx() *Comp {
	c.El.Style(styles.WPx())
	return c
}

func (c *Comp) WScreen() *Comp {
	c.El.Style(styles.WScreen())
	return c
}

func (c *Comp) WSm() *Comp {
	c.El.Style(styles.WSm())
	return c
}

func (c *Comp) WSvh() *Comp {
	c.El.Style(styles.WSvh())
	return c
}

func (c *Comp) WSvw() *Comp {
	c.El.Style(styles.WSvw())
	return c
}

func (c *Comp) WXl() *Comp {
	c.El.Style(styles.WXl())
	return c
}

func (c *Comp) WXs() *Comp {
	c.El.Style(styles.WXs())
	return c
}

func (c *Comp) WhitespaceBreakSpaces() *Comp {
	c.El.Style(styles.WhitespaceBreakSpaces())
	return c
}

func (c *Comp) WhitespaceNormal() *Comp {
	c.El.Style(styles.WhitespaceNormal())
	return c
}

func (c *Comp) WhitespaceNowrap() *Comp {
	c.El.Style(styles.WhitespaceNowrap())
	return c
}

func (c *Comp) WhitespacePre() *Comp {
	c.El.Style(styles.WhitespacePre())
	return c
}

func (c *Comp) WhitespacePreLine() *Comp {
	c.El.Style(styles.WhitespacePreLine())
	return c
}

func (c *Comp) WhitespacePreWrap() *Comp {
	c.El.Style(styles.WhitespacePreWrap())
	return c
}

func (c *Comp) WillChange(val value.Value) *Comp {
	c.El.Style(styles.WillChange(val))
	return c
}

func (c *Comp) WillChangeAuto() *Comp {
	c.El.Style(styles.WillChangeAuto())
	return c
}

func (c *Comp) WillChangeContents() *Comp {
	c.El.Style(styles.WillChangeContents())
	return c
}

func (c *Comp) WillChangeScroll() *Comp {
	c.El.Style(styles.WillChangeScroll())
	return c
}

func (c *Comp) WillChangeTransform() *Comp {
	c.El.Style(styles.WillChangeTransform())
	return c
}

func (c *Comp) WrapAnywhere() *Comp {
	c.El.Style(styles.WrapAnywhere())
	return c
}

func (c *Comp) WrapBreakWord() *Comp {
	c.El.Style(styles.WrapBreakWord())
	return c
}

func (c *Comp) WrapNormal() *Comp {
	c.El.Style(styles.WrapNormal())
	return c
}

func (c *Comp) Z(index int) *Comp {
	c.El.Style(styles.Z(index))
	return c
}

func (c *Comp) ZAuto() *Comp {
	c.El.Style(styles.ZAuto())
	return c
}

func (c *Comp) ZBy(val value.Value) *Comp {
	c.El.Style(styles.ZBy(val))
	return c
}
