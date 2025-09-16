package animations

type Animation string

const (
	Animation_Fade_In  Animation = "--animation-fade-in"
	Animation_Fade_Out Animation = "--animation-fade-out"

	Animation_Fade_In_Bloom  Animation = "--animation-fade-in-bloom"
	Animation_Fade_Out_Bloom Animation = "--animation-fade-out-bloom"

	Animation_Shake_X Animation = "--animation-shake-x"
	Animation_Shake_Y Animation = "--animation-shake-y"
	Animation_Shake_Z Animation = "--animation-shake-z"

	Animation_Slide_Out_Up    Animation = "--animation-slide-out-up"
	Animation_Slide_Out_Down  Animation = "--animation-slide-out-down"
	Animation_Slide_Out_Left  Animation = "--animation-slide-out-left"
	Animation_Slide_Out_Right Animation = "--animation-slide-out-right"

	Animation_Slide_In_Up    Animation = "--animation-slide-in-up"
	Animation_Slide_In_Down  Animation = "--animation-slide-in-down"
	Animation_Slide_In_Left  Animation = "--animation-slide-in-left"
	Animation_Slide_In_Right Animation = "--animation-slide-in-right"

	Animation_Spin   Animation = "--animation-spin"
	Animation_Ping   Animation = "--animation-ping"
	Animation_Blink  Animation = "--animation-blink"
	Animation_Flout  Animation = "--animation-flout"
	Animation_Bounce Animation = "--animation-bounce"
	Animation_Pulse  Animation = "--animation-pulse"
)
