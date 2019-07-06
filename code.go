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
