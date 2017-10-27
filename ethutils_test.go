package ethcontract

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

var testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

func TestParseTruffleContract(t *testing.T) {
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		crypto.PubkeyToAddress(testKey.PublicKey): {Balance: big.NewInt(10000000000)},
	})
	//	eclient, err := NewEthUtil("http://localhost:8545")
	eclient := &EClient{conn: backend, LastTranasction: &ETransact{}, auth: bind.NewKeyedTransactor(testKey)}

	truffleFile := "examples/truffle_contract/blockssh.json"
	address, err := eclient.DeployContractTruffleFromFile(truffleFile)
	if err != nil {
		t.Fatalf("Failed to deploy contract: %v", err)
	}
	fmt.Printf("Deployed truffle contract -%s to address 0x%x \n\n", truffleFile, address)
}
