package instaccntschannel

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (instAccntsChannelClient *CerberusClient) QueryRecords(selectorKey, selectorValue string) (string, error) {

	// channel instance -> create
	err := instAccntsChannelClient.setupInstitutionAccountsChannelClient()

	if err != nil {
		return "", err
	}
	defer sdkInstance.Close()

	instAccntsChannelClient.channelClient, err = channel.New(instAccntsChannelClient.channelCtx)

	if err != nil {
		return "", err
	}

	// request -> prepare
	request := channel.Request{
		ChaincodeID: InstitutionAccountsChannelChainCode,
		Fcn:         "queryRecords",
		Args:        [][]byte{[]byte(selectorKey), []byte(selectorValue)},
	}

	response, err := instAccntsChannelClient.channelClient.Query(request)
	// or:
	//response, err := instAccntsChannelClient.channelClient.Query(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return "", err
	}

	if len(response.Payload) < 5 { // small random number of bytes
		fmt.Println("No records with " + selectorKey + ":" + selectorValue + " exist.")
		return "", nil
	}

	return string(response.Payload), nil
}

func (instAccntsChannelClient *CerberusClient) QueryAccountData(queryType, publicId string) (string, error) {

	// channel instance -> create
	err := instAccntsChannelClient.setupInstitutionAccountsChannelClient()

	if err != nil {
		return "", nil
	}
	defer sdkInstance.Close()

	instAccntsChannelClient.channelClient, err = channel.New(instAccntsChannelClient.channelCtx)

	if err != nil {
		return "", err
	}

	// request -> prepare
	request := channel.Request{
		ChaincodeID: InstitutionAccountsChannelChainCode,
		Fcn:         "queryAccountData",
		Args:        [][]byte{[]byte(queryType), []byte(publicId)},
	}

	response, err := instAccntsChannelClient.channelClient.Query(request)
	// or:
	//response, err := instAccntsChannelClient.channelClient.Query(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return "", err
	}

	if len(response.Payload) < 5 { // small random number of bytes
		fmt.Println("No records with id: " + publicId + " exist.")
		return "", nil
	}

	return string(response.Payload), nil
}
