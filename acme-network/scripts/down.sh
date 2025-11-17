#Este script baja la red de Hyperledger Fabric

set -e

function safe_rm() {
    sudo rm -rf "$1" 2>/dev/null || true
}
cd .. && docker-compose -f docker-compose-root-ca.yaml -f docker-compose-int-ca.yaml -f docker-compose-cli-couchdb.yaml down
cd scripts && ./cleancerts.sh
safe_rm "../channel-artifacts"
safe_rm "../crypto-config/org1.acme.com/*/*.pem"
safe_rm "../crypto-config/org2.acme.com/*/*.pem"
safe_rm "../crypto-config/org3.acme.com/*/*.pem"
safe_rm "../crypto-config/org1.acme.com/users"
safe_rm "../crypto-config/org2.acme.com/users"
safe_rm "../crypto-config/org3.acme.com/users"

echo "La red de Hyperledger Fabric ha sido detenida y los certificados han sido limpiados."

