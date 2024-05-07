package test

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/assert"
	"testing"
	"web3_practice/week_one/day_two"
)

var (
	HelloWorldHex = "0x47173285a8d7341e5e972fc677286384f802f8ef42a5ec5f03bbfa254cb01fad"
	msg           = "hello world"
)

func TestHashMessage(t *testing.T) {
	t.Run("helloWorld", func(t *testing.T) {
		hash := day_two.HashMessage(msg)
		log.Debug("hash : %s", hash)

		assert.Equal(t, hash.String(), HelloWorldHex)
	})
}

func TestSignMessage(t *testing.T) {
	t.Run("should not return an error", func(t *testing.T) {
		_, err := day_two.SignMessage("hello world")

		assert.Equal(t, nil, err)
	})

	t.Run("should have been signed by the same private key", func(t *testing.T) {
		sig, err := day_two.SignMessage(msg)
		assert.Equal(t, nil, err)

		messageHash := day_two.HashMessage(msg)
		assert.Equal(t, HelloWorldHex, messageHash.String())

		pvtKey, _ := crypto.HexToECDSA(day_two.PRIVATE_KEY)
		publicKey := pvtKey.Public()
		publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

		recoveredKey, err := crypto.Ecrecover(messageHash.Bytes(), sig)
		assert.Equal(t, nil, err)
		assert.Equal(t, recoveredKey, publicKeyBytes)

	})
}

func TestRecoverMessage(t *testing.T) {
	t.Run("should recover the public key from a signed message", func(t *testing.T) {
		sig, err := day_two.SignMessage(msg)
		assert.Equal(t, nil, err)

		pvtKey, _ := crypto.HexToECDSA(day_two.PRIVATE_KEY)
		publicKey := pvtKey.Public()
		publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

		recoveredKey, err := day_two.RecoverKey(msg, sig)
		assert.Equal(t, nil, err)

		assert.Equal(t, recoveredKey, publicKeyBytes)
	})
}

func TestGetAddress(t *testing.T) {
	var (
		privateKey       = "6b911fd37cdf5c81d4c0adb1ab7fa822ed253ab0ad9aa18d77257c88b29b718e"
		EXPECTED_ADDRESS = "0x16bB6031CBF3a12B899aB99D96B64b7bbD719705"
	)
	t.Run("Get Address", func(t *testing.T) {
		pvtKey, _ := crypto.HexToECDSA(privateKey)
		publicKey := pvtKey.Public()
		publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
		//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

		addr := day_two.GetAddress(*publicKeyECDSA)

		assert.Equal(t, addr.String(), EXPECTED_ADDRESS)
	})
}
