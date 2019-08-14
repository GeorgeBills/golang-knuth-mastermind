package main

import (
	"testing"
)

func TestFeedbackString(t *testing.T) {
	tests := []struct {
		feedback feedback
		expected string
	}{
		{
			feedback{numBlack: 0, numWhite: 0},
			"-",
		},
		{
			feedback{numBlack: 4, numWhite: 0},
			"BBBB",
		},
		{
			feedback{numBlack: 0, numWhite: 3},
			"WWW",
		},
		{
			feedback{numBlack: 2, numWhite: 2},
			"BBWW",
		},
	}
	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			s := tt.feedback.String()
			if s != tt.expected {
				t.Errorf("Expected %s string for feedback %+v, got %s", tt.expected, tt.feedback, s)
			}
		})
	}
}
