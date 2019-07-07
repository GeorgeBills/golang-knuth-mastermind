package main

import (
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
