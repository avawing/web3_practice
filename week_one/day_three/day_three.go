package day_three

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

var (
	TARGET_DIFFICULTY = "0fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	MAX_TRANSACTIONS  = 10
	Mempool           []Transaction
	Blocks            []BlockChain
)

type Transaction struct {
	To     string
	Sender string
}
type Block struct {
	ID           int
	Transactions []Transaction
	Nonce        uint
}

type BlockChain struct {
	Block
	Hash string
}

func AddTransaction(transaction Transaction) {
	Mempool = append(Mempool, transaction)
}

func Mine() {
	var transactions []Transaction
	for len(transactions) < MAX_TRANSACTIONS && len(Mempool) > 0 {
		transactions = append(transactions, Mempool[len(Mempool)-1])
		Mempool = Mempool[:len(Mempool)-1]
	}

	block := Block{
		ID:           len(Blocks),
		Transactions: transactions,
	}

	block.Nonce = 0
	for {
		hash, _ := json.Marshal(block)
		h := sha256.New()
		h.Write(hash)
		s := hex.EncodeToString(h.Sum(nil))
		if s < TARGET_DIFFICULTY {
			chain := BlockChain{
				Block: block,
				Hash:  s,
			}
			Blocks = append(Blocks, chain)
			return
		}
		block.Nonce++
	}
}
