package quadratic

import (
	"testing"
)

func TestSolveQuadraticCoefficients(t *testing.T) {
	want := [3]int{1, -3, 2}
	if got := SolveQuadraticCoefficients(1, 2); got != want {
		t.Errorf("SolveQuadraticCoefficients() = %q, want %q", got, want)
	}
}

func TestSolveQuadraticCoefficientsZero(t *testing.T) {
	want := [3]int{1, 0, 0}
	if got := SolveQuadraticCoefficients(0, 0); got != want {
		t.Errorf("SolveQuadraticCoefficients() = %d, want %d", got, want)
	}
}
