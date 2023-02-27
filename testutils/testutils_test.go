
package testutils

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	bindingsm "github.com/postables/cryptovendingmachine/bindings/vendingMachine"
	bindingsf "github.com/postables/cryptovendingmachine/bindings/vendorFactory"
	bindingsvm "github.com/postables/cryptovendingmachine/bindings/vendorManagement"
)

var (
	// blake2b-264 with cidv1
	hash1 = "bafy2dzaceeylnokurfh2lhglows4kyrhua4iw2wb7h423qc66vozap26pivgr7i"
	// bke2b-264 with cidv1
	hash2           = "bafy2dzaceegxix47qjggdvh66kduvov677rncikxbgukyk47rngou733miejgmq"
	key             = `{"address":"7e4a2359c745a982a54653128085eac69e446de1","crypto":{"cipher":"aes-128-ctr","ciphertext":"eea2004c17292a9e94217bf53efbc31ff4ae62f3dd57f0938ab61c949a565dc1","cipherparams":{"iv":"6f6a7a89b556604940ac87ab1e78cfd1"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8088e943ac0f37c8b4d01592d8bee96468853b6f1f13ca64d201cd68e7dc7b12"},"mac":"f856d734705f35e2acf854a44eb40796518730bd835ecaec01d1f3e7a7037813"},"id":"99e2cd49-4b51-4f01-b34c-aaa0efd332c3","version":3}`
	keypass         = "password123"
	ethendpoint     = "https://rinkeby.rpc.rivet.cloud/14a054c2fd6f41a2bf26c7e875a7449f"
	ropstenendpoint = "https://ropsten.rpc.rivet.cloud/14a054c2fd6f41a2bf26c7e875a7449f"
	// ropstenendpoint = "ropsten.infura.io/v3/22894d54ea1d4101ae951b14a69bc3ec"
)

func Test_Rivet(t *testing.T) {
	if testing.Short() {
		t.Skip("short testing skipped")
	}
	client, err := ethclient.Dial(ropstenendpoint)
	if err != nil {
		t.Fatal(err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), keypass)
	if err != nil {
		log.Fatalf("Error unlocking account %v", err)
	}
	auth.Nonce = big.NewInt(100)
	fmt.Println("deploying vendor factory")
	addr, tx, _, err := bindingsf.DeployVendorfactory(auth, client)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("vendor factory deploy tx: ", tx.Hash().String())
	if _, err := bind.WaitMined(context.Background(), client, tx); err != nil {
		t.Fatal(err)
	}
	fmt.Println("vendor factory: ", addr)
}
func getHash(t *testing.T, hash string) [2][32]byte {
	cid, err := FormatCID(hash)
	if err != nil {
		t.Fatal(err)
	}
	return cid
}

func Test_SumKeccak256(t *testing.T) {
	hash := SumKeccak256([]byte(hash1))
	if len(hash) != 32 {
		t.Fatal("bad hash")
	}
}

func Test_FormatCID(t *testing.T) {
	cid := getHash(t, hash1)
	fmt.Println(len(hash1))
	if len(cid) != 2 {
		t.Fatal("bad formatted hash")
	}
}

func Test_NewBlockchain(t *testing.T) {
	_, _ = NewBlockchain(t)
}

func Test_DeployVendorManagement(t *testing.T) {
	auth, sim := NewBlockchain(t)
	_, addr := DeployVendorManagement(t, sim, auth)
	if _, err := bindingsvm.NewVendormanagement(addr, sim); err != nil {
		t.Fatal(err)
	}
}

func Test_VendorManagement(t *testing.T) {
	auth, sim := NewBlockchain(t)
	contract, _ := DeployVendorManagement(t, sim, auth)
	retID, err := contract.Id(nil)
	if err != nil {
		t.Fatal(err)
	}