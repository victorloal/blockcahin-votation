#certificados para una red Hyperledger Fabric, registrando y matriculando identidades con las Autoridades de Certificación. 
#Se define para distintas entidades (admins, clientes, peers, orderers) y maneja tanto certificados regulares como TLS. 
#Utiliza funciones para registrar y matricular con parámetros como nombres de organización y tipos de identidad.
#!/bin/bash
set -x

function issueCertificates() {
    ca=$1
    ca_port=$2
    org=$3
    id_name=$4
    id_secret=$5
    id_type=$6
    csr_names=$7
    csr_hosts=$8


    # register identity with CA admin
    export FABRIC_CA_CLIENT_HOME=../crypto-config/$org/$ca/clients/admin
    fabric-ca-client register --id.name $id_name --id.secret $id_secret --id.type $id_type -u http://admin:adminpw@localhost:$ca_port
    # enroll registered identity
    export FABRIC_CA_CLIENT_HOME=../crypto-config/$org/$ca/clients/$id_name
    fabric-ca-client enroll -u http://$id_name:$id_secret@localhost:$ca_port --csr.names "$csr_names" --csr.hosts "$csr_hosts"
}

function issueTLSCertificates() {
    ca=$1
    ca_port=$2
    org=$3
    id_name=$4
    id_secret=$5
    id_type=$6
    csr_names=$7
    csr_hosts=$8

    
    # register identity with CA admin
    export FABRIC_CA_CLIENT_HOME=../crypto-config/$org/$ca/clients/admin
    fabric-ca-client register --id.name $id_name --id.secret $id_secret --id.type $id_type -u http://admin:adminpw@localhost:$ca_port
    # enroll registered identity
    export FABRIC_CA_CLIENT_HOME=../crypto-config/$org/$ca/clients/$id_name
    fabric-ca-client enroll -u http://$id_name:$id_secret@localhost:$ca_port --csr.names "$csr_names" --csr.hosts "$csr_hosts" --enrollment.profile tls
}

# Udenar
export CSR_NAMES_UDENAR="C=CO,ST=Narino,L=Pasto,O=org1,OU=Hyperledger Fabric"
# issue certificates for admin user identity
issueCertificates int 7056 org1.acme.com admin@org1.acme.com adminpw admin "$CSR_NAMES_UDENAR" ""
issueTLSCertificates tls-int 7057 org1.acme.com admin@org1.acme.com adminpw admin "$CSR_NAMES_UDENAR" "admin@org1.acme.com,localhost"
# issue certificates for client identity and for client tls
# issueCertificates int 7056 org1.acme.com client@org1.acme.com clientpw client "$CSR_NAMES_UDENAR" ""
# issueTLSCertificates tls-int 7057 org1.acme.com client@org1.acme.com clientpw client "$CSR_NAMES_UDENAR" "client@org1.acme.com,localhost"
# issue certificates for peer node identity and for peer server tls
issueCertificates int 7056 org1.acme.com peer0.org1.acme.com peerpw peer "$CSR_NAMES_UDENAR" ""
issueTLSCertificates tls-int 7057 org1.acme.com peer0.org1.acme.com peerpw peer "$CSR_NAMES_UDENAR" "peer0.org1.acme.com,localhost"
# issue certificates for orderer node identity and for orderer server tls
issueCertificates int 7056 org1.acme.com orderer.org1.acme.com ordererpw orderer "$CSR_NAMES_UDENAR" ""
issueTLSCertificates tls-int 7057 org1.acme.com orderer.org1.acme.com ordererpw orderer "$CSR_NAMES_UDENAR" "orderer.org1.acme.com,localhost"

# UMariana
export CSR_NAMES_UMARIANA="C=CO,ST=Narino,L=Pasto,O=org2,OU=Hyperledger Fabric"
# issue certificates for admin user identity
issueCertificates int 8056 org2.acme.com admin@org2.acme.com adminpw admin "$CSR_NAMES_UMARIANA" ""
issueTLSCertificates tls-int 8057 org2.acme.com admin@org2.acme.com adminpw admin "$CSR_NAMES_UMARIANA" "admin@org2.acme.com,localhost"
# issue certificates for client identity and for client tls
# issueCertificates int 8056 org2.acme.com client@org2.acme.com clientpw client "$CSR_NAMES_UMARIANA" ""
# issueTLSCertificates tls-int 8057 org2.acme.com client@org2.acme.com clientpw client "$CSR_NAMES_UMARIANA" "client@org2.acme.com,localhost"
# issue certificates for peer node identity and for peer server tls
issueCertificates int 8056 org2.acme.com peer0.org2.acme.com peerpw peer "$CSR_NAMES_UMARIANA" ""
issueTLSCertificates tls-int 8057 org2.acme.com peer0.org2.acme.com peerpw peer "$CSR_NAMES_UMARIANA" "peer0.org2.acme.com,localhost"
# issue certificates for orderer node identity and for orderer server tls
issueCertificates int 8056 org2.acme.com orderer.org2.acme.com ordererpw orderer "$CSR_NAMES_UMARIANA" ""
issueTLSCertificates tls-int 8057 org2.acme.com orderer.org2.acme.com ordererpw orderer "$CSR_NAMES_UMARIANA" "orderer.org2.acme.com,localhost"

# UCC
export CSR_NAMES_UCC="C=CO,ST=Narino,L=Pasto,O=org3,OU=Hyperledger Fabric"
# issue certificates for admin user identity
issueCertificates int 9056 org3.acme.com admin@org3.acme.com adminpw admin "$CSR_NAMES_UCC" ""
issueTLSCertificates tls-int 9057 org3.acme.com admin@org3.acme.com adminpw admin "$CSR_NAMES_UCC" "admin@org3.acme.com,localhost"
# issue certificates for client identity and for client tls
# issueCertificates int 9056 org3.acme.com client@org3.acme.com clientpw client "$CSR_NAMES_UCC" ""
# issueTLSCertificates tls-int 9057 org3.acme.com client@org3.acme.com clientpw client "$CSR_NAMES_UCC" "client@org3.acme.com,localhost"
# issue certificates for peer node identity and for peer server tls
issueCertificates int 9056 org3.acme.com peer0.org3.acme.com peerpw peer "$CSR_NAMES_UCC"
issueTLSCertificates tls-int 9057 org3.acme.com peer0.org3.acme.com peerpw peer "$CSR_NAMES_UCC" "peer0.org3.acme.com,localhost"
# issue certificates for orderer node identity and for orderer server tls
issueCertificates int 9056 org3.acme.com orderer.org3.acme.com ordererpw orderer "$CSR_NAMES_UCC" ""
issueTLSCertificates tls-int 9057 org3.acme.com orderer.org3.acme.com ordererpw orderer "$CSR_NAMES_UCC" "orderer.org3.acme.com,localhost"

