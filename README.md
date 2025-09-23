# Hyperledger Fabric Network - ACME Network
A production-ready Hyperledger Fabric network with 3 organizations, each with complete PKI infrastructure and distributed services.

## 🏗️ Architecture 
### Organizations 

Org1 - org1.acme.com

Org2 - org2.acme.com

Org3 - org3.acme.com

### Services per Organization
Each organization runs:

Peer - Transaction node (peer0.orgX.acme.com)

Orderer - Ordering service (orderer.orgX.acme.com)

Root CA - Root certificate authority

TLS Root CA - TLS root CA

Intermediate CA - Intermediate issuing CA

TLS Intermediate CA - TLS intermediate CA

### 🚀 Quick Start
#### Prerequisites
- bash
- docker 
- docker-compose 
## Deploy Network
- bash
./scripts/up.sh
## Stop and Clean
- bash
./scripts/down.sh

## 🔐 Security Features
TLS encryption for all communications

Multi-level PKI hierarchy

Role-based access control

Organization-specific endorsement policies

## 🌐 Endpoints
Org1
- Peer: peer0.org1.acme.com:7051

- Orderer: orderer.org1.acme.com:7050

- CA: ca.org1.acme.com:7054

Org2
- Peer: peer0.org2.acme.com:8051

- Orderer: orderer.org2.acme.com:8050

- CA: ca.org2.acme.com:8054

Org3
- Peer: peer0.org3.acme.com:9051

- Orderer: orderer.org3.acme.com:9050

- CA: ca.org3.acme.com:9054

## 🐛 Debug Mode
