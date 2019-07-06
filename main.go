package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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
	return c.slot1.String() + " " + c.slot2.String() + " " + c.slot3.String() + " " + c.slot4.String()
}

// feedback is given through the codemaster placing up to 4 black and white pegs
type feedback struct {
	numBlack uint8 // black pegs indicate a correctly guessed colour in the correct slot
	numWhite uint8 // white pegs indicate a correctly guessed colour in an incorrect slot
}

func (r feedback) isCorrect() bool {
	return r.numBlack == numSlots
}

func main() {
	seed := time.Now().Unix()
	guessch := make(chan code)
	feedbackch := make(chan feedback)
	go codemaker(seed, guessch, feedbackch)
	go codebreaker(guessch, feedbackch)
}

func codemaker(seed int64, guessch <-chan code, feedbackch chan<- feedback) {
	source := rand.NewSource(seed)
	rnd := rand.New(source)
	code := randomCode(rnd)
	fmt.Printf("Code: %s\n", code)
}

func codebreaker(guessch chan<- code, feedbackch <-chan feedback) {
	possibles := getPossibleCodes()
	fmt.Printf("%d possible codes\n", len(possibles))
}

func randomCode(rnd *rand.Rand) code {
	randomPeg := func() codePeg {
		return codePeg(rnd.Intn(int(maxCodePeg + 1)))
	}
	return code{
		slot1: randomPeg(),
		slot2: randomPeg(),
		slot3: randomPeg(),
		slot4: randomPeg(),
	}
}

func getPossibleCodes() []code {
	numCodes := int(math.Pow(float64(maxCodePeg), numSlots))
	codes := make([]code, 0, numCodes)
	var i, j, k, l codePeg
	for i = 0; i <= maxCodePeg; i++ {
		for j = 0; j <= maxCodePeg; j++ {
			for k = 0; k <= maxCodePeg; k++ {
				for l = 0; l <= maxCodePeg; l++ {
					codes = append(codes, code{i, j, k, l})
				}
			}
		}
	}
	return codes
}
