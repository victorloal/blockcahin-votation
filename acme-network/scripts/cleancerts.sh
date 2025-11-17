# Este script limpia los certificados generados por la red de Hyperledger Fabric, y los archivos generados por la red de Hyperledger Fabric.

set -e

function safe_rm() {
    sudo rm -rf "$1" 2>/dev/null || true
}


function cleanCA(){
    org=$1
    ca=$2

    CA_PATH=../crypto-config/$org/$ca
    safe_rm $CA_PATH/clients
    safe_rm $CA_PATH/msp
    safe_rm $CA_PATH/ca-cert.pem
    safe_rm $CA_PATH/fabric-ca-server.db
    safe_rm $CA_PATH/IssuerPublicKey
    safe_rm $CA_PATH/IssuerRevocationPublicKey
    CA_CHAIN_FILE=$CA_PATH/ca-chain.pem
    if test -f "$CA_CHAIN_FILE"; then
        safe_rm $CA_CHAIN_FILE
    fi
}

    

function cleanOrgMSP() {
    org=$1

    MSP_PATH=../crypto-config/$org/msp
    safe_rm $MSP_PATH/cacerts
    safe_rm $MSP_PATH/intermediatecerts
    safe_rm $MSP_PATH/tlscacerts
    safe_rm $MSP_PATH/tlsintermediatecerts
}

function cleanLocalMSP() {
    org=$1
    name=$2
    type=$3

    LOCAL_MSP_PATH=../crypto-config/$org/${type}s/$name/msp
    TLS_FOLDER_PATH=../crypto-config/$org/${type}s/$name/tls

    safe_rm $LOCAL_MSP_PATH
    safe_rm $TLS_FOLDER_PATH
}

#cleanCA acme.com root
#cleanCA acme.com int
#cleanCA acme.com tls-root
#cleanCA acme.com tls-int
cleanCA org1.acme.com root
cleanCA org1.acme.com int
cleanCA org1.acme.com tls-root
cleanCA org1.acme.com tls-int
cleanCA org2.acme.com root
cleanCA org2.acme.com int
cleanCA org2.acme.com tls-root
cleanCA org2.acme.com tls-int
cleanCA org3.acme.com root
cleanCA org3.acme.com int
cleanCA org3.acme.com tls-root
cleanCA org3.acme.com tls-int

cleanOrgMSP org1.acme.com
cleanOrgMSP org2.acme.com
cleanOrgMSP org3.acme.com
#cleanOrgMSP acme.com

cleanLocalMSP org1.acme.com orderer.org1.acme.com orderer
cleanLocalMSP org2.acme.com orderer.org2.acme.com orderer
cleanLocalMSP org3.acme.com orderer.org3.acme.com orderer

cleanLocalMSP org1.acme.com peer0.org1.acme.com peer
cleanLocalMSP org2.acme.com peer0.org2.acme.com peer
cleanLocalMSP org3.acme.com peer0.org3.acme.com peer
#cleanLocalMSP acme.com orderer.acme.com orderer

cleanLocalMSP org1.acme.com admin@org1.acme.com user
cleanLocalMSP org2.acme.com admin@org2.acme.com user
cleanLocalMSP org3.acme.com admin@org3.acme.com user
#cleanLocalMSP acme.com admin@acme.com user