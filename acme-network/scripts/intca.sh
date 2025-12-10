
#Este script automatiza la configuración de Autoridades de Certificación intermedias y sus equivalentes TLS para varias organizaciones en Hyperledger Fabric. 
#Para cada entidad, inscribe identidades iniciales para CA intermedias y TLS, utilizando credenciales predefinidas y comunicándose con las CA raíz.

set -x
# UDENAR
export CSR_NAMES_UDENAR="C=CO,ST=Nariño,L=Pasto,O=org1,OU=Hyperledger Fabric"
# Enroll bootstrap identity of int CA
export FABRIC_CA_CLIENT_HOME=../crypto-config/org1.acme.com/int/clients/admin
fabric-ca-client enroll -u http://admin:adminpw@localhost:7056 --csr.names "$CSR_NAMES_UDENAR"
# Enroll bootstrap identity of tls int CA
export FABRIC_CA_CLIENT_HOME=../crypto-config/org1.acme.com/tls-int/clients/admin
fabric-ca-client enroll -u http://admin:adminpw@localhost:7057 --csr.names "$CSR_NAMES_UDENAR"

# MARIAN
export CSR_NAMES_UMARIANA="C=CO,ST=Nariño,L=Pasto,O=org2,OU=Hyperledger Fabric"
# Enroll bootstrap identity of int CA
export FABRIC_CA_CLIENT_HOME=../crypto-config/org2.acme.com/int/clients/admin
fabric-ca-client enroll -u http://admin:adminpw@localhost:8056 --csr.names "$CSR_NAMES_UMARIANA"
# Enroll bootstrap identity of tls int CA
export FABRIC_CA_CLIENT_HOME=../crypto-config/org2.acme.com/tls-int/clients/admin
fabric-ca-client enroll -u http://admin:adminpw@localhost:8057 --csr.names "$CSR_NAMES_UMARIANA"

# UCC
export CSR_NAMES_UCC="C=CO,ST=Nariño,L=Pasto,O=org3,OU=Hyperledger Fabric"
# Enroll bootstrap identity of int CA
export FABRIC_CA_CLIENT_HOME=../crypto-config/org3.acme.com/int/clients/admin
fabric-ca-client enroll -u http://admin:adminpw@localhost:9056 --csr.names "$CSR_NAMES_UCC"
# Enroll bootstrap identity of tls int CA
export FABRIC_CA_CLIENT_HOME=../crypto-config/org3.acme.com/tls-int/clients/admin
fabric-ca-client enroll -u http://admin:adminpw@localhost:9057 --csr.names "$CSR_NAMES_UCC"

