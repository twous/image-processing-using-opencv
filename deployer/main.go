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
	hash1 = "bafy2dzaceeylnokurfh2lhglows4kyrhua4iw2wb7h423qc66vozap26pivgr7i"
	// bke2b-264 with cidv1
	hash2           = "bafy2dzaceegxix47qjggdvh66kduvov677rncikxbgukyk47rngou733miejgmq"
	key             = `{"address":"7e4a2359c745a982a54653128085eac69e446de1","crypto":{"cipher":"aes-128-ctr"