
#Este script configura autoridades de certificación para varias organizaciones en Hyperledger Fabric, 
#estableciendo CA raíz y TLS, y registrando CAs intermedios. 
#Automatiza el enroll de identidades administrativas y el registro de nuevas CAs intermedias. 

set -x


# UDENAR
export CSR_NAMES_UDENAR="C=CO,ST=Nariño,L=Pasto,O=org1,OU=Hyperledger Fabric"
# Enroll bootstrap identity of root CA
export FABRIC_CA_CLIENT_HOME=../crypto-config/org1.acme.com/root/clients/admin
fabric-ca-client enroll -u http://admin:adminpw@localhost:7054 --csr.names "$CSR_NAMES_UDENAR"
# Register intermediate CA in the root CA
fabric-ca-client register --id.name int.ca.org1.acme.com --id.secret password --id.attrs 'hf.IntermediateCA=true' -u http://admin:adminpw@localhost:7054
# Enroll bootstrap identity of tls root CA
export FABRIC_CA_CLIENT_HOME=../crypto-config/org1.acme.com/tls-root/clients/admin
fabric-ca-client enroll -u http://admin:adminpw@localhost:7055 --csr.names "$CSR_NAMES_UDENAR"
# Register intermediate CA in the tls root CA
fabric-ca-client register --id.name tls.int.ca.org1.acme.com --id.secret password --id.attrs 'hf.IntermediateCA=true' -u http://admin:adminpw@localhost:7055

# UMARIANA
export CSR_NAMES_UMARIANA="C=CO,ST=Nariño,L=Pasto,O=org2,OU=Hyperledger Fabric"
# Enroll bootstrap identity of root CA
export FABRIC_CA_CLIENT_HOME=../crypto-config/org2.acme.com/root/clients/admin
fabric-ca-client enroll -u http://admin:adminpw@localhost:8054 --csr.names "$CSR_NAMES_UMARIANA"
# Register intermediate CA in the root CA
fabric-ca-client register --id.name int.ca.org2.acme.com --id.secret password --id.attrs 'hf.IntermediateCA=true' -u http://admin:adminpw@localhost:8054
# Enroll bootstrap identity of tls root CA
export FABRIC_CA_CLIENT_HOME=../crypto-config/org2.acme.com/tls-root/clients/admin
fabric-ca-client enroll -u http://admin:adminpw@localhost:8055 --csr.names "$CSR_NAMES_UMARIANA"
# Register intermediate CA in the tls root CA
fabric-ca-client register --id.name tls.int.ca.org2.acme.com --id.secret password --id.attrs 'hf.IntermediateCA=true' -u http://admin:adminpw@localhost:8055

# UCC
export CSR_NAMES_UCC="C=CO,ST=Nariño,L=Pasto,O=org3,OU=Hyperledger Fabric"
# Enroll bootstrap identity of root CA
export FABRIC_CA_CLIENT_HOME=../crypto-config/org3.acme.com/root/clients/admin
fabric-ca-client enroll -u http://admin:adminpw@localhost:9054 --csr.names "$CSR_NAMES_UCC"
# Register intermediate CA in the root CA
fabric-ca-client register --id.name int.ca.org3.acme.com --id.secret password --id.attrs 'hf.IntermediateCA=true' -u http://admin:adminpw@localhost:9054
# Enroll bootstrap identity of tls root CA
export FABRIC_CA_CLIENT_HOME=../crypto-config/org3.acme.com/tls-root/clients/admin
fabric-ca-client enroll -u http://admin:adminpw@localhost:9055 --csr.names "$CSR_NAMES_UCC"
# Register intermediate CA in the tls root CA
fabric-ca-client register --id.name tls.int.ca.org3.acme.com --id.secret password --id.attrs 'hf.IntermediateCA=true' -u http://admin:adminpw@localhost:9055
