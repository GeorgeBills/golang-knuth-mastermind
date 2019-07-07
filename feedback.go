package main

import (
	"strings"
)

// feedback is given through the codemaster placing up to 4 black and white pegs
type feedback struct {
	numBlack uint8 // black pegs indicate a correctly guessed colour in the correct slot
	numWhite uint8 // white pegs indicate a correctly guessed colour in an incorrect slot
}

func (r feedback) isCorrect() bool {
	return r.numBlack == numSlots
}

func (r feedback) String() string {
	if r.numBlack == 0 && r.numWhite == 0 {
		return "-"
	}
	return strings.Repeat("B", int(r.numBlack)) + strings.Repeat("W", int(r.numWhite))
}
