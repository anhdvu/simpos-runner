package main

import (
	"fmt"
	"math/rand"
	"time"
)

func formatAcquirer(s string, length int) string {
	if len(s) > length {
		return s[0:length]
	}
	return fmt.Sprintf("%-*v", length, s)
}

func randomizeAmount(s SharedConfig) float64 {
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	r := s.AmountMin + randomizer.Float64()*(s.AmountMax-s.AmountMin)

	return r
}

func makePartialAmount(amount float64) float64 {
	return amount * 0.1
}
