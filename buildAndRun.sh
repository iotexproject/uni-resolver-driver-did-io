export HOST=0.0.0.0
export PORT=8080
export uniresolver_driver_did_io_IoTexConnections=api.testnet.iotex.one:443
export uniresolver_driver_did_io_IoTexContract=io1j2af3s4f7qjk8eudzx6a6kdhekr7zt2k5y5qlk
go build ./src/cmd/resolver-driver-server
./resolver-driver-server