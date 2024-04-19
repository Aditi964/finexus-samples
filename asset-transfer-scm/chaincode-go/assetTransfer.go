/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
        //"github.com/Aditi964/finexus-samples/main/asset-transfer-scm/chaincode-go/chaincode"
	"github.com/hyperledger/fabric-samples/asset-transfer-scm/chaincode-go/chaincode"
)

// Define main function for testing
func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Error creating asset-transfer-scm chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting asset-transfer-scm chaincode: %s", err.Error())
	}
}
