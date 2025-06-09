package core

const (
	FocusURL focusableComponent = iota
	FocusViewport
	FocusReqMethod
	NumOfFocusableComponents
)

type reqMethod string

func (r reqMethod) FilterValue() string { return string(r) }
func (r reqMethod) String() string      { return string(r) }

type focusableComponent int
type httpResMsg string
type errMsg error

type methodDelegate struct{}

func (d methodDelegate) Height() int  { return 1 }
func (d methodDelegate) Spacing() int { return 0 }
