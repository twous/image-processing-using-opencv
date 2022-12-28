package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	bindingsf "github.com/postables/cryptovendingmachine/bindings/vendorFactory"
)

var (
	// blake2b-264 with cidv1
	hash1 = "bafy2dzaceeylnoku