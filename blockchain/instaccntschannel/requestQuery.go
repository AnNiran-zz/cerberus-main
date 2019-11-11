package instaccntschannel

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (instAccntsChannelClient *CerberusClient) QueryRequestData(idType, id string) (string, error) {

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
		Fcn:         "queryRequestData",
		Args:        [][]byte{[]byte(idType), []byte(id)},
	}

	//response, err := instAccntsChannelClient.channelClient.Query(request)
	// or:
	response, err := instAccntsChannelClient.channelClient.Query(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return "", err
	}

	return string(response.Payload), nil
}

func (instAccntsChannelClient *CerberusClient) QueryRequests(queryType, requestType, selectorKey, selectorValue string) (string, error) {

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
		Fcn:         "queryRequests",
		Args:        [][]byte{[]byte(queryType), []byte(requestType), []byte(selectorKey), []byte(selectorValue)},
	}

	//response, err := instAccntsChannelClient.channelClient.Query(request)
	// or:
	response, err := instAccntsChannelClient.channelClient.Query(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return "", err
	}

	return string(response.Payload), nil
}
