package styles

import (
	"github.com/raitucarp/gomix/value"
)

func AnimateSpin() ApplyProp {
	return func(s *Style) StyleProp {

		return &Properties{
			string(animationProp): s.theme.UseVarKey(themeAnimate, "spin"),
			string(rawProp): `
				@keyframes spin {
					to {
						transform: rotate(360deg);
					}
				}
			`,
		}
	}
}

func AnimatePing() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(animationProp): s.theme.UseVarKey(themeAnimate, "ping"),
			string(rawProp): `
				@keyframes ping {
					75%, 100% {
						transform: scale(2);
						opacity: 0;
					}
				}
			`,
		}
	}
}

func AnimatePulse() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(animationProp): s.theme.UseVarKey(themeAnimate, "pulse"),
			string(rawProp): `
				@keyframes pulse {
					50% {
						opacity: 0.5;
					}
				}
			`,
		}
	}
}

func AnimateBounce() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(animationProp): s.theme.UseVarKey(themeAnimate, "bounce"),
			string(rawProp): `
				@keyframes bounce {
					0%, 100% {
						transform: translateY(-25%);
						animation-timing-function: cubic-bezier(0.8, 0, 1, 1);
					}
					50% {
						transform: none;
						animation-timing-function: cubic-bezier(0, 0, 0.2, 1);
					}
				}
			`,
		}
	}
}

func AnimateNone() ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(animationProp): "none",
		}
	}
}

func Animate(val value.Value) ApplyProp {
	return func(s *Style) StyleProp {
		return &Properties{
			string(animationProp): val.Value(),
		}
	}
}
