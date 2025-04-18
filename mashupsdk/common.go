package mashupsdk

import "google.golang.org/protobuf/types/known/emptypb"

type DisplayElementState int64

const (
	Immutable      DisplayElementState = 0 // For libraries
	Mutable        DisplayElementState = 1 // For
	Init           DisplayElementState = 2
	Clicked        DisplayElementState = 4
	Moved          DisplayElementState = 8
	Hidden         DisplayElementState = 16  // Hidden objects are not clickable
	Recursive      DisplayElementState = 32  // Apply attributes recursively
	SourceExternal DisplayElementState = 64  // Event source external
	RightClick     DisplayElementState = 128 // Click was a right mouse click
	ControlClicked DisplayElementState = 256 // Click was a Ctrl mouse click
)

type MashupDisplayState int

const (
	Configured     MashupDisplayState = 1
	Position       MashupDisplayState = 2
	Frame          MashupDisplayState = 4
	AppInitted     MashupDisplayState = 8
	DisplaySettled MashupDisplayState = 31
)

type MashupDisplayContext struct {
	MainWinDisplay *MashupDisplayHint
	yOffset        int
	settled        MashupDisplayState
}

func (m *MashupDisplayContext) GetYoffset() int {
	return m.yOffset
}

func (m *MashupDisplayContext) SetYoffset(a int) int {
	m.yOffset = a
	return m.yOffset
}

func (m *MashupDisplayContext) GetSettled() MashupDisplayState {
	return m.settled
}

func (m *MashupDisplayContext) ApplySettled(s MashupDisplayState, override bool) MashupDisplayState {
	if override {
		m.settled = s
	} else {
		m.settled |= s
	}
	return m.settled
}

func (m *MashupDisplayContext) OnDisplayChange(displayHint *MashupDisplayHint) bool {
	resize := false

	if m.MainWinDisplay == nil {
		resize = true
		m.MainWinDisplay = &MashupDisplayHint{}
	}

	if displayHint == nil {
		return false
	}

	if displayHint.Xpos != 0 && (*m.MainWinDisplay).Xpos != displayHint.Xpos {
		m.MainWinDisplay.Xpos = displayHint.Xpos
		resize = true
	}
	if displayHint.Ypos != 0 && (*m.MainWinDisplay).Ypos != displayHint.Ypos {
		m.MainWinDisplay.Ypos = displayHint.Ypos + int64(m.yOffset)
		resize = true
	}
	if displayHint.Width != 0 && (*m.MainWinDisplay).Width != displayHint.Width {
		m.MainWinDisplay.Width = displayHint.Width
		resize = true
	}
	if displayHint.Height != 0 && (*m.MainWinDisplay).Height != displayHint.Height+int64(m.yOffset) {
		m.MainWinDisplay.Height = displayHint.Height
		resize = true
	}

	if displayHint.Focused {
		resize = true
	}

	if m.settled < (Configured | Position | Frame | AppInitted) {
		return false
	} else if m.settled == (Configured | Position | Frame | AppInitted) {
		resize = true
		m.settled = DisplaySettled
	}
	return resize
}

// MashupApiHandler -- mashups implement this to handle all events sent from
// other mashups.
type MashupApiHandler interface {
	OnDisplayChange(displayHint *MashupDisplayHint)
	GetElements() (*MashupDetailedElementBundle, error)
	UpsertElements(detailedElementBundle *MashupDetailedElementBundle) (*MashupDetailedElementBundle, error)
	TweakStates(elementStateBundle *MashupElementStateBundle) (*MashupElementStateBundle, error)
	ResetStates()
	TweakStatesByMotiv(*Motiv) (*emptypb.Empty, error)
}

type MashupContextInitHandler interface {
	RegisterContext(*MashupContext)
}
