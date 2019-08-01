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

	makeGuess := func(guess code) feedback {
		fmt.Printf("Guessing %s\n", guess)
		guessch <- guess
		return <-feedbackch
	}

	// loop until we've guessed correctly
	for {
		// take the next guess
		guess := pickGuess(possibles, unguessed)
		feedback := makeGuess(guess)

		fmt.Printf("Feedback: %v\n", feedback)

		if feedback.isCorrect() {
			break
		}

		// eliminate any codes that wouldn't have produced this result
		possibles = eliminateCodes(possibles, guess, feedback)
		fmt.Printf("%d possible codes remaining\n", len(possibles))
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
		isPossible   bool
	}
	best := candidate{
		minRemaining: len(possibleCodes) + 1,
	}

	// for each guess we could make...
	for _, candidateGuess := range unguessedCodes {
		var maxRemaining int
		var isPossible bool

		// for each feedback the codemaker could possibly give for that guess..
		for _, feedback := range possibleFeedbacks {

			// how many possible codes would remain after that feedback?
			numRemaining := 0
			for _, possibleCode := range possibleCodes {

				if candidateGuess == possibleCode {
					// our guess is in the set of possible codes
					isPossible = true
				}

				actual := possibleCode.assess(candidateGuess)
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

		// if the maximum remaining possibilities after making this guess are
		// less than the current best minimum then this is the new best guess.
		// this is also the new best guess if the number of possiblities are
		// equal and this guess is possibly the code (i.e. if we could win with
		// this answer).
		if maxRemaining < best.minRemaining ||
			(!best.isPossible && isPossible && maxRemaining == best.minRemaining) {
			best = candidate{candidateGuess, maxRemaining, isPossible}
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
