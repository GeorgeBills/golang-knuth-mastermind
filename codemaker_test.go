package main

import (
	"sync"
	"testing"
)

func TestCodemaker(t *testing.T) {
	var seed int64 = 12345678 // code should be ⚪⚫💚⚪
	guessch := make(chan code)
	feedbackch := make(chan feedback)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go codemaker(seed, guessch, feedbackch, wg)

	// feed in a sequence of guesses, finishing with a correct guess
	sequence := []struct {
		guess    code
		expected feedback
	}{
		{
			code{cpRed, cpRed, cpRed, cpRed},
			feedback{numBlack: 0, numWhite: 0},
		},
		{
			code{cpWhite, cpGreen, cpBlack, cpWhite},
			feedback{numBlack: 2, numWhite: 2},
		},
		{
			code{cpWhite, cpBlack, cpGreen, cpWhite},
			feedback{numBlack: 4},
		},
	}
	for _, ge := range sequence {
		guessch <- ge.guess
		fb := <-feedbackch
		if fb != ge.expected {
			t.Errorf("Wrong feedback: %s; expected: %s", fb, ge.expected)
		}
	}

	// if this deadlocks then that's a bug you should fix: the codemaker should
	// signal wg.Done() on the final correct answer
	wg.Wait()
}
