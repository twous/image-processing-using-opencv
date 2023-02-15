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