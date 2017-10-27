package main

import (
	"fmt"
	"log"

	"github.com/loomnetwork/ethcontract"
)

const truffleFile = "blockssh.json"

// You need to replace keydata with your own wallet address
const keydata = `0b6abf9f0c659d64798f0231871c4b6928c02f60ad35e83404dc0df5f355541b`

func main() {
	eclient, err := ethcontract.NewEthUtil("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	eclient.SetWalletPrivateKey(keydata)

	address, err := eclient.DeployContractTruffleFromFile(truffleFile)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}
	fmt.Printf("Deployed truffle contract -%s to address 0x%x", truffleFile, address)
}
