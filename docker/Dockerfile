FROM golang:1.13.5-stretch
ENV GO111MODULE=on

WORKDIR apps/uni-resolver-driver-did-io
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /usr/local/bin/resolver-driver-server ./src/cmd/resolver-driver-server
CMD [ "resolver-driver-server"]