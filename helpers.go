package main

import (
	"fmt"
	"math/rand"
	"time"
)

func formatAcquirer(s string, length int) string {
	if len(s) > length {
		return s[0:length]
	} else if len(s) < length {
		pad := length - len(s)
		return fmt.Sprintf("%-*v", pad, s)
	}
	return s
}

func randomizeAmount(stc *SharedTestConfig) float64 {
	source := rand.NewSource(time.Now().Unix())
	randomizer := rand.New(source)
	r := stc.AmountMin + randomizer.Float64()*(stc.AmountMax-stc.AmountMin)

	return r
}

func makePartialAmount(amount float64) float64 {
	return 0
}
