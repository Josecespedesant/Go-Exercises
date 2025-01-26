package de_emojify

import (
	"testing"
)

func TestSolveQuadraticCoefficients(t *testing.T) {
	want := "hi"
	if got := De_Emojify(":D :) :/  :D :) :|"); got != want {
		t.Errorf("De_Emojify() = %s, want %s", got, want)
	}
}
