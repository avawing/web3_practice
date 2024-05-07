package day_two

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var PRIVATE_KEY = "6b911fd37cdf5c81d4c0adb1ab7fa822ed253ab0ad9aa18d77257c88b29b718e"

func HashMessage(message string) common.Hash {
	return crypto.Keccak256Hash([]byte(message))
}

func SignMessage(message string) ([]byte, error) {
	messageHash := HashMessage(message)
	pvtKey, _ := crypto.HexToECDSA(PRIVATE_KEY)
	return crypto.Sign(messageHash.Bytes(), pvtKey)
}

func RecoverKey(message string, signature []byte) ([]byte, error) {
	messageHash := HashMessage(message)
	return crypto.Ecrecover(messageHash.Bytes(), signature)
}
