package main

import (
	"testing"
)

func TestAssess(t *testing.T) {
	tests := []struct {
		name        string
		code, guess code
		expected    feedback
	}{
		{
			name:     "all incorrect",
			code:     code{cpRed, cpGreen, cpYellow, cpBlue},
			guess:    code{cpBlack, cpBlack, cpBlack, cpBlack},
			expected: feedback{numBlack: 0, numWhite: 0},
		},
		{
			name:     "all correct colours and correct positions",
			code:     code{cpRed, cpGreen, cpYellow, cpBlue},
			guess:    code{cpRed, cpGreen, cpYellow, cpBlue},
			expected: feedback{numBlack: 4, numWhite: 0},
		},
		{
			name:     "all correct colours, but incorrect positions",
			code:     code{cpRed, cpGreen, cpYellow, cpBlue},
			guess:    code{cpBlue, cpYellow, cpGreen, cpRed},
			expected: feedback{numBlack: 0, numWhite: 4},
		},
		{
			name:     "one guess shouldn't give two key pegs",
			code:     code{cpRed, cpRed, cpBlue, cpGreen},
			guess:    code{cpRed, cpRed, cpRed, cpRed},
			expected: feedback{numBlack: 2, numWhite: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			feedback := tt.code.assess(tt.guess)
			if feedback != tt.expected {
				t.Errorf("Expected %+v but was %+v", tt.expected, feedback)
			}
		})
	}
}

func BenchmarkAssess(b *testing.B) {
	c := code{cpRed, cpYellow, cpGreen, cpBlue}
	guess := code{cpRed, cpGreen, cpYellow, cpWhite}
	for i := 0; i < b.N; i++ {
		_ = c.assess(guess)
	}
}
