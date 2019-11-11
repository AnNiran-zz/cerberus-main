package institution

import (
	"cerberus/blockchain/instaccntschannel"
	"cerberus/services/cryptography"
	"cerberus/services/ipfs"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

func GetAccountById(accountId string) (string, error) {

	if accountId == "" {
		return "", errors.New("Account Id value cannot be an empty string")
	}

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	accountData, err := instAccntsChannelClient.QueryAccountData("getAccountRecords", accountId)

	if err != nil {
		return "", err
	}

	return string(accountData), nil
}

func GetAccountsByEmail(email string) ([]string, error) {

	if email == "" {
		return nil, errors.New("Email value cannot be an empty string")
	}

	selectorKey := "email"

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	accountData, err := instAccntsChannelClient.QueryRecords(selectorKey, email)

	if err != nil {
		return nil, err
	}

	return []string{string(accountData)}, nil
}

func GetAccountsByOrgName(orgName string) ([]string, error) {

	if orgName == "" {
		return nil, errors.New("Organization name cannot be an empty string")
	}

	selectorKey := "organizationName"

	instAccntsChanelClient := instaccntschannel.CerberusClient{}
	accountData, err := instAccntsChanelClient.QueryRecords(selectorKey, orgName)

	if err != nil {
		return nil, err
	}

	return []string{string(accountData)}, nil
}

func GetAccountsByContactPerson(contactPerson string) ([]string, error) {

	if contactPerson == "" {
		return nil, errors.New("Contact person value cannot be an empty string")
	}

	selectorKey := "contactPerson"

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	accountData, err := instAccntsChannelClient.QueryRecords(selectorKey, contactPerson)

	if err != nil {
		return nil, err
	}

	return []string{string(accountData)}, nil
}

func GetAccountHistory(accountId string) ([]string, error) {

	if accountId == "" {
		return nil, errors.New("Account Id value cannot be an empty string")
	}

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	accountData, err := instAccntsChannelClient.QueryAccountData("getAccountHistory", accountId)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return []string{string(accountData)}, nil
}

// Selectors:
/*
- email
- address
- organizationName
- contactPerson
*/
func GetAccountsBySelector(selectorKey, selectorValue string) ([]string, error) {

	if selectorKey == "" {
		return nil, errors.New("Selector key value cannot be an empty string")
	}

	if selectorValue == "" {
		return nil, errors.New("Selector value cannot be an empty string")
	}

	instAccntsChanelClient := instaccntschannel.CerberusClient{}
	accountData, err := instAccntsChanelClient.QueryRecords(selectorKey, selectorValue)

	if err != nil {
		return nil, err
	}

	return []string{string(accountData)}, nil
}

func GetAccountDocument(accountId, documentName string) ([]string, error) {

	if accountId == "" {
		return nil, errors.New("Account Id value cannot be an empty string")
	}

	if documentName == "" {
		return nil, errors.New("Document name value cannot be an empty string")
	}

	documentName = strings.ToLower(documentName)

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	accountData, err := instAccntsChannelClient.QueryAccountData("getAccountRecords", accountId)

	if err != nil {
		return nil, err
	}

	record := &institutionAccount{}
	err = json.Unmarshal([]byte(accountData), record)

	if _, ok := record.Documents[documentName]; !ok {
		return nil, errors.New("Document with name " + documentName + " does not exist")
	}

	documentDataAsBytes, err := json.Marshal(record.Documents[documentName])

	if err != nil {
		return nil, err
	}

	return []string{string(documentDataAsBytes)}, nil
}

func GetAccountDocumentVersion(accountId, documentName, documentVersion string) ([]string, error) {

	if accountId == "" {
		return nil, errors.New("Account Id value cannot be an empty string")
	}

	if documentName == "" {
		return nil, errors.New("Document name value cannot be an empty value")
	}

	if documentVersion == "" {
		return nil, errors.New("Document version value cannot be an empty string")
	}

	documentName = strings.ToLower(documentName)

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	accountData, err := instAccntsChannelClient.QueryAccountData("getAccountRecords", accountId)

	if err != nil {
		return nil, err
	}

	record := &institutionAccount{}
	err = json.Unmarshal([]byte(accountData), record)

	if _, ok := record.Documents[documentName]; !ok {
		return nil, errors.New("Document with name " + documentName + " does not exist")
	}

	ver, err := strconv.Atoi(documentVersion)

	if err != nil {
		return nil, err
	}

	document := record.Documents[documentName]
	version := document.IpfsDocumentVersionsData[ver]

	if _, ok := record.Documents[documentName].IpfsDocumentVersionsData[ver]; !ok {
		return nil, errors.New("Document version " + documentVersion + " for document " + documentName + " does not exist")
	}

	// get content from ipfs
	ipfsTempDocumentPath, err := ipfs.GetDocumentIpfsTempDirectory(institutionAccountsIpfsTempPath, record.PublicId, documentName)

	if err != nil {
		return nil, err
	}

	// get cipher key
	cipherKeyPath := filepath.Join(ipfsTempDocumentPath, "cipher", strconv.Itoa(version.Name))

	cipherKey, err := cryptography.ReadCipherKey(cipherKeyPath)

	if err != nil {
		return nil, err
	}

	filename, err := ipfs.ExportFileFromIpfs(version.IpfsData.ObjectHash, strconv.Itoa(version.Name), ipfsTempDocumentPath, cipherKey)

	if err != nil {
		return nil, err
	}

	versionAsBytes, err := json.Marshal(version)

	if err != nil {
		return nil, err
	}

	return []string{string(versionAsBytes), filename}, nil
}

func GetAccountDocumentVersions(accountId, documentName string) ([]string, error) {

	if accountId == "" {
		return nil, errors.New("Account Id value cannot be an empty string")
	}

	if documentName == "" {
		return nil, errors.New("Document name value cannot be an empty string")
	}

	documentName = strings.ToLower(documentName)

	instAccntsChannelClient := instaccntschannel.CerberusClient{}
	accountData, err := instAccntsChannelClient.QueryAccountData("getAccountRecords", accountId)

	if err != nil {
		return nil, err
	}

	record := &institutionAccount{}
	err = json.Unmarshal([]byte(accountData), record)

	if _, ok := record.Documents[documentName]; !ok {
		return nil, errors.New("Document with name " + documentName + " does not exist")
	}

	// get content from ipfs
	ipfsTempDocumentPath, err := ipfs.GetDocumentIpfsTempDirectory(institutionAccountsIpfsTempPath, record.PublicId, documentName)

	if err != nil {
		return nil, err
	}

	var versions []string
	for versionNumber, version := range record.Documents[documentName].IpfsDocumentVersionsData {

		// get cipher key
		cipherKeyPath := filepath.Join(ipfsTempDocumentPath, "cipher", strconv.Itoa(versionNumber))

		cipherKey, err := cryptography.ReadCipherKey(cipherKeyPath)

		if err != nil {
			return nil, err
		}

		_, err = ipfs.ExportFileFromIpfs(version.IpfsData.ObjectHash, strconv.Itoa(version.Name), ipfsTempDocumentPath, cipherKey)

		if err != nil {
			return nil, err
		}

		versionAsBytes, err := json.Marshal(version)

		if err != nil {
			return nil, err
		}

		versions = append(versions, string(versionAsBytes))
	}

	return versions, nil
}
