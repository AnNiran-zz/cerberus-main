package institution

import (

	//"cerberus/services/cryptography"

	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func TestInst() {

	// test institution
	//result, err := ipfs.CreateGroupAccountsIpfsDirectory("institutionAccounts")
	//fmt.Println(result)
	//fmt.Println(err)

	filename := os.Getenv("GOPATH") + "/src/cerberus/app/create-account-personal.png"
	fmt.Println(filename)

	id1 := "fb8a76b781c51d10c353d28e2b328450" // document holder
	fmt.Println(id1)

	id2 := "4fa1a1d4dd7df2fa912b201b28f2424b"
	fmt.Println(id2)

	id3 := "43147fc9dd9d5852b5aa089eec8215f8"
	fmt.Println(id3)

	//response, data, err := institution.CreateAccount("myOrg", "anna", "my address", "angelowwa@gmail.com", "123456")
	//response, err := institution.DeleteAccount(id3)

	//fmt.Println(response)
	//fmt.Println(data)
	//fmt.Println(err)

	//data, response, err := institution.UpdateAccountName(id1, "myNEWORGANIZATIONNAME")
	//data, response, err := institution.UpdateAccountEmail(id1, "newEmail")
	//
	//data, response, err := institution.UpdateAccountPhone(id1, "123")
	//data, response, err := institution.UpdateAccountContactPerson(id2, "New Contact Person name")
	//data, response, err := institution.UpdateAccountAddress(id1, "newEmailAddress")
	//data, response, err := institution.UpdateAccountBySelector(id2, "ContactPerson", "NEWCONTACTPERSON")

	//fmt.Println(data)
	//fmt.Println(response)
	//fmt.Println(err)

	//response, err := institution.DeleteAccount(id3)

	//fmt.Println(response)
	//fmt.Println(err)

	// *******************************
	//record, err := institution.GetAccountById(id2)
	//record, err := institution.GetAccountsByEmail("ANgelowwa@gmail.com")
	//record, err := institution.GetAccountsByOrgName("myNEwOrganizationName") // check this
	//record, err := institution.GetAccountsBySelector("email", "AngelowwA@gmail.com")
	//record, err := institution.GetAccountsByContactPerson("aNNa")
	//record, err := institution.GetAccountHistory(id1)

	//fmt.Println(record)
	//fmt.Println(err)

	// **************************
	//rsaLink := "/hdd/server/go/src/cerberus/ipfs/institutionAccounts/0c40f4d8baff287145e133dca563834c/mynewdocument11/rsa/rsa_key.pem"

	//data, response, rsaLink, err := institution.CreateNewDocument(id1, "myNewDocument11", "myOrg", "bulgaria", filename)
	//data, response, err := institution.CreateDocumentVersion(id1, "myNewDocUMENT11", filename, rsaLink)

	//fmt.Println(data)
	//fmt.Println(response)
	//fmt.Println(rsaLink)
	//fmt.Println(err)

	//record, response, err := institution.UpdateDocumentHolderName(id1, "myNewDocument10", "newHolderName")
	//record, response, err := institution.UpdateDocumentCountryIssue(id1, "myNewDocument15", "newCountry")
	//record, response, err := institution.DeleteDocumentVersion(id1, "myNewDocument11", 1)
	//record, response, err := institution.DeleteDocument(id1, "myNewDocument11")

	//fmt.Println(record)
	//fmt.Println(response)
	//fmt.Println(err)

	//record, err := institution.GetAccountDocument(id1, "myNewDocument11")
	//record, err := institution.GetAccountDocumentVersion(id1, "MyNEWdocument11", "1")
	//record, err := institution.GetAccountDocumentVersions(id1, "myNewDocument11")
	//record, response, err := institution.DeleteDocument(id1, "myNewDocument4")

	//fmt.Println(record)
	//fmt.Println(response)
	//fmt.Println(err)

	//account, _, err := person.UpdateAccountBySelector(id1, "FirstName", "123")
	//account, _, err := institution.UpdateAccountBySelector(id1, "OrganizationName", "new")

	//fmt.Println(account)
	//fmt.Println(err)

	// *********************************
	//id, record, err := institution.CreateAccountDataRequest(id1, id3, []string{"organizationName", "address", "email"})
	//id, record, err := institution.UpdateAccountDataRequest(id1, id3, "41c97489e66c56743fcef9ebbd65d5e7", []string{"organizationName", "email"})

	//fmt.Println(id)
	//fmt.Println(string(record))
	//fmt.Println(err)

	//response, result, err := institution.AcceptAccountDataRequest(id3, "36b7f4e494df90a13f6b3da9d81291ff", []string{"organizationName", "email"})
	//response, result, err := institution.RejectAccountDataRequest(id3, "578cee828a0818cbf67655a837525547")

	//fmt.Println(response)
	//fmt.Println(result)
	//fmt.Println(err)

	// ********************************

	//id, record, err := institution.CreateDocumentDataRequest(id3, id1, "myNewDocument11", []string{"holder", "countryIssue", "documentId"}, true)

	//fmt.Println(id)
	//fmt.Println(string(record))
	//fmt.Println(err)

	//response, result, _, err := institution.AcceptDocumentDataRequest(id1, "c8ed244fa59fe1c1cd9017864e78bc0d", "1", []string{"holder", "documentId"})
	//response, result, err := institution.RejectDocumentDataRequest(id1, "8bb0c181353b80d9e16b1119b912492b")
	//response, result, err := institution.UpdateDocumentDataRequest(id3, id1, "bd3fd8cd0ea22251277780ab57429734", "myNewDocument11", []string{"countryIssue", "holder"}, true)

	//fmt.Println(response)
	//fmt.Println(string(result))
	//fmt.Println(err)

	// ***************************************
	//data, err := institution.GetRequestsObjectsBySelector("any", "status", "pending")
	//data, err := institution.GetRequestsPublicIdsBySelector("any", "status", "accepted")

	//data, err := institution.GetRequestsByStatus("publicIds", "accountData", "rejected")
	//data, err := institution.GetRequestsByRecipient("objects", "documentData", id1)
	//data, err := institution.GetRequestsByRequester("objects", "any", id3)
	//data, err := institution.GetRequestsByDocumentName("objects", "myNEWDocument12")
	//data, err := institution.GetRequestsByStatus("objects", "any", "accepted")
	//data, err := institution.GetRequestObject("publicId", "188f46048105757dda4b5b43f2883d48")

	//fmt.Println(data)
	//fmt.Println(err)

	//result, err := person.GetRequestsPublicIdsBySelector("documentData", "documentName", "mydocument")
	//result, err := person.GetRequestPublicId("5d09fe7b2d4e0121d8e52d49", "accountData")
	//fmt.Println(result)
	//fmt.Println(err)

	//result, err := person.GetRequestsByStatus("objects", "accountData", "pending")
	//fmt.Println(result)
	//fmt.Println(err)

	//data, err := institution.GetRequestObject("publicId", "250e75e3b536d8b7d62b0afc0542e44b")
	//fmt.Println(data)
	//fmt.Println(err)

	//response, record, err := institution.RejectDocumentDataRequest("5d0665a12d4e01328ab8bf58-5d0664aa2d4e012e83652af1")
	//fmt.Println(response)
	//fmt.Println(record)
	//fmt.Println(err)

	// **************************************
}

// Person Accounts:
// QmbSSWSv6jfWFp2Cwix3Lx8Rfn39HNBHGncekEdkr31ECn

func updateFields(currentData, newData string) map[string]string {

	updateFields := make(map[string]string)
	err := json.Unmarshal([]byte(newData), &updateFields)
	if err != nil {
		fmt.Println(err)
	}

	currentDataSplit := strings.Split(currentData, "+")

	currentFields := make(map[string]string)
	for _, value := range currentDataSplit {
		currentFields[value] = ""
	}

	//newResult := make(map[string]string)
	for key, _ := range updateFields {

		if _, ok := currentFields[key]; !ok {
			currentFields[key] = ""
		}
	}

	return currentFields
}
