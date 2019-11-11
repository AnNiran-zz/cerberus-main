package iachannel

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (iaChannelClient *CerberusClient) QueryAccounts(selectorKey, selectorValue string) (string, error) {

	// channel instance -> create
	err := iaChannelClient.setupIntegrationAccountsChannelClient()

	if err != nil {
		return "", nil
	}
	defer sdkInstance.Close()

	iaChannelClient.channelClient, err = channel.New(iaChannelClient.channelCtx)

	if err != nil {
		return "", err
	}

	// request prepare
	request := channel.Request{
		ChaincodeID:  IAChannelCC,
		Fcn:         "queryIntegrationAccountByName",
		Args:        [][]byte{[]byte(selectorKey), []byte(selectorValue)},
	}

	response, err := iaChannelClient.channelClient.Query(request)
	// or:
	//response, err := iaChannelClient.channelClient.Query(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return "", err
	}

	if len(response.Payload) < 5 { // small random number of bytes
		fmt.Println("No records with " + selectorKey + ":" + selectorValue + " exist.")
		return "", nil
	}

	return string(response.Payload), nil
}

func (iaChannelClient *CerberusClient) GetAccountHistory(id string) (string, error) {

	// channel instance -> create
	err := iaChannelClient.setupIntegrationAccountsChannelClient()

	if err != nil {
		return "", err
	}
	defer sdkInstance.Close()

	iaChannelClient.channelClient, err = channel.New(iaChannelClient.channelCtx)

	if err != nil {
		return "", err
	}

	// request -> prepare
	request := channel.Request{
		ChaincodeID: IAChannelCC,
		Fcn:         "getAccountHistory",
		Args:        [][]byte{[]byte(id)},
	}

	//response, err := paChannelClient.channelClient.Query(request)
	// or:
	response, err := iaChannelClient.channelClient.Query(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return "", err
	}

	if len(response.Payload) < 5 { // small random number of bytes
		fmt.Println("No records with id: " + id + " exist.")
		return "", nil
	}

	return string(response.Payload), nil
}

func (iaChannelClient *CerberusClient) GetAccountRecords(id string) (string, error) {

	// channel instance -> create
	err := iaChannelClient.setupIntegrationAccountsChannelClient()

	if err != nil {
		return "", err
	}

	defer sdkInstance.Close()

	iaChannelClient.channelClient, err = channel.New(iaChannelClient.channelCtx)

	if err != nil {
		return "", err
	}

	// request -> create
	request := channel.Request{
		ChaincodeID: IAChannelCC,
		Fcn:         "getAccountRecords",
		Args:        [][]byte{[]byte(id)},
	}

	//response, err := paChannelClient.channelClient.Query(request)
	// or:
	response, err := iaChannelClient.channelClient.Query(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return "", err
	}

	if len(response.Payload) < 5 { // small random number of bytes
		fmt.Println("No records with id: " + id + " exist.")
		return "", nil
	}

	return string(response.Payload), nil
}
