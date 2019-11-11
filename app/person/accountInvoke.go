package person

import (
	"cerberus/blockchain/persaccntschannel"
	"cerberus/services/cryptography"
	"cerberus/services/ipfs"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func CreateAccount(firstName, lastName, email, phone string) ([]string, string, error) {

	if firstName == "" {
		return nil, "", errors.New("First name value cannot be an empty string")
	}

	if lastName == "" {
		return nil, "", errors.New("Last name value cannot be an empty string")
	}

	if email == "" {
		return nil, "", errors.New("Email value cannot be an empty string")
	}

	if phone == "" {
		return nil, "", errors.New("Phone value cannot be an empty string")
	}

	// create object
	firstName = strings.ToLower(firstName)
	lastName = strings.ToLower(lastName)
	email = strings.ToLower(email)

	id := bson.NewObjectId().Hex()
	publicId := cryptography.GetMD5Hash(id)

	documents := make(map[string]*documentDirectory)

	accountData := &accountData{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	accountObject := &personAccount{
		Id:          id,
		PublicId:    publicId,
		ObjectType:  "person",
		AccountData: accountData,
		Documents:   documents,
	}

	// create personAccount folder in ipfs
	//linkReference := "/personAccounts/" + email
	ipfsData, _, err := ipfs.CreateIpfsAccountDirectory(email, personAccountsIpfsHash)

	if err != nil {
		return nil, "", err
	}

	accountObject.IpfsAccountData = ipfsData

	accountObjectAsBytes, err := json.Marshal(accountObject)

	if err != nil {
		return nil, "", err
	}

	persAccntsChannelClient := persaccntschannel.CerberusClient{}
	response, newAccountData, err := persAccntsChannelClient.CreateAccount(accountObjectAsBytes)

	if err != nil {
		ipfs.DeleteDirectoryFromIpfs(ipfsData.ObjectHash, ipfsData.LinkObjectHash)

		return nil, "", err
	}

	record := &personAccount{}
	err = json.Unmarshal(newAccountData, record)

	if err != nil {
		return nil, "", err
	}

	return response, string(newAccountData), nil
}

// Update account:
/*
Selectors:
- FirstName
- LastName
- Phone
- etc
*/
func UpdateAccountBySelector(accountPublicId, selectorName, selectorValue string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Account Public Id value cannot be an empty string")
	}

	if selectorName == "" {
		return nil, nil, errors.New("SelectorName value cannot be an empty string")
	}

	if selectorValue == "" {
		return nil, nil, errors.New("SelectorValue cannot be an empty string")
	}

	selector := strings.ToLower(selectorName)

	switch selector {
	case "firstname":
		selectorName = "FirstName"

	case "lastname":
		selectorName = "LastName"

	case "email":
		selectorName = "Email"

	case "phone":
		selectorName = "Phone"
	}
	persAccntsChannelClient := persaccntschannel.CerberusClient{}
	response, newAccountData, err := persAccntsChannelClient.UpdateRecords("updateAccount", accountPublicId, []string{selectorName, strings.ToLower(selectorValue)})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(newAccountData)}, response, nil
}

func UpdateAccountFirstName(accountPublicId, firstName string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Account Public Id cannot be an empty string")
	}

	if firstName == "" {
		return nil, nil, errors.New("First name value cannot be an empty string")
	}

	dataField := "FirstName"

	persAccntsChannelClient := persaccntschannel.CerberusClient{}
	response, newAccountData, err := persAccntsChannelClient.UpdateRecords("updateAccount", accountPublicId, []string{dataField, strings.ToLower(firstName)})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(newAccountData)}, response, nil
}

func UpdateAccountLastName(accountPublicId, lastName string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Account Public Id cannot be an empty string")
	}

	if lastName == "" {
		return nil, nil, errors.New("Last name value cannot be an empty string")
	}

	dataField := "LastName"

	persAccntsChannelClient := persaccntschannel.CerberusClient{}
	response, newAccountData, err := persAccntsChannelClient.UpdateRecords("updateAccount", accountPublicId, []string{dataField, strings.ToLower(lastName)})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(newAccountData)}, response, nil
}

func UpdateAccountPhone(accountPublicId, phone string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Account Public Id cannot be an empty string")
	}

	if phone == "" {
		return nil, nil, errors.New("Phome name value cannot be an empty string")
	}

	dataField := "Phone"

	persAccntsChannelClient := persaccntschannel.CerberusClient{}
	response, newAccountData, err := persAccntsChannelClient.UpdateRecords("updateAccount", accountPublicId, []string{dataField, phone})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(newAccountData)}, response, nil
}

func UpdateAccountEmail(accountPublicId, email string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Account Public Id cannot be an empty string")
	}

	if email == "" {
		return nil, nil, errors.New("Phome name value cannot be an empty string")
	}

	dataField := "Email"

	persAccntsChannelClient := persaccntschannel.CerberusClient{}
	response, newAccountData, err := persAccntsChannelClient.UpdateRecords("updateAccount", accountPublicId, []string{dataField, email})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(newAccountData)}, response, nil
}

func DeleteAccount(accountPublicId string) ([]string, error) {

	if accountPublicId == "" {
		return nil, errors.New("Account Public Id cannot be an empty string")
	}

	persAccntsChannelClient := persaccntschannel.CerberusClient{}
	response, deletedRecord, err := persAccntsChannelClient.DeleteAccount(accountPublicId)

	if err != nil {
		return nil, err
	}

	record := &personAccount{}
	err = json.Unmarshal([]byte(deletedRecord), record)

	if err != nil {
		return nil, err
	}

	// delete records from ipfs
	ipfs.DeleteDirectoryFromIpfs(record.IpfsAccountData.ObjectHash, record.IpfsAccountData.LinkObjectHash)

	for _, directory := range record.Documents {

		_, response, err = DeleteDocument(accountPublicId, directory.DocumentData.DocumentName)

		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func CreateNewDocument(accountPublicId, documentName, holderName, countryIssue, filename string) ([]string, []string, string, error) {

	if accountPublicId == "" {
		return nil, nil, "", errors.New("Id value cannot be an empty string")
	}

	if documentName == "" {
		return nil, nil, "", errors.New("Document name value cannot be an empty string")
	}

	if holderName == "" {
		return nil, nil, "", errors.New("Holder name value cannot be an empty string")
	}

	if countryIssue == "" {
		return nil, nil, "", errors.New("Country issue value cannot be an empty string")
	}

	if filename == "" {
		return nil, nil, "", errors.New("Filename value cannot be an empty string")
	}

	documentName = strings.ToLower(documentName)
	holderName = strings.ToLower(holderName)
	countryIssue = strings.ToLower(countryIssue)

	persAccntsChannelClient := persaccntschannel.CerberusClient{}
	accountRecords, err := persAccntsChannelClient.QueryAccountData("getAccountRecords", accountPublicId)

	if err != nil {
		return nil, nil, "", err
	}

	recordUpdate := &personAccount{}
	err = json.Unmarshal([]byte(accountRecords), recordUpdate)

	if err != nil {
		return nil, nil, "", err
	}

	// check if document folder already exists in account record
	if _, ok := recordUpdate.Documents[documentName]; ok {
		return nil, nil, "", errors.New("Document with name " + documentName + " already exists. ")
	}

	// create ipfs temp directory
	newDocumentIpfsDirectory, newDocumentVersion, updatedAccountIpfsLinks, rsaLink, err := createFirstDocumentVersion(filename, documentName, recordUpdate.PublicId, recordUpdate.IpfsAccountData.ObjectHash, recordUpdate.IpfsAccountData.LinkObjectHash)

	if err != nil {
		return nil, nil, "", err
	}

	// add new document directory to record
	newDocument := &documentDirectory{
		Id:         bson.NewObjectId().Hex(),
		ObjectType: "docType",
		DocumentData: &documentData{
			DocumentName: documentName,
			Holder:       holderName,
			CountryIssue: countryIssue,
			CreatedAt:    getTime(),
		},
		IpfsDocumentDirectoryData: newDocumentIpfsDirectory,
		IpfsDocumentVersionsData:  make(map[int]*documentVersion),
	}

	// add new document version to the folder
	newDocument.IpfsDocumentVersionsData[newDocumentVersion.Name] = newDocumentVersion

	recordUpdate.Documents[documentName] = newDocument
	recordUpdate.IpfsAccountData.LinkObjectHash = updatedAccountIpfsLinks
	recordUpdate.AccountData.CreatedAt = getTime()

	recordUpdateAsBytes, err := json.Marshal(recordUpdate)

	if err != nil {
		return nil, nil, "", err
	}

	response, updatedAccount, err := persAccntsChannelClient.UpdateRecords("updateDocumentRecords", accountPublicId, []string{string(recordUpdateAsBytes)})

	if err != nil {
		return nil, nil, "", err
	}

	return []string{string(updatedAccount)}, response, rsaLink, nil
}

func CreateDocumentVersion(accountPublicId, documentName, filename, rsaLink string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Id value cannot be an empty string")
	}

	if documentName == "" {
		return nil, nil, errors.New("Document name value cannot be an empty string")
	}

	if filename == "" {
		return nil, nil, errors.New("Filename value cannot be an empty string")
	}

	if rsaLink == "" {
		return nil, nil, errors.New("Rsa link value cannot be an empty string")
	}

	documentName = strings.ToLower(documentName)

	persAccntsChannelClient := persaccntschannel.CerberusClient{}
	accountRecords, err := persAccntsChannelClient.QueryAccountData("getAccountRecords", accountPublicId)

	if err != nil {
		return nil, nil, err
	}

	recordUpdate := &personAccount{}
	err = json.Unmarshal([]byte(accountRecords), recordUpdate)

	if err != nil {
		return nil, nil, err
	}

	// check if document folder already exists in account record
	if _, ok := recordUpdate.Documents[documentName]; !ok {
		return nil, nil, errors.New("Document with name " + documentName + " does not exist")
	}

	// get existing document directory
	document := recordUpdate.Documents[documentName] // type documentDirectory

	// create document next version name
	newVersionNumber := getNextDocumentVersion(document.IpfsDocumentVersionsData)

	newDocumentVersion, updatedDirectoryLinks, err := createNewDocumentVersion(newVersionNumber, filename, rsaLink, recordUpdate.PublicId, documentName, document.IpfsDocumentDirectoryData.ObjectHash, document.IpfsDocumentDirectoryData.LinkObjectHash)

	if err != nil {
		return nil, nil, err
	}

	document.IpfsDocumentVersionsData[newDocumentVersion.Name] = newDocumentVersion
	document.IpfsDocumentDirectoryData.LinkObjectHash = updatedDirectoryLinks
	document.UpdatedAt = getTime()

	// add new version to the record
	recordUpdate.Documents[documentName].IpfsDocumentVersionsData[newDocumentVersion.Name] = newDocumentVersion

	recordUpdateAsBytes, err := json.Marshal(recordUpdate)

	if err != nil {
		return nil, nil, err
	}

	response, updatedAccount, err := persAccntsChannelClient.UpdateRecords("updateDocumentRecords", accountPublicId, []string{string(recordUpdateAsBytes)})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(updatedAccount)}, response, nil
}

func UpdateDocumentCountryIssue(accountPublicId, documentName, countryIssueUpdate string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Id value cannot be an empty string")
	}

	if documentName == "" {
		return nil, nil, errors.New("Document name value cannot be an empty string")
	}

	if countryIssueUpdate == "" {
		return nil, nil, errors.New("Country issue update value cannot be an empty string")
	}

	persAccntsChanelClient := persaccntschannel.CerberusClient{}
	accountRecords, err := persAccntsChanelClient.QueryAccountData("getAccountRecords", accountPublicId)

	if err != nil {
		return nil, nil, err
	}

	recordUpdate := &personAccount{}
	err = json.Unmarshal([]byte(accountRecords), recordUpdate)

	if err != nil {
		return nil, nil, err
	}

	if _, ok := recordUpdate.Documents[documentName]; !ok {
		return nil, nil, errors.New("Document with name " + documentName + " does not exist")
	}

	recordUpdate.Documents[documentName].DocumentData.CountryIssue = countryIssueUpdate
	recordUpdate.Documents[documentName].UpdatedAt = getTime()

	recordUpdateAsBytes, err := json.Marshal(recordUpdate)

	if err != nil {
		return nil, nil, err
	}

	response, updatedRecord, err := persAccntsChanelClient.UpdateRecords("updateDocumentRecords", accountPublicId, []string{string(recordUpdateAsBytes)})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(updatedRecord)}, response, nil
}

func UpdateDocumentHolderName(accountPublicId, documentName, personNameUpdate string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Id value cannot be an empty string")
	}

	if documentName == "" {
		return nil, nil, errors.New("Document name value cannot be an empty string")
	}

	if personNameUpdate == "" {
		return nil, nil, errors.New("Person name update value cannot be an empty string")
	}

	documentName = strings.ToLower(documentName)
	personNameUpdate = strings.ToLower(personNameUpdate)

	persAccntsChannelClient := persaccntschannel.CerberusClient{}
	accountRecords, err := persAccntsChannelClient.QueryAccountData("getAccountRecords", accountPublicId)

	if err != nil {
		return nil, nil, err
	}

	recordUpdate := &personAccount{}
	err = json.Unmarshal([]byte(accountRecords), recordUpdate)

	if err != nil {
		return nil, nil, err
	}

	if _, ok := recordUpdate.Documents[documentName]; !ok {
		return nil, nil, errors.New("Document with name " + documentName + " does not exist")
	}

	recordUpdate.Documents[documentName].DocumentData.Holder = personNameUpdate
	recordUpdate.Documents[documentName].UpdatedAt = getTime()

	recordUpdateAsBytes, err := json.Marshal(recordUpdate)

	if err != nil {
		return nil, nil, err
	}

	response, updatedRecord, err := persAccntsChannelClient.UpdateRecords("updateDocumentRecords", accountPublicId, []string{string(recordUpdateAsBytes)})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(updatedRecord)}, response, nil
}

func DeleteDocument(accountPublicId, documentName string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Account Id value cannot be an empty string")
	}

	if documentName == "" {
		return nil, nil, errors.New("Document name value cannot be an empty string")
	}

	documentName = strings.ToLower(documentName)

	persAccntsChannelClient := persaccntschannel.CerberusClient{}
	accountRecords, err := persAccntsChannelClient.QueryAccountData("getAccountRecords", accountPublicId)

	if err != nil {
		return nil, nil, err
	}

	recordUpdate := &personAccount{}
	err = json.Unmarshal([]byte(accountRecords), recordUpdate)

	if err != nil {
		return nil, nil, err
	}

	// check if document exists
	documents := recordUpdate.Documents

	if _, ok := documents[documentName]; !ok {
		return nil, nil, errors.New("Document with name: " + documentName + " does not exist.")
	}

	documentToDelete := documents[documentName]

	delete(recordUpdate.Documents, documentName)

	recordUpdateAsBytes, err := json.Marshal(recordUpdate)

	if err != nil {
		return nil, nil, err
	}

	response, updatedRecord, err := persAccntsChannelClient.UpdateRecords("updateDocumentRecords", accountPublicId, []string{string(recordUpdateAsBytes)})

	if err != nil {
		return nil, nil, err
	}

	// delete records from ipfs
	ipfs.DeleteDirectoryFromIpfs(documentToDelete.IpfsDocumentDirectoryData.ObjectHash, documentToDelete.IpfsDocumentDirectoryData.LinkObjectHash)

	for _, version := range documentToDelete.IpfsDocumentVersionsData {
		ipfs.DeleteDocumentObjectFromIpfs(version.IpfsData)

		ipfsTempDocumentPath := filepath.Join(personAccountsIpfsTempPath, accountPublicId, documentName)

		err = cryptography.DeleteCipherKeyFile(ipfsTempDocumentPath, strconv.Itoa(version.Name))

		fmt.Println(err)

		if err != nil {
			return nil, nil, err
		}

		err = cryptography.DeleteRsaDirectory(ipfsTempDocumentPath)

		if err != nil {
			return nil, nil, err
		}

		fmt.Println(err)
	}

	// delete folder from ipfs temp directory
	documentIpfsTempPath := filepath.Join(personAccountsIpfsTempPath)
	_, err = ipfs.DeleteDocumentIpfsTempDirectory(documentIpfsTempPath, accountPublicId, documentName)

	fmt.Println(err)

	return []string{string(updatedRecord)}, response, nil
}

func DeleteDocumentVersion(accountPublicId, documentName string, documentVersion int) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Account Id value cannot be an empty string")
	}

	if documentName == "" {
		return nil, nil, errors.New("Document name value cannot be an empty string")
	}

	if documentVersion < 1 {
		return nil, nil, errors.New("Document version value must be a valid version number")
	}

	documentName = strings.ToLower(documentName)

	persAccntsChannelClient := persaccntschannel.CerberusClient{}
	accountRecords, err := persAccntsChannelClient.QueryAccountData("getAccountRecords", accountPublicId)

	if err != nil {
		return nil, nil, err
	}

	recordUpdate := &personAccount{}
	err = json.Unmarshal([]byte(accountRecords), recordUpdate)

	if err != nil {
		return nil, nil, err
	}

	// check if document exists
	documents := recordUpdate.Documents

	if _, ok := documents[documentName]; !ok {
		return nil, nil, errors.New("Document with name: " + documentName + " does not exist.")
	}

	if _, ok := documents[documentName].IpfsDocumentVersionsData[documentVersion]; !ok {
		return nil, nil, errors.New("Document version " + strconv.Itoa(documentVersion) + " does not exist for " + documentName)
	}

	documentVersionToDelete := documents[documentName].IpfsDocumentVersionsData[documentVersion]

	delete(recordUpdate.Documents[documentName].IpfsDocumentVersionsData, documentVersion)
	recordUpdate.Documents[documentName].UpdatedAt = getTime()

	recordUpdateAsBytes, err := json.Marshal(recordUpdate)

	if err != nil {
		return nil, nil, err
	}

	response, updatedRecord, err := persAccntsChannelClient.UpdateRecords("updateDocumentRecords", accountPublicId, []string{string(recordUpdateAsBytes)})

	if err != nil {
		return nil, nil, err
	}

	// delete records from ipfs
	ipfs.DeleteDocumentObjectFromIpfs(documentVersionToDelete.IpfsData)

	// delete cipher key
	ipfsTempDocumentPath := filepath.Join(personAccountsIpfsTempPath, accountPublicId, documentName)
	err = cryptography.DeleteCipherKeyFile(ipfsTempDocumentPath, strconv.Itoa(documentVersion))

	if err != nil {
		return nil, nil, err
	}

	return []string{string(updatedRecord)}, response, nil
}

//
func createNewDocumentVersion(newVersion int, filename, rsaFile, accountPublicId, documentName, parentDirHash, parentDirObjectLinkHash string) (*documentVersion, string, error) {

	nextDocumentVersionString := strconv.Itoa(newVersion)

	// get Ipfs temporary document directory
	ipfsTempDocumentPath, err := ipfs.GetDocumentIpfsTempDirectory(personAccountsIpfsTempPath, accountPublicId, documentName)

	if err != nil {
		return nil, "", err
	}

	// encrypt scanned document and get data for ipfs upload
	// filename is the location of the scanned image before encryption
	rsaPath := filepath.Join(ipfsTempDocumentPath, "rsa")
	encryptedDocument, cipherKey, err := cryptography.EncryptDocument(filename, rsaPath, rsaFile)

	if err != nil {
		return nil, "", err
	}

	// save cipherKey
	// save encrypted cipher key to a file - temporary solution
	err = cryptography.SaveCipherKey(cipherKey, ipfsTempDocumentPath, nextDocumentVersionString)

	if err != nil {
		return nil, "", err
	}

	documentReference := filepath.Join(documentName, nextDocumentVersionString)

	documentVersionIpfsData, updatedDirectoryLinks, err := ipfs.UploadFileToIpfs(encryptedDocument, nextDocumentVersionString, documentReference, parentDirHash, parentDirObjectLinkHash)

	if err != nil {
		return nil, "", err
	}

	// create new document version
	newDocumentVersion := &documentVersion{
		Id:        bson.NewObjectId().Hex(),
		Name:      newVersion,
		IpfsData:  documentVersionIpfsData,
		CreatedAt: getTime(),
	}

	if err != nil {
		return nil, "", err
	}

	return newDocumentVersion, updatedDirectoryLinks, nil
}

func createFirstDocumentVersion(filename, documentName, accountPublicId, accountHash, accountObjectLinkHash string) (*ipfs.IpfsDirectoryData, *documentVersion, string, string, error) {

	// create new document directory in Ipfs network
	directoryName := documentName

	documentDirIpfsData, updatedAccountIpfsLinks, _, err := ipfs.CreateIpfsDocumentDirectory(directoryName, accountHash, accountObjectLinkHash)

	if err != nil {
		return nil, nil, "", "", err
	}

	// create document first version
	newDocumentVersionString := strconv.Itoa(1)

	// create Ipfs temporary document directory for saving rsa data on the server
	ipfsTempDocumentPath, err := ipfs.GetDocumentIpfsTempDirectory(personAccountsIpfsTempPath, accountPublicId, directoryName)

	if err != nil {
		return nil, nil, "", "", err
	}

	rsaPath := filepath.Join(ipfsTempDocumentPath, "rsa")
	rsaLink, err := cryptography.GenerateRSAKeyPair(rsaPath)

	if err != nil {
		return nil, nil, "", "", err
	}

	// filename is the location of the scanned image before encryption
	// rsaPath - filename of rsa keys
	encryptedDocument, cipherKey, err := cryptography.EncryptDocument(filename, rsaPath, rsaLink)

	if err != nil {
		return nil, nil, "", "", err
	}

	// save cipherKey
	// save encrypted cipher key to a file - temporary solution
	err = cryptography.SaveCipherKey(cipherKey, ipfsTempDocumentPath, newDocumentVersionString)

	if err != nil {
		return nil, nil, "", "", err
	}

	documentReference := filepath.Join(documentName, newDocumentVersionString)
	documentVersionIpfsData, updatedDirectoryIpfsLinks, err := ipfs.UploadFileToIpfs(encryptedDocument, newDocumentVersionString, documentReference, documentDirIpfsData.ObjectHash, documentDirIpfsData.LinkObjectHash)

	if err != nil {
		return nil, nil, "", "", err
	}

	// create document version
	documentVersion := &documentVersion{
		Id:        bson.NewObjectId().Hex(),
		Name:      1,
		IpfsData:  documentVersionIpfsData,
		CreatedAt: getTime(),
	}

	documentDirIpfsData.LinkObjectHash = updatedDirectoryIpfsLinks

	return documentDirIpfsData, documentVersion, updatedAccountIpfsLinks, rsaLink, nil
}

func getNextDocumentVersion(documentVersions map[int]*documentVersion) int {

	keys := make([]int, 0, len(documentVersions))
	for k := range documentVersions {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	highestVersion := keys[len(keys)-1]
	nextVersion := highestVersion + 1

	return nextVersion
}

func getTime() string {

	currentDateTime := time.Now()
	CurrentDateTime := currentDateTime.Format("2006-01-02 15:04:05")

	return CurrentDateTime
}
