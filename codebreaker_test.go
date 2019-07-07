package main

import (
	"testing"
)

func TestEliminateCodes(t *testing.T) {
	// toMap converts codes to a map
	toMap := func(codes []code) map[code]struct{} {
		m := make(map[code]struct{}, len(codes))
		for _, c := range codes {
			m[c] = struct{}{}
		}
		return m
	}

	// setDifference returns codes that are present in set a but not in set b
	setDifference := func(a, b []code) []code {
		bmap := toMap(b)
		diff := make([]code, 0)
		for _, code := range a {
			_, contains := bmap[code]
			if !contains {
				diff = append(diff, code)
			}
		}
		return diff
	}

	tests := []struct {
		name     string
		guess    code
		feedback feedback
		original []code
		expected []code
	}{
		{
			name:     "1 correct colour, 0 correct positions",
			guess:    code{cpBlue, cpGreen, cpYellow, cpRed},
			feedback: feedback{numWhite: 1},
			original: []code{
				{cpBlack, cpBlack, cpBlack, cpBlack}, // no matching colours
				{cpWhite, cpWhite, cpWhite, cpBlue},  // ✓
				{cpGreen, cpWhite, cpGreen, cpGreen}, // ✓
				{cpBlack, cpGreen, cpBlack, cpBlack}, // would be black key peg, not white
			},
			expected: []code{
				{cpWhite, cpWhite, cpWhite, cpBlue},
				{cpGreen, cpWhite, cpGreen, cpGreen},
			},
		},
		{
			name:     "4 correct colours, 0 correct positions",
			guess:    code{cpBlue, cpGreen, cpYellow, cpRed},
			feedback: feedback{numWhite: 4},
			original: []code{
				{cpBlack, cpBlack, cpBlack, cpBlack}, // no matching colours
				{cpBlue, cpGreen, cpRed, cpYellow},   // would be 2 black key pegs and 2 white
				{cpRed, cpYellow, cpGreen, cpBlue},   // ✓
				{cpGreen, cpBlue, cpRed, cpYellow},   // ✓
			},
			expected: []code{
				{cpRed, cpYellow, cpGreen, cpBlue},
				{cpGreen, cpBlue, cpRed, cpYellow},
			},
		},
		{
			name:     "1 code peg of correct colour in the correct position",
			guess:    code{cpBlue, cpGreen, cpYellow, cpRed},
			feedback: feedback{numBlack: 1},
			original: []code{
				{cpBlack, cpBlack, cpBlack, cpBlack},   // no matching pegs
				{cpBlue, cpRed, cpYellow, cpGreen},     // would be 1 black key peg and 3 white
				{cpBlue, cpBlack, cpBlack, cpBlack},    // ✓
				{cpWhite, cpYellow, cpYellow, cpWhite}, // ✓
			},
			expected: []code{
				{cpBlue, cpBlack, cpBlack, cpBlack},
				{cpWhite, cpYellow, cpYellow, cpWhite},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := eliminateCodes(tt.original, tt.guess, tt.feedback)
			for _, c := range setDifference(tt.expected, result) {
				t.Errorf("Expected code %s wasn't in result", c)
			}
			for _, c := range setDifference(result, tt.expected) {
				t.Errorf("Unexpected code %s found in result", c)
			}
		})
	}
}
