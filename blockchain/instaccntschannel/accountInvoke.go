package instaccntschannel

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (instAccntsChannelClient *CerberusClient) CreateAccount(accountObject []byte) ([]string, []byte, error) {

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
		Fcn:         "createAccount",
		Args:        [][]byte{accountObject},
	}

	response, err := instAccntsChannelClient.channelClient.Execute(request)
	// or:
	//response, err := instAccntsChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Successfully updated records.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}

func (instAccntsChannelClient *CerberusClient) DeleteAccount(publicId string) ([]string, []byte, error) {

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
		Fcn:         "deleteAccount",
		Args:        [][]byte{[]byte(publicId)},
	}

	//response, err := instAccntsChannelClient.channelClient.Query(request)
	// or:
	response, err := instAccntsChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Successfully updated records.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}

func (instAccntsChannelClient *CerberusClient) UpdateRecords(updateType, publicId string, updateArgs []string) ([]string, []byte, error) {

	var args [][]byte

	args = append(args, []byte(updateType))
	args = append(args, []byte(publicId))

	for _, arg := range updateArgs {
		args = append(args, []byte(arg))
	}

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
		Fcn:         "updateRecords",
		Args:        args,
	}

	//response, err := instAccntsChannelClient.channelClient.Query(request)
	// or:
	response, err := instAccntsChannelClient.channelClient.Execute(request, channel.WithTargetEndpoints(AnchorPrSipher))

	if err != nil {
		return nil, nil, err
	}

	if response.ChaincodeStatus == 200 {
		fmt.Println("Successfully updated records.")
		fmt.Println("Transaction ID is: " + response.TransactionID)
	}

	return []string{"200", string(response.TransactionID)}, response.Payload, nil
}
