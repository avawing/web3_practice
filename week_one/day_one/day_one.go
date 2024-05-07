package day_one

import (
	"crypto/sha256"
)

// the possible colors that the hash could represent
var colors = []string{"red", "green", "blue", "yellow", "pink", "orange"}

// given a hash, return the color that created the hash
func FindColor(hash []byte) string {
	for _, color := range colors {
		stepOne := []byte(color)
		h := sha256.New()
		s := h.Sum(stepOne)

		if string(s) == string(hash) {
			return color
		}
	}
	return "OOPS"
}
