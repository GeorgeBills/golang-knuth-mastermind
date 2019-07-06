package main

// codePeg is one of the coloured pegs that can be placed in a slot.
type codePeg uint8

const (
	cpBlack codePeg = iota
	cpWhite
	cpRed
	cpYellow
	cpGreen
	cpBlue
)

// code represents the four slots that each contain a codePeg.
type code struct {
	slot1 codePeg
	slot2 codePeg
	slot3 codePeg
	slot4 codePeg
}

// keyPeg is one of the pegs that can be placed to indicate a correctly guessed
// colour, which may or may not be in the correct slot.
type keyPeg uint8

const (
	// kpBlack indicates a correctly guessed color in the correct slot.
	kpBlack keyPeg = iota

	// kpWhite indicates a correctly guessed colour in an incorrect slot.
	kpWhite
)

func main() {

}
