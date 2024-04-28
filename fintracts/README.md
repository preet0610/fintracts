# Fintracts

To start using the app follow the given instructions

1. download fabric-samples place `fintracts` directory in it and `cd` to `test-network` directory

```bash
./network.sh down
./network.sh up -cs -s couchdb
./network.sh createChannel -c bankchannel
./network.sh deployCC -ccn bank -ccp ../fintracts/chaincode-go-bank/ -ccl go -ccep "OR('Org1MSP.peer','Org2MSP.peer')" -cccg ../fintracts/chaincode-go-bank/collections_config.json -c bankchannel
./network.sh deployCC -ccn emp -ccp ../fintracts/chaincode-go-emp/ -ccl go -ccep "OR('Org1MSP.peer','Org2MSP.peer')" -cccg ../fintracts/chaincode-go-bank/collections_config.json -c bankchannel
```

this will start your hyperledger network

2. in new terminal, go to `fabric-samples/fintracts/rest-api-go` directory and type following commands

```
go mod download
go run main.go
```

this will start api server for Org1

similarly, in new terminal go to `fabric-samples/fintracts/rest-api-go-org2` directory and type following commands

```
go mod download
go run main.go
```

this will start api server for Org2

3. In new terminal go to `fabric-samples/fintracts/frontend/fintracts` directory and type

```
yarn install
yarn start
```

this will start the ReactJS app to interact with Hyperledger Network using golang APIs.
