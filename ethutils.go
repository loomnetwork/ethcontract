package ethcontract

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ETransact allows complex transaction info to be assesible if needed, so we dont
// need to constantly return it
type ETransact struct {
	Address  *common.Address
	Contract *bind.BoundContract
	TxHash   *common.Hash
}

type EClient struct {
	conn            *ethclient.Client
	auth            *bind.TransactOpts
	LastTranasction *ETransact
}

func NewEthUtil(dialpath string) (*EClient, error) {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := ethclient.Dial(dialpath)
	if err != nil {
		return nil, err
	}
	return &EClient{conn: conn, LastTranasction: &ETransact{}}, err
}

// SetWalletPrivateKey takes in a key string
func (e *EClient) SetWalletPrivateKey(keydata string) {
	keyBytes := common.FromHex(keydata)
	key := crypto.ToECDSAUnsafe(keyBytes)

	e.auth = bind.NewKeyedTransactor(key)
}

// DeployContractSimple deploys a simple contract with an abi and bin data. If the contract
// has contructor arguments, we will automatically generate zero values for them
func (e *EClient) DeployContractSimple(contractAbi string, contractBin string) (*common.Address, error) {
	// Deploy a new awesome contract for the binding demo
	parsed, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}
	return e.deployContractSimple(parsed, contractBin)
}

func (e *EClient) deployContractSimple(parsed abi.ABI, contractBin string) (*common.Address, error) {
	inputs := parsed.Constructor.Inputs

	var dataInputs []interface{} = make([]interface{}, 0)

	//Default all the parameters
	for _, i := range inputs {
		t := reflect.Zero(i.Type.Type)
		v := t.Interface()
		if t.Kind() == reflect.Ptr && t.IsNil() {
			elem := reflect.TypeOf(v).Elem()
			v2 := reflect.New(elem)
			v = v2.Interface()
		}
		dataInputs = append(dataInputs, v)
	}

	address, tx, contract, err := bind.DeployContract(e.auth, parsed, common.FromHex(contractBin), e.conn, dataInputs...)
	if err != nil {
		return nil, err
	}

	//	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	//	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	//	fmt.Printf("Contract Object-%v", contract)
	e.LastTranasction.Address = &address
	e.LastTranasction.Contract = contract
	h := tx.Hash()
	e.LastTranasction.TxHash = &h

	return &address, nil
}

type TruffleContract struct {
	Abi        abi.ABI `json:"abi"`
	BinaryData string  `json:"unlinked_binary"`
}

func (e *EClient) DeployContractTruffle(truffleJsonData string) (*common.Address, error) {
	var truffleContract TruffleContract
	err := json.Unmarshal([]byte(truffleJsonData), &truffleContract)
	if err != nil {
		return nil, err
	}
	return e.deployContractSimple(truffleContract.Abi, truffleContract.BinaryData)
}

func (e *EClient) DeployContractTruffleFromFile(filename string) (*common.Address, error) {
	// read file into string
	data, err := ioutil.ReadFile(filename) // just pass the file name
	if err != nil {
		return nil, err
	}
	//Pass to deploy truffle contract
	return e.DeployContractTruffle(string(data))
}
