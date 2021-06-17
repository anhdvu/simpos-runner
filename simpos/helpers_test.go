package simpos

import "testing"

func TestFormatAcquirer(t *testing.T) {
	want := "dace      "

	got := formatAcquirer("dace", 10)

	if want != got {
		t.Errorf("Wanted %q but got %q", want, got)
	}
}

func TestRandomizeAmount(t *testing.T) {
	s := SharedConfig{
		AmountMin: 10.0,
		AmountMax: 15.0,
	}

	got := randomizeAmount(s)

	if got < s.AmountMin || got > s.AmountMax {
		t.Errorf("Wanted a random number that is in range (10.0, 15.0), but got %f", got)
	}
}

func TestMakePartialAmount(t *testing.T) {
	want := 1.0
	got := makePartialAmount(10.0)

	if want != got {
		t.Errorf("Wanted %f but got %f", want, got)
	}
}
