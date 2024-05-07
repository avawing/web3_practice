package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"web3_practice/week_one/day_three"
)

func TestAddTransaction(t *testing.T) {
	t.Run("should add the transaction to the mempool", func(t *testing.T) {
		txn := day_three.Transaction{To: "bob", Sender: "alice"}
		day_three.AddTransaction(txn)
		assert.Equal(t, 1, len(day_three.Mempool))
		assert.Equal(t, day_three.Mempool[0], txn)
	})
}

func TestMine(t *testing.T) {
	t.Run("first block", func(t *testing.T) {
		day_three.Mine()

		assert.Equal(t, 1, len(day_three.Blocks))
		assert.Equal(t, 0, day_three.Blocks[0].ID)
		assert.NotEqual(t, 0, day_three.Blocks[0].Nonce)
	})
	t.Run("second block", func(t *testing.T) {
		day_three.Mine()

		assert.Equal(t, 2, len(day_three.Blocks))
		assert.Equal(t, 1, day_three.Blocks[1].ID)
		assert.NotEqual(t, 0, day_three.Blocks[1].Nonce)
	})
	t.Run("with 5 mempool transactions", func(t *testing.T) {
		// setup
		day_three.Mempool = []day_three.Transaction{}
		day_three.Blocks = []day_three.BlockChain{}

		for i := 0; i < 5; i++ {
			day_three.Mempool = append(day_three.Mempool, day_three.Transaction{
				To:     "Bob",
				Sender: "Alice",
			})
		}

		assert.Equal(t, len(day_three.Mempool), 5)
		assert.Equal(t, len(day_three.Blocks), 0)

		day_three.Mine()

		assert.Equal(t, len(day_three.Mempool), 0)
		assert.Equal(t, len(day_three.Blocks), 1)
		assert.Equal(t, 0, day_three.Blocks[0].ID)
		assert.NotEqual(t, 0, day_three.Blocks[0].Nonce)

	})
	t.Run("with 15 mempool transactions", func(t *testing.T) {
		day_three.Mempool = []day_three.Transaction{}

		for i := 0; i < 15; i++ {
			day_three.Mempool = append(day_three.Mempool, day_three.Transaction{
				To:     "Bob",
				Sender: "Alice",
			})
		}

		assert.Equal(t, len(day_three.Mempool), 15)
		assert.Equal(t, len(day_three.Blocks), 1)

		day_three.Mine()
		assert.Equal(t, len(day_three.Blocks[1].Transactions), 10)
		assert.Equal(t, len(day_three.Blocks), 2)
		assert.Equal(t, len(day_three.Mempool), 5)
		assert.NotEqual(t, 0, day_three.Blocks[1].Nonce)

	})
	t.Run("after mining again", func(t *testing.T) {
		day_three.Mine()
		assert.Equal(t, 3, len(day_three.Blocks))
		assert.Equal(t, len(day_three.Blocks[2].Transactions), 5)
		assert.Equal(t, len(day_three.Mempool), 0)
		assert.NotEqual(t, 0, day_three.Blocks[2].Nonce)

	})
}
