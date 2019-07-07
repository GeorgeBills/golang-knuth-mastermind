package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {
	seed := time.Now().Unix()
	guessch := make(chan code)
	feedbackch := make(chan feedback)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go codemaker(seed, guessch, feedbackch, wg)
	go codebreaker(guessch, feedbackch, wg)
	wg.Wait()
}

func codemaker(seed int64, guessch <-chan code, feedbackch chan<- feedback, wg *sync.WaitGroup) {
	source := rand.NewSource(seed)
	rnd := rand.New(source)
	code := randomCode(rnd)
	fmt.Printf("Code: %s\n", code)

	// assess guesses until the codebreaker guesses correctly
	numGuesses := 0
	for guess := range guessch {
		numGuesses++
		feedback := code.assess(guess)
		feedbackch <- feedback
		if feedback.isCorrect() {
			break
		}
	}

	fmt.Printf("Finished: guesser took %d guesses\n", numGuesses)
	close(feedbackch)
	wg.Done()
}

func codebreaker(guessch chan<- code, feedbackch <-chan feedback, wg *sync.WaitGroup) {
	possibles := getPossibleCodes()
	fmt.Printf("%d possible codes\n", len(possibles))

	// make initial guess
	guessch <- code{cpBlack, cpBlack, cpWhite, cpWhite}
	feedback := <-feedbackch

	// loop until we've guessed correctly
	for !feedback.isCorrect() {
		// literally the worst guessing algorithm possible
		guess := possibles[0]
		possibles = possibles[1:]
		fmt.Printf("Guessing %s\n", guess)
		guessch <- guess
		feedback = <-feedbackch
	}

	close(guessch)
	wg.Done()
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
