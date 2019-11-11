package institution

import (
	"cerberus/blockchain/instaccntschannel"
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

func CreateAccount(organizationName, contactPerson, address, email, phone string) ([]string, []string, error) {

	if organizationName == "" {
		return nil, nil, errors.New("Organization name value cannot be an empty string")
	}

	if contactPerson == "" {
		return nil, nil, errors.New("Contact person value cannot be an empty string")
	}

	if address == "" {
		return nil, nil, errors.New("Address value cannot be an empty string")
	}

	if email == "" {
		return nil, nil, errors.New("Email value cannot be an empty string")
	}

	if phone == "" {
		return nil, nil, errors.New("Phone value cannot be an empty string")
	}

	// create object
	organizationName = strings.ToLower(organizationName)
	contactPerson = strings.ToLower(contactPerson)
	address = strings.ToLower(address)
	email = strings.ToLower(email)

	id := bson.NewObjectId().Hex()
	publicId := cryptography.GetMD5Hash(id)

	documents := make(map[string]*documentDirectory)

	accountData := &accountData{
		OrganizationName: organizationName,
		ContactPerson:    contactPerson,
		Address:          address,
		Email:            email,
		Phone:            phone,
	}

	accountObject := &institutionAccount{
		Id:          id,
		PublicId:    publicId,
		ObjectType:  "institution",
		AccountData: accountData,
		Documents:   documents,
	}

	// create personAccount folder in ipfs
	//linkReference := "/organizationAccounts/" + organizationName
	ipfsData, _, err := ipfs.CreateIpfsAccountDirectory(organizationName, institutionAccountsIpfsHash)

	if err != nil {
		return nil, nil, err
	}

	accountObject.IpfsAccountData = ipfsData

	accountObjectAsBytes, err := json.Marshal(accountObject)

	if err != nil {
		return nil, nil, err
	}

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	response, newAccountData, err := instAccntsChannelClient.CreateAccount(accountObjectAsBytes)

	if err != nil {
		ipfs.DeleteDirectoryFromIpfs(ipfsData.ObjectHash, ipfsData.LinkObjectHash)

		return nil, nil, err
	}

	record := &institutionAccount{}
	err = json.Unmarshal(newAccountData, record)

	if err != nil {
		return nil, nil, err
	}

	return response, []string{string(newAccountData)}, nil
}

// Update account:
/*
Selectors:
- OrganizationName
- Email
- Phone
- ContactPerson
- Address
*/
func UpdateAccountBySelector(accountPublicId, selectorName, selectorValue string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Account Public Id value cannot be an empty string")
	}

	if selectorName == "" {
		return nil, nil, errors.New("Selector name value cannot be an empty string")
	}

	if selectorValue == "" {
		return nil, nil, errors.New("Selector value cannot be an empty string")
	}

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	response, newAccountData, err := instAccntsChannelClient.UpdateRecords("updateAccount", accountPublicId, []string{selectorName, strings.ToLower(selectorValue)})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(newAccountData)}, response, nil
}

func UpdateAccountName(accountPublicId, name string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Id value cannot be an empty string")
	}

	if name == "" {
		return nil, nil, errors.New("Name value cannot be an empty string")
	}

	dataField := "OrganizationName"

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	response, newAccountData, err := instAccntsChannelClient.UpdateRecords("updateAccount", accountPublicId, []string{dataField, strings.ToLower(name)})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(newAccountData)}, response, nil
}

func UpdateAccountEmail(accountPublicId, email string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Id value cannot be an empty string")
	}

	if email == "" {
		return nil, nil, errors.New("Email value cannot be an empty string")
	}

	dataField := "Email"

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	response, newAccountData, err := instAccntsChannelClient.UpdateRecords("updateAccount", accountPublicId, []string{dataField, strings.ToLower(email)})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(newAccountData)}, response, nil
}

func UpdateAccountPhone(accountPublicId, phone string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Id value cannot be an empty string")
	}

	if phone == "" {
		return nil, nil, errors.New("Phone value cannot be an empty string")
	}

	dataField := "Phone"

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	response, newAccountData, err := instAccntsChannelClient.UpdateRecords("updateAccount", accountPublicId, []string{dataField, phone})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(newAccountData)}, response, nil
}

func UpdateAccountContactPerson(accountPublicId, contactPerson string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Id value cannot be an empty string")
	}

	if contactPerson == "" {
		return nil, nil, errors.New("Contact person value cannot be an empty string")
	}

	dataField := "ContactPerson"

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	response, newAccountData, err := instAccntsChannelClient.UpdateRecords("updateAccount", accountPublicId, []string{dataField, strings.ToLower(contactPerson)})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(newAccountData)}, response, nil
}

func UpdateAccountAddress(accountPublicId, address string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Id value cannot be an empty string")
	}

	if address == "" {
		return nil, nil, errors.New("Address value cannot be an empty string")
	}

	dataField := "Address"

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	response, newAccountData, err := instAccntsChannelClient.UpdateRecords("updateAccount", accountPublicId, []string{dataField, strings.ToLower(address)})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(newAccountData)}, response, nil
}

func DeleteAccount(accountPublicId string) ([]string, error) {

	if accountPublicId == "" {
		return nil, errors.New("Id value cannot be an empty string")
	}

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	response, deletedRecord, err := instAccntsChannelClient.DeleteAccount(accountPublicId)

	if err != nil {
		return nil, err
	}

	record := &institutionAccount{}
	err = json.Unmarshal([]byte(deletedRecord), record)

	if err != nil {
		return nil, err
	}

	// delete records from ipfs
	ipfs.DeleteDirectoryFromIpfs(record.IpfsAccountData.ObjectHash, record.IpfsAccountData.LinkObjectHash)

	if len(record.Documents) > 0 {
		for documentName, _ := range record.Documents {

			record, response, err := DeleteDocument(accountPublicId, documentName)

			if err != nil {
				return nil, err
			}

			fmt.Println(response)
			fmt.Println(record)
		}
	}

	return response, nil
}

func CreateNewDocument(accountPublicId, documentName, holder, countryIssue, filename string) ([]string, []string, string, error) {

	if accountPublicId == "" {
		return nil, nil, "", errors.New("Id value cannot be an empty string")
	}

	if documentName == "" {
		return nil, nil, "", errors.New("Document name value cannot be an empty string")
	}

	if holder == "" {
		return nil, nil, "", errors.New("Holder value cannot be an empty string")
	}

	if countryIssue == "" {
		return nil, nil, "", errors.New("Country issue value cannot be an empty string")
	}

	if filename == "" {
		return nil, nil, "", errors.New("Filename value cannot be an empty string")
	}

	documentName = strings.ToLower(documentName)
	holder = strings.ToLower(holder)
	countryIssue = strings.ToLower(countryIssue)

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	accountRecords, err := instAccntsChannelClient.QueryAccountData("getAccountRecords", accountPublicId)

	if err != nil {
		return nil, nil, "", err
	}

	recordUpdate := &institutionAccount{}
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
			Holder:       holder,
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

	response, updatedAccount, err := instAccntsChannelClient.UpdateRecords("updateDocumentRecords", accountPublicId, []string{string(recordUpdateAsBytes)})

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

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	accountRecords, err := instAccntsChannelClient.QueryAccountData("getAccountRecords", accountPublicId)

	if err != nil {
		return nil, nil, err
	}

	recordUpdate := &institutionAccount{}
	err = json.Unmarshal([]byte(accountRecords), recordUpdate)

	if err != nil {
		return nil, nil, err
	}

	// check if document folder already exists in account record
	if _, ok := recordUpdate.Documents[documentName]; !ok {
		return nil, nil, errors.New("Document with name " + documentName + " does not exist")
	}

	// get existing document directory
	document := recordUpdate.Documents[documentName] // type documentdirectory

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

	response, updatedAccount, err := instAccntsChannelClient.UpdateRecords("updateDocumentRecords", accountPublicId, []string{string(recordUpdateAsBytes)})

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

	documentName = strings.ToLower(documentName)
	countryIssueUpdate = strings.ToLower(countryIssueUpdate)

	instAccntsChanelClient := instaccntschannel.CerberusClient{}
	accountRecords, err := instAccntsChanelClient.QueryAccountData("getAccountRecords", accountPublicId)

	if err != nil {
		return nil, nil, err
	}

	recordUpdate := &institutionAccount{}
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

	response, updatedRecord, err := instAccntsChanelClient.UpdateRecords("updateDocumentRecords", accountPublicId, []string{string(recordUpdateAsBytes)})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(updatedRecord)}, response, nil
}

func UpdateDocumentHolderName(accountPublicId, documentName, holderNameUpdate string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Id value cannot be an empty string")
	}

	if documentName == "" {
		return nil, nil, errors.New("Document name value cannot be an empty string")
	}

	if holderNameUpdate == "" {
		return nil, nil, errors.New("Holder name update value cannot be an empty string")
	}

	documentName = strings.ToLower(documentName)
	holderNameUpdate = strings.ToLower(holderNameUpdate)

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	accountRecords, err := instAccntsChannelClient.QueryAccountData("getAccountRecords", accountPublicId)

	if err != nil {
		return nil, nil, err
	}

	recordUpdate := &institutionAccount{}
	err = json.Unmarshal([]byte(accountRecords), recordUpdate)

	if err != nil {
		return nil, nil, err
	}

	if _, ok := recordUpdate.Documents[documentName]; !ok {
		return nil, nil, errors.New("Document with name " + documentName + " does not exist")
	}

	recordUpdate.Documents[documentName].DocumentData.Holder = holderNameUpdate
	recordUpdate.Documents[documentName].UpdatedAt = getTime()

	recordUpdateAsBytes, err := json.Marshal(recordUpdate)

	if err != nil {
		return nil, nil, err
	}

	response, updatedRecord, err := instAccntsChannelClient.UpdateRecords("updateDocumentRecords", accountPublicId, []string{string(recordUpdateAsBytes)})

	if err != nil {
		return nil, nil, err
	}

	return []string{string(updatedRecord)}, response, nil
}

// test again with fully functional ipfs
func DeleteDocument(accountPublicId, documentName string) ([]string, []string, error) {

	if accountPublicId == "" {
		return nil, nil, errors.New("Account Id value cannot be an empty string")
	}

	if documentName == "" {
		return nil, nil, errors.New("Document name value cannot be an empty string")
	}

	documentName = strings.ToLower(documentName)

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	accountRecords, err := instAccntsChannelClient.QueryAccountData("getAccountRecords", accountPublicId)

	if err != nil {
		return nil, nil, err
	}

	recordUpdate := &institutionAccount{}
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

	response, updatedRecord, err := instAccntsChannelClient.UpdateRecords("updateDocumentRecords", accountPublicId, []string{string(recordUpdateAsBytes)})

	if err != nil {
		return nil, nil, err
	}

	// delete records from ipfs
	ipfs.DeleteDirectoryFromIpfs(documentToDelete.IpfsDocumentDirectoryData.ObjectHash, documentToDelete.IpfsDocumentDirectoryData.LinkObjectHash)

	for _, version := range documentToDelete.IpfsDocumentVersionsData {
		ipfs.DeleteDocumentObjectFromIpfs(version.IpfsData)

		ipfsTempDocumentPath := filepath.Join(institutionAccountsIpfsTempPath, accountPublicId, documentName)

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
	documentIpfsTempPath := filepath.Join(institutionAccountsIpfsTempPath)
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

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	accountRecords, err := instAccntsChannelClient.QueryAccountData("getAccountRecords", accountPublicId)

	if err != nil {
		return nil, nil, err
	}

	recordUpdate := &institutionAccount{}
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

	response, updatedRecord, err := instAccntsChannelClient.UpdateRecords("updateDocumentRecords", accountPublicId, []string{string(recordUpdateAsBytes)})

	if err != nil {
		return nil, nil, err
	}

	// delete records from ipfs
	ipfs.DeleteDocumentObjectFromIpfs(documentVersionToDelete.IpfsData)

	// delete cipher key
	ipfsTempDocumentPath := filepath.Join(institutionAccountsIpfsTempPath, accountPublicId, documentName)
	err = cryptography.DeleteCipherKeyFile(ipfsTempDocumentPath, strconv.Itoa(documentVersion))

	if err != nil {
		return nil, nil, err
	}

	return []string{string(updatedRecord)}, response, nil
}

func createNewDocumentVersion(newVersion int, filename, rsaFile, accountPublicId, documentName, parentDirHash, parentDirObjectLinkHash string) (*documentVersion, string, error) {

	nextDocumentVersionString := strconv.Itoa(newVersion)

	// get Ipfs temporary document directory
	ipfsTempDocumentPath, err := ipfs.GetDocumentIpfsTempDirectory(institutionAccountsIpfsTempPath, accountPublicId, documentName)

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
	ipfsTempDocumentPath, err := ipfs.GetDocumentIpfsTempDirectory(institutionAccountsIpfsTempPath, accountPublicId, documentName)

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
