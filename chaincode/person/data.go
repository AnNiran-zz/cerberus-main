package main

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
	RequestedData     string            `json:"requestedData"`
	DocumentName      string            `json:"documentName"`
	DocumentData      map[string]string `json:"documentData"`
	DocumentCopy      bool              `json:"documentCopy"`
	CreatedAt         string            `json:"createdAt"`
	UpdatedAt         string            `json:"updatedAt"`
	Status            string            `json:"status"`
}

type ipfsDocumentVersionData struct {
	ContentIdentifier        string `json:"contentIdentifier"`
	ObjectHash               string `json:"objectHash"`
	Reference                string `json:"reference"`
	ParentDirectoryHash      string `json:"parentDirectoryHash"`
	ParentDirectoryReference string `json:"parentDirectoryReference"`
}

type documentVersion struct {
	Id        string                   `json:"id"`
	Name      int                      `json:"name"`
	IpfsData  *ipfsDocumentVersionData `json:"ipfsData"`
	CipherKey string                   `json:"cipherKey"`
	CreatedAt string                   `json:"createdAt"`
	UpdateAt  string                   `json:"updatedAt"`
}

type ipfsDirectoryData struct {
	ContentIdentifier   string `json:"contentIdentifier"`
	ObjectHash          string `json:"objectHash"`
	Reference           string `json:"reference"`
	LinkObjectHash      string `json:"linkObjectHash"`
	ParentDirectoryHash string `json:"parentDirectoryHash"`
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
	Id                       string                   `json:"id"`
	ObjectType               string                   `json:"docType"`
	DocumentData             *documentData            `json:"documentData"`
	IpfsDocumentData         *ipfsDirectoryData       `json:"ipfsDocumentData"`
	IpfsDocumentVersionsData map[int]*documentVersion `json:"ipfsDocumentVersionsData"`
	UpdatedAt                string                   `json:"updatedAt"`
}

type accountData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type personAccount struct {
	Id              string                        `json:"id"`
	PublicId        string                        `json:"publicId"`
	ObjectType      string                        `json:"docType"`
	AccountData     *accountData                  `json:"accountData"`
	IpfsAccountData *ipfsDirectoryData            `json:"ipfsAccountData"`
	Documents       map[string]*documentDirectory `json:"documents"`
}
