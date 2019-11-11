package iachannel

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"gopkg.in/mgo.v2/bson"
)

func (iaChannelClient *CerberusClient) CreateAccount(organizationName, contactPerson, address, email1, email2, phone1 string, ipfsData []byte) ([]string, []byte, error) {

	// channel instance -> create
	err := iaChannelClient.setupIntegrationAccountsChannelClient()

	if err != nil {
		return nil, nil, err
	}
	defer sdkInstance.Close()

	iaChannelClient.channelClient, err = channel.New(iaChannelClient.channelCtx)

	if err != nil {
		return nil, nil, err
	}

	newAccountId := bson.NewObjectId().Hex()

	// request -> prepare
	request := channel.Request{
		ChaincodeID: IAChannelCC,
		Fcn:         "createAccount",
		Args:        [][]byte{[]byte(newAccountId), []byte(organizationName), []byte(contactPerson), []byte(address), []byte(email1), []byte(email2), []byte(phone1), ipfsData},
	}

	response, err := iaChannelClient.channelClient.Query(request)
	// or:
	//response, err := iaChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Successfully updated records.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}

func (iaChannelClient *CerberusClient) UpdateAccount(id, dataField, value string) ([]string, []byte, error) {

	// channel instance -> create
	err := iaChannelClient.setupIntegrationAccountsChannelClient()

	if err != nil {
		return nil, nil, err
	}
	defer sdkInstance.Close()

	iaChannelClient.channelClient, err = channel.New(iaChannelClient.channelCtx)

	if err != nil {
		return nil, nil, err
	}

	// request -> prepare
	request := channel.Request{
		ChaincodeID: IAChannelCC,
		Fcn:         "updateAccount",
		Args:        [][]byte{[]byte(id), []byte(dataField)},
	}

	//response, err := iaChannelClient.channelClient.Query(request)
	// or:
	response, err := iaChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Successfully updated records.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}

func (iaChannelClient *CerberusClient) DeleteAccount(id string) ([]string, []byte, error) {

	// channel instance -> create
	err := iaChannelClient.setupIntegrationAccountsChannelClient()

	if err != nil {
		return nil, nil, err
	}
	defer sdkInstance.Close()

	iaChannelClient.channelClient, err = channel.New(iaChannelClient.channelCtx)

	if err != nil {
		return nil, nil, err
	}

	// request prepare
	request := channel.Request{
		ChaincodeID: IAChannelCC,
		Fcn:         "deleteAccount",
		Args:        [][]byte{[]byte(id)},
	}

	//response, err := paChannelClient.channelClient.Query(request)
	// or:
	response, err := iaChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Successfully updated records.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}

func (iaChannelClient *CerberusClient) CreateAccountDocument(id string, recordUpdate []byte) ([]string, []byte, error) {

	// channel instance -> create
	err := iaChannelClient.setupIntegrationAccountsChannelClient()

	if err != nil {
		return nil, nil, err
	}

	defer sdkInstance.Close()

	iaChannelClient.channelClient, err = channel.New(iaChannelClient.channelCtx)

	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	// request prepare
	request := channel.Request{
		ChaincodeID: IAChannelCC,
		Fcn:         "createAccountDocument",
		Args:        [][]byte{[]byte(id), recordUpdate},
	}

	//response, err := paChannelClient.channelClient.Query(request)
	// or:
	response, err := iaChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Successfully updated records.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}

func (iaChannelClient *CerberusClient) UpdateDocumentRecords(id string, updateData []byte) ([]string, []byte, error) {

	// channel instance -> create
	err := iaChannelClient.setupIntegrationAccountsChannelClient()

	if err != nil {
		return nil, nil, err
	}
	defer sdkInstance.Close()

	iaChannelClient.channelClient, err = channel.New(iaChannelClient.channelCtx)

	if err != nil {
		return nil, nil, err
	}

	// request -> prepare
	request := channel.Request{
		ChaincodeID: IAChannelCC,
		Fcn:         "updateDocumentRecords",
		Args:        [][]byte{[]byte(id), []byte(updateData)},
	}

	//response, err := paChannelClient.channelClient.Query(request)
	// or:
	response, err := iaChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Successfully updated records.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}
