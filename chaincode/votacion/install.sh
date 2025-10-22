#Business Blockchain Training & Consulting SpA. All Rights Reserved.
#www.blockchainempresarial.com
#email: ricardo@blockchainempresarial.com

export CHANNEL_NAME=marketplace
export CHAINCODE_NAME=votacion
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

#Actualizar este  valor con el que obtengan al empaquetar el chaincode: votacion_1:





export CC_PACKAGEID=3ed61b0cd15e6663ce1963d3412bccab8f849b665d50d86ad7823164d7a260f4




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

peer lifecycle chaincode approveformyorg --tls --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --waitForEvent --signature-policy "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')" --package-id votacion_1:$CC_PACKAGEID -o $ORDERER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses $CORE_PEER_ADDRESS --waitForEventTimeout 120s


#definir la politica de endorsamiento para el chaincode en org2
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/orderers/orderer.org2.acme.com/tls/ca.crt
export ORDERER_ADDRESS=orderer.org2.acme.com:7050
export CORE_PEER_ADDRESS=peer0.org2.acme.com:7051
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/peers/peer0.org2.acme.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/users/admin@org2.acme.com/msp
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/users/admin@org2.acme.com/msp CORE_PEER_LOCALMSPID="Org2MSP" peer lifecycle chaincode approveformyorg --tls --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --waitForEvent --signature-policy "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')" --package-id votacion_1:$CC_PACKAGEID -o $ORDERER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses $CORE_PEER_ADDRESS


#definir la politica de endorsamiento para el chaincode en org3
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/orderers/orderer.org3.acme.com/tls/ca.crt
export ORDERER_ADDRESS=orderer.org3.acme.com:7050
export CORE_PEER_ADDRESS=peer0.org3.acme.com:7051
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/peers/peer0.org3.acme.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/users/admin@org3.acme.com/msp
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/users/admin@org3.acme.com/msp CORE_PEER_LOCALMSPID="Org3MSP" peer lifecycle chaincode approveformyorg --tls --cafile $ORDERER_CA  --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --waitForEvent --signature-policy "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')" --package-id votacion_1:$CC_PACKAGEID -o $ORDERER_ADDRESS --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses $CORE_PEER_ADDRESS


#Realizar el commit del chaincode en el canal para todos los peers

export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/users/admin@org1.acme.com/msp
export CORE_PEER_ADDRESS=peer0.org1.acme.com:7051
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt

peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 --signature-policy "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')"  --output json


export CORE_PEER_ADDRESS=peer0.org1.acme.com:7051
export ORDERER_ADDRESS=orderer.acme.com:7050
# peer lifecycle chaincode commit -o $ORDERER_ADDRESS  --tls --cafile $ORDERER_CA  --peerAddresses peer0.org1.acme.com:7051  --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt --peerAddresses peer0.org2.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/peers/peer0.org2.acme.com/tls/ca.crt --peerAddresses peer0.org3.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/peers/peer0.org3.acme.com/tls/ca.crt --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version 1 --sequence 1 --signature-policy "OR ('Org1MSP.peer','Or2MSP.peer','Org3MSP.peer')" 
export CHANNEL_NAME=marketplace
export CHAINCODE_NAME=votacion
export CHAINCODE_VERSION=1
export CC_RUNTIME_LANGUAGE=golang
export CC_SRC_PATH="../../../chaincode/$CHAINCODE_NAME/"
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/orderers/orderer.org1.acme.com/tls/ca.crt
export ORDERER_ADDRESS=orderer.org1.acme.com:7050
 
export VERBOSE=true
export CLI_DELAY=3
export CLI_TIMEOUT=10


peer lifecycle chaincode commit -o $ORDERER_ADDRESS --tls --cafile $ORDERER_CA \
  --peerAddresses peer0.org1.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt \
  --peerAddresses peer0.org2.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org2.acme.com/peers/peer0.org2.acme.com/tls/ca.crt \
  --peerAddresses peer0.org3.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org3.acme.com/peers/peer0.org3.acme.com/tls/ca.crt \
  --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --version $CHAINCODE_VERSION --sequence 1 \
  --signature-policy "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')"


#verificar que el chaincode se instalo correctamente
export CORE_PEER_ADDRESS=peer0.org1.acme.com:7051 
export ORDERER_ADDRESS=orderer.org1.acme.com:7050
peer lifecycle chaincode querycommitted --channelID $CHANNEL_NAME --name $CHAINCODE_NAME --output json --peerAddresses $CORE_PEER_ADDRESS --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt





# InitLedger
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt
peer chaincode invoke -o $ORDERER_ADDRESS --tls --cafile $ORDERER_CA -C  $CHANNEL_NAME  -n $CHAINCODE_NAME -c '{"Args":["NuevaVotacion","2","prueba","2025-10-17:12.00","2025-10-17:02.40","true"]}' --peerAddresses peer0.org1.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt


# CrearListaCandidatos
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt
peer chaincode invoke -o $ORDERER_ADDRESS --tls --cafile $ORDERER_CA -C  $CHANNEL_NAME  -n $CHAINCODE_NAME -c '{"Args":["CrearListaCandidatos","2"]}' --peerAddresses peer0.org1.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt

#AgregarCandidato
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt
peer chaincode invoke -o $ORDERER_ADDRESS --tls --cafile $ORDERER_CA -C  $CHANNEL_NAME  -n $CHAINCODE_NAME -c '{"Args":["AgregarCandidato","2","1","1193458050"]}' --peerAddresses peer0.org1.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt


#AgregarCandidato
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt
peer chaincode invoke -o $ORDERER_ADDRESS --tls --cafile $ORDERER_CA -C  $CHANNEL_NAME  -n $CHAINCODE_NAME -c '{"Args":["AgregarCandidato","2","2","1193458052"]}' --peerAddresses peer0.org1.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt



#AgregarCandidato
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt
peer chaincode invoke -o $ORDERER_ADDRESS --tls --cafile $ORDERER_CA -C  $CHANNEL_NAME  -n $CHAINCODE_NAME -c '{"Args":["AgregarCandidato","2","3","1193458053"]}' --peerAddresses peer0.org1.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt



#PruebaVotacion
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt
peer chaincode invoke -o $ORDERER_ADDRESS --tls --cafile $ORDERER_CA -C  $CHANNEL_NAME  -n $CHAINCODE_NAME -c '{"Args":["PruebaVotacion","2","[0,0,1,0,1,2]","1193458053","{\"N\":31911166551840822647174294053386726304103157927073240083725948790393534002107034654477032954888565198136435751832056451130879121107890693495357841926723161163285401065796567852510584335261087462523287823670244230206418312008499695599346964503027679678826665817438847361586019548682538955077448241541196750332922063548103499494444866731805525281454828482275314920059196597493463678358138014241102030474869215538785066450704892711412666647652379305145574365495285681599221280316899303900544362802877345293067362429830645019480144069691766529285736446749263288231135420724062504498800596847089588086443700737859786331701,\"G\":31911166551840822647174294053386726304103157927073240083725948790393534002107034654477032954888565198136435751832056451130879121107890693495357841926723161163285401065796567852510584335261087462523287823670244230206418312008499695599346964503027679678826665817438847361586019548682538955077448241541196750332922063548103499494444866731805525281454828482275314920059196597493463678358138014241102030474869215538785066450704892711412666647652379305145574365495285681599221280316899303900544362802877345293067362429830645019480144069691766529285736446749263288231135420724062504498800596847089588086443700737859786331702,\"NSquared\":1018322550699324498669370196850946604789457446060187215991614639446322647586299237257745248147597296042240871895775553102656225373921182294157658241223091417497785278828842908769831066356622964919157286148526051254760248365431921825375802559694951075301003807766501642201679072212597885341166905237250419512941050543511537652114575966332793514895529671132103471543931990671189980789173549642897291892339399281833191443220048441850569144983603571431085299324744006459903784703198452767087577872671273625319962643587634488055465285596021702476107509289834655797683218779113547920537298206354354110220413635355964756043114511118696595946387742977380775155732622994317494699163662991101781459812817525802172361835089204023695644647672159650268824844635879332620074632892124535585652979983623301461391771428038101665639381264116893880711464989995368471852718045065518075097270056231645527821470935811444079013590637737036863471966223816489683719728002475546160943557253388193232108775699214247540909938286446386602756246960266371569876255797679984308625436177972536322374626728581693865906044610421338715085315552741202336756487227101862721178135195820113728983683021237351361500754024753865345908623195825923426256387752863453861997553401}","{\"N\":31911166551840822647174294053386726304103157927073240083725948790393534002107034654477032954888565198136435751832056451130879121107890693495357841926723161163285401065796567852510584335261087462523287823670244230206418312008499695599346964503027679678826665817438847361586019548682538955077448241541196750332922063548103499494444866731805525281454828482275314920059196597493463678358138014241102030474869215538785066450704892711412666647652379305145574365495285681599221280316899303900544362802877345293067362429830645019480144069691766529285736446749263288231135420724062504498800596847089588086443700737859786331701,\"G\":31911166551840822647174294053386726304103157927073240083725948790393534002107034654477032954888565198136435751832056451130879121107890693495357841926723161163285401065796567852510584335261087462523287823670244230206418312008499695599346964503027679678826665817438847361586019548682538955077448241541196750332922063548103499494444866731805525281454828482275314920059196597493463678358138014241102030474869215538785066450704892711412666647652379305145574365495285681599221280316899303900544362802877345293067362429830645019480144069691766529285736446749263288231135420724062504498800596847089588086443700737859786331702,\"NSquared\":1018322550699324498669370196850946604789457446060187215991614639446322647586299237257745248147597296042240871895775553102656225373921182294157658241223091417497785278828842908769831066356622964919157286148526051254760248365431921825375802559694951075301003807766501642201679072212597885341166905237250419512941050543511537652114575966332793514895529671132103471543931990671189980789173549642897291892339399281833191443220048441850569144983603571431085299324744006459903784703198452767087577872671273625319962643587634488055465285596021702476107509289834655797683218779113547920537298206354354110220413635355964756043114511118696595946387742977380775155732622994317494699163662991101781459812817525802172361835089204023695644647672159650268824844635879332620074632892124535585652979983623301461391771428038101665639381264116893880711464989995368471852718045065518075097270056231645527821470935811444079013590637737036863471966223816489683719728002475546160943557253388193232108775699214247540909938286446386602756246960266371569876255797679984308625436177972536322374626728581693865906044610421338715085315552741202336756487227101862721178135195820113728983683021237351361500754024753865345908623195825923426256387752863453861997553401}"]}' --peerAddresses peer0.org1.acme.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/org1.acme.com/peers/peer0.org1.acme.com/tls/ca.crt