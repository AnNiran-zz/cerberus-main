package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (t *CerberusInstitutionAccounts) createAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Start Institution account initialization.")

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	accountObject := args[0]

	newAccount := &institutionAccount{}
	err := json.Unmarshal([]byte(accountObject), newAccount)

	if err != nil {
		return shim.Error(err.Error())
	}

	// check if account exists
	queryResultBytes, _, err := t.readAccount(stub, []string{newAccount.PublicId})

	if err != nil {
		return shim.Error(err.Error())
	}

	if queryResultBytes != nil {
		return shim.Error("Record with public id: " + newAccount.PublicId + " already exists.")
	}

	newAccount.AccountData.CreatedAt = getTime()

	institutionAccountAsBytes, err := json.Marshal(newAccount)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ledger invoke operation
	err = stub.PutState(newAccount.PublicId, institutionAccountAsBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end createAccount")
	return shim.Success(institutionAccountAsBytes)
}

func (t *CerberusInstitutionAccounts) updateRecords(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Initialize updateRecords")

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}

	// values -> assign
	updateFunction := args[0]
	updateArgs := args[1:]

	switch updateFunction {

	case "updateAccount":
		return t.updateAccount(stub, updateArgs)

	case "updateDocumentRecords":
		return t.updateDocumentRecords(stub, updateArgs)

	default:
		return shim.Error("Function name not found.")
	}
}

func (t *CerberusInstitutionAccounts) updateAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// assign values
	publicId := args[0]
	dataField := args[1]
	updateValue := strings.ToLower(args[2])

	// check if account exists
	queryResultBytes, _, err := t.readAccount(stub, []string{publicId})

	if err != nil {
		return shim.Error(err.Error())
	}

	if queryResultBytes == nil {
		return shim.Error("No records with provided id exist.")
	}

	// object -> get
	recordUpdate := &institutionAccount{}
	err = json.Unmarshal(queryResultBytes, recordUpdate)

	if err != nil {
		return shim.Error(err.Error())
	}

	// object -> update
	value := reflect.ValueOf(recordUpdate.AccountData).Elem().FieldByName(dataField)
	if value.IsValid() {
		value.SetString(updateValue)
	}

	recordUpdate.AccountData.UpdatedAt = getTime()

	recordUpdateAsBytes, err := json.Marshal(recordUpdate)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ledger invoke operation
	err = stub.PutState(publicId, recordUpdateAsBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end updateAccount: " + string(recordUpdateAsBytes))
	return shim.Success(recordUpdateAsBytes)
}

func (t *CerberusInstitutionAccounts) deleteAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// input sanitation
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// assign values
	publicId := args[0]

	// check if account exists
	queryResultBytes, _, err := t.readAccount(stub, []string{publicId})

	if err != nil {
		return shim.Error(err.Error())
	}

	if queryResultBytes == nil {
		return shim.Error("No records with provided id exist.")
	}

	// ledger invoke operation
	err = stub.DelState(publicId)

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end deleteAccount")
	return shim.Success(queryResultBytes)
}

func (t *CerberusInstitutionAccounts) updateDocumentRecords(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// assign values
	publicId := args[0]
	data := args[1]

	// check if account exists
	queryResultBytes, _, err := t.readAccount(stub, []string{publicId})

	if err != nil {
		return shim.Error(err.Error())
	}

	if queryResultBytes == nil {
		return shim.Error("No records with provided id exist.")
	}

	// ledger invoke operation
	err = stub.PutState(publicId, []byte(data))

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end updateDocumentRecords: " + string(data))
	return shim.Success([]byte(data))
}
