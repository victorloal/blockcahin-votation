Configuración y ejecución de la red Hyperledger Fabric
=============
acme-network
-------------



Este curso usa el siguiente repositorio:

https://github.com/blockchainempresarial/curso-hyperledger-fabric.git
Asegúrese de tener el control de acceso correcto a este repositorio; si tiene problemas, envíe un correo electrónico a ricardo@blockchainempresarial.com

### Directorio de trabajo
Siga y ejecute los siguientes comandos


```shell
 cd $HOME
 git clone https://github.com/blockchainempresarial/curso-hyperledger-fabric.git
 cd curso-hyperledger-fabric/acme-network

```
'../crypto/org1.acme.com/tlsca/tlsca.org1.acme.com-cert.pem'
### Parámetros globales
Ejecute el siguiente comando para definir parámetros globales en la consola de Linux. (temporales)


```shell
export CHANNEL_NAME=marketplace
export VERBOSE=false
export FABRIC_CFG_PATH=$PWD
```

### Certificados
Siga y ejecute los siguientes comandos para generar certificados utilizando la herramienta de cifrado

Primero, cargue los siguientes archivos de configuración en el directorio de trabajo.
crypto-config.yaml

#### Generar certificados de los pers y orderers
Generar certificado, el siguiente comando creará un directorio de configuración de cifrado que contiene varios certificados y claves para pedidos y pares

```shell
cryptogen generate --config=./crypto-config.yaml

```
#### Generando el bloque Orderer Genesis

Más información aquí:
https://hyperledger-fabric.readthedocs.io/en/release-2.2/configtx.html

Primero, cargue los siguientes archivos de configuración en el directorio de trabajo.
configtx.yaml

Para crear el bloque orderer genesis, es necesario utilizar y ejecutar la herramienta configtxgen
El bloque Génesis es el primer bloque de nuestra cadena de bloques. Se utiliza para iniciar el servicio de ordenamiento y asegura la configuración del canal.


```shell
mkdir channel-artifacts
configtxgen  -profile ThreeOrgsOrdererGenesis 	-channelID system-channel -outputBlock ./channel-artifacts/genesis.block
```


#### Generando transacción de configuración de canal 'channel.tx'

Para la transacción de configuración del canal: channel.tx es la transacción que le permite crear el canal Hyperledger Fabric. El canal es la ubicación donde existe el libro mayor y el mecanismo que permite a los pares unirse a las redes de negocio.

```shell
configtxgen -profile ThreeOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
```


#### Anchor peers transactions
Las transacciones de los anchos peers   especifican el Peer de anclaje de cada organización en este canal. Ejecute los siguientes tres comandos para definir el anchor peer para cada organización


##### Generando el  anchor peer update para  Org1MSP 
```shell
configtxgen -profile ThreeOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP

```
##### Generando anchor peer update para Org2MSP

```shell
configtxgen -profile ThreeOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP
```

##### Generando anchor peer update para Org3MSP

```shell
configtxgen -profile ThreeOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org3MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org3MSP
```

### Inspecciona los artefactos

Para inspeccionar el channel.tx en formato json, siga las siguientes instrucciones.

```shell
configtxgen --inspectChannelCreateTx ./channel-artifacts/channel.tx >> ./channel-artifacts/channel.tx.json
```

Para ver el archivo de resultados:

```shell
vi channel-artifacts/channel.tx.json

```
Para inspeccionar genesis.block en formato json, siga las siguientes instrucciones.

```shell
configtxgen --inspectBlock ./channel-artifacts/genesis.block >> ./channel-artifacts/genesis.block.json

```
Para ver el archivo de resultados:

```shell
vi channel-artifacts/genesis.block.json

```

docker volume create portainer_data



CHANNEL_NAME=$CHANNEL_NAME docker-compose -f docker-compose-cli-couchdb.yaml down -v

CHANNEL_NAME=$CHANNEL_NAME docker-compose -f docker-compose-cli-couchdb.yaml up -d


-- Nuevos usuarios 


fabric-ca-client register --id.name User2 --id.secret 123123 --id.type client --url https://localhost:7054 --mspdir /etc/hyperledger/fabric-ca-server/msp/ 





















export FABRIC_CA_CLIENT_HOME=./crypto-config/org1.acme.com/ca/

export FABRIC_CA_CLIENT_TLS_CERTFILES=./crypto-config/org1.acme.com/tlsca/tlsca.org1.acme.com-cert.pem

export FABRIC_CA_CLIENT_HOME=/home/victor/Documentos/votaciones-/votaciones/acme-network/crypto-config/org1.acme.com/users/admin@org1.acme.com/

export FABRIC_CA_CLIENT_TLS_CERTFILES=/home/victor/Documentos/votaciones-/votaciones/acme-network/crypto-config/org1.acme.com/users/admin@org1.acme.com/msp/tlscacerts/tlsca.org1.acme.com-cert.pem

export FABRIC_CA_CLIENT_TLS_ENABLE=false
 
export FABRIC_CA_CLIENT_CERTIFILES= '/home/victor/Documentos/votaciones-/votaciones/acme-network/crypto-config/org1.acme.com/users/admin@org1.acme.com/msp/ca/'


fabric-ca-client register --id.name user2 --id.secret password123 --id.type client -u https://admin:adminpw@localhost:7054 --tls.certfiles $FABRIC_CA_CLIENT_TLS_CERTFILES 

fabric-ca-client register --id.name user2 --id.secret password123 --id.type client -u https://admin:adminpw@localhost:7054 --tls.certfiles /path/to/ca-cert.pem


fabric-ca-client register -u https://admin:adminpw@ca.org1.acme.com:7054 --insecure
fabric-ca-client enroll -u https://admin:adminpw@victor-A520M-DS3H:7054 --insecure

# enroll registered identity

export FABRIC_CA_CLIENT_HOME=../crypto-config/$org/$ca/clients/$id_name fabric-ca-client enroll -u http://$id_name:$id_secret@localhost:$ca_port --csr.names "$csr_names" --csr.hosts "$csr_hosts"


export FABRIC_CA_CLIENT_HOME=/home/victor/Documentos/votaciones/votaciones/acme-network/crypto-config/org1.acme.com/ca/msp