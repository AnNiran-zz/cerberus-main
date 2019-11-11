package integration

import (
	"encoding/json"
	"fmt"
	"cerberus/services/ipfs"
	"cerberus/blockchain/iachannel"
)

type documentVersion struct {
	Name      int                   `json:"name"`
	IpfsData  *ipfs.DocumentVersion `json:"IpfsData"`
	CipherKey string                `json:"cipherKey"`
	CreatedAt string                `json:"createdAt"`
	UpdateAt  string                `json:"updatedAt"`
}

type documentDirectory struct {
	ObjectType                string                   `json:"docType"`
	DocumentName              string                   `json:"documentName"`
	PersonName                string                   `json:"personName"`
	CountryIssue              string                   `json:"countryIssue"`
	IpfsDocumentDirectoryData *ipfs.Directory          `json:"ipfsDocumentDirectoryData"`
	IpfsDocumentVersionsData  map[int]*documentVersion `json:"ipfsDocumentVersionsData"`
	CreatedAt                 string                   `json:"createdAt"`
	UpdatedAt                 string                   `json:"updatedAt"`
}

type organization struct {
	ObjectType          string                        `json:"docType"`
	OrganizationName    string                        `json:"organizationName"`
	ContactPerson       string                        `json:"contactPerson"`
	Address             string                        `json:"address"`
	Email1              string                        `json:"email1"`
	Email2              string                        `json:"email2"`
	Phone1              string                        `json:"phone1"`
	Phone2              string                        `json:"phone2"`
	IpfsAccountData     *ipfs.Directory               `json:"ipfsAccountData"`
	CreatedAt           string                        `json:"createdAt"`
	UpdatedAt           string                        `json:"updatedAt"`
	DocumentDirectories map[string]*documentDirectory `json:"documentDirectories"`
}

var organizationAccountsHash string

func CreateAccount(organizationName, contactPerson, address, email1, email2, phone1 string) ([]string, []string, error) {

	// create personAccount folder in ipfs
	linkReference := "/organizationAccounts/" + organizationName
	ipfsData, _, err := ipfs.CreateIpfsAccountDirectory(organizationName, linkReference, organizationAccountsHash)

	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	ipfsDataAsBytes, err := json.Marshal(ipfsData)

	iaChannelClient := iachannel.CerberusClient{}
	response, newAccountData, err := iaChannelClient.CreateIntegrationAccount(organizationName, contactPerson, address, email1, email2, phone1, ipfsDataAsBytes)

	if err != nil {
		fmt.Println(err)
		ipfs.DeleteDirectoryFromIpfs(ipfsData.Hash, ipfsData.LinkObjectHash)
		return nil, nil, err
	}

	record := &organization{}
	err = json.Unmarshal(newAccountData, record)
	fmt.Println(record)
	fmt.Println(record.IpfsAccountData)

	if err != nil {
		return nil, nil, err
	}

	return nil, response, nil
}

func UpdateOrganizationName() {}

func UpdateOrganizationEmail1() {}

func UpdateOrganizationEmail2() {}

func UpdateOrganizationPhone1() {}

func UpdateOrganizationPhone2() {}

func UpdateOrganizationContactPerson() {}

func UpdateOrganizationAddress() {}

func UpdateOrganizationDocumentCountryIssue() {}

func CreateOrganizationDocument() {}

func UpdateOrganizationDocumentName() {}

func DeleteOrganizationAccount() {}

func DeleteOrganizationDocument() {}

func DeleteOrganizationDocumentVersion() {}

// ...