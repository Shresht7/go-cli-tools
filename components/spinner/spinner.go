package spinner

type Spinner struct {
	// The frames of the spinner
	frames []string

	// The current frame of the spinner
	currentFrame int

	// The text to display next to the spinner
	text string
}

// Instantiates a new Spinner with the given frames
func New(frames []string) *Spinner {
	return &Spinner{
		frames: frames,
	}
}

// Sets the text of the spinner
func (s *Spinner) WithText(text string) *Spinner {
	s.text = text
	return s
}

// Returns the string representation of the spinner
func (s *Spinner) String() string {
	// Get the current frame
	frame := s.frames[s.currentFrame]

	// Increment the current frame
	s.currentFrame = (s.currentFrame + 1) % len(s.frames)

	// Return the spinner frame
	return frame + " " + s.text
}
