package testutils

import (
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
)

// SumKeccak256 is used to hash input data against
// keccak256 and turn it into a [32]