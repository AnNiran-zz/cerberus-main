package instaccntschannel

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (instAccntsChannelClient *CerberusClient) CreateAccountDataRequest(newRequest, requestData []byte) ([]string, []byte, error) {

	// channel instance -> create
	err := instAccntsChannelClient.setupInstitutionAccountsChannelClient()

	if err != nil {
		return nil, nil, err
	}
	defer sdkInstance.Close()

	instAccntsChannelClient.channelClient, err = channel.New(instAccntsChannelClient.channelCtx)

	if err != nil {
		return nil, nil, err
	}

	// request -> prepare
	request := channel.Request{
		ChaincodeID: InstitutionAccountsChannelChainCode,
		Fcn:         "createRequest",
		Args:        [][]byte{[]byte("accountData"), newRequest, requestData},
	}

	//response, err := instAccntsChannelClient.channelClient.Execute(request)
	// or:
	response, err := instAccntsChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Request created successfully.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}

func (instAccntsChannelClient *CerberusClient) CreateDocumentDataRequest(newRequest, requestData []byte) ([]string, []byte, error) {

	// channel instance -> create
	err := instAccntsChannelClient.setupInstitutionAccountsChannelClient()

	if err != nil {
		return nil, nil, err
	}
	defer sdkInstance.Close()

	instAccntsChannelClient.channelClient, err = channel.New(instAccntsChannelClient.channelCtx)

	if err != nil {
		return nil, nil, err
	}

	// request -> prepare
	request := channel.Request{
		ChaincodeID: InstitutionAccountsChannelChainCode,
		Fcn:         "createRequest",
		Args:        [][]byte{[]byte("documentData"), newRequest, requestData},
	}

	//response, err := instAccntsChannelClient.channelClient.Query(request)
	// or:
	response, err := instAccntsChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Request created successfully.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}

func (instAccntsChannelClient *CerberusClient) AcceptRequest(requestType, requestPublicId, recipientPublicId string, acceptedData []byte) ([]string, []byte, error) {

	// channel instance -> create
	err := instAccntsChannelClient.setupInstitutionAccountsChannelClient()

	if err != nil {
		return nil, nil, err
	}
	defer sdkInstance.Close()

	instAccntsChannelClient.channelClient, err = channel.New(instAccntsChannelClient.channelCtx)

	if err != nil {
		return nil, nil, err
	}

	// request -> prepare
	request := channel.Request{
		ChaincodeID: InstitutionAccountsChannelChainCode,
		Fcn:         "acceptRequest",
		Args:        [][]byte{[]byte(requestType), []byte(requestPublicId), []byte(recipientPublicId), acceptedData},
	}

	//response, err := instAccntsChannelClient.channelClient.Execute(request)
	// or:
	response, err := instAccntsChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Request accepted successfully.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}

func (instAccntsChannelClient *CerberusClient) RejectRequest(requestType, requestPublicId, recipientPublicId string) ([]string, []byte, error) {

	// channel instance -> create
	err := instAccntsChannelClient.setupInstitutionAccountsChannelClient()

	if err != nil {
		return nil, nil, err
	}
	defer sdkInstance.Close()

	instAccntsChannelClient.channelClient, err = channel.New(instAccntsChannelClient.channelCtx)

	if err != nil {
		return nil, nil, err
	}

	// request -> prepare
	request := channel.Request{
		ChaincodeID: InstitutionAccountsChannelChainCode,
		Fcn:         "rejectRequest",
		Args:        [][]byte{[]byte(requestType), []byte(requestPublicId), []byte(recipientPublicId)},
	}

	//response, err := instAccntsChannelClient.channelClient.Query(request)
	// or:
	response, err := instAccntsChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Request accepted successfully.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}

func (instAccntsChannelClient *CerberusClient) UpdateRequest(requestType, publicId, requesterId, recipientId string, updatedData []byte) ([]string, []byte, error) {

	// channel instance -> create
	err := instAccntsChannelClient.setupInstitutionAccountsChannelClient()

	if err != nil {
		return nil, nil, err
	}
	defer sdkInstance.Close()

	instAccntsChannelClient.channelClient, err = channel.New(instAccntsChannelClient.channelCtx)
	if err != nil {
		return nil, nil, err
	}

	// request -> prepare
	request := channel.Request{
		ChaincodeID: InstitutionAccountsChannelChainCode,
		Fcn:         "updateRequest",
		Args:        [][]byte{[]byte(requestType), []byte(publicId), []byte(requesterId), []byte(recipientId), []byte(updatedData)},
	}

	//response, err := instAccntsChannelClient.channelClient.Query(request)
	// or:
	response, err := instAccntsChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Request updated successfully.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}
