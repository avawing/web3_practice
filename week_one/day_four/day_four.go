package day_four

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Block struct {
	Data         string
	PreviousHash string
}

func (b *Block) ToHash() string {
	hash, _ := json.Marshal(b)
	h := sha256.New()
	h.Write(hash)
	return hex.EncodeToString(h.Sum(nil))
}

type BlockChain struct {
	Chain []Block
}

func (bc *BlockChain) Genesis() *Block {
	return &bc.Chain[0]
}

func (bc *BlockChain) AddBlock(block Block) {
	block.PreviousHash = bc.Chain[len(bc.Chain)-1].ToHash()
	bc.Chain = append(bc.Chain, block)
}

func (bc *BlockChain) IsValid() bool {
	for i := len(bc.Chain) - 1; i > 0; i-- {
		if bc.Chain[i].PreviousHash != bc.Chain[i-1].ToHash() {
			return false
		}
	}
	return true
}
