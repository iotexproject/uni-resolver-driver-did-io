FROM golang:1.13.5-stretch

WORKDIR $GOPATH/github.com/iotexproject/uni-resolver-driver-did-io

RUN mkdir -p $GOPATH/github.com/iotexproject&& \
    cd $GOPATH/github.com/iotexproject&& \
    git clone https://github.com/iotexproject/uni-resolver-driver-did-io&& \
    cd uni-resolver-driver-did-io&& \
    go mod download&& \
    go build ./src/cmd/resolver-driver-server&& \
    cp resolver-driver-server /usr/local/bin

CMD [ "resolver-driver-server"]