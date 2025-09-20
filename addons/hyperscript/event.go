package hyperscript

import (
	"fmt"
	"strconv"
	"strings"
)

type HtmlEvent string

const (
	EventAbort              HtmlEvent = "abort"              //	The loading of a media is aborted	UiEvent, Event
	EventAfterPrint         HtmlEvent = "afterprint"         //	A page has started printing	Event
	EventAnimationEnd       HtmlEvent = "animationend"       //	A CSS animation has completed	AnimationEvent
	EventAnimationIteration HtmlEvent = "animationiteration" //	A CSS animation is repeated	AnimationEvent
	EventAnimationStart     HtmlEvent = "animationstart"     //	A CSS animation has started	AnimationEvent
	EventBeforePrint        HtmlEvent = "beforeprint"        //	A page is about to be printed	Event
	EventBeforeUnload       HtmlEvent = "beforeunload"       //	Before a document is about to be unloaded	UiEvent, Event
	EventBlur               HtmlEvent = "blur"               //	An element loses focus	FocusEvent
	EventCanPlay            HtmlEvent = "canplay"            //	The browser can start playing a media (has buffered enough to begin)	Event
	EventCanplayThrough     HtmlEvent = "canplaythrough"     //	The browser can play through a media without stopping for buffering	Event
	EventChange             HtmlEvent = "change"             //	The content of a form element has changed	Event
	EventClick              HtmlEvent = "click"              //	An element is clicked on	MouseEvent
	EventContextMenu        HtmlEvent = "contextmenu"        //	An element is right-clicked to open a context menu	MouseEvent
	EventCopy               HtmlEvent = "copy"               //	The content of an element is copied	ClipboardEvent
	EventCut                HtmlEvent = "cut"                //	The content of an element is cut	ClipboardEvent
	EventDblclick           HtmlEvent = "dblclick"           //	An element is double-clicked	MouseEvent
	EventDrag               HtmlEvent = "drag"               //	An element is being dragged	DragEvent
	EventDragEnd            HtmlEvent = "dragend"            //	Dragging of an element has ended	DragEvent
	EventDragEnter          HtmlEvent = "dragenter"          //	A dragged element enters the drop target	DragEvent
	EventDragLeave          HtmlEvent = "dragleave"          //	A dragged element leaves the drop target	DragEvent
	EventDragOver           HtmlEvent = "dragover"           //	A dragged element is over the drop target	DragEvent
	EventDragStart          HtmlEvent = "dragstart"          //	Dragging of an element has started	DragEvent
	EventDrop               HtmlEvent = "drop"               //	A dragged element is dropped on the target	DragEvent
	EventDurationChange     HtmlEvent = "durationchange"     //	The duration of a media is changed	Event
	EventEnded              HtmlEvent = "ended"              //	A media has reach the end ("thanks for listening")	Event
	EventError              HtmlEvent = "error"              //	An error has occurred while loading a file	ProgressEvent, UiEvent, Event
	EventFocus              HtmlEvent = "focus"              //	An element gets focus	FocusEvent
	EventFocusIn            HtmlEvent = "focusin"            //	An element is about to get focus	FocusEvent
	EventFocusOut           HtmlEvent = "focusout"           //	An element is about to lose focus	FocusEvent
	EventFullScreenChange   HtmlEvent = "fullscreenchange"   //	An element is displayed in fullscreen mode	Event
	EventFullScreenError    HtmlEvent = "fullscreenerror"    //	An element can not be displayed in fullscreen mode	Event
	EventHashChange         HtmlEvent = "hashchange"         //	There has been changes to the anchor part of a URL	HashChangeEvent
	EventInput              HtmlEvent = "input"              //	An element gets user input	InputEvent, Event
	EventInvalid            HtmlEvent = "invalid"            //	An element is invalid	Event
	EventKeyDown            HtmlEvent = "keydown"            //	A key is down	KeyboardEvent
	EventKeyPress           HtmlEvent = "keypress"           //	A key is pressed	KeyboardEvent
	EventKeyUp              HtmlEvent = "keyup"              //	A key is released	KeyboardEvent
	EventLoad               HtmlEvent = "load"               //	An object has loaded	UiEvent, Event
	EventLoadedData         HtmlEvent = "loadeddata"         //	Media data is loaded	Event
	EventLoadedMetadata     HtmlEvent = "loadedmetadata"     //	Meta data (like dimensions and duration) are loaded	Event
	EventLoadStart          HtmlEvent = "loadstart"          //	The browser starts looking for the specified media	ProgressEvent
	EventMessage            HtmlEvent = "message"            //	A message is received through the event source	Event
	EventMouseDown          HtmlEvent = "mousedown"          //	The mouse button is pressed over an element	MouseEvent
	EventMouseEnter         HtmlEvent = "mouseenter"         //	The pointer is moved onto an element	MouseEvent
	EventMouseLeave         HtmlEvent = "mouseleave"         //	The pointer is moved out of an element	MouseEvent
	EventMouseMove          HtmlEvent = "mousemove"          //	The pointer is moved over an element	MouseEvent
	EventMouseOver          HtmlEvent = "mouseover"          //	The pointer is moved onto an element	MouseEvent
	EventMouseOut           HtmlEvent = "mouseout"           //	The pointer is moved out of an element	MouseEvent
	EventMouseUp            HtmlEvent = "mouseup"            //	A user releases a mouse button over an element	MouseEvent
	EventMouseWheel         HtmlEvent = "mousewheel"         //	Deprecated. Use the wheel event instead	WheelEvent
	EventOffline            HtmlEvent = "offline"            //	The browser starts working offline	Event
	EventOnline             HtmlEvent = "online"             //	The browser starts working online	Event
	EventOpen               HtmlEvent = "open"               //	A connection with the event source is opened	Event
	EventPageHide           HtmlEvent = "pagehide"           //	User navigates away from a webpage	PageTransitionEvent
	EventPageShow           HtmlEvent = "pageshow"           //	User navigates to a webpage	PageTransitionEvent
	EventPaste              HtmlEvent = "paste"              //	Some content is pasted in an element	ClipboardEvent
	EventPause              HtmlEvent = "pause"              //	A media is paused	Event
	EventPlay               HtmlEvent = "play"               //	The media has started or is no longer paused	Event
	EventPlaying            HtmlEvent = "playing"            //	The media is playing after being paused or buffered	Event
	EventPopState           HtmlEvent = "popstate"           //	The window's history changes	PopStateEvent
	EventProgress           HtmlEvent = "progress"           //	The browser is downloading media data	Event
	EventRateChange         HtmlEvent = "ratechange"         //	The playing speed of a media is changed	Event
	EventResize             HtmlEvent = "resize"             //	The document view is resized	UiEvent, Event
	EventReset              HtmlEvent = "reset"              //	A form is reset	Event
	EventScroll             HtmlEvent = "scroll"             //	A scrollbar is being scrolled	UiEvent, Event
	EventSearch             HtmlEvent = "search"             //	Something is written in a search field	Event
	EventSeeked             HtmlEvent = "seeked"             //	Skipping to a media position is finished	Event
	EventSeeking            HtmlEvent = "seeking"            //	Skipping to a media position is started	Event
	EventSelect             HtmlEvent = "select"             //	User selects some text	UiEvent, Event
	EventShow               HtmlEvent = "show"               //	A <menu> element is shown as a context menu	Event
	EventStalled            HtmlEvent = "stalled"            //	The browser is trying to get unavailable media data	Event
	EventStorage            HtmlEvent = "storage"            //	A Web Storage area is updated	StorageEvent
	EventSubmit             HtmlEvent = "submit"             //	A form is submitted	Event
	EventSuspend            HtmlEvent = "suspend"            //	The browser is intentionally not getting media data	Event
	EventTimeUpdate         HtmlEvent = "timeupdate"         //	The playing position has changed (the user moves to a different point in the media)	Event
	EventToggle             HtmlEvent = "toggle"             //	The user opens or closes the <details> element	Event
	EventTouchCancel        HtmlEvent = "touchcancel"        //	The touch is interrupted	TouchEvent
	EventTouchEnd           HtmlEvent = "touchend"           //	A finger is removed from a touch screen	TouchEvent
	EventTouchMove          HtmlEvent = "touchmove"          //	A finger is dragged across the screen	TouchEvent
	EventTouchStart         HtmlEvent = "touchstart"         //	A finger is placed on a touch screen	TouchEvent
	EventTransitionEnd      HtmlEvent = "transitionend"      //	A CSS transition has completed	TransitionEvent
	EventUnload             HtmlEvent = "unload"             //	A page has unloaded	UiEvent, Event
	EventVolumeChange       HtmlEvent = "volumechange"       //	The volume of a media is changed (includes muting)	Event
	EventWaiting            HtmlEvent = "waiting"            //	A media is paused but is expected to resume (e.g. buffering)	Event
	EventWheel              HtmlEvent = "wheel"              //	The mouse wheel rolls up or down over an element	WheelEvent
)

func (htmlEv *HtmlEvent) String() string {
	return string(*htmlEv)
}

func EventCustom(name string) HtmlEvent {
	return HtmlEvent(name)
}

type ev struct {
	event    HtmlEvent
	params   []string
	counts   []int
	from     string
	debounce *Time
	throttle *Time
}

type EventHandler struct {
	events  []*ev
	end     bool
	command Command
}

func (e *EventHandler) Or(event HtmlEvent, params ...string) *EventHandler {
	e.events = append(e.events, &ev{event: event, params: params})
	return e
}

func (e *EventHandler) End() *EventHandler {
	if len(e.events) > 0 {
		e.end = true
	}
	return e
}

func (e *EventHandler) Count(count ...int) *EventHandler {
	e.events[len(e.events)-1].counts = count
	return e
}

func (e *EventHandler) From(expr string) *EventHandler {
	e.events[len(e.events)-1].from = expr

	return e
}

func (e *EventHandler) DebouncedAt(time *Time) *EventHandler {

	e.events[len(e.events)-1].debounce = time

	return e
}

func (e *EventHandler) ThrottledAt(time *Time) *EventHandler {
	e.events[len(e.events)-1].throttle = time

	return e
}

func (e *EventHandler) Command(command Command) *EventHandler {
	e.command = command

	return e
}

func (e *EventHandler) String() string {
	scripts := []string{"on"}
	for i, ev := range e.events {
		if i >= 1 {
			scripts = append(scripts, "or")
		}

		eventName := ev.event.String()
		if len(ev.params) > 0 {
			eventName = fmt.Sprintf("%s(%s)", eventName, strings.Join(ev.params, ", "))
		}

		s := []string{eventName}
		if len(ev.counts) == 1 {
			s = append(s, strconv.Itoa(ev.counts[0]))
		}
		if len(ev.counts) > 1 {
			s = append(s, strconv.Itoa(ev.counts[0]), "to", strconv.Itoa(ev.counts[len(ev.counts)-1]))
		}

		if ev.from != "" {
			s = append(s, "from", ev.from)
		}

		if ev.debounce != nil {
			s = append(s, "debounced at "+ev.debounce.String())
		}

		if ev.throttle != nil {
			s = append(s, "throttled at "+ev.throttle.String())
		}

		scripts = append(scripts, s...)
	}

	if e.command != nil {
		scripts = append(scripts, e.command.String())

	}

	if e.end {
		scripts = append(scripts, "end")
	}

	return strings.Join(scripts, " ")
}

func On(event HtmlEvent, params ...string) *EventHandler {
	e := &EventHandler{}
	e.events = append(e.events, &ev{event: event, params: params})
	return e
}
