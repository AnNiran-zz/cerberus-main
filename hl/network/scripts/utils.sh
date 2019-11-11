#!/usr/bin/env bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a collection of bash functions used by different scripts

ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/cerberus.dev/orderers/orderer.cerberus.dev/msp/tlscacerts/tlsca.cerberus.dev-cert.pem

OSINSTANCE0_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/cerberus.dev/orderers/osinstance0.cerberus.dev/msp/tlscacerts/tlsca.cerberus.dev-cert.pem
OSINSTANCE1_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/cerberus.dev/orderers/osinstance1.cerberus.dev/msp/tlscacerts/tlsca.cerberus.dev-cert.pem
OSINSTANCE2_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/cerberus.dev/orderers/osinstance2.cerberus.dev/msp/tlscacerts/tlsca.cerberus.dev-cert.pem
OSINSTANCE3_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/cerberus.dev/orderers/osinstance3.cerberus.dev/msp/tlscacerts/tlsca.cerberus.dev-cert.pem
OSINSTANCE4_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/cerberus.dev/orderers/osinstance4.cerberus.dev/msp/tlscacerts/tlsca.cerberus.dev-cert.pem

PEER0_ORG1_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/sipher.cerberus.dev/peers/anchorpr.sipher.cerberus.dev/tls/ca.crt
PEER0_ORG2_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/whitebox.cerberus.dev/peers/anchorpr.whitebox.cerberus.dev/tls/ca.crt


# verify the result of the end-to-end test
verifyResult() {
  if [ $1 -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "========= ERROR !!! FAILED to execute End-2-End Scenario ==========="
    echo
    exit 1
  fi
}

# Set OrdererOrg.Admin globals
setOrdererGlobals() {
  CORE_PEER_LOCALMSPID="OrdererMSP"
  CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/cerberus.dev/orderers/orderer.cerberus.dev/msp/tlscacerts/tlsca.cerberus.dev-cert.pem
  CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/cerberus.dev/users/Admin@cerberus.dev/msp
}

setGlobals() {
  PEER=$1
  ORG=$2
  if [ $ORG -eq 1 ]; then
    CORE_PEER_LOCALMSPID="SipherMSP"
    ORGANIZATION_NAME="sipher"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/sipher.cerberus.dev/users/Admin@sipher.cerberus.dev/msp
    if [ $PEER -eq 0 ]; then
      PEER_NAME="anchorpr"
      CORE_PEER_ADDRESS=anchorpr.sipher.cerberus.dev:7051
    elif [ $PEER -eq 1 ]; then
      PEER_NAME="lead0pr"
      CORE_PEER_ADDRESS=lead0pr.sipher.cerberus.dev:7051
    elif [ $PEER -eq 2 ]; then
      PEER_NAME="lead1pr"
      CORE_PEER_ADDRESS=lead1pr.sipher.cerberus.dev:7051
    elif [ $PEER -eq 3 ]; then
      PEER_NAME="communicatepr"
      CORE_PEER_ADDRESS=communicatepr.sipher.cerberus.dev:7051
    elif [ $PEER -eq 4 ]; then
      PEER_NAME="execute0pr"
      CORE_PEER_ADDRESS=execute0pr.sipher.cerberus.dev:7051
    elif [ $PEER -eq 5 ]; then
      PEER_NAME="execute1pr"
      CORE_PEER_ADDRESS=execute1pr.sipher.cerberus.dev:7051
    else
      PEER_NAME="fallback0pr"
      CORE_PEER_ADDRESS=fallback0pr.sipher.cerberus.dev:7051
    fi

  elif [ $ORG -eq 2 ]; then
    CORE_PEER_LOCALMSPID="WhiteBoxMSP"
    ORGANIZATION_NAME="whitebox"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/whitebox.cerberus.dev/users/Admin@whitebox.cerberus.dev/msp
    if [ $PEER -eq 0 ]; then
      PEER_NAME="anchorpr"
      CORE_PEER_ADDRESS=anchorpr.whitebox.cerberus.dev:7051
    elif [ $PEER -eq 1 ]; then
      PEER_NAME="lead0pr"
      CORE_PEER_ADDRESS=lead0pr.whitebox.cerberus.dev:7051
    elif [ $PEER -eq 2 ]; then
      PEER_NAME="lead1pr"
      CORE_PEER_ADDRESS=lead1pr.whitebox.cerberus.dev:7051
    elif [ $PEER -eq 3 ]; then
      PEER_NAME="communicatepr"
      CORE_PEER_ADDRESS=communicatepr.whitebox.cerberus.dev:7051
    elif [ $PEER -eq 4 ]; then
      PEER_NAME="execute0pr"
      CORE_PEER_ADDRESS=execute0pr.whitebox.cerberus.dev:7051
    elif [ $PEER -eq 5 ]; then
      PEER_NAME="execute1pr"
      CORE_PEER_ADDRESS=execute1pr.whitebox.cerberus.dev:7051
    else
      PEER_NAME="fallback0pr"
      CORE_PEER_ADDRESS=fallback0pr.whitebox.cerberus.dev:7051
    fi

  else
    echo "================== ERROR !!! ORG Unknown =================="
  fi

  if [ "$VERBOSE" == "true" ]; then
    env | grep CORE
  fi
}

setOrganizationGlobals() {
  ORG=$1
  if [ $ORG -eq 1 ]; then
    ORGANIZATION_NAME="sipher"
    CORE_PEER_LOCALMSPID="SipherMSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/sipher.cerberus.dev/users/Admin@sipher.cerberus.dev/msp
  elif [ $ORG -eq 2 ]; then
    ORGANIZATION_NAME="whitebox"
    CORE_PEER_LOCALMSPID="WhiteBoxMSP"
    CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/whitebox.cerberus.dev/users/Admin@whitebox.cerberus.dev/msp
  else
    echo "================== ERROR !!! ORG Unknown =================="
  fi
}

updateAnchorPeersForPerAccntsChannel() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG

  if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
    set -x
    peer channel update -o osinstance0.cerberus.dev:7050 -c $PERSON_ACCOUNTS_CHANNEL -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}persaccntschannelAnchors.tx >&log.txt
    res=$?
    set +x
  else
    set -x
    peer channel update -o osinstance0.cerberus.dev:7050 -c $PERSON_ACCOUNTS_CHANNEL -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}persaccntschannelAnchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $OSINSTANCE0_CA >&log.txt
    res=$?
    set +x
  fi
  cat log.txt
  verifyResult $res "Anchor peer update failed"
  echo "===================== Anchor peers updated for org '$CORE_PEER_LOCALMSPID' on channel '$PERSON_ACCOUNTS_CHANNEL' ===================== "
  sleep $DELAY
  echo
}

updateAnchorPeersForInstAccntsChannel() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG

  if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
    set -x
    peer channel update -o osinstance1.cerberus.dev:7050 -c $INSTITUTION_ACCOUNTS_CHANNEL -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}instaccntschannelAnchors.tx >&log.txt
    res=$?
    set +x
  else
    set -x
    peer channel update -o osinstance1.cerberus.dev:7050 -c $INSTITUTION_ACCOUNTS_CHANNEL -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}instaccntschannelAnchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $OSINSTANCE1_CA >&log.txt
    res=$?
    set +x
  fi
  cat log.txt
  verifyResult $res "Anchor peer update failed"
  echo "===================== Anchor peers updated for org '$CORE_PEER_LOCALMSPID' on channel '$INSTITUTION_ACCOUNTS_CHANNEL' ===================== "
  sleep $DELAY
  echo
}

updateAnchorPeersForIntegrAccntsChannel() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG

  if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
    set -x
    peer channel update -o osinstance2.cerberus.dev:7050 -c $INTEGRATION_ACCOUNTS_CHANNEL -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}integraccntschannelAnchors.tx >&log.txt
    res=$?
    set +x
  else
    set -x
    peer channel update -o osinstance2.cerberus.dev:7050 -c $INTEGRATION_ACCOUNTS_CHANNEL -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}integraccntschannelAnchors.tx --tls $CORE_PEER_TLS_ENABLED --cafile $OSINSTANCE2_CA >&log.txt
    res=$?
    set +x
  fi
  cat log.txt
  verifyResult $res "Anchor peer update failed"
  echo "===================== Anchor peers updated for org '$CORE_PEER_LOCALMSPID' on channel '$INTEGRATION_ACCOUNTS_CHANNEL' ===================== "
  sleep $DELAY
  echo
}

## Sometimes Join takes time hence RETRY at least 5 times
joinPerAccntsChannelWithRetry() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG

  set -x
  peer channel join -b $PERSON_ACCOUNTS_CHANNEL.block >&log.txt
  res=$?
  set +x
  cat log.txt
  if [ $res -ne 0 -a $COUNTER -lt $MAX_RETRY ]; then
    COUNTER=$(expr $COUNTER + 1)
    echo "${PEER_NAME}.${ORGANIZATION_NAME} failed to join the channel, Retry after $DELAY seconds"
    sleep $DELAY
    joinPerAccntsChannelWithRetry $PEER $ORG
  else
    COUNTER=1
  fi
  verifyResult $res "After $MAX_RETRY attempts, ${PEER_NAME}.${ORGANIZATION_NAME} has failed to join channel '$PERSON_ACCOUNTS_CHANNEL' "
  echo "===================== ${PEER_NAME}.${ORGANIZATION_NAME} joined channel '$PERSON_ACCOUNTS_CHANNEL' ===================== "
}

joinInstAccntsChannelWithRetry() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG

  set -x
  peer channel join -b $INSTITUTION_ACCOUNTS_CHANNEL.block >&log.txt
  res=$?
  set +x
  cat log.txt
  if [ $res -ne 0 -a $COUNTER -lt $MAX_RETRY ]; then
    COUNTER=$(expr $COUNTER + 1)
    echo "${PEER_NAME}.${ORGANIZATION_NAME} failed to join the channel, Retry after $DELAY seconds"
    sleep $DELAY
    joinInstAccntsChannelWithRetry $PEER $ORG
  else
    COUNTER=1
  fi
  verifyResult $res "After $MAX_RETRY attempts, ${PEER_NAME}.${ORGANIZATION_NAME} has failed to join channel '$INSTITUTION_ACCOUNTS_CHANNEL' "
  echo "===================== ${PEER_NAME}.${ORGANIZATION_NAME} joined channel '$INSTITUTION_ACCOUNTS_CHANNEL' ===================== "
}

joinIntegrAccntsChannelWithRetry() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG

  set -x
  peer channel join -b $INTEGRATION_ACCOUNTS_CHANNEL.block >&log.txt
  res=$?
  set +x
  cat log.txt
  if [ $res -ne 0 -a $COUNTER -lt $MAX_RETRY ]; then
    COUNTER=$(expr $COUNTER + 1)
    echo "${PEER_NAME}.${ORGANIZATION_NAME} failed to join the channel, Retry after $DELAY seconds"
    sleep $DELAY
    joinIntegrAccntsChannelWithRetry $PEER $ORG
  else
    COUNTER=1
  fi
  verifyResult $res "After $MAX_RETRY attempts, ${PEER_NAME}.${ORGANIZATION_NAME} has failed to join channel '$INTEGRATION_ACCOUNTS_CHANNEL' "
  echo "===================== ${PEER_NAME}.${ORGANIZATION_NAME} joined channel '$INTEGRATION_ACCOUNTS_CHANNEL' ===================== "

}

installPersonAccountsChaincode() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG
  VERSION=${3:-1.0}
  set -x
  peer chaincode install -n persaccntschannelcc -v ${VERSION} -l ${LANGUAGE} -p ${PERSON_ACCOUNTS_CC_SRC_PATH} >&log.txt
  res=$?
  set +x
  cat log.txt
  verifyResult $res "Chaincode installation on ${PEER_NAME}.${ORGANIZATION_NAME} has failed"
  echo "===================== Chaincode is installed on ${PEER_NAME}.${ORGANIZATION_NAME} ===================== "
  echo
}

installInstitutionAccountsChaincode() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG
  VERSION=${3:-1.0}
  set -x
  peer chaincode install -n instaccntschannelcc -v ${VERSION} -l ${LANGUAGE} -p ${INSTITUTION_ACCOUNTS_CC_SRC_PATH} >&log.txt
  res=$?
  set +x
  cat log.txt
  verifyResult $res "Chaincode installation on ${PEER_NAME}.${ORGANIZATION_NAME} has failed"
  echo "===================== Chaincode is installed on ${PEER_NAME}.${ORGANIZATION_NAME} ===================== "
  echo
}

installIntegrationAccountsChaincode() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG
  VERSION=${3:-1.0}
  set -x
  peer chaincode install -n integraccntschannelcc -v ${VERSION} -l ${LANGUAGE} -p ${INTEGRATION_ACCOUNTS_CC_SRC_PATH} >&log.txt
  res=$?
  set +x
  cat log.txt
  verifyResult $res "Chaincode installation on ${PEER_NAME}.${ORGANIZATION_NAME} has failed"
  echo "===================== Chaincode is installed on ${PEER_NAME}.${ORGANIZATION_NAME} ===================== "
  echo
}

instantiatePersonAccountsChaincode() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG
  VERSION=${3:-1.0}

  # while 'peer chaincode' command can get the orderer endpoint from the peer
  # (if join was successful), let's supply it directly as we know it using
  # the "-o" option
  if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
    set -x
    peer chaincode instantiate -o osinstance0.cerberus.dev:7050 -C $PERSON_ACCOUNTS_CHANNEL -n persaccntschannelcc -l ${LANGUAGE} -v ${VERSION} -c '{"Args":["init", "123456", "anna", "angelova", "angelowwa@gmail.com", "0877150173", "someData"]}' -P "OR ('SipherMSP.peer','WhiteBoxMSP.peer')" >&log.txt
    res=$?
    set +x
  else
    set -x
    peer chaincode instantiate -o osinstance0.cerberus.dev:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $OSINSTANCE0_CA -C $PERSON_ACCOUNTS_CHANNEL -n persaccntschannelcc -l ${LANGUAGE} -v ${VERSION} -c '{"Args":["init", "123456", "anna", "angelova", "angelowwa@gmail.com", "0877150173", "some data"]}' -P "OR ('SipherMSP.peer','WhiteBoxMSP.peer')" >&log.txt
    res=$?
    set +x
  fi
  cat log.txt
  verifyResult $res "Chaincode instantiation on ${PEER_NAME}.${ORGANIZATION_NAME} on channel '$PERSON_ACCOUNTS_CHANNEL' failed"
  echo "===================== Chaincode is instantiated on ${PEER_NAME}.${ORGANIZATION_NAME} on channel '$PERSON_ACCOUNTS_CHANNEL' ===================== "
  echo
}

instantiateInstitutionAccountsChaincode() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG

  # while 'peer chaincode' command can get the orderer endpoint from the peer
  # (if join was successful), let's supply it directly as we know it using
  # the "-o" option
  if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
    set -x
    peer chaincode instantiate -o osinstance1.cerberus.dev:7050 -C $INSTITUTION_ACCOUNTS_CHANNEL -n instaccntschannelcc -l ${LANGUAGE} -v ${VERSION} -c '{"Args":["init", "myOrganization", "Anna", "myAddress", "angelowwa@gmail.com", "angelowwa@gmail.com", "0877150173"]}' -P "OR ('SipherMSP.peer','WhiteBoxMSP.peer')" >&log.txt
    res=$?
    set +x
  else
    set -x
    peer chaincode instantiate -o osinstance1.cerberus.dev:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $OSINSTANCE1_CA -C $INSTITUTION_ACCOUNTS_CHANNEL -n instaccntschannelcc -l ${LANGUAGE} -v ${VERSION} -c '{"Args":["init", "myOrganization", "Anna", "myAddress", "angelowwa@gmail.com", "angelowwa@gmail.com", "0877150173"]}' -P "OR ('SipherMSP.peer','WhiteBoxMSP.peer')" >&log.txt
    res=$?
    set +x
  fi
  cat log.txt
  verifyResult $res "Chaincode instantiation on ${PEER_NAME}.${ORGANIZATION_NAME} on channel '$INSTITUTION_ACCOUNTS_CHANNEL' failed"
  echo "===================== Chaincode is instantiated on ${PEER_NAME}.${ORGANIZATION_NAME} on channel '$INSTITUTION_ACCOUNTS_CHANNEL' ===================== "
  echo
}

instantiateIntegrationAccountsChaincode() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG

  # while 'peer chaincode' command can get the orderer endpoint from the peer
  # (if join was successful), let's supply it directly as we know it using
  # the "-o" option
  if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
    set -x
    peer chaincode instantiate -o osinstance2.cerberus.dev:7050 -C $INTEGRATION_ACCOUNTS_CHANNEL -n integraccntschannelcc -l ${LANGUAGE} -v ${VERSION} -c '{"Args":["init", "myOrganization", "Anna", "myAddress", "angelowwa@gmail.com", "angelowwa@gmail.com", "0877150173"]}' -P "OR ('SipherMSP.peer','WhiteBoxMSP.peer')" >&log.txt
    res=$?
    set +x
  else
    set -x
    peer chaincode instantiate -o osinstance2.cerberus.dev:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $OSINSTANCE2_CA -C $INTEGRATION_ACCOUNTS_CHANNEL -n integraccntschannelcc -l ${LANGUAGE} -v ${VERSION} -c '{"Args":["init", "myOrganization", "Anna", "myAddress", "angelowwa@gmail.com", "angelowwa@gmail.com", "0877150173"]}' -P "OR ('SipherMSP.peer','WhiteBoxMSP.peer')" >&log.txt
    res=$?
    set +x
  fi
  cat log.txt
  verifyResult $res "Chaincode instantiation on ${PEER_NAME}.${ORGANIZATION_NAME} on channel '$INTEGRATION_ACCOUNTS_CHANNEL' failed"
  echo "===================== Chaincode is instantiated on ${PEER_NAME}.${ORGANIZATION_NAME} on channel '$INTEGRATION_ACCOUNTS_CHANNEL' ===================== "
  echo
}

chaincodeQuery() {
  PEER=$1
  ORG=$2
  setGlobals $PEER $ORG
  #EXPECTED_RESULT=$3
  echo "===================== Querying on ${PEER_NAME}.${ORGANIZATION_NAME} on channel '$PERSON_ACCOUNTS_CHANNEL'... ===================== "
  local rc=1
  local starttime=$(date +%s)

  # continue to poll
  # we either get a successful response, or reach TIMEOUT
  while
    test "$(($(date +%s) - starttime))" -lt "$TIMEOUT" -a $rc -ne 0
  do
    sleep $DELAY
    echo "Attempting to Query ${PEER_NAME}.${ORGANIZATION_NAME} ...$(($(date +%s) - starttime)) secs"
    set -x
    peer chaincode query -C $PERSON_ACCOUNTS_CHANNEL -n personaccountscc -c '{"Args":["queryPersonAccountByEmail","angelowwa@gmail.com"]}' >&log.txt
    res=$?
    set +x

    test $res -eq 0 && VALUE=$(cat log.txt | awk '/Query Result/ {print $NF}')
    echo $VALUE
    echo "Value printed"

    #test "$VALUE" = "$EXPECTED_RESULT" && let rc=0
    # removed the string "Query Result" from peer chaincode query command
    # result. as a result, have to support both options until the change
    # is merged.
    #test $rc -ne 0 && VALUE=$(cat log.txt | egrep '^[0-9]+$')
    #test "$VALUE" = "$EXPECTED_RESULT" && let rc=0
  done

  echo
  cat log.txt
  verifyResult $res "!!!!!!!!!!!!!!! Query result on ${PEER_NAME}.${ORGANIZATION_NAME} is INVALID !!!!!!!!!!!!!!!!"
  echo "===================== Query successful on ${PEER_NAME}.${ORGANIZATION_NAME} on channel '$PERSON_ACCOUNTS_CHANNEL' ===================== "


  #echo "================== ERROR !!! FAILED to execute End-2-End Scenario =================="
  #echo
}
