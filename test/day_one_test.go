package test

import (
	"crypto/sha256"
	"github.com/stretchr/testify/assert"
	"testing"
	"web3_practice/week_one/day_one"
)

var colors = []string{"red", "green", "blue", "yellow", "pink", "orange"}

func TestFindColor(t *testing.T) {
	t.Run("findColor", func(t *testing.T) {
		for _, color := range colors {
			stepOne := []byte(color)
			h := sha256.New()
			s := h.Sum(stepOne)
			response := day_one.FindColor(s)
			assert.Equal(t, color, response)
		}
	})
}
