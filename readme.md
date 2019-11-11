Cerberus
Project dependencies:
Hyperledger binaries and images - v. 1.4.0:
curl -sSL http://bit.ly/2ysbOFE | bash -s 1.4.0

Rest:
go 1.11
install docker 18.09.4
install docker-compose 1.21.2
ipfs 0.4.19

bring up the containers: ./cerberus/hl/network/cerberus.sh up

bring down the containers: ./cerberus/hl/network/cerberus.sh down

bash settings files are: /cerberus/hl/network/scripts and /cerberus/hl/network/cerberusntw.sh

to work with the project functions: start the network with ./cerberusntw.sh up start ipfs server with ipfs daemon --writable=true

call querying and invoking functions from folder /cerberus/app for each type of accounts all requested arguments are specified
