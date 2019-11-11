package institution

import (
	"cerberus/services/ipfs"
	"os"
)

type documentShareableData struct {
	documentId   string
	documentName string
	holder       string
	countryIssue string
	// some other data
}

type accountShareableData struct { // better name?
	organizationName string
	contactPerson    string
	address          string
	email            string
	phone            string
	// some other data
}

type accountDataRequest struct {
	Id                string            `json:"id"`
	PublicId          string            `json:"publicId"`
	RequestType       string            `json:"requestType"`
	ObjectType        string            `json:"docType"`
	RequesterPublicId string            `json:"requesterPublicId"`
	RecipientPublicId string            `json:"recipientPublicId"`
	RequestedData     string            `json:"requestedData"`
	AccountData       map[string]string `json:"accountData"`
	CreatedAt         string            `json:"createdAt"`
	UpdatedAt         string            `json:"updatedAt"`
	Status            string            `json:"status"`
}

type documentDataRequest struct {
	Id                string            `json:"id"`
	PublicId          string            `json:"publicId"`
	RequestType       string            `json:"requestType"`
	ObjectType        string            `json:"docType"`
	RequesterPublicId string            `json:"requesterPublicId"`
	RecipientPublicId string            `json:"recipientPublicId"`
	DocumentName      string            `json:"documentName"`
	DocumentData      map[string]string `json:"documentData"`
	DocumentCopy      bool              `json:"documentCopy"`
	CreatedAt         string            `json:"createdAt"`
	UpdatedAt         string            `json:"updatedAt"`
	Status            string            `json:"status"`
}

type documentVersion struct {
	Id        string                        `json:"id"`
	Name      int                           `json:"name"`
	IpfsData  *ipfs.IpfsDocumentVersionData `json:"ipfsData"`
	CreatedAt string                        `json:"createdAt"`
	UpdateAt  string                        `json:"updatedAt"`
}

type documentData struct {
	DocumentId   string `json:"documentId"`
	DocumentName string `json:"documentName"`
	Holder       string `json:"holder"`
	CountryIssue string `json:"countryIssue"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type documentDirectory struct {
	Id                        string                   `json:"id"`
	ObjectType                string                   `json:"docType"`
	DocumentData              *documentData            `json:"documentData"`
	IpfsDocumentDirectoryData *ipfs.IpfsDirectoryData  `json:"ipfsDocumentDirectoryData"`
	IpfsDocumentVersionsData  map[int]*documentVersion `json:"ipfsDocumentVersionsData"`
	UpdatedAt                 string                   `json:"updatedAt"`
}

type accountData struct {
	OrganizationName string `json:"organizationName"`
	ContactPerson    string `json:"contactPerson"`
	Address          string `json:"address"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}

type institutionAccount struct {
	Id              string                        `json:"id"`
	PublicId        string                        `json:"publicId"`
	ObjectType      string                        `json:"docType"`
	AccountData     *accountData                  `json:"accountData"`
	IpfsAccountData *ipfs.IpfsDirectoryData       `json:"ipfsAccountData"`
	Documents       map[string]*documentDirectory `json:"documents"`
}

var institutionAccountsIpfsHash = "QmbhFWeAGVTxzFTPwHtTWCG4NVdhjsHbX9YFfvAUNFjAzx"
var institutionAccountsIpfsTempPath = os.Getenv("GOPATH") + "/src/cerberus/ipfs/institutionAccounts"
