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
type code [numSlots]codePeg

const numSlots = 4

func (c code) String() string {
	return strings.Join([]string{c[0].String(), c[1].String(), c[2].String(), c[3].String()}, " ")
}

func (c code) assess(guess code) feedback {
	f := feedback{}

	// count the colours in the code
	var colourCount [maxCodePeg + 1]uint8
	for i := 0; i < numSlots; i++ {
		colour := c[i]
		colourCount[colour]++
	}

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
	for i := 0; i < numSlots; i++ {
		assessSlot(guess[i], c[i])
	}

	return f
}
