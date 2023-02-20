package testutils

import (
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
)

// SumKeccak256 is used to hash input data against
// keccak256 and turn it into a [32]byte suitable
// for gas efficient smart contract storage
func SumKeccak256(data []byte) [32]byte {
	hashedData := crypto.Keccak256Hash(data)
	var hash [32]byte
	for i := 0; i < len(hashedData.Bytes()); i++ {
		hash[i] = hashedData.Bytes()[i]
	}
	return hash
}

// FormatCID is used to format a CID to be
// store on a smart contract in a manner that
// allows for taking up exactly 2 slots of
// smart con