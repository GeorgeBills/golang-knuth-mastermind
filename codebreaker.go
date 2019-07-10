package main

import (
	"fmt"
	"math"
	"sync"
)

func codebreaker(guessch chan<- code, feedbackch <-chan feedback, wg *sync.WaitGroup) {
	// possibles is the set of codes that are still viable answers
	possibles := getPossibleCodes()

	// unguessed is the full set of codes, minus codes we've already guessed
	unguessed := make([]code, len(possibles))
	copy(unguessed, possibles)

	// make initial guess
	fmt.Printf("%d possible codes\n", len(possibles))
	guess := code{cpBlack, cpBlack, cpWhite, cpWhite}
	guessch <- guess
	feedback := <-feedbackch

	// loop until we've guessed correctly
	for !feedback.isCorrect() {
		// eliminate any codes that wouldn't have produced this result
		possibles = eliminateCodes(possibles, guess, feedback)

		// take the next guess
		fmt.Printf("%d possible codes\n", len(possibles))
		guess = pickGuess(possibles, unguessed)
		fmt.Printf("Guessing %s\n", guess)
		guessch <- guess
		feedback = <-feedbackch
		fmt.Printf("Feedback: %v\n", feedback)
	}

	close(guessch)
	wg.Done()
}

func eliminateCodes(codes []code, guess code, fb feedback) []code {
	var ret []code
	for _, code := range codes {
		actual := code.assess(guess)
		if actual == fb {
			ret = append(ret, code)
		}
	}
	return ret
}

// pickGuess chooses the next code to guess.
//
// We pick the guess that "minimizes the maximum number of remaining
// possibilities over all 15 responses by the codemaker", i.e. for each
// candidate guess we assess the maximum number of remaining possibilities that
// may remain after making that guess, and then pick the candidate guess that
// minimizes that number.
//
// pickGuess may choose a guess from unguessed that isn't contained within
// possibles; there are codes we could guess that can't possibly be the answer,
// but that would eliminate the greatest number of possibles if guessed. "A
// codeword which cannot possibly win in four is necessary here in order to win
// in five".
func pickGuess(possibleCodes, unguessedCodes []code) code {
	possibleFeedbacks := []feedback{
		{numBlack: 0, numWhite: 0},
		{numBlack: 0, numWhite: 1},
		{numBlack: 0, numWhite: 2},
		{numBlack: 0, numWhite: 3},
		{numBlack: 0, numWhite: 4},
		{numBlack: 1, numWhite: 0},
		{numBlack: 1, numWhite: 1},
		{numBlack: 1, numWhite: 2},
		{numBlack: 1, numWhite: 3},
		{numBlack: 2, numWhite: 0},
		{numBlack: 2, numWhite: 1},
		{numBlack: 2, numWhite: 2},
		{numBlack: 3, numWhite: 0},
		{numBlack: 3, numWhite: 1},
		{numBlack: 4, numWhite: 0},
	}

	type candidate struct {
		guess        code
		minRemaining int
	}
	best := candidate{
		minRemaining: len(possibleCodes) + 1,
	}

	// if it's a coin flip then just pick at random. this is necessary to avoid
	// a pathological case where the algorithm finds that almost any possible
	// guess will at best knock out one of the two alternatives, and then keeps
	// picking guesses that aren't possible answers.
	// TODO: fix the algorithm to prefer guesses that are also in possibleCodes
	if len(possibleCodes) <= 2 {
		return possibleCodes[0]
	}

	// for each guess we could make...
	for _, possibleGuess := range unguessedCodes {
		var maxRemaining int

		// for each feedback the codemaker could possibly give for that guess..
		for _, feedback := range possibleFeedbacks {

			// how many possible codes would remain after that feedback?
			numRemaining := 0
			for _, possibleCode := range possibleCodes {
				actual := possibleCode.assess(possibleGuess)
				if actual == feedback {
					numRemaining++
				}
			}
			// numRemaining is how many possible codes would remain possible
			// after our guess followed by the codemakers feedback.
			if numRemaining > maxRemaining {
				maxRemaining = numRemaining
			}
		}

		// if the maximum remaining possiblities after making this guess are
		// less than the current best minimum then it's the new best guess.
		if maxRemaining < best.minRemaining {
			best = candidate{possibleGuess, maxRemaining}
		}
	}

	return best.guess
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
