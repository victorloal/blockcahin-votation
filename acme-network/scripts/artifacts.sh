#genera archivos de configuración de anclaje para cada MSP en Hyperledger Fabric. 
#Utiliza configtxgen para crear un bloque genesis y un archivo de configuración de canal, 
#luego llama a la función generateAnchorConfigurationTx para cada MSP, generando archivos de anclaje para cada uno. 
#También verifica la disponibilidad de configtxgen.

set -x
function generateAnchorConfigurationTx() {
    msp=$1
    configtxgen -outputAnchorPeersUpdate ../channel-artifacts/${msp}anchors.tx -profile ThreeOrgsChannel -asOrg $msp -channelID marketplace
}

which configtxgen
if [ "$?" -ne 0 ]; then
    echo "configtxgen tool not found. exiting"
    exit 1
fi
export FABRIC_CFG_PATH=$(cd ../ && pwd)
configtxgen -profile ThreeOrgsOrdererGenesis -channelID sys-channel -outputBlock ../channel-artifacts/genesis.block
configtxgen -profile ThreeOrgsChannel -outputCreateChannelTx ../channel-artifacts/channel.tx -channelID marketplace
generateAnchorConfigurationTx Org1MSP
generateAnchorConfigurationTx Org2MSP
generateAnchorConfigurationTx Org3MSP

