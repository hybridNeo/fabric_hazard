package main

import (
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SimpleAsset struct{}

// Init
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// initState initialize the data state.
func initState(stub shim.ChaincodeStubInterface) peer.Response {
	stub.PutState("a", []byte("This is A's secret JNSDLJNSD"))
	stub.PutState("b", []byte("This is B's secret EREFKSNKS"))
	stub.PutState("readToken", []byte("0"))
	return shim.Success([]byte("State initialized"))
}

// readAorB takes a choice for either a or b and allows the user to read one of them.
// the read token must be set to 0. This means only one of them can be read.
func readAorB(choice string, stub shim.ChaincodeStubInterface) peer.Response {
	data, _ := stub.GetState("readToken") // gets value of readtoken from the eldger
	if string(data) == "0" {              // Checks if read token is set to 0
		stub.PutState("readToken", []byte("1")) // Set the readToken to 1 indicating that read is complete
		// Check the choice.
		if strings.Compare(choice, "a") == 0 {
			data, _ := stub.GetState("a")
			return shim.Success(data)
		} else {
			data, _ := stub.GetState("b")
			return shim.Success(data)
		}
	}
	return shim.Success([]byte("Read Not Allowed"))
}

// Invoke
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()
	// We test our vulnerability demonstration here
	if fn == "initialize" {
		// Initialize the state of secret and token
		initState(stub)
		return shim.Success([]byte("Initialized State"))
	} else if fn == "readAorB" {
		if len(args) >= 1 {
			readAorB(string(args[0]), stub)
		}
		return shim.Success([]byte("Invalid arguments."))
	}
	return shim.Success([]byte("Invalid endpoint"))
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
