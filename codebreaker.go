package main

import (
	"fmt"
	"math"
	"sync"
)

func codebreaker(guessch chan<- code, feedbackch <-chan feedback, wg *sync.WaitGroup) {
	possibles := getPossibleCodes()
	fmt.Printf("%d possible codes\n", len(possibles))

	// make initial guess
	guess := code{cpBlack, cpBlack, cpWhite, cpWhite}
	guessch <- guess
	feedback := <-feedbackch

	// loop until we've guessed correctly
	for !feedback.isCorrect() {
		// literally the worst guessing algorithm possible
		guess, possibles = possibles[0], possibles[1:]
		fmt.Printf("Guessing %s\n", guess)
		guessch <- guess
		feedback = <-feedbackch
	}

	close(guessch)
	wg.Done()
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
