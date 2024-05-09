package test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"web3_practice/week_one/day_four"
)

func TestHashData(t *testing.T) {
	t.Run("should hash some random data", func(t *testing.T) {
		email := "bob@gmail.com"
		hash, _ := json.Marshal(email)
		h := sha256.New()
		h.Write(hash)

		expected := hex.EncodeToString(h.Sum(nil))

		block := day_four.Block{Data: email}
		retHash := block.ToHash()

		assert.Equal(t, expected, retHash)
	})
}

func TestBlockChain(t *testing.T) {
	bc := day_four.BlockChain{Chain: make([]day_four.Block, 1)}
	t.Run("should return the block chain", func(t *testing.T) {
		block1 := day_four.Block{Data: "Some Data"}
		block2 := day_four.Block{Data: "Some Other Data"}

		bc.AddBlock(block1)
		bc.AddBlock(block2)

		assert.Equal(t, block1.Data, bc.Chain[1].Data)
		assert.Equal(t, block2.Data, bc.Chain[2].Data)
		assert.Equal(t, bc.Chain[2].PreviousHash, bc.Chain[1].ToHash())
		assert.Equal(t, bc.Chain[1].PreviousHash, bc.Chain[0].ToHash())
		assert.Equal(t, 3, len(bc.Chain))

		isValid := bc.IsValid()
		assert.Equal(t, true, isValid)
	})
	t.Run("tampering with previous hash", func(t *testing.T) {
		gibberish := "gibberish"
		hash, _ := json.Marshal(gibberish)
		h := sha256.New()
		h.Write(hash)

		tamper := hex.EncodeToString(h.Sum(nil))
		bc.Chain[1].PreviousHash = tamper

		isValid := bc.IsValid()
		assert.Equal(t, false, isValid)
	})
}
