#Business Blockchain Training & Consulting SpA. All Rights Reserved.
#www.blockchainempresarial.com
#email: ricardo@blockchainempresarial.com

export CHANNEL_NAME=marketplace
export CHAINCODE_NAME=votaciones
export CHAINCODE_VERSION=1
export CC_RUNTIME_LANGUAGE=golang
export CC_SRC_PATH="../../../chaincode/$CHAINCODE_NAME/"
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/orderers/orderer.org1.acme.com/tls/ca.crt
export ORDERER_ADDRESS=orderer.org1.acme.com:7050
 
export VERBOSE=true
export CLI_DELAY=3
export CLI_TIMEOUT=10


#Empaqueta el chaincode
peer lifecycle chaincode package ${CHAINCODE_NAME}.tar.gz --path ${CC_SRC_PATH} --lang ${CC_RUNTIME_LANGUAGE} --label ${CHAINCODE_NAME}_${CHAINCODE_VERSION} >&log.txt

#peer lifecycle chaincode install example
#first peer peer0.org1.acme.com
peer lifecycle chaincode install ${CHAINCODE_NAME}.tar.gz 

#Actualizar este  valor con el que obtengan al empaquetar el chaincode: votaciones_1:





export CC_PACKAGEID=e7baa7bf12c50aeace123a5db0516e8c0d687003019e2209afb4c7544b5eb2ed




# peer0.org2
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/users/admin@org2.acme.com/msp CORE_PEER_ADDRESS=peer0.org2.acme.com:7051 CORE_PEER_LOCALMSPID="Org2MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/peers/peer0.org2.acme.com/tls/ca.crt peer lifecycle chaincode install  ${CHAINCODE_NAME}.tar.gz



# peer0.org3
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/users/admin@org3.acme.com/msp CORE_PEER_ADDRESS=peer0.org3.acme.com:7051 CORE_PEER_LOCALMSPID="Org3MSP" CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/peers/peer0.org3.acme.com/tls/ca.crt peer lifecycle chaincode install  ${CHAINCODE_NAME}.tar.gz






#Endorsement policy for lifecycle chaincode 
export CHAINCODE_VERSION=1
export CC_RUNTIME_LANGUAGE=golang
export CC_SRC_PATH="../../../chaincode/$CHAINCODE_NAME/"
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/orderers/orderer.org1.acme.com/tls/ca.crt
export ORDERER_ADDRESS=orderer.org1.acme.com:7050
export CORE_PEER_ADDRESS=peer0.org1.acme.com:7051
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/users/admin@org1.acme.com/msp
export CORE_PEER_ADDRESS=peer0.org1.acme.com:7051
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt

peer lifecycle chaincode approveformyorg --tls --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --waitForEvent --signature-policy "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')" --package-id votaciones_1:$CC_PACKAGEID -o $ORDERER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses $CORE_PEER_ADDRESS --waitForEventTimeout 120s


#definir la politica de endorsamiento para el chaincode en org2
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/orderers/orderer.org2.acme.com/tls/ca.crt
export ORDERER_ADDRESS=orderer.org2.acme.com:7050
export CORE_PEER_ADDRESS=peer0.org2.acme.com:7051
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/peers/peer0.org2.acme.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/users/admin@org2.acme.com/msp
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/users/admin@org2.acme.com/msp CORE_PEER_LOCALMSPID="Org2MSP" peer lifecycle chaincode approveformyorg --tls --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --waitForEvent --signature-policy "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')" --package-id votaciones_1:$CC_PACKAGEID -o $ORDERER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses $CORE_PEER_ADDRESS


#definir la politica de endorsamiento para el chaincode en org3
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/orderers/orderer.org3.acme.com/tls/ca.crt
export ORDERER_ADDRESS=orderer.org3.acme.com:7050
export CORE_PEER_ADDRESS=peer0.org3.acme.com:7051
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/peers/peer0.org3.acme.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/users/admin@org3.acme.com/msp
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/users/admin@org3.acme.com/msp CORE_PEER_LOCALMSPID="Org3MSP" peer lifecycle chaincode approveformyorg --tls --cafile $ORDERER_CA  --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --waitForEvent --signature-policy "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')" --package-id votaciones_1:$CC_PACKAGEID -o $ORDERER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses $CORE_PEER_ADDRESS


#Realizar el commit del chaincode en el canal para todos los peers

export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/users/admin@org1.acme.com/msp
export CORE_PEER_ADDRESS=peer0.org1.acme.com:7051
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt

peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --signature-policy "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')"  --output json


export CORE_PEER_ADDRESS=peer0.org1.acme.com:7051
export ORDERER_ADDRESS=orderer.acme.com:7050
# peer lifecycle chaincode commit -o $ORDERER_ADDRESS  --tls --cafile $ORDERER_CA  --peerAddresses peer0.org1.acme.com:7051  --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt --peerAddresses peer0.org2.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/peers/peer0.org2.acme.com/tls/ca.crt --peerAddresses peer0.org3.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/peers/peer0.org3.acme.com/tls/ca.crt --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version 1 --sequence 1 --signature-policy "OR ('Org1MSP.peer','Or2MSP.peer','Org3MSP.peer')" 

peer lifecycle chaincode commit -o $ORDERER_ADDRESS --tls --cafile $ORDERER_CA \
  --peerAddresses peer0.org1.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt \
  --peerAddresses peer0.org2.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/peers/peer0.org2.acme.com/tls/ca.crt \
  --peerAddresses peer0.org3.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/peers/peer0.org3.acme.com/tls/ca.crt \
  --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 \
  --signature-policy "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')"


#verificar que el chaincode se instalo correctamente
export CORE_PEER_ADDRESS=peer0.org1.acme.com:7051 
export ORDERER_ADDRESS=orderer.acme.com:7050
peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --output json --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt





# InitLedger
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt
peer chaincode invoke -o $ORDERER_ADDRESS --tls --cafile $ORDERER_CA -C  $CHANNEL_NAME  -n $CHAINCODE_NAME -c '{"Args":["InitLedger"]}' --peerAddresses peer0.org1.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt



#definir la politica de endorsamiento para el chaincode en org3
export ORDERER_ADDRESS=orderer.acme.com:7050
export CORE_PEER_ADDRESS=peer0.org1.acme.com:7051
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/users/admin@org1.acme.com/msp
# PostulateCandidate
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt
peer chaincode invoke -o $ORDERER_ADDRESS --tls --cafile $ORDERER_CA -C  $CHANNEL_NAME  -n $CHAINCODE_NAME -c '{"Args":["PostulateCandidate","victor"]}' --peerAddresses peer0.org1.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt

# GetCandidateVotes
#definir la politica de endorsamiento para el chaincode en org3

export CORE_PEER_LOCALMSPID="Org2MSP"
export ORDERER_ADDRESS=orderer.acme.com:7050
export CORE_PEER_ADDRESS=peer0.org2.acme.com:7051
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/peers/peer0.org2.acme.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/users/admin@org2.acme.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/peers/peer0.org2.acme.com/tls/ca.crt
peer chaincode invoke -o $ORDERER_ADDRESS --tls --cafile $ORDERER_CA -C  $CHANNEL_NAME  -n $CHAINCODE_NAME -c '{"Args":["GetCandidateVotes","victor"]}' --peerAddresses peer0.org2.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/peers/peer0.org2.acme.com/tls/ca.crt

# VoteForCandidate
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/peers/peer0.org2.acme.com/tls/ca.crt
peer chaincode invoke -o $ORDERER_ADDRESS --tls --cafile $ORDERER_CA -C  $CHANNEL_NAME  -n $CHAINCODE_NAME -c '{"Args":["VoteForCandidate","orlando","victor"]}' --peerAddresses peer0.org2.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/peers/peer0.org2.acme.com/tls/ca.crt