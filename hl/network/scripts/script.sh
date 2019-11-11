#!/bin/bash

echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "Build Cerberus Network end-to-end test"
echo
PERSON_ACCOUNTS_CHANNEL="$1"
INSTITUTION_ACCOUNTS_CHANNEL="$2"
INTEGRATION_ACCOUNTS_CHANNEL="$3"
DELAY="$4"
LANGUAGE="$5"
TIMEOUT="$6"
VERBOSE="$7"
: ${PERSON_ACCOUNTS_CHANNEL:="persaccntschannel"}
: ${INSTITUTION_ACCOUNTS_CHANNEL:="instaccntschannel"}
: ${INTEGRATION_ACCOUNTS_CHANNEL:="integraccntschannel"}
: ${DELAY:="20"}
: ${LANGUAGE:="golang"}
: ${TIMEOUT:="10"}
: ${VERBOSE:="false"}
LANGUAGE=`echo "$LANGUAGE" | tr [:upper:] [:lower:]`
COUNTER=1
MAX_RETRY=10

PERSON_ACCOUNTS_CC_SRC_PATH="github.com/chaincode/person/"
INSTITUTION_ACCOUNTS_CC_SRC_PATH="github.com/chaincode/institution/"
INTEGRATION_ACCOUNTS_CC_SRC_PATH="github.com/chaincode/integration/"

echo "Channels: "
echo $PERSON_ACCOUNTS_CHANNEL
echo $INSTITUTION_ACCOUNTS_CHANNEL
echo $INTEGRATION_ACCOUNTS_CHANNEL

# import utils
. scripts/utils.sh

createChannels() {

	setGlobals 0 1

	if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
                set -x
		peer channel create -o osinstance0.cerberus.dev:7050 -c $PERSON_ACCOUNTS_CHANNEL -f ./channel-artifacts/persaccntschannel.tx >&log.txt
		res=$?
                set +x
	else
				set -x
		peer channel create -o osinstance0.cerberus.dev:7050 -c $PERSON_ACCOUNTS_CHANNEL -f ./channel-artifacts/persaccntschannel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $OSINSTANCE0_CA >&log.txt
		res=$?
				set +x
	fi
	cat log.txt
	verifyResult $res "Channel creation failed"
	echo "===================== PersonAccounts Channel created ===================== "
	echo

	if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
                set -x
		peer channel create -o osinstance1.cerberus.dev:7050 -c $INSTITUTION_ACCOUNTS_CHANNEL -f ./channel-artifacts/instaccntschannel.tx >&log.txt
		res=$?
                set +x
	else
				set -x
		peer channel create -o osinstance1.cerberus.dev:7050 -c $INSTITUTION_ACCOUNTS_CHANNEL -f ./channel-artifacts/instaccntschannel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $OSINSTANCE1_CA >&log.txt
		res=$?
				set +x
	fi
	cat log.txt
	verifyResult $res "Channel creation failed"
	echo "===================== InstitutionAccounts Channel created ===================== "
	echo

	if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
                set -x
		peer channel create -o osinstance2.cerberus.dev:7050 -c $INTEGRATION_ACCOUNTS_CHANNEL -f ./channel-artifacts/integraccntschannel.tx >&log.txt
		res=$?
                set +x
	else
				set -x
		peer channel create -o osinstance2.cerberus.dev:7050 -c $INTEGRATION_ACCOUNTS_CHANNEL -f ./channel-artifacts/integraccntschannel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $OSINSTANCE2_CA >&log.txt
		res=$?
				set +x
	fi
	cat log.txt
	verifyResult $res "Channel creation failed"
	echo "===================== IntegrationAccounts Channel created ===================== "
	echo
}

joinPerAccntsChannel () {
	for org in 1 2; do
	    #for peer in 0 1 2 3 4; do
	    for peer in 0 1 2 3; do
		    joinPerAccntsChannelWithRetry $peer $org
		    sleep $DELAY
	    done
	done
}

joinInstAccntsChannel () {
	for org in 1 2; do
	    #for peer in 0 1 2 3 4; do
	    for peer in 0 1 2 3; do
		    joinInstAccntsChannelWithRetry $peer $org
		    sleep $DELAY
	    done
	done
}

joinIntegrAccntsChannel () {
	for org in 1 2; do
	    #for peer in 0 1 2 3 4; do
	    for peer in 0 1 2 3; do
		    joinIntegrAccntsChannelWithRetry $peer $org
		    sleep $DELAY
	    done
	done
}

## Create channel
echo "Creating channel..."
createChannels

## Join all the peers to the channel



echo "Having all peers join the channel..."
joinPerAccntsChannel
#joinInstAccntsChannel
#joinIntegrAccntsChannel

## Set the anchor peers for each org in the channel
echo "Updating anchor peers for Sipher..."
updateAnchorPeersForPerAccntsChannel 0 1
#updateAnchorPeersForInstAccntsChannel 0 1
#updateAnchorPeersForIntegrAccntsChannel 0 1

echo "Updating anchor peers for WhiteBox..."
updateAnchorPeersForPerAccntsChannel 0 2
#updateAnchorPeersForInstAccntsChannel 0 2
#updateAnchorPeersForIntegrAccntsChannel 0 2


# Install chaincode on anchorpr.sipher and anchorpr.whitebox
echo "Installing person accounts chaincode on anchorpr.sipher..."
installPersonAccountsChaincode 0 1

#echo "Installing institution accounts chaincode on anchorpr.sipher..."
#installInstitutionAccountsChaincode 0 1

#echo "Installing integration accounts chaincode on anchorpr.sipher..."
#installIntegrationAccountsChaincode 0 1

echo "Installing person accounts chaincode on lead0pr.sipher..."
installPersonAccountsChaincode 1 1

#echo "Installing institution accounts chaincode on lead0pr.sipher..."
#installInstitutionAccountsChaincode 1 1

#echo "Installing integration accounts chaincode on lead0pr.sipher..."
#installIntegrationAccountsChaincode 1 1

echo "Installing person accounts chaincode on lead1pr.sipher..."
installPersonAccountsChaincode 2 1

#echo "Installing institution accounts chaincode on lead1pr.sipher..."
#installInstitutionAccountsChaincode 2 1

#echo "Installing integration accounts chaincode on lead1pr.sipher..."
#installIntegrationAccountsChaincode 2 1

echo "Installing person accounts chaincode on communicatepr.sipher..."
installPersonAccountsChaincode 3 1

#echo "Installing institution accounts chaincode on communicatepr.sipher..."
#installInstitutionAccountsChaincode 3 1

#echo "Installing integration accounts chaincode on communicatepr.sipher..."
#installIntegrationAccountsChaincode 3 1


echo "Installing person accounts chaincode on anchorpr.whitebox..."
installPersonAccountsChaincode 0 2

#echo "Installing institution accounts chaincode on anchorpr.whitebox..."
#installInstitutionAccountsChaincode 0 2

#echo "Installing integration accounts chaincode on anchorpr.whitebox..."
#installIntegrationAccountsChaincode 0 2

echo "Installing person accounts chaincode on lead0pr.whitebox..."
installPersonAccountsChaincode 1 2

#echo "Installing institution accounts chaincode on lead0pr.whitebox..."
#installInstitutionAccountsChaincode 1 2

#echo "Installing integration accounts chaincode on lead0pr.whitebox..."
#installIntegrationAccountsChaincode 1 2

echo "Installing person accounts chaincode on lead1pr.whitebox..."
installPersonAccountsChaincode 2 2

#echo "Installing institution accounts chaincode on lead1pr.whitebox..."
#installInstitutionAccountsChaincode 2 2

#echo "Installing integration accounts chaincode on lead1pr.whitebox..."
#installIntegrationAccountsChaincode 2 2

echo "Installing pachaincode on communicatepr.whitebox..."
installPersonAccountsChaincode 3 2

#echo "Installing oachaincode on communicatepr.whitebox..."
#installInstitutionAccountsChaincode 3 2

#echo "Installing iachaincode on communicatepr.whitebox..."
#installIntegrationAccountsChaincode 3 2


echo "Instantiating chaincode on anchorpr.sipher..."
instantiatePersonAccountsChaincode 0 1

#instantiateInstitutionAccountsChaincode 0 1

#instantiateIntegrationAccountsChaincode 0 1

#echo "Querying chaincode on anchorpr.sipher..."
#chaincodeQuery 0 1

# Invoke chaincode on peer0.org1 and peer0.org2
#echo "Sending invoke transaction on peer0.org1 peer0.org2..."
#chaincodeInvoke 0 1 0 2

#echo "Installing chaincode on leadpr.sipher..."
#installchaincode 0 2

## Install chaincode on peer1.org2
#echo "Installing chaincode on leadpr.whitebox..."
#installChaincode 1 2

# Query on chaincode on peer1.org2, check if the result is 90
#echo "Querying chaincode on leadpr.whitebox..."
#chaincodeQuery 1 2 100

#echo "Querying chaincode on peer1.org1..."
#chaincodeQuery 1 1 90

echo
echo "========= All GOOD, Cerberus Network build execution completed =========== "
echo

echo
echo " _____   _   _   ____   "
echo "| ____| | \ | | |  _ \  "
echo "|  _|   |  \| | | | | | "
echo "| |___  | |\  | | |_| | "
echo "|_____| |_| \_| |____/  "
echo

exit 0
