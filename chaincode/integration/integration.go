package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type CerberusIntegrationAccounts struct {}


type IpfsDocumentVersionData struct {
	ContentIdentifier        string `json:"contentIdentifier"`
	ObjectHash               string `json:"objectHash"`
	Reference                string `json:"reference"`
	ParentDirectoryHash      string `json:"parentDirectoryHash"`
	ParentDirectoryReference string `json:"parentDirectoryReference"`
}

type documentVersion struct {
	Id        string                   `json:"id"`
	Name      int                      `json:"name"`
	IpfsData  *IpfsDocumentVersionData `json:"IpfsData"`
	CipherKey string                   `json:"cipherKey"`
	CreatedAt string                   `json:"createdAt"`
	UpdateAt  string                   `json:"updatedAt"`
}

type IpfsDirectoryData struct {
	ContentIdentifier   string `json:"contentIdentifier"`
	ObjectHash          string `json:"objectHash"`
	Reference           string `json:"reference"`
	LinkObjectHash      string `json:"linkObjectHash"`
	ParentDirectoryHash string `json:"parentDirectoryHash"`
}

type documentDirectory struct {
	ObjectType                string                   `json:"docType"`
	DocumentName              string                   `json:"documentName"`
	PersonName                string                   `json:"personName"`
	CountryIssue              string                   `json:"countryIssue"`
	IpfsDocumentData          *IpfsDirectoryData       `json:"ipfsDocumentData"`
	IpfsDocumentVersionsData  map[int]*documentVersion `json:"ipfsDocumentVersionsData"`
	CreatedAt                 string                   `json:"createdAt"`
	UpdatedAt                 string                   `json:"updatedAt"`
}

type integration struct {
	Id               string                        `json:"id"`
	ObjectType       string                        `json:"docType"`
	OrganizationName string                        `json:"organizationName"`
	ContactPerson    string                        `json:"contactPerson"`
	Address          string                        `json:"address"`
	Email1           string                        `json:"email1"`
	Email2           string                        `json:"email2"`
	Phone1           string                        `json:"phone1"`
	Phone2           string                        `json:"phone2"`
	IpfsAccountData  *IpfsDirectoryData            `json:"ipfsAccountData"`
	CreatedAt        string                        `json:"createdAt"`
	UpdatedAt        string                        `json:"updatedAt"`
	Documents        map[string]*documentDirectory `json:"documentDirectories"`
}

func (t *CerberusIntegrationAccounts) Init(stub shim.ChaincodeStubInterface) pb.Response {

	// initialize chaincode
	fmt.Println("Cerberus chaincode instantiation.")

	_, args := stub.GetFunctionAndParameters()

	t.createAccount(stub, args)
	return shim.Success(nil)
}

func (t *CerberusIntegrationAccounts) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()
	fmt.Println("Invoke is running " + function)

	// Handle different functions
	switch function {

	case "createAccount":
		return t.createAccount(stub, args)

	case "queryAccounts":
		return t.queryAccounts(stub, args)

	case "getAccountHistory":
		return t.getAccountHistory(stub, args)

	case "updateAccount":
		return t.updateAccount(stub, args)

	case "deleteAccount":
		return t.deleteAccount(stub, args) // Ok

	case "getAccountRecords":
		return t.getAccountRecords(stub, args)

	case "updateDocumentRecords":
		return t.updateDocumentRecords(stub, args)

	default:
		return shim.Error("Function name not found.")
	}

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(CerberusIntegrationAccounts))
	if err != nil {
		fmt.Println("Error starting Integration chaincode: %s", err.Error())
	}
}

func (t *CerberusIntegrationAccounts) createAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Start Integration account initialization.")

	if len(args) < 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8.")
	}

	// Input sanitation
	fmt.Println("Start Integration account initialization.")

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}

	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}

	if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}

	if len(args[5]) <= 0 {
		return shim.Error("6th argument must be a non-empty string")
	}

	if len(args[6]) <= 0 {
		return shim.Error("7th argument must be a non-empty string")
	}

	if len(args[7]) <= 0 {
		return shim.Error("8th argument must be a non-empty string")
	}

	// values -> assign
	id                := strings.ToLower(args[0])
	organizationName  := strings.ToLower(args[0])
	contactPerson     := strings.ToLower(args[1])
	address           := string(args[2])
	email1            := strings.ToLower(args[3])
	email2            := strings.ToLower(args[4])
	phone1            := strings.ToLower(args[5])
	ipfsAccountData   := args[6]

	// check if account exists
	queryResult := t.readAccount(stub, []string{id})

	if queryResult.Payload != nil {
		return shim.Error("Record with id: " + id + " already exists.")
	}

	// check if organization is registered
	record := &integration{}
	err := json.Unmarshal(queryResult.Payload, record)

	if err != nil {
		return shim.Error(err.Error())
	}

	if record.OrganizationName == organizationName {
		return shim.Error("Record with organization name: " + record.OrganizationName + " already exists.")

	}

	// object -> create
	ipfsData := &IpfsDirectoryData{}
	err = json.Unmarshal([]byte(ipfsAccountData), ipfsData)

	if err != nil {
		return shim.Error(err.Error())
	}

	documents := make(map[string]*documentDirectory)

	organization := &integration{
		Id:               id,
		ObjectType:       "docType",
		OrganizationName: organizationName,
		ContactPerson:    contactPerson,
		Address:          address,
		Email1:           email1,
		Email2:           email2,
		Phone1:           phone1,
		IpfsAccountData:  ipfsData,
		CreatedAt:        getTime(),
		Documents:        documents,
	}

	organizationasBytes, err := json.Marshal(organization)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ledger invoke operation
	err = stub.PutState(id, organizationasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end createAccount")
	return shim.Success(nil)
}

func (t *CerberusIntegrationAccounts) queryAccounts(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2.")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 1 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	// assign values
	selectorKey   := args[0]
	selectorValue := strings.ToLower(args[1])

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"docType\",\"%s\":\"%s\"}}", selectorKey, selectorValue)

	// obtain records
	queryResults, err := getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end queryAccounts: " + string(queryResults))
	return shim.Success(queryResults)
}

func (t *CerberusIntegrationAccounts) getAccountHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// assign values
	id := strings.ToLower(args[0])

	// obtain records
	resultsIterator, err := stub.GetHistoryForKey(id)

	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the account
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON account)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getIntegrationAccountHistory returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

func (t *CerberusIntegrationAccounts) updateAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

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
	id          := strings.ToLower(args[0])
	dataField   := args[1]
	updateValue := strings.ToLower(args[2])

	// check if account exists
	queryResult  := t.readAccount(stub, []string{id})

	if queryResult.Payload == nil {
		return shim.Error("No records with provided if exist.")
	}

	// object -> get
	recordUpdate := &integration{}
	err := json.Unmarshal([]byte(queryResult.Payload), recordUpdate)

	if err != nil {
		return shim.Error(err.Error())
	}

	// object -> update
	value := reflect.ValueOf(recordUpdate).Elem().FieldByName(dataField)
	if value.IsValid() {
		value.SetString(updateValue)
	}

	recordUpdate.UpdatedAt = getTime()

	recordUpdateAsBytes, err := json.Marshal(recordUpdate)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ledger invoke operation
	err = stub.PutState(id, recordUpdateAsBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end updateAccount: " + string(recordUpdateAsBytes))
	return shim.Success(recordUpdateAsBytes)
}

func (t *CerberusIntegrationAccounts) deleteAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// input sanitation
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// assign values
	id := strings.ToLower(args[0])

	// check if account exists
	queryResult := t.readAccount(stub, []string{id})

	if queryResult.Payload == nil {
		return shim.Error("No records with provided id exist.")
	}

	// ledger invoke operation
	err := stub.DelState(id)

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end deleteAccount")
	return shim.Success(nil)
}

func (t *CerberusIntegrationAccounts) readAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var result []byte
	var id     string

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1.")
	}

	id           = strings.ToLower(args[0])
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"docType\",\"id\":\"%s\"}}", id)

	// obtain records
	resultsIterator, err := stub.GetQueryResult(queryString)

	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()

		if err != nil {
			return shim.Error(err.Error())
		}

		result = response.Value
	}

	return shim.Success(result)
}

func (t *CerberusIntegrationAccounts) getAccountRecords(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// input sanitation
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	// assign values
	id := strings.ToLower(args[0])

	// check if account exists
	queryResult  := t.readAccount(stub, []string{id})

	if queryResult.Payload == nil {
		return shim.Error("No records with provided if exist.")
	}

	fmt.Println("- end getAccountRecords: " + string(queryResult.Payload))
	return shim.Success(nil)
}

func (t *CerberusIntegrationAccounts) updateDocumentRecords(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	// input sanitation
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	// assign values
	id   := strings.ToLower(args[0])
	data := args[1]

	// check if account exists
	queryResult  := t.readAccount(stub, []string{id})

	if queryResult.Payload == nil {
		return shim.Error("No records with provided if exist.")
	}

	// ledger invoke operation
	err := stub.PutState(id, []byte(data))

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end updateDocumentRecords: " + string(data))
	return shim.Success(nil)
}

func getTime() string {

	currentDateTime := time.Now()
	CurrentDateTime := currentDateTime.Format("2006-01-02 15:04:05")

	return CurrentDateTime
}