#Este script baja la red de Hyperledger Fabric

set -e
cd .. && docker-compose -f docker-compose-root-ca.yaml -f docker-compose-int-ca.yaml -f docker-compose-cli-couchdb.yaml down
cd scripts && ./cleancerts.sh
rm -r ../channel-artifacts/*
rm -r ../crypto-config/org1.acme.com/*/*.pem
rm -r ../crypto-config/org2.acme.com/*/*.pem
rm -r ../crypto-config/org3.acme.com/*/*.pem
rm -r ../crypto-config/org1.acme.com/users
rm -r ../crypto-config/org2.acme.com/users
rm -r ../crypto-config/org3.acme.com/users
