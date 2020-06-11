# uni-resolver-driver-did-iotx

![IoTex Logo](logo/IoTeX.png)

# Universal Resolver Driver: IoTex

## Specifications

* [Decentralized Identifiers](https://w3c.github.io/did-core/)

## Example DIDs

```
did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426
```
## Example request:
curl -X GET http://127.0.0.1:8080/1.0/identifiers/did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426

## Build and Run (Docker)
please change uniresolver_driver_did_io_IoTexConnections and uniresolver_driver_did_io_IoTexContract to the real chain endpoint and contract address
```
docker build -f ./docker/Dockerfile . -t iotexproject/uni-resolver-driver-did-iotx
docker run -p 8080:8080 -e "HOST=0.0.0.0" -e "PORT=8080" -e "uniresolver_driver_did_io_IoTexConnections=api.testnet.iotex.one:443" -e "uniresolver_driver_did_io_IoTexContract=io1j2af3s4f7qjk8eudzx6a6kdhekr7zt2k5y5qlk" iotexproject/uni-resolver-driver-did-iotx

```

## Build and Run
please modify buildAndRun.sh to change uniresolver_driver_did_io_IoTexConnections and uniresolver_driver_did_io_IoTexContract to the real endpoint and contract address
```
./buildAndRun.sh
```

## Driver Environment Variables

The driver need to set the following environment variables:

HOST=0.0.0.0 default 0.0.0.0
PORT=8080 default 8080
uniresolver_driver_did_io_IoTexConnections=api.testnet.iotex.one:443
uniresolver_driver_did_io_IoTexContract=io1j2af3s4f7qjk8eudzx6a6kdhekr7zt2k5y5qlk

## License
This project is licensed under the [Apache License 2.0](LICENSE).
