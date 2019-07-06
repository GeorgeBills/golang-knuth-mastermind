package main

// feedback is given through the codemaster placing up to 4 black and white pegs
type feedback struct {
	numBlack uint8 // black pegs indicate a correctly guessed colour in the correct slot
	numWhite uint8 // white pegs indicate a correctly guessed colour in an incorrect slot
}

func (r feedback) isCorrect() bool {
	return r.numBlack == numSlots
}
