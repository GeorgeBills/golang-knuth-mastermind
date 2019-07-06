package main

import (
	"strings"
)

//go:generate stringer -type=codePeg -linecomment=true

// codePeg is one of the coloured pegs that can be placed in a slot.
type codePeg uint8

const (
	cpBlack    codePeg = iota // ⚫
	cpWhite                   // ⚪
	cpRed                     // ❤️
	cpYellow                  // 💛
	cpGreen                   // 💚
	cpBlue                    // 🔵
	maxCodePeg = cpBlue
)

// code represents the four slots that each contain a codePeg.
type code struct {
	slot1 codePeg
	slot2 codePeg
	slot3 codePeg
	slot4 codePeg
}

const numSlots = 4

func (c code) String() string {
	return strings.Join([]string{c.slot1.String(), c.slot2.String(), c.slot3.String(), c.slot4.String()}, " ")
}

func (c code) assess(guess code) feedback {
	f := feedback{}

	// count the colours in the code
	var colourCount [maxCodePeg + 1]uint8
	colourCount[c.slot1]++
	colourCount[c.slot2]++
	colourCount[c.slot3]++
	colourCount[c.slot4]++

	// assess each code peg in the guess
	assessSlot := func(guess, real codePeg) {
		if guess == real {
			// if the guess code peg matches the real code peg in the same slot,
			// then we indicate that with a black key peg
			f.numBlack++
			colourCount[guess]--
		} else if colourCount[guess] > 0 {
			// if the guess peg matches one of the colours we have, but wasn't
			// in the correct slot, then we indicate that with a white key peg
			f.numWhite++
			colourCount[guess]--
		}
	}
	assessSlot(guess.slot1, c.slot1)
	assessSlot(guess.slot2, c.slot2)
	assessSlot(guess.slot3, c.slot3)
	assessSlot(guess.slot4, c.slot4)

	return f
}
