# EthContract Go

A set utilities for golang that help you manage smart contracts on Ethereuem with GO. These use the standard go-ethereum library, but give niceties for developers building native apps interacting with the blockchain. Including helper methods to deal with Truffle compiled contracts. 

See examples folder for more samples. View the docs at readthedocs.


## Example Simple tokem contract
see examples/simple_token/main.go

```
// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "Insert Here"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `Insert here```

// You need to replace keydata with your own wallet address
const keydata = `3958dcf7ffda44ed0540ab65d05999a56671b10137f2d1fe551c07331b740f70`

func main() {
	eclient, err := ethutils.NewEthUtil("http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	eclient.SetWalletPrivateKey(keydata)

	address, err := eclient.DeployContractSimple(TokenABI, TokenBin)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}
	fmt.Printf("Deployed contract to address 0x%x", address)
}



## Example Deploy truffle contract
see examples/simple_token/main.go

```
const truffleFile = "blockssh.json"

// You need to replace keydata with your own wallet address
const keydata = `3958dcf7ffda44ed0540ab65d05999a56671b10137f2d1fe551c07331b740f70`

func main() {
	eclient, err := ethutils.NewEthUtil("http://localhost:8545")
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

```
