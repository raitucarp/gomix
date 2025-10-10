package gomix

type Scope string

const (
	AppScope      Scope = "app"
	WebScope      Scope = "web"
	PageScope     Scope = "page"
	FragmentScope Scope = "fragment"
	RestScope     Scope = "rest"
	SocketScope   Scope = "ws"
	SSEScope      Scope = "sse"
	AssetsScope   Scope = "assets"
)
