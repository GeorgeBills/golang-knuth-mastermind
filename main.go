package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

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
