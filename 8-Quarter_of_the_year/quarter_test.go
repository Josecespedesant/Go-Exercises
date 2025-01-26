package quarter

import (
	"testing"
)

func TestGetQuarterFirst(t *testing.T) {
	want := 1
	if got := GetQuarter(1); got != want {
		t.Errorf("GetQuarter() = %d, want %d", got, want)
	}
}

func TestGetQuarterSecond(t *testing.T) {
	want := 2
	if got := GetQuarter(4); got != want {
		t.Errorf("GetQuarter() = %d, want %d", got, want)
	}
}

func TestGetQuarterThird(t *testing.T) {
	want := 3
	if got := GetQuarter(7); got != want {
		t.Errorf("GetQuarter() = %d, want %d", got, want)
	}
}

func TestGetQuarterFourth(t *testing.T) {
	want := 4
	if got := GetQuarter(12); got != want {
		t.Errorf("GetQuarter() = %d, want %d", got, want)
	}
}

func TestGetQuarterOutOfBounds(t *testing.T) {
	want := -1
	if got := GetQuarter(13); got != want {
		t.Errorf("GetQuarter() = %d, want %d", got, want)
	}
}
