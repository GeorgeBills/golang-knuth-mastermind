package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func codemaker(seed int64, guessch <-chan code, feedbackch chan<- feedback, wg *sync.WaitGroup) {
	source := rand.NewSource(seed)
	rnd := rand.New(source)
	code := randomCode(rnd)
	fmt.Printf("Code: %s (seed %d)\n", code, seed)

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

func randomCode(rnd *rand.Rand) code {
	randomPeg := func() codePeg {
		return codePeg(rnd.Intn(int(maxCodePeg + 1)))
	}
	c := code{}
	for i := 0; i < numSlots; i++ {
		c[i] = randomPeg()
	}
	return c
}
