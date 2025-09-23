set -e

export PATH=$PATH:/home/victor/Documentos/fabric-samples/bin


cd .. && docker-compose -f docker-compose-root-ca.yaml up -d
sleep 5
cd scripts && ./rootca.sh
cd .. && docker-compose -f docker-compose-int-ca.yaml up -d
sleep 5
cd scripts && ./intca.sh
sleep 5
./identities.sh
sleep 5
./msp.sh
sleep 5
./artifacts.sh
sleep 5
cd .. && docker-compose -f docker-compose-cli-couchdb.yaml up -d
sleep 5
cd scripts && ./channels.sh


# Run Portainer container
#docker run -d -p 10000:10000 -p 9000:9000 -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer

